package vsz

// API Version: v9_0

type SwitchMLagconfigCreate struct {
	// GroupId
	// Switch Group Id
	GroupId *string `json:"groupId,omitempty"`

	// Name
	// The Name of LAG Config
	Name *string `json:"name,omitempty"`

	// Ports
	// The Switch ports would like to join together
	Ports []string `json:"ports,omitempty"`

	// SwitchId
	// Switch Id
	SwitchId *string `json:"switchId,omitempty"`

	// Type
	// The LAG Types in STATIC or DYNAMIC
	// Constraints:
	//    - oneof:[STATIC,DYNAMIC]
	Type *string `json:"type,omitempty"`
}

func NewSwitchMLagconfigCreate() *SwitchMLagconfigCreate {
	m := new(SwitchMLagconfigCreate)
	return m
}

type SwitchMLagconfigCreateResult interface{}

func MakeSwitchMLagconfigCreateResult() SwitchMLagconfigCreateResult {
	m := new(SwitchMLagconfigCreateResult)
	return m
}

type SwitchMLagconfig struct {
	// CreatedTime
	// The LAG Types in STATIC or DYNAMIC
	CreatedTime *int `json:"createdTime,omitempty"`

	// GroupId
	// The ID of  Switch Group
	GroupId *string `json:"groupId,omitempty"`

	// Id
	// The ID of LAG Config
	Id *string `json:"id,omitempty"`

	// Name
	// The Name of LAG Config
	Name *string `json:"name,omitempty"`

	// Ports
	// The Switch ports would like to join together
	Ports []string `json:"ports,omitempty"`

	// SwitchId
	// The ID of Switch
	SwitchId *string `json:"switchId,omitempty"`

	// Type
	// The LAG Types in STATIC or DYNAMIC
	// Constraints:
	//    - oneof:[STATIC,DYNAMIC]
	Type *string `json:"type,omitempty"`

	// UpdatedTime
	// The LAG Types in STATIC or DYNAMIC
	UpdatedTime *int `json:"updatedTime,omitempty"`
}

func NewSwitchMLagconfig() *SwitchMLagconfig {
	m := new(SwitchMLagconfig)
	return m
}

type SwitchMLagconfigList struct {
	// FirstIndex
	// Index of the first config returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more configs after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// The list of configs in this response
	List []*SwitchMLagconfig `json:"list,omitempty"`

	// TotalCount
	// Total configs count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMLagconfigList() *SwitchMLagconfigList {
	m := new(SwitchMLagconfigList)
	return m
}

type SwitchMLagconfigModify struct {
	// Name
	// The Name of LAG Config
	Name *string `json:"name,omitempty"`

	// Ports
	// The Switch ports would like to join together
	Ports []string `json:"ports,omitempty"`
}

func NewSwitchMLagconfigModify() *SwitchMLagconfigModify {
	m := new(SwitchMLagconfigModify)
	return m
}
