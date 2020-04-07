package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGL2AccessControlService struct {
	apiClient *APIClient
}

func NewWSGL2AccessControlService(c *APIClient) *WSGL2AccessControlService {
	s := new(WSGL2AccessControlService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGL2AccessControlService() *WSGL2AccessControlService {
	return NewWSGL2AccessControlService(ss.apiClient)
}

type WSGL2AccessControlCreateL2AccessControl struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// EtherTypeRestriction
	// restriction of EtherType rule of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all stations listed below
	// Constraints:
	//    - required
	//    - oneof:[ALLOW,BLOCK]
	EtherTypeRestriction *string `json:"etherTypeRestriction" validate:"required,oneof=ALLOW BLOCK"`

	// EtherTypes
	// Constraints:
	//    - nullable
	EtherTypes []*WSGL2AccessControlEtherTypeObject `json:"etherTypes,omitempty" validate:"omitempty,dive"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// Restriction
	// restriction of mac rule of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all stations listed below
	// Constraints:
	//    - required
	//    - oneof:[ALLOW,BLOCK]
	Restriction *string `json:"restriction" validate:"required,oneof=ALLOW BLOCK"`

	// Rules
	// Constraints:
	//    - nullable
	Rules []*WSGL2AccessControlRuleObject `json:"rules,omitempty" validate:"omitempty,dive"`
}

func NewWSGL2AccessControlCreateL2AccessControl() *WSGL2AccessControlCreateL2AccessControl {
	m := new(WSGL2AccessControlCreateL2AccessControl)
	return m
}

type WSGL2AccessControlEtherTypeObject struct {
	// EtherType
	// Constraints:
	//    - nullable
	EtherType *WSGCommonEtherType `json:"etherType,omitempty"`
}

func NewWSGL2AccessControlEtherTypeObject() *WSGL2AccessControlEtherTypeObject {
	m := new(WSGL2AccessControlEtherTypeObject)
	return m
}

type WSGL2AccessControl struct {
	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// EtherTypeRestriction
	// restriction of EtherType rule of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all stations listed below
	// Constraints:
	//    - nullable
	//    - oneof:[ALLOW,BLOCK]
	EtherTypeRestriction *string `json:"etherTypeRestriction,omitempty" validate:"omitempty,oneof=ALLOW BLOCK"`

	// EtherTypes
	// Constraints:
	//    - nullable
	EtherTypes []*WSGL2AccessControlEtherTypeObject `json:"etherTypes,omitempty" validate:"omitempty,dive"`

	// Id
	// identifier of the L2 Access Control
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Restriction
	// restriction of mac rule of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all stations listed below
	// Constraints:
	//    - nullable
	//    - oneof:[ALLOW,BLOCK]
	Restriction *string `json:"restriction,omitempty" validate:"omitempty,oneof=ALLOW BLOCK"`

	// Rules
	// Constraints:
	//    - nullable
	Rules []*WSGL2AccessControlRuleObject `json:"rules,omitempty" validate:"omitempty,dive"`
}

func NewWSGL2AccessControl() *WSGL2AccessControl {
	m := new(WSGL2AccessControl)
	return m
}

type WSGL2AccessControlList struct {
	// Extra
	// Constraints:
	//    - nullable
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGL2AccessControl `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGL2AccessControlList() *WSGL2AccessControlList {
	m := new(WSGL2AccessControlList)
	return m
}

type WSGL2AccessControlModifyL2AccessControl struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// EtherTypeRestriction
	// restriction of EtherType rule of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all stations listed below
	// Constraints:
	//    - required
	//    - oneof:[ALLOW,BLOCK]
	EtherTypeRestriction *string `json:"etherTypeRestriction" validate:"required,oneof=ALLOW BLOCK"`

	// EtherTypes
	// Constraints:
	//    - nullable
	EtherTypes []*WSGL2AccessControlEtherTypeObject `json:"etherTypes,omitempty" validate:"omitempty,dive"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Restriction
	// restriction of mac rule of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all stations listed below
	// Constraints:
	//    - required
	//    - oneof:[ALLOW,BLOCK]
	Restriction *string `json:"restriction" validate:"required,oneof=ALLOW BLOCK"`

	// Rules
	// Constraints:
	//    - nullable
	Rules []*WSGL2AccessControlRuleObject `json:"rules,omitempty" validate:"omitempty,dive"`
}

func NewWSGL2AccessControlModifyL2AccessControl() *WSGL2AccessControlModifyL2AccessControl {
	m := new(WSGL2AccessControlModifyL2AccessControl)
	return m
}

type WSGL2AccessControlRuleObject struct {
	// Mac
	// Constraints:
	//    - required
	Mac *WSGCommonMac `json:"mac" validate:"required"`
}

func NewWSGL2AccessControlRuleObject() *WSGL2AccessControlRuleObject {
	m := new(WSGL2AccessControlRuleObject)
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
func (s *WSGL2AccessControlService) AddRkszonesL2ACLByZoneId(ctx context.Context, body *WSGPortalServiceCreateL2ACL, zoneId string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesL2ACLByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
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
func (s *WSGL2AccessControlService) DeleteRkszonesL2ACLById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
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
func (s *WSGL2AccessControlService) FindRkszonesL2ACLById(ctx context.Context, id string, zoneId string) (*WSGPortalServiceL2ACL, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
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
func (s *WSGL2AccessControlService) FindRkszonesL2ACLByZoneId(ctx context.Context, zoneId string, optionalParams map[string][]string) (*WSGPortalServiceList, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesL2ACLByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	if v, ok := optionalParams["index"]; ok {
		req.AddQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok {
		req.AddQueryParameter("listSize", v)
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
func (s *WSGL2AccessControlService) FindServicesL2AccessControlByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (interface{}, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesL2AccessControlByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
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
func (s *WSGL2AccessControlService) PartialUpdateRkszonesL2ACLById(ctx context.Context, body *WSGPortalServiceModifyL2ACL, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesL2ACLById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
