package apinfo

// API Version: v8_0

type ApInfo struct {
	ApMac *string `json:"apMac,omitempty"`

	ApName *string `json:"apName,omitempty"`

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

type ApName string
