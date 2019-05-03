package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/wlangroup"
)

type WSGWLANGroupService struct {
    client *Client
}

func NewWSGWLANGroupService (client *Client) *WSGWLANGroupService {
    s := new(WSGWLANGroupService)
    s.client = client
    return s
}

func (ss *WSGService) WSGWLANGroupService () *WSGWLANGroupService {
    serv := new(WSGWLANGroupService)
    serv.client = ss.client
    return serv
}

func (s *WSGWLANGroupService) AddRkszonesWlangroupsMembersById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGWLANGroupService) DeleteRkszonesWlangroupsMembersByMemberId (ctx context.Context, id string, memberId string, zoneId string) error {
}

func (s *WSGWLANGroupService) DeleteRkszonesWlangroupsMembersNasIdByMemberId (ctx context.Context, id string, memberId string, zoneId string) error {
}

func (s *WSGWLANGroupService) DeleteRkszonesWlangroupsMembersVlanOverrideByMemberId (ctx context.Context, id string, memberId string, zoneId string) error {
}

func (s *WSGWLANGroupService) FindRkszonesWlangroupsById (ctx context.Context, id string, zoneId string) (wlangroup.WlanGroup, error) {
}

func (s *WSGWLANGroupService) FindRkszonesWlangroupsByZoneId (ctx context.Context, zoneId string) (wlangroup.WlanGroupList, error) {
}

