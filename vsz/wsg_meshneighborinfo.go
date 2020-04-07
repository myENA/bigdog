package vsz

// API Version: v9_0

import (
	"encoding/json"
)

type WSGMeshNeighborInfo struct {
	// ApFirmware
	// The firmware of the neighbor AP
	// Constraints:
	//    - nullable
	ApFirmware *string `json:"apFirmware,omitempty"`

	// ApMac
	// The MAC of the neighbor AP
	// Constraints:
	//    - nullable
	ApMac *string `json:"apMac,omitempty"`

	// ApModel
	// The model of the neighbor AP
	// Constraints:
	//    - nullable
	ApModel *string `json:"apModel,omitempty"`

	// ApName
	// The name of the neighbor AP
	// Constraints:
	//    - nullable
	ApName *string `json:"apName,omitempty"`

	// Channel
	// The channel of the neighbor AP
	// Constraints:
	//    - nullable
	Channel *string `json:"channel,omitempty"`

	// ConnectionStatus
	// The connection status of the neighbor AP
	// Constraints:
	//    - nullable
	ConnectionStatus *string `json:"connectionStatus,omitempty"`

	// ExternalIPAddress
	// The external IP of the neighbor AP
	// Constraints:
	//    - nullable
	ExternalIPAddress *string `json:"externalIPAddress,omitempty"`

	// IpAddress
	// The IP of the neighbor AP
	// Constraints:
	//    - nullable
	IpAddress *string `json:"ipAddress,omitempty"`

	// Rssi
	// The RSSI of the neighbor AP
	// Constraints:
	//    - nullable
	Rssi *int `json:"rssi,omitempty"`

	// ZoneName
	// The zone name of the neighbor AP
	// Constraints:
	//    - nullable
	ZoneName *string `json:"zoneName,omitempty"`
}

func NewWSGMeshNeighborInfo() *WSGMeshNeighborInfo {
	m := new(WSGMeshNeighborInfo)
	return m
}

type WSGMeshNeighborInfoList struct {
	// Extra
	// Any additional response data.
	// Constraints:
	//    - nullable
	Extra *WSGMeshNeighborInfoListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first MeshNeighborInfo returned out of the complete Rogue AP list
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more MeshNeighborInfo after the current displayed list
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGMeshNeighborInfo `json:"list,omitempty" validate:"omitempty,dive"`

	// RawDataTotalCount
	// MeshNeighborInfos count.
	// Constraints:
	//    - nullable
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// MeshNeighborInfos count in this response.
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGMeshNeighborInfoList() *WSGMeshNeighborInfoList {
	m := new(WSGMeshNeighborInfoList)
	return m
}

// WSGMeshNeighborInfoListExtraType
//
// Any additional response data.
// Constraints:
//    - nullable
type WSGMeshNeighborInfoListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGMeshNeighborInfoListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGMeshNeighborInfoListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGMeshNeighborInfoListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGMeshNeighborInfoListExtraType() *WSGMeshNeighborInfoListExtraType {
	m := new(WSGMeshNeighborInfoListExtraType)
	return m
}
