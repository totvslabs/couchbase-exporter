package collector

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"github.com/totvslabs/couchbase-exporter/client"
)

type ftsCollector struct {
	mutex  sync.Mutex
	client client.Client

	up             *prometheus.Desc
	scrapeDuration *prometheus.Desc

	batchBytesAdded               *prometheus.Desc
	batchBytesRemoved             *prometheus.Desc
	currBatchesBlockedByHerder    *prometheus.Desc
	numBytesUsedRam               *prometheus.Desc
	pctCpuGc                      *prometheus.Desc
	totBatchesFlushedOnMaxops     *prometheus.Desc
	totatchesFlushedOnTimer       *prometheus.Desc
	totBleveDestClosed            *prometheus.Desc
	totBleveDestOpened            *prometheus.Desc
	totGrpcListenersClosed        *prometheus.Desc
	totGrpcListenersOpened        *prometheus.Desc
	totGrpcQueryrejectOnMemquota  *prometheus.Desc
	totGrpcsListenersClosed       *prometheus.Desc
	totGrpcsListenersOpened       *prometheus.Desc
	totHttpLimitlistenersClosed   *prometheus.Desc
	totHttpLimitlistenersOpened   *prometheus.Desc
	totHttpsLimitlistenersClosed  *prometheus.Desc
	totHttpsLimitlistenersOpened  *prometheus.Desc
	totQueryrejectOnMemquota      *prometheus.Desc
	totRemoteGrpc                 *prometheus.Desc
	totRemoteGrpcTls              *prometheus.Desc
	totRemoteHttp                 *prometheus.Desc
	totRemoteHttp2                *prometheus.Desc
	totalGc                       *prometheus.Desc
	totalQueriesRejectedByHerder  *prometheus.Desc
	avgGrpcInternalQueriesLatency *prometheus.Desc
	avgGrpcQueriesLatency         *prometheus.Desc
	avgInternalQueriesLatency     *prometheus.Desc
	avgQueriesLatency             *prometheus.Desc
	batchMergeCount               *prometheus.Desc
	docCount                      *prometheus.Desc
	iteratorNextCount             *prometheus.Desc
	iteratorSeekCount             *prometheus.Desc
	numBytesLiveData              *prometheus.Desc
	numBytesUsedDisk              *prometheus.Desc
	numBytesUsedDiskByRoot        *prometheus.Desc
	numFilesOnDisk                *prometheus.Desc
	numMutationToIndex            *prometheus.Desc
	numPersisterNapMergerBreak    *prometheus.Desc
	numPersisterNapPauseComleted  *prometheus.Desc
	numPindexesActual             *prometheus.Desc
	numPindexesTarget             *prometheus.Desc
	numRecsToPersist              *prometheus.Desc
	numRootFilesegments           *prometheus.Desc
	numRootMemorysegments         *prometheus.Desc
	readerGetCount                *prometheus.Desc
	readerMultiGetCount           *prometheus.Desc
	readerPrefixIteratorCount     *prometheus.Desc
	readerRangeIteratorCount      *prometheus.Desc
	timerBatchStoreCount          *prometheus.Desc
	timerDataDeleteCount          *prometheus.Desc
	timerDataUpdateCount          *prometheus.Desc
	timerOpaqueGetCount           *prometheus.Desc
	timerOpaqueSetCount           *prometheus.Desc
	timerRollbackCount            *prometheus.Desc
	timerSnapshotStarCount        *prometheus.Desc
	totalBytesIndexed             *prometheus.Desc
	totalBytesQueryResults        *prometheus.Desc
	totalCompactionWrittenBytes   *prometheus.Desc
	totalCompactions              *prometheus.Desc
	totalGrpcInternalQueries      *prometheus.Desc
	totalGrpcQueries              *prometheus.Desc
	totalGrpcQueriesError         *prometheus.Desc
	totalGrpcQueriesSlow          *prometheus.Desc
	totalGrpcQueriesTimeout       *prometheus.Desc
	totalGrpcQueriesTime          *prometheus.Desc
	totalInternalQueries          *prometheus.Desc
	totalQueries                  *prometheus.Desc
	totalQueriesError             *prometheus.Desc
	totalQueriesSlow              *prometheus.Desc
	totalQueriesTimeout           *prometheus.Desc
	totalRequestTime              *prometheus.Desc
	totalTermSearchers            *prometheus.Desc
	totalTermSearchersFinished    *prometheus.Desc
	writerEcecuteBatchCount       *prometheus.Desc
}

