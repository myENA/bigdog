package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/domain"
)

type WSGDomainService struct {
    c *Client
}

func NewWSGDomainService (c *Client) *WSGDomainService {
    s := new(WSGDomainService)
    s.c = c
    return s
}

