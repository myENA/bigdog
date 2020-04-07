package vsz

// API Version: v9_0

import (
	"encoding/json"
)

type SwitchMVlanConfigCreateVlanConfig struct {
	// ArpInspections
	// Constraints:
	//    - nullable
	ArpInspections []*SwitchMVlanConfigVlanArpInspections `json:"arpInspections,omitempty" validate:"omitempty,dive"`

	// ArpInspectionTrustPort
	// ARP Inspection Trust Port
	// Constraints:
	//    - nullable
	ArpInspectionTrustPort *string `json:"arpInspectionTrustPort,omitempty"`

	// EnableArpInspection
	// ARP Inspection Enabled
	// Constraints:
	//    - nullable
	EnableArpInspection *bool `json:"enableArpInspection,omitempty"`

	// EnableIpv4DhcpSnooping
	// IPv4 DHCP Snooping Enabled
	// Constraints:
	//    - nullable
	EnableIpv4DhcpSnooping *bool `json:"enableIpv4DhcpSnooping,omitempty"`

	// FamilyId
	// Family Id
	// Constraints:
	//    - nullable
	FamilyId *string `json:"familyId,omitempty"`

	// GroupId
	// Group Id
	// Constraints:
	//    - nullable
	GroupId *string `json:"groupId,omitempty"`

	// IgmpSnooping
	// IGMP Snooping
	// Constraints:
	//    - nullable
	//    - oneof:[IST_NONE,ACTIVE,PASSIVE]
	IgmpSnooping *string `json:"igmpSnooping,omitempty" validate:"omitempty,oneof=IST_NONE ACTIVE PASSIVE"`

	// Ipv4DhcpSnoopingTrustPort
	// IPv4 DHCP Snooping Trust Port
	// Constraints:
	//    - nullable
	Ipv4DhcpSnoopingTrustPort *string `json:"ipv4DhcpSnoopingTrustPort,omitempty"`

	// MulticastVersion
	// Mutilcast Version
	// Constraints:
	//    - nullable
	MulticastVersion *int `json:"multicastVersion,omitempty"`

	// Name
	// Vlan Name
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// Ports
	// Constraints:
	//    - nullable
	Ports []*SwitchMVlanConfigVlanPorts `json:"ports,omitempty" validate:"omitempty,dive"`

	// PushTime
	// Puch Schedule Time
	// Constraints:
	//    - nullable
	PushTime *int `json:"pushTime,omitempty"`

	// PushTimeType
	// Puch Config Type
	// Constraints:
	//    - nullable
	//    - oneof:[NOW,SCHEDULE]
	PushTimeType *string `json:"pushTimeType,omitempty" validate:"omitempty,oneof=NOW SCHEDULE"`

	// RootBridgeFamilyId
	// Constraints:
	//    - nullable
	RootBridgeFamilyId *string `json:"rootBridgeFamilyId,omitempty"`

	// SpanningTree
	// Spanning Tree
	// Constraints:
	//    - nullable
	//    - oneof:[STT_NONE,STP,RSTP]
	SpanningTree *string `json:"spanningTree,omitempty" validate:"omitempty,oneof=STT_NONE STP RSTP"`

	// SpanningTreePriority
	// Spanning Tree Priority
	// Constraints:
	//    - nullable
	SpanningTreePriority *int `json:"spanningTreePriority,omitempty"`

	// SwitchId
	// Switch Id
	// Constraints:
	//    - nullable
	SwitchId *string `json:"switchId,omitempty"`

	// VlanId
	// Vlan Id
	// Constraints:
	//    - nullable
	VlanId *int `json:"vlanId,omitempty"`
}

func NewSwitchMVlanConfigCreateVlanConfig() *SwitchMVlanConfigCreateVlanConfig {
	m := new(SwitchMVlanConfigCreateVlanConfig)
	return m
}

