package meshnodeinfo

// API Version: v8_1

import (
	"encoding/json"
)

type HelperZoneInfo struct {
	HelperAPZoneId *string `json:"helperAPZoneId,omitempty"`

	HelperAPZoneName *string `json:"helperAPZoneName,omitempty"`
}

type MeshNodeInfo struct {
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

	HelperZoneInfo []*HelperZoneInfo `json:"helperZoneInfo,omitempty"`

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

type MeshNodeInfoArray []*MeshNodeInfo

type MeshNodeInfoList struct {
	// Extra
	// Any additional response data.
	Extra *MeshNodeInfoListExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first MeshNodeInfo returned out of the complete Rogue AP list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more MeshNodeInfo after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*MeshNodeInfo `json:"list,omitempty"`

	// RawDataTotalCount
	// MeshNodeInfos count.
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// MeshNodeInfos count in this response.
	TotalCount *int `json:"totalCount,omitempty"`
}

// MeshNodeInfoListExtraType
//
// Any additional response data.
type MeshNodeInfoListExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *MeshNodeInfoListExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = MeshNodeInfoListExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *MeshNodeInfoListExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

type UpdateAPZeroTouch struct {
	ApMac *string `json:"apMac,omitempty"`

	HelperZoneId *string `json:"helperZoneId,omitempty"`

	SerialNumber *string `json:"serialNumber,omitempty"`

	Status *string `json:"status,omitempty" validate:"oneof=Accept Reject"`
}
