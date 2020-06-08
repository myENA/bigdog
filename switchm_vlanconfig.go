package ruckus

// API Version: v9_0

import (
	"encoding/json"
)

type SwitchMSwitchVLANConfigCreateVlanConfig struct {
	ArpInspections []*SwitchMSwitchVLANConfigVlanArpInspections `json:"arpInspections,omitempty"`

	// ArpInspectionTrustPort
	// ARP Inspection Trust Port
	ArpInspectionTrustPort *string `json:"arpInspectionTrustPort,omitempty"`

	// EnableArpInspection
	// ARP Inspection Enabled
	EnableArpInspection *bool `json:"enableArpInspection,omitempty"`

	// EnableIpv4DhcpSnooping
	// IPv4 DHCP Snooping Enabled
	EnableIpv4DhcpSnooping *bool `json:"enableIpv4DhcpSnooping,omitempty"`

	// FamilyId
	// Family Id
	FamilyId *string `json:"familyId,omitempty"`

	// GroupId
	// Group Id
	GroupId *string `json:"groupId,omitempty"`

	// IgmpSnooping
	// IGMP Snooping
	// Constraints:
	//    - oneof:[IST_NONE,ACTIVE,PASSIVE]
	IgmpSnooping *string `json:"igmpSnooping,omitempty"`

	// Ipv4DhcpSnoopingTrustPort
	// IPv4 DHCP Snooping Trust Port
	Ipv4DhcpSnoopingTrustPort *string `json:"ipv4DhcpSnoopingTrustPort,omitempty"`

	// MulticastVersion
	// Mutilcast Version
	MulticastVersion *int `json:"multicastVersion,omitempty"`

	// Name
	// Vlan Name
	Name *string `json:"name,omitempty"`

	Ports []*SwitchMSwitchVLANConfigVlanPorts `json:"ports,omitempty"`

	// PushTime
	// Puch Schedule Time
	PushTime *int `json:"pushTime,omitempty"`

	// PushTimeType
	// Puch Config Type
	// Constraints:
	//    - oneof:[NOW,SCHEDULE]
	PushTimeType *string `json:"pushTimeType,omitempty"`

	RootBridgeFamilyId *string `json:"rootBridgeFamilyId,omitempty"`

	// SpanningTree
	// Spanning Tree
	// Constraints:
	//    - oneof:[STT_NONE,STP,RSTP]
	SpanningTree *string `json:"spanningTree,omitempty"`

	// SpanningTreePriority
	// Spanning Tree Priority
	SpanningTreePriority *int `json:"spanningTreePriority,omitempty"`

	// SwitchId
	// Switch Id
	SwitchId *string `json:"switchId,omitempty"`

	// VlanId
	// Vlan Id
	VlanId *int `json:"vlanId,omitempty"`
}

func NewSwitchMSwitchVLANConfigCreateVlanConfig() *SwitchMSwitchVLANConfigCreateVlanConfig {
	m := new(SwitchMSwitchVLANConfigCreateVlanConfig)
	return m
}

type SwitchMSwitchVLANConfigUpdateVlanConfig struct {
	ArpInspections []*SwitchMSwitchVLANConfigVlanArpInspections `json:"arpInspections,omitempty"`

	// ArpInspectionTrustPort
	// ARP Inspection Trust Port
	ArpInspectionTrustPort *string `json:"arpInspectionTrustPort,omitempty"`

	// EnableArpInspection
	// ARP Inspection Enabled
	EnableArpInspection *bool `json:"enableArpInspection,omitempty"`

	// EnableIpv4DhcpSnooping
	// IPv4 DHCP Snooping Enabled
	EnableIpv4DhcpSnooping *bool `json:"enableIpv4DhcpSnooping,omitempty"`

	// IgmpSnooping
	// IGMP Snooping
	// Constraints:
	//    - oneof:[IST_NONE,ACTIVE,PASSIVE]
	IgmpSnooping *string `json:"igmpSnooping,omitempty"`

	// Ipv4DhcpSnoopingTrustPort
	// IPv4 DHCP Snooping Trust Port
	Ipv4DhcpSnoopingTrustPort *string `json:"ipv4DhcpSnoopingTrustPort,omitempty"`

	// MulticastVersion
	// Mutilcast Version
	MulticastVersion *int `json:"multicastVersion,omitempty"`

	// Name
	// Vlan Name
	Name *string `json:"name,omitempty"`

	Ports []*SwitchMSwitchVLANConfigVlanPorts `json:"ports,omitempty"`

	// PushTime
	// Puch Schedule Time
	PushTime *int `json:"pushTime,omitempty"`

	// PushTimeType
	// Puch Config Type
	// Constraints:
	//    - oneof:[NOW,SCHEDULE]
	PushTimeType *string `json:"pushTimeType,omitempty"`

	RootBridgeFamilyId *string `json:"rootBridgeFamilyId,omitempty"`

	// SpanningTree
	// Spanning Tree
	// Constraints:
	//    - oneof:[STT_NONE,STP,RSTP]
	SpanningTree *string `json:"spanningTree,omitempty"`

	// SpanningTreePriority
	// Spanning Tree Priority
	SpanningTreePriority *int `json:"spanningTreePriority,omitempty"`
}

