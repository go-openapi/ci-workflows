# Dependabot + GitHub Actions automation

Patterns, gotchas, and known limitations for automating Dependabot
via GitHub Actions in the go-openapi organization.

Use this skill alongside `github-actions.md` (general workflow patterns)
and `.claude/rules/github-workflows-conventions.md` (formatting rules).

## When to use this skill

Invoke when designing or debugging a workflow that:

- Posts `@dependabot` comment commands (rebase, recreate, merge, ignore, etc.)
- Auto-merges Dependabot PRs
- Reads Dependabot PR metadata (update type, ecosystem, dependency name)
- Rebases stalled Dependabot branches
- Reacts to `dependabot[bot]` push or PR events

Do NOT use this skill for:
- Configuring `.github/dependabot.yml` (the dependency-update schedule itself)
  — that is Dependabot's input config, not GHA automation.
- General `update-branch` / auto-merge patterns that aren't Dependabot-specific.

## The identity rule (most important gotcha)

Dependabot inspects every comment it receives via the `issue_comment`
webhook and decides whether to honor commands based on
**`comment.user.type`**, NOT on the username.

| `user.type` | Examples | Dependabot decision |
|---|---|---|
| `"Bot"` | `github-actions[bot]`, `bot-go-openapi[bot]`, any GitHub App installation token | **Rejected — always.** "Sorry, only users with push access can use that command." |
| `"User"` | A real user account (human or dedicated bot account like `bot-go-openapi`), commenting via PAT or OAuth token | Accepted if the account has push/maintain/admin on the repo. |

