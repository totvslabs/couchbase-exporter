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
	FtsMemoryQuota   int    `json:"ftsMemoryQuota"`
	IndexMemoryQuota int    `json:"indexMemoryQuota"`
	MemoryQuota      int    `json:"memoryQuota"`
	RebalanceStatus  string `json:"rebalanceStatus"`
	MaxBucketCount   int    `json:"maxBucketCount"`
	Counters         struct {
		FailoverNode     int `json:"failover_node"`
		RebalanceSuccess int `json:"rebalance_success"`
		RebalanceStart   int `json:"rebalance_start"`
		RebalanceFail    int `json:"rebalance_fail"`
	} `json:"counters"`
	Balanced bool `json:"balanced"` // couchbase 5.1.1
}
