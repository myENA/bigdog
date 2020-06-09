package bigdog

// API Version: v9_0

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

type SwitchMRegistrationRulesClientObjectID struct {
	// ExtraValues
	// Extra values of the client
	ExtraValues *SwitchMRegistrationRulesClientObjectIDExtraValuesType `json:"extraValues,omitempty"`

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

// SwitchMRegistrationRulesClientObjectIDExtraValuesType
//
// Extra values of the client
type SwitchMRegistrationRulesClientObjectIDExtraValuesType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMRegistrationRulesClientObjectIDExtraValuesType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMRegistrationRulesClientObjectIDExtraValuesType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMRegistrationRulesClientObjectIDExtraValuesType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMRegistrationRulesClientObjectIDExtraValuesType() *SwitchMRegistrationRulesClientObjectIDExtraValuesType {
	m := new(SwitchMRegistrationRulesClientObjectIDExtraValuesType)
	return m
}

type SwitchMRegistrationRulesCreateResult struct {
	Data *SwitchMRegistrationRulesClientObjectID `json:"data,omitempty"`

	Error *SwitchMRegistrationRulesErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra *SwitchMRegistrationRulesCreateResultExtraType `json:"extra,omitempty"`

	// MetaData
	// Matadata of Rule create result
	MetaData *SwitchMRegistrationRulesCreateResultMetaDataType `json:"metaData,omitempty"`

	// Success
	// Create result success or not
	Success *bool `json:"success,omitempty"`
}

func NewSwitchMRegistrationRulesCreateResult() *SwitchMRegistrationRulesCreateResult {
	m := new(SwitchMRegistrationRulesCreateResult)
	return m
}

// SwitchMRegistrationRulesCreateResultExtraType
//
// Any additional response
type SwitchMRegistrationRulesCreateResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMRegistrationRulesCreateResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMRegistrationRulesCreateResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMRegistrationRulesCreateResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMRegistrationRulesCreateResultExtraType() *SwitchMRegistrationRulesCreateResultExtraType {
	m := new(SwitchMRegistrationRulesCreateResultExtraType)
	return m
}

// SwitchMRegistrationRulesCreateResultMetaDataType
//
// Matadata of Rule create result
type SwitchMRegistrationRulesCreateResultMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMRegistrationRulesCreateResultMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMRegistrationRulesCreateResultMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMRegistrationRulesCreateResultMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMRegistrationRulesCreateResultMetaDataType() *SwitchMRegistrationRulesCreateResultMetaDataType {
	m := new(SwitchMRegistrationRulesCreateResultMetaDataType)
	return m
}

type SwitchMRegistrationRulesDeleteMultipleResult struct {
	Data *SwitchMRegistrationRulesList `json:"data,omitempty"`

	Error *SwitchMRegistrationRulesErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra *SwitchMRegistrationRulesDeleteMultipleResultExtraType `json:"extra,omitempty"`

	// MetaData
	// Matadata of delete multiple rules result
	MetaData *SwitchMRegistrationRulesDeleteMultipleResultMetaDataType `json:"metaData,omitempty"`

	// Success
	// Delete multiple result success or not
	Success *bool `json:"success,omitempty"`
}

func NewSwitchMRegistrationRulesDeleteMultipleResult() *SwitchMRegistrationRulesDeleteMultipleResult {
	m := new(SwitchMRegistrationRulesDeleteMultipleResult)
	return m
}

// SwitchMRegistrationRulesDeleteMultipleResultExtraType
//
// Any additional response
type SwitchMRegistrationRulesDeleteMultipleResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMRegistrationRulesDeleteMultipleResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMRegistrationRulesDeleteMultipleResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMRegistrationRulesDeleteMultipleResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMRegistrationRulesDeleteMultipleResultExtraType() *SwitchMRegistrationRulesDeleteMultipleResultExtraType {
	m := new(SwitchMRegistrationRulesDeleteMultipleResultExtraType)
	return m
}

// SwitchMRegistrationRulesDeleteMultipleResultMetaDataType
//
// Matadata of delete multiple rules result
type SwitchMRegistrationRulesDeleteMultipleResultMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMRegistrationRulesDeleteMultipleResultMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMRegistrationRulesDeleteMultipleResultMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMRegistrationRulesDeleteMultipleResultMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMRegistrationRulesDeleteMultipleResultMetaDataType() *SwitchMRegistrationRulesDeleteMultipleResultMetaDataType {
	m := new(SwitchMRegistrationRulesDeleteMultipleResultMetaDataType)
	return m
}

type SwitchMRegistrationRulesDeleteResult struct {
	Data *SwitchMRegistrationRulesClientObjectID `json:"data,omitempty"`

	Error *SwitchMRegistrationRulesErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra *SwitchMRegistrationRulesDeleteResultExtraType `json:"extra,omitempty"`

	// MetaData
	// Matadata of Rule delete result
	MetaData *SwitchMRegistrationRulesDeleteResultMetaDataType `json:"metaData,omitempty"`

	// Success
	// Delete result success or not
	Success *bool `json:"success,omitempty"`
}

func NewSwitchMRegistrationRulesDeleteResult() *SwitchMRegistrationRulesDeleteResult {
	m := new(SwitchMRegistrationRulesDeleteResult)
	return m
}

// SwitchMRegistrationRulesDeleteResultExtraType
//
// Any additional response
type SwitchMRegistrationRulesDeleteResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMRegistrationRulesDeleteResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMRegistrationRulesDeleteResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMRegistrationRulesDeleteResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMRegistrationRulesDeleteResultExtraType() *SwitchMRegistrationRulesDeleteResultExtraType {
	m := new(SwitchMRegistrationRulesDeleteResultExtraType)
	return m
}

// SwitchMRegistrationRulesDeleteResultMetaDataType
//
// Matadata of Rule delete result
type SwitchMRegistrationRulesDeleteResultMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMRegistrationRulesDeleteResultMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMRegistrationRulesDeleteResultMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMRegistrationRulesDeleteResultMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMRegistrationRulesDeleteResultMetaDataType() *SwitchMRegistrationRulesDeleteResultMetaDataType {
	m := new(SwitchMRegistrationRulesDeleteResultMetaDataType)
	return m
}

type SwitchMRegistrationRulesErrorObject struct {
	List []string `json:"list,omitempty"`

	Message *string `json:"message,omitempty"`

	MsgKey *string `json:"msgKey,omitempty"`
}

func NewSwitchMRegistrationRulesErrorObject() *SwitchMRegistrationRulesErrorObject {
	m := new(SwitchMRegistrationRulesErrorObject)
	return m
}

type SwitchMRegistrationRulesList struct {
	// Extra
	// Any additional response data
	Extra *SwitchMRegistrationRulesListExtraType `json:"extra,omitempty"`

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

// SwitchMRegistrationRulesListExtraType
//
// Any additional response data
type SwitchMRegistrationRulesListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMRegistrationRulesListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMRegistrationRulesListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMRegistrationRulesListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMRegistrationRulesListExtraType() *SwitchMRegistrationRulesListExtraType {
	m := new(SwitchMRegistrationRulesListExtraType)
	return m
}

type SwitchMRegistrationRulesModifyResult struct {
	Data *SwitchMRegistrationRulesClientObjectID `json:"data,omitempty"`

	Error *SwitchMRegistrationRulesErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra *SwitchMRegistrationRulesModifyResultExtraType `json:"extra,omitempty"`

	// MetaData
	// Matadata of rule modify result
	MetaData *SwitchMRegistrationRulesModifyResultMetaDataType `json:"metaData,omitempty"`

	// Success
	// Modify result success or not
	Success *bool `json:"success,omitempty"`
}

func NewSwitchMRegistrationRulesModifyResult() *SwitchMRegistrationRulesModifyResult {
	m := new(SwitchMRegistrationRulesModifyResult)
	return m
}

// SwitchMRegistrationRulesModifyResultExtraType
//
// Any additional response
type SwitchMRegistrationRulesModifyResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMRegistrationRulesModifyResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMRegistrationRulesModifyResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMRegistrationRulesModifyResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMRegistrationRulesModifyResultExtraType() *SwitchMRegistrationRulesModifyResultExtraType {
	m := new(SwitchMRegistrationRulesModifyResultExtraType)
	return m
}

// SwitchMRegistrationRulesModifyResultMetaDataType
//
// Matadata of rule modify result
type SwitchMRegistrationRulesModifyResultMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMRegistrationRulesModifyResultMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMRegistrationRulesModifyResultMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMRegistrationRulesModifyResultMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMRegistrationRulesModifyResultMetaDataType() *SwitchMRegistrationRulesModifyResultMetaDataType {
	m := new(SwitchMRegistrationRulesModifyResultMetaDataType)
	return m
}

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

type SwitchMRegistrationRulesRuleQueryResultList struct {
	Data *SwitchMRegistrationRulesList `json:"data,omitempty"`

	Error *SwitchMRegistrationRulesErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra *SwitchMRegistrationRulesRuleQueryResultListExtraType `json:"extra,omitempty"`

	// MetaData
	// Matadata of Rule query result
	MetaData *SwitchMRegistrationRulesRuleQueryResultListMetaDataType `json:"metaData,omitempty"`

	// Success
	// Rule query result success or not
	Success *bool `json:"success,omitempty"`
}

func NewSwitchMRegistrationRulesRuleQueryResultList() *SwitchMRegistrationRulesRuleQueryResultList {
	m := new(SwitchMRegistrationRulesRuleQueryResultList)
	return m
}

// SwitchMRegistrationRulesRuleQueryResultListExtraType
//
// Any additional response
type SwitchMRegistrationRulesRuleQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMRegistrationRulesRuleQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMRegistrationRulesRuleQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMRegistrationRulesRuleQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMRegistrationRulesRuleQueryResultListExtraType() *SwitchMRegistrationRulesRuleQueryResultListExtraType {
	m := new(SwitchMRegistrationRulesRuleQueryResultListExtraType)
	return m
}

// SwitchMRegistrationRulesRuleQueryResultListMetaDataType
//
// Matadata of Rule query result
type SwitchMRegistrationRulesRuleQueryResultListMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMRegistrationRulesRuleQueryResultListMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMRegistrationRulesRuleQueryResultListMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMRegistrationRulesRuleQueryResultListMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMRegistrationRulesRuleQueryResultListMetaDataType() *SwitchMRegistrationRulesRuleQueryResultListMetaDataType {
	m := new(SwitchMRegistrationRulesRuleQueryResultListMetaDataType)
	return m
}

type SwitchMRegistrationRulesRuleUUIDs []string

func MakeSwitchMRegistrationRulesRuleUUIDs() SwitchMRegistrationRulesRuleUUIDs {
	m := make(SwitchMRegistrationRulesRuleUUIDs, 0)
	return m
}

// AddRegistrationRules
//
// Use this API command to create new switch registration rule.
//
// Request Body:
//	 - body *SwitchMRegistrationRulesRegistrationRule
func (s *SwitchMRegistrationRulesService) AddRegistrationRules(ctx context.Context, body *SwitchMRegistrationRulesRegistrationRule) (*SwitchMRegistrationRulesCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddRegistrationRules, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMRegistrationRulesCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteRegistrationRules
//
// Use this API command to delete multiple switch registration rules.
//
// Request Body:
//	 - body SwitchMRegistrationRulesRuleUUIDs
func (s *SwitchMRegistrationRulesService) DeleteRegistrationRules(ctx context.Context, body SwitchMRegistrationRulesRuleUUIDs) (*SwitchMRegistrationRulesDeleteMultipleResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteRegistrationRules, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMRegistrationRulesDeleteMultipleResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteRegistrationRulesById
//
// Use this API command to delete a switch registration rule.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMRegistrationRulesService) DeleteRegistrationRulesById(ctx context.Context, id string) (*SwitchMRegistrationRulesDeleteResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteRegistrationRulesById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMRegistrationRulesDeleteResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRegistrationRules
//
// Use this API command to retrieves all the registration rules configured in SmartZone.
func (s *SwitchMRegistrationRulesService) FindRegistrationRules(ctx context.Context) (*SwitchMRegistrationRulesRuleQueryResultList, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindRegistrationRules, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMRegistrationRulesRuleQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateRegistrationRulesById
//
// Use this API command to modify the registration rule.
//
// Request Body:
//	 - body *SwitchMRegistrationRulesRegistrationRule
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMRegistrationRulesService) UpdateRegistrationRulesById(ctx context.Context, body *SwitchMRegistrationRulesRegistrationRule, id string) (*SwitchMRegistrationRulesModifyResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPut, RouteSwitchMUpdateRegistrationRulesById, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMRegistrationRulesModifyResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
