package vlanconfig

// API Version: v8_1

import (
	"encoding/json"
)

type CreateVlanConfig struct {
	ArpInspections []*VlanArpInspections `json:"arpInspections,omitempty"`

	// ArpInspectionTrustPort
	// ARP Inspection Trust Port
	ArpInspectionTrustPort *string `json:"arpInspectionTrustPort,omitempty"`

	// EnableArpInspection
	// ARP Inspection Enabled
	EnableArpInspection *bool `json:"enableArpInspection,omitempty"`

	// EnableIpv4DHCPSnooping
	// IPv4 DHCP Snooping Enabled
	EnableIpv4DHCPSnooping *bool `json:"enableIpv4DhcpSnooping,omitempty"`

	// FamilyID
	// Family Id
	FamilyID *string `json:"familyId,omitempty"`

	// GroupID
	// Group Id
	GroupID *string `json:"groupId,omitempty"`

	// IgmpSnooping
	// IGMP Snooping
	IgmpSnooping *string `json:"igmpSnooping,omitempty" validate:"oneof=IST_NONE ACTIVE PASSIVE"`

	// Ipv4DHCPSnoopingTrustPort
	// IPv4 DHCP Snooping Trust Port
	Ipv4DHCPSnoopingTrustPort *string `json:"ipv4DhcpSnoopingTrustPort,omitempty"`

	// MulticastVersion
	// Mutilcast Version
	MulticastVersion *int `json:"multicastVersion,omitempty"`

	// Name
	// Vlan Name
	Name *string `json:"name,omitempty"`

	Ports []*VlanPorts `json:"ports,omitempty"`

	// PushTime
	// Puch Schedule Time
	PushTime *int `json:"pushTime,omitempty"`

	// PushTimeType
	// Puch Config Type
	PushTimeType *string `json:"pushTimeType,omitempty" validate:"oneof=NOW SCHEDULE"`

	// SpanningTree
	// Spanning Tree
	SpanningTree *string `json:"spanningTree,omitempty" validate:"oneof=STT_NONE STP RSTP"`

	// SpanningTreePriority
	// Spanning Tree Priority
	SpanningTreePriority *int `json:"spanningTreePriority,omitempty"`

	// SwitchID
	// Switch Id
	SwitchID *string `json:"switchId,omitempty"`

	// VlanID
	// Vlan Id
	VlanID *int `json:"vlanId,omitempty"`
}

type EmptyResult struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *EmptyResult) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = EmptyResult{XAdditionalProperties: tmp}
	return nil
}

func (t *EmptyResult) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

type UpdateVlanConfig struct {
	ArpInspections []*VlanArpInspections `json:"arpInspections,omitempty"`

	// ArpInspectionTrustPort
	// ARP Inspection Trust Port
	ArpInspectionTrustPort *string `json:"arpInspectionTrustPort,omitempty"`

	// EnableArpInspection
	// ARP Inspection Enabled
	EnableArpInspection *bool `json:"enableArpInspection,omitempty"`

	// EnableIpv4DHCPSnooping
	// IPv4 DHCP Snooping Enabled
	EnableIpv4DHCPSnooping *bool `json:"enableIpv4DhcpSnooping,omitempty"`

	// IgmpSnooping
	// IGMP Snooping
	IgmpSnooping *string `json:"igmpSnooping,omitempty" validate:"oneof=IST_NONE ACTIVE PASSIVE"`

	// Ipv4DHCPSnoopingTrustPort
	// IPv4 DHCP Snooping Trust Port
	Ipv4DHCPSnoopingTrustPort *string `json:"ipv4DhcpSnoopingTrustPort,omitempty"`

	// MulticastVersion
	// Mutilcast Version
	MulticastVersion *int `json:"multicastVersion,omitempty"`

	// Name
	// Vlan Name
	Name *string `json:"name,omitempty"`

	Ports []*VlanPorts `json:"ports,omitempty"`

	// PushTime
	// Puch Schedule Time
	PushTime *int `json:"pushTime,omitempty"`

	// PushTimeType
	// Puch Config Type
	PushTimeType *string `json:"pushTimeType,omitempty" validate:"oneof=NOW SCHEDULE"`

	// SpanningTree
	// Spanning Tree
	SpanningTree *string `json:"spanningTree,omitempty" validate:"oneof=STT_NONE STP RSTP"`

	// SpanningTreePriority
	// Spanning Tree Priority
	SpanningTreePriority *int `json:"spanningTreePriority,omitempty"`
}

