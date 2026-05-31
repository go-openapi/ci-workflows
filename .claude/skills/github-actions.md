# GitHub Actions Skills for go-openapi

This skill captures best practices, patterns, and learnings for working with GitHub Actions workflows in the go-openapi organization.

## Repository Architecture

### Two-Repository Pattern

**go-openapi/ci-workflows** (.github/workflows/):
- Contains **reusable workflows** called by other repos
- Examples: `go-test.yml`, `bump-release.yml`, `contributors.yml`, `auto-merge.yml`
- These are called via `uses: go-openapi/ci-workflows/.github/workflows/workflow-name.yml@master`

**go-openapi/gh-actions** (ci-jobs/ and install/):
- Contains **composite actions** for reusable steps
- Examples: `bot-credentials`, `wait-pending-jobs`, tool installers
- These are called via `uses: go-openapi/gh-actions/ci-jobs/action-name@master`

### When to Create New Actions

**IMPORTANT**: When proposing new reusable GitHub Actions functionality:
- **Always create composite actions in `go-openapi/gh-actions`**, NOT in ci-workflows
- Structure: `gh-actions/ci-jobs/{action-name}/action.yml`
- Document in the main `gh-actions/README.md` (NOT a separate README in the action folder)
- Follow the pattern established by `bot-credentials` and `wait-pending-jobs`

## Backward Compatibility Pattern

Use the `||` operator to provide fallback to default secret names:

```yaml
# Works for go-openapi repos with secrets: inherit
# Also works for other orgs that pass explicit secrets
secrets.custom-name || secrets.GO_OPENAPI_DEFAULT_NAME
```

## Race Condition Handling

### The Problem: Check-Then-Act (TOCTOU)

**NEVER** use check-then-act patterns for concurrent operations:

```yaml
# ❌ RACE CONDITION (Time-of-check to time-of-use vulnerability)
- name: Check if merged
  id: check
  run: |
    if gh pr view "$PR" --json state | grep -q MERGED; then
      echo "already-merged=true" >> "$GITHUB_OUTPUT"
    fi

- name: Merge PR
  if: steps.check.outputs.already-merged != 'true'
  run: gh pr merge --auto "$PR"  # Can fail if merged between check and act
```

### The Solution: Optimistic Execution with Error Handling

**ALWAYS** attempt the operation and handle expected errors gracefully:

```yaml
# ✅ CORRECT - Handles race condition properly
- name: Enable auto-merge
  run: |
    set +e
    OUTPUT=$(gh pr merge --auto --rebase "$PR_URL" 2>&1)
    EXIT_CODE=$?
    set -e

    if [ $EXIT_CODE -eq 0 ]; then
      echo "::notice title=auto-merge::Auto-merge enabled successfully"
      exit 0
    fi

    # Check for expected race condition error
    # GitHub GraphQL API returns: "GraphQL: Merge already in progress (mergePullRequest)"
    if echo "$OUTPUT" | grep -q "Merge already in progress"; then
      echo "::warning title=auto-merge::Auto-merge already handled by another workflow"
      exit 0
    fi

    # Unexpected error - fail properly
    echo "::error title=auto-merge::Failed to enable auto-merge"
    echo "$OUTPUT"
    exit $EXIT_CODE
```

### Known GitHub GraphQL Error Messages

When handling race conditions, check for these specific error messages:

- **Merge already in progress**: `"Merge already in progress"` (from `gh pr merge --auto`)
- Use exact strings from GitHub's GraphQL API responses (check action logs to verify)

## Common Patterns & Solutions

### wait-pending-jobs Action

**Purpose**: Wait for ALL workflow runs (including non-required jobs) to complete before merging.

**Why it's needed**: Without this, auto-merge can delete the branch while non-required jobs (like coverage upload) are still running, causing them to fail.

```yaml
- name: Wait for all workflow runs to complete
  uses: go-openapi/gh-actions/ci-jobs/wait-pending-jobs@master
  with:
    pr-url: ${{ github.event.pull_request.html_url }}
    github-token: ${{ secrets.GITHUB_TOKEN }}
    # Optional: exclude patterns to prevent deadlocks
    exclude-workflow-patterns: 'auto-merge,contributors'
```

### bot-credentials Action

**Purpose**: Securely configure GPG signing and GitHub App authentication.

**Solves**: The `secrets[inputs.name]` security vulnerability.

```yaml
# For GPG signing only
- uses: go-openapi/gh-actions/ci-jobs/bot-credentials@master
  with:
    enable-gpg-signing: 'true'
    gpg-private-key: ${{ secrets.gpg-private-key || secrets.CI_BOT_GPG_PRIVATE_KEY }}
    gpg-passphrase: ${{ secrets.gpg-passphrase || secrets.CI_BOT_GPG_PASSPHRASE }}
    gpg-fingerprint: ${{ secrets.gpg-fingerprint || secrets.CI_BOT_SIGNING_KEY }}

# For GitHub App + GPG
- uses: go-openapi/gh-actions/ci-jobs/bot-credentials@master
  id: bot
  with:
    enable-github-app: 'true'
    github-app-id: ${{ secrets.github-app-id || secrets.CI_BOT_APP_ID }}
    github-app-private-key: ${{ secrets.github-app-private-key || secrets.CI_BOT_APP_PRIVATE_KEY }}
    enable-gpg-signing: 'true'
    gpg-private-key: ${{ secrets.gpg-private-key || secrets.CI_BOT_GPG_PRIVATE_KEY }}
    gpg-passphrase: ${{ secrets.gpg-passphrase || secrets.CI_BOT_GPG_PASSPHRASE }}
    gpg-fingerprint: ${{ secrets.gpg-fingerprint || secrets.CI_BOT_SIGNING_KEY }}

- name: Use bot token
  run: gh pr create --token "${{ steps.bot.outputs.app-token }}"
```

