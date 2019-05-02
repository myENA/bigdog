package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/service"
)

type WSGGGSNPGWServiceService struct {
    client *Client
}

func NewWSGGGSNPGWServiceService (client *Client) *WSGGGSNPGWServiceService {
    s := new(WSGGGSNPGWServiceService)
    s.client = client
    return s
}

func (ss *WSGService) WSGGGSNPGWServiceService () *WSGGGSNPGWServiceService {
    serv := new(WSGGGSNPGWServiceService)
    serv.client = ss.client
    return serv
}

