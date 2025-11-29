# CI workflows

## Shared workflows

### Dependencies automation

* auto-merge.yml:
  * auto-merge dependabot updates,  with dependency group rules
  * auto-merge go-openapi bot updates

### Test automation

* go-test.yml: go unit tests
* monorepo-go-test.yml: go unit tests for monorepos

* collect-coverage.yml: (common) collect & publish test coverage (to codecov)
* collect-reports.yml: (common) collect & publish test reports (to codecov and github)

### Security 

* codeql.yml: CodeQL workflow for go and github actions
* scanner.yml: trivy & govulncheck scans

### Release automation

* bump-release.yml: manually triggered workflow to cut a release
* tag-release.yml: cut a release on push tag
* release.yml: (common) release & release notes build

### Code quality

* collect-coverage.yml: common collect & publish test coverage (to codecov)
* collect-reports.yml: common collect & publish test reports (to codecov and github)

### Documentation quality

* contributors.yml: updates CONTRIBUTORS.md
* doc-update.yml: lint & spellcheck on markdown updates
* pr-comment.yml: common PR commment workflow

## Test workflows

* local-auto-merge.yml
* local-bump-release.yml
* local-codeql.yml
* local-contributors.yml
* local-doc-update.yml
* local-go-test.yml
* local-monorepo-go-test.yml
* local-release.yml
* local-scanner.yml
* local-tag-release.yml

## Configuration files

* .cliff

scripts
