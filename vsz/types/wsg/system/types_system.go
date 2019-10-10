package system

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type ApNumberLimitSettingOfDomain struct {
	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// DomainName
	// Domain Name
	DomainName *string `json:"domainName,omitempty"`

	// NumberLimit
	// Number of Limit
	NumberLimit *float64 `json:"numberLimit,omitempty"`

	// Shared
	// Share mode
	Shared *bool `json:"shared,omitempty"`
}

type ApNumberLimitSettingOfZone struct {
	// DomainId
	// Admin Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// DomainName
	// Admin Domain Name
	DomainName *string `json:"domainName,omitempty"`

	// NumberLimit
	// Number of Limit
	NumberLimit *float64 `json:"numberLimit,omitempty"`

	// Shared
	// Share mode
	Shared *bool `json:"shared,omitempty"`

	// ZoneId
	// Zone Id
	ZoneId *string `json:"zoneId,omitempty"`

	// ZoneName
	// Zone Name
	ZoneName *string `json:"zoneName,omitempty"`
}

type AuthenticationKey struct {
	// Key
	// Authentication Key value
	Key *string `json:"key,omitempty"`

	// KeyId
	// Authentication Key ID
	KeyId *int `json:"keyId,omitempty"`

	// KeyType
	// Authentication Key Type
	KeyType *string `json:"keyType,omitempty" validate:"oneof=SHA1 MD5"`
}

type CaptchaSetting struct {
	// CaptchaEnabled
	// Captcha setting
	CaptchaEnabled *bool `json:"captchaEnabled,omitempty"`
}

type ControllerList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ControllerListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type ControllerListType struct {
	// ApVersion
	// AP version
	ApVersion *string `json:"apVersion,omitempty"`

	// ClusterIp
	// Cluster IP
	ClusterIp *string `json:"clusterIp,omitempty"`

	// ClusterIpv6
	// Cluster IPv6
	ClusterIpv6 *string `json:"clusterIpv6,omitempty"`

	// ClusterRole
	// Indicator the role of the controller
	ClusterRole *string `json:"clusterRole,omitempty"`

	// ControlIp
	// Control IP
	ControlIp *string `json:"controlIp,omitempty"`

	// ControlIpv6
	// Control IPv6
	ControlIpv6 *string `json:"controlIpv6,omitempty"`

	// ControlNatIp
	// Control NAT IP address settings
	ControlNatIp *string `json:"controlNatIp,omitempty"`

	// Description
	// Description of the controller
	Description *string `json:"description,omitempty"`

	// HostName
	// Host name of the controller
	HostName *string `json:"hostName,omitempty"`

	// Id
	// Identifier of the controller
	Id *string `json:"id,omitempty"`

	// Mac
	// MAC address of the controller
	Mac *string `json:"mac,omitempty"`

	// ManagementIp
	// Management IP
	ManagementIp *string `json:"managementIp,omitempty"`

	// ManagementIpv6
	// Management IPv6
	ManagementIpv6 *string `json:"managementIpv6,omitempty"`

	// Model
	// Product model
	Model *string `json:"model,omitempty"`

	// Name
	// Name of the controller
	Name *string `json:"name,omitempty"`

	// SerialNumber
	// Serial number of the controller
	SerialNumber *string `json:"serialNumber,omitempty"`

	// UptimeInSec
	// Uptime (in seconds) of the controller
	UptimeInSec *int `json:"uptimeInSec,omitempty"`

	// Version
	// SCG version
	Version *string `json:"version,omitempty"`
}

type ControlPlaneConfiguration struct {
	// IpMode
	// IP support version
	IpMode *string `json:"ipMode,omitempty" validate:"oneof=IPV4 IPV4_IPV6"`

	Ipv4AccessAndCoreSeparation *Ipv4AccessAndCoreSeparation `json:"ipv4AccessAndCoreSeparation,omitempty"`

	Ipv4ClusterInterface *Ipv4ClusterInterface `json:"ipv4ClusterInterface,omitempty"`

	Ipv4ControlInterface *Ipv4ControlInterface `json:"ipv4ControlInterface,omitempty"`

	Ipv4ManagementInterface *Ipv4ManagementInterface `json:"ipv4ManagementInterface,omitempty"`

	Ipv6AccessAndCoreSeparation *Ipv6AccessAndCoreSeparation `json:"ipv6AccessAndCoreSeparation,omitempty"`

	Ipv6ControlInterface *Ipv6ControlInterface `json:"ipv6ControlInterface,omitempty"`

	Ipv6ManagementInterface *Ipv6ManagementInterface `json:"ipv6ManagementInterface,omitempty"`
}

