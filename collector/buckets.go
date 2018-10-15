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

	basicstatsDataused                        *prometheus.Desc
	basicstatsDiskfetches                     *prometheus.Desc
	basicstatsDiskused                        *prometheus.Desc
	basicstatsItemcount                       *prometheus.Desc
	basicstatsMemused                         *prometheus.Desc
	basicstatsOpspersec                       *prometheus.Desc
	basicstatsQuotapercentused                *prometheus.Desc
	statsAvgBgWaitTime                        *prometheus.Desc
	statsAvgDiskCommitTime                    *prometheus.Desc
	statsAvgDiskUpdateTime                    *prometheus.Desc
	statsBgWaitCount                          *prometheus.Desc
	statsBgWaitTotal                          *prometheus.Desc
	statsBytesRead                            *prometheus.Desc
	statsBytesWritten                         *prometheus.Desc
	statsCasBadval                            *prometheus.Desc
	statsCasHits                              *prometheus.Desc
	statsCasMisses                            *prometheus.Desc
	statsCmdGet                               *prometheus.Desc
	statsCmdSet                               *prometheus.Desc
	statsCouchTotalDiskSize                   *prometheus.Desc
	statsCouchDocsDataSize                    *prometheus.Desc
	statsCouchDocsDiskSize                    *prometheus.Desc
	statsCouchDocsActualDiskSize              *prometheus.Desc
	statsCouchDocsFragmentation               *prometheus.Desc
	statsCpuIdleMs                            *prometheus.Desc
	statsCpuLocalMs                           *prometheus.Desc
	statsCpuUtilizationRate                   *prometheus.Desc
	statsCurrConnections                      *prometheus.Desc
	statsCurrItems                            *prometheus.Desc
	statsCurrItemsTot                         *prometheus.Desc
	statsDecrHits                             *prometheus.Desc
	statsDecrMisses                           *prometheus.Desc
	statsDeleteHits                           *prometheus.Desc
	statsDeleteMisses                         *prometheus.Desc
	statsDiskCommitCount                      *prometheus.Desc
	statsDiskCommitTotal                      *prometheus.Desc
	statsDiskUpdateCount                      *prometheus.Desc
	statsDiskUpdateTotal                      *prometheus.Desc
	statsDiskWriteQueue                       *prometheus.Desc
	statsEpBgFetched                          *prometheus.Desc
	statsEpCacheMissRate                      *prometheus.Desc
	statsEpDcp2iBackoff                       *prometheus.Desc
	statsEpDcp2iCount                         *prometheus.Desc
	statsEpDcp2iItemsRemaining                *prometheus.Desc
	statsEpDcp2iItemsSent                     *prometheus.Desc
	statsEpDcp2iProducerCount                 *prometheus.Desc
	statsEpDcp2iTotalBacklogSize              *prometheus.Desc
	statsEpDcp2iTotalBytes                    *prometheus.Desc
	statsEpDcpOtherBackoff                    *prometheus.Desc
	statsEpDcpOtherCount                      *prometheus.Desc
	statsEpDcpOtherItemsRemaining             *prometheus.Desc
	statsEpDcpOtherItemsSent                  *prometheus.Desc
	statsEpDcpOtherProducerCount              *prometheus.Desc
	statsEpDcpOtherTotalBacklogSize           *prometheus.Desc
	statsEpDcpOtherTotalBytes                 *prometheus.Desc
	statsEpDcpReplicaBackoff                  *prometheus.Desc
	statsEpDcpReplicaCount                    *prometheus.Desc
	statsEpDcpReplicaItemsRemaining           *prometheus.Desc
	statsEpDcpReplicaItemsSent                *prometheus.Desc
	statsEpDcpReplicaProducerCount            *prometheus.Desc
	statsEpDcpReplicaTotalBacklogSize         *prometheus.Desc
	statsEpDcpReplicaTotalBytes               *prometheus.Desc
	statsEpDcpViewsBackoff                    *prometheus.Desc
	statsEpDcpViewsCount                      *prometheus.Desc
	statsEpDcpViewsItemsRemaining             *prometheus.Desc
	statsEpDcpViewsItemsSent                  *prometheus.Desc
	statsEpDcpViewsProducerCount              *prometheus.Desc
	statsEpDcpViewsTotalBacklogSize           *prometheus.Desc
	statsEpDcpViewsTotalBytes                 *prometheus.Desc
	statsEpDcpXdcrBackoff                     *prometheus.Desc
	statsEpDcpXdcrCount                       *prometheus.Desc
	statsEpDcpXdcrItemsRemaining              *prometheus.Desc
	statsEpDcpXdcrItemsSent                   *prometheus.Desc
	statsEpDcpXdcrProducerCount               *prometheus.Desc
	statsEpDcpXdcrTotalBacklogSize            *prometheus.Desc
	statsEpDcpXdcrTotalBytes                  *prometheus.Desc
	statsEpDiskqueueDrain                     *prometheus.Desc
	statsEpDiskqueueFill                      *prometheus.Desc
	statsEpDiskqueueItems                     *prometheus.Desc
	statsEpFlusherTodo                        *prometheus.Desc
	statsEpItemCommitFailed                   *prometheus.Desc
	statsEpKvSize                             *prometheus.Desc
	statsEpMaxSize                            *prometheus.Desc
	statsEpMemHighWat                         *prometheus.Desc
	statsEpMemLowWat                          *prometheus.Desc
	statsEpMetaDataMemory                     *prometheus.Desc
	statsEpNumNonResident                     *prometheus.Desc
	statsEpNumOpsDelMeta                      *prometheus.Desc
	statsEpNumOpsDelRetMeta                   *prometheus.Desc
	statsEpNumOpsGetMeta                      *prometheus.Desc
	statsEpNumOpsSetMeta                      *prometheus.Desc
	statsEpNumOpsSetRetMeta                   *prometheus.Desc
	statsEpNumValueEjects                     *prometheus.Desc
	statsEpOomErrors                          *prometheus.Desc
	statsEpOpsCreate                          *prometheus.Desc
	statsEpOpsUpdate                          *prometheus.Desc
	statsEpOverhead                           *prometheus.Desc
	statsEpQueueSize                          *prometheus.Desc
	statsEpResidentItemsRate                  *prometheus.Desc
	statsEpTapRebalanceCount                  *prometheus.Desc
	statsEpTapRebalanceQlen                   *prometheus.Desc
	statsEpTapRebalanceQueueBackfillremaining *prometheus.Desc
	statsEpTapRebalanceQueueBackoff           *prometheus.Desc
	statsEpTapRebalanceQueueDrain             *prometheus.Desc
	statsEpTapRebalanceQueueFill              *prometheus.Desc
	statsEpTapRebalanceQueueItemondisk        *prometheus.Desc
	statsEpTapRebalanceTotalBacklogSize       *prometheus.Desc
	statsEpTapReplicaCount                    *prometheus.Desc
	statsEpTapReplicaQlen                     *prometheus.Desc
	statsEpTapReplicaQueueBackfillremaining   *prometheus.Desc
	statsEpTapReplicaQueueBackoff             *prometheus.Desc
	statsEpTapReplicaQueueDrain               *prometheus.Desc
	statsEpTapReplicaQueueFill                *prometheus.Desc
	statsEpTapReplicaQueueItemondisk          *prometheus.Desc
	statsEpTapReplicaTotalBacklogSize         *prometheus.Desc
	statsEpTapTotalCount                      *prometheus.Desc
	statsEpTapTotalQlen                       *prometheus.Desc
	statsEpTapTotalQueueBackfillremaining     *prometheus.Desc
	statsEpTapTotalQueueBackoff               *prometheus.Desc
	statsEpTapTotalQueueDrain                 *prometheus.Desc
	statsEpTapTotalQueueFill                  *prometheus.Desc
	statsEpTapTotalQueueItemondisk            *prometheus.Desc
	statsEpTapTotalTotalBacklogSize           *prometheus.Desc
	statsEpTapUserCount                       *prometheus.Desc
	statsEpTapUserQlen                        *prometheus.Desc
	statsEpTapUserQueueBackfillremaining      *prometheus.Desc
	statsEpTapUserQueueBackoff                *prometheus.Desc
	statsEpTapUserQueueDrain                  *prometheus.Desc
	statsEpTapUserQueueFill                   *prometheus.Desc
	statsEpTapUserQueueItemondisk             *prometheus.Desc
	statsEpTapUserTotalBacklogSize            *prometheus.Desc
	statsEpTmpOomErrors                       *prometheus.Desc
	statsEpVbTotal                            *prometheus.Desc
	statsEvictions                            *prometheus.Desc
	statsGetHits                              *prometheus.Desc
	statsGetMisses                            *prometheus.Desc
	statsHibernatedRequests                   *prometheus.Desc
	statsHibernatedWaked                      *prometheus.Desc
	statsHitRatio                             *prometheus.Desc
	statsIncrHits                             *prometheus.Desc
	statsIncrMisses                           *prometheus.Desc
	statsMemActualFree                        *prometheus.Desc
	statsMemActualUsed                        *prometheus.Desc
	statsMemFree                              *prometheus.Desc
	statsMemTotal                             *prometheus.Desc
	statsMemUsed                              *prometheus.Desc
	statsMemUsedSys                           *prometheus.Desc
	statsMisses                               *prometheus.Desc
	statsOps                                  *prometheus.Desc
	statsRestRequests                         *prometheus.Desc
	statsSwapTotal                            *prometheus.Desc
	statsSwapUsed                             *prometheus.Desc
	statsTimestamp                            *prometheus.Desc
	statsVbActiveEject                        *prometheus.Desc
	statsVbActiveItmMemory                    *prometheus.Desc
	statsVbActiveMetaDataMemory               *prometheus.Desc
	statsVbActiveNum                          *prometheus.Desc
	statsVbActiveNumNonResident               *prometheus.Desc
	statsVbActiveOpsCreate                    *prometheus.Desc
	statsVbActiveOpsUpdate                    *prometheus.Desc
	statsVbActiveQueueAge                     *prometheus.Desc
	statsVbActiveQueueDrain                   *prometheus.Desc
	statsVbActiveQueueFill                    *prometheus.Desc
	statsVbActiveQueueSize                    *prometheus.Desc
	statsVbActiveResidentItemsRatio           *prometheus.Desc
	statsVbAvgActiveQueueAge                  *prometheus.Desc
	statsVbAvgPendingQueueAge                 *prometheus.Desc
	statsVbAvgReplicaQueueAge                 *prometheus.Desc
	statsVbAvgTotalQueueAge                   *prometheus.Desc
	statsVbPendingCurrItems                   *prometheus.Desc
	statsVbPendingEject                       *prometheus.Desc
	statsVbPendingItmMemory                   *prometheus.Desc
	statsVbPendingMetaDataMemory              *prometheus.Desc
	statsVbPendingNum                         *prometheus.Desc
	statsVbPendingNumNonResident              *prometheus.Desc
	statsVbPendingOpsCreate                   *prometheus.Desc
	statsVbPendingOpsUpdate                   *prometheus.Desc
	statsVbPendingQueueAge                    *prometheus.Desc
	statsVbPendingQueueDrain                  *prometheus.Desc
	statsVbPendingQueueFill                   *prometheus.Desc
	statsVbPendingQueueSize                   *prometheus.Desc
	statsVbPendingResidentItemsRatio          *prometheus.Desc
	statsVbReplicaCurrItems                   *prometheus.Desc
	statsVbReplicaEject                       *prometheus.Desc
	statsVbReplicaItmMemory                   *prometheus.Desc
	statsVbReplicaMetaDataMemory              *prometheus.Desc
	statsVbReplicaNum                         *prometheus.Desc
	statsVbReplicaNumNonResident              *prometheus.Desc
	statsVbReplicaOpsCreate                   *prometheus.Desc
	statsVbReplicaOpsUpdate                   *prometheus.Desc
	statsVbReplicaQueueAge                    *prometheus.Desc
	statsVbReplicaQueueDrain                  *prometheus.Desc
	statsVbReplicaQueueFill                   *prometheus.Desc
	statsVbReplicaQueueSize                   *prometheus.Desc
	statsVbReplicaResidentItemsRatio          *prometheus.Desc
	statsVbTotalQueueAge                      *prometheus.Desc
	statsXdcOps                               *prometheus.Desc
}

