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