type ControlPlaneList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ControlPlaneListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type ControlPlaneListType struct {
	// ClusterIp
	// Cluster IP
	ClusterIp *string `json:"clusterIp,omitempty"`

	// ClusterRole
	// Cluster Role
	ClusterRole *string `json:"clusterRole,omitempty"`

	// ControlIp
	// Control IP
	ControlIp *string `json:"controlIp,omitempty"`

	// Description
	// Description
	Description *string `json:"description,omitempty"`

	// Firmware
	// Firmware
	Firmware *string `json:"firmware,omitempty"`

	// Id
	// Identifier of the control plane
	Id *string `json:"id,omitempty"`

	// ManagementIp
	// Management IP
	ManagementIp *string `json:"managementIp,omitempty"`

	// Model
	// Model
	Model *string `json:"model,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// NumOfAps
	// Total Connected APs
	NumOfAps *string `json:"numOfAps,omitempty"`

	// SerialNumber
	// Serial Number
	SerialNumber *string `json:"serialNumber,omitempty"`

	// UpTime
	// Uptime
	UpTime *string `json:"upTime,omitempty"`
}

type CpStaticRoute struct {
	// Gateway
	// Gateway
	Gateway *string `json:"gateway,omitempty"`

	// InterfaceMode
	// Interface Mode
	InterfaceMode *string `json:"interfaceMode,omitempty"`

	// Metric
	// Metric
	Metric *int `json:"metric,omitempty" validate:"gte=0,lte=999"`

	// NetworkAddress
	// Network Address
	NetworkAddress *string `json:"networkAddress,omitempty"`

	// SubnetMask
	// Subnet Mask
	SubnetMask *string `json:"subnetMask,omitempty"`
}

type CpUserDefinedInterface struct {
	// Gateway
	// Gateway
	Gateway *string `json:"gateway,omitempty"`

	// IpAddress
	// IP Address
	IpAddress *string `json:"ipAddress,omitempty"`

	// Name
	// Name
	Name *string `json:"name,omitempty"`

	// PhysicalInterface
	// Physical interface
	PhysicalInterface *string `json:"physicalInterface,omitempty" validate:"oneof=Control Management Cluster"`

	// Service
	// Service
	Service *string `json:"service,omitempty" validate:"oneof=NotSpecified Hotspot"`

	// SubnetMask
	// Subnet Mask
	SubnetMask *string `json:"subnetMask,omitempty"`

	// Vlan
	// Vlan
	Vlan *string `json:"vlan,omitempty"`
}

type DataPlaneConfiguration struct {
	InterfaceMode *string `json:"interfaceMode,omitempty" validate:"oneof=SINGLE ACCESS_AND_CORE"`

	Ipv6PrimaryInterface *Ipv6PrimaryInterface `json:"ipv6PrimaryInterface,omitempty"`

	IsDataCenter *bool `json:"isDataCenter,omitempty"`

	KeepConfig *bool `json:"keepConfig,omitempty"`

	PrimaryInterface *PrimaryInterface `json:"primaryInterface,omitempty"`

	SecondaryInterface *SecondaryInterface `json:"secondaryInterface,omitempty"`

	StaticRoute []*StaticRoute `json:"staticRoute,omitempty"`
}

type DataPlaneList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*DataPlaneListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type DataPlaneListType struct {
	BladeName *common.NormalName `json:"bladeName,omitempty"`

	// DpStatus
	// Status
	DpStatus *string `json:"dpStatus,omitempty"`

	// FwVersion
	// Firmware
	FwVersion *string `json:"fwVersion,omitempty"`

	// GreTunnels
	// # of Ruckus GRE Tunnels
	GreTunnels *string `json:"greTunnels,omitempty"`

	// Id
	// Identifier of the data plane
	Id *string `json:"id,omitempty"`

	// Ip
	// IP address
	Ip *string `json:"ip,omitempty"`

	// Ipv6
	// IP address of ipv6
	Ipv6 *string `json:"ipv6,omitempty"`

	// LastSeen
	// Last Seen On
	LastSeen *string `json:"lastSeen,omitempty"`

	// Mac
	// DP MAC Address
	Mac *string `json:"mac,omitempty"`

	// ManagedBy
	// Managed By
	ManagedBy *string `json:"managedBy,omitempty"`

	// Model
	// Model
	Model *string `json:"model,omitempty"`

	// SerialNumber
	// Serial Number
	SerialNumber *string `json:"serialNumber,omitempty"`

	// Type
	// Type
	Type *string `json:"type,omitempty"`

	// Uptime
	// Uptime
	Uptime *string `json:"uptime,omitempty"`
}

type DeleteBulkFtp struct {
	IdList common.IdList `json:"idList,omitempty"`
}

type FriendlyNameLang struct {
	// Display
	// Display name
	Display *string `json:"display,omitempty"`

	// Value
	// value of language used on create Hotspot 2.0 Identity provider (Language in OSU Service Description)
	// profile
	Value *string `json:"value,omitempty"`
}

type FriendlyNameLangList struct {
	// FirstIndex
	// Index of the first FriendlyName of language returned out of the complete FTP list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more FriendlyName of language after the list that is currently displayed
	HasMore *bool `json:"hasMore,omitempty"`

	List []*FriendlyNameLang `json:"list,omitempty"`

	// TotalCount
	// Total count of FriendlyName of language
	TotalCount *int `json:"totalCount,omitempty"`
}

type Ftp struct {
	// CreateDatetime
	// entry create time
	CreateDatetime *int `json:"createDatetime,omitempty"`

	// CreatorUUID
	// creator id
	CreatorUUID *string `json:"creatorUUID,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// FtpHost
	// IP/DN of FTP
	FtpHost *string `json:"ftpHost,omitempty"`

	// FtpName
	// FTP name
	FtpName *string `json:"ftpName,omitempty"`

	// FtpPassword
	// Password for login
	FtpPassword *string `json:"ftpPassword,omitempty"`

	// FtpPort
	// Port used by FTP
	FtpPort *int `json:"ftpPort,omitempty" validate:"gte=21,lte=65535"`

	// FtpProtocol
	// Protocol used
	FtpProtocol *string `json:"ftpProtocol,omitempty" validate:"oneof=FTP SFTP"`

	// FtpRemoteDirectory
	// Destination directory used for file upload
	FtpRemoteDirectory *string `json:"ftpRemoteDirectory,omitempty"`

	// FtpUserName
	// Username for login
	FtpUserName *string `json:"ftpUserName,omitempty"`

	// Id
	// FTP Id
	Id *string `json:"id,omitempty"`

	// LastModifiedBy
	// last modified user
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedOn
	// last modified time
	LastModifiedOn *int `json:"lastModifiedOn,omitempty"`

	// TenantId
	// Tenant Id
	TenantId *string `json:"tenantId,omitempty"`
}

