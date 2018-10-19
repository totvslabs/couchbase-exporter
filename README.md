# couchbase-exporter

A prometheus couchbase exporter!

All others I found seem to be abandoned and/or very incomplete. This is an attempt to
have all metrics exported, including task-related metrics!

## Usage

```console
$ couchbase-exporter --couchbase.username adm --couchbase.password secret
```

> check `couchbase-exporter --help` for more options!

## Roadmap

- [x] export task metrics
- [x] export bucket metrics
- [ ] export node metrics
- [x] export cluster metrics
- [ ] provide grafana dashboards (maybe use jsonnet (https://github.com/grafana/grafonnet-lib)?)
- [ ] check other TODOs
- [ ] improve metric names (add `_bytes`, `_seconds`, `_total`, etc)
~~- [ ] do the `if cb5` and `if cb4` accordingly~~: we will support CB5 only.
