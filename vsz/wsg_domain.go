package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGDomainService struct {
	apiClient *APIClient
}

func NewWSGDomainService(c *APIClient) *WSGDomainService {
	s := new(WSGDomainService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDomainService() *WSGDomainService {
	return NewWSGDomainService(ss.apiClient)
}

type WSGDomainCreateDomain struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainType
	// domain type
	// Constraints:
	//    - default:'REGULAR'
	//    - oneof:[PARTNER,MVNO,REGULAR]
	DomainType *string `json:"domainType,omitempty" validate:"oneof=PARTNER MVNO REGULAR"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	// ParentDomainId
	// parent domain id
	ParentDomainId *string `json:"parentDomainId,omitempty"`

	ZeroTouchStatus *bool `json:"zeroTouchStatus,omitempty"`
}

func NewWSGDomainCreateDomain() *WSGDomainCreateDomain {
	m := new(WSGDomainCreateDomain)
	return m
}

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

func NewWSGDomainConfiguration() *WSGDomainConfiguration {
	m := new(WSGDomainConfiguration)
	return m
}

type WSGDomainList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGDomainConfiguration `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGDomainList() *WSGDomainList {
	m := new(WSGDomainList)
	return m
}

type WSGDomainModifyDomain struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainType
	// domain type
	DomainType *string `json:"domainType,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// ParentDomainId
	// parent domain id
	ParentDomainId *string `json:"parentDomainId,omitempty"`

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
// Request Body:
//	 - body *WSGDomainCreateDomain
//
// Optional Parameters:
// - parentDomainId string
//		- nullable
func (s *WSGDomainService) AddDomains(ctx context.Context, body *WSGDomainCreateDomain, optionalParams map[string][]string) (*WSGCommonCreateResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddDomains, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	if v, ok := optionalParams["parentDomainId"]; ok {
		req.AddQueryParameter("parentDomainId", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	return resp, handleResponse(req, http.StatusCreated, httpResp, &resp, err)
}

// DeleteDomainsById
//
// Use this API command to delete domain.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDomainService) DeleteDomainsById(ctx context.Context, id string) error {
	var (
		req      *APIRequest
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteDomainsById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	return handleResponse(req, http.StatusNoContent, httpResp, nil, err)
}

// FindDomains
//
// Use this API command to retrieve a list of domain under Administration Domain.
//
// Optional Parameters:
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
func (s *WSGDomainService) FindDomains(ctx context.Context, optionalParams map[string][]string) (*WSGDomainList, error) {
	var (
		req      *APIRequest
		resp     *WSGDomainList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDomains, true)
	if v, ok := optionalParams["excludeRegularDomain"]; ok {
		req.AddQueryParameter("excludeRegularDomain", v)
	}
	if v, ok := optionalParams["includeSelf"]; ok {
		req.AddQueryParameter("includeSelf", v)
	}
	if v, ok := optionalParams["index"]; ok {
		req.AddQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok {
		req.AddQueryParameter("listSize", v)
	}
	if v, ok := optionalParams["recursively"]; ok {
		req.AddQueryParameter("recursively", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDomainList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindDomainsById
//
// Use this API command to retrieve domain by specified Domain ID.
//
// Required Parameters:
// - id string
//		- required
//
// Optional Parameters:
// - recursively string
//		- nullable
func (s *WSGDomainService) FindDomainsById(ctx context.Context, id string, optionalParams map[string][]string) (*WSGDomainConfiguration, error) {
	var (
		req      *APIRequest
		resp     *WSGDomainConfiguration
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDomainsById, true)
	req.SetPathParameter("id", id)
	if v, ok := optionalParams["recursively"]; ok {
		req.AddQueryParameter("recursively", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDomainConfiguration()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindDomainsByNameByDomainName
//
// Use this API command to retrieve a list of domain by specified Domain name.
//
// Required Parameters:
// - domainName string
//		- required
func (s *WSGDomainService) FindDomainsByNameByDomainName(ctx context.Context, domainName string) (*WSGDomainList, error) {
	var (
		req      *APIRequest
		resp     *WSGDomainList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, domainName, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDomainsByNameByDomainName, true)
	req.SetPathParameter("domainName", domainName)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDomainList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindDomainsSubdomainById
//
// Use this API command to retrieve a list of subdomain by specified Domain ID.
//
// Required Parameters:
// - id string
//		- required
//
// Optional Parameters:
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
func (s *WSGDomainService) FindDomainsSubdomainById(ctx context.Context, id string, optionalParams map[string][]string) (*WSGDomainList, error) {
	var (
		req      *APIRequest
		resp     *WSGDomainList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDomainsSubdomainById, true)
	req.SetPathParameter("id", id)
	if v, ok := optionalParams["excludeRegularDomain"]; ok {
		req.AddQueryParameter("excludeRegularDomain", v)
	}
	if v, ok := optionalParams["includeSelf"]; ok {
		req.AddQueryParameter("includeSelf", v)
	}
	if v, ok := optionalParams["index"]; ok {
		req.AddQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok {
		req.AddQueryParameter("listSize", v)
	}
	if v, ok := optionalParams["recursively"]; ok {
		req.AddQueryParameter("recursively", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDomainList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// PartialUpdateDomainsById
//
// Use this API command to modify the configuration of domain.
//
// Request Body:
//	 - body *WSGDomainModifyDomain
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDomainService) PartialUpdateDomainsById(ctx context.Context, body *WSGDomainModifyDomain, id string) error {
	var (
		req      *APIRequest
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateDomainsById, true)
	if err = req.SetBody(body); err != nil {
		return err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	return handleResponse(req, http.StatusNoContent, httpResp, nil, err)
}