type FtpGlobalSetting struct {
	// Enabled
	// enable logging to remote syslog server
	Enabled *bool `json:"enabled,omitempty"`

	// FtpId
	// Identifier of the FTP Server
	FtpId *string `json:"ftpId,omitempty"`

	// FtpInterval
	// ftpInterval
	FtpInterval *string `json:"ftpInterval,omitempty" validate:"oneof=Hourly"`
}

type FtpList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first FTP returned out of the complete FTP list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more FTPs after the list that is currently displayed
	HasMore *bool `json:"hasMore,omitempty"`

	List []*Ftp `json:"list,omitempty"`

	// TotalCount
	// Total FTP count
	TotalCount *int `json:"totalCount,omitempty"`
}

type FtpTestResponse struct {
	// Data
	// The testing result
	Data *bool `json:"data,omitempty"`

	// Error
	// The error message of http request
	Error *string `json:"error,omitempty"`

	// Extra
	// The extra info
	Extra *string `json:"extra,omitempty"`

	// Success
	// The status of http request
	Success *bool `json:"success,omitempty"`
}

type GatewayAdvanced struct {
	// AllowSessionOnAccountingFail
	// Allow session on accounting fail
	AllowSessionOnAccountingFail *bool `json:"allowSessionOnAccountingFail,omitempty"`

	// EcgiInGtpV2
	// Include ECGI in GTPv2 messages
	EcgiInGtpV2 *bool `json:"ecgiInGtpV2,omitempty"`

	// GtpInterfaceType
	// GTPv2 interface type
	GtpInterfaceType *string `json:"gtpInterfaceType,omitempty" validate:"oneof=S2A S5_S8"`

	// GtpNetworkServiceAcessPointIdentifier
	// GTP network service access point identifier (NSAPI)
	GtpNetworkServiceAcessPointIdentifier *int `json:"gtpNetworkServiceAcessPointIdentifier,omitempty" validate:"gte=0,lte=5"`

	// ImeiInGtp
	// Include IMEI IE in GTP messages
	ImeiInGtp *bool `json:"imeiInGtp,omitempty"`

	// ScgRaiInGtpV2
	// Include SCG-RAI in GTPv2 messages
	ScgRaiInGtpV2 *bool `json:"scgRaiInGtpV2,omitempty"`

	// ScgSaiInGtpV2
	// Include SCG-SAI in GTPv2 messages
	ScgSaiInGtpV2 *bool `json:"scgSaiInGtpV2,omitempty"`

	// TaiInGtpV2
	// Include TAI in GTPv2 messages
	TaiInGtpV2 *bool `json:"taiInGtpV2,omitempty"`
}

