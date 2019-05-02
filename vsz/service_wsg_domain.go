package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/domain"
)

type WSGDomainService struct {
    client *Client
}

func NewWSGDomainService (client *Client) *WSGDomainService {
    s := new(WSGDomainService)
    s.client = client
    return s
}

func (ss *WSGService) WSGDomainService () *WSGDomainService {
    serv := new(WSGDomainService)
    serv.client = ss.client
    return serv
}

