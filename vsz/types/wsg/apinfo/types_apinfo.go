package apinfo

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type ApInfo struct {
	ApMac *common.Mac `json:"apMac,omitempty"`

	ApName *ApName `json:"apName,omitempty"`

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