type GetDataPlaneMeshTunnelSetting struct {
	// Encrypted
	// Data Plane mesh tunnel encrypted.
	Encrypted *bool `json:"encrypted,omitempty"`
}

type InventoryList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*InventoryListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type InventoryListType struct {
	ApFirmwareVersion *string `json:"apFirmwareVersion,omitempty"`

	Clients *int `json:"clients,omitempty"`

	ConnectedAPs *int `json:"connectedAPs,omitempty"`

	ConnectedDownMeshAPs *int `json:"connectedDownMeshAPs,omitempty"`

	ConnectedeMeshAPs *int `json:"connectedeMeshAPs,omitempty"`

	ConnectedMeshAPs *int `json:"connectedMeshAPs,omitempty"`

	ConnectedMeshDisabledAPs *int `json:"connectedMeshDisabledAPs,omitempty"`

	ConnectedRootAPs *int `json:"connectedRootAPs,omitempty"`

	DisconnectedAPs *int `json:"disconnectedAPs,omitempty"`

	DisconnectedDownMeshAPs *int `json:"disconnectedDownMeshAPs,omitempty"`

	DisconnectedeMeshAPs *int `json:"disconnectedeMeshAPs,omitempty"`

	DisconnectedMeshAPs *int `json:"disconnectedMeshAPs,omitempty"`

	DisconnectedMeshDisabledAPs *int `json:"disconnectedMeshDisabledAPs,omitempty"`

	DisconnectedRootAPs *int `json:"disconnectedRootAPs,omitempty"`

	DiscoveryAPs *int `json:"discoveryAPs,omitempty"`

	MeshEnabled *bool `json:"meshEnabled,omitempty"`

	MeshSSID *string `json:"meshSSID,omitempty"`

	ProvisionedAPs *int `json:"provisionedAPs,omitempty"`

	RebootingDownMeshAPs *int `json:"rebootingDownMeshAPs,omitempty"`

	RebootingeMeshAPs *int `json:"rebootingeMeshAPs,omitempty"`

	RebootingMeshAPs *int `json:"rebootingMeshAPs,omitempty"`

	RebootingRootAPs *int `json:"rebootingRootAPs,omitempty"`

	TotalAPs *int `json:"totalAPs,omitempty"`

	ZoneId *string `json:"zoneId,omitempty"`

	ZoneName *string `json:"zoneName,omitempty"`
}

type Ipv4AccessAndCoreSeparation struct {
	// DefaultGateway
	// Gateway
	DefaultGateway *string `json:"defaultGateway,omitempty" validate:"oneof=Control Management Cluster"`

	// PrimaryDNSServer
	// Primary DNS server
	PrimaryDNSServer *string `json:"primaryDNSServer,omitempty"`

	// SecondaryDNSServer
	// Secondary DNS server
	SecondaryDNSServer *string `json:"secondaryDNSServer,omitempty"`
}

type Ipv4ClusterInterface struct {
	// Gateway
	// Gateway
	Gateway *string `json:"gateway,omitempty"`

	// IpAddress
	// IP address
	IpAddress *string `json:"ipAddress,omitempty"`

	// IpMode
	// IP mode
	IpMode *string `json:"ipMode,omitempty" validate:"oneof=STATIC DHCP"`

	// SubnetMask
	// Subnet mask
	SubnetMask *string `json:"subnetMask,omitempty"`
}

type Ipv4ControlInterface struct {
	// Gateway
	// Gateway
	Gateway *string `json:"gateway,omitempty"`

	// IpAddress
	// IP address
	IpAddress *string `json:"ipAddress,omitempty"`

	// IpMode
	// IP mode
	IpMode *string `json:"ipMode,omitempty" validate:"oneof=STATIC DHCP"`

	// NatIp
	// NAT IP
	NatIp *string `json:"natIp,omitempty"`

	// SubnetMask
	// Subnet mask
	SubnetMask *string `json:"subnetMask,omitempty"`
}

type Ipv4ManagementInterface struct {
	// Gateway
	// Gateway
	Gateway *string `json:"gateway,omitempty"`

	// IpAddress
	// IP address
	IpAddress *string `json:"ipAddress,omitempty"`

	// IpMode
	// IP mode
	IpMode *string `json:"ipMode,omitempty" validate:"oneof=STATIC DHCP"`

	// SubnetMask
	// Subnet mask
	SubnetMask *string `json:"subnetMask,omitempty"`
}

type Ipv6AccessAndCoreSeparation struct {
	// DefaultGateway
	// Gateway
	DefaultGateway *string `json:"defaultGateway,omitempty" validate:"oneof=Control Management Cluster"`

	// PrimaryDNSServer
	// Primary DNS server
	PrimaryDNSServer *string `json:"primaryDNSServer,omitempty"`

	// SecondaryDNSServer
	// Secondary DNS server
	SecondaryDNSServer *string `json:"secondaryDNSServer,omitempty"`
}

