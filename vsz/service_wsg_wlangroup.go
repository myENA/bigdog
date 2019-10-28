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

// AddRkszonesWlangroupsMembersById
//
// Use this API command to add a member to a WLAN group.
//
// Request Body:
//	 - body *wlangroup.WlanMember
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWLANGroupService) AddRkszonesWlangroupsMembersById(ctx context.Context, body *wlangroup.WlanMember, pId string, pZoneId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesWlangroupsMembersByMemberId
//
// Use this API command to remove a member from a WLAN group.
//
// Path Parameters:
// - pId string
//		- required
// - pMemberId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWLANGroupService) DeleteRkszonesWlangroupsMembersByMemberId(ctx context.Context, pId string, pMemberId string, pZoneId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesWlangroupsMembersNasIdByMemberId
//
// Use this API command to disable a member NAS-ID override of a WLAN group.
//
// Path Parameters:
// - pId string
//		- required
// - pMemberId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWLANGroupService) DeleteRkszonesWlangroupsMembersNasIdByMemberId(ctx context.Context, pId string, pMemberId string, pZoneId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesWlangroupsMembersVlanOverrideByMemberId
//
// Use this API command to disable a member VLAN override of a WLAN group.
//
// Path Parameters:
// - pId string
//		- required
// - pMemberId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWLANGroupService) DeleteRkszonesWlangroupsMembersVlanOverrideByMemberId(ctx context.Context, pId string, pMemberId string, pZoneId string) (*common.EmptyResult, error) {
}

// FindRkszonesWlangroupsById
//
// Use this API command to retrieve the WLAN group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWLANGroupService) FindRkszonesWlangroupsById(ctx context.Context, pId string, pZoneId string) (*wlangroup.WlanGroup, error) {
}

// FindRkszonesWlangroupsByZoneId
//
// Use this API command to retrieve the list of WLAN groups within a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGWLANGroupService) FindRkszonesWlangroupsByZoneId(ctx context.Context, pZoneId string, qIndex string, qListSize string) (*wlangroup.WlanGroupList, error) {
}