### Auto-Merge Pattern for Bot PRs

**Complete flow** for bot-created PRs:

```yaml
jobs:
  create-pr:
    steps:
      - name: Create PR
        id: create-pr
        uses: peter-evans/create-pull-request@v8
        with:
          token: ${{ steps.bot.outputs.app-token }}

  auto-merge:
    needs: [create-pr]
    steps:
      - name: Auto-approve PR
        run: gh pr review --approve "$PR_URL"

      - name: Wait for all workflow runs to complete
        uses: go-openapi/gh-actions/ci-jobs/wait-pending-jobs@master
        with:
          pr-url: ${{ env.PR_URL }}
          github-token: ${{ secrets.GITHUB_TOKEN }}

      - name: Enable auto-merge (with race condition handling)
        run: |
          set +e
          OUTPUT=$(gh pr merge --auto --rebase "$PR_URL" 2>&1)
          EXIT_CODE=$?
          set -e

          if [ $EXIT_CODE -eq 0 ]; then
            echo "::notice title=auto-merge::Auto-merge enabled successfully"
            exit 0
          fi

          if echo "$OUTPUT" | grep -q "Merge already in progress"; then
            echo "::warning title=auto-merge::Auto-merge already handled by another workflow"
            exit 0
          fi

          echo "::error title=auto-merge::Failed to enable auto-merge"
          echo "$OUTPUT"
          exit $EXIT_CODE
```

## Action Definition Best Practices

### Composite Action Structure

```yaml
# SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
# SPDX-License-Identifier: Apache-2.0

name: action-name
description: |
  Clear description of what this action does.

  Multi-line descriptions are fine.

author: go-openapi

inputs:
  my-input:
    description: |
      Description of the input.

      Pass it as: secrets.MY_SECRET_NAME

      Required when some-condition is true.
    required: false
    default: 'false'

outputs:
  my-output:
    description: Description of the output
    value: ${{ steps.step-id.outputs.value }}

runs:
  using: composite
  steps:
    - name: Do something
      shell: bash
      run: |
        echo "Action logic here"

branding:
  icon: 'key'
  color: 'green'
```

## Reusable Workflow Patterns

### Defining Secrets vs Inputs

```yaml
on:
  workflow_call:
    inputs:
      # Use inputs for configuration
      # IMPORTANT: Use type: string for boolean-like values (never type: boolean)
      enable-signing:
        type: string          # ✅ Use string, not boolean
        required: false
        default: 'true'       # String value
      bump-major:
        type: string          # ✅ Use string, not boolean
        required: false
        default: 'false'      # String value

    secrets:
      # Use secrets for sensitive data
      gpg-private-key:
        description: |
          GPG private key in armored format.

          Default for go-openapi: CI_BOT_GPG_PRIVATE_KEY

          Required when enable-signing is true.
        required: false
```

### Calling Reusable Workflows

**From go-openapi repos** (with secrets: inherit):

```yaml
jobs:
  release:
    uses: go-openapi/ci-workflows/.github/workflows/bump-release.yml@master
    secrets: inherit  # Inherits all secrets, workflow uses fallback pattern
```

**From other organizations** (with explicit secrets):

```yaml
jobs:
  release:
    uses: go-openapi/ci-workflows/.github/workflows/bump-release.yml@master
    secrets:
      gpg-private-key: ${{ secrets.MY_ORG_GPG_KEY }}
      gpg-passphrase: ${{ secrets.MY_ORG_GPG_PASS }}
      gpg-fingerprint: ${{ secrets.MY_ORG_GPG_FP }}
```

## Documentation Standards

### README Structure for Actions

Document new actions in the main `gh-actions/README.md`:

```markdown
### action-name

Brief description of what the action does.

**Features:**
- Feature 1
- Feature 2

**Usage example 1: go-openapi repos (using default secret names)**

[Show example with default go-openapi secret names]

**Usage example 2: Other organizations (using custom secret names)**

[Show example with custom secret names for other orgs]

**Background:** Explain why this action exists and what problem it solves.
```

### Local Testing Pattern

ci-workflows has `local-*` workflows that test the shared workflows:

```yaml
# local-bump-release.yml
on:
  workflow_dispatch:
    inputs:
      # Mirror the inputs from the reusable workflow

jobs:
  test:
    uses: ./.github/workflows/bump-release.yml  # Call local version for testing
    with:
      bump-patch: ${{ inputs.bump-patch }}
    secrets: inherit
```

This allows testing changes before they're consumed by other repos.
