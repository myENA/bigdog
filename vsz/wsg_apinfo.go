package vsz

// API Version: v9_0

type WSGAPInfo struct {
	// ApMac
	// Constraints:
	//    - nullable
	ApMac *WSGCommonMac `json:"apMac,omitempty"`

	// ApName
	// Constraints:
	//    - nullable
	ApName *WSGAPInfoApName `json:"apName,omitempty"`

	// LastDetected
	// Timestamp of the AP
	// Constraints:
	//    - nullable
	LastDetected *int `json:"lastDetected,omitempty"`

	// MainDetector
	// To identify whether this is main instance for UI
	// Constraints:
	//    - nullable
	MainDetector *bool `json:"mainDetector,omitempty"`

	// RogueType
	// Rogue type
	// Constraints:
	//    - nullable
	RogueType *string `json:"rogueType,omitempty"`

	// Rssi
	// RSSI of the rogue AP
	// Constraints:
	//    - nullable
	Rssi *string `json:"rssi,omitempty"`

	// ZoneName
	// Zone name
	// Constraints:
	//    - nullable
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
