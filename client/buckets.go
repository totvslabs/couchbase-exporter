package client

import "github.com/pkg/errors"

// Buckets returns the results of /pools/default/buckets
func (c Client) Buckets() ([]Bucket, error) {
	var buckets []Bucket
	err := c.get("/pools/default/buckets", &buckets)
	return buckets, errors.Wrap(err, "failed to get buckets")
}

// Bucket response
type Bucket struct {
	Name          string `json:"name,omitempty"`
	ReplicaNumber int64  `json:"replica_number,omitempty"`
	ThreadsNumber int64  `json:"threads_number,omitempty"`
	Quota         struct {
		RAM    int64 `json:"ram,omitempty"`
		RawRAM int64 `json:"rawRAM,omitempty"`
	} `json:"quota,omitempty"`
	BasicStats struct {
		QuotaPercentUsed float64 `json:"quotaPercentUsed,omitempty"`
		OpsPerSec        int64   `json:"opsPerSec,omitempty"`
		DiskFetchs       int64   `json:"diskFetchs,omitempty"`
		ItemCount        int64   `json:"itemCount,omitempty"`
		DiskUsed         int64   `json:"diskUsed,omitempty"`
		DataUsed         int64   `json:"dataUsed,omitempty"`
		MemUsed          int64   `json:"memUsed,omitempty"`
	} `json:"basicStats,omitempty"`
	Nodes []struct {
		SystemStats struct {
			SwapTotal int64 `json:"swap_total,omitempty"`
			SwapUsed  int64 `json:"swap_used,omitempty"`
			MemTotal  int64 `json:"mem_total,omitempty"`
			MemFree   int64 `json:"mem_free,omitempty"`
		} `json:"system_stats,omitempty"`
		InterestingStats struct {
			CmdGet                   int64 `json:"cmd_get,omitempty"`
			CouchDocsActualDiskSize  int64 `json:"couch_docs_actual_disk_size,omitempty"`
			CouchDocsDataSize        int64 `json:"couch_docs_data_size,omitempty"`
			CouchSpatialDataSize     int64 `json:"couch_spatial_data_size,omitempty"`
			CouchSpatialDiskSize     int64 `json:"couch_spatial_disk_size,omitempty"`
			CouchViewsActualDiskSize int64 `json:"couch_views_actual_disk_size,omitempty"`
			CouchViewsDataSize       int64 `json:"couch_views_data_size,omitempty"`
			CurrItems                int64 `json:"curr_items,omitempty"`
			CurrItemsTot             int64 `json:"curr_items_tot,omitempty"`
			EpBgFetched              int64 `json:"ep_bg_fetched,omitempty"`
			GetHits                  int64 `json:"get_hits,omitempty"`
			MemUsed                  int64 `json:"mem_used,omitempty"`
			Ops                      int64 `json:"ops,omitempty"`
			VbReplicaCurrItems       int64 `json:"vb_replica_curr_items,omitempty"`
		} `json:"interesting_stats,omitempty"`
		MemoryTotal int64  `json:"memory_total,omitempty"`
		MemoryFree  int64  `json:"memory_free,omitempty"`
		Replication int64  `json:"replication,omitempty"`
		Status      string `json:"status,omitempty"`
		OptNode     string `json:"opt_node,omitempty"`
		Hostname    string `json:"hostname,omitempty"`
	} `json:"nodes,omitempty"`
}
