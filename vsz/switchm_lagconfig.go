package vsz

// API Version: v9_0

type SwitchMLagConfigCreate struct {
	// GroupId
	// Switch Group Id
	// Constraints:
	//    - nullable
	GroupId *string `json:"groupId,omitempty"`

	// Name
	// The Name of LAG Config
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// Ports
	// The Switch ports would like to join together
	// Constraints:
	//    - nullable
	Ports []string `json:"ports,omitempty" validate:"omitempty,dive"`

	// SwitchId
	// Switch Id
	// Constraints:
	//    - nullable
	SwitchId *string `json:"switchId,omitempty"`

	// Type
	// The LAG Types in STATIC or DYNAMIC
	// Constraints:
	//    - nullable
	//    - oneof:[STATIC,DYNAMIC]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=STATIC DYNAMIC"`
}

func NewSwitchMLagConfigCreate() *SwitchMLagConfigCreate {
	m := new(SwitchMLagConfigCreate)
	return m
}

type SwitchMLagConfigCreateResult interface{}

func MakeSwitchMLagConfigCreateResult() SwitchMLagConfigCreateResult {
	m := new(SwitchMLagConfigCreateResult)
	return m
}

type SwitchMLagConfig struct {
	// CreatedTime
	// The LAG Types in STATIC or DYNAMIC
	// Constraints:
	//    - nullable
	CreatedTime *int `json:"createdTime,omitempty"`

	// GroupId
	// The ID of  Switch Group
	// Constraints:
	//    - nullable
	GroupId *string `json:"groupId,omitempty"`

	// Id
	// The ID of LAG Config
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// The Name of LAG Config
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// Ports
	// The Switch ports would like to join together
	// Constraints:
	//    - nullable
	Ports []string `json:"ports,omitempty" validate:"omitempty,dive"`

	// SwitchId
	// The ID of Switch
	// Constraints:
	//    - nullable
	SwitchId *string `json:"switchId,omitempty"`

	// Type
	// The LAG Types in STATIC or DYNAMIC
	// Constraints:
	//    - nullable
	//    - oneof:[STATIC,DYNAMIC]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=STATIC DYNAMIC"`

	// UpdatedTime
	// The LAG Types in STATIC or DYNAMIC
	// Constraints:
	//    - nullable
	UpdatedTime *int `json:"updatedTime,omitempty"`
}

func NewSwitchMLagConfig() *SwitchMLagConfig {
	m := new(SwitchMLagConfig)
	return m
}

type SwitchMLagConfigList struct {
	// FirstIndex
	// Index of the first config returned out of the complete list
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more configs after the current displayed list
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// The list of configs in this response
	// Constraints:
	//    - nullable
	List []*SwitchMLagConfig `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Total configs count in this response
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMLagConfigList() *SwitchMLagConfigList {
	m := new(SwitchMLagConfigList)
	return m
}

type SwitchMLagConfigModify struct {
	// Name
	// The Name of LAG Config
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// Ports
	// The Switch ports would like to join together
	// Constraints:
	//    - nullable
	Ports []string `json:"ports,omitempty" validate:"omitempty,dive"`
}

func NewSwitchMLagConfigModify() *SwitchMLagConfigModify {
	m := new(SwitchMLagConfigModify)
	return m
}
