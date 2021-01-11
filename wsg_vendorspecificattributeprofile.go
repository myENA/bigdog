package bigdog

// API Version: v9_1

import (
	"context"
	"errors"
	"io"
	"net/http"
)

type WSGVendorSpecificAttributeProfileService struct {
	apiClient *VSZClient
}

func NewWSGVendorSpecificAttributeProfileService(c *VSZClient) *WSGVendorSpecificAttributeProfileService {
	s := new(WSGVendorSpecificAttributeProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGVendorSpecificAttributeProfileService() *WSGVendorSpecificAttributeProfileService {
	return NewWSGVendorSpecificAttributeProfileService(ss.apiClient)
}

// WSGVendorSpecificAttributeProfileCreateResult
//
// Definition: vendorSpecificAttributeProfile_createResult
type WSGVendorSpecificAttributeProfileCreateResult struct {
	// Id
	// Identifier of the vendor specific attribute profile
	Id *string `json:"id,omitempty"`
}

type WSGVendorSpecificAttributeProfileCreateResultAPIResponse struct {
	*RawAPIResponse
	Data *WSGVendorSpecificAttributeProfileCreateResult
}

func newWSGVendorSpecificAttributeProfileCreateResultAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGVendorSpecificAttributeProfileCreateResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGVendorSpecificAttributeProfileCreateResultAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGVendorSpecificAttributeProfileCreateResult)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGVendorSpecificAttributeProfileCreateResult() *WSGVendorSpecificAttributeProfileCreateResult {
	m := new(WSGVendorSpecificAttributeProfileCreateResult)
	return m
}

// WSGVendorSpecificAttributeProfileDeleteBulk
//
// Definition: vendorSpecificAttributeProfile_deleteBulk
type WSGVendorSpecificAttributeProfileDeleteBulk struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGVendorSpecificAttributeProfileDeleteBulk() *WSGVendorSpecificAttributeProfileDeleteBulk {
	m := new(WSGVendorSpecificAttributeProfileDeleteBulk)
	return m
}

// WSGVendorSpecificAttributeProfileGet
//
// Definition: vendorSpecificAttributeProfile_get
type WSGVendorSpecificAttributeProfileGet struct {
	// Attributes
	// Vendor specific attribute list for Radius protocol
	Attributes []*WSGVendorSpecificAttributeProfileVendorSpecificAttribute `json:"attributes,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of the vendor specific attribute profile
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// ZoneId
	// Zone Id
	ZoneId *string `json:"zoneId,omitempty"`
}

type WSGVendorSpecificAttributeProfileGetAPIResponse struct {
	*RawAPIResponse
	Data *WSGVendorSpecificAttributeProfileGet
}

func newWSGVendorSpecificAttributeProfileGetAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGVendorSpecificAttributeProfileGetAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGVendorSpecificAttributeProfileGetAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGVendorSpecificAttributeProfileGet)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGVendorSpecificAttributeProfileGet() *WSGVendorSpecificAttributeProfileGet {
	m := new(WSGVendorSpecificAttributeProfileGet)
	return m
}

// WSGVendorSpecificAttributeProfileList
//
// Definition: vendorSpecificAttributeProfile_list
type WSGVendorSpecificAttributeProfileList struct {
	// FirstIndex
	// Index of the first profile returned out of the profile list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more profiles after the list that is currently displayed
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Information list of the vendor specific attribute profile
	List []*WSGVendorSpecificAttributeProfileListType `json:"list,omitempty"`

	// TotalCount
	// Total number of the profiles
	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGVendorSpecificAttributeProfileListAPIResponse struct {
	*RawAPIResponse
	Data *WSGVendorSpecificAttributeProfileList
}

func newWSGVendorSpecificAttributeProfileListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGVendorSpecificAttributeProfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGVendorSpecificAttributeProfileListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGVendorSpecificAttributeProfileList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGVendorSpecificAttributeProfileList() *WSGVendorSpecificAttributeProfileList {
	m := new(WSGVendorSpecificAttributeProfileList)
	return m
}

// WSGVendorSpecificAttributeProfileListType
//
// Definition: vendorSpecificAttributeProfile_listType
type WSGVendorSpecificAttributeProfileListType struct {
	// Id
	// Identifier of the vendor specific attribute profile
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGVendorSpecificAttributeProfileListType() *WSGVendorSpecificAttributeProfileListType {
	m := new(WSGVendorSpecificAttributeProfileListType)
	return m
}

// WSGVendorSpecificAttributeProfilePersist
//
// Definition: vendorSpecificAttributeProfile_persist
type WSGVendorSpecificAttributeProfilePersist struct {
	// Attributes
	// Vendor specific attribute list for Radius protocol
	// Constraints:
	//    - required
	Attributes []*WSGVendorSpecificAttributeProfileVendorSpecificAttribute `json:"attributes"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`
}

func NewWSGVendorSpecificAttributeProfilePersist() *WSGVendorSpecificAttributeProfilePersist {
	m := new(WSGVendorSpecificAttributeProfilePersist)
	return m
}

// WSGVendorSpecificAttributeProfileQueryCriteriaResult
//
// Definition: vendorSpecificAttributeProfile_queryCriteriaResult
type WSGVendorSpecificAttributeProfileQueryCriteriaResult struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGVendorSpecificAttributeProfileGet `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGVendorSpecificAttributeProfileQueryCriteriaResultAPIResponse struct {
	*RawAPIResponse
	Data *WSGVendorSpecificAttributeProfileQueryCriteriaResult
}

func newWSGVendorSpecificAttributeProfileQueryCriteriaResultAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGVendorSpecificAttributeProfileQueryCriteriaResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGVendorSpecificAttributeProfileQueryCriteriaResultAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGVendorSpecificAttributeProfileQueryCriteriaResult)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGVendorSpecificAttributeProfileQueryCriteriaResult() *WSGVendorSpecificAttributeProfileQueryCriteriaResult {
	m := new(WSGVendorSpecificAttributeProfileQueryCriteriaResult)
	return m
}

// WSGVendorSpecificAttributeProfileVendorSpecificAttribute
//
// Definition: vendorSpecificAttributeProfile_vendorSpecificAttribute
type WSGVendorSpecificAttributeProfileVendorSpecificAttribute struct {
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

func NewWSGVendorSpecificAttributeProfileVendorSpecificAttribute() *WSGVendorSpecificAttributeProfileVendorSpecificAttribute {
	m := new(WSGVendorSpecificAttributeProfileVendorSpecificAttribute)
	return m
}

// AddRkszonesVendorSpecificAttributeProfilesByZoneId
//
// Create a vendor specific attribute profile.
//
// Operation ID: addRkszonesVendorSpecificAttributeProfilesByZoneId
// Operation path: /rkszones/{zoneId}/vendorSpecificAttributeProfiles
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGVendorSpecificAttributeProfilePersist
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGVendorSpecificAttributeProfileService) AddRkszonesVendorSpecificAttributeProfilesByZoneId(ctx context.Context, body *WSGVendorSpecificAttributeProfilePersist, zoneId string, mutators ...RequestMutator) (*WSGVendorSpecificAttributeProfileCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGVendorSpecificAttributeProfileCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGVendorSpecificAttributeProfileCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesVendorSpecificAttributeProfilesByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGVendorSpecificAttributeProfileCreateResultAPIResponse), err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGVendorSpecificAttributeProfileCreateResultAPIResponse), err
}

