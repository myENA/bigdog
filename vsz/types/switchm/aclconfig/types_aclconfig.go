package aclconfig

// API Version: v8_1

import (
	"encoding/json"
)

type ACLConfig struct {
	AclRule []*ACLRule `json:"aclRule,omitempty"`

	// AclType
	// ACL Type
	AclType *string `json:"aclType,omitempty"`

	// CreatedTime
	// The create time of the AccessControl
	CreatedTime *int `json:"createdTime,omitempty"`

	// FamilyId
	// Family Id
	FamilyId *string `json:"familyId,omitempty"`

	// GroupId
	// Switch Group Id
	GroupId *string `json:"groupId,omitempty"`

	// Id
	// Id
	Id *string `json:"id,omitempty"`

	// Name
	// ACL Name/ID
	Name *string `json:"name,omitempty"`

	// PushTime
	// Puch Schedule Time
	PushTime *int `json:"pushTime,omitempty"`

	// PushTimeType
	// Puch Config Type
	PushTimeType *string `json:"pushTimeType,omitempty"`

	// SwitchId
	// Switch Id
	SwitchId *string `json:"switchId,omitempty"`

	// UpdatedTime
	// The modify time of the AccessControl
	UpdatedTime *int `json:"updatedTime,omitempty"`
}

type ACLConfigsQueryResult struct {
	// Extra
	// Any additional response data
	Extra *ACLConfigsQueryResultExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first AccessControl returned out of the complete AccessControl list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more AccessControl after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*ACLConfig `json:"list,omitempty"`

	// RawDataTotalCount
	// Total AccessControl count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total AccessControl count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

// ACLConfigsQueryResultExtraType
//
// Any additional response data
type ACLConfigsQueryResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *ACLConfigsQueryResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = ACLConfigsQueryResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *ACLConfigsQueryResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

type ACLRule struct {
	// Action
	// The action of AccessControl Rule
	Action *string `json:"action,omitempty"`

	// DestNetwork
	// The destination network of AccessControl Rule
	DestNetwork *string `json:"destNetwork,omitempty"`

	// DestPort
	// The destination port of AccessControl Rule
	DestPort *string `json:"destPort,omitempty"`

	// Protocol
	// The protocol of AccessControl Rule
	Protocol *string `json:"protocol,omitempty"`

	// Seq
	// The seq of AccessControl Rule
	Seq *int `json:"seq,omitempty"`

	// SrcNetwork
	// The source network of AccessControl Rule
	SrcNetwork *string `json:"srcNetwork,omitempty"`

	// SrcPort
	// The source port of AccessControl Rule
	SrcPort *string `json:"srcPort,omitempty"`
}

type CreateACLConfig struct {
	AclRule []*ACLRule `json:"aclRule,omitempty"`

	// AclType
	// ACL Type
	AclType *string `json:"aclType,omitempty"`

	// FamilyId
	// Family Id
	FamilyId *string `json:"familyId,omitempty"`

	// GroupId
	// Switch Group Id
	GroupId *string `json:"groupId,omitempty"`

	// Name
	// ACL Name/ID
	Name *string `json:"name,omitempty"`

	// PushTime
	// Puch Schedule Time
	PushTime *int `json:"pushTime,omitempty"`

	// PushTimeType
	// Puch Config Type
	PushTimeType *string `json:"pushTimeType,omitempty"`

	// SwitchId
	// Switch Id
	SwitchId *string `json:"switchId,omitempty"`
}

type EmptyResult struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *EmptyResult) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = EmptyResult{XAdditionalProperties: tmp}
	return nil
}

func (t *EmptyResult) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

type UpdateACLConfig struct {
	AclRule []*ACLRule `json:"aclRule,omitempty"`

	// AclType
	// ACL Type
	AclType *string `json:"aclType,omitempty"`

	// PushTime
	// Puch Schedule Time
	PushTime *int `json:"pushTime,omitempty"`

	// PushTimeType
	// Puch Config Type
	PushTimeType *string `json:"pushTimeType,omitempty"`
}
