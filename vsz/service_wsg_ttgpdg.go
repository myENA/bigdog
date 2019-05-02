package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGTTGPDGService struct {
    client *Client
}

func NewWSGTTGPDGService (client *Client) *WSGTTGPDGService {
    s := new(WSGTTGPDGService)
    s.client = client
    return s
}

func (ss *WSGService) WSGTTGPDGService () *WSGTTGPDGService {
    serv := new(WSGTTGPDGService)
    serv.client = ss.client
    return serv
}

