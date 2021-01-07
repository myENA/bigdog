package bigdog

// API Version: v9_1

import (
	"context"
	"encoding/json"
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

// WSGDNSSpoofingProfile
//
// Definition: dnsSpoofingProfile_dnsSpoofingProfile
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

// WSGDNSSpoofingProfileDetail
//
// Definition: dnsSpoofingProfile_dnsSpoofingProfileDetail
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

type WSGDNSSpoofingProfileDetailAPIResponse struct {
	*RawAPIResponse
	Data *WSGDNSSpoofingProfileDetail
}

func newWSGDNSSpoofingProfileDetailAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGDNSSpoofingProfileDetailAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGDNSSpoofingProfileDetailAPIResponse) Hydrate() error {
	r.Data = new(WSGDNSSpoofingProfileDetail)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGDNSSpoofingProfileDetail() *WSGDNSSpoofingProfileDetail {
	m := new(WSGDNSSpoofingProfileDetail)
	return m
}

// WSGDNSSpoofingProfileSummary
//
// Definition: dnsSpoofingProfile_dnsSpoofingProfileSummary
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

// WSGDNSSpoofingProfileDnsSpoofingRule
//
// Definition: dnsSpoofingProfile_dnsSpoofingRule
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

// WSGDNSSpoofingProfileGetDnsSpoofingProfileList
//
// Definition: dnsSpoofingProfile_getDnsSpoofingProfileList
type WSGDNSSpoofingProfileGetDnsSpoofingProfileList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGDNSSpoofingProfileSummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGDNSSpoofingProfileGetDnsSpoofingProfileListAPIResponse struct {
	*RawAPIResponse
	Data *WSGDNSSpoofingProfileGetDnsSpoofingProfileList
}

func newWSGDNSSpoofingProfileGetDnsSpoofingProfileListAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGDNSSpoofingProfileGetDnsSpoofingProfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGDNSSpoofingProfileGetDnsSpoofingProfileListAPIResponse) Hydrate() error {
	r.Data = new(WSGDNSSpoofingProfileGetDnsSpoofingProfileList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGDNSSpoofingProfileGetDnsSpoofingProfileList() *WSGDNSSpoofingProfileGetDnsSpoofingProfileList {
	m := new(WSGDNSSpoofingProfileGetDnsSpoofingProfileList)
	return m
}

// AddRkszonesDnsSpoofingProfilesByZoneId
//
// Operation ID: addRkszonesDnsSpoofingProfilesByZoneId
//
// Use this API command to create DNS Spoofing profile.
//
// Request Body:
//	 - body *WSGDNSSpoofingProfile
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDNSSpoofingProfileService) AddRkszonesDnsSpoofingProfilesByZoneId(ctx context.Context, body *WSGDNSSpoofingProfile, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesDnsSpoofingProfilesByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonCreateResult()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteRkszonesDnsSpoofingProfiles
//
// Operation ID: deleteRkszonesDnsSpoofingProfiles
//
// Use this API command to delete bulk DNS Spoofing profile.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGDNSSpoofingProfileService) DeleteRkszonesDnsSpoofingProfiles(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesDnsSpoofingProfiles, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesDnsSpoofingProfilesById
//
// Operation ID: deleteRkszonesDnsSpoofingProfilesById
//
// Use this API command to delete DNS Spoofing profile.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDNSSpoofingProfileService) DeleteRkszonesDnsSpoofingProfilesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesDnsSpoofingProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindRkszonesDnsSpoofingProfilesById
//
// Operation ID: findRkszonesDnsSpoofingProfilesById
//
// Use this API command to retrieve DNS Spoofing profile.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDNSSpoofingProfileService) FindRkszonesDnsSpoofingProfilesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGDNSSpoofingProfileDetailAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesDnsSpoofingProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDNSSpoofingProfileDetail()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesDnsSpoofingProfilesByZoneId
//
// Operation ID: findRkszonesDnsSpoofingProfilesByZoneId
//
// Use this API command to retrieve a list of DNS Spoofing profile.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDNSSpoofingProfileService) FindRkszonesDnsSpoofingProfilesByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGDNSSpoofingProfileGetDnsSpoofingProfileListAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesDnsSpoofingProfilesByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDNSSpoofingProfileGetDnsSpoofingProfileList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateRkszonesDnsSpoofingProfilesById
//
// Operation ID: updateRkszonesDnsSpoofingProfilesById
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
func (s *WSGDNSSpoofingProfileService) UpdateRkszonesDnsSpoofingProfilesById(ctx context.Context, body *WSGDNSSpoofingProfile, id string, zoneId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateRkszonesDnsSpoofingProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