type SwitchMVlanConfigUpdateVlanConfig struct {
	// ArpInspections
	// Constraints:
	//    - nullable
	ArpInspections []*SwitchMVlanConfigVlanArpInspections `json:"arpInspections,omitempty" validate:"omitempty,dive"`

	// ArpInspectionTrustPort
	// ARP Inspection Trust Port
	// Constraints:
	//    - nullable
	ArpInspectionTrustPort *string `json:"arpInspectionTrustPort,omitempty"`

	// EnableArpInspection
	// ARP Inspection Enabled
	// Constraints:
	//    - nullable
	EnableArpInspection *bool `json:"enableArpInspection,omitempty"`

	// EnableIpv4DhcpSnooping
	// IPv4 DHCP Snooping Enabled
	// Constraints:
	//    - nullable
	EnableIpv4DhcpSnooping *bool `json:"enableIpv4DhcpSnooping,omitempty"`

	// IgmpSnooping
	// IGMP Snooping
	// Constraints:
	//    - nullable
	//    - oneof:[IST_NONE,ACTIVE,PASSIVE]
	IgmpSnooping *string `json:"igmpSnooping,omitempty" validate:"omitempty,oneof=IST_NONE ACTIVE PASSIVE"`

	// Ipv4DhcpSnoopingTrustPort
	// IPv4 DHCP Snooping Trust Port
	// Constraints:
	//    - nullable
	Ipv4DhcpSnoopingTrustPort *string `json:"ipv4DhcpSnoopingTrustPort,omitempty"`

	// MulticastVersion
	// Mutilcast Version
	// Constraints:
	//    - nullable
	MulticastVersion *int `json:"multicastVersion,omitempty"`

	// Name
	// Vlan Name
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// Ports
	// Constraints:
	//    - nullable
	Ports []*SwitchMVlanConfigVlanPorts `json:"ports,omitempty" validate:"omitempty,dive"`

	// PushTime
	// Puch Schedule Time
	// Constraints:
	//    - nullable
	PushTime *int `json:"pushTime,omitempty"`

	// PushTimeType
	// Puch Config Type
	// Constraints:
	//    - nullable
	//    - oneof:[NOW,SCHEDULE]
	PushTimeType *string `json:"pushTimeType,omitempty" validate:"omitempty,oneof=NOW SCHEDULE"`

	// RootBridgeFamilyId
	// Constraints:
	//    - nullable
	RootBridgeFamilyId *string `json:"rootBridgeFamilyId,omitempty"`

	// SpanningTree
	// Spanning Tree
	// Constraints:
	//    - nullable
	//    - oneof:[STT_NONE,STP,RSTP]
	SpanningTree *string `json:"spanningTree,omitempty" validate:"omitempty,oneof=STT_NONE STP RSTP"`

	// SpanningTreePriority
	// Spanning Tree Priority
	// Constraints:
	//    - nullable
	SpanningTreePriority *int `json:"spanningTreePriority,omitempty"`
}

func NewSwitchMVlanConfigUpdateVlanConfig() *SwitchMVlanConfigUpdateVlanConfig {
	m := new(SwitchMVlanConfigUpdateVlanConfig)
	return m
}

type SwitchMVlanConfigVlanArpInspections struct {
	// Ip
	// The IP of ArpInspections
	// Constraints:
	//    - nullable
	Ip *string `json:"ip,omitempty"`

	// Mac
	// The MAC of ArpInspections
	// Constraints:
	//    - nullable
	Mac *string `json:"mac,omitempty"`
}

func NewSwitchMVlanConfigVlanArpInspections() *SwitchMVlanConfigVlanArpInspections {
	m := new(SwitchMVlanConfigVlanArpInspections)
	return m
}

