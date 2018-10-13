package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
)

func (c Client) Tasks() ([]Task, error) {
	resp, err := c.client.Get(c.url("/pools/default/tasks"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tasks")
	}
	fmt.Println(resp)
	var tasks []Task
	bts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}
	defer resp.Body.Close()
	fmt.Println(string(bts))
	if err := json.Unmarshal(bts, &tasks); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshall tasks")
	}
	return tasks, nil
}

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

type NodeProgress struct {
	DocsTotal           int64 `json:"docsTotal,omitempty"`
	DocsTransferred     int64 `json:"docsTransferred,omitempty"`
	ActiveVBucketsLeft  int64 `json:"activeVBucketsLeft,omitempty"`
	ReplicaVBucketsLeft int64 `json:"replicaVBucketsLeft,omitempty"`
}
