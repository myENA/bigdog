package bigdog

// API Version: v9_1

import (
	"context"
	"errors"
	"io"
	"net/http"
)

type SwitchMGroupService struct {
	apiClient *VSZClient
}

func NewSwitchMGroupService(c *VSZClient) *SwitchMGroupService {
	s := new(SwitchMGroupService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMGroupService() *SwitchMGroupService {
	return NewSwitchMGroupService(ss.apiClient)
}

// SwitchMGroupAuditId
//
// Definition: group_auditId
type SwitchMGroupAuditId struct {
	// Id
	// Audit Id
	Id *string `json:"id,omitempty"`

	// Name
	// Audit name
	Name *string `json:"name,omitempty"`
}

type SwitchMGroupAuditIdAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMGroupAuditId
}

func newSwitchMGroupAuditIdAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMGroupAuditIdAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMGroupAuditIdAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMGroupAuditId)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMGroupAuditId() *SwitchMGroupAuditId {
	m := new(SwitchMGroupAuditId)
	return m
}

// SwitchMGroupAuditIdList
//
// Definition: group_auditIdList
type SwitchMGroupAuditIdList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMGroupAuditId `json:"list,omitempty"`

	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchMGroupAuditIdListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMGroupAuditIdList
}

func newSwitchMGroupAuditIdListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMGroupAuditIdListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMGroupAuditIdListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMGroupAuditIdList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMGroupAuditIdList() *SwitchMGroupAuditIdList {
	m := new(SwitchMGroupAuditIdList)
	return m
}

// SwitchMGroupClientObjectID
//
// Definition: group_clientObjectID
type SwitchMGroupClientObjectID struct {
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

func NewSwitchMGroupClientObjectID() *SwitchMGroupClientObjectID {
	m := new(SwitchMGroupClientObjectID)
	return m
}

// SwitchMGroupCloneConfigByGroup
//
// Definition: group_cloneConfigByGroup
type SwitchMGroupCloneConfigByGroup struct {
	Destination []string `json:"destination,omitempty"`

	// Source
	// Source
	Source *string `json:"source,omitempty"`
}

func NewSwitchMGroupCloneConfigByGroup() *SwitchMGroupCloneConfigByGroup {
	m := new(SwitchMGroupCloneConfigByGroup)
	return m
}

// SwitchMGroupCloneConfigBySwitch
//
// Definition: group_cloneConfigBySwitch
type SwitchMGroupCloneConfigBySwitch struct {
	// Config
	// Config
	Config *string `json:"config,omitempty"`

	Destination []string `json:"destination,omitempty"`

	// Source
	// Source
	Source *string `json:"source,omitempty"`
}

func NewSwitchMGroupCloneConfigBySwitch() *SwitchMGroupCloneConfigBySwitch {
	m := new(SwitchMGroupCloneConfigBySwitch)
	return m
}

// SwitchMGroupErrorObject
//
// Definition: group_errorObject
type SwitchMGroupErrorObject struct {
	List []string `json:"list,omitempty"`

	Message *string `json:"message,omitempty"`

	MsgKey *string `json:"msgKey,omitempty"`
}

func NewSwitchMGroupErrorObject() *SwitchMGroupErrorObject {
	m := new(SwitchMGroupErrorObject)
	return m
}

// SwitchMGroupGetConfigBySwitch
//
// Definition: group_getConfigBySwitch
type SwitchMGroupGetConfigBySwitch struct {
	// Id
	// ID
	Id *string `json:"id,omitempty"`

	// Source
	// Source
	Source *string `json:"source,omitempty"`
}

func NewSwitchMGroupGetConfigBySwitch() *SwitchMGroupGetConfigBySwitch {
	m := new(SwitchMGroupGetConfigBySwitch)
	return m
}

// SwitchMGroupsByIdsQueryResultList
//
// Definition: group_groupsByIdsQueryResultList
type SwitchMGroupsByIdsQueryResultList struct {
	Data *SwitchMGroupQueryResultList `json:"data,omitempty"`

	Error *SwitchMGroupErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra interface{} `json:"extra,omitempty"`

	// MetaData
	// Metadata of query result list
	MetaData interface{} `json:"metaData,omitempty"`

	// Success
	// Query result success or not
	Success *bool `json:"success,omitempty"`
}

type SwitchMGroupsByIdsQueryResultListAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMGroupsByIdsQueryResultList
}

func newSwitchMGroupsByIdsQueryResultListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMGroupsByIdsQueryResultListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMGroupsByIdsQueryResultListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMGroupsByIdsQueryResultList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMGroupsByIdsQueryResultList() *SwitchMGroupsByIdsQueryResultList {
	m := new(SwitchMGroupsByIdsQueryResultList)
	return m
}

// SwitchMGroupQueryResultList
//
// Definition: group_queryResultList
type SwitchMGroupQueryResultList struct {
	// Extra
	// Extra information for query result list
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first query result returned out of the complete query result list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more query result after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMGroupClientObjectID `json:"list,omitempty"`

	// RawDataTotalCount
	// Query result count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total query result count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMGroupQueryResultList() *SwitchMGroupQueryResultList {
	m := new(SwitchMGroupQueryResultList)
	return m
}

