package appackcapture

// API Version: v8_1

type ApPacketCaptureReq struct {
	// CaptureInterface
	// Capture interface
	CaptureInterface *string `json:"captureInterface,omitempty"`

	// HostIp
	// Wireshark host IP
	HostIp *string `json:"hostIp,omitempty"`

	// IncludedFrameTypes
	// Frame type filter
	IncludedFrameTypes []string `json:"includedFrameTypes,omitempty"`

	// IncludedMac
	// MAC filter
	IncludedMac *string `json:"includedMac,omitempty"`
}

type ApPacketCaptureRes struct {
	// ApMac
	// AP MAC address
	ApMac *string `json:"apMac,omitempty"`

	// CaptureInterface
	// Capture interface
	CaptureInterface *string `json:"captureInterface,omitempty"`

	// CaptureMode
	// Capture mode
	CaptureMode *string `json:"captureMode,omitempty"`

	// CaptureState
	// Capture state
	CaptureState *string `json:"captureState,omitempty"`

	// HostIp
	// Wireshark host IP
	HostIp *string `json:"hostIp,omitempty"`

	// IncludedFrameTypes
	// Frame type filter
	IncludedFrameTypes []string `json:"includedFrameTypes,omitempty"`

	// IncludedMac
	// MAC filter
	IncludedMac *string `json:"includedMac,omitempty"`
}
