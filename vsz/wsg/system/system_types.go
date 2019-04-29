package system

// API Version: v8_0

type ApNumberLimitSettingOfDomain struct {
	DomainID    *string  `json:"domainId,omitempty"`
	DomainName  *string  `json:"domainName,omitempty"`
	NumberLimit *float64 `json:"numberLimit,omitempty"`
	Shared      *bool    `json:"shared,omitempty"`
}

type ApNumberLimitSettingOfZone struct {
	DomainID    *string  `json:"domainId,omitempty"`
	DomainName  *string  `json:"domainName,omitempty"`
	NumberLimit *float64 `json:"numberLimit,omitempty"`
	Shared      *bool    `json:"shared,omitempty"`
	ZoneID      *string  `json:"zoneId,omitempty"`
	ZoneName    *string  `json:"zoneName,omitempty"`
}

type CaptchaSetting struct {
	CaptchaEnabled *bool `json:"captchaEnabled,omitempty"`
}

type ControllerList struct {
	FirstIndex *int    `json:"firstIndex,omitempty"`
	HasMore    *bool   `json:"hasMore,omitempty"`
	List       []*List `json:"list,omitempty"`
	TotalCount *int    `json:"totalCount,omitempty"`
}

type ControllerListType struct {
	ApVersion      *string `json:"apVersion,omitempty"`
	ClusterIP      *string `json:"clusterIp,omitempty"`
	ClusterIpv6    *string `json:"clusterIpv6,omitempty"`
	ClusterRole    *string `json:"clusterRole,omitempty"`
	ControlIP      *string `json:"controlIp,omitempty"`
	ControlIpv6    *string `json:"controlIpv6,omitempty"`
	ControlNatIP   *string `json:"controlNatIp,omitempty"`
	Description    *string `json:"description,omitempty"`
	HostName       *string `json:"hostName,omitempty"`
	ID             *string `json:"id,omitempty"`
	Mac            *string `json:"mac,omitempty"`
	ManagementIP   *string `json:"managementIp,omitempty"`
	ManagementIpv6 *string `json:"managementIpv6,omitempty"`
	Model          *string `json:"model,omitempty"`
	Name           *string `json:"name,omitempty"`
	SerialNumber   *string `json:"serialNumber,omitempty"`
	UptimeInSec    *int    `json:"uptimeInSec,omitempty"`
	Version        *string `json:"version,omitempty"`
}

type ControlPlaneConfiguration struct {
	IPMode                      *string                      `json:"ipMode,omitempty"`
	Ipv4AccessAndCoreSeparation *Ipv4AccessAndCoreSeparation `json:"ipv4AccessAndCoreSeparation,omitempty"`
	Ipv4ClusterInterface        *Ipv4ClusterInterface        `json:"ipv4ClusterInterface,omitempty"`
	Ipv4ControlInterface        *Ipv4ControlInterface        `json:"ipv4ControlInterface,omitempty"`
	Ipv4ManagementInterface     *Ipv4ManagementInterface     `json:"ipv4ManagementInterface,omitempty"`
	Ipv6AccessAndCoreSeparation *Ipv6AccessAndCoreSeparation `json:"ipv6AccessAndCoreSeparation,omitempty"`
	Ipv6ControlInterface        *Ipv6ControlInterface        `json:"ipv6ControlInterface,omitempty"`
	Ipv6ManagementInterface     *Ipv6ManagementInterface     `json:"ipv6ManagementInterface,omitempty"`
}

type ControlPlaneList struct {
	FirstIndex *int    `json:"firstIndex,omitempty"`
	HasMore    *bool   `json:"hasMore,omitempty"`
	List       []*List `json:"list,omitempty"`
	TotalCount *int    `json:"totalCount,omitempty"`
}

type ControlPlaneListType struct {
	ClusterIP    *string `json:"clusterIp,omitempty"`
	ClusterRole  *string `json:"clusterRole,omitempty"`
	ControlIP    *string `json:"controlIp,omitempty"`
	Description  *string `json:"description,omitempty"`
	Firmware     *string `json:"firmware,omitempty"`
	ID           *string `json:"id,omitempty"`
	ManagementIP *string `json:"managementIp,omitempty"`
	Model        *string `json:"model,omitempty"`
	Name         *string `json:"name,omitempty"`
	NumOfAps     *string `json:"numOfAps,omitempty"`
	SerialNumber *string `json:"serialNumber,omitempty"`
	UpTime       *string `json:"upTime,omitempty"`
}

