package client

import "github.com/pkg/errors"

// Nodes returns the results of /pools/nodes
func (c Client) Nodes() (Nodes, error) {
	var nodes Nodes
	err := c.get("/pools/nodes", &nodes)
	return nodes, errors.Wrap(err, "failed to get nodes")
}

// Nodes (/pools/nodes)
type Nodes struct {
	StorageTotals struct {
		RAM struct {
			Total             int64 `json:"total"`
			QuotaTotal        int   `json:"quotaTotal"`
			QuotaUsed         int   `json:"quotaUsed"`
			Used              int64 `json:"used"`
			UsedByData        int   `json:"usedByData"`
			QuotaUsedPerNode  int   `json:"quotaUsedPerNode"`
			QuotaTotalPerNode int   `json:"quotaTotalPerNode"`
		} `json:"ram"`
		Hdd struct {
			Total      int64 `json:"total"`
			QuotaTotal int64 `json:"quotaTotal"`
			Used       int64 `json:"used"`
			UsedByData int   `json:"usedByData"`
			Free       int64 `json:"free"`
		} `json:"hdd"`
	} `json:"storageTotals"`
	FtsMemoryQuota   int           `json:"ftsMemoryQuota"`
	IndexMemoryQuota int           `json:"indexMemoryQuota"`
	MemoryQuota      int           `json:"memoryQuota"`
	Name             string        `json:"name"`
	Alerts           []interface{} `json:"alerts"`
	Nodes            []struct {
		SystemStats struct {
			CPUUtilizationRate float64 `json:"cpu_utilization_rate"`
			SwapTotal          int     `json:"swap_total"`
			SwapUsed           int     `json:"swap_used"`
			MemTotal           int64   `json:"mem_total"`
			MemFree            int64   `json:"mem_free"`
		} `json:"systemStats"`
		InterestingStats     InterestingStats `json:"interestingStats"`
		Uptime               string           `json:"uptime"`
		MemoryTotal          int64            `json:"memoryTotal"`
		MemoryFree           int64            `json:"memoryFree"`
		McdMemoryReserved    int              `json:"mcdMemoryReserved"`
		McdMemoryAllocated   int              `json:"mcdMemoryAllocated"`
		CouchAPIBase         string           `json:"couchApiBase"`
		OtpCookie            string           `json:"otpCookie"`
		ClusterMembership    string           `json:"clusterMembership"`
		RecoveryType         string           `json:"recoveryType"`
		Status               string           `json:"status"`
		OtpNode              string           `json:"otpNode"`
		ThisNode             bool             `json:"thisNode,omitempty"`
		Hostname             string           `json:"hostname"`
		ClusterCompatibility int              `json:"clusterCompatibility"`
		Version              string           `json:"version"`
		Os                   string           `json:"os"`
	} `json:"nodes"`
	RebalanceStatus string   `json:"rebalanceStatus"`
	MaxBucketCount  int      `json:"maxBucketCount"`
	Counters        Counters `json:"counters"`
	ClusterName     string   `json:"clusterName"`
	Balanced        bool     `json:"balanced"`
}

type Counters struct {
	RebalanceStart          int64 `json:"rebalance_start"`
	RebalanceSuccess        int64 `json:"rebalance_success"`
	RebalanceFail           int64 `json:"rebalance_fail"`
	FailoverNode            int64 `json:"failover_node"`
	GracefulFailoverStart   int64 `json:"graceful_failover_start"`
	GracefulFailoverSuccess int64 `json:"graceful_failover_success"`
	GracefulFailoverFail    int64 `json:"graceful_failover_fail"`
}

type InterestingStats struct {
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
}
