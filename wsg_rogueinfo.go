package bigdog

// API Version: v9_0

// WSGRogueInfo
//
// Definition: rogueInfo_rogueInfo
type WSGRogueInfo struct {
	// Channel
	// Channel of the rogue AP
	Channel *int `json:"channel,omitempty"`

	// Classification
	// Rogue classification by policy
	Classification *string `json:"classification,omitempty"`

	// DetectedByAP
	// The list of the AP found this Rogue AP.
	DetectedByAP []*WSGAPInfo `json:"detectedByAP,omitempty"`

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

	RogueAPMac *WSGCommonMac `json:"rogueAPMac,omitempty"`

	RogueMac *WSGCommonMac `json:"rogueMac,omitempty"`

	// Ssid
	// SSID of the rogue AP
	Ssid *string `json:"ssid,omitempty"`

	// Type
	// Type of the rogue AP
	Type *string `json:"type,omitempty"`
}

func NewWSGRogueInfo() *WSGRogueInfo {
	m := new(WSGRogueInfo)
	return m
}

// WSGRogueInfoList
//
// Definition: rogueInfo_rogueInfoList
type WSGRogueInfoList struct {
	// Extra
	// Any additional response data.
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first Rogue AP returned out of the complete Rogue AP list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Rogue AP after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGRogueInfo `json:"list,omitempty"`

	// RawDataTotalCount
	// Total Rogue APs count.
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Rogue APs count in this response.
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGRogueInfoList() *WSGRogueInfoList {
	m := new(WSGRogueInfoList)
	return m
}