type Ipv6ControlInterface struct {
	// Gateway
	// Gateway
	Gateway *string `json:"gateway,omitempty"`

	// IpAddress
	// IP address
	IpAddress *string `json:"ipAddress,omitempty"`

	// IpMode
	// IP mode
	IpMode *string `json:"ipMode,omitempty" validate:"oneof=STATIC AUTO"`
}

type Ipv6ManagementInterface struct {
	// Gateway
	// Gateway
	Gateway *string `json:"gateway,omitempty"`

	// IpAddress
	// IP address
	IpAddress *string `json:"ipAddress,omitempty"`

	// IpMode
	// IP mode
	IpMode *string `json:"ipMode,omitempty" validate:"oneof=STATIC AUTO"`
}

type Ipv6PrimaryInterface struct {
	// Gateway
	// Gateway
	Gateway *string `json:"gateway" validate:"required"`

	// IpAddress
	// IP address
	IpAddress *string `json:"ipAddress" validate:"required"`

	// IpMode
	// IP mode
	IpMode *string `json:"ipMode" validate:"required,oneof=AUTO STATIC"`

	// PrimaryDNSServer
	// Primary DNS server
	PrimaryDNSServer *string `json:"primaryDNSServer,omitempty"`

	// SecondaryDNSServer
	// Secondary DNS server
	SecondaryDNSServer *string `json:"secondaryDNSServer,omitempty"`
}

type Lwapp2scgConfiguration struct {
	// Acl
	// acl  of the lwapp
	Acl *string `json:"acl,omitempty"`

	// ApList
	// accessPoint List  of the lwapp
	ApList []string `json:"apList,omitempty"`

	// NatIpTranslation
	// natIpTranslation of the lwapp
	NatIpTranslation *bool `json:"natIpTranslation,omitempty"`

	// PasvMaxPort
	// pasvMaxPort of the lwapp
	PasvMaxPort *int `json:"pasvMaxPort,omitempty" validate:"gte=16384,lte=65000"`

	// PasvMinPort
	// pasvMinPort of the lwapp
	PasvMinPort *int `json:"pasvMinPort,omitempty" validate:"gte=16384,lte=65000"`

	// Policy
	// policy of the lwapp
	Policy *string `json:"policy,omitempty" validate:"oneof=DENY ACCEPT DENY_ALL ACCEPT_ALL"`
}

type ModifyControlPlane struct {
	// EnableAccessAndCoreSeparation
	// Enable Access & Core Separation
	EnableAccessAndCoreSeparation *bool `json:"enableAccessAndCoreSeparation,omitempty"`

	Ipv4AccessAndCoreSeparation *Ipv4AccessAndCoreSeparation `json:"ipv4AccessAndCoreSeparation,omitempty"`

	Ipv4ClusterInterface *Ipv4ClusterInterface `json:"ipv4ClusterInterface,omitempty"`

	Ipv4ControlInterface *Ipv4ControlInterface `json:"ipv4ControlInterface,omitempty"`

	Ipv4ManagementInterface *Ipv4ManagementInterface `json:"ipv4ManagementInterface,omitempty"`

	Ipv6AccessAndCoreSeparation *Ipv6AccessAndCoreSeparation `json:"ipv6AccessAndCoreSeparation,omitempty"`

	Ipv6ControlInterface *Ipv6ControlInterface `json:"ipv6ControlInterface,omitempty"`

	Ipv6ManagementInterface *Ipv6ManagementInterface `json:"ipv6ManagementInterface,omitempty"`
}

type ModifyDataPlane struct {
	// InterfaceMode
	// Interface mode
	InterfaceMode *string `json:"interfaceMode" validate:"required,oneof=ACCESS_AND_CORE SINGLE"`

	Ipv6PrimaryInterface *Ipv6PrimaryInterface `json:"ipv6PrimaryInterface" validate:"required"`

	KeepConfig *bool `json:"keepConfig,omitempty"`

	PrimaryInterface *PrimaryInterface `json:"primaryInterface" validate:"required"`

	SecondaryInterface *SecondaryInterface `json:"secondaryInterface,omitempty"`

	// StaticRoute
	// Primary(Access) interface
	StaticRoute []*StaticRoute `json:"staticRoute,omitempty"`
}

type ModifyDataPlaneState struct {
	// IsDataCenter
	// Mark this Data Plane as a CALEA Relay
	IsDataCenter *bool `json:"isDataCenter,omitempty"`
}

