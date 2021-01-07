package bigdog

// API Version: v9_1

import (
	"context"
	"encoding/json"
	"net/http"
)

type WSGVLANPoolingService struct {
	apiClient *VSZClient
}

func NewWSGVLANPoolingService(c *VSZClient) *WSGVLANPoolingService {
	s := new(WSGVLANPoolingService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGVLANPoolingService() *WSGVLANPoolingService {
	return NewWSGVLANPoolingService(ss.apiClient)
}

// WSGVLANPoolingCreateVlanPooling
//
// Definition: vlanpooling_createVlanPooling
type WSGVLANPoolingCreateVlanPooling struct {
	// Algo
	// Algorithm of the VLAN pooling profile
	// Constraints:
	//    - required
	//    - default:'MAC_HASH'
	//    - oneof:[MAC_HASH]
	Algo *string `json:"algo"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// identifier of the domain
	DomainId *string `json:"domainId,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// Pool
	// VLANs of the VLAN pooling profile
	// Constraints:
	//    - required
	Pool *string `json:"pool"`
}

func NewWSGVLANPoolingCreateVlanPooling() *WSGVLANPoolingCreateVlanPooling {
	m := new(WSGVLANPoolingCreateVlanPooling)
	return m
}

// WSGVLANPoolingDeleteBulkVlanPooling
//
// Definition: vlanpooling_deleteBulkVlanPooling
type WSGVLANPoolingDeleteBulkVlanPooling struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGVLANPoolingDeleteBulkVlanPooling() *WSGVLANPoolingDeleteBulkVlanPooling {
	m := new(WSGVLANPoolingDeleteBulkVlanPooling)
	return m
}

// WSGVLANPoolingModifyVlanPooling
//
// Definition: vlanpooling_modifyVlanPooling
type WSGVLANPoolingModifyVlanPooling struct {
	// Algo
	// Algorithm of the VLAN pooling profile
	// Constraints:
	//    - oneof:[MAC_HASH]
	Algo *string `json:"algo,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// identifier of the domain
	DomainId *string `json:"domainId,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Pool
	// VLANs of the VLAN pooling profile
	Pool *string `json:"pool,omitempty"`
}

func NewWSGVLANPoolingModifyVlanPooling() *WSGVLANPoolingModifyVlanPooling {
	m := new(WSGVLANPoolingModifyVlanPooling)
	return m
}

// WSGVLANPooling
//
// Definition: vlanpooling_vlanPooling
type WSGVLANPooling struct {
	// Algo
	// Algorithm of the VLAN pooling profile
	// Constraints:
	//    - oneof:[MAC_HASH]
	Algo *string `json:"algo,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Identifier of the domain which the VLAN pooling profile belongs to
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of the VLAN pooling profile
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Pool
	// VLANs of the VLAN pooling profile
	Pool *string `json:"pool,omitempty"`
}

type WSGVLANPoolingAPIResponse struct {
	*RawAPIResponse
	Data *WSGVLANPooling
}

func newWSGVLANPoolingAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGVLANPoolingAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGVLANPoolingAPIResponse) Hydrate() error {
	r.Data = new(WSGVLANPooling)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGVLANPooling() *WSGVLANPooling {
	m := new(WSGVLANPooling)
	return m
}

// WSGVLANPoolingList
//
// Definition: vlanpooling_vlanPoolingList
type WSGVLANPoolingList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGVLANPoolingListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGVLANPoolingListAPIResponse struct {
	*RawAPIResponse
	Data *WSGVLANPoolingList
}

func newWSGVLANPoolingListAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGVLANPoolingListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGVLANPoolingListAPIResponse) Hydrate() error {
	r.Data = new(WSGVLANPoolingList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGVLANPoolingList() *WSGVLANPoolingList {
	m := new(WSGVLANPoolingList)
	return m
}

// WSGVLANPoolingListType
//
// Definition: vlanpooling_vlanPoolingListType
type WSGVLANPoolingListType struct {
	// Algo
	// Algorithm of the VLAN pooling profile
	// Constraints:
	//    - default:'MAC_HASH'
	//    - oneof:[MAC_HASH]
	Algo *string `json:"algo,omitempty"`

	// Description
	// Description of the service
	Description *string `json:"description,omitempty"`

	// DomainId
	// identifier of the domain
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of the service
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Pool
	// VLANs of the VLAN pooling profile
	Pool *string `json:"pool,omitempty"`
}

func NewWSGVLANPoolingListType() *WSGVLANPoolingListType {
	m := new(WSGVLANPoolingListType)
	return m
}

// AddVlanpoolings
//
// Operation ID: addVlanpoolings
//
// Use this API command to create new VLAN pooling.
//
// Request Body:
//	 - body *WSGVLANPoolingCreateVlanPooling
func (s *WSGVLANPoolingService) AddVlanpoolings(ctx context.Context, body *WSGVLANPoolingCreateVlanPooling, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddVlanpoolings, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteVlanpoolings
//
// Operation ID: deleteVlanpoolings
//
// Use this API command to bulk delete VLAN pooling.
//
// Request Body:
//	 - body *WSGVLANPoolingDeleteBulkVlanPooling
func (s *WSGVLANPoolingService) DeleteVlanpoolings(ctx context.Context, body *WSGVLANPoolingDeleteBulkVlanPooling, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteVlanpoolings, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteVlanpoolingsById
//
// Operation ID: deleteVlanpoolingsById
//
// Use this API command to delete VLAN pooling.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGVLANPoolingService) DeleteVlanpoolingsById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteVlanpoolingsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindServicesVlanPoolingByQueryCriteria
//
// Operation ID: findServicesVlanPoolingByQueryCriteria
//
// Query Vlan Pooling Profiles with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGVLANPoolingService) FindServicesVlanPoolingByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindServicesVlanPoolingByQueryCriteria, true)
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

// FindVlanpoolingsById
//
// Operation ID: findVlanpoolingsById
//
// Use this API command to retrieve VLAN pooling.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGVLANPoolingService) FindVlanpoolingsById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGVLANPoolingAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGVLANPoolingAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGVLANPoolingAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindVlanpoolingsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGVLANPoolingAPIResponse), err
}

// FindVlanpoolingsByQueryCriteria
//
// Operation ID: findVlanpoolingsByQueryCriteria
//
// Use this API command to retrieve a list of VLAN poolings.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGVLANPoolingService) FindVlanpoolingsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGVLANPoolingListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGVLANPoolingListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGVLANPoolingListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindVlanpoolingsByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGVLANPoolingListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGVLANPoolingListAPIResponse), err
}

// PartialUpdateVlanpoolingsById
//
// Operation ID: partialUpdateVlanpoolingsById
//
// Use this API command to modify the configuration on VLAN pooling.
//
// Request Body:
//	 - body *WSGVLANPoolingModifyVlanPooling
//
// Required Parameters:
// - id string
//		- required
func (s *WSGVLANPoolingService) PartialUpdateVlanpoolingsById(ctx context.Context, body *WSGVLANPoolingModifyVlanPooling, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateVlanpoolingsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
