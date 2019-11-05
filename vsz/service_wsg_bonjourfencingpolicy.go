package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGBonjourFencingPolicyService struct {
	apiClient *APIClient
}

func NewWSGBonjourFencingPolicyService(c *APIClient) *WSGBonjourFencingPolicyService {
	s := new(WSGBonjourFencingPolicyService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGBonjourFencingPolicyService() *WSGBonjourFencingPolicyService {
	serv := new(WSGBonjourFencingPolicyService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddRkszonesBonjourFencingPolicyByZoneId
//
// Use this API command to create Bonjour Fencing Policy.
//
// Request Body:
//	 - body *profile.CreateBonjourFencingPolicy
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGBonjourFencingPolicyService) AddRkszonesBonjourFencingPolicyByZoneId(ctx context.Context, body *profile.CreateBonjourFencingPolicy, pZoneId string) (*common.CreateResult, error) {
}

// DeleteRkszonesBonjourFencingPolicy
//
// Use this API command to delete Bulk Bonjour Fencing Policy.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
func (s *WSGBonjourFencingPolicyService) DeleteRkszonesBonjourFencingPolicy(ctx context.Context, body *common.BulkDeleteRequest) (*common.EmptyResult, error) {
}

// DeleteRkszonesBonjourFencingPolicyById
//
// Use this API command to delete Bonjour Fencing Policy.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGBonjourFencingPolicyService) DeleteRkszonesBonjourFencingPolicyById(ctx context.Context, pId string) error {
}

// FindApsBonjourFencingStatisticByApMac
//
// Use this API command to get Bonjour Fencing Statistic per AP.
//
// Path Parameters:
// - pApMac string
//		- required
func (s *WSGBonjourFencingPolicyService) FindApsBonjourFencingStatisticByApMac(ctx context.Context, pApMac string) (*profile.BonjourFencingStatistic, error) {
}

// FindRkszonesBonjourFencingPolicyById
//
// Use this API command to retrieve Bonjour Fencing Policy.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGBonjourFencingPolicyService) FindRkszonesBonjourFencingPolicyById(ctx context.Context, pId string, pZoneId string) (*profile.BonjourFencingPolicy, error) {
}

// FindRkszonesBonjourFencingPolicyByZoneId
//
// Use this API command to retrieve a list of Bonjour Fencing Policy.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGBonjourFencingPolicyService) FindRkszonesBonjourFencingPolicyByZoneId(ctx context.Context, pZoneId string) (*profile.BonjourFencingPolicyList, error) {
}

// FindServicesBonjourFencingPolicyByQueryCriteria
//
// Use this API command to retrieve a list of Bonjour Fencing Policy.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGBonjourFencingPolicyService) FindServicesBonjourFencingPolicyByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.BonjourFencingPolicyList, error) {
}

// PartialUpdateRkszonesBonjourFencingPolicyById
//
// Use this API command to modify the basic information of Bonjour Fencing Policy.
//
// Request Body:
//	 - body *profile.ModifyBonjourFencingPolicy
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGBonjourFencingPolicyService) PartialUpdateRkszonesBonjourFencingPolicyById(ctx context.Context, body *profile.ModifyBonjourFencingPolicy, pId string, pZoneId string) (*common.EmptyResult, error) {
}
