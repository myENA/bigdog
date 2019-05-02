package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGRealTimeLocationServiceProfileService struct {
    client *Client
}

func NewWSGRealTimeLocationServiceProfileService (client *Client) *WSGRealTimeLocationServiceProfileService {
    s := new(WSGRealTimeLocationServiceProfileService)
    s.client = client
    return s
}

func (ss *WSGService) WSGRealTimeLocationServiceProfileService () *WSGRealTimeLocationServiceProfileService {
    serv := new(WSGRealTimeLocationServiceProfileService)
    serv.client = ss.client
    return serv
}

