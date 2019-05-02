package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/system"
)

type WSGControlPlanesService struct {
    client *Client
}

func NewWSGControlPlanesService (client *Client) *WSGControlPlanesService {
    s := new(WSGControlPlanesService)
    s.client = client
    return s
}

func (ss *WSGService) WSGControlPlanesService () *WSGControlPlanesService {
    serv := new(WSGControlPlanesService)
    serv.client = ss.client
    return serv
}

