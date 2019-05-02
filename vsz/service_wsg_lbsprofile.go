package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGLBSprofileService struct {
    client *Client
}

func NewWSGLBSprofileService (client *Client) *WSGLBSprofileService {
    s := new(WSGLBSprofileService)
    s.client = client
    return s
}

func (ss *WSGService) WSGLBSprofileService () *WSGLBSprofileService {
    serv := new(WSGLBSprofileService)
    serv.client = ss.client
    return serv
}