// NewBucketsCollector buckets collector
func NewFTSCollector(client client.Client) prometheus.Collector {
	const subsystem = "fts"
	// nolint: lll
	return &ftsCollector{
		client: client,
		up: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "up"),
			"Couchbase cluster API is responding",
			nil,
			nil,
		),
		scrapeDuration: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "scrape_duration_seconds"),
			"Scrape duration in seconds",
			nil,
			nil,
		),
		batchBytesAdded: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "batch_bytes_added"),
			"FTS batch_bytes_added",
			nil,
			nil,
		),
		batchBytesRemoved: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "batch_bytes_removed"),
			"FTS batch_bytes_removed",
			nil,
			nil,
		),
		currBatchesBlockedByHerder: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "curr_batches_blocked_by_herder"),
			"FTS curr_batches_blocked_by_herder",
			nil,
			nil,
		),
		numBytesUsedRam: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_bytes_used_ram"),
			"FTS num_bytes_used_ram",
			nil,
			nil,
		),
		pctCpuGc: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "pct_cpu_gc"),
			"FTS pct_cpu_gc",
			nil,
			nil,
		),
		totBatchesFlushedOnMaxops: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "tot_batches_flushed_on_maxops"),
			"tot_batches_flushed_on_maxops",
			nil,
			nil,
		),
		totatchesFlushedOnTimer: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "tot_batches_flushed_on_timer"),
			"tot_batches_flushed_on_timer",
			nil,
			nil,
		),
		totBleveDestClosed: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "tot_bleve_dest_closed"),
			"tot_bleve_dest_opened",
			nil,
			nil,
		),
		totBleveDestOpened: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "tot_bleve_dest_opened"),
			"tot_bleve_dest_opened",
			nil,
			nil,
		),
		totGrpcListenersClosed: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "tot_grpc_listeners_closed"),
			"tot_grpc_listeners_closed",
			nil,
			nil,
		),
		totGrpcListenersOpened: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "tot_grpc_listeners_opened"),
			"tot_grpc_listeners_opened",
			nil,
			nil,
		),
		totGrpcQueryrejectOnMemquota: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "tot_grpc_queryreject_on_memquota"),
			"tot_grpc_queryreject_on_memquota",
			nil,
			nil,
		),
		totGrpcsListenersClosed: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "tot_grpcs_listeners_closed"),
			"tot_grpcs_listeners_closed",
			nil,
			nil,
		),
		totGrpcsListenersOpened: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "tot_grpcs_listeners_opened"),
			"tot_grpcs_listeners_opened",
			nil,
			nil,
		),
		totHttpLimitlistenersClosed: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "tot_http_limitlisteners_closed"),
			"tot_http_limitlisteners_closed",
			nil,
			nil,
		),
		totHttpLimitlistenersOpened: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "tot_http_limitlisteners_opened"),
			"tot_http_limitlisteners_opened",
			nil,
			nil,
		),
		totHttpsLimitlistenersClosed: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "tot_https_limitlisteners_closed"),
			"tot_https_limitlisteners_closed",
			nil,
			nil,
		),
		totHttpsLimitlistenersOpened: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "tot_https_limitlisteners_opened"),
			"tot_https_limitlisteners_opened",
			nil,
			nil,
		),
		totQueryrejectOnMemquota: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "tot_queryreject_on_memquota"),
			"tot_queryreject_on_memquota",
			nil,
			nil,
		),
		totRemoteGrpc: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "tot_remote_grpc"),
			"tot_remote_grpc",
			nil,
			nil,
		),
		totRemoteGrpcTls: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "tot_remote_grpc_tls"),
			"tot_remote_grpc_tls",
			nil,
			nil,
		),
		totRemoteHttp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "tot_remote_http"),
			"tot_remote_http",
			nil,
			nil,
		),
		totRemoteHttp2: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "tot_remote_http2"),
			"tot_remote_http2",
			nil,
			nil,
		),
		totalGc: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "total_gc"),
			"total_gc",
			nil,
			nil,
		),
		totalQueriesRejectedByHerder: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "total_queries_rejected_by_herder"),
			"total_queries_rejected_by_herder",
			nil,
			nil,
		),
		avgGrpcInternalQueriesLatency: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "avg_grpc_internal_queries_latency"),
			"avg_grpc_internal_queries_latency",
			[]string{"fts"},
			nil,
		),
		avgGrpcQueriesLatency: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "avg_grpc_queries_latency"),
			"avg_grpc_queries_latency",
			[]string{"fts"},
			nil,
		),
		avgInternalQueriesLatency: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "avg_internal_queries_latency"),
			"avg_internal_queries_latency",
			[]string{"fts"},
			nil,
		),
		avgQueriesLatency: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "avg_queries_latency"),
			"avg_queries_latency",
			[]string{"fts"},
			nil,
		),
		batchMergeCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "batch_merge_count"),
			"batch_merge_count",
			[]string{"fts"},
			nil,
		),
		docCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "doc_count"),
			"doc_count",
			[]string{"fts"},
			nil,
		),
		iteratorNextCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "iterator_next_count"),
			"iterator_next_count",
			[]string{"fts"},
			nil,
		),
		iteratorSeekCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "iterator_seek_count"),
			"iterator_seek_count",
			[]string{"fts"},
			nil,
		),
		numBytesLiveData: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_bytes_live_data"),
			"num_bytes_live_data",
			[]string{"fts"},
			nil,
		),
		numBytesUsedDisk: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_bytes_used_disk"),
			"num_bytes_used_disk",
			[]string{"fts"},
			nil,
		),
		numBytesUsedDiskByRoot: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_bytes_used_disk_by_root"),
			"num_bytes_used_disk_by_root",
			[]string{"fts"},
			nil,
		),
		numFilesOnDisk: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_files_on_disk"),
			"num_files_on_disk",
			[]string{"fts"},
			nil,
		),
		numMutationToIndex: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_mutations_to_index"),
			"num_mutations_to_index",
			[]string{"fts"},
			nil,
		),
		numPersisterNapMergerBreak: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_persister_nap_merger_break"),
			"num_persister_nap_merger_break",
			[]string{"fts"},
			nil,
		),
		numPersisterNapPauseComleted: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_persister_nap_pause_completed"),
			"num_persister_nap_pause_completed",
			[]string{"fts"},
			nil,
		),
		numPindexesActual: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_pindexes_actual"),
			"num_pindexes_actual",
			[]string{"fts"},
			nil,
		),
		numPindexesTarget: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_pindexes_target"),
			"num_pindexes_target",
			[]string{"fts"},
			nil,
		),
		numRecsToPersist: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_recs_to_persist"),
			"num_recs_to_persist",
			[]string{"fts"},
			nil,
		),
		numRootFilesegments: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_root_filesegments"),
			"num_root_filesegments",
			[]string{"fts"},
			nil,
		),
		numRootMemorysegments: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "num_root_memorysegments"),
			"num_root_memorysegments",
			[]string{"fts"},
			nil,
		),
		readerGetCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "reader_get_count"),
			"reader_get_count",
			[]string{"fts"},
			nil,
		),
		readerMultiGetCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "reader_multi_get_count"),
			"reader_multi_get_count",
			[]string{"fts"},
			nil,
		),
		readerPrefixIteratorCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "reader_prefix_iterator_count"),
			"reader_prefix_iterator_count",
			[]string{"fts"},
			nil,
		),
		readerRangeIteratorCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "reader_range_iterator_count"),
			"reader_range_iterator_count",
			[]string{"fts"},
			nil,
		),
		timerBatchStoreCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "timer_batch_store_count"),
			"timer_batch_store_count",
			[]string{"fts"},
			nil,
		),
		timerDataDeleteCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "timer_data_delete_count"),
			"timer_data_delete_count",
			[]string{"fts"},
			nil,
		),
		timerDataUpdateCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "timer_data_update_count"),
			"timer_data_update_count",
			[]string{"fts"},
			nil,
		),
		timerOpaqueGetCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "timer_opaque_get_count"),
			"timer_opaque_get_count",
			[]string{"fts"},
			nil,
		),
		timerOpaqueSetCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "timer_opaque_set_count"),
			"timer_opaque_set_count",
			[]string{"fts"},
			nil,
		),
		timerRollbackCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "timer_rollback_count"),
			"timer_rollback_count",
			[]string{"fts"},
			nil,
		),
		timerSnapshotStarCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "timer_snapshot_start_count"),
			"timer_snapshot_start_count",
			[]string{"fts"},
			nil,
		),
		totalBytesIndexed: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "total_bytes_indexed"),
			"total_bytes_indexed",
			[]string{"fts"},
			nil,
		),
		totalBytesQueryResults: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "total_bytes_query_results"),
			"total_bytes_query_results",
			[]string{"fts"},
			nil,
		),
		totalCompactionWrittenBytes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "total_compaction_written_bytes"),
			"total_compaction_written_bytes",
			[]string{"fts"},
			nil,
		),
		totalCompactions: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "total_compactions"),
			"total_compactions",
			[]string{"fts"},
			nil,
		),
		totalGrpcInternalQueries: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "total_grpc_internal_queries"),
			"total_grpc_internal_queries",
			[]string{"fts"},
			nil,
		),
		totalGrpcQueries: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "total_grpc_queries"),
			"total_grpc_queries",
			[]string{"fts"},
			nil,
		),
		totalGrpcQueriesError: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "total_grpc_queries_error"),
			"total_grpc_queries_error",
			[]string{"fts"},
			nil,
		),
		totalGrpcQueriesSlow: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "total_grpc_queries_slow"),
			"total_grpc_queries_slow",
			[]string{"fts"},
			nil,
		),
		totalGrpcQueriesTimeout: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "total_grpc_queries_timeout"),
			"total_grpc_queries_timeout",
			[]string{"fts"},
			nil,
		),
		totalGrpcQueriesTime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "total_grpc_request_time"),
			"total_grpc_request_time",
			[]string{"fts"},
			nil,
		),
		totalInternalQueries: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "total_internal_queries"),
			"total_internal_queries",
			[]string{"fts"},
			nil,
		),
		totalQueries: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "total_queries"),
			"total_queries",
			[]string{"fts"},
			nil,
		),
		totalQueriesError: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "total_queries_error"),
			"total_queries_error",
			[]string{"fts"},
			nil,
		),
		totalQueriesSlow: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "total_queries_slow"),
			"total_queries_slow",
			[]string{"fts"},
			nil,
		),
		totalQueriesTimeout: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "total_queries_timeout"),
			"total_queries_timeout",
			[]string{"fts"},
			nil,
		),
		totalRequestTime: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "total_request_time"),
			"total_request_time",
			[]string{"fts"},
			nil,
		),
		totalTermSearchers: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "total_term_searchers"),
			"total_term_searchers",
			[]string{"fts"},
			nil,
		),
		totalTermSearchersFinished: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "total_term_searchers_finished"),
			"total_term_searchers_finished",
			[]string{"fts"},
			nil,
		),
		writerEcecuteBatchCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "writer_execute_batch_count"),
			"writer_execute_batch_count",
			[]string{"fts"},
			nil,
		),
	}
}

