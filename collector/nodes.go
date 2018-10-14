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

	up             *prometheus.Desc
	scrapeDuration *prometheus.Desc
}

// NewNodesCollector nodes collector
//
// TODO: implement this
//
func NewNodesCollector(client client.Client) prometheus.Collector {
	const ns = "nodes"
	return &nodesCollector{
		client: client,
		up: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "up"),
			"Couchbase nodes API is responding",
			nil,
			nil,
		),
		scrapeDuration: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "scrape_duration_seconds"),
			"Scrape duration in seconds",
			nil,
			nil,
		),
	}
}

// Describe all metrics
func (c *nodesCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.up
	ch <- c.scrapeDuration
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

	// for _, node := range nodes {
	log.Info(nodes.Hostname)
	// }

	ch <- prometheus.MustNewConstMetric(c.scrapeDuration, prometheus.GaugeValue, time.Since(start).Seconds())
}
