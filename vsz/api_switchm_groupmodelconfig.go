package vsz

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SwitchMGroupmodelconfigService struct {
	apiClient *APIClient
}

func NewSwitchMGroupmodelconfigService(c *APIClient) *SwitchMGroupmodelconfigService {
	s := new(SwitchMGroupmodelconfigService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMGroupmodelconfigService() *SwitchMGroupmodelconfigService {
	return NewSwitchMGroupmodelconfigService(ss.apiClient)
}

type SwitchMGroupmodelconfigGroupConfigAuditId struct {
	// Id
	// The identifier of the Group Model Config
	Id *string `json:"id,omitempty"`

	// Name
	// The name of the Group Model Config
	Name *string `json:"name,omitempty"`
}

func NewSwitchMGroupmodelconfigGroupConfigAuditId() *SwitchMGroupmodelconfigGroupConfigAuditId {
	m := new(SwitchMGroupmodelconfigGroupConfigAuditId)
	return m
}

type SwitchMGroupmodelconfig struct {
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

func NewSwitchMGroupmodelconfig() *SwitchMGroupmodelconfig {
	m := new(SwitchMGroupmodelconfig)
	return m
}

type SwitchMGroupmodelconfigQueryResult struct {
	// Extra
	// Any additional response data
	Extra *SwitchMGroupmodelconfigQueryResultExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first Group Model Config returned out of the complete Group Model Configs list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Group Model Configs after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMGroupmodelconfig `json:"list,omitempty"`

	// RawDataTotalCount
	// Total Group Model Configs count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Group Model Configs count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMGroupmodelconfigQueryResult() *SwitchMGroupmodelconfigQueryResult {
	m := new(SwitchMGroupmodelconfigQueryResult)
	return m
}

// SwitchMGroupmodelconfigQueryResultExtraType
//
// Any additional response data
type SwitchMGroupmodelconfigQueryResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMGroupmodelconfigQueryResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMGroupmodelconfigQueryResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMGroupmodelconfigQueryResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMGroupmodelconfigQueryResultExtraType() *SwitchMGroupmodelconfigQueryResultExtraType {
	m := new(SwitchMGroupmodelconfigQueryResultExtraType)
	return m
}

type SwitchMGroupmodelconfigSelectedIds struct {
	SelectedIdList []string `json:"selectedIdList,omitempty"`
}

func NewSwitchMGroupmodelconfigSelectedIds() *SwitchMGroupmodelconfigSelectedIds {
	m := new(SwitchMGroupmodelconfigSelectedIds)
	return m
}

type SwitchMGroupmodelconfigUpdateGroupConfigResultList struct {
	// Extra
	// Any additional response data
	Extra *SwitchMGroupmodelconfigUpdateGroupConfigResultListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first updated Group Model Config returned out of the complete Group Model Configs list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more updated Group Model Configs after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMGroupmodelconfigGroupConfigAuditId `json:"list,omitempty"`

	// RawDataTotalCount
	// Total updated Group Model Configs count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total updated Group Model Configs count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMGroupmodelconfigUpdateGroupConfigResultList() *SwitchMGroupmodelconfigUpdateGroupConfigResultList {
	m := new(SwitchMGroupmodelconfigUpdateGroupConfigResultList)
	return m
}

// SwitchMGroupmodelconfigUpdateGroupConfigResultListExtraType
//
// Any additional response data
type SwitchMGroupmodelconfigUpdateGroupConfigResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMGroupmodelconfigUpdateGroupConfigResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMGroupmodelconfigUpdateGroupConfigResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMGroupmodelconfigUpdateGroupConfigResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMGroupmodelconfigUpdateGroupConfigResultListExtraType() *SwitchMGroupmodelconfigUpdateGroupConfigResultListExtraType {
	m := new(SwitchMGroupmodelconfigUpdateGroupConfigResultListExtraType)
	return m
}

// FindGroupModelConfigsByQueryCriteria
//
// Use this API command to retrieve the list of group model configs.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMGroupmodelconfigService) FindGroupModelConfigsByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMGroupmodelconfigQueryResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMGroupmodelconfigQueryResult
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
	resp = NewSwitchMGroupmodelconfigQueryResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateGroupModelConfigsByGroupId
//
// Use this API command to add or remove the model family of a group config.
//
// Request Body:
//	 - body *SwitchMGroupmodelconfigSelectedIds
//
// Required Parameters:
// - groupId string
//		- required
func (s *SwitchMGroupmodelconfigService) UpdateGroupModelConfigsByGroupId(ctx context.Context, body *SwitchMGroupmodelconfigSelectedIds, groupId string) (*SwitchMGroupmodelconfigUpdateGroupConfigResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMGroupmodelconfigUpdateGroupConfigResultList
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
	resp = NewSwitchMGroupmodelconfigUpdateGroupConfigResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