type CpStaticRoute struct {
	Gateway        *string `json:"gateway,omitempty"`
	InterfaceMode  *string `json:"interfaceMode,omitempty"`
	Metric         *int    `json:"metric,omitempty"`
	NetworkAddress *string `json:"networkAddress,omitempty"`
	SubnetMask     *string `json:"subnetMask,omitempty"`
}

type DataPlaneConfiguration struct {
	InterfaceMode        *string               `json:"interfaceMode,omitempty"`
	Ipv6PrimaryInterface *Ipv6PrimaryInterface `json:"ipv6PrimaryInterface,omitempty"`
	IsDataCenter         *bool                 `json:"isDataCenter,omitempty"`
	KeepConfig           *bool                 `json:"keepConfig,omitempty"`
	PrimaryInterface     *PrimaryInterface     `json:"primaryInterface,omitempty"`
	SecondaryInterface   *SecondaryInterface   `json:"secondaryInterface,omitempty"`
	StaticRoute          []*StaticRoute        `json:"staticRoute,omitempty"`
}

type DataPlaneList struct {
	FirstIndex *int    `json:"firstIndex,omitempty"`
	HasMore    *bool   `json:"hasMore,omitempty"`
	List       []*List `json:"list,omitempty"`
	TotalCount *int    `json:"totalCount,omitempty"`
}

type DataPlaneListType struct {
	BladeName    *string `json:"bladeName,omitempty"`
	DpStatus     *string `json:"dpStatus,omitempty"`
	FwVersion    *string `json:"fwVersion,omitempty"`
	GreTunnels   *string `json:"greTunnels,omitempty"`
	ID           *string `json:"id,omitempty"`
	IP           *string `json:"ip,omitempty"`
	Ipv6         *string `json:"ipv6,omitempty"`
	LastSeen     *string `json:"lastSeen,omitempty"`
	Mac          *string `json:"mac,omitempty"`
	ManagedBy    *string `json:"managedBy,omitempty"`
	Model        *string `json:"model,omitempty"`
	SerialNumber *string `json:"serialNumber,omitempty"`
	Type         *string `json:"type,omitempty"`
	Uptime       *string `json:"uptime,omitempty"`
}

type DeleteBulkFtp struct {
	IDList common.IDList `json:"idList,omitempty"`
}

type FriendlyNameLang struct {
	Display *string `json:"display,omitempty"`
	Value   *string `json:"value,omitempty"`
}

type FriendlyNameLangList struct {
	FirstIndex *int    `json:"firstIndex,omitempty"`
	HasMore    *bool   `json:"hasMore,omitempty"`
	List       []*List `json:"list,omitempty"`
	TotalCount *int    `json:"totalCount,omitempty"`
}

type Ftp struct {
	CreateDatetime     *int    `json:"createDatetime,omitempty"`
	CreatorUUID        *string `json:"creatorUUID,omitempty"`
	DomainID           *string `json:"domainId,omitempty"`
	FtpHost            *string `json:"ftpHost,omitempty"`
	FtpName            *string `json:"ftpName,omitempty"`
	FtpPassword        *string `json:"ftpPassword,omitempty"`
	FtpPort            *int    `json:"ftpPort,omitempty"`
	FtpProtocol        *string `json:"ftpProtocol,omitempty"`
	FtpRemoteDirectory *string `json:"ftpRemoteDirectory,omitempty"`
	FtpUserName        *string `json:"ftpUserName,omitempty"`
	ID                 *string `json:"id,omitempty"`
	LastModifiedBy     *string `json:"lastModifiedBy,omitempty"`
	LastModifiedOn     *int    `json:"lastModifiedOn,omitempty"`
	TenantID           *string `json:"tenantId,omitempty"`
}

type FtpGlobalSetting struct {
	Enabled     *bool   `json:"enabled,omitempty"`
	FtpID       *string `json:"ftpId,omitempty"`
	FtpInterval *string `json:"ftpInterval,omitempty"`
}

type FtpList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}

type FtpTestResponse struct {
	Data    *bool   `json:"data,omitempty"`
	Error   *string `json:"error,omitempty"`
	Extra   *string `json:"extra,omitempty"`
	Success *bool   `json:"success,omitempty"`
}

