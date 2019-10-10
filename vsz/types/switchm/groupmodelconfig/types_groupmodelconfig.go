package groupmodelconfig

// API Version: v8_1

import (
	"encoding/json"
)

type GroupConfigAuditID struct {
	// ID
	// The identifier of the Group Model Config
	ID *string `json:"id,omitempty"`

	// Name
	// The name of the Group Model Config
	Name *string `json:"name,omitempty"`
}

type GroupModelConfig struct {
	// CreatedTime
	// The create time of the Group Model Config
	CreatedTime *int `json:"createdTime,omitempty"`

	// FamilyID
	// Family Id
	FamilyID *string `json:"familyId,omitempty"`

	// GroupID
	// Group Id
	GroupID *string `json:"groupId,omitempty"`

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

	// ID
	// Id
	ID *string `json:"id,omitempty"`

	// UpdatedTime
	// The update time of the Group Model Config
	UpdatedTime *int `json:"updatedTime,omitempty"`
}

type GroupModelConfigQueryResult struct {
	// Extra
	// Any additional response data
	Extra *GroupModelConfigQueryResultExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first Group Model Config returned out of the complete Group Model Configs list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Group Model Configs after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*GroupModelConfig `json:"list,omitempty"`

	// RawDataTotalCount
	// Total Group Model Configs count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Group Model Configs count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

// GroupModelConfigQueryResultExtraType
//
// Any additional response data
type GroupModelConfigQueryResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *GroupModelConfigQueryResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = GroupModelConfigQueryResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *GroupModelConfigQueryResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

type SelectedIds struct {
	SelectedIDList []string `json:"selectedIdList,omitempty"`
}

type UpdateGroupConfigResultList struct {
	// Extra
	// Any additional response data
	Extra *UpdateGroupConfigResultListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first updated Group Model Config returned out of the complete Group Model Configs list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more updated Group Model Configs after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*GroupConfigAuditID `json:"list,omitempty"`

	// RawDataTotalCount
	// Total updated Group Model Configs count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total updated Group Model Configs count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

// UpdateGroupConfigResultListExtraType
//
// Any additional response data
type UpdateGroupConfigResultListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *UpdateGroupConfigResultListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = UpdateGroupConfigResultListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *UpdateGroupConfigResultListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}
