package main

import (
	"fmt"
	"net/http"

	"github.com/caarlos0/couchbase-exporter/collector"
	"github.com/caarlos0/couchbase-exporter/client"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/prometheus/common/log"
)

var (
	version       = "dev"
	listenAddress = kingpin.Flag("web.listen-address", "Address to listen on for web interface and telemetry").Default(":9420").String()
	metricsPath   = kingpin.Flag("web.telemetry-path", "Path under which to expose metrics").Default("/metrics").String()
	couchbaseURL  = kingpin.Flag("couchbase.url", "Couchbase URL to scrape").Default("http://localhost:8091").String()
	couchbaseUsername  = kingpin.Flag("couchbase.username", "Couchbase username").String()
	couchbasePassword  = kingpin.Flag("couchbase.password", "Couchbase password").String()
)

func main() {
	kingpin.Version(version)
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	log.Info("Starting couchbase-exporter ", version)


	var client = client.New(*couchbaseURL, *couchbaseUsername, *couchbasePassword)

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

	log.Infof("Server listening on %s", *listenAddress)
	if err := http.ListenAndServe(*listenAddress, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

