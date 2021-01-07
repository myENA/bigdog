package bigdog

// API Version: v9_1

import (
	"context"
	"encoding/json"
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

func newSwitchMRegistrationRulesCreateResultAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(SwitchMRegistrationRulesCreateResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *SwitchMRegistrationRulesCreateResultAPIResponse) Hydrate() error {
	r.Data = new(SwitchMRegistrationRulesCreateResult)
	return json.NewDecoder(r).Decode(r.Data)
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

func newSwitchMRegistrationRulesDeleteMultipleResultAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(SwitchMRegistrationRulesDeleteMultipleResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *SwitchMRegistrationRulesDeleteMultipleResultAPIResponse) Hydrate() error {
	r.Data = new(SwitchMRegistrationRulesDeleteMultipleResult)
	return json.NewDecoder(r).Decode(r.Data)
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

func newSwitchMRegistrationRulesDeleteResultAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(SwitchMRegistrationRulesDeleteResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *SwitchMRegistrationRulesDeleteResultAPIResponse) Hydrate() error {
	r.Data = new(SwitchMRegistrationRulesDeleteResult)
	return json.NewDecoder(r).Decode(r.Data)
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

func newSwitchMRegistrationRulesModifyResultAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(SwitchMRegistrationRulesModifyResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *SwitchMRegistrationRulesModifyResultAPIResponse) Hydrate() error {
	r.Data = new(SwitchMRegistrationRulesModifyResult)
	return json.NewDecoder(r).Decode(r.Data)
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

func newSwitchMRegistrationRulesRuleQueryResultListAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(SwitchMRegistrationRulesRuleQueryResultListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *SwitchMRegistrationRulesRuleQueryResultListAPIResponse) Hydrate() error {
	r.Data = new(SwitchMRegistrationRulesRuleQueryResultList)
	return json.NewDecoder(r).Decode(r.Data)
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
// Operation ID: addRegistrationRules
//
// Use this API command to create new switch registration rule.
//
// Request Body:
//	 - body *SwitchMRegistrationRulesRegistrationRule
func (s *SwitchMRegistrationRulesService) AddRegistrationRules(ctx context.Context, body *SwitchMRegistrationRulesRegistrationRule, mutators ...RequestMutator) (*SwitchMRegistrationRulesCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMRegistrationRulesCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMAddRegistrationRules, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMRegistrationRulesCreateResult()
	rm, err = handleAPIResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteRegistrationRules
//
// Operation ID: deleteRegistrationRules
//
// Use this API command to delete multiple switch registration rules.
//
// Request Body:
//	 - body SwitchMRegistrationRulesRuleUUIDs
func (s *SwitchMRegistrationRulesService) DeleteRegistrationRules(ctx context.Context, body SwitchMRegistrationRulesRuleUUIDs, mutators ...RequestMutator) (*SwitchMRegistrationRulesDeleteMultipleResultAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMRegistrationRulesDeleteMultipleResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteSwitchMDeleteRegistrationRules, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMRegistrationRulesDeleteMultipleResult()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteRegistrationRulesById
//
// Operation ID: deleteRegistrationRulesById
//
// Use this API command to delete a switch registration rule.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMRegistrationRulesService) DeleteRegistrationRulesById(ctx context.Context, id string, mutators ...RequestMutator) (*SwitchMRegistrationRulesDeleteResultAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMRegistrationRulesDeleteResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteSwitchMDeleteRegistrationRulesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMRegistrationRulesDeleteResult()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRegistrationRules
//
// Operation ID: findRegistrationRules
//
// Use this API command to retrieves all the registration rules configured in SmartZone.
func (s *SwitchMRegistrationRulesService) FindRegistrationRules(ctx context.Context, mutators ...RequestMutator) (*SwitchMRegistrationRulesRuleQueryResultListAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMRegistrationRulesRuleQueryResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSwitchMFindRegistrationRules, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMRegistrationRulesRuleQueryResultList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateRegistrationRulesById
//
// Operation ID: updateRegistrationRulesById
//
// Use this API command to modify the registration rule.
//
// Request Body:
//	 - body *SwitchMRegistrationRulesRegistrationRule
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMRegistrationRulesService) UpdateRegistrationRulesById(ctx context.Context, body *SwitchMRegistrationRulesRegistrationRule, id string, mutators ...RequestMutator) (*SwitchMRegistrationRulesModifyResultAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMRegistrationRulesModifyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPut, RouteSwitchMUpdateRegistrationRulesById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMRegistrationRulesModifyResult()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
