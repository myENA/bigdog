package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/avc"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGApplicationVisibilityControlService struct {
    client *Client
}

func NewWSGApplicationVisibilityControlService (client *Client) *WSGApplicationVisibilityControlService {
    s := new(WSGApplicationVisibilityControlService)
    s.client = client
    return s
}

func (ss *WSGService) WSGApplicationVisibilityControlService () *WSGApplicationVisibilityControlService {
    serv := new(WSGApplicationVisibilityControlService)
    serv.client = ss.client
    return serv
}

