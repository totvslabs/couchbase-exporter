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
	RemoteClusters struct {
		URI         string `json:"uri"`
		ValidateURI string `json:"validateURI"`
	} `json:"remoteClusters"`
	RebalanceStatus        string  `json:"rebalanceStatus"`
	MaxBucketCount         float64 `json:"maxBucketCount"`
	AutoCompactionSettings struct {
		ParallelDBAndViewCompaction    bool `json:"parallelDBAndViewCompaction"`
		DatabaseFragmentationThreshold struct {
			Percentage float64 `json:"percentage"`
			Size       string  `json:"size"`
		} `json:"databaseFragmentationThreshold"`
		ViewFragmentationThreshold struct {
			Percentage float64 `json:"percentage"`
			Size       string  `json:"size"`
		} `json:"viewFragmentationThreshold"`
		IndexCompactionMode     string `json:"indexCompactionMode"`
		IndexCircularCompaction struct {
			DaysOfWeek string `json:"daysOfWeek"`
			Interval   struct {
				FromHour     float64 `json:"fromHour"`
				ToHour       float64 `json:"toHour"`
				FromMinute   float64 `json:"fromMinute"`
				ToMinute     float64 `json:"toMinute"`
				AbortOutside bool    `json:"abortOutside"`
			} `json:"interval"`
		} `json:"indexCircularCompaction"`
		IndexFragmentationThreshold struct {
			Percentage float64 `json:"percentage"`
		} `json:"indexFragmentationThreshold"`
	} `json:"autoCompactionSettings"`
	Tasks struct {
		URI string `json:"uri"`
	} `json:"tasks"`
	Counters            Counters `json:"counters"`
	IndexStatusURI      string   `json:"indexStatusURI"`
	CheckPermissionsURI string   `json:"checkPermissionsURI"`
	ServerGroupsURI     string   `json:"serverGroupsUri"`
	ClusterName         string   `json:"clusterName"`
	Balanced            bool     `json:"balanced"`
}
