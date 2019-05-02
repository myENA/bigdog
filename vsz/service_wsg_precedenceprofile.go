package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGPrecedenceProfileService struct {
    client *Client
}

func NewWSGPrecedenceProfileService (client *Client) *WSGPrecedenceProfileService {
    s := new(WSGPrecedenceProfileService)
    s.client = client
    return s
}

func (ss *WSGService) WSGPrecedenceProfileService () *WSGPrecedenceProfileService {
    serv := new(WSGPrecedenceProfileService)
    serv.client = ss.client
    return serv
}

