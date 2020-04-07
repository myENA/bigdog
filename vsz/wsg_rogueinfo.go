package vsz

// API Version: v9_0

import (
	"encoding/json"
)

type WSGRogueInfo struct {
	// Channel
	// Channel of the rogue AP
	// Constraints:
	//    - nullable
	Channel *int `json:"channel,omitempty"`

	// Classification
	// Rogue classification by policy
	// Constraints:
	//    - nullable
	Classification *string `json:"classification,omitempty"`

	// DetectedByAP
	// The list of the AP found this Rogue AP.
	// Constraints:
	//    - nullable
	DetectedByAP []*WSGAPInfo `json:"detectedByAP,omitempty" validate:"omitempty,dive"`

	// Encryption
	// Encryption of the rogue AP
	// Constraints:
	//    - nullable
	Encryption *string `json:"encryption,omitempty"`

	// LastDetected
	// Timestamp of the rogue AP
	// Constraints:
	//    - nullable
	LastDetected *int `json:"lastDetected,omitempty"`

	// MatchResult
	// What policy and rule matched when system doing classification by rogue policy
	// Constraints:
	//    - nullable
	MatchResult *string `json:"matchResult,omitempty"`

	// Radio
	// Radio of the rogue AP
	// Constraints:
	//    - nullable
	Radio *string `json:"radio,omitempty"`

	// RogueAPMac
	// Constraints:
	//    - nullable
	RogueAPMac *WSGCommonMac `json:"rogueAPMac,omitempty"`

	// RogueMac
	// Constraints:
	//    - nullable
	RogueMac *WSGCommonMac `json:"rogueMac,omitempty"`

	// Ssid
	// SSID of the rogue AP
	// Constraints:
	//    - nullable
	Ssid *string `json:"ssid,omitempty"`

	// Type
	// Type of the rogue AP
	// Constraints:
	//    - nullable
	Type *string `json:"type,omitempty"`
}

func NewWSGRogueInfo() *WSGRogueInfo {
	m := new(WSGRogueInfo)
	return m
}

type WSGRogueInfoList struct {
	// Extra
	// Any additional response data.
	// Constraints:
	//    - nullable
	Extra *WSGRogueInfoListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first Rogue AP returned out of the complete Rogue AP list
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Rogue AP after the current displayed list
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGRogueInfo `json:"list,omitempty" validate:"omitempty,dive"`

	// RawDataTotalCount
	// Total Rogue APs count.
	// Constraints:
	//    - nullable
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Rogue APs count in this response.
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGRogueInfoList() *WSGRogueInfoList {
	m := new(WSGRogueInfoList)
	return m
}

// WSGRogueInfoListExtraType
//
// Any additional response data.
// Constraints:
//    - nullable
type WSGRogueInfoListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGRogueInfoListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGRogueInfoListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGRogueInfoListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGRogueInfoListExtraType() *WSGRogueInfoListExtraType {
	m := new(WSGRogueInfoListExtraType)
	return m
}