// SwitchMGroupSwitchGroup
//
// Definition: group_switchGroup
type SwitchMGroupSwitchGroup struct {
	// CreateDatetime
	// Create datetime of the switch group
	CreateDatetime *int `json:"createDatetime,omitempty"`

	// CreatorId
	// Creator Id of the switch group
	CreatorId *string `json:"creatorId,omitempty"`

	// Description
	// Description of the switch group
	Description *string `json:"description,omitempty"`

	// DomainId
	// Identifier of the management domain to which the switch group belong
	DomainId *string `json:"domainId,omitempty"`

	// EnableTwoFactor
	// Enable two factor authentication. Only support FIPS mode
	EnableTwoFactor *bool `json:"enableTwoFactor,omitempty"`

	// Firmware
	// Firmware of the switch group
	Firmware *string `json:"firmware,omitempty"`

	// Id
	// Identifier of the switch group
	Id *string `json:"id,omitempty"`

	// LevelOne
	// Level one  of the switch group
	LevelOne *bool `json:"levelOne,omitempty"`

	// LevelTwo
	// Level two of the switch group
	LevelTwo *bool `json:"levelTwo,omitempty"`

	// Name
	// Name of the switch group
	Name *string `json:"name,omitempty"`

	// SampledInstant
	// Sampled instant of the switch group
	SampledInstant interface{} `json:"sampledInstant,omitempty"`

	// SwitchGroupLevelOneId
	// Level one Id of the switch group
	SwitchGroupLevelOneId *string `json:"switchGroupLevelOneId,omitempty"`

	// SwitchGroupLevelTwoId
	// Level two Id of the switch group
	SwitchGroupLevelTwoId *string `json:"switchGroupLevelTwoId,omitempty"`

	// TenantId
	// Tenant Id of the switch group
	TenantId *string `json:"tenantId,omitempty"`
}

type SwitchMGroupSwitchGroupAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMGroupSwitchGroup
}

func newSwitchMGroupSwitchGroupAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMGroupSwitchGroupAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMGroupSwitchGroupAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMGroupSwitchGroup)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMGroupSwitchGroup() *SwitchMGroupSwitchGroup {
	m := new(SwitchMGroupSwitchGroup)
	return m
}

// SwitchMGroupSwitchGroupFirmwareByDomain
//
// Definition: group_switchGroupFirmwareByDomain
type SwitchMGroupSwitchGroupFirmwareByDomain struct {
	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// Firmware
	// Firmware Version
	Firmware *string `json:"firmware,omitempty"`
}

func NewSwitchMGroupSwitchGroupFirmwareByDomain() *SwitchMGroupSwitchGroupFirmwareByDomain {
	m := new(SwitchMGroupSwitchGroupFirmwareByDomain)
	return m
}

// SwitchMGroupUpdateSwitchGroupByPut
//
// Definition: group_updateSwitchGroupByPut
type SwitchMGroupUpdateSwitchGroupByPut struct {
	// CreateDatetime
	// Create datetime of the switch group
	CreateDatetime *int `json:"createDatetime,omitempty"`

	// CreatorId
	// Creator Id of the switch group
	CreatorId *string `json:"creatorId,omitempty"`

	// Description
	// Description of the switch group
	Description *string `json:"description,omitempty"`

	// DomainId
	// Identifier of the management domain to which the switch group belong
	DomainId *string `json:"domainId,omitempty"`

	// EnableTwoFactor
	// Enable two factor authentication. Only support FIPS mode
	EnableTwoFactor *bool `json:"enableTwoFactor,omitempty"`

	// Firmware
	// Firmware of the switch group
	Firmware *string `json:"firmware,omitempty"`

	Id *string `json:"id,omitempty"`

	// LevelOne
	// Level one  of the switch group
	LevelOne *bool `json:"levelOne,omitempty"`

	// LevelTwo
	// Level two of the switch group
	LevelTwo *bool `json:"levelTwo,omitempty"`

	// Name
	// Name of the switch group
	Name *string `json:"name,omitempty"`

	// SampledInstant
	// Sampled instant of the switch group
	SampledInstant interface{} `json:"sampledInstant,omitempty"`

	// SwitchGroupLevelOneId
	// Level one Id of the switch group
	SwitchGroupLevelOneId *string `json:"switchGroupLevelOneId,omitempty"`

	// SwitchGroupLevelTwoId
	// Level two Id of the switch group
	SwitchGroupLevelTwoId *string `json:"switchGroupLevelTwoId,omitempty"`

	// TenantId
	// Tenant Id of the switch group
	TenantId *string `json:"tenantId,omitempty"`
}

func NewSwitchMGroupUpdateSwitchGroupByPut() *SwitchMGroupUpdateSwitchGroupByPut {
	m := new(SwitchMGroupUpdateSwitchGroupByPut)
	return m
}