This was changed by GitHub in early 2023 and is still tracked as open
in [dependabot-core#9147](https://github.com/dependabot/dependabot-core/issues/9147)
("Support Dependabot commands by GitHub apps"). See also
[community#48010](https://github.com/orgs/community/discussions/48010).
There is no per-App allowlist. Granting an App `contents: write` +
`pull-requests: write` does NOT bypass the check.

### What this means in practice

| You want to … | Use this identity | Source |
|---|---|---|
| Post `@dependabot rebase` (or any `@dependabot ...` command) | **A real user account's PAT.** In go-openapi: `bot-go-openapi` user's PAT, stored as `CI_BOT_PAT`. | secret |
| Rebase / update a branch directly (not via Dependabot command) | `GITHUB_TOKEN` with `contents: write` + `pull-requests: write`. Calls `updatePullRequestBranch` GraphQL mutation; the `[bot]` identity check does NOT apply here. | built-in |
| Approve / merge a Dependabot PR | `GITHUB_TOKEN` (cannot self-approve, but can merge PRs by others) or the bot user's PAT (can approve as user). | built-in or PAT |
| Create a PR (e.g. for follow-up changes) | GitHub App installation token (so commits/PR are attributed to the bot identity). In go-openapi: `bot-credentials` action via `CI_BOT_APP_*`. | App |

### Why we can't use the GitHub App for Dependabot commands

`bot-go-openapi` exists as **two separate entities**:
- A **GitHub App** (`bot-go-openapi[bot]`, `user.type == "Bot"`) — used by
  `bot-credentials` for App installation tokens. Rejected by Dependabot.
- A **real user account** (`bot-go-openapi`, `user.type == "User"`) —
  separate login, must be a team member with write access. Accepted by
  Dependabot when commenting via its PAT.

These two identities share a name root but are not interchangeable.
**For Dependabot commands, only the user account works.**

## Patterns

### Posting `@dependabot rebase` on stalled PRs

The shared workflow is `.github/workflows/monitor-bot-pr.yml` (the
`dependabot-prs` job). Key shape:

```yaml
secrets:
  bot-pat:
    description: |
      PAT of a real user with push access. go-openapi default: CI_BOT_PAT.
    required: false

jobs:
  dependabot-prs:
    permissions:
      contents: write
      pull-requests: write
      statuses: read
    env:
      GH_TOKEN: ${{ secrets.bot-pat || secrets.CI_BOT_PAT }}
    steps:
      -
        name: Request a rebase on stalled dependabot PRs
        run: |
          # detect stalled PRs (BEHIND, auto-merge enabled, no failed checks)
          # ...
          gh pr comment "$pr_url" --body "@dependabot rebase"
```

The bot user PAT scopes needed (minimal):
- **Classic:** `public_repo` only (or `repo` for private repos).
- **Fine-grained:** `Pull requests: Read and write` + auto-included
  `Metadata: Read`. Resource owner = the org (go-openapi).

### Rebasing a non-Dependabot bot's PRs

For PRs opened by an organization bot (e.g. `bot-go-openapi[bot]`
opening a PR via the App), use `gh pr update-branch --rebase` directly.
This calls the `updatePullRequestBranch` GraphQL mutation and works
with `GITHUB_TOKEN` (no PAT needed) because Dependabot's identity check
does NOT apply — we're not asking Dependabot to do anything.

See `monitor-bot-pr.yml` `organization-bot-prs` job for the implementation.

### Reading Dependabot PR metadata

Use `dependabot/fetch-metadata` (a Dependabot-published action). It
runs as part of a workflow triggered by `pull_request` events on
Dependabot PRs and exposes outputs like `update-type`, `dependency-names`,
`package-ecosystem`. Useful for conditional auto-merge ("only auto-merge
patch and minor updates"). No identity concerns — it just reads
metadata, doesn't comment.

### Auto-merging Dependabot PRs

Two approaches, both work without the identity gotcha:
- `gh pr merge --auto --rebase "$pr_url"` with `GITHUB_TOKEN` (needs
  `contents: write` + `pull-requests: write`). Triggered from a workflow
  reacting to the Dependabot PR.
- An explicit approval step using the bot user's PAT before merge, if
  required-reviews are enforced (since `GITHUB_TOKEN` cannot approve PRs
  it didn't open — and Dependabot opened this one).

## Anti-patterns / things tried that don't work

**Granting more permissions to `GITHUB_TOKEN`.** The check is on
`user.type`, not on token scopes. Even `contents: admin` + `pull-requests:
admin` wouldn't change the answer for `github-actions[bot]`.

**Swapping in a GitHub App installation token (via `bot-credentials`
with `enable-github-app: 'true'`).** The comment is then attributed to
`<app-name>[bot]`, `user.type == "Bot"` — still rejected. This was tried
on strfmt PR #256; commit `7370d55` introduced it; commit `980b99d`
replaced it with the PAT approach when it failed.

**Using `gh pr update-branch --rebase` instead of `@dependabot rebase`.**
This force-pushes the rebased branch externally. Looks tempting (no PAT
needed) but:
- Dependabot treats externally-pushed branches as "user-edited" and
  stops managing the PR.
- Likely triggers a `go.sum` merge conflict (or whatever lockfile the
  ecosystem uses), since both base and head modified it. Dependabot's
  own rebase regenerates the lockfile cleanly — a raw git rebase
  doesn't.
- Loses the Dependabot pipeline orchestration (re-resolve, re-tidy,
  close if obsolete). End result: more babysitting, not less.

**Adding `bot-go-openapi` as a collaborator via the GitHub App.** The
App can be granted repo write access, but that authorizes the App
identity (`user.type == "Bot"`), not the user account. The user account
is a separate entity that needs separate team/collab membership.

## Anatomy of the `bot-go-openapi` identities

| Identity | `user.login` | `user.type` | How to authenticate | Used for |
|---|---|---|---|---|
| GitHub App | `bot-go-openapi[bot]` | `Bot` | App installation token via `bot-credentials` action (`CI_BOT_APP_ID` + `CI_BOT_APP_PRIVATE_KEY`) | Opening PRs, pushing commits, creating tags. Comments will be rejected by Dependabot. |
| User account | `bot-go-openapi` | `User` | PAT (`CI_BOT_PAT`) | Posting Dependabot commands. Approving PRs the App opened. |

When choosing the identity for a new automation step, look at the
target operation's identity check, not at what feels natural.

## External references

- [dependabot-core#9147 — Support Dependabot commands by GitHub apps](https://github.com/dependabot/dependabot-core/issues/9147)
- [community#48010 — Sudden change in access required for dependabot commands](https://github.com/orgs/community/discussions/48010)
- [GitHub Docs: Managing pull requests for dependency updates](https://docs.github.com/en/code-security/dependabot/working-with-dependabot/managing-pull-requests-for-dependency-updates)
- [GitHub Docs: Automating Dependabot with GitHub Actions](https://docs.github.com/en/code-security/dependabot/working-with-dependabot/automating-dependabot-with-github-actions)
- `bot-credentials` action: `go-openapi/gh-actions/ci-jobs/bot-credentials`
- Reference workflow: `.github/workflows/monitor-bot-pr.yml` in this repo
