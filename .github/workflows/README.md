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
* bump-release-monorepo.yml: manually triggered workflow to cut a release for a mono-repo
* prepare-release-monorepo.yml: (common) update the inter-module dependencies of a mono-repo before tagging
* tag-release.yml: cut a release on push tag
* release.yml: (common) release & release notes build

#### Releasing assets with goreleaser

By default a release only carries its release notes. Repositories which ship binaries or packages may opt in to
[goreleaser](https://goreleaser.com) with `enable-goreleaser: 'true'`, on any of the release workflows above:

```yaml
jobs:
  gh-release:
    name: Create release
    permissions:
      contents: write
    uses: go-openapi/ci-workflows/.github/workflows/tag-release.yml@master
    with:
      tag: ${{ github.ref_name }}
      enable-goreleaser: 'true'
    secrets: inherit
```

The tagging and the release notes don't change: the workflow still resolves the tag and still builds the notes with
git-cliff, then hands the notes file over to goreleaser, which builds the assets and publishes the github release.

There is no shared goreleaser config: each repository declares its own binaries, packages, signatures and announcements
in a local config file (`.goreleaser.yml` by default, see the `goreleaser-config` input).

| input | default | purpose |
|---|---|---|
| `enable-goreleaser` | `'false'` | opt into the goreleaser stage |
| `goreleaser-config` | `'.goreleaser.yml'` | path to the local goreleaser config |
| `goreleaser-version` | `'~> v2'` | version constraint of the goreleaser distribution |
| `goreleaser-args` | `'release --clean --fail-fast'` | base arguments (`--config` and `--release-notes` are appended) |
| `goreleaser-env` | `''` | extra **non-secret** environment, as newline-separated `KEY=VALUE` pairs |
| `enable-artifact-signing` | `'true'` | import the bot PGP key, for configs with a `signs` section |
| `enable-upx` | `'false'` | install UPX, for configs with an `upx` section |

Secrets, all optional and all defaulting to their go-openapi counterpart when the caller uses `secrets: inherit`:
`gpg-private-key`, `gpg-passphrase`, `gpg-fingerprint`, `nfpm-gpg-private-key`, `discord-webhook-id`,
`discord-webhook-token`. Callers from another organisation must pass them explicitly, as secrets are not inherited
across organisations.

##### What the local goreleaser config must respect

* **Do not let goreleaser generate the changelog.** The release notes come from git-cliff through `--release-notes`.
  Suppress goreleaser's own changelog with a filter, and *not* with `changelog.disable`, which would also discard the
  notes handed over by this workflow:

  ```yaml
  changelog:
    filters:
      exclude: [".*"]
  ```

* **Expect the assets to be signed with the bot key** when `enable-artifact-signing` is on. The passphrase and the
  fingerprint reach the config as `GPG_PASSPHRASE` and `GPG_FINGERPRINT`.

* **deb and rpm signing uses a separate key.** nfpm relies on go's openpgp library, which handles neither DSA keys nor
  subkey-only exports, so it needs a standalone RSA key. Its path reaches the config as `GPG_NFPM_KEY_FILE`.

* **Consider `release.mode: replace`** if a release job may be re-run after a partial failure, otherwise the second run
  fails on the already existing github release.

* **Mono-repos release from the root module tag.** goreleaser runs once, at the root of the repository, and the version
  is pinned to the root tag through `GORELEASER_CURRENT_TAG` — a mono-repo release pushes several tags on the same
  commit, and goreleaser would otherwise resolve an arbitrary one. The mono-repo support of goreleaser is not used.
  Binaries living in a nested module need a `dir:` in their `builds` entry.

* **Docker images are out of scope** for now: build them from a separate workflow.

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
