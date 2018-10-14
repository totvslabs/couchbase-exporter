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
- [ ] provide grafana dashboards (maybe use jsonnet?)
- [ ] check other TODOs
