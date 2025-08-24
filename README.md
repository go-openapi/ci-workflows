# ci-workflows

Common CI workflows and setup for go-openapi repos.

* shared github action workflows
* shared dependabot config

NOTE: at this moment, it is difficult to share the golangci-lint config, so that one is not shared yet.

## Motivation

It took a while, but we eventually managed to align all checks, tests and dependabot rules declared in the
family of go-openapi repos.

Now we'd like to be able to maintain, enrich and improve these checks without worrying too much about
the burden to replicate it about a dozen times.

## Contemplated enhancements

In no particular order:

* [x] ui: enrich github actions UI with a job summary
* [x] doc: add markdown linting for docs
* [ ] doc: add spellcheck for docs (and code?)
* [ ] version common workflows, so we can limit the impact of a change
* [ ] build: verify that go.sum cache for tests works (should be enabled)
* [ ] share mono repo workflows (see github.com/go-openapi/swag/.github/workflows)
* [ ] lint: manage somehow to share golangci config (with local merge)
* [ ] deps: manage somehow to share / replicate dependabot config
* [ ] lint: golangci-lint: check valid PR comments etc
* [ ] lint: use non-blocking, scheduled, proactive full linting to check for the impact of new linters, new go versions etc
* [ ] doc: (possibility) take over hugo & doc gen part from go-swagger
* [ ] (possibility) take over release part from go-swagger
* [ ] doc: produce hugo github page with all latest tagged versions (incl. mono repo)
* [ ] add bot to filter PRs, issues
* [ ] check with github API that all repo settings (branch protection rules, etc) are identical
* [ ] comment PRs and issues
