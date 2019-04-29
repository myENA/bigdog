package switchmswitch

// API Version: v8_0

type AuditID struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type BarChart struct {
	ID    *string  `json:"id,omitempty"`
	Key   *string  `json:"key,omitempty"`
	Value *float64 `json:"value,omitempty"`
}

type ConnectedAPsQueryList struct {
	Extra             interface{} `json:"extra,omitempty"`
	FirstIndex        *int        `json:"firstIndex,omitempty"`
	HasMore           *bool       `json:"hasMore,omitempty"`
	List              []*List     `json:"list,omitempty"`
	RawDataTotalCount *int        `json:"rawDataTotalCount,omitempty"`
	TotalCount        *int        `json:"totalCount,omitempty"`
}

type ConnectedDevice struct {
	DomainID              *string                `json:"domainId,omitempty"`
	ID                    *string                `json:"id,omitempty"`
	IsRuckusAP            *string                `json:"isRuckusAP,omitempty"`
	LocalPort             *string                `json:"localPort,omitempty"`
	LocalPortIfaceName    *string                `json:"localPortIfaceName,omitempty"`
	LocalPortMac          *string                `json:"localPortMac,omitempty"`
	RemoteDeviceName      *string                `json:"remoteDeviceName,omitempty"`
	RemotePort            *string                `json:"remotePort,omitempty"`
	RemotePortDesc        *string                `json:"remotePortDesc,omitempty"`
	RemotePortMac         *string                `json:"remotePortMac,omitempty"`
	RemotePortType        *string                `json:"remotePortType,omitempty"`
	SampledInstant        map[string]interface{} `json:"sampledInstant,omitempty"`
	SwitchGroup           *string                `json:"switchGroup,omitempty"`
	SwitchGroupLevelOneID *string                `json:"switchGroupLevelOneId,omitempty"`
	SwitchGroupLevelTwoID *string                `json:"switchGroupLevelTwoId,omitempty"`
	SwitchID              *string                `json:"switchId,omitempty"`
	SwitchName            *string                `json:"switchName,omitempty"`
	TenantID              *string                `json:"tenantId,omitempty"`
	UnitID                *string                `json:"unitId,omitempty"`
}

type ConnectedDevicesQueryList struct {
	Extra             interface{} `json:"extra,omitempty"`
	FirstIndex        *int        `json:"firstIndex,omitempty"`
	HasMore           *bool       `json:"hasMore,omitempty"`
	List              []*List     `json:"list,omitempty"`
	RawDataTotalCount *int        `json:"rawDataTotalCount,omitempty"`
	TotalCount        *int        `json:"totalCount,omitempty"`
}

type DeleteSwitchesResultList struct {
	Extra             interface{} `json:"extra,omitempty"`
	FirstIndex        *int        `json:"firstIndex,omitempty"`
	HasMore           *bool       `json:"hasMore,omitempty"`
	List              []*List     `json:"list,omitempty"`
	RawDataTotalCount *int        `json:"rawDataTotalCount,omitempty"`
	TotalCount        *int        `json:"totalCount,omitempty"`
}

type Firmware struct {
	FromVersion *string  `json:"fromVersion,omitempty"`
	Timestamp   *float64 `json:"timestamp,omitempty"`
	ToVersion   *string  `json:"toVersion,omitempty"`
}

type FirmwareHistoryQueryResultList struct {
	Extra             interface{} `json:"extra,omitempty"`
	FirstIndex        *int        `json:"firstIndex,omitempty"`
	HasMore           *bool       `json:"hasMore,omitempty"`
	List              []*List     `json:"list,omitempty"`
	RawDataTotalCount *int        `json:"rawDataTotalCount,omitempty"`
	TotalCount        *int        `json:"totalCount,omitempty"`
}

