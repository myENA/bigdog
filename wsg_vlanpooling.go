package bigdog

// API Version: v9_0

import (
	"context"
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
func (s *WSGVLANPoolingService) AddVlanpoolings(ctx context.Context, body *WSGVLANPoolingCreateVlanPooling, mutators ...RequestMutator) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddVlanpoolings, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteVlanpoolings
//
// Operation ID: deleteVlanpoolings
//
// Use this API command to bulk delete VLAN pooling.
//
// Request Body:
//	 - body *WSGVLANPoolingDeleteBulkVlanPooling
func (s *WSGVLANPoolingService) DeleteVlanpoolings(ctx context.Context, body *WSGVLANPoolingDeleteBulkVlanPooling, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteVlanpoolings, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
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
func (s *WSGVLANPoolingService) DeleteVlanpoolingsById(ctx context.Context, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteVlanpoolingsById, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindServicesVlanPoolingByQueryCriteria
//
// Operation ID: findServicesVlanPoolingByQueryCriteria
//
// Query Vlan Pooling Profiles with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGVLANPoolingService) FindServicesVlanPoolingByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesVlanPoolingByQueryCriteria, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *WSGVLANPoolingService) FindVlanpoolingsById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGVLANPooling, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGVLANPooling
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindVlanpoolingsById, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGVLANPooling()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindVlanpoolingsByQueryCriteria
//
// Operation ID: findVlanpoolingsByQueryCriteria
//
// Use this API command to retrieve a list of VLAN poolings.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGVLANPoolingService) FindVlanpoolingsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGVLANPoolingList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGVLANPoolingList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindVlanpoolingsByQueryCriteria, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGVLANPoolingList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *WSGVLANPoolingService) PartialUpdateVlanpoolingsById(ctx context.Context, body *WSGVLANPoolingModifyVlanPooling, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateVlanpoolingsById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
