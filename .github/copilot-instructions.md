# Copilot Instructions

This repository provides shared, reusable GitHub Actions workflows for the go-openapi organization.

When editing workflow files, follow the conventions documented in `.github/copilot/` (formatting, security, common gotchas).

## Key Rules

- Pin all action versions to commit SHAs (e.g., `uses: actions/checkout@8e8c483db84b4bee98b60c0593521ed34d9990e8 # v6.0.1`)
- Set minimal permissions at workflow level, elevate at job level only when needed
- Use `${{ }}` with spaces inside braces in all expressions and `if:` conditions
- Use `type: string` for boolean-like workflow inputs (never `type: boolean`)
- Format step arrays with `-` on its own line before `name:`

## Repository Structure

- **Shared workflows** (`.github/workflows/*.yml`): called by other go-openapi repos
- **Local test workflows** (`local-*.yml`): test the shared workflows within this repo
- New reusable actions belong in `go-openapi/gh-actions`, not here
