package vsz

// API Version: v9_0

import (
	"encoding/json"
)

type SwitchMACLConfig struct {
	// AclRule
	// Constraints:
	//    - nullable
	AclRule []*SwitchMACLConfigACLRule `json:"aclRule,omitempty" validate:"omitempty,dive"`

	// AclType
	// ACL Type
	// Constraints:
	//    - nullable
	//    - oneof:[STANDARD,EXTENDED]
	AclType *string `json:"aclType,omitempty" validate:"omitempty,oneof=STANDARD EXTENDED"`

	// CreatedTime
	// The create time of the AccessControl
	// Constraints:
	//    - nullable
	CreatedTime *int `json:"createdTime,omitempty"`

	// FamilyId
	// Family Id
	// Constraints:
	//    - nullable
	FamilyId *string `json:"familyId,omitempty"`

	// GroupId
	// Switch Group Id
	// Constraints:
	//    - nullable
	GroupId *string `json:"groupId,omitempty"`

	// Id
	// Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// ACL Name/ID
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// PushTime
	// Puch Schedule Time
	// Constraints:
	//    - nullable
	PushTime *int `json:"pushTime,omitempty"`

	// PushTimeType
	// Puch Config Type
	// Constraints:
	//    - nullable
	//    - oneof:[NOW,SCHEDULE]
	PushTimeType *string `json:"pushTimeType,omitempty" validate:"omitempty,oneof=NOW SCHEDULE"`

	// SwitchId
	// Switch Id
	// Constraints:
	//    - nullable
	SwitchId *string `json:"switchId,omitempty"`

	// UpdatedTime
	// The modify time of the AccessControl
	// Constraints:
	//    - nullable
	UpdatedTime *int `json:"updatedTime,omitempty"`
}

func NewSwitchMACLConfig() *SwitchMACLConfig {
	m := new(SwitchMACLConfig)
	return m
}

type SwitchMACLConfigsQueryResult struct {
	// Extra
	// Any additional response data
	// Constraints:
	//    - nullable
	Extra *SwitchMACLConfigsQueryResultExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first AccessControl returned out of the complete AccessControl list
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more AccessControl after the current displayed list
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*SwitchMACLConfig `json:"list,omitempty" validate:"omitempty,dive"`

	// RawDataTotalCount
	// Total AccessControl count
	// Constraints:
	//    - nullable
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total AccessControl count in this response
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMACLConfigsQueryResult() *SwitchMACLConfigsQueryResult {
	m := new(SwitchMACLConfigsQueryResult)
	return m
}

// SwitchMACLConfigsQueryResultExtraType
//
// Any additional response data
// Constraints:
//    - nullable
type SwitchMACLConfigsQueryResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMACLConfigsQueryResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMACLConfigsQueryResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMACLConfigsQueryResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMACLConfigsQueryResultExtraType() *SwitchMACLConfigsQueryResultExtraType {
	m := new(SwitchMACLConfigsQueryResultExtraType)
	return m
}

type SwitchMACLConfigACLRule struct {
	// Action
	// The action of AccessControl Rule
	// Constraints:
	//    - nullable
	//    - oneof:[PERMIT,DENY]
	Action *string `json:"action,omitempty" validate:"omitempty,oneof=PERMIT DENY"`

	// DestNetwork
	// The destination network of AccessControl Rule
	// Constraints:
	//    - nullable
	DestNetwork *string `json:"destNetwork,omitempty"`

	// DestPort
	// The destination port of AccessControl Rule
	// Constraints:
	//    - nullable
	DestPort *string `json:"destPort,omitempty"`

	// Protocol
	// The protocol of AccessControl Rule
	// Constraints:
	//    - nullable
	//    - oneof:[IP,TCP,UDP]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=IP TCP UDP"`

	// Seq
	// The seq of AccessControl Rule
	// Constraints:
	//    - nullable
	Seq *int `json:"seq,omitempty"`

	// SrcNetwork
	// The source network of AccessControl Rule
	// Constraints:
	//    - nullable
	SrcNetwork *string `json:"srcNetwork,omitempty"`

	// SrcPort
	// The source port of AccessControl Rule
	// Constraints:
	//    - nullable
	SrcPort *string `json:"srcPort,omitempty"`
}

func NewSwitchMACLConfigACLRule() *SwitchMACLConfigACLRule {
	m := new(SwitchMACLConfigACLRule)
	return m
}

type SwitchMACLConfigCreateACLConfig struct {
	// AclRule
	// Constraints:
	//    - nullable
	AclRule []*SwitchMACLConfigACLRule `json:"aclRule,omitempty" validate:"omitempty,dive"`

	// AclType
	// ACL Type
	// Constraints:
	//    - nullable
	//    - oneof:[STANDARD,EXTENDED]
	AclType *string `json:"aclType,omitempty" validate:"omitempty,oneof=STANDARD EXTENDED"`

	// FamilyId
	// Family Id
	// Constraints:
	//    - nullable
	FamilyId *string `json:"familyId,omitempty"`

	// GroupId
	// Switch Group Id
	// Constraints:
	//    - nullable
	GroupId *string `json:"groupId,omitempty"`

	// Name
	// ACL Name/ID
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// PushTime
	// Puch Schedule Time
	// Constraints:
	//    - nullable
	PushTime *int `json:"pushTime,omitempty"`

	// PushTimeType
	// Puch Config Type
	// Constraints:
	//    - nullable
	//    - oneof:[NOW,SCHEDULE]
	PushTimeType *string `json:"pushTimeType,omitempty" validate:"omitempty,oneof=NOW SCHEDULE"`

	// SwitchId
	// Switch Id
	// Constraints:
	//    - nullable
	SwitchId *string `json:"switchId,omitempty"`
}

func NewSwitchMACLConfigCreateACLConfig() *SwitchMACLConfigCreateACLConfig {
	m := new(SwitchMACLConfigCreateACLConfig)
	return m
}

type SwitchMACLConfigUpdateACLConfig struct {
	// AclRule
	// Constraints:
	//    - nullable
	AclRule []*SwitchMACLConfigACLRule `json:"aclRule,omitempty" validate:"omitempty,dive"`

	// AclType
	// ACL Type
	// Constraints:
	//    - nullable
	//    - oneof:[STANDARD,EXTENDED]
	AclType *string `json:"aclType,omitempty" validate:"omitempty,oneof=STANDARD EXTENDED"`

	// PushTime
	// Puch Schedule Time
	// Constraints:
	//    - nullable
	PushTime *int `json:"pushTime,omitempty"`

	// PushTimeType
	// Puch Config Type
	// Constraints:
	//    - nullable
	//    - oneof:[NOW,SCHEDULE]
	PushTimeType *string `json:"pushTimeType,omitempty" validate:"omitempty,oneof=NOW SCHEDULE"`
}

func NewSwitchMACLConfigUpdateACLConfig() *SwitchMACLConfigUpdateACLConfig {
	m := new(SwitchMACLConfigUpdateACLConfig)
	return m
}
