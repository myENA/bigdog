package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/northbounddatastreaming"
)

type WSGNorthboundDataStreamingService struct {
    c *Client
}

func NewWSGNorthboundDataStreamingService (c *Client) *WSGNorthboundDataStreamingService {
    s := new(WSGNorthboundDataStreamingService)
    s.c = c
    return s
}

