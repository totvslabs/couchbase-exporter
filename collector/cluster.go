package collector

import (
	"sync"
	"time"

	"github.com/caarlos0/couchbase-exporter/client"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

type clusterCollector struct {
	mutex  sync.Mutex
	client client.Client

	up             *prometheus.Desc
	scrapeDuration *prometheus.Desc
}

func NewClusterCollector(client client.Client) prometheus.Collector {
	const ns = "cluster"
	return &clusterCollector{
		client: client,
		up: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "up"),
			"Couchbase cluster API is responding",
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
func (c *clusterCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.up
	ch <- c.scrapeDuration
}

// Collect all metrics
func (c *clusterCollector) Collect(ch chan<- prometheus.Metric) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	start := time.Now()
	log.Info("Collecting tasks metrics...")

	cluster, err := c.client.Cluster()
	if err != nil {
		ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 0)
		log.With("error", err).Error("failed to scrape cluster")
		return
	}
	ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 1)

	log.Info(cluster)

	ch <- prometheus.MustNewConstMetric(c.scrapeDuration, prometheus.GaugeValue, time.Since(start).Seconds())
}