type ModifyGatewayAdvanced struct {
	// AllowSessionOnAccountingFail
	// Allow session on accounting fail
	AllowSessionOnAccountingFail *bool `json:"allowSessionOnAccountingFail,omitempty"`

	// EcgiInGtpV2
	// Include ECGI in GTPv2 messages
	EcgiInGtpV2 *bool `json:"ecgiInGtpV2,omitempty"`

	// GtpInterfaceType
	// GTPv2 interface type
	GtpInterfaceType *string `json:"gtpInterfaceType,omitempty" validate:"oneof=S2A S5_S8"`

	// GtpNetworkServiceAcessPointIdentifier
	// GTP network service access point identifier (NSAPI)
	GtpNetworkServiceAcessPointIdentifier *int `json:"gtpNetworkServiceAcessPointIdentifier,omitempty" validate:"gte=0,lte=5"`

	// ImeiInGtp
	// Include IMEI IE in GTP messages
	ImeiInGtp *bool `json:"imeiInGtp,omitempty"`

	// ScgRaiInGtpV2
	// Include SCG-RAI in GTPv2 messages
	ScgRaiInGtpV2 *bool `json:"scgRaiInGtpV2,omitempty"`

	// ScgSaiInGtpV2
	// Include SCG-SAI in GTPv2 messages
	ScgSaiInGtpV2 *bool `json:"scgSaiInGtpV2,omitempty"`

	// TaiInGtpV2
	// Include TAI in GTPv2 messages
	TaiInGtpV2 *bool `json:"taiInGtpV2,omitempty"`
}

type ModifyIpSupportType struct {
	// IpMode
	// IP support version
	IpMode *string `json:"ipMode" validate:"required,oneof=IPV4 IPV4_IPV6"`
}

type ModifyLwapp2scg struct {
	// ApList
	// accessPoint List  of the lwapp
	ApList []string `json:"apList,omitempty"`

	// NatIpTranslation
	// natIpTranslation of the lwapp
	NatIpTranslation *bool `json:"natIpTranslation,omitempty"`

	// PasvMaxPort
	// pasvMaxPort of the lwapp
	PasvMaxPort *int `json:"pasvMaxPort,omitempty" validate:"gte=16384,lte=65000"`

	// PasvMinPort
	// pasvMinPort of the lwapp
	PasvMinPort *int `json:"pasvMinPort,omitempty" validate:"gte=16384,lte=65000"`

	// Policy
	// policy of the lwapp
	Policy *string `json:"policy,omitempty" validate:"oneof=DENY ACCEPT DENY_ALL ACCEPT_ALL"`
}

type ModifySnmpAgent struct {
	// SnmpNotificationEnabled
	// Enable SNMP Notifications Globally (If SNMP Notification is disabled globally, no Notification message
	// is sent out.)
	SnmpNotificationEnabled *bool `json:"snmpNotificationEnabled" validate:"required"`

	// SnmpV2Agent
	// Community List of the SNMP V2 Agent.
	SnmpV2Agent []*common.SnmpCommunity `json:"snmpV2Agent,omitempty"`

	// SnmpV3Agent
	// User List of the SNMP V2 Agent.
	SnmpV3Agent []*common.SnmpUser `json:"snmpV3Agent,omitempty"`
}

type ModifySystemTimeSetting struct {
	AuthenticationKey *AuthenticationKey `json:"authenticationKey,omitempty"`

	// NtpServer
	// NtpServer address
	NtpServer *string `json:"ntpServer,omitempty"`

	// Timezone
	// System defined time zone, please refer to the “Overview > Time Zone” list
	Timezone *string `json:"timezone,omitempty"`
}

type NorthboundInterface struct {
	Password *common.ApLoginPassword `json:"password,omitempty"`

	// RadiusAuthType
	// AuthType of the Radius used in Northbound Interface, the value should be "PAP" or "CHAP".
	RadiusAuthType *string `json:"radiusAuthType,omitempty" validate:"oneof=PAP CHAP"`

	UserName *common.ApLoginName `json:"userName,omitempty"`
}

type PortalLang struct {
	// Display
	// Display name
	Display *string `json:"display,omitempty"`

	// Value
	// value of language used on create Hotspot 2.0 Identity provider (Language in OSU Service Description)
	// profile
	Value *string `json:"value,omitempty"`
}

type PortalLangList struct {
	// FirstIndex
	// Index of the first portal names returned out of the complete portal names list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more portal names after the list that is currently displayed
	HasMore *bool `json:"hasMore,omitempty"`

	List []*PortalLang `json:"list,omitempty"`

	// TotalCount
	// Total portal name count
	TotalCount *int `json:"totalCount,omitempty"`
}

