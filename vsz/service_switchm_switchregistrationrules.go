package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/registration"
)

type SwitchMSwitchRegistrationRulesService struct {
	apiClient *APIClient
}

func NewSwitchMSwitchRegistrationRulesService(c *APIClient) *SwitchMSwitchRegistrationRulesService {
	s := new(SwitchMSwitchRegistrationRulesService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchRegistrationRulesService() *SwitchMSwitchRegistrationRulesService {
	serv := new(SwitchMSwitchRegistrationRulesService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddRegistrationRules
//
// Use this API command to create new switch registration rule.
//
// Request Body:
//	 - body *registration.RegistrationRule
func (s *SwitchMSwitchRegistrationRulesService) AddRegistrationRules(ctx context.Context, body *registration.RegistrationRule) (*registration.CreateResult, error) {
}

// DeleteRegistrationRules
//
// Use this API command to delete multiple switch registration rules.
//
// Request Body:
//	 - body registration.RuleUUIDs
func (s *SwitchMSwitchRegistrationRulesService) DeleteRegistrationRules(ctx context.Context, body registration.RuleUUIDs) (*registration.DeleteMultipleResult, error) {
}

// DeleteRegistrationRulesById
//
// Use this API command to delete a switch registration rule.
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchRegistrationRulesService) DeleteRegistrationRulesById(ctx context.Context, pId string) (*registration.DeleteResult, error) {
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
//
// Path Parameters:
// - pId string
//		- required
func (s *SwitchMSwitchRegistrationRulesService) UpdateRegistrationRulesById(ctx context.Context, body *registration.RegistrationRule, pId string) (*registration.ModifyResult, error) {
}
