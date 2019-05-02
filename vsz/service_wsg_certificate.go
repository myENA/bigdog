package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/certificate"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGCertificateService struct {
    c *Client
}

func NewWSGCertificateService (c *Client) *WSGCertificateService {
    s := new(WSGCertificateService)
    s.c = c
    return s
}

