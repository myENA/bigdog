package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/aaa"
)

type WSGTestAAAServerService struct {
    client *Client
}

func NewWSGTestAAAServerService (client *Client) *WSGTestAAAServerService {
    s := new(WSGTestAAAServerService)
    s.client = client
    return s
}

func (ss *WSGService) WSGTestAAAServerService () *WSGTestAAAServerService {
    serv := new(WSGTestAAAServerService)
    serv.client = ss.client
    return serv
}

func (s *WSGTestAAAServerService) AddSystemAaaTest (ctx context.Context) (*aaa.TestAAAServerResult, error) {
}

