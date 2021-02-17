package bigdog

// API Version: v9_1

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"
)

type SwitchMGroupModelConfigService struct {
	apiClient *VSZClient
}

func NewSwitchMGroupModelConfigService(c *VSZClient) *SwitchMGroupModelConfigService {
	s := new(SwitchMGroupModelConfigService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMGroupModelConfigService() *SwitchMGroupModelConfigService {
	return NewSwitchMGroupModelConfigService(ss.apiClient)
}

// SwitchMGroupModelConfigGroupConfigAuditId
//
// Definition: groupModelConfig_groupConfigAuditId
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

// SwitchMGroupModelConfig
//
// Definition: groupModelConfig_groupModelConfig
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

// SwitchMGroupModelConfigQueryResult
//
// Definition: groupModelConfig_groupModelConfigQueryResult
type SwitchMGroupModelConfigQueryResult struct {
	// Extra
	// Any additional response data
	Extra interface{} `json:"extra,omitempty"`

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

type SwitchMGroupModelConfigQueryResultAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMGroupModelConfigQueryResult
}

func newSwitchMGroupModelConfigQueryResultAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMGroupModelConfigQueryResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMGroupModelConfigQueryResultAPIResponse) Hydrate() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return nil
		}
		return r.err
	}
	data := new(SwitchMGroupModelConfigQueryResult)
	if err := r.doHydrate(data); err != nil {
		return err
	}
	r.Data = data
	return nil
}
func NewSwitchMGroupModelConfigQueryResult() *SwitchMGroupModelConfigQueryResult {
	m := new(SwitchMGroupModelConfigQueryResult)
	return m
}

// SwitchMGroupModelConfigSelectedIds
//
// Definition: groupModelConfig_selectedIds
type SwitchMGroupModelConfigSelectedIds struct {
	SelectedIdList []string `json:"selectedIdList,omitempty"`
}

func NewSwitchMGroupModelConfigSelectedIds() *SwitchMGroupModelConfigSelectedIds {
	m := new(SwitchMGroupModelConfigSelectedIds)
	return m
}

// SwitchMGroupModelConfigUpdateGroupConfigResultList
//
// Definition: groupModelConfig_updateGroupConfigResultList
type SwitchMGroupModelConfigUpdateGroupConfigResultList struct {
	// Extra
	// Any additional response data
	Extra interface{} `json:"extra,omitempty"`

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

type SwitchMGroupModelConfigUpdateGroupConfigResultListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMGroupModelConfigUpdateGroupConfigResultList
}

func newSwitchMGroupModelConfigUpdateGroupConfigResultListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMGroupModelConfigUpdateGroupConfigResultListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMGroupModelConfigUpdateGroupConfigResultListAPIResponse) Hydrate() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return nil
		}
		return r.err
	}
	data := new(SwitchMGroupModelConfigUpdateGroupConfigResultList)
	if err := r.doHydrate(data); err != nil {
		return err
	}
	r.Data = data
	return nil
}
func NewSwitchMGroupModelConfigUpdateGroupConfigResultList() *SwitchMGroupModelConfigUpdateGroupConfigResultList {
	m := new(SwitchMGroupModelConfigUpdateGroupConfigResultList)
	return m
}

// FindGroupModelConfigsByQueryCriteria
//
// Use this API command to retrieve the list of group model configs.
//
// Operation ID: findGroupModelConfigsByQueryCriteria
// Operation path: /groupModelConfigs/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMGroupModelConfigService) FindGroupModelConfigsByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMGroupModelConfigQueryResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMGroupModelConfigQueryResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMFindGroupModelConfigsByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMGroupModelConfigQueryResultAPIResponse), err
}

// UpdateGroupModelConfigsByGroupId
//
// Use this API command to add or remove the model family of a group config.
//
// Operation ID: updateGroupModelConfigsByGroupId
// Operation path: /groupModelConfigs/{groupId}
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMGroupModelConfigSelectedIds
//
// Required parameters:
// - groupId string
//		- required
func (s *SwitchMGroupModelConfigService) UpdateGroupModelConfigsByGroupId(ctx context.Context, body *SwitchMGroupModelConfigSelectedIds, groupId string, mutators ...RequestMutator) (*SwitchMGroupModelConfigUpdateGroupConfigResultListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMGroupModelConfigUpdateGroupConfigResultListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteSwitchMUpdateGroupModelConfigsByGroupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("groupId", groupId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMGroupModelConfigUpdateGroupConfigResultListAPIResponse), err
}