type PortStatistic struct {
	// RxBps
	// rxBps
	RxBps *float64 `json:"rxBps,omitempty"`

	// RxBpsMax
	// rxBpsMax
	RxBpsMax *float64 `json:"rxBpsMax,omitempty"`

	// RxBpsMin
	// rxBpsMin
	RxBpsMin *float64 `json:"rxBpsMin,omitempty"`

	// RxBytes
	// rxBytes
	RxBytes *float64 `json:"rxBytes,omitempty"`

	// RxDropped
	// rxDropped
	RxDropped *float64 `json:"rxDropped,omitempty"`

	// RxPackets
	// rxPackets
	RxPackets *float64 `json:"rxPackets,omitempty"`

	// TxBps
	// txBps
	TxBps *float64 `json:"txBps,omitempty"`

	// TxBpsMax
	// txBpsMax
	TxBpsMax *float64 `json:"txBpsMax,omitempty"`

	// TxBpsMin
	// txBpsMin
	TxBpsMin *float64 `json:"txBpsMin,omitempty"`

	// TxBytes
	// txBytes
	TxBytes *float64 `json:"txBytes,omitempty"`

	// TxDropped
	// txDropped
	TxDropped *float64 `json:"txDropped,omitempty"`

	// TxPackets
	// txPackets
	TxPackets *float64 `json:"txPackets,omitempty"`
}

type PrimaryInterface struct {
	// Gateway
	// Gateway
	Gateway *string `json:"gateway" validate:"required"`

	// IpAddress
	// IP address
	IpAddress *string `json:"ipAddress" validate:"required"`

	// IpMode
	// IP mode
	IpMode *string `json:"ipMode" validate:"required,oneof=DHCP STATIC"`

	// NatIp
	// NAT IP
	NatIp *string `json:"natIp,omitempty"`

	// PrimaryDNSServer
	// Primary DNS server
	PrimaryDNSServer *string `json:"primaryDNSServer,omitempty"`

	// SecondaryDNSServer
	// Secondary DNS server
	SecondaryDNSServer *string `json:"secondaryDNSServer,omitempty"`

	// SubnetMask
	// Subnet mask
	SubnetMask *string `json:"subnetMask" validate:"required"`

	// Vlan
	// Vlan
	Vlan *string `json:"vlan,omitempty"`
}

type SaveApNumberLimitSettingOfDomain struct {
	DomainId *string `json:"domainId" validate:"required"`

	NumberLimit *float64 `json:"numberLimit" validate:"required"`

	Shared *bool `json:"shared" validate:"required"`
}

type SaveApNumberLimitSettingOfZone struct {
	DomainId *string `json:"domainId" validate:"required"`

	NumberLimit *float64 `json:"numberLimit" validate:"required"`

	Shared *bool `json:"shared" validate:"required"`

	ZoneId *string `json:"zoneId" validate:"required"`
}

type SaveSystemSettings struct {
	ApNumberLimitEnabled *bool `json:"apNumberLimitEnabled,omitempty"`

	ApNumberLimitSettingsOfDomain []*SaveApNumberLimitSettingOfDomain `json:"apNumberLimitSettingsOfDomain,omitempty"`

	ApNumberLimitSettingsOfZone []*SaveApNumberLimitSettingOfZone `json:"apNumberLimitSettingsOfZone,omitempty"`
}

type SecondaryInterface struct {
	// IpAddress
	// IP address
	IpAddress *string `json:"ipAddress" validate:"required"`

	// SubnetMask
	// Subnet mask
	SubnetMask *string `json:"subnetMask" validate:"required"`

	// Vlan
	// vlan
	Vlan *string `json:"vlan,omitempty"`
}

