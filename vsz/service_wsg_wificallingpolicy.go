package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/wificalling"
)

type WSGWiFiCallingPolicyService struct {
	apiClient *APIClient
}

func NewWSGWiFiCallingPolicyService(c *APIClient) *WSGWiFiCallingPolicyService {
	s := new(WSGWiFiCallingPolicyService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGWiFiCallingPolicyService() *WSGWiFiCallingPolicyService {
	serv := new(WSGWiFiCallingPolicyService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddWifiCallingWifiCallingPolicy
//
// Use this API command to Create Wi-Fi Calling Policy.
//
// Request Body:
//	 - body *wificalling.CreateWifiCallingPolicy
func (s *WSGWiFiCallingPolicyService) AddWifiCallingWifiCallingPolicy(ctx context.Context, body *wificalling.CreateWifiCallingPolicy) (*common.CreateResult, error) {
}

// DeleteWifiCallingWifiCallingPolicy
//
// Use this API command to Delete bulk Wi-Fi Calling policies.
//
// Request Body:
//	 - body *wificalling.DeleteBulk
func (s *WSGWiFiCallingPolicyService) DeleteWifiCallingWifiCallingPolicy(ctx context.Context, body *wificalling.DeleteBulk) error {
}

// DeleteWifiCallingWifiCallingPolicyById
//
// Use this API command to Delete a Wi-Fi Calling policy by ID.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGWiFiCallingPolicyService) DeleteWifiCallingWifiCallingPolicyById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// FindWifiCallingByQueryCriteria
//
// Use this API command to Query Wi-Fi Calling Policy List.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGWiFiCallingPolicyService) FindWifiCallingByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*wificalling.WifiCallingPolicyList, error) {
}

// FindWifiCallingWifiCallingPolicy
//
// Use this API command to Retrieve List of Wi-Fi Calling Policy.
//
// Query Parameters:
// - qDomainId string
// - qIndex string
// - qListSize string
func (s *WSGWiFiCallingPolicyService) FindWifiCallingWifiCallingPolicy(ctx context.Context, qDomainId string, qIndex string, qListSize string) (*wificalling.WifiCallingPolicyList, error) {
}

// FindWifiCallingWifiCallingPolicyById
//
// Use this API command to Retrieve Wi-Fi Calling Policy.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGWiFiCallingPolicyService) FindWifiCallingWifiCallingPolicyById(ctx context.Context, pId string) (*wificalling.WifiCallingPolicy, error) {
}

// PartialUpdateWifiCallingWifiCallingPolicyById
//
// Use this API command to Modify a Wi-Fi Calling policy.
//
// Request Body:
//	 - body *wificalling.ModifyWifiCallingPolicy
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGWiFiCallingPolicyService) PartialUpdateWifiCallingWifiCallingPolicyById(ctx context.Context, body *wificalling.ModifyWifiCallingPolicy, pId string) (*common.EmptyResult, error) {
}

// UpdateWifiCallingWifiCallingPolicyById
//
// Use this API command to Modify Entire Wi-Fi Calling policy.
//
// Request Body:
//	 - body *wificalling.ModifyEntireWifiCallingPolicy
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGWiFiCallingPolicyService) UpdateWifiCallingWifiCallingPolicyById(ctx context.Context, body *wificalling.ModifyEntireWifiCallingPolicy, pId string) (*common.EmptyResult, error) {
}
