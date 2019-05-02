package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/aprules"
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

func (s *WSGAPRegistrationRulesService) FindApRules (ctx context.Context) (aprules.ApRuleList, error) {
}

func (s *WSGAPRegistrationRulesService) FindApRulesById (ctx context.Context, id string) (aprules.ApRuleConfiguration, error) {
}

func (s *WSGAPRegistrationRulesService) FindApRulesPriorityDownById (ctx context.Context, id string) error {
}

func (s *WSGAPRegistrationRulesService) FindApRulesPriorityUpById (ctx context.Context, id string) error {
}