type SwitchMVlanConfig struct {
	// ArpInspections
	// Constraints:
	//    - nullable
	ArpInspections []*SwitchMVlanConfigVlanArpInspections `json:"arpInspections,omitempty" validate:"omitempty,dive"`

	// ArpInspectionTrustPort
	// ARP Inspection Trust Port
	// Constraints:
	//    - nullable
	ArpInspectionTrustPort *string `json:"arpInspectionTrustPort,omitempty"`

	// CreatedTime
	// The create time of the Vlan Config
	// Constraints:
	//    - nullable
	CreatedTime *int `json:"createdTime,omitempty"`

	// EnableArpInspection
	// ARP Inspection Enabled
	// Constraints:
	//    - nullable
	EnableArpInspection *bool `json:"enableArpInspection,omitempty"`

	// EnableIpv4DhcpSnooping
	// IPv4 DHCP Snooping Enabled
	// Constraints:
	//    - nullable
	EnableIpv4DhcpSnooping *bool `json:"enableIpv4DhcpSnooping,omitempty"`

	// FamilyId
	// Family Id
	// Constraints:
	//    - nullable
	FamilyId *string `json:"familyId,omitempty"`

	// GroupId
	// Group Id
	// Constraints:
	//    - nullable
	GroupId *string `json:"groupId,omitempty"`

	// Id
	// Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// IgmpSnooping
	// IGMP Snooping
	// Constraints:
	//    - nullable
	IgmpSnooping *string `json:"igmpSnooping,omitempty"`

	// Ipv4DhcpSnoopingTrustPort
	// IPv4 DHCP Snooping Trust Port
	// Constraints:
	//    - nullable
	Ipv4DhcpSnoopingTrustPort *string `json:"ipv4DhcpSnoopingTrustPort,omitempty"`

	// MulticastVersion
	// Mutilcast Version
	// Constraints:
	//    - nullable
	MulticastVersion *int `json:"multicastVersion,omitempty"`

	// Name
	// Vlan Name
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// Ports
	// Constraints:
	//    - nullable
	Ports []*SwitchMVlanConfigVlanPorts `json:"ports,omitempty" validate:"omitempty,dive"`

	// PushTime
	// Puch Schedule Time
	// Constraints:
	//    - nullable
	PushTime *int `json:"pushTime,omitempty"`

	// PushTimeType
	// Puch Config Type
	// Constraints:
	//    - nullable
	//    - oneof:[NOW,SCHEDULE]
	PushTimeType *string `json:"pushTimeType,omitempty" validate:"omitempty,oneof=NOW SCHEDULE"`

	// RootBridgeFamilyId
	// Constraints:
	//    - nullable
	RootBridgeFamilyId *string `json:"rootBridgeFamilyId,omitempty"`

	// SpanningTree
	// Spanning Tree
	// Constraints:
	//    - nullable
	//    - oneof:[STT_NONE,STP,RSTP]
	SpanningTree *string `json:"spanningTree,omitempty" validate:"omitempty,oneof=STT_NONE STP RSTP"`

	// SpanningTreePriority
	// Spanning Tree Priority
	// Constraints:
	//    - nullable
	SpanningTreePriority *int `json:"spanningTreePriority,omitempty"`

	// SwitchId
	// Switch Id
	// Constraints:
	//    - nullable
	SwitchId *string `json:"switchId,omitempty"`

	// UpdatedTime
	// The update time of the Vlan Config
	// Constraints:
	//    - nullable
	UpdatedTime *int `json:"updatedTime,omitempty"`

	// VlanId
	// Vlan Id
	// Constraints:
	//    - nullable
	VlanId *int `json:"vlanId,omitempty"`
}

func NewSwitchMVlanConfig() *SwitchMVlanConfig {
	m := new(SwitchMVlanConfig)
	return m
}

type SwitchMVlanConfigQueryResult struct {
	// Extra
	// Any additional response data
	// Constraints:
	//    - nullable
	Extra *SwitchMVlanConfigQueryResultExtraType `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first Vlan Configs returned out of the complete Vlan Configs list
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more Vlan Configs after the current displayed list
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*SwitchMVlanConfig `json:"list,omitempty" validate:"omitempty,dive"`

	// RawDataTotalCount
	// Total Vlan Configs count
	// Constraints:
	//    - nullable
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Vlan Configs count in this response
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMVlanConfigQueryResult() *SwitchMVlanConfigQueryResult {
	m := new(SwitchMVlanConfigQueryResult)
	return m
}

// SwitchMVlanConfigQueryResultExtraType
//
// Any additional response data
// Constraints:
//    - nullable
type SwitchMVlanConfigQueryResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMVlanConfigQueryResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMVlanConfigQueryResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMVlanConfigQueryResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMVlanConfigQueryResultExtraType() *SwitchMVlanConfigQueryResultExtraType {
	m := new(SwitchMVlanConfigQueryResultExtraType)
	return m
}

type SwitchMVlanConfigVlanPorts struct {
	// SwitchModel
	// The Switch Model of Ports
	// Constraints:
	//    - nullable
	SwitchModel *string `json:"switchModel,omitempty"`

	// TaggedPorts
	// The Tagged Ports of Ports
	// Constraints:
	//    - nullable
	TaggedPorts *string `json:"taggedPorts,omitempty"`

	// UntaggedPorts
	// The Untagged Ports of Ports
	// Constraints:
	//    - nullable
	UntaggedPorts *string `json:"untaggedPorts,omitempty"`
}

func NewSwitchMVlanConfigVlanPorts() *SwitchMVlanConfigVlanPorts {
	m := new(SwitchMVlanConfigVlanPorts)
	return m
}
