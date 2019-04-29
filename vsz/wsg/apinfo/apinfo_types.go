package apinfo

// API Version: v8_0

type ApInfo struct {
	ApMac        *string `json:"apMac,omitempty"`
	ApName       *string `json:"apName,omitempty"`
	LastDetected *int    `json:"lastDetected,omitempty"`
	MainDetector *bool   `json:"mainDetector,omitempty"`
	RogueType    *string `json:"rogueType,omitempty"`
	Rssi         *string `json:"rssi,omitempty"`
	ZoneName     *string `json:"zoneName,omitempty"`
}