type NetworkSwitch struct {
	Alarm              *int            `json:"alarm,omitempty"`
	DefaultGateway     *string         `json:"defaultGateway,omitempty"`
	DomainID           *string         `json:"domainId,omitempty"`
	FirmwareUpdate     *FirmwareUpdate `json:"firmwareUpdate,omitempty"`
	FirmwareVersion    *string         `json:"firmwareVersion,omitempty"`
	GroupID            *string         `json:"groupId,omitempty"`
	GroupName          *string         `json:"groupName,omitempty"`
	ID                 *string         `json:"id,omitempty"`
	IPAddress          *string         `json:"ipAddress,omitempty"`
	IPAddressType      *string         `json:"ipAddressType,omitempty"`
	LastBackupStatus   *string         `json:"lastBackupStatus,omitempty"`
	LastBackupTime     *int            `json:"lastBackupTime,omitempty"`
	LastRestoreStatus  *string         `json:"lastRestoreStatus,omitempty"`
	LastRestoreTime    *int            `json:"lastRestoreTime,omitempty"`
	MacAddress         *string         `json:"macAddress,omitempty"`
	Model              *string         `json:"model,omitempty"`
	Modules            *string         `json:"modules,omitempty"`
	NumOfUnits         *int            `json:"numOfUnits,omitempty"`
	Poe                *Poe            `json:"poe,omitempty"`
	Ports              *int            `json:"ports,omitempty"`
	PortStatus         *PortStatus     `json:"portStatus,omitempty"`
	RegistrationStatus *string         `json:"registrationStatus,omitempty"`
	SerialNumber       *string         `json:"serialNumber,omitempty"`
	StackID            *string         `json:"stackId,omitempty"`
	Status             *string         `json:"status,omitempty"`
	SwitchName         *string         `json:"switchName,omitempty"`
	UpTime             *string         `json:"upTime,omitempty"`
}

type NetworkSwitchFirmwareUpdateType struct {
	ModifiedTime  *string `json:"modifiedTime,omitempty"`
	ScheduledTime *string `json:"scheduledTime,omitempty"`
	ScheduleID    *string `json:"scheduleId,omitempty"`
	Status        *string `json:"status,omitempty"`
	ToVersion     *string `json:"toVersion,omitempty"`
}

type NetworkSwitchPoeType struct {
	Free    *int     `json:"free,omitempty"`
	Percent *float64 `json:"percent,omitempty"`
	Total   *int     `json:"total,omitempty"`
}

type NetworkSwitchPortStatusType struct {
	AdminDown *int    `json:"adminDown,omitempty"`
	Down      *int    `json:"down,omitempty"`
	Speed     *string `json:"speed,omitempty"`
	SpeedInt  *int    `json:"speedInt,omitempty"`
	Total     *int    `json:"total,omitempty"`
	Up        *int    `json:"up,omitempty"`
	Warning   *int    `json:"warning,omitempty"`
}

