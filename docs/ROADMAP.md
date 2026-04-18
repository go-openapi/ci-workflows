## Roadmap

### Contemplated enhancements

#### maintenance automation

* [x] version common workflows, so we can limit the impact of a change
* [ ] dependencies: manage somehow to share / replicate dependabot config
* [ ] build: verify that go.sum cache for tests works (should be enabled)

#### tests

* [x] mono-repo test
* [x] share mono repo workflows (see github.com/go-openapi/swag/.github/workflows)
* [x] ui: enrich github actions UI with a job summary

#### releases

* [x] mono-repo release
* [ ] (possibility) take over release part from go-swagger

#### security

* [x] security: separate PR / issue comments as a trusted bot workflow, acting on request artifacts

#### code quality

* [ ] lint: manage somehow to share golangci config (with local merge)
* [ ] lint: golangci-lint: check valid PR comments etc
* [ ] lint: use non-blocking, scheduled, proactive full linting to check for
      the impact of new linters, new go versions etc

#### doc quality checks

* [x] ~doc: experiment LanguageTool for grammar checks ( -> a github action / docker image exists)~
* [x] ~doc: checkout vale style-check guide (vale-action exists)~ **vale is not up to the task**
* [x] ~introduce config file specific checkout (markdownlint, spellcheck)~
* [ ] doc: (possibility) take over hugo & doc gen part from go-swagger
* [ ] doc: produce hugo github page with all latest tagged versions
      (incl. mono repo)
* [ ] github pages w/ hugo (like go-swagger, experiment another theme and json data)
* [ ] check with github API that all repo settings (branch protection rules, etc)
      are identical
* [ ] doc: experiment LLM from github model, using embeddings
* [ ] doc: add markdown linting for docs
* [ ] doc: add spellcheck for docs (and code?)

#### github & support

* [ ] comment PRs and issues
* [ ] add bot to filter PRs, issues
* [ ] issues: experiment LLM from github model, using embeddings ( -> show related issues)

