# ci-workflows

<!-- Badges: status  -->
[![Tests][test-badge]][test-url] [![Coverage][cov-badge]][cov-url] [![CI vuln scan][vuln-scan-badge]][vuln-scan-url] [![CodeQL][codeql-badge]][codeql-url]
<!-- Badges: release & docker images  -->
<!-- Badges: code quality  -->
<!-- Badges: license & compliance -->
[![Release][release-badge]][release-url] [![Go Report Card][gocard-badge]][gocard-url] [![CodeFactor Grade][codefactor-badge]][codefactor-url] [![License][license-badge]][license-url]
<!-- Badges: documentation & support -->
<!-- Badges: others & stats -->
[![GoDoc][godoc-badge]][godoc-url] [![Slack Channel][slack-logo]![slack-badge]][slack-url] [![go version][goversion-badge]][goversion-url] ![Top language][top-badge] ![Commits since latest release][commits-badge]

---

Common Continuous Integration (`CI`) workflows and setup for go-openapi repos.

* shared github action workflows
* shared `dependabot` configuration (**BLOCKED**)
* shared `golangci-lint` configuration (**BLOCKED**)

## Status

Unreleased.

Initial setup. Content may evolve with breaking changes.

> NOTE: at this moment, it is difficult to share the dependabot and golangci-lint configurations,
> so these are not shared yet.

## Basic usage

## Motivation

It took a while (well a something like 10 years...), but we eventually managed to align all checks, tests and
dependabot rules declared in the family of go-openapi repos.

Now we'd like to be able to maintain, enrich and improve these checks without
worrying too much about the burden of replicating the stuff about a dozen times.

## Contemplated enhancements

In no particular order:

* [x] ui: enrich github actions UI with a job summary
* [x] introduce config file specific checkout (markdownlint, spellcheck)
* [x] security: separate PR / issue comments as a trusted bot workflow, acting on request artifacts
* [x] version common workflows, so we can limit the impact of a change
* [ ] build: verify that go.sum cache for tests works (should be enabled)
* [ ] share mono repo workflows (see github.com/go-openapi/swag/.github/workflows)
* [ ] lint: manage somehow to share golangci config (with local merge)
* [ ] dependencies: manage somehow to share / replicate dependabot config
* [ ] lint: golangci-lint: check valid PR comments etc
* [ ] lint: use non-blocking, scheduled, proactive full linting to check for
      the impact of new linters, new go versions etc
* [ ] doc: (possibility) take over hugo & doc gen part from go-swagger
* [ ] (possibility) take over release part from go-swagger
* [ ] doc: produce hugo github page with all latest tagged versions
      (incl. mono repo)
* [ ] add bot to filter PRs, issues
* [ ] check with github API that all repo settings (branch protection rules, etc)
      are identical
* [ ] comment PRs and issues
* [ ] doc: checkout vale style-check guide (vale-action exists)
* [x] ~doc: experiment LanguageTool for grammar checks ( -> a github action / docker image exists)~
* [ ] doc: experiment LLM from github model, using embeddings ( -> 
* [ ] issues: experiment LLM from github model, using embeddings ( -> show related issues)
* [ ] github pages w/ hugo (like go-swagger, experiment another theme and json data)

To be reworked:
* [ ] doc: add markdown linting for docs
* [ ] doc: add spellcheck for docs (and code?)

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

## Cutting a new release

Maintainers can cut a new release by either:

* running [this workflow](https://github.com/go-openapi/gh-actions/actions/workflows/local-bump-release.yml)
* or pushing a semver tag
  * signed tags are preferred
  * The tag message is prepended to release notes

<!-- Badges: status  -->
[test-badge]: https://github.com/go-openapi/ci-workflows/actions/workflows/go-test.yml/badge.svg
[test-url]: https://github.com/go-openapi/ci-workflows/actions/workflows/go-test.yml
[cov-badge]: https://codecov.io/gh/go-openapi/ci-workflows/branch/master/graph/badge.svg
[cov-url]: https://codecov.io/gh/go-openapi/ci-workflows
[vuln-scan-badge]: https://github.com/go-openapi/ci-workflows/actions/workflows/scanner.yml/badge.svg
[vuln-scan-url]: https://github.com/go-openapi/ci-workflows/actions/workflows/scanner.yml
[codeql-badge]: https://github.com/go-openapi/ci-workflows/actions/workflows/codeql.yml/badge.svg
[codeql-url]: https://github.com/go-openapi/ci-workflows/actions/workflows/codeql.yml
<!-- Badges: release & docker images  -->
[release-badge]: https://badge.fury.io/gh/go-openapi%2Fci-workflowser.svg
[release-url]: https://badge.fury.io/gh/go-openapi%2Fci-workflowser
[gomod-badge]: https://badge.fury.io/go/github.com%2Fgo-openapi%2Fci-workflowser.svg
[gomod-url]: https://badge.fury.io/go/github.com%2Fgo-openapi%2Fci-workflowser
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
<!-- Badges: license & compliance -->
[license-badge]: http://img.shields.io/badge/license-Apache%20v2-orange.svg
[license-url]: https://github.com/go-openapi/ci-workflows/?tab=Apache-2.0-1-ov-file#readme
<!-- Badges: others & stats -->
[goversion-badge]: https://img.shields.io/github/go-mod/go-version/go-openapi/ci-workflows
[goversion-url]: https://github.com/go-openapi/ci-workflows/blob/master/go.mod
[top-badge]: https://img.shields.io/github/languages/top/go-openapi/ci-workflows
[commits-badge]: https://img.shields.io/github/commits-since/go-openapi/ci-workflows/latest
