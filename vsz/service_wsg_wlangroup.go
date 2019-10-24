package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/wlangroup"
)

type WSGWLANGroupService struct {
	client *Client
}

func NewWSGWLANGroupService(client *Client) *WSGWLANGroupService {
	s := new(WSGWLANGroupService)
	s.client = client
	return s
}

func (ss *WSGService) WSGWLANGroupService() *WSGWLANGroupService {
	serv := new(WSGWLANGroupService)
	serv.client = ss.client
	return serv
}

func (s *WSGWLANGroupService) AddRkszonesWlangroupsMembersById(ctx context.Context, body *wlangroup.WlanMember, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGWLANGroupService) DeleteRkszonesWlangroupsMembersByMemberId(ctx context.Context, pId string, pMemberId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGWLANGroupService) DeleteRkszonesWlangroupsMembersNasIdByMemberId(ctx context.Context, pId string, pMemberId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGWLANGroupService) DeleteRkszonesWlangroupsMembersVlanOverrideByMemberId(ctx context.Context, pId string, pMemberId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGWLANGroupService) FindRkszonesWlangroupsById(ctx context.Context, pId string, pZoneId string) (*wlangroup.WlanGroup, error) {
}

func (s *WSGWLANGroupService) FindRkszonesWlangroupsByZoneId(ctx context.Context, pZoneId string, qIndex string, qListSize string) (*wlangroup.WlanGroupList, error) {
}
