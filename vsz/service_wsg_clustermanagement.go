package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/administration"
	"github.com/myENA/ruckus-client/vsz/types/wsg/clusterblade"
	"github.com/myENA/ruckus-client/vsz/types/wsg/clusterredundancy"
)

type WSGClusterManagementService struct {
    client *Client
}

func NewWSGClusterManagementService (client *Client) *WSGClusterManagementService {
    s := new(WSGClusterManagementService)
    s.client = client
    return s
}

func (ss *WSGService) WSGClusterManagementService () *WSGClusterManagementService {
    serv := new(WSGClusterManagementService)
    serv.client = ss.client
    return serv
}

