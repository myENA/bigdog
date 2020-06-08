package ruckus

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SwitchMSwitchGroupService struct {
	apiClient *VSZClient
}

func NewSwitchMSwitchGroupService(c *VSZClient) *SwitchMSwitchGroupService {
	s := new(SwitchMSwitchGroupService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchGroupService() *SwitchMSwitchGroupService {
	return NewSwitchMSwitchGroupService(ss.apiClient)
}

type SwitchMSwitchGroupAuditId struct {
	// Id
	// Audit Id
	Id *string `json:"id,omitempty"`

	// Name
	// Audit name
	Name *string `json:"name,omitempty"`
}

func NewSwitchMSwitchGroupAuditId() *SwitchMSwitchGroupAuditId {
	m := new(SwitchMSwitchGroupAuditId)
	return m
}

type SwitchMSwitchGroupClientObjectID struct {
	// ExtraValues
	// Extra values of the client
	ExtraValues *SwitchMSwitchGroupClientObjectIDExtraValuesType `json:"extraValues,omitempty"`

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

func NewSwitchMSwitchGroupClientObjectID() *SwitchMSwitchGroupClientObjectID {
	m := new(SwitchMSwitchGroupClientObjectID)
	return m
}

// SwitchMSwitchGroupClientObjectIDExtraValuesType
//
// Extra values of the client
type SwitchMSwitchGroupClientObjectIDExtraValuesType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchGroupClientObjectIDExtraValuesType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchGroupClientObjectIDExtraValuesType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchGroupClientObjectIDExtraValuesType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchGroupClientObjectIDExtraValuesType() *SwitchMSwitchGroupClientObjectIDExtraValuesType {
	m := new(SwitchMSwitchGroupClientObjectIDExtraValuesType)
	return m
}

type SwitchMSwitchGroupCloneConfigByGroup struct {
	Destination []string `json:"destination,omitempty"`

	// Source
	// Source
	Source *string `json:"source,omitempty"`
}

func NewSwitchMSwitchGroupCloneConfigByGroup() *SwitchMSwitchGroupCloneConfigByGroup {
	m := new(SwitchMSwitchGroupCloneConfigByGroup)
	return m
}

type SwitchMSwitchGroupCloneConfigBySwitch struct {
	// Config
	// Config
	Config *string `json:"config,omitempty"`

	Destination []string `json:"destination,omitempty"`

	// Source
	// Source
	Source *string `json:"source,omitempty"`
}

func NewSwitchMSwitchGroupCloneConfigBySwitch() *SwitchMSwitchGroupCloneConfigBySwitch {
	m := new(SwitchMSwitchGroupCloneConfigBySwitch)
	return m
}

type SwitchMSwitchGroupErrorObject struct {
	List []string `json:"list,omitempty"`

	Message *string `json:"message,omitempty"`

	MsgKey *string `json:"msgKey,omitempty"`
}

func NewSwitchMSwitchGroupErrorObject() *SwitchMSwitchGroupErrorObject {
	m := new(SwitchMSwitchGroupErrorObject)
	return m
}

type SwitchMSwitchGroupGetConfigBySwitch struct {
	// Id
	// ID
	Id *string `json:"id,omitempty"`

	// Source
	// Source
	Source *string `json:"source,omitempty"`
}

func NewSwitchMSwitchGroupGetConfigBySwitch() *SwitchMSwitchGroupGetConfigBySwitch {
	m := new(SwitchMSwitchGroupGetConfigBySwitch)
	return m
}

type SwitchMSwitchGroupGroupsByIdsQueryResultList struct {
	Data *SwitchMSwitchGroupQueryResultList `json:"data,omitempty"`

	Error *SwitchMSwitchGroupErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra *SwitchMSwitchGroupGroupsByIdsQueryResultListExtraType `json:"extra,omitempty"`

	// MetaData
	// Metadata of query result list
	MetaData *SwitchMSwitchGroupGroupsByIdsQueryResultListMetaDataType `json:"metaData,omitempty"`

	// Success
	// Query result success or not
	Success *bool `json:"success,omitempty"`
}

func NewSwitchMSwitchGroupGroupsByIdsQueryResultList() *SwitchMSwitchGroupGroupsByIdsQueryResultList {
	m := new(SwitchMSwitchGroupGroupsByIdsQueryResultList)
	return m
}

// SwitchMSwitchGroupGroupsByIdsQueryResultListExtraType
//
// Any additional response
type SwitchMSwitchGroupGroupsByIdsQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchGroupGroupsByIdsQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchGroupGroupsByIdsQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchGroupGroupsByIdsQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchGroupGroupsByIdsQueryResultListExtraType() *SwitchMSwitchGroupGroupsByIdsQueryResultListExtraType {
	m := new(SwitchMSwitchGroupGroupsByIdsQueryResultListExtraType)
	return m
}

// SwitchMSwitchGroupGroupsByIdsQueryResultListMetaDataType
//
// Metadata of query result list
type SwitchMSwitchGroupGroupsByIdsQueryResultListMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchGroupGroupsByIdsQueryResultListMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchGroupGroupsByIdsQueryResultListMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchGroupGroupsByIdsQueryResultListMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchGroupGroupsByIdsQueryResultListMetaDataType() *SwitchMSwitchGroupGroupsByIdsQueryResultListMetaDataType {
	m := new(SwitchMSwitchGroupGroupsByIdsQueryResultListMetaDataType)
	return m
}

type SwitchMSwitchGroupQueryResultList struct {
	// Extra
	// Extra information for query result list
	Extra *SwitchMSwitchGroupQueryResultListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first query result returned out of the complete query result list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more query result after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMSwitchGroupClientObjectID `json:"list,omitempty"`

	// RawDataTotalCount
	// Query result count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total query result count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchGroupQueryResultList() *SwitchMSwitchGroupQueryResultList {
	m := new(SwitchMSwitchGroupQueryResultList)
	return m
}

// SwitchMSwitchGroupQueryResultListExtraType
//
// Extra information for query result list
type SwitchMSwitchGroupQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchGroupQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchGroupQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchGroupQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchGroupQueryResultListExtraType() *SwitchMSwitchGroupQueryResultListExtraType {
	m := new(SwitchMSwitchGroupQueryResultListExtraType)
	return m
}

type SwitchMSwitchGroup struct {
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
	SampledInstant *SwitchMSwitchGroupSampledInstantType `json:"sampledInstant,omitempty"`

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

func NewSwitchMSwitchGroup() *SwitchMSwitchGroup {
	m := new(SwitchMSwitchGroup)
	return m
}

// SwitchMSwitchGroupSampledInstantType
//
// Sampled instant of the switch group
type SwitchMSwitchGroupSampledInstantType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchGroupSampledInstantType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchGroupSampledInstantType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchGroupSampledInstantType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchGroupSampledInstantType() *SwitchMSwitchGroupSampledInstantType {
	m := new(SwitchMSwitchGroupSampledInstantType)
	return m
}

// AddGroup
//
// Use this API command to create a new switch group under an existing domain or switch group.
//
// Request Body:
//	 - body *SwitchMSwitchGroup
func (s *SwitchMSwitchGroupService) AddGroup(ctx context.Context, body *SwitchMSwitchGroup) (*SwitchMSwitchGroupAuditId, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchGroupAuditId
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddGroup, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchGroupAuditId()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteGroupBySwitchGroupId
//
// Use this API command to delete a switch group.
//
// Required Parameters:
// - switchGroupId string
//		- required
func (s *SwitchMSwitchGroupService) DeleteGroupBySwitchGroupId(ctx context.Context, switchGroupId string) (*SwitchMSwitchGroupAuditId, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchGroupAuditId
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteGroupBySwitchGroupId, true)
	req.SetPathParameter("switchGroupId", switchGroupId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchGroupAuditId()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindGroupBySwitchGroupId
//
// Use this API command to retrieve switch group detail.
//
// Required Parameters:
// - switchGroupId string
//		- required
func (s *SwitchMSwitchGroupService) FindGroupBySwitchGroupId(ctx context.Context, switchGroupId string) (*SwitchMSwitchGroup, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchGroup
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindGroupBySwitchGroupId, true)
	req.SetPathParameter("switchGroupId", switchGroupId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchGroup()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindGroupIdsByDomainByDomainId
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
func (s *SwitchMSwitchGroupService) FindGroupIdsByDomainByDomainId(ctx context.Context, domainId string, optionalParams map[string][]string) (*SwitchMSwitchGroupGroupsByIdsQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchGroupGroupsByIdsQueryResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindGroupIdsByDomainByDomainId, true)
	req.SetPathParameter("domainId", domainId)
	if v, ok := optionalParams["showStagingGroup"]; ok && len(v) > 0 {
		req.SetQueryParameter("showStagingGroup", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchGroupGroupsByIdsQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindSwitchClientVisibilityByQueryCriteria
//
// Use this API command to Retrieve the switch client list.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMSwitchGroupService) FindSwitchClientVisibilityByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) ([]byte, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     []byte
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMFindSwitchClientVisibilityByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = make([]byte, 0)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// PartialUpdateGroupBySwitchGroupId
//
// Use this API command to update an existing switch group name or description.
//
// Request Body:
//	 - body *SwitchMSwitchGroup
//
// Required Parameters:
// - switchGroupId string
//		- required
func (s *SwitchMSwitchGroupService) PartialUpdateGroupBySwitchGroupId(ctx context.Context, body *SwitchMSwitchGroup, switchGroupId string) (*SwitchMSwitchGroupAuditId, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchGroupAuditId
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteSwitchMPartialUpdateGroupBySwitchGroupId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("switchGroupId", switchGroupId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchGroupAuditId()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
