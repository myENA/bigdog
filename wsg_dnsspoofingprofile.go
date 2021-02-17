package bigdog

// API Version: v9_1

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"
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

func newWSGDNSSpoofingProfileDetailAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGDNSSpoofingProfileDetailAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGDNSSpoofingProfileDetailAPIResponse) Hydrate() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return nil
		}
		return r.err
	}
	data := new(WSGDNSSpoofingProfileDetail)
	if err := r.doHydrate(data); err != nil {
		return err
	}
	r.Data = data
	return nil
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

func newWSGDNSSpoofingProfileGetDnsSpoofingProfileListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGDNSSpoofingProfileGetDnsSpoofingProfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGDNSSpoofingProfileGetDnsSpoofingProfileListAPIResponse) Hydrate() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return nil
		}
		return r.err
	}
	data := new(WSGDNSSpoofingProfileGetDnsSpoofingProfileList)
	if err := r.doHydrate(data); err != nil {
		return err
	}
	r.Data = data
	return nil
}
func NewWSGDNSSpoofingProfileGetDnsSpoofingProfileList() *WSGDNSSpoofingProfileGetDnsSpoofingProfileList {
	m := new(WSGDNSSpoofingProfileGetDnsSpoofingProfileList)
	return m
}

// AddRkszonesDnsSpoofingProfilesByZoneId
//
// Use this API command to create DNS Spoofing profile.
//
// Operation ID: addRkszonesDnsSpoofingProfilesByZoneId
// Operation path: /rkszones/{zoneId}/dnsSpoofingProfiles
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGDNSSpoofingProfile
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGDNSSpoofingProfileService) AddRkszonesDnsSpoofingProfilesByZoneId(ctx context.Context, body *WSGDNSSpoofingProfile, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddRkszonesDnsSpoofingProfilesByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteRkszonesDnsSpoofingProfiles
//
// Use this API command to delete bulk DNS Spoofing profile.
//
// Operation ID: deleteRkszonesDnsSpoofingProfiles
// Operation path: /rkszones/dnsSpoofingProfiles
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGDNSSpoofingProfileService) DeleteRkszonesDnsSpoofingProfiles(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesDnsSpoofingProfiles, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesDnsSpoofingProfilesById
//
// Use this API command to delete DNS Spoofing profile.
//
// Operation ID: deleteRkszonesDnsSpoofingProfilesById
// Operation path: /rkszones/{zoneId}/dnsSpoofingProfiles/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDNSSpoofingProfileService) DeleteRkszonesDnsSpoofingProfilesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesDnsSpoofingProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindRkszonesDnsSpoofingProfilesById
//
// Use this API command to retrieve DNS Spoofing profile.
//
// Operation ID: findRkszonesDnsSpoofingProfilesById
// Operation path: /rkszones/{zoneId}/dnsSpoofingProfiles/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDNSSpoofingProfileService) FindRkszonesDnsSpoofingProfilesById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGDNSSpoofingProfileDetailAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGDNSSpoofingProfileDetailAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesDnsSpoofingProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDNSSpoofingProfileDetailAPIResponse), err
}

// FindRkszonesDnsSpoofingProfilesByZoneId
//
// Use this API command to retrieve a list of DNS Spoofing profile.
//
// Operation ID: findRkszonesDnsSpoofingProfilesByZoneId
// Operation path: /rkszones/{zoneId}/dnsSpoofingProfiles
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGDNSSpoofingProfileService) FindRkszonesDnsSpoofingProfilesByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGDNSSpoofingProfileGetDnsSpoofingProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGDNSSpoofingProfileGetDnsSpoofingProfileListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesDnsSpoofingProfilesByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDNSSpoofingProfileGetDnsSpoofingProfileListAPIResponse), err
}

// UpdateRkszonesDnsSpoofingProfilesById
//
// Use this API command to update DNS Spoofing profile.
//
// Operation ID: updateRkszonesDnsSpoofingProfilesById
// Operation path: /rkszones/{zoneId}/dnsSpoofingProfiles/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGDNSSpoofingProfile
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDNSSpoofingProfileService) UpdateRkszonesDnsSpoofingProfilesById(ctx context.Context, body *WSGDNSSpoofingProfile, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateRkszonesDnsSpoofingProfilesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