// NewBucketsCollector buckets collector
//
// TODO: add help to the metrics
//
func NewBucketsCollector(client client.Client) prometheus.Collector {
	const ns = "bucket"
	const nsBsic = "bucket_basicstats"
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
		basicstatsDataused: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "basicstats_dataused"),
			"basicstats_dataused",
			[]string{"bucket"},
			nil,
		),
		basicstatsDiskfetches: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "basicstats_diskfetches"),
			"basicstats_diskfetches",
			[]string{"bucket"},
			nil,
		),
		basicstatsDiskused: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "basicstats_diskused"),
			"basicstats_diskused",
			[]string{"bucket"},
			nil,
		),
		basicstatsItemcount: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "basicstats_itemcount"),
			"basicstats_itemcount",
			[]string{"bucket"},
			nil,
		),
		basicstatsMemused: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "basicstats_memused"),
			"basicstats_memused",
			[]string{"bucket"},
			nil,
		),
		basicstatsOpspersec: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "basicstats_opspersec"),
			"basicstats_opspersec",
			[]string{"bucket"},
			nil,
		),
		basicstatsQuotapercentused: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "basicstats_quotapercentused"),
			"basicstats_quotapercentused",
			[]string{"bucket"},
			nil,
		),
		statsAvgBgWaitTime: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_avg_bg_wait_time"),
			"Average background fetch time in microseconds",
			[]string{"bucket"},
			nil,
		),
		statsAvgDiskCommitTime: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_avg_disk_commit_time"),
			"Average disk commit time in seconds as from disk_update histogram of timings",
			[]string{"bucket"},
			nil,
		),
		statsAvgDiskUpdateTime: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_avg_disk_update_time"),
			"Average disk update time in microseconds as from disk_update histogram of timings",
			[]string{"bucket"},
			nil,
		),
		statsBgWaitCount: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_bg_wait_count"),
			"stats_bg_wait_count",
			[]string{"bucket"},
			nil,
		),
		statsBgWaitTotal: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_bg_wait_total"),
			"stats_bg_wait_total",
			[]string{"bucket"},
			nil,
		),
		statsBytesRead: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_bytes_read"),
			"stats_bytes_read",
			[]string{"bucket"},
			nil,
		),
		statsBytesWritten: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_bytes_written"),
			"stats_bytes_written",
			[]string{"bucket"},
			nil,
		),
		statsCasBadval: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_cas_badval"),
			"stats_cas_badval",
			[]string{"bucket"},
			nil,
		),
		statsCasHits: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_cas_hits"),
			"Number of operations with a CAS id per second for this bucket",
			[]string{"bucket"},
			nil,
		),
		statsCasMisses: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_cas_misses"),
			"stats_cas_misses",
			[]string{"bucket"},
			nil,
		),
		statsCmdGet: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_cmd_get"),
			"Number of reads (get operations) per second from this bucket",
			[]string{"bucket"},
			nil,
		),
		statsCmdSet: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_cmd_set"),
			"Number of writes (set operations) per second to this bucket",
			[]string{"bucket"},
			nil,
		),
		statsCouchTotalDiskSize: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "couch_total_disk_size"),
			"The total size on disk of all data and view files for this bucket",
			[]string{"bucket"},
			nil,
		),
		statsCouchDocsFragmentation: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "couch_docs_fragmentation"),
			"How much fragmented data there is to be compacted compared to real data for the data files in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsCouchDocsActualDiskSize: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "couch_docs_actual_disk_size"),
			"The size of all data files for this bucket, including the data itself, meta data and temporary files",
			[]string{"bucket"},
			nil,
		),
		statsCouchDocsDataSize: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_couch_docs_data_size"),
			"The size of active data in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsCouchDocsDiskSize: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_couch_docs_disk_size"),
			"The size of all data files for this bucket, including the data itself, meta data and temporary files",
			[]string{"bucket"},
			nil,
		),
		statsCpuIdleMs: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_cpu_idle_ms"),
			"stats_cpu_idle_ms",
			[]string{"bucket"},
			nil,
		),
		statsCpuLocalMs: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_cpu_local_ms"),
			"stats_cpu_local_ms",
			[]string{"bucket"},
			nil,
		),
		statsCpuUtilizationRate: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_cpu_utilization_rate"),
			"Percentage of CPU in use across all available cores on this server",
			[]string{"bucket"},
			nil,
		),
		statsCurrConnections: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_curr_connections"),
			"Number of connections to this server includingconnections from external client SDKs, proxies, TAP requests and internal statistic gathering",
			[]string{"bucket"},
			nil,
		),
		statsCurrItems: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_curr_items"),
			"Number of unique items in this bucket - only active items, not replica",
			[]string{"bucket"},
			nil,
		),
		statsCurrItemsTot: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_curr_items_tot"),
			"Total number of items in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsDecrHits: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_decr_hits"),
			"stats_decr_hits",
			[]string{"bucket"},
			nil,
		),
		statsDecrMisses: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_decr_misses"),
			"stats_decr_misses",
			[]string{"bucket"},
			nil,
		),
		statsDeleteHits: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_delete_hits"),
			"Number of delete operations per second for this bucket",
			[]string{"bucket"},
			nil,
		),
		statsDeleteMisses: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_delete_misses"),
			"stats_delete_misses",
			[]string{"bucket"},
			nil,
		),
		statsDiskCommitCount: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_disk_commit_count"),
			"stats_disk_commit_count",
			[]string{"bucket"},
			nil,
		),
		statsDiskCommitTotal: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_disk_commit_total"),
			"stats_disk_commit_total",
			[]string{"bucket"},
			nil,
		),
		statsDiskUpdateCount: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_disk_update_count"),
			"stats_disk_update_count",
			[]string{"bucket"},
			nil,
		),
		statsDiskUpdateTotal: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_disk_update_total"),
			"stats_disk_update_total",
			[]string{"bucket"},
			nil,
		),
		statsDiskWriteQueue: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_disk_write_queue"),
			"Number of items waiting to be written to disk in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpBgFetched: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_bg_fetched"),
			"Number of reads per second from disk for this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpCacheMissRate: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_cache_miss_rate"),
			"Percentage of reads per second to this bucket from disk as opposed to RAM",
			[]string{"bucket"},
			nil,
		),
		statsEpDcp2iBackoff: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_2i_backoff"),
			"stats_ep_dcp_2i_backoff",
			[]string{"bucket"},
			nil,
		),
		statsEpDcp2iCount: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_2i_count"),
			"stats_ep_dcp_2i_count",
			[]string{"bucket"},
			nil,
		),
		statsEpDcp2iItemsRemaining: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_2i_items_remaining"),
			"stats_ep_dcp_2i_items_remaining",
			[]string{"bucket"},
			nil,
		),
		statsEpDcp2iItemsSent: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_2i_items_sent"),
			"stats_ep_dcp_2i_items_sent",
			[]string{"bucket"},
			nil,
		),
		statsEpDcp2iProducerCount: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_2i_producer_count"),
			"stats_ep_dcp_2i_producer_count",
			[]string{"bucket"},
			nil,
		),
		statsEpDcp2iTotalBacklogSize: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_2i_total_backlog_size"),
			"stats_ep_dcp_2i_total_backlog_size",
			[]string{"bucket"},
			nil,
		),
		statsEpDcp2iTotalBytes: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_2i_total_bytes"),
			"stats_ep_dcp_2i_total_bytes",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpOtherBackoff: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_other_backoff"),
			"stats_ep_dcp_other_backoff",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpOtherCount: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_other_count"),
			"stats_ep_dcp_other_count",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpOtherItemsRemaining: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_other_items_remaining"),
			"stats_ep_dcp_other_items_remaining",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpOtherItemsSent: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_other_items_sent"),
			"stats_ep_dcp_other_items_sent",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpOtherProducerCount: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_other_producer_count"),
			"stats_ep_dcp_other_producer_count",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpOtherTotalBacklogSize: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_other_total_backlog_size"),
			"stats_ep_dcp_other_total_backlog_size",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpOtherTotalBytes: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_other_total_bytes"),
			"stats_ep_dcp_other_total_bytes",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpReplicaBackoff: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_replica_backoff"),
			"stats_ep_dcp_replica_backoff",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpReplicaCount: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_replica_count"),
			"stats_ep_dcp_replica_count",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpReplicaItemsRemaining: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_replica_items_remaining"),
			"stats_ep_dcp_replica_items_remaining",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpReplicaItemsSent: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_replica_items_sent"),
			"stats_ep_dcp_replica_items_sent",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpReplicaProducerCount: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_replica_producer_count"),
			"stats_ep_dcp_replica_producer_count",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpReplicaTotalBacklogSize: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_replica_total_backlog_size"),
			"stats_ep_dcp_replica_total_backlog_size",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpReplicaTotalBytes: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_replica_total_bytes"),
			"stats_ep_dcp_replica_total_bytes",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpViewsBackoff: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_views_backoff"),
			"stats_ep_dcp_views_backoff",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpViewsCount: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_views_count"),
			"stats_ep_dcp_views_count",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpViewsItemsRemaining: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_views_items_remaining"),
			"stats_ep_dcp_views_items_remaining",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpViewsItemsSent: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_views_items_sent"),
			"stats_ep_dcp_views_items_sent",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpViewsProducerCount: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_views_producer_count"),
			"stats_ep_dcp_views_producer_count",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpViewsTotalBacklogSize: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_views_total_backlog_size"),
			"stats_ep_dcp_views_total_backlog_size",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpViewsTotalBytes: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_views_total_bytes"),
			"stats_ep_dcp_views_total_bytes",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpXdcrBackoff: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_xdcr_backoff"),
			"stats_ep_dcp_xdcr_backoff",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpXdcrCount: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_xdcr_count"),
			"stats_ep_dcp_xdcr_count",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpXdcrItemsRemaining: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_xdcr_items_remaining"),
			"stats_ep_dcp_xdcr_items_remaining",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpXdcrItemsSent: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_xdcr_items_sent"),
			"stats_ep_dcp_xdcr_items_sent",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpXdcrProducerCount: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_xdcr_producer_count"),
			"stats_ep_dcp_xdcr_producer_count",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpXdcrTotalBacklogSize: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_xdcr_total_backlog_size"),
			"stats_ep_dcp_xdcr_total_backlog_size",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpXdcrTotalBytes: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_dcp_xdcr_total_bytes"),
			"stats_ep_dcp_xdcr_total_bytes",
			[]string{"bucket"},
			nil,
		),
		statsEpDiskqueueDrain: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_diskqueue_drain"),
			"stats_ep_diskqueue_drain",
			[]string{"bucket"},
			nil,
		),
		statsEpDiskqueueFill: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_diskqueue_fill"),
			"stats_ep_diskqueue_fill",
			[]string{"bucket"},
			nil,
		),
		statsEpDiskqueueItems: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_diskqueue_items"),
			"Total number of items waiting to be written to disk in the bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpFlusherTodo: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_flusher_todo"),
			"stats_ep_flusher_todo",
			[]string{"bucket"},
			nil,
		),
		statsEpItemCommitFailed: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_item_commit_failed"),
			"stats_ep_item_commit_failed",
			[]string{"bucket"},
			nil,
		),
		statsEpKvSize: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_kv_size"),
			"stats_ep_kv_size",
			[]string{"bucket"},
			nil,
		),
		statsEpMaxSize: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_max_size"),
			"stats_ep_max_size",
			[]string{"bucket"},
			nil,
		),
		statsEpMemHighWat: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_mem_high_wat"),
			"High water mark for auto-evictions",
			[]string{"bucket"},
			nil,
		),
		statsEpMemLowWat: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_mem_low_wat"),
			"Low water mark for auto-evictions",
			[]string{"bucket"},
			nil,
		),
		statsEpMetaDataMemory: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_meta_data_memory"),
			"stats_ep_meta_data_memory",
			[]string{"bucket"},
			nil,
		),
		statsEpNumNonResident: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_num_non_resident"),
			"stats_ep_num_non_resident",
			[]string{"bucket"},
			nil,
		),
		statsEpNumOpsDelMeta: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_num_ops_del_meta"),
			"stats_ep_num_ops_del_meta",
			[]string{"bucket"},
			nil,
		),
		statsEpNumOpsDelRetMeta: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_num_ops_del_ret_meta"),
			"stats_ep_num_ops_del_ret_meta",
			[]string{"bucket"},
			nil,
		),
		statsEpNumOpsGetMeta: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_num_ops_get_meta"),
			"stats_ep_num_ops_get_meta",
			[]string{"bucket"},
			nil,
		),
		statsEpNumOpsSetMeta: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_num_ops_set_meta"),
			"stats_ep_num_ops_set_meta",
			[]string{"bucket"},
			nil,
		),
		statsEpNumOpsSetRetMeta: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_num_ops_set_ret_meta"),
			"stats_ep_num_ops_set_ret_meta",
			[]string{"bucket"},
			nil,
		),
		statsEpNumValueEjects: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_num_value_ejects"),
			"stats_ep_num_value_ejects",
			[]string{"bucket"},
			nil,
		),
		statsEpOomErrors: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_oom_errors"),
			"stats_ep_oom_errors",
			[]string{"bucket"},
			nil,
		),
		statsEpOpsCreate: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_ops_create"),
			"Number of new items created on disk per second for this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpOpsUpdate: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_ops_update"),
			"Number of items updated on disk per second for this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpOverhead: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_overhead"),
			"stats_ep_overhead",
			[]string{"bucket"},
			nil,
		),
		statsEpQueueSize: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_queue_size"),
			"stats_ep_queue_size",
			[]string{"bucket"},
			nil,
		),
		statsEpResidentItemsRate: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_resident_items_rate"),
			"stats_ep_resident_items_rate",
			[]string{"bucket"},
			nil,
		),
		statsEpTapRebalanceCount: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_rebalance_count"),
			"stats_ep_tap_rebalance_count",
			[]string{"bucket"},
			nil,
		),
		statsEpTapRebalanceQlen: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_rebalance_qlen"),
			"stats_ep_tap_rebalance_qlen",
			[]string{"bucket"},
			nil,
		),
		statsEpTapRebalanceQueueBackfillremaining: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_rebalance_queue_backfillremaining"),
			"stats_ep_tap_rebalance_queue_backfillremaining",
			[]string{"bucket"},
			nil,
		),
		statsEpTapRebalanceQueueBackoff: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_rebalance_queue_backoff"),
			"stats_ep_tap_rebalance_queue_backoff",
			[]string{"bucket"},
			nil,
		),
		statsEpTapRebalanceQueueDrain: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_rebalance_queue_drain"),
			"stats_ep_tap_rebalance_queue_drain",
			[]string{"bucket"},
			nil,
		),
		statsEpTapRebalanceQueueFill: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_rebalance_queue_fill"),
			"stats_ep_tap_rebalance_queue_fill",
			[]string{"bucket"},
			nil,
		),
		statsEpTapRebalanceQueueItemondisk: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_rebalance_queue_itemondisk"),
			"stats_ep_tap_rebalance_queue_itemondisk",
			[]string{"bucket"},
			nil,
		),
		statsEpTapRebalanceTotalBacklogSize: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_rebalance_total_backlog_size"),
			"stats_ep_tap_rebalance_total_backlog_size",
			[]string{"bucket"},
			nil,
		),
		statsEpTapReplicaCount: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_replica_count"),
			"stats_ep_tap_replica_count",
			[]string{"bucket"},
			nil,
		),
		statsEpTapReplicaQlen: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_replica_qlen"),
			"stats_ep_tap_replica_qlen",
			[]string{"bucket"},
			nil,
		),
		statsEpTapReplicaQueueBackfillremaining: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_replica_queue_backfillremaining"),
			"stats_ep_tap_replica_queue_backfillremaining",
			[]string{"bucket"},
			nil,
		),
		statsEpTapReplicaQueueBackoff: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_replica_queue_backoff"),
			"stats_ep_tap_replica_queue_backoff",
			[]string{"bucket"},
			nil,
		),
		statsEpTapReplicaQueueDrain: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_replica_queue_drain"),
			"stats_ep_tap_replica_queue_drain",
			[]string{"bucket"},
			nil,
		),
		statsEpTapReplicaQueueFill: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_replica_queue_fill"),
			"stats_ep_tap_replica_queue_fill",
			[]string{"bucket"},
			nil,
		),
		statsEpTapReplicaQueueItemondisk: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_replica_queue_itemondisk"),
			"stats_ep_tap_replica_queue_itemondisk",
			[]string{"bucket"},
			nil,
		),
		statsEpTapReplicaTotalBacklogSize: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_replica_total_backlog_size"),
			"stats_ep_tap_replica_total_backlog_size",
			[]string{"bucket"},
			nil,
		),
		statsEpTapTotalCount: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_total_count"),
			"stats_ep_tap_total_count",
			[]string{"bucket"},
			nil,
		),
		statsEpTapTotalQlen: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_total_qlen"),
			"stats_ep_tap_total_qlen",
			[]string{"bucket"},
			nil,
		),
		statsEpTapTotalQueueBackfillremaining: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_total_queue_backfillremaining"),
			"stats_ep_tap_total_queue_backfillremaining",
			[]string{"bucket"},
			nil,
		),
		statsEpTapTotalQueueBackoff: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_total_queue_backoff"),
			"stats_ep_tap_total_queue_backoff",
			[]string{"bucket"},
			nil,
		),
		statsEpTapTotalQueueDrain: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_total_queue_drain"),
			"stats_ep_tap_total_queue_drain",
			[]string{"bucket"},
			nil,
		),
		statsEpTapTotalQueueFill: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_total_queue_fill"),
			"stats_ep_tap_total_queue_fill",
			[]string{"bucket"},
			nil,
		),
		statsEpTapTotalQueueItemondisk: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_total_queue_itemondisk"),
			"stats_ep_tap_total_queue_itemondisk",
			[]string{"bucket"},
			nil,
		),
		statsEpTapTotalTotalBacklogSize: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_total_total_backlog_size"),
			"stats_ep_tap_total_total_backlog_size",
			[]string{"bucket"},
			nil,
		),
		statsEpTapUserCount: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_user_count"),
			"stats_ep_tap_user_count",
			[]string{"bucket"},
			nil,
		),
		statsEpTapUserQlen: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_user_qlen"),
			"stats_ep_tap_user_qlen",
			[]string{"bucket"},
			nil,
		),
		statsEpTapUserQueueBackfillremaining: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_user_queue_backfillremaining"),
			"stats_ep_tap_user_queue_backfillremaining",
			[]string{"bucket"},
			nil,
		),
		statsEpTapUserQueueBackoff: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_user_queue_backoff"),
			"stats_ep_tap_user_queue_backoff",
			[]string{"bucket"},
			nil,
		),
		statsEpTapUserQueueDrain: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_user_queue_drain"),
			"stats_ep_tap_user_queue_drain",
			[]string{"bucket"},
			nil,
		),
		statsEpTapUserQueueFill: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_user_queue_fill"),
			"stats_ep_tap_user_queue_fill",
			[]string{"bucket"},
			nil,
		),
		statsEpTapUserQueueItemondisk: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_user_queue_itemondisk"),
			"stats_ep_tap_user_queue_itemondisk",
			[]string{"bucket"},
			nil,
		),
		statsEpTapUserTotalBacklogSize: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tap_user_total_backlog_size"),
			"stats_ep_tap_user_total_backlog_size",
			[]string{"bucket"},
			nil,
		),
		statsEpTmpOomErrors: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_tmp_oom_errors"),
			"Number of back-offs sent per second to client SDKs due to OOM situations from this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpVbTotal: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ep_vb_total"),
			"stats_ep_vb_total",
			[]string{"bucket"},
			nil,
		),
		statsEvictions: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_evictions"),
			"stats_evictions",
			[]string{"bucket"},
			nil,
		),
		statsGetHits: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_get_hits"),
			"stats_get_hits",
			[]string{"bucket"},
			nil,
		),
		statsGetMisses: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_get_misses"),
			"stats_get_misses",
			[]string{"bucket"},
			nil,
		),
		statsHibernatedRequests: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_hibernated_requests"),
			"Number of streaming requests on port 8091 now idle",
			[]string{"bucket"},
			nil,
		),
		statsHibernatedWaked: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_hibernated_waked"),
			"Rate of streaming request wakeups on port 8091",
			[]string{"bucket"},
			nil,
		),
		statsHitRatio: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_hit_ratio"),
			"stats_hit_ratio",
			[]string{"bucket"},
			nil,
		),
		statsIncrHits: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_incr_hits"),
			"stats_incr_hits",
			[]string{"bucket"},
			nil,
		),
		statsIncrMisses: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_incr_misses"),
			"stats_incr_misses",
			[]string{"bucket"},
			nil,
		),
		statsMemActualFree: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_mem_actual_free"),
			"Amount of RAM available on this server",
			[]string{"bucket"},
			nil,
		),
		statsMemActualUsed: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_mem_actual_used"),
			"stats_mem_actual_used",
			[]string{"bucket"},
			nil,
		),
		statsMemFree: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_mem_free"),
			"stats_mem_free",
			[]string{"bucket"},
			nil,
		),
		statsMemTotal: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_mem_total"),
			"stats_mem_total",
			[]string{"bucket"},
			nil,
		),
		statsMemUsed: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_mem_used"),
			"Memory used",
			[]string{"bucket"},
			nil,
		),
		statsMemUsedSys: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_mem_used_sys"),
			"stats_mem_used_sys",
			[]string{"bucket"},
			nil,
		),
		statsMisses: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_misses"),
			"stats_misses",
			[]string{"bucket"},
			nil,
		),
		statsOps: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_ops"),
			"Total amount of operations per second to this bucket",
			[]string{"bucket"},
			nil,
		),
		statsRestRequests: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_rest_requests"),
			"Rate of http requests on port 8091",
			[]string{"bucket"},
			nil,
		),
		statsSwapTotal: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_swap_total"),
			"stats_swap_total",
			[]string{"bucket"},
			nil,
		),
		statsSwapUsed: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_swap_used"),
			"Amount of swap space in use on this server",
			[]string{"bucket"},
			nil,
		),
		statsTimestamp: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_timestamp"),
			"stats_timestamp",
			[]string{"bucket"},
			nil,
		),
		statsVbActiveEject: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_active_eject"),
			"stats_vb_active_eject",
			[]string{"bucket"},
			nil,
		),
		statsVbActiveItmMemory: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_active_itm_memory"),
			"stats_vb_active_itm_memory",
			[]string{"bucket"},
			nil,
		),
		statsVbActiveMetaDataMemory: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_active_meta_data_memory"),
			"stats_vb_active_meta_data_memory",
			[]string{"bucket"},
			nil,
		),
		statsVbActiveNum: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_active_num"),
			"stats_vb_active_num",
			[]string{"bucket"},
			nil,
		),
		statsVbActiveNumNonResident: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_active_num_non_resident"),
			"stats_vb_active_num_non_resident",
			[]string{"bucket"},
			nil,
		),
		statsVbActiveOpsCreate: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_active_ops_create"),
			"stats_vb_active_ops_create",
			[]string{"bucket"},
			nil,
		),
		statsVbActiveOpsUpdate: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_active_ops_update"),
			"stats_vb_active_ops_update",
			[]string{"bucket"},
			nil,
		),
		statsVbActiveQueueAge: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_active_queue_age"),
			"stats_vb_active_queue_age",
			[]string{"bucket"},
			nil,
		),
		statsVbActiveQueueDrain: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_active_queue_drain"),
			"stats_vb_active_queue_drain",
			[]string{"bucket"},
			nil,
		),
		statsVbActiveQueueFill: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_active_queue_fill"),
			"stats_vb_active_queue_fill",
			[]string{"bucket"},
			nil,
		),
		statsVbActiveQueueSize: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_active_queue_size"),
			"stats_vb_active_queue_size",
			[]string{"bucket"},
			nil,
		),
		statsVbActiveResidentItemsRatio: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_active_resident_items_ratio"),
			"Percentage of active items cached in RAM in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsVbAvgActiveQueueAge: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_avg_active_queue_age"),
			"stats_vb_avg_active_queue_age",
			[]string{"bucket"},
			nil,
		),
		statsVbAvgPendingQueueAge: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_avg_pending_queue_age"),
			"stats_vb_avg_pending_queue_age",
			[]string{"bucket"},
			nil,
		),
		statsVbAvgReplicaQueueAge: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_avg_replica_queue_age"),
			"stats_vb_avg_replica_queue_age",
			[]string{"bucket"},
			nil,
		),
		statsVbAvgTotalQueueAge: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_avg_total_queue_age"),
			"stats_vb_avg_total_queue_age",
			[]string{"bucket"},
			nil,
		),
		statsVbPendingCurrItems: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_pending_curr_items"),
			"stats_vb_pending_curr_items",
			[]string{"bucket"},
			nil,
		),
		statsVbPendingEject: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_pending_eject"),
			"stats_vb_pending_eject",
			[]string{"bucket"},
			nil,
		),
		statsVbPendingItmMemory: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_pending_itm_memory"),
			"stats_vb_pending_itm_memory",
			[]string{"bucket"},
			nil,
		),
		statsVbPendingMetaDataMemory: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_pending_meta_data_memory"),
			"stats_vb_pending_meta_data_memory",
			[]string{"bucket"},
			nil,
		),
		statsVbPendingNum: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_pending_num"),
			"stats_vb_pending_num",
			[]string{"bucket"},
			nil,
		),
		statsVbPendingNumNonResident: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_pending_num_non_resident"),
			"stats_vb_pending_num_non_resident",
			[]string{"bucket"},
			nil,
		),
		statsVbPendingOpsCreate: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_pending_ops_create"),
			"stats_vb_pending_ops_create",
			[]string{"bucket"},
			nil,
		),
		statsVbPendingOpsUpdate: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_pending_ops_update"),
			"stats_vb_pending_ops_update",
			[]string{"bucket"},
			nil,
		),
		statsVbPendingQueueAge: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_pending_queue_age"),
			"stats_vb_pending_queue_age",
			[]string{"bucket"},
			nil,
		),
		statsVbPendingQueueDrain: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_pending_queue_drain"),
			"stats_vb_pending_queue_drain",
			[]string{"bucket"},
			nil,
		),
		statsVbPendingQueueFill: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_pending_queue_fill"),
			"stats_vb_pending_queue_fill",
			[]string{"bucket"},
			nil,
		),
		statsVbPendingQueueSize: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_pending_queue_size"),
			"stats_vb_pending_queue_size",
			[]string{"bucket"},
			nil,
		),
		statsVbPendingResidentItemsRatio: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_pending_resident_items_ratio"),
			"stats_vb_pending_resident_items_ratio",
			[]string{"bucket"},
			nil,
		),
		statsVbReplicaCurrItems: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_replica_curr_items"),
			"stats_vb_replica_curr_items",
			[]string{"bucket"},
			nil,
		),
		statsVbReplicaEject: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_replica_eject"),
			"stats_vb_replica_eject",
			[]string{"bucket"},
			nil,
		),
		statsVbReplicaItmMemory: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_replica_itm_memory"),
			"stats_vb_replica_itm_memory",
			[]string{"bucket"},
			nil,
		),
		statsVbReplicaMetaDataMemory: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_replica_meta_data_memory"),
			"stats_vb_replica_meta_data_memory",
			[]string{"bucket"},
			nil,
		),
		statsVbReplicaNum: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_replica_num"),
			"stats_vb_replica_num",
			[]string{"bucket"},
			nil,
		),
		statsVbReplicaNumNonResident: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_replica_num_non_resident"),
			"stats_vb_replica_num_non_resident",
			[]string{"bucket"},
			nil,
		),
		statsVbReplicaOpsCreate: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_replica_ops_create"),
			"stats_vb_replica_ops_create",
			[]string{"bucket"},
			nil,
		),
		statsVbReplicaOpsUpdate: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_replica_ops_update"),
			"stats_vb_replica_ops_update",
			[]string{"bucket"},
			nil,
		),
		statsVbReplicaQueueAge: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_replica_queue_age"),
			"stats_vb_replica_queue_age",
			[]string{"bucket"},
			nil,
		),
		statsVbReplicaQueueDrain: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_replica_queue_drain"),
			"stats_vb_replica_queue_drain",
			[]string{"bucket"},
			nil,
		),
		statsVbReplicaQueueFill: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_replica_queue_fill"),
			"stats_vb_replica_queue_fill",
			[]string{"bucket"},
			nil,
		),
		statsVbReplicaQueueSize: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_replica_queue_size"),
			"stats_vb_replica_queue_size",
			[]string{"bucket"},
			nil,
		),
		statsVbReplicaResidentItemsRatio: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_replica_resident_items_ratio"),
			"stats_vb_replica_resident_items_ratio",
			[]string{"bucket"},
			nil,
		),
		statsVbTotalQueueAge: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_vb_total_queue_age"),
			"stats_vb_total_queue_age",
			[]string{"bucket"},
			nil,
		),
		statsXdcOps: prometheus.NewDesc(
			prometheus.BuildFQName(globalNamespace, ns, "stats_xdc_ops"),
			"stats_xdc_ops",
			[]string{"bucket"},
			nil,
		),
	}
}