// AddGroup
//
// Use this API command to create a new switch group under an existing domain or switch group.
//
// Operation ID: addGroup
// Operation path: /group
// Success code: 201 (Created)
//
// Request body:
//	 - body *SwitchMGroupSwitchGroup
func (s *SwitchMGroupService) AddGroup(ctx context.Context, body *SwitchMGroupSwitchGroup, mutators ...RequestMutator) (*SwitchMGroupAuditIdAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMGroupAuditIdAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddGroup, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMGroupAuditIdAPIResponse), err
}

// DeleteGroupBySwitchGroupId
//
// Use this API command to delete a switch group.
//
// Operation ID: deleteGroupBySwitchGroupId
// Operation path: /group/{switchGroupId}
// Success code: 200 (OK)
//
// Required parameters:
// - switchGroupId string
//		- required
func (s *SwitchMGroupService) DeleteGroupBySwitchGroupId(ctx context.Context, switchGroupId string, mutators ...RequestMutator) (*SwitchMGroupAuditIdAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMGroupAuditIdAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteSwitchMDeleteGroupBySwitchGroupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("switchGroupId", switchGroupId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMGroupAuditIdAPIResponse), err
}

// FindGroupBySwitchGroupId
//
// Use this API command to retrieve switch group detail.
//
// Operation ID: findGroupBySwitchGroupId
// Operation path: /group/{switchGroupId}
// Success code: 200 (OK)
//
// Required parameters:
// - switchGroupId string
//		- required
func (s *SwitchMGroupService) FindGroupBySwitchGroupId(ctx context.Context, switchGroupId string, mutators ...RequestMutator) (*SwitchMGroupSwitchGroupAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMGroupSwitchGroupAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteSwitchMFindGroupBySwitchGroupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("switchGroupId", switchGroupId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMGroupSwitchGroupAPIResponse), err
}

// FindGroupIdsByDomainByDomainId
//
// Use this API command to retrieve the switch groups by domain ID.
//
// Operation ID: findGroupIdsByDomainByDomainId
// Operation path: /group/ids/byDomain/{domainId}
// Success code: 200 (OK)
//
// Required parameters:
// - domainId string
//		- required
//
// Optional parameters:
// - showStagingGroup string
//		- nullable
func (s *SwitchMGroupService) FindGroupIdsByDomainByDomainId(ctx context.Context, domainId string, optionalParams map[string][]string, mutators ...RequestMutator) (*SwitchMGroupsByIdsQueryResultListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMGroupsByIdsQueryResultListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteSwitchMFindGroupIdsByDomainByDomainId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("domainId", domainId)
	if v, ok := optionalParams["showStagingGroup"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("showStagingGroup", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMGroupsByIdsQueryResultListAPIResponse), err
}

// FindSwitchClientVisibilityByQueryCriteria
//
// Use this API command to Retrieve the switch client list.
//
// Operation ID: findSwitchClientVisibilityByQueryCriteria
// Operation path: /switchClientVisibility/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMGroupService) FindSwitchClientVisibilityByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*FileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newFileAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMFindSwitchClientVisibilityByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*FileAPIResponse), err
}

// PartialUpdateGroupBySwitchGroupId
//
// Use this API command to update an existing switch group name or description.
//
// Operation ID: partialUpdateGroupBySwitchGroupId
// Operation path: /group/{switchGroupId}
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMGroupSwitchGroup
//
// Required parameters:
// - switchGroupId string
//		- required
func (s *SwitchMGroupService) PartialUpdateGroupBySwitchGroupId(ctx context.Context, body *SwitchMGroupSwitchGroup, switchGroupId string, mutators ...RequestMutator) (*SwitchMGroupAuditIdAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMGroupAuditIdAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteSwitchMPartialUpdateGroupBySwitchGroupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("switchGroupId", switchGroupId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMGroupAuditIdAPIResponse), err
}

// UpdateGroupBySwitchGroupId
//
// Use this API command to update an existing switch group name, description.
//
// Operation ID: updateGroupBySwitchGroupId
// Operation path: /group/{switchGroupId}
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMGroupUpdateSwitchGroupByPut
//
// Required parameters:
// - switchGroupId string
//		- required
func (s *SwitchMGroupService) UpdateGroupBySwitchGroupId(ctx context.Context, body *SwitchMGroupUpdateSwitchGroupByPut, switchGroupId string, mutators ...RequestMutator) (*SwitchMGroupAuditIdAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMGroupAuditIdAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteSwitchMUpdateGroupBySwitchGroupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("switchGroupId", switchGroupId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMGroupAuditIdAPIResponse), err
}

// UpdateGroupFirmwareByDomain
//
// Use this API command to update default firmware of groups under a specific domain
//
// Operation ID: updateGroupFirmwareByDomain
// Operation path: /group/firmware/byDomain
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMGroupSwitchGroupFirmwareByDomain
func (s *SwitchMGroupService) UpdateGroupFirmwareByDomain(ctx context.Context, body *SwitchMGroupSwitchGroupFirmwareByDomain, mutators ...RequestMutator) (*SwitchMGroupAuditIdListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMGroupAuditIdListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteSwitchMUpdateGroupFirmwareByDomain, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMGroupAuditIdListAPIResponse), err
}
