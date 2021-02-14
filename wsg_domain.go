package bigdog

// API Version: v9_1

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"
)

type WSGDomainService struct {
	apiClient *VSZClient
}

func NewWSGDomainService(c *VSZClient) *WSGDomainService {
	s := new(WSGDomainService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDomainService() *WSGDomainService {
	return NewWSGDomainService(ss.apiClient)
}

// WSGDomainCreateDomain
//
// Definition: domain_createDomain
type WSGDomainCreateDomain struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainType
	// domain type
	// Constraints:
	//    - default:'REGULAR'
	//    - oneof:[PARTNER,MVNO,REGULAR]
	DomainType *string `json:"domainType,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// ParentDomainId
	// parent domain id
	ParentDomainId *string `json:"parentDomainId,omitempty"`

	ZeroTouchStatus *bool `json:"zeroTouchStatus,omitempty"`
}

func NewWSGDomainCreateDomain() *WSGDomainCreateDomain {
	m := new(WSGDomainCreateDomain)
	return m
}

// WSGDomainConfiguration
//
// Definition: domain_domainConfiguration
type WSGDomainConfiguration struct {
	// AdministratorCount
	// # of Subdomains
	AdministratorCount *int `json:"administratorCount,omitempty"`

	// ApCount
	// # of Subdomains
	ApCount *int `json:"apCount,omitempty"`

	// CreateDatetime
	// Created by
	CreateDatetime *string `json:"createDatetime,omitempty"`

	// CreatedBy
	// Created by
	CreatedBy *string `json:"createdBy,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainType
	// domain type
	DomainType *string `json:"domainType,omitempty"`

	// Id
	// Identifier of the domain
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// ParentDomainId
	// Parent Domain Id
	ParentDomainId *string `json:"parentDomainId,omitempty"`

	// SubDomainCount
	// # of Subdomains
	SubDomainCount *int `json:"subDomainCount,omitempty"`

	ZeroTouchStatus *bool `json:"zeroTouchStatus,omitempty"`

	// ZoneCount
	// # of Zones
	ZoneCount *int `json:"zoneCount,omitempty"`
}

type WSGDomainConfigurationAPIResponse struct {
	*RawAPIResponse
	Data *WSGDomainConfiguration
}

func newWSGDomainConfigurationAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGDomainConfigurationAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGDomainConfigurationAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGDomainConfiguration)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGDomainConfiguration() *WSGDomainConfiguration {
	m := new(WSGDomainConfiguration)
	return m
}

// WSGDomainList
//
// Definition: domain_domainList
type WSGDomainList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGDomainConfiguration `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGDomainListAPIResponse struct {
	*RawAPIResponse
	Data *WSGDomainList
}

func newWSGDomainListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGDomainListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGDomainListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGDomainList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGDomainList() *WSGDomainList {
	m := new(WSGDomainList)
	return m
}

// WSGDomainModifyDomain
//
// Definition: domain_modifyDomain
type WSGDomainModifyDomain struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainType
	// domain type
	DomainType *string `json:"domainType,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// ZeroTouchStatus
	// Zero Touch enable/disable
	ZeroTouchStatus *bool `json:"zeroTouchStatus,omitempty"`
}

func NewWSGDomainModifyDomain() *WSGDomainModifyDomain {
	m := new(WSGDomainModifyDomain)
	return m
}

// AddDomains
//
// Use this API command to create new domain.
//
// Operation ID: addDomains
// Operation path: /domains
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGDomainCreateDomain
//
// Optional parameters:
// - parentDomainId string
//		- nullable
func (s *WSGDomainService) AddDomains(ctx context.Context, body *WSGDomainCreateDomain, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddDomains, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	if v, ok := optionalParams["parentDomainId"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("parentDomainId", v)
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteDomainsById
//
// Use this API command to delete domain.
//
// Operation ID: deleteDomainsById
// Operation path: /domains/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGDomainService) DeleteDomainsById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteDomainsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindDomains
//
// Use this API command to retrieve a list of domain under Administration Domain.
//
// Operation ID: findDomains
// Operation path: /domains
// Success code: 200 (OK)
//
// Optional parameters:
// - excludeRegularDomain string
//		- nullable
// - includeSelf string
//		- nullable
// - index string
//		- nullable
// - listSize string
//		- nullable
// - recursively string
//		- nullable
func (s *WSGDomainService) FindDomains(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGDomainListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGDomainListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindDomains, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["excludeRegularDomain"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("excludeRegularDomain", v)
	}
	if v, ok := optionalParams["includeSelf"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("includeSelf", v)
	}
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	if v, ok := optionalParams["recursively"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("recursively", v)
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDomainListAPIResponse), err
}

// FindDomainsById
//
// Use this API command to retrieve domain by specified Domain ID.
//
// Operation ID: findDomainsById
// Operation path: /domains/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
//
// Optional parameters:
// - recursively string
//		- nullable
func (s *WSGDomainService) FindDomainsById(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGDomainConfigurationAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGDomainConfigurationAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindDomainsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	if v, ok := optionalParams["recursively"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("recursively", v)
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDomainConfigurationAPIResponse), err
}

// FindDomainsByNameByDomainName
//
// Use this API command to retrieve a list of domain by specified Domain name.
//
// Operation ID: findDomainsByNameByDomainName
// Operation path: /domains/byName/{domainName}
// Success code: 200 (OK)
//
// Required parameters:
// - domainName string
//		- required
func (s *WSGDomainService) FindDomainsByNameByDomainName(ctx context.Context, domainName string, mutators ...RequestMutator) (*WSGDomainListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGDomainListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindDomainsByNameByDomainName, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("domainName", domainName)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDomainListAPIResponse), err
}

// FindDomainsSubdomainById
//
// Use this API command to retrieve a list of subdomain by specified Domain ID.
//
// Operation ID: findDomainsSubdomainById
// Operation path: /domains/{id}/subdomain
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
//
// Optional parameters:
// - excludeRegularDomain string
//		- nullable
// - includeSelf string
//		- nullable
// - index string
//		- nullable
// - listSize string
//		- nullable
// - recursively string
//		- nullable
func (s *WSGDomainService) FindDomainsSubdomainById(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGDomainListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGDomainListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindDomainsSubdomainById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	if v, ok := optionalParams["excludeRegularDomain"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("excludeRegularDomain", v)
	}
	if v, ok := optionalParams["includeSelf"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("includeSelf", v)
	}
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	if v, ok := optionalParams["recursively"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("recursively", v)
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDomainListAPIResponse), err
}

// PartialUpdateDomainsById
//
// Use this API command to modify the configuration of domain.
//
// Operation ID: partialUpdateDomainsById
// Operation path: /domains/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGDomainModifyDomain
//
// Required parameters:
// - id string
//		- required
func (s *WSGDomainService) PartialUpdateDomainsById(ctx context.Context, body *WSGDomainModifyDomain, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateDomainsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