func NewSwitchMSwitchVLANConfigUpdateVlanConfig() *SwitchMSwitchVLANConfigUpdateVlanConfig {
	m := new(SwitchMSwitchVLANConfigUpdateVlanConfig)
	return m
}

type SwitchMSwitchVLANConfigVlanArpInspections struct {
	// Ip
	// The IP of ArpInspections
	Ip *string `json:"ip,omitempty"`

	// Mac
	// The MAC of ArpInspections
	Mac *string `json:"mac,omitempty"`
}

func NewSwitchMSwitchVLANConfigVlanArpInspections() *SwitchMSwitchVLANConfigVlanArpInspections {
	m := new(SwitchMSwitchVLANConfigVlanArpInspections)
	return m
}

type SwitchMSwitchVLANConfigVlanConfig struct {
	ArpInspections []*SwitchMSwitchVLANConfigVlanArpInspections `json:"arpInspections,omitempty"`

	// ArpInspectionTrustPort
	// ARP Inspection Trust Port
	ArpInspectionTrustPort *string `json:"arpInspectionTrustPort,omitempty"`

	// CreatedTime
	// The create time of the Vlan Config
	CreatedTime *int `json:"createdTime,omitempty"`

	// EnableArpInspection
	// ARP Inspection Enabled
	EnableArpInspection *bool `json:"enableArpInspection,omitempty"`

	// EnableIpv4DhcpSnooping
	// IPv4 DHCP Snooping Enabled
	EnableIpv4DhcpSnooping *bool `json:"enableIpv4DhcpSnooping,omitempty"`

	// FamilyId
	// Family Id
	FamilyId *string `json:"familyId,omitempty"`

	// GroupId
	// Group Id
	GroupId *string `json:"groupId,omitempty"`

	// Id
	// Id
	Id *string `json:"id,omitempty"`

	// IgmpSnooping
	// IGMP Snooping
	IgmpSnooping *string `json:"igmpSnooping,omitempty"`

	// Ipv4DhcpSnoopingTrustPort
	// IPv4 DHCP Snooping Trust Port
	Ipv4DhcpSnoopingTrustPort *string `json:"ipv4DhcpSnoopingTrustPort,omitempty"`

	// MulticastVersion
	// Mutilcast Version
	MulticastVersion *int `json:"multicastVersion,omitempty"`

	// Name
	// Vlan Name
	Name *string `json:"name,omitempty"`

	Ports []*SwitchMSwitchVLANConfigVlanPorts `json:"ports,omitempty"`

	// PushTime
	// Puch Schedule Time
	PushTime *int `json:"pushTime,omitempty"`

	// PushTimeType
	// Puch Config Type
	// Constraints:
	//    - oneof:[NOW,SCHEDULE]
	PushTimeType *string `json:"pushTimeType,omitempty"`

	RootBridgeFamilyId *string `json:"rootBridgeFamilyId,omitempty"`

	// SpanningTree
	// Spanning Tree
	// Constraints:
	//    - oneof:[STT_NONE,STP,RSTP]
	SpanningTree *string `json:"spanningTree,omitempty"`

	// SpanningTreePriority
	// Spanning Tree Priority
	SpanningTreePriority *int `json:"spanningTreePriority,omitempty"`

	// SwitchId
	// Switch Id
	SwitchId *string `json:"switchId,omitempty"`

	// UpdatedTime
	// The update time of the Vlan Config
	UpdatedTime *int `json:"updatedTime,omitempty"`

	// VlanId
	// Vlan Id
	VlanId *int `json:"vlanId,omitempty"`
}

func NewSwitchMSwitchVLANConfigVlanConfig() *SwitchMSwitchVLANConfigVlanConfig {
	m := new(SwitchMSwitchVLANConfigVlanConfig)
	return m
}

type SwitchMSwitchVLANConfigVlanConfigQueryResult struct {
	// Extra
	// Any additional response data
	Extra *SwitchMSwitchVLANConfigVlanConfigQueryResultExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first Vlan Configs returned out of the complete Vlan Configs list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Vlan Configs after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMSwitchVLANConfigVlanConfig `json:"list,omitempty"`

	// RawDataTotalCount
	// Total Vlan Configs count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Vlan Configs count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMSwitchVLANConfigVlanConfigQueryResult() *SwitchMSwitchVLANConfigVlanConfigQueryResult {
	m := new(SwitchMSwitchVLANConfigVlanConfigQueryResult)
	return m
}

// SwitchMSwitchVLANConfigVlanConfigQueryResultExtraType
//
// Any additional response data
type SwitchMSwitchVLANConfigVlanConfigQueryResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchVLANConfigVlanConfigQueryResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchVLANConfigVlanConfigQueryResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchVLANConfigVlanConfigQueryResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchVLANConfigVlanConfigQueryResultExtraType() *SwitchMSwitchVLANConfigVlanConfigQueryResultExtraType {
	m := new(SwitchMSwitchVLANConfigVlanConfigQueryResultExtraType)
	return m
}

type SwitchMSwitchVLANConfigVlanPorts struct {
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

func NewSwitchMSwitchVLANConfigVlanPorts() *SwitchMSwitchVLANConfigVlanPorts {
	m := new(SwitchMSwitchVLANConfigVlanPorts)
	return m
}
