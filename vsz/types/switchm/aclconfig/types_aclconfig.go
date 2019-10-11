package aclconfig

// API Version: v8_1

import (
	"encoding/json"
)

type ACLConfig struct {
	AclRule []*ACLRule `json:"aclRule,omitempty"`

	// AclType
	// ACL Type
	// Constraints:
	//    - nullable
	//    - oneof:[STANDARD,EXTENDED]
	AclType *string `json:"aclType,omitempty" validate:"omitempty,oneof=STANDARD EXTENDED"`

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
	// Constraints:
	//    - nullable
	//    - oneof:[NOW,SCHEDULE]
	PushTimeType *string `json:"pushTimeType,omitempty" validate:"omitempty,oneof=NOW SCHEDULE"`

	// SwitchId
	// Switch Id
	SwitchId *string `json:"switchId,omitempty"`

	// UpdatedTime
	// The modify time of the AccessControl
	UpdatedTime *int `json:"updatedTime,omitempty"`
}

func NewACLConfig() *ACLConfig {
	aCLConfigType := new(ACLConfig)
	return aCLConfigType
}

func NewDefaultACLConfig() *ACLConfig {
	aCLConfigType := new(ACLConfig)
	return aCLConfigType
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

func NewACLConfigsQueryResult() *ACLConfigsQueryResult {
	aCLConfigsQueryResultType := new(ACLConfigsQueryResult)
	return aCLConfigsQueryResultType
}

func NewDefaultACLConfigsQueryResult() *ACLConfigsQueryResult {
	aCLConfigsQueryResultType := new(ACLConfigsQueryResult)
	return aCLConfigsQueryResultType
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

func NewACLConfigsQueryResultExtraType() *ACLConfigsQueryResultExtraType {
	aCLConfigsQueryResultExtraTypeType := new(ACLConfigsQueryResultExtraType)
	return aCLConfigsQueryResultExtraTypeType
}

func NewDefaultACLConfigsQueryResultExtraType() *ACLConfigsQueryResultExtraType {
	aCLConfigsQueryResultExtraTypeType := new(ACLConfigsQueryResultExtraType)
	return aCLConfigsQueryResultExtraTypeType
}

type ACLRule struct {
	// Action
	// The action of AccessControl Rule
	// Constraints:
	//    - nullable
	//    - oneof:[PERMIT,DENY]
	Action *string `json:"action,omitempty" validate:"omitempty,oneof=PERMIT DENY"`

	// DestNetwork
	// The destination network of AccessControl Rule
	DestNetwork *string `json:"destNetwork,omitempty"`

	// DestPort
	// The destination port of AccessControl Rule
	DestPort *string `json:"destPort,omitempty"`

	// Protocol
	// The protocol of AccessControl Rule
	// Constraints:
	//    - nullable
	//    - oneof:[IP,TCP,UDP]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=IP TCP UDP"`

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

func NewACLRule() *ACLRule {
	aCLRuleType := new(ACLRule)
	return aCLRuleType
}

func NewDefaultACLRule() *ACLRule {
	aCLRuleType := new(ACLRule)
	return aCLRuleType
}

type CreateACLConfig struct {
	AclRule []*ACLRule `json:"aclRule,omitempty"`

	// AclType
	// ACL Type
	// Constraints:
	//    - nullable
	//    - oneof:[STANDARD,EXTENDED]
	AclType *string `json:"aclType,omitempty" validate:"omitempty,oneof=STANDARD EXTENDED"`

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
	// Constraints:
	//    - nullable
	//    - oneof:[NOW,SCHEDULE]
	PushTimeType *string `json:"pushTimeType,omitempty" validate:"omitempty,oneof=NOW SCHEDULE"`

	// SwitchId
	// Switch Id
	SwitchId *string `json:"switchId,omitempty"`
}

func NewCreateACLConfig() *CreateACLConfig {
	createACLConfigType := new(CreateACLConfig)
	return createACLConfigType
}

func NewDefaultCreateACLConfig() *CreateACLConfig {
	createACLConfigType := new(CreateACLConfig)
	return createACLConfigType
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

func NewEmptyResult() *EmptyResult {
	emptyResultType := new(EmptyResult)
	return emptyResultType
}

func NewDefaultEmptyResult() *EmptyResult {
	emptyResultType := new(EmptyResult)
	return emptyResultType
}

type UpdateACLConfig struct {
	AclRule []*ACLRule `json:"aclRule,omitempty"`

	// AclType
	// ACL Type
	// Constraints:
	//    - nullable
	//    - oneof:[STANDARD,EXTENDED]
	AclType *string `json:"aclType,omitempty" validate:"omitempty,oneof=STANDARD EXTENDED"`

	// PushTime
	// Puch Schedule Time
	PushTime *int `json:"pushTime,omitempty"`

	// PushTimeType
	// Puch Config Type
	// Constraints:
	//    - nullable
	//    - oneof:[NOW,SCHEDULE]
	PushTimeType *string `json:"pushTimeType,omitempty" validate:"omitempty,oneof=NOW SCHEDULE"`
}

func NewUpdateACLConfig() *UpdateACLConfig {
	updateACLConfigType := new(UpdateACLConfig)
	return updateACLConfigType
}

func NewDefaultUpdateACLConfig() *UpdateACLConfig {
	updateACLConfigType := new(UpdateACLConfig)
	return updateACLConfigType
}
