package bigdog

// API Version: v9_1

import (
	"context"
	"errors"
	"io"
	"net/http"
)

type SwitchMRegistrationRulesService struct {
	apiClient *VSZClient
}

func NewSwitchMRegistrationRulesService(c *VSZClient) *SwitchMRegistrationRulesService {
	s := new(SwitchMRegistrationRulesService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMRegistrationRulesService() *SwitchMRegistrationRulesService {
	return NewSwitchMRegistrationRulesService(ss.apiClient)
}

// SwitchMRegistrationRulesClientObjectID
//
// Definition: registration_clientObjectID
type SwitchMRegistrationRulesClientObjectID struct {
	// ExtraValues
	// Extra values of the client
	ExtraValues interface{} `json:"extraValues,omitempty"`

	// Id
	// Identifier of the client
	Id *string `json:"id,omitempty"`

	// Label
	// Label of the client
	Label *string `json:"label,omitempty"`

	// Type
	// Type of the client
	Type *string `json:"type,omitempty"`
}

func NewSwitchMRegistrationRulesClientObjectID() *SwitchMRegistrationRulesClientObjectID {
	m := new(SwitchMRegistrationRulesClientObjectID)
	return m
}

// SwitchMRegistrationRulesCreateResult
//
// Definition: registration_createResult
type SwitchMRegistrationRulesCreateResult struct {
	Data *SwitchMRegistrationRulesClientObjectID `json:"data,omitempty"`

	Error *SwitchMRegistrationRulesErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra interface{} `json:"extra,omitempty"`

	// MetaData
	// Matadata of Rule create result
	MetaData interface{} `json:"metaData,omitempty"`

	// Success
	// Create result success or not
	Success *bool `json:"success,omitempty"`
}

type SwitchMRegistrationRulesCreateResultAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMRegistrationRulesCreateResult
}

func newSwitchMRegistrationRulesCreateResultAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMRegistrationRulesCreateResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMRegistrationRulesCreateResultAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMRegistrationRulesCreateResult)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMRegistrationRulesCreateResult() *SwitchMRegistrationRulesCreateResult {
	m := new(SwitchMRegistrationRulesCreateResult)
	return m
}

// SwitchMRegistrationRulesDeleteMultipleResult
//
// Definition: registration_deleteMultipleResult
type SwitchMRegistrationRulesDeleteMultipleResult struct {
	Data *SwitchMRegistrationRulesList `json:"data,omitempty"`

	Error *SwitchMRegistrationRulesErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra interface{} `json:"extra,omitempty"`

	// MetaData
	// Matadata of delete multiple rules result
	MetaData interface{} `json:"metaData,omitempty"`

	// Success
	// Delete multiple result success or not
	Success *bool `json:"success,omitempty"`
}

type SwitchMRegistrationRulesDeleteMultipleResultAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMRegistrationRulesDeleteMultipleResult
}

func newSwitchMRegistrationRulesDeleteMultipleResultAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMRegistrationRulesDeleteMultipleResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMRegistrationRulesDeleteMultipleResultAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMRegistrationRulesDeleteMultipleResult)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMRegistrationRulesDeleteMultipleResult() *SwitchMRegistrationRulesDeleteMultipleResult {
	m := new(SwitchMRegistrationRulesDeleteMultipleResult)
	return m
}

// SwitchMRegistrationRulesDeleteResult
//
// Definition: registration_deleteResult
type SwitchMRegistrationRulesDeleteResult struct {
	Data *SwitchMRegistrationRulesClientObjectID `json:"data,omitempty"`

	Error *SwitchMRegistrationRulesErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra interface{} `json:"extra,omitempty"`

	// MetaData
	// Matadata of Rule delete result
	MetaData interface{} `json:"metaData,omitempty"`

	// Success
	// Delete result success or not
	Success *bool `json:"success,omitempty"`
}

type SwitchMRegistrationRulesDeleteResultAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMRegistrationRulesDeleteResult
}

func newSwitchMRegistrationRulesDeleteResultAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMRegistrationRulesDeleteResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMRegistrationRulesDeleteResultAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMRegistrationRulesDeleteResult)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMRegistrationRulesDeleteResult() *SwitchMRegistrationRulesDeleteResult {
	m := new(SwitchMRegistrationRulesDeleteResult)
	return m
}

