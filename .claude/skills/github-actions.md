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

## Code Style & Formatting

### Expression Spacing

**REQUIRED**: All GitHub Actions expressions must have spaces inside the braces:

```yaml
# ✅ CORRECT
env:
  PR_URL: ${{ github.event.pull_request.html_url }}
  TOKEN: ${{ secrets.GITHUB_TOKEN }}

# ❌ WRONG
env:
  PR_URL: ${{github.event.pull_request.html_url}}
  TOKEN: ${{secrets.GITHUB_TOKEN}}
```

### Conditional Syntax

Always use `${{ }}` in `if:` conditions:

```yaml
# ✅ CORRECT
if: ${{ inputs.enable-signing == 'true' }}
if: ${{ github.event.pull_request.user.login == 'dependabot[bot]' }}

# ❌ WRONG (works but inconsistent)
if: inputs.enable-signing == 'true'
```

### GitHub Workflow Commands

Use workflow commands for user-visible messages:

```yaml
# ✅ CORRECT - Shows as annotation in GitHub UI
echo "::notice title=build::Build completed successfully"
echo "::warning title=race-condition::Merge already in progress"
echo "::error title=deployment::Failed to deploy"

# ❌ WRONG - Just logs to console
echo "Build completed"
```

## Security Best Practices

### The secrets[inputs.name] Vulnerability

**NEVER** use dynamic secret access in workflows:

```yaml
# ❌ SECURITY VULNERABILITY
# This exposes ALL organization and repository secrets to the runner
on:
  workflow_call:
    inputs:
      secret-name:
        type: string
jobs:
  my-job:
    steps:
      - uses: some-action@v1
        with:
          token: ${{ secrets[inputs.secret-name] }}  # ❌ DANGEROUS!
```

**SOLUTION**: Use explicit secret parameters with fallback for defaults:

```yaml
# ✅ SECURE
on:
  workflow_call:
    secrets:
      gpg-private-key:
        required: false
jobs:
  my-job:
    steps:
      - uses: go-openapi/gh-actions/ci-jobs/bot-credentials@master
        with:
          # Falls back to go-openapi default if not explicitly passed
          gpg-private-key: ${{ secrets.gpg-private-key || secrets.CI_BOT_GPG_PRIVATE_KEY }}
```

### Backward Compatibility Pattern

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

### Important Notes on action.yml

**DO NOT** use `${{ }}` expressions in description fields:

```yaml
# ❌ WRONG - Can cause YAML parsing errors
description: |
  Pass it as: gpg-private-key: ${{ secrets.MY_KEY }}

# ✅ CORRECT
description: |
  Pass it as: secrets.MY_KEY
```

## Reusable Workflow Patterns

### Defining Secrets vs Inputs

```yaml
on:
  workflow_call:
    inputs:
      # Use inputs for configuration
      enable-signing:
        type: boolean
        required: false
        default: true
      bump-major:
        type: boolean
        required: false
        default: false

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

## Security Pinning

Always pin action versions to commit SHAs:

```yaml
# ✅ CORRECT - Pinned to commit SHA with version comment
uses: actions/checkout@8e8c483db84b4bee98b60c0593521ed34d9990e8 # v6.0.1
uses: crazy-max/ghaction-import-gpg@e89d40939c28e39f97cf32126055eeae86ba74ec # v6.3.0

# ❌ WRONG - Mutable tag reference
uses: actions/checkout@v6
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

## Common Gotchas

1. **Boolean input comparisons**: GitHub Actions inputs are strongly typed, with no "JS-like" truthy logic
   ```yaml
   # ❌ WRONG - Boolean true is NOT equal to string 'true'
   on:
     workflow_call:
       inputs:
         enable-feature:
           type: boolean
           default: true

   jobs:
     my-job:
       if: ${{ inputs.enable-feature == 'true' }}  # FALSE when input is boolean true!

   # ✅ CORRECT - Handle both boolean and string values
   if: ${{ inputs.enable-feature == 'true' || inputs.enable-feature == true }}

   # Note: In bash, this works fine because bash converts to string:
   if [[ '${{ inputs.enable-feature }}' == 'true' ]]; then  # Works in bash
   ```

2. **Expression evaluation in descriptions**: Don't use `${{ }}` in action.yml description fields
3. **Race conditions**: Always use optimistic execution + error handling, never check-then-act
4. **Secret exposure**: Never use `secrets[inputs.name]` - always use explicit secret parameters
5. **Branch deletion**: Use `wait-pending-jobs` before merging to prevent failures in non-required jobs
6. **Idempotency**: `gh pr merge --auto` is NOT idempotent - handle "Merge already in progress" error
7. **TOCTOU vulnerabilities**: State can change between check and action - handle at runtime

## Testing Workflows

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
