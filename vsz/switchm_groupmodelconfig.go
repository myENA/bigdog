package vsz

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SwitchMGroupModelConfigService struct {
	apiClient *APIClient
}

func NewSwitchMGroupModelConfigService(c *APIClient) *SwitchMGroupModelConfigService {
	s := new(SwitchMGroupModelConfigService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMGroupModelConfigService() *SwitchMGroupModelConfigService {
	return NewSwitchMGroupModelConfigService(ss.apiClient)
}

type SwitchMGroupModelConfigGroupConfigAuditId struct {
	// Id
	// The identifier of the Group Model Config
	Id *string `json:"id,omitempty"`

	// Name
	// The name of the Group Model Config
	Name *string `json:"name,omitempty"`
}

func NewSwitchMGroupModelConfigGroupConfigAuditId() *SwitchMGroupModelConfigGroupConfigAuditId {
	m := new(SwitchMGroupModelConfigGroupConfigAuditId)
	return m
}

type SwitchMGroupModelConfig struct {
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

func NewSwitchMGroupModelConfig() *SwitchMGroupModelConfig {
	m := new(SwitchMGroupModelConfig)
	return m
}

type SwitchMGroupModelConfigQueryResult struct {
	// Extra
	// Any additional response data
	Extra *SwitchMGroupModelConfigQueryResultExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first Group Model Config returned out of the complete Group Model Configs list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Group Model Configs after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMGroupModelConfig `json:"list,omitempty"`

	// RawDataTotalCount
	// Total Group Model Configs count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Group Model Configs count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMGroupModelConfigQueryResult() *SwitchMGroupModelConfigQueryResult {
	m := new(SwitchMGroupModelConfigQueryResult)
	return m
}

// SwitchMGroupModelConfigQueryResultExtraType
//
// Any additional response data
type SwitchMGroupModelConfigQueryResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMGroupModelConfigQueryResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMGroupModelConfigQueryResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMGroupModelConfigQueryResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMGroupModelConfigQueryResultExtraType() *SwitchMGroupModelConfigQueryResultExtraType {
	m := new(SwitchMGroupModelConfigQueryResultExtraType)
	return m
}

type SwitchMGroupModelConfigSelectedIds struct {
	SelectedIdList []string `json:"selectedIdList,omitempty"`
}

func NewSwitchMGroupModelConfigSelectedIds() *SwitchMGroupModelConfigSelectedIds {
	m := new(SwitchMGroupModelConfigSelectedIds)
	return m
}

type SwitchMGroupModelConfigUpdateGroupConfigResultList struct {
	// Extra
	// Any additional response data
	Extra *SwitchMGroupModelConfigUpdateGroupConfigResultListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first updated Group Model Config returned out of the complete Group Model Configs list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more updated Group Model Configs after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMGroupModelConfigGroupConfigAuditId `json:"list,omitempty"`

	// RawDataTotalCount
	// Total updated Group Model Configs count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total updated Group Model Configs count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMGroupModelConfigUpdateGroupConfigResultList() *SwitchMGroupModelConfigUpdateGroupConfigResultList {
	m := new(SwitchMGroupModelConfigUpdateGroupConfigResultList)
	return m
}

// SwitchMGroupModelConfigUpdateGroupConfigResultListExtraType
//
// Any additional response data
type SwitchMGroupModelConfigUpdateGroupConfigResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMGroupModelConfigUpdateGroupConfigResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMGroupModelConfigUpdateGroupConfigResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMGroupModelConfigUpdateGroupConfigResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMGroupModelConfigUpdateGroupConfigResultListExtraType() *SwitchMGroupModelConfigUpdateGroupConfigResultListExtraType {
	m := new(SwitchMGroupModelConfigUpdateGroupConfigResultListExtraType)
	return m
}

// FindGroupModelConfigsByQueryCriteria
//
// Use this API command to retrieve the list of group model configs.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMGroupModelConfigService) FindGroupModelConfigsByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMGroupModelConfigQueryResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMGroupModelConfigQueryResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMFindGroupModelConfigsByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMGroupModelConfigQueryResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateGroupModelConfigsByGroupId
//
// Use this API command to add or remove the model family of a group config.
//
// Request Body:
//	 - body *SwitchMGroupModelConfigSelectedIds
//
// Required Parameters:
// - groupId string
//		- required
func (s *SwitchMGroupModelConfigService) UpdateGroupModelConfigsByGroupId(ctx context.Context, body *SwitchMGroupModelConfigSelectedIds, groupId string) (*SwitchMGroupModelConfigUpdateGroupConfigResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMGroupModelConfigUpdateGroupConfigResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, groupId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSwitchMUpdateGroupModelConfigsByGroupId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("groupId", groupId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMGroupModelConfigUpdateGroupConfigResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
