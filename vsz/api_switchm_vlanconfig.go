package vsz

// API Version: v9_0

import (
	"encoding/json"
)

type SwitchMVlanconfigCreateVlanConfig struct {
	ArpInspections []*SwitchMVlanconfigVlanArpInspections `json:"arpInspections,omitempty"`

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

	Ports []*SwitchMVlanconfigVlanPorts `json:"ports,omitempty"`

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

func NewSwitchMVlanconfigCreateVlanConfig() *SwitchMVlanconfigCreateVlanConfig {
	m := new(SwitchMVlanconfigCreateVlanConfig)
	return m
}

type SwitchMVlanconfigUpdateVlanConfig struct {
	ArpInspections []*SwitchMVlanconfigVlanArpInspections `json:"arpInspections,omitempty"`

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

	Ports []*SwitchMVlanconfigVlanPorts `json:"ports,omitempty"`

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

func NewSwitchMVlanconfigUpdateVlanConfig() *SwitchMVlanconfigUpdateVlanConfig {
	m := new(SwitchMVlanconfigUpdateVlanConfig)
	return m
}

type SwitchMVlanconfigVlanArpInspections struct {
	// Ip
	// The IP of ArpInspections
	Ip *string `json:"ip,omitempty"`

	// Mac
	// The MAC of ArpInspections
	Mac *string `json:"mac,omitempty"`
}

func NewSwitchMVlanconfigVlanArpInspections() *SwitchMVlanconfigVlanArpInspections {
	m := new(SwitchMVlanconfigVlanArpInspections)
	return m
}

type SwitchMVlanconfig struct {
	ArpInspections []*SwitchMVlanconfigVlanArpInspections `json:"arpInspections,omitempty"`

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

	Ports []*SwitchMVlanconfigVlanPorts `json:"ports,omitempty"`

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

func NewSwitchMVlanconfig() *SwitchMVlanconfig {
	m := new(SwitchMVlanconfig)
	return m
}

type SwitchMVlanconfigQueryResult struct {
	// Extra
	// Any additional response data
	Extra *SwitchMVlanconfigQueryResultExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first Vlan Configs returned out of the complete Vlan Configs list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Vlan Configs after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMVlanconfig `json:"list,omitempty"`

	// RawDataTotalCount
	// Total Vlan Configs count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Vlan Configs count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMVlanconfigQueryResult() *SwitchMVlanconfigQueryResult {
	m := new(SwitchMVlanconfigQueryResult)
	return m
}

// SwitchMVlanconfigQueryResultExtraType
//
// Any additional response data
type SwitchMVlanconfigQueryResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMVlanconfigQueryResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMVlanconfigQueryResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMVlanconfigQueryResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMVlanconfigQueryResultExtraType() *SwitchMVlanconfigQueryResultExtraType {
	m := new(SwitchMVlanconfigQueryResultExtraType)
	return m
}

type SwitchMVlanconfigVlanPorts struct {
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

func NewSwitchMVlanconfigVlanPorts() *SwitchMVlanconfigVlanPorts {
	m := new(SwitchMVlanconfigVlanPorts)
	return m
}
