package rogueinfo

// API Version: v8_1

import (
	"encoding/json"
	"github.com/myENA/ruckus-client/vsz/types/wsg/apinfo"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type RogueInfo struct {
	// Channel
	// Channel of the rogue AP
	Channel *int `json:"channel,omitempty"`

	// Classification
	// Rogue classification by policy
	Classification *string `json:"classification,omitempty"`

	// DetectedByAP
	// The list of the AP found this Rogue AP.
	DetectedByAP []*apinfo.ApInfo `json:"detectedByAP,omitempty"`

	// Encryption
	// Encryption of the rogue AP
	Encryption *string `json:"encryption,omitempty"`

	// LastDetected
	// Timestamp of the rogue AP
	LastDetected *int `json:"lastDetected,omitempty"`

	// MatchResult
	// What policy and rule matched when system doing classification by rogue policy
	MatchResult *string `json:"matchResult,omitempty"`

	// Radio
	// Radio of the rogue AP
	Radio *string `json:"radio,omitempty"`

	RogueAPMac *common.Mac `json:"rogueAPMac,omitempty"`

	RogueMac *common.Mac `json:"rogueMac,omitempty"`

	// Ssid
	// SSID of the rogue AP
	Ssid *string `json:"ssid,omitempty"`

	// Type
	// Type of the rogue AP
	Type *string `json:"type,omitempty"`
}

func NewRogueInfo() *RogueInfo {
	rogueInfoType := new(RogueInfo)
	return rogueInfoType
}

func NewDefaultRogueInfo() *RogueInfo {
	rogueInfoType := new(RogueInfo)
	return rogueInfoType
}

type RogueInfoList struct {
	// Extra
	// Any additional response data.
	Extra *RogueInfoListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first Rogue AP returned out of the complete Rogue AP list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Rogue AP after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*RogueInfo `json:"list,omitempty"`

	// RawDataTotalCount
	// Total Rogue APs count.
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Rogue APs count in this response.
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewRogueInfoList() *RogueInfoList {
	rogueInfoListType := new(RogueInfoList)
	return rogueInfoListType
}

func NewDefaultRogueInfoList() *RogueInfoList {
	rogueInfoListType := new(RogueInfoList)
	return rogueInfoListType
}

// RogueInfoListExtraType
//
// Any additional response data.
type RogueInfoListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *RogueInfoListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = RogueInfoListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *RogueInfoListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewRogueInfoListExtraType() *RogueInfoListExtraType {
	rogueInfoListExtraTypeType := new(RogueInfoListExtraType)
	return rogueInfoListExtraTypeType
}

func NewDefaultRogueInfoListExtraType() *RogueInfoListExtraType {
	rogueInfoListExtraTypeType := new(RogueInfoListExtraType)
	return rogueInfoListExtraTypeType
}
