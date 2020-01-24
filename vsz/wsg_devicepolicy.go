package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGDevicePolicyService struct {
	apiClient *APIClient
}

func NewWSGDevicePolicyService(c *APIClient) *WSGDevicePolicyService {
	s := new(WSGDevicePolicyService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDevicePolicyService() *WSGDevicePolicyService {
	return NewWSGDevicePolicyService(ss.apiClient)
}

type WSGDevicePolicyCreateDevicePolicy struct {
	// DefaultAction
	// defaultAction of the device policy cofig
	// Constraints:
	//    - required
	//    - oneof:[ALLOW,BLOCK]
	DefaultAction *string `json:"defaultAction" validate:"required,oneof=ALLOW BLOCK"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`
}

func NewWSGDevicePolicyCreateDevicePolicy() *WSGDevicePolicyCreateDevicePolicy {
	m := new(WSGDevicePolicyCreateDevicePolicy)
	return m
}

type WSGDevicePolicyPorfile struct {
	// DefaultAction
	// defaultAction of the device policy cofig
	// Constraints:
	//    - oneof:[ALLOW,BLOCK]
	DefaultAction *string `json:"defaultAction,omitempty" validate:"oneof=ALLOW BLOCK"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// identifier of the device policy cofig
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Rule
	// rule of the device policy cofig
	Rule []*WSGDevicePolicyRule `json:"rule,omitempty"`
}

func NewWSGDevicePolicyPorfile() *WSGDevicePolicyPorfile {
	m := new(WSGDevicePolicyPorfile)
	return m
}

type WSGDevicePolicyRule struct {
	// Action
	// defaultAction of the device policy cofig
	// Constraints:
	//    - oneof:[ALLOW,BLOCK]
	Action *string `json:"action,omitempty" validate:"oneof=ALLOW BLOCK"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DeviceType
	// deviceType of the device policy rule
	// Constraints:
	//    - oneof:[Windows,Android,Apple_iOS,Mac_OS,Linux,VoIP,Gaming,Printers,BlackBerry,Chrome_OS]
	DeviceType *string `json:"deviceType,omitempty" validate:"oneof=Windows Android Apple_iOS Mac_OS Linux VoIP Gaming Printers BlackBerry Chrome_OS"`

	// Downlink
	// downlink rate Limiting of the device policy rule, rate unit is mbps, value must include be 0, 0.10, 0.25, 0.50, 0.75, 1.00, 1.25,1.50, 1.75, 2.00, 2.25, 2.50, 2.75, 3.00, 3.20, 3.50, 3.75, 4.00, 4.25, 4.50, 4.75, 5.00, 5.25, 5.50, 5.75, 6.00, 6.25, 6.50, 6.75,  7.00, 7.25, 7.50, 7.75, 8.00, 8.25, 8.50, 8.75,  9.00, 9.25, 9.50, 9.75,  10.00, 10.25, 10.50, 10.75, 11.00, 11.25, 11.50, 11.75, 12.00, 12.25, 12.50, 12.75, 13.00, 13.25, 13.50, 13.75, 14.00, 14.25, 14.50, 14.75, 15.00, 15.25, 15.50, 15.75, 16.00, 16.25, 16.50, 16.75, 17.00, 17.25, 17.50, 17.75, 18.00, 18.25, 18.50, 18.75, 19.00, 19.25, 19.50, 19.75, 20.00
	// Constraints:
	//    - min:0.000000
	//    - max:20.000000
	Downlink *float64 `json:"downlink,omitempty" validate:"gte=0.000000,lte=20.000000"`

	// Uplink
	// uplink rate Limiting of the device policy rule, rate unit is mbps, value must include be 0, 0.10, 0.25, 0.50, 0.75, 1.00, 1.25,1.50, 1.75, 2.00, 2.25, 2.50, 2.75, 3.00, 3.20, 3.50, 3.75, 4.00, 4.25, 4.50, 4.75, 5.00, 5.25, 5.50, 5.75, 6.00, 6.25, 6.50, 6.75,  7.00, 7.25, 7.50, 7.75, 8.00, 8.25, 8.50, 8.75,  9.00, 9.25, 9.50, 9.75,  10.00, 10.25, 10.50, 10.75, 11.00, 11.25, 11.50, 11.75, 12.00, 12.25, 12.50, 12.75, 13.00, 13.25, 13.50, 13.75, 14.00, 14.25, 14.50, 14.75, 15.00, 15.25, 15.50, 15.75, 16.00, 16.25, 16.50, 16.75, 17.00, 17.25, 17.50, 17.75, 18.00, 18.25, 18.50, 18.75, 19.00, 19.25, 19.50, 19.75, 20.00
	// Constraints:
	//    - min:0.000000
	//    - max:20.000000
	Uplink *float64 `json:"uplink,omitempty" validate:"gte=0.000000,lte=20.000000"`

	// Vlan
	// VLAN Members of the ethernet port profile
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:4094
	Vlan *int `json:"vlan,omitempty" validate:"omitempty,gte=1,lte=4094"`
}

func NewWSGDevicePolicyRule() *WSGDevicePolicyRule {
	m := new(WSGDevicePolicyRule)
	return m
}

type WSGDevicePolicyModifyDevicePolicy struct {
	// DefaultAction
	// defaultAction of the device policy cofig
	// Constraints:
	//    - oneof:[ALLOW,BLOCK]
	DefaultAction *string `json:"defaultAction,omitempty" validate:"oneof=ALLOW BLOCK"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Rule
	// rule of the device policy cofig
	Rule []*WSGDevicePolicyRule `json:"rule,omitempty"`
}

func NewWSGDevicePolicyModifyDevicePolicy() *WSGDevicePolicyModifyDevicePolicy {
	m := new(WSGDevicePolicyModifyDevicePolicy)
	return m
}

type WSGDevicePolicyPorfileList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGDevicePolicyPorfileListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGDevicePolicyPorfileList() *WSGDevicePolicyPorfileList {
	m := new(WSGDevicePolicyPorfileList)
	return m
}

type WSGDevicePolicyPorfileListType struct {
	// Id
	// Identifier of the service
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGDevicePolicyPorfileListType() *WSGDevicePolicyPorfileListType {
	m := new(WSGDevicePolicyPorfileListType)
	return m
}

// AddRkszonesDevicePolicyByZoneId
//
// Create a new Device Policy Porfile.
//
// Request Body:
//	 - body *WSGDevicePolicyCreateDevicePolicy
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDevicePolicyService) AddRkszonesDevicePolicyByZoneId(ctx context.Context, body *WSGDevicePolicyCreateDevicePolicy, zoneId string) (*WSGCommonCreateResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonCreateResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesDevicePolicyByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteRkszonesDevicePolicyById
//
// Delete Device Policy Porfile.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDevicePolicyService) DeleteRkszonesDevicePolicyById(ctx context.Context, id string, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesDevicePolicyById, true)
}

// FindRkszonesDevicePolicyById
//
// Retrieve a Device Policy Porfile.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDevicePolicyService) FindRkszonesDevicePolicyById(ctx context.Context, id string, zoneId string) (*WSGDevicePolicyPorfile, error) {
	var (
		req  *APIRequest
		resp *WSGDevicePolicyPorfile
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesDevicePolicyById, true)
}

// FindRkszonesDevicePolicyByZoneId
//
// Retrieve a list of Device Policy Porfiles within a zone.
//
// Required Parameters:
// - zoneId string
//		- required
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGDevicePolicyService) FindRkszonesDevicePolicyByZoneId(ctx context.Context, zoneId string, optionalParams map[string]interface{}) (*WSGDevicePolicyPorfileList, error) {
	var (
		req  *APIRequest
		resp *WSGDevicePolicyPorfileList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesDevicePolicyByZoneId, true)
}

// PartialUpdateRkszonesDevicePolicyById
//
// Modify a specific Device Policy Porfile.
//
// Request Body:
//	 - body *WSGDevicePolicyModifyDevicePolicy
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDevicePolicyService) PartialUpdateRkszonesDevicePolicyById(ctx context.Context, body *WSGDevicePolicyModifyDevicePolicy, id string, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesDevicePolicyById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}
