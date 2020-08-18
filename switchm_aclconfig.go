package bigdog

// API Version: v9_1

// SwitchMACLConfig
//
// Definition: aclConfig_ACLConfig
type SwitchMACLConfig struct {
	AclRule []*SwitchMACLConfigACLRule `json:"aclRule,omitempty"`

	// AclType
	// ACL Type
	// Constraints:
	//    - oneof:[STANDARD,EXTENDED]
	AclType *string `json:"aclType,omitempty"`

	AltoId *string `json:"altoId,omitempty"`

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
	//    - oneof:[NOW,SCHEDULE]
	PushTimeType *string `json:"pushTimeType,omitempty"`

	// SwitchId
	// Switch Id
	SwitchId *string `json:"switchId,omitempty"`

	// UpdatedTime
	// The modify time of the AccessControl
	UpdatedTime *int `json:"updatedTime,omitempty"`
}

func NewSwitchMACLConfig() *SwitchMACLConfig {
	m := new(SwitchMACLConfig)
	return m
}

// SwitchMACLConfigsQueryResult
//
// Definition: aclConfig_ACLConfigsQueryResult
type SwitchMACLConfigsQueryResult struct {
	// Extra
	// Any additional response data
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first AccessControl returned out of the complete AccessControl list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more AccessControl after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMACLConfig `json:"list,omitempty"`

	// RawDataTotalCount
	// Total AccessControl count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total AccessControl count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMACLConfigsQueryResult() *SwitchMACLConfigsQueryResult {
	m := new(SwitchMACLConfigsQueryResult)
	return m
}

// SwitchMACLConfigACLRule
//
// Definition: aclConfig_ACLRule
type SwitchMACLConfigACLRule struct {
	// Action
	// The action of AccessControl Rule
	// Constraints:
	//    - oneof:[PERMIT,DENY]
	Action *string `json:"action,omitempty"`

	// DestNetwork
	// The destination network of AccessControl Rule
	DestNetwork *string `json:"destNetwork,omitempty"`

	// DestPort
	// The destination port of AccessControl Rule
	DestPort *string `json:"destPort,omitempty"`

	// DscpMarking
	// The DSCP value to be used for marking of AccessControl Rule
	DscpMarking *int `json:"dscpMarking,omitempty"`

	// DscpMatching
	// The DSCP value for matching of AccessControl Rule
	DscpMatching *int `json:"dscpMatching,omitempty"`

	// InternalPriority
	// The QoS priority of AccessControl Rule
	InternalPriority *int `json:"internalPriority,omitempty"`

	// Protocol
	// The protocol of AccessControl Rule
	// Constraints:
	//    - oneof:[IP,TCP,UDP]
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

func NewSwitchMACLConfigACLRule() *SwitchMACLConfigACLRule {
	m := new(SwitchMACLConfigACLRule)
	return m
}

// SwitchMACLConfigCreateACLConfig
//
// Definition: aclConfig_createACLConfig
type SwitchMACLConfigCreateACLConfig struct {
	AclRule []*SwitchMACLConfigACLRule `json:"aclRule,omitempty"`

	// AclType
	// ACL Type
	// Constraints:
	//    - oneof:[STANDARD,EXTENDED]
	AclType *string `json:"aclType,omitempty"`

	AltoId *string `json:"altoId,omitempty"`

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
	//    - oneof:[NOW,SCHEDULE]
	PushTimeType *string `json:"pushTimeType,omitempty"`

	// SwitchId
	// Switch Id
	SwitchId *string `json:"switchId,omitempty"`
}

func NewSwitchMACLConfigCreateACLConfig() *SwitchMACLConfigCreateACLConfig {
	m := new(SwitchMACLConfigCreateACLConfig)
	return m
}

// SwitchMACLConfigUpdateACLConfig
//
// Definition: aclConfig_updateACLConfig
type SwitchMACLConfigUpdateACLConfig struct {
	AclRule []*SwitchMACLConfigACLRule `json:"aclRule,omitempty"`

	// AclType
	// ACL Type
	// Constraints:
	//    - oneof:[STANDARD,EXTENDED]
	AclType *string `json:"aclType,omitempty"`

	AltoId *string `json:"altoId,omitempty"`

	// PushTime
	// Puch Schedule Time
	PushTime *int `json:"pushTime,omitempty"`

	// PushTimeType
	// Puch Config Type
	// Constraints:
	//    - oneof:[NOW,SCHEDULE]
	PushTimeType *string `json:"pushTimeType,omitempty"`
}

func NewSwitchMACLConfigUpdateACLConfig() *SwitchMACLConfigUpdateACLConfig {
	m := new(SwitchMACLConfigUpdateACLConfig)
	return m
}