type GatewayAdvanced struct {
	AllowSessionOnAccountingFail          *bool   `json:"allowSessionOnAccountingFail,omitempty"`
	EcgiInGtpV2                           *bool   `json:"ecgiInGtpV2,omitempty"`
	GtpInterfaceType                      *string `json:"gtpInterfaceType,omitempty"`
	GtpNetworkServiceAcessPointIdentifier *int    `json:"gtpNetworkServiceAcessPointIdentifier,omitempty"`
	ImeiInGtp                             *bool   `json:"imeiInGtp,omitempty"`
	ScgRaiInGtpV2                         *bool   `json:"scgRaiInGtpV2,omitempty"`
	ScgSaiInGtpV2                         *bool   `json:"scgSaiInGtpV2,omitempty"`
	TaiInGtpV2                            *bool   `json:"taiInGtpV2,omitempty"`
}

type GetDataPlaneMeshTunnelSetting struct {
	Encrypted *bool `json:"encrypted,omitempty"`
}

type InventoryList struct {
	FirstIndex *int    `json:"firstIndex,omitempty"`
	HasMore    *bool   `json:"hasMore,omitempty"`
	List       []*List `json:"list,omitempty"`
	TotalCount *int    `json:"totalCount,omitempty"`
}

type InventoryListType struct {
	ApFirmwareVersion           *string `json:"apFirmwareVersion,omitempty"`
	Clients                     *int    `json:"clients,omitempty"`
	ConnectedAPs                *int    `json:"connectedAPs,omitempty"`
	ConnectedDownMeshAPs        *int    `json:"connectedDownMeshAPs,omitempty"`
	ConnectedeMeshAPs           *int    `json:"connectedeMeshAPs,omitempty"`
	ConnectedMeshAPs            *int    `json:"connectedMeshAPs,omitempty"`
	ConnectedMeshDisabledAPs    *int    `json:"connectedMeshDisabledAPs,omitempty"`
	ConnectedRootAPs            *int    `json:"connectedRootAPs,omitempty"`
	DisconnectedAPs             *int    `json:"disconnectedAPs,omitempty"`
	DisconnectedDownMeshAPs     *int    `json:"disconnectedDownMeshAPs,omitempty"`
	DisconnectedeMeshAPs        *int    `json:"disconnectedeMeshAPs,omitempty"`
	DisconnectedMeshAPs         *int    `json:"disconnectedMeshAPs,omitempty"`
	DisconnectedMeshDisabledAPs *int    `json:"disconnectedMeshDisabledAPs,omitempty"`
	DisconnectedRootAPs         *int    `json:"disconnectedRootAPs,omitempty"`
	DiscoveryAPs                *int    `json:"discoveryAPs,omitempty"`
	MeshEnabled                 *bool   `json:"meshEnabled,omitempty"`
	MeshSSID                    *string `json:"meshSSID,omitempty"`
	ProvisionedAPs              *int    `json:"provisionedAPs,omitempty"`
	RebootingDownMeshAPs        *int    `json:"rebootingDownMeshAPs,omitempty"`
	RebootingeMeshAPs           *int    `json:"rebootingeMeshAPs,omitempty"`
	RebootingMeshAPs            *int    `json:"rebootingMeshAPs,omitempty"`
	RebootingRootAPs            *int    `json:"rebootingRootAPs,omitempty"`
	TotalAPs                    *int    `json:"totalAPs,omitempty"`
	ZoneID                      *string `json:"zoneId,omitempty"`
	ZoneName                    *string `json:"zoneName,omitempty"`
}

type Ipv4AccessAndCoreSeparation struct {
	DefaultGateway     *string `json:"defaultGateway,omitempty"`
	PrimaryDNSServer   *string `json:"primaryDNSServer,omitempty"`
	SecondaryDNSServer *string `json:"secondaryDNSServer,omitempty"`
}

type Ipv4ClusterInterface struct {
	Gateway    *string `json:"gateway,omitempty"`
	IPAddress  *string `json:"ipAddress,omitempty"`
	IPMode     *string `json:"ipMode,omitempty"`
	SubnetMask *string `json:"subnetMask,omitempty"`
}

