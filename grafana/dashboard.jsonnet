local grafana = import 'grafonnet/grafonnet/grafana.libsonnet';
local dashboard = grafana.dashboard;
local row = grafana.row;
local singlestat = grafana.singlestat;
local graphPanel = grafana.graphPanel;
local prometheus = grafana.prometheus;

dashboard.new(
	'Couchbase5',
	refresh='10s',
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
			format='decbytes',
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
	.addPanel(
		singlestat.new(
			'Server Count',
			format='none',
			datasource='Prometheus',
			span=2,
			valueName='current',
			sparklineFull=true,
			sparklineShow=true,
		)
		.addTarget(
			prometheus.target(
				'count(couchbase_node_interestingstats_ops{instance=~"$instance"})',
			)
		)
	)
	.addPanel(
		singlestat.new(
			'Rebalance Progress',
			format='precent',
			datasource='Prometheus',
			span=2,
			valueName='current',
			sparklineFull=true,
			sparklineShow=true,
		)
		.addTarget(
			prometheus.target(
				'couchbase_task_rebalance_progress{instance=~"$instance"}',
			)
		)
	)
	.addPanel(
		singlestat.new(
			'Bucket QPS',
			format='none',
			datasource='Prometheus',
			span=2,
			valueName='current',
			sparklineFull=true,
			sparklineShow=true,
		)
		.addTarget(
			prometheus.target(
				'avg(sum by (bucket) (couchbase_bucket_stats_cmd_set{bucket=~"$bucket",instance=~"$instance"}) + sum by (bucket) (couchbase_bucket_stats_cmd_get{bucket=~"$bucket",instance=~"$instance"}))',
			)
		)
	)
)
.addRow(
	row.new(
		title='Details'
	)
	.addPanel(
		graphPanel.new(
			'QPS',
			span=12,
			legend_alignAsTable=true,
			legend_rightSide=true,
			legend_values=true,
			legend_current=true,
			legend_sort='current',
			legend_sortDesc=true,
			min=0,
		)
		.addTarget(
			prometheus.target(
				'sum by (bucket) (couchbase_bucket_stats_cmd_set{bucket=~"$bucket",instance=~"$instance"}) + sum by (bucket) (couchbase_bucket_stats_cmd_get{bucket=~"$bucket",instance=~"$instance"})',
				legendFormat='{{ bucket }}',
			)
		)
	)
	.addPanel(
		graphPanel.new(
			'Cache Miss Rate',
			span=12,
			legend_alignAsTable=true,
			legend_rightSide=true,
			legend_values=true,
			legend_current=true,
			legend_sort='current',
			legend_sortDesc=true,
			format='percent',
			min=0,
			max=100,
		)
		.addTarget(
			prometheus.target(
				'couchbase_bucket_stats_ep_cache_miss_rate{bucket=~"$bucket",instance=~"$instance"}',
				legendFormat='{{ bucket }}',
			)
		)
	)
)
.addRow(
	row.new(
		title='Compacting'
	)
	.addPanel(
		graphPanel.new(
			'Fragmentation',
			span=6,
			legend_alignAsTable=true,
			legend_rightSide=true,
			legend_values=true,
			legend_current=true,
			legend_sort='current',
			legend_sortDesc=true,
			format='percent',
			min=0,
			max=100,
		)
		.addTarget(
			prometheus.target(
				'couchbase_bucket_couch_docs_fragmentation{instance=~"$instance", bucket=~"$bucket"}',
				legendFormat='{{ bucket }}',
			)
		)
	)
	.addPanel(
		graphPanel.new(
			'Compaction Progress',
			span=6,
			legend_alignAsTable=true,
			legend_rightSide=true,
			legend_values=true,
			legend_current=true,
			legend_sort='current',
			legend_sortDesc=true,
			format='percent',
			min=0,
			max=100,
		)
		.addTarget(
			prometheus.target(
				'couchbase_task_compacting_progress{instance=~"$instance", bucket=~"$bucket"}',
				legendFormat='{{ bucket }}',
			)
		)
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
	.addPanel(
		graphPanel.new(
			'Memory Used',
			span=12,
			legend_alignAsTable=true,
			legend_rightSide=true,
			legend_values=true,
			legend_current=true,
			legend_sort='current',
			legend_sortDesc=true,
			format='decbytes',
			min=0,
		)
		.addTarget(
			prometheus.target(
				'couchbase_bucket_basicstats_memused{bucket=~"$bucket",instance=~"$instance"}',
				legendFormat='{{ bucket }}.Usage',
			)
		)
		.addTarget(
			prometheus.target(
				'couchbase_bucket_stats_ep_mem_high_wat{bucket=~"$bucket",instance=~"$instance"}',
				legendFormat='{{ bucket }}.HighWatermark',
			)
		)
		.addTarget(
			prometheus.target(
				'couchbase_bucket_stats_ep_mem_low_wat{bucket=~"$bucket",instance=~"$instance"}',
				legendFormat='{{ bucket }}.LowWatermark',
			)
		)
	)
	.addPanel(
		graphPanel.new(
			'Items Count',
			span=6,
			legend_alignAsTable=true,
			legend_rightSide=true,
			legend_values=true,
			legend_current=true,
			legend_sort='current',
			legend_sortDesc=true,
			min=0,
		)
		.addTarget(
			prometheus.target(
				'couchbase_bucket_stats_curr_items{bucket=~"$bucket",instance=~"$instance"}',
				legendFormat='{{ bucket }}',
			)
		)
	)
	.addPanel(
		graphPanel.new(
			'Hard Out of Memory Errors',
			span=6,
			legend_alignAsTable=true,
			legend_rightSide=true,
			legend_values=true,
			legend_current=true,
			legend_sort='current',
			legend_sortDesc=true,
			min=0,
		)
		.addTarget(
			prometheus.target(
				'rate(couchbase_bucket_stats_ep_oom_errors{instance="$instance", bucket=~"$bucket"}[5m])',
				legendFormat='{{ bucket }}',
			)
		)
	)
)
.addRow(
	row.new(
		title='Queries'
	)
	.addPanel(
		graphPanel.new(
			'Gets / Sets',
			span=12,
			legend_alignAsTable=true,
			legend_rightSide=true,
			legend_values=true,
			legend_current=true,
			legend_sort='current',
			legend_sortDesc=true,
		)
		.addSeriesOverride(
			{
				"alias": "/Get.*/",
          		"transform": "negative-Y"
            }
		)
		.addTarget(
			prometheus.target(
				'couchbase_bucket_stats_cmd_set{bucket=~"$bucket",instance=~"$instance"}',
				legendFormat='Sets on {{ bucket }}',
			)
		)
		.addTarget(
			prometheus.target(
				'couchbase_bucket_stats_cmd_get{bucket=~"$bucket",instance=~"$instance"}',
				legendFormat='Gets on {{ bucket }}',
			)
		)
	)
	.addPanel(
		graphPanel.new(
			'Evictions',
			span=6,
			legend_alignAsTable=true,
			legend_rightSide=true,
			legend_values=true,
			legend_current=true,
			legend_sort='current',
			legend_sortDesc=true,
			min=0,
		)
		.addTarget(
			prometheus.target(
				'couchbase_bucket_stats_evictions{bucket=~"$bucket",instance=~"$instance"}',
				legendFormat='{{ bucket }}',
			)
		)
	)
	.addPanel(
		graphPanel.new(
			'Miss Rate',
			span=6,
			legend_alignAsTable=true,
			legend_rightSide=true,
			legend_values=true,
			legend_current=true,
			legend_sort='current',
			legend_sortDesc=true,
			format='percent',
			min=0,
			max=100,
		)
		.addTarget(
			prometheus.target(
				'couchbase_bucket_stats_ep_cache_miss_rate{bucket=~"$bucket",instance=~"$instance"}',
				legendFormat='{{ bucket }}',
			)
		)
	)
)
.addRow(
	row.new(
		title='XDCR'
	)
)
