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
		InterestingStats struct {
		} `json:"interestingStats"`
		Uptime               string `json:"uptime"`
		MemoryTotal          int64  `json:"memoryTotal"`
		MemoryFree           int64  `json:"memoryFree"`
		McdMemoryReserved    int    `json:"mcdMemoryReserved"`
		McdMemoryAllocated   int    `json:"mcdMemoryAllocated"`
		CouchAPIBase         string `json:"couchApiBase"`
		OtpCookie            string `json:"otpCookie"`
		ClusterMembership    string `json:"clusterMembership"`
		RecoveryType         string `json:"recoveryType"`
		Status               string `json:"status"`
		OtpNode              string `json:"otpNode"`
		ThisNode             bool   `json:"thisNode,omitempty"`
		Hostname             string `json:"hostname"`
		ClusterCompatibility int    `json:"clusterCompatibility"`
		Version              string `json:"version"`
		Os                   string `json:"os"`
	} `json:"nodes"`
	RebalanceStatus string `json:"rebalanceStatus"`
	MaxBucketCount  int    `json:"maxBucketCount"`
	Counters        struct {
		RebalanceSuccess int `json:"rebalance_success"`
		RebalanceStart   int `json:"rebalance_start"`
	} `json:"counters"`
	ClusterName string `json:"clusterName"`
	Balanced    bool   `json:"balanced"`
}
