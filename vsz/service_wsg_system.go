package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/system"
)

type WSGSystemService struct {
    client *Client
}

func NewWSGSystemService (client *Client) *WSGSystemService {
    s := new(WSGSystemService)
    s.client = client
    return s
}

func (ss *WSGService) WSGSystemService () *WSGSystemService {
    serv := new(WSGSystemService)
    serv.client = ss.client
    return serv
}

