package bigdog

// API Version: v9_1

import (
	"errors"
	"io"
)

// WSGMeshNodeInfoHelperZoneInfo
//
// Definition: meshNodeInfo_helperZoneInfo
type WSGMeshNodeInfoHelperZoneInfo struct {
	HelperAPZoneId *string `json:"helperAPZoneId,omitempty"`

	HelperAPZoneName *string `json:"helperAPZoneName,omitempty"`
}

func NewWSGMeshNodeInfoHelperZoneInfo() *WSGMeshNodeInfoHelperZoneInfo {
	m := new(WSGMeshNodeInfoHelperZoneInfo)
	return m
}

// WSGMeshNodeInfo
//
// Definition: meshNodeInfo_meshNodeInfo
type WSGMeshNodeInfo struct {
	// ApMac
	// The MAC of the mesh AP
	ApMac *string `json:"apMac,omitempty"`

	// ApModel
	// The model of the mesh AP
	ApModel *string `json:"apModel,omitempty"`

	// ApName
	// The name of the mesh AP
	ApName *string `json:"apName,omitempty"`

	// Channel
	// The channel of the mesh AP
	Channel *string `json:"channel,omitempty"`

	// ClientCount
	// The count of clients of the mesh AP
	ClientCount *int `json:"clientCount,omitempty"`

	// DownlinkRssi
	// The downlinkRssi of the mesh AP
	DownlinkRssi *int `json:"downlinkRssi,omitempty"`

	// ExternalIPAddress
	// The external IP of the mesh AP
	ExternalIPAddress *string `json:"externalIPAddress,omitempty"`

	// HasDownLink
	// The hasDownLink of the mesh AP
	HasDownLink *bool `json:"hasDownLink,omitempty"`

	HelperZoneInfo []*WSGMeshNodeInfoHelperZoneInfo `json:"helperZoneInfo,omitempty"`

	// Hops
	// The hop count of this mesh AP
	Hops *int `json:"hops,omitempty"`

	// IpAddress
	// The IP of the mesh AP
	IpAddress *int `json:"ipAddress,omitempty"`

	// MeshRole
	// The Role of the mesh AP
	MeshRole *string `json:"meshRole,omitempty"`

	// UplinkRssi
	// The uplinkRssi of the mesh AP
	UplinkRssi *int `json:"uplinkRssi,omitempty"`
}

func NewWSGMeshNodeInfo() *WSGMeshNodeInfo {
	m := new(WSGMeshNodeInfo)
	return m
}

// WSGMeshNodeInfoArray
//
// Definition: meshNodeInfo_meshNodeInfoArray
type WSGMeshNodeInfoArray []*WSGMeshNodeInfo

type WSGMeshNodeInfoArrayAPIResponse struct {
	*RawAPIResponse
	Data WSGMeshNodeInfoArray
}

func newWSGMeshNodeInfoArrayAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGMeshNodeInfoArrayAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGMeshNodeInfoArrayAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := make(WSGMeshNodeInfoArray, 0)
	if err := r.doHydrate(&data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func MakeWSGMeshNodeInfoArray() WSGMeshNodeInfoArray {
	m := make(WSGMeshNodeInfoArray, 0)
	return m
}

// WSGMeshNodeInfoList
//
// Definition: meshNodeInfo_meshNodeInfoList
type WSGMeshNodeInfoList struct {
	// Extra
	// Any additional response data.
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first MeshNodeInfo returned out of the complete Rogue AP list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more MeshNodeInfo after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGMeshNodeInfo `json:"list,omitempty"`

	// RawDataTotalCount
	// MeshNodeInfos count.
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// MeshNodeInfos count in this response.
	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGMeshNodeInfoListAPIResponse struct {
	*RawAPIResponse
	Data *WSGMeshNodeInfoList
}

func newWSGMeshNodeInfoListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGMeshNodeInfoListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGMeshNodeInfoListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGMeshNodeInfoList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGMeshNodeInfoList() *WSGMeshNodeInfoList {
	m := new(WSGMeshNodeInfoList)
	return m
}

// WSGMeshNodeInfoUpdateAPZeroTouch
//
// Definition: meshNodeInfo_updateAPZeroTouch
type WSGMeshNodeInfoUpdateAPZeroTouch struct {
	ApMac *string `json:"apMac,omitempty"`

	HelperZoneId *string `json:"helperZoneId,omitempty"`

	SerialNumber *string `json:"serialNumber,omitempty"`

	// Status
	// Constraints:
	//    - oneof:[Accept,Reject]
	Status *string `json:"status,omitempty"`
}

func NewWSGMeshNodeInfoUpdateAPZeroTouch() *WSGMeshNodeInfoUpdateAPZeroTouch {
	m := new(WSGMeshNodeInfoUpdateAPZeroTouch)
	return m
}
