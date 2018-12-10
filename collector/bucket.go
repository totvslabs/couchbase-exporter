package collector

import (
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"github.com/totvslabs/couchbase-exporter/client"
)

type bucketsCollector struct {
	mutex  sync.Mutex
	client client.Client

	up             *prometheus.Desc
	scrapeDuration *prometheus.Desc

	basicstatsDataused         *prometheus.Desc
	basicstatsDiskfetches      *prometheus.Desc
	basicstatsDiskused         *prometheus.Desc
	basicstatsItemcount        *prometheus.Desc
	basicstatsMemused          *prometheus.Desc
	basicstatsOpspersec        *prometheus.Desc
	basicstatsQuotapercentused *prometheus.Desc

	statsAvgBgWaitTime                    *prometheus.Desc
	statsAvgActiveTimestampDrift          *prometheus.Desc
	statsAvgReplicaTimestampDrift         *prometheus.Desc
	statsAvgDiskCommitTime                *prometheus.Desc
	statsAvgDiskUpdateTime                *prometheus.Desc
	statsBytesRead                        *prometheus.Desc
	statsBytesWritten                     *prometheus.Desc
	statsCasBadval                        *prometheus.Desc
	statsCasHits                          *prometheus.Desc
	statsCasMisses                        *prometheus.Desc
	statsCmdGet                           *prometheus.Desc
	statsCmdSet                           *prometheus.Desc
	statsCouchTotalDiskSize               *prometheus.Desc
	statsCouchViewsDataSize               *prometheus.Desc
	statsCouchViewsActualDiskSize         *prometheus.Desc
	statsCouchViewsFragmentation          *prometheus.Desc
	statsCouchViewsOps                    *prometheus.Desc
	statsCouchDocsDataSize                *prometheus.Desc
	statsCouchDocsDiskSize                *prometheus.Desc
	statsCouchDocsActualDiskSize          *prometheus.Desc
	statsCouchDocsFragmentation           *prometheus.Desc
	statsCPUIdleMs                        *prometheus.Desc
	statsCPULocalMs                       *prometheus.Desc
	statsCPUUtilizationRate               *prometheus.Desc
	statsCurrConnections                  *prometheus.Desc
	statsCurrItems                        *prometheus.Desc
	statsCurrItemsTot                     *prometheus.Desc
	statsDecrHits                         *prometheus.Desc
	statsDecrMisses                       *prometheus.Desc
	statsDeleteHits                       *prometheus.Desc
	statsDeleteMisses                     *prometheus.Desc
	statsDiskCommitCount                  *prometheus.Desc
	statsDiskUpdateCount                  *prometheus.Desc
	statsDiskWriteQueue                   *prometheus.Desc
	statsEpActiveAheadExceptions          *prometheus.Desc
	statsEpActiveHlcDrift                 *prometheus.Desc
	statsEpClockCasDriftThresholdExceeded *prometheus.Desc
	statsEpBgFetched                      *prometheus.Desc
	statsEpCacheMissRate                  *prometheus.Desc
	statsEpDcp2iBackoff                   *prometheus.Desc
	statsEpDcp2iCount                     *prometheus.Desc
	statsEpDcp2iItemsRemaining            *prometheus.Desc
	statsEpDcp2iItemsSent                 *prometheus.Desc
	statsEpDcp2iProducerCount             *prometheus.Desc
	statsEpDcp2iTotalBacklogSize          *prometheus.Desc
	statsEpDcp2iTotalBytes                *prometheus.Desc
	statsEpDcpOtherBackoff                *prometheus.Desc
	statsEpDcpOtherCount                  *prometheus.Desc
	statsEpDcpOtherItemsRemaining         *prometheus.Desc
	statsEpDcpOtherItemsSent              *prometheus.Desc
	statsEpDcpOtherProducerCount          *prometheus.Desc
	statsEpDcpOtherTotalBacklogSize       *prometheus.Desc
	statsEpDcpOtherTotalBytes             *prometheus.Desc
	statsEpDcpReplicaBackoff              *prometheus.Desc
	statsEpDcpReplicaCount                *prometheus.Desc
	statsEpDcpReplicaItemsRemaining       *prometheus.Desc
	statsEpDcpReplicaItemsSent            *prometheus.Desc
	statsEpDcpReplicaProducerCount        *prometheus.Desc
	statsEpDcpReplicaTotalBacklogSize     *prometheus.Desc
	statsEpDcpReplicaTotalBytes           *prometheus.Desc
	statsEpDcpViewsBackoff                *prometheus.Desc
	statsEpDcpViewsCount                  *prometheus.Desc
	statsEpDcpViewsItemsRemaining         *prometheus.Desc
	statsEpDcpViewsItemsSent              *prometheus.Desc
	statsEpDcpViewsProducerCount          *prometheus.Desc
	statsEpDcpViewsTotalBacklogSize       *prometheus.Desc
	statsEpDcpViewsTotalBytes             *prometheus.Desc
	statsEpDcpXdcrBackoff                 *prometheus.Desc
	statsEpDcpXdcrCount                   *prometheus.Desc
	statsEpDcpXdcrItemsRemaining          *prometheus.Desc
	statsEpDcpXdcrItemsSent               *prometheus.Desc
	statsEpDcpXdcrProducerCount           *prometheus.Desc
	statsEpDcpXdcrTotalBacklogSize        *prometheus.Desc
	statsEpDcpXdcrTotalBytes              *prometheus.Desc
	statsEpDiskqueueDrain                 *prometheus.Desc
	statsEpDiskqueueFill                  *prometheus.Desc
	statsEpDiskqueueItems                 *prometheus.Desc
	statsEpFlusherTodo                    *prometheus.Desc
	statsEpItemCommitFailed               *prometheus.Desc
	statsEpKvSize                         *prometheus.Desc
	statsEpMaxSize                        *prometheus.Desc
	statsEpMemHighWat                     *prometheus.Desc
	statsEpMemLowWat                      *prometheus.Desc
	statsEpMetaDataMemory                 *prometheus.Desc
	statsEpNumNonResident                 *prometheus.Desc
	statsEpNumOpsDelMeta                  *prometheus.Desc
	statsEpNumOpsDelRetMeta               *prometheus.Desc
	statsEpNumOpsGetMeta                  *prometheus.Desc
	statsEpNumOpsSetMeta                  *prometheus.Desc
	statsEpNumOpsSetRetMeta               *prometheus.Desc
	statsEpNumValueEjects                 *prometheus.Desc
	statsEpOomErrors                      *prometheus.Desc
	statsEpOpsCreate                      *prometheus.Desc
	statsEpOpsUpdate                      *prometheus.Desc
	statsEpOverhead                       *prometheus.Desc
	statsEpQueueSize                      *prometheus.Desc
	statsEpResidentItemsRate              *prometheus.Desc
	statsEpReplicaAheadExceptions         *prometheus.Desc
	statsEpReplicaHlcDrift                *prometheus.Desc
	statsEpTmpOomErrors                   *prometheus.Desc
	statsEpVbTotal                        *prometheus.Desc
	statsEvictions                        *prometheus.Desc
	statsGetHits                          *prometheus.Desc
	statsGetMisses                        *prometheus.Desc
	statsHibernatedRequests               *prometheus.Desc
	statsHibernatedWaked                  *prometheus.Desc
	statsHitRatio                         *prometheus.Desc
	statsIncrHits                         *prometheus.Desc
	statsIncrMisses                       *prometheus.Desc
	statsMemActualFree                    *prometheus.Desc
	statsMemActualUsed                    *prometheus.Desc
	statsMemFree                          *prometheus.Desc
	statsMemTotal                         *prometheus.Desc
	statsMemUsed                          *prometheus.Desc
	statsMemUsedSys                       *prometheus.Desc
	statsMisses                           *prometheus.Desc
	statsOps                              *prometheus.Desc
	statsRestRequests                     *prometheus.Desc
	statsSwapTotal                        *prometheus.Desc
	statsSwapUsed                         *prometheus.Desc
	statsTimestamp                        *prometheus.Desc
	statsVbActiveEject                    *prometheus.Desc
	statsVbActiveItmMemory                *prometheus.Desc
	statsVbActiveMetaDataMemory           *prometheus.Desc
	statsVbActiveNum                      *prometheus.Desc
	statsVbActiveNumNonResident           *prometheus.Desc
	statsVbActiveOpsCreate                *prometheus.Desc
	statsVbActiveOpsUpdate                *prometheus.Desc
	statsVbActiveQueueAge                 *prometheus.Desc
	statsVbActiveQueueDrain               *prometheus.Desc
	statsVbActiveQueueFill                *prometheus.Desc
	statsVbActiveQueueSize                *prometheus.Desc
	statsVbActiveResidentItemsRatio       *prometheus.Desc
	statsVbAvgActiveQueueAge              *prometheus.Desc
	statsVbAvgPendingQueueAge             *prometheus.Desc
	statsVbAvgReplicaQueueAge             *prometheus.Desc
	statsVbAvgTotalQueueAge               *prometheus.Desc
	statsVbPendingCurrItems               *prometheus.Desc
	statsVbPendingEject                   *prometheus.Desc
	statsVbPendingItmMemory               *prometheus.Desc
	statsVbPendingMetaDataMemory          *prometheus.Desc
	statsVbPendingNum                     *prometheus.Desc
	statsVbPendingNumNonResident          *prometheus.Desc
	statsVbPendingOpsCreate               *prometheus.Desc
	statsVbPendingOpsUpdate               *prometheus.Desc
	statsVbPendingQueueAge                *prometheus.Desc
	statsVbPendingQueueDrain              *prometheus.Desc
	statsVbPendingQueueFill               *prometheus.Desc
	statsVbPendingQueueSize               *prometheus.Desc
	statsVbPendingResidentItemsRatio      *prometheus.Desc
	statsVbReplicaCurrItems               *prometheus.Desc
	statsVbReplicaEject                   *prometheus.Desc
	statsVbReplicaItmMemory               *prometheus.Desc
	statsVbReplicaMetaDataMemory          *prometheus.Desc
	statsVbReplicaNum                     *prometheus.Desc
	statsVbReplicaNumNonResident          *prometheus.Desc
	statsVbReplicaOpsCreate               *prometheus.Desc
	statsVbReplicaOpsUpdate               *prometheus.Desc
	statsVbReplicaQueueAge                *prometheus.Desc
	statsVbReplicaQueueDrain              *prometheus.Desc
	statsVbReplicaQueueFill               *prometheus.Desc
	statsVbReplicaQueueSize               *prometheus.Desc
	statsVbReplicaResidentItemsRatio      *prometheus.Desc
	statsVbTotalQueueAge                  *prometheus.Desc
	statsXdcOps                           *prometheus.Desc
}

