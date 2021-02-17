package bigdog

// API Version: v9_1

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"
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

func newWSGVLANPoolingAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGVLANPoolingAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGVLANPoolingAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGVLANPooling)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
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

func newWSGVLANPoolingListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGVLANPoolingListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGVLANPoolingListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGVLANPoolingList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
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
// Use this API command to create new VLAN pooling.
//
// Operation ID: addVlanpoolings
// Operation path: /vlanpoolings
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGVLANPoolingCreateVlanPooling
func (s *WSGVLANPoolingService) AddVlanpoolings(ctx context.Context, body *WSGVLANPoolingCreateVlanPooling, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddVlanpoolings, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteVlanpoolings
//
// Use this API command to bulk delete VLAN pooling.
//
// Operation ID: deleteVlanpoolings
// Operation path: /vlanpoolings
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGVLANPoolingDeleteBulkVlanPooling
func (s *WSGVLANPoolingService) DeleteVlanpoolings(ctx context.Context, body *WSGVLANPoolingDeleteBulkVlanPooling, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteVlanpoolings, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteVlanpoolingsById
//
// Use this API command to delete VLAN pooling.
//
// Operation ID: deleteVlanpoolingsById
// Operation path: /vlanpoolings/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGVLANPoolingService) DeleteVlanpoolingsById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteVlanpoolingsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// FindServicesVlanPoolingByQueryCriteria
//
// Query Vlan Pooling Profiles with specified filters.
//
// Operation ID: findServicesVlanPoolingByQueryCriteria
// Operation path: /query/services/vlanPooling
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGVLANPoolingService) FindServicesVlanPoolingByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindServicesVlanPoolingByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*RawAPIResponse), err
}

// FindVlanpoolingsById
//
// Use this API command to retrieve VLAN pooling.
//
// Operation ID: findVlanpoolingsById
// Operation path: /vlanpoolings/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGVLANPoolingService) FindVlanpoolingsById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGVLANPoolingAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGVLANPoolingAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindVlanpoolingsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGVLANPoolingAPIResponse), err
}

// FindVlanpoolingsByQueryCriteria
//
// Use this API command to retrieve a list of VLAN poolings.
//
// Operation ID: findVlanpoolingsByQueryCriteria
// Operation path: /vlanpoolings/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGVLANPoolingService) FindVlanpoolingsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGVLANPoolingListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGVLANPoolingListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindVlanpoolingsByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGVLANPoolingListAPIResponse), err
}

// PartialUpdateVlanpoolingsById
//
// Use this API command to modify the configuration on VLAN pooling.
//
// Operation ID: partialUpdateVlanpoolingsById
// Operation path: /vlanpoolings/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGVLANPoolingModifyVlanPooling
//
// Required parameters:
// - id string
//		- required
func (s *WSGVLANPoolingService) PartialUpdateVlanpoolingsById(ctx context.Context, body *WSGVLANPoolingModifyVlanPooling, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateVlanpoolingsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}
