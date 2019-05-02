package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/administration"
)

type WSGApplicationLogAndStatusService struct {
    client *Client
}

func NewWSGApplicationLogAndStatusService (client *Client) *WSGApplicationLogAndStatusService {
    s := new(WSGApplicationLogAndStatusService)
    s.client = client
    return s
}

func (ss *WSGService) WSGApplicationLogAndStatusService () *WSGApplicationLogAndStatusService {
    serv := new(WSGApplicationLogAndStatusService)
    serv.client = ss.client
    return serv
}

