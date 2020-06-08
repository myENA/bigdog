package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGSystemService struct {
	apiClient *VSZClient
}

func NewWSGSystemService(c *VSZClient) *WSGSystemService {
	s := new(WSGSystemService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSystemService() *WSGSystemService {
	return NewWSGSystemService(ss.apiClient)
}

type WSGSystemApMacOUI struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	Oui *WSGCommonOui `json:"oui,omitempty"`
}

func NewWSGSystemApMacOUI() *WSGSystemApMacOUI {
	m := new(WSGSystemApMacOUI)
	return m
}

type WSGSystemApMacOUIList struct {
	// FirstIndex
	// Index of the first AP MAC OUI returned out of the complete AP MAC OUI list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more AP MAC OUIs after the list that is currently displayed
	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSystemApMacOUI `json:"list,omitempty"`

	// TotalCount
	// Total AP MAC OUI count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGSystemApMacOUIList() *WSGSystemApMacOUIList {
	m := new(WSGSystemApMacOUIList)
	return m
}

type WSGSystemApNumberLimitSettingOfDomain struct {
	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// DomainName
	// Domain Name
	DomainName *string `json:"domainName,omitempty"`

	// NumberLimit
	// Number of Limit
	NumberLimit *int `json:"numberLimit,omitempty"`

	// Shared
	// Share mode
	Shared *bool `json:"shared,omitempty"`
}

func NewWSGSystemApNumberLimitSettingOfDomain() *WSGSystemApNumberLimitSettingOfDomain {
	m := new(WSGSystemApNumberLimitSettingOfDomain)
	return m
}

type WSGSystemApNumberLimitSettingOfZone struct {
	// DomainId
	// Admin Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// DomainName
	// Admin Domain Name
	DomainName *string `json:"domainName,omitempty"`

	// NumberLimit
	// Number of Limit
	NumberLimit *int `json:"numberLimit,omitempty"`

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

func NewWSGSystemApNumberLimitSettingOfZone() *WSGSystemApNumberLimitSettingOfZone {
	m := new(WSGSystemApNumberLimitSettingOfZone)
	return m
}

type WSGSystemAuthenticationKey struct {
	// Key
	// Authentication Key value
	Key *string `json:"key,omitempty"`

	// KeyId
	// Authentication Key ID
	KeyId *int `json:"keyId,omitempty"`

	// KeyType
	// Authentication Key Type
	// Constraints:
	//    - oneof:[SHA1,MD5]
	KeyType *string `json:"keyType,omitempty"`
}

func NewWSGSystemAuthenticationKey() *WSGSystemAuthenticationKey {
	m := new(WSGSystemAuthenticationKey)
	return m
}

type WSGSystemControllerList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSystemControllerListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGSystemControllerList() *WSGSystemControllerList {
	m := new(WSGSystemControllerList)
	return m
}

type WSGSystemControllerListType struct {
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

	ReservedPorts []*WSGSystemReservedPort `json:"reservedPorts,omitempty"`

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

func NewWSGSystemControllerListType() *WSGSystemControllerListType {
	m := new(WSGSystemControllerListType)
	return m
}

type WSGSystemControlPlaneConfiguration struct {
	// IpMode
	// IP support version
	// Constraints:
	//    - oneof:[IPV4,IPV4_IPV6]
	IpMode *string `json:"ipMode,omitempty"`

	Ipv4AccessAndCoreSeparation *WSGSystemIpv4AccessAndCoreSeparation `json:"ipv4AccessAndCoreSeparation,omitempty"`

	Ipv4ClusterInterface *WSGSystemIpv4ClusterInterface `json:"ipv4ClusterInterface,omitempty"`

	Ipv4ControlInterface *WSGSystemIpv4ControlInterface `json:"ipv4ControlInterface,omitempty"`

	Ipv4ManagementInterface *WSGSystemIpv4ManagementInterface `json:"ipv4ManagementInterface,omitempty"`

	Ipv6AccessAndCoreSeparation *WSGSystemIpv6AccessAndCoreSeparation `json:"ipv6AccessAndCoreSeparation,omitempty"`

	Ipv6ControlInterface *WSGSystemIpv6ControlInterface `json:"ipv6ControlInterface,omitempty"`

	Ipv6ManagementInterface *WSGSystemIpv6ManagementInterface `json:"ipv6ManagementInterface,omitempty"`
}

func NewWSGSystemControlPlaneConfiguration() *WSGSystemControlPlaneConfiguration {
	m := new(WSGSystemControlPlaneConfiguration)
	return m
}

type WSGSystemControlPlaneInterface struct {
	// Id
	// Interface Id
	Id *string `json:"id,omitempty"`

	// Name
	// Interface Name
	Name *string `json:"name,omitempty"`
}

func NewWSGSystemControlPlaneInterface() *WSGSystemControlPlaneInterface {
	m := new(WSGSystemControlPlaneInterface)
	return m
}

type WSGSystemControlPlaneInterfaceList struct {
	// ControlPlaneInterfaces
	// Interface list
	ControlPlaneInterfaces []*WSGSystemControlPlaneInterface `json:"controlPlaneInterfaces,omitempty"`
}

func NewWSGSystemControlPlaneInterfaceList() *WSGSystemControlPlaneInterfaceList {
	m := new(WSGSystemControlPlaneInterfaceList)
	return m
}

type WSGSystemControlPlaneList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSystemControlPlaneListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGSystemControlPlaneList() *WSGSystemControlPlaneList {
	m := new(WSGSystemControlPlaneList)
	return m
}

type WSGSystemControlPlaneListType struct {
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

	Name *WSGCommonNormalName `json:"name,omitempty"`

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

func NewWSGSystemControlPlaneListType() *WSGSystemControlPlaneListType {
	m := new(WSGSystemControlPlaneListType)
	return m
}

type WSGSystemCpStaticRoute struct {
	// Gateway
	// Gateway
	Gateway *string `json:"gateway,omitempty"`

	// InterfaceMode
	// Interface Type or UserDefined Interface Name
	InterfaceMode *string `json:"interfaceMode,omitempty"`

	// Metric
	// Metric
	// Constraints:
	//    - min:0
	//    - max:999
	Metric *int `json:"metric,omitempty"`

	// NetworkAddress
	// Network Address
	NetworkAddress *string `json:"networkAddress,omitempty"`

	// SubnetMask
	// Subnet Mask
	SubnetMask *string `json:"subnetMask,omitempty"`
}

func NewWSGSystemCpStaticRoute() *WSGSystemCpStaticRoute {
	m := new(WSGSystemCpStaticRoute)
	return m
}

type WSGSystemCpUserDefinedInterface struct {
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
	// Constraints:
	//    - oneof:[Control,Management,Cluster]
	PhysicalInterface *string `json:"physicalInterface,omitempty"`

	// Service
	// Service
	// Constraints:
	//    - oneof:[NotSpecified,Hotspot,SecondManagementInterface]
	Service *string `json:"service,omitempty"`

	// SubnetMask
	// Subnet Mask
	SubnetMask *string `json:"subnetMask,omitempty"`

	// Vlan
	// Vlan
	Vlan *string `json:"vlan,omitempty"`
}

func NewWSGSystemCpUserDefinedInterface() *WSGSystemCpUserDefinedInterface {
	m := new(WSGSystemCpUserDefinedInterface)
	return m
}

type WSGSystemCreateApMacOUI struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Oui
	// Constraints:
	//    - required
	Oui *WSGCommonOui `json:"oui"`
}

func NewWSGSystemCreateApMacOUI() *WSGSystemCreateApMacOUI {
	m := new(WSGSystemCreateApMacOUI)
	return m
}

type WSGSystemDataPlaneConfiguration struct {
	// InterfaceMode
	// Interface mode
	// Constraints:
	//    - oneof:[SINGLE,ACCESS_AND_CORE]
	InterfaceMode *string `json:"interfaceMode,omitempty"`

	Ipv6PrimaryInterface *WSGSystemIpv6PrimaryInterface `json:"ipv6PrimaryInterface,omitempty"`

	IsDataCenter *bool `json:"isDataCenter,omitempty"`

	KeepConfig *bool `json:"keepConfig,omitempty"`

	PrimaryInterface *WSGSystemPrimaryInterface `json:"primaryInterface,omitempty"`

	SecondaryInterface *WSGSystemSecondaryInterface `json:"secondaryInterface,omitempty"`

	// StaticRoute
	// Primary(Access) interface
	StaticRoute []*WSGSystemStaticRoute `json:"staticRoute,omitempty"`
}

func NewWSGSystemDataPlaneConfiguration() *WSGSystemDataPlaneConfiguration {
	m := new(WSGSystemDataPlaneConfiguration)
	return m
}

type WSGSystemDataPlaneList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSystemDataPlaneListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGSystemDataPlaneList() *WSGSystemDataPlaneList {
	m := new(WSGSystemDataPlaneList)
	return m
}

type WSGSystemDataPlaneListType struct {
	BladeName *WSGCommonNormalName `json:"bladeName,omitempty"`

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
	// Identifier of data plane
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

func NewWSGSystemDataPlaneListType() *WSGSystemDataPlaneListType {
	m := new(WSGSystemDataPlaneListType)
	return m
}

type WSGSystemDeleteBulkFtp struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGSystemDeleteBulkFtp() *WSGSystemDeleteBulkFtp {
	m := new(WSGSystemDeleteBulkFtp)
	return m
}

type WSGSystemFriendlyNameLang struct {
	// Display
	// Display name
	Display *string `json:"display,omitempty"`

	// Value
	// value of language used on create Hotspot 2.0 Identity provider (Language in OSU Service Description) profile
	Value *string `json:"value,omitempty"`
}

func NewWSGSystemFriendlyNameLang() *WSGSystemFriendlyNameLang {
	m := new(WSGSystemFriendlyNameLang)
	return m
}

type WSGSystemFriendlyNameLangList struct {
	// FirstIndex
	// Index of the first FriendlyName of language returned out of the complete FTP list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more FriendlyName of language after the list that is currently displayed
	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSystemFriendlyNameLang `json:"list,omitempty"`

	// TotalCount
	// Total count of FriendlyName of language
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGSystemFriendlyNameLangList() *WSGSystemFriendlyNameLangList {
	m := new(WSGSystemFriendlyNameLangList)
	return m
}

type WSGSystemFtp struct {
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
	// Constraints:
	//    - min:20
	//    - max:65535
	FtpPort *int `json:"ftpPort,omitempty"`

	// FtpProtocol
	// Protocol used
	// Constraints:
	//    - oneof:[FTP,SFTP]
	FtpProtocol *string `json:"ftpProtocol,omitempty"`

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

func NewWSGSystemFtp() *WSGSystemFtp {
	m := new(WSGSystemFtp)
	return m
}

type WSGSystemFtpGlobalSetting struct {
	// Enabled
	// enable logging to remote syslog server
	Enabled *bool `json:"enabled,omitempty"`

	// FtpId
	// Identifier of the FTP Server
	FtpId *string `json:"ftpId,omitempty"`

	// FtpInterval
	// ftpInterval
	// Constraints:
	//    - oneof:[Hourly]
	FtpInterval *string `json:"ftpInterval,omitempty"`
}

func NewWSGSystemFtpGlobalSetting() *WSGSystemFtpGlobalSetting {
	m := new(WSGSystemFtpGlobalSetting)
	return m
}

type WSGSystemFtpList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first FTP returned out of the complete FTP list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more FTPs after the list that is currently displayed
	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSystemFtp `json:"list,omitempty"`

	// TotalCount
	// Total FTP count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGSystemFtpList() *WSGSystemFtpList {
	m := new(WSGSystemFtpList)
	return m
}

type WSGSystemFtpTestResponse struct {
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

func NewWSGSystemFtpTestResponse() *WSGSystemFtpTestResponse {
	m := new(WSGSystemFtpTestResponse)
	return m
}

type WSGSystemGatewayAdvanced struct {
	// AllowSessionOnAccountingFail
	// Allow session on accounting fail
	AllowSessionOnAccountingFail *bool `json:"allowSessionOnAccountingFail,omitempty"`

	// EcgiInGtpV2
	// Include ECGI in GTPv2 messages
	EcgiInGtpV2 *bool `json:"ecgiInGtpV2,omitempty"`

	// GtpInterfaceType
	// GTPv2 interface type
	// Constraints:
	//    - oneof:[S2A,S5_S8]
	GtpInterfaceType *string `json:"gtpInterfaceType,omitempty"`

	// GtpNetworkServiceAcessPointIdentifier
	// GTP network service access point identifier (NSAPI)
	// Constraints:
	//    - min:0
	//    - max:5
	GtpNetworkServiceAcessPointIdentifier *int `json:"gtpNetworkServiceAcessPointIdentifier,omitempty"`

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

func NewWSGSystemGatewayAdvanced() *WSGSystemGatewayAdvanced {
	m := new(WSGSystemGatewayAdvanced)
	return m
}

type WSGSystemGetDataPlaneMeshTunnelSetting struct {
	// Encrypted
	// Data Plane mesh tunnel encrypted.
	Encrypted *bool `json:"encrypted,omitempty"`
}

func NewWSGSystemGetDataPlaneMeshTunnelSetting() *WSGSystemGetDataPlaneMeshTunnelSetting {
	m := new(WSGSystemGetDataPlaneMeshTunnelSetting)
	return m
}

type WSGSystemInventoryList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSystemInventoryListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGSystemInventoryList() *WSGSystemInventoryList {
	m := new(WSGSystemInventoryList)
	return m
}

type WSGSystemInventoryListType struct {
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

func NewWSGSystemInventoryListType() *WSGSystemInventoryListType {
	m := new(WSGSystemInventoryListType)
	return m
}

type WSGSystemIpv4AccessAndCoreSeparation struct {
	// DefaultGateway
	// Gateway
	// Constraints:
	//    - oneof:[Control,Management,Cluster]
	DefaultGateway *string `json:"defaultGateway,omitempty"`

	// PrimaryDNSServer
	// Primary DNS server
	PrimaryDNSServer *string `json:"primaryDNSServer,omitempty"`

	// SecondaryDNSServer
	// Secondary DNS server
	SecondaryDNSServer *string `json:"secondaryDNSServer,omitempty"`
}

func NewWSGSystemIpv4AccessAndCoreSeparation() *WSGSystemIpv4AccessAndCoreSeparation {
	m := new(WSGSystemIpv4AccessAndCoreSeparation)
	return m
}

type WSGSystemIpv4ClusterInterface struct {
	// Gateway
	// Gateway
	Gateway *string `json:"gateway,omitempty"`

	// IpAddress
	// IP address
	IpAddress *string `json:"ipAddress,omitempty"`

	// IpMode
	// IP mode
	// Constraints:
	//    - oneof:[STATIC,DHCP]
	IpMode *string `json:"ipMode,omitempty"`

	// SubnetMask
	// Subnet mask
	SubnetMask *string `json:"subnetMask,omitempty"`
}

func NewWSGSystemIpv4ClusterInterface() *WSGSystemIpv4ClusterInterface {
	m := new(WSGSystemIpv4ClusterInterface)
	return m
}

type WSGSystemIpv4ControlInterface struct {
	// Gateway
	// Gateway
	Gateway *string `json:"gateway,omitempty"`

	// IpAddress
	// IP address
	IpAddress *string `json:"ipAddress,omitempty"`

	// IpMode
	// IP mode
	// Constraints:
	//    - oneof:[STATIC,DHCP]
	IpMode *string `json:"ipMode,omitempty"`

	// NatIp
	// NAT IP
	NatIp *string `json:"natIp,omitempty"`

	// SubnetMask
	// Subnet mask
	SubnetMask *string `json:"subnetMask,omitempty"`
}

func NewWSGSystemIpv4ControlInterface() *WSGSystemIpv4ControlInterface {
	m := new(WSGSystemIpv4ControlInterface)
	return m
}

type WSGSystemIpv4ManagementInterface struct {
	// Gateway
	// Gateway
	Gateway *string `json:"gateway,omitempty"`

	// IpAddress
	// IP address
	IpAddress *string `json:"ipAddress,omitempty"`

	// IpMode
	// IP mode
	// Constraints:
	//    - oneof:[STATIC,DHCP]
	IpMode *string `json:"ipMode,omitempty"`

	// SubnetMask
	// Subnet mask
	SubnetMask *string `json:"subnetMask,omitempty"`
}

func NewWSGSystemIpv4ManagementInterface() *WSGSystemIpv4ManagementInterface {
	m := new(WSGSystemIpv4ManagementInterface)
	return m
}

type WSGSystemIpv6AccessAndCoreSeparation struct {
	// DefaultGateway
	// Gateway
	// Constraints:
	//    - oneof:[Control,Management,Cluster]
	DefaultGateway *string `json:"defaultGateway,omitempty"`

	// PrimaryDNSServer
	// Primary DNS server
	PrimaryDNSServer *string `json:"primaryDNSServer,omitempty"`

	// SecondaryDNSServer
	// Secondary DNS server
	SecondaryDNSServer *string `json:"secondaryDNSServer,omitempty"`
}

func NewWSGSystemIpv6AccessAndCoreSeparation() *WSGSystemIpv6AccessAndCoreSeparation {
	m := new(WSGSystemIpv6AccessAndCoreSeparation)
	return m
}

type WSGSystemIpv6ControlInterface struct {
	// Gateway
	// Gateway
	Gateway *string `json:"gateway,omitempty"`

	// IpAddress
	// IP address
	IpAddress *string `json:"ipAddress,omitempty"`

	// IpMode
	// IP mode
	// Constraints:
	//    - oneof:[STATIC,AUTO]
	IpMode *string `json:"ipMode,omitempty"`
}

func NewWSGSystemIpv6ControlInterface() *WSGSystemIpv6ControlInterface {
	m := new(WSGSystemIpv6ControlInterface)
	return m
}

type WSGSystemIpv6ManagementInterface struct {
	// Gateway
	// Gateway
	Gateway *string `json:"gateway,omitempty"`

	// IpAddress
	// IP address
	IpAddress *string `json:"ipAddress,omitempty"`

	// IpMode
	// IP mode
	// Constraints:
	//    - oneof:[STATIC,AUTO]
	IpMode *string `json:"ipMode,omitempty"`
}

func NewWSGSystemIpv6ManagementInterface() *WSGSystemIpv6ManagementInterface {
	m := new(WSGSystemIpv6ManagementInterface)
	return m
}

type WSGSystemIpv6PrimaryInterface struct {
	// Gateway
	// Gateway
	// Constraints:
	//    - required
	Gateway *string `json:"gateway"`

	// IpAddress
	// IP address
	// Constraints:
	//    - required
	IpAddress *string `json:"ipAddress"`

	// IpMode
	// IP mode
	// Constraints:
	//    - required
	//    - oneof:[STATIC,AUTO]
	IpMode *string `json:"ipMode"`

	// PrimaryDNSServer
	// Primary DNS server
	PrimaryDNSServer *string `json:"primaryDNSServer,omitempty"`

	// SecondaryDNSServer
	// Secondary DNS server
	SecondaryDNSServer *string `json:"secondaryDNSServer,omitempty"`
}

func NewWSGSystemIpv6PrimaryInterface() *WSGSystemIpv6PrimaryInterface {
	m := new(WSGSystemIpv6PrimaryInterface)
	return m
}

type WSGSystemLwapp2scgConfiguration struct {
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
	// Constraints:
	//    - min:16384
	//    - max:65000
	PasvMaxPort *int `json:"pasvMaxPort,omitempty"`

	// PasvMinPort
	// pasvMinPort of the lwapp
	// Constraints:
	//    - min:16384
	//    - max:65000
	PasvMinPort *int `json:"pasvMinPort,omitempty"`

	// Policy
	// policy of the lwapp
	// Constraints:
	//    - oneof:[DENY,ACCEPT,DENY_ALL,ACCEPT_ALL]
	Policy *string `json:"policy,omitempty"`
}

func NewWSGSystemLwapp2scgConfiguration() *WSGSystemLwapp2scgConfiguration {
	m := new(WSGSystemLwapp2scgConfiguration)
	return m
}

type WSGSystemModifyControlPlane struct {
	// EnableAccessAndCoreSeparation
	// Enable Access & Core Separation
	EnableAccessAndCoreSeparation *bool `json:"enableAccessAndCoreSeparation,omitempty"`

	Ipv4AccessAndCoreSeparation *WSGSystemIpv4AccessAndCoreSeparation `json:"ipv4AccessAndCoreSeparation,omitempty"`

	Ipv4ClusterInterface *WSGSystemIpv4ClusterInterface `json:"ipv4ClusterInterface,omitempty"`

	Ipv4ControlInterface *WSGSystemIpv4ControlInterface `json:"ipv4ControlInterface,omitempty"`

	Ipv4ManagementInterface *WSGSystemIpv4ManagementInterface `json:"ipv4ManagementInterface,omitempty"`

	Ipv6AccessAndCoreSeparation *WSGSystemIpv6AccessAndCoreSeparation `json:"ipv6AccessAndCoreSeparation,omitempty"`

	Ipv6ControlInterface *WSGSystemIpv6ControlInterface `json:"ipv6ControlInterface,omitempty"`

	Ipv6ManagementInterface *WSGSystemIpv6ManagementInterface `json:"ipv6ManagementInterface,omitempty"`
}

func NewWSGSystemModifyControlPlane() *WSGSystemModifyControlPlane {
	m := new(WSGSystemModifyControlPlane)
	return m
}

type WSGSystemModifyCPStaticRoute struct {
	// StaticRoutes
	// Static route for control plane.
	StaticRoutes []*WSGSystemCpStaticRoute `json:"staticRoutes,omitempty"`
}

func NewWSGSystemModifyCPStaticRoute() *WSGSystemModifyCPStaticRoute {
	m := new(WSGSystemModifyCPStaticRoute)
	return m
}

type WSGSystemModifyCPUserDefinedInterface struct {
	// UserDefinedInterface
	// User defined interface for Control Plane
	UserDefinedInterface []*WSGSystemCpUserDefinedInterface `json:"userDefinedInterface,omitempty"`
}

func NewWSGSystemModifyCPUserDefinedInterface() *WSGSystemModifyCPUserDefinedInterface {
	m := new(WSGSystemModifyCPUserDefinedInterface)
	return m
}

type WSGSystemModifyDataPlane struct {
	// InterfaceMode
	// Interface mode
	// Constraints:
	//    - required
	//    - oneof:[SINGLE,ACCESS_AND_CORE]
	InterfaceMode *string `json:"interfaceMode"`

	// Ipv6PrimaryInterface
	// Constraints:
	//    - required
	Ipv6PrimaryInterface *WSGSystemIpv6PrimaryInterface `json:"ipv6PrimaryInterface"`

	IsDataCenter *bool `json:"isDataCenter,omitempty"`

	KeepConfig *bool `json:"keepConfig,omitempty"`

	// PrimaryInterface
	// Constraints:
	//    - required
	PrimaryInterface *WSGSystemPrimaryInterface `json:"primaryInterface"`

	SecondaryInterface *WSGSystemSecondaryInterface `json:"secondaryInterface,omitempty"`

	// StaticRoute
	// Primary(Access) interface
	StaticRoute []*WSGSystemStaticRoute `json:"staticRoute,omitempty"`
}

func NewWSGSystemModifyDataPlane() *WSGSystemModifyDataPlane {
	m := new(WSGSystemModifyDataPlane)
	return m
}

type WSGSystemModifyDataPlaneState struct {
	// IsDataCenter
	// Mark this Data Plane as a CALEA Relay
	IsDataCenter *bool `json:"isDataCenter,omitempty"`
}

func NewWSGSystemModifyDataPlaneState() *WSGSystemModifyDataPlaneState {
	m := new(WSGSystemModifyDataPlaneState)
	return m
}

type WSGSystemModifyGatewayAdvanced struct {
	// AllowSessionOnAccountingFail
	// Allow session on accounting fail
	AllowSessionOnAccountingFail *bool `json:"allowSessionOnAccountingFail,omitempty"`

	// EcgiInGtpV2
	// Include ECGI in GTPv2 messages
	EcgiInGtpV2 *bool `json:"ecgiInGtpV2,omitempty"`

	// GtpInterfaceType
	// GTPv2 interface type
	// Constraints:
	//    - oneof:[S2A,S5_S8]
	GtpInterfaceType *string `json:"gtpInterfaceType,omitempty"`

	// GtpNetworkServiceAcessPointIdentifier
	// GTP network service access point identifier (NSAPI)
	// Constraints:
	//    - min:0
	//    - max:5
	GtpNetworkServiceAcessPointIdentifier *int `json:"gtpNetworkServiceAcessPointIdentifier,omitempty"`

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

func NewWSGSystemModifyGatewayAdvanced() *WSGSystemModifyGatewayAdvanced {
	m := new(WSGSystemModifyGatewayAdvanced)
	return m
}

type WSGSystemModifyIpSupportType struct {
	// IpMode
	// IP support version
	// Constraints:
	//    - required
	//    - oneof:[IPV4,IPV4_IPV6]
	IpMode *string `json:"ipMode"`
}

func NewWSGSystemModifyIpSupportType() *WSGSystemModifyIpSupportType {
	m := new(WSGSystemModifyIpSupportType)
	return m
}

type WSGSystemModifyLwapp2scg struct {
	// ApList
	// accessPoint List  of the lwapp
	ApList []string `json:"apList,omitempty"`

	// NatIpTranslation
	// natIpTranslation of the lwapp
	NatIpTranslation *bool `json:"natIpTranslation,omitempty"`

	// PasvMaxPort
	// pasvMaxPort of the lwapp
	// Constraints:
	//    - min:16384
	//    - max:65000
	PasvMaxPort *int `json:"pasvMaxPort,omitempty"`

	// PasvMinPort
	// pasvMinPort of the lwapp
	// Constraints:
	//    - min:16384
	//    - max:65000
	PasvMinPort *int `json:"pasvMinPort,omitempty"`

	// Policy
	// policy of the lwapp
	// Constraints:
	//    - oneof:[DENY,ACCEPT,DENY_ALL,ACCEPT_ALL]
	Policy *string `json:"policy,omitempty"`
}

func NewWSGSystemModifyLwapp2scg() *WSGSystemModifyLwapp2scg {
	m := new(WSGSystemModifyLwapp2scg)
	return m
}

type WSGSystemModifySnmpAgent struct {
	// SnmpNotificationEnabled
	// Enable SNMP Notifications Globally (If SNMP Notification is disabled globally, no Notification message is sent out.)
	// Constraints:
	//    - required
	SnmpNotificationEnabled *bool `json:"snmpNotificationEnabled"`

	// SnmpV2Agent
	// Community List of the SNMP V2 Agent.
	SnmpV2Agent []*WSGCommonSnmpCommunity `json:"snmpV2Agent,omitempty"`

	// SnmpV3Agent
	// User List of the SNMP V2 Agent.
	SnmpV3Agent []*WSGCommonSnmpUser `json:"snmpV3Agent,omitempty"`
}

func NewWSGSystemModifySnmpAgent() *WSGSystemModifySnmpAgent {
	m := new(WSGSystemModifySnmpAgent)
	return m
}

type WSGSystemModifySystemTimeSetting struct {
	AuthenticationKey *WSGSystemAuthenticationKey `json:"authenticationKey,omitempty"`

	// NtpServer
	// NtpServer address
	NtpServer *string `json:"ntpServer,omitempty"`

	SecondaryAuthenticationKey *WSGSystemAuthenticationKey `json:"secondaryAuthenticationKey,omitempty"`

	// SecondaryNtpServer
	// Backup NtpServer address
	SecondaryNtpServer *string `json:"secondaryNtpServer,omitempty"`

	// Timezone
	// System defined time zone, please refer to the 'Overview > Time Zone' list
	Timezone *string `json:"timezone,omitempty"`
}

func NewWSGSystemModifySystemTimeSetting() *WSGSystemModifySystemTimeSetting {
	m := new(WSGSystemModifySystemTimeSetting)
	return m
}

type WSGSystemNorthboundInterface struct {
	Password *WSGCommonApLoginPassword `json:"password,omitempty"`

	// RadiusAuthType
	// AuthType of the Radius used in Northbound Interface, the value should be "PAP" or "CHAP".
	// Constraints:
	//    - oneof:[PAP,CHAP]
	RadiusAuthType *string `json:"radiusAuthType,omitempty"`

	UserName *WSGCommonApLoginName `json:"userName,omitempty"`
}

func NewWSGSystemNorthboundInterface() *WSGSystemNorthboundInterface {
	m := new(WSGSystemNorthboundInterface)
	return m
}

type WSGSystemNtpServerValidation struct {
	AuthenticationKey *WSGSystemAuthenticationKey `json:"authenticationKey,omitempty"`

	// NtpServer
	// NTP Server address for validation
	// Constraints:
	//    - required
	NtpServer *string `json:"ntpServer"`
}

func NewWSGSystemNtpServerValidation() *WSGSystemNtpServerValidation {
	m := new(WSGSystemNtpServerValidation)
	return m
}

type WSGSystemNtpServerValidationMessage struct {
	// Message
	// NTP Server Validation Message
	Message *string `json:"message,omitempty"`
}

func NewWSGSystemNtpServerValidationMessage() *WSGSystemNtpServerValidationMessage {
	m := new(WSGSystemNtpServerValidationMessage)
	return m
}

type WSGSystemPortalLang struct {
	// Display
	// Display name
	Display *string `json:"display,omitempty"`

	// Value
	// value of language used on create Hotspot 2.0 Identity provider (Language in OSU Service Description) profile
	Value *string `json:"value,omitempty"`
}

func NewWSGSystemPortalLang() *WSGSystemPortalLang {
	m := new(WSGSystemPortalLang)
	return m
}

type WSGSystemPortalLangList struct {
	// FirstIndex
	// Index of the first portal names returned out of the complete portal names list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more portal names after the list that is currently displayed
	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSystemPortalLang `json:"list,omitempty"`

	// TotalCount
	// Total portal name count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGSystemPortalLangList() *WSGSystemPortalLangList {
	m := new(WSGSystemPortalLangList)
	return m
}

type WSGSystemPortStatistic struct {
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

func NewWSGSystemPortStatistic() *WSGSystemPortStatistic {
	m := new(WSGSystemPortStatistic)
	return m
}

type WSGSystemPrimaryInterface struct {
	// Gateway
	// Gateway
	// Constraints:
	//    - required
	Gateway *string `json:"gateway"`

	// IpAddress
	// IP address
	// Constraints:
	//    - required
	IpAddress *string `json:"ipAddress"`

	// IpMode
	// IP mode
	// Constraints:
	//    - required
	//    - oneof:[STATIC,DHCP]
	IpMode *string `json:"ipMode"`

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
	// Constraints:
	//    - required
	SubnetMask *string `json:"subnetMask"`

	// Vlan
	// Vlan
	Vlan *string `json:"vlan,omitempty"`
}

func NewWSGSystemPrimaryInterface() *WSGSystemPrimaryInterface {
	m := new(WSGSystemPrimaryInterface)
	return m
}

type WSGSystemReservedPort struct {
	// BindingInterface
	// The binding interfaces, ["Control", "Cluster", "Management"]
	BindingInterface *string `json:"bindingInterface,omitempty"`

	// Description
	// The purpose of reserved port range
	Description *string `json:"description,omitempty"`

	// Destination
	// The traffic destination (IP Address)
	Destination *string `json:"destination,omitempty"`

	// From
	// Rule from System or User
	From *string `json:"from,omitempty"`

	// PortRange
	// Reserved port range for SZ service
	PortRange *string `json:"portRange,omitempty"`

	// Protocol
	// TCP/UDP
	Protocol *string `json:"protocol,omitempty"`

	// TrafficDirection
	// Inbound/Outbound
	TrafficDirection *string `json:"trafficDirection,omitempty"`
}

func NewWSGSystemReservedPort() *WSGSystemReservedPort {
	m := new(WSGSystemReservedPort)
	return m
}

type WSGSystemSecondaryInterface struct {
	// IpAddress
	// IP address
	// Constraints:
	//    - required
	IpAddress *string `json:"ipAddress"`

	// SubnetMask
	// Subnet mask
	// Constraints:
	//    - required
	SubnetMask *string `json:"subnetMask"`

	// Vlan
	// vlan
	Vlan *string `json:"vlan,omitempty"`
}

func NewWSGSystemSecondaryInterface() *WSGSystemSecondaryInterface {
	m := new(WSGSystemSecondaryInterface)
	return m
}

type WSGSystemSecuritySetting struct {
	// AbsoluteSessionTimeout
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:1440
	AbsoluteSessionTimeout *int `json:"absoluteSessionTimeout,omitempty"`

	// CaptchaEnabled
	// Captcha setting
	CaptchaEnabled *bool `json:"captchaEnabled,omitempty"`

	// MaxInteractiveConcurrentSessions
	// Constraints:
	//    - nullable
	//    - min:3
	//    - max:10
	MaxInteractiveConcurrentSessions *int `json:"maxInteractiveConcurrentSessions,omitempty"`

	// MaxPublicApiConcurrentSessions
	// Constraints:
	//    - nullable
	//    - min:64
	//    - max:2048
	MaxPublicApiConcurrentSessions *int `json:"maxPublicApiConcurrentSessions,omitempty"`
}

func NewWSGSystemSecuritySetting() *WSGSystemSecuritySetting {
	m := new(WSGSystemSecuritySetting)
	return m
}

type WSGSystemSms struct {
	// AccountSid
	// Account SID
	AccountSid *string `json:"accountSid,omitempty"`

	// AuthToken
	// Auth Token
	AuthToken *string `json:"authToken,omitempty"`

	// DomainId
	// Domain Id
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Enabled
	// Enabled SMS server or not
	// Constraints:
	//    - oneof:[0,1]
	Enabled *int `json:"enabled,omitempty"`

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
	// Constraints:
	//    - oneof:[Twilio]
	ServerType *string `json:"serverType,omitempty"`
}

func NewWSGSystemSms() *WSGSystemSms {
	m := new(WSGSystemSms)
	return m
}

type WSGSystemSmsList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first SMS gateway returned out of the complete SMS gateway list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more SMS gateway after the list that is currently displayed
	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSystemSms `json:"list,omitempty"`

	// TotalCount
	// Total SMS gateway count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGSystemSmsList() *WSGSystemSmsList {
	m := new(WSGSystemSmsList)
	return m
}

type WSGSystemSnmpAgentConfiguration struct {
	// SnmpNotificationEnabled
	// Enable SNMP Notifications Globally (If SNMP Notification is disabled globally, no Notification message is sent out.)
	SnmpNotificationEnabled *bool `json:"snmpNotificationEnabled,omitempty"`

	// SnmpV2Agent
	// Community List of the SNMP V2 Agent.
	SnmpV2Agent []*WSGCommonSnmpCommunity `json:"snmpV2Agent,omitempty"`

	// SnmpV3Agent
	// User List of the SNMP V2 Agent.
	SnmpV3Agent []*WSGCommonSnmpUser `json:"snmpV3Agent,omitempty"`
}

func NewWSGSystemSnmpAgentConfiguration() *WSGSystemSnmpAgentConfiguration {
	m := new(WSGSystemSnmpAgentConfiguration)
	return m
}

type WSGSystemStaticRoute struct {
	// Gateway
	// Gateway
	// Constraints:
	//    - required
	Gateway *string `json:"gateway"`

	// NetworkAddress
	// Network address
	// Constraints:
	//    - required
	NetworkAddress *string `json:"networkAddress"`

	// SubnetMask
	// Subnet mask
	// Constraints:
	//    - required
	SubnetMask *string `json:"subnetMask"`
}

func NewWSGSystemStaticRoute() *WSGSystemStaticRoute {
	m := new(WSGSystemStaticRoute)
	return m
}

type WSGSystemStaticRouteList struct {
	// StaticRoutes
	// Static route for Control Plane
	StaticRoutes []*WSGSystemCpStaticRoute `json:"staticRoutes,omitempty"`
}

func NewWSGSystemStaticRouteList() *WSGSystemStaticRouteList {
	m := new(WSGSystemStaticRouteList)
	return m
}

type WSGSystemStatisticList []*WSGSystemStatisticListType

func MakeWSGSystemStatisticList() WSGSystemStatisticList {
	m := make(WSGSystemStatisticList, 0)
	return m
}

type WSGSystemStatisticListType struct {
	Cluster *WSGSystemPortStatistic `json:"cluster,omitempty"`

	Control *WSGSystemPortStatistic `json:"control,omitempty"`

	Cpu *WSGSystemStatisticListTypeCpuType `json:"cpu,omitempty"`

	Disk *WSGSystemStatisticListTypeDiskType `json:"disk,omitempty"`

	Management *WSGSystemPortStatistic `json:"management,omitempty"`

	Memory *WSGSystemStatisticListTypeMemoryType `json:"memory,omitempty"`

	Port0 *WSGSystemPortStatistic `json:"port0,omitempty"`

	Port1 *WSGSystemPortStatistic `json:"port1,omitempty"`

	Port2 *WSGSystemPortStatistic `json:"port2,omitempty"`

	Port3 *WSGSystemPortStatistic `json:"port3,omitempty"`

	Port4 *WSGSystemPortStatistic `json:"port4,omitempty"`

	Port5 *WSGSystemPortStatistic `json:"port5,omitempty"`

	// Timestamp
	// timestamp
	Timestamp *float64 `json:"timestamp,omitempty"`
}

func NewWSGSystemStatisticListType() *WSGSystemStatisticListType {
	m := new(WSGSystemStatisticListType)
	return m
}

type WSGSystemStatisticListTypeCpuType struct {
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

func NewWSGSystemStatisticListTypeCpuType() *WSGSystemStatisticListTypeCpuType {
	m := new(WSGSystemStatisticListTypeCpuType)
	return m
}

type WSGSystemStatisticListTypeDiskType struct {
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

func NewWSGSystemStatisticListTypeDiskType() *WSGSystemStatisticListTypeDiskType {
	m := new(WSGSystemStatisticListTypeDiskType)
	return m
}

type WSGSystemStatisticListTypeMemoryType struct {
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

func NewWSGSystemStatisticListTypeMemoryType() *WSGSystemStatisticListTypeMemoryType {
	m := new(WSGSystemStatisticListTypeMemoryType)
	return m
}

type WSGSystemSettings struct {
	// ApMacOUIEnabled
	// Enabled AP Mac OUI feature or no
	ApMacOUIEnabled *bool `json:"apMacOUIEnabled,omitempty"`

	// ApNumberLimitEnabled
	// Enabled AP number limit feature or not
	ApNumberLimitEnabled *bool `json:"apNumberLimitEnabled,omitempty"`

	ApNumberLimitSettingsOfDomain []*WSGSystemApNumberLimitSettingOfDomain `json:"apNumberLimitSettingsOfDomain,omitempty"`

	ApNumberLimitSettingsOfZone []*WSGSystemApNumberLimitSettingOfZone `json:"apNumberLimitSettingsOfZone,omitempty"`
}

func NewWSGSystemSettings() *WSGSystemSettings {
	m := new(WSGSystemSettings)
	return m
}

type WSGSystemTimeSetting struct {
	AuthenticationKey *WSGSystemAuthenticationKey `json:"authenticationKey,omitempty"`

	// CurrentSystemTimeString
	// System Time
	CurrentSystemTimeString *string `json:"currentSystemTimeString,omitempty"`

	// CurrentSystemTimeUTCString
	// System UTC Time
	CurrentSystemTimeUTCString *string `json:"currentSystemTimeUTCString,omitempty"`

	// NtpServer
	// NtpServer address
	NtpServer *string `json:"ntpServer,omitempty"`

	SecondaryAuthenticationKey *WSGSystemAuthenticationKey `json:"secondaryAuthenticationKey,omitempty"`

	// SecondaryNtpServer
	// Backup NtpServer address
	SecondaryNtpServer *string `json:"secondaryNtpServer,omitempty"`

	// Timezone
	// System defined time zone, please refer to the 'Overview > Time Zone' list
	Timezone *string `json:"timezone,omitempty"`
}

func NewWSGSystemTimeSetting() *WSGSystemTimeSetting {
	m := new(WSGSystemTimeSetting)
	return m
}

type WSGSystemUpdateApMacOUI struct {
	Description *WSGCommonDescription `json:"description,omitempty"`
}

func NewWSGSystemUpdateApMacOUI() *WSGSystemUpdateApMacOUI {
	m := new(WSGSystemUpdateApMacOUI)
	return m
}

type WSGSystemUpdateDpMeshTunnelSetting struct {
	// Encrypted
	// Data Plane mesh tunnel encrypted.
	Encrypted *bool `json:"encrypted,omitempty"`
}

func NewWSGSystemUpdateDpMeshTunnelSetting() *WSGSystemUpdateDpMeshTunnelSetting {
	m := new(WSGSystemUpdateDpMeshTunnelSetting)
	return m
}

type WSGSystemUserDefinedInterfaceList struct {
	// UserDefinedInterface
	// User defined interface for Control Plane
	UserDefinedInterface []*WSGSystemCpUserDefinedInterface `json:"userDefinedInterface,omitempty"`
}

func NewWSGSystemUserDefinedInterfaceList() *WSGSystemUserDefinedInterfaceList {
	m := new(WSGSystemUserDefinedInterfaceList)
	return m
}

// AddGlobalSettingsSystemTimeValidate
//
// Use this API command to validate a NTP server.
//
// Request Body:
//	 - body *WSGSystemNtpServerValidation
func (s *WSGSystemService) AddGlobalSettingsSystemTimeValidate(ctx context.Context, body *WSGSystemNtpServerValidation) (*WSGSystemNtpServerValidationMessage, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSystemNtpServerValidationMessage
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddGlobalSettingsSystemTimeValidate, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemNtpServerValidationMessage()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddSystemApBalance
//
// Execute ap balance.
func (s *WSGSystemService) AddSystemApBalance(ctx context.Context) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddSystemApBalance, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddSystemApMacOUIs
//
// Use this API command to create AP Mac OUI.
//
// Request Body:
//	 - body *WSGSystemCreateApMacOUI
func (s *WSGSystemService) AddSystemApMacOUIs(ctx context.Context, body *WSGSystemCreateApMacOUI) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddSystemApMacOUIs, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// AddSystemApRoutineConfigInterval
//
// Use this API command to get AP routine configuration interval setting.
//
// Request Body:
//	 - body *WSGAPRoutineConfigIntervalReq
func (s *WSGSystemService) AddSystemApRoutineConfigInterval(ctx context.Context, body *WSGAPRoutineConfigIntervalReq) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddSystemApRoutineConfigInterval, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// AddSystemApRoutineStatusIntervalSlowdown
//
// Use this API command to set AP routine status interval setting to 900 seconds.
func (s *WSGSystemService) AddSystemApRoutineStatusIntervalSlowdown(ctx context.Context) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddSystemApRoutineStatusIntervalSlowdown, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// AddSystemApRoutineStatusIntervalSpeedup
//
// Use this API command to set AP routine status interval setting to 180 seconds.
func (s *WSGSystemService) AddSystemApRoutineStatusIntervalSpeedup(ctx context.Context) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddSystemApRoutineStatusIntervalSpeedup, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteSystemApMacOUIsByOUI
//
// Use this API command to delete AP Mac OUI.
//
// Required Parameters:
// - OUI string
//		- required
func (s *WSGSystemService) DeleteSystemApMacOUIsByOUI(ctx context.Context, OUI string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteSystemApMacOUIsByOUI, true)
	req.SetPathParameter("OUI", OUI)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteSystemNbi
//
// Use this API command to disable the user information by Northbound Portal Interface.
//
// Optional Parameters:
// - domainId string
//		- nullable
func (s *WSGSystemService) DeleteSystemNbi(ctx context.Context, optionalParams map[string][]string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteSystemNbi, true)
	if v, ok := optionalParams["domainId"]; ok && len(v) > 0 {
		req.SetQueryParameter("domainId", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindController
//
// Use this API command to retrieve the system summary.
func (s *WSGSystemService) FindController(ctx context.Context) (*WSGSystemControllerList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSystemControllerList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindController, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemControllerList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindControllerStatisticsById
//
// Use this API command to retrieve the system statistics.
//
// Required Parameters:
// - id string
//		- required
//
// Optional Parameters:
// - interval string
//		- nullable
// - size float64
//		- nullable
func (s *WSGSystemService) FindControllerStatisticsById(ctx context.Context, id string, optionalParams map[string][]string) (WSGSystemStatisticList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     WSGSystemStatisticList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindControllerStatisticsById, true)
	req.SetPathParameter("id", id)
	if v, ok := optionalParams["interval"]; ok && len(v) > 0 {
		req.SetQueryParameter("interval", v)
	}
	if v, ok := optionalParams["size"]; ok && len(v) > 0 {
		req.SetQueryParameter("size", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = MakeWSGSystemStatisticList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindSystem
//
// Use this API command to get settings of system. Currently, Only can get settings about AP number limit.
func (s *WSGSystemService) FindSystem(ctx context.Context) (*WSGSystemSettings, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSystemSettings
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindSystem, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemSettings()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindSystemApMacOUIs
//
// Use this API command to retrieve a list of AP Mac OUIs.
func (s *WSGSystemService) FindSystemApMacOUIs(ctx context.Context) (*WSGSystemApMacOUIList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSystemApMacOUIList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindSystemApMacOUIs, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemApMacOUIList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindSystemApmodels
//
// Use this API command to retrieve support AP models for the current installed SZ version's default AP firmware.
func (s *WSGSystemService) FindSystemApmodels(ctx context.Context) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindSystemApmodels, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindSystemApmodelsByFirmwareVersion
//
// Use this API command to retrieve support AP models from input firmware version.
//
// Required Parameters:
// - firmwareVersion string
//		- required
func (s *WSGSystemService) FindSystemApmodelsByFirmwareVersion(ctx context.Context, firmwareVersion string) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindSystemApmodelsByFirmwareVersion, true)
	req.SetPathParameter("firmwareVersion", firmwareVersion)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindSystemApRoutineConfigInterval
//
// Use this API command to get AP routine configuration interval setting.
func (s *WSGSystemService) FindSystemApRoutineConfigInterval(ctx context.Context) (*WSGAPRoutineConfigIntervalRsp, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAPRoutineConfigIntervalRsp
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindSystemApRoutineConfigInterval, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAPRoutineConfigIntervalRsp()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindSystemApRoutineStatusInterval
//
// Use this API command to get AP routine status interval setting.
func (s *WSGSystemService) FindSystemApRoutineStatusInterval(ctx context.Context) (*WSGAPRoutineStatusIntervalRsp, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAPRoutineStatusIntervalRsp
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindSystemApRoutineStatusInterval, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAPRoutineStatusIntervalRsp()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindSystemByQueryCriteria
//
// Use this API command to query settings of system. Currently, Only can get settings about AP number limit by query criteria with domain and zone filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGSystemService) FindSystemByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGSystemSettings, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSystemSettings
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindSystemByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemSettings()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindSystemDevicesSummary
//
// Use this API command to retrieve devices summary.
func (s *WSGSystemService) FindSystemDevicesSummary(ctx context.Context) (*WSGDeviceCapacityDevicesSummary, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDeviceCapacityDevicesSummary
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindSystemDevicesSummary, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDeviceCapacityDevicesSummary()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindSystemGatewayAdvanced
//
// Use this API command to retrieve gateway advanced setting.
func (s *WSGSystemService) FindSystemGatewayAdvanced(ctx context.Context) (*WSGSystemGatewayAdvanced, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSystemGatewayAdvanced
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindSystemGatewayAdvanced, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemGatewayAdvanced()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindSystemInventory
//
// Use this API command to retrieve the system inventory with current logon user domain.
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGSystemService) FindSystemInventory(ctx context.Context, optionalParams map[string][]string) (*WSGSystemInventoryList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSystemInventoryList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindSystemInventory, true)
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.SetQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.SetQueryParameter("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemInventoryList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindSystemNbi
//
// Use this API command to retrieve user information by Northbound Portal Interface.
//
// Optional Parameters:
// - domainId string
//		- nullable
func (s *WSGSystemService) FindSystemNbi(ctx context.Context, optionalParams map[string][]string) (*WSGSystemNorthboundInterface, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSystemNorthboundInterface
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindSystemNbi, true)
	if v, ok := optionalParams["domainId"]; ok && len(v) > 0 {
		req.SetQueryParameter("domainId", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemNorthboundInterface()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindSystemSecuritySetting
//
// Use this API command to retrieve the security setting.
func (s *WSGSystemService) FindSystemSecuritySetting(ctx context.Context) (*WSGSystemSecuritySetting, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSystemSecuritySetting
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindSystemSecuritySetting, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemSecuritySetting()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindSystemSystemTime
//
// Retrieve System Time Setting.
func (s *WSGSystemService) FindSystemSystemTime(ctx context.Context) (*WSGSystemTimeSetting, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSystemTimeSetting
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindSystemSystemTime, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemTimeSetting()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateSystem
//
// Use this API command to modify settings of system. Currently, Only can modify settings about AP number limit by query criteria with domain and zone filters.
//
// Request Body:
//	 - body *WSGSystemSettings
func (s *WSGSystemService) PartialUpdateSystem(ctx context.Context, body *WSGSystemSettings) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateSystem, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PartialUpdateSystemGatewayAdvanced
//
// Use this API command to modify the gateway advanced setting.
//
// Request Body:
//	 - body *WSGSystemModifyGatewayAdvanced
func (s *WSGSystemService) PartialUpdateSystemGatewayAdvanced(ctx context.Context, body *WSGSystemModifyGatewayAdvanced) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateSystemGatewayAdvanced, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PartialUpdateSystemNbi
//
// Use this API command to modify the user information by Northbound Portal Interface.
//
// Request Body:
//	 - body *WSGSystemNorthboundInterface
//
// Optional Parameters:
// - domainId string
//		- nullable
func (s *WSGSystemService) PartialUpdateSystemNbi(ctx context.Context, body *WSGSystemNorthboundInterface, optionalParams map[string][]string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateSystemNbi, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	if v, ok := optionalParams["domainId"]; ok && len(v) > 0 {
		req.SetQueryParameter("domainId", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// PartialUpdateSystemSystemTime
//
// Modify System Time Setting.
//
// Request Body:
//	 - body *WSGSystemModifySystemTimeSetting
func (s *WSGSystemService) PartialUpdateSystemSystemTime(ctx context.Context, body *WSGSystemModifySystemTimeSetting) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateSystemSystemTime, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// UpdateSystemApMacOUIsByOUI
//
// Use this API command to update AP Mac OUI.
//
// Request Body:
//	 - body *WSGSystemUpdateApMacOUI
//
// Required Parameters:
// - OUI string
//		- required
func (s *WSGSystemService) UpdateSystemApMacOUIsByOUI(ctx context.Context, body *WSGSystemUpdateApMacOUI, OUI string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateSystemApMacOUIsByOUI, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("OUI", OUI)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// UpdateSystemSecuritySetting
//
// Use this API command to retrieve the security setting.
//
// Request Body:
//	 - body *WSGSystemSecuritySetting
func (s *WSGSystemService) UpdateSystemSecuritySetting(ctx context.Context, body *WSGSystemSecuritySetting) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateSystemSecuritySetting, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
