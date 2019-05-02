package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/wlan"
)

type WSGWLANService struct {
    client *Client
}

func NewWSGWLANService (client *Client) *WSGWLANService {
    s := new(WSGWLANService)
    s.client = client
    return s
}

func (ss *WSGService) WSGWLANService () *WSGWLANService {
    serv := new(WSGWLANService)
    serv.client = ss.client
    return serv
}

func (s *WSGWLANService) AddRkszonesWlansGuestByZoneId (ctx context.Context, zoneId string) (common.CreateResult, error) {
}

func (s *WSGWLANService) AddRkszonesWlansHotspot20ByZoneId (ctx context.Context, zoneId string) (common.CreateResult, error) {
}

func (s *WSGWLANService) AddRkszonesWlansHotspot20openByZoneId (ctx context.Context, zoneId string) (common.CreateResult, error) {
}

func (s *WSGWLANService) AddRkszonesWlansHotspot20osenByZoneId (ctx context.Context, zoneId string) (common.CreateResult, error) {
}

func (s *WSGWLANService) AddRkszonesWlansQosMapsById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGWLANService) AddRkszonesWlansStandard8021XByZoneId (ctx context.Context, zoneId string) (common.CreateResult, error) {
}

func (s *WSGWLANService) AddRkszonesWlansStandard8021XmacByZoneId (ctx context.Context, zoneId string) (common.CreateResult, error) {
}

func (s *WSGWLANService) AddRkszonesWlansStandardmacByZoneId (ctx context.Context, zoneId string) (common.CreateResult, error) {
}

func (s *WSGWLANService) AddRkszonesWlansWebauthByZoneId (ctx context.Context, zoneId string) (common.CreateResult, error) {
}

func (s *WSGWLANService) AddRkszonesWlansWechatByZoneId (ctx context.Context, zoneId string) (common.CreateResult, error) {
}

func (s *WSGWLANService) AddRkszonesWlansWispr8021XByZoneId (ctx context.Context, zoneId string) (common.CreateResult, error) {
}

func (s *WSGWLANService) AddRkszonesWlansWisprByZoneId (ctx context.Context, zoneId string) (common.CreateResult, error) {
}

func (s *WSGWLANService) AddRkszonesWlansWisprmacByZoneId (ctx context.Context, zoneId string) (common.CreateResult, error) {
}

func (s *WSGWLANService) DeleteRkszonesWlansAccountingServiceOrProfileById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGWLANService) DeleteRkszonesWlansDevicePolicyById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGWLANService) DeleteRkszonesWlansDiffServProfileById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGWLANService) DeleteRkszonesWlansDnsServerProfileById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGWLANService) DeleteRkszonesWlansL2ACLById (ctx context.Context, id string, zoneId string) error {
}

func (s *WSGWLANService) FindRkszonesWlansById (ctx context.Context, id string, zoneId string) (wlan.WlanConfiguration, error) {
}

func (s *WSGWLANService) FindRkszonesWlansByZoneId (ctx context.Context, zoneId string) (wlan.WlanList, error) {
}

