package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/certificate"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGCertificateService struct {
    client *Client
}

func NewWSGCertificateService (client *Client) *WSGCertificateService {
    s := new(WSGCertificateService)
    s.client = client
    return s
}

func (ss *WSGService) WSGCertificateService () *WSGCertificateService {
    serv := new(WSGCertificateService)
    serv.client = ss.client
    return serv
}

