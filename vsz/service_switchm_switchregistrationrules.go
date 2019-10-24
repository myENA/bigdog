package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/registration"
)

type SwitchMSwitchRegistrationRulesService struct {
	client *Client
}

func NewSwitchMSwitchRegistrationRulesService(client *Client) *SwitchMSwitchRegistrationRulesService {
	s := new(SwitchMSwitchRegistrationRulesService)
	s.client = client
	return s
}

func (ss *SwitchMService) SwitchMSwitchRegistrationRulesService() *SwitchMSwitchRegistrationRulesService {
	serv := new(SwitchMSwitchRegistrationRulesService)
	serv.client = ss.client
	return serv
}

// FindRegistrationRules
//
// Use this API command to retrieves all the registration rules configured in SmartZone.
func (s *SwitchMSwitchRegistrationRulesService) FindRegistrationRules(ctx context.Context) (*registration.RuleQueryResultList, error) {
}

// UpdateRegistrationRulesById
//
// Use this API command to modify the registration rule.
//
// Request Body:
//	 - body *registration.RegistrationRule
func (s *SwitchMSwitchRegistrationRulesService) UpdateRegistrationRulesById(ctx context.Context, body *registration.RegistrationRule, pId string) (*registration.ModifyResult, error) {
}
