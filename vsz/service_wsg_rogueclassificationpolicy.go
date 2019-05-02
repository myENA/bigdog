package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGRogueClassificationPolicyService struct {
    client *Client
}

func NewWSGRogueClassificationPolicyService (client *Client) *WSGRogueClassificationPolicyService {
    s := new(WSGRogueClassificationPolicyService)
    s.client = client
    return s
}

func (ss *WSGService) WSGRogueClassificationPolicyService () *WSGRogueClassificationPolicyService {
    serv := new(WSGRogueClassificationPolicyService)
    serv.client = ss.client
    return serv
}

func (s *WSGRogueClassificationPolicyService) FindRkszonesRogueApPoliciesById (ctx context.Context, id string, zoneId string) (profile.RogueApPolicy, error) {
}

func (s *WSGRogueClassificationPolicyService) FindRkszonesRogueApPoliciesByZoneId (ctx context.Context, zoneId string) (profile.RogueApPolicyList, error) {
}

