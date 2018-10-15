package main

import (
	"fmt"
	"net/http"

	"github.com/caarlos0/couchbase-exporter/client"
	"github.com/caarlos0/couchbase-exporter/collector"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	version           = "dev"
	listenAddress     = kingpin.Flag("web.listen-address", "Address to listen on for web interface and telemetry").Default(":9420").String()
	metricsPath       = kingpin.Flag("web.telemetry-path", "Path under which to expose metrics").Default("/metrics").String()
	couchbaseURL      = kingpin.Flag("couchbase.url", "Couchbase URL to scrape").Default("http://localhost:8091").String()
	couchbaseUsername = kingpin.Flag("couchbase.username", "Couchbase username").String()
	couchbasePassword = kingpin.Flag("couchbase.password", "Couchbase password").String()
)

func main() {
	kingpin.Version(version)
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	log.Info("starting couchbase-exporter", version)

	client, err := client.New(*couchbaseURL, *couchbaseUsername, *couchbasePassword)
	if err != nil {
		log.Fatalf("failed to create couchbase client: %v", err)
	}

	prometheus.MustRegister(collector.NewTasksCollector(client))
	prometheus.MustRegister(collector.NewBucketsCollector(client))
	prometheus.MustRegister(collector.NewNodesCollector(client))
	prometheus.MustRegister(collector.NewClusterCollector(client))

	http.Handle(*metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,
			`
			<html>
			<head><title>Couchbase Exporter</title></head>
			<body>
				<h1>Couchbase Exporter</h1>
				<p><a href="`+*metricsPath+`">Metrics</a></p>
			</body>
			</html>
			`)
	})

	log.Infof("server listening on %s", *listenAddress)
	if err := http.ListenAndServe(*listenAddress, nil); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
