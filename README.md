# couchbase-exporter

A prometheus couchbase exporter!

All others I found seem to be abandoned and/or very incomplete. This is an attempt to
have all metrics exported, including task-related metrics!

## Goal

Innitially, have the same metrics as [our old exporter](https://github.com/brunopsoares/prometheus_couchbase_exporter),
to facilitate the migration.

Once that's done, we can better document all metrics, eventually improve their
naming (like adding `_total`, `_seconds`, `_byte` prefixes), and finally,
provide example alerting rules and grafana dashboards, so it's easier to
proper monitor a couchbase cluster.

It's worth saying that we will only support Couchbase 5+ for now.

## Installation

### Building from source
This requires the user to have `go` installed on their system, preferably version1.13 and onwards.  
1. Run `git clone https://github.com/totvslabs/couchbase-exporter.git`.
2. Change directory by running `cd couchbase-exporter`.
3. Create a go executable by running `go build -o couchbase-exporter main.go`. You might have to prefix `sudo` if you encounter permission issues.
4. Start the exporter using `./couchbase-exporter` followed by specific flags and their values.

### Using the Releases

1. Pick from the couchbase-exporter releases available [here](https://github.com/totvslabs/couchbase-exporter/releases) depending on whether you're a macOS or Linux user.
2. Download the release binary and extract it.
3. The exporter is now ready to be used!

## Usage

```console
$ export COUCHBASE_PASSWORD=secret
$ couchbase-exporter --couchbase.username adm
```

> check `couchbase-exporter --help` for more options!

## What's included

- the exporter itself
- a grafana dashboard
- an example of alerting rules

## Roadmap

- [x] export task metrics
- [x] export bucket metrics
- [x] export node metrics
- [x] export cluster metrics
- [x] provide alerting rules examples
- [x] provide grafana dashboards (maybe use jsonnet (https://github.com/grafana/grafonnet-lib)?)
- [ ] check other TODOs
- [ ] improve metric names (add `_bytes`, `_seconds`, `_total`, etc)
- [ ] add some sort of tests

<!--
TODO: when the needed PRs get merged into grafonnet, stop using my fork
 -->
