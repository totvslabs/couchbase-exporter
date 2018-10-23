package collector

import (
	"sync"
	"time"

	"github.com/caarlos0/couchbase-exporter/client"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

type nodesCollector struct {
	mutex  sync.Mutex
	client client.Client

	up                                       *prometheus.Desc
	scrapeDuration                           *prometheus.Desc
	healthy                                  *prometheus.Desc
	systemStatsCPUUtilizationRate            *prometheus.Desc
	systemStatsSwapTotal                     *prometheus.Desc
	systemStatsSwapUsed                      *prometheus.Desc
	systemStatsMemTotal                      *prometheus.Desc
	systemStatsMemFree                       *prometheus.Desc
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
	const subsystem = "nodes"
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
		systemStatsCPUUtilizationRate: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "systemStatsCPUUtilizationRate"),
			"systemStatsCPUUtilizationRate",
			[]string{"node"},
			nil,
		),
		systemStatsSwapTotal: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "systemStatsSwapTotal"),
			"systemStatsSwapTotal",
			[]string{"node"},
			nil,
		),
		systemStatsSwapUsed: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "systemStatsSwapUsed"),
			"systemStatsSwapUsed",
			[]string{"node"},
			nil,
		),
		systemStatsMemTotal: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "systemStatsMemTotal"),
			"systemStatsMemTotal",
			[]string{"node"},
			nil,
		),
		systemStatsMemFree: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "systemStatsMemFree"),
			"systemStatsMemFree",
			[]string{"node"},
			nil,
		),
		interestingStatsCouchDocsActualDiskSize: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "interestingStatsCouchDocsActualDiskSize"),
			"interestingStatsCouchDocsActualDiskSize",
			[]string{"node"},
			nil,
		),
		interestingStatsCouchDocsDataSize: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "interestingStatsCouchDocsDataSize"),
			"interestingStatsCouchDocsDataSize",
			[]string{"node"},
			nil,
		),
		interestingStatsCouchViewsActualDiskSize: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "interestingStatsCouchViewsActualDiskSize"),
			"interestingStatsCouchViewsActualDiskSize",
			[]string{"node"},
			nil,
		),
		interestingStatsCouchViewsDataSize: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "interestingStatsCouchViewsDataSize"),
			"interestingStatsCouchViewsDataSize",
			[]string{"node"},
			nil,
		),
		interestingStatsMemUsed: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "interestingStatsMemUsed"),
			"interestingStatsMemUsed",
			[]string{"node"},
			nil,
		),
		interestingStatsOps: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "interestingStatsOps"),
			"interestingStatsOps",
			[]string{"node"},
			nil,
		),
		interestingStatsCurrItems: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "interestingStatsCurrItems"),
			"interestingStatsCurrItems",
			[]string{"node"},
			nil,
		),
		interestingStatsCurrItemsTot: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "interestingStatsCurrItemsTot"),
			"interestingStatsCurrItemsTot",
			[]string{"node"},
			nil,
		),
		interestingStatsVbReplicaCurrItems: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "interestingStatsVbReplicaCurrItems"),
			"interestingStatsVbReplicaCurrItems",
			[]string{"node"},
			nil,
		),
		interestingStatsCouchSpatialDiskSize: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "interestingStatsCouchSpatialDiskSize"),
			"interestingStatsCouchSpatialDiskSize",
			[]string{"node"},
			nil,
		),
		interestingStatsCouchSpatialDataSize: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "interestingStatsCouchSpatialDataSize"),
			"interestingStatsCouchSpatialDataSize",
			[]string{"node"},
			nil,
		),
		interestingStatsCmdGet: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "interestingStatsCmdGet"),
			"interestingStatsCmdGet",
			[]string{"node"},
			nil,
		),
		interestingStatsGetHits: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "interestingStatsGetHits"),
			"interestingStatsGetHits",
			[]string{"node"},
			nil,
		),
		interestingStatsEpBgFetched: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "interestingStatsEpBgFetched"),
			"interestingStatsEpBgFetched",
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
	ch <- c.systemStatsCPUUtilizationRate
	ch <- c.systemStatsSwapTotal
	ch <- c.systemStatsSwapUsed
	ch <- c.systemStatsMemTotal
	ch <- c.systemStatsMemFree
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
	ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 1)

	for _, node := range nodes.Nodes {
		log.Infof("Collecting %s node metrics...", node.Hostname)
		ch <- prometheus.MustNewConstMetric(c.healthy, prometheus.GaugeValue, fromBool(node.Status == "healthy"), node.Hostname)
		ch <- prometheus.MustNewConstMetric(c.systemStatsCPUUtilizationRate, prometheus.GaugeValue, node.SystemStats.CPUUtilizationRate, node.Hostname)
		ch <- prometheus.MustNewConstMetric(c.systemStatsSwapTotal, prometheus.GaugeValue, float64(node.SystemStats.SwapTotal), node.Hostname)
		ch <- prometheus.MustNewConstMetric(c.systemStatsSwapUsed, prometheus.GaugeValue, float64(node.SystemStats.SwapUsed), node.Hostname)
		ch <- prometheus.MustNewConstMetric(c.systemStatsMemTotal, prometheus.GaugeValue, float64(node.SystemStats.MemTotal), node.Hostname)
		ch <- prometheus.MustNewConstMetric(c.systemStatsMemFree, prometheus.GaugeValue, float64(node.SystemStats.MemFree), node.Hostname)
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

	ch <- prometheus.MustNewConstMetric(c.scrapeDuration, prometheus.GaugeValue, time.Since(start).Seconds())
}
