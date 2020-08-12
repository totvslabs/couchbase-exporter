package client

import "github.com/pkg/errors"

// FTS returns the results of /pools/default
func (c Client) FTS() (FTS, error) {
	var fts FTS
	err := c.get("/api/nsstats", &fts)
	return fts, errors.Wrap(err, "failed to get fts")
}

// FTS (/api/nsstats)
type FTS interface{}
