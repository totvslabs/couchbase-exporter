package client

import "github.com/pkg/errors"

// Cluster returns the results of /pools/default
func (c Client) Cluster() (Cluster, error) {
	var cluster Cluster
	err := c.get("/pools/default", &cluster)
	return cluster, errors.Wrap(err, "failed to get cluster")
}

// Cluster (/pools/default)
type Cluster struct {
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
	FtsMemoryQuota   int64  `json:"ftsMemoryQuota"`
	IndexMemoryQuota int64  `json:"indexMemoryQuota"`
	MemoryQuota      int64  `json:"memoryQuota"`
	RebalanceStatus  string `json:"rebalanceStatus"`
	MaxBucketCount   int64  `json:"maxBucketCount"`
	Counters         struct {
		FailoverNode     int64 `json:"failover_node"`
		RebalanceSuccess int64 `json:"rebalance_success"`
		RebalanceStart   int64 `json:"rebalance_start"`
		RebalanceFail    int64 `json:"rebalance_fail"`
	} `json:"counters"`
	Balanced bool `json:"balanced"` // couchbase 5.1.1
}
