package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"github.com/totvslabs/couchbase-exporter/client"
	"github.com/totvslabs/couchbase-exporter/collector"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

// nolint: gochecknoglobals,lll
var (
	version           = "dev"
	app               = kingpin.New("couchbase-exporter", "exports couchbase metrics in the prometheus format")
	listenAddress     = app.Flag("web.listen-address", "Address to listen on for web interface and telemetry").Default(":9420").String()
	metricsPath       = app.Flag("web.telemetry-path", "Path under which to expose metrics").Default("/metrics").String()
	couchbaseURL      = app.Flag("couchbase.url", "Couchbase URL to scrape").Default("http://localhost:8091").String()
	couchbaseUsername = app.Flag("couchbase.username", "Couchbase username").String()
	couchbasePassword = app.Flag("couchbase.password", "Couchbase password").OverrideDefaultFromEnvar("COUCHBASE_PASSWORD").String()

	tasks   = app.Flag("collectors.tasks", "Whether to collect tasks metrics").Default("true").Bool()
	buckets = app.Flag("collectors.buckets", "Whether to collect buckets metrics").Default("true").Bool()
	nodes   = app.Flag("collectors.nodes", "Whether to collect nodes metrics").Default("true").Bool()
	cluster = app.Flag("collectors.cluster", "Whether to collect cluster metrics").Default("true").Bool()
)

func main() {
	log.AddFlags(app)
	app.Version(version)
	app.HelpFlag.Short('h')
	kingpin.MustParse(app.Parse(os.Args[1:]))

	log.Infof("starting couchbase-exporter %s...", version)

	var client = client.New(*couchbaseURL, *couchbaseUsername, *couchbasePassword)

	if *tasks {
		prometheus.MustRegister(collector.NewTasksCollector(client))
	}
	if *buckets {
		prometheus.MustRegister(collector.NewBucketsCollector(client))
	}
	if *nodes {
		prometheus.MustRegister(collector.NewNodesCollector(client))
	}
	if *cluster {
		prometheus.MustRegister(collector.NewClusterCollector(client))
	}

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
