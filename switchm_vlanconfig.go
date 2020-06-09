package bigdog

// API Version: v9_0

import (
	"encoding/json"
)

type SwitchMVLANConfigCreateVlanConfig struct {
	ArpInspections []*SwitchMVLANConfigVlanArpInspections `json:"arpInspections,omitempty"`

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

	Ports []*SwitchMVLANConfigVlanPorts `json:"ports,omitempty"`

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

func NewSwitchMVLANConfigCreateVlanConfig() *SwitchMVLANConfigCreateVlanConfig {
	m := new(SwitchMVLANConfigCreateVlanConfig)
	return m
}

type SwitchMVLANConfigUpdateVlanConfig struct {
	ArpInspections []*SwitchMVLANConfigVlanArpInspections `json:"arpInspections,omitempty"`

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

	Ports []*SwitchMVLANConfigVlanPorts `json:"ports,omitempty"`

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

func NewSwitchMVLANConfigUpdateVlanConfig() *SwitchMVLANConfigUpdateVlanConfig {
	m := new(SwitchMVLANConfigUpdateVlanConfig)
	return m
}

type SwitchMVLANConfigVlanArpInspections struct {
	// Ip
	// The IP of ArpInspections
	Ip *string `json:"ip,omitempty"`

	// Mac
	// The MAC of ArpInspections
	Mac *string `json:"mac,omitempty"`
}

func NewSwitchMVLANConfigVlanArpInspections() *SwitchMVLANConfigVlanArpInspections {
	m := new(SwitchMVLANConfigVlanArpInspections)
	return m
}

type SwitchMVLANConfig struct {
	ArpInspections []*SwitchMVLANConfigVlanArpInspections `json:"arpInspections,omitempty"`

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

	Ports []*SwitchMVLANConfigVlanPorts `json:"ports,omitempty"`

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

func NewSwitchMVLANConfig() *SwitchMVLANConfig {
	m := new(SwitchMVLANConfig)
	return m
}

type SwitchMVLANConfigQueryResult struct {
	// Extra
	// Any additional response data
	Extra *SwitchMVLANConfigQueryResultExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first Vlan Configs returned out of the complete Vlan Configs list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Vlan Configs after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMVLANConfig `json:"list,omitempty"`

	// RawDataTotalCount
	// Total Vlan Configs count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Vlan Configs count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMVLANConfigQueryResult() *SwitchMVLANConfigQueryResult {
	m := new(SwitchMVLANConfigQueryResult)
	return m
}

// SwitchMVLANConfigQueryResultExtraType
//
// Any additional response data
type SwitchMVLANConfigQueryResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMVLANConfigQueryResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMVLANConfigQueryResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMVLANConfigQueryResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMVLANConfigQueryResultExtraType() *SwitchMVLANConfigQueryResultExtraType {
	m := new(SwitchMVLANConfigQueryResultExtraType)
	return m
}

type SwitchMVLANConfigVlanPorts struct {
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

func NewSwitchMVLANConfigVlanPorts() *SwitchMVLANConfigVlanPorts {
	m := new(SwitchMVLANConfigVlanPorts)
	return m
}
