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

// AddRkszonesWlansGuestByZoneId
//
// Use this API command to create a new guest access WLAN.
//
// Request Body:
//	 - body *wlan.CreateGuestAccessWlan
func (s *WSGWLANService) AddRkszonesWlansGuestByZoneId(ctx context.Context, body *wlan.CreateGuestAccessWlan, pZoneId string) (*common.CreateResult, error) {
}

// AddRkszonesWlansHotspot20ByZoneId
//
// Use this API command to create a new Hotspot 2.0 access WLAN.
//
// Request Body:
//	 - body *wlan.CreateHotspot20Wlan
func (s *WSGWLANService) AddRkszonesWlansHotspot20ByZoneId(ctx context.Context, body *wlan.CreateHotspot20Wlan, pZoneId string) (*common.CreateResult, error) {
}

// AddRkszonesWlansHotspot20openByZoneId
//
// Use this API command to create a new Hotspot 2.0 Onboarding WLAN with Authentication Method as 'Open'.
//
// Request Body:
//	 - body *wlan.CreateHotspot20OpenWlan
func (s *WSGWLANService) AddRkszonesWlansHotspot20openByZoneId(ctx context.Context, body *wlan.CreateHotspot20OpenWlan, pZoneId string) (*common.CreateResult, error) {
}

// AddRkszonesWlansHotspot20osenByZoneId
//
// Use this API command to create a new Hotspot 2.0 Onboarding WLAN with Authentication Method as '802.1X'.
//
// Request Body:
//	 - body *wlan.CreateHotspot20OpenWlan
func (s *WSGWLANService) AddRkszonesWlansHotspot20osenByZoneId(ctx context.Context, body *wlan.CreateHotspot20OpenWlan, pZoneId string) (*common.CreateResult, error) {
}

// AddRkszonesWlansQosMapsById
//
// Use this API command to enable Qos Map Set of a WLAN.
func (s *WSGWLANService) AddRkszonesWlansQosMapsById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

// AddRkszonesWlansStandard8021XByZoneId
//
// Use this API command to create a new standard, 802.1X and non-tunneled WLAN.
//
// Request Body:
//	 - body *wlan.CreateStandard80211Wlan
func (s *WSGWLANService) AddRkszonesWlansStandard8021XByZoneId(ctx context.Context, body *wlan.CreateStandard80211Wlan, pZoneId string) (*common.CreateResult, error) {
}

// AddRkszonesWlansStandard8021XmacByZoneId
//
// Use this API command to create a new standard, 802.1X with MAC address and non-tunneled WLAN.
//
// Request Body:
//	 - body *wlan.CreateStandard80211Wlan
func (s *WSGWLANService) AddRkszonesWlansStandard8021XmacByZoneId(ctx context.Context, body *wlan.CreateStandard80211Wlan, pZoneId string) (*common.CreateResult, error) {
}

// AddRkszonesWlansStandardmacByZoneId
//
// Use this API command to create a new standard, MAC auth and non-tunneled WLAN.
//
// Request Body:
//	 - body *wlan.CreateStandard80211Wlan
func (s *WSGWLANService) AddRkszonesWlansStandardmacByZoneId(ctx context.Context, body *wlan.CreateStandard80211Wlan, pZoneId string) (*common.CreateResult, error) {
}

// AddRkszonesWlansWebauthByZoneId
//
// Use this API command to creates new web authentication WLAN.
//
// Request Body:
//	 - body *wlan.CreateWebAuthWlan
func (s *WSGWLANService) AddRkszonesWlansWebauthByZoneId(ctx context.Context, body *wlan.CreateWebAuthWlan, pZoneId string) (*common.CreateResult, error) {
}

// AddRkszonesWlansWechatByZoneId
//
// Use this API command to create a new wechat WLAN.
//
// Request Body:
//	 - body *wlan.CreateWechatWlan
func (s *WSGWLANService) AddRkszonesWlansWechatByZoneId(ctx context.Context, body *wlan.CreateWechatWlan, pZoneId string) (*common.CreateResult, error) {
}

// AddRkszonesWlansWispr8021XByZoneId
//
// Use this API command to create a new hotspot (WISPr) with 802.1X WLAN.
//
// Request Body:
//	 - body *wlan.CreateHotspotWlan
func (s *WSGWLANService) AddRkszonesWlansWispr8021XByZoneId(ctx context.Context, body *wlan.CreateHotspotWlan, pZoneId string) (*common.CreateResult, error) {
}

// AddRkszonesWlansWisprByZoneId
//
// Use this API command to create new hotspot (WISPr) WLAN.
//
// Request Body:
//	 - body *wlan.CreateHotspotWlan
func (s *WSGWLANService) AddRkszonesWlansWisprByZoneId(ctx context.Context, body *wlan.CreateHotspotWlan, pZoneId string) (*common.CreateResult, error) {
}

// AddRkszonesWlansWisprmacByZoneId
//
// Use this API command to create a new hotspot (WISPr) with MAC bypass WLAN.
//
// Request Body:
//	 - body *wlan.CreateHotspotWlan
func (s *WSGWLANService) AddRkszonesWlansWisprmacByZoneId(ctx context.Context, body *wlan.CreateHotspotWlan, pZoneId string) (*common.CreateResult, error) {
}

// DeleteRkszonesWlansAccountingServiceOrProfileById
//
// Use this API command to disable the accounting of a WLAN.
func (s *WSGWLANService) DeleteRkszonesWlansAccountingServiceOrProfileById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesWlansDevicePolicyById
//
// Use this API command to disable the device policy of a WLAN.
func (s *WSGWLANService) DeleteRkszonesWlansDevicePolicyById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesWlansDiffServProfileById
//
// Use this API command to disable the DiffServ profile of a WLAN.
func (s *WSGWLANService) DeleteRkszonesWlansDiffServProfileById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesWlansDnsServerProfileById
//
// Use this API command to disable DNS server profile of a WLAN.
func (s *WSGWLANService) DeleteRkszonesWlansDnsServerProfileById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesWlansL2ACLById
//
// Use this API command to disable the layer 2 access control list (ACL) configuration of a WLAN.
func (s *WSGWLANService) DeleteRkszonesWlansL2ACLById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

// FindRkszonesWlansById
//
// Use this API command to retrieve a WLAN.
func (s *WSGWLANService) FindRkszonesWlansById(ctx context.Context, pId string, pZoneId string) (*wlan.WlanConfiguration, error) {
}

// FindRkszonesWlansByZoneId
//
// Use this API command to retrieve a list of WLANs within a zone.
func (s *WSGWLANService) FindRkszonesWlansByZoneId(ctx context.Context, pZoneId string, qIndex string, qListSize string) (*wlan.WlanList, error) {
}
