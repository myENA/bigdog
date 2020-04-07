package vsz

// API Version: v9_0

type WSGAPPackCaptureApPacketCaptureReq struct {
	// CaptureInterface
	// Capture interface
	// Constraints:
	//    - required
	//    - oneof:[RADIO24,RADIO50,ETH0,ETH1,ETH2,ETH3,ETH4,ETH5,ETH6,ETH7]
	CaptureInterface *string `json:"captureInterface" validate:"required,oneof=RADIO24 RADIO50 ETH0 ETH1 ETH2 ETH3 ETH4 ETH5 ETH6 ETH7"`

	// HostIp
	// Wireshark host IP
	// Constraints:
	//    - nullable
	HostIp *string `json:"hostIp,omitempty"`

	// IncludedFrameTypes
	// Frame type filter
	// Constraints:
	//    - nullable
	IncludedFrameTypes []string `json:"includedFrameTypes,omitempty" validate:"omitempty,dive"`

	// IncludedMac
	// MAC filter
	// Constraints:
	//    - nullable
	IncludedMac *string `json:"includedMac,omitempty"`
}

func NewWSGAPPackCaptureApPacketCaptureReq() *WSGAPPackCaptureApPacketCaptureReq {
	m := new(WSGAPPackCaptureApPacketCaptureReq)
	return m
}

type WSGAPPackCaptureApPacketCaptureRes struct {
	// ApMac
	// AP MAC address
	// Constraints:
	//    - nullable
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
	// Constraints:
	//    - nullable
	CaptureState *string `json:"captureState,omitempty"`

	// HostIp
	// Wireshark host IP
	// Constraints:
	//    - nullable
	HostIp *string `json:"hostIp,omitempty"`

	// IncludedFrameTypes
	// Frame type filter
	// Constraints:
	//    - nullable
	IncludedFrameTypes []string `json:"includedFrameTypes,omitempty" validate:"omitempty,dive"`

	// IncludedMac
	// MAC filter
	// Constraints:
	//    - nullable
	IncludedMac *string `json:"includedMac,omitempty"`
}

func NewWSGAPPackCaptureApPacketCaptureRes() *WSGAPPackCaptureApPacketCaptureRes {
	m := new(WSGAPPackCaptureApPacketCaptureRes)
	return m
}