type Sms struct {
	// AccountSid
	// Account SID
	AccountSid *string `json:"accountSid,omitempty"`

	// AuthToken
	// Auth Token
	AuthToken *string `json:"authToken,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// Enabled
	// Enabled SMS server or not
	Enabled *int `json:"enabled,omitempty" validate:"oneof=0 1"`

	// From
	// From
	From *string `json:"from,omitempty"`

	// Id
	// SMS Id
	Id *string `json:"id,omitempty"`

	// ServerName
	// Server Name
	ServerName *string `json:"serverName,omitempty"`

	// ServerType
	// Server type
	ServerType *string `json:"serverType,omitempty" validate:"oneof=Twilio"`
}

type SmsList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first SMS gateway returned out of the complete SMS gateway list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more SMS gateway after the list that is currently displayed
	HasMore *bool `json:"hasMore,omitempty"`

	List []*Sms `json:"list,omitempty"`

	// TotalCount
	// Total SMS gateway count
	TotalCount *int `json:"totalCount,omitempty"`
}

type SnmpAgentConfiguration struct {
	// SnmpNotificationEnabled
	// Enable SNMP Notifications Globally (If SNMP Notification is disabled globally, no Notification message
	// is sent out.)
	SnmpNotificationEnabled *bool `json:"snmpNotificationEnabled,omitempty"`

	// SnmpV2Agent
	// Community List of the SNMP V2 Agent.
	SnmpV2Agent []*common.SnmpCommunity `json:"snmpV2Agent,omitempty"`

	// SnmpV3Agent
	// User List of the SNMP V2 Agent.
	SnmpV3Agent []*common.SnmpUser `json:"snmpV3Agent,omitempty"`
}

type StaticRoute struct {
	// Gateway
	// Gateway
	Gateway *string `json:"gateway" validate:"required"`

	// NetworkAddress
	// Network address
	NetworkAddress *string `json:"networkAddress" validate:"required"`

	// SubnetMask
	// Subnet mask
	SubnetMask *string `json:"subnetMask" validate:"required"`
}

type StaticRouteList struct {
	// StaticRoutes
	// Static route for Control Plane
	StaticRoutes []*CpStaticRoute `json:"staticRoutes,omitempty"`
}

type StatisticList []*StatisticListType

type StatisticListType struct {
	Cluster *PortStatistic `json:"cluster,omitempty"`

	Control *PortStatistic `json:"control,omitempty"`

	Cpu *StatisticListTypeCpuType `json:"cpu,omitempty"`

	Disk *StatisticListTypeDiskType `json:"disk,omitempty"`

	Management *PortStatistic `json:"management,omitempty"`

	Memory *StatisticListTypeMemoryType `json:"memory,omitempty"`

	Port0 *PortStatistic `json:"port0,omitempty"`

	Port1 *PortStatistic `json:"port1,omitempty"`

	Port2 *PortStatistic `json:"port2,omitempty"`

	Port3 *PortStatistic `json:"port3,omitempty"`

	Port4 *PortStatistic `json:"port4,omitempty"`

	Port5 *PortStatistic `json:"port5,omitempty"`

	// Timestamp
	// timestamp
	Timestamp *float64 `json:"timestamp,omitempty"`
}

type StatisticListTypeCpuType struct {
	// MaxPercent
	// maxPercent
	MaxPercent *float64 `json:"maxPercent,omitempty"`

	// MinPercent
	// minPercent
	MinPercent *float64 `json:"minPercent,omitempty"`

	// Percent
	// percent
	Percent *float64 `json:"percent,omitempty"`
}

type StatisticListTypeDiskType struct {
	// Free
	// free
	Free *float64 `json:"free,omitempty"`

	// MaxFree
	// maxFree
	MaxFree *float64 `json:"maxFree,omitempty"`

	// MinFree
	// minFree
	MinFree *float64 `json:"minFree,omitempty"`

	// Total
	// total
	Total *float64 `json:"total,omitempty"`
}

type StatisticListTypeMemoryType struct {
	// MaxPercent
	// maxPercent
	MaxPercent *float64 `json:"maxPercent,omitempty"`

	// MinPercent
	// minPercent
	MinPercent *float64 `json:"minPercent,omitempty"`

	// Percent
	// percent
	Percent *float64 `json:"percent,omitempty"`
}

type SystemSettings struct {
	// ApNumberLimitEnabled
	// Enabled AP number limit feature or not
	ApNumberLimitEnabled *bool `json:"apNumberLimitEnabled,omitempty"`

	ApNumberLimitSettingsOfDomain []*ApNumberLimitSettingOfDomain `json:"apNumberLimitSettingsOfDomain,omitempty"`

	ApNumberLimitSettingsOfZone []*ApNumberLimitSettingOfZone `json:"apNumberLimitSettingsOfZone,omitempty"`
}

type SystemTimeSetting struct {
	AuthenticationKey *AuthenticationKey `json:"authenticationKey,omitempty"`

	// CurrentSystemTimeString
	// System Time
	CurrentSystemTimeString *string `json:"currentSystemTimeString,omitempty"`

	// CurrentSystemTimeUTCString
	// System UTC Time
	CurrentSystemTimeUTCString *string `json:"currentSystemTimeUTCString,omitempty"`

	// NtpServer
	// NtpServer address
	NtpServer *string `json:"ntpServer,omitempty"`

	// Timezone
	// System defined time zone, please refer to the “Overview > Time Zone” list
	Timezone *string `json:"timezone,omitempty"`
}

type UpdateDpMeshTunnelSetting struct {
	// Encrypted
	// Data Plane mesh tunnel encrypted.
	Encrypted *bool `json:"encrypted,omitempty"`
}

type UserDefinedInterfaceList struct {
	// UserDefinedInterface
	// User defined interface for Control Plane
	UserDefinedInterface []*CpUserDefinedInterface `json:"userDefinedInterface,omitempty"`
}
