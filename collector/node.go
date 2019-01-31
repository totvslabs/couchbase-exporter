package collector

import (
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"github.com/totvslabs/couchbase-exporter/client"
)

type nodesCollector struct {
	mutex  sync.Mutex
	client client.Client

	up                                       *prometheus.Desc
	scrapeDuration                           *prometheus.Desc
	healthy                                  *prometheus.Desc
	interestingStatsCouchDocsActualDiskSize  *prometheus.Desc
	interestingStatsCouchDocsDataSize        *prometheus.Desc
	interestingStatsCouchViewsActualDiskSize *prometheus.Desc
	interestingStatsCouchViewsDataSize       *prometheus.Desc
	interestingStatsMemUsed                  *prometheus.Desc
	interestingStatsOps                      *prometheus.Desc
	interestingStatsCurrItems                *prometheus.Desc
	interestingStatsCurrItemsTot             *prometheus.Desc
	interestingStatsVbReplicaCurrItems       *prometheus.Desc
	interestingStatsCouchSpatialDiskSize     *prometheus.Desc
	interestingStatsCouchSpatialDataSize     *prometheus.Desc
	interestingStatsCmdGet                   *prometheus.Desc
	interestingStatsGetHits                  *prometheus.Desc
	interestingStatsEpBgFetched              *prometheus.Desc
}

// NewNodesCollector nodes collector
func NewNodesCollector(client client.Client) prometheus.Collector {
	const subsystem = "node"
	// nolint: lll
	return &nodesCollector{
		client: client,
		up: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "up"),
			"Couchbase nodes API is responding",
			nil,
			nil,
		),
		scrapeDuration: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "scrape_duration_seconds"),
			"Scrape duration in seconds",
			nil,
			nil,
		),
		healthy: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "healthy"),
			"Is this node healthy",
			[]string{"node"},
			nil,
		),
		interestingStatsCouchDocsActualDiskSize: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "interestingstats_couch_docs_actual_disk_size"),
			"interestingstats_couch_docs_actual_disk_size",
			[]string{"node"},
			nil,
		),
		interestingStatsCouchDocsDataSize: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "interestingstats_couch_docs_data_size"),
			"interestingstats_couch_docs_data_size",
			[]string{"node"},
			nil,
		),
		interestingStatsCouchViewsActualDiskSize: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "interestingstats_couch_views_actual_disk_size"),
			"interestingstats_couch_views_actual_disk_size",
			[]string{"node"},
			nil,
		),
		interestingStatsCouchViewsDataSize: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "interestingstats_couch_views_data_size"),
			"interestingstats_couch_views_data_size",
			[]string{"node"},
			nil,
		),
		interestingStatsMemUsed: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "interestingstats_mem_used"),
			"interestingstats_mem_used",
			[]string{"node"},
			nil,
		),
		interestingStatsOps: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "interestingstats_ops"),
			"interestingstats_ops",
			[]string{"node"},
			nil,
		),
		interestingStatsCurrItems: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "interestingstats_curr_items"),
			"interestingstats_curr_items",
			[]string{"node"},
			nil,
		),
		interestingStatsCurrItemsTot: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "interestingstats_curr_items_tot"),
			"interestingstats_curr_items_tot",
			[]string{"node"},
			nil,
		),
		interestingStatsVbReplicaCurrItems: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "interestingstats_vb_replica_curr_items"),
			"interestingstats_vb_replica_curr_items",
			[]string{"node"},
			nil,
		),
		interestingStatsCouchSpatialDiskSize: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "interestingstats_couch_spatial_disk_size"),
			"interestingstats_couch_spatial_disk_size",
			[]string{"node"},
			nil,
		),
		interestingStatsCouchSpatialDataSize: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "interestingstats_couch_spatial_data_size"),
			"interestingstats_couch_spatial_data_size",
			[]string{"node"},
			nil,
		),
		interestingStatsCmdGet: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "interestingstats_cmd_get"),
			"interestingstats_cmd_get",
			[]string{"node"},
			nil,
		),
		interestingStatsGetHits: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "interestingstats_get_hits"),
			"interestingstats_get_hits",
			[]string{"node"},
			nil,
		),
		interestingStatsEpBgFetched: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "interestingstats_ep_bg_fetched"),
			"interestingstats_ep_bg_fetched",
			[]string{"node"},
			nil,
		),
	}
}

