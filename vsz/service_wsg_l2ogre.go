package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGL2oGREService struct {
    client *Client
}

func NewWSGL2oGREService (client *Client) *WSGL2oGREService {
    s := new(WSGL2oGREService)
    s.client = client
    return s
}

func (ss *WSGService) WSGL2oGREService () *WSGL2oGREService {
    serv := new(WSGL2oGREService)
    serv.client = ss.client
    return serv
}

