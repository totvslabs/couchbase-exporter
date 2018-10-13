package client

import "github.com/pkg/errors"

// Tasks returns the results of /pools/default/tasks
func (c Client) Tasks() ([]Task, error) {
	var tasks []Task
	err := c.get("/pools/default/tasks", &tasks)
	return tasks, errors.Wrap(err, "failed to get tasks")
}

// Task is a couchbase task
type Task struct {
	Type     string  `json:"type,omitempty"`
	Status   string  `json:"status,omitempty"`
	Progress float64 `json:"progress,omitempty"`

	// compacting stuff
	Bucket       string `json:"bucket,omitempty"`
	ChangesDone  int64  `json:"changesDone,omitempty"`
	TotalChanges int64  `json:"totalChanges,omitempty"`

	// Rebalance stuff
	PerNode map[string]struct {
		Progress float64 `json:"progress,omitempty"`
	} `json:"perNode,omitempty"`
	DetailedProgress struct {
		Bucket       string `json:"bucket,omitempty"`
		BucketNumber int    `json:"bucketNumber,omitempty"`
		BucketCount  int    `json:"bucketCount,omitempty"`
		PerNode      map[string]struct {
			Ingoing  NodeProgress `json:"ingoing,omitempty"`
			Outgoing NodeProgress `json:"outgoing,omitempty"`
		} `json:"perNode,omitempty"`
	} `json:"detailedProgress,omitempty"`
}

// NodeProgress is the ingoing/outgoing detailed progress of a task in a node
type NodeProgress struct {
	DocsTotal           int64 `json:"docsTotal,omitempty"`
	DocsTransferred     int64 `json:"docsTransferred,omitempty"`
	ActiveVBucketsLeft  int64 `json:"activeVBucketsLeft,omitempty"`
	ReplicaVBucketsLeft int64 `json:"replicaVBucketsLeft,omitempty"`
}
