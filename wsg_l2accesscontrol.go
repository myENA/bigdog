package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGL2accesscontrolService struct {
	apiClient *VSZClient
}

func NewWSGL2accesscontrolService(c *VSZClient) *WSGL2accesscontrolService {
	s := new(WSGL2accesscontrolService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGL2accesscontrolService() *WSGL2accesscontrolService {
	return NewWSGL2accesscontrolService(ss.apiClient)
}

type WSGL2accesscontrolCreateL2AccessControl struct {
	Description **WSGL2accesscontrolCreateL2AccessControl `json:"description,omitempty"`

	DomainId *string `json:"domainId,omitempty"`

	// EtherTypeRestriction
	// restriction of EtherType rule of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all stations listed below
	// Constraints:
	//    - required
	//    - oneof:[ALLOW,BLOCK]
	EtherTypeRestriction *string `json:"etherTypeRestriction"`

	EtherTypes []**WSGL2accesscontrolCreateL2AccessControl `json:"etherTypes,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name **WSGL2accesscontrolCreateL2AccessControl `json:"name"`

	// Restriction
	// restriction of mac rule of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all stations listed below
	// Constraints:
	//    - required
	//    - oneof:[ALLOW,BLOCK]
	Restriction *string `json:"restriction"`

	Rules []**WSGL2accesscontrolCreateL2AccessControl `json:"rules,omitempty"`
}

func NewWSGL2accesscontrolCreateL2AccessControl() *WSGL2accesscontrolCreateL2AccessControl {
	m := new(WSGL2accesscontrolCreateL2AccessControl)
	return m
}

type WSGL2accesscontrolEtherTypeObject struct {
	EtherType **WSGL2accesscontrolEtherTypeObject `json:"etherType,omitempty"`
}

func NewWSGL2accesscontrolEtherTypeObject() *WSGL2accesscontrolEtherTypeObject {
	m := new(WSGL2accesscontrolEtherTypeObject)
	return m
}

type WSGL2accesscontrol struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description **WSGL2accesscontrol `json:"description,omitempty"`

	DomainId *string `json:"domainId,omitempty"`

	// EtherTypeRestriction
	// restriction of EtherType rule of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all stations listed below
	// Constraints:
	//    - oneof:[ALLOW,BLOCK]
	EtherTypeRestriction *string `json:"etherTypeRestriction,omitempty"`

	EtherTypes []**WSGL2accesscontrol `json:"etherTypes,omitempty"`

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

	Name **WSGL2accesscontrol `json:"name,omitempty"`

	// Restriction
	// restriction of mac rule of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all stations listed below
	// Constraints:
	//    - oneof:[ALLOW,BLOCK]
	Restriction *string `json:"restriction,omitempty"`

	Rules []**WSGL2accesscontrol `json:"rules,omitempty"`
}

func NewWSGL2accesscontrol() *WSGL2accesscontrol {
	m := new(WSGL2accesscontrol)
	return m
}

type WSGL2accesscontrolList struct {
	Extra **WSGL2accesscontrolList `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []**WSGL2accesscontrolList `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGL2accesscontrolList() *WSGL2accesscontrolList {
	m := new(WSGL2accesscontrolList)
	return m
}

type WSGL2accesscontrolModifyL2AccessControl struct {
	Description **WSGL2accesscontrolModifyL2AccessControl `json:"description,omitempty"`

	// EtherTypeRestriction
	// restriction of EtherType rule of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all stations listed below
	// Constraints:
	//    - required
	//    - oneof:[ALLOW,BLOCK]
	EtherTypeRestriction *string `json:"etherTypeRestriction"`

	EtherTypes []**WSGL2accesscontrolModifyL2AccessControl `json:"etherTypes,omitempty"`

	Name **WSGL2accesscontrolModifyL2AccessControl `json:"name,omitempty"`

	// Restriction
	// restriction of mac rule of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all stations listed below
	// Constraints:
	//    - required
	//    - oneof:[ALLOW,BLOCK]
	Restriction *string `json:"restriction"`

	Rules []**WSGL2accesscontrolModifyL2AccessControl `json:"rules,omitempty"`
}

func NewWSGL2accesscontrolModifyL2AccessControl() *WSGL2accesscontrolModifyL2AccessControl {
	m := new(WSGL2accesscontrolModifyL2AccessControl)
	return m
}

type WSGL2accesscontrolRuleObject struct {
	// Mac
	// Constraints:
	//    - required
	Mac **WSGL2accesscontrolRuleObject `json:"mac"`
}

func NewWSGL2accesscontrolRuleObject() *WSGL2accesscontrolRuleObject {
	m := new(WSGL2accesscontrolRuleObject)
	return m
}

// AddRkszonesL2ACLByZoneId
//
// Create a new L2 Access Control (for Firmware Versions less than 5.2).
//
// Request Body:
//	 - body *WSGPortalServiceCreateL2ACL
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGL2accesscontrolService) AddRkszonesL2ACLByZoneId(ctx context.Context, body *WSGPortalServiceCreateL2ACL, zoneId string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesL2ACLByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteRkszonesL2ACLById
//
// Delete an L2 Access Control (for Firmware Versions less than 5.2).
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGL2accesscontrolService) DeleteRkszonesL2ACLById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesL2ACLById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindRkszonesL2ACLById
//
// Retrieve an L2 Access Control (for Firmware Versions less than 5.2).
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGL2accesscontrolService) FindRkszonesL2ACLById(ctx context.Context, id string, zoneId string) (*WSGPortalServiceL2ACL, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesL2ACLById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGPortalServiceL2ACL()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesL2ACLByZoneId
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
func (s *WSGL2accesscontrolService) FindRkszonesL2ACLByZoneId(ctx context.Context, zoneId string, optionalParams map[string][]string) (*WSGPortalServiceList, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesL2ACLByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.SetQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.SetQueryParameter("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGPortalServiceList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesL2AccessControlByQueryCriteria
//
// Query L2 AccessControl Profiles with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGL2accesscontrolService) FindServicesL2AccessControlByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesL2AccessControlByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateRkszonesL2ACLById
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
func (s *WSGL2accesscontrolService) PartialUpdateRkszonesL2ACLById(ctx context.Context, body *WSGPortalServiceModifyL2ACL, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesL2ACLById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
