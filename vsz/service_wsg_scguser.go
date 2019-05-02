package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/scguser"
)

type WSGSCGUserService struct {
    client *Client
}

func NewWSGSCGUserService (client *Client) *WSGSCGUserService {
    s := new(WSGSCGUserService)
    s.client = client
    return s
}

func (ss *WSGService) WSGSCGUserService () *WSGSCGUserService {
    serv := new(WSGSCGUserService)
    serv.client = ss.client
    return serv
}

