package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/sci"
)

type WSGSCIService struct {
    client *Client
}

func NewWSGSCIService (client *Client) *WSGSCIService {
    s := new(WSGSCIService)
    s.client = client
    return s
}

func (ss *WSGService) WSGSCIService () *WSGSCIService {
    serv := new(WSGSCIService)
    serv.client = ss.client
    return serv
}

