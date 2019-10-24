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

// FindApRules
//
// Use this API command to retrieve a list of AP Registration Rules profile.
func (s *WSGAPRegistrationRulesService) FindApRules(ctx context.Context) (*aprules.ApRuleList, error) {
}

// FindApRulesById
//
// Use this API command to retrieve AP Registration Rules profile by ID.
func (s *WSGAPRegistrationRulesService) FindApRulesById(ctx context.Context, pId string) (*aprules.ApRuleConfiguration, error) {
}

// FindApRulesPriorityDownById
//
// Use this API command to move Priority Down of AP Registration Rules profile.
func (s *WSGAPRegistrationRulesService) FindApRulesPriorityDownById(ctx context.Context, pId string) error {
}

// FindApRulesPriorityUpById
//
// Use this API command to move Priority Up of AP Registration Rules profile.
func (s *WSGAPRegistrationRulesService) FindApRulesPriorityUpById(ctx context.Context, pId string) error {
}
