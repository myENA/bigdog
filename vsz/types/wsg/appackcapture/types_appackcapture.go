package appackcapture

// API Version: v8_1

type ApPacketCaptureReq struct {
	// CaptureInterface
	// Capture interface
	CaptureInterface *string `json:"captureInterface,omitempty" validate:"required,oneof=RADIO24 RADIO50 ETH0 ETH1 ETH2 ETH3 ETH4 ETH5 ETH6 ETH7"`

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
	CaptureInterface *string `json:"captureInterface,omitempty" validate:"oneof=RADIO24 RADIO50 ETH0 ETH1 ETH2 ETH3 ETH4 ETH5 ETH6 ETH7"`

	// CaptureMode
	// Capture mode
	CaptureMode *string `json:"captureMode,omitempty" validate:"oneof=STREAMING FILE_CAPTURE"`

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