// Describe all metrics
func (c *bucketsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.up
	ch <- c.scrapeDuration

	ch <- c.basicstatsDataused
	ch <- c.basicstatsDiskfetches
	ch <- c.basicstatsDiskused
	ch <- c.basicstatsItemcount
	ch <- c.basicstatsMemused
	ch <- c.basicstatsOpspersec
	ch <- c.basicstatsQuotapercentused
	ch <- c.statsAvgBgWaitTime
	ch <- c.statsAvgDiskCommitTime
	ch <- c.statsAvgDiskUpdateTime
	ch <- c.statsBgWaitCount
	ch <- c.statsBgWaitTotal
	ch <- c.statsBytesRead
	ch <- c.statsBytesWritten
	ch <- c.statsCasBadval
	ch <- c.statsCasHits
	ch <- c.statsCasMisses
	ch <- c.statsCmdGet
	ch <- c.statsCmdSet
	ch <- c.statsCouchDocsDataSize
	ch <- c.statsCouchDocsActualDiskSize
	ch <- c.statsCouchTotalDiskSize
	ch <- c.statsCouchDocsFragmentation
	ch <- c.statsCouchDocsDiskSize
	ch <- c.statsCpuIdleMs
	ch <- c.statsCpuLocalMs
	ch <- c.statsCpuUtilizationRate
	ch <- c.statsCurrConnections
	ch <- c.statsCurrItems
	ch <- c.statsCurrItemsTot
	ch <- c.statsDecrHits
	ch <- c.statsDecrMisses
	ch <- c.statsDeleteHits
	ch <- c.statsDeleteMisses
	ch <- c.statsDiskCommitCount
	ch <- c.statsDiskCommitTotal
	ch <- c.statsDiskUpdateCount
	ch <- c.statsDiskUpdateTotal
	ch <- c.statsDiskWriteQueue
	ch <- c.statsEpBgFetched
	ch <- c.statsEpCacheMissRate
	ch <- c.statsEpDcp2iBackoff
	ch <- c.statsEpDcp2iCount
	ch <- c.statsEpDcp2iItemsRemaining
	ch <- c.statsEpDcp2iItemsSent
	ch <- c.statsEpDcp2iProducerCount
	ch <- c.statsEpDcp2iTotalBacklogSize
	ch <- c.statsEpDcp2iTotalBytes
	ch <- c.statsEpDcpOtherBackoff
	ch <- c.statsEpDcpOtherCount
	ch <- c.statsEpDcpOtherItemsRemaining
	ch <- c.statsEpDcpOtherItemsSent
	ch <- c.statsEpDcpOtherProducerCount
	ch <- c.statsEpDcpOtherTotalBacklogSize
	ch <- c.statsEpDcpOtherTotalBytes
	ch <- c.statsEpDcpReplicaBackoff
	ch <- c.statsEpDcpReplicaCount
	ch <- c.statsEpDcpReplicaItemsRemaining
	ch <- c.statsEpDcpReplicaItemsSent
	ch <- c.statsEpDcpReplicaProducerCount
	ch <- c.statsEpDcpReplicaTotalBacklogSize
	ch <- c.statsEpDcpReplicaTotalBytes
	ch <- c.statsEpDcpViewsBackoff
	ch <- c.statsEpDcpViewsCount
	ch <- c.statsEpDcpViewsItemsRemaining
	ch <- c.statsEpDcpViewsItemsSent
	ch <- c.statsEpDcpViewsProducerCount
	ch <- c.statsEpDcpViewsTotalBacklogSize
	ch <- c.statsEpDcpViewsTotalBytes
	ch <- c.statsEpDcpXdcrBackoff
	ch <- c.statsEpDcpXdcrCount
	ch <- c.statsEpDcpXdcrItemsRemaining
	ch <- c.statsEpDcpXdcrItemsSent
	ch <- c.statsEpDcpXdcrProducerCount
	ch <- c.statsEpDcpXdcrTotalBacklogSize
	ch <- c.statsEpDcpXdcrTotalBytes
	ch <- c.statsEpDiskqueueDrain
	ch <- c.statsEpDiskqueueFill
	ch <- c.statsEpDiskqueueItems
	ch <- c.statsEpFlusherTodo
	ch <- c.statsEpItemCommitFailed
	ch <- c.statsEpKvSize
	ch <- c.statsEpMaxSize
	ch <- c.statsEpMemHighWat
	ch <- c.statsEpMemLowWat
	ch <- c.statsEpMetaDataMemory
	ch <- c.statsEpNumNonResident
	ch <- c.statsEpNumOpsDelMeta
	ch <- c.statsEpNumOpsDelRetMeta
	ch <- c.statsEpNumOpsGetMeta
	ch <- c.statsEpNumOpsSetMeta
	ch <- c.statsEpNumOpsSetRetMeta
	ch <- c.statsEpNumValueEjects
	ch <- c.statsEpOomErrors
	ch <- c.statsEpOpsCreate
	ch <- c.statsEpOpsUpdate
	ch <- c.statsEpOverhead
	ch <- c.statsEpQueueSize
	ch <- c.statsEpResidentItemsRate
	ch <- c.statsEpTapRebalanceCount
	ch <- c.statsEpTapRebalanceQlen
	ch <- c.statsEpTapRebalanceQueueBackfillremaining
	ch <- c.statsEpTapRebalanceQueueBackoff
	ch <- c.statsEpTapRebalanceQueueDrain
	ch <- c.statsEpTapRebalanceQueueFill
	ch <- c.statsEpTapRebalanceQueueItemondisk
	ch <- c.statsEpTapRebalanceTotalBacklogSize
	ch <- c.statsEpTapReplicaCount
	ch <- c.statsEpTapReplicaQlen
	ch <- c.statsEpTapReplicaQueueBackfillremaining
	ch <- c.statsEpTapReplicaQueueBackoff
	ch <- c.statsEpTapReplicaQueueDrain
	ch <- c.statsEpTapReplicaQueueFill
	ch <- c.statsEpTapReplicaQueueItemondisk
	ch <- c.statsEpTapReplicaTotalBacklogSize
	ch <- c.statsEpTapTotalCount
	ch <- c.statsEpTapTotalQlen
	ch <- c.statsEpTapTotalQueueBackfillremaining
	ch <- c.statsEpTapTotalQueueBackoff
	ch <- c.statsEpTapTotalQueueDrain
	ch <- c.statsEpTapTotalQueueFill
	ch <- c.statsEpTapTotalQueueItemondisk
	ch <- c.statsEpTapTotalTotalBacklogSize
	ch <- c.statsEpTapUserCount
	ch <- c.statsEpTapUserQlen
	ch <- c.statsEpTapUserQueueBackfillremaining
	ch <- c.statsEpTapUserQueueBackoff
	ch <- c.statsEpTapUserQueueDrain
	ch <- c.statsEpTapUserQueueFill
	ch <- c.statsEpTapUserQueueItemondisk
	ch <- c.statsEpTapUserTotalBacklogSize
	ch <- c.statsEpTmpOomErrors
	ch <- c.statsEpVbTotal
	ch <- c.statsEvictions
	ch <- c.statsGetHits
	ch <- c.statsGetMisses
	ch <- c.statsHibernatedRequests
	ch <- c.statsHibernatedWaked
	ch <- c.statsHitRatio
	ch <- c.statsIncrHits
	ch <- c.statsIncrMisses
	ch <- c.statsMemActualFree
	ch <- c.statsMemActualUsed
	ch <- c.statsMemFree
	ch <- c.statsMemTotal
	ch <- c.statsMemUsed
	ch <- c.statsMemUsedSys
	ch <- c.statsMisses
	ch <- c.statsOps
	ch <- c.statsRestRequests
	ch <- c.statsSwapTotal
	ch <- c.statsSwapUsed
	ch <- c.statsTimestamp
	ch <- c.statsVbActiveEject
	ch <- c.statsVbActiveItmMemory
	ch <- c.statsVbActiveMetaDataMemory
	ch <- c.statsVbActiveNum
	ch <- c.statsVbActiveNumNonResident
	ch <- c.statsVbActiveOpsCreate
	ch <- c.statsVbActiveOpsUpdate
	ch <- c.statsVbActiveQueueAge
	ch <- c.statsVbActiveQueueDrain
	ch <- c.statsVbActiveQueueFill
	ch <- c.statsVbActiveQueueSize
	ch <- c.statsVbActiveResidentItemsRatio
	ch <- c.statsVbAvgActiveQueueAge
	ch <- c.statsVbAvgPendingQueueAge
	ch <- c.statsVbAvgReplicaQueueAge
	ch <- c.statsVbAvgTotalQueueAge
	ch <- c.statsVbPendingCurrItems
	ch <- c.statsVbPendingEject
	ch <- c.statsVbPendingItmMemory
	ch <- c.statsVbPendingMetaDataMemory
	ch <- c.statsVbPendingNum
	ch <- c.statsVbPendingNumNonResident
	ch <- c.statsVbPendingOpsCreate
	ch <- c.statsVbPendingOpsUpdate
	ch <- c.statsVbPendingQueueAge
	ch <- c.statsVbPendingQueueDrain
	ch <- c.statsVbPendingQueueFill
	ch <- c.statsVbPendingQueueSize
	ch <- c.statsVbPendingResidentItemsRatio
	ch <- c.statsVbReplicaCurrItems
	ch <- c.statsVbReplicaEject
	ch <- c.statsVbReplicaItmMemory
	ch <- c.statsVbReplicaMetaDataMemory
	ch <- c.statsVbReplicaNum
	ch <- c.statsVbReplicaNumNonResident
	ch <- c.statsVbReplicaOpsCreate
	ch <- c.statsVbReplicaOpsUpdate
	ch <- c.statsVbReplicaQueueAge
	ch <- c.statsVbReplicaQueueDrain
	ch <- c.statsVbReplicaQueueFill
	ch <- c.statsVbReplicaQueueSize
	ch <- c.statsVbReplicaResidentItemsRatio
	ch <- c.statsVbTotalQueueAge
	ch <- c.statsXdcOps
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

	for _, bucket := range buckets {
		log.Infof("Collecting %s bucket metrics...", bucket.Name)
		stats, err := c.client.BucketStats(bucket.Name)
		if err != nil {
			ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 0)
			log.With("error", err).Error("failed to scrape buckets stats")
			return
		}

		// TODO: collect bucket.Nodes metrics as well

		ch <- prometheus.MustNewConstMetric(c.basicstatsDataused, prometheus.GaugeValue, float64(bucket.BasicStats.DataUsed), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.basicstatsDiskfetches, prometheus.GaugeValue, float64(bucket.BasicStats.DiskFetchs), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.basicstatsDiskused, prometheus.GaugeValue, float64(bucket.BasicStats.DiskUsed), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.basicstatsItemcount, prometheus.GaugeValue, float64(bucket.BasicStats.ItemCount), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.basicstatsMemused, prometheus.GaugeValue, float64(bucket.BasicStats.MemUsed), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.basicstatsOpspersec, prometheus.GaugeValue, bucket.BasicStats.OpsPerSec, bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.basicstatsQuotapercentused, prometheus.GaugeValue, bucket.BasicStats.QuotaPercentUsed, bucket.Name)

		ch <- prometheus.MustNewConstMetric(c.statsAvgBgWaitTime, prometheus.GaugeValue, last(stats.Op.Samples.AvgBgWaitTime), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsAvgDiskCommitTime, prometheus.GaugeValue, last(stats.Op.Samples.AvgDiskCommitTime), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsAvgDiskUpdateTime, prometheus.GaugeValue, last(stats.Op.Samples.AvgDiskUpdateTime), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsBgWaitCount, prometheus.GaugeValue, last(stats.Op.Samples.BgWaitCount), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsBgWaitTotal, prometheus.GaugeValue, last(stats.Op.Samples.BgWaitTotal), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsBytesRead, prometheus.GaugeValue, last(stats.Op.Samples.BytesRead), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsBytesWritten, prometheus.GaugeValue, last(stats.Op.Samples.BytesWritten), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCasBadval, prometheus.GaugeValue, last(stats.Op.Samples.CasBadval), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCasHits, prometheus.GaugeValue, last(stats.Op.Samples.CasHits), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCasMisses, prometheus.GaugeValue, last(stats.Op.Samples.CasMisses), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCmdGet, prometheus.GaugeValue, last(stats.Op.Samples.CmdGet), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCmdSet, prometheus.GaugeValue, last(stats.Op.Samples.CmdSet), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCouchTotalDiskSize, prometheus.GaugeValue, last(stats.Op.Samples.CouchTotalDiskSize), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCouchDocsActualDiskSize, prometheus.GaugeValue, last(stats.Op.Samples.CouchDocsActualDiskSize), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCouchDocsFragmentation, prometheus.GaugeValue, last(stats.Op.Samples.CouchDocsFragmentation), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCouchDocsDataSize, prometheus.GaugeValue, last(stats.Op.Samples.CouchDocsDataSize), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCouchDocsDiskSize, prometheus.GaugeValue, last(stats.Op.Samples.CouchDocsDiskSize), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCpuIdleMs, prometheus.GaugeValue, last(stats.Op.Samples.CPUIdleMs), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCpuLocalMs, prometheus.GaugeValue, last(stats.Op.Samples.CPULocalMs), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCpuUtilizationRate, prometheus.GaugeValue, last(stats.Op.Samples.CPUUtilizationRate), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCurrConnections, prometheus.GaugeValue, last(stats.Op.Samples.CurrConnections), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCurrItems, prometheus.GaugeValue, last(stats.Op.Samples.CurrItems), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCurrItemsTot, prometheus.GaugeValue, last(stats.Op.Samples.CurrItemsTot), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsDecrHits, prometheus.GaugeValue, last(stats.Op.Samples.DecrHits), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsDecrMisses, prometheus.GaugeValue, last(stats.Op.Samples.DecrMisses), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsDeleteHits, prometheus.GaugeValue, last(stats.Op.Samples.DeleteHits), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsDeleteMisses, prometheus.GaugeValue, last(stats.Op.Samples.DeleteMisses), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsDiskCommitCount, prometheus.GaugeValue, last(stats.Op.Samples.DiskCommitCount), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsDiskCommitTotal, prometheus.GaugeValue, last(stats.Op.Samples.DiskCommitTotal), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsDiskUpdateCount, prometheus.GaugeValue, last(stats.Op.Samples.DiskUpdateCount), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsDiskUpdateTotal, prometheus.GaugeValue, last(stats.Op.Samples.DiskUpdateTotal), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsDiskWriteQueue, prometheus.GaugeValue, last(stats.Op.Samples.DiskWriteQueue), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpBgFetched, prometheus.GaugeValue, last(stats.Op.Samples.EpBgFetched), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpCacheMissRate, prometheus.GaugeValue, last(stats.Op.Samples.EpCacheMissRate), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcp2iBackoff, prometheus.GaugeValue, last(stats.Op.Samples.EpDcp2IBackoff), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcp2iCount, prometheus.GaugeValue, last(stats.Op.Samples.EpDcp2ICount), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcp2iItemsRemaining, prometheus.GaugeValue, last(stats.Op.Samples.EpDcp2IItemsRemaining), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcp2iItemsSent, prometheus.GaugeValue, last(stats.Op.Samples.EpDcp2IItemsSent), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcp2iProducerCount, prometheus.GaugeValue, last(stats.Op.Samples.EpDcp2IProducerCount), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcp2iTotalBacklogSize, prometheus.GaugeValue, last(stats.Op.Samples.EpDcp2ITotalBacklogSize), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcp2iTotalBytes, prometheus.GaugeValue, last(stats.Op.Samples.EpDcp2ITotalBytes), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcpOtherBackoff, prometheus.GaugeValue, last(stats.Op.Samples.EpDcpOtherBackoff), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcpOtherCount, prometheus.GaugeValue, last(stats.Op.Samples.EpDcpOtherCount), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcpOtherItemsRemaining, prometheus.GaugeValue, last(stats.Op.Samples.EpDcpOtherItemsRemaining), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcpOtherItemsSent, prometheus.GaugeValue, last(stats.Op.Samples.EpDcpOtherItemsSent), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcpOtherProducerCount, prometheus.GaugeValue, last(stats.Op.Samples.EpDcpOtherProducerCount), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcpOtherTotalBacklogSize, prometheus.GaugeValue, last(stats.Op.Samples.EpDcpOtherTotalBacklogSize), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcpOtherTotalBytes, prometheus.GaugeValue, last(stats.Op.Samples.EpDcpOtherTotalBytes), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcpReplicaBackoff, prometheus.GaugeValue, last(stats.Op.Samples.EpDcpReplicaBackoff), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcpReplicaCount, prometheus.GaugeValue, last(stats.Op.Samples.EpDcpReplicaCount), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcpReplicaItemsRemaining, prometheus.GaugeValue, last(stats.Op.Samples.EpDcpReplicaItemsRemaining), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcpReplicaItemsSent, prometheus.GaugeValue, last(stats.Op.Samples.EpDcpReplicaItemsSent), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcpReplicaProducerCount, prometheus.GaugeValue, last(stats.Op.Samples.EpDcpReplicaProducerCount), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcpReplicaTotalBacklogSize, prometheus.GaugeValue, last(stats.Op.Samples.EpDcpReplicaTotalBacklogSize), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcpReplicaTotalBytes, prometheus.GaugeValue, last(stats.Op.Samples.EpDcpReplicaTotalBytes), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcpViewsBackoff, prometheus.GaugeValue, last(stats.Op.Samples.EpDcpViewsBackoff), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcpViewsCount, prometheus.GaugeValue, last(stats.Op.Samples.EpDcpViewsCount), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcpViewsItemsRemaining, prometheus.GaugeValue, last(stats.Op.Samples.EpDcpViewsItemsRemaining), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcpViewsItemsSent, prometheus.GaugeValue, last(stats.Op.Samples.EpDcpViewsItemsSent), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcpViewsProducerCount, prometheus.GaugeValue, last(stats.Op.Samples.EpDcpViewsProducerCount), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcpViewsTotalBacklogSize, prometheus.GaugeValue, last(stats.Op.Samples.EpDcpViewsTotalBacklogSize), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcpViewsTotalBytes, prometheus.GaugeValue, last(stats.Op.Samples.EpDcpViewsTotalBytes), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcpXdcrBackoff, prometheus.GaugeValue, last(stats.Op.Samples.EpDcpXdcrBackoff), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcpXdcrCount, prometheus.GaugeValue, last(stats.Op.Samples.EpDcpXdcrCount), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcpXdcrItemsRemaining, prometheus.GaugeValue, last(stats.Op.Samples.EpDcpXdcrItemsRemaining), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcpXdcrItemsSent, prometheus.GaugeValue, last(stats.Op.Samples.EpDcpXdcrItemsSent), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcpXdcrProducerCount, prometheus.GaugeValue, last(stats.Op.Samples.EpDcpXdcrProducerCount), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcpXdcrTotalBacklogSize, prometheus.GaugeValue, last(stats.Op.Samples.EpDcpXdcrTotalBacklogSize), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDcpXdcrTotalBytes, prometheus.GaugeValue, last(stats.Op.Samples.EpDcpXdcrTotalBytes), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDiskqueueDrain, prometheus.GaugeValue, last(stats.Op.Samples.EpDiskqueueDrain), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDiskqueueFill, prometheus.GaugeValue, last(stats.Op.Samples.EpDiskqueueFill), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpDiskqueueItems, prometheus.GaugeValue, last(stats.Op.Samples.EpDiskqueueItems), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpFlusherTodo, prometheus.GaugeValue, last(stats.Op.Samples.EpFlusherTodo), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpItemCommitFailed, prometheus.GaugeValue, last(stats.Op.Samples.EpItemCommitFailed), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpKvSize, prometheus.GaugeValue, last(stats.Op.Samples.EpKvSize), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpMaxSize, prometheus.GaugeValue, last(stats.Op.Samples.EpMaxSize), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpMemHighWat, prometheus.GaugeValue, last(stats.Op.Samples.EpMemHighWat), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpMemLowWat, prometheus.GaugeValue, last(stats.Op.Samples.EpMemLowWat), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpMetaDataMemory, prometheus.GaugeValue, last(stats.Op.Samples.EpMetaDataMemory), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpNumNonResident, prometheus.GaugeValue, last(stats.Op.Samples.EpNumNonResident), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpNumOpsDelMeta, prometheus.GaugeValue, last(stats.Op.Samples.EpNumOpsDelMeta), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpNumOpsDelRetMeta, prometheus.GaugeValue, last(stats.Op.Samples.EpNumOpsDelRetMeta), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpNumOpsGetMeta, prometheus.GaugeValue, last(stats.Op.Samples.EpNumOpsGetMeta), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpNumOpsSetMeta, prometheus.GaugeValue, last(stats.Op.Samples.EpNumOpsSetMeta), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpNumOpsSetRetMeta, prometheus.GaugeValue, last(stats.Op.Samples.EpNumOpsSetRetMeta), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpNumValueEjects, prometheus.GaugeValue, last(stats.Op.Samples.EpNumValueEjects), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpOomErrors, prometheus.GaugeValue, last(stats.Op.Samples.EpOomErrors), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpOpsCreate, prometheus.GaugeValue, last(stats.Op.Samples.EpOpsCreate), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpOpsUpdate, prometheus.GaugeValue, last(stats.Op.Samples.EpOpsUpdate), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpOverhead, prometheus.GaugeValue, last(stats.Op.Samples.EpOverhead), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpQueueSize, prometheus.GaugeValue, last(stats.Op.Samples.EpQueueSize), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpResidentItemsRate, prometheus.GaugeValue, last(stats.Op.Samples.EpResidentItemsRate), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapRebalanceCount, prometheus.GaugeValue, last(stats.Op.Samples.EpTapRebalanceCount), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapRebalanceQlen, prometheus.GaugeValue, last(stats.Op.Samples.EpTapRebalanceQlen), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapRebalanceQueueBackfillremaining, prometheus.GaugeValue, last(stats.Op.Samples.EpTapRebalanceQueueBackfillremaining), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapRebalanceQueueBackoff, prometheus.GaugeValue, last(stats.Op.Samples.EpTapRebalanceQueueBackoff), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapRebalanceQueueDrain, prometheus.GaugeValue, last(stats.Op.Samples.EpTapRebalanceQueueDrain), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapRebalanceQueueFill, prometheus.GaugeValue, last(stats.Op.Samples.EpTapRebalanceQueueFill), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapRebalanceQueueItemondisk, prometheus.GaugeValue, last(stats.Op.Samples.EpTapRebalanceQueueItemondisk), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapRebalanceTotalBacklogSize, prometheus.GaugeValue, last(stats.Op.Samples.EpTapRebalanceTotalBacklogSize), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapReplicaCount, prometheus.GaugeValue, last(stats.Op.Samples.EpTapReplicaCount), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapReplicaQlen, prometheus.GaugeValue, last(stats.Op.Samples.EpTapReplicaQlen), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapReplicaQueueBackfillremaining, prometheus.GaugeValue, last(stats.Op.Samples.EpTapReplicaQueueBackfillremaining), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapReplicaQueueBackoff, prometheus.GaugeValue, last(stats.Op.Samples.EpTapReplicaQueueBackoff), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapReplicaQueueDrain, prometheus.GaugeValue, last(stats.Op.Samples.EpTapReplicaQueueDrain), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapReplicaQueueFill, prometheus.GaugeValue, last(stats.Op.Samples.EpTapReplicaQueueFill), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapReplicaQueueItemondisk, prometheus.GaugeValue, last(stats.Op.Samples.EpTapReplicaQueueItemondisk), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapReplicaTotalBacklogSize, prometheus.GaugeValue, last(stats.Op.Samples.EpTapReplicaTotalBacklogSize), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapTotalCount, prometheus.GaugeValue, last(stats.Op.Samples.EpTapTotalCount), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapTotalQlen, prometheus.GaugeValue, last(stats.Op.Samples.EpTapTotalQlen), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapTotalQueueBackfillremaining, prometheus.GaugeValue, last(stats.Op.Samples.EpTapTotalQueueBackfillremaining), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapTotalQueueBackoff, prometheus.GaugeValue, last(stats.Op.Samples.EpTapTotalQueueBackoff), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapTotalQueueDrain, prometheus.GaugeValue, last(stats.Op.Samples.EpTapTotalQueueDrain), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapTotalQueueFill, prometheus.GaugeValue, last(stats.Op.Samples.EpTapTotalQueueFill), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapTotalQueueItemondisk, prometheus.GaugeValue, last(stats.Op.Samples.EpTapTotalQueueItemondisk), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapTotalTotalBacklogSize, prometheus.GaugeValue, last(stats.Op.Samples.EpTapTotalTotalBacklogSize), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapUserCount, prometheus.GaugeValue, last(stats.Op.Samples.EpTapUserCount), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapUserQlen, prometheus.GaugeValue, last(stats.Op.Samples.EpTapUserQlen), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapUserQueueBackfillremaining, prometheus.GaugeValue, last(stats.Op.Samples.EpTapUserQueueBackfillremaining), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapUserQueueBackoff, prometheus.GaugeValue, last(stats.Op.Samples.EpTapUserQueueBackoff), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapUserQueueDrain, prometheus.GaugeValue, last(stats.Op.Samples.EpTapUserQueueDrain), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapUserQueueFill, prometheus.GaugeValue, last(stats.Op.Samples.EpTapUserQueueFill), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapUserQueueItemondisk, prometheus.GaugeValue, last(stats.Op.Samples.EpTapUserQueueItemondisk), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTapUserTotalBacklogSize, prometheus.GaugeValue, last(stats.Op.Samples.EpTapUserTotalBacklogSize), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpTmpOomErrors, prometheus.GaugeValue, last(stats.Op.Samples.EpTmpOomErrors), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpVbTotal, prometheus.GaugeValue, last(stats.Op.Samples.EpVbTotal), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEvictions, prometheus.GaugeValue, last(stats.Op.Samples.Evictions), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsGetHits, prometheus.GaugeValue, last(stats.Op.Samples.GetHits), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsGetMisses, prometheus.GaugeValue, last(stats.Op.Samples.GetMisses), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsHibernatedRequests, prometheus.GaugeValue, last(stats.Op.Samples.HibernatedRequests), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsHibernatedWaked, prometheus.GaugeValue, last(stats.Op.Samples.HibernatedWaked), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsHitRatio, prometheus.GaugeValue, last(stats.Op.Samples.HitRatio), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsIncrHits, prometheus.GaugeValue, last(stats.Op.Samples.IncrHits), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsIncrMisses, prometheus.GaugeValue, last(stats.Op.Samples.IncrMisses), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsMemActualFree, prometheus.GaugeValue, last(stats.Op.Samples.MemActualFree), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsMemActualUsed, prometheus.GaugeValue, last(stats.Op.Samples.MemActualUsed), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsMemFree, prometheus.GaugeValue, last(stats.Op.Samples.MemFree), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsMemTotal, prometheus.GaugeValue, last(stats.Op.Samples.MemTotal), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsMemUsed, prometheus.GaugeValue, last(stats.Op.Samples.MemUsed), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsMemUsedSys, prometheus.GaugeValue, last(stats.Op.Samples.MemUsedSys), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsMisses, prometheus.GaugeValue, last(stats.Op.Samples.Misses), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsOps, prometheus.GaugeValue, last(stats.Op.Samples.Ops), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsRestRequests, prometheus.GaugeValue, last(stats.Op.Samples.RestRequests), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsSwapTotal, prometheus.GaugeValue, last(stats.Op.Samples.SwapTotal), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsSwapUsed, prometheus.GaugeValue, last(stats.Op.Samples.SwapUsed), bucket.Name)
		// ch <- prometheus.MustNewConstMetric(c.statsTimestamp, prometheus.GaugeValue, last(stats.Op.Samples.Timestamp), bucket.Name) TODO
		ch <- prometheus.MustNewConstMetric(c.statsVbActiveEject, prometheus.GaugeValue, last(stats.Op.Samples.VbActiveEject), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbActiveItmMemory, prometheus.GaugeValue, last(stats.Op.Samples.VbActiveItmMemory), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbActiveMetaDataMemory, prometheus.GaugeValue, last(stats.Op.Samples.VbActiveMetaDataMemory), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbActiveNum, prometheus.GaugeValue, last(stats.Op.Samples.VbActiveNum), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbActiveNumNonResident, prometheus.GaugeValue, last(stats.Op.Samples.VbActiveNumNonResident), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbActiveOpsCreate, prometheus.GaugeValue, last(stats.Op.Samples.VbActiveOpsCreate), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbActiveOpsUpdate, prometheus.GaugeValue, last(stats.Op.Samples.VbActiveOpsUpdate), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbActiveQueueAge, prometheus.GaugeValue, last(stats.Op.Samples.VbActiveQueueAge), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbActiveQueueDrain, prometheus.GaugeValue, last(stats.Op.Samples.VbActiveQueueDrain), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbActiveQueueFill, prometheus.GaugeValue, last(stats.Op.Samples.VbActiveQueueFill), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbActiveQueueSize, prometheus.GaugeValue, last(stats.Op.Samples.VbActiveQueueSize), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbActiveResidentItemsRatio, prometheus.GaugeValue, last(stats.Op.Samples.VbActiveResidentItemsRatio), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbAvgActiveQueueAge, prometheus.GaugeValue, last(stats.Op.Samples.VbAvgActiveQueueAge), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbAvgPendingQueueAge, prometheus.GaugeValue, last(stats.Op.Samples.VbAvgPendingQueueAge), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbAvgReplicaQueueAge, prometheus.GaugeValue, last(stats.Op.Samples.VbAvgReplicaQueueAge), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbAvgTotalQueueAge, prometheus.GaugeValue, last(stats.Op.Samples.VbAvgTotalQueueAge), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbPendingCurrItems, prometheus.GaugeValue, last(stats.Op.Samples.VbPendingCurrItems), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbPendingEject, prometheus.GaugeValue, last(stats.Op.Samples.VbPendingEject), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbPendingItmMemory, prometheus.GaugeValue, last(stats.Op.Samples.VbPendingItmMemory), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbPendingMetaDataMemory, prometheus.GaugeValue, last(stats.Op.Samples.VbPendingMetaDataMemory), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbPendingNum, prometheus.GaugeValue, last(stats.Op.Samples.VbPendingNum), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbPendingNumNonResident, prometheus.GaugeValue, last(stats.Op.Samples.VbPendingNumNonResident), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbPendingOpsCreate, prometheus.GaugeValue, last(stats.Op.Samples.VbPendingOpsCreate), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbPendingOpsUpdate, prometheus.GaugeValue, last(stats.Op.Samples.VbPendingOpsUpdate), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbPendingQueueAge, prometheus.GaugeValue, last(stats.Op.Samples.VbPendingQueueAge), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbPendingQueueDrain, prometheus.GaugeValue, last(stats.Op.Samples.VbPendingQueueDrain), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbPendingQueueFill, prometheus.GaugeValue, last(stats.Op.Samples.VbPendingQueueFill), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbPendingQueueSize, prometheus.GaugeValue, last(stats.Op.Samples.VbPendingQueueSize), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbPendingResidentItemsRatio, prometheus.GaugeValue, last(stats.Op.Samples.VbPendingResidentItemsRatio), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbReplicaCurrItems, prometheus.GaugeValue, last(stats.Op.Samples.VbReplicaCurrItems), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbReplicaEject, prometheus.GaugeValue, last(stats.Op.Samples.VbReplicaEject), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbReplicaItmMemory, prometheus.GaugeValue, last(stats.Op.Samples.VbReplicaItmMemory), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbReplicaMetaDataMemory, prometheus.GaugeValue, last(stats.Op.Samples.VbReplicaMetaDataMemory), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbReplicaNum, prometheus.GaugeValue, last(stats.Op.Samples.VbReplicaNum), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbReplicaNumNonResident, prometheus.GaugeValue, last(stats.Op.Samples.VbReplicaNumNonResident), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbReplicaOpsCreate, prometheus.GaugeValue, last(stats.Op.Samples.VbReplicaOpsCreate), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbReplicaOpsUpdate, prometheus.GaugeValue, last(stats.Op.Samples.VbReplicaOpsUpdate), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbReplicaQueueAge, prometheus.GaugeValue, last(stats.Op.Samples.VbReplicaQueueAge), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbReplicaQueueDrain, prometheus.GaugeValue, last(stats.Op.Samples.VbReplicaQueueDrain), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbReplicaQueueFill, prometheus.GaugeValue, last(stats.Op.Samples.VbReplicaQueueFill), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbReplicaQueueSize, prometheus.GaugeValue, last(stats.Op.Samples.VbReplicaQueueSize), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbReplicaResidentItemsRatio, prometheus.GaugeValue, last(stats.Op.Samples.VbReplicaResidentItemsRatio), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsVbTotalQueueAge, prometheus.GaugeValue, last(stats.Op.Samples.VbTotalQueueAge), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsXdcOps, prometheus.GaugeValue, last(stats.Op.Samples.XdcOps), bucket.Name)
	}

	ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 1)
	ch <- prometheus.MustNewConstMetric(c.scrapeDuration, prometheus.GaugeValue, time.Since(start).Seconds())
}
