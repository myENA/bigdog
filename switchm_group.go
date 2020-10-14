package bigdog

// API Version: v9_1

import (
	"context"
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
// Operation ID: addGroup
//
// Use this API command to create a new switch group under an existing domain or switch group.
//
// Request Body:
//	 - body *SwitchMGroupSwitchGroup
func (s *SwitchMGroupService) AddGroup(ctx context.Context, body *SwitchMGroupSwitchGroup, mutators ...RequestMutator) (*SwitchMGroupAuditId, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMGroupAuditId
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddGroup, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMGroupAuditId()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteGroupBySwitchGroupId
//
// Operation ID: deleteGroupBySwitchGroupId
//
// Use this API command to delete a switch group.
//
// Required Parameters:
// - switchGroupId string
//		- required
func (s *SwitchMGroupService) DeleteGroupBySwitchGroupId(ctx context.Context, switchGroupId string, mutators ...RequestMutator) (*SwitchMGroupAuditId, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMGroupAuditId
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteGroupBySwitchGroupId, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("switchGroupId", switchGroupId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMGroupAuditId()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindGroupBySwitchGroupId
//
// Operation ID: findGroupBySwitchGroupId
//
// Use this API command to retrieve switch group detail.
//
// Required Parameters:
// - switchGroupId string
//		- required
func (s *SwitchMGroupService) FindGroupBySwitchGroupId(ctx context.Context, switchGroupId string, mutators ...RequestMutator) (*SwitchMGroupSwitchGroup, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMGroupSwitchGroup
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindGroupBySwitchGroupId, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("switchGroupId", switchGroupId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMGroupSwitchGroup()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindGroupIdsByDomainByDomainId
//
// Operation ID: findGroupIdsByDomainByDomainId
//
// Use this API command to retrieve the switch groups by domain ID.
//
// Required Parameters:
// - domainId string
//		- required
//
// Optional Parameters:
// - showStagingGroup string
//		- nullable
func (s *SwitchMGroupService) FindGroupIdsByDomainByDomainId(ctx context.Context, domainId string, optionalParams map[string][]string, mutators ...RequestMutator) (*SwitchMGroupsByIdsQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMGroupsByIdsQueryResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindGroupIdsByDomainByDomainId, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("domainId", domainId)
	if v, ok := optionalParams["showStagingGroup"]; ok && len(v) > 0 {
		req.SetQueryParameterValues("showStagingGroup", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMGroupsByIdsQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindSwitchClientVisibilityByQueryCriteria
//
// Operation ID: findSwitchClientVisibilityByQueryCriteria
//
// Use this API command to Retrieve the switch client list.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMGroupService) FindSwitchClientVisibilityByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMFindSwitchClientVisibilityByQueryCriteria, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateGroupBySwitchGroupId
//
// Operation ID: partialUpdateGroupBySwitchGroupId
//
// Use this API command to update an existing switch group name or description.
//
// Request Body:
//	 - body *SwitchMGroupSwitchGroup
//
// Required Parameters:
// - switchGroupId string
//		- required
func (s *SwitchMGroupService) PartialUpdateGroupBySwitchGroupId(ctx context.Context, body *SwitchMGroupSwitchGroup, switchGroupId string, mutators ...RequestMutator) (*SwitchMGroupAuditId, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMGroupAuditId
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteSwitchMPartialUpdateGroupBySwitchGroupId, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("switchGroupId", switchGroupId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMGroupAuditId()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateGroupBySwitchGroupId
//
// Operation ID: updateGroupBySwitchGroupId
//
// Use this API command to update an existing switch group name, description.
//
// Request Body:
//	 - body *SwitchMGroupUpdateSwitchGroupByPut
//
// Required Parameters:
// - switchGroupId string
//		- required
func (s *SwitchMGroupService) UpdateGroupBySwitchGroupId(ctx context.Context, body *SwitchMGroupUpdateSwitchGroupByPut, switchGroupId string, mutators ...RequestMutator) (*SwitchMGroupAuditId, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMGroupAuditId
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSwitchMUpdateGroupBySwitchGroupId, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("switchGroupId", switchGroupId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMGroupAuditId()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateGroupFirmwareByDomain
//
// Operation ID: updateGroupFirmwareByDomain
//
// Use this API command to update default firmware of groups under a specific domain
//
// Request Body:
//	 - body *SwitchMGroupSwitchGroupFirmwareByDomain
func (s *SwitchMGroupService) UpdateGroupFirmwareByDomain(ctx context.Context, body *SwitchMGroupSwitchGroupFirmwareByDomain, mutators ...RequestMutator) (*SwitchMGroupAuditIdList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMGroupAuditIdList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSwitchMUpdateGroupFirmwareByDomain, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMGroupAuditIdList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
