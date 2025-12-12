# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

This repository provides shared, reusable GitHub Actions workflows for the go-openapi organization. The workflows are designed to be called from other go-openapi repositories to standardize CI/CD processes across the entire project family.

## GitHub Actions Skills
       
**IMPORTANT**: When working with GitHub Actions workflows in this repository, refer to the comprehensive GitHub Actions skill:
         
📖 **See `.claude/skills/github-actions.md`** for:
- Code style and formatting requirements (expression spacing, workflow commands)
- Security best practices (avoiding `secrets[inputs.name]` vulnerability)
- Race condition handling patterns (optimistic execution with error handling)
- Common workflow patterns (bot-credentials, wait-pending-jobs, auto-merge)
- Action definition best practices
- Documentation standards

When proposing new reusable actions, always create them in `go-openapi/gh-actions` (not in this repo).

## Testing & Development Commands

### Running Tests
```bash
# Run all unit tests with coverage
gotestsum --jsonfile 'unit.report.json' -- -race -p 2 -count 1 -timeout=20m -coverprofile='unit.coverage.out' -covermode=atomic -coverpkg="$(go list)"/... ./...

# Run a single test
go test -v -run TestName ./path/to/package

# Run tests locally (as CI would)
# Uses the local-* workflows which mirror how other repos consume these workflows
```

### Linting
```bash
# Run golangci-lint (uses .golangci.yml configuration)
golangci-lint run

# The linter configuration is highly customized with many linters disabled
# to match go-openapi's established code style
```

### Fuzz Testing
```bash
# List available fuzz tests across all packages
go test -list Fuzz ./...

# Run a specific fuzz test for 1m30s
go test -fuzz=FuzzTestName -fuzztime=1m30s ./path/to/package

# Fuzz corpus is cached in $(go env GOCACHE)/fuzz
```

## Architecture & Design

### Workflow Types

This repository contains two types of workflows:

1. **Shared Workflows** (called by other repos):
   - `go-test.yml` - Complete test pipeline with linting, unit tests, fuzz tests, coverage
   - `auto-merge.yml` - Auto-approves and merges dependabot and bot PRs
   - `bump-release.yml` - Manually triggered release workflow (creates signed tags)
   - `tag-release.yml` - Release workflow triggered by pushing tags
   - `release.yml` - Common release building workflow (called by other release workflows)
   - `codeql.yml` - CodeQL security scanning for Go and GitHub Actions
   - `scanner.yml` - Trivy and govulncheck vulnerability scanning
   - `contributors.yml` - Automatically updates CONTRIBUTORS.md
   - `collect-coverage.yml` - Collects and publishes coverage to codecov
   - `collect-reports.yml` - Collects and publishes test reports
   - `fuzz-test.yml` - Orchestrates fuzz testing with cached corpus

2. **Local Test Workflows** (prefixed with `local-*`):
   - These workflows test the shared workflows within this repository
   - They mimic how consuming repos would invoke the shared workflows
   - Example: `local-go-test.yml` calls `./.github/workflows/go-test.yml`

### How Shared Workflows Are Used

Other go-openapi repos consume these workflows like:

```yaml
jobs:
  test:
    uses: go-openapi/ci-workflows/.github/workflows/go-test.yml@master
    secrets: inherit
```

Recommended practice: pin to a commit SHA and let dependabot update it:
```yaml
uses: go-openapi/ci-workflows/.github/workflows/go-test.yml@b28a8b978a5ee5b7f4241ffafd6cc6163edb5dfd # v0.1.0
```

### Fuzz Test Architecture

Fuzz testing has a unique multi-stage design due to Go's limitation that `go test -fuzz` cannot run across multiple packages:

1. **fuzz-matrix job**: Discovers all fuzz tests using `go test -list Fuzz -json` and creates a matrix
2. **fuzz-test job**: Runs each discovered fuzz test in parallel with:
   - Cached corpus stored in GitHub Actions cache (max 250MB)
   - Automatic cache purging to maintain size limits
   - Failed corpus uploaded as artifacts for 60 days
   - Default fuzz time: 1m30s, minimize time: 5m

### Release Process

Releases can be triggered in two ways:

1. **Manual bump** via `bump-release.yml`:
   - Select patch/minor/major bump
   - Creates GPG-signed tag using bot credentials
   - Triggers release build automatically

2. **Direct tag push** via `tag-release.yml`:
   - Push a semver tag (signed tags preferred)
   - Tag message is prepended to release notes
   - Triggers release build

Release notes are generated using [git-cliff](https://git-cliff.org/) with configuration in `.cliff.toml`.

### Auto-merge Logic

The `auto-merge.yml` workflow handles two bot types:
- **dependabot[bot]**: Auto-approves all PRs, auto-merges the following groups:
    - `development-dependencies`
    - `go-openapi-dependencies` (for minor and patch updates only)
    - `golang-org-dependencies`
- **bot-go-openapi[bot]**: Auto-approves and auto-merges all PRs (for contributor updates, etc.)

## Key Configuration Files

- `.golangci.yml` - Linter configuration (many linters disabled to match go-openapi style)
- `.cliff.toml` - Release notes generation configuration
- `.github/dependabot.yaml` - Dependency update configuration
- `go.mod` - Requires Go 1.24.0

## Important Notes

- All workflow action versions are pinned to commit SHAs for security
- Permissions are explicitly granted at job level to follow least-privilege principle
- This repo itself uses minimal Go code (just sample tests); it's primarily YAML workflows
- The `local-*` workflows serve as both tests and documentation of proper usage
