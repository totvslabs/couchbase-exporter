# Contributing

## Setup your machine

`couchbase-exporter` is written in [Go](https://golang.org/), version 1.11 and beyond.

Prerequisites:

- `make`
- [Go 1.11+](https://golang.org/doc/install)
- `jsonnet` (to generate grafana dashboards)

Clone the project anywhere:

```sh
$ git clone git@github.com:totvslabs/couchbase-exporter.git
```

Install the build and lint dependencies:

```console
$ make setup
```

A good way of making sure everything is all right is running the test suite:

```console
$ make test
```

## Test your change

You can create a branch for your changes and try to build from the source as you go:

```console
$ make build
```

When you are satisfied with the changes, we suggest you run:

```console
$ make ci
```

Which runs all the linters and tests.
