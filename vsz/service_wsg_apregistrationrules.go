package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/aprules"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGAPRegistrationRulesService struct {
    client *Client
}

func NewWSGAPRegistrationRulesService (client *Client) *WSGAPRegistrationRulesService {
    s := new(WSGAPRegistrationRulesService)
    s.client = client
    return s
}

func (ss *WSGService) WSGAPRegistrationRulesService () *WSGAPRegistrationRulesService {
    serv := new(WSGAPRegistrationRulesService)
    serv.client = ss.client
    return serv
}

