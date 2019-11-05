package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/aprules"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGAPRegistrationRulesService struct {
	apiClient *APIClient
}

func NewWSGAPRegistrationRulesService(c *APIClient) *WSGAPRegistrationRulesService {
	s := new(WSGAPRegistrationRulesService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAPRegistrationRulesService() *WSGAPRegistrationRulesService {
	serv := new(WSGAPRegistrationRulesService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddApRules
//
// Use this API command to create AP Registration Rules profile.
//
// Request Body:
//	 - body *aprules.CreateApRule
func (s *WSGAPRegistrationRulesService) AddApRules(ctx context.Context, body *aprules.CreateApRule) (*common.CreateResult, error) {
}

// DeleteApRulesById
//
// Use this API command to delete AP Registration Rules profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAPRegistrationRulesService) DeleteApRulesById(ctx context.Context, pId string) error {
}

// FindApRules
//
// Use this API command to retrieve a list of AP Registration Rules profile.
func (s *WSGAPRegistrationRulesService) FindApRules(ctx context.Context) (*aprules.ApRuleList, error) {
}

// FindApRulesById
//
// Use this API command to retrieve AP Registration Rules profile by ID.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAPRegistrationRulesService) FindApRulesById(ctx context.Context, pId string) (*aprules.ApRuleConfiguration, error) {
}

// FindApRulesPriorityDownById
//
// Use this API command to move Priority Down of AP Registration Rules profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAPRegistrationRulesService) FindApRulesPriorityDownById(ctx context.Context, pId string) error {
}

// FindApRulesPriorityUpById
//
// Use this API command to move Priority Up of AP Registration Rules profile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAPRegistrationRulesService) FindApRulesPriorityUpById(ctx context.Context, pId string) error {
}

// PartialUpdateApRulesById
//
// Use this API command to modify the basic information of AP Registration Rules profile.
//
// Request Body:
//	 - body *aprules.ModifyApRule
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAPRegistrationRulesService) PartialUpdateApRulesById(ctx context.Context, body *aprules.ModifyApRule, pId string) (*common.EmptyResult, error) {
}
