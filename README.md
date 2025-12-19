# ci-workflows

<!-- Badges: status  -->
[![Tests][test-badge]][test-url] [![Coverage][cov-badge]][cov-url] [![CI vuln scan][vuln-scan-badge]][vuln-scan-url] [![CodeQL][codeql-badge]][codeql-url]
<!-- Badges: release & docker images  -->
<!-- Badges: code quality  -->
<!-- Badges: license & compliance -->
[![Release][release-badge]][release-url] [![Go Report Card][gocard-badge]][gocard-url] [![CodeFactor Grade][codefactor-badge]][codefactor-url] [![License][license-badge]][license-url]
<!-- Badges: documentation & support -->
<!-- Badges: others & stats -->
[![GoDoc][godoc-badge]][godoc-url] [![Discord Channel][discord-badge]][discord-url] [![go version][goversion-badge]][goversion-url] ![Top language][top-badge] ![Commits since latest release][commits-badge]

---

Common Continuous Integration (`CI`) workflows and setup for go-openapi repos.

* shared github action workflows
* shared `dependabot` configuration (**BLOCKED**)
* shared `golangci-lint` configuration (**BLOCKED**)

## Announcements

* **2025-12-19** : new community chat on discord
  * a new discord community channel is available to be notified of changes and support users
  * our venerable Slack channel remains open, and will be eventually discontinued on **2026-03-31**

You may join the discord community by clicking the invite link on the discord badge (also above). [![Discord Channel][discord-badge]][discord-url]

Or join our Slack channel: [![Slack Channel][slack-logo]![slack-badge]][slack-url]

## Status

Development is active. We are regularly adding more shared workflows to standardize CI across go-openapi repos.

> NOTE: at this moment, it is difficult to share the configurations for dependabot and golangci-lint,
> so these are not shared yet.

## Example

`go-test.yml`

![go-test workflow](./docs/images/go-test.png)

## Basic usage

You reuse a workflow like so:

```yaml
name: go test

permissions:
  pull-requests: read
  contents: read

on:
  push:
    branches:
      - master

  pull_request:

jobs:
  test:
    uses: go-openapi/ci-workflows/.github/workflows/go-test.yml@master
    secrets: inherit
```

It is recommended to pin the git ref `master` with a commit sha, and let dependabot keep you up to date. Like so:

```yaml
    uses: go-openapi/ci-workflows/.github/workflows/go-test.yml@b28a8b978a5ee5b7f4241ffafd6cc6163edb5dfd # v0.1.0
```

### Permissions

Make sure your job permissions match the requirements of the called shared workflow.

Example:
```yaml
name: "CodeQL"

on:
  push:
    branches: [ "master" ]
  pull_request:
    branches: [ "master" ]
    paths-ignore: # remove this clause if CodeQL is a required check
      - '**/*.md'
  schedule:
    - cron: '39 19 * * 5'

permissions:
  contents: read

jobs:
  codeql:
    permissions:  # <- grant permissions at the job level that match the requirements of the called workflow
      contents: read
      security-events: write
    uses: ./.github/workflows/codeql.yml
    secrets: inherit
```

## Available workflows `[v0.1.0]`

### Dependencies automation

* auto-merge.yml:
  * auto-merge dependabot updates,  with dependency group rules
  * auto-merge go-openapi bot updates

### Test automation

* go-test.yml: go unit tests
  * includes:
    * fuzz-test.yml: orchestrates fuzz testing with a cached corpus
    * collect-coverage.yml: (common) collect & publish test coverage (to codecov)
    * collect-reports.yml: (common) collect & publish test reports (to codecov and github)

* go-test-monorepo.yml: go unit tests, with support for go mono-repos (same features)

>NOTE: for mono-repos, the workflow works best with go1.25 and go.work declaring all your modules and committed to git.

### Security 

* codeql.yml: CodeQL workflow for go and github actions
* scanner.yml: trivy & govulncheck scans

### Release automation

* bump-release.yml: manually triggered workflow to cut a release
* tag-release.yml: cut a release on push tag
* release.yml: (common) release & release notes build

>NOTE: mono-repos are not supported yet.

Release notes are produced using `git-cliff`. The configuration may be set using a `.cliff.toml` file.
The default configuration is the `.cliff.toml` in this repo (uses remote config).

### Documentation quality

* contributors.yml: updates CONTRIBUTORS.md

## Motivation

