package bigdog

// API Version: v9_1

import (
	"errors"
	"io"
)

// WSGAPPackCaptureApPacketCaptureReq
//
// Definition: apPacketCapture_apPacketCaptureReq
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

// WSGAPPackCaptureApPacketCaptureRes
//
// Definition: apPacketCapture_apPacketCaptureRes
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

type WSGAPPackCaptureApPacketCaptureResAPIResponse struct {
	*RawAPIResponse
	Data *WSGAPPackCaptureApPacketCaptureRes
}

func newWSGAPPackCaptureApPacketCaptureResAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAPPackCaptureApPacketCaptureResAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAPPackCaptureApPacketCaptureResAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAPPackCaptureApPacketCaptureRes)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAPPackCaptureApPacketCaptureRes() *WSGAPPackCaptureApPacketCaptureRes {
	m := new(WSGAPPackCaptureApPacketCaptureRes)
	return m
}