func (c *ftsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.up
	ch <- c.scrapeDuration

	ch <- c.batchBytesAdded
	ch <- c.batchBytesRemoved
	ch <- c.currBatchesBlockedByHerder
	ch <- c.numBytesUsedRam
	ch <- c.pctCpuGc
	ch <- c.totBatchesFlushedOnMaxops
	ch <- c.totatchesFlushedOnTimer
	ch <- c.totBleveDestClosed
	ch <- c.totBleveDestOpened
	ch <- c.totGrpcListenersClosed
	ch <- c.totGrpcListenersOpened
	ch <- c.totGrpcQueryrejectOnMemquota
	ch <- c.totGrpcsListenersClosed
	ch <- c.totGrpcsListenersOpened
	ch <- c.totHttpLimitlistenersClosed
	ch <- c.totHttpLimitlistenersOpened
	ch <- c.totHttpsLimitlistenersClosed
	ch <- c.totHttpsLimitlistenersOpened
	ch <- c.totQueryrejectOnMemquota
	ch <- c.totRemoteGrpc
	ch <- c.totRemoteGrpcTls
	ch <- c.totRemoteHttp
	ch <- c.totRemoteHttp2
	ch <- c.totalGc
	ch <- c.totalQueriesRejectedByHerder
	ch <- c.avgGrpcInternalQueriesLatency
	ch <- c.avgGrpcQueriesLatency
	ch <- c.avgInternalQueriesLatency
	ch <- c.avgQueriesLatency
	ch <- c.batchMergeCount
	ch <- c.docCount
	ch <- c.iteratorNextCount
	ch <- c.iteratorSeekCount
	ch <- c.numBytesLiveData
	ch <- c.numBytesUsedDisk
	ch <- c.numBytesUsedDiskByRoot
	ch <- c.numFilesOnDisk
	ch <- c.numMutationToIndex
	ch <- c.numPersisterNapMergerBreak
	ch <- c.numPersisterNapPauseComleted
	ch <- c.numPindexesActual
	ch <- c.numPindexesTarget
	ch <- c.numRecsToPersist
	ch <- c.numRootFilesegments
	ch <- c.numRootMemorysegments
	ch <- c.readerGetCount
	ch <- c.readerMultiGetCount
	ch <- c.readerPrefixIteratorCount
	ch <- c.readerRangeIteratorCount
	ch <- c.timerBatchStoreCount
	ch <- c.timerDataDeleteCount
	ch <- c.timerDataUpdateCount
	ch <- c.timerOpaqueGetCount
	ch <- c.timerOpaqueSetCount
	ch <- c.timerRollbackCount
	ch <- c.timerSnapshotStarCount
	ch <- c.totalBytesIndexed
	ch <- c.totalBytesQueryResults
	ch <- c.totalCompactionWrittenBytes
	ch <- c.totalCompactions
	ch <- c.totalGrpcInternalQueries
	ch <- c.totalGrpcQueries
	ch <- c.totalGrpcQueriesError
	ch <- c.totalGrpcQueriesSlow
	ch <- c.totalGrpcQueriesTimeout
	ch <- c.totalGrpcQueriesTime
	ch <- c.totalInternalQueries
	ch <- c.totalQueries
	ch <- c.totalQueriesError
	ch <- c.totalQueriesSlow
	ch <- c.totalQueriesTimeout
	ch <- c.totalRequestTime
	ch <- c.totalTermSearchers
	ch <- c.totalTermSearchersFinished
	ch <- c.writerEcecuteBatchCount
}

