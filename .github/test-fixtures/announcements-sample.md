# Announcements test fixture

This file exists solely to exercise `webhook-announcements.yml` via
`local-webhook-announcements.yml`. Editing the `## Announcements` section below
and pushing to `master` triggers the local test, which dry-run-prints the
payloads it would post (no webhook secret required).

The format mirrors how go-openapi repos write announcements in their real
`README.md`: newest first, one top-level bullet per announcement
(`* **DATE** : summary`), with indented sub-bullets for detail.

## Announcements

* **2026-04-15** : added support for trailing "-" for arrays (v0.23.0)
  * this brings full support of [RFC6901][RFC6901]
  * API semantics remain essentially unaltered, with one documented exception around
    in-place mutation of arrays via a trailing "-"
  * types that implement the `JSONSetable` interface keep their behavior

* **2026-04-15** : added support for optional alternate JSON name providers
  * the default name provider is not fully aligned with the Go JSON stdlib
  * a new alternate provider (imported from `go-openapi/swag/jsonname`) is available

## Status

(end of fixture)

[RFC6901]: https://www.rfc-editor.org/rfc/rfc6901
