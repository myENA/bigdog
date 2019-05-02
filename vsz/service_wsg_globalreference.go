package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/system"
)

type WSGGlobalreferenceService struct {
    client *Client
}

func NewWSGGlobalreferenceService (client *Client) *WSGGlobalreferenceService {
    s := new(WSGGlobalreferenceService)
    s.client = client
    return s
}

func (ss *WSGService) WSGGlobalreferenceService () *WSGGlobalreferenceService {
    serv := new(WSGGlobalreferenceService)
    serv.client = ss.client
    return serv
}

