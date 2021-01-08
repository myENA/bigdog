package bigdog

// API Version: v9_1

import (
	"context"
	"encoding/json"
	"io"
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

// WSGSystemApMacOUI
//
// Definition: system_apMacOUI
type WSGSystemApMacOUI struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	Oui *WSGCommonOui `json:"oui,omitempty"`
}

func NewWSGSystemApMacOUI() *WSGSystemApMacOUI {
	m := new(WSGSystemApMacOUI)
	return m
}

// WSGSystemApMacOUIList
//
// Definition: system_apMacOUIList
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

type WSGSystemApMacOUIListAPIResponse struct {
	*RawAPIResponse
	Data *WSGSystemApMacOUIList
}

func newWSGSystemApMacOUIListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSystemApMacOUIListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSystemApMacOUIListAPIResponse) Hydrate() error {
	r.Data = new(WSGSystemApMacOUIList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSystemApMacOUIList() *WSGSystemApMacOUIList {
	m := new(WSGSystemApMacOUIList)
	return m
}

// WSGSystemApNumberLimitSettingOfDomain
//
// Definition: system_apNumberLimitSettingOfDomain
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

// WSGSystemApNumberLimitSettingOfZone
//
// Definition: system_apNumberLimitSettingOfZone
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

// WSGSystemAuthenticationKey
//
// Definition: system_authenticationKey
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

// WSGSystemControllerList
//
// Definition: system_controllerList
type WSGSystemControllerList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSystemControllerListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGSystemControllerListAPIResponse struct {
	*RawAPIResponse
	Data *WSGSystemControllerList
}

func newWSGSystemControllerListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSystemControllerListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSystemControllerListAPIResponse) Hydrate() error {
	r.Data = new(WSGSystemControllerList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSystemControllerList() *WSGSystemControllerList {
	m := new(WSGSystemControllerList)
	return m
}

// WSGSystemControllerListType
//
// Definition: system_controllerListType
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

// WSGSystemControlPlaneConfiguration
//
// Definition: system_controlPlaneConfiguration
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

type WSGSystemControlPlaneConfigurationAPIResponse struct {
	*RawAPIResponse
	Data *WSGSystemControlPlaneConfiguration
}

func newWSGSystemControlPlaneConfigurationAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSystemControlPlaneConfigurationAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSystemControlPlaneConfigurationAPIResponse) Hydrate() error {
	r.Data = new(WSGSystemControlPlaneConfiguration)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSystemControlPlaneConfiguration() *WSGSystemControlPlaneConfiguration {
	m := new(WSGSystemControlPlaneConfiguration)
	return m
}

// WSGSystemControlPlaneInterface
//
// Definition: system_controlPlaneInterface
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

// WSGSystemControlPlaneInterfaceList
//
// Definition: system_controlPlaneInterfaceList
type WSGSystemControlPlaneInterfaceList struct {
	// ControlPlaneInterfaces
	// Interface list
	ControlPlaneInterfaces []*WSGSystemControlPlaneInterface `json:"controlPlaneInterfaces,omitempty"`
}

type WSGSystemControlPlaneInterfaceListAPIResponse struct {
	*RawAPIResponse
	Data *WSGSystemControlPlaneInterfaceList
}

func newWSGSystemControlPlaneInterfaceListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSystemControlPlaneInterfaceListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSystemControlPlaneInterfaceListAPIResponse) Hydrate() error {
	r.Data = new(WSGSystemControlPlaneInterfaceList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSystemControlPlaneInterfaceList() *WSGSystemControlPlaneInterfaceList {
	m := new(WSGSystemControlPlaneInterfaceList)
	return m
}

// WSGSystemControlPlaneList
//
// Definition: system_controlPlaneList
type WSGSystemControlPlaneList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSystemControlPlaneListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGSystemControlPlaneListAPIResponse struct {
	*RawAPIResponse
	Data *WSGSystemControlPlaneList
}

func newWSGSystemControlPlaneListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSystemControlPlaneListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSystemControlPlaneListAPIResponse) Hydrate() error {
	r.Data = new(WSGSystemControlPlaneList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSystemControlPlaneList() *WSGSystemControlPlaneList {
	m := new(WSGSystemControlPlaneList)
	return m
}

// WSGSystemControlPlaneListType
//
// Definition: system_controlPlaneListType
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

// WSGSystemCpStaticRoute
//
// Definition: system_cpStaticRoute
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

// WSGSystemCpUserDefinedInterface
//
// Definition: system_cpUserDefinedInterface
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

// WSGSystemCreateApMacOUI
//
// Definition: system_createApMacOUI
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

// WSGSystemDataPlaneConfiguration
//
// Definition: system_dataPlaneConfiguration
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

type WSGSystemDataPlaneConfigurationAPIResponse struct {
	*RawAPIResponse
	Data *WSGSystemDataPlaneConfiguration
}

func newWSGSystemDataPlaneConfigurationAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSystemDataPlaneConfigurationAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSystemDataPlaneConfigurationAPIResponse) Hydrate() error {
	r.Data = new(WSGSystemDataPlaneConfiguration)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSystemDataPlaneConfiguration() *WSGSystemDataPlaneConfiguration {
	m := new(WSGSystemDataPlaneConfiguration)
	return m
}

// WSGSystemDataPlaneList
//
// Definition: system_dataPlaneList
type WSGSystemDataPlaneList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSystemDataPlaneListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGSystemDataPlaneListAPIResponse struct {
	*RawAPIResponse
	Data *WSGSystemDataPlaneList
}

func newWSGSystemDataPlaneListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSystemDataPlaneListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSystemDataPlaneListAPIResponse) Hydrate() error {
	r.Data = new(WSGSystemDataPlaneList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSystemDataPlaneList() *WSGSystemDataPlaneList {
	m := new(WSGSystemDataPlaneList)
	return m
}

// WSGSystemDataPlaneListType
//
// Definition: system_dataPlaneListType
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

// WSGSystemDeleteBulkFtp
//
// Definition: system_deleteBulkFtp
type WSGSystemDeleteBulkFtp struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGSystemDeleteBulkFtp() *WSGSystemDeleteBulkFtp {
	m := new(WSGSystemDeleteBulkFtp)
	return m
}

// WSGSystemFriendlyNameLang
//
// Definition: system_friendlyNameLang
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

// WSGSystemFriendlyNameLangList
//
// Definition: system_friendlyNameLangList
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

type WSGSystemFriendlyNameLangListAPIResponse struct {
	*RawAPIResponse
	Data *WSGSystemFriendlyNameLangList
}

func newWSGSystemFriendlyNameLangListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSystemFriendlyNameLangListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSystemFriendlyNameLangListAPIResponse) Hydrate() error {
	r.Data = new(WSGSystemFriendlyNameLangList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSystemFriendlyNameLangList() *WSGSystemFriendlyNameLangList {
	m := new(WSGSystemFriendlyNameLangList)
	return m
}

// WSGSystemFtp
//
// Definition: system_ftp
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

type WSGSystemFtpAPIResponse struct {
	*RawAPIResponse
	Data *WSGSystemFtp
}

func newWSGSystemFtpAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSystemFtpAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSystemFtpAPIResponse) Hydrate() error {
	r.Data = new(WSGSystemFtp)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSystemFtp() *WSGSystemFtp {
	m := new(WSGSystemFtp)
	return m
}

// WSGSystemFtpGlobalSetting
//
// Definition: system_ftpGlobalSetting
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

type WSGSystemFtpGlobalSettingAPIResponse struct {
	*RawAPIResponse
	Data *WSGSystemFtpGlobalSetting
}

func newWSGSystemFtpGlobalSettingAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSystemFtpGlobalSettingAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSystemFtpGlobalSettingAPIResponse) Hydrate() error {
	r.Data = new(WSGSystemFtpGlobalSetting)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSystemFtpGlobalSetting() *WSGSystemFtpGlobalSetting {
	m := new(WSGSystemFtpGlobalSetting)
	return m
}

// WSGSystemFtpList
//
// Definition: system_ftpList
type WSGSystemFtpList struct {
	Extra interface{} `json:"extra,omitempty"`

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

type WSGSystemFtpListAPIResponse struct {
	*RawAPIResponse
	Data *WSGSystemFtpList
}

func newWSGSystemFtpListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSystemFtpListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSystemFtpListAPIResponse) Hydrate() error {
	r.Data = new(WSGSystemFtpList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSystemFtpList() *WSGSystemFtpList {
	m := new(WSGSystemFtpList)
	return m
}

// WSGSystemFtpTestResponse
//
// Definition: system_ftpTestResponse
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

type WSGSystemFtpTestResponseAPIResponse struct {
	*RawAPIResponse
	Data *WSGSystemFtpTestResponse
}

func newWSGSystemFtpTestResponseAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSystemFtpTestResponseAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSystemFtpTestResponseAPIResponse) Hydrate() error {
	r.Data = new(WSGSystemFtpTestResponse)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSystemFtpTestResponse() *WSGSystemFtpTestResponse {
	m := new(WSGSystemFtpTestResponse)
	return m
}

// WSGSystemGatewayAdvanced
//
// Definition: system_gatewayAdvanced
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

type WSGSystemGatewayAdvancedAPIResponse struct {
	*RawAPIResponse
	Data *WSGSystemGatewayAdvanced
}

func newWSGSystemGatewayAdvancedAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSystemGatewayAdvancedAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSystemGatewayAdvancedAPIResponse) Hydrate() error {
	r.Data = new(WSGSystemGatewayAdvanced)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSystemGatewayAdvanced() *WSGSystemGatewayAdvanced {
	m := new(WSGSystemGatewayAdvanced)
	return m
}

// WSGSystemGetDataPlaneMeshTunnelSetting
//
// Definition: system_getDataPlaneMeshTunnelSetting
type WSGSystemGetDataPlaneMeshTunnelSetting struct {
	// Encrypted
	// Data Plane mesh tunnel encrypted.
	Encrypted *bool `json:"encrypted,omitempty"`
}

type WSGSystemGetDataPlaneMeshTunnelSettingAPIResponse struct {
	*RawAPIResponse
	Data *WSGSystemGetDataPlaneMeshTunnelSetting
}

func newWSGSystemGetDataPlaneMeshTunnelSettingAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSystemGetDataPlaneMeshTunnelSettingAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSystemGetDataPlaneMeshTunnelSettingAPIResponse) Hydrate() error {
	r.Data = new(WSGSystemGetDataPlaneMeshTunnelSetting)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSystemGetDataPlaneMeshTunnelSetting() *WSGSystemGetDataPlaneMeshTunnelSetting {
	m := new(WSGSystemGetDataPlaneMeshTunnelSetting)
	return m
}

// WSGSystemInventoryList
//
// Definition: system_inventoryList
type WSGSystemInventoryList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSystemInventoryListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGSystemInventoryListAPIResponse struct {
	*RawAPIResponse
	Data *WSGSystemInventoryList
}

func newWSGSystemInventoryListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSystemInventoryListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSystemInventoryListAPIResponse) Hydrate() error {
	r.Data = new(WSGSystemInventoryList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSystemInventoryList() *WSGSystemInventoryList {
	m := new(WSGSystemInventoryList)
	return m
}

// WSGSystemInventoryListType
//
// Definition: system_inventoryListType
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

// WSGSystemIpv4AccessAndCoreSeparation
//
// Definition: system_ipv4AccessAndCoreSeparation
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

// WSGSystemIpv4ClusterInterface
//
// Definition: system_ipv4ClusterInterface
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

// WSGSystemIpv4ControlInterface
//
// Definition: system_ipv4ControlInterface
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

// WSGSystemIpv4ManagementInterface
//
// Definition: system_ipv4ManagementInterface
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

// WSGSystemIpv6AccessAndCoreSeparation
//
// Definition: system_ipv6AccessAndCoreSeparation
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

// WSGSystemIpv6ControlInterface
//
// Definition: system_ipv6ControlInterface
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

// WSGSystemIpv6ManagementInterface
//
// Definition: system_ipv6ManagementInterface
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

// WSGSystemIpv6PrimaryInterface
//
// Definition: system_ipv6PrimaryInterface
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

// WSGSystemLwapp2scgConfiguration
//
// Definition: system_lwapp2scgConfiguration
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

type WSGSystemLwapp2scgConfigurationAPIResponse struct {
	*RawAPIResponse
	Data *WSGSystemLwapp2scgConfiguration
}

func newWSGSystemLwapp2scgConfigurationAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSystemLwapp2scgConfigurationAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSystemLwapp2scgConfigurationAPIResponse) Hydrate() error {
	r.Data = new(WSGSystemLwapp2scgConfiguration)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSystemLwapp2scgConfiguration() *WSGSystemLwapp2scgConfiguration {
	m := new(WSGSystemLwapp2scgConfiguration)
	return m
}

// WSGSystemModifyControlPlane
//
// Definition: system_modifyControlPlane
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

// WSGSystemModifyCPStaticRoute
//
// Definition: system_modifyCPStaticRoute
type WSGSystemModifyCPStaticRoute struct {
	// StaticRoutes
	// Static route for control plane.
	StaticRoutes []*WSGSystemCpStaticRoute `json:"staticRoutes,omitempty"`
}

func NewWSGSystemModifyCPStaticRoute() *WSGSystemModifyCPStaticRoute {
	m := new(WSGSystemModifyCPStaticRoute)
	return m
}

// WSGSystemModifyCPUserDefinedInterface
//
// Definition: system_modifyCPUserDefinedInterface
type WSGSystemModifyCPUserDefinedInterface struct {
	// UserDefinedInterface
	// User defined interface for Control Plane
	UserDefinedInterface []*WSGSystemCpUserDefinedInterface `json:"userDefinedInterface,omitempty"`
}

func NewWSGSystemModifyCPUserDefinedInterface() *WSGSystemModifyCPUserDefinedInterface {
	m := new(WSGSystemModifyCPUserDefinedInterface)
	return m
}

// WSGSystemModifyDataPlane
//
// Definition: system_modifyDataPlane
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

// WSGSystemModifyDataPlaneState
//
// Definition: system_modifyDataPlaneState
type WSGSystemModifyDataPlaneState struct {
	// IsDataCenter
	// Mark this Data Plane as a CALEA Relay
	IsDataCenter *bool `json:"isDataCenter,omitempty"`
}

func NewWSGSystemModifyDataPlaneState() *WSGSystemModifyDataPlaneState {
	m := new(WSGSystemModifyDataPlaneState)
	return m
}

// WSGSystemModifyGatewayAdvanced
//
// Definition: system_modifyGatewayAdvanced
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

// WSGSystemModifyIpSupportType
//
// Definition: system_modifyIpSupportType
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

// WSGSystemModifyLwapp2scg
//
// Definition: system_modifyLwapp2scg
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

// WSGSystemModifySnmpAgent
//
// Definition: system_modifySnmpAgent
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

// WSGSystemModifySystemTimeSetting
//
// Definition: system_modifySystemTimeSetting
type WSGSystemModifySystemTimeSetting struct {
	AuthenticationKey *WSGSystemAuthenticationKey `json:"authenticationKey,omitempty"`

	// NtpServer
	// Primary NtpServer address
	NtpServer *string `json:"ntpServer,omitempty"`

	SecondaryAuthenticationKey *WSGSystemAuthenticationKey `json:"secondaryAuthenticationKey,omitempty"`

	// SecondaryNtpServer
	// Secondary NtpServer address
	SecondaryNtpServer *string `json:"secondaryNtpServer,omitempty"`

	ThirdAuthenticationKey *WSGSystemAuthenticationKey `json:"thirdAuthenticationKey,omitempty"`

	// ThirdNtpServer
	// Third NtpServer address
	ThirdNtpServer *string `json:"thirdNtpServer,omitempty"`

	// Timezone
	// System defined time zone, please refer to the 'Overview > Time Zone' list
	Timezone *string `json:"timezone,omitempty"`
}

func NewWSGSystemModifySystemTimeSetting() *WSGSystemModifySystemTimeSetting {
	m := new(WSGSystemModifySystemTimeSetting)
	return m
}

// WSGSystemNorthboundInterface
//
// Definition: system_northboundInterface
type WSGSystemNorthboundInterface struct {
	Password *WSGCommonApLoginPassword `json:"password,omitempty"`

	// RadiusAuthType
	// AuthType of the Radius used in Northbound Interface, the value should be "PAP" or "CHAP".
	// Constraints:
	//    - oneof:[PAP,CHAP]
	RadiusAuthType *string `json:"radiusAuthType,omitempty"`

	UserName *WSGCommonApLoginName `json:"userName,omitempty"`
}

type WSGSystemNorthboundInterfaceAPIResponse struct {
	*RawAPIResponse
	Data *WSGSystemNorthboundInterface
}

func newWSGSystemNorthboundInterfaceAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSystemNorthboundInterfaceAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSystemNorthboundInterfaceAPIResponse) Hydrate() error {
	r.Data = new(WSGSystemNorthboundInterface)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSystemNorthboundInterface() *WSGSystemNorthboundInterface {
	m := new(WSGSystemNorthboundInterface)
	return m
}

// WSGSystemNtpServerValidation
//
// Definition: system_ntpServerValidation
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

// WSGSystemNtpServerValidationMessage
//
// Definition: system_ntpServerValidationMessage
type WSGSystemNtpServerValidationMessage struct {
	// Message
	// NTP Server Validation Message
	Message *string `json:"message,omitempty"`
}

type WSGSystemNtpServerValidationMessageAPIResponse struct {
	*RawAPIResponse
	Data *WSGSystemNtpServerValidationMessage
}

func newWSGSystemNtpServerValidationMessageAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSystemNtpServerValidationMessageAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSystemNtpServerValidationMessageAPIResponse) Hydrate() error {
	r.Data = new(WSGSystemNtpServerValidationMessage)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSystemNtpServerValidationMessage() *WSGSystemNtpServerValidationMessage {
	m := new(WSGSystemNtpServerValidationMessage)
	return m
}

// WSGSystemPortalLang
//
// Definition: system_portalLang
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

// WSGSystemPortalLangList
//
// Definition: system_portalLangList
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

type WSGSystemPortalLangListAPIResponse struct {
	*RawAPIResponse
	Data *WSGSystemPortalLangList
}

func newWSGSystemPortalLangListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSystemPortalLangListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSystemPortalLangListAPIResponse) Hydrate() error {
	r.Data = new(WSGSystemPortalLangList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSystemPortalLangList() *WSGSystemPortalLangList {
	m := new(WSGSystemPortalLangList)
	return m
}

// WSGSystemPortStatistic
//
// Definition: system_portStatistic
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

// WSGSystemPrimaryInterface
//
// Definition: system_primaryInterface
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

// WSGSystemReservedPort
//
// Definition: system_reservedPort
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

// WSGSystemSecondaryInterface
//
// Definition: system_secondaryInterface
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

// WSGSystemSecuritySetting
//
// Definition: system_securitySetting
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

type WSGSystemSecuritySettingAPIResponse struct {
	*RawAPIResponse
	Data *WSGSystemSecuritySetting
}

func newWSGSystemSecuritySettingAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSystemSecuritySettingAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSystemSecuritySettingAPIResponse) Hydrate() error {
	r.Data = new(WSGSystemSecuritySetting)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSystemSecuritySetting() *WSGSystemSecuritySetting {
	m := new(WSGSystemSecuritySetting)
	return m
}

// WSGSystemSms
//
// Definition: system_sms
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

type WSGSystemSmsAPIResponse struct {
	*RawAPIResponse
	Data *WSGSystemSms
}

func newWSGSystemSmsAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSystemSmsAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSystemSmsAPIResponse) Hydrate() error {
	r.Data = new(WSGSystemSms)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSystemSms() *WSGSystemSms {
	m := new(WSGSystemSms)
	return m
}

// WSGSystemSmsList
//
// Definition: system_smsList
type WSGSystemSmsList struct {
	Extra interface{} `json:"extra,omitempty"`

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

type WSGSystemSmsListAPIResponse struct {
	*RawAPIResponse
	Data *WSGSystemSmsList
}

func newWSGSystemSmsListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSystemSmsListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSystemSmsListAPIResponse) Hydrate() error {
	r.Data = new(WSGSystemSmsList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSystemSmsList() *WSGSystemSmsList {
	m := new(WSGSystemSmsList)
	return m
}

// WSGSystemSnmpAgentConfiguration
//
// Definition: system_snmpAgentConfiguration
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

type WSGSystemSnmpAgentConfigurationAPIResponse struct {
	*RawAPIResponse
	Data *WSGSystemSnmpAgentConfiguration
}

func newWSGSystemSnmpAgentConfigurationAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSystemSnmpAgentConfigurationAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSystemSnmpAgentConfigurationAPIResponse) Hydrate() error {
	r.Data = new(WSGSystemSnmpAgentConfiguration)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSystemSnmpAgentConfiguration() *WSGSystemSnmpAgentConfiguration {
	m := new(WSGSystemSnmpAgentConfiguration)
	return m
}

// WSGSystemStaticRoute
//
// Definition: system_staticRoute
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

// WSGSystemStaticRouteList
//
// Definition: system_staticRouteList
type WSGSystemStaticRouteList struct {
	// StaticRoutes
	// Static route for Control Plane
	StaticRoutes []*WSGSystemCpStaticRoute `json:"staticRoutes,omitempty"`
}

type WSGSystemStaticRouteListAPIResponse struct {
	*RawAPIResponse
	Data *WSGSystemStaticRouteList
}

func newWSGSystemStaticRouteListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSystemStaticRouteListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSystemStaticRouteListAPIResponse) Hydrate() error {
	r.Data = new(WSGSystemStaticRouteList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSystemStaticRouteList() *WSGSystemStaticRouteList {
	m := new(WSGSystemStaticRouteList)
	return m
}

// WSGSystemStatisticList
//
// Definition: system_statisticList
type WSGSystemStatisticList []*WSGSystemStatisticListType

func MakeWSGSystemStatisticList() WSGSystemStatisticList {
	m := make(WSGSystemStatisticList, 0)
	return m
}

// WSGSystemStatisticListType
//
// Definition: system_statisticListType
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

// WSGSystemStatisticListTypeCpuType
//
// Definition: system_statisticListTypeCpuType
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

// WSGSystemStatisticListTypeDiskType
//
// Definition: system_statisticListTypeDiskType
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

// WSGSystemStatisticListTypeMemoryType
//
// Definition: system_statisticListTypeMemoryType
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

// WSGSystemSettings
//
// Definition: system_systemSettings
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

type WSGSystemSettingsAPIResponse struct {
	*RawAPIResponse
	Data *WSGSystemSettings
}

func newWSGSystemSettingsAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSystemSettingsAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSystemSettingsAPIResponse) Hydrate() error {
	r.Data = new(WSGSystemSettings)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSystemSettings() *WSGSystemSettings {
	m := new(WSGSystemSettings)
	return m
}

// WSGSystemTimeSetting
//
// Definition: system_systemTimeSetting
type WSGSystemTimeSetting struct {
	AuthenticationKey *WSGSystemAuthenticationKey `json:"authenticationKey,omitempty"`

	// CurrentSystemTimeString
	// System Time
	CurrentSystemTimeString *string `json:"currentSystemTimeString,omitempty"`

	// CurrentSystemTimeUTCString
	// System UTC Time
	CurrentSystemTimeUTCString *string `json:"currentSystemTimeUTCString,omitempty"`

	// NtpServer
	// Primary NtpServer address
	NtpServer *string `json:"ntpServer,omitempty"`

	SecondaryAuthenticationKey *WSGSystemAuthenticationKey `json:"secondaryAuthenticationKey,omitempty"`

	// SecondaryNtpServer
	// Secondary NtpServer address
	SecondaryNtpServer *string `json:"secondaryNtpServer,omitempty"`

	ThirdAuthenticationKey *WSGSystemAuthenticationKey `json:"thirdAuthenticationKey,omitempty"`

	// ThirdNtpServer
	// Third NtpServer address
	ThirdNtpServer *string `json:"thirdNtpServer,omitempty"`

	// Timezone
	// System defined time zone, please refer to the 'Overview > Time Zone' list
	Timezone *string `json:"timezone,omitempty"`
}

type WSGSystemTimeSettingAPIResponse struct {
	*RawAPIResponse
	Data *WSGSystemTimeSetting
}

func newWSGSystemTimeSettingAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSystemTimeSettingAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSystemTimeSettingAPIResponse) Hydrate() error {
	r.Data = new(WSGSystemTimeSetting)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSystemTimeSetting() *WSGSystemTimeSetting {
	m := new(WSGSystemTimeSetting)
	return m
}

// WSGSystemUpdateApMacOUI
//
// Definition: system_updateApMacOUI
type WSGSystemUpdateApMacOUI struct {
	Description *WSGCommonDescription `json:"description,omitempty"`
}

func NewWSGSystemUpdateApMacOUI() *WSGSystemUpdateApMacOUI {
	m := new(WSGSystemUpdateApMacOUI)
	return m
}

// WSGSystemUpdateDpMeshTunnelSetting
//
// Definition: system_updateDpMeshTunnelSetting
type WSGSystemUpdateDpMeshTunnelSetting struct {
	// Encrypted
	// Data Plane mesh tunnel encrypted.
	Encrypted *bool `json:"encrypted,omitempty"`
}

func NewWSGSystemUpdateDpMeshTunnelSetting() *WSGSystemUpdateDpMeshTunnelSetting {
	m := new(WSGSystemUpdateDpMeshTunnelSetting)
	return m
}

// WSGSystemUserDefinedInterfaceList
//
// Definition: system_userDefinedInterfaceList
type WSGSystemUserDefinedInterfaceList struct {
	// UserDefinedInterface
	// User defined interface for Control Plane
	UserDefinedInterface []*WSGSystemCpUserDefinedInterface `json:"userDefinedInterface,omitempty"`
}

type WSGSystemUserDefinedInterfaceListAPIResponse struct {
	*RawAPIResponse
	Data *WSGSystemUserDefinedInterfaceList
}

func newWSGSystemUserDefinedInterfaceListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSystemUserDefinedInterfaceListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSystemUserDefinedInterfaceListAPIResponse) Hydrate() error {
	r.Data = new(WSGSystemUserDefinedInterfaceList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSystemUserDefinedInterfaceList() *WSGSystemUserDefinedInterfaceList {
	m := new(WSGSystemUserDefinedInterfaceList)
	return m
}

// AddGlobalSettingsSystemTimeValidate
//
// Operation ID: addGlobalSettingsSystemTimeValidate
//
// Use this API command to validate a NTP server.
//
// Request Body:
//	 - body *WSGSystemNtpServerValidation
func (s *WSGSystemService) AddGlobalSettingsSystemTimeValidate(ctx context.Context, body *WSGSystemNtpServerValidation, mutators ...RequestMutator) (*WSGSystemNtpServerValidationMessageAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSystemNtpServerValidationMessageAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSystemNtpServerValidationMessageAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddGlobalSettingsSystemTimeValidate, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGSystemNtpServerValidationMessageAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemNtpServerValidationMessageAPIResponse), err
}

// AddSystemApBalance
//
// Operation ID: addSystemAp_balance
//
// Execute ap balance.
func (s *WSGSystemService) AddSystemApBalance(ctx context.Context, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddSystemApBalance, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// AddSystemApMacOUIs
//
// Operation ID: addSystemApMacOUIs
//
// Use this API command to create AP Mac OUI.
//
// Request Body:
//	 - body *WSGSystemCreateApMacOUI
func (s *WSGSystemService) AddSystemApMacOUIs(ctx context.Context, body *WSGSystemCreateApMacOUI, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddSystemApMacOUIs, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// AddSystemApRoutineConfigInterval
//
// Operation ID: addSystemApRoutineConfigInterval
//
// Use this API command to get AP routine configuration interval setting.
//
// Request Body:
//	 - body *WSGAPRoutineConfigIntervalReq
func (s *WSGSystemService) AddSystemApRoutineConfigInterval(ctx context.Context, body *WSGAPRoutineConfigIntervalReq, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddSystemApRoutineConfigInterval, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// AddSystemApRoutineStatusIntervalSlowdown
//
// Operation ID: addSystemApRoutineStatusIntervalSlowdown
//
// Use this API command to set AP routine status interval setting to 900 seconds.
func (s *WSGSystemService) AddSystemApRoutineStatusIntervalSlowdown(ctx context.Context, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddSystemApRoutineStatusIntervalSlowdown, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// AddSystemApRoutineStatusIntervalSpeedup
//
// Operation ID: addSystemApRoutineStatusIntervalSpeedup
//
// Use this API command to set AP routine status interval setting to 180 seconds.
func (s *WSGSystemService) AddSystemApRoutineStatusIntervalSpeedup(ctx context.Context, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddSystemApRoutineStatusIntervalSpeedup, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteSystemApMacOUIsByOUI
//
// Operation ID: deleteSystemApMacOUIsByOUI
//
// Use this API command to delete AP Mac OUI.
//
// Required Parameters:
// - OUI string
//		- required
func (s *WSGSystemService) DeleteSystemApMacOUIsByOUI(ctx context.Context, OUI string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteSystemApMacOUIsByOUI, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("OUI", OUI)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteSystemNbi
//
// Operation ID: deleteSystemNbi
//
// Use this API command to disable the user information by Northbound Portal Interface.
//
// Optional Parameters:
// - domainId string
//		- nullable
func (s *WSGSystemService) DeleteSystemNbi(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteSystemNbi, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["domainId"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("domainId", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindController
//
// Operation ID: findController
//
// Use this API command to retrieve the system summary.
func (s *WSGSystemService) FindController(ctx context.Context, mutators ...RequestMutator) (*WSGSystemControllerListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSystemControllerListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSystemControllerListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindController, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemControllerListAPIResponse), err
}

// FindControllerStatisticsById
//
// Operation ID: findControllerStatisticsById
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
func (s *WSGSystemService) FindControllerStatisticsById(ctx context.Context, id string, optionalParams map[string][]string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindControllerStatisticsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	if v, ok := optionalParams["interval"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("interval", v)
	}
	if v, ok := optionalParams["size"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("size", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindSystem
//
// Operation ID: findSystem
//
// Use this API command to get settings of system. Currently, Only can get settings about AP number limit.
func (s *WSGSystemService) FindSystem(ctx context.Context, mutators ...RequestMutator) (*WSGSystemSettingsAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSystemSettingsAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSystemSettingsAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindSystem, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemSettingsAPIResponse), err
}

// FindSystemApMacOUIs
//
// Operation ID: findSystemApMacOUIs
//
// Use this API command to retrieve a list of AP Mac OUIs.
func (s *WSGSystemService) FindSystemApMacOUIs(ctx context.Context, mutators ...RequestMutator) (*WSGSystemApMacOUIListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSystemApMacOUIListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSystemApMacOUIListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindSystemApMacOUIs, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemApMacOUIListAPIResponse), err
}

// FindSystemApmodels
//
// Operation ID: findSystemApmodels
//
// Use this API command to retrieve support AP models for the current installed SZ version's default AP firmware.
func (s *WSGSystemService) FindSystemApmodels(ctx context.Context, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindSystemApmodels, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindSystemApmodelsByFirmwareVersion
//
// Operation ID: findSystemApmodelsByFirmwareVersion
//
// Use this API command to retrieve support AP models from input firmware version.
//
// Required Parameters:
// - firmwareVersion string
//		- required
func (s *WSGSystemService) FindSystemApmodelsByFirmwareVersion(ctx context.Context, firmwareVersion string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindSystemApmodelsByFirmwareVersion, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("firmwareVersion", firmwareVersion)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindSystemApRoutineConfigInterval
//
// Operation ID: findSystemApRoutineConfigInterval
//
// Use this API command to get AP routine configuration interval setting.
func (s *WSGSystemService) FindSystemApRoutineConfigInterval(ctx context.Context, mutators ...RequestMutator) (*WSGAPRoutineConfigIntervalRspAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAPRoutineConfigIntervalRspAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAPRoutineConfigIntervalRspAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindSystemApRoutineConfigInterval, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAPRoutineConfigIntervalRspAPIResponse), err
}

// FindSystemApRoutineStatusInterval
//
// Operation ID: findSystemApRoutineStatusInterval
//
// Use this API command to get AP routine status interval setting.
func (s *WSGSystemService) FindSystemApRoutineStatusInterval(ctx context.Context, mutators ...RequestMutator) (*WSGAPRoutineStatusIntervalRspAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAPRoutineStatusIntervalRspAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAPRoutineStatusIntervalRspAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindSystemApRoutineStatusInterval, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAPRoutineStatusIntervalRspAPIResponse), err
}

// FindSystemByQueryCriteria
//
// Operation ID: findSystemByQueryCriteria
//
// Use this API command to query settings of system. Currently, Only can get settings about AP number limit by query criteria with domain and zone filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGSystemService) FindSystemByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGSystemSettingsAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSystemSettingsAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSystemSettingsAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindSystemByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGSystemSettingsAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemSettingsAPIResponse), err
}

// FindSystemDevicesSummary
//
// Operation ID: findSystemDevicesSummary
//
// Use this API command to retrieve devices summary.
func (s *WSGSystemService) FindSystemDevicesSummary(ctx context.Context, mutators ...RequestMutator) (*WSGDeviceCapacityDevicesSummaryAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGDeviceCapacityDevicesSummaryAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGDeviceCapacityDevicesSummaryAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindSystemDevicesSummary, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGDeviceCapacityDevicesSummaryAPIResponse), err
}

// FindSystemGatewayAdvanced
//
// Operation ID: findSystemGatewayAdvanced
//
// Use this API command to retrieve gateway advanced setting.
func (s *WSGSystemService) FindSystemGatewayAdvanced(ctx context.Context, mutators ...RequestMutator) (*WSGSystemGatewayAdvancedAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSystemGatewayAdvancedAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSystemGatewayAdvancedAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindSystemGatewayAdvanced, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemGatewayAdvancedAPIResponse), err
}

// FindSystemInventory
//
// Operation ID: findSystemInventory
//
// Use this API command to retrieve the system inventory with current logon user domain.
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGSystemService) FindSystemInventory(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGSystemInventoryListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSystemInventoryListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSystemInventoryListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindSystemInventory, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemInventoryListAPIResponse), err
}

// FindSystemNbi
//
// Operation ID: findSystemNbi
//
// Use this API command to retrieve user information by Northbound Portal Interface.
//
// Optional Parameters:
// - domainId string
//		- nullable
func (s *WSGSystemService) FindSystemNbi(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGSystemNorthboundInterfaceAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSystemNorthboundInterfaceAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSystemNorthboundInterfaceAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindSystemNbi, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["domainId"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("domainId", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemNorthboundInterfaceAPIResponse), err
}

// FindSystemSecuritySetting
//
// Operation ID: findSystemSecuritySetting
//
// Use this API command to retrieve the security setting.
func (s *WSGSystemService) FindSystemSecuritySetting(ctx context.Context, mutators ...RequestMutator) (*WSGSystemSecuritySettingAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSystemSecuritySettingAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSystemSecuritySettingAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindSystemSecuritySetting, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemSecuritySettingAPIResponse), err
}

// FindSystemSystemTime
//
// Operation ID: findSystemSystemTime
//
// Retrieve System Time Setting.
func (s *WSGSystemService) FindSystemSystemTime(ctx context.Context, mutators ...RequestMutator) (*WSGSystemTimeSettingAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSystemTimeSettingAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSystemTimeSettingAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindSystemSystemTime, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemTimeSettingAPIResponse), err
}

// PartialUpdateSystem
//
// Operation ID: partialUpdateSystem
//
// Use this API command to modify settings of system. Currently, Only can modify settings about AP number limit by query criteria with domain and zone filters.
//
// Request Body:
//	 - body *WSGSystemSettings
func (s *WSGSystemService) PartialUpdateSystem(ctx context.Context, body *WSGSystemSettings, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateSystem, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// PartialUpdateSystemGatewayAdvanced
//
// Operation ID: partialUpdateSystemGatewayAdvanced
//
// Use this API command to modify the gateway advanced setting.
//
// Request Body:
//	 - body *WSGSystemModifyGatewayAdvanced
func (s *WSGSystemService) PartialUpdateSystemGatewayAdvanced(ctx context.Context, body *WSGSystemModifyGatewayAdvanced, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateSystemGatewayAdvanced, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// PartialUpdateSystemNbi
//
// Operation ID: partialUpdateSystemNbi
//
// Use this API command to modify the user information by Northbound Portal Interface.
//
// Request Body:
//	 - body *WSGSystemNorthboundInterface
//
// Optional Parameters:
// - domainId string
//		- nullable
func (s *WSGSystemService) PartialUpdateSystemNbi(ctx context.Context, body *WSGSystemNorthboundInterface, optionalParams map[string][]string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateSystemNbi, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	if v, ok := optionalParams["domainId"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("domainId", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// PartialUpdateSystemSystemTime
//
// Operation ID: partialUpdateSystemSystemTime
//
// Modify System Time Setting.
//
// Request Body:
//	 - body *WSGSystemModifySystemTimeSetting
func (s *WSGSystemService) PartialUpdateSystemSystemTime(ctx context.Context, body *WSGSystemModifySystemTimeSetting, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateSystemSystemTime, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// UpdateSystemApMacOUIsByOUI
//
// Operation ID: updateSystemApMacOUIsByOUI
//
// Use this API command to update AP Mac OUI.
//
// Request Body:
//	 - body *WSGSystemUpdateApMacOUI
//
// Required Parameters:
// - OUI string
//		- required
func (s *WSGSystemService) UpdateSystemApMacOUIsByOUI(ctx context.Context, body *WSGSystemUpdateApMacOUI, OUI string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateSystemApMacOUIsByOUI, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("OUI", OUI)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// UpdateSystemSecuritySetting
//
// Operation ID: updateSystemSecuritySetting
//
// Use this API command to retrieve the security setting.
//
// Request Body:
//	 - body *WSGSystemSecuritySetting
func (s *WSGSystemService) UpdateSystemSecuritySetting(ctx context.Context, body *WSGSystemSecuritySetting, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateSystemSecuritySetting, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
