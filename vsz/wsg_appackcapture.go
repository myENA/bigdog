package vsz

// API Version: v8_1

type WSGAPPackCaptureApPacketCaptureReq struct {
	// CaptureInterface
	// Capture interface
	// Constraints:
	//    - required
	//    - oneof:[RADIO24,RADIO50,ETH0,ETH1,ETH2,ETH3,ETH4,ETH5,ETH6,ETH7]
	CaptureInterface *string `json:"captureInterface" validate:"required,oneof=RADIO24 RADIO50 ETH0 ETH1 ETH2 ETH3 ETH4 ETH5 ETH6 ETH7"`

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

type WSGAPPackCaptureApPacketCaptureRes struct {
	// ApMac
	// AP MAC address
	ApMac *string `json:"apMac,omitempty"`

	// CaptureInterface
	// Capture interface
	// Constraints:
	//    - nullable
	//    - oneof:[RADIO24,RADIO50,ETH0,ETH1,ETH2,ETH3,ETH4,ETH5,ETH6,ETH7]
	CaptureInterface *string `json:"captureInterface,omitempty" validate:"omitempty,oneof=RADIO24 RADIO50 ETH0 ETH1 ETH2 ETH3 ETH4 ETH5 ETH6 ETH7"`

	// CaptureMode
	// Capture mode
	// Constraints:
	//    - nullable
	//    - oneof:[STREAMING,FILE_CAPTURE]
	CaptureMode *string `json:"captureMode,omitempty" validate:"omitempty,oneof=STREAMING FILE_CAPTURE"`

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
