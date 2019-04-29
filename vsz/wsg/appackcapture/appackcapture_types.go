package appackcapture

// API Version: v8_0

type ApPacketCaptureReq struct {
	CaptureInterface   *string  `json:"captureInterface,omitempty"`
	HostIP             *string  `json:"hostIp,omitempty"`
	IncludedFrameTypes []string `json:"includedFrameTypes,omitempty"`
	IncludedMac        *string  `json:"includedMac,omitempty"`
}

type ApPacketCaptureRes struct {
	ApMac              *string  `json:"apMac,omitempty"`
	CaptureInterface   *string  `json:"captureInterface,omitempty"`
	CaptureMode        *string  `json:"captureMode,omitempty"`
	CaptureState       *string  `json:"captureState,omitempty"`
	HostIP             *string  `json:"hostIp,omitempty"`
	IncludedFrameTypes []string `json:"includedFrameTypes,omitempty"`
	IncludedMac        *string  `json:"includedMac,omitempty"`
}
