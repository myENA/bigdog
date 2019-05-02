package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/administration"
	"github.com/myENA/ruckus-client/vsz/types/wsg/clusterblade"
	"github.com/myENA/ruckus-client/vsz/types/wsg/clusterredundancy"
)

type WSGClusterManagementService struct {
    c *Client
}

func NewWSGClusterManagementService (c *Client) *WSGClusterManagementService {
    s := new(WSGClusterManagementService)
    s.c = c
    return s
}