func (c *ftsCollector) Collect(ch chan<- prometheus.Metric) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	start := time.Now()
	log.Info("Collecting fts metrics...")

	fts, err := c.client.FTS()
	if err != nil {
		ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 0)
		log.With("error", err).Error("failed to scrape cluster")
		return
	}

	file, err := os.Open("fts.txt")
	var txtlines []string
	if err != nil {
		log.Warnf("failed opening file: %s", err)
	} else {
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			txtlines = append(txtlines, scanner.Text())
		}
	}

	defer file.Close()
	for _, index := range txtlines {
		log.Debugf("Collecting %s fts metrics...", index)
		ch <- prometheus.MustNewConstMetric(c.avgGrpcInternalQueriesLatency, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":avg_grpc_internal_queries_latency")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.avgGrpcQueriesLatency, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":avg_grpc_queries_latency")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.avgInternalQueriesLatency, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":avg_internal_queries_latency")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.avgQueriesLatency, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":avg_queries_latency")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.batchMergeCount, prometheus.CounterValue, fts.(map[string]interface{})[fmt.Sprint(index, ":batch_merge_count")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.docCount, prometheus.CounterValue, fts.(map[string]interface{})[fmt.Sprint(index, ":doc_count")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.iteratorNextCount, prometheus.CounterValue, fts.(map[string]interface{})[fmt.Sprint(index, ":iterator_next_count")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.iteratorSeekCount, prometheus.CounterValue, fts.(map[string]interface{})[fmt.Sprint(index, ":iterator_seek_count")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.numBytesLiveData, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":num_bytes_live_data")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.numBytesUsedDisk, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":num_bytes_used_disk")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.numBytesUsedDiskByRoot, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":num_bytes_used_disk_by_root")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.numFilesOnDisk, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":num_files_on_disk")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.numMutationToIndex, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":num_mutations_to_index")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.numPersisterNapMergerBreak, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":num_persister_nap_merger_break")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.numPersisterNapPauseComleted, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":num_persister_nap_pause_completed")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.numPindexesActual, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":num_pindexes_actual")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.numPindexesTarget, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":num_pindexes_target")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.numRecsToPersist, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":num_recs_to_persist")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.numRootFilesegments, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":num_root_filesegments")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.numRootMemorysegments, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":num_root_memorysegments")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.readerGetCount, prometheus.CounterValue, fts.(map[string]interface{})[fmt.Sprint(index, ":reader_get_count")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.readerPrefixIteratorCount, prometheus.CounterValue, fts.(map[string]interface{})[fmt.Sprint(index, ":reader_multi_get_count")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.readerRangeIteratorCount, prometheus.CounterValue, fts.(map[string]interface{})[fmt.Sprint(index, ":reader_prefix_iterator_count")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.timerBatchStoreCount, prometheus.CounterValue, fts.(map[string]interface{})[fmt.Sprint(index, ":timer_batch_store_count")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.timerDataDeleteCount, prometheus.CounterValue, fts.(map[string]interface{})[fmt.Sprint(index, ":timer_data_delete_count")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.timerDataUpdateCount, prometheus.CounterValue, fts.(map[string]interface{})[fmt.Sprint(index, ":timer_data_update_count")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.timerOpaqueGetCount, prometheus.CounterValue, fts.(map[string]interface{})[fmt.Sprint(index, ":timer_opaque_get_count")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.timerOpaqueSetCount, prometheus.CounterValue, fts.(map[string]interface{})[fmt.Sprint(index, ":timer_opaque_set_count")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.timerRollbackCount, prometheus.CounterValue, fts.(map[string]interface{})[fmt.Sprint(index, ":timer_rollback_count")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.timerSnapshotStarCount, prometheus.CounterValue, fts.(map[string]interface{})[fmt.Sprint(index, ":timer_snapshot_start_count")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.totalBytesIndexed, prometheus.CounterValue, fts.(map[string]interface{})[fmt.Sprint(index, ":total_bytes_indexed")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.totalBytesQueryResults, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":total_bytes_query_results")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.totalCompactionWrittenBytes, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":total_compaction_written_bytes")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.totalCompactions, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":total_compactions")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.totalGrpcInternalQueries, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":total_grpc_internal_queries")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.totalGrpcQueries, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":total_grpc_queries")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.totalGrpcQueriesError, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":total_grpc_queries_error")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.totalGrpcQueriesSlow, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":total_grpc_queries_slow")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.totalGrpcQueriesTimeout, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":total_grpc_queries_timeout")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.totalGrpcQueriesTime, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":total_grpc_request_time")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.totalInternalQueries, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":total_internal_queries")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.totalQueries, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":total_queries")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.totalQueriesError, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":total_queries_error")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.totalQueriesSlow, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":total_queries_slow")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.totalQueriesTimeout, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":total_queries_timeout")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.totalRequestTime, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":total_request_time")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.totalTermSearchers, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":total_term_searchers")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.totalTermSearchersFinished, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":total_term_searchers_finished")].(float64), index)
		ch <- prometheus.MustNewConstMetric(c.writerEcecuteBatchCount, prometheus.GaugeValue, fts.(map[string]interface{})[fmt.Sprint(index, ":writer_execute_batch_count")].(float64), index)
	}

	ch <- prometheus.MustNewConstMetric(c.batchBytesAdded, prometheus.GaugeValue, fts.(map[string]interface{})["batch_bytes_added"].(float64))
	ch <- prometheus.MustNewConstMetric(c.batchBytesRemoved, prometheus.GaugeValue, fts.(map[string]interface{})["batch_bytes_removed"].(float64))
	ch <- prometheus.MustNewConstMetric(c.currBatchesBlockedByHerder, prometheus.GaugeValue, fts.(map[string]interface{})["curr_batches_blocked_by_herder"].(float64))
	ch <- prometheus.MustNewConstMetric(c.numBytesUsedRam, prometheus.GaugeValue, fts.(map[string]interface{})["num_bytes_used_ram"].(float64))
	ch <- prometheus.MustNewConstMetric(c.pctCpuGc, prometheus.GaugeValue, fts.(map[string]interface{})["pct_cpu_gc"].(float64))
	ch <- prometheus.MustNewConstMetric(c.totBatchesFlushedOnMaxops, prometheus.GaugeValue, fts.(map[string]interface{})["tot_batches_flushed_on_maxops"].(float64))
	ch <- prometheus.MustNewConstMetric(c.totatchesFlushedOnTimer, prometheus.GaugeValue, fts.(map[string]interface{})["tot_batches_flushed_on_timer"].(float64))
	ch <- prometheus.MustNewConstMetric(c.totBleveDestOpened, prometheus.GaugeValue, fts.(map[string]interface{})["tot_bleve_dest_opened"].(float64))
	ch <- prometheus.MustNewConstMetric(c.totBleveDestClosed, prometheus.GaugeValue, fts.(map[string]interface{})["tot_bleve_dest_closed"].(float64))
	ch <- prometheus.MustNewConstMetric(c.totGrpcListenersClosed, prometheus.GaugeValue, fts.(map[string]interface{})["tot_grpc_listeners_closed"].(float64))
	ch <- prometheus.MustNewConstMetric(c.totGrpcListenersOpened, prometheus.GaugeValue, fts.(map[string]interface{})["tot_grpc_listeners_opened"].(float64))
	ch <- prometheus.MustNewConstMetric(c.totGrpcQueryrejectOnMemquota, prometheus.GaugeValue, fts.(map[string]interface{})["tot_grpc_queryreject_on_memquota"].(float64))
	ch <- prometheus.MustNewConstMetric(c.totGrpcsListenersClosed, prometheus.GaugeValue, fts.(map[string]interface{})["tot_grpcs_listeners_closed"].(float64))
	ch <- prometheus.MustNewConstMetric(c.totGrpcsListenersOpened, prometheus.GaugeValue, fts.(map[string]interface{})["tot_grpcs_listeners_opened"].(float64))
	ch <- prometheus.MustNewConstMetric(c.totHttpLimitlistenersClosed, prometheus.GaugeValue, fts.(map[string]interface{})["tot_http_limitlisteners_closed"].(float64))
	ch <- prometheus.MustNewConstMetric(c.totHttpLimitlistenersOpened, prometheus.GaugeValue, fts.(map[string]interface{})["tot_http_limitlisteners_opened"].(float64))
	ch <- prometheus.MustNewConstMetric(c.totHttpsLimitlistenersClosed, prometheus.GaugeValue, fts.(map[string]interface{})["tot_https_limitlisteners_closed"].(float64))
	ch <- prometheus.MustNewConstMetric(c.totHttpsLimitlistenersOpened, prometheus.GaugeValue, fts.(map[string]interface{})["tot_https_limitlisteners_opened"].(float64))
	ch <- prometheus.MustNewConstMetric(c.totQueryrejectOnMemquota, prometheus.GaugeValue, fts.(map[string]interface{})["tot_queryreject_on_memquota"].(float64))
	ch <- prometheus.MustNewConstMetric(c.totRemoteGrpc, prometheus.GaugeValue, fts.(map[string]interface{})["tot_remote_grpc"].(float64))
	ch <- prometheus.MustNewConstMetric(c.totRemoteGrpcTls, prometheus.GaugeValue, fts.(map[string]interface{})["tot_remote_grpc_tls"].(float64))
	ch <- prometheus.MustNewConstMetric(c.totRemoteHttp, prometheus.GaugeValue, fts.(map[string]interface{})["tot_remote_http"].(float64))
	ch <- prometheus.MustNewConstMetric(c.totRemoteHttp2, prometheus.GaugeValue, fts.(map[string]interface{})["tot_remote_http2"].(float64))
	ch <- prometheus.MustNewConstMetric(c.totalGc, prometheus.GaugeValue, fts.(map[string]interface{})["total_gc"].(float64))

	ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 1)
	ch <- prometheus.MustNewConstMetric(c.scrapeDuration, prometheus.GaugeValue, time.Since(start).Seconds())
}
