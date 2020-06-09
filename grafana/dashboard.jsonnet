local grafana = import 'grafonnet/grafonnet/grafana.libsonnet';
local dashboard = grafana.dashboard;
local row = grafana.row;
local singlestat = grafana.singlestat;
local graphPanel = grafana.graphPanel;
local prometheus = grafana.prometheus;

dashboard.new(
  'Couchbase5',
  description='Couchbase 5+ dashboard, created for totvslabs/couchbase-exporter',
  refresh='10s',
  time_from='now-1h',
  tags=['couchbase'],
  editable=false,
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
    'label_values(couchbase_bucket_basicstats_dataused_bytes, instance)',
    label='Instance',
    refresh='load',
  )
)
.addTemplate(
  grafana.template.new(
    'bucket',
    '$PROMETHEUS_DS',
    'label_values(couchbase_bucket_basicstats_dataused_bytes{instance="$instance"}, bucket)',
    label='Bucket',
    refresh='load',
    multi=true,
    includeAll=true,
  )
)
.addRow(
  row.new(
    title='General',
    collapse=false,
  )
  .addPanel(
    singlestat.new(
      'Balanced',
      span=2,
      datasource='$PROMETHEUS_DS',
      valueName='current',
      colorBackground=true,
      valueFontSize='200%',
      thresholds='0,1',
      colors=[
        '#d44a3a',
        '#d44a3a',
        '#299c46',
      ],
      valueMaps=[
        {
          value: 'null',
          op: '=',
          text: 'N/A',
        },
        {
          value: '1',
          op: '=',
          text: 'YES',
        },
        {
          value: '0',
          op: '=',
          text: 'NO',
        },
      ]
    )
    .addTarget(
      prometheus.target(
        'couchbase_cluster_balanced{instance=~"$instance"}',
      )
    )
  )
  .addPanel(
    singlestat.new(
      'Rebalance Progress',
      format='percent',
      span=2,
      datasource='$PROMETHEUS_DS',
      decimals=1,
      valueName='current',
      gaugeShow=true,
      gaugeThresholdMarkers=false,
      valueMaps=[
        {
          value: 'null',
          op: '=',
          text: 'N/A',
        },
        {
          value: '0',
          op: '=',
          text: 'OK',
        },
      ]
    )
    .addTarget(
      prometheus.target(
        'couchbase_task_rebalance_progress{instance=~"$instance"}',
      )
    )
  )
  .addPanel(
    singlestat.new(
      'Compacting Progress',
      format='percent',
      span=2,
      datasource='$PROMETHEUS_DS',
      decimals=1,
      valueName='current',
      gaugeShow=true,
      gaugeThresholdMarkers=false,
      valueMaps=[
        {
          value: 'null',
          op: '=',
          text: 'N/A',
        },
        {
          value: '0',
          op: '=',
          text: 'OK',
        },
      ]
    )
    .addTarget(
      prometheus.target(
        'avg(couchbase_task_compacting_progress{instance=~"$instance"})',
      )
    )
  )
  .addPanel(
    singlestat.new(
      'Bucket RAM Usage',
      span=2,
      datasource='$PROMETHEUS_DS',
      decimals=1,
      valueName='current',
      gaugeShow=true,
      gaugeThresholdMarkers=true,
      thresholds='70,90',
      format='percent',
    )
    .addTarget(
      prometheus.target(
        'avg(100 * (sum by (bucket) (couchbase_bucket_basicstats_memused_bytes{bucket=~"$bucket",instance=~"$instance"})) / sum by (bucket) (couchbase_bucket_stats_ep_max_size_bytes{bucket=~"$bucket",instance=~"$instance"}))',
      )
    )
  )
  .addPanel(
    singlestat.new(
      'Server Count',
      format='none',
      span=2,
      datasource='$PROMETHEUS_DS',
      valueName='current',
      valueFontSize='200%',
      sparklineFull=true,
      sparklineShow=true,
    )
    .addTarget(
      prometheus.target(
        'count(couchbase_node_healthy{instance=~"$instance"})',
      )
    )
  )
  .addPanel(
    singlestat.new(
      'Unhealthy Server Count',
      format='none',
      span=2,
      datasource='$PROMETHEUS_DS',
      valueName='current',
      valueFontSize='200%',
      thresholds='0,1',
      colorBackground=true,
      colors=[
        '#299c46',
        '#299c46',
        '#d44a3a',
      ],
    )
    .addTarget(
      prometheus.target(
        'count(couchbase_node_healthy{instance=~"$instance"} == 0) or (1-absent(couchbase_node_healthy{instance=~"$instance"} == 0))',
      )
    )
  )
)
.addRow(
  row.new(
    title='Details',
    collapse=false,
  )
  .addPanel(
    graphPanel.new(
      'QPS',
      datasource='$PROMETHEUS_DS',
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
      datasource='$PROMETHEUS_DS',
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
  .addPanel(
    graphPanel.new(
      'Connections',
      datasource='$PROMETHEUS_DS',
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
        'couchbase_bucket_stats_curr_connections{bucket=~"$bucket",instance=~"$instance"}',
        legendFormat='{{ bucket }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Primary items total',
      datasource='$PROMETHEUS_DS',
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
)
.addRow(
  row.new(
    title='Queries',
    collapse=false,
  )
  .addPanel(
    graphPanel.new(
      'Gets / Sets',
      datasource='$PROMETHEUS_DS',
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
        alias: '/Sets/',
        transform: 'negative-Y',
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
      datasource='$PROMETHEUS_DS',
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
      'Resident Ratio',
      datasource='$PROMETHEUS_DS',
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
      thresholds=[
        {
          value: 20,
          colorMode: 'critical',
          op: 'lt',
          fill: true,
          line: true,
          yaxis: 'left',
        },
      ],
    )
    .addTarget(
      prometheus.target(
        'couchbase_bucket_stats_vbuckets_active_resident_items_ratio{bucket=~"$bucket",instance=~"$instance"}',
        legendFormat='{{ bucket }}',
      )
    )
  )
)
.addRow(
  row.new(
    title='Memory',
    collapse=true,
  )
  .addPanel(
    graphPanel.new(
      'Memory Used',
      datasource='$PROMETHEUS_DS',
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
      thresholds=[
        {
          value: 90,
          colorMode: 'critical',
          op: 'gt',
          fill: true,
          line: true,
          yaxis: 'left',
        },
      ],
    )
    .addTarget(
      prometheus.target(
        '100 * couchbase_bucket_basicstats_memused_bytes{bucket=~"$bucket",instance=~"$instance"} / couchbase_bucket_stats_ep_max_size_bytes',
        legendFormat='{{ bucket }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Items Count',
      datasource='$PROMETHEUS_DS',
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
      datasource='$PROMETHEUS_DS',
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
    title='Disk',
    collapse=true,
  )
  .addPanel(
    graphPanel.new(
      'Disk Fetches',
      datasource='$PROMETHEUS_DS',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
    )
    .addTarget(
      prometheus.target(
        'couchbase_bucket_basicstats_diskfetches{instance=~"$instance", bucket=~"$bucket"}',
        legendFormat='{{ bucket }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Disk Write Queue',
      datasource='$PROMETHEUS_DS',
      span=6,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
    )
    .addTarget(
      prometheus.target(
        'couchbase_bucket_stats_disk_write_queue{instance=~"$instance", bucket=~"$bucket"}',
        legendFormat='{{ bucket }}',
      )
    )
  )
)
.addRow(
  row.new(
    title='Compacting',
    collapse=true,
  )
  .addPanel(
    graphPanel.new(
      'Fragmentation',
      datasource='$PROMETHEUS_DS',
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
      datasource='$PROMETHEUS_DS',
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
      nullPointMode='null as zero',
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
    title='Rebalance',
    collapse=true,
  )
  .addPanel(
    graphPanel.new(
      'Rebalance Progress',
      datasource='$PROMETHEUS_DS',
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
      nullPointMode='null as zero',
    )
    .addTarget(
      prometheus.target(
        'couchbase_task_rebalance_progress{instance=~"$instance"}',
        legendFormat='{{ instance }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'DCP Replication',
      datasource='$PROMETHEUS_DS',
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
        'couchbase_bucket_stats_ep_dcp_replica_producers{instance=~"$instance", bucket=~"$bucket"}',
        legendFormat='{{ bucket }}: DCP Senders',
      )
    )
    .addTarget(
      prometheus.target(
        'couchbase_bucket_stats_ep_dcp_replicas{instance=~"$instance", bucket=~"$bucket"}',
        legendFormat='{{ bucket }}: DCP Connections',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Items Sent / Remaining',
      datasource='$PROMETHEUS_DS',
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
        alias: '/Remaining/',
        transform: 'negative-Y',
      }
    )
    .addTarget(
      prometheus.target(
        'couchbase_bucket_stats_ep_dcp_replica_items_sent{instance=~"$instance", bucket=~"$bucket"}',
        legendFormat='Sent on {{ bucket }}',
      )
    )
    .addTarget(
      prometheus.target(
        'couchbase_bucket_stats_ep_dcp_replica_items_remaining{instance=~"$instance", bucket=~"$bucket"}',
        legendFormat='Remaining on {{ bucket }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Speed',
      datasource='$PROMETHEUS_DS',
      span=12,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='Bps',
      min=0,
    )
    .addTarget(
      prometheus.target(
        'couchbase_bucket_stats_ep_dcp_replica_total_bytes{instance=~"$instance", bucket=~"$bucket"}',
        legendFormat='{{ bucket }}',
      )
    )
  )
)
.addRow(
  row.new(
    title='XDCR',
    collapse=true,
  )
  .addPanel(
    graphPanel.new(
      'Items Sent / Remaining',
      datasource='$PROMETHEUS_DS',
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
        alias: '/Remaining/',
        transform: 'negative-Y',
      }
    )
    .addTarget(
      prometheus.target(
        'couchbase_bucket_stats_ep_dcp_xdcr_items_remaining{bucket=~"$bucket",instance=~"$instance"}',
        legendFormat='{{ bucket }}: Remaining',
      )
    )
    .addTarget(
      prometheus.target(
        'couchbase_bucket_stats_ep_dcp_xdcr_items_sent{bucket=~"$bucket", instance=~"$instance"}',
        legendFormat='{{ bucket }}: Sent'
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Speed',
      datasource='$PROMETHEUS_DS',
      span=12,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='Bps',
      min=0,
    )
    .addTarget(
      prometheus.target(
        'couchbase_bucket_stats_ep_dcp_xdcr_total_bytes{bucket=~"$bucket", instance=~"$instance"}',
        legendFormat='{{ bucket }}',
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'Backlog Size',
      datasource='$PROMETHEUS_DS',
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
        'couchbase_bucket_stats_ep_dcp_xdcr_total_backlog_size{bucket=~"$bucket", instance=~"$instance"}',
        legendFormat='{{ bucket }}'
      )
    )
  )
  .addPanel(
    graphPanel.new(
      'DCP',
      datasource='$PROMETHEUS_DS',
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
        'couchbase_bucket_stats_ep_dcp_xdcr_connections{bucket=~"$bucket", instance=~"$instance"}',
        legendFormat='{{ bucket }}: connections'
      )
    )
    .addTarget(
      prometheus.target(
        'couchbase_bucket_stats_ep_dcp_xdcr_producers{bucket=~"$bucket", instance=~"$instance"}',
        legendFormat='{{ bucket }}: producers'
      )
    )
  )
)
.addRow(
  row.new(
    title='Scrape times',
    collapse=true,
  )
  .addPanel(
    graphPanel.new(
      'Scrape durations',
      datasource='$PROMETHEUS_DS',
      span=12,
      legend_alignAsTable=true,
      legend_rightSide=true,
      legend_values=true,
      legend_current=true,
      legend_sort='current',
      legend_sortDesc=true,
      format='s',
    )
    .addTarget(
      prometheus.target(
        'couchbase_cluster_scrape_duration_seconds{instance="$instance"}',
        legendFormat='cluster',
      )
    )
    .addTarget(
      prometheus.target(
        'couchbase_bucket_scrape_duration_seconds{instance="$instance"}',
        legendFormat='bucket',
      )
    )
    .addTarget(
      prometheus.target(
        'couchbase_node_scrape_duration_seconds{instance="$instance"}',
        legendFormat='node',
      )
    )
    .addTarget(
      prometheus.target(
        'couchbase_task_scrape_duration_seconds{instance="$instance"}',
        legendFormat='task',
      )
    )
  )
)