type PortDetails struct {
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
	Packets            *Packets         `json:"packets,omitempty"`
	Poe                *Poe             `json:"poe,omitempty"`
	PoeEnabled         *bool            `json:"poeEnabled,omitempty"`
	PoeType            *string          `json:"poeType,omitempty"`
	PortError          *PortError       `json:"portError,omitempty"`
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

type PortDetailsConnectedDeviceType struct {
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

type PortDetailsPacketsType struct {
	BroadcastIn  *int `json:"broadcastIn,omitempty"`
	BroadcastOut *int `json:"broadcastOut,omitempty"`
	MulticastIn  *int `json:"multicastIn,omitempty"`
	MulticastOut *int `json:"multicastOut,omitempty"`
}

type PortDetailsPoeType struct {
	Free    *int     `json:"free,omitempty"`
	Percent *float64 `json:"percent,omitempty"`
	Total   *int     `json:"total,omitempty"`
}

type PortDetailsPortErrorType struct {
	CrcError  *int `json:"crcError,omitempty"`
	InDiscard *int `json:"inDiscard,omitempty"`
	InError   *int `json:"inError,omitempty"`
	OutError  *int `json:"outError,omitempty"`
}

type PortDetailsQueryResultList struct {
	Extra             interface{} `json:"extra,omitempty"`
	FirstIndex        *int        `json:"firstIndex,omitempty"`
	HasMore           *bool       `json:"hasMore,omitempty"`
	List              []*List     `json:"list,omitempty"`
	RawDataTotalCount *int        `json:"rawDataTotalCount,omitempty"`
	TotalCount        *int        `json:"totalCount,omitempty"`
}

type PortDetailsTrafficUsageType struct {
	Rx *int `json:"rx,omitempty"`
	Tx *int `json:"tx,omitempty"`
}

type PortStatus struct {
	AdminDown *int    `json:"adminDown,omitempty"`
	Down      *int    `json:"down,omitempty"`
	Speed     *string `json:"speed,omitempty"`
	SpeedInt  *int    `json:"speedInt,omitempty"`
	Total     *int    `json:"total,omitempty"`
	Up        *int    `json:"up,omitempty"`
	Warning   *int    `json:"warning,omitempty"`
}

type StackMember struct {
	ActiveMode   *string      `json:"activeMode,omitempty"`
	Model        *string      `json:"model,omitempty"`
	Ports        *int         `json:"ports,omitempty"`
	PortStatus   *PortStatus  `json:"portStatus,omitempty"`
	SerialNumber *string      `json:"serialNumber,omitempty"`
	SwitchModule *string      `json:"switchModule,omitempty"`
	SwitchName   *string      `json:"switchName,omitempty"`
	SwitchPorts  *SwitchPorts `json:"switchPorts,omitempty"`
	SwitchUnit   *string      `json:"switchUnit,omitempty"`
}

type StackMemberQueryResult struct {
	Extra             interface{} `json:"extra,omitempty"`
	FirstIndex        *int        `json:"firstIndex,omitempty"`
	HasMore           *bool       `json:"hasMore,omitempty"`
	List              []*List     `json:"list,omitempty"`
	RawDataTotalCount *int        `json:"rawDataTotalCount,omitempty"`
	TotalCount        *int        `json:"totalCount,omitempty"`
}

type SwitchPortsSummaryQueryResultList struct {
	Extra             interface{} `json:"extra,omitempty"`
	FirstIndex        *int        `json:"firstIndex,omitempty"`
	HasMore           *bool       `json:"hasMore,omitempty"`
	List              []*List     `json:"list,omitempty"`
	RawDataTotalCount *int        `json:"rawDataTotalCount,omitempty"`
	TotalCount        *int        `json:"totalCount,omitempty"`
}

type SwitchQueryResultList struct {
	Extra             *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex        *int                 `json:"firstIndex,omitempty"`
	HasMore           *bool                `json:"hasMore,omitempty"`
	List              []*List              `json:"list,omitempty"`
	RawDataTotalCount *int                 `json:"rawDataTotalCount,omitempty"`
	TotalCount        *int                 `json:"totalCount,omitempty"`
}

type TopSwitchesByFirmwareQueryResultList struct {
	Extra             interface{} `json:"extra,omitempty"`
	FirstIndex        *int        `json:"firstIndex,omitempty"`
	HasMore           *bool       `json:"hasMore,omitempty"`
	List              []*List     `json:"list,omitempty"`
	RawDataTotalCount *int        `json:"rawDataTotalCount,omitempty"`
	TotalCount        *int        `json:"totalCount,omitempty"`
}

type TopSwitchesByModelQueryResultList struct {
	Extra             interface{} `json:"extra,omitempty"`
	FirstIndex        *int        `json:"firstIndex,omitempty"`
	HasMore           *bool       `json:"hasMore,omitempty"`
	List              []*List     `json:"list,omitempty"`
	RawDataTotalCount *int        `json:"rawDataTotalCount,omitempty"`
	TotalCount        *int        `json:"totalCount,omitempty"`
}
