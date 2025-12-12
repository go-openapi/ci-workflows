## Roadmap

### Contemplated enhancements

Most urgent:

* [ ] mono-repo test
* [ ] mono-repo release

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
* [ ] check with github API that all repo settings (branch protection rules, etc)
      are identical
* [ ] comment PRs and issues
* [ ] doc: checkout vale style-check guide (vale-action exists)
* [x] ~doc: experiment LanguageTool for grammar checks ( -> a github action / docker image exists)~
* [ ] github pages w/ hugo (like go-swagger, experiment another theme and json data)

AI-driven experiments:
* [ ] add bot to filter PRs, issues
* [ ] doc: experiment LLM from github model, using embeddings ( -> 
* [ ] issues: experiment LLM from github model, using embeddings ( -> show related issues)

To be reworked:
* [ ] doc: add markdown linting for docs
* [ ] doc: add spellcheck for docs (and code?)
