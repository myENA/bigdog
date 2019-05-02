package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/dpsk"
)

type WSGDynamicPSKService struct {
    client *Client
}

func NewWSGDynamicPSKService (client *Client) *WSGDynamicPSKService {
    s := new(WSGDynamicPSKService)
    s.client = client
    return s
}

func (ss *WSGService) WSGDynamicPSKService () *WSGDynamicPSKService {
    serv := new(WSGDynamicPSKService)
    serv.client = ss.client
    return serv
}

