package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/aprules"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGAPRegistrationRulesService struct {
    c *Client
}

func NewWSGAPRegistrationRulesService (c *Client) *WSGAPRegistrationRulesService {
    s := new(WSGAPRegistrationRulesService)
    s.c = c
    return s
}

