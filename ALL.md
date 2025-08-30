This is a test of the spellcheck dictionary

Bump md

# Swagger 2.0 specification schema

This folder contains the Swagger 2.0 specification schema files maintained here:

https://github.com/reverb/swagger-spec/blob/master/schemas/v2.0# OpenAPI v2 object model [![Build Status](https://github.com/go-openapi/spec/actions/workflows/go-test.yml/badge.svg)](https://github.com/go-openapi/spec/actions?query=workflow%3A"go+test") [![codecov](https://codecov.io/gh/go-openapi/spec/branch/master/graph/badge.svg)](https://codecov.io/gh/go-openapi/spec)

[![Slack Status](https://slackin.goswagger.io/badge.svg)](https://slackin.goswagger.io)
[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/go-openapi/spec/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-openapi/spec.svg)](https://pkg.go.dev/github.com/go-openapi/spec)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-openapi/spec)](https://goreportcard.com/report/github.com/go-openapi/spec)

The object model for OpenAPI specification documents.

### FAQ

* What does this do?

> 1. This package knows how to marshal and unmarshal Swagger API specifications into a golang object model
> 2. It knows how to resolve $ref and expand them to make a single root document

* How does it play with the rest of the go-openapi packages ?

> 1. This package is at the core of the go-openapi suite of packages and [code generator](https://github.com/go-swagger/go-swagger)
> 2. There is a [spec loading package](https://github.com/go-openapi/loads) to fetch specs as JSON or YAML from local or remote locations
> 3. There is a [spec validation package](https://github.com/go-openapi/validate) built on top of it
> 4. There is a [spec analysis package](https://github.com/go-openapi/analysis) built on top of it, to analyze, flatten, fix and merge spec documents

* Does this library support OpenAPI 3?

> No.
> This package currently only supports OpenAPI 2.0 (aka Swagger 2.0).
> There is no plan to make it evolve toward supporting OpenAPI 3.x.
> This [discussion thread](https://github.com/go-openapi/spec/issues/21) relates the full story.
>
> An early attempt to support Swagger 3 may be found at: https://github.com/go-openapi/spec3

* Does the unmarshaling support YAML?

> Not directly. The exposed types know only how to unmarshal from JSON.
>
> In order to load a YAML document as a Swagger spec, you need to use the loaders provided by
> github.com/go-openapi/loads
>
> Take a look at the example there: https://pkg.go.dev/github.com/go-openapi/loads#example-Spec
>
> See also https://github.com/go-openapi/spec/issues/164

* How can I validate a spec?

> Validation is provided by [the validate package](http://github.com/go-openapi/validate)

* Why do we have an `ID` field for `Schema` which is not part of the swagger spec?

> We found jsonschema compatibility more important: since `id` in jsonschema influences
> how `$ref` are resolved.
> This `id` does not conflict with any property named `id`.
>
> See also https://github.com/go-openapi/spec/issues/23
# Contributor Covenant Code of Conduct

## Our Pledge

In the interest of fostering an open and welcoming environment, we as
contributors and maintainers pledge to making participation in our project and
our community a harassment-free experience for everyone, regardless of age, body
size, disability, ethnicity, gender identity and expression, level of experience,
nationality, personal appearance, race, religion, or sexual identity and
orientation.

## Our Standards

Examples of behavior that contributes to creating a positive environment
include:

* Using welcoming and inclusive language
* Being respectful of differing viewpoints and experiences
* Gracefully accepting constructive criticism
* Focusing on what is best for the community
* Showing empathy towards other community members

Examples of unacceptable behavior by participants include:

* The use of sexualized language or imagery and unwelcome sexual attention or
advances
* Trolling, insulting/derogatory comments, and personal or political attacks
* Public or private harassment
* Publishing others' private information, such as a physical or electronic
  address, without explicit permission
* Other conduct which could reasonably be considered inappropriate in a
  professional setting

## Our Responsibilities

Project maintainers are responsible for clarifying the standards of acceptable
behavior and are expected to take appropriate and fair corrective action in
response to any instances of unacceptable behavior.

Project maintainers have the right and responsibility to remove, edit, or
reject comments, commits, code, wiki edits, issues, and other contributions
that are not aligned to this Code of Conduct, or to ban temporarily or
permanently any contributor for other behaviors that they deem inappropriate,
threatening, offensive, or harmful.

## Scope

This Code of Conduct applies both within project spaces and in public spaces
when an individual is representing the project or its community. Examples of
representing a project or community include using an official project e-mail
address, posting via an official social media account, or acting as an appointed
representative at an online or offline event. Representation of a project may be
further defined and clarified by project maintainers.

## Enforcement

Instances of abusive, harassing, or otherwise unacceptable behavior may be
reported by contacting the project team at ivan+abuse@flanders.co.nz. All
complaints will be reviewed and investigated and will result in a response that
is deemed necessary and appropriate to the circumstances. The project team is
obligated to maintain confidentiality with regard to the reporter of an incident.
Further details of specific enforcement policies may be posted separately.

Project maintainers who do not follow or enforce the Code of Conduct in good
faith may face temporary or permanent repercussions as determined by other
members of the project's leadership.

## Attribution

This Code of Conduct is adapted from the [Contributor Covenant][homepage], version 1.4,
available at [http://contributor-covenant.org/version/1/4][version]

[homepage]: http://contributor-covenant.org
[version]: http://contributor-covenant.org/version/1/4/
## Contribution Guidelines

### Pull requests are always welcome

We are always thrilled to receive pull requests, and do our best to
process them as fast as possible. Not sure if that typo is worth a pull
request? Do it! We will appreciate it.

If your pull request is not accepted on the first try, don't be
discouraged! If there's a problem with the implementation, hopefully you
received feedback on what to improve.

We're trying very hard to keep go-swagger lean and focused. We don't want it
to do everything for everybody. This means that we might decide against
incorporating a new feature. However, there might be a way to implement
that feature *on top of* go-swagger.


### Conventions

Fork the repo and make changes on your fork in a feature branch:

- If it's a bugfix branch, name it XXX-something where XXX is the number of the
  issue
- If it's a feature branch, create an enhancement issue to announce your
  intentions, and name it XXX-something where XXX is the number of the issue.

Submit unit tests for your changes.  Go has a great test framework built in; use
it! Take a look at existing tests for inspiration. Run the full test suite on
your branch before submitting a pull request.

Update the documentation when creating or modifying features. Test
your documentation changes for clarity, concision, and correctness, as
well as a clean documentation build. See ``docs/README.md`` for more
information on building the docs and how docs get released.

Write clean code. Universally formatted code promotes ease of writing, reading,
and maintenance. Always run `gofmt -s -w file.go` on each changed file before
committing your changes. Most editors have plugins that do this automatically.

Pull requests descriptions should be as clear as possible and include a
reference to all the issues that they address.

Pull requests must not contain commits from other users or branches.

Commit messages must start with a capitalized and short summary (max. 50
chars) written in the imperative, followed by an optional, more detailed
explanatory text which is separated from the summary by an empty line.

Code review comments may be added to your pull request. Discuss, then make the
suggested modifications and push additional commits to your feature branch. Be
sure to post a comment after pushing. The new commits will show up in the pull
request automatically, but the reviewers will not be notified unless you
comment.

Before the pull request is merged, make sure that you squash your commits into
logical units of work using `git rebase -i` and `git push -f`. After every
commit the test suite should be passing. Include documentation changes in the
same commit so that a revert would remove all traces of the feature or fix.

Commits that fix or close an issue should include a reference like `Closes #XXX`
or `Fixes #XXX`, which will automatically close the issue when merged.

### Sign your work

The sign-off is a simple line at the end of the explanation for the
patch, which certifies that you wrote it or otherwise have the right to
pass it on as an open-source patch.  The rules are pretty simple: if you
can certify the below (from
[developercertificate.org](http://developercertificate.org/)):

```
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
660 York Street, Suite 102,
San Francisco, CA 94110 USA

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.


Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

then you just add a line to every git commit message:

    Signed-off-by: Joe Smith <joe@gmail.com>

using your real name (sorry, no pseudonyms or anonymous contributions.)

You can add the sign off when creating the git commit via `git commit -s`.
# gojsonreference [![Build Status](https://github.com/go-openapi/jsonreference/actions/workflows/go-test.yml/badge.svg)](https://github.com/go-openapi/jsonreference/actions?query=workflow%3A"go+test") [![codecov](https://codecov.io/gh/go-openapi/jsonreference/branch/master/graph/badge.svg)](https://codecov.io/gh/go-openapi/jsonreference)

[![Slack Status](https://slackin.goswagger.io/badge.svg)](https://slackin.goswagger.io)
[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/go-openapi/jsonreference/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-openapi/jsonreference.svg)](https://pkg.go.dev/github.com/go-openapi/jsonreference)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-openapi/jsonreference)](https://goreportcard.com/report/github.com/go-openapi/jsonreference)

An implementation of JSON Reference - Go language

## Status
Feature complete. Stable API

## Dependencies
* https://github.com/go-openapi/jsonpointer

## References

* http://tools.ietf.org/html/draft-ietf-appsawg-json-pointer-07
* http://tools.ietf.org/html/draft-pbryan-zyp-json-ref-03
# Contributor Covenant Code of Conduct

## Our Pledge

In the interest of fostering an open and welcoming environment, we as
contributors and maintainers pledge to making participation in our project and
our community a harassment-free experience for everyone, regardless of age, body
size, disability, ethnicity, gender identity and expression, level of experience,
nationality, personal appearance, race, religion, or sexual identity and
orientation.

## Our Standards

Examples of behavior that contributes to creating a positive environment
include:

* Using welcoming and inclusive language
* Being respectful of differing viewpoints and experiences
* Gracefully accepting constructive criticism
* Focusing on what is best for the community
* Showing empathy towards other community members

Examples of unacceptable behavior by participants include:

* The use of sexualized language or imagery and unwelcome sexual attention or
advances
* Trolling, insulting/derogatory comments, and personal or political attacks
* Public or private harassment
* Publishing others' private information, such as a physical or electronic
  address, without explicit permission
* Other conduct which could reasonably be considered inappropriate in a
  professional setting

## Our Responsibilities

Project maintainers are responsible for clarifying the standards of acceptable
behavior and are expected to take appropriate and fair corrective action in
response to any instances of unacceptable behavior.

Project maintainers have the right and responsibility to remove, edit, or
reject comments, commits, code, wiki edits, issues, and other contributions
that are not aligned to this Code of Conduct, or to ban temporarily or
permanently any contributor for other behaviors that they deem inappropriate,
threatening, offensive, or harmful.

## Scope

This Code of Conduct applies both within project spaces and in public spaces
when an individual is representing the project or its community. Examples of
representing a project or community include using an official project e-mail
address, posting via an official social media account, or acting as an appointed
representative at an online or offline event. Representation of a project may be
further defined and clarified by project maintainers.

## Enforcement

Instances of abusive, harassing, or otherwise unacceptable behavior may be
reported by contacting the project team at ivan+abuse@flanders.co.nz. All
complaints will be reviewed and investigated and will result in a response that
is deemed necessary and appropriate to the circumstances. The project team is
obligated to maintain confidentiality with regard to the reporter of an incident.
Further details of specific enforcement policies may be posted separately.

Project maintainers who do not follow or enforce the Code of Conduct in good
faith may face temporary or permanent repercussions as determined by other
members of the project's leadership.

## Attribution

This Code of Conduct is adapted from the [Contributor Covenant][homepage], version 1.4,
available at [http://contributor-covenant.org/version/1/4][version]

[homepage]: http://contributor-covenant.org
[version]: http://contributor-covenant.org/version/1/4/
## Contribution Guidelines

### Pull requests are always welcome

We are always thrilled to receive pull requests, and do our best to
process them as fast as possible. Not sure if that typo is worth a pull
request? Do it! We will appreciate it.

If your pull request is not accepted on the first try, don't be
discouraged! If there's a problem with the implementation, hopefully you
received feedback on what to improve.

We're trying very hard to keep go-swagger lean and focused. We don't want it
to do everything for everybody. This means that we might decide against
incorporating a new feature. However, there might be a way to implement
that feature *on top of* go-swagger.


### Conventions

Fork the repo and make changes on your fork in a feature branch:

- If it's a bugfix branch, name it XXX-something where XXX is the number of the
  issue
- If it's a feature branch, create an enhancement issue to announce your
  intentions, and name it XXX-something where XXX is the number of the issue.

Submit unit tests for your changes.  Go has a great test framework built in; use
it! Take a look at existing tests for inspiration. Run the full test suite on
your branch before submitting a pull request.

Update the documentation when creating or modifying features. Test
your documentation changes for clarity, concision, and correctness, as
well as a clean documentation build. See ``docs/README.md`` for more
information on building the docs and how docs get released.

Write clean code. Universally formatted code promotes ease of writing, reading,
and maintenance. Always run `gofmt -s -w file.go` on each changed file before
committing your changes. Most editors have plugins that do this automatically.

Pull requests descriptions should be as clear as possible and include a
reference to all the issues that they address.

Pull requests must not contain commits from other users or branches.

Commit messages must start with a capitalized and short summary (max. 50
chars) written in the imperative, followed by an optional, more detailed
explanatory text which is separated from the summary by an empty line.

Code review comments may be added to your pull request. Discuss, then make the
suggested modifications and push additional commits to your feature branch. Be
sure to post a comment after pushing. The new commits will show up in the pull
request automatically, but the reviewers will not be notified unless you
comment.

Before the pull request is merged, make sure that you squash your commits into
logical units of work using `git rebase -i` and `git push -f`. After every
commit the test suite should be passing. Include documentation changes in the
same commit so that a revert would remove all traces of the feature or fix.

Commits that fix or close an issue should include a reference like `Closes #XXX`
or `Fixes #XXX`, which will automatically close the issue when merged.

### Sign your work

The sign-off is a simple line at the end of the explanation for the
patch, which certifies that you wrote it or otherwise have the right to
pass it on as an open-source patch.  The rules are pretty simple: if you
can certify the below (from
[developercertificate.org](http://developercertificate.org/)):

```
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
660 York Street, Suite 102,
San Francisco, CA 94110 USA

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.


Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

then you just add a line to every git commit message:

    Signed-off-by: Joe Smith <joe@gmail.com>

using your real name (sorry, no pseudonyms or anonymous contributions.)

You can add the sign off when creating the git commit via `git commit -s`.
# doc

Documentation check and generation tools.
# ci-workflows[![Build Status](https://github.com/go-openapi/ci-workflows/actions/workflows/go-test.yml/badge.svg)](https://github.com/go-openapi/ci-workflows/actions?query=workflow%3A"go+test") [![codecov](https://codecov.io/gh/go-openapi/ci-workflows/branch/master/graph/badge.svg)](https://codecov.io/gh/go-openapi/ci-workflows)

[![Slack Status](https://slackin.goswagger.io/badge.svg)](https://slackin.goswagger.io)
[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/go-openapi/ci-workflows/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-openapi/ci-workflows.svg)](https://pkg.go.dev/github.com/go-openapi/ci-workflows)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-openapi/ci-workflows)](https://goreportcard.com/report/github.com/go-openapi/ci-workflows)

Common Continuous Integration (`CI`) workflows and setup for go-openapi repos.

* shared github action workflows
* shared `dependabot` configuration (TODO)
* shared `golangci-lint` configuration (**BLOCKED**)

> NOTE: at this moment, it is difficult to share the golangci-lint config,
> so that one is not shared yet.

## Motivation

It took a while, but we eventually managed to align all checks, tests and
dependabot rules declared in the family of go-openapi repos.

Now we'd like to be able to maintain, enrich and improve these checks without
worrying too much about the burden to replicate it about a dozen times.

## Contemplated enhancements

In no particular order:

* [x] ui: enrich github actions UI with a job summary
* [x] doc: add markdown linting for docs
* [x] doc: add spellcheck for docs (and code?)
* [ ] version common workflows, so we can limit the impact of a change
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
# Contributor Covenant Code of Conduct

## Our Pledge

In the interest of fostering an open and welcoming environment, we as
contributors and maintainers pledge to making participation in our project and
our community a harassment-free experience for everyone, regardless of age, body
size, disability, ethnicity, gender identity and expression, level of experience,
nationality, personal appearance, race, religion, or sexual identity and
orientation.

## Our Standards

Examples of behavior that contributes to creating a positive environment
include:

* Using welcoming and inclusive language
* Being respectful of differing viewpoints and experiences
* Gracefully accepting constructive criticism
* Focusing on what is best for the community
* Showing empathy towards other community members

Examples of unacceptable behavior by participants include:

* The use of sexualized language or imagery and unwelcome sexual attention or
advances
* Trolling, insulting/derogatory comments, and personal or political attacks
* Public or private harassment
* Publishing others' private information, such as a physical or electronic
  address, without explicit permission
* Other conduct which could reasonably be considered inappropriate in a
  professional setting

## Our Responsibilities

Project maintainers are responsible for clarifying the standards of acceptable
behavior and are expected to take appropriate and fair corrective action in
response to any instances of unacceptable behavior.

Project maintainers have the right and responsibility to remove, edit, or
reject comments, commits, code, wiki edits, issues, and other contributions
that are not aligned to this Code of Conduct, or to ban temporarily or
permanently any contributor for other behaviors that they deem inappropriate,
threatening, offensive, or harmful.

## Scope

This Code of Conduct applies both within project spaces and in public spaces
when an individual is representing the project or its community. Examples of
representing a project or community include using an official project e-mail
address, posting via an official social media account, or acting as an appointed
representative at an online or offline event. Representation of a project may be
further defined and clarified by project maintainers.

## Enforcement

Instances of abusive, harassing, or otherwise unacceptable behavior may be
reported by contacting the project team at ivan+abuse@flanders.co.nz. All
complaints will be reviewed and investigated and will result in a response that
is deemed necessary and appropriate to the circumstances. The project team is
obligated to maintain confidentiality with regard to the reporter of an incident.
Further details of specific enforcement policies may be posted separately.

Project maintainers who do not follow or enforce the Code of Conduct in good
faith may face temporary or permanent repercussions as determined by other
members of the project's leadership.

## Attribution

This Code of Conduct is adapted from the [Contributor Covenant][homepage], version 1.4,
available at [http://contributor-covenant.org/version/1/4][version]

[homepage]: http://contributor-covenant.org
[version]: http://contributor-covenant.org/version/1/4/
## Contribution Guidelines

### Pull requests are always welcome

We are always thrilled to receive pull requests, and do our best to
process them as fast as possible. Not sure if that typo is worth a pull
request? Do it! We will appreciate it.

If your pull request is not accepted on the first try, don't be
discouraged! If there's a problem with the implementation, hopefully you
received feedback on what to improve.

We're trying very hard to keep go-swagger lean and focused. We don't want it
to do everything for everybody. This means that we might decide against
incorporating a new feature. However, there might be a way to implement
that feature *on top of* go-swagger.


### Conventions

Fork the repo and make changes on your fork in a feature branch:

- If it's a bugfix branch, name it XXX-something where XXX is the number of the
  issue
- If it's a feature branch, create an enhancement issue to announce your
  intentions, and name it XXX-something where XXX is the number of the issue.

Submit unit tests for your changes.  Go has a great test framework built in; use
it! Take a look at existing tests for inspiration. Run the full test suite on
your branch before submitting a pull request.

Update the documentation when creating or modifying features. Test
your documentation changes for clarity, concision, and correctness, as
well as a clean documentation build. See ``docs/README.md`` for more
information on building the docs and how docs get released.

Write clean code. Universally formatted code promotes ease of writing, reading,
and maintenance. Always run `gofmt -s -w file.go` on each changed file before
committing your changes. Most editors have plugins that do this automatically.

Pull requests descriptions should be as clear as possible and include a
reference to all the issues that they address.

Pull requests must not contain commits from other users or branches.

Commit messages must start with a capitalized and short summary (max. 50
chars) written in the imperative, followed by an optional, more detailed
explanatory text which is separated from the summary by an empty line.

Code review comments may be added to your pull request. Discuss, then make the
suggested modifications and push additional commits to your feature branch. Be
sure to post a comment after pushing. The new commits will show up in the pull
request automatically, but the reviewers will not be notified unless you
comment.

Before the pull request is merged, make sure that you squash your commits into
logical units of work using `git rebase -i` and `git push -f`. After every
commit the test suite should be passing. Include documentation changes in the
same commit so that a revert would remove all traces of the feature or fix.

Commits that fix or close an issue should include a reference like `Closes #XXX`
or `Fixes #XXX`, which will automatically close the issue when merged.

### Sign your work

The sign-off is a simple line at the end of the explanation for the
patch, which certifies that you wrote it or otherwise have the right to
pass it on as an open-source patch.  The rules are pretty simple: if you
can certify the below (from
[developercertificate.org](http://developercertificate.org/)):

```
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
660 York Street, Suite 102,
San Francisco, CA 94110 USA

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.


Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

then you just add a line to every git commit message:

    Signed-off-by: Joe Smith <joe@gmail.com>

using your real name (sorry, no pseudonyms or anonymous contributions.)

You can add the sign off when creating the git commit via `git commit -s`.
#  Audit

Audit of github repos to check for inconsistencies.
# inflect [![Build Status](https://github.com/go-openapi/inflect/actions/workflows/go-test.yml/badge.svg)](https://github.com/go-openapi/inflect/actions?query=workflow%3A"go+test") [![codecov](https://codecov.io/gh/go-openapi/inflect/branch/master/graph/badge.svg)](https://codecov.io/gh/go-openapi/inflect)

[![Slack Status](https://slackin.goswagger.io/badge.svg)](https://slackin.goswagger.io)
[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/go-openapi/inflect/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-openapi/inflect.svg)](https://pkg.go.dev/github.com/go-openapi/inflect)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-openapi/inflect)](https://goreportcard.com/report/github.com/go-openapi/inflect)

A package to pluralize words.

Originally forked from fork of https://bitbucket.org/pkg/inflect under a MIT License.

A golang library applying grammar rules to English words.

> This package provides a basic set of functions applying
> grammar rules to inflect English words, modify case style
> (Capitalize, camelCase, snake_case, etc.).
>
> Acronyms are properly handled. A common use case is word pluralization.


This library is not used at all by other go-openapi packages and is somewhat redundant with
go-openapi/swag/mangling (for camelcase etc).

Currently we have one single dependency in one place in a go-swagger template (used as a funcmap).

## Note to maintainers

* Aug. 2025: CI workflows have now moved to a shared repository go-openapi/ci-workflows
## Contribution Guidelines

### Pull requests are always welcome

We are always thrilled to receive pull requests, and do our best to
process them as fast as possible. Not sure if that typo is worth a pull
request? Do it! We will appreciate it.

If your pull request is not accepted on the first try, don't be
discouraged! If there's a problem with the implementation, hopefully you
received feedback on what to improve.

We're trying very hard to keep go-swagger lean and focused. We don't want it
to do everything for everybody. This means that we might decide against
incorporating a new feature. However, there might be a way to implement
that feature *on top of* go-swagger.


### Conventions

Fork the repo and make changes on your fork in a feature branch:

- If it's a bugfix branch, name it XXX-something where XXX is the number of the
  issue
- If it's a feature branch, create an enhancement issue to announce your
  intentions, and name it XXX-something where XXX is the number of the issue.

Submit unit tests for your changes.  Go has a great test framework built in; use
it! Take a look at existing tests for inspiration. Run the full test suite on
your branch before submitting a pull request.

Update the documentation when creating or modifying features. Test
your documentation changes for clarity, concision, and correctness, as
well as a clean documentation build. See ``docs/README.md`` for more
information on building the docs and how docs get released.

Write clean code. Universally formatted code promotes ease of writing, reading,
and maintenance. Always run `gofmt -s -w file.go` on each changed file before
committing your changes. Most editors have plugins that do this automatically.

Pull requests descriptions should be as clear as possible and include a
reference to all the issues that they address.

Pull requests must not contain commits from other users or branches.

Commit messages must start with a capitalized and short summary (max. 50
chars) written in the imperative, followed by an optional, more detailed
explanatory text which is separated from the summary by an empty line.

Code review comments may be added to your pull request. Discuss, then make the
suggested modifications and push additional commits to your feature branch. Be
sure to post a comment after pushing. The new commits will show up in the pull
request automatically, but the reviewers will not be notified unless you
comment.

Before the pull request is merged, make sure that you squash your commits into
logical units of work using `git rebase -i` and `git push -f`. After every
commit the test suite should be passing. Include documentation changes in the
same commit so that a revert would remove all traces of the feature or fix.

Commits that fix or close an issue should include a reference like `Closes #XXX`
or `Fixes #XXX`, which will automatically close the issue when merged.

### Sign your work

The sign-off is a simple line at the end of the explanation for the
patch, which certifies that you wrote it or otherwise have the right to
pass it on as an open-source patch.  The rules are pretty simple: if you
can certify the below (from
[developercertificate.org](http://developercertificate.org/)):

```
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
660 York Street, Suite 102,
San Francisco, CA 94110 USA

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.


Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

then you just add a line to every git commit message:

    Signed-off-by: Joe Smith <joe@gmail.com>

using your real name (sorry, no pseudonyms or anonymous contributions.)

You can add the sign off when creating the git commit via `git commit -s`.
# Validation helpers [![Build Status](https://github.com/go-openapi/validate/actions/workflows/go-test.yml/badge.svg)](https://github.com/go-openapi/validate/actions?query=workflow%3A"go+test") [![codecov](https://codecov.io/gh/go-openapi/validate/branch/master/graph/badge.svg)](https://codecov.io/gh/go-openapi/validate)

[![Slack Status](https://slackin.goswagger.io/badge.svg)](https://slackin.goswagger.io)
[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/go-openapi/validate/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-openapi/validate.svg)](https://pkg.go.dev/github.com/go-openapi/validate)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-openapi/validate)](https://goreportcard.com/report/github.com/go-openapi/validate)

This package provides helpers to validate Swagger 2.0. specification (aka OpenAPI 2.0). 

Reference can be found here: https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md.

## What's inside?

* A validator for Swagger specifications
* A validator for JSON schemas draft4
* Helper functions to validate individual values (used by code generated by [go-swagger](https://github.com/go-swagger/go-swagger)).
  * Required, RequiredNumber, RequiredString
  * ReadOnly
  * UniqueItems, MaxItems, MinItems
  * Enum, EnumCase
  * Pattern, MinLength, MaxLength
  * Minimum, Maximum, MultipleOf
  * FormatOf

[Documentation](https://pkg.go.dev/github.com/go-openapi/validate)

## FAQ

* Does this library support OpenAPI 3?

> No.
> This package currently only supports OpenAPI 2.0 (aka Swagger 2.0).
> There is no plan to make it evolve toward supporting OpenAPI 3.x.
> This [discussion thread](https://github.com/go-openapi/spec/issues/21) relates the full story.
>
> An early attempt to support Swagger 3 may be found at: https://github.com/go-openapi/spec3
# Contributor Covenant Code of Conduct

## Our Pledge

In the interest of fostering an open and welcoming environment, we as
contributors and maintainers pledge to making participation in our project and
our community a harassment-free experience for everyone, regardless of age, body
size, disability, ethnicity, gender identity and expression, level of experience,
nationality, personal appearance, race, religion, or sexual identity and
orientation.

## Our Standards

Examples of behavior that contributes to creating a positive environment
include:

* Using welcoming and inclusive language
* Being respectful of differing viewpoints and experiences
* Gracefully accepting constructive criticism
* Focusing on what is best for the community
* Showing empathy towards other community members

Examples of unacceptable behavior by participants include:

* The use of sexualized language or imagery and unwelcome sexual attention or
advances
* Trolling, insulting/derogatory comments, and personal or political attacks
* Public or private harassment
* Publishing others' private information, such as a physical or electronic
  address, without explicit permission
* Other conduct which could reasonably be considered inappropriate in a
  professional setting

## Our Responsibilities

Project maintainers are responsible for clarifying the standards of acceptable
behavior and are expected to take appropriate and fair corrective action in
response to any instances of unacceptable behavior.

Project maintainers have the right and responsibility to remove, edit, or
reject comments, commits, code, wiki edits, issues, and other contributions
that are not aligned to this Code of Conduct, or to ban temporarily or
permanently any contributor for other behaviors that they deem inappropriate,
threatening, offensive, or harmful.

## Scope

This Code of Conduct applies both within project spaces and in public spaces
when an individual is representing the project or its community. Examples of
representing a project or community include using an official project e-mail
address, posting via an official social media account, or acting as an appointed
representative at an online or offline event. Representation of a project may be
further defined and clarified by project maintainers.

## Enforcement

Instances of abusive, harassing, or otherwise unacceptable behavior may be
reported by contacting the project team at ivan+abuse@flanders.co.nz. All
complaints will be reviewed and investigated and will result in a response that
is deemed necessary and appropriate to the circumstances. The project team is
obligated to maintain confidentiality with regard to the reporter of an incident.
Further details of specific enforcement policies may be posted separately.

Project maintainers who do not follow or enforce the Code of Conduct in good
faith may face temporary or permanent repercussions as determined by other
members of the project's leadership.

## Attribution

This Code of Conduct is adapted from the [Contributor Covenant][homepage], version 1.4,
available at [http://contributor-covenant.org/version/1/4][version]

[homepage]: http://contributor-covenant.org
[version]: http://contributor-covenant.org/version/1/4/
## Contribution Guidelines

### Pull requests are always welcome

We are always thrilled to receive pull requests, and do our best to
process them as fast as possible. Not sure if that typo is worth a pull
request? Do it! We will appreciate it.

If your pull request is not accepted on the first try, don't be
discouraged! If there's a problem with the implementation, hopefully you
received feedback on what to improve.

We're trying very hard to keep go-swagger lean and focused. We don't want it
to do everything for everybody. This means that we might decide against
incorporating a new feature. However, there might be a way to implement
that feature *on top of* go-swagger.


### Conventions

Fork the repo and make changes on your fork in a feature branch:

- If it's a bugfix branch, name it XXX-something where XXX is the number of the
  issue
- If it's a feature branch, create an enhancement issue to announce your
  intentions, and name it XXX-something where XXX is the number of the issue.

Submit unit tests for your changes.  Go has a great test framework built in; use
it! Take a look at existing tests for inspiration. Run the full test suite on
your branch before submitting a pull request.

Update the documentation when creating or modifying features. Test
your documentation changes for clarity, concision, and correctness, as
well as a clean documentation build. See ``docs/README.md`` for more
information on building the docs and how docs get released.

Write clean code. Universally formatted code promotes ease of writing, reading,
and maintenance. Always run `gofmt -s -w file.go` on each changed file before
committing your changes. Most editors have plugins that do this automatically.

Pull requests descriptions should be as clear as possible and include a
reference to all the issues that they address.

Pull requests must not contain commits from other users or branches.

Commit messages must start with a capitalized and short summary (max. 50
chars) written in the imperative, followed by an optional, more detailed
explanatory text which is separated from the summary by an empty line.

Code review comments may be added to your pull request. Discuss, then make the
suggested modifications and push additional commits to your feature branch. Be
sure to post a comment after pushing. The new commits will show up in the pull
request automatically, but the reviewers will not be notified unless you
comment.

Before the pull request is merged, make sure that you squash your commits into
logical units of work using `git rebase -i` and `git push -f`. After every
commit the test suite should be passing. Include documentation changes in the
same commit so that a revert would remove all traces of the feature or fix.

Commits that fix or close an issue should include a reference like `Closes #XXX`
or `Fixes #XXX`, which will automatically close the issue when merged.

### Sign your work

The sign-off is a simple line at the end of the explanation for the
patch, which certifies that you wrote it or otherwise have the right to
pass it on as an open-source patch.  The rules are pretty simple: if you
can certify the below (from
[developercertificate.org](http://developercertificate.org/)):

```
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
660 York Street, Suite 102,
San Francisco, CA 94110 USA

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.


Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

then you just add a line to every git commit message:

    Signed-off-by: Joe Smith <joe@gmail.com>

using your real name (sorry, no pseudonyms or anonymous contributions.)

You can add the sign off when creating the git commit via `git commit -s`.
# Benchmark

Validating the Kubernetes Swagger API

## v0.22.6: 60,000,000 allocs
```
goos: linux
goarch: amd64
pkg: github.com/go-openapi/validate
cpu: AMD Ryzen 7 5800X 8-Core Processor
Benchmark_KubernetesSpec/validating_kubernetes_API-16         	       1	8549863982 ns/op	7067424936 B/op	59583275 allocs/op
```

## After refact PR: minor but noticable improvements: 25,000,000 allocs
```
go test -bench Spec
goos: linux
goarch: amd64
pkg: github.com/go-openapi/validate
cpu: AMD Ryzen 7 5800X 8-Core Processor
Benchmark_KubernetesSpec/validating_kubernetes_API-16         	       1	4064535557 ns/op	3379715592 B/op	25320330 allocs/op
```

## After reduce GC pressure PR: 17,000,000 allocs
```
goos: linux
goarch: amd64
pkg: github.com/go-openapi/validate
cpu: AMD Ryzen 7 5800X 8-Core Processor             
Benchmark_KubernetesSpec/validating_kubernetes_API-16         	       1	3758414145 ns/op	2593881496 B/op	17111373 allocs/op
```
# Denco [![Build Status](https://travis-ci.org/naoina/denco.png?branch=master)](https://travis-ci.org/naoina/denco)

The fast and flexible HTTP request router for [Go](http://golang.org).

Denco is based on Double-Array implementation of [Kocha-urlrouter](https://github.com/naoina/kocha-urlrouter).
However, Denco is optimized and some features added.

## Features

* Fast (See [go-http-routing-benchmark](https://github.com/naoina/go-http-routing-benchmark))
* [URL patterns](#url-patterns) (`/foo/:bar` and `/foo/*wildcard`)
* Small (but enough) URL router API
* HTTP request multiplexer like `http.ServeMux`

## Installation

    go get -u github.com/go-openapi/runtime/middleware/denco

## Using as HTTP request multiplexer

```go
package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/go-openapi/runtime/middleware/denco"
)

func Index(w http.ResponseWriter, r *http.Request, params denco.Params) {
    fmt.Fprintf(w, "Welcome to Denco!\n")
}

func User(w http.ResponseWriter, r *http.Request, params denco.Params) {
    fmt.Fprintf(w, "Hello %s!\n", params.Get("name"))
}

func main() {
    mux := denco.NewMux()
    handler, err := mux.Build([]denco.Handler{
        mux.GET("/", Index),
        mux.GET("/user/:name", User),
        mux.POST("/user/:name", User),
    })
    if err != nil {
        panic(err)
    }
    log.Fatal(http.ListenAndServe(":8080", handler))
}
```

## Using as URL router

```go
package main

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware/denco"
)

type route struct {
	name string
}

func main() {
	router := denco.New()
	router.Build([]denco.Record{
		{"/", &route{"root"}},
		{"/user/:id", &route{"user"}},
		{"/user/:name/:id", &route{"username"}},
		{"/static/*filepath", &route{"static"}},
	})

	data, params, found := router.Lookup("/")
	// print `&main.route{name:"root"}, denco.Params(nil), true`.
	fmt.Printf("%#v, %#v, %#v\n", data, params, found)

	data, params, found = router.Lookup("/user/hoge")
	// print `&main.route{name:"user"}, denco.Params{denco.Param{Name:"id", Value:"hoge"}}, true`.
	fmt.Printf("%#v, %#v, %#v\n", data, params, found)

	data, params, found = router.Lookup("/user/hoge/7")
	// print `&main.route{name:"username"}, denco.Params{denco.Param{Name:"name", Value:"hoge"}, denco.Param{Name:"id", Value:"7"}}, true`.
	fmt.Printf("%#v, %#v, %#v\n", data, params, found)

	data, params, found = router.Lookup("/static/path/to/file")
	// print `&main.route{name:"static"}, denco.Params{denco.Param{Name:"filepath", Value:"path/to/file"}}, true`.
	fmt.Printf("%#v, %#v, %#v\n", data, params, found)
}
```

See [Godoc](http://godoc.org/github.com/go-openapi/runtime/middleware/denco) for more details.

## Getting the value of path parameter

You can get the value of path parameter by 2 ways.

1. Using [`denco.Params.Get`](http://godoc.org/github.com/go-openapi/runtime/middleware/denco#Params.Get) method
2. Find by loop

```go
package main

import (
    "fmt"

    "github.com/go-openapi/runtime/middleware/denco"
)

func main() {
    router := denco.New()
    if err := router.Build([]denco.Record{
        {"/user/:name/:id", "route1"},
    }); err != nil {
        panic(err)
    }

    // 1. Using denco.Params.Get method.
    _, params, _ := router.Lookup("/user/alice/1")
    name := params.Get("name")
    if name != "" {
        fmt.Printf("Hello %s.\n", name) // prints "Hello alice.".
    }

    // 2. Find by loop.
    for _, param := range params {
        if param.Name == "name" {
            fmt.Printf("Hello %s.\n", name) // prints "Hello alice.".
        }
    }
}
```

## URL patterns

Denco's route matching strategy is "most nearly matching".

When routes `/:name` and `/alice` have been built, URI `/alice` matches the route `/alice`, not `/:name`.
Because URI `/alice` is more match with the route `/alice` than `/:name`.

For more example, when routes below have been built:

```
/user/alice
/user/:name
/user/:name/:id
/user/alice/:id
/user/:id/bob
```

Routes matching are:

```
/user/alice      => "/user/alice" (no match with "/user/:name")
/user/bob        => "/user/:name"
/user/naoina/1   => "/user/:name/1"
/user/alice/1    => "/user/alice/:id" (no match with "/user/:name/:id")
/user/1/bob      => "/user/:id/bob"   (no match with "/user/:name/:id")
/user/alice/bob  => "/user/alice/:id" (no match with "/user/:name/:id" and "/user/:id/bob")
```

## Limitation

Denco has some limitations below.

* Number of param records (such as `/:name`) must be less than 2^22
* Number of elements of internal slice must be less than 2^22

## Benchmarks

    cd $GOPATH/github.com/go-openapi/runtime/middleware/denco
    go test -bench . -benchmem

## License

Denco is licensed under the MIT License.
# runtime [![Build Status](https://github.com/go-openapi/runtime/actions/workflows/go-test.yml/badge.svg)](https://github.com/go-openapi/runtime/actions?query=workflow%3A"go+test") [![codecov](https://codecov.io/gh/go-openapi/runtime/branch/master/graph/badge.svg)](https://codecov.io/gh/go-openapi/runtime)

[![Slack Status](https://slackin.goswagger.io/badge.svg)](https://slackin.goswagger.io)
[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/go-openapi/runtime/master/LICENSE) 
[![Go Reference](https://pkg.go.dev/badge/github.com/go-openapi/runtime.svg)](https://pkg.go.dev/github.com/go-openapi/runtime)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-openapi/runtime)](https://goreportcard.com/report/github.com/go-openapi/runtime)

# go OpenAPI toolkit runtime

The runtime component for use in code generation or as untyped usage.
# Contributor Covenant Code of Conduct

## Our Pledge

In the interest of fostering an open and welcoming environment, we as
contributors and maintainers pledge to making participation in our project and
our community a harassment-free experience for everyone, regardless of age, body
size, disability, ethnicity, gender identity and expression, level of experience,
nationality, personal appearance, race, religion, or sexual identity and
orientation.

## Our Standards

Examples of behavior that contributes to creating a positive environment
include:

* Using welcoming and inclusive language
* Being respectful of differing viewpoints and experiences
* Gracefully accepting constructive criticism
* Focusing on what is best for the community
* Showing empathy towards other community members

Examples of unacceptable behavior by participants include:

* The use of sexualized language or imagery and unwelcome sexual attention or
advances
* Trolling, insulting/derogatory comments, and personal or political attacks
* Public or private harassment
* Publishing others' private information, such as a physical or electronic
  address, without explicit permission
* Other conduct which could reasonably be considered inappropriate in a
  professional setting

## Our Responsibilities

Project maintainers are responsible for clarifying the standards of acceptable
behavior and are expected to take appropriate and fair corrective action in
response to any instances of unacceptable behavior.

Project maintainers have the right and responsibility to remove, edit, or
reject comments, commits, code, wiki edits, issues, and other contributions
that are not aligned to this Code of Conduct, or to ban temporarily or
permanently any contributor for other behaviors that they deem inappropriate,
threatening, offensive, or harmful.

## Scope

This Code of Conduct applies both within project spaces and in public spaces
when an individual is representing the project or its community. Examples of
representing a project or community include using an official project e-mail
address, posting via an official social media account, or acting as an appointed
representative at an online or offline event. Representation of a project may be
further defined and clarified by project maintainers.

## Enforcement

Instances of abusive, harassing, or otherwise unacceptable behavior may be
reported by contacting the project team at ivan+abuse@flanders.co.nz. All
complaints will be reviewed and investigated and will result in a response that
is deemed necessary and appropriate to the circumstances. The project team is
obligated to maintain confidentiality with regard to the reporter of an incident.
Further details of specific enforcement policies may be posted separately.

Project maintainers who do not follow or enforce the Code of Conduct in good
faith may face temporary or permanent repercussions as determined by other
members of the project's leadership.

## Attribution

This Code of Conduct is adapted from the [Contributor Covenant][homepage], version 1.4,
available at [http://contributor-covenant.org/version/1/4][version]

[homepage]: http://contributor-covenant.org
[version]: http://contributor-covenant.org/version/1/4/
## Contribution Guidelines

### Pull requests are always welcome

We are always thrilled to receive pull requests, and do our best to
process them as fast as possible. Not sure if that typo is worth a pull
request? Do it! We will appreciate it.

If your pull request is not accepted on the first try, don't be
discouraged! If there's a problem with the implementation, hopefully you
received feedback on what to improve.

We're trying very hard to keep go-swagger lean and focused. We don't want it
to do everything for everybody. This means that we might decide against
incorporating a new feature. However, there might be a way to implement
that feature *on top of* go-swagger.


### Conventions

Fork the repo and make changes on your fork in a feature branch:

- If it's a bugfix branch, name it XXX-something where XXX is the number of the
  issue
- If it's a feature branch, create an enhancement issue to announce your
  intentions, and name it XXX-something where XXX is the number of the issue.

Submit unit tests for your changes.  Go has a great test framework built in; use
it! Take a look at existing tests for inspiration. Run the full test suite on
your branch before submitting a pull request.

Update the documentation when creating or modifying features. Test
your documentation changes for clarity, concision, and correctness, as
well as a clean documentation build. See ``docs/README.md`` for more
information on building the docs and how docs get released.

Write clean code. Universally formatted code promotes ease of writing, reading,
and maintenance. Always run `gofmt -s -w file.go` on each changed file before
committing your changes. Most editors have plugins that do this automatically.

Pull requests descriptions should be as clear as possible and include a
reference to all the issues that they address.

Pull requests must not contain commits from other users or branches.

Commit messages must start with a capitalized and short summary (max. 50
chars) written in the imperative, followed by an optional, more detailed
explanatory text which is separated from the summary by an empty line.

Code review comments may be added to your pull request. Discuss, then make the
suggested modifications and push additional commits to your feature branch. Be
sure to post a comment after pushing. The new commits will show up in the pull
request automatically, but the reviewers will not be notified unless you
comment.

Before the pull request is merged, make sure that you squash your commits into
logical units of work using `git rebase -i` and `git push -f`. After every
commit the test suite should be passing. Include documentation changes in the
same commit so that a revert would remove all traces of the feature or fix.

Commits that fix or close an issue should include a reference like `Closes #XXX`
or `Fixes #XXX`, which will automatically close the issue when merged.

### Sign your work

The sign-off is a simple line at the end of the explanation for the
patch, which certifies that you wrote it or otherwise have the right to
pass it on as an open-source patch.  The rules are pretty simple: if you
can certify the below (from
[developercertificate.org](http://developercertificate.org/)):

```
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
660 York Street, Suite 102,
San Francisco, CA 94110 USA

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.


Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

then you just add a line to every git commit message:

    Signed-off-by: Joe Smith <joe@gmail.com>

using your real name (sorry, no pseudonyms or anonymous contributions.)

You can add the sign off when creating the git commit via `git commit -s`.
# Loads OAI specs [![Build Status](https://github.com/go-openapi/loads/actions/workflows/go-test.yml/badge.svg)](https://github.com/go-openapi/loads/actions?query=workflow%3A"go+test") [![codecov](https://codecov.io/gh/go-openapi/loads/branch/master/graph/badge.svg)](https://codecov.io/gh/go-openapi/loads)

[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/go-openapi/loads/master/LICENSE) [![GoDoc](https://godoc.org/github.com/go-openapi/loads?status.svg)](http://godoc.org/github.com/go-openapi/loads)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-openapi/loads)](https://goreportcard.com/report/github.com/go-openapi/loads)

Loading of OAI specification documents from local or remote locations. Supports JSON and YAML documents.
# Contributor Covenant Code of Conduct

## Our Pledge

In the interest of fostering an open and welcoming environment, we as
contributors and maintainers pledge to making participation in our project and
our community a harassment-free experience for everyone, regardless of age, body
size, disability, ethnicity, gender identity and expression, level of experience,
nationality, personal appearance, race, religion, or sexual identity and
orientation.

## Our Standards

Examples of behavior that contributes to creating a positive environment
include:

* Using welcoming and inclusive language
* Being respectful of differing viewpoints and experiences
* Gracefully accepting constructive criticism
* Focusing on what is best for the community
* Showing empathy towards other community members

Examples of unacceptable behavior by participants include:

* The use of sexualized language or imagery and unwelcome sexual attention or
advances
* Trolling, insulting/derogatory comments, and personal or political attacks
* Public or private harassment
* Publishing others' private information, such as a physical or electronic
  address, without explicit permission
* Other conduct which could reasonably be considered inappropriate in a
  professional setting

## Our Responsibilities

Project maintainers are responsible for clarifying the standards of acceptable
behavior and are expected to take appropriate and fair corrective action in
response to any instances of unacceptable behavior.

Project maintainers have the right and responsibility to remove, edit, or
reject comments, commits, code, wiki edits, issues, and other contributions
that are not aligned to this Code of Conduct, or to ban temporarily or
permanently any contributor for other behaviors that they deem inappropriate,
threatening, offensive, or harmful.

## Scope

This Code of Conduct applies both within project spaces and in public spaces
when an individual is representing the project or its community. Examples of
representing a project or community include using an official project e-mail
address, posting via an official social media account, or acting as an appointed
representative at an online or offline event. Representation of a project may be
further defined and clarified by project maintainers.

## Enforcement

Instances of abusive, harassing, or otherwise unacceptable behavior may be
reported by contacting the project team at ivan+abuse@flanders.co.nz. All
complaints will be reviewed and investigated and will result in a response that
is deemed necessary and appropriate to the circumstances. The project team is
obligated to maintain confidentiality with regard to the reporter of an incident.
Further details of specific enforcement policies may be posted separately.

Project maintainers who do not follow or enforce the Code of Conduct in good
faith may face temporary or permanent repercussions as determined by other
members of the project's leadership.

## Attribution

This Code of Conduct is adapted from the [Contributor Covenant][homepage], version 1.4,
available at [http://contributor-covenant.org/version/1/4][version]

[homepage]: http://contributor-covenant.org
[version]: http://contributor-covenant.org/version/1/4/
## Contribution Guidelines

### Pull requests are always welcome

We are always thrilled to receive pull requests, and do our best to
process them as fast as possible. Not sure if that typo is worth a pull
request? Do it! We will appreciate it.

If your pull request is not accepted on the first try, don't be
discouraged! If there's a problem with the implementation, hopefully you
received feedback on what to improve.

We're trying very hard to keep go-swagger lean and focused. We don't want it
to do everything for everybody. This means that we might decide against
incorporating a new feature. However, there might be a way to implement
that feature *on top of* go-swagger.


### Conventions

Fork the repo and make changes on your fork in a feature branch:

- If it's a bugfix branch, name it XXX-something where XXX is the number of the
  issue
- If it's a feature branch, create an enhancement issue to announce your
  intentions, and name it XXX-something where XXX is the number of the issue.

Submit unit tests for your changes.  Go has a great test framework built in; use
it! Take a look at existing tests for inspiration. Run the full test suite on
your branch before submitting a pull request.

Update the documentation when creating or modifying features. Test
your documentation changes for clarity, concision, and correctness, as
well as a clean documentation build. See ``docs/README.md`` for more
information on building the docs and how docs get released.

Write clean code. Universally formatted code promotes ease of writing, reading,
and maintenance. Always run `gofmt -s -w file.go` on each changed file before
committing your changes. Most editors have plugins that do this automatically.

Pull requests descriptions should be as clear as possible and include a
reference to all the issues that they address.

Pull requests must not contain commits from other users or branches.

Commit messages must start with a capitalized and short summary (max. 50
chars) written in the imperative, followed by an optional, more detailed
explanatory text which is separated from the summary by an empty line.

Code review comments may be added to your pull request. Discuss, then make the
suggested modifications and push additional commits to your feature branch. Be
sure to post a comment after pushing. The new commits will show up in the pull
request automatically, but the reviewers will not be notified unless you
comment.

Before the pull request is merged, make sure that you squash your commits into
logical units of work using `git rebase -i` and `git push -f`. After every
commit the test suite should be passing. Include documentation changes in the
same commit so that a revert would remove all traces of the feature or fix.

Commits that fix or close an issue should include a reference like `Closes #XXX`
or `Fixes #XXX`, which will automatically close the issue when merged.

### Sign your work

The sign-off is a simple line at the end of the explanation for the
patch, which certifies that you wrote it or otherwise have the right to
pass it on as an open-source patch.  The rules are pretty simple: if you
can certify the below (from
[developercertificate.org](http://developercertificate.org/)):

```
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
660 York Street, Suite 102,
San Francisco, CA 94110 USA

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.


Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

then you just add a line to every git commit message:

    Signed-off-by: Joe Smith <joe@gmail.com>

using your real name (sorry, no pseudonyms or anonymous contributions.)

You can add the sign off when creating the git commit via `git commit -s`.
# OpenAPI analysis [![Build Status](https://github.com/go-openapi/analysis/actions/workflows/go-test.yml/badge.svg)](https://github.com/go-openapi/analysis/actions?query=workflow%3A"go+test") [![codecov](https://codecov.io/gh/go-openapi/analysis/branch/master/graph/badge.svg)](https://codecov.io/gh/go-openapi/analysis)

[![Slack Status](https://slackin.goswagger.io/badge.svg)](https://slackin.goswagger.io)
[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/go-openapi/analysis/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-openapi/analysis.svg)](https://pkg.go.dev/github.com/go-openapi/analysis)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-openapi/analysis)](https://goreportcard.com/report/github.com/go-openapi/analysis)


A foundational library to analyze an OAI specification document for easier reasoning about the content.

## What's inside?

* An analyzer providing methods to walk the functional content of a specification
* A spec flattener producing a self-contained document bundle, while preserving `$ref`s
* A spec merger ("mixin") to merge several spec documents into a primary spec
* A spec "fixer" ensuring that response descriptions are non empty

[Documentation](https://pkg.go.dev/github.com/go-openapi/analysis)

## FAQ

* Does this library support OpenAPI 3?

> No.
> This package currently only supports OpenAPI 2.0 (aka Swagger 2.0).
> There is no plan to make it evolve toward supporting OpenAPI 3.x.
> This [discussion thread](https://github.com/go-openapi/spec/issues/21) relates the full story.
# Contributor Covenant Code of Conduct

## Our Pledge

In the interest of fostering an open and welcoming environment, we as
contributors and maintainers pledge to making participation in our project and
our community a harassment-free experience for everyone, regardless of age, body
size, disability, ethnicity, gender identity and expression, level of experience,
nationality, personal appearance, race, religion, or sexual identity and
orientation.

## Our Standards

Examples of behavior that contributes to creating a positive environment
include:

* Using welcoming and inclusive language
* Being respectful of differing viewpoints and experiences
* Gracefully accepting constructive criticism
* Focusing on what is best for the community
* Showing empathy towards other community members

Examples of unacceptable behavior by participants include:

* The use of sexualized language or imagery and unwelcome sexual attention or
advances
* Trolling, insulting/derogatory comments, and personal or political attacks
* Public or private harassment
* Publishing others' private information, such as a physical or electronic
  address, without explicit permission
* Other conduct which could reasonably be considered inappropriate in a
  professional setting

## Our Responsibilities

Project maintainers are responsible for clarifying the standards of acceptable
behavior and are expected to take appropriate and fair corrective action in
response to any instances of unacceptable behavior.

Project maintainers have the right and responsibility to remove, edit, or
reject comments, commits, code, wiki edits, issues, and other contributions
that are not aligned to this Code of Conduct, or to ban temporarily or
permanently any contributor for other behaviors that they deem inappropriate,
threatening, offensive, or harmful.

## Scope

This Code of Conduct applies both within project spaces and in public spaces
when an individual is representing the project or its community. Examples of
representing a project or community include using an official project e-mail
address, posting via an official social media account, or acting as an appointed
representative at an online or offline event. Representation of a project may be
further defined and clarified by project maintainers.

## Enforcement

Instances of abusive, harassing, or otherwise unacceptable behavior may be
reported by contacting the project team at ivan+abuse@flanders.co.nz. All
complaints will be reviewed and investigated and will result in a response that
is deemed necessary and appropriate to the circumstances. The project team is
obligated to maintain confidentiality with regard to the reporter of an incident.
Further details of specific enforcement policies may be posted separately.

Project maintainers who do not follow or enforce the Code of Conduct in good
faith may face temporary or permanent repercussions as determined by other
members of the project's leadership.

## Attribution

This Code of Conduct is adapted from the [Contributor Covenant][homepage], version 1.4,
available at [http://contributor-covenant.org/version/1/4][version]

[homepage]: http://contributor-covenant.org
[version]: http://contributor-covenant.org/version/1/4/
## Contribution Guidelines

### Pull requests are always welcome

We are always thrilled to receive pull requests, and do our best to
process them as fast as possible. Not sure if that typo is worth a pull
request? Do it! We will appreciate it.

If your pull request is not accepted on the first try, don't be
discouraged! If there's a problem with the implementation, hopefully you
received feedback on what to improve.

We're trying very hard to keep go-swagger lean and focused. We don't want it
to do everything for everybody. This means that we might decide against
incorporating a new feature. However, there might be a way to implement
that feature *on top of* go-swagger.


### Conventions

Fork the repo and make changes on your fork in a feature branch:

- If it's a bugfix branch, name it XXX-something where XXX is the number of the
  issue
- If it's a feature branch, create an enhancement issue to announce your
  intentions, and name it XXX-something where XXX is the number of the issue.

Submit unit tests for your changes.  Go has a great test framework built in; use
it! Take a look at existing tests for inspiration. Run the full test suite on
your branch before submitting a pull request.

Update the documentation when creating or modifying features. Test
your documentation changes for clarity, concision, and correctness, as
well as a clean documentation build. See ``docs/README.md`` for more
information on building the docs and how docs get released.

Write clean code. Universally formatted code promotes ease of writing, reading,
and maintenance. Always run `gofmt -s -w file.go` on each changed file before
committing your changes. Most editors have plugins that do this automatically.

Pull requests descriptions should be as clear as possible and include a
reference to all the issues that they address.

Pull requests must not contain commits from other users or branches.

Commit messages must start with a capitalized and short summary (max. 50
chars) written in the imperative, followed by an optional, more detailed
explanatory text which is separated from the summary by an empty line.

Code review comments may be added to your pull request. Discuss, then make the
suggested modifications and push additional commits to your feature branch. Be
sure to post a comment after pushing. The new commits will show up in the pull
request automatically, but the reviewers will not be notified unless you
comment.

Before the pull request is merged, make sure that you squash your commits into
logical units of work using `git rebase -i` and `git push -f`. After every
commit the test suite should be passing. Include documentation changes in the
same commit so that a revert would remove all traces of the feature or fix.

Commits that fix or close an issue should include a reference like `Closes #XXX`
or `Fixes #XXX`, which will automatically close the issue when merged.

### Sign your work

The sign-off is a simple line at the end of the explanation for the
patch, which certifies that you wrote it or otherwise have the right to
pass it on as an open-source patch.  The rules are pretty simple: if you
can certify the below (from
[developercertificate.org](http://developercertificate.org/)):

```
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
660 York Street, Suite 102,
San Francisco, CA 94110 USA

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.


Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

then you just add a line to every git commit message:

    Signed-off-by: Joe Smith <joe@gmail.com>

using your real name (sorry, no pseudonyms or anonymous contributions.)

You can add the sign off when creating the git commit via `git commit -s`.
# Benchmarking name mangling utilities

```bash
go test -bench XXX -run XXX -benchtime 30s
```

## Benchmarks at b3e7a5386f996177e4808f11acb2aa93a0f660df

```
goos: linux
goarch: amd64
pkg: github.com/go-openapi/swag
cpu: Intel(R) Core(TM) i5-6200U CPU @ 2.30GHz
BenchmarkToXXXName/ToGoName-4         	  862623	     44101 ns/op	   10450 B/op	     732 allocs/op
BenchmarkToXXXName/ToVarName-4        	  853656	     40728 ns/op	   10468 B/op	     734 allocs/op
BenchmarkToXXXName/ToFileName-4       	 1268312	     27813 ns/op	    9785 B/op	     617 allocs/op
BenchmarkToXXXName/ToCommandName-4    	 1276322	     27903 ns/op	    9785 B/op	     617 allocs/op
BenchmarkToXXXName/ToHumanNameLower-4 	  895334	     40354 ns/op	   10472 B/op	     731 allocs/op
BenchmarkToXXXName/ToHumanNameTitle-4 	  882441	     40678 ns/op	   10566 B/op	     749 allocs/op
```

## Benchmarks after PR #79

~ x10 performance improvement and ~ /100 memory allocations.

```
goos: linux
goarch: amd64
pkg: github.com/go-openapi/swag
cpu: Intel(R) Core(TM) i5-6200U CPU @ 2.30GHz
BenchmarkToXXXName/ToGoName-4         	 9595830	      3991 ns/op	      42 B/op	       5 allocs/op
BenchmarkToXXXName/ToVarName-4        	 9194276	      3984 ns/op	      62 B/op	       7 allocs/op
BenchmarkToXXXName/ToFileName-4       	17002711	      2123 ns/op	     147 B/op	       7 allocs/op
BenchmarkToXXXName/ToCommandName-4    	16772926	      2111 ns/op	     147 B/op	       7 allocs/op
BenchmarkToXXXName/ToHumanNameLower-4 	 9788331	      3749 ns/op	      92 B/op	       6 allocs/op
BenchmarkToXXXName/ToHumanNameTitle-4 	 9188260	      3941 ns/op	     104 B/op	       6 allocs/op
```

```
goos: linux
goarch: amd64
pkg: github.com/go-openapi/swag
cpu: AMD Ryzen 7 5800X 8-Core Processor             
BenchmarkToXXXName/ToGoName-16         	18527378	      1972 ns/op	      42 B/op	       5 allocs/op
BenchmarkToXXXName/ToVarName-16        	15552692	      2093 ns/op	      62 B/op	       7 allocs/op
BenchmarkToXXXName/ToFileName-16       	32161176	      1117 ns/op	     147 B/op	       7 allocs/op
BenchmarkToXXXName/ToCommandName-16    	32256634	      1137 ns/op	     147 B/op	       7 allocs/op
BenchmarkToXXXName/ToHumanNameLower-16 	18599661	      1946 ns/op	      92 B/op	       6 allocs/op
BenchmarkToXXXName/ToHumanNameTitle-16 	17581353	      2054 ns/op	     105 B/op	       6 allocs/op
```

## Benchmarks at d7d2d1b895f5b6747afaff312dd2a402e69e818b

go1.24

```
goos: linux
goarch: amd64
pkg: github.com/go-openapi/swag
cpu: AMD Ryzen 7 5800X 8-Core Processor             
BenchmarkToXXXName/ToGoName-16         	19757858	      1881 ns/op	      42 B/op	       5 allocs/op
BenchmarkToXXXName/ToVarName-16        	17494111	      2094 ns/op	      74 B/op	       7 allocs/op
BenchmarkToXXXName/ToFileName-16       	28161226	      1492 ns/op	     158 B/op	       7 allocs/op
BenchmarkToXXXName/ToCommandName-16    	23787333	      1489 ns/op	     158 B/op	       7 allocs/op
BenchmarkToXXXName/ToHumanNameLower-16 	17537257	      2030 ns/op	     103 B/op	       6 allocs/op
BenchmarkToXXXName/ToHumanNameTitle-16 	16977453	      2156 ns/op	     105 B/op	       6 allocs/op
```

## Benchmarks after PR #106

Moving the scope of everything down to a struct allowed to reduce a bit garbage and pooling.

On top of that, ToGoName (and thus ToVarName) have been subject to a minor optimization, removing a few allocations.

Overall timings improve by ~ -10%.

go1.24

```
goos: linux
goarch: amd64
pkg: github.com/go-openapi/swag/mangling
cpu: AMD Ryzen 7 5800X 8-Core Processor             
BenchmarkToXXXName/ToGoName-16         	22496130	      1618 ns/op	      31 B/op	       3 allocs/op
BenchmarkToXXXName/ToVarName-16        	22538068	      1618 ns/op	      33 B/op	       3 allocs/op
BenchmarkToXXXName/ToFileName-16       	27722977	      1236 ns/op	     105 B/op	       6 allocs/op
BenchmarkToXXXName/ToCommandName-16    	27967395	      1258 ns/op	     105 B/op	       6 allocs/op
BenchmarkToXXXName/ToHumanNameLower-16 	18587901	      1917 ns/op	     103 B/op	       6 allocs/op
BenchmarkToXXXName/ToHumanNameTitle-16 	17193208	      2019 ns/op	     108 B/op	       7 allocs/op
```
# Swag [![Build Status](https://github.com/go-openapi/swag/actions/workflows/go-test.yml/badge.svg)](https://github.com/go-openapi/swag/actions?query=workflow%3A"go+test") [![codecov](https://codecov.io/gh/go-openapi/swag/branch/master/graph/badge.svg)](https://codecov.io/gh/go-openapi/swag)

[![Slack Status](https://slackin.goswagger.io/badge.svg)](https://slackin.goswagger.io)
[![license](https://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/go-openapi/swag/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-openapi/swag.svg)](https://pkg.go.dev/github.com/go-openapi/swag)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-openapi/swag)](https://goreportcard.com/report/github.com/go-openapi/swag)

Package `swag` contains a bunch of helper functions for go-openapi and go-swagger projects.

You may also use it standalone for your projects.

> `swag` is one of the foundational building blocks of the go-openapi initiative.
>
> Most repositories in `github.com/go-openapi/...` depend on it in some way.
> So does the CLI tool `github.com/go-swagger/go-swagger`,
> and the code generated by this tool.

## Contents

`go-openapi/swag` now exposes a collection of relatively independent modules.

Here is what is inside:

*	Module `cmdutils`

  * [x] utilities to work with CLIs

*	Module `conv`

  * [x] convert between values and pointers for any types
  * [x] convert from string to builtin types (wraps `strconv`)
  * [x] require `./typeutils` (test dependency)

*	Module `fileutils`

  * [x] file upload type
  * [x] search in path (deprecated)

* Module `jsonname`

  * [x] infer JSON names from go properties

* Module `jsonutils`

  * [x] fast json concatenation
  * [x] read and write JSON from and to dynamic go data structures
  * [x] require `github.com/mailru/easyjson`

* Module `loading`

  * [x] load from file or http
  * [x] require `./yamlutils`

* Module `mangling`

  * [x] name mangling for go

* Module `netutils`

  * [x] host, port from address

*	Module `stringutils`

  * [x] search in slice (with case-insensitive)
  * [x] split/join query parameters as arrays

*	Module `typeutils`

  * [x] check the zero value for any type

* Module `yamlutils`

  * [x] converting YAML to JSON
  * [x] loading YAML into a dynamic YAML document
  * [x] require `./jsonutils`
  * [x] require `github.com/mailru/easyjson`
  * [x] require `gopkg.in/yaml.v3`

---

The root module `github.com/go-openapi/swag` at the repo level maintains a few 
dependencies outside of the standard library:

* YAML utilities depend on `gopkg.in/yaml.v3`
* JSON utilities `github.com/mailru/easyjson`

This is not necessarily the case for all sub-modules.

## Release notes

### v0.24.0 [Unreleased]

With this release, we have largely modernized the API of `swag`:

* The traditional `swag` API is still supported: code that imports `swag` will still
  compile and work the same.
* A deprecation notice is published to encourage consumers of this library to adopt
  the newer API
* **Deprecation notice** 
  * configuration through global variables is now deprecated, in favor of options passed as parameters
  * all helper functions are moved to more specialized packages, which are exposed as 
    go modules. Importing such a module would reduce the footprint of dependencies.
  * _all_ functions, variables, constants exposed by the deprecated API have now moved, so
    that consumers of the new API no longer need to import github.com/go-openapi/swag, but
    should import the desired sub-module(s).

**New with this release**:

* [x] type converters and pointer to value helpers now support generic types
* [x] name mangling now support pluralized initialisms (issue #46)
      Strings like "contact IDs" are now recognized as such a plural form and mangled as a linter would expect.
* [x] performance: small improvements to reduce the overhead of convert/format wrappers (see issues #110, or PR #108)
* [x] performance: name mangling utilities run ~ 10% faster (PR #115)

---

Moving forward, no additional feature will be added to the `swag` API directly.

However, child modules will continue to evolve or some new ones may be added in the future.


#### Note to contributors

The mono-repo structure comes with some unavoidable extra pains...

* Testing

> The usual `go test ./...` command, run from the root of this repo won't work any longer to test all submodules.
>
> Each module constitutes an independant unit of test. So you have to run `go test` inside each module.
> Or you may take a look at how this is achieved by CI
> [here] https://github.com/go-openapi/swag/blob/master/.github/workflows/go-test.yml).
>
> There are also some alternative tricks using `go work`, for local development, if you feel comfortable with
> go workspaces. Perhaps some day, we'll have a `go work test` to run all tests without any hack.

* Releasing

> Each module follows its own independant module versioning.
>
> So you have tags like `mangling/v0.24.0`, `fileutils/v0.24.0` etc that are used by `go mod` and `go get`
> to refer to the tagged version of each module specifically.
>
> This means we may release patches etc to each module independently.
>
> We'd like to adopt the rule that modules in this repo would only differ by a patch version
> (e.g. `v0.24.5` vs `v0.24.3`), and we'll level all modules whenever a minor version is introduced.
>
> A script in `./hack` is provided to tag all modules in one go at the same level in one go.

## Todos, suggestions and plans

All kinds of contributions are welcome.

A few ideas:

* [ ] Complete the split of dependencies to isolate easyjson from the rest
* [ ] Improve mangling utilities (improve readability, support for capitalized words,
      better word substitution for non-letter symbols...)
* [ ] Move back to this common shared pot a few of the technical features introduced by go-swagger independently
      (e.g. mangle go package names, search package with go modules support, ...)
* [ ] Apply a similar mono-repo approach to go-openapi/strfmt which suffer from similar woes: bloated API,
      imposed dependency to some database driver.

# Contributor Covenant Code of Conduct

## Our Pledge

In the interest of fostering an open and welcoming environment, we as
contributors and maintainers pledge to making participation in our project and
our community a harassment-free experience for everyone, regardless of age, body
size, disability, ethnicity, gender identity and expression, level of experience,
nationality, personal appearance, race, religion, or sexual identity and
orientation.

## Our Standards

Examples of behavior that contributes to creating a positive environment
include:

* Using welcoming and inclusive language
* Being respectful of differing viewpoints and experiences
* Gracefully accepting constructive criticism
* Focusing on what is best for the community
* Showing empathy towards other community members

Examples of unacceptable behavior by participants include:

* The use of sexualized language or imagery and unwelcome sexual attention or
advances
* Trolling, insulting/derogatory comments, and personal or political attacks
* Public or private harassment
* Publishing others' private information, such as a physical or electronic
  address, without explicit permission
* Other conduct which could reasonably be considered inappropriate in a
  professional setting

## Our Responsibilities

Project maintainers are responsible for clarifying the standards of acceptable
behavior and are expected to take appropriate and fair corrective action in
response to any instances of unacceptable behavior.

Project maintainers have the right and responsibility to remove, edit, or
reject comments, commits, code, wiki edits, issues, and other contributions
that are not aligned to this Code of Conduct, or to ban temporarily or
permanently any contributor for other behaviors that they deem inappropriate,
threatening, offensive, or harmful.

## Scope

This Code of Conduct applies both within project spaces and in public spaces
when an individual is representing the project or its community. Examples of
representing a project or community include using an official project e-mail
address, posting via an official social media account, or acting as an appointed
representative at an online or offline event. Representation of a project may be
further defined and clarified by project maintainers.

## Enforcement

Instances of abusive, harassing, or otherwise unacceptable behavior may be
reported by contacting the project team at ivan+abuse@flanders.co.nz. All
complaints will be reviewed and investigated and will result in a response that
is deemed necessary and appropriate to the circumstances. The project team is
obligated to maintain confidentiality with regard to the reporter of an incident.
Further details of specific enforcement policies may be posted separately.

Project maintainers who do not follow or enforce the Code of Conduct in good
faith may face temporary or permanent repercussions as determined by other
members of the project's leadership.

## Attribution

This Code of Conduct is adapted from the [Contributor Covenant][homepage], version 1.4,
available at [http://contributor-covenant.org/version/1/4][version]

[homepage]: http://contributor-covenant.org
[version]: http://contributor-covenant.org/version/1/4/
## Contribution Guidelines

### Pull requests are always welcome

We are always thrilled to receive pull requests, and do our best to
process them as fast as possible. Not sure if that typo is worth a pull
request? Do it! We will appreciate it.

If your pull request is not accepted on the first try, don't be
discouraged! If there's a problem with the implementation, hopefully you
received feedback on what to improve.

We're trying very hard to keep go-swagger lean and focused. We don't want it
to do everything for everybody. This means that we might decide against
incorporating a new feature. However, there might be a way to implement
that feature *on top of* go-swagger.


### Conventions

Fork the repo and make changes on your fork in a feature branch:

- If it's a bugfix branch, name it XXX-something where XXX is the number of the
  issue
- If it's a feature branch, create an enhancement issue to announce your
  intentions, and name it XXX-something where XXX is the number of the issue.

Submit unit tests for your changes.  Go has a great test framework built in; use
it! Take a look at existing tests for inspiration. Run the full test suite on
your branch before submitting a pull request.

Update the documentation when creating or modifying features. Test
your documentation changes for clarity, concision, and correctness, as
well as a clean documentation build. See ``docs/README.md`` for more
information on building the docs and how docs get released.

Write clean code. Universally formatted code promotes ease of writing, reading,
and maintenance. Always run `gofmt -s -w file.go` on each changed file before
committing your changes. Most editors have plugins that do this automatically.

Pull requests descriptions should be as clear as possible and include a
reference to all the issues that they address.

Pull requests must not contain commits from other users or branches.

Commit messages must start with a capitalized and short summary (max. 50
chars) written in the imperative, followed by an optional, more detailed
explanatory text which is separated from the summary by an empty line.

Code review comments may be added to your pull request. Discuss, then make the
suggested modifications and push additional commits to your feature branch. Be
sure to post a comment after pushing. The new commits will show up in the pull
request automatically, but the reviewers will not be notified unless you
comment.

Before the pull request is merged, make sure that you squash your commits into
logical units of work using `git rebase -i` and `git push -f`. After every
commit the test suite should be passing. Include documentation changes in the
same commit so that a revert would remove all traces of the feature or fix.

Commits that fix or close an issue should include a reference like `Closes #XXX`
or `Fixes #XXX`, which will automatically close the issue when merged.

### Sign your work

The sign-off is a simple line at the end of the explanation for the
patch, which certifies that you wrote it or otherwise have the right to
pass it on as an open-source patch.  The rules are pretty simple: if you
can certify the below (from
[developercertificate.org](http://developercertificate.org/)):

```
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
660 York Street, Suite 102,
San Francisco, CA 94110 USA

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.


Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

then you just add a line to every git commit message:

    Signed-off-by: Joe Smith <joe@gmail.com>

using your real name (sorry, no pseudonyms or anonymous contributions.)

You can add the sign off when creating the git commit via `git commit -s`.
# Strfmt [![Build Status](https://github.com/go-openapi/strfmt/actions/workflows/go-test.yml/badge.svg)](https://github.com/go-openapi/strfmt/actions?query=workflow%3A"go+test") [![codecov](https://codecov.io/gh/go-openapi/strfmt/branch/master/graph/badge.svg)](https://codecov.io/gh/go-openapi/strfmt)
[![Slack Status](https://slackin.goswagger.io/badge.svg)](https://slackin.goswagger.io)
[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/go-openapi/strfmt/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/go-openapi/strfmt?status.svg)](http://godoc.org/github.com/go-openapi/strfmt)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-openapi/strfmt)](https://goreportcard.com/report/github.com/go-openapi/strfmt)

This package exposes a registry of data types to support string formats in the go-openapi toolkit.

strfmt represents a well known string format such as credit card or email. The go toolkit for OpenAPI specifications knows how to deal with those.

## Supported data formats
go-openapi/strfmt follows the swagger 2.0 specification with the following formats
defined [here](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types).

It also provides convenient extensions to go-openapi users.

- [x] JSON-schema draft 4 formats
  - date-time
  - email
  - hostname
  - ipv4
  - ipv6
  - uri
- [x] swagger 2.0 format extensions
  - binary
  - byte (e.g. base64 encoded string)
  - date (e.g. "1970-01-01")
  - password
- [x] go-openapi custom format extensions
  - bsonobjectid (BSON objectID)
  - creditcard
  - duration (e.g. "3 weeks", "1ms")
  - hexcolor (e.g. "#FFFFFF")
  - isbn, isbn10, isbn13
  - mac (e.g "01:02:03:04:05:06")
  - rgbcolor (e.g. "rgb(100,100,100)")
  - ssn
  - uuid, uuid3, uuid4, uuid5
  - cidr (e.g. "192.0.2.1/24", "2001:db8:a0b:12f0::1/32")
  - ulid (e.g. "00000PP9HGSBSSDZ1JTEXBJ0PW", [spec](https://github.com/ulid/spec))

> NOTE: as the name stands for, this package is intended to support string formatting only.
> It does not provide validation for numerical values with swagger format extension for JSON types "number" or
> "integer" (e.g. float, double, int32...).

## Type conversion

All types defined here are stringers and may be converted to strings with `.String()`.
Note that most types defined by this package may be converted directly to string like `string(Email{})`.

`Date` and `DateTime` may be converted directly to `time.Time` like `time.Time(Time{})`.
Similarly, you can convert `Duration` to `time.Duration` as in `time.Duration(Duration{})`

## Using pointers

The `conv` subpackage provides helpers to convert the types to and from pointers, just like `go-openapi/swag` does
with primitive types.

## Format types
Types defined in strfmt expose marshaling and validation capabilities.

List of defined types:
- Base64
- CreditCard
- Date
- DateTime
- Duration
- Email
- HexColor
- Hostname
- IPv4
- IPv6
- CIDR
- ISBN
- ISBN10
- ISBN13
- MAC
- ObjectId
- Password
- RGBColor
- SSN
- URI
- UUID
- UUID3
- UUID4
- UUID5
- [ULID](https://github.com/ulid/spec)
# Contributor Covenant Code of Conduct

## Our Pledge

In the interest of fostering an open and welcoming environment, we as
contributors and maintainers pledge to making participation in our project and
our community a harassment-free experience for everyone, regardless of age, body
size, disability, ethnicity, gender identity and expression, level of experience,
nationality, personal appearance, race, religion, or sexual identity and
orientation.

## Our Standards

Examples of behavior that contributes to creating a positive environment
include:

* Using welcoming and inclusive language
* Being respectful of differing viewpoints and experiences
* Gracefully accepting constructive criticism
* Focusing on what is best for the community
* Showing empathy towards other community members

Examples of unacceptable behavior by participants include:

* The use of sexualized language or imagery and unwelcome sexual attention or
advances
* Trolling, insulting/derogatory comments, and personal or political attacks
* Public or private harassment
* Publishing others' private information, such as a physical or electronic
  address, without explicit permission
* Other conduct which could reasonably be considered inappropriate in a
  professional setting

## Our Responsibilities

Project maintainers are responsible for clarifying the standards of acceptable
behavior and are expected to take appropriate and fair corrective action in
response to any instances of unacceptable behavior.

Project maintainers have the right and responsibility to remove, edit, or
reject comments, commits, code, wiki edits, issues, and other contributions
that are not aligned to this Code of Conduct, or to ban temporarily or
permanently any contributor for other behaviors that they deem inappropriate,
threatening, offensive, or harmful.

## Scope

This Code of Conduct applies both within project spaces and in public spaces
when an individual is representing the project or its community. Examples of
representing a project or community include using an official project e-mail
address, posting via an official social media account, or acting as an appointed
representative at an online or offline event. Representation of a project may be
further defined and clarified by project maintainers.

## Enforcement

Instances of abusive, harassing, or otherwise unacceptable behavior may be
reported by contacting the project team at ivan+abuse@flanders.co.nz. All
complaints will be reviewed and investigated and will result in a response that
is deemed necessary and appropriate to the circumstances. The project team is
obligated to maintain confidentiality with regard to the reporter of an incident.
Further details of specific enforcement policies may be posted separately.

Project maintainers who do not follow or enforce the Code of Conduct in good
faith may face temporary or permanent repercussions as determined by other
members of the project's leadership.

## Attribution

This Code of Conduct is adapted from the [Contributor Covenant][homepage], version 1.4,
available at [http://contributor-covenant.org/version/1/4][version]

[homepage]: http://contributor-covenant.org
[version]: http://contributor-covenant.org/version/1/4/
## Contribution Guidelines

### Pull requests are always welcome

We are always thrilled to receive pull requests, and do our best to
process them as fast as possible. Not sure if that typo is worth a pull
request? Do it! We will appreciate it.

If your pull request is not accepted on the first try, don't be
discouraged! If there's a problem with the implementation, hopefully you
received feedback on what to improve.

We're trying very hard to keep go-swagger lean and focused. We don't want it
to do everything for everybody. This means that we might decide against
incorporating a new feature. However, there might be a way to implement
that feature *on top of* go-swagger.


### Conventions

Fork the repo and make changes on your fork in a feature branch:

- If it's a bugfix branch, name it XXX-something where XXX is the number of the
  issue
- If it's a feature branch, create an enhancement issue to announce your
  intentions, and name it XXX-something where XXX is the number of the issue.

Submit unit tests for your changes.  Go has a great test framework built in; use
it! Take a look at existing tests for inspiration. Run the full test suite on
your branch before submitting a pull request.

Update the documentation when creating or modifying features. Test
your documentation changes for clarity, concision, and correctness, as
well as a clean documentation build. See ``docs/README.md`` for more
information on building the docs and how docs get released.

Write clean code. Universally formatted code promotes ease of writing, reading,
and maintenance. Always run `gofmt -s -w file.go` on each changed file before
committing your changes. Most editors have plugins that do this automatically.

Pull requests descriptions should be as clear as possible and include a
reference to all the issues that they address.

Pull requests must not contain commits from other users or branches.

Commit messages must start with a capitalized and short summary (max. 50
chars) written in the imperative, followed by an optional, more detailed
explanatory text which is separated from the summary by an empty line.

Code review comments may be added to your pull request. Discuss, then make the
suggested modifications and push additional commits to your feature branch. Be
sure to post a comment after pushing. The new commits will show up in the pull
request automatically, but the reviewers will not be notified unless you
comment.

Before the pull request is merged, make sure that you squash your commits into
logical units of work using `git rebase -i` and `git push -f`. After every
commit the test suite should be passing. Include documentation changes in the
same commit so that a revert would remove all traces of the feature or fix.

Commits that fix or close an issue should include a reference like `Closes #XXX`
or `Fixes #XXX`, which will automatically close the issue when merged.

### Sign your work

The sign-off is a simple line at the end of the explanation for the
patch, which certifies that you wrote it or otherwise have the right to
pass it on as an open-source patch.  The rules are pretty simple: if you
can certify the below (from
[developercertificate.org](http://developercertificate.org/)):

```
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
660 York Street, Suite 102,
San Francisco, CA 94110 USA

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.


Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

then you just add a line to every git commit message:

    Signed-off-by: Joe Smith <joe@gmail.com>

using your real name (sorry, no pseudonyms or anonymous contributions.)

You can add the sign off when creating the git commit via `git commit -s`.
# blag blah
# blag blah

Ordered list:

 1. xyz     
 2. abc     
 
First list:

 - a
 - b
 - c
 
Second list:

 * A
 * B
 * C

Another Ordered list:

 0. xyz     
 0. abc     
# OpenAPI errors [![Build Status](https://github.com/go-openapi/errors/actions/workflows/go-test.yml/badge.svg)](https://github.com/go-openapi/errors/actions?query=workflow%3A"go+test") [![codecov](https://codecov.io/gh/go-openapi/errors/branch/master/graph/badge.svg)](https://codecov.io/gh/go-openapi/errors)

[![Slack Status](https://slackin.goswagger.io/badge.svg)](https://slackin.goswagger.io)
[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/go-openapi/errors/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-openapi/errors.svg)](https://pkg.go.dev/github.com/go-openapi/errors)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-openapi/errors)](https://goreportcard.com/report/github.com/go-openapi/errors)

Shared errors and error interface used throughout the various libraries found in the go-openapi toolkit.
# Contributor Covenant Code of Conduct

## Our Pledge

In the interest of fostering an open and welcoming environment, we as
contributors and maintainers pledge to making participation in our project and
our community a harassment-free experience for everyone, regardless of age, body
size, disability, ethnicity, gender identity and expression, level of experience,
nationality, personal appearance, race, religion, or sexual identity and
orientation.

## Our Standards

Examples of behavior that contributes to creating a positive environment
include:

* Using welcoming and inclusive language
* Being respectful of differing viewpoints and experiences
* Gracefully accepting constructive criticism
* Focusing on what is best for the community
* Showing empathy towards other community members

Examples of unacceptable behavior by participants include:

* The use of sexualized language or imagery and unwelcome sexual attention or
advances
* Trolling, insulting/derogatory comments, and personal or political attacks
* Public or private harassment
* Publishing others' private information, such as a physical or electronic
  address, without explicit permission
* Other conduct which could reasonably be considered inappropriate in a
  professional setting

## Our Responsibilities

Project maintainers are responsible for clarifying the standards of acceptable
behavior and are expected to take appropriate and fair corrective action in
response to any instances of unacceptable behavior.

Project maintainers have the right and responsibility to remove, edit, or
reject comments, commits, code, wiki edits, issues, and other contributions
that are not aligned to this Code of Conduct, or to ban temporarily or
permanently any contributor for other behaviors that they deem inappropriate,
threatening, offensive, or harmful.

## Scope

This Code of Conduct applies both within project spaces and in public spaces
when an individual is representing the project or its community. Examples of
representing a project or community include using an official project e-mail
address, posting via an official social media account, or acting as an appointed
representative at an online or offline event. Representation of a project may be
further defined and clarified by project maintainers.

## Enforcement

Instances of abusive, harassing, or otherwise unacceptable behavior may be
reported by contacting the project team at ivan+abuse@flanders.co.nz. All
complaints will be reviewed and investigated and will result in a response that
is deemed necessary and appropriate to the circumstances. The project team is
obligated to maintain confidentiality with regard to the reporter of an incident.
Further details of specific enforcement policies may be posted separately.

Project maintainers who do not follow or enforce the Code of Conduct in good
faith may face temporary or permanent repercussions as determined by other
members of the project's leadership.

## Attribution

This Code of Conduct is adapted from the [Contributor Covenant][homepage], version 1.4,
available at [http://contributor-covenant.org/version/1/4][version]

[homepage]: http://contributor-covenant.org
[version]: http://contributor-covenant.org/version/1/4/
## Contribution Guidelines

### Pull requests are always welcome

We are always thrilled to receive pull requests, and do our best to
process them as fast as possible. Not sure if that typo is worth a pull
request? Do it! We will appreciate it.

If your pull request is not accepted on the first try, don't be
discouraged! If there's a problem with the implementation, hopefully you
received feedback on what to improve.

We're trying very hard to keep go-swagger lean and focused. We don't want it
to do everything for everybody. This means that we might decide against
incorporating a new feature. However, there might be a way to implement
that feature *on top of* go-swagger.


### Conventions

Fork the repo and make changes on your fork in a feature branch:

- If it's a bugfix branch, name it XXX-something where XXX is the number of the
  issue
- If it's a feature branch, create an enhancement issue to announce your
  intentions, and name it XXX-something where XXX is the number of the issue.

Submit unit tests for your changes.  Go has a great test framework built in; use
it! Take a look at existing tests for inspiration. Run the full test suite on
your branch before submitting a pull request.

Update the documentation when creating or modifying features. Test
your documentation changes for clarity, concision, and correctness, as
well as a clean documentation build. See ``docs/README.md`` for more
information on building the docs and how docs get released.

Write clean code. Universally formatted code promotes ease of writing, reading,
and maintenance. Always run `gofmt -s -w file.go` on each changed file before
committing your changes. Most editors have plugins that do this automatically.

Pull requests descriptions should be as clear as possible and include a
reference to all the issues that they address.

Pull requests must not contain commits from other users or branches.

Commit messages must start with a capitalized and short summary (max. 50
chars) written in the imperative, followed by an optional, more detailed
explanatory text which is separated from the summary by an empty line.

Code review comments may be added to your pull request. Discuss, then make the
suggested modifications and push additional commits to your feature branch. Be
sure to post a comment after pushing. The new commits will show up in the pull
request automatically, but the reviewers will not be notified unless you
comment.

Before the pull request is merged, make sure that you squash your commits into
logical units of work using `git rebase -i` and `git push -f`. After every
commit the test suite should be passing. Include documentation changes in the
same commit so that a revert would remove all traces of the feature or fix.

Commits that fix or close an issue should include a reference like `Closes #XXX`
or `Fixes #XXX`, which will automatically close the issue when merged.

### Sign your work

The sign-off is a simple line at the end of the explanation for the
patch, which certifies that you wrote it or otherwise have the right to
pass it on as an open-source patch.  The rules are pretty simple: if you
can certify the below (from
[developercertificate.org](http://developercertificate.org/)):

```
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
660 York Street, Suite 102,
San Francisco, CA 94110 USA

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.


Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

then you just add a line to every git commit message:

    Signed-off-by: Joe Smith <joe@gmail.com>

using your real name (sorry, no pseudonyms or anonymous contributions.)

You can add the sign off when creating the git commit via `git commit -s`.
Evaluate:

* markdownlint
  * assessment: go
  * used by opentelemetry/opentelemetry-go
  * packaged as github action
* misspell
* spellcheck
* govulncheck
* [x] godoc-lint:
  * assessment: no go
  * too simplistic: no real value added
  * not integrated into golangci's suite of linters
  * no packaged github action
<h1>dfkjfkfj<h1>


ERGGLGKLK
==========
this is a title
# .github
Default github settings for go-openapi repositories
---
name: '❓ Question'
about: I have a question about this project
title: ''
labels: [question]
assignees: ''

---

**Question**
<!-- A clear and concise question -->

**Additional context**
<!-- Add any other context or screenshots about the feature request here. -->
---
name: '➕ Feature request'
about: Suggest an idea for this project
title: ''
labels: [enhancement]
assignees: ''

---

**Is your feature request related to a problem? Please describe.**
<!-- A clear and concise description of what the problem is. Ex. I'm always frustrated when [...] -->

**Describe the solution you'd like**
<!-- A clear and concise description of what you want to happen. -->

**Describe alternatives you've considered**
<!-- A clear and concise description of any alternative solutions or features you've considered. -->

**Additional context**
<!-- Add any other context or screenshots about the feature request here. -->
---
name: '🐞 Bug report'
about: Create a report to help us make go-openapi a great project
title: ''
labels: [bug]
assignees: ''

---

**Describe the bug**
<!-- A clear and concise description of what the bug is. -->

**Steps to Reproduce**
<!--
Ideally provide a short repro case in go.
-->

<!--
Steps to reproduce the behavior:
1. Go to '...'
2. Run this '....'
3. Check that '....'
4. See error
-->

**Expected behavior**
<!-- A clear and concise description of what you expected to happen. -->

**Actual behavior**
<!-- A clear and concise description of what is the behavior you observe. -->

**Environment**
<!-- Please paste here the output of `go env`. -->

**Additional context**
<!-- Add any other context or references about the problem here. -->
## Change type

Please select: 🆕 New feature or enhancement|🔧 Bug fix'|📃 Documentation update

## Short description
<!-- Please provide a short description of your change -->

## Fixes
<!-- 
Example:
* fixes #123

Avoid cross-repository fixes, e.g.
* fixes go-openapi/spec#123

Prefer instead:
* contributes go-openapi/spec#123

This means will be solved, but when releases and dependencies updates have been carried out
-->

## Full description
<!-- If needed, please add here more details about your implementation etc -->

<!-- Since this is a bug fix, try your best not to mix this change with extra features or potentially breaking changes -->

## Checklist
<!--- Go over all the following points, and put an `x` in all the boxes that apply. -->
<!-- If you don't qualify for all of the below check list items, please mark your PR in a draft status, so it may be discussed or reviewed with lighter requirements. -->

* [ ] I have signed all my commits with my name and email (see [DCO](https://github.com/apps/dco). **This does not require a PGP-signed commit**
* [ ] I have rebased and squashed my work, so only one commit remains
* [ ] I have added tests to cover my changes.
* [ ] I have properly enriched go doc comments in code.
* [ ] I have properly documented any breaking change.
## Change type: 📃 Documentation update

## Short description
<!-- clear and concise description of your change -->

## Fixes
<!-- 
Example:
* fixes #123

Avoid cross-repository fixes, e.g.
* fixes go-openapi/spec#123

Prefer instead:
* contributes go-openapi/spec#123

This means will be solved, but when releases and dependencies updates have been carried out
-->

## Checklist
<!--- Go over all the following points, and put an `x` in all the boxes that apply. -->
<!-- If you don't qualify for all of the below check list items, please mark your PR in a draft status, so it may be discussed or reviewed with lighter requirements. -->

* [ ] I have signed all my commits with my name and email (see [DCO](https://github.com/apps/dco). **This does not require a PGP-signed commit**
* [ ] I have rebased and squashed my work, so only one commit remains
* [ ] I have signed all my commits with my name and email (see DCO: https://github.com/apps/dco). **This does not require a PGP-signed commit**
* [ ] I have rebased and squash my work, so only one commit remains
* [ ] I have only modified comments in code or markdown files
## Change type: 🆕 New feature or enhancement

## Short description
<!-- Please provide a short description of your change -->

## Fixes
<!-- 
Example:
* fixes #123

Avoid cross-repository fixes, e.g.
* fixes go-openapi/spec#123

Prefer instead:
* contributes go-openapi/spec#123

This means will be solved, but when releases and dependencies updates have been carried out
-->

## Checklist
<!--- Go over all the following points, and put an `x` in all the boxes that apply. -->
<!-- If you don't qualify for all of the below check list items, please mark your PR in a draft status, so it may be discussed or reviewed with lighter requirements. -->

* [ ] I have signed all my commits with my name and email (see [DCO](https://github.com/apps/dco). **This does not require a PGP-signed commit**
* [ ] I have rebased and squashed my work, so only one commit remains
* [ ] I have added tests to cover my changes.
* [ ] I have properly enriched go doc comments in code.
* [ ] I have properly documented any breaking change.
## Change type: 🔧 Bug fix

## Short description
<!-- clear and concise description of your change -->

## Fixes
<!-- 
Example:
* fixes #123

Avoid cross-repository fixes, e.g.
* fixes go-openapi/spec#123

Prefer instead:
* contributes go-openapi/spec#123

This means will be solved, but when releases and dependencies updates have been carried out
-->

## Full description
<!-- If needed, please add here more details about your implementation etc -->

<!-- Since this is a bug fix, try your best not to mix this change with extra features or potentially breaking changes -->

## Checklist
<!--- Go over all the following points, and put an `x` in all the boxes that apply. -->
<!-- If you don't qualify for all of the below check list items, please mark your PR in a draft status, so it may be discussed or reviewed with lighter requirements. -->

* [ ] I have signed all my commits with my name and email (see [DCO](https://github.com/apps/dco). **This does not require a PGP-signed commit**
* [ ] I have rebased and squashed my work, so only one commit remains
* [ ] I have added tests to cover my changes.
* [ ] I have made sure that my code remains portable
* [ ] I have properly enriched go doc comments in code.
* [ ] I have properly documented any breaking change.
# OpenAPI initiative random data generator [![Build Status](https://travis-ci.org/go-openapi/stubs.svg?branch=master)](https://travis-ci.org/go-openapi/stubs) [![codecov](https://codecov.io/gh/go-openapi/stubs/branch/master/graph/badge.svg)](https://codecov.io/gh/go-openapi/stubs) [![Slack Status](https://slackin.goswagger.io/badge.svg)](https://slackin.goswagger.io)

[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/go-openapi/stubs/master/LICENSE) [![GoDoc](https://godoc.org/github.com/go-openapi/stubs?status.svg)](http://godoc.org/github.com/go-openapi/stubs) 


A library to generate random data for a swagger specification.
This is a building block for generating stubs for your API as well as tests.
# Contributor Covenant Code of Conduct

## Our Pledge

In the interest of fostering an open and welcoming environment, we as
contributors and maintainers pledge to making participation in our project and
our community a harassment-free experience for everyone, regardless of age, body
size, disability, ethnicity, gender identity and expression, level of experience,
nationality, personal appearance, race, religion, or sexual identity and
orientation.

## Our Standards

Examples of behavior that contributes to creating a positive environment
include:

* Using welcoming and inclusive language
* Being respectful of differing viewpoints and experiences
* Gracefully accepting constructive criticism
* Focusing on what is best for the community
* Showing empathy towards other community members

Examples of unacceptable behavior by participants include:

* The use of sexualized language or imagery and unwelcome sexual attention or
advances
* Trolling, insulting/derogatory comments, and personal or political attacks
* Public or private harassment
* Publishing others' private information, such as a physical or electronic
  address, without explicit permission
* Other conduct which could reasonably be considered inappropriate in a
  professional setting

## Our Responsibilities

Project maintainers are responsible for clarifying the standards of acceptable
behavior and are expected to take appropriate and fair corrective action in
response to any instances of unacceptable behavior.

Project maintainers have the right and responsibility to remove, edit, or
reject comments, commits, code, wiki edits, issues, and other contributions
that are not aligned to this Code of Conduct, or to ban temporarily or
permanently any contributor for other behaviors that they deem inappropriate,
threatening, offensive, or harmful.

## Scope

This Code of Conduct applies both within project spaces and in public spaces
when an individual is representing the project or its community. Examples of
representing a project or community include using an official project e-mail
address, posting via an official social media account, or acting as an appointed
representative at an online or offline event. Representation of a project may be
further defined and clarified by project maintainers.

## Enforcement

Instances of abusive, harassing, or otherwise unacceptable behavior may be
reported by contacting the project team at ivan+abuse@flanders.co.nz. All
complaints will be reviewed and investigated and will result in a response that
is deemed necessary and appropriate to the circumstances. The project team is
obligated to maintain confidentiality with regard to the reporter of an incident.
Further details of specific enforcement policies may be posted separately.

Project maintainers who do not follow or enforce the Code of Conduct in good
faith may face temporary or permanent repercussions as determined by other
members of the project's leadership.

## Attribution

This Code of Conduct is adapted from the [Contributor Covenant][homepage], version 1.4,
available at [http://contributor-covenant.org/version/1/4][version]

[homepage]: http://contributor-covenant.org
[version]: http://contributor-covenant.org/version/1/4/
# Stubs design

The goal of this library is to be able to generate somewhat good looking random data for structures defined in an openapi or jsonschema document.
These can be used to provide an api with stubs so that you can verify valid and invalid request/responses without an actually meaningful implementation.
Some areas where this is useful are collaboration between different teams so that all teams can do work without having to wait on the delivery of the completed implementation of the API.
A second area where this is obviously useful is for doing automated tests and filling up datastructures for use in those tests.


In the openapi 2.0 specification there are 2 families of types that can be used as targets for a random data generator.

1. Path Item related types
   1. Non-Schema Parameters
   2. Response headers
   3. Collection Items on parameters or headers  
2. Schemas

These 2 families warrant slightly different strategies for getting the necessary parameters for a datagenerator.

## Functionality

To generate meaningful random data the library looks at the schema and based on properties for that schema it picks an appropriate random value generator.
This means that when it is generating a string for an object property that represents a first name it should pick a data generator that generates first names.
Similarly should a property be called city it should generate a city name.

It's possible that the validity of a particular property is dependent on the value of another property, for example a creation date is typically never larger than a modified date.

A data generator for a schema should be able to generate valid and invalid data, ideally there is control for specifying the value is invalid for a particular validation. This will help in tests to validate error codes

## How does it work

The main components in this library are a registry of datagenerators so they are addressable by keys, aliases for those keys to aid with inferring which datagenerator to use.
And of course a value generator, which will generate either a valid or an invalid value.

In the openapi API document and in a json schema document there is a vendor extension that can be used to customize the generation process.

### Value generator function

There are 2 types of data generators but they both have the same function signature.
There are value generators and there are composite value generators.

* A value generator generates a single value for a simple type.
* A composite generator generates a value that is built out of one or more generators eg. object schemas with properties.

The signature for the value generator is:

```go
type ValueGenerator func(GeneratorOpts) (interface{}, error)
```

A value generator is the innermost component in the library and the arguments to the function are used to configure the generator.
The generation process for a generator is configured through a GeneratorOpts interface.

### GeneratorOpts interface

The generator options describe the type and potentially the format for the value that needs to be generated.
In addition to the type information it also captures the field name or definition name of the value that needs to be generated.

### StubMode

The stubmode bitmask allows for configuring which validations should fail for a given value generator.

### Generator

The generator is the main entry point for the library and its Generate method is what will generate the random value for the descriptor.
# Use cases for stubs

The stubs package is intended to support the following use cases:

1. Generate unit test cases for code generated by go-swagger, especially validation code
2. Generate test cases for API servers and clients
3. Generate realistic examples

## Supported swagger 2.0 constructs

- parameter 
- header 
- response
- schema, including support for:
  - [] AllOf 
  - [] Not
  - [] additionalProperties
  - [] patternProperties 
  - [] additionalItems

Formated strings and numbers support go-openapi/strfmt formats, including:
  - [] binary (strfmt.Base64)
  - [] bsonobjectid

The "file" datatype is not supported at the moment.

The generation assumes the provide swagger spec is valid. 

Warnings regarding test case (2): 
- stubs is not concerned about security definitions and credentials
- stubs generates values independently for each parameter, or response: it does not know about consistent sets of values beyond validation constraints, let alone 
  proper chaining of request with responses.

## Sample applications for stubs

- process a spec to include examples 
- generate fixtures for codegen unit testing
- generate fixtures for API server benchmarking
- ...

# Reproducible stubs 

Generated stubs may be regenerated using the same initial seed. 

There are two seeding modes to run the stubs generator: 

- AutoSeed: reseeds itself based on current timestamp 
- SetSeed: generation occurs with a single seed, provided by the caller

The package uses a default seeder to carry on this task. Alternate seeders might be provided.

This also allows for unit testing this very package...

# Setting generation goals 

Goals select the general use case and allow for fine tuning of the sample generation.

We use GenerateMode and StubMode flags to define goals.

Stubmode is used to tune individually every single generated value.

GenerateMode sets general directives and may define StubMode autonomously according to the generation goal.

## GenerateMode

target:
- unitTest: any value value may be produced, not necessarily a realistic definevalue. 
  By default in the UnitTest mode, non-formated strings are loren ipsums, numbers, dates and times may be very small or very large
- examples: the generator attempts to figure out realistic values to be used as examples

GenerateMode accepts arguments to further specify the expected generation: 

args:
- valid:    all generated values are valid. This is the default.
  - count: number of samples to be generated. May be zero to get only default, examples, edge cases or validation checks. Default is 1.
- invalid:  all generated values are invalid, with failures chosen at random.
  - count: number of samples to be generated. Default is 1.
- skipTilting: tilting rules are omitted (only applies to invalid values). Default is true in unit test mode and false in example mode.
- skipFuzzying: fuzzying rules are omitted
- withTilting: add tilted invalid values, with failures chosen at random. This forces a tilting strategy on a new set of values.
  - count: number of tilted samples to be generated. Default is 1.
- withValidationCheck: for each validation, an invalid value against this validation is generated (e.g. if min and max are specified, we have one value below min and one value above max)
   - count: this sets a limit on the count of generated values. If you have n validations, that makes n generated values. With the limit set, validations are checked width-first (outermost validations come first)
- withAllValidationChecks: all combinations of failed validations are generated. This includes one valid value and single failure from withValidation.
   - count: this sets a limit on the count of generated values. If you have n validations, that makes 2^n generated values, which may raise quickly when nesting schema structures. With the limit set, validations are checked width-first (outermost validations come first)
- withDefault: add default value to sample values (not included in "count"). This simulates a user providing exactly the same value as the default. Since defaults may be set at different levels of the structure, a new value is generated for each applicable default.
- withExample: add example value to sample values (not included in "count"), whenever applicable. Examples are assumed valid, but may be invalid.
   - tiltDefault: constructs invalid values by tilting default
   - tiltExample: constructs invalid values by tilting examples  
   - count: number of title
- withEdgeCase:
    - standard: standard preconfigured edge cases (e.g. boundary value for min constraint with boundary included, Feb 29th dates for date types, ...)
    - value: a customizable value added as an edge case. May be any value, array or object. Severa values may be specified.
- max: an upper limit on number generation (floats, ints, ...). By default, the full range of the data type is used.
- min: a lower limit on number generation.
- precision: rounding specification for floats
- multipleOf: a multiple specification for numbers
- length:
  - on string values: an upper limit on string size, to avoid long texts being unwittingly generated
  - on array values: an upper limit on the number of generated items
  - on additionalProperties: an upper limit on the number of generated additionalProperties
  - on patternProperties: an upper limit on the number of generated patternProperties
- words: a limit on the number of words in generated texts, paragraphs and sentences
- lang: specifies the target language(s) for strings and text. 
  - [lang]
  - all: special value to tell the generator to pick text at random from all supported languages

This sets an overall target for all generations. For every single spec object, you may hint further and override global settings on a case by case basis.

# Hinting

The `x-datagen` extension may be used in the swagger spec to hint the stubs generator.

Examples:
``` 
parameters:
    phone:
        in: query 
        required: true
        type: string
        x-datagen:
          name: mobile
```
will generate a mobile phone number for query parameter "phone".

``` 
parameters:
    price:
        in: query 
        required: true
        type: string
        x-datagen:
          name: float
```
will generate float numbers for the price parameter. 
It might be not realistic to have 3.1344344E35 as a price, but it is a valid value...

`x-datagen` accepts args to further hint the generator. 
All Args accepted by GenerateMode or StubMode may be used here to locally override generation options. 

Example:
``` 
parameters:
    vat:
        in: query 
        required: true
        type: number
        x-datagen:
          args:
            min: 100
            max: 10000
            
```
will generate floats between 100 and 10000. 
Again, it might not be realistic to have prices such as 503.14561434566...

Some generators are already configured to suit some targeted usage.

Example: 
``` 
parameters:
    price:
        in: query 
        required: true
        type: string
        x-datagen:
          name: small-amounts
```
and you get more realistic values like 3.23, 104.30...

## x-datagen args

- name: overrides the generator's name

Supported args, not already mentioned in the GenerateMode section:
- args:

# Fuzzying

The generation takes basic decisions based on type and available validation constraints.
Obviously, this does not work well when the goal is to generate realistic samples for your data.
Of course, everyhing may be fine-tuned by using the x-datagen extension.
But this is tedious work on large specs. 

We introduce fuzzying as a way to carry on some of the guesswork when deciding which kind of fake value would best match expectations. 
Fuzzying relies on names, titles and description to look for recognizable keywords, then elects an adequate random generator.

Examples for the above samples:
- with name "price", the fuzzying guesses that it is not some arbitrary float to be generated, but a decimal between 0 and 1,000,000 (max overrides this limit).
- with a description such as "The item price expressed in USD. Items are sold at a low price", fuzzying rules should guess that we want to generate small-amount values (so between 0 and 1,000).

Fuzzying also works for non-strings data:
- a date field named "birthday" would hint toward the birthdate value generator (hence dates are not in the future and only in a recent past)
- a datetime named "timestamp" or described as "The update date" would generate a recent datetime value. Further, edge cases with this guess would add 01/01/1970 to the generated value

# Tilting

Tilting is a strategy to generate invalid values, by tilting a valid value just enough so it becomes invalid.
This generates interesting test cases since values are supposed to be "almost" valid but actually are not.

Tilting strategies depend on validations. 

Examples:
- tilting on max will generate a value at or slightly above the lowest invalid value 
- tilting on regexp will change the smallest number of characters in a string so the regexp no more matches
- tilting on enum will slightly modify a valid value and make it invalid
- ...

Options skipTilting and withTilting tune the generator behavior regarding tilting.
# gojsonpointer [![Build Status](https://github.com/go-openapi/jsonpointer/actions/workflows/go-test.yml/badge.svg)](https://github.com/go-openapi/jsonpointer/actions?query=workflow%3A"go+test") [![codecov](https://codecov.io/gh/go-openapi/jsonpointer/branch/master/graph/badge.svg)](https://codecov.io/gh/go-openapi/jsonpointer)

[![Slack Status](https://slackin.goswagger.io/badge.svg)](https://slackin.goswagger.io)
[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/go-openapi/jsonpointer/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-openapi/jsonpointer.svg)](https://pkg.go.dev/github.com/go-openapi/jsonpointer)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-openapi/jsonpointer)](https://goreportcard.com/report/github.com/go-openapi/jsonpointer)

An implementation of JSON Pointer - Go language

## Status
Completed YES

Tested YES

## References
http://tools.ietf.org/html/draft-ietf-appsawg-json-pointer-07

### Note
The 4.Evaluation part of the previous reference, starting with 'If the currently referenced value is a JSON array, the reference token MUST contain either...' is not implemented.
# Contributor Covenant Code of Conduct

## Our Pledge

In the interest of fostering an open and welcoming environment, we as
contributors and maintainers pledge to making participation in our project and
our community a harassment-free experience for everyone, regardless of age, body
size, disability, ethnicity, gender identity and expression, level of experience,
nationality, personal appearance, race, religion, or sexual identity and
orientation.

## Our Standards

Examples of behavior that contributes to creating a positive environment
include:

* Using welcoming and inclusive language
* Being respectful of differing viewpoints and experiences
* Gracefully accepting constructive criticism
* Focusing on what is best for the community
* Showing empathy towards other community members

Examples of unacceptable behavior by participants include:

* The use of sexualized language or imagery and unwelcome sexual attention or
advances
* Trolling, insulting/derogatory comments, and personal or political attacks
* Public or private harassment
* Publishing others' private information, such as a physical or electronic
  address, without explicit permission
* Other conduct which could reasonably be considered inappropriate in a
  professional setting

## Our Responsibilities

Project maintainers are responsible for clarifying the standards of acceptable
behavior and are expected to take appropriate and fair corrective action in
response to any instances of unacceptable behavior.

Project maintainers have the right and responsibility to remove, edit, or
reject comments, commits, code, wiki edits, issues, and other contributions
that are not aligned to this Code of Conduct, or to ban temporarily or
permanently any contributor for other behaviors that they deem inappropriate,
threatening, offensive, or harmful.

## Scope

This Code of Conduct applies both within project spaces and in public spaces
when an individual is representing the project or its community. Examples of
representing a project or community include using an official project e-mail
address, posting via an official social media account, or acting as an appointed
representative at an online or offline event. Representation of a project may be
further defined and clarified by project maintainers.

## Enforcement

Instances of abusive, harassing, or otherwise unacceptable behavior may be
reported by contacting the project team at ivan+abuse@flanders.co.nz. All
complaints will be reviewed and investigated and will result in a response that
is deemed necessary and appropriate to the circumstances. The project team is
obligated to maintain confidentiality with regard to the reporter of an incident.
Further details of specific enforcement policies may be posted separately.

Project maintainers who do not follow or enforce the Code of Conduct in good
faith may face temporary or permanent repercussions as determined by other
members of the project's leadership.

## Attribution

This Code of Conduct is adapted from the [Contributor Covenant][homepage], version 1.4,
available at [http://contributor-covenant.org/version/1/4][version]

[homepage]: http://contributor-covenant.org
[version]: http://contributor-covenant.org/version/1/4/
## Contribution Guidelines

### Pull requests are always welcome

We are always thrilled to receive pull requests, and do our best to
process them as fast as possible. Not sure if that typo is worth a pull
request? Do it! We will appreciate it.

If your pull request is not accepted on the first try, don't be
discouraged! If there's a problem with the implementation, hopefully you
received feedback on what to improve.

We're trying very hard to keep go-swagger lean and focused. We don't want it
to do everything for everybody. This means that we might decide against
incorporating a new feature. However, there might be a way to implement
that feature *on top of* go-swagger.


### Conventions

Fork the repo and make changes on your fork in a feature branch:

- If it's a bugfix branch, name it XXX-something where XXX is the number of the
  issue
- If it's a feature branch, create an enhancement issue to announce your
  intentions, and name it XXX-something where XXX is the number of the issue.

Submit unit tests for your changes.  Go has a great test framework built in; use
it! Take a look at existing tests for inspiration. Run the full test suite on
your branch before submitting a pull request.

Update the documentation when creating or modifying features. Test
your documentation changes for clarity, concision, and correctness, as
well as a clean documentation build. See ``docs/README.md`` for more
information on building the docs and how docs get released.

Write clean code. Universally formatted code promotes ease of writing, reading,
and maintenance. Always run `gofmt -s -w file.go` on each changed file before
committing your changes. Most editors have plugins that do this automatically.

Pull requests descriptions should be as clear as possible and include a
reference to all the issues that they address.

Pull requests must not contain commits from other users or branches.

Commit messages must start with a capitalized and short summary (max. 50
chars) written in the imperative, followed by an optional, more detailed
explanatory text which is separated from the summary by an empty line.

Code review comments may be added to your pull request. Discuss, then make the
suggested modifications and push additional commits to your feature branch. Be
sure to post a comment after pushing. The new commits will show up in the pull
request automatically, but the reviewers will not be notified unless you
comment.

Before the pull request is merged, make sure that you squash your commits into
logical units of work using `git rebase -i` and `git push -f`. After every
commit the test suite should be passing. Include documentation changes in the
same commit so that a revert would remove all traces of the feature or fix.

Commits that fix or close an issue should include a reference like `Closes #XXX`
or `Fixes #XXX`, which will automatically close the issue when merged.

### Sign your work

The sign-off is a simple line at the end of the explanation for the
patch, which certifies that you wrote it or otherwise have the right to
pass it on as an open-source patch.  The rules are pretty simple: if you
can certify the below (from
[developercertificate.org](http://developercertificate.org/)):

```
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
660 York Street, Suite 102,
San Francisco, CA 94110 USA

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.


Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

then you just add a line to every git commit message:

    Signed-off-by: Joe Smith <joe@gmail.com>

using your real name (sorry, no pseudonyms or anonymous contributions.)

You can add the sign off when creating the git commit via `git commit -s`.
# Swagger 2.0 specification schema

This folder contains the Swagger 2.0 specification schema files maintained here:

https://github.com/reverb/swagger-spec/blob/master/schemas/v2.0# OpenAPI v2 object model [![Build Status](https://github.com/go-openapi/spec/actions/workflows/go-test.yml/badge.svg)](https://github.com/go-openapi/spec/actions?query=workflow%3A"go+test") [![codecov](https://codecov.io/gh/go-openapi/spec/branch/master/graph/badge.svg)](https://codecov.io/gh/go-openapi/spec)

[![Slack Status](https://slackin.goswagger.io/badge.svg)](https://slackin.goswagger.io)
[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/go-openapi/spec/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-openapi/spec.svg)](https://pkg.go.dev/github.com/go-openapi/spec)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-openapi/spec)](https://goreportcard.com/report/github.com/go-openapi/spec)

The object model for OpenAPI specification documents.

### FAQ

* What does this do?

> 1. This package knows how to marshal and unmarshal Swagger API specifications into a golang object model
> 2. It knows how to resolve $ref and expand them to make a single root document

* How does it play with the rest of the go-openapi packages ?

> 1. This package is at the core of the go-openapi suite of packages and [code generator](https://github.com/go-swagger/go-swagger)
> 2. There is a [spec loading package](https://github.com/go-openapi/loads) to fetch specs as JSON or YAML from local or remote locations
> 3. There is a [spec validation package](https://github.com/go-openapi/validate) built on top of it
> 4. There is a [spec analysis package](https://github.com/go-openapi/analysis) built on top of it, to analyze, flatten, fix and merge spec documents

* Does this library support OpenAPI 3?

> No.
> This package currently only supports OpenAPI 2.0 (aka Swagger 2.0).
> There is no plan to make it evolve toward supporting OpenAPI 3.x.
> This [discussion thread](https://github.com/go-openapi/spec/issues/21) relates the full story.
>
> An early attempt to support Swagger 3 may be found at: https://github.com/go-openapi/spec3

* Does the unmarshaling support YAML?

> Not directly. The exposed types know only how to unmarshal from JSON.
>
> In order to load a YAML document as a Swagger spec, you need to use the loaders provided by
> github.com/go-openapi/loads
>
> Take a look at the example there: https://pkg.go.dev/github.com/go-openapi/loads#example-Spec
>
> See also https://github.com/go-openapi/spec/issues/164

* How can I validate a spec?

> Validation is provided by [the validate package](http://github.com/go-openapi/validate)

* Why do we have an `ID` field for `Schema` which is not part of the swagger spec?

> We found jsonschema compatibility more important: since `id` in jsonschema influences
> how `$ref` are resolved.
> This `id` does not conflict with any property named `id`.
>
> See also https://github.com/go-openapi/spec/issues/23
# Contributor Covenant Code of Conduct

## Our Pledge

In the interest of fostering an open and welcoming environment, we as
contributors and maintainers pledge to making participation in our project and
our community a harassment-free experience for everyone, regardless of age, body
size, disability, ethnicity, gender identity and expression, level of experience,
nationality, personal appearance, race, religion, or sexual identity and
orientation.

## Our Standards

Examples of behavior that contributes to creating a positive environment
include:

* Using welcoming and inclusive language
* Being respectful of differing viewpoints and experiences
* Gracefully accepting constructive criticism
* Focusing on what is best for the community
* Showing empathy towards other community members

Examples of unacceptable behavior by participants include:

* The use of sexualized language or imagery and unwelcome sexual attention or
advances
* Trolling, insulting/derogatory comments, and personal or political attacks
* Public or private harassment
* Publishing others' private information, such as a physical or electronic
  address, without explicit permission
* Other conduct which could reasonably be considered inappropriate in a
  professional setting

## Our Responsibilities

Project maintainers are responsible for clarifying the standards of acceptable
behavior and are expected to take appropriate and fair corrective action in
response to any instances of unacceptable behavior.

Project maintainers have the right and responsibility to remove, edit, or
reject comments, commits, code, wiki edits, issues, and other contributions
that are not aligned to this Code of Conduct, or to ban temporarily or
permanently any contributor for other behaviors that they deem inappropriate,
threatening, offensive, or harmful.

## Scope

This Code of Conduct applies both within project spaces and in public spaces
when an individual is representing the project or its community. Examples of
representing a project or community include using an official project e-mail
address, posting via an official social media account, or acting as an appointed
representative at an online or offline event. Representation of a project may be
further defined and clarified by project maintainers.

## Enforcement

Instances of abusive, harassing, or otherwise unacceptable behavior may be
reported by contacting the project team at ivan+abuse@flanders.co.nz. All
complaints will be reviewed and investigated and will result in a response that
is deemed necessary and appropriate to the circumstances. The project team is
obligated to maintain confidentiality with regard to the reporter of an incident.
Further details of specific enforcement policies may be posted separately.

Project maintainers who do not follow or enforce the Code of Conduct in good
faith may face temporary or permanent repercussions as determined by other
members of the project's leadership.

## Attribution

This Code of Conduct is adapted from the [Contributor Covenant][homepage], version 1.4,
available at [http://contributor-covenant.org/version/1/4][version]

[homepage]: http://contributor-covenant.org
[version]: http://contributor-covenant.org/version/1/4/
## Contribution Guidelines

### Pull requests are always welcome

We are always thrilled to receive pull requests, and do our best to
process them as fast as possible. Not sure if that typo is worth a pull
request? Do it! We will appreciate it.

If your pull request is not accepted on the first try, don't be
discouraged! If there's a problem with the implementation, hopefully you
received feedback on what to improve.

We're trying very hard to keep go-swagger lean and focused. We don't want it
to do everything for everybody. This means that we might decide against
incorporating a new feature. However, there might be a way to implement
that feature *on top of* go-swagger.


### Conventions

Fork the repo and make changes on your fork in a feature branch:

- If it's a bugfix branch, name it XXX-something where XXX is the number of the
  issue
- If it's a feature branch, create an enhancement issue to announce your
  intentions, and name it XXX-something where XXX is the number of the issue.

Submit unit tests for your changes.  Go has a great test framework built in; use
it! Take a look at existing tests for inspiration. Run the full test suite on
your branch before submitting a pull request.

Update the documentation when creating or modifying features. Test
your documentation changes for clarity, concision, and correctness, as
well as a clean documentation build. See ``docs/README.md`` for more
information on building the docs and how docs get released.

Write clean code. Universally formatted code promotes ease of writing, reading,
and maintenance. Always run `gofmt -s -w file.go` on each changed file before
committing your changes. Most editors have plugins that do this automatically.

Pull requests descriptions should be as clear as possible and include a
reference to all the issues that they address.

Pull requests must not contain commits from other users or branches.

Commit messages must start with a capitalized and short summary (max. 50
chars) written in the imperative, followed by an optional, more detailed
explanatory text which is separated from the summary by an empty line.

Code review comments may be added to your pull request. Discuss, then make the
suggested modifications and push additional commits to your feature branch. Be
sure to post a comment after pushing. The new commits will show up in the pull
request automatically, but the reviewers will not be notified unless you
comment.

Before the pull request is merged, make sure that you squash your commits into
logical units of work using `git rebase -i` and `git push -f`. After every
commit the test suite should be passing. Include documentation changes in the
same commit so that a revert would remove all traces of the feature or fix.

Commits that fix or close an issue should include a reference like `Closes #XXX`
or `Fixes #XXX`, which will automatically close the issue when merged.

### Sign your work

The sign-off is a simple line at the end of the explanation for the
patch, which certifies that you wrote it or otherwise have the right to
pass it on as an open-source patch.  The rules are pretty simple: if you
can certify the below (from
[developercertificate.org](http://developercertificate.org/)):

```
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
660 York Street, Suite 102,
San Francisco, CA 94110 USA

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.


Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

then you just add a line to every git commit message:

    Signed-off-by: Joe Smith <joe@gmail.com>

using your real name (sorry, no pseudonyms or anonymous contributions.)

You can add the sign off when creating the git commit via `git commit -s`.
# gojsonreference [![Build Status](https://github.com/go-openapi/jsonreference/actions/workflows/go-test.yml/badge.svg)](https://github.com/go-openapi/jsonreference/actions?query=workflow%3A"go+test") [![codecov](https://codecov.io/gh/go-openapi/jsonreference/branch/master/graph/badge.svg)](https://codecov.io/gh/go-openapi/jsonreference)

[![Slack Status](https://slackin.goswagger.io/badge.svg)](https://slackin.goswagger.io)
[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/go-openapi/jsonreference/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-openapi/jsonreference.svg)](https://pkg.go.dev/github.com/go-openapi/jsonreference)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-openapi/jsonreference)](https://goreportcard.com/report/github.com/go-openapi/jsonreference)

An implementation of JSON Reference - Go language

## Status
Feature complete. Stable API

## Dependencies
* https://github.com/go-openapi/jsonpointer

## References

* http://tools.ietf.org/html/draft-ietf-appsawg-json-pointer-07
* http://tools.ietf.org/html/draft-pbryan-zyp-json-ref-03
# Contributor Covenant Code of Conduct

## Our Pledge

In the interest of fostering an open and welcoming environment, we as
contributors and maintainers pledge to making participation in our project and
our community a harassment-free experience for everyone, regardless of age, body
size, disability, ethnicity, gender identity and expression, level of experience,
nationality, personal appearance, race, religion, or sexual identity and
orientation.

## Our Standards

Examples of behavior that contributes to creating a positive environment
include:

* Using welcoming and inclusive language
* Being respectful of differing viewpoints and experiences
* Gracefully accepting constructive criticism
* Focusing on what is best for the community
* Showing empathy towards other community members

Examples of unacceptable behavior by participants include:

* The use of sexualized language or imagery and unwelcome sexual attention or
advances
* Trolling, insulting/derogatory comments, and personal or political attacks
* Public or private harassment
* Publishing others' private information, such as a physical or electronic
  address, without explicit permission
* Other conduct which could reasonably be considered inappropriate in a
  professional setting

## Our Responsibilities

Project maintainers are responsible for clarifying the standards of acceptable
behavior and are expected to take appropriate and fair corrective action in
response to any instances of unacceptable behavior.

Project maintainers have the right and responsibility to remove, edit, or
reject comments, commits, code, wiki edits, issues, and other contributions
that are not aligned to this Code of Conduct, or to ban temporarily or
permanently any contributor for other behaviors that they deem inappropriate,
threatening, offensive, or harmful.

## Scope

This Code of Conduct applies both within project spaces and in public spaces
when an individual is representing the project or its community. Examples of
representing a project or community include using an official project e-mail
address, posting via an official social media account, or acting as an appointed
representative at an online or offline event. Representation of a project may be
further defined and clarified by project maintainers.

## Enforcement

Instances of abusive, harassing, or otherwise unacceptable behavior may be
reported by contacting the project team at ivan+abuse@flanders.co.nz. All
complaints will be reviewed and investigated and will result in a response that
is deemed necessary and appropriate to the circumstances. The project team is
obligated to maintain confidentiality with regard to the reporter of an incident.
Further details of specific enforcement policies may be posted separately.

Project maintainers who do not follow or enforce the Code of Conduct in good
faith may face temporary or permanent repercussions as determined by other
members of the project's leadership.

## Attribution

This Code of Conduct is adapted from the [Contributor Covenant][homepage], version 1.4,
available at [http://contributor-covenant.org/version/1/4][version]

[homepage]: http://contributor-covenant.org
[version]: http://contributor-covenant.org/version/1/4/
## Contribution Guidelines

### Pull requests are always welcome

We are always thrilled to receive pull requests, and do our best to
process them as fast as possible. Not sure if that typo is worth a pull
request? Do it! We will appreciate it.

If your pull request is not accepted on the first try, don't be
discouraged! If there's a problem with the implementation, hopefully you
received feedback on what to improve.

We're trying very hard to keep go-swagger lean and focused. We don't want it
to do everything for everybody. This means that we might decide against
incorporating a new feature. However, there might be a way to implement
that feature *on top of* go-swagger.


### Conventions

Fork the repo and make changes on your fork in a feature branch:

- If it's a bugfix branch, name it XXX-something where XXX is the number of the
  issue
- If it's a feature branch, create an enhancement issue to announce your
  intentions, and name it XXX-something where XXX is the number of the issue.

Submit unit tests for your changes.  Go has a great test framework built in; use
it! Take a look at existing tests for inspiration. Run the full test suite on
your branch before submitting a pull request.

Update the documentation when creating or modifying features. Test
your documentation changes for clarity, concision, and correctness, as
well as a clean documentation build. See ``docs/README.md`` for more
information on building the docs and how docs get released.

Write clean code. Universally formatted code promotes ease of writing, reading,
and maintenance. Always run `gofmt -s -w file.go` on each changed file before
committing your changes. Most editors have plugins that do this automatically.

Pull requests descriptions should be as clear as possible and include a
reference to all the issues that they address.

Pull requests must not contain commits from other users or branches.

Commit messages must start with a capitalized and short summary (max. 50
chars) written in the imperative, followed by an optional, more detailed
explanatory text which is separated from the summary by an empty line.

Code review comments may be added to your pull request. Discuss, then make the
suggested modifications and push additional commits to your feature branch. Be
sure to post a comment after pushing. The new commits will show up in the pull
request automatically, but the reviewers will not be notified unless you
comment.

Before the pull request is merged, make sure that you squash your commits into
logical units of work using `git rebase -i` and `git push -f`. After every
commit the test suite should be passing. Include documentation changes in the
same commit so that a revert would remove all traces of the feature or fix.

Commits that fix or close an issue should include a reference like `Closes #XXX`
or `Fixes #XXX`, which will automatically close the issue when merged.

### Sign your work

The sign-off is a simple line at the end of the explanation for the
patch, which certifies that you wrote it or otherwise have the right to
pass it on as an open-source patch.  The rules are pretty simple: if you
can certify the below (from
[developercertificate.org](http://developercertificate.org/)):

```
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
660 York Street, Suite 102,
San Francisco, CA 94110 USA

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.


Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

then you just add a line to every git commit message:

    Signed-off-by: Joe Smith <joe@gmail.com>

using your real name (sorry, no pseudonyms or anonymous contributions.)

You can add the sign off when creating the git commit via `git commit -s`.
# doc

Documentation check and generation tools.
# ci-workflows[![Build Status](https://github.com/go-openapi/ci-workflows/actions/workflows/go-test.yml/badge.svg)](https://github.com/go-openapi/ci-workflows/actions?query=workflow%3A"go+test") [![codecov](https://codecov.io/gh/go-openapi/ci-workflows/branch/master/graph/badge.svg)](https://codecov.io/gh/go-openapi/ci-workflows)

[![Slack Status](https://slackin.goswagger.io/badge.svg)](https://slackin.goswagger.io)
[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/go-openapi/ci-workflows/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-openapi/ci-workflows.svg)](https://pkg.go.dev/github.com/go-openapi/ci-workflows)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-openapi/ci-workflows)](https://goreportcard.com/report/github.com/go-openapi/ci-workflows)

Common Continuous Integration (`CI`) workflows and setup for go-openapi repos.

* shared github action workflows
* shared `dependabot` configuration (TODO)
* shared `golangci-lint` configuration (**BLOCKED**)

> NOTE: at this moment, it is difficult to share the golangci-lint config,
> so that one is not shared yet.

## Motivation

It took a while, but we eventually managed to align all checks, tests and
dependabot rules declared in the family of go-openapi repos.

Now we'd like to be able to maintain, enrich and improve these checks without
worrying too much about the burden to replicate it about a dozen times.

## Contemplated enhancements

In no particular order:

* [x] ui: enrich github actions UI with a job summary
* [x] doc: add markdown linting for docs
* [x] doc: add spellcheck for docs (and code?)
* [ ] version common workflows, so we can limit the impact of a change
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
# Contributor Covenant Code of Conduct

## Our Pledge

In the interest of fostering an open and welcoming environment, we as
contributors and maintainers pledge to making participation in our project and
our community a harassment-free experience for everyone, regardless of age, body
size, disability, ethnicity, gender identity and expression, level of experience,
nationality, personal appearance, race, religion, or sexual identity and
orientation.

## Our Standards

Examples of behavior that contributes to creating a positive environment
include:

* Using welcoming and inclusive language
* Being respectful of differing viewpoints and experiences
* Gracefully accepting constructive criticism
* Focusing on what is best for the community
* Showing empathy towards other community members

Examples of unacceptable behavior by participants include:

* The use of sexualized language or imagery and unwelcome sexual attention or
advances
* Trolling, insulting/derogatory comments, and personal or political attacks
* Public or private harassment
* Publishing others' private information, such as a physical or electronic
  address, without explicit permission
* Other conduct which could reasonably be considered inappropriate in a
  professional setting

## Our Responsibilities

Project maintainers are responsible for clarifying the standards of acceptable
behavior and are expected to take appropriate and fair corrective action in
response to any instances of unacceptable behavior.

Project maintainers have the right and responsibility to remove, edit, or
reject comments, commits, code, wiki edits, issues, and other contributions
that are not aligned to this Code of Conduct, or to ban temporarily or
permanently any contributor for other behaviors that they deem inappropriate,
threatening, offensive, or harmful.

## Scope

This Code of Conduct applies both within project spaces and in public spaces
when an individual is representing the project or its community. Examples of
representing a project or community include using an official project e-mail
address, posting via an official social media account, or acting as an appointed
representative at an online or offline event. Representation of a project may be
further defined and clarified by project maintainers.

## Enforcement

Instances of abusive, harassing, or otherwise unacceptable behavior may be
reported by contacting the project team at ivan+abuse@flanders.co.nz. All
complaints will be reviewed and investigated and will result in a response that
is deemed necessary and appropriate to the circumstances. The project team is
obligated to maintain confidentiality with regard to the reporter of an incident.
Further details of specific enforcement policies may be posted separately.

Project maintainers who do not follow or enforce the Code of Conduct in good
faith may face temporary or permanent repercussions as determined by other
members of the project's leadership.

## Attribution

This Code of Conduct is adapted from the [Contributor Covenant][homepage], version 1.4,
available at [http://contributor-covenant.org/version/1/4][version]

[homepage]: http://contributor-covenant.org
[version]: http://contributor-covenant.org/version/1/4/
## Contribution Guidelines

### Pull requests are always welcome

We are always thrilled to receive pull requests, and do our best to
process them as fast as possible. Not sure if that typo is worth a pull
request? Do it! We will appreciate it.

If your pull request is not accepted on the first try, don't be
discouraged! If there's a problem with the implementation, hopefully you
received feedback on what to improve.

We're trying very hard to keep go-swagger lean and focused. We don't want it
to do everything for everybody. This means that we might decide against
incorporating a new feature. However, there might be a way to implement
that feature *on top of* go-swagger.


### Conventions

Fork the repo and make changes on your fork in a feature branch:

- If it's a bugfix branch, name it XXX-something where XXX is the number of the
  issue
- If it's a feature branch, create an enhancement issue to announce your
  intentions, and name it XXX-something where XXX is the number of the issue.

Submit unit tests for your changes.  Go has a great test framework built in; use
it! Take a look at existing tests for inspiration. Run the full test suite on
your branch before submitting a pull request.

Update the documentation when creating or modifying features. Test
your documentation changes for clarity, concision, and correctness, as
well as a clean documentation build. See ``docs/README.md`` for more
information on building the docs and how docs get released.

Write clean code. Universally formatted code promotes ease of writing, reading,
and maintenance. Always run `gofmt -s -w file.go` on each changed file before
committing your changes. Most editors have plugins that do this automatically.

Pull requests descriptions should be as clear as possible and include a
reference to all the issues that they address.

Pull requests must not contain commits from other users or branches.

Commit messages must start with a capitalized and short summary (max. 50
chars) written in the imperative, followed by an optional, more detailed
explanatory text which is separated from the summary by an empty line.

Code review comments may be added to your pull request. Discuss, then make the
suggested modifications and push additional commits to your feature branch. Be
sure to post a comment after pushing. The new commits will show up in the pull
request automatically, but the reviewers will not be notified unless you
comment.

Before the pull request is merged, make sure that you squash your commits into
logical units of work using `git rebase -i` and `git push -f`. After every
commit the test suite should be passing. Include documentation changes in the
same commit so that a revert would remove all traces of the feature or fix.

Commits that fix or close an issue should include a reference like `Closes #XXX`
or `Fixes #XXX`, which will automatically close the issue when merged.

### Sign your work

The sign-off is a simple line at the end of the explanation for the
patch, which certifies that you wrote it or otherwise have the right to
pass it on as an open-source patch.  The rules are pretty simple: if you
can certify the below (from
[developercertificate.org](http://developercertificate.org/)):

```
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
660 York Street, Suite 102,
San Francisco, CA 94110 USA

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.


Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

then you just add a line to every git commit message:

    Signed-off-by: Joe Smith <joe@gmail.com>

using your real name (sorry, no pseudonyms or anonymous contributions.)

You can add the sign off when creating the git commit via `git commit -s`.
#  Audit

Audit of github repos to check for inconsistencies.
# inflect [![Build Status](https://github.com/go-openapi/inflect/actions/workflows/go-test.yml/badge.svg)](https://github.com/go-openapi/inflect/actions?query=workflow%3A"go+test") [![codecov](https://codecov.io/gh/go-openapi/inflect/branch/master/graph/badge.svg)](https://codecov.io/gh/go-openapi/inflect)

[![Slack Status](https://slackin.goswagger.io/badge.svg)](https://slackin.goswagger.io)
[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/go-openapi/inflect/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-openapi/inflect.svg)](https://pkg.go.dev/github.com/go-openapi/inflect)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-openapi/inflect)](https://goreportcard.com/report/github.com/go-openapi/inflect)

A package to pluralize words.

Originally forked from fork of https://bitbucket.org/pkg/inflect under a MIT License.

A golang library applying grammar rules to English words.

> This package provides a basic set of functions applying
> grammar rules to inflect English words, modify case style
> (Capitalize, camelCase, snake_case, etc.).
>
> Acronyms are properly handled. A common use case is word pluralization.


This library is not used at all by other go-openapi packages and is somewhat redundant with
go-openapi/swag/mangling (for camelcase etc).

Currently we have one single dependency in one place in a go-swagger template (used as a funcmap).

## Note to maintainers

* Aug. 2025: CI workflows have now moved to a shared repository go-openapi/ci-workflows
## Contribution Guidelines

### Pull requests are always welcome

We are always thrilled to receive pull requests, and do our best to
process them as fast as possible. Not sure if that typo is worth a pull
request? Do it! We will appreciate it.

If your pull request is not accepted on the first try, don't be
discouraged! If there's a problem with the implementation, hopefully you
received feedback on what to improve.

We're trying very hard to keep go-swagger lean and focused. We don't want it
to do everything for everybody. This means that we might decide against
incorporating a new feature. However, there might be a way to implement
that feature *on top of* go-swagger.


### Conventions

Fork the repo and make changes on your fork in a feature branch:

- If it's a bugfix branch, name it XXX-something where XXX is the number of the
  issue
- If it's a feature branch, create an enhancement issue to announce your
  intentions, and name it XXX-something where XXX is the number of the issue.

Submit unit tests for your changes.  Go has a great test framework built in; use
it! Take a look at existing tests for inspiration. Run the full test suite on
your branch before submitting a pull request.

Update the documentation when creating or modifying features. Test
your documentation changes for clarity, concision, and correctness, as
well as a clean documentation build. See ``docs/README.md`` for more
information on building the docs and how docs get released.

Write clean code. Universally formatted code promotes ease of writing, reading,
and maintenance. Always run `gofmt -s -w file.go` on each changed file before
committing your changes. Most editors have plugins that do this automatically.

Pull requests descriptions should be as clear as possible and include a
reference to all the issues that they address.

Pull requests must not contain commits from other users or branches.

Commit messages must start with a capitalized and short summary (max. 50
chars) written in the imperative, followed by an optional, more detailed
explanatory text which is separated from the summary by an empty line.

Code review comments may be added to your pull request. Discuss, then make the
suggested modifications and push additional commits to your feature branch. Be
sure to post a comment after pushing. The new commits will show up in the pull
request automatically, but the reviewers will not be notified unless you
comment.

Before the pull request is merged, make sure that you squash your commits into
logical units of work using `git rebase -i` and `git push -f`. After every
commit the test suite should be passing. Include documentation changes in the
same commit so that a revert would remove all traces of the feature or fix.

Commits that fix or close an issue should include a reference like `Closes #XXX`
or `Fixes #XXX`, which will automatically close the issue when merged.

### Sign your work

The sign-off is a simple line at the end of the explanation for the
patch, which certifies that you wrote it or otherwise have the right to
pass it on as an open-source patch.  The rules are pretty simple: if you
can certify the below (from
[developercertificate.org](http://developercertificate.org/)):

```
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
660 York Street, Suite 102,
San Francisco, CA 94110 USA

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.


Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

then you just add a line to every git commit message:

    Signed-off-by: Joe Smith <joe@gmail.com>

using your real name (sorry, no pseudonyms or anonymous contributions.)

You can add the sign off when creating the git commit via `git commit -s`.
# Validation helpers [![Build Status](https://github.com/go-openapi/validate/actions/workflows/go-test.yml/badge.svg)](https://github.com/go-openapi/validate/actions?query=workflow%3A"go+test") [![codecov](https://codecov.io/gh/go-openapi/validate/branch/master/graph/badge.svg)](https://codecov.io/gh/go-openapi/validate)

[![Slack Status](https://slackin.goswagger.io/badge.svg)](https://slackin.goswagger.io)
[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/go-openapi/validate/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-openapi/validate.svg)](https://pkg.go.dev/github.com/go-openapi/validate)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-openapi/validate)](https://goreportcard.com/report/github.com/go-openapi/validate)

This package provides helpers to validate Swagger 2.0. specification (aka OpenAPI 2.0). 

Reference can be found here: https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md.

## What's inside?

* A validator for Swagger specifications
* A validator for JSON schemas draft4
* Helper functions to validate individual values (used by code generated by [go-swagger](https://github.com/go-swagger/go-swagger)).
  * Required, RequiredNumber, RequiredString
  * ReadOnly
  * UniqueItems, MaxItems, MinItems
  * Enum, EnumCase
  * Pattern, MinLength, MaxLength
  * Minimum, Maximum, MultipleOf
  * FormatOf

[Documentation](https://pkg.go.dev/github.com/go-openapi/validate)

## FAQ

* Does this library support OpenAPI 3?

> No.
> This package currently only supports OpenAPI 2.0 (aka Swagger 2.0).
> There is no plan to make it evolve toward supporting OpenAPI 3.x.
> This [discussion thread](https://github.com/go-openapi/spec/issues/21) relates the full story.
>
> An early attempt to support Swagger 3 may be found at: https://github.com/go-openapi/spec3
# Contributor Covenant Code of Conduct

## Our Pledge

In the interest of fostering an open and welcoming environment, we as
contributors and maintainers pledge to making participation in our project and
our community a harassment-free experience for everyone, regardless of age, body
size, disability, ethnicity, gender identity and expression, level of experience,
nationality, personal appearance, race, religion, or sexual identity and
orientation.

## Our Standards

Examples of behavior that contributes to creating a positive environment
include:

* Using welcoming and inclusive language
* Being respectful of differing viewpoints and experiences
* Gracefully accepting constructive criticism
* Focusing on what is best for the community
* Showing empathy towards other community members

Examples of unacceptable behavior by participants include:

* The use of sexualized language or imagery and unwelcome sexual attention or
advances
* Trolling, insulting/derogatory comments, and personal or political attacks
* Public or private harassment
* Publishing others' private information, such as a physical or electronic
  address, without explicit permission
* Other conduct which could reasonably be considered inappropriate in a
  professional setting

## Our Responsibilities

Project maintainers are responsible for clarifying the standards of acceptable
behavior and are expected to take appropriate and fair corrective action in
response to any instances of unacceptable behavior.

Project maintainers have the right and responsibility to remove, edit, or
reject comments, commits, code, wiki edits, issues, and other contributions
that are not aligned to this Code of Conduct, or to ban temporarily or
permanently any contributor for other behaviors that they deem inappropriate,
threatening, offensive, or harmful.

## Scope

This Code of Conduct applies both within project spaces and in public spaces
when an individual is representing the project or its community. Examples of
representing a project or community include using an official project e-mail
address, posting via an official social media account, or acting as an appointed
representative at an online or offline event. Representation of a project may be
further defined and clarified by project maintainers.

## Enforcement

Instances of abusive, harassing, or otherwise unacceptable behavior may be
reported by contacting the project team at ivan+abuse@flanders.co.nz. All
complaints will be reviewed and investigated and will result in a response that
is deemed necessary and appropriate to the circumstances. The project team is
obligated to maintain confidentiality with regard to the reporter of an incident.
Further details of specific enforcement policies may be posted separately.

Project maintainers who do not follow or enforce the Code of Conduct in good
faith may face temporary or permanent repercussions as determined by other
members of the project's leadership.

## Attribution

This Code of Conduct is adapted from the [Contributor Covenant][homepage], version 1.4,
available at [http://contributor-covenant.org/version/1/4][version]

[homepage]: http://contributor-covenant.org
[version]: http://contributor-covenant.org/version/1/4/
## Contribution Guidelines

### Pull requests are always welcome

We are always thrilled to receive pull requests, and do our best to
process them as fast as possible. Not sure if that typo is worth a pull
request? Do it! We will appreciate it.

If your pull request is not accepted on the first try, don't be
discouraged! If there's a problem with the implementation, hopefully you
received feedback on what to improve.

We're trying very hard to keep go-swagger lean and focused. We don't want it
to do everything for everybody. This means that we might decide against
incorporating a new feature. However, there might be a way to implement
that feature *on top of* go-swagger.


### Conventions

Fork the repo and make changes on your fork in a feature branch:

- If it's a bugfix branch, name it XXX-something where XXX is the number of the
  issue
- If it's a feature branch, create an enhancement issue to announce your
  intentions, and name it XXX-something where XXX is the number of the issue.

Submit unit tests for your changes.  Go has a great test framework built in; use
it! Take a look at existing tests for inspiration. Run the full test suite on
your branch before submitting a pull request.

Update the documentation when creating or modifying features. Test
your documentation changes for clarity, concision, and correctness, as
well as a clean documentation build. See ``docs/README.md`` for more
information on building the docs and how docs get released.

Write clean code. Universally formatted code promotes ease of writing, reading,
and maintenance. Always run `gofmt -s -w file.go` on each changed file before
committing your changes. Most editors have plugins that do this automatically.

Pull requests descriptions should be as clear as possible and include a
reference to all the issues that they address.

Pull requests must not contain commits from other users or branches.

Commit messages must start with a capitalized and short summary (max. 50
chars) written in the imperative, followed by an optional, more detailed
explanatory text which is separated from the summary by an empty line.

Code review comments may be added to your pull request. Discuss, then make the
suggested modifications and push additional commits to your feature branch. Be
sure to post a comment after pushing. The new commits will show up in the pull
request automatically, but the reviewers will not be notified unless you
comment.

Before the pull request is merged, make sure that you squash your commits into
logical units of work using `git rebase -i` and `git push -f`. After every
commit the test suite should be passing. Include documentation changes in the
same commit so that a revert would remove all traces of the feature or fix.

Commits that fix or close an issue should include a reference like `Closes #XXX`
or `Fixes #XXX`, which will automatically close the issue when merged.

### Sign your work

The sign-off is a simple line at the end of the explanation for the
patch, which certifies that you wrote it or otherwise have the right to
pass it on as an open-source patch.  The rules are pretty simple: if you
can certify the below (from
[developercertificate.org](http://developercertificate.org/)):

```
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
660 York Street, Suite 102,
San Francisco, CA 94110 USA

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.


Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

then you just add a line to every git commit message:

    Signed-off-by: Joe Smith <joe@gmail.com>

using your real name (sorry, no pseudonyms or anonymous contributions.)

You can add the sign off when creating the git commit via `git commit -s`.
# Benchmark

Validating the Kubernetes Swagger API

## v0.22.6: 60,000,000 allocs
```
goos: linux
goarch: amd64
pkg: github.com/go-openapi/validate
cpu: AMD Ryzen 7 5800X 8-Core Processor
Benchmark_KubernetesSpec/validating_kubernetes_API-16         	       1	8549863982 ns/op	7067424936 B/op	59583275 allocs/op
```

## After refact PR: minor but noticable improvements: 25,000,000 allocs
```
go test -bench Spec
goos: linux
goarch: amd64
pkg: github.com/go-openapi/validate
cpu: AMD Ryzen 7 5800X 8-Core Processor
Benchmark_KubernetesSpec/validating_kubernetes_API-16         	       1	4064535557 ns/op	3379715592 B/op	25320330 allocs/op
```

## After reduce GC pressure PR: 17,000,000 allocs
```
goos: linux
goarch: amd64
pkg: github.com/go-openapi/validate
cpu: AMD Ryzen 7 5800X 8-Core Processor             
Benchmark_KubernetesSpec/validating_kubernetes_API-16         	       1	3758414145 ns/op	2593881496 B/op	17111373 allocs/op
```
# Denco [![Build Status](https://travis-ci.org/naoina/denco.png?branch=master)](https://travis-ci.org/naoina/denco)

The fast and flexible HTTP request router for [Go](http://golang.org).

Denco is based on Double-Array implementation of [Kocha-urlrouter](https://github.com/naoina/kocha-urlrouter).
However, Denco is optimized and some features added.

## Features

* Fast (See [go-http-routing-benchmark](https://github.com/naoina/go-http-routing-benchmark))
* [URL patterns](#url-patterns) (`/foo/:bar` and `/foo/*wildcard`)
* Small (but enough) URL router API
* HTTP request multiplexer like `http.ServeMux`

## Installation

    go get -u github.com/go-openapi/runtime/middleware/denco

## Using as HTTP request multiplexer

```go
package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/go-openapi/runtime/middleware/denco"
)

func Index(w http.ResponseWriter, r *http.Request, params denco.Params) {
    fmt.Fprintf(w, "Welcome to Denco!\n")
}

func User(w http.ResponseWriter, r *http.Request, params denco.Params) {
    fmt.Fprintf(w, "Hello %s!\n", params.Get("name"))
}

func main() {
    mux := denco.NewMux()
    handler, err := mux.Build([]denco.Handler{
        mux.GET("/", Index),
        mux.GET("/user/:name", User),
        mux.POST("/user/:name", User),
    })
    if err != nil {
        panic(err)
    }
    log.Fatal(http.ListenAndServe(":8080", handler))
}
```

## Using as URL router

```go
package main

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware/denco"
)

type route struct {
	name string
}

func main() {
	router := denco.New()
	router.Build([]denco.Record{
		{"/", &route{"root"}},
		{"/user/:id", &route{"user"}},
		{"/user/:name/:id", &route{"username"}},
		{"/static/*filepath", &route{"static"}},
	})

	data, params, found := router.Lookup("/")
	// print `&main.route{name:"root"}, denco.Params(nil), true`.
	fmt.Printf("%#v, %#v, %#v\n", data, params, found)

	data, params, found = router.Lookup("/user/hoge")
	// print `&main.route{name:"user"}, denco.Params{denco.Param{Name:"id", Value:"hoge"}}, true`.
	fmt.Printf("%#v, %#v, %#v\n", data, params, found)

	data, params, found = router.Lookup("/user/hoge/7")
	// print `&main.route{name:"username"}, denco.Params{denco.Param{Name:"name", Value:"hoge"}, denco.Param{Name:"id", Value:"7"}}, true`.
	fmt.Printf("%#v, %#v, %#v\n", data, params, found)

	data, params, found = router.Lookup("/static/path/to/file")
	// print `&main.route{name:"static"}, denco.Params{denco.Param{Name:"filepath", Value:"path/to/file"}}, true`.
	fmt.Printf("%#v, %#v, %#v\n", data, params, found)
}
```

See [Godoc](http://godoc.org/github.com/go-openapi/runtime/middleware/denco) for more details.

## Getting the value of path parameter

You can get the value of path parameter by 2 ways.

1. Using [`denco.Params.Get`](http://godoc.org/github.com/go-openapi/runtime/middleware/denco#Params.Get) method
2. Find by loop

```go
package main

import (
    "fmt"

    "github.com/go-openapi/runtime/middleware/denco"
)

func main() {
    router := denco.New()
    if err := router.Build([]denco.Record{
        {"/user/:name/:id", "route1"},
    }); err != nil {
        panic(err)
    }

    // 1. Using denco.Params.Get method.
    _, params, _ := router.Lookup("/user/alice/1")
    name := params.Get("name")
    if name != "" {
        fmt.Printf("Hello %s.\n", name) // prints "Hello alice.".
    }

    // 2. Find by loop.
    for _, param := range params {
        if param.Name == "name" {
            fmt.Printf("Hello %s.\n", name) // prints "Hello alice.".
        }
    }
}
```

## URL patterns

Denco's route matching strategy is "most nearly matching".

When routes `/:name` and `/alice` have been built, URI `/alice` matches the route `/alice`, not `/:name`.
Because URI `/alice` is more match with the route `/alice` than `/:name`.

For more example, when routes below have been built:

```
/user/alice
/user/:name
/user/:name/:id
/user/alice/:id
/user/:id/bob
```

Routes matching are:

```
/user/alice      => "/user/alice" (no match with "/user/:name")
/user/bob        => "/user/:name"
/user/naoina/1   => "/user/:name/1"
/user/alice/1    => "/user/alice/:id" (no match with "/user/:name/:id")
/user/1/bob      => "/user/:id/bob"   (no match with "/user/:name/:id")
/user/alice/bob  => "/user/alice/:id" (no match with "/user/:name/:id" and "/user/:id/bob")
```

## Limitation

Denco has some limitations below.

* Number of param records (such as `/:name`) must be less than 2^22
* Number of elements of internal slice must be less than 2^22

## Benchmarks

    cd $GOPATH/github.com/go-openapi/runtime/middleware/denco
    go test -bench . -benchmem

## License

Denco is licensed under the MIT License.
# runtime [![Build Status](https://github.com/go-openapi/runtime/actions/workflows/go-test.yml/badge.svg)](https://github.com/go-openapi/runtime/actions?query=workflow%3A"go+test") [![codecov](https://codecov.io/gh/go-openapi/runtime/branch/master/graph/badge.svg)](https://codecov.io/gh/go-openapi/runtime)

[![Slack Status](https://slackin.goswagger.io/badge.svg)](https://slackin.goswagger.io)
[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/go-openapi/runtime/master/LICENSE) 
[![Go Reference](https://pkg.go.dev/badge/github.com/go-openapi/runtime.svg)](https://pkg.go.dev/github.com/go-openapi/runtime)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-openapi/runtime)](https://goreportcard.com/report/github.com/go-openapi/runtime)

# go OpenAPI toolkit runtime

The runtime component for use in code generation or as untyped usage.
# Contributor Covenant Code of Conduct

## Our Pledge

In the interest of fostering an open and welcoming environment, we as
contributors and maintainers pledge to making participation in our project and
our community a harassment-free experience for everyone, regardless of age, body
size, disability, ethnicity, gender identity and expression, level of experience,
nationality, personal appearance, race, religion, or sexual identity and
orientation.

## Our Standards

Examples of behavior that contributes to creating a positive environment
include:

* Using welcoming and inclusive language
* Being respectful of differing viewpoints and experiences
* Gracefully accepting constructive criticism
* Focusing on what is best for the community
* Showing empathy towards other community members

Examples of unacceptable behavior by participants include:

* The use of sexualized language or imagery and unwelcome sexual attention or
advances
* Trolling, insulting/derogatory comments, and personal or political attacks
* Public or private harassment
* Publishing others' private information, such as a physical or electronic
  address, without explicit permission
* Other conduct which could reasonably be considered inappropriate in a
  professional setting

## Our Responsibilities

Project maintainers are responsible for clarifying the standards of acceptable
behavior and are expected to take appropriate and fair corrective action in
response to any instances of unacceptable behavior.

Project maintainers have the right and responsibility to remove, edit, or
reject comments, commits, code, wiki edits, issues, and other contributions
that are not aligned to this Code of Conduct, or to ban temporarily or
permanently any contributor for other behaviors that they deem inappropriate,
threatening, offensive, or harmful.

## Scope

This Code of Conduct applies both within project spaces and in public spaces
when an individual is representing the project or its community. Examples of
representing a project or community include using an official project e-mail
address, posting via an official social media account, or acting as an appointed
representative at an online or offline event. Representation of a project may be
further defined and clarified by project maintainers.

## Enforcement

Instances of abusive, harassing, or otherwise unacceptable behavior may be
reported by contacting the project team at ivan+abuse@flanders.co.nz. All
complaints will be reviewed and investigated and will result in a response that
is deemed necessary and appropriate to the circumstances. The project team is
obligated to maintain confidentiality with regard to the reporter of an incident.
Further details of specific enforcement policies may be posted separately.

Project maintainers who do not follow or enforce the Code of Conduct in good
faith may face temporary or permanent repercussions as determined by other
members of the project's leadership.

## Attribution

This Code of Conduct is adapted from the [Contributor Covenant][homepage], version 1.4,
available at [http://contributor-covenant.org/version/1/4][version]

[homepage]: http://contributor-covenant.org
[version]: http://contributor-covenant.org/version/1/4/
## Contribution Guidelines

### Pull requests are always welcome

We are always thrilled to receive pull requests, and do our best to
process them as fast as possible. Not sure if that typo is worth a pull
request? Do it! We will appreciate it.

If your pull request is not accepted on the first try, don't be
discouraged! If there's a problem with the implementation, hopefully you
received feedback on what to improve.

We're trying very hard to keep go-swagger lean and focused. We don't want it
to do everything for everybody. This means that we might decide against
incorporating a new feature. However, there might be a way to implement
that feature *on top of* go-swagger.


### Conventions

Fork the repo and make changes on your fork in a feature branch:

- If it's a bugfix branch, name it XXX-something where XXX is the number of the
  issue
- If it's a feature branch, create an enhancement issue to announce your
  intentions, and name it XXX-something where XXX is the number of the issue.

Submit unit tests for your changes.  Go has a great test framework built in; use
it! Take a look at existing tests for inspiration. Run the full test suite on
your branch before submitting a pull request.

Update the documentation when creating or modifying features. Test
your documentation changes for clarity, concision, and correctness, as
well as a clean documentation build. See ``docs/README.md`` for more
information on building the docs and how docs get released.

Write clean code. Universally formatted code promotes ease of writing, reading,
and maintenance. Always run `gofmt -s -w file.go` on each changed file before
committing your changes. Most editors have plugins that do this automatically.

Pull requests descriptions should be as clear as possible and include a
reference to all the issues that they address.

Pull requests must not contain commits from other users or branches.

Commit messages must start with a capitalized and short summary (max. 50
chars) written in the imperative, followed by an optional, more detailed
explanatory text which is separated from the summary by an empty line.

Code review comments may be added to your pull request. Discuss, then make the
suggested modifications and push additional commits to your feature branch. Be
sure to post a comment after pushing. The new commits will show up in the pull
request automatically, but the reviewers will not be notified unless you
comment.

Before the pull request is merged, make sure that you squash your commits into
logical units of work using `git rebase -i` and `git push -f`. After every
commit the test suite should be passing. Include documentation changes in the
same commit so that a revert would remove all traces of the feature or fix.

Commits that fix or close an issue should include a reference like `Closes #XXX`
or `Fixes #XXX`, which will automatically close the issue when merged.

### Sign your work

The sign-off is a simple line at the end of the explanation for the
patch, which certifies that you wrote it or otherwise have the right to
pass it on as an open-source patch.  The rules are pretty simple: if you
can certify the below (from
[developercertificate.org](http://developercertificate.org/)):

```
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
660 York Street, Suite 102,
San Francisco, CA 94110 USA

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.


Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

then you just add a line to every git commit message:

    Signed-off-by: Joe Smith <joe@gmail.com>

using your real name (sorry, no pseudonyms or anonymous contributions.)

You can add the sign off when creating the git commit via `git commit -s`.
# Loads OAI specs [![Build Status](https://github.com/go-openapi/loads/actions/workflows/go-test.yml/badge.svg)](https://github.com/go-openapi/loads/actions?query=workflow%3A"go+test") [![codecov](https://codecov.io/gh/go-openapi/loads/branch/master/graph/badge.svg)](https://codecov.io/gh/go-openapi/loads)

[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/go-openapi/loads/master/LICENSE) [![GoDoc](https://godoc.org/github.com/go-openapi/loads?status.svg)](http://godoc.org/github.com/go-openapi/loads)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-openapi/loads)](https://goreportcard.com/report/github.com/go-openapi/loads)

Loading of OAI specification documents from local or remote locations. Supports JSON and YAML documents.
# Contributor Covenant Code of Conduct

## Our Pledge

In the interest of fostering an open and welcoming environment, we as
contributors and maintainers pledge to making participation in our project and
our community a harassment-free experience for everyone, regardless of age, body
size, disability, ethnicity, gender identity and expression, level of experience,
nationality, personal appearance, race, religion, or sexual identity and
orientation.

## Our Standards

Examples of behavior that contributes to creating a positive environment
include:

* Using welcoming and inclusive language
* Being respectful of differing viewpoints and experiences
* Gracefully accepting constructive criticism
* Focusing on what is best for the community
* Showing empathy towards other community members

Examples of unacceptable behavior by participants include:

* The use of sexualized language or imagery and unwelcome sexual attention or
advances
* Trolling, insulting/derogatory comments, and personal or political attacks
* Public or private harassment
* Publishing others' private information, such as a physical or electronic
  address, without explicit permission
* Other conduct which could reasonably be considered inappropriate in a
  professional setting

## Our Responsibilities

Project maintainers are responsible for clarifying the standards of acceptable
behavior and are expected to take appropriate and fair corrective action in
response to any instances of unacceptable behavior.

Project maintainers have the right and responsibility to remove, edit, or
reject comments, commits, code, wiki edits, issues, and other contributions
that are not aligned to this Code of Conduct, or to ban temporarily or
permanently any contributor for other behaviors that they deem inappropriate,
threatening, offensive, or harmful.

## Scope

This Code of Conduct applies both within project spaces and in public spaces
when an individual is representing the project or its community. Examples of
representing a project or community include using an official project e-mail
address, posting via an official social media account, or acting as an appointed
representative at an online or offline event. Representation of a project may be
further defined and clarified by project maintainers.

## Enforcement

Instances of abusive, harassing, or otherwise unacceptable behavior may be
reported by contacting the project team at ivan+abuse@flanders.co.nz. All
complaints will be reviewed and investigated and will result in a response that
is deemed necessary and appropriate to the circumstances. The project team is
obligated to maintain confidentiality with regard to the reporter of an incident.
Further details of specific enforcement policies may be posted separately.

Project maintainers who do not follow or enforce the Code of Conduct in good
faith may face temporary or permanent repercussions as determined by other
members of the project's leadership.

## Attribution

This Code of Conduct is adapted from the [Contributor Covenant][homepage], version 1.4,
available at [http://contributor-covenant.org/version/1/4][version]

[homepage]: http://contributor-covenant.org
[version]: http://contributor-covenant.org/version/1/4/
## Contribution Guidelines

### Pull requests are always welcome

We are always thrilled to receive pull requests, and do our best to
process them as fast as possible. Not sure if that typo is worth a pull
request? Do it! We will appreciate it.

If your pull request is not accepted on the first try, don't be
discouraged! If there's a problem with the implementation, hopefully you
received feedback on what to improve.

We're trying very hard to keep go-swagger lean and focused. We don't want it
to do everything for everybody. This means that we might decide against
incorporating a new feature. However, there might be a way to implement
that feature *on top of* go-swagger.


### Conventions

Fork the repo and make changes on your fork in a feature branch:

- If it's a bugfix branch, name it XXX-something where XXX is the number of the
  issue
- If it's a feature branch, create an enhancement issue to announce your
  intentions, and name it XXX-something where XXX is the number of the issue.

Submit unit tests for your changes.  Go has a great test framework built in; use
it! Take a look at existing tests for inspiration. Run the full test suite on
your branch before submitting a pull request.

Update the documentation when creating or modifying features. Test
your documentation changes for clarity, concision, and correctness, as
well as a clean documentation build. See ``docs/README.md`` for more
information on building the docs and how docs get released.

Write clean code. Universally formatted code promotes ease of writing, reading,
and maintenance. Always run `gofmt -s -w file.go` on each changed file before
committing your changes. Most editors have plugins that do this automatically.

Pull requests descriptions should be as clear as possible and include a
reference to all the issues that they address.

Pull requests must not contain commits from other users or branches.

Commit messages must start with a capitalized and short summary (max. 50
chars) written in the imperative, followed by an optional, more detailed
explanatory text which is separated from the summary by an empty line.

Code review comments may be added to your pull request. Discuss, then make the
suggested modifications and push additional commits to your feature branch. Be
sure to post a comment after pushing. The new commits will show up in the pull
request automatically, but the reviewers will not be notified unless you
comment.

Before the pull request is merged, make sure that you squash your commits into
logical units of work using `git rebase -i` and `git push -f`. After every
commit the test suite should be passing. Include documentation changes in the
same commit so that a revert would remove all traces of the feature or fix.

Commits that fix or close an issue should include a reference like `Closes #XXX`
or `Fixes #XXX`, which will automatically close the issue when merged.

### Sign your work

The sign-off is a simple line at the end of the explanation for the
patch, which certifies that you wrote it or otherwise have the right to
pass it on as an open-source patch.  The rules are pretty simple: if you
can certify the below (from
[developercertificate.org](http://developercertificate.org/)):

```
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
660 York Street, Suite 102,
San Francisco, CA 94110 USA

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.


Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

then you just add a line to every git commit message:

    Signed-off-by: Joe Smith <joe@gmail.com>

using your real name (sorry, no pseudonyms or anonymous contributions.)

You can add the sign off when creating the git commit via `git commit -s`.
# OpenAPI analysis [![Build Status](https://github.com/go-openapi/analysis/actions/workflows/go-test.yml/badge.svg)](https://github.com/go-openapi/analysis/actions?query=workflow%3A"go+test") [![codecov](https://codecov.io/gh/go-openapi/analysis/branch/master/graph/badge.svg)](https://codecov.io/gh/go-openapi/analysis)

[![Slack Status](https://slackin.goswagger.io/badge.svg)](https://slackin.goswagger.io)
[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/go-openapi/analysis/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-openapi/analysis.svg)](https://pkg.go.dev/github.com/go-openapi/analysis)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-openapi/analysis)](https://goreportcard.com/report/github.com/go-openapi/analysis)


A foundational library to analyze an OAI specification document for easier reasoning about the content.

## What's inside?

* An analyzer providing methods to walk the functional content of a specification
* A spec flattener producing a self-contained document bundle, while preserving `$ref`s
* A spec merger ("mixin") to merge several spec documents into a primary spec
* A spec "fixer" ensuring that response descriptions are non empty

[Documentation](https://pkg.go.dev/github.com/go-openapi/analysis)

## FAQ

* Does this library support OpenAPI 3?

> No.
> This package currently only supports OpenAPI 2.0 (aka Swagger 2.0).
> There is no plan to make it evolve toward supporting OpenAPI 3.x.
> This [discussion thread](https://github.com/go-openapi/spec/issues/21) relates the full story.
# Contributor Covenant Code of Conduct

## Our Pledge

In the interest of fostering an open and welcoming environment, we as
contributors and maintainers pledge to making participation in our project and
our community a harassment-free experience for everyone, regardless of age, body
size, disability, ethnicity, gender identity and expression, level of experience,
nationality, personal appearance, race, religion, or sexual identity and
orientation.

## Our Standards

Examples of behavior that contributes to creating a positive environment
include:

* Using welcoming and inclusive language
* Being respectful of differing viewpoints and experiences
* Gracefully accepting constructive criticism
* Focusing on what is best for the community
* Showing empathy towards other community members

Examples of unacceptable behavior by participants include:

* The use of sexualized language or imagery and unwelcome sexual attention or
advances
* Trolling, insulting/derogatory comments, and personal or political attacks
* Public or private harassment
* Publishing others' private information, such as a physical or electronic
  address, without explicit permission
* Other conduct which could reasonably be considered inappropriate in a
  professional setting

## Our Responsibilities

Project maintainers are responsible for clarifying the standards of acceptable
behavior and are expected to take appropriate and fair corrective action in
response to any instances of unacceptable behavior.

Project maintainers have the right and responsibility to remove, edit, or
reject comments, commits, code, wiki edits, issues, and other contributions
that are not aligned to this Code of Conduct, or to ban temporarily or
permanently any contributor for other behaviors that they deem inappropriate,
threatening, offensive, or harmful.

## Scope

This Code of Conduct applies both within project spaces and in public spaces
when an individual is representing the project or its community. Examples of
representing a project or community include using an official project e-mail
address, posting via an official social media account, or acting as an appointed
representative at an online or offline event. Representation of a project may be
further defined and clarified by project maintainers.

## Enforcement

Instances of abusive, harassing, or otherwise unacceptable behavior may be
reported by contacting the project team at ivan+abuse@flanders.co.nz. All
complaints will be reviewed and investigated and will result in a response that
is deemed necessary and appropriate to the circumstances. The project team is
obligated to maintain confidentiality with regard to the reporter of an incident.
Further details of specific enforcement policies may be posted separately.

Project maintainers who do not follow or enforce the Code of Conduct in good
faith may face temporary or permanent repercussions as determined by other
members of the project's leadership.

## Attribution

This Code of Conduct is adapted from the [Contributor Covenant][homepage], version 1.4,
available at [http://contributor-covenant.org/version/1/4][version]

[homepage]: http://contributor-covenant.org
[version]: http://contributor-covenant.org/version/1/4/
## Contribution Guidelines

### Pull requests are always welcome

We are always thrilled to receive pull requests, and do our best to
process them as fast as possible. Not sure if that typo is worth a pull
request? Do it! We will appreciate it.

If your pull request is not accepted on the first try, don't be
discouraged! If there's a problem with the implementation, hopefully you
received feedback on what to improve.

We're trying very hard to keep go-swagger lean and focused. We don't want it
to do everything for everybody. This means that we might decide against
incorporating a new feature. However, there might be a way to implement
that feature *on top of* go-swagger.


### Conventions

Fork the repo and make changes on your fork in a feature branch:

- If it's a bugfix branch, name it XXX-something where XXX is the number of the
  issue
- If it's a feature branch, create an enhancement issue to announce your
  intentions, and name it XXX-something where XXX is the number of the issue.

Submit unit tests for your changes.  Go has a great test framework built in; use
it! Take a look at existing tests for inspiration. Run the full test suite on
your branch before submitting a pull request.

Update the documentation when creating or modifying features. Test
your documentation changes for clarity, concision, and correctness, as
well as a clean documentation build. See ``docs/README.md`` for more
information on building the docs and how docs get released.

Write clean code. Universally formatted code promotes ease of writing, reading,
and maintenance. Always run `gofmt -s -w file.go` on each changed file before
committing your changes. Most editors have plugins that do this automatically.

Pull requests descriptions should be as clear as possible and include a
reference to all the issues that they address.

Pull requests must not contain commits from other users or branches.

Commit messages must start with a capitalized and short summary (max. 50
chars) written in the imperative, followed by an optional, more detailed
explanatory text which is separated from the summary by an empty line.

Code review comments may be added to your pull request. Discuss, then make the
suggested modifications and push additional commits to your feature branch. Be
sure to post a comment after pushing. The new commits will show up in the pull
request automatically, but the reviewers will not be notified unless you
comment.

Before the pull request is merged, make sure that you squash your commits into
logical units of work using `git rebase -i` and `git push -f`. After every
commit the test suite should be passing. Include documentation changes in the
same commit so that a revert would remove all traces of the feature or fix.

Commits that fix or close an issue should include a reference like `Closes #XXX`
or `Fixes #XXX`, which will automatically close the issue when merged.

### Sign your work

The sign-off is a simple line at the end of the explanation for the
patch, which certifies that you wrote it or otherwise have the right to
pass it on as an open-source patch.  The rules are pretty simple: if you
can certify the below (from
[developercertificate.org](http://developercertificate.org/)):

```
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
660 York Street, Suite 102,
San Francisco, CA 94110 USA

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.


Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

then you just add a line to every git commit message:

    Signed-off-by: Joe Smith <joe@gmail.com>

using your real name (sorry, no pseudonyms or anonymous contributions.)

You can add the sign off when creating the git commit via `git commit -s`.
# Benchmarking name mangling utilities

```bash
go test -bench XXX -run XXX -benchtime 30s
```

## Benchmarks at b3e7a5386f996177e4808f11acb2aa93a0f660df

```
goos: linux
goarch: amd64
pkg: github.com/go-openapi/swag
cpu: Intel(R) Core(TM) i5-6200U CPU @ 2.30GHz
BenchmarkToXXXName/ToGoName-4         	  862623	     44101 ns/op	   10450 B/op	     732 allocs/op
BenchmarkToXXXName/ToVarName-4        	  853656	     40728 ns/op	   10468 B/op	     734 allocs/op
BenchmarkToXXXName/ToFileName-4       	 1268312	     27813 ns/op	    9785 B/op	     617 allocs/op
BenchmarkToXXXName/ToCommandName-4    	 1276322	     27903 ns/op	    9785 B/op	     617 allocs/op
BenchmarkToXXXName/ToHumanNameLower-4 	  895334	     40354 ns/op	   10472 B/op	     731 allocs/op
BenchmarkToXXXName/ToHumanNameTitle-4 	  882441	     40678 ns/op	   10566 B/op	     749 allocs/op
```

## Benchmarks after PR #79

~ x10 performance improvement and ~ /100 memory allocations.

```
goos: linux
goarch: amd64
pkg: github.com/go-openapi/swag
cpu: Intel(R) Core(TM) i5-6200U CPU @ 2.30GHz
BenchmarkToXXXName/ToGoName-4         	 9595830	      3991 ns/op	      42 B/op	       5 allocs/op
BenchmarkToXXXName/ToVarName-4        	 9194276	      3984 ns/op	      62 B/op	       7 allocs/op
BenchmarkToXXXName/ToFileName-4       	17002711	      2123 ns/op	     147 B/op	       7 allocs/op
BenchmarkToXXXName/ToCommandName-4    	16772926	      2111 ns/op	     147 B/op	       7 allocs/op
BenchmarkToXXXName/ToHumanNameLower-4 	 9788331	      3749 ns/op	      92 B/op	       6 allocs/op
BenchmarkToXXXName/ToHumanNameTitle-4 	 9188260	      3941 ns/op	     104 B/op	       6 allocs/op
```

```
goos: linux
goarch: amd64
pkg: github.com/go-openapi/swag
cpu: AMD Ryzen 7 5800X 8-Core Processor             
BenchmarkToXXXName/ToGoName-16         	18527378	      1972 ns/op	      42 B/op	       5 allocs/op
BenchmarkToXXXName/ToVarName-16        	15552692	      2093 ns/op	      62 B/op	       7 allocs/op
BenchmarkToXXXName/ToFileName-16       	32161176	      1117 ns/op	     147 B/op	       7 allocs/op
BenchmarkToXXXName/ToCommandName-16    	32256634	      1137 ns/op	     147 B/op	       7 allocs/op
BenchmarkToXXXName/ToHumanNameLower-16 	18599661	      1946 ns/op	      92 B/op	       6 allocs/op
BenchmarkToXXXName/ToHumanNameTitle-16 	17581353	      2054 ns/op	     105 B/op	       6 allocs/op
```

## Benchmarks at d7d2d1b895f5b6747afaff312dd2a402e69e818b

go1.24

```
goos: linux
goarch: amd64
pkg: github.com/go-openapi/swag
cpu: AMD Ryzen 7 5800X 8-Core Processor             
BenchmarkToXXXName/ToGoName-16         	19757858	      1881 ns/op	      42 B/op	       5 allocs/op
BenchmarkToXXXName/ToVarName-16        	17494111	      2094 ns/op	      74 B/op	       7 allocs/op
BenchmarkToXXXName/ToFileName-16       	28161226	      1492 ns/op	     158 B/op	       7 allocs/op
BenchmarkToXXXName/ToCommandName-16    	23787333	      1489 ns/op	     158 B/op	       7 allocs/op
BenchmarkToXXXName/ToHumanNameLower-16 	17537257	      2030 ns/op	     103 B/op	       6 allocs/op
BenchmarkToXXXName/ToHumanNameTitle-16 	16977453	      2156 ns/op	     105 B/op	       6 allocs/op
```

## Benchmarks after PR #106

Moving the scope of everything down to a struct allowed to reduce a bit garbage and pooling.

On top of that, ToGoName (and thus ToVarName) have been subject to a minor optimization, removing a few allocations.

Overall timings improve by ~ -10%.

go1.24

```
goos: linux
goarch: amd64
pkg: github.com/go-openapi/swag/mangling
cpu: AMD Ryzen 7 5800X 8-Core Processor             
BenchmarkToXXXName/ToGoName-16         	22496130	      1618 ns/op	      31 B/op	       3 allocs/op
BenchmarkToXXXName/ToVarName-16        	22538068	      1618 ns/op	      33 B/op	       3 allocs/op
BenchmarkToXXXName/ToFileName-16       	27722977	      1236 ns/op	     105 B/op	       6 allocs/op
BenchmarkToXXXName/ToCommandName-16    	27967395	      1258 ns/op	     105 B/op	       6 allocs/op
BenchmarkToXXXName/ToHumanNameLower-16 	18587901	      1917 ns/op	     103 B/op	       6 allocs/op
BenchmarkToXXXName/ToHumanNameTitle-16 	17193208	      2019 ns/op	     108 B/op	       7 allocs/op
```
# Swag [![Build Status](https://github.com/go-openapi/swag/actions/workflows/go-test.yml/badge.svg)](https://github.com/go-openapi/swag/actions?query=workflow%3A"go+test") [![codecov](https://codecov.io/gh/go-openapi/swag/branch/master/graph/badge.svg)](https://codecov.io/gh/go-openapi/swag)

[![Slack Status](https://slackin.goswagger.io/badge.svg)](https://slackin.goswagger.io)
[![license](https://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/go-openapi/swag/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-openapi/swag.svg)](https://pkg.go.dev/github.com/go-openapi/swag)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-openapi/swag)](https://goreportcard.com/report/github.com/go-openapi/swag)

Package `swag` contains a bunch of helper functions for go-openapi and go-swagger projects.

You may also use it standalone for your projects.

> `swag` is one of the foundational building blocks of the go-openapi initiative.
>
> Most repositories in `github.com/go-openapi/...` depend on it in some way.
> So does the CLI tool `github.com/go-swagger/go-swagger`,
> and the code generated by this tool.

## Contents

`go-openapi/swag` now exposes a collection of relatively independent modules.

Here is what is inside:

*	Module `cmdutils`

  * [x] utilities to work with CLIs

*	Module `conv`

  * [x] convert between values and pointers for any types
  * [x] convert from string to builtin types (wraps `strconv`)
  * [x] require `./typeutils` (test dependency)

*	Module `fileutils`

  * [x] file upload type
  * [x] search in path (deprecated)

* Module `jsonname`

  * [x] infer JSON names from go properties

* Module `jsonutils`

  * [x] fast json concatenation
  * [x] read and write JSON from and to dynamic go data structures
  * [x] require `github.com/mailru/easyjson`

* Module `loading`

  * [x] load from file or http
  * [x] require `./yamlutils`

* Module `mangling`

  * [x] name mangling for go

* Module `netutils`

  * [x] host, port from address

*	Module `stringutils`

  * [x] search in slice (with case-insensitive)
  * [x] split/join query parameters as arrays

*	Module `typeutils`

  * [x] check the zero value for any type

* Module `yamlutils`

  * [x] converting YAML to JSON
  * [x] loading YAML into a dynamic YAML document
  * [x] require `./jsonutils`
  * [x] require `github.com/mailru/easyjson`
  * [x] require `gopkg.in/yaml.v3`

---

The root module `github.com/go-openapi/swag` at the repo level maintains a few 
dependencies outside of the standard library:

* YAML utilities depend on `gopkg.in/yaml.v3`
* JSON utilities `github.com/mailru/easyjson`

This is not necessarily the case for all sub-modules.

## Release notes

### v0.24.0 [Unreleased]

With this release, we have largely modernized the API of `swag`:

* The traditional `swag` API is still supported: code that imports `swag` will still
  compile and work the same.
* A deprecation notice is published to encourage consumers of this library to adopt
  the newer API
* **Deprecation notice** 
  * configuration through global variables is now deprecated, in favor of options passed as parameters
  * all helper functions are moved to more specialized packages, which are exposed as 
    go modules. Importing such a module would reduce the footprint of dependencies.
  * _all_ functions, variables, constants exposed by the deprecated API have now moved, so
    that consumers of the new API no longer need to import github.com/go-openapi/swag, but
    should import the desired sub-module(s).

**New with this release**:

* [x] type converters and pointer to value helpers now support generic types
* [x] name mangling now support pluralized initialisms (issue #46)
      Strings like "contact IDs" are now recognized as such a plural form and mangled as a linter would expect.
* [x] performance: small improvements to reduce the overhead of convert/format wrappers (see issues #110, or PR #108)
* [x] performance: name mangling utilities run ~ 10% faster (PR #115)

---

Moving forward, no additional feature will be added to the `swag` API directly.

However, child modules will continue to evolve or some new ones may be added in the future.


#### Note to contributors

The mono-repo structure comes with some unavoidable extra pains...

* Testing

> The usual `go test ./...` command, run from the root of this repo won't work any longer to test all submodules.
>
> Each module constitutes an independant unit of test. So you have to run `go test` inside each module.
> Or you may take a look at how this is achieved by CI
> [here] https://github.com/go-openapi/swag/blob/master/.github/workflows/go-test.yml).
>
> There are also some alternative tricks using `go work`, for local development, if you feel comfortable with
> go workspaces. Perhaps some day, we'll have a `go work test` to run all tests without any hack.

* Releasing

> Each module follows its own independant module versioning.
>
> So you have tags like `mangling/v0.24.0`, `fileutils/v0.24.0` etc that are used by `go mod` and `go get`
> to refer to the tagged version of each module specifically.
>
> This means we may release patches etc to each module independently.
>
> We'd like to adopt the rule that modules in this repo would only differ by a patch version
> (e.g. `v0.24.5` vs `v0.24.3`), and we'll level all modules whenever a minor version is introduced.
>
> A script in `./hack` is provided to tag all modules in one go at the same level in one go.

## Todos, suggestions and plans

All kinds of contributions are welcome.

A few ideas:

* [ ] Complete the split of dependencies to isolate easyjson from the rest
* [ ] Improve mangling utilities (improve readability, support for capitalized words,
      better word substitution for non-letter symbols...)
* [ ] Move back to this common shared pot a few of the technical features introduced by go-swagger independently
      (e.g. mangle go package names, search package with go modules support, ...)
* [ ] Apply a similar mono-repo approach to go-openapi/strfmt which suffer from similar woes: bloated API,
      imposed dependency to some database driver.

# Contributor Covenant Code of Conduct

## Our Pledge

In the interest of fostering an open and welcoming environment, we as
contributors and maintainers pledge to making participation in our project and
our community a harassment-free experience for everyone, regardless of age, body
size, disability, ethnicity, gender identity and expression, level of experience,
nationality, personal appearance, race, religion, or sexual identity and
orientation.

## Our Standards

Examples of behavior that contributes to creating a positive environment
include:

* Using welcoming and inclusive language
* Being respectful of differing viewpoints and experiences
* Gracefully accepting constructive criticism
* Focusing on what is best for the community
* Showing empathy towards other community members

Examples of unacceptable behavior by participants include:

* The use of sexualized language or imagery and unwelcome sexual attention or
advances
* Trolling, insulting/derogatory comments, and personal or political attacks
* Public or private harassment
* Publishing others' private information, such as a physical or electronic
  address, without explicit permission
* Other conduct which could reasonably be considered inappropriate in a
  professional setting

## Our Responsibilities

Project maintainers are responsible for clarifying the standards of acceptable
behavior and are expected to take appropriate and fair corrective action in
response to any instances of unacceptable behavior.

Project maintainers have the right and responsibility to remove, edit, or
reject comments, commits, code, wiki edits, issues, and other contributions
that are not aligned to this Code of Conduct, or to ban temporarily or
permanently any contributor for other behaviors that they deem inappropriate,
threatening, offensive, or harmful.

## Scope

This Code of Conduct applies both within project spaces and in public spaces
when an individual is representing the project or its community. Examples of
representing a project or community include using an official project e-mail
address, posting via an official social media account, or acting as an appointed
representative at an online or offline event. Representation of a project may be
further defined and clarified by project maintainers.

## Enforcement

Instances of abusive, harassing, or otherwise unacceptable behavior may be
reported by contacting the project team at ivan+abuse@flanders.co.nz. All
complaints will be reviewed and investigated and will result in a response that
is deemed necessary and appropriate to the circumstances. The project team is
obligated to maintain confidentiality with regard to the reporter of an incident.
Further details of specific enforcement policies may be posted separately.

Project maintainers who do not follow or enforce the Code of Conduct in good
faith may face temporary or permanent repercussions as determined by other
members of the project's leadership.

## Attribution

This Code of Conduct is adapted from the [Contributor Covenant][homepage], version 1.4,
available at [http://contributor-covenant.org/version/1/4][version]

[homepage]: http://contributor-covenant.org
[version]: http://contributor-covenant.org/version/1/4/
## Contribution Guidelines

### Pull requests are always welcome

We are always thrilled to receive pull requests, and do our best to
process them as fast as possible. Not sure if that typo is worth a pull
request? Do it! We will appreciate it.

If your pull request is not accepted on the first try, don't be
discouraged! If there's a problem with the implementation, hopefully you
received feedback on what to improve.

We're trying very hard to keep go-swagger lean and focused. We don't want it
to do everything for everybody. This means that we might decide against
incorporating a new feature. However, there might be a way to implement
that feature *on top of* go-swagger.


### Conventions

Fork the repo and make changes on your fork in a feature branch:

- If it's a bugfix branch, name it XXX-something where XXX is the number of the
  issue
- If it's a feature branch, create an enhancement issue to announce your
  intentions, and name it XXX-something where XXX is the number of the issue.

Submit unit tests for your changes.  Go has a great test framework built in; use
it! Take a look at existing tests for inspiration. Run the full test suite on
your branch before submitting a pull request.

Update the documentation when creating or modifying features. Test
your documentation changes for clarity, concision, and correctness, as
well as a clean documentation build. See ``docs/README.md`` for more
information on building the docs and how docs get released.

Write clean code. Universally formatted code promotes ease of writing, reading,
and maintenance. Always run `gofmt -s -w file.go` on each changed file before
committing your changes. Most editors have plugins that do this automatically.

Pull requests descriptions should be as clear as possible and include a
reference to all the issues that they address.

Pull requests must not contain commits from other users or branches.

Commit messages must start with a capitalized and short summary (max. 50
chars) written in the imperative, followed by an optional, more detailed
explanatory text which is separated from the summary by an empty line.

Code review comments may be added to your pull request. Discuss, then make the
suggested modifications and push additional commits to your feature branch. Be
sure to post a comment after pushing. The new commits will show up in the pull
request automatically, but the reviewers will not be notified unless you
comment.

Before the pull request is merged, make sure that you squash your commits into
logical units of work using `git rebase -i` and `git push -f`. After every
commit the test suite should be passing. Include documentation changes in the
same commit so that a revert would remove all traces of the feature or fix.

Commits that fix or close an issue should include a reference like `Closes #XXX`
or `Fixes #XXX`, which will automatically close the issue when merged.

### Sign your work

The sign-off is a simple line at the end of the explanation for the
patch, which certifies that you wrote it or otherwise have the right to
pass it on as an open-source patch.  The rules are pretty simple: if you
can certify the below (from
[developercertificate.org](http://developercertificate.org/)):

```
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
660 York Street, Suite 102,
San Francisco, CA 94110 USA

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.


Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

then you just add a line to every git commit message:

    Signed-off-by: Joe Smith <joe@gmail.com>

using your real name (sorry, no pseudonyms or anonymous contributions.)

You can add the sign off when creating the git commit via `git commit -s`.
# Strfmt [![Build Status](https://github.com/go-openapi/strfmt/actions/workflows/go-test.yml/badge.svg)](https://github.com/go-openapi/strfmt/actions?query=workflow%3A"go+test") [![codecov](https://codecov.io/gh/go-openapi/strfmt/branch/master/graph/badge.svg)](https://codecov.io/gh/go-openapi/strfmt)
[![Slack Status](https://slackin.goswagger.io/badge.svg)](https://slackin.goswagger.io)
[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/go-openapi/strfmt/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/go-openapi/strfmt?status.svg)](http://godoc.org/github.com/go-openapi/strfmt)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-openapi/strfmt)](https://goreportcard.com/report/github.com/go-openapi/strfmt)

This package exposes a registry of data types to support string formats in the go-openapi toolkit.

strfmt represents a well known string format such as credit card or email. The go toolkit for OpenAPI specifications knows how to deal with those.

## Supported data formats
go-openapi/strfmt follows the swagger 2.0 specification with the following formats
defined [here](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types).

It also provides convenient extensions to go-openapi users.

- [x] JSON-schema draft 4 formats
  - date-time
  - email
  - hostname
  - ipv4
  - ipv6
  - uri
- [x] swagger 2.0 format extensions
  - binary
  - byte (e.g. base64 encoded string)
  - date (e.g. "1970-01-01")
  - password
- [x] go-openapi custom format extensions
  - bsonobjectid (BSON objectID)
  - creditcard
  - duration (e.g. "3 weeks", "1ms")
  - hexcolor (e.g. "#FFFFFF")
  - isbn, isbn10, isbn13
  - mac (e.g "01:02:03:04:05:06")
  - rgbcolor (e.g. "rgb(100,100,100)")
  - ssn
  - uuid, uuid3, uuid4, uuid5
  - cidr (e.g. "192.0.2.1/24", "2001:db8:a0b:12f0::1/32")
  - ulid (e.g. "00000PP9HGSBSSDZ1JTEXBJ0PW", [spec](https://github.com/ulid/spec))

> NOTE: as the name stands for, this package is intended to support string formatting only.
> It does not provide validation for numerical values with swagger format extension for JSON types "number" or
> "integer" (e.g. float, double, int32...).

## Type conversion

All types defined here are stringers and may be converted to strings with `.String()`.
Note that most types defined by this package may be converted directly to string like `string(Email{})`.

`Date` and `DateTime` may be converted directly to `time.Time` like `time.Time(Time{})`.
Similarly, you can convert `Duration` to `time.Duration` as in `time.Duration(Duration{})`

## Using pointers

The `conv` subpackage provides helpers to convert the types to and from pointers, just like `go-openapi/swag` does
with primitive types.

## Format types
Types defined in strfmt expose marshaling and validation capabilities.

List of defined types:
- Base64
- CreditCard
- Date
- DateTime
- Duration
- Email
- HexColor
- Hostname
- IPv4
- IPv6
- CIDR
- ISBN
- ISBN10
- ISBN13
- MAC
- ObjectId
- Password
- RGBColor
- SSN
- URI
- UUID
- UUID3
- UUID4
- UUID5
- [ULID](https://github.com/ulid/spec)
# Contributor Covenant Code of Conduct

## Our Pledge

In the interest of fostering an open and welcoming environment, we as
contributors and maintainers pledge to making participation in our project and
our community a harassment-free experience for everyone, regardless of age, body
size, disability, ethnicity, gender identity and expression, level of experience,
nationality, personal appearance, race, religion, or sexual identity and
orientation.

## Our Standards

Examples of behavior that contributes to creating a positive environment
include:

* Using welcoming and inclusive language
* Being respectful of differing viewpoints and experiences
* Gracefully accepting constructive criticism
* Focusing on what is best for the community
* Showing empathy towards other community members

Examples of unacceptable behavior by participants include:

* The use of sexualized language or imagery and unwelcome sexual attention or
advances
* Trolling, insulting/derogatory comments, and personal or political attacks
* Public or private harassment
* Publishing others' private information, such as a physical or electronic
  address, without explicit permission
* Other conduct which could reasonably be considered inappropriate in a
  professional setting

## Our Responsibilities

Project maintainers are responsible for clarifying the standards of acceptable
behavior and are expected to take appropriate and fair corrective action in
response to any instances of unacceptable behavior.

Project maintainers have the right and responsibility to remove, edit, or
reject comments, commits, code, wiki edits, issues, and other contributions
that are not aligned to this Code of Conduct, or to ban temporarily or
permanently any contributor for other behaviors that they deem inappropriate,
threatening, offensive, or harmful.

## Scope

This Code of Conduct applies both within project spaces and in public spaces
when an individual is representing the project or its community. Examples of
representing a project or community include using an official project e-mail
address, posting via an official social media account, or acting as an appointed
representative at an online or offline event. Representation of a project may be
further defined and clarified by project maintainers.

## Enforcement

Instances of abusive, harassing, or otherwise unacceptable behavior may be
reported by contacting the project team at ivan+abuse@flanders.co.nz. All
complaints will be reviewed and investigated and will result in a response that
is deemed necessary and appropriate to the circumstances. The project team is
obligated to maintain confidentiality with regard to the reporter of an incident.
Further details of specific enforcement policies may be posted separately.

Project maintainers who do not follow or enforce the Code of Conduct in good
faith may face temporary or permanent repercussions as determined by other
members of the project's leadership.

## Attribution

This Code of Conduct is adapted from the [Contributor Covenant][homepage], version 1.4,
available at [http://contributor-covenant.org/version/1/4][version]

[homepage]: http://contributor-covenant.org
[version]: http://contributor-covenant.org/version/1/4/
## Contribution Guidelines

### Pull requests are always welcome

We are always thrilled to receive pull requests, and do our best to
process them as fast as possible. Not sure if that typo is worth a pull
request? Do it! We will appreciate it.

If your pull request is not accepted on the first try, don't be
discouraged! If there's a problem with the implementation, hopefully you
received feedback on what to improve.

We're trying very hard to keep go-swagger lean and focused. We don't want it
to do everything for everybody. This means that we might decide against
incorporating a new feature. However, there might be a way to implement
that feature *on top of* go-swagger.


### Conventions

Fork the repo and make changes on your fork in a feature branch:

- If it's a bugfix branch, name it XXX-something where XXX is the number of the
  issue
- If it's a feature branch, create an enhancement issue to announce your
  intentions, and name it XXX-something where XXX is the number of the issue.

Submit unit tests for your changes.  Go has a great test framework built in; use
it! Take a look at existing tests for inspiration. Run the full test suite on
your branch before submitting a pull request.

Update the documentation when creating or modifying features. Test
your documentation changes for clarity, concision, and correctness, as
well as a clean documentation build. See ``docs/README.md`` for more
information on building the docs and how docs get released.

Write clean code. Universally formatted code promotes ease of writing, reading,
and maintenance. Always run `gofmt -s -w file.go` on each changed file before
committing your changes. Most editors have plugins that do this automatically.

Pull requests descriptions should be as clear as possible and include a
reference to all the issues that they address.

Pull requests must not contain commits from other users or branches.

Commit messages must start with a capitalized and short summary (max. 50
chars) written in the imperative, followed by an optional, more detailed
explanatory text which is separated from the summary by an empty line.

Code review comments may be added to your pull request. Discuss, then make the
suggested modifications and push additional commits to your feature branch. Be
sure to post a comment after pushing. The new commits will show up in the pull
request automatically, but the reviewers will not be notified unless you
comment.

Before the pull request is merged, make sure that you squash your commits into
logical units of work using `git rebase -i` and `git push -f`. After every
commit the test suite should be passing. Include documentation changes in the
same commit so that a revert would remove all traces of the feature or fix.

Commits that fix or close an issue should include a reference like `Closes #XXX`
or `Fixes #XXX`, which will automatically close the issue when merged.

### Sign your work

The sign-off is a simple line at the end of the explanation for the
patch, which certifies that you wrote it or otherwise have the right to
pass it on as an open-source patch.  The rules are pretty simple: if you
can certify the below (from
[developercertificate.org](http://developercertificate.org/)):

```
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
660 York Street, Suite 102,
San Francisco, CA 94110 USA

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.


Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

then you just add a line to every git commit message:

    Signed-off-by: Joe Smith <joe@gmail.com>

using your real name (sorry, no pseudonyms or anonymous contributions.)

You can add the sign off when creating the git commit via `git commit -s`.
# blag blah
# blag blah

Ordered list:

 1. xyz     
 2. abc     
 
First list:

 - a
 - b
 - c
 
Second list:

 * A
 * B
 * C

Another Ordered list:

 0. xyz     
 0. abc     
# OpenAPI errors [![Build Status](https://github.com/go-openapi/errors/actions/workflows/go-test.yml/badge.svg)](https://github.com/go-openapi/errors/actions?query=workflow%3A"go+test") [![codecov](https://codecov.io/gh/go-openapi/errors/branch/master/graph/badge.svg)](https://codecov.io/gh/go-openapi/errors)

[![Slack Status](https://slackin.goswagger.io/badge.svg)](https://slackin.goswagger.io)
[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/go-openapi/errors/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-openapi/errors.svg)](https://pkg.go.dev/github.com/go-openapi/errors)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-openapi/errors)](https://goreportcard.com/report/github.com/go-openapi/errors)

Shared errors and error interface used throughout the various libraries found in the go-openapi toolkit.
# Contributor Covenant Code of Conduct

## Our Pledge

In the interest of fostering an open and welcoming environment, we as
contributors and maintainers pledge to making participation in our project and
our community a harassment-free experience for everyone, regardless of age, body
size, disability, ethnicity, gender identity and expression, level of experience,
nationality, personal appearance, race, religion, or sexual identity and
orientation.

## Our Standards

Examples of behavior that contributes to creating a positive environment
include:

* Using welcoming and inclusive language
* Being respectful of differing viewpoints and experiences
* Gracefully accepting constructive criticism
* Focusing on what is best for the community
* Showing empathy towards other community members

Examples of unacceptable behavior by participants include:

* The use of sexualized language or imagery and unwelcome sexual attention or
advances
* Trolling, insulting/derogatory comments, and personal or political attacks
* Public or private harassment
* Publishing others' private information, such as a physical or electronic
  address, without explicit permission
* Other conduct which could reasonably be considered inappropriate in a
  professional setting

## Our Responsibilities

Project maintainers are responsible for clarifying the standards of acceptable
behavior and are expected to take appropriate and fair corrective action in
response to any instances of unacceptable behavior.

Project maintainers have the right and responsibility to remove, edit, or
reject comments, commits, code, wiki edits, issues, and other contributions
that are not aligned to this Code of Conduct, or to ban temporarily or
permanently any contributor for other behaviors that they deem inappropriate,
threatening, offensive, or harmful.

## Scope

This Code of Conduct applies both within project spaces and in public spaces
when an individual is representing the project or its community. Examples of
representing a project or community include using an official project e-mail
address, posting via an official social media account, or acting as an appointed
representative at an online or offline event. Representation of a project may be
further defined and clarified by project maintainers.

## Enforcement

Instances of abusive, harassing, or otherwise unacceptable behavior may be
reported by contacting the project team at ivan+abuse@flanders.co.nz. All
complaints will be reviewed and investigated and will result in a response that
is deemed necessary and appropriate to the circumstances. The project team is
obligated to maintain confidentiality with regard to the reporter of an incident.
Further details of specific enforcement policies may be posted separately.

Project maintainers who do not follow or enforce the Code of Conduct in good
faith may face temporary or permanent repercussions as determined by other
members of the project's leadership.

## Attribution

This Code of Conduct is adapted from the [Contributor Covenant][homepage], version 1.4,
available at [http://contributor-covenant.org/version/1/4][version]

[homepage]: http://contributor-covenant.org
[version]: http://contributor-covenant.org/version/1/4/
## Contribution Guidelines

### Pull requests are always welcome

We are always thrilled to receive pull requests, and do our best to
process them as fast as possible. Not sure if that typo is worth a pull
request? Do it! We will appreciate it.

If your pull request is not accepted on the first try, don't be
discouraged! If there's a problem with the implementation, hopefully you
received feedback on what to improve.

We're trying very hard to keep go-swagger lean and focused. We don't want it
to do everything for everybody. This means that we might decide against
incorporating a new feature. However, there might be a way to implement
that feature *on top of* go-swagger.


### Conventions

Fork the repo and make changes on your fork in a feature branch:

- If it's a bugfix branch, name it XXX-something where XXX is the number of the
  issue
- If it's a feature branch, create an enhancement issue to announce your
  intentions, and name it XXX-something where XXX is the number of the issue.

Submit unit tests for your changes.  Go has a great test framework built in; use
it! Take a look at existing tests for inspiration. Run the full test suite on
your branch before submitting a pull request.

Update the documentation when creating or modifying features. Test
your documentation changes for clarity, concision, and correctness, as
well as a clean documentation build. See ``docs/README.md`` for more
information on building the docs and how docs get released.

Write clean code. Universally formatted code promotes ease of writing, reading,
and maintenance. Always run `gofmt -s -w file.go` on each changed file before
committing your changes. Most editors have plugins that do this automatically.

Pull requests descriptions should be as clear as possible and include a
reference to all the issues that they address.

Pull requests must not contain commits from other users or branches.

Commit messages must start with a capitalized and short summary (max. 50
chars) written in the imperative, followed by an optional, more detailed
explanatory text which is separated from the summary by an empty line.

Code review comments may be added to your pull request. Discuss, then make the
suggested modifications and push additional commits to your feature branch. Be
sure to post a comment after pushing. The new commits will show up in the pull
request automatically, but the reviewers will not be notified unless you
comment.

Before the pull request is merged, make sure that you squash your commits into
logical units of work using `git rebase -i` and `git push -f`. After every
commit the test suite should be passing. Include documentation changes in the
same commit so that a revert would remove all traces of the feature or fix.

Commits that fix or close an issue should include a reference like `Closes #XXX`
or `Fixes #XXX`, which will automatically close the issue when merged.

### Sign your work

The sign-off is a simple line at the end of the explanation for the
patch, which certifies that you wrote it or otherwise have the right to
pass it on as an open-source patch.  The rules are pretty simple: if you
can certify the below (from
[developercertificate.org](http://developercertificate.org/)):

```
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
660 York Street, Suite 102,
San Francisco, CA 94110 USA

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.


Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

then you just add a line to every git commit message:

    Signed-off-by: Joe Smith <joe@gmail.com>

using your real name (sorry, no pseudonyms or anonymous contributions.)

You can add the sign off when creating the git commit via `git commit -s`.
Evaluate:

* markdownlint
  * assessment: go
  * used by opentelemetry/opentelemetry-go
  * packaged as github action
* misspell
* spellcheck
* govulncheck
* [x] godoc-lint:
  * assessment: no go
  * too simplistic: no real value added
  * not integrated into golangci's suite of linters
  * no packaged github action
<h1>dfkjfkfj<h1>


ERGGLGKLK
==========
this is a title
# .github
Default github settings for go-openapi repositories
---
name: '❓ Question'
about: I have a question about this project
title: ''
labels: [question]
assignees: ''

---

**Question**
<!-- A clear and concise question -->

**Additional context**
<!-- Add any other context or screenshots about the feature request here. -->
---
name: '➕ Feature request'
about: Suggest an idea for this project
title: ''
labels: [enhancement]
assignees: ''

---

**Is your feature request related to a problem? Please describe.**
<!-- A clear and concise description of what the problem is. Ex. I'm always frustrated when [...] -->

**Describe the solution you'd like**
<!-- A clear and concise description of what you want to happen. -->

**Describe alternatives you've considered**
<!-- A clear and concise description of any alternative solutions or features you've considered. -->

**Additional context**
<!-- Add any other context or screenshots about the feature request here. -->
---
name: '🐞 Bug report'
about: Create a report to help us make go-openapi a great project
title: ''
labels: [bug]
assignees: ''

---

**Describe the bug**
<!-- A clear and concise description of what the bug is. -->

**Steps to Reproduce**
<!--
Ideally provide a short repro case in go.
-->

<!--
Steps to reproduce the behavior:
1. Go to '...'
2. Run this '....'
3. Check that '....'
4. See error
-->

**Expected behavior**
<!-- A clear and concise description of what you expected to happen. -->

**Actual behavior**
<!-- A clear and concise description of what is the behavior you observe. -->

**Environment**
<!-- Please paste here the output of `go env`. -->

**Additional context**
<!-- Add any other context or references about the problem here. -->
## Change type

Please select: 🆕 New feature or enhancement|🔧 Bug fix'|📃 Documentation update

## Short description
<!-- Please provide a short description of your change -->

## Fixes
<!-- 
Example:
* fixes #123

Avoid cross-repository fixes, e.g.
* fixes go-openapi/spec#123

Prefer instead:
* contributes go-openapi/spec#123

This means will be solved, but when releases and dependencies updates have been carried out
-->

## Full description
<!-- If needed, please add here more details about your implementation etc -->

<!-- Since this is a bug fix, try your best not to mix this change with extra features or potentially breaking changes -->

## Checklist
<!--- Go over all the following points, and put an `x` in all the boxes that apply. -->
<!-- If you don't qualify for all of the below check list items, please mark your PR in a draft status, so it may be discussed or reviewed with lighter requirements. -->

* [ ] I have signed all my commits with my name and email (see [DCO](https://github.com/apps/dco). **This does not require a PGP-signed commit**
* [ ] I have rebased and squashed my work, so only one commit remains
* [ ] I have added tests to cover my changes.
* [ ] I have properly enriched go doc comments in code.
* [ ] I have properly documented any breaking change.
## Change type: 📃 Documentation update

## Short description
<!-- clear and concise description of your change -->

## Fixes
<!-- 
Example:
* fixes #123

Avoid cross-repository fixes, e.g.
* fixes go-openapi/spec#123

Prefer instead:
* contributes go-openapi/spec#123

This means will be solved, but when releases and dependencies updates have been carried out
-->

## Checklist
<!--- Go over all the following points, and put an `x` in all the boxes that apply. -->
<!-- If you don't qualify for all of the below check list items, please mark your PR in a draft status, so it may be discussed or reviewed with lighter requirements. -->

* [ ] I have signed all my commits with my name and email (see [DCO](https://github.com/apps/dco). **This does not require a PGP-signed commit**
* [ ] I have rebased and squashed my work, so only one commit remains
* [ ] I have signed all my commits with my name and email (see DCO: https://github.com/apps/dco). **This does not require a PGP-signed commit**
* [ ] I have rebased and squash my work, so only one commit remains
* [ ] I have only modified comments in code or markdown files
## Change type: 🆕 New feature or enhancement

## Short description
<!-- Please provide a short description of your change -->

## Fixes
<!-- 
Example:
* fixes #123

Avoid cross-repository fixes, e.g.
* fixes go-openapi/spec#123

Prefer instead:
* contributes go-openapi/spec#123

This means will be solved, but when releases and dependencies updates have been carried out
-->

## Checklist
<!--- Go over all the following points, and put an `x` in all the boxes that apply. -->
<!-- If you don't qualify for all of the below check list items, please mark your PR in a draft status, so it may be discussed or reviewed with lighter requirements. -->

* [ ] I have signed all my commits with my name and email (see [DCO](https://github.com/apps/dco). **This does not require a PGP-signed commit**
* [ ] I have rebased and squashed my work, so only one commit remains
* [ ] I have added tests to cover my changes.
* [ ] I have properly enriched go doc comments in code.
* [ ] I have properly documented any breaking change.
## Change type: 🔧 Bug fix

## Short description
<!-- clear and concise description of your change -->

## Fixes
<!-- 
Example:
* fixes #123

Avoid cross-repository fixes, e.g.
* fixes go-openapi/spec#123

Prefer instead:
* contributes go-openapi/spec#123

This means will be solved, but when releases and dependencies updates have been carried out
-->

## Full description
<!-- If needed, please add here more details about your implementation etc -->

<!-- Since this is a bug fix, try your best not to mix this change with extra features or potentially breaking changes -->

## Checklist
<!--- Go over all the following points, and put an `x` in all the boxes that apply. -->
<!-- If you don't qualify for all of the below check list items, please mark your PR in a draft status, so it may be discussed or reviewed with lighter requirements. -->

* [ ] I have signed all my commits with my name and email (see [DCO](https://github.com/apps/dco). **This does not require a PGP-signed commit**
* [ ] I have rebased and squashed my work, so only one commit remains
* [ ] I have added tests to cover my changes.
* [ ] I have made sure that my code remains portable
* [ ] I have properly enriched go doc comments in code.
* [ ] I have properly documented any breaking change.
# OpenAPI initiative random data generator [![Build Status](https://travis-ci.org/go-openapi/stubs.svg?branch=master)](https://travis-ci.org/go-openapi/stubs) [![codecov](https://codecov.io/gh/go-openapi/stubs/branch/master/graph/badge.svg)](https://codecov.io/gh/go-openapi/stubs) [![Slack Status](https://slackin.goswagger.io/badge.svg)](https://slackin.goswagger.io)

[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/go-openapi/stubs/master/LICENSE) [![GoDoc](https://godoc.org/github.com/go-openapi/stubs?status.svg)](http://godoc.org/github.com/go-openapi/stubs) 


A library to generate random data for a swagger specification.
This is a building block for generating stubs for your API as well as tests.
# Contributor Covenant Code of Conduct

## Our Pledge

In the interest of fostering an open and welcoming environment, we as
contributors and maintainers pledge to making participation in our project and
our community a harassment-free experience for everyone, regardless of age, body
size, disability, ethnicity, gender identity and expression, level of experience,
nationality, personal appearance, race, religion, or sexual identity and
orientation.

## Our Standards

Examples of behavior that contributes to creating a positive environment
include:

* Using welcoming and inclusive language
* Being respectful of differing viewpoints and experiences
* Gracefully accepting constructive criticism
* Focusing on what is best for the community
* Showing empathy towards other community members

Examples of unacceptable behavior by participants include:

* The use of sexualized language or imagery and unwelcome sexual attention or
advances
* Trolling, insulting/derogatory comments, and personal or political attacks
* Public or private harassment
* Publishing others' private information, such as a physical or electronic
  address, without explicit permission
* Other conduct which could reasonably be considered inappropriate in a
  professional setting

## Our Responsibilities

Project maintainers are responsible for clarifying the standards of acceptable
behavior and are expected to take appropriate and fair corrective action in
response to any instances of unacceptable behavior.

Project maintainers have the right and responsibility to remove, edit, or
reject comments, commits, code, wiki edits, issues, and other contributions
that are not aligned to this Code of Conduct, or to ban temporarily or
permanently any contributor for other behaviors that they deem inappropriate,
threatening, offensive, or harmful.

## Scope

This Code of Conduct applies both within project spaces and in public spaces
when an individual is representing the project or its community. Examples of
representing a project or community include using an official project e-mail
address, posting via an official social media account, or acting as an appointed
representative at an online or offline event. Representation of a project may be
further defined and clarified by project maintainers.

## Enforcement

Instances of abusive, harassing, or otherwise unacceptable behavior may be
reported by contacting the project team at ivan+abuse@flanders.co.nz. All
complaints will be reviewed and investigated and will result in a response that
is deemed necessary and appropriate to the circumstances. The project team is
obligated to maintain confidentiality with regard to the reporter of an incident.
Further details of specific enforcement policies may be posted separately.

Project maintainers who do not follow or enforce the Code of Conduct in good
faith may face temporary or permanent repercussions as determined by other
members of the project's leadership.

## Attribution

This Code of Conduct is adapted from the [Contributor Covenant][homepage], version 1.4,
available at [http://contributor-covenant.org/version/1/4][version]

[homepage]: http://contributor-covenant.org
[version]: http://contributor-covenant.org/version/1/4/
# Stubs design

The goal of this library is to be able to generate somewhat good looking random data for structures defined in an openapi or jsonschema document.
These can be used to provide an api with stubs so that you can verify valid and invalid request/responses without an actually meaningful implementation.
Some areas where this is useful are collaboration between different teams so that all teams can do work without having to wait on the delivery of the completed implementation of the API.
A second area where this is obviously useful is for doing automated tests and filling up datastructures for use in those tests.


In the openapi 2.0 specification there are 2 families of types that can be used as targets for a random data generator.

1. Path Item related types
   1. Non-Schema Parameters
   2. Response headers
   3. Collection Items on parameters or headers  
2. Schemas

These 2 families warrant slightly different strategies for getting the necessary parameters for a datagenerator.

## Functionality

To generate meaningful random data the library looks at the schema and based on properties for that schema it picks an appropriate random value generator.
This means that when it is generating a string for an object property that represents a first name it should pick a data generator that generates first names.
Similarly should a property be called city it should generate a city name.

It's possible that the validity of a particular property is dependent on the value of another property, for example a creation date is typically never larger than a modified date.

A data generator for a schema should be able to generate valid and invalid data, ideally there is control for specifying the value is invalid for a particular validation. This will help in tests to validate error codes

## How does it work

The main components in this library are a registry of datagenerators so they are addressable by keys, aliases for those keys to aid with inferring which datagenerator to use.
And of course a value generator, which will generate either a valid or an invalid value.

In the openapi API document and in a json schema document there is a vendor extension that can be used to customize the generation process.

### Value generator function

There are 2 types of data generators but they both have the same function signature.
There are value generators and there are composite value generators.

* A value generator generates a single value for a simple type.
* A composite generator generates a value that is built out of one or more generators eg. object schemas with properties.

The signature for the value generator is:

```go
type ValueGenerator func(GeneratorOpts) (interface{}, error)
```

A value generator is the innermost component in the library and the arguments to the function are used to configure the generator.
The generation process for a generator is configured through a GeneratorOpts interface.

### GeneratorOpts interface

The generator options describe the type and potentially the format for the value that needs to be generated.
In addition to the type information it also captures the field name or definition name of the value that needs to be generated.

### StubMode

The stubmode bitmask allows for configuring which validations should fail for a given value generator.

### Generator

The generator is the main entry point for the library and its Generate method is what will generate the random value for the descriptor.
# Use cases for stubs

The stubs package is intended to support the following use cases:

1. Generate unit test cases for code generated by go-swagger, especially validation code
2. Generate test cases for API servers and clients
3. Generate realistic examples

## Supported swagger 2.0 constructs

- parameter 
- header 
- response
- schema, including support for:
  - [] AllOf 
  - [] Not
  - [] additionalProperties
  - [] patternProperties 
  - [] additionalItems

Formated strings and numbers support go-openapi/strfmt formats, including:
  - [] binary (strfmt.Base64)
  - [] bsonobjectid

The "file" datatype is not supported at the moment.

The generation assumes the provide swagger spec is valid. 

Warnings regarding test case (2): 
- stubs is not concerned about security definitions and credentials
- stubs generates values independently for each parameter, or response: it does not know about consistent sets of values beyond validation constraints, let alone 
  proper chaining of request with responses.

## Sample applications for stubs

- process a spec to include examples 
- generate fixtures for codegen unit testing
- generate fixtures for API server benchmarking
- ...

# Reproducible stubs 

Generated stubs may be regenerated using the same initial seed. 

There are two seeding modes to run the stubs generator: 

- AutoSeed: reseeds itself based on current timestamp 
- SetSeed: generation occurs with a single seed, provided by the caller

The package uses a default seeder to carry on this task. Alternate seeders might be provided.

This also allows for unit testing this very package...

# Setting generation goals 

Goals select the general use case and allow for fine tuning of the sample generation.

We use GenerateMode and StubMode flags to define goals.

Stubmode is used to tune individually every single generated value.

GenerateMode sets general directives and may define StubMode autonomously according to the generation goal.

## GenerateMode

target:
- unitTest: any value value may be produced, not necessarily a realistic definevalue. 
  By default in the UnitTest mode, non-formated strings are loren ipsums, numbers, dates and times may be very small or very large
- examples: the generator attempts to figure out realistic values to be used as examples

GenerateMode accepts arguments to further specify the expected generation: 

args:
- valid:    all generated values are valid. This is the default.
  - count: number of samples to be generated. May be zero to get only default, examples, edge cases or validation checks. Default is 1.
- invalid:  all generated values are invalid, with failures chosen at random.
  - count: number of samples to be generated. Default is 1.
- skipTilting: tilting rules are omitted (only applies to invalid values). Default is true in unit test mode and false in example mode.
- skipFuzzying: fuzzying rules are omitted
- withTilting: add tilted invalid values, with failures chosen at random. This forces a tilting strategy on a new set of values.
  - count: number of tilted samples to be generated. Default is 1.
- withValidationCheck: for each validation, an invalid value against this validation is generated (e.g. if min and max are specified, we have one value below min and one value above max)
   - count: this sets a limit on the count of generated values. If you have n validations, that makes n generated values. With the limit set, validations are checked width-first (outermost validations come first)
- withAllValidationChecks: all combinations of failed validations are generated. This includes one valid value and single failure from withValidation.
   - count: this sets a limit on the count of generated values. If you have n validations, that makes 2^n generated values, which may raise quickly when nesting schema structures. With the limit set, validations are checked width-first (outermost validations come first)
- withDefault: add default value to sample values (not included in "count"). This simulates a user providing exactly the same value as the default. Since defaults may be set at different levels of the structure, a new value is generated for each applicable default.
- withExample: add example value to sample values (not included in "count"), whenever applicable. Examples are assumed valid, but may be invalid.
   - tiltDefault: constructs invalid values by tilting default
   - tiltExample: constructs invalid values by tilting examples  
   - count: number of title
- withEdgeCase:
    - standard: standard preconfigured edge cases (e.g. boundary value for min constraint with boundary included, Feb 29th dates for date types, ...)
    - value: a customizable value added as an edge case. May be any value, array or object. Severa values may be specified.
- max: an upper limit on number generation (floats, ints, ...). By default, the full range of the data type is used.
- min: a lower limit on number generation.
- precision: rounding specification for floats
- multipleOf: a multiple specification for numbers
- length:
  - on string values: an upper limit on string size, to avoid long texts being unwittingly generated
  - on array values: an upper limit on the number of generated items
  - on additionalProperties: an upper limit on the number of generated additionalProperties
  - on patternProperties: an upper limit on the number of generated patternProperties
- words: a limit on the number of words in generated texts, paragraphs and sentences
- lang: specifies the target language(s) for strings and text. 
  - [lang]
  - all: special value to tell the generator to pick text at random from all supported languages

This sets an overall target for all generations. For every single spec object, you may hint further and override global settings on a case by case basis.

# Hinting

The `x-datagen` extension may be used in the swagger spec to hint the stubs generator.

Examples:
``` 
parameters:
    phone:
        in: query 
        required: true
        type: string
        x-datagen:
          name: mobile
```
will generate a mobile phone number for query parameter "phone".

``` 
parameters:
    price:
        in: query 
        required: true
        type: string
        x-datagen:
          name: float
```
will generate float numbers for the price parameter. 
It might be not realistic to have 3.1344344E35 as a price, but it is a valid value...

`x-datagen` accepts args to further hint the generator. 
All Args accepted by GenerateMode or StubMode may be used here to locally override generation options. 

Example:
``` 
parameters:
    vat:
        in: query 
        required: true
        type: number
        x-datagen:
          args:
            min: 100
            max: 10000
            
```
will generate floats between 100 and 10000. 
Again, it might not be realistic to have prices such as 503.14561434566...

Some generators are already configured to suit some targeted usage.

Example: 
``` 
parameters:
    price:
        in: query 
        required: true
        type: string
        x-datagen:
          name: small-amounts
```
and you get more realistic values like 3.23, 104.30...

## x-datagen args

- name: overrides the generator's name

Supported args, not already mentioned in the GenerateMode section:
- args:

# Fuzzying

The generation takes basic decisions based on type and available validation constraints.
Obviously, this does not work well when the goal is to generate realistic samples for your data.
Of course, everyhing may be fine-tuned by using the x-datagen extension.
But this is tedious work on large specs. 

We introduce fuzzying as a way to carry on some of the guesswork when deciding which kind of fake value would best match expectations. 
Fuzzying relies on names, titles and description to look for recognizable keywords, then elects an adequate random generator.

Examples for the above samples:
- with name "price", the fuzzying guesses that it is not some arbitrary float to be generated, but a decimal between 0 and 1,000,000 (max overrides this limit).
- with a description such as "The item price expressed in USD. Items are sold at a low price", fuzzying rules should guess that we want to generate small-amount values (so between 0 and 1,000).

Fuzzying also works for non-strings data:
- a date field named "birthday" would hint toward the birthdate value generator (hence dates are not in the future and only in a recent past)
- a datetime named "timestamp" or described as "The update date" would generate a recent datetime value. Further, edge cases with this guess would add 01/01/1970 to the generated value

# Tilting

Tilting is a strategy to generate invalid values, by tilting a valid value just enough so it becomes invalid.
This generates interesting test cases since values are supposed to be "almost" valid but actually are not.

Tilting strategies depend on validations. 

Examples:
- tilting on max will generate a value at or slightly above the lowest invalid value 
- tilting on regexp will change the smallest number of characters in a string so the regexp no more matches
- tilting on enum will slightly modify a valid value and make it invalid
- ...

Options skipTilting and withTilting tune the generator behavior regarding tilting.
# gojsonpointer [![Build Status](https://github.com/go-openapi/jsonpointer/actions/workflows/go-test.yml/badge.svg)](https://github.com/go-openapi/jsonpointer/actions?query=workflow%3A"go+test") [![codecov](https://codecov.io/gh/go-openapi/jsonpointer/branch/master/graph/badge.svg)](https://codecov.io/gh/go-openapi/jsonpointer)

[![Slack Status](https://slackin.goswagger.io/badge.svg)](https://slackin.goswagger.io)
[![license](http://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/go-openapi/jsonpointer/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-openapi/jsonpointer.svg)](https://pkg.go.dev/github.com/go-openapi/jsonpointer)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-openapi/jsonpointer)](https://goreportcard.com/report/github.com/go-openapi/jsonpointer)

An implementation of JSON Pointer - Go language

## Status
Completed YES

Tested YES

## References
http://tools.ietf.org/html/draft-ietf-appsawg-json-pointer-07

### Note
The 4.Evaluation part of the previous reference, starting with 'If the currently referenced value is a JSON array, the reference token MUST contain either...' is not implemented.
# Contributor Covenant Code of Conduct

## Our Pledge

In the interest of fostering an open and welcoming environment, we as
contributors and maintainers pledge to making participation in our project and
our community a harassment-free experience for everyone, regardless of age, body
size, disability, ethnicity, gender identity and expression, level of experience,
nationality, personal appearance, race, religion, or sexual identity and
orientation.

## Our Standards

Examples of behavior that contributes to creating a positive environment
include:

* Using welcoming and inclusive language
* Being respectful of differing viewpoints and experiences
* Gracefully accepting constructive criticism
* Focusing on what is best for the community
* Showing empathy towards other community members

Examples of unacceptable behavior by participants include:

* The use of sexualized language or imagery and unwelcome sexual attention or
advances
* Trolling, insulting/derogatory comments, and personal or political attacks
* Public or private harassment
* Publishing others' private information, such as a physical or electronic
  address, without explicit permission
* Other conduct which could reasonably be considered inappropriate in a
  professional setting

## Our Responsibilities

Project maintainers are responsible for clarifying the standards of acceptable
behavior and are expected to take appropriate and fair corrective action in
response to any instances of unacceptable behavior.

Project maintainers have the right and responsibility to remove, edit, or
reject comments, commits, code, wiki edits, issues, and other contributions
that are not aligned to this Code of Conduct, or to ban temporarily or
permanently any contributor for other behaviors that they deem inappropriate,
threatening, offensive, or harmful.

## Scope

This Code of Conduct applies both within project spaces and in public spaces
when an individual is representing the project or its community. Examples of
representing a project or community include using an official project e-mail
address, posting via an official social media account, or acting as an appointed
representative at an online or offline event. Representation of a project may be
further defined and clarified by project maintainers.

## Enforcement

Instances of abusive, harassing, or otherwise unacceptable behavior may be
reported by contacting the project team at ivan+abuse@flanders.co.nz. All
complaints will be reviewed and investigated and will result in a response that
is deemed necessary and appropriate to the circumstances. The project team is
obligated to maintain confidentiality with regard to the reporter of an incident.
Further details of specific enforcement policies may be posted separately.

Project maintainers who do not follow or enforce the Code of Conduct in good
faith may face temporary or permanent repercussions as determined by other
members of the project's leadership.

## Attribution

This Code of Conduct is adapted from the [Contributor Covenant][homepage], version 1.4,
available at [http://contributor-covenant.org/version/1/4][version]

[homepage]: http://contributor-covenant.org
[version]: http://contributor-covenant.org/version/1/4/
## Contribution Guidelines

### Pull requests are always welcome

We are always thrilled to receive pull requests, and do our best to
process them as fast as possible. Not sure if that typo is worth a pull
request? Do it! We will appreciate it.

If your pull request is not accepted on the first try, don't be
discouraged! If there's a problem with the implementation, hopefully you
received feedback on what to improve.

We're trying very hard to keep go-swagger lean and focused. We don't want it
to do everything for everybody. This means that we might decide against
incorporating a new feature. However, there might be a way to implement
that feature *on top of* go-swagger.


### Conventions

Fork the repo and make changes on your fork in a feature branch:

- If it's a bugfix branch, name it XXX-something where XXX is the number of the
  issue
- If it's a feature branch, create an enhancement issue to announce your
  intentions, and name it XXX-something where XXX is the number of the issue.

Submit unit tests for your changes.  Go has a great test framework built in; use
it! Take a look at existing tests for inspiration. Run the full test suite on
your branch before submitting a pull request.

Update the documentation when creating or modifying features. Test
your documentation changes for clarity, concision, and correctness, as
well as a clean documentation build. See ``docs/README.md`` for more
information on building the docs and how docs get released.

Write clean code. Universally formatted code promotes ease of writing, reading,
and maintenance. Always run `gofmt -s -w file.go` on each changed file before
committing your changes. Most editors have plugins that do this automatically.

Pull requests descriptions should be as clear as possible and include a
reference to all the issues that they address.

Pull requests must not contain commits from other users or branches.

Commit messages must start with a capitalized and short summary (max. 50
chars) written in the imperative, followed by an optional, more detailed
explanatory text which is separated from the summary by an empty line.

Code review comments may be added to your pull request. Discuss, then make the
suggested modifications and push additional commits to your feature branch. Be
sure to post a comment after pushing. The new commits will show up in the pull
request automatically, but the reviewers will not be notified unless you
comment.

Before the pull request is merged, make sure that you squash your commits into
logical units of work using `git rebase -i` and `git push -f`. After every
commit the test suite should be passing. Include documentation changes in the
same commit so that a revert would remove all traces of the feature or fix.

Commits that fix or close an issue should include a reference like `Closes #XXX`
or `Fixes #XXX`, which will automatically close the issue when merged.

### Sign your work

The sign-off is a simple line at the end of the explanation for the
patch, which certifies that you wrote it or otherwise have the right to
pass it on as an open-source patch.  The rules are pretty simple: if you
can certify the below (from
[developercertificate.org](http://developercertificate.org/)):

```
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
660 York Street, Suite 102,
San Francisco, CA 94110 USA

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.


Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```

then you just add a line to every git commit message:

    Signed-off-by: Joe Smith <joe@gmail.com>

using your real name (sorry, no pseudonyms or anonymous contributions.)

You can add the sign off when creating the git commit via `git commit -s`.
