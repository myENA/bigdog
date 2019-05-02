package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/wlanscheduler"
)

type WSGWLANSchedulerService struct {
    c *Client
}

func NewWSGWLANSchedulerService (c *Client) *WSGWLANSchedulerService {
    s := new(WSGWLANSchedulerService)
    s.c = c
    return s
}

