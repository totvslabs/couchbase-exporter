package collector

import (
	"sync"
	"time"

	"github.com/caarlos0/couchbase-exporter/client"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

type bucketsCollector struct {
	mutex  sync.Mutex
	client client.Client

	up             *prometheus.Desc
	scrapeDuration *prometheus.Desc
}

func NewBucketsCollector(client client.Client) prometheus.Collector {
	const ns = "buckets"
	return &bucketsCollector{
		client: client,
		up: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "up"),
			"Couchbase buckets API is responding",
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
func (c *bucketsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.up
	ch <- c.scrapeDuration
}

// Collect all metrics
func (c *bucketsCollector) Collect(ch chan<- prometheus.Metric) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	start := time.Now()
	log.Info("Collecting buckets metrics...")

	buckets, err := c.client.Buckets()
	if err != nil {
		ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 0)
		log.With("error", err).Error("failed to scrape buckets")
		return
	}
	ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 1)

	for _, bucket := range buckets {
		log.Info(bucket.Name)
	}

	ch <- prometheus.MustNewConstMetric(c.scrapeDuration, prometheus.GaugeValue, time.Since(start).Seconds())
}
