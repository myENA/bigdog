package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGRogueClassificationPolicyService struct {
	client *Client
}

func NewWSGRogueClassificationPolicyService(client *Client) *WSGRogueClassificationPolicyService {
	s := new(WSGRogueClassificationPolicyService)
	s.client = client
	return s
}

func (ss *WSGService) WSGRogueClassificationPolicyService() *WSGRogueClassificationPolicyService {
	serv := new(WSGRogueClassificationPolicyService)
	serv.client = ss.client
	return serv
}

// FindRkszonesRogueApPoliciesById
//
// Use this API command to retrieve rogue AP policy.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGRogueClassificationPolicyService) FindRkszonesRogueApPoliciesById(ctx context.Context, pId string, pZoneId string) (*profile.RogueApPolicy, error) {
}

// FindRkszonesRogueApPoliciesByZoneId
//
// Use this API command to retrieve a list of rogue AP policy.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGRogueClassificationPolicyService) FindRkszonesRogueApPoliciesByZoneId(ctx context.Context, pZoneId string) (*profile.RogueApPolicyList, error) {
}
