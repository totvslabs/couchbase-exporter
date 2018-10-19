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
	Nodes            []struct {
		SystemStats struct {
			CPUUtilizationRate float64 `json:"cpu_utilization_rate"`
			SwapTotal          float64 `json:"swap_total"`
			SwapUsed           float64 `json:"swap_used"`
			MemTotal           float64 `json:"mem_total"`
			MemFree            float64 `json:"mem_free"`
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
		Uptime               string  `json:"uptime"`
		MemoryTotal          float64 `json:"memoryTotal"`
		MemoryFree           float64 `json:"memoryFree"`
		McdMemoryReserved    float64 `json:"mcdMemoryReserved"`
		McdMemoryAllocated   float64 `json:"mcdMemoryAllocated"`
		CouchAPIBase         string  `json:"couchApiBase"`
		OtpCookie            string  `json:"otpCookie"`
		ClusterMembership    string  `json:"clusterMembership"`
		RecoveryType         string  `json:"recoveryType"`
		Status               string  `json:"status"`
		OtpNode              string  `json:"otpNode"`
		ThisNode             bool    `json:"thisNode,omitempty"`
		Hostname             string  `json:"hostname"`
		ClusterCompatibility float64 `json:"clusterCompatibility"`
		Version              string  `json:"version"`
		Os                   string  `json:"os"`
		Ports                struct {
			Proxy  float64 `json:"proxy"`
			Direct float64 `json:"direct"`
		} `json:"ports"`
		Services []string `json:"services"`
	} `json:"nodes"`
	Buckets struct {
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
	Counters struct {
		RebalanceStart          int64 `json:"rebalance_start"`
		RebalanceSuccess        int64 `json:"rebalance_success"`
		RebalanceFail           int64 `json:"rebalance_fail"`
		FailoverNode            int64 `json:"failover_node"`
		GracefulFailoverStart   int64 `json:"graceful_failover_start"`
		GracefulFailoverSuccess int64 `json:"graceful_failover_success"`
		GracefulFailoverFail    int64 `json:"graceful_failover_fail"`
	} `json:"counters"`
	IndexStatusURI      string `json:"indexStatusURI"`
	CheckPermissionsURI string `json:"checkPermissionsURI"`
	ServerGroupsURI     string `json:"serverGroupsUri"`
	ClusterName         string `json:"clusterName"`
	Balanced            bool   `json:"balanced"`
}
