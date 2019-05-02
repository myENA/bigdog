package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/wlanscheduler"
)

type WSGWLANSchedulerService struct {
    client *Client
}

func NewWSGWLANSchedulerService (client *Client) *WSGWLANSchedulerService {
    s := new(WSGWLANSchedulerService)
    s.client = client
    return s
}

func (ss *WSGService) WSGWLANSchedulerService () *WSGWLANSchedulerService {
    serv := new(WSGWLANSchedulerService)
    serv.client = ss.client
    return serv
}

