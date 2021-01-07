package bigdog

// API Version: v9_1

import (
	"context"
	"encoding/json"
	"net/http"
)

type WSGL2AccessControlService struct {
	apiClient *VSZClient
}

func NewWSGL2AccessControlService(c *VSZClient) *WSGL2AccessControlService {
	s := new(WSGL2AccessControlService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGL2AccessControlService() *WSGL2AccessControlService {
	return NewWSGL2AccessControlService(ss.apiClient)
}

// WSGL2AccessControlCreateL2AccessControl
//
// Definition: l2AccessControl_createL2AccessControl
type WSGL2AccessControlCreateL2AccessControl struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	DomainId *string `json:"domainId,omitempty"`

	// EtherTypeRestriction
	// restriction of EtherType rule of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all stations listed below
	// Constraints:
	//    - required
	//    - oneof:[ALLOW,BLOCK]
	EtherTypeRestriction *string `json:"etherTypeRestriction"`

	EtherTypes []*WSGL2AccessControlEtherTypeObject `json:"etherTypes,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// Restriction
	// restriction of mac rule of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all stations listed below
	// Constraints:
	//    - required
	//    - oneof:[ALLOW,BLOCK]
	Restriction *string `json:"restriction"`

	Rules []*WSGL2AccessControlRuleObject `json:"rules,omitempty"`

	UserDefinedEtherTypes []*WSGL2AccessControlUserDefinedEtherTypeObject `json:"userDefinedEtherTypes,omitempty"`
}

func NewWSGL2AccessControlCreateL2AccessControl() *WSGL2AccessControlCreateL2AccessControl {
	m := new(WSGL2AccessControlCreateL2AccessControl)
	return m
}

// WSGL2AccessControlEtherTypeObject
//
// Definition: l2AccessControl_etherTypeObject
type WSGL2AccessControlEtherTypeObject struct {
	EtherType *WSGCommonEtherType `json:"etherType,omitempty"`
}

func NewWSGL2AccessControlEtherTypeObject() *WSGL2AccessControlEtherTypeObject {
	m := new(WSGL2AccessControlEtherTypeObject)
	return m
}

// WSGL2AccessControl
//
// Definition: l2AccessControl_l2AccessControl
type WSGL2AccessControl struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	DomainId *string `json:"domainId,omitempty"`

	// EtherTypeRestriction
	// restriction of EtherType rule of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all stations listed below
	// Constraints:
	//    - oneof:[ALLOW,BLOCK]
	EtherTypeRestriction *string `json:"etherTypeRestriction,omitempty"`

	EtherTypes []*WSGL2AccessControlEtherTypeObject `json:"etherTypes,omitempty"`

	// Id
	// identifier of the L2 Access Control
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Restriction
	// restriction of mac rule of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all stations listed below
	// Constraints:
	//    - oneof:[ALLOW,BLOCK]
	Restriction *string `json:"restriction,omitempty"`

	Rules []*WSGL2AccessControlRuleObject `json:"rules,omitempty"`

	UserDefinedEtherTypes []*WSGL2AccessControlUserDefinedEtherTypeObject `json:"userDefinedEtherTypes,omitempty"`
}

type WSGL2AccessControlAPIResponse struct {
	*RawAPIResponse
	Data *WSGL2AccessControl
}

func newWSGL2AccessControlAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGL2AccessControlAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGL2AccessControlAPIResponse) Hydrate() error {
	r.Data = new(WSGL2AccessControl)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGL2AccessControl() *WSGL2AccessControl {
	m := new(WSGL2AccessControl)
	return m
}

// WSGL2AccessControlList
//
// Definition: l2AccessControl_l2AccessControlList
type WSGL2AccessControlList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGL2AccessControl `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGL2AccessControlListAPIResponse struct {
	*RawAPIResponse
	Data *WSGL2AccessControlList
}

func newWSGL2AccessControlListAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGL2AccessControlListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGL2AccessControlListAPIResponse) Hydrate() error {
	r.Data = new(WSGL2AccessControlList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGL2AccessControlList() *WSGL2AccessControlList {
	m := new(WSGL2AccessControlList)
	return m
}

// WSGL2AccessControlModifyL2AccessControl
//
// Definition: l2AccessControl_modifyL2AccessControl
type WSGL2AccessControlModifyL2AccessControl struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// EtherTypeRestriction
	// restriction of EtherType rule of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all stations listed below
	// Constraints:
	//    - required
	//    - oneof:[ALLOW,BLOCK]
	EtherTypeRestriction *string `json:"etherTypeRestriction"`

	EtherTypes []*WSGL2AccessControlEtherTypeObject `json:"etherTypes,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Restriction
	// restriction of mac rule of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all stations listed below
	// Constraints:
	//    - required
	//    - oneof:[ALLOW,BLOCK]
	Restriction *string `json:"restriction"`

	Rules []*WSGL2AccessControlRuleObject `json:"rules,omitempty"`

	UserDefinedEtherTypes []*WSGL2AccessControlUserDefinedEtherTypeObject `json:"userDefinedEtherTypes,omitempty"`
}

func NewWSGL2AccessControlModifyL2AccessControl() *WSGL2AccessControlModifyL2AccessControl {
	m := new(WSGL2AccessControlModifyL2AccessControl)
	return m
}

// WSGL2AccessControlRuleObject
//
// Definition: l2AccessControl_ruleObject
type WSGL2AccessControlRuleObject struct {
	// Mac
	// Constraints:
	//    - required
	Mac *WSGCommonMac `json:"mac"`
}

func NewWSGL2AccessControlRuleObject() *WSGL2AccessControlRuleObject {
	m := new(WSGL2AccessControlRuleObject)
	return m
}

// WSGL2AccessControlUserDefinedEtherTypeObject
//
// Definition: l2AccessControl_userDefinedEtherTypeObject
type WSGL2AccessControlUserDefinedEtherTypeObject struct {
	// EtherType
	// Constraints:
	//    - required
	//    - max:6
	//    - min:1
	EtherType *string `json:"etherType"`

	// ProtocolName
	// Constraints:
	//    - required
	//    - max:32
	//    - min:1
	ProtocolName *string `json:"protocolName"`
}

func NewWSGL2AccessControlUserDefinedEtherTypeObject() *WSGL2AccessControlUserDefinedEtherTypeObject {
	m := new(WSGL2AccessControlUserDefinedEtherTypeObject)
	return m
}

// AddRkszonesL2ACLByZoneId
//
// Operation ID: addRkszonesL2ACLByZoneId
//
// Create a new L2 Access Control (for Firmware Versions less than 5.2).
//
// Request Body:
//	 - body *WSGPortalServiceCreateL2ACL
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGL2AccessControlService) AddRkszonesL2ACLByZoneId(ctx context.Context, body *WSGPortalServiceCreateL2ACL, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesL2ACLByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonCreateResult()
	rm, err = handleAPIResponse(req, http.StatusCreated, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}

// DeleteRkszonesL2ACLById
//
// Operation ID: deleteRkszonesL2ACLById
//
// Delete an L2 Access Control (for Firmware Versions less than 5.2).
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGL2AccessControlService) DeleteRkszonesL2ACLById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesL2ACLById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, s.apiClient.autoHydrate, err)
	return rm, err
}

// FindRkszonesL2ACLById
//
// Operation ID: findRkszonesL2ACLById
//
// Retrieve an L2 Access Control (for Firmware Versions less than 5.2).
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGL2AccessControlService) FindRkszonesL2ACLById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGPortalServiceL2ACLAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGPortalServiceL2ACL
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesL2ACLById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGPortalServiceL2ACL()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}

// FindRkszonesL2ACLByZoneId
//
// Operation ID: findRkszonesL2ACLByZoneId
//
// Retrieve a list of L2 Access Control (for Firmware Versions less than 5.2).
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
func (s *WSGL2AccessControlService) FindRkszonesL2ACLByZoneId(ctx context.Context, zoneId string, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGPortalServiceListAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGPortalServiceList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesL2ACLByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("zoneId", zoneId)
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGPortalServiceList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}

// FindServicesL2AccessControlByQueryCriteria
//
// Operation ID: findServicesL2AccessControlByQueryCriteria
//
// Query L2 AccessControl Profiles with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGL2AccessControlService) FindServicesL2AccessControlByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawAPIResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindServicesL2AccessControlByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawAPIResponse)
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}

// PartialUpdateRkszonesL2ACLById
//
// Operation ID: partialUpdateRkszonesL2ACLById
//
// Modify a specific L2 Access Control basic (for Firmware Versions less than 5.2).
//
// Request Body:
//	 - body *WSGPortalServiceModifyL2ACL
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGL2AccessControlService) PartialUpdateRkszonesL2ACLById(ctx context.Context, body *WSGPortalServiceModifyL2ACL, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateRkszonesL2ACLById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, s.apiClient.autoHydrate, err)
	return rm, err
}
