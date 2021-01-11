package bigdog

// API Version: v9_1

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type WSGDevicePolicyService struct {
	apiClient *VSZClient
}

func NewWSGDevicePolicyService(c *VSZClient) *WSGDevicePolicyService {
	s := new(WSGDevicePolicyService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDevicePolicyService() *WSGDevicePolicyService {
	return NewWSGDevicePolicyService(ss.apiClient)
}

// WSGDevicePolicyCreateDevicePolicy
//
// Definition: devicePolicy_createDevicePolicy
type WSGDevicePolicyCreateDevicePolicy struct {
	// DefaultAction
	// defaultAction of the device policy config
	// Constraints:
	//    - required
	//    - oneof:[ALLOW,BLOCK]
	DefaultAction *string `json:"defaultAction"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`
}

func NewWSGDevicePolicyCreateDevicePolicy() *WSGDevicePolicyCreateDevicePolicy {
	m := new(WSGDevicePolicyCreateDevicePolicy)
	return m
}

// WSGDevicePolicyPorfile
//
// Definition: devicePolicy_devicePolicyPorfile
type WSGDevicePolicyPorfile struct {
	// DefaultAction
	// defaultAction of the device policy config
	// Constraints:
	//    - oneof:[ALLOW,BLOCK]
	DefaultAction *string `json:"defaultAction,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// identifier of the device policy config
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Rule
	// rule of the device policy config
	Rule []*WSGDevicePolicyRule `json:"rule,omitempty"`
}

type WSGDevicePolicyPorfileAPIResponse struct {
	*RawAPIResponse
	Data *WSGDevicePolicyPorfile
}

func newWSGDevicePolicyPorfileAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGDevicePolicyPorfileAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGDevicePolicyPorfileAPIResponse) Hydrate() error {
	r.Data = new(WSGDevicePolicyPorfile)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGDevicePolicyPorfile() *WSGDevicePolicyPorfile {
	m := new(WSGDevicePolicyPorfile)
	return m
}

// WSGDevicePolicyRule
//
// Definition: devicePolicy_devicePolicyRule
type WSGDevicePolicyRule struct {
	// Action
	// defaultAction of the device policy config
	// Constraints:
	//    - oneof:[ALLOW,BLOCK]
	Action *string `json:"action,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DeviceType
	// deviceType of the device policy rule
	// Constraints:
	//    - oneof:[Windows,Android,Apple_iOS,Mac_OS,Linux,VoIP,Gaming,Printers,BlackBerry,Chrome_OS]
	DeviceType *string `json:"deviceType,omitempty"`

	// Downlink
	// downlink rate Limiting of the device policy rule, rate unit is mbps, value must include be 0, 0.10, 0.25, 0.50, 0.75, 1.00, 1.25,1.50, 1.75, 2.00, 2.25, 2.50, 2.75, 3.00, 3.20, 3.50, 3.75, 4.00, 4.25, 4.50, 4.75, 5.00, 5.25, 5.50, 5.75, 6.00, 6.25, 6.50, 6.75,  7.00, 7.25, 7.50, 7.75, 8.00, 8.25, 8.50, 8.75,  9.00, 9.25, 9.50, 9.75,  10.00, 10.25, 10.50, 10.75, 11.00, 11.25, 11.50, 11.75, 12.00, 12.25, 12.50, 12.75, 13.00, 13.25, 13.50, 13.75, 14.00, 14.25, 14.50, 14.75, 15.00, 15.25, 15.50, 15.75, 16.00, 16.25, 16.50, 16.75, 17.00, 17.25, 17.50, 17.75, 18.00, 18.25, 18.50, 18.75, 19.00, 19.25, 19.50, 19.75, 20.00
	// Constraints:
	//    - min:0.000000
	//    - max:20.000000
	Downlink *float64 `json:"downlink,omitempty"`

	// Uplink
	// uplink rate Limiting of the device policy rule, rate unit is mbps, value must include be 0, 0.10, 0.25, 0.50, 0.75, 1.00, 1.25,1.50, 1.75, 2.00, 2.25, 2.50, 2.75, 3.00, 3.20, 3.50, 3.75, 4.00, 4.25, 4.50, 4.75, 5.00, 5.25, 5.50, 5.75, 6.00, 6.25, 6.50, 6.75,  7.00, 7.25, 7.50, 7.75, 8.00, 8.25, 8.50, 8.75,  9.00, 9.25, 9.50, 9.75,  10.00, 10.25, 10.50, 10.75, 11.00, 11.25, 11.50, 11.75, 12.00, 12.25, 12.50, 12.75, 13.00, 13.25, 13.50, 13.75, 14.00, 14.25, 14.50, 14.75, 15.00, 15.25, 15.50, 15.75, 16.00, 16.25, 16.50, 16.75, 17.00, 17.25, 17.50, 17.75, 18.00, 18.25, 18.50, 18.75, 19.00, 19.25, 19.50, 19.75, 20.00
	// Constraints:
	//    - min:0.000000
	//    - max:20.000000
	Uplink *float64 `json:"uplink,omitempty"`

	// Vlan
	// VLAN Members of the ethernet port profile
	// Constraints:
	//    - nullable
	//    - min:1
	//    - max:4094
	Vlan *int `json:"vlan,omitempty"`
}

func NewWSGDevicePolicyRule() *WSGDevicePolicyRule {
	m := new(WSGDevicePolicyRule)
	return m
}

// WSGDevicePolicyModifyDevicePolicy
//
// Definition: devicePolicy_modifyDevicePolicy
type WSGDevicePolicyModifyDevicePolicy struct {
	// DefaultAction
	// defaultAction of the device policy config
	// Constraints:
	//    - oneof:[ALLOW,BLOCK]
	DefaultAction *string `json:"defaultAction,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Rule
	// rule of the device policy config
	Rule []*WSGDevicePolicyRule `json:"rule,omitempty"`
}

func NewWSGDevicePolicyModifyDevicePolicy() *WSGDevicePolicyModifyDevicePolicy {
	m := new(WSGDevicePolicyModifyDevicePolicy)
	return m
}

// WSGDevicePolicyPorfileList
//
// Definition: devicePolicy_porfileList
type WSGDevicePolicyPorfileList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGDevicePolicyPorfileListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGDevicePolicyPorfileListAPIResponse struct {
	*RawAPIResponse
	Data *WSGDevicePolicyPorfileList
}

func newWSGDevicePolicyPorfileListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGDevicePolicyPorfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGDevicePolicyPorfileListAPIResponse) Hydrate() error {
	r.Data = new(WSGDevicePolicyPorfileList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGDevicePolicyPorfileList() *WSGDevicePolicyPorfileList {
	m := new(WSGDevicePolicyPorfileList)
	return m
}

// WSGDevicePolicyPorfileListType
//
// Definition: devicePolicy_porfileListType
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
// Create a new Device Policy Profile (for Firmware Versions less than 5.2).
//
// Operation ID: addRkszonesDevicePolicyByZoneId
// Operation path: /rkszones/{zoneId}/devicePolicy
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGDevicePolicyCreateDevicePolicy
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGDevicePolicyService) AddRkszonesDevicePolicyByZoneId(ctx context.Context, body *WSGDevicePolicyCreateDevicePolicy, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesDevicePolicyByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteRkszonesDevicePolicyById
//
// Delete Device Policy Profile (for Firmware Versions less than 5.2).
//
// Operation ID: deleteRkszonesDevicePolicyById
// Operation path: /rkszones/{zoneId}/devicePolicy/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDevicePolicyService) DeleteRkszonesDevicePolicyById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesDevicePolicyById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindRkszonesDevicePolicyById
//
// Retrieve a Device Policy Profile (for Firmware Versions less than 5.2).
//
// Operation ID: findRkszonesDevicePolicyById
// Operation path: /rkszones/{zoneId}/devicePolicy/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDevicePolicyService) FindRkszonesDevicePolicyById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGDevicePolicyPorfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDevicePolicyPorfileAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDevicePolicyPorfileAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesDevicePolicyById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDevicePolicyPorfileAPIResponse), err
}

// FindRkszonesDevicePolicyByZoneId
//
// Retrieve a list of Device Policy Profiles within a zone (for Firmware Versions less than 5.2).
//
// Operation ID: findRkszonesDevicePolicyByZoneId
// Operation path: /rkszones/{zoneId}/devicePolicy
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
//
// Optional parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGDevicePolicyService) FindRkszonesDevicePolicyByZoneId(ctx context.Context, zoneId string, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGDevicePolicyPorfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDevicePolicyPorfileListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDevicePolicyPorfileListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesDevicePolicyByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDevicePolicyPorfileListAPIResponse), err
}

// FindServicesDevicePolicyByQueryCriteria
//
// Query Device Policy Profiles with specified filters.
//
// Operation ID: findServicesDevicePolicyByQueryCriteria
// Operation path: /query/services/devicePolicy
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDevicePolicyService) FindServicesDevicePolicyByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindServicesDevicePolicyByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// PartialUpdateRkszonesDevicePolicyById
//
// Modify a specific Device Policy Profile (for Firmware Versions less than 5.2).
//
// Operation ID: partialUpdateRkszonesDevicePolicyById
// Operation path: /rkszones/{zoneId}/devicePolicy/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGDevicePolicyModifyDevicePolicy
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDevicePolicyService) PartialUpdateRkszonesDevicePolicyById(ctx context.Context, body *WSGDevicePolicyModifyDevicePolicy, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateRkszonesDevicePolicyById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
