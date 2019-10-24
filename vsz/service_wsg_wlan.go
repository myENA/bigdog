package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/wlan"
)

type WSGWLANService struct {
	client *Client
}

func NewWSGWLANService(client *Client) *WSGWLANService {
	s := new(WSGWLANService)
	s.client = client
	return s
}

func (ss *WSGService) WSGWLANService() *WSGWLANService {
	serv := new(WSGWLANService)
	serv.client = ss.client
	return serv
}

func (s *WSGWLANService) AddRkszonesWlansGuestByZoneId(ctx context.Context, body *wlan.CreateGuestAccessWlan, pZoneId string) (*common.CreateResult, error) {
}

func (s *WSGWLANService) AddRkszonesWlansHotspot20ByZoneId(ctx context.Context, body *wlan.CreateHotspot20Wlan, pZoneId string) (*common.CreateResult, error) {
}

func (s *WSGWLANService) AddRkszonesWlansHotspot20openByZoneId(ctx context.Context, body *wlan.CreateHotspot20OpenWlan, pZoneId string) (*common.CreateResult, error) {
}

func (s *WSGWLANService) AddRkszonesWlansHotspot20osenByZoneId(ctx context.Context, body *wlan.CreateHotspot20OpenWlan, pZoneId string) (*common.CreateResult, error) {
}

func (s *WSGWLANService) AddRkszonesWlansQosMapsById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGWLANService) AddRkszonesWlansStandard8021XByZoneId(ctx context.Context, body *wlan.CreateStandard80211Wlan, pZoneId string) (*common.CreateResult, error) {
}

func (s *WSGWLANService) AddRkszonesWlansStandard8021XmacByZoneId(ctx context.Context, body *wlan.CreateStandard80211Wlan, pZoneId string) (*common.CreateResult, error) {
}

func (s *WSGWLANService) AddRkszonesWlansStandardmacByZoneId(ctx context.Context, body *wlan.CreateStandard80211Wlan, pZoneId string) (*common.CreateResult, error) {
}

func (s *WSGWLANService) AddRkszonesWlansWebauthByZoneId(ctx context.Context, body *wlan.CreateWebAuthWlan, pZoneId string) (*common.CreateResult, error) {
}

func (s *WSGWLANService) AddRkszonesWlansWechatByZoneId(ctx context.Context, body *wlan.CreateWechatWlan, pZoneId string) (*common.CreateResult, error) {
}

func (s *WSGWLANService) AddRkszonesWlansWispr8021XByZoneId(ctx context.Context, body *wlan.CreateHotspotWlan, pZoneId string) (*common.CreateResult, error) {
}

func (s *WSGWLANService) AddRkszonesWlansWisprByZoneId(ctx context.Context, body *wlan.CreateHotspotWlan, pZoneId string) (*common.CreateResult, error) {
}

func (s *WSGWLANService) AddRkszonesWlansWisprmacByZoneId(ctx context.Context, body *wlan.CreateHotspotWlan, pZoneId string) (*common.CreateResult, error) {
}

func (s *WSGWLANService) DeleteRkszonesWlansAccountingServiceOrProfileById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGWLANService) DeleteRkszonesWlansDevicePolicyById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGWLANService) DeleteRkszonesWlansDiffServProfileById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGWLANService) DeleteRkszonesWlansDnsServerProfileById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGWLANService) DeleteRkszonesWlansL2ACLById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGWLANService) FindRkszonesWlansById(ctx context.Context, pId string, pZoneId string) (*wlan.WlanConfiguration, error) {
}

func (s *WSGWLANService) FindRkszonesWlansByZoneId(ctx context.Context, pZoneId string, qIndex string, qListSize string) (*wlan.WlanList, error) {
}
