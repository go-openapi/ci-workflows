# {{ .Repo }}

<!-- Badges: status  -->
[![Tests][test-badge]][test-url] [![Coverage][cov-badge]][cov-url] [![CI vuln scan][vuln-scan-badge]][vuln-scan-url] [![CodeQL][codeql-badge]][codeql-url]
<!-- Badges: release & docker images  -->
<!-- Badges: code quality  -->
<!-- Badges: license & compliance -->
[![Release][release-badge]][release-url] [![Go Report Card][gocard-badge]][gocard-url] [![CodeFactor Grade][codefactor-badge]][codefactor-url] [![License][license-badge]][license-url]
<!-- Badges: documentation & support -->
<!-- Badges: others & stats -->
<!-- Slack badge disabled until I am able to restore a valid link to the chat -->
[![GoDoc][godoc-badge]][godoc-url] <!-- [![Slack Channel][slack-badge]][slack-url] -->[![go version][goversion-badge]][goversion-url] ![Top language][top-badge] ![Commits since latest release][commits-badge]

---

{{ .Title }}

## Status

API is stable.

## Import this library in your project

```cmd
go get github.com/{{ .Owner }}/{{ .Repo }}
```

## Basic usage

## Change log

See <https://github.com/{{ .Owner }}/{{ .Repo }}/releases>

<!--
## References
-->

## Licensing

This library ships under the [SPDX-License-Identifier: Apache-2.0](./LICENSE).

<!--
See the license [NOTICE](./NOTICE), which recalls the licensing terms of all the pieces of software
on top of which it has been built.
-->

<!--
## Limitations
-->

## Other documentation

* [All-time contributors](./CONTRIBUTORS.md)
* [Contributing guidelines](.github/CONTRIBUTING.md)
* [Maintainers documentation](docs/MAINTAINERS.md)
* [Code style](docs/STYLE.md)

<!-- Badges: status  -->
[test-badge]: https://github.com/{{ .Owner }}/{{ .Repo }}/actions/workflows/go-test.yml/badge.svg
[test-url]: https://github.com/{{ .Owner }}/{{ .Repo }}/actions/workflows/go-test.yml
[cov-badge]: https://codecov.io/gh/{{ .Owner }}/{{ .Repo }}/branch/master/graph/badge.svg
[cov-url]: https://codecov.io/gh/{{ .Owner }}/{{ .Repo }}
[vuln-scan-badge]: https://github.com/{{ .Owner }}/{{ .Repo }}/actions/workflows/scanner.yml/badge.svg
[vuln-scan-url]: https://github.com/{{ .Owner }}/{{ .Repo }}/actions/workflows/scanner.yml
[codeql-badge]: https://github.com/{{ .Owner }}/{{ .Repo }}/actions/workflows/codeql.yml/badge.svg
[codeql-url]: https://github.com/{{ .Owner }}/{{ .Repo }}/actions/workflows/codeql.yml
<!-- Badges: release & docker images  -->
[release-badge]: https://badge.fury.io/gh/{{ .Owner }}%2F{{ .Repo }}.svg
[release-url]: https://badge.fury.io/gh/{{ .Owner }}%2F{{ .Repo }}
[gomod-badge]: https://badge.fury.io/go/github.com%2F{{ .Owner }}%2F{{ .Repo }}.svg
[gomod-url]: https://badge.fury.io/go/github.com%2F{{ .Owner }}%2F{{ .Repo }}
<!-- Badges: code quality  -->
[gocard-badge]: https://goreportcard.com/badge/github.com/{{ .Owner }}/{{ .Repo }}
[gocard-url]: https://goreportcard.com/report/github.com/{{ .Owner }}/{{ .Repo }}
[codefactor-badge]: https://img.shields.io/codefactor/grade/github/{{ .Owner }}/{{ .Repo }}
[codefactor-url]: https://www.codefactor.io/repository/github/{{ .Owner }}/{{ .Repo }}
<!-- Badges: documentation & support -->
[doc-badge]: https://img.shields.io/badge/doc-site-blue?link=https%3A%2F%2Fgoswagger.io%2F{{ .Owner }}%2F
[doc-url]: https://goswagger.io/{{ .Owner }}
[godoc-badge]: https://pkg.go.dev/badge/github.com/{{ .Owner }}/{{ .Repo }}
[godoc-url]: http://pkg.go.dev/github.com/{{ .Owner }}/{{ .Repo }}
[slack-badge]: https://slackin.goswagger.io/badge.svg
[slack-url]: https://slackin.goswagger.io
<!-- Badges: license & compliance -->
[license-badge]: http://img.shields.io/badge/license-Apache%20v2-orange.svg
[license-url]: https://github.com/{{ .Owner }}/{{ .Repo }}/?tab=Apache-2.0-1-ov-file#readme
<!-- Badges: others & stats -->
[goversion-badge]: https://img.shields.io/github/go-mod/go-version/{{ .Owner }}/{{ .Repo }}
[goversion-url]: https://github.com/{{ .Owner }}/{{ .Repo }}/blob/master/go.mod
[top-badge]: https://img.shields.io/github/languages/top/{{ .Owner }}/{{ .Repo }}
[commits-badge]: https://img.shields.io/github/commits-since/{{ .Owner }}/{{ .Repo }}/latest