// SwitchMRegistrationRulesErrorObject
//
// Definition: registration_errorObject
type SwitchMRegistrationRulesErrorObject struct {
	List []string `json:"list,omitempty"`

	Message *string `json:"message,omitempty"`

	MsgKey *string `json:"msgKey,omitempty"`
}

func NewSwitchMRegistrationRulesErrorObject() *SwitchMRegistrationRulesErrorObject {
	m := new(SwitchMRegistrationRulesErrorObject)
	return m
}

// SwitchMRegistrationRulesList
//
// Definition: registration_list
type SwitchMRegistrationRulesList struct {
	// Extra
	// Any additional response data
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first registration rule returned out of the complete registration rule list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more registration rule after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMRegistrationRulesRegistrationRule `json:"list,omitempty"`

	// RawDataTotalCount
	// Registration rule list count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Registration rule list count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMRegistrationRulesList() *SwitchMRegistrationRulesList {
	m := new(SwitchMRegistrationRulesList)
	return m
}

// SwitchMRegistrationRulesModifyResult
//
// Definition: registration_modifyResult
type SwitchMRegistrationRulesModifyResult struct {
	Data *SwitchMRegistrationRulesClientObjectID `json:"data,omitempty"`

	Error *SwitchMRegistrationRulesErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra interface{} `json:"extra,omitempty"`

	// MetaData
	// Matadata of rule modify result
	MetaData interface{} `json:"metaData,omitempty"`

	// Success
	// Modify result success or not
	Success *bool `json:"success,omitempty"`
}

type SwitchMRegistrationRulesModifyResultAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMRegistrationRulesModifyResult
}

func newSwitchMRegistrationRulesModifyResultAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMRegistrationRulesModifyResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMRegistrationRulesModifyResultAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMRegistrationRulesModifyResult)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMRegistrationRulesModifyResult() *SwitchMRegistrationRulesModifyResult {
	m := new(SwitchMRegistrationRulesModifyResult)
	return m
}

// SwitchMRegistrationRulesRegistrationRule
//
// Definition: registration_registrationRule
type SwitchMRegistrationRulesRegistrationRule struct {
	// CreateDatetime
	// Create datetime of the registration rule
	CreateDatetime *int `json:"createDatetime,omitempty"`

	// CreatorId
	// Creator Id of the registration rule
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorName
	// Creator name of the registration rule
	CreatorName *string `json:"creatorName,omitempty"`

	// Description
	// Description of the registration rule
	Description *string `json:"description,omitempty"`

	// GroupName
	// Switch group name of the registration rule
	GroupName *string `json:"groupName,omitempty"`

	// Id
	// Identifier of the registration rule
	Id *string `json:"id,omitempty"`

	// IpFrom
	// Start IP range of the registration rule
	IpFrom *string `json:"ipFrom,omitempty"`

	// IpRange
	// IP range of the registration rule
	IpRange *string `json:"ipRange,omitempty"`

	// IpTo
	// End IP range of the registration rule
	IpTo *string `json:"ipTo,omitempty"`

	// Label
	// Lable of the registration rule
	Label *string `json:"label,omitempty"`

	// ModelNumber
	// Switch Model number of the registration rule
	ModelNumber *string `json:"modelNumber,omitempty"`

	// Network
	// Network of the registration rule
	Network *string `json:"network,omitempty"`

	// Rank
	// Rank of the registration rule
	Rank *int `json:"rank,omitempty"`

	// SubnetMask
	// Subnet mask of the registration rule
	SubnetMask *string `json:"subnetMask,omitempty"`

	// SwitchGroupId
	// Switch group Id of the registration rule
	SwitchGroupId *string `json:"switchGroupId,omitempty"`

	// Type
	// Type of the registration rule
	// Constraints:
	//    - oneof:[IP_RANGE,SUBNET,MODEL_NUMBER]
	Type *string `json:"type,omitempty"`
}

func NewSwitchMRegistrationRulesRegistrationRule() *SwitchMRegistrationRulesRegistrationRule {
	m := new(SwitchMRegistrationRulesRegistrationRule)
	return m
}

// SwitchMRegistrationRulesRuleQueryResultList
//
// Definition: registration_ruleQueryResultList
type SwitchMRegistrationRulesRuleQueryResultList struct {
	Data *SwitchMRegistrationRulesList `json:"data,omitempty"`

	Error *SwitchMRegistrationRulesErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra interface{} `json:"extra,omitempty"`

	// MetaData
	// Matadata of Rule query result
	MetaData interface{} `json:"metaData,omitempty"`

	// Success
	// Rule query result success or not
	Success *bool `json:"success,omitempty"`
}

