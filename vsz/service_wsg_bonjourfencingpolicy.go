package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGBonjourFencingPolicyService struct {
    client *Client
}

func NewWSGBonjourFencingPolicyService (client *Client) *WSGBonjourFencingPolicyService {
    s := new(WSGBonjourFencingPolicyService)
    s.client = client
    return s
}

func (ss *WSGService) WSGBonjourFencingPolicyService () *WSGBonjourFencingPolicyService {
    serv := new(WSGBonjourFencingPolicyService)
    serv.client = ss.client
    return serv
}

func (s *WSGBonjourFencingPolicyService) DeleteRkszonesBonjourFencingPolicy (ctx context.Context) (*common.EmptyResult, error) {
}

func (s *WSGBonjourFencingPolicyService) DeleteRkszonesBonjourFencingPolicyById (ctx context.Context, id string) error {
}

func (s *WSGBonjourFencingPolicyService) FindApsBonjourFencingStatisticByApMac (ctx context.Context, apMac string) (*profile.BonjourFencingStatistic, error) {
}

func (s *WSGBonjourFencingPolicyService) FindRkszonesBonjourFencingPolicyById (ctx context.Context, id string, zoneId string) (*profile.BonjourFencingPolicy, error) {
}

func (s *WSGBonjourFencingPolicyService) FindRkszonesBonjourFencingPolicyByZoneId (ctx context.Context, zoneId string) (*profile.BonjourFencingPolicyList, error) {
}

func (s *WSGBonjourFencingPolicyService) FindServicesBonjourFencingPolicyByQueryCriteria (ctx context.Context) (*profile.BonjourFencingPolicyList, error) {
}

