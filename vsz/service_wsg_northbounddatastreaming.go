package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/northbounddatastreaming"
)

type WSGNorthboundDataStreamingService struct {
    client *Client
}

func NewWSGNorthboundDataStreamingService (client *Client) *WSGNorthboundDataStreamingService {
    s := new(WSGNorthboundDataStreamingService)
    s.client = client
    return s
}

func (ss *WSGService) WSGNorthboundDataStreamingService () *WSGNorthboundDataStreamingService {
    serv := new(WSGNorthboundDataStreamingService)
    serv.client = ss.client
    return serv
}

