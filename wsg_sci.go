package bigdog

// API Version: v9_1

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type WSGSCIService struct {
	apiClient *VSZClient
}

func NewWSGSCIService(c *VSZClient) *WSGSCIService {
	s := new(WSGSCIService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSCIService() *WSGSCIService {
	return NewWSGSCIService(ss.apiClient)
}

// WSGSCICreateSciProfile
//
// Definition: sci_createSciProfile
type WSGSCICreateSciProfile struct {
	// SciPassword
	// SCI password of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciPassword *string `json:"sciPassword"`

	// SciProfile
	// Profile name of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciProfile *string `json:"sciProfile"`

	// SciServerHost
	// SCI server host of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciServerHost *string `json:"sciServerHost"`

	// SciServerPort
	// SCI server port of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciServerPort *string `json:"sciServerPort"`

	// SciSystemId
	// SCI system UUID of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciSystemId *string `json:"sciSystemId"`

	// SciUser
	// SCI user name of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciUser *string `json:"sciUser"`
}

func NewWSGSCICreateSciProfile() *WSGSCICreateSciProfile {
	m := new(WSGSCICreateSciProfile)
	return m
}

// WSGSCIDeleteSciProfile
//
// Definition: sci_deleteSciProfile
type WSGSCIDeleteSciProfile struct {
	// Id
	// UUID of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	Id *string `json:"id"`
}

func NewWSGSCIDeleteSciProfile() *WSGSCIDeleteSciProfile {
	m := new(WSGSCIDeleteSciProfile)
	return m
}

// WSGSCIDeleteSciProfileList
//
// Definition: sci_deleteSciProfileList
type WSGSCIDeleteSciProfileList struct {
	List []*WSGSCIDeleteSciProfile `json:"list,omitempty"`
}

func NewWSGSCIDeleteSciProfileList() *WSGSCIDeleteSciProfileList {
	m := new(WSGSCIDeleteSciProfileList)
	return m
}

// WSGSCIModifyEventCode
//
// Definition: sci_modifyEventCode
type WSGSCIModifyEventCode struct {
	// SciAcceptedEventCodes
	// Constraints:
	//    - required
	SciAcceptedEventCodes []int `json:"sciAcceptedEventCodes"`
}

func NewWSGSCIModifyEventCode() *WSGSCIModifyEventCode {
	m := new(WSGSCIModifyEventCode)
	return m
}

// WSGSCIModifySciEnabled
//
// Definition: sci_modifySciEnabled
type WSGSCIModifySciEnabled struct {
	// SciEnabled
	// Is SZ/SCI interface enabled or disabled
	// Constraints:
	//    - required
	SciEnabled *bool `json:"sciEnabled"`
}

func NewWSGSCIModifySciEnabled() *WSGSCIModifySciEnabled {
	m := new(WSGSCIModifySciEnabled)
	return m
}

// WSGSCIModifySciProfile
//
// Definition: sci_modifySciProfile
type WSGSCIModifySciProfile struct {
	// Id
	// UUID of the SCI profile for SZ/SCI interface
	Id *string `json:"id,omitempty"`

	// SciPassword
	// SCI password of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciPassword *string `json:"sciPassword"`

	// SciProfile
	// Profile name of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciProfile *string `json:"sciProfile"`

	// SciServerHost
	// SCI server host of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciServerHost *string `json:"sciServerHost"`

	// SciServerPort
	// SCI server port of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciServerPort *string `json:"sciServerPort"`

	// SciSystemId
	// SCI system UUID of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciSystemId *string `json:"sciSystemId"`

	// SciUser
	// SCI user name of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciUser *string `json:"sciUser"`
}

func NewWSGSCIModifySciProfile() *WSGSCIModifySciProfile {
	m := new(WSGSCIModifySciProfile)
	return m
}

// WSGSCIEventCode
//
// Definition: sci_sciEventCode
type WSGSCIEventCode struct {
	// FirstIndex
	// Index of the first event code returned from the complete event code set
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more SCI accepted event codes after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSCIEventCodeListType `json:"list,omitempty"`

	// TotalCount
	// Total SCI accepted event code count
	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGSCIEventCodeAPIResponse struct {
	*RawAPIResponse
	Data *WSGSCIEventCode
}

func newWSGSCIEventCodeAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSCIEventCodeAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSCIEventCodeAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGSCIEventCode)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGSCIEventCode() *WSGSCIEventCode {
	m := new(WSGSCIEventCode)
	return m
}

// WSGSCIEventCodeListType
//
// Definition: sci_sciEventCodeListType
type WSGSCIEventCodeListType struct {
	// Code
	// SCI accepted event code
	Code *int `json:"code,omitempty"`

	// Type
	// SCI accepted event type
	Type *string `json:"type,omitempty"`
}

func NewWSGSCIEventCodeListType() *WSGSCIEventCodeListType {
	m := new(WSGSCIEventCodeListType)
	return m
}

// WSGSCIProfile
//
// Definition: sci_sciProfile
type WSGSCIProfile struct {
	// Id
	// UUID of the SCI profile for SZ/SCI interface
	Id *string `json:"id,omitempty"`

	// SciPassword
	// SCI password of the SCI profile for SZ/SCI interface
	SciPassword *string `json:"sciPassword,omitempty"`

	// SciPriority
	// Priority of the SCI profile for SZ/SCI interface
	SciPriority *int `json:"sciPriority,omitempty"`

	// SciProfile
	// Profile name of the SCI profile for SZ/SCI interface
	SciProfile *string `json:"sciProfile,omitempty"`

	// SciServerHost
	// SCI server host of the SCI profile for SZ/SCI interface
	SciServerHost *string `json:"sciServerHost,omitempty"`

	// SciServerPort
	// SCI server port of the SCI profile for SZ/SCI interface
	SciServerPort *string `json:"sciServerPort,omitempty"`

	// SciSystemId
	// SCI system UUID of the SCI profile for SZ/SCI interface
	SciSystemId *string `json:"sciSystemId,omitempty"`

	// SciUser
	// SCI user name of the SCI profile for SZ/SCI interface
	SciUser *string `json:"sciUser,omitempty"`
}

type WSGSCIProfileAPIResponse struct {
	*RawAPIResponse
	Data *WSGSCIProfile
}

func newWSGSCIProfileAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSCIProfileAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSCIProfileAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGSCIProfile)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGSCIProfile() *WSGSCIProfile {
	m := new(WSGSCIProfile)
	return m
}

// WSGSCIProfileList
//
// Definition: sci_sciProfileList
type WSGSCIProfileList struct {
	Extra *WSGSCIProfileListExtraType `json:"extra,omitempty"`

	List []*WSGSCIProfile `json:"list,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGSCIProfileList) UnmarshalJSON(b []byte) error {
	type _WSGSCIProfileList WSGSCIProfileList
	tmpType := new(_WSGSCIProfileList)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "extra")
	delete(tmpType.XAdditionalProperties, "list")
	*t = WSGSCIProfileList(*tmpType)
	return nil
}

func (t *WSGSCIProfileList) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Extra != nil {
		tmp["extra"] = t.Extra
	}
	if t.List != nil {
		tmp["list"] = t.List
	}
	return json.Marshal(tmp)
}

type WSGSCIProfileListAPIResponse struct {
	*RawAPIResponse
	Data *WSGSCIProfileList
}

func newWSGSCIProfileListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSCIProfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSCIProfileListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGSCIProfileList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGSCIProfileList() *WSGSCIProfileList {
	m := new(WSGSCIProfileList)
	return m
}

// WSGSCIProfileListExtraType
//
// Definition: sci_sciProfileListExtraType
type WSGSCIProfileListExtraType struct {
	// SciEnabled
	// SCI password of the SCI profile for SZ/SCI interface
	SciEnabled *bool `json:"sciEnabled,omitempty"`
}

func NewWSGSCIProfileListExtraType() *WSGSCIProfileListExtraType {
	m := new(WSGSCIProfileListExtraType)
	return m
}

// AddSciSciEventCode
//
// Use this API command to modify SciAcceptedEventCodes.
//
// Operation ID: addSciSciEventCode
// Operation path: /sci/sciEventCode
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGSCIModifyEventCode
func (s *WSGSCIService) AddSciSciEventCode(ctx context.Context, body *WSGSCIModifyEventCode, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddSciSciEventCode, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*RawAPIResponse), err
}

// AddSciSciProfile
//
// Use this API command to create sciProfile.
//
// Operation ID: addSciSciProfile
// Operation path: /sci/sciProfile
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGSCICreateSciProfile
func (s *WSGSCIService) AddSciSciProfile(ctx context.Context, body *WSGSCICreateSciProfile, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddSciSciProfile, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteSciSciProfile
//
// Use this API command to delete sciProfile list.
//
// Operation ID: deleteSciSciProfile
// Operation path: /sci/sciProfile
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGSCIDeleteSciProfileList
func (s *WSGSCIService) DeleteSciSciProfile(ctx context.Context, body *WSGSCIDeleteSciProfileList, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteSciSciProfile, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*RawAPIResponse), err
}

// DeleteSciSciProfileById
//
// Use this API command to delete sciProfile.
//
// Operation ID: deleteSciSciProfileById
// Operation path: /sci/sciProfile/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGSCIService) DeleteSciSciProfileById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteSciSciProfileById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// FindSciSciEventCode
//
// Use this API command to retrieve SciAcceptedEventCodes.
//
// Operation ID: findSciSciEventCode
// Operation path: /sci/sciEventCode
// Success code: 200 (OK)
func (s *WSGSCIService) FindSciSciEventCode(ctx context.Context, mutators ...RequestMutator) (*WSGSCIEventCodeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGSCIEventCodeAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindSciSciEventCode, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGSCIEventCodeAPIResponse), err
}

// FindSciSciProfile
//
// Use this API command to retrieve sciProfile list.
//
// Operation ID: findSciSciProfile
// Operation path: /sci/sciProfile
// Success code: 200 (OK)
func (s *WSGSCIService) FindSciSciProfile(ctx context.Context, mutators ...RequestMutator) (*WSGSCIProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGSCIProfileListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindSciSciProfile, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGSCIProfileListAPIResponse), err
}

// FindSciSciProfileById
//
// Use this API command to retrieve sciProfile.
//
// Operation ID: findSciSciProfileById
// Operation path: /sci/sciProfile/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGSCIService) FindSciSciProfileById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGSCIProfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGSCIProfileAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindSciSciProfileById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGSCIProfileAPIResponse), err
}

// PartialUpdateSciSciEnabled
//
// Use this API command to modify SCI settings is enabled or not.
//
// Operation ID: partialUpdateSciSciEnabled
// Operation path: /sci/sciEnabled
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGSCIModifySciEnabled
func (s *WSGSCIService) PartialUpdateSciSciEnabled(ctx context.Context, body *WSGSCIModifySciEnabled, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateSciSciEnabled, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*RawAPIResponse), err
}

// PartialUpdateSciSciProfileById
//
// Use this API command to modify sciProfile.
//
// Operation ID: partialUpdateSciSciProfileById
// Operation path: /sci/sciProfile/{id}
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGSCIModifySciProfile
//
// Required parameters:
// - id string
//		- required
func (s *WSGSCIService) PartialUpdateSciSciProfileById(ctx context.Context, body *WSGSCIModifySciProfile, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateSciSciProfileById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*RawAPIResponse), err
}
