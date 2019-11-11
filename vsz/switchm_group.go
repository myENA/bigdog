package vsz

// API Version: v8_1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type SwitchMGroupService struct {
	apiClient *APIClient
}

func NewSwitchMGroupService(c *APIClient) *SwitchMGroupService {
	s := new(SwitchMGroupService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMGroupService() *SwitchMGroupService {
	return NewSwitchMGroupService(ss.apiClient)
}

type SwitchMGroupAuditId struct {
	// Id
	// Audit Id
	Id *string `json:"id,omitempty"`

	// Name
	// Audit name
	Name *string `json:"name,omitempty"`
}

type SwitchMGroupClientObjectID struct {
	// ExtraValues
	// Extra values of the client
	ExtraValues *SwitchMGroupClientObjectIDExtraValuesType `json:"extraValues,omitempty"`

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

// SwitchMGroupClientObjectIDExtraValuesType
//
// Extra values of the client
type SwitchMGroupClientObjectIDExtraValuesType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMGroupClientObjectIDExtraValuesType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMGroupClientObjectIDExtraValuesType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMGroupClientObjectIDExtraValuesType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

type SwitchMGroupCloneConfigByGroup struct {
	Destination []string `json:"destination,omitempty"`

	// Source
	// Source
	Source *string `json:"source,omitempty"`
}

type SwitchMGroupCloneConfigBySwitch struct {
	// Config
	// Config
	Config *string `json:"config,omitempty"`

	Destination []string `json:"destination,omitempty"`

	// Source
	// Source
	Source *string `json:"source,omitempty"`
}

type SwitchMGroupDeleteSwitchGroupResult struct {
	*SwitchMGroupAuditId
}

type SwitchMGroupErrorObject struct {
	List []string `json:"list,omitempty"`

	Message *string `json:"message,omitempty"`

	MsgKey *string `json:"msgKey,omitempty"`
}

type SwitchMGroupGetConfigBySwitch struct {
	// Id
	// ID
	Id *string `json:"id,omitempty"`

	// Source
	// Source
	Source *string `json:"source,omitempty"`
}

type SwitchMGroupsByIdsQueryResultList struct {
	Data *SwitchMGroupQueryResultList `json:"data,omitempty"`

	Error *SwitchMGroupErrorObject `json:"error,omitempty"`

	// Extra
	// Any additional response
	Extra *SwitchMGroupsByIdsQueryResultListExtraType `json:"extra,omitempty"`

	// MetaData
	// Metadata of query result list
	MetaData *SwitchMGroupsByIdsQueryResultListMetaDataType `json:"metaData,omitempty"`

	// Success
	// Query result success or not
	Success *bool `json:"success,omitempty"`
}

// SwitchMGroupsByIdsQueryResultListExtraType
//
// Any additional response
type SwitchMGroupsByIdsQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMGroupsByIdsQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMGroupsByIdsQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMGroupsByIdsQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

// SwitchMGroupsByIdsQueryResultListMetaDataType
//
// Metadata of query result list
type SwitchMGroupsByIdsQueryResultListMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMGroupsByIdsQueryResultListMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMGroupsByIdsQueryResultListMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMGroupsByIdsQueryResultListMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

type SwitchMGroupQueryResultList struct {
	// Extra
	// Extra information for query result list
	Extra *SwitchMGroupQueryResultListExtraType `json:"extra,omitempty"`

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

// SwitchMGroupQueryResultListExtraType
//
// Extra information for query result list
type SwitchMGroupQueryResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMGroupQueryResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMGroupQueryResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMGroupQueryResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

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
	SampledInstant *SwitchMGroupSwitchGroupSampledInstantType `json:"sampledInstant,omitempty"`

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

type SwitchMGroupSwitchGroupQueryResult struct {
	*SwitchMGroupSwitchGroup
}

// SwitchMGroupSwitchGroupSampledInstantType
//
// Sampled instant of the switch group
type SwitchMGroupSwitchGroupSampledInstantType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMGroupSwitchGroupSampledInstantType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMGroupSwitchGroupSampledInstantType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMGroupSwitchGroupSampledInstantType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

type SwitchMGroupUpdateSwitchGroup struct {
	*SwitchMGroupSwitchGroup
}

type SwitchMGroupUpdateSwitchGroupResult struct {
	*SwitchMGroupAuditId
}

// AddGroup
//
// Use this API command to create a new switch group under an existing domain or switch group.
//
// Request Body:
//	 - body *SwitchMGroupSwitchGroup
func (s *SwitchMGroupService) AddGroup(ctx context.Context, body *SwitchMGroupSwitchGroup) (*SwitchMGroupAuditId, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteGroupBySwitchGroupId
//
// Use this API command to delete a switch group.
//
// Path Parameters:
// - pSwitchGroupId string
//		- required
func (s *SwitchMGroupService) DeleteGroupBySwitchGroupId(ctx context.Context, pSwitchGroupId string) (*SwitchMGroupDeleteSwitchGroupResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindGroupBySwitchGroupId
//
// Use this API command to retrieve switch group detail.
//
// Path Parameters:
// - pSwitchGroupId string
//		- required
func (s *SwitchMGroupService) FindGroupBySwitchGroupId(ctx context.Context, pSwitchGroupId string) (*SwitchMGroupSwitchGroupQueryResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindGroupIdsByDomainByDomainId
//
// Use this API command to retrieve the switch groups by domain ID.
//
// Path Parameters:
// - pDomainId string
//		- required
func (s *SwitchMGroupService) FindGroupIdsByDomainByDomainId(ctx context.Context, pDomainId string) (*SwitchMGroupsByIdsQueryResultList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateGroupBySwitchGroupId
//
// Use this API command to update an existing switch group name or description.
//
// Request Body:
//	 - body *SwitchMGroupUpdateSwitchGroup
//
// Path Parameters:
// - pSwitchGroupId string
//		- required
func (s *SwitchMGroupService) PartialUpdateGroupBySwitchGroupId(ctx context.Context, body *SwitchMGroupUpdateSwitchGroup, pSwitchGroupId string) (*SwitchMGroupUpdateSwitchGroupResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
