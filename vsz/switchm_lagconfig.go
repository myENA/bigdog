package vsz

// API Version: v8_1

import (
	"encoding/json"
)

type SwitchMLagConfigCreate struct {
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
	//    - nullable
	//    - oneof:[STATIC,DYNAMIC]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=STATIC DYNAMIC"`
}

type SwitchMLagConfigCreateResult interface{}

type SwitchMLagConfigEmptyResult struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMLagConfigEmptyResult) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMLagConfigEmptyResult{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMLagConfigEmptyResult) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

type SwitchMLagConfig struct {
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
	//    - nullable
	//    - oneof:[STATIC,DYNAMIC]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=STATIC DYNAMIC"`

	// UpdatedTime
	// The LAG Types in STATIC or DYNAMIC
	UpdatedTime *int `json:"updatedTime,omitempty"`
}

type SwitchMLagConfigList struct {
	// FirstIndex
	// Index of the first config returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more configs after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// The list of configs in this response
	List []*SwitchMLagConfig `json:"list,omitempty"`

	// TotalCount
	// Total configs count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchMLagConfigModify struct {
	// Name
	// The Name of LAG Config
	Name *string `json:"name,omitempty"`

	// Ports
	// The Switch ports would like to join together
	Ports []string `json:"ports,omitempty"`
}
