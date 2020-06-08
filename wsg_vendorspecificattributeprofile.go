package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGVendorspecificattributeprofileService struct {
	apiClient *VSZClient
}

func NewWSGVendorspecificattributeprofileService(c *VSZClient) *WSGVendorspecificattributeprofileService {
	s := new(WSGVendorspecificattributeprofileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGVendorspecificattributeprofileService() *WSGVendorspecificattributeprofileService {
	return NewWSGVendorspecificattributeprofileService(ss.apiClient)
}

type WSGVendorspecificattributeprofileCreateResult struct {
	// Id
	// Identifier of the vendor specific attribute profile
	Id *string `json:"id,omitempty"`
}

func NewWSGVendorspecificattributeprofileCreateResult() *WSGVendorspecificattributeprofileCreateResult {
	m := new(WSGVendorspecificattributeprofileCreateResult)
	return m
}

type WSGVendorspecificattributeprofileDeleteBulk struct {
	IdList *WSGVendorspecificattributeprofileDeleteBulk `json:"idList,omitempty"`
}

func NewWSGVendorspecificattributeprofileDeleteBulk() *WSGVendorspecificattributeprofileDeleteBulk {
	m := new(WSGVendorspecificattributeprofileDeleteBulk)
	return m
}

type WSGVendorspecificattributeprofileGet struct {
	// Attributes
	// Vendor specific attribute list for Radius protocol
	Attributes []**WSGVendorspecificattributeprofileGet `json:"attributes,omitempty"`

	Description **WSGVendorspecificattributeprofileGet `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of the vendor specific attribute profile
	Id *string `json:"id,omitempty"`

	Name **WSGVendorspecificattributeprofileGet `json:"name,omitempty"`

	// ZoneId
	// Zone Id
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGVendorspecificattributeprofileGet() *WSGVendorspecificattributeprofileGet {
	m := new(WSGVendorspecificattributeprofileGet)
	return m
}

type WSGVendorspecificattributeprofileList struct {
	// FirstIndex
	// Index of the first profile returned out of the profile list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more profiles after the list that is currently displayed
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Information list of the vendor specific attribute profile
	List []**WSGVendorspecificattributeprofileList `json:"list,omitempty"`

	// TotalCount
	// Total number of the profiles
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGVendorspecificattributeprofileList() *WSGVendorspecificattributeprofileList {
	m := new(WSGVendorspecificattributeprofileList)
	return m
}

type WSGVendorspecificattributeprofileListType struct {
	// Id
	// Identifier of the vendor specific attribute profile
	Id *string `json:"id,omitempty"`

	Name **WSGVendorspecificattributeprofileListType `json:"name,omitempty"`
}

func NewWSGVendorspecificattributeprofileListType() *WSGVendorspecificattributeprofileListType {
	m := new(WSGVendorspecificattributeprofileListType)
	return m
}

type WSGVendorspecificattributeprofilePersist struct {
	// Attributes
	// Vendor specific attribute list for Radius protocol
	// Constraints:
	//    - required
	Attributes []**WSGVendorspecificattributeprofilePersist `json:"attributes"`

	Description **WSGVendorspecificattributeprofilePersist `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name **WSGVendorspecificattributeprofilePersist `json:"name"`
}

func NewWSGVendorspecificattributeprofilePersist() *WSGVendorspecificattributeprofilePersist {
	m := new(WSGVendorspecificattributeprofilePersist)
	return m
}

type WSGVendorspecificattributeprofileQueryCriteriaResult struct {
	Extra **WSGVendorspecificattributeprofileQueryCriteriaResult `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []**WSGVendorspecificattributeprofileQueryCriteriaResult `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGVendorspecificattributeprofileQueryCriteriaResult() *WSGVendorspecificattributeprofileQueryCriteriaResult {
	m := new(WSGVendorspecificattributeprofileQueryCriteriaResult)
	return m
}

type WSGVendorspecificattributeprofileVendorSpecificAttribute struct {
	// KeyId
	// Key ID of vendor specific attribute
	// Constraints:
	//    - required
	KeyId *int `json:"keyId"`

	// SupportedRadiusProtocol
	// The radius protocol to which this given vendor specific attribute will attach
	// Constraints:
	//    - required
	//    - oneof:[ACCOUNTING,AUTHENTICATION,BOTH_AUTHENTICATION_AND_ACCOUNTING]
	SupportedRadiusProtocol *string `json:"supportedRadiusProtocol"`

	// Type
	// Type of vendor specific attribute
	// Constraints:
	//    - required
	//    - oneof:[INTEGER,STRING]
	Type *string `json:"type"`

	// Value
	// Value of vendor specific attribute
	// Constraints:
	//    - required
	Value *string `json:"value"`

	// VendorId
	// Vendor ID of vendor specific attribute
	// Constraints:
	//    - required
	VendorId *int `json:"vendorId"`
}

func NewWSGVendorspecificattributeprofileVendorSpecificAttribute() *WSGVendorspecificattributeprofileVendorSpecificAttribute {
	m := new(WSGVendorspecificattributeprofileVendorSpecificAttribute)
	return m
}

// AddRkszonesVendorSpecificAttributeProfilesByZoneId
//
// Create a vendor specific attribute profile.
//
// Request Body:
//	 - body *WSGVendorspecificattributeprofilePersist
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGVendorspecificattributeprofileService) AddRkszonesVendorSpecificAttributeProfilesByZoneId(ctx context.Context, body *WSGVendorspecificattributeprofilePersist, zoneId string) (*WSGVendorspecificattributeprofileCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGVendorspecificattributeprofileCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesVendorSpecificAttributeProfilesByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGVendorspecificattributeprofileCreateResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteRkszonesVendorSpecificAttributeProfilesById
//
// Use this API command to delete a vendor specific attribute profile by ID.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGVendorspecificattributeprofileService) DeleteRkszonesVendorSpecificAttributeProfilesById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesVendorSpecificAttributeProfilesById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesVendorSpecificAttributeProfilesByZoneId
//
// Use this API command to delete a list of vendor specific attribute profile.
//
// Request Body:
//	 - body *WSGVendorspecificattributeprofileDeleteBulk
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGVendorspecificattributeprofileService) DeleteRkszonesVendorSpecificAttributeProfilesByZoneId(ctx context.Context, body *WSGVendorspecificattributeprofileDeleteBulk, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesVendorSpecificAttributeProfilesByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindRkszonesVendorSpecificAttributeProfilesById
//
// Get a vendor specific attribute profile by ID.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGVendorspecificattributeprofileService) FindRkszonesVendorSpecificAttributeProfilesById(ctx context.Context, id string, zoneId string) (*WSGVendorspecificattributeprofileGet, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGVendorspecificattributeprofileGet
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesVendorSpecificAttributeProfilesById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGVendorspecificattributeprofileGet()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesVendorSpecificAttributeProfilesByQueryCriteria
//
// Use this API command to retrieve a list of vendor specific attribute profile by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGVendorspecificattributeprofileService) FindRkszonesVendorSpecificAttributeProfilesByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGVendorspecificattributeprofileQueryCriteriaResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGVendorspecificattributeprofileQueryCriteriaResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindRkszonesVendorSpecificAttributeProfilesByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGVendorspecificattributeprofileQueryCriteriaResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesVendorSpecificAttributeProfilesByZoneId
//
// Get a ID list of vendor specific attribute profile in this Zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGVendorspecificattributeprofileService) FindRkszonesVendorSpecificAttributeProfilesByZoneId(ctx context.Context, zoneId string) (*WSGVendorspecificattributeprofileList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGVendorspecificattributeprofileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesVendorSpecificAttributeProfilesByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGVendorspecificattributeprofileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateRkszonesVendorSpecificAttributeProfilesById
//
// Use this API command to modify entire information of a vendor specific attribute profile.
//
// Request Body:
//	 - body *WSGVendorspecificattributeprofilePersist
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGVendorspecificattributeprofileService) UpdateRkszonesVendorSpecificAttributeProfilesById(ctx context.Context, body *WSGVendorspecificattributeprofilePersist, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateRkszonesVendorSpecificAttributeProfilesById, true)
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