type Ipv4ControlInterface struct {
	Gateway    *string `json:"gateway,omitempty"`
	IPAddress  *string `json:"ipAddress,omitempty"`
	IPMode     *string `json:"ipMode,omitempty"`
	NatIP      *string `json:"natIp,omitempty"`
	SubnetMask *string `json:"subnetMask,omitempty"`
}

type Ipv4ManagementInterface struct {
	Gateway    *string `json:"gateway,omitempty"`
	IPAddress  *string `json:"ipAddress,omitempty"`
	IPMode     *string `json:"ipMode,omitempty"`
	SubnetMask *string `json:"subnetMask,omitempty"`
}

type Ipv6AccessAndCoreSeparation struct {
	DefaultGateway     *string `json:"defaultGateway,omitempty"`
	PrimaryDNSServer   *string `json:"primaryDNSServer,omitempty"`
	SecondaryDNSServer *string `json:"secondaryDNSServer,omitempty"`
}

type Ipv6ControlInterface struct {
	Gateway   *string `json:"gateway,omitempty"`
	IPAddress *string `json:"ipAddress,omitempty"`
	IPMode    *string `json:"ipMode,omitempty"`
}

type Ipv6ManagementInterface struct {
	Gateway   *string `json:"gateway,omitempty"`
	IPAddress *string `json:"ipAddress,omitempty"`
	IPMode    *string `json:"ipMode,omitempty"`
}

type Ipv6PrimaryInterface struct {
	Gateway            *string `json:"gateway,omitempty"`
	IPAddress          *string `json:"ipAddress,omitempty"`
	IPMode             *string `json:"ipMode,omitempty"`
	PrimaryDNSServer   *string `json:"primaryDNSServer,omitempty"`
	SecondaryDNSServer *string `json:"secondaryDNSServer,omitempty"`
}

type LwappScgConfiguration struct {
	Acl              *string  `json:"acl,omitempty"`
	ApList           []string `json:"apList,omitempty"`
	NatIPTranslation *bool    `json:"natIpTranslation,omitempty"`
	PasvMaxPort      *int     `json:"pasvMaxPort,omitempty"`
	PasvMinPort      *int     `json:"pasvMinPort,omitempty"`
	Policy           *string  `json:"policy,omitempty"`
}

type ModifyControlPlane struct {
	EnableAccessAndCoreSeparation *bool                        `json:"enableAccessAndCoreSeparation,omitempty"`
	Ipv4AccessAndCoreSeparation   *Ipv4AccessAndCoreSeparation `json:"ipv4AccessAndCoreSeparation,omitempty"`
	Ipv4ClusterInterface          *Ipv4ClusterInterface        `json:"ipv4ClusterInterface,omitempty"`
	Ipv4ControlInterface          *Ipv4ControlInterface        `json:"ipv4ControlInterface,omitempty"`
	Ipv4ManagementInterface       *Ipv4ManagementInterface     `json:"ipv4ManagementInterface,omitempty"`
	Ipv6AccessAndCoreSeparation   *Ipv6AccessAndCoreSeparation `json:"ipv6AccessAndCoreSeparation,omitempty"`
	Ipv6ControlInterface          *Ipv6ControlInterface        `json:"ipv6ControlInterface,omitempty"`
	Ipv6ManagementInterface       *Ipv6ManagementInterface     `json:"ipv6ManagementInterface,omitempty"`
}

type ModifyDataPlane struct {
	InterfaceMode        *string               `json:"interfaceMode,omitempty"`
	Ipv6PrimaryInterface *Ipv6PrimaryInterface `json:"ipv6PrimaryInterface,omitempty"`
	KeepConfig           *bool                 `json:"keepConfig,omitempty"`
	PrimaryInterface     *PrimaryInterface     `json:"primaryInterface,omitempty"`
	SecondaryInterface   *SecondaryInterface   `json:"secondaryInterface,omitempty"`
	StaticRoute          []*StaticRoute        `json:"staticRoute,omitempty"`
}

type ModifyDataPlaneState struct {
	IsDataCenter *bool `json:"isDataCenter,omitempty"`
}