It took a while (well something like 10 years...), but we eventually managed to align all checks, tests and
dependabot rules declared in the family of go-openapi repos.

Now we'd like to be able to maintain, enrich and improve these checks without
worrying too much about the burden of replicating the stuff about a dozen times.

## Change log

See <https://github.com/go-openapi/ci-workflows/releases>

## Licensing

This content ships under the [SPDX-License-Identifier: Apache-2.0](./LICENSE).

<!--
## Limitations
-->

## Other documentation

* [All-time contributors](./CONTRIBUTORS.md)
* [Contributing guidelines](.github/CONTRIBUTING.md)
* [Maintainers documentation](docs/MAINTAINERS.md)
* [Code style](docs/STYLE.md)
* [Roadmap](docs/ROADMAP.md)

## Cutting a new release

Maintainers can cut a new release by either:

* running [this workflow](https://github.com/go-openapi/ci-workflows/actions/workflows/local-bump-release.yml)
* or pushing a semver tag
  * signed tags are preferred
  * The tag message is prepended to release notes

<!-- Badges: status  -->
[test-badge]: https://github.com/go-openapi/ci-workflows/actions/workflows/local-go-test.yml/badge.svg
[test-url]: https://github.com/go-openapi/ci-workflows/actions/workflows/local-go-test.yml
[cov-badge]: https://codecov.io/gh/go-openapi/ci-workflows/branch/master/graph/badge.svg
[cov-url]: https://codecov.io/gh/go-openapi/ci-workflows
[vuln-scan-badge]: https://github.com/go-openapi/ci-workflows/actions/workflows/local-scanner.yml/badge.svg
[vuln-scan-url]: https://github.com/go-openapi/ci-workflows/actions/workflows/local-scanner.yml
[codeql-badge]: https://github.com/go-openapi/ci-workflows/actions/workflows/local-codeql.yml/badge.svg
[codeql-url]: https://github.com/go-openapi/ci-workflows/actions/workflows/local-codeql.yml
<!-- Badges: release & docker images  -->
[release-badge]: https://badge.fury.io/gh/go-openapi%2Fci-workflows.svg
[release-url]: https://badge.fury.io/gh/go-openapi%2Fci-workflows
[gomod-badge]: https://badge.fury.io/go/github.com%2Fgo-openapi%2Fci-workflows.svg
[gomod-url]: https://badge.fury.io/go/github.com%2Fgo-openapi%2Fci-workflows
<!-- Badges: code quality  -->
[gocard-badge]: https://goreportcard.com/badge/github.com/go-openapi/ci-workflows
[gocard-url]: https://goreportcard.com/report/github.com/go-openapi/ci-workflows
[codefactor-badge]: https://img.shields.io/codefactor/grade/github/go-openapi/ci-workflows
[codefactor-url]: https://www.codefactor.io/repository/github/go-openapi/ci-workflows
<!-- Badges: documentation & support -->
[doc-badge]: https://img.shields.io/badge/doc-site-blue?link=https%3A%2F%2Fgoswagger.io%2Fgo-openapi%2F
[doc-url]: https://goswagger.io/go-openapi
[godoc-badge]: https://pkg.go.dev/badge/github.com/go-openapi/ci-workflows
[godoc-url]: http://pkg.go.dev/github.com/go-openapi/ci-workflows
[slack-logo]: https://a.slack-edge.com/e6a93c1/img/icons/favicon-32.png
[slack-badge]: https://img.shields.io/badge/slack-blue?link=https%3A%2F%2Fgoswagger.slack.com%2Farchives%2FC04R30YM
[slack-url]: https://goswagger.slack.com/archives/C04R30YMU
[discord-badge]: https://img.shields.io/discord/1446918742398341256?logo=discord&label=discord&color=blue
[discord-url]: https://discord.gg/DrafRmZx
<!-- Badges: license & compliance -->
[license-badge]: http://img.shields.io/badge/license-Apache%20v2-orange.svg
[license-url]: https://github.com/go-openapi/ci-workflows/?tab=Apache-2.0-1-ov-file#readme
<!-- Badges: others & stats -->
[goversion-badge]: https://img.shields.io/github/go-mod/go-version/go-openapi/ci-workflows
[goversion-url]: https://github.com/go-openapi/ci-workflows/blob/master/go.mod
[top-badge]: https://img.shields.io/github/languages/top/go-openapi/ci-workflows
[commits-badge]: https://img.shields.io/github/commits-since/go-openapi/ci-workflows/latest
