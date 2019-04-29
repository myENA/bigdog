package stack

// API Version: v8_0

type List struct {
	FirstIndex        *int    `json:"firstIndex,omitempty"`
	HasMore           *bool   `json:"hasMore,omitempty"`
	List              []*List `json:"list,omitempty"`
	RawDataTotalCount *int    `json:"rawDataTotalCount,omitempty"`
	TotalCount        *int    `json:"totalCount,omitempty"`
}

type Member struct {
	ActiveMode   *string        `json:"activeMode,omitempty"`
	Model        *string        `json:"model,omitempty"`
	Ports        *int           `json:"ports,omitempty"`
	PortStatus   *PortStatus    `json:"portStatus,omitempty"`
	SerialNumber *string        `json:"serialNumber,omitempty"`
	SwitchModule *string        `json:"switchModule,omitempty"`
	SwitchName   *string        `json:"switchName,omitempty"`
	SwitchPorts  []*SwitchPorts `json:"switchPorts,omitempty"`
	SwitchUnit   *string        `json:"switchUnit,omitempty"`
}

type MemberPortStatusType struct {
	AdminDown *int    `json:"adminDown,omitempty"`
	Down      *int    `json:"down,omitempty"`
	Speed     *string `json:"speed,omitempty"`
	Total     *int    `json:"total,omitempty"`
	Up        *int    `json:"up,omitempty"`
	Warning   *int    `json:"warning,omitempty"`
}

type MemberSwitchPortsType struct {
	AdminStatus        *string          `json:"adminStatus,omitempty"`
	ConnectedDevice    *ConnectedDevice `json:"connectedDevice,omitempty"`
	ID                 *string          `json:"id,omitempty"`
	InUtilization      *float64         `json:"inUtilization,omitempty"`
	LagName            *string          `json:"lagName,omitempty"`
	Mac                *string          `json:"mac,omitempty"`
	Name               *string          `json:"name,omitempty"`
	NeighborName       *string          `json:"neighborName,omitempty"`
	OpticsType         *string          `json:"opticsType,omitempty"`
	OutUtilization     *float64         `json:"outUtilization,omitempty"`
	Poe                *Poe             `json:"poe,omitempty"`
	PoeEnabled         *bool            `json:"poeEnabled,omitempty"`
	PortIdentifier     *string          `json:"portIdentifier,omitempty"`
	PortSpeed          *string          `json:"portSpeed,omitempty"`
	SampledInstant     *string          `json:"sampledInstant,omitempty"`
	Status             *string          `json:"status,omitempty"`
	StpState           *int             `json:"stpState,omitempty"`
	SwitchGroup        *string          `json:"switchGroup,omitempty"`
	SwitchName         *string          `json:"switchName,omitempty"`
	TrafficUsage       *TrafficUsage    `json:"trafficUsage,omitempty"`
	Type               *string          `json:"type,omitempty"`
	UnTaggedVlan       *string          `json:"unTaggedVlan,omitempty"`
	UsedInFormingStack *bool            `json:"usedInFormingStack,omitempty"`
	Vlans              *string          `json:"vlans,omitempty"`
}

type MemberSwitchPortsTypeConnectedDeviceType struct {
	DomainID              *string `json:"domainId,omitempty"`
	ID                    *string `json:"id,omitempty"`
	IsRuckusAP            *string `json:"isRuckusAP,omitempty"`
	LocalPort             *string `json:"localPort,omitempty"`
	LocalPortIfaceName    *string `json:"localPortIfaceName,omitempty"`
	LocalPortMac          *string `json:"localPortMac,omitempty"`
	RemoteDeviceName      *string `json:"remoteDeviceName,omitempty"`
	RemotePort            *string `json:"remotePort,omitempty"`
	RemotePortDesc        *string `json:"remotePortDesc,omitempty"`
	RemotePortMac         *string `json:"remotePortMac,omitempty"`
	RemotePortType        *string `json:"remotePortType,omitempty"`
	SwitchGroup           *string `json:"switchGroup,omitempty"`
	SwitchGroupLevelOneID *string `json:"switchGroupLevelOneId,omitempty"`
	SwitchGroupLevelTwoID *string `json:"switchGroupLevelTwoId,omitempty"`
	SwitchID              *string `json:"switchId,omitempty"`
	SwitchName            *string `json:"switchName,omitempty"`
	TenantID              *string `json:"tenantId,omitempty"`
	UnitID                *string `json:"unitId,omitempty"`
}

type MemberSwitchPortsTypePoeType struct {
	Free    *int     `json:"free,omitempty"`
	Percent *float64 `json:"percent,omitempty"`
	Total   *int     `json:"total,omitempty"`
}

type MemberSwitchPortsTypeTrafficUsageType struct {
	Rx *int `json:"rx,omitempty"`
	Tx *int `json:"tx,omitempty"`
}
