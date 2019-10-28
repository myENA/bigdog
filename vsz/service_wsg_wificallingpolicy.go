package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/wificalling"
)

type WSGWiFiCallingPolicyService struct {
	client *Client
}

func NewWSGWiFiCallingPolicyService(client *Client) *WSGWiFiCallingPolicyService {
	s := new(WSGWiFiCallingPolicyService)
	s.client = client
	return s
}

func (ss *WSGService) WSGWiFiCallingPolicyService() *WSGWiFiCallingPolicyService {
	serv := new(WSGWiFiCallingPolicyService)
	serv.client = ss.client
	return serv
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
