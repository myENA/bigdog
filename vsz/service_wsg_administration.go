package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/administration"
)

type WSGAdministrationService struct {
    client *Client
}

func NewWSGAdministrationService (client *Client) *WSGAdministrationService {
    s := new(WSGAdministrationService)
    s.client = client
    return s
}

func (ss *WSGService) WSGAdministrationService () *WSGAdministrationService {
    serv := new(WSGAdministrationService)
    serv.client = ss.client
    return serv
}

