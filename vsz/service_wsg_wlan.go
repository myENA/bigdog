package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/wlan"
	"gopkg.in/go-playground/validator.v9"
)

type WSGWLANService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGWLANService(c *APIClient) *WSGWLANService {
	s := new(WSGWLANService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGWLANService() *WSGWLANService {
	return NewWSGWLANService(ss.apiClient)
}

// AddRkszonesWlansByZoneId
//
// Use this API command to create a new standard, open and non-tunneled basic WLAN.
//
// Request Body:
//	 - body *wlan.CreateStandardOpenWlan
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansByZoneId(ctx context.Context, body *wlan.CreateStandardOpenWlan, pZoneId string) (*common.CreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}

// AddRkszonesWlansGuestByZoneId
//
// Use this API command to create a new guest access WLAN.
//
// Request Body:
//	 - body *wlan.CreateGuestAccessWlan
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansGuestByZoneId(ctx context.Context, body *wlan.CreateGuestAccessWlan, pZoneId string) (*common.CreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}

// AddRkszonesWlansHotspot20ByZoneId
//
// Use this API command to create a new Hotspot 2.0 access WLAN.
//
// Request Body:
//	 - body *wlan.CreateHotspot20Wlan
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansHotspot20ByZoneId(ctx context.Context, body *wlan.CreateHotspot20Wlan, pZoneId string) (*common.CreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}

// AddRkszonesWlansHotspot20openByZoneId
//
// Use this API command to create a new Hotspot 2.0 Onboarding WLAN with Authentication Method as 'Open'.
//
// Request Body:
//	 - body *wlan.CreateHotspot20OpenWlan
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansHotspot20openByZoneId(ctx context.Context, body *wlan.CreateHotspot20OpenWlan, pZoneId string) (*common.CreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}

// AddRkszonesWlansHotspot20osenByZoneId
//
// Use this API command to create a new Hotspot 2.0 Onboarding WLAN with Authentication Method as '802.1X'.
//
// Request Body:
//	 - body *wlan.CreateHotspot20OpenWlan
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansHotspot20osenByZoneId(ctx context.Context, body *wlan.CreateHotspot20OpenWlan, pZoneId string) (*common.CreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}

// AddRkszonesWlansQosMapsById
//
// Use this API command to enable Qos Map Set of a WLAN.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansQosMapsById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddRkszonesWlansStandard8021XByZoneId
//
// Use this API command to create a new standard, 802.1X and non-tunneled WLAN.
//
// Request Body:
//	 - body *wlan.CreateStandard80211Wlan
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansStandard8021XByZoneId(ctx context.Context, body *wlan.CreateStandard80211Wlan, pZoneId string) (*common.CreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}

// AddRkszonesWlansStandard8021XmacByZoneId
//
// Use this API command to create a new standard, 802.1X with MAC address and non-tunneled WLAN.
//
// Request Body:
//	 - body *wlan.CreateStandard80211Wlan
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansStandard8021XmacByZoneId(ctx context.Context, body *wlan.CreateStandard80211Wlan, pZoneId string) (*common.CreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}

// AddRkszonesWlansStandardmacByZoneId
//
// Use this API command to create a new standard, MAC auth and non-tunneled WLAN.
//
// Request Body:
//	 - body *wlan.CreateStandard80211Wlan
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansStandardmacByZoneId(ctx context.Context, body *wlan.CreateStandard80211Wlan, pZoneId string) (*common.CreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}

// AddRkszonesWlansWebauthByZoneId
//
// Use this API command to creates new web authentication WLAN.
//
// Request Body:
//	 - body *wlan.CreateWebAuthWlan
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansWebauthByZoneId(ctx context.Context, body *wlan.CreateWebAuthWlan, pZoneId string) (*common.CreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}

// AddRkszonesWlansWechatByZoneId
//
// Use this API command to create a new wechat WLAN.
//
// Request Body:
//	 - body *wlan.CreateWechatWlan
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansWechatByZoneId(ctx context.Context, body *wlan.CreateWechatWlan, pZoneId string) (*common.CreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}

// AddRkszonesWlansWispr8021XByZoneId
//
// Use this API command to create a new hotspot (WISPr) with 802.1X WLAN.
//
// Request Body:
//	 - body *wlan.CreateHotspotWlan
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansWispr8021XByZoneId(ctx context.Context, body *wlan.CreateHotspotWlan, pZoneId string) (*common.CreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}

// AddRkszonesWlansWisprByZoneId
//
// Use this API command to create new hotspot (WISPr) WLAN.
//
// Request Body:
//	 - body *wlan.CreateHotspotWlan
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansWisprByZoneId(ctx context.Context, body *wlan.CreateHotspotWlan, pZoneId string) (*common.CreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}

// AddRkszonesWlansWisprmacByZoneId
//
// Use this API command to create a new hotspot (WISPr) with MAC bypass WLAN.
//
// Request Body:
//	 - body *wlan.CreateHotspotWlan
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGWLANService) AddRkszonesWlansWisprmacByZoneId(ctx context.Context, body *wlan.CreateHotspotWlan, pZoneId string) (*common.CreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}

// DeleteRkszonesWlansAccountingServiceOrProfileById
//
// Use this API command to disable the accounting of a WLAN.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWLANService) DeleteRkszonesWlansAccountingServiceOrProfileById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesWlansById
//
// Use this API command to delete a WLAN.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWLANService) DeleteRkszonesWlansById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesWlansDevicePolicyById
//
// Use this API command to disable the device policy of a WLAN.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWLANService) DeleteRkszonesWlansDevicePolicyById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesWlansDiffServProfileById
//
// Use this API command to disable the DiffServ profile of a WLAN.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWLANService) DeleteRkszonesWlansDiffServProfileById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesWlansDnsServerProfileById
//
// Use this API command to disable DNS server profile of a WLAN.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWLANService) DeleteRkszonesWlansDnsServerProfileById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesWlansL2ACLById
//
// Use this API command to disable the layer 2 access control list (ACL) configuration of a WLAN.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWLANService) DeleteRkszonesWlansL2ACLById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesWlansQosMapsById
//
// Use this API command to disable Qos Map Set of a WLAN.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWLANService) DeleteRkszonesWlansQosMapsById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesWlansById
//
// Use this API command to retrieve a WLAN.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWLANService) FindRkszonesWlansById(ctx context.Context, pId string, pZoneId string) (*wlan.WlanConfiguration, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesWlansByZoneId
//
// Use this API command to retrieve a list of WLANs within a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGWLANService) FindRkszonesWlansByZoneId(ctx context.Context, pZoneId string, qIndex string, qListSize string) (*wlan.WlanList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateRkszonesWlansById
//
// Use this API command to modify the basic information of a WLAN.
//
// Request Body:
//	 - body *wlan.ModifyWlan
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWLANService) PartialUpdateRkszonesWlansById(ctx context.Context, body *wlan.ModifyWlan, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}

// UpdateRkszonesWlansById
//
// Use this API command to modify entire information of a WLAN.
//
// Request Body:
//	 - body *wlan.ModifyWlan
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWLANService) UpdateRkszonesWlansById(ctx context.Context, body *wlan.ModifyWlan, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}
