package bigdog

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGDNSSpoofingProfileService struct {
	apiClient *VSZClient
}

func NewWSGDNSSpoofingProfileService(c *VSZClient) *WSGDNSSpoofingProfileService {
	s := new(WSGDNSSpoofingProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDNSSpoofingProfileService() *WSGDNSSpoofingProfileService {
	return NewWSGDNSSpoofingProfileService(ss.apiClient)
}

type WSGDNSSpoofingProfile struct {
	// Description
	// DNS Spoofing Profile's description
	Description *string `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName2to64 `json:"name"`

	// Rules
	// DNS Spoofing Profile's rules (At least one rule)
	Rules []*WSGDNSSpoofingProfileDnsSpoofingRule `json:"rules,omitempty"`
}

func NewWSGDNSSpoofingProfile() *WSGDNSSpoofingProfile {
	m := new(WSGDNSSpoofingProfile)
	return m
}

type WSGDNSSpoofingProfileDetail struct {
	// CreateDateTime
	// The user who create the DNS Spoofing Profile
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorUsername
	// The time when the user create the DNS Spoofing Profile
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// DNS Spoofing Profile's description
	Description *string `json:"description,omitempty"`

	// Id
	// DNS Spoofing Profile's id
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// The time when the user modify the DNS Spoofing Profile
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierUsername
	// The user who modify the DNS Spoofing Profile
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// DNS Spoofing Profile's name
	Name *string `json:"name,omitempty"`

	// Rules
	// DNS Spoofing Profile's rules (At least one rule)
	Rules []*WSGDNSSpoofingProfileDnsSpoofingRule `json:"rules,omitempty"`

	// ZoneId
	// The zone which DNS Spoofing Profile belong to
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGDNSSpoofingProfileDetail() *WSGDNSSpoofingProfileDetail {
	m := new(WSGDNSSpoofingProfileDetail)
	return m
}

type WSGDNSSpoofingProfileSummary struct {
	// Id
	// DNS Spoofing Profile's id
	Id *string `json:"id,omitempty"`

	// Name
	// DNS Spoofing Profile's name
	Name *string `json:"name,omitempty"`
}

func NewWSGDNSSpoofingProfileSummary() *WSGDNSSpoofingProfileSummary {
	m := new(WSGDNSSpoofingProfileSummary)
	return m
}

type WSGDNSSpoofingProfileDnsSpoofingRule struct {
	// DomainName
	// Rule's Domain Name
	// Constraints:
	//    - required
	DomainName *string `json:"domainName"`

	// IpList
	// Rule's Ip addresses
	// Constraints:
	//    - required
	IpList []string `json:"ipList"`
}

func NewWSGDNSSpoofingProfileDnsSpoofingRule() *WSGDNSSpoofingProfileDnsSpoofingRule {
	m := new(WSGDNSSpoofingProfileDnsSpoofingRule)
	return m
}

type WSGDNSSpoofingProfileGetDnsSpoofingProfileList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGDNSSpoofingProfileSummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGDNSSpoofingProfileGetDnsSpoofingProfileList() *WSGDNSSpoofingProfileGetDnsSpoofingProfileList {
	m := new(WSGDNSSpoofingProfileGetDnsSpoofingProfileList)
	return m
}

// AddRkszonesDnsSpoofingProfilesByZoneId
//
// Use this API command to create DNS Spoofing profile.
//
// Request Body:
//	 - body *WSGDNSSpoofingProfile
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDNSSpoofingProfileService) AddRkszonesDnsSpoofingProfilesByZoneId(ctx context.Context, body *WSGDNSSpoofingProfile, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesDnsSpoofingProfilesByZoneId, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteRkszonesDnsSpoofingProfiles
//
// Use this API command to delete bulk DNS Spoofing profile.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGDNSSpoofingProfileService) DeleteRkszonesDnsSpoofingProfiles(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesDnsSpoofingProfiles, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesDnsSpoofingProfilesById
//
// Use this API command to delete DNS Spoofing profile.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDNSSpoofingProfileService) DeleteRkszonesDnsSpoofingProfilesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesDnsSpoofingProfilesById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindRkszonesDnsSpoofingProfilesById
//
// Use this API command to retrieve DNS Spoofing profile.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDNSSpoofingProfileService) FindRkszonesDnsSpoofingProfilesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGDNSSpoofingProfileDetail, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDNSSpoofingProfileDetail
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesDnsSpoofingProfilesById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDNSSpoofingProfileDetail()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesDnsSpoofingProfilesByZoneId
//
// Use this API command to retrieve a list of DNS Spoofing profile.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDNSSpoofingProfileService) FindRkszonesDnsSpoofingProfilesByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGDNSSpoofingProfileGetDnsSpoofingProfileList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDNSSpoofingProfileGetDnsSpoofingProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesDnsSpoofingProfilesByZoneId, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDNSSpoofingProfileGetDnsSpoofingProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateRkszonesDnsSpoofingProfilesById
//
// Use this API command to update DNS Spoofing profile.
//
// Request Body:
//	 - body *WSGDNSSpoofingProfile
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDNSSpoofingProfileService) UpdateRkszonesDnsSpoofingProfilesById(ctx context.Context, body *WSGDNSSpoofingProfile, id string, zoneId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateRkszonesDnsSpoofingProfilesById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
