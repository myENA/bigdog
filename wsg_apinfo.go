package ruckus

// API Version: v9_0

type WSGAPInfo struct {
	ApMac **WSGAPInfo `json:"apMac,omitempty"`

	ApName **WSGAPInfo `json:"apName,omitempty"`

	// LastDetected
	// Timestamp of the AP
	LastDetected *int `json:"lastDetected,omitempty"`

	// MainDetector
	// To identify whether this is main instance for UI
	MainDetector *bool `json:"mainDetector,omitempty"`

	// RogueType
	// Rogue type
	RogueType *string `json:"rogueType,omitempty"`

	// Rssi
	// RSSI of the rogue AP
	Rssi *string `json:"rssi,omitempty"`

	// ZoneName
	// Zone name
	ZoneName *string `json:"zoneName,omitempty"`
}

func NewWSGAPInfo() *WSGAPInfo {
	m := new(WSGAPInfo)
	return m
}

type WSGAPInfoApName string

func NewWSGAPInfoApName() *WSGAPInfoApName {
	m := new(WSGAPInfoApName)
	return m
}
