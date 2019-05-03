package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/dp"
)

type WSGDataPlaneOperationalService struct {
    client *Client
}

func NewWSGDataPlaneOperationalService (client *Client) *WSGDataPlaneOperationalService {
    s := new(WSGDataPlaneOperationalService)
    s.client = client
    return s
}

func (ss *WSGService) WSGDataPlaneOperationalService () *WSGDataPlaneOperationalService {
    serv := new(WSGDataPlaneOperationalService)
    serv.client = ss.client
    return serv
}

func (s *WSGDataPlaneOperationalService) AddDpsSwitchoverCluster (ctx context.Context) (*dp.EmptyResult, error) {
}

