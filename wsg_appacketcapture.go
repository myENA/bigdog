package ruckus

// API Version: v9_0

type WSGAPPackCaptureApPacketCaptureReq struct {
	// CaptureInterface
	// Capture interface
	// Constraints:
	//    - required
	//    - oneof:[RADIO24,RADIO50,ETH0,ETH1,ETH2,ETH3,ETH4,ETH5,ETH6,ETH7]
	CaptureInterface *string `json:"captureInterface"`

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

func NewWSGAPPackCaptureApPacketCaptureReq() *WSGAPPackCaptureApPacketCaptureReq {
	m := new(WSGAPPackCaptureApPacketCaptureReq)
	return m
}

type WSGAPPackCaptureApPacketCaptureRes struct {
	// ApMac
	// AP MAC address
	ApMac *string `json:"apMac,omitempty"`

	// CaptureInterface
	// Capture interface
	// Constraints:
	//    - oneof:[RADIO24,RADIO50,ETH0,ETH1,ETH2,ETH3,ETH4,ETH5,ETH6,ETH7]
	CaptureInterface *string `json:"captureInterface,omitempty"`

	// CaptureMode
	// Capture mode
	// Constraints:
	//    - oneof:[STREAMING,FILE_CAPTURE]
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

func NewWSGAPPackCaptureApPacketCaptureRes() *WSGAPPackCaptureApPacketCaptureRes {
	m := new(WSGAPPackCaptureApPacketCaptureRes)
	return m
}
