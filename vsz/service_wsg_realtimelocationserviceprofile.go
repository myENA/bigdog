package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGRealTimeLocationServiceProfileService struct {
    c *Client
}

func NewWSGRealTimeLocationServiceProfileService (c *Client) *WSGRealTimeLocationServiceProfileService {
    s := new(WSGRealTimeLocationServiceProfileService)
    s.c = c
    return s
}