// NewBucketsCollector buckets collector
func NewBucketsCollector(client client.Client) prometheus.Collector {
	const subsystem = "bucket"
	// nolint: lll
	return &bucketsCollector{
		client: client,
		up: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "up"),
			"Couchbase buckets API is responding",
			nil,
			nil,
		),
		scrapeDuration: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "scrape_duration_seconds"),
			"Scrape duration in seconds",
			nil,
			nil,
		),
		basicstatsDataused: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "basicstats_dataused_bytes"),
			"basicstats_dataused",
			[]string{"bucket"},
			nil,
		),
		basicstatsDiskfetches: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "basicstats_diskfetches"),
			"basicstats_diskfetches",
			[]string{"bucket"},
			nil,
		),
		basicstatsDiskused: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "basicstats_diskused_bytes"),
			"basicstats_diskused",
			[]string{"bucket"},
			nil,
		),
		basicstatsItemcount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "basicstats_itemcount"),
			"basicstats_itemcount",
			[]string{"bucket"},
			nil,
		),
		basicstatsMemused: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "basicstats_memused_bytes"),
			"basicstats_memused",
			[]string{"bucket"},
			nil,
		),
		basicstatsOpspersec: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "basicstats_opspersec"),
			"basicstats_opspersec",
			[]string{"bucket"},
			nil,
		),
		basicstatsQuotapercentused: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "basicstats_quota_user_percent"),
			"basicstats_quotapercentused",
			[]string{"bucket"},
			nil,
		),
		statsAvgBgWaitTime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_avg_bg_wait_seconds"),
			"Average background fetch time in seconds",
			[]string{"bucket"},
			nil,
		),
		statsAvgActiveTimestampDrift: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_avg_active_timestamp_drift"),
			"avg_active_timestamp_drift",
			[]string{"bucket"},
			nil,
		),
		statsAvgReplicaTimestampDrift: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "avg_replica_timestamp_drift"),
			"avg_replica_timestamp_drift",
			[]string{"bucket"},
			nil,
		),
		statsAvgDiskCommitTime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_avg_disk_commit_time"),
			"Average disk commit time in seconds as from disk_update histogram of timings",
			[]string{"bucket"},
			nil,
		),
		statsAvgDiskUpdateTime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_avg_disk_update_time"),
			"Average disk update time in microseconds as from disk_update histogram of timings",
			[]string{"bucket"},
			nil,
		),
		statsBytesRead: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_read_bytes"),
			"Bytes read",
			[]string{"bucket"},
			nil,
		),
		statsBytesWritten: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_written_bytes"),
			"Bytes written",
			[]string{"bucket"},
			nil,
		),
		statsCasBadval: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_cas_badval"),
			"Compare and Swap bad values",
			[]string{"bucket"},
			nil,
		),
		statsCasHits: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_cas_hits"),
			"Number of operations with a CAS id per second for this bucket",
			[]string{"bucket"},
			nil,
		),
		statsCasMisses: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_cas_misses"),
			"Compare and Swap misses",
			[]string{"bucket"},
			nil,
		),
		statsCmdGet: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_cmd_get"),
			"Number of reads (get operations) per second from this bucket",
			[]string{"bucket"},
			nil,
		),
		statsCmdSet: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_cmd_set"),
			"Number of writes (set operations) per second to this bucket",
			[]string{"bucket"},
			nil,
		),
		statsCouchTotalDiskSize: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "couch_total_disk_size"),
			"The total size on disk of all data and view files for this bucket",
			[]string{"bucket"},
			nil,
		),
		statsCouchViewsFragmentation: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "couch_views_fragmentation"),
			"How much fragmented data there is to be compacted compared to real data for the view index files in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsCouchViewsOps: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "couch_views_ops"),
			"All the view reads for all design documents including scatter gather",
			[]string{"bucket"},
			nil,
		),
		statsCouchViewsDataSize: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "couch_views_data_size"),
			"The size of active data on for all the indexes in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsCouchViewsActualDiskSize: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "couch_views_actual_disk_size"),
			"The size of all active items in all the indexes for this bucket on disk",
			[]string{"bucket"},
			nil,
		),
		statsCouchDocsFragmentation: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "couch_docs_fragmentation"),
			"How much fragmented data there is to be compacted compared to real data for the data files in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsCouchDocsActualDiskSize: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "couch_docs_actual_disk_size"),
			"The size of all data files for this bucket, including the data itself, meta data and temporary files",
			[]string{"bucket"},
			nil,
		),
		statsCouchDocsDataSize: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_couch_docs_data_size"),
			"The size of active data in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsCouchDocsDiskSize: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_couch_docs_disk_size"),
			"The size of all data files for this bucket, including the data itself, meta data and temporary files",
			[]string{"bucket"},
			nil,
		),
		statsCPUIdleMs: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_cpu_idle_ms"),
			"CPU idle milliseconds",
			[]string{"bucket"},
			nil,
		),
		statsCPULocalMs: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_cpu_local_ms"),
			"stats_cpu_local_ms",
			[]string{"bucket"},
			nil,
		),
		statsCPUUtilizationRate: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_cpu_utilization_rate"),
			"Percentage of CPU in use across all available cores on this server",
			[]string{"bucket"},
			nil,
		),
		statsCurrConnections: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_curr_connections"),
			"Number of connections to this server includingconnections from external client SDKs, proxies, TAP requests and internal statistic gathering",
			[]string{"bucket"},
			nil,
		),
		statsCurrItems: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_curr_items"),
			"Number of items in active vBuckets in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsCurrItemsTot: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_curr_items_tot"),
			"Total number of items in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsDecrHits: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_decr_hits"),
			"Decrement hits",
			[]string{"bucket"},
			nil,
		),
		statsDecrMisses: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_decr_misses"),
			"Decrement misses",
			[]string{"bucket"},
			nil,
		),
		statsDeleteHits: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_delete_hits"),
			"Number of delete operations per second for this bucket",
			[]string{"bucket"},
			nil,
		),
		statsDeleteMisses: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_delete_misses"),
			"Delete misses",
			[]string{"bucket"},
			nil,
		),
		statsDiskCommitCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_disk_commits"),
			"Disk commits",
			[]string{"bucket"},
			nil,
		),
		statsDiskUpdateCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_disk_updates"),
			"Disk updates",
			[]string{"bucket"},
			nil,
		),
		statsDiskWriteQueue: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_disk_write_queue"),
			"Number of items waiting to be written to disk in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpActiveAheadExceptions: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_active_ahead_exceptions"),
			"stats_ep_active_ahead_exceptions",
			[]string{"bucket"},
			nil,
		),
		statsEpActiveHlcDrift: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_active_hlc_drift"),
			"stats_ep_active_hlc_drift",
			[]string{"bucket"},
			nil,
		),
		statsEpClockCasDriftThresholdExceeded: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_clock_cas_drift_threshold_exceeded"),
			"stats_ep_clock_cas_drift_threshold_exceeded",
			[]string{"bucket"},
			nil,
		),
		statsEpBgFetched: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_bg_fetched"),
			"Number of reads per second from disk for this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpCacheMissRate: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_cache_miss_rate"),
			"Percentage of reads per second to this bucket from disk as opposed to RAM",
			[]string{"bucket"},
			nil,
		),
		statsEpDcp2iBackoff: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_2i_backoff"),
			"Number of backoffs for indexes DCP connections",
			[]string{"bucket"},
			nil,
		),
		statsEpDcp2iCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_2i_connections"),
			"Number of indexes DCP connections",
			[]string{"bucket"},
			nil,
		),
		statsEpDcp2iItemsRemaining: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_2i_items_remaining"),
			"Number of indexes items remaining to be sent",
			[]string{"bucket"},
			nil,
		),
		statsEpDcp2iItemsSent: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_2i_items_sent"),
			"Number of indexes items sent",
			[]string{"bucket"},
			nil,
		),
		statsEpDcp2iProducerCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_2i_producers"),
			"Number of indexes producers",
			[]string{"bucket"},
			nil,
		),
		statsEpDcp2iTotalBacklogSize: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_2i_total_backlog_size"),
			"stats_ep_dcp_2i_total_backlog_size",
			[]string{"bucket"},
			nil,
		),
		statsEpDcp2iTotalBytes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_2i_total_bytes"),
			"Number bytes per second being sent for indexes DCP connections",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpOtherBackoff: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_other_backoff"),
			"Number of backoffs for other DCP connections",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpOtherCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_others"),
			"Number of other DCP connections in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpOtherItemsRemaining: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_other_items_remaining"),
			"Number of items remaining to be sent to consumer in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpOtherItemsSent: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_other_items_sent"),
			"Number of items per second being sent for a producer for this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpOtherProducerCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_other_producers"),
			"Number of other senders for this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpOtherTotalBacklogSize: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_other_total_backlog_size"),
			"stats_ep_dcp_other_total_backlog_size",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpOtherTotalBytes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_other_total_bytes"),
			"Number of bytes per second being sent for other DCP connections for this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpReplicaBackoff: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_replica_backoff"),
			"Number of backoffs for replication DCP connections",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpReplicaCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_replicas"),
			"Number of internal replication DCP connections in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpReplicaItemsRemaining: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_replica_items_remaining"),
			"Number of items remaining to be sent to consumer in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpReplicaItemsSent: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_replica_items_sent"),
			"Number of items per second being sent for a producer for this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpReplicaProducerCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_replica_producers"),
			"Number of replication senders for this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpReplicaTotalBacklogSize: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_replica_total_backlog_size"),
			"stats_ep_dcp_replica_total_backlog_size",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpReplicaTotalBytes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_replica_total_bytes"),
			"Number of bytes per second being sent for replication DCP connections for this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpViewsBackoff: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_views_backoffs"),
			"Number of backoffs for views DCP connections",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpViewsCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_view_connections"),
			"Number of views DCP connections",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpViewsItemsRemaining: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_views_items_remaining"),
			"Number of views items remaining to be sent",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpViewsItemsSent: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_views_items_sent"),
			"Number of views items sent",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpViewsProducerCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_views_producers"),
			"Number of views producers",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpViewsTotalBacklogSize: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_views_total_backlog_size"),
			"stats_ep_dcp_views_total_backlog_size",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpViewsTotalBytes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_views_total_bytes"),
			"Number bytes per second being sent for views DCP connections",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpXdcrBackoff: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_xdcr_backoff"),
			"Number of backoffs for XDCR DCP connections",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpXdcrCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_xdcr_connections"),
			"Number of internal XDCR DCP connections in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpXdcrItemsRemaining: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_xdcr_items_remaining"),
			"Number of items remaining to be sent to consumer in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpXdcrItemsSent: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_xdcr_items_sent"),
			"Number of items per second being sent for a producer for this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpXdcrProducerCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_xdcr_producers"),
			"Number of XDCR senders for this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpXdcrTotalBacklogSize: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_xdcr_total_backlog_size"),
			"stats_ep_dcp_xdcr_total_backlog_size",
			[]string{"bucket"},
			nil,
		),
		statsEpDcpXdcrTotalBytes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_dcp_xdcr_total_bytes"),
			"Number of bytes per second being sent for XDCR DCP connections for this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpDiskqueueDrain: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_diskqueue_drain"),
			"Total number of items per second being written to disk in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpDiskqueueFill: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_diskqueue_fill"),
			"Total number of items per second being put on the disk queue in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpDiskqueueItems: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_diskqueue_items"),
			"Total number of items waiting to be written to disk in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpFlusherTodo: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_flusher_todo"),
			"Number of items currently being written",
			[]string{"bucket"},
			nil,
		),
		statsEpItemCommitFailed: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_item_commit_failed"),
			"Number of times a transaction failed to commit due to storage errors",
			[]string{"bucket"},
			nil,
		),
		statsEpKvSize: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_kv_size"),
			"Total amount of user data cached in RAM in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpMaxSize: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_max_size_bytes"),
			"The maximum amount of memory this bucket can use",
			[]string{"bucket"},
			nil,
		),
		statsEpMemHighWat: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_mem_high_wat_bytes"),
			"High water mark for auto-evictions",
			[]string{"bucket"},
			nil,
		),
		statsEpMemLowWat: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_mem_low_wat_bytes"),
			"Low water mark for auto-evictions",
			[]string{"bucket"},
			nil,
		),
		statsEpMetaDataMemory: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_meta_data_memory"),
			"Total amount of item metadata consuming RAM in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpNumNonResident: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_num_non_resident"),
			"Number of non-resident items",
			[]string{"bucket"},
			nil,
		),
		statsEpNumOpsDelMeta: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_num_ops_del_meta"),
			"Number of delete operations per second for this bucket as the target for XDCR",
			[]string{"bucket"},
			nil,
		),
		statsEpNumOpsDelRetMeta: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_num_ops_del_ret_meta"),
			"Number of delRetMeta operations per second for this bucket as the target for XDCR",
			[]string{"bucket"},
			nil,
		),
		statsEpNumOpsGetMeta: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_num_ops_get_meta"),
			"Number of metadata read operations per second for this bucket as the target for XDCR",
			[]string{"bucket"},
			nil,
		),
		statsEpNumOpsSetMeta: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_num_ops_set_meta"),
			"Number of set operations per second for this bucket as the target for XDCR",
			[]string{"bucket"},
			nil,
		),
		statsEpNumOpsSetRetMeta: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_num_ops_set_ret_meta"),
			"Number of setRetMeta operations per second for this bucket as the target for XDCR",
			[]string{"bucket"},
			nil,
		),
		statsEpNumValueEjects: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_num_value_ejects"),
			"Total number of items per second being ejected to disk in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpOomErrors: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_oom_errors"),
			"Number of times unrecoverable OOMs happened while processing operations",
			[]string{"bucket"},
			nil,
		),
		statsEpOpsCreate: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_ops_create"),
			"Total number of new items being inserted into this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpOpsUpdate: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_ops_update"),
			"Number of items updated on disk per second for this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpOverhead: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_overhead"),
			"Extra memory used by transient data like persistence queues or checkpoints",
			[]string{"bucket"},
			nil,
		),
		statsEpQueueSize: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_queue_size"),
			"Number of items queued for storage",
			[]string{"bucket"},
			nil,
		),
		statsEpResidentItemsRate: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_resident_items_rate"),
			"Percentage of all items cached in RAM in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpReplicaAheadExceptions: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_replica_ahead_exceptions"),
			"ep_replica_ahead_exceptions",
			[]string{"bucket"},
			nil,
		),
		statsEpReplicaHlcDrift: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_replica_hlc_drift"),
			"The sum of the total Absolute Drift, which is the accumulated drift observed by the vBucket. Drift is always accumulated as an absolute value.",
			[]string{"bucket"},
			nil,
		),
		statsEpTmpOomErrors: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_tmp_oom_errors"),
			"Number of back-offs sent per second to client SDKs due to OOM situations from this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEpVbTotal: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ep_vbuckets"),
			"Total number of vBuckets for this bucket",
			[]string{"bucket"},
			nil,
		),
		statsEvictions: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_evictions"),
			"Number of evictions",
			[]string{"bucket"},
			nil,
		),
		statsGetHits: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_get_hits"),
			"Number of get hits",
			[]string{"bucket"},
			nil,
		),
		statsGetMisses: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_get_misses"),
			"Number of get misses",
			[]string{"bucket"},
			nil,
		),
		statsHibernatedRequests: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_hibernated_requests"),
			"Number of streaming requests on port 8091 now idle",
			[]string{"bucket"},
			nil,
		),
		statsHibernatedWaked: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_hibernated_waked"),
			"Rate of streaming request wakeups on port 8091",
			[]string{"bucket"},
			nil,
		),
		statsHitRatio: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_hit_ratio"),
			"Hit ratio",
			[]string{"bucket"},
			nil,
		),
		statsIncrHits: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_incr_hits"),
			"Number of increment hits",
			[]string{"bucket"},
			nil,
		),
		statsIncrMisses: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_incr_misses"),
			"Number of increment misses",
			[]string{"bucket"},
			nil,
		),
		statsMemActualFree: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_mem_actual_free"),
			"Amount of RAM available on this server",
			[]string{"bucket"},
			nil,
		),
		statsMemActualUsed: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_mem_actual_used_bytes"),
			"stats_mem_actual_used",
			[]string{"bucket"},
			nil,
		),
		statsMemFree: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_mem_free_bytes"),
			"Amount of Memory free",
			[]string{"bucket"},
			nil,
		),
		statsMemTotal: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_mem_bytes"),
			"Total amount of memory available",
			[]string{"bucket"},
			nil,
		),
		statsMemUsed: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_mem_used_bytes"),
			"Amount of memory used",
			[]string{"bucket"},
			nil,
		),
		statsMemUsedSys: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_mem_used_sys_bytes"),
			"stats_mem_used_sys",
			[]string{"bucket"},
			nil,
		),
		statsMisses: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_misses"),
			"Number of misses",
			[]string{"bucket"},
			nil,
		),
		statsOps: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_ops"),
			"Total amount of operations per second to this bucket",
			[]string{"bucket"},
			nil,
		),
		statsRestRequests: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_rest_requests"),
			"Rate of http requests on port 8091",
			[]string{"bucket"},
			nil,
		),
		statsSwapTotal: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_swap_bytes"),
			"Total amount of swap available",
			[]string{"bucket"},
			nil,
		),
		statsSwapUsed: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_swap_used_bytes"),
			"Amount of swap space in use on this server",
			[]string{"bucket"},
			nil,
		),
		statsTimestamp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_timestamp"),
			"stats_timestamp",
			[]string{"bucket"},
			nil,
		),
		statsVbActiveEject: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_active_eject"),
			"Number of items per second being ejected to disk from active vBuckets in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsVbActiveItmMemory: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_active_itm_memory"),
			"Amount of active user data cached in RAM in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsVbActiveMetaDataMemory: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_active_meta_data_memory"),
			"Amount of active item metadata consuming RAM in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsVbActiveNum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_active_num"),
			"Number of vBuckets in the active state for this bucket",
			[]string{"bucket"},
			nil,
		),
		statsVbActiveNumNonResident: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_active_num_non_resident"),
			"Number of non resident vBuckets in the active state for this bucket",
			[]string{"bucket"},
			nil,
		),
		statsVbActiveOpsCreate: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_active_ops_create"),
			"New items per second being inserted into active vBuckets in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsVbActiveOpsUpdate: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_active_ops_update"),
			"Number of items updated on active vBucket per second for this bucket",
			[]string{"bucket"},
			nil,
		),
		statsVbActiveQueueAge: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_active_queue_age"),
			"Sum of disk queue item age in milliseconds",
			[]string{"bucket"},
			nil,
		),
		statsVbActiveQueueDrain: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_active_queue_drain"),
			"Number of active items per second being written to disk in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsVbActiveQueueFill: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_active_queue_fill"),
			"Number of active items per second being put on the active item disk queue in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsVbActiveQueueSize: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_active_queue_size"),
			"Number of active items waiting to be written to disk in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsVbActiveResidentItemsRatio: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_active_resident_items_ratio"),
			"Percentage of active items cached in RAM in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsVbAvgActiveQueueAge: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_avg_active_queue_age"),
			"Average age in seconds of active items in the active item queue for this bucket",
			[]string{"bucket"},
			nil,
		),
		statsVbAvgPendingQueueAge: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_avg_pending_queue_age"),
			"Average age in seconds of pending items in the pending item queue for this bucket and should be transient during rebalancing",
			[]string{"bucket"},
			nil,
		),
		statsVbAvgReplicaQueueAge: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_avg_replica_queue_age"),
			"Average age in seconds of replica items in the replica item queue for this bucket",
			[]string{"bucket"},
			nil,
		),
		statsVbAvgTotalQueueAge: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_avg_total_queue_age"),
			"Average age in seconds of all items in the disk write queue for this bucket",
			[]string{"bucket"},
			nil,
		),
		statsVbPendingCurrItems: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_pending_curr_items"),
			"Number of items in pending vBuckets in this bucket and should be transient during rebalancing",
			[]string{"bucket"},
			nil,
		),
		statsVbPendingEject: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_pending_eject"),
			"Number of items per second being ejected to disk from pending vBuckets in this bucket and should be transient during rebalancing",
			[]string{"bucket"},
			nil,
		),
		statsVbPendingItmMemory: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_pending_itm_memory"),
			"Amount of pending user data cached in RAM in this bucket and should be transient during rebalancing",
			[]string{"bucket"},
			nil,
		),
		statsVbPendingMetaDataMemory: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_pending_meta_data_memory"),
			"Amount of pending item metadata consuming RAM in this bucket and should be transient during rebalancing",
			[]string{"bucket"},
			nil,
		),
		statsVbPendingNum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_pending_num"),
			"Number of vBuckets in the pending state for this bucket and should be transient during rebalancing",
			[]string{"bucket"},
			nil,
		),
		statsVbPendingNumNonResident: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_pending_num_non_resident"),
			"Number of non resident vBuckets in the pending state for this bucket",
			[]string{"bucket"},
			nil,
		),
		statsVbPendingOpsCreate: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_pending_ops_create"),
			"New items per second being instead into pending vBuckets in this bucket and should be transient during rebalancing",
			[]string{"bucket"},
			nil,
		),
		statsVbPendingOpsUpdate: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_pending_ops_update"),
			"Number of items updated on pending vBucket per second for this bucket",
			[]string{"bucket"},
			nil,
		),
		statsVbPendingQueueAge: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_pending_queue_age"),
			"Sum of disk pending queue item age in milliseconds",
			[]string{"bucket"},
			nil,
		),
		statsVbPendingQueueDrain: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_pending_queue_drain"),
			"Number of pending items per second being written to disk in this bucket and should be transient during rebalancing",
			[]string{"bucket"},
			nil,
		),
		statsVbPendingQueueFill: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_pending_queue_fill"),
			"Number of pending items per second being put on the pending item disk queue in this bucket and should be transient during rebalancing",
			[]string{"bucket"},
			nil,
		),
		statsVbPendingQueueSize: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_pending_queue_size"),
			"Number of pending items waiting to be written to disk in this bucket and should be transient during rebalancing",
			[]string{"bucket"},
			nil,
		),
		statsVbPendingResidentItemsRatio: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_pending_resident_items_ratio"),
			"Percentage of items in pending state vbuckets cached in RAM in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsVbReplicaCurrItems: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_replica_curr_items"),
			"Number of items in replica vBuckets in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsVbReplicaEject: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_replica_eject"),
			"Number of items per second being ejected to disk from replica vBuckets in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsVbReplicaItmMemory: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_replica_itm_memory"),
			"Amount of replica user data cached in RAM in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsVbReplicaMetaDataMemory: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_replica_meta_data_memory"),
			"Amount of replica item metadata consuming in RAM in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsVbReplicaNum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_replica_num"),
			"Number of vBuckets in the replica state for this bucket",
			[]string{"bucket"},
			nil,
		),
		statsVbReplicaNumNonResident: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_replica_num_non_resident"),
			"stats_vb_replica_num_non_resident",
			[]string{"bucket"},
			nil,
		),
		statsVbReplicaOpsCreate: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_replica_ops_create"),
			"New items per second being inserted into replica vBuckets in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsVbReplicaOpsUpdate: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_replica_ops_update"),
			"Number of items updated on replica vBucket per second for this bucket",
			[]string{"bucket"},
			nil,
		),
		statsVbReplicaQueueAge: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_replica_queue_age"),
			"Sum of disk replica queue item age in milliseconds",
			[]string{"bucket"},
			nil,
		),
		statsVbReplicaQueueDrain: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_replica_queue_drain"),
			"Number of replica items per second being written to disk in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsVbReplicaQueueFill: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_replica_queue_fill"),
			"Number of replica items per second being put on the replica item disk queue in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsVbReplicaQueueSize: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_replica_queue_size"),
			"Number of replica items waiting to be written to disk in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsVbReplicaResidentItemsRatio: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_replica_resident_items_ratio"),
			"Percentage of replica items cached in RAM in this bucket",
			[]string{"bucket"},
			nil,
		),
		statsVbTotalQueueAge: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_vbuckets_total_queue_age"),
			"Sum of disk queue item age in milliseconds",
			[]string{"bucket"},
			nil,
		),
		statsXdcOps: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_xdc_ops"),
			"Total XDCR operations per second for this bucket",
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
	ch <- c.statsAvgActiveTimestampDrift
	ch <- c.statsAvgReplicaTimestampDrift
	ch <- c.statsAvgDiskCommitTime
	ch <- c.statsAvgDiskUpdateTime
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
	ch <- c.statsCouchViewsDataSize
	ch <- c.statsCouchViewsActualDiskSize
	ch <- c.statsCouchViewsFragmentation
	ch <- c.statsCouchViewsOps
	ch <- c.statsCouchDocsFragmentation
	ch <- c.statsCouchDocsDiskSize
	ch <- c.statsCPUIdleMs
	ch <- c.statsCPULocalMs
	ch <- c.statsCPUUtilizationRate
	ch <- c.statsCurrConnections
	ch <- c.statsCurrItems
	ch <- c.statsCurrItemsTot
	ch <- c.statsDecrHits
	ch <- c.statsDecrMisses
	ch <- c.statsDeleteHits
	ch <- c.statsDeleteMisses
	ch <- c.statsDiskCommitCount
	ch <- c.statsDiskUpdateCount
	ch <- c.statsDiskWriteQueue
	ch <- c.statsEpActiveAheadExceptions
	ch <- c.statsEpActiveHlcDrift
	ch <- c.statsEpClockCasDriftThresholdExceeded
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
	ch <- c.statsEpReplicaAheadExceptions
	ch <- c.statsEpReplicaHlcDrift
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

	// nolint: lll
	for _, bucket := range buckets {
		log.Debugf("Collecting %s bucket metrics...", bucket.Name)
		stats, err := c.client.BucketStats(bucket.Name)
		if err != nil {
			ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 0)
			log.With("error", err).Error("failed to scrape buckets stats")
			return
		}

		// TODO: collect bucket.Nodes metrics as well

		ch <- prometheus.MustNewConstMetric(c.basicstatsDataused, prometheus.GaugeValue, bucket.BasicStats.DataUsed, bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.basicstatsDiskfetches, prometheus.GaugeValue, bucket.BasicStats.DiskFetches, bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.basicstatsDiskused, prometheus.GaugeValue, bucket.BasicStats.DiskUsed, bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.basicstatsItemcount, prometheus.GaugeValue, bucket.BasicStats.ItemCount, bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.basicstatsMemused, prometheus.GaugeValue, bucket.BasicStats.MemUsed, bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.basicstatsOpspersec, prometheus.GaugeValue, bucket.BasicStats.OpsPerSec, bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.basicstatsQuotapercentused, prometheus.GaugeValue, bucket.BasicStats.QuotaPercentUsed, bucket.Name)

		ch <- prometheus.MustNewConstMetric(c.statsAvgBgWaitTime, prometheus.GaugeValue, last(stats.Op.Samples.AvgBgWaitTime)/1000000, bucket.Name) // this comes as microseconds from CB
		ch <- prometheus.MustNewConstMetric(c.statsAvgActiveTimestampDrift, prometheus.GaugeValue, last(stats.Op.Samples.AvgActiveTimestampDrift), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsAvgReplicaTimestampDrift, prometheus.GaugeValue, last(stats.Op.Samples.AvgReplicaTimestampDrift), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsAvgDiskCommitTime, prometheus.GaugeValue, last(stats.Op.Samples.AvgDiskCommitTime), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsAvgDiskUpdateTime, prometheus.GaugeValue, last(stats.Op.Samples.AvgDiskUpdateTime), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsBytesRead, prometheus.GaugeValue, last(stats.Op.Samples.BytesRead), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsBytesWritten, prometheus.GaugeValue, last(stats.Op.Samples.BytesWritten), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCasBadval, prometheus.GaugeValue, last(stats.Op.Samples.CasBadval), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCasHits, prometheus.GaugeValue, last(stats.Op.Samples.CasHits), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCasMisses, prometheus.GaugeValue, last(stats.Op.Samples.CasMisses), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCmdGet, prometheus.GaugeValue, last(stats.Op.Samples.CmdGet), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCmdSet, prometheus.GaugeValue, last(stats.Op.Samples.CmdSet), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCouchTotalDiskSize, prometheus.GaugeValue, last(stats.Op.Samples.CouchTotalDiskSize), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCouchViewsDataSize, prometheus.GaugeValue, last(stats.Op.Samples.CouchViewsDataSize), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCouchViewsActualDiskSize, prometheus.GaugeValue, last(stats.Op.Samples.CouchViewsActualDiskSize), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCouchViewsFragmentation, prometheus.GaugeValue, last(stats.Op.Samples.CouchViewsFragmentation), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCouchViewsOps, prometheus.GaugeValue, last(stats.Op.Samples.CouchViewsOps), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCouchDocsActualDiskSize, prometheus.GaugeValue, last(stats.Op.Samples.CouchDocsActualDiskSize), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCouchDocsFragmentation, prometheus.GaugeValue, last(stats.Op.Samples.CouchDocsFragmentation), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCouchDocsDataSize, prometheus.GaugeValue, last(stats.Op.Samples.CouchDocsDataSize), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCouchDocsDiskSize, prometheus.GaugeValue, last(stats.Op.Samples.CouchDocsDiskSize), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCPUIdleMs, prometheus.GaugeValue, last(stats.Op.Samples.CPUIdleMs), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCPULocalMs, prometheus.GaugeValue, last(stats.Op.Samples.CPULocalMs), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCPUUtilizationRate, prometheus.GaugeValue, last(stats.Op.Samples.CPUUtilizationRate), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCurrConnections, prometheus.GaugeValue, last(stats.Op.Samples.CurrConnections), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCurrItems, prometheus.GaugeValue, last(stats.Op.Samples.CurrItems), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsCurrItemsTot, prometheus.GaugeValue, last(stats.Op.Samples.CurrItemsTot), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsDecrHits, prometheus.GaugeValue, last(stats.Op.Samples.DecrHits), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsDecrMisses, prometheus.GaugeValue, last(stats.Op.Samples.DecrMisses), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsDeleteHits, prometheus.GaugeValue, last(stats.Op.Samples.DeleteHits), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsDeleteMisses, prometheus.GaugeValue, last(stats.Op.Samples.DeleteMisses), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsDiskCommitCount, prometheus.GaugeValue, last(stats.Op.Samples.DiskCommitCount), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsDiskUpdateCount, prometheus.GaugeValue, last(stats.Op.Samples.DiskUpdateCount), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsDiskWriteQueue, prometheus.GaugeValue, last(stats.Op.Samples.DiskWriteQueue), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpActiveAheadExceptions, prometheus.GaugeValue, last(stats.Op.Samples.EpActiveAheadExceptions), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpActiveHlcDrift, prometheus.GaugeValue, last(stats.Op.Samples.EpActiveHlcDrift), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpClockCasDriftThresholdExceeded, prometheus.GaugeValue, last(stats.Op.Samples.EpClockCasDriftThresholdExceeded), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpBgFetched, prometheus.GaugeValue, last(stats.Op.Samples.EpBgFetched), bucket.Name)
		// for some reason, this ratio can be > 100, so we added a `min` function
		ch <- prometheus.MustNewConstMetric(c.statsEpCacheMissRate, prometheus.GaugeValue, min(last(stats.Op.Samples.EpCacheMissRate), 100), bucket.Name)
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
		ch <- prometheus.MustNewConstMetric(c.statsEpReplicaAheadExceptions, prometheus.GaugeValue, last(stats.Op.Samples.EpReplicaAheadExceptions), bucket.Name)
		ch <- prometheus.MustNewConstMetric(c.statsEpReplicaHlcDrift, prometheus.GaugeValue, last(stats.Op.Samples.EpReplicaHlcDrift), bucket.Name)
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