// Describe all metrics
func (c *nodesCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.up
	ch <- c.scrapeDuration
	ch <- c.healthy
	ch <- c.interestingStatsCouchDocsActualDiskSize
	ch <- c.interestingStatsCouchDocsDataSize
	ch <- c.interestingStatsCouchViewsActualDiskSize
	ch <- c.interestingStatsCouchViewsDataSize
	ch <- c.interestingStatsMemUsed
	ch <- c.interestingStatsOps
	ch <- c.interestingStatsCurrItems
	ch <- c.interestingStatsCurrItemsTot
	ch <- c.interestingStatsVbReplicaCurrItems
	ch <- c.interestingStatsCouchSpatialDiskSize
	ch <- c.interestingStatsCouchSpatialDataSize
	ch <- c.interestingStatsCmdGet
	ch <- c.interestingStatsGetHits
	ch <- c.interestingStatsEpBgFetched
}

// Collect all metrics
func (c *nodesCollector) Collect(ch chan<- prometheus.Metric) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	start := time.Now()
	log.Info("Collecting nodes metrics...")

	nodes, err := c.client.Nodes()
	if err != nil {
		ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 0)
		log.With("error", err).Error("failed to scrape nodes")
		return
	}

	// nolint: lll
	for _, node := range nodes.Nodes {
		log.Debugf("Collecting %s node metrics...", node.Hostname)
		ch <- prometheus.MustNewConstMetric(c.healthy, prometheus.GaugeValue, fromBool(node.Status == "healthy"), node.Hostname)
		ch <- prometheus.MustNewConstMetric(c.interestingStatsCouchDocsActualDiskSize, prometheus.GaugeValue, node.InterestingStats.CouchDocsActualDiskSize, node.Hostname)
		ch <- prometheus.MustNewConstMetric(c.interestingStatsCouchDocsDataSize, prometheus.GaugeValue, node.InterestingStats.CouchDocsDataSize, node.Hostname)
		ch <- prometheus.MustNewConstMetric(c.interestingStatsCouchViewsActualDiskSize, prometheus.GaugeValue, node.InterestingStats.CouchViewsActualDiskSize, node.Hostname)
		ch <- prometheus.MustNewConstMetric(c.interestingStatsCouchViewsDataSize, prometheus.GaugeValue, node.InterestingStats.CouchViewsDataSize, node.Hostname)
		ch <- prometheus.MustNewConstMetric(c.interestingStatsMemUsed, prometheus.GaugeValue, node.InterestingStats.MemUsed, node.Hostname)
		ch <- prometheus.MustNewConstMetric(c.interestingStatsOps, prometheus.GaugeValue, node.InterestingStats.Ops, node.Hostname)
		ch <- prometheus.MustNewConstMetric(c.interestingStatsCurrItems, prometheus.GaugeValue, node.InterestingStats.CurrItems, node.Hostname)
		ch <- prometheus.MustNewConstMetric(c.interestingStatsCurrItemsTot, prometheus.GaugeValue, node.InterestingStats.CurrItemsTot, node.Hostname)
		ch <- prometheus.MustNewConstMetric(c.interestingStatsVbReplicaCurrItems, prometheus.GaugeValue, node.InterestingStats.VbReplicaCurrItems, node.Hostname)
		ch <- prometheus.MustNewConstMetric(c.interestingStatsCouchSpatialDiskSize, prometheus.GaugeValue, node.InterestingStats.CouchSpatialDiskSize, node.Hostname)
		ch <- prometheus.MustNewConstMetric(c.interestingStatsCouchSpatialDataSize, prometheus.GaugeValue, node.InterestingStats.CouchSpatialDataSize, node.Hostname)
		ch <- prometheus.MustNewConstMetric(c.interestingStatsCmdGet, prometheus.GaugeValue, node.InterestingStats.CmdGet, node.Hostname)
		ch <- prometheus.MustNewConstMetric(c.interestingStatsGetHits, prometheus.GaugeValue, node.InterestingStats.GetHits, node.Hostname)
		ch <- prometheus.MustNewConstMetric(c.interestingStatsEpBgFetched, prometheus.GaugeValue, node.InterestingStats.EpBgFetched, node.Hostname)
	}

	ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 1)
	// nolint: lll
	ch <- prometheus.MustNewConstMetric(c.scrapeDuration, prometheus.GaugeValue, time.Since(start).Seconds())
}
