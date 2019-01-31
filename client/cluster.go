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
			Total             float64 `json:"total"`
			QuotaTotal        float64 `json:"quotaTotal"`
			QuotaUsed         float64 `json:"quotaUsed"`
			Used              float64 `json:"used"`
			UsedByData        float64 `json:"usedByData"`
			QuotaUsedPerNode  float64 `json:"quotaUsedPerNode"`
			QuotaTotalPerNode float64 `json:"quotaTotalPerNode"`
		} `json:"ram"`
		Hdd struct {
			Total      float64 `json:"total"`
			QuotaTotal float64 `json:"quotaTotal"`
			Used       float64 `json:"used"`
			UsedByData float64 `json:"usedByData"`
			Free       float64 `json:"free"`
		} `json:"hdd"`
	} `json:"storageTotals"`
	FtsMemoryQuota   float64       `json:"ftsMemoryQuota"`
	IndexMemoryQuota float64       `json:"indexMemoryQuota"`
	MemoryQuota      float64       `json:"memoryQuota"`
	Name             string        `json:"name"`
	Alerts           []interface{} `json:"alerts"`
	AlertsSilenceURL string        `json:"alertsSilenceURL"`
	Buckets          struct {
		URI                       string `json:"uri"`
		TerseBucketsBase          string `json:"terseBucketsBase"`
		TerseStreamingBucketsBase string `json:"terseStreamingBucketsBase"`
	} `json:"buckets"`
	RebalanceStatus     string   `json:"rebalanceStatus"`
	MaxBucketCount      float64  `json:"maxBucketCount"`
	Counters            Counters `json:"counters"`
	IndexStatusURI      string   `json:"indexStatusURI"`
	CheckPermissionsURI string   `json:"checkPermissionsURI"`
	ServerGroupsURI     string   `json:"serverGroupsUri"`
	ClusterName         string   `json:"clusterName"`
	Balanced            bool     `json:"balanced"`
}
