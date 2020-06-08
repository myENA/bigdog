package ruckus

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SwitchMSwitchGroupModelConfigService struct {
	apiClient *VSZClient
}

func NewSwitchMSwitchGroupModelConfigService(c *VSZClient) *SwitchMSwitchGroupModelConfigService {
	s := new(SwitchMSwitchGroupModelConfigService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchGroupModelConfigService() *SwitchMSwitchGroupModelConfigService {
	return NewSwitchMSwitchGroupModelConfigService(ss.apiClient)
}

type SwitchMSwitchGroupModelConfigGroupConfigAuditId struct {
	// Id
	// The identifier of the Group Model Config
	Id *string `json:"id,omitempty"`

	// Name
	// The name of the Group Model Config
	Name *string `json:"name,omitempty"`
}

func NewSwitchMSwitchGroupModelConfigGroupConfigAuditId() *SwitchMSwitchGroupModelConfigGroupConfigAuditId {
	m := new(SwitchMSwitchGroupModelConfigGroupConfigAuditId)
	return m
}

type SwitchMSwitchGroupModelConfigGroupModelConfig struct {
	// CreatedTime
	// The create time of the Group Model Config
	CreatedTime *int `json:"createdTime,omitempty"`

	// FamilyId
	// Family Id
	FamilyId *string `json:"familyId,omitempty"`

	// GroupId
	// Group Id
	GroupId *string `json:"groupId,omitempty"`

	// HasAclConfig
	// Indicates if there is any ACLConfig in the Group Model Config
	HasAclConfig *bool `json:"hasAclConfig,omitempty"`

	// HasConfigured
	// Indicates if there is any config in the Group Model Config
	HasConfigured *bool `json:"hasConfigured,omitempty"`

	// HasSelected
	// Indicates if this Group Model Config has been selected
	HasSelected *bool `json:"hasSelected,omitempty"`

	// HasStaticRouteConfig
	// Indicates if there is any StaticRoute in the Group Model Config
	HasStaticRouteConfig *bool `json:"hasStaticRouteConfig,omitempty"`

	// HasVlanConfig
	// Indicates if there is any VlanConfig in the Group Model Config
	HasVlanConfig *bool `json:"hasVlanConfig,omitempty"`

	// Id
	// Id
	Id *string `json:"id,omitempty"`

	// UpdatedTime
	// The update time of the Group Model Config
	UpdatedTime *int `json:"updatedTime,omitempty"`
}

func NewSwitchMSwitchGroupModelConfigGroupModelConfig() *SwitchMSwitchGroupModelConfigGroupModelConfig {
	m := new(SwitchMSwitchGroupModelConfigGroupModelConfig)
	return m
}

type SwitchMSwitchGroupModelConfigGroupModelConfigQueryResult struct {
	// Extra
	// Any additional response data
	Extra *SwitchMSwitchGroupModelConfigGroupModelConfigQueryResult `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first Group Model Config returned out of the complete Group Model Configs list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Group Model Configs after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMSwitchGroupModelConfigGroupModelConfigQueryResult `json:"list,omitempty"`

	// RawDataTotalCount
	// Total Group Model Configs count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Group Model Configs count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchGroupModelConfigGroupModelConfigQueryResult() *SwitchMSwitchGroupModelConfigGroupModelConfigQueryResult {
	m := new(SwitchMSwitchGroupModelConfigGroupModelConfigQueryResult)
	return m
}

// SwitchMSwitchGroupModelConfigGroupModelConfigQueryResultExtraType
//
// Any additional response data
type SwitchMSwitchGroupModelConfigGroupModelConfigQueryResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchGroupModelConfigGroupModelConfigQueryResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchGroupModelConfigGroupModelConfigQueryResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchGroupModelConfigGroupModelConfigQueryResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchGroupModelConfigGroupModelConfigQueryResultExtraType() *SwitchMSwitchGroupModelConfigGroupModelConfigQueryResultExtraType {
	m := new(SwitchMSwitchGroupModelConfigGroupModelConfigQueryResultExtraType)
	return m
}

type SwitchMSwitchGroupModelConfigSelectedIds struct {
	SelectedIdList []string `json:"selectedIdList,omitempty"`
}

func NewSwitchMSwitchGroupModelConfigSelectedIds() *SwitchMSwitchGroupModelConfigSelectedIds {
	m := new(SwitchMSwitchGroupModelConfigSelectedIds)
	return m
}

type SwitchMSwitchGroupModelConfigUpdateGroupConfigResultList struct {
	// Extra
	// Any additional response data
	Extra *SwitchMSwitchGroupModelConfigUpdateGroupConfigResultList `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first updated Group Model Config returned out of the complete Group Model Configs list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more updated Group Model Configs after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMSwitchGroupModelConfigUpdateGroupConfigResultList `json:"list,omitempty"`

	// RawDataTotalCount
	// Total updated Group Model Configs count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total updated Group Model Configs count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchGroupModelConfigUpdateGroupConfigResultList() *SwitchMSwitchGroupModelConfigUpdateGroupConfigResultList {
	m := new(SwitchMSwitchGroupModelConfigUpdateGroupConfigResultList)
	return m
}

// SwitchMSwitchGroupModelConfigUpdateGroupConfigResultListExtraType
//
// Any additional response data
type SwitchMSwitchGroupModelConfigUpdateGroupConfigResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchGroupModelConfigUpdateGroupConfigResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchGroupModelConfigUpdateGroupConfigResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchGroupModelConfigUpdateGroupConfigResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchGroupModelConfigUpdateGroupConfigResultListExtraType() *SwitchMSwitchGroupModelConfigUpdateGroupConfigResultListExtraType {
	m := new(SwitchMSwitchGroupModelConfigUpdateGroupConfigResultListExtraType)
	return m
}

// FindGroupModelConfigsByQueryCriteria
//
// Use this API command to retrieve the list of group model configs.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMSwitchGroupModelConfigService) FindGroupModelConfigsByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchGroupModelConfigGroupModelConfigQueryResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchGroupModelConfigGroupModelConfigQueryResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMFindGroupModelConfigsByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchGroupModelConfigGroupModelConfigQueryResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateGroupModelConfigsByGroupId
//
// Use this API command to add or remove the model family of a group config.
//
// Request Body:
//	 - body *SwitchMSwitchGroupModelConfigSelectedIds
//
// Required Parameters:
// - groupId string
//		- required
func (s *SwitchMSwitchGroupModelConfigService) UpdateGroupModelConfigsByGroupId(ctx context.Context, body *SwitchMSwitchGroupModelConfigSelectedIds, groupId string) (*SwitchMSwitchGroupModelConfigUpdateGroupConfigResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchGroupModelConfigUpdateGroupConfigResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSwitchMUpdateGroupModelConfigsByGroupId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("groupId", groupId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchGroupModelConfigUpdateGroupConfigResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
