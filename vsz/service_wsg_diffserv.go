package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/zone"
)

type WSGDiffServService struct {
    client *Client
}

func NewWSGDiffServService (client *Client) *WSGDiffServService {
    s := new(WSGDiffServService)
    s.client = client
    return s
}

func (ss *WSGService) WSGDiffServService () *WSGDiffServService {
    serv := new(WSGDiffServService)
    serv.client = ss.client
    return serv
}