// DeleteRkszonesVendorSpecificAttributeProfilesById
//
// Use this API command to delete a vendor specific attribute profile by ID.
//
// Operation ID: deleteRkszonesVendorSpecificAttributeProfilesById
// Operation path: /rkszones/{zoneId}/vendorSpecificAttributeProfiles/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGVendorSpecificAttributeProfileService) DeleteRkszonesVendorSpecificAttributeProfilesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesVendorSpecificAttributeProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesVendorSpecificAttributeProfilesByZoneId
//
// Use this API command to delete a list of vendor specific attribute profile.
//
// Operation ID: deleteRkszonesVendorSpecificAttributeProfilesByZoneId
// Operation path: /rkszones/{zoneId}/vendorSpecificAttributeProfiles
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGVendorSpecificAttributeProfileDeleteBulk
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGVendorSpecificAttributeProfileService) DeleteRkszonesVendorSpecificAttributeProfilesByZoneId(ctx context.Context, body *WSGVendorSpecificAttributeProfileDeleteBulk, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesVendorSpecificAttributeProfilesByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindRkszonesVendorSpecificAttributeProfilesById
//
// Get a vendor specific attribute profile by ID.
//
// Operation ID: findRkszonesVendorSpecificAttributeProfilesById
// Operation path: /rkszones/{zoneId}/vendorSpecificAttributeProfiles/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGVendorSpecificAttributeProfileService) FindRkszonesVendorSpecificAttributeProfilesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGVendorSpecificAttributeProfileGetAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGVendorSpecificAttributeProfileGetAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGVendorSpecificAttributeProfileGetAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesVendorSpecificAttributeProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGVendorSpecificAttributeProfileGetAPIResponse), err
}

// FindRkszonesVendorSpecificAttributeProfilesByQueryCriteria
//
// Use this API command to retrieve a list of vendor specific attribute profile by query criteria.
//
// Operation ID: findRkszonesVendorSpecificAttributeProfilesByQueryCriteria
// Operation path: /rkszones/vendorSpecificAttributeProfiles/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGVendorSpecificAttributeProfileService) FindRkszonesVendorSpecificAttributeProfilesByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGVendorSpecificAttributeProfileQueryCriteriaResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGVendorSpecificAttributeProfileQueryCriteriaResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGVendorSpecificAttributeProfileQueryCriteriaResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindRkszonesVendorSpecificAttributeProfilesByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGVendorSpecificAttributeProfileQueryCriteriaResultAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGVendorSpecificAttributeProfileQueryCriteriaResultAPIResponse), err
}

// FindRkszonesVendorSpecificAttributeProfilesByZoneId
//
// Get a ID list of vendor specific attribute profile in this Zone.
//
// Operation ID: findRkszonesVendorSpecificAttributeProfilesByZoneId
// Operation path: /rkszones/{zoneId}/vendorSpecificAttributeProfiles
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGVendorSpecificAttributeProfileService) FindRkszonesVendorSpecificAttributeProfilesByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGVendorSpecificAttributeProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGVendorSpecificAttributeProfileListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGVendorSpecificAttributeProfileListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesVendorSpecificAttributeProfilesByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGVendorSpecificAttributeProfileListAPIResponse), err
}

// UpdateRkszonesVendorSpecificAttributeProfilesById
//
// Use this API command to modify entire information of a vendor specific attribute profile.
//
// Operation ID: updateRkszonesVendorSpecificAttributeProfilesById
// Operation path: /rkszones/{zoneId}/vendorSpecificAttributeProfiles/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGVendorSpecificAttributeProfilePersist
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGVendorSpecificAttributeProfileService) UpdateRkszonesVendorSpecificAttributeProfilesById(ctx context.Context, body *WSGVendorSpecificAttributeProfilePersist, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateRkszonesVendorSpecificAttributeProfilesById, true)
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
