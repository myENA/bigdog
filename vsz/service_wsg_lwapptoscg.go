package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/system"
)

type WSGLWAPPTOSCGService struct {
    client *Client
}

func NewWSGLWAPPTOSCGService (client *Client) *WSGLWAPPTOSCGService {
    s := new(WSGLWAPPTOSCGService)
    s.client = client
    return s
}

func (ss *WSGService) WSGLWAPPTOSCGService () *WSGLWAPPTOSCGService {
    serv := new(WSGLWAPPTOSCGService)
    serv.client = ss.client
    return serv
}

func (s *WSGLWAPPTOSCGService) FindLwapp2scg (ctx context.Context) (system.LwappScgConfiguration, error) {
}

func (s *WSGLWAPPTOSCGService) PartialUpdateLwapp2scgApList (ctx context.Context) error {
}
