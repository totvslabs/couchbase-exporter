package client

import "github.com/pkg/errors"

// Nodes returns the results of /nodes/self
func (c Client) Nodes() (Node, error) {
	var nodes Node
	err := c.get("/nodes/self", &nodes)
	return nodes, errors.Wrap(err, "failed to get nodes")
}

// Node (/nodes/self)
type Node struct {
	StorageTotals struct {
		RAM struct {
			Total      int `json:"total"`
			QuotaTotal int `json:"quotaTotal"`
			QuotaUsed  int `json:"quotaUsed"`
			Used       int `json:"used"`
			UsedByData int `json:"usedByData"`
		} `json:"ram"`
		Hdd struct {
			Total      int64 `json:"total"`
			QuotaTotal int64 `json:"quotaTotal"`
			Used       int64 `json:"used"`
			UsedByData int   `json:"usedByData"`
			Free       int64 `json:"free"`
		} `json:"hdd"`
	} `json:"storageTotals"`
	SystemStats struct {
		CPUUtilizationRate float64 `json:"cpu_utilization_rate"`
		SwapTotal          int     `json:"swap_total"`
		SwapUsed           int     `json:"swap_used"`
	} `json:"systemStats"`
	InterestingStats struct {
		CmdGet                       int `json:"cmd_get"`
		CouchDocsActualDiskSize      int `json:"couch_docs_actual_disk_size"`
		CouchDocsDataSize            int `json:"couch_docs_data_size"`
		CouchSpatialDataSize         int `json:"couch_spatial_data_size"`
		CouchSpatialDiskSize         int `json:"couch_spatial_disk_size"`
		CouchViewsActualDiskSize     int `json:"couch_views_actual_disk_size"`
		CouchViewsDataSize           int `json:"couch_views_data_size"`
		CurrItems                    int `json:"curr_items"`
		CurrItemsTot                 int `json:"curr_items_tot"`
		EpBgFetched                  int `json:"ep_bg_fetched"`
		GetHits                      int `json:"get_hits"`
		MemUsed                      int `json:"mem_used"`
		Ops                          int `json:"ops"`
		VbReplicaCurrItems           int `json:"vb_replica_curr_items"`
		VbActiveNumNonResidentNumber int `json:"vb_active_num_non_residentNumber"` // couchbase 5.1.1
	} `json:"interestingStats"`
	Uptime               string   `json:"uptime"`
	McdMemoryReserved    int      `json:"mcdMemoryReserved"`
	McdMemoryAllocated   int      `json:"mcdMemoryAllocated"`
	ClusterMembership    string   `json:"clusterMembership"`
	RecoveryType         string   `json:"recoveryType"`
	Status               string   `json:"status"`
	Hostname             string   `json:"hostname"`
	ClusterCompatibility int      `json:"clusterCompatibility"`
	Version              string   `json:"version"`
	Os                   string   `json:"os"`
	Services             []string `json:"services"`
	FtsMemoryQuota       int      `json:"ftsMemoryQuota"`
	IndexMemoryQuota     int      `json:"indexMemoryQuota"`
	MemoryQuota          int      `json:"memoryQuota"`
}