type VlanArpInspections struct {
	// IP
	// The IP of ArpInspections
	IP *string `json:"ip,omitempty"`

	// Mac
	// The MAC of ArpInspections
	Mac *string `json:"mac,omitempty"`
}

type VlanConfig struct {
	ArpInspections []*VlanArpInspections `json:"arpInspections,omitempty"`

	// ArpInspectionTrustPort
	// ARP Inspection Trust Port
	ArpInspectionTrustPort *string `json:"arpInspectionTrustPort,omitempty"`

	// CreatedTime
	// The create time of the Vlan Config
	CreatedTime *int `json:"createdTime,omitempty"`

	// EnableArpInspection
	// ARP Inspection Enabled
	EnableArpInspection *bool `json:"enableArpInspection,omitempty"`

	// EnableIpv4DHCPSnooping
	// IPv4 DHCP Snooping Enabled
	EnableIpv4DHCPSnooping *bool `json:"enableIpv4DhcpSnooping,omitempty"`

	// FamilyID
	// Family Id
	FamilyID *string `json:"familyId,omitempty"`

	// GroupID
	// Group Id
	GroupID *string `json:"groupId,omitempty"`

	// ID
	// Id
	ID *string `json:"id,omitempty"`

	// IgmpSnooping
	// IGMP Snooping
	IgmpSnooping *string `json:"igmpSnooping,omitempty"`

	// Ipv4DHCPSnoopingTrustPort
	// IPv4 DHCP Snooping Trust Port
	Ipv4DHCPSnoopingTrustPort *string `json:"ipv4DhcpSnoopingTrustPort,omitempty"`

	// MulticastVersion
	// Mutilcast Version
	MulticastVersion *int `json:"multicastVersion,omitempty"`

	// Name
	// Vlan Name
	Name *string `json:"name,omitempty"`

	Ports []*VlanPorts `json:"ports,omitempty"`

	// PushTime
	// Puch Schedule Time
	PushTime *int `json:"pushTime,omitempty"`

	// PushTimeType
	// Puch Config Type
	PushTimeType *string `json:"pushTimeType,omitempty" validate:"oneof=NOW SCHEDULE"`

	// SpanningTree
	// Spanning Tree
	SpanningTree *string `json:"spanningTree,omitempty" validate:"oneof=STT_NONE STP RSTP"`

	// SpanningTreePriority
	// Spanning Tree Priority
	SpanningTreePriority *int `json:"spanningTreePriority,omitempty"`

	// SwitchID
	// Switch Id
	SwitchID *string `json:"switchId,omitempty"`

	// UpdatedTime
	// The update time of the Vlan Config
	UpdatedTime *int `json:"updatedTime,omitempty"`

	// VlanID
	// Vlan Id
	VlanID *int `json:"vlanId,omitempty"`
}

type VlanConfigQueryResult struct {
	// Extra
	// Any additional response data
	Extra *VlanConfigQueryResultExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first Vlan Configs returned out of the complete Vlan Configs list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Vlan Configs after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*VlanConfig `json:"list,omitempty"`

	// RawDataTotalCount
	// Total Vlan Configs count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Vlan Configs count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

// VlanConfigQueryResultExtraType
//
// Any additional response data
type VlanConfigQueryResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *VlanConfigQueryResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = VlanConfigQueryResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *VlanConfigQueryResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

type VlanPorts struct {
	// SwitchModel
	// The Switch Model of Ports
	SwitchModel *string `json:"switchModel,omitempty"`

	// TaggedPorts
	// The Tagged Ports of Ports
	TaggedPorts *string `json:"taggedPorts,omitempty"`

	// UntaggedPorts
	// The Untagged Ports of Ports
	UntaggedPorts *string `json:"untaggedPorts,omitempty"`
}
