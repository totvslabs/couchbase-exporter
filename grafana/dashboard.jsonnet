local grafana = import 'grafonnet/grafonnet/grafana.libsonnet';
local dashboard = grafana.dashboard;
local row = grafana.row;
local singlestat = grafana.singlestat;
local prometheus = grafana.prometheus;

dashboard.new(
	'Couchbase5',
	refresh='30s',
	time_from='now-1h',
	tags=['couchbase'],
	editable=true,
)
.addTemplate(
	grafana.template.datasource(
		'PROMETHEUS_DS',
		'prometheus',
		'Prometheus',
		hide='label',
	)
)
.addTemplate(
	grafana.template.new(
		'instance',
		'$PROMETHEUS_DS',
		'label_values(couchbase_bucket_basicstats_dataused, instance)',
		label='Instance',
		refresh='load',
	)
)
.addTemplate(
	grafana.template.new(
		'bucket',
		'$PROMETHEUS_DS',
		'label_values(couchbase_bucket_basicstats_dataused{instance="$instance"}, bucket)',
		label='Bucket',
		refresh='load',
		multi=true,
		includeAll=true,
	)
)
.addRow(
	row.new(
		title='General'
	)
	.addPanel(
		singlestat.new(
			'Bucket RAM Usage',
			datasource='Prometheus',
			span=2,
			valueName='current',
			gaugeShow=true,
			gaugeThresholdMarkers=true,
			thresholds='70,90',
			format='percent',
		)
		.addTarget(
			prometheus.target(
				'avg(100 * (sum by (bucket) (couchbase_bucket_basicstats_memused{bucket=~"$bucket",instance=~"$instance"})) / sum by (bucket) (couchbase_bucket_stats_ep_max_size{bucket=~"$bucket",instance=~"$instance"}))',
			)
		)
	)
	.addPanel(
		singlestat.new(
			'Bucket RAM Size',
			format='bytes',
			datasource='Prometheus',
			span=2,
			valueName='current',
			sparklineFull=true,
			sparklineShow=true,
		)
		.addTarget(
			prometheus.target(
				'avg(sum(couchbase_bucket_stats_ep_max_size{bucket=~"$bucket",instance=~"$instance"}))',
			)
		)
	)
)
.addRow(
	row.new(
		title='Compacting'
	)
)
.addRow(
	row.new(
		title='Rebalance'
	)
)
.addRow(
	row.new(
		title='Memory'
	)
)
.addRow(
	row.new(
		title='Queries'
	)
)
.addRow(
	row.new(
		title='XDCR'
	)
)
