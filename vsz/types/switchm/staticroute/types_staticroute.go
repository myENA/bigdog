package staticroute

// API Version: v8_1

import (
	"encoding/json"
)

type CreateStaticRoute struct {
	// AdminDistance
	// Admin Distance
	AdminDistance *int `json:"adminDistance,omitempty"`

	// DestinationIP
	// Destination IP
	DestinationIP *string `json:"destinationIp,omitempty"`

	// FamilyID
	// Family Id
	FamilyID *string `json:"familyId,omitempty"`

	// GroupID
	// Switch Group Id
	GroupID *string `json:"groupId,omitempty"`

	// NextHop
	// Next Hop
	NextHop *string `json:"nextHop,omitempty"`

	// PushTime
	// Puch Schedule Time
	PushTime *int `json:"pushTime,omitempty"`

	// PushTimeType
	// Puch Config Type
	PushTimeType *string `json:"pushTimeType,omitempty" validate:"oneof=NOW SCHEDULE"`

	// SwitchID
	// Switch Id
	SwitchID *string `json:"switchId,omitempty"`
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

type StaticRoute struct {
	// AdminDistance
	// Admin Distance
	AdminDistance *string `json:"adminDistance,omitempty"`

	// CreatedTime
	// The create time of the Static Route
	CreatedTime *int `json:"createdTime,omitempty"`

	// DestinationIP
	// Destination IP
	DestinationIP *string `json:"destinationIp,omitempty"`

	// FamilyID
	// Family Id
	FamilyID *string `json:"familyId,omitempty"`

	// GroupID
	// Switch Group Id
	GroupID *string `json:"groupId,omitempty"`

	// ID
	// Static Route Id
	ID *string `json:"id,omitempty"`

	// NextHop
	// Next Hop
	NextHop *string `json:"nextHop,omitempty"`

	// PushTime
	// Puch Schedule Time
	PushTime *int `json:"pushTime,omitempty"`

	// PushTimeType
	// Puch Config Type
	PushTimeType *string `json:"pushTimeType,omitempty" validate:"oneof=NOW SCHEDULE"`

	// SwitchID
	// Switch Id
	SwitchID *string `json:"switchId,omitempty"`

	// UpdatedTime
	// The modify time of the Static Route
	UpdatedTime *int `json:"updatedTime,omitempty"`
}

type StaticRoutesQueryResult struct {
	// Extra
	// Any additional response data
	Extra *StaticRoutesQueryResultExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first Static Route returned out of the complete Static Route list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Static Routes after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*StaticRoute `json:"list,omitempty"`

	// RawDataTotalCount
	// Total Static Route count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Static Route Servers count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

// StaticRoutesQueryResultExtraType
//
// Any additional response data
type StaticRoutesQueryResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *StaticRoutesQueryResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = StaticRoutesQueryResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *StaticRoutesQueryResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

type UpdateStaticRoute struct {
	// AdminDistance
	// Admin Distance
	AdminDistance *int `json:"adminDistance,omitempty"`

	// DestinationIP
	// Destination IP
	DestinationIP *string `json:"destinationIp,omitempty"`

	// NextHop
	// Next Hop
	NextHop *string `json:"nextHop,omitempty"`

	// PushTime
	// Puch Schedule Time
	PushTime *int `json:"pushTime,omitempty"`

	// PushTimeType
	// Puch Config Type
	PushTimeType *string `json:"pushTimeType,omitempty" validate:"oneof=NOW SCHEDULE"`
}
