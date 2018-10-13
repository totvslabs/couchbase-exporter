package collector

import (
	"sync"
	"time"

	"github.com/caarlos0/couchbase-exporter/client"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

type taskCollector struct {
	mutex  sync.Mutex
	client client.Client

	up               *prometheus.Desc
	scrapeDuration   *prometheus.Desc
	rebalance        *prometheus.Desc
	rebalancePerNode *prometheus.Desc
	compacting       *prometheus.Desc
}

func NewTasksCollector(client client.Client) prometheus.Collector {
	const ns = "task"
	return &taskCollector{
		client: client,
		up: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "up"),
			"Couchbase task API is responding",
			nil,
			nil,
		),
		scrapeDuration: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "scrape_duration_seconds"),
			"Scrape duration in seconds",
			nil,
			nil,
		),
		rebalance: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "rebalance_progress"),
			"Progress of a rebalance task",
			nil,
			nil,
		),
		rebalancePerNode: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "node_rebalance_progress"),
			"Progress of a rebalance task per node",
			[]string{"node"},
			nil,
		),
		compacting: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "compacting_progress"),
			"Progress of a bucket compaction task",
			[]string{"bucket"},
			nil,
		),
	}
}

// Describe all metrics
func (c *taskCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.rebalance
	ch <- c.compacting
}

// Collect all metrics
func (c *taskCollector) Collect(ch chan<- prometheus.Metric) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	start := time.Now()
	log.Info("Collecting tasks metrics...")

	tasks, err := c.client.Tasks()
	if err != nil {
		ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 0)
		log.With("error", err).Error("failed to scrape tasks")
		return
	}
	ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 1)

	for _, task := range tasks {
		if task.Status != "running" {
			continue
		}
		switch task.Type {
		case "rebalance":
			ch <- prometheus.MustNewConstMetric(c.rebalance, prometheus.GaugeValue, task.Progress)
			for node, progress := range task.PerNode {
				ch <- prometheus.MustNewConstMetric(c.rebalancePerNode, prometheus.GaugeValue, progress.Progress, node)
			}
		case "bucket_compaction":
			ch <- prometheus.MustNewConstMetric(c.compacting, prometheus.GaugeValue, task.Progress, task.Bucket)
		default:
			log.With("type", task.Type).Error("not implemented")
		}
	}

	ch <- prometheus.MustNewConstMetric(c.scrapeDuration, prometheus.GaugeValue, time.Since(start).Seconds())
}