type ModifyGatewayAdvanced struct {
	AllowSessionOnAccountingFail          *bool   `json:"allowSessionOnAccountingFail,omitempty"`
	EcgiInGtpV2                           *bool   `json:"ecgiInGtpV2,omitempty"`
	GtpInterfaceType                      *string `json:"gtpInterfaceType,omitempty"`
	GtpNetworkServiceAcessPointIdentifier *int    `json:"gtpNetworkServiceAcessPointIdentifier,omitempty"`
	ImeiInGtp                             *bool   `json:"imeiInGtp,omitempty"`
	ScgRaiInGtpV2                         *bool   `json:"scgRaiInGtpV2,omitempty"`
	ScgSaiInGtpV2                         *bool   `json:"scgSaiInGtpV2,omitempty"`
	TaiInGtpV2                            *bool   `json:"taiInGtpV2,omitempty"`
}

type ModifyIPSupportType struct {
	IPMode *string `json:"ipMode,omitempty"`
}

type ModifyLwappScg struct {
	ApList           []string `json:"apList,omitempty"`
	NatIPTranslation *bool    `json:"natIpTranslation,omitempty"`
	PasvMaxPort      *int     `json:"pasvMaxPort,omitempty"`
	PasvMinPort      *int     `json:"pasvMinPort,omitempty"`
	Policy           *string  `json:"policy,omitempty"`
}

type ModifySNMPAgent struct {
	SNMPNotificationEnabled *bool          `json:"snmpNotificationEnabled,omitempty"`
	SNMPV2Agent             []*SNMPV2Agent `json:"snmpV2Agent,omitempty"`
	SNMPV3Agent             []*SNMPV3Agent `json:"snmpV3Agent,omitempty"`
}

type ModifySystemTimeSetting struct {
	NtpServer *string `json:"ntpServer,omitempty"`
	Timezone  *string `json:"timezone,omitempty"`
}

type NorthboundInterface struct {
	Password *string `json:"password,omitempty"`
	UserName *string `json:"userName,omitempty"`
}

type PortalLang struct {
	Display *string `json:"display,omitempty"`
	Value   *string `json:"value,omitempty"`
}

type PortalLangList struct {
	FirstIndex *int    `json:"firstIndex,omitempty"`
	HasMore    *bool   `json:"hasMore,omitempty"`
	List       []*List `json:"list,omitempty"`
	TotalCount *int    `json:"totalCount,omitempty"`
}

type PortStatistic struct {
	RxBps     *float64 `json:"rxBps,omitempty"`
	RxBpsMax  *float64 `json:"rxBpsMax,omitempty"`
	RxBpsMin  *float64 `json:"rxBpsMin,omitempty"`
	RxBytes   *float64 `json:"rxBytes,omitempty"`
	RxDropped *float64 `json:"rxDropped,omitempty"`
	RxPackets *float64 `json:"rxPackets,omitempty"`
	TxBps     *float64 `json:"txBps,omitempty"`
	TxBpsMax  *float64 `json:"txBpsMax,omitempty"`
	TxBpsMin  *float64 `json:"txBpsMin,omitempty"`
	TxBytes   *float64 `json:"txBytes,omitempty"`
	TxDropped *float64 `json:"txDropped,omitempty"`
	TxPackets *float64 `json:"txPackets,omitempty"`
}

type PrimaryInterface struct {
	Gateway            *string `json:"gateway,omitempty"`
	IPAddress          *string `json:"ipAddress,omitempty"`
	IPMode             *string `json:"ipMode,omitempty"`
	NatIP              *string `json:"natIp,omitempty"`
	PrimaryDNSServer   *string `json:"primaryDNSServer,omitempty"`
	SecondaryDNSServer *string `json:"secondaryDNSServer,omitempty"`
	SubnetMask         *string `json:"subnetMask,omitempty"`
	Vlan               *string `json:"vlan,omitempty"`
}

type SaveApNumberLimitSettingOfDomain struct {
	DomainID    *string  `json:"domainId,omitempty"`
	NumberLimit *float64 `json:"numberLimit,omitempty"`
	Shared      *bool    `json:"shared,omitempty"`
}

type SaveApNumberLimitSettingOfZone struct {
	DomainID    *string  `json:"domainId,omitempty"`
	NumberLimit *float64 `json:"numberLimit,omitempty"`
	Shared      *bool    `json:"shared,omitempty"`
	ZoneID      *string  `json:"zoneId,omitempty"`
}