type SwitchMRegistrationRulesRuleQueryResultListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMRegistrationRulesRuleQueryResultList
}

func newSwitchMRegistrationRulesRuleQueryResultListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMRegistrationRulesRuleQueryResultListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMRegistrationRulesRuleQueryResultListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMRegistrationRulesRuleQueryResultList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMRegistrationRulesRuleQueryResultList() *SwitchMRegistrationRulesRuleQueryResultList {
	m := new(SwitchMRegistrationRulesRuleQueryResultList)
	return m
}

// SwitchMRegistrationRulesRuleUUIDs
//
// Definition: registration_ruleUUIDs
type SwitchMRegistrationRulesRuleUUIDs []string

func MakeSwitchMRegistrationRulesRuleUUIDs() SwitchMRegistrationRulesRuleUUIDs {
	m := make(SwitchMRegistrationRulesRuleUUIDs, 0)
	return m
}

// AddRegistrationRules
//
// Use this API command to create new switch registration rule.
//
// Operation ID: addRegistrationRules
// Operation path: /registrationRules
// Success code: 201 (Created)
//
// Request body:
//	 - body *SwitchMRegistrationRulesRegistrationRule
func (s *SwitchMRegistrationRulesService) AddRegistrationRules(ctx context.Context, body *SwitchMRegistrationRulesRegistrationRule, mutators ...RequestMutator) (*SwitchMRegistrationRulesCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMRegistrationRulesCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMRegistrationRulesCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddRegistrationRules, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMRegistrationRulesCreateResultAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMRegistrationRulesCreateResultAPIResponse), err
}

// DeleteRegistrationRules
//
// Use this API command to delete multiple switch registration rules.
//
// Operation ID: deleteRegistrationRules
// Operation path: /registrationRules
// Success code: 200 (OK)
//
// Request body:
//	 - body SwitchMRegistrationRulesRuleUUIDs
func (s *SwitchMRegistrationRulesService) DeleteRegistrationRules(ctx context.Context, body SwitchMRegistrationRulesRuleUUIDs, mutators ...RequestMutator) (*SwitchMRegistrationRulesDeleteMultipleResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMRegistrationRulesDeleteMultipleResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMRegistrationRulesDeleteMultipleResultAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteSwitchMDeleteRegistrationRules, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMRegistrationRulesDeleteMultipleResultAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMRegistrationRulesDeleteMultipleResultAPIResponse), err
}

// DeleteRegistrationRulesById
//
// Use this API command to delete a switch registration rule.
//
// Operation ID: deleteRegistrationRulesById
// Operation path: /registrationRules/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *SwitchMRegistrationRulesService) DeleteRegistrationRulesById(ctx context.Context, id string, mutators ...RequestMutator) (*SwitchMRegistrationRulesDeleteResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMRegistrationRulesDeleteResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMRegistrationRulesDeleteResultAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteSwitchMDeleteRegistrationRulesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMRegistrationRulesDeleteResultAPIResponse), err
}

// FindRegistrationRules
//
// Use this API command to retrieves all the registration rules configured in SmartZone.
//
// Operation ID: findRegistrationRules
// Operation path: /registrationRules
// Success code: 200 (OK)
func (s *SwitchMRegistrationRulesService) FindRegistrationRules(ctx context.Context, mutators ...RequestMutator) (*SwitchMRegistrationRulesRuleQueryResultListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMRegistrationRulesRuleQueryResultListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMRegistrationRulesRuleQueryResultListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteSwitchMFindRegistrationRules, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMRegistrationRulesRuleQueryResultListAPIResponse), err
}

// UpdateRegistrationRulesById
//
// Use this API command to modify the registration rule.
//
// Operation ID: updateRegistrationRulesById
// Operation path: /registrationRules/{id}
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMRegistrationRulesRegistrationRule
//
// Required parameters:
// - id string
//		- required
func (s *SwitchMRegistrationRulesService) UpdateRegistrationRulesById(ctx context.Context, body *SwitchMRegistrationRulesRegistrationRule, id string, mutators ...RequestMutator) (*SwitchMRegistrationRulesModifyResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMRegistrationRulesModifyResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMRegistrationRulesModifyResultAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteSwitchMUpdateRegistrationRulesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMRegistrationRulesModifyResultAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMRegistrationRulesModifyResultAPIResponse), err
}
