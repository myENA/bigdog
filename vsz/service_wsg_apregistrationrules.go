package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/aprules"
)

type WSGAPRegistrationRulesService struct {
	client *Client
}

func NewWSGAPRegistrationRulesService(client *Client) *WSGAPRegistrationRulesService {
	s := new(WSGAPRegistrationRulesService)
	s.client = client
	return s
}

func (ss *WSGService) WSGAPRegistrationRulesService() *WSGAPRegistrationRulesService {
	serv := new(WSGAPRegistrationRulesService)
	serv.client = ss.client
	return serv
}

func (s *WSGAPRegistrationRulesService) FindApRules(ctx context.Context) (*aprules.ApRuleList, error) {
}

func (s *WSGAPRegistrationRulesService) FindApRulesById(ctx context.Context, pId string) (*aprules.ApRuleConfiguration, error) {
}

func (s *WSGAPRegistrationRulesService) FindApRulesPriorityDownById(ctx context.Context, pId string) error {
}

func (s *WSGAPRegistrationRulesService) FindApRulesPriorityUpById(ctx context.Context, pId string) error {
}
