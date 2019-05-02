package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGIPSECProfileService struct {
    client *Client
}

func NewWSGIPSECProfileService (client *Client) *WSGIPSECProfileService {
    s := new(WSGIPSECProfileService)
    s.client = client
    return s
}

func (ss *WSGService) WSGIPSECProfileService () *WSGIPSECProfileService {
    serv := new(WSGIPSECProfileService)
    serv.client = ss.client
    return serv
}