type SaveSystemSettings struct {
	ApNumberLimitEnabled          *bool                            `json:"apNumberLimitEnabled,omitempty"`
	ApNumberLimitSettingsOfDomain []*ApNumberLimitSettingsOfDomain `json:"apNumberLimitSettingsOfDomain,omitempty"`
	ApNumberLimitSettingsOfZone   []*ApNumberLimitSettingsOfZone   `json:"apNumberLimitSettingsOfZone,omitempty"`
}

type SecondaryInterface struct {
	IPAddress  *string `json:"ipAddress,omitempty"`
	SubnetMask *string `json:"subnetMask,omitempty"`
	Vlan       *string `json:"vlan,omitempty"`
}

type Sms struct {
	AccountSid *string `json:"accountSid,omitempty"`
	AuthToken  *string `json:"authToken,omitempty"`
	DomainID   *string `json:"domainId,omitempty"`
	Enabled    *int    `json:"enabled,omitempty"`
	From       *string `json:"from,omitempty"`
	ID         *string `json:"id,omitempty"`
	ServerName *string `json:"serverName,omitempty"`
	ServerType *string `json:"serverType,omitempty"`
}

type SmsList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}

type SNMPAgentConfiguration struct {
	SNMPNotificationEnabled *bool          `json:"snmpNotificationEnabled,omitempty"`
	SNMPV2Agent             []*SNMPV2Agent `json:"snmpV2Agent,omitempty"`
	SNMPV3Agent             []*SNMPV3Agent `json:"snmpV3Agent,omitempty"`
}

type StaticRoute struct {
	Gateway        *string `json:"gateway,omitempty"`
	NetworkAddress *string `json:"networkAddress,omitempty"`
	SubnetMask     *string `json:"subnetMask,omitempty"`
}

type StaticRouteList struct {
	StaticRoutes []*StaticRoutes `json:"staticRoutes,omitempty"`
}

type StatisticListType struct {
	Cluster    *Cluster    `json:"cluster,omitempty"`
	Control    *Control    `json:"control,omitempty"`
	CPU        *CPU        `json:"cpu,omitempty"`
	Disk       *Disk       `json:"disk,omitempty"`
	Management *Management `json:"management,omitempty"`
	Memory     *Memory     `json:"memory,omitempty"`
	Port0      *Port0      `json:"port0,omitempty"`
	Port1      *Port1      `json:"port1,omitempty"`
	Port2      *Port2      `json:"port2,omitempty"`
	Port3      *Port3      `json:"port3,omitempty"`
	Port4      *Port4      `json:"port4,omitempty"`
	Port5      *Port5      `json:"port5,omitempty"`
	Timestamp  *float64    `json:"timestamp,omitempty"`
}

type StatisticListTypeCPUType struct {
	MaxPercent *float64 `json:"maxPercent,omitempty"`
	MinPercent *float64 `json:"minPercent,omitempty"`
	Percent    *float64 `json:"percent,omitempty"`
}

type StatisticListTypeDiskType struct {
	Free    *float64 `json:"free,omitempty"`
	MaxFree *float64 `json:"maxFree,omitempty"`
	MinFree *float64 `json:"minFree,omitempty"`
	Total   *float64 `json:"total,omitempty"`
}

type StatisticListTypeMemoryType struct {
	MaxPercent *float64 `json:"maxPercent,omitempty"`
	MinPercent *float64 `json:"minPercent,omitempty"`
	Percent    *float64 `json:"percent,omitempty"`
}

type SystemSettings struct {
	ApNumberLimitEnabled          *bool                            `json:"apNumberLimitEnabled,omitempty"`
	ApNumberLimitSettingsOfDomain []*ApNumberLimitSettingsOfDomain `json:"apNumberLimitSettingsOfDomain,omitempty"`
	ApNumberLimitSettingsOfZone   []*ApNumberLimitSettingsOfZone   `json:"apNumberLimitSettingsOfZone,omitempty"`
}

type SystemTimeSetting struct {
	CurrentSystemTimeString    *string `json:"currentSystemTimeString,omitempty"`
	CurrentSystemTimeUTCString *string `json:"currentSystemTimeUTCString,omitempty"`
	NtpServer                  *string `json:"ntpServer,omitempty"`
	Timezone                   *string `json:"timezone,omitempty"`
}

type UpdateDpMeshTunnelSetting struct {
	Encrypted *bool `json:"encrypted,omitempty"`
}
