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
			Total             int64 `json:"total"`
			QuotaTotal        int64 `json:"quotaTotal"`
			QuotaUsed         int64 `json:"quotaUsed"`
			Used              int64 `json:"used"`
			UsedByData        int64 `json:"usedByData"`
			QuotaUsedPerNode  int64 `json:"quotaUsedPerNode"`
			QuotaTotalPerNode int64 `json:"quotaTotalPerNode"`
		} `json:"ram"`
		Hdd struct {
			Total      int64 `json:"total"`
			QuotaTotal int64 `json:"quotaTotal"`
			Used       int64 `json:"used"`
			UsedByData int64 `json:"usedByData"`
			Free       int64 `json:"free"`
		} `json:"hdd"`
	} `json:"storageTotals"`
	SystemStats struct {
		CPUUtilizationRate float64 `json:"cpu_utilization_rate"`
		SwapTotal          int64   `json:"swap_total"`
		SwapUsed           int64   `json:"swap_used"`
		MemTotal           int64   `json:"mem_total"`
		MemFree            int64   `json:"mem_free"`
	} `json:"systemStats"`
	InterestingStats struct {
		CmdGet                   float64 `json:"cmd_get"`
		CouchDocsActualDiskSize  float64 `json:"couch_docs_actual_disk_size"`
		CouchDocsDataSize        float64 `json:"couch_docs_data_size"`
		CouchSpatialDataSize     float64 `json:"couch_spatial_data_size"`
		CouchSpatialDiskSize     float64 `json:"couch_spatial_disk_size"`
		CouchViewsActualDiskSize float64 `json:"couch_views_actual_disk_size"`
		CouchViewsDataSize       float64 `json:"couch_views_data_size"`
		CurrItems                float64 `json:"curr_items"`
		CurrItemsTot             float64 `json:"curr_items_tot"`
		EpBgFetched              float64 `json:"ep_bg_fetched"`
		GetHits                  float64 `json:"get_hits"`
		MemUsed                  float64 `json:"mem_used"`
		Ops                      float64 `json:"ops"`
		VbActiveNumNonResident   float64 `json:"vb_active_num_non_resident"`
		VbReplicaCurrItems       float64 `json:"vb_replica_curr_items"`
	} `json:"interestingStats"`
	MemoryTotal          int64    `json:"memoryTotal"`
	MemoryFree           int64    `json:"memoryFree"`
	McdMemoryReserved    int64    `json:"mcdMemoryReserved"`
	McdMemoryAllocated   int64    `json:"mcdMemoryAllocated"`
	CouchAPIBase         string   `json:"couchApiBase"`
	OtpCookie            string   `json:"otpCookie"`
	ClusterMembership    string   `json:"clusterMembership"`
	RecoveryType         string   `json:"recoveryType"`
	Status               string   `json:"status"`
	OtpNode              string   `json:"otpNode"`
	ThisNode             bool     `json:"thisNode"`
	Hostname             string   `json:"hostname"`
	ClusterCompatibility int      `json:"clusterCompatibility"`
	Version              string   `json:"version"`
	Os                   string   `json:"os"`
	Services             []string `json:"services"`
	FtsMemoryQuota       int      `json:"ftsMemoryQuota"`
	IndexMemoryQuota     int      `json:"indexMemoryQuota"`
	MemoryQuota          int      `json:"memoryQuota"`
}
