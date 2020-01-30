package vsz

// API Version: v9_0

import (
	"encoding/json"
)

type WSGCommonAltitude struct {
	// AltitudeUnit
	// altitude unit
	// Constraints:
	//    - default:'meters'
	//    - oneof:[meters,floor]
	AltitudeUnit *string `json:"altitudeUnit,omitempty" validate:"oneof=meters floor"`

	// AltitudeValue
	// altitude value
	AltitudeValue *int `json:"altitudeValue,omitempty"`
}

func NewWSGCommonAltitude() *WSGCommonAltitude {
	m := new(WSGCommonAltitude)
	return m
}

// WSGCommonApGpsSource
//
// GPS Source of the AP
// Constraints:
//    - oneof:[GPS,MANUAL]
type WSGCommonApGpsSource string

func NewWSGCommonApGpsSource() *WSGCommonApGpsSource {
	m := new(WSGCommonApGpsSource)
	return m
}

type WSGCommonApLatencyInterval struct {
	// PingEnabled
	// AP ping latency enabled
	PingEnabled *bool `json:"pingEnabled,omitempty"`
}

func NewWSGCommonApLatencyInterval() *WSGCommonApLatencyInterval {
	m := new(WSGCommonApLatencyInterval)
	return m
}

type WSGCommonApLoginName string

func NewWSGCommonApLoginName() *WSGCommonApLoginName {
	m := new(WSGCommonApLoginName)
	return m
}

type WSGCommonApLoginPassword string

func NewWSGCommonApLoginPassword() *WSGCommonApLoginPassword {
	m := new(WSGCommonApLoginPassword)
	return m
}

type WSGCommonApManagementVlan struct {
	// Id
	// Vlan id of the zone
	Id *int `json:"id,omitempty"`

	// Mode
	// Vlan Mode of the zone
	// Constraints:
	//    - default:'KEEP'
	//    - oneof:[KEEP,USER_DEFINED]
	Mode *string `json:"mode,omitempty" validate:"oneof=KEEP USER_DEFINED"`
}

func NewWSGCommonApManagementVlan() *WSGCommonApManagementVlan {
	m := new(WSGCommonApManagementVlan)
	return m
}

type WSGCommonApRadio50 struct {
	// AutoCellSizing
	// Auto Cell Sizing
	AutoCellSizing *bool `json:"autoCellSizing,omitempty"`

	// Channel
	// channel number
	Channel *int `json:"channel,omitempty"`

	// ChannelRange
	// channel range options
	ChannelRange []int `json:"channelRange,omitempty"`

	// ChannelWidth
	// channel width, 0 mean Auto, 8080 means 80+80MHz
	// Constraints:
	//    - oneof:[0,20,40,80,8080,160]
	ChannelWidth *int `json:"channelWidth,omitempty" validate:"oneof=0 20 40 80 8080 160"`

	// SecondaryChannel
	// channel number (channelWidth is 80+80MHz only)
	SecondaryChannel *int `json:"secondaryChannel,omitempty"`

	TxPower *WSGCommonTxPower `json:"txPower,omitempty"`
}

func NewWSGCommonApRadio50() *WSGCommonApRadio50 {
	m := new(WSGCommonApRadio50)
	return m
}

type WSGCommonApRebootTimeout struct {
	// GatewayLossTimeoutInSec
	// Gateway loss timeout in second
	// Constraints:
	//    - default:1800
	//    - oneof:[0,1800,3600,5400,7200,9000,10800,12600,14400,16200,18000,19800,23400,25200,27000,28800,30600,32400,34200,36000,37800,39600,41400,43200,45000,46800,48600,50400,52200,54000,55800,57600,59400,61200,63000,64800,66600,68400,70200,72000,73800,75600,77400,79200,81000,82800,84600,86400]
	GatewayLossTimeoutInSec *int `json:"gatewayLossTimeoutInSec,omitempty" validate:"oneof=0 1800 3600 5400 7200 9000 10800 12600 14400 16200 18000 19800 23400 25200 27000 28800 30600 32400 34200 36000 37800 39600 41400 43200 45000 46800 48600 50400 52200 54000 55800 57600 59400 61200 63000 64800 66600 68400 70200 72000 73800 75600 77400 79200 81000 82800 84600 86400"`

	// ServerLossTimeoutInSec
	// Server loss timeout in second
	// Constraints:
	//    - default:7200
	//    - oneof:[0,7200,14400,21600,28800,36000,43200,50400,57600,64800,72000,79200,86400]
	ServerLossTimeoutInSec *int `json:"serverLossTimeoutInSec,omitempty" validate:"oneof=0 7200 14400 21600 28800 36000 43200 50400 57600 64800 72000 79200 86400"`
}

func NewWSGCommonApRebootTimeout() *WSGCommonApRebootTimeout {
	m := new(WSGCommonApRebootTimeout)
	return m
}

type WSGCommonAutoChannelSelection struct {
	// ChannelFlyMtbc
	// ChannelFly MTBC
	// Constraints:
	//    - default:480
	//    - min:100
	//    - max:1440
	ChannelFlyMtbc *int `json:"channelFlyMtbc,omitempty" validate:"gte=100,lte=1440"`

	// ChannelSelectMode
	// Channel Select Mode
	// Constraints:
	//    - default:'BackgroundScanning'
	//    - oneof:[None,BackgroundScanning,ChannelFly]
	ChannelSelectMode *string `json:"channelSelectMode,omitempty" validate:"oneof=None BackgroundScanning ChannelFly"`
}

func NewWSGCommonAutoChannelSelection() *WSGCommonAutoChannelSelection {
	m := new(WSGCommonAutoChannelSelection)
	return m
}

// WSGCommonAwsVenue
//
// Venue Code
// Constraints:
//    - max:64
//    - min:0
type WSGCommonAwsVenue string

func NewWSGCommonAwsVenue() *WSGCommonAwsVenue {
	m := new(WSGCommonAwsVenue)
	return m
}

type WSGCommonBaseServiceInfo struct {
	// Id
	// ID of service
	Id *string `json:"id,omitempty"`

	// ServiceId
	// ID of service
	ServiceId *string `json:"serviceId,omitempty"`

	// ServiceName
	// Name of service
	ServiceName *string `json:"serviceName,omitempty"`

	// ServiceType
	// Type of service
	ServiceType *string `json:"serviceType,omitempty"`
}

func NewWSGCommonBaseServiceInfo() *WSGCommonBaseServiceInfo {
	m := new(WSGCommonBaseServiceInfo)
	return m
}

type WSGCommonBulkDeleteRequest struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGCommonBulkDeleteRequest() *WSGCommonBulkDeleteRequest {
	m := new(WSGCommonBulkDeleteRequest)
	return m
}

type WSGCommonClientAdmissionControl struct {
	// MaxRadioLoadPercent
	// Maximum radio load percentage.
	// Constraints:
	//    - default:75
	//    - min:50
	//    - max:100
	MaxRadioLoadPercent *int `json:"maxRadioLoadPercent,omitempty" validate:"gte=50,lte=100"`

	// MinClientCount
	// Minimum client count number.
	// Constraints:
	//    - default:10
	//    - min:0
	//    - max:100
	MinClientCount *int `json:"minClientCount,omitempty" validate:"gte=0,lte=100"`

	// MinClientThroughputMbps
	// Minimum client throughput in Mbps.
	// Constraints:
	//    - default:0
	//    - min:0.000000
	//    - max:100.000000
	MinClientThroughputMbps *float64 `json:"minClientThroughputMbps,omitempty" validate:"gte=0.000000,lte=100.000000"`
}

func NewWSGCommonClientAdmissionControl() *WSGCommonClientAdmissionControl {
	m := new(WSGCommonClientAdmissionControl)
	return m
}

type WSGCommonCreateResult struct {
	Id *string `json:"id,omitempty"`
}

func NewWSGCommonCreateResult() *WSGCommonCreateResult {
	m := new(WSGCommonCreateResult)
	return m
}

type WSGCommonCreateResultIdName struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

func NewWSGCommonCreateResultIdName() *WSGCommonCreateResultIdName {
	m := new(WSGCommonCreateResultIdName)
	return m
}

type WSGCommonDescription string

func NewWSGCommonDescription() *WSGCommonDescription {
	m := new(WSGCommonDescription)
	return m
}

type WSGCommonDescriptionTo128 string

func NewWSGCommonDescriptionTo128() *WSGCommonDescriptionTo128 {
	m := new(WSGCommonDescriptionTo128)
	return m
}

type WSGCommonDhcpProfileRef struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the DHCP Profile
	Id *string `json:"id,omitempty"`

	// LeaseTimeHours
	// Lease time in hours of the DHCP Profile
	// Constraints:
	//    - min:0
	//    - max:24
	LeaseTimeHours *int `json:"leaseTimeHours,omitempty" validate:"gte=0,lte=24"`

	// LeaseTimeMinutes
	// Lease time in minutes of the DHCP Profile
	// Constraints:
	//    - min:0
	//    - max:59
	LeaseTimeMinutes *int `json:"leaseTimeMinutes,omitempty" validate:"gte=0,lte=59"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	PoolEndIp *WSGCommonIpAddress `json:"poolEndIp,omitempty"`

	PoolStartIp *WSGCommonIpAddress `json:"poolStartIp,omitempty"`

	PrimaryDnsIp *WSGCommonIpAddress `json:"primaryDnsIp,omitempty"`

	SecondaryDnsIp *WSGCommonIpAddress `json:"secondaryDnsIp,omitempty"`

	SubnetMask *WSGCommonIpAddress `json:"subnetMask,omitempty"`

	SubnetNetworkIp *WSGCommonIpAddress `json:"subnetNetworkIp,omitempty"`

	// VlanId
	// VLAN ID of the DHCP Profile
	// Constraints:
	//    - min:1
	//    - max:4094
	VlanId *int `json:"vlanId,omitempty" validate:"gte=1,lte=4094"`

	// ZoneId
	// Zone Id of DHCP Profile
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGCommonDhcpProfileRef() *WSGCommonDhcpProfileRef {
	m := new(WSGCommonDhcpProfileRef)
	return m
}

type WSGCommonDhcpSiteConfigListRef struct {
	// DwpdEnabled
	// DHCP Service Dynamic WAN Port Detection
	DwpdEnabled *bool `json:"dwpdEnabled,omitempty"`

	Eth0ProfileId *int `json:"eth0ProfileId,omitempty"`

	Eth1ProfileId *int `json:"eth1ProfileId,omitempty"`

	// ManualSelect
	// DHCP Service AP Selection Mode
	ManualSelect *bool `json:"manualSelect,omitempty"`

	SiteAps []*WSGCommonDhcpSiteConfigListRefSiteApsType `json:"siteAps,omitempty"`

	// SiteEnabled
	// DHCP Service Enabling Status
	SiteEnabled *bool `json:"siteEnabled,omitempty"`

	// SiteMode
	// DHCP Service mode
	// Constraints:
	//    - oneof:[EnableOnEachAPs,EnableOnMultipleAPs,EnableOnHierarchicalAPs]
	SiteMode *string `json:"siteMode,omitempty" validate:"oneof=EnableOnEachAPs EnableOnMultipleAPs EnableOnHierarchicalAPs"`

	SiteProfiles []*WSGCommonDhcpProfileRef `json:"siteProfiles,omitempty"`

	// ZoneName
	// DHCP Service Zone Name
	ZoneName *string `json:"zoneName,omitempty"`
}

func NewWSGCommonDhcpSiteConfigListRef() *WSGCommonDhcpSiteConfigListRef {
	m := new(WSGCommonDhcpSiteConfigListRef)
	return m
}

// WSGCommonDhcpSiteConfigListRefSiteApsType
//
// DHCP Site selected APs
type WSGCommonDhcpSiteConfigListRefSiteApsType struct {
	ApGatewayIp *string `json:"apGatewayIp,omitempty"`

	ApMac *string `json:"apMac,omitempty"`

	ApName *string `json:"apName,omitempty"`

	ApServerEnabled *bool `json:"apServerEnabled,omitempty"`

	ApServerIp *string `json:"apServerIp,omitempty"`

	ApServerPrimary *bool `json:"apServerPrimary,omitempty"`

	ApServerType *string `json:"apServerType,omitempty"`

	// ApStatus
	// Constraints:
	//    - oneof:[Online,Offline,Flagged]
	ApStatus *string `json:"apStatus,omitempty" validate:"oneof=Online Offline Flagged"`
}

func NewWSGCommonDhcpSiteConfigListRefSiteApsType() *WSGCommonDhcpSiteConfigListRefSiteApsType {
	m := new(WSGCommonDhcpSiteConfigListRefSiteApsType)
	return m
}

type WSGCommonDhcpSiteConfigRef struct {
	// DwpdEnabled
	// DHCP Service Dynamic WAN Port Detection
	DwpdEnabled *bool `json:"dwpdEnabled,omitempty"`

	Eth0ProfileId *int `json:"eth0ProfileId,omitempty"`

	Eth1ProfileId *int `json:"eth1ProfileId,omitempty"`

	// ManualSelect
	// DHCP Service AP Selection Mode. This value is effective when the siteMode is EnableOnMultipleAPs.
	ManualSelect *bool `json:"manualSelect,omitempty"`

	SiteAps []*WSGCommonDhcpSiteConfigRefSiteApsType `json:"siteAps,omitempty"`

	// SiteEnabled
	// DHCP Service Enabling Status
	SiteEnabled *bool `json:"siteEnabled,omitempty"`

	// SiteMode
	// DHCP Service mode
	// Constraints:
	//    - oneof:[EnableOnEachAPs,EnableOnMultipleAPs,EnableOnHierarchicalAPs]
	SiteMode *string `json:"siteMode,omitempty" validate:"oneof=EnableOnEachAPs EnableOnMultipleAPs EnableOnHierarchicalAPs"`

	SiteProfileIds []string `json:"siteProfileIds,omitempty"`
}

func NewWSGCommonDhcpSiteConfigRef() *WSGCommonDhcpSiteConfigRef {
	m := new(WSGCommonDhcpSiteConfigRef)
	return m
}

// WSGCommonDhcpSiteConfigRefSiteApsType
//
// DHCP Site selected APs. The content is effective when the siteMode is EnableOnMultipleAPs.
type WSGCommonDhcpSiteConfigRefSiteApsType struct {
	ApGatewayIp *string `json:"apGatewayIp,omitempty"`

	ApMac *string `json:"apMac,omitempty"`

	ApName *string `json:"apName,omitempty"`

	ApServerEnabled *bool `json:"apServerEnabled,omitempty"`

	// ApServerIp
	// Constraints:
	//    - nullable
	ApServerIp *string `json:"apServerIp,omitempty" validate:"omitempty"`

	ApServerPrimary *bool `json:"apServerPrimary,omitempty"`

	ApServerType *string `json:"apServerType,omitempty"`

	// ApStatus
	// Constraints:
	//    - oneof:[Online,Offline,Flagged]
	ApStatus *string `json:"apStatus,omitempty" validate:"oneof=Online Offline Flagged"`
}

func NewWSGCommonDhcpSiteConfigRefSiteApsType() *WSGCommonDhcpSiteConfigRefSiteApsType {
	m := new(WSGCommonDhcpSiteConfigRefSiteApsType)
	return m
}

type WSGCommonDoAssignIp struct {
	// DwpdEnabled
	// DHCP Service Dynamic WAN Port Detection
	DwpdEnabled *bool `json:"dwpdEnabled,omitempty"`

	// ManualSelect
	// DHCP Service AP Selection Mode. This value is effective when the siteMode is EnableOnMultipleAPs.
	ManualSelect *bool `json:"manualSelect,omitempty"`

	SiteAps []*WSGCommonDoAssignIpSiteApsType `json:"siteAps,omitempty"`

	// SiteEnabled
	// DHCP Service Enabling Status
	SiteEnabled *bool `json:"siteEnabled,omitempty"`

	// SiteMode
	// DHCP Service mode
	// Constraints:
	//    - oneof:[EnableOnEachAPs,EnableOnMultipleAPs,EnableOnHierarchicalAPs]
	SiteMode *string `json:"siteMode,omitempty" validate:"oneof=EnableOnEachAPs EnableOnMultipleAPs EnableOnHierarchicalAPs"`

	SiteProfileIds []string `json:"siteProfileIds,omitempty"`
}

func NewWSGCommonDoAssignIp() *WSGCommonDoAssignIp {
	m := new(WSGCommonDoAssignIp)
	return m
}

// WSGCommonDoAssignIpSiteApsType
//
// DHCP Site selected APs. The content is effective when the siteMode is EnableOnMultipleAPs.
type WSGCommonDoAssignIpSiteApsType struct {
	ApMac *string `json:"apMac,omitempty"`

	ApServerEnabled *bool `json:"apServerEnabled,omitempty"`

	ApServerPrimary *bool `json:"apServerPrimary,omitempty"`
}

func NewWSGCommonDoAssignIpSiteApsType() *WSGCommonDoAssignIpSiteApsType {
	m := new(WSGCommonDoAssignIpSiteApsType)
	return m
}

type WSGCommonEmail string

func NewWSGCommonEmail() *WSGCommonEmail {
	m := new(WSGCommonEmail)
	return m
}

type WSGCommonEmptyResult struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGCommonEmptyResult) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGCommonEmptyResult{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGCommonEmptyResult) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGCommonEmptyResult() *WSGCommonEmptyResult {
	m := new(WSGCommonEmptyResult)
	return m
}

type WSGCommonEtherType string

func NewWSGCommonEtherType() *WSGCommonEtherType {
	m := new(WSGCommonEtherType)
	return m
}

type WSGCommonFilterOperator string

func NewWSGCommonFilterOperator() *WSGCommonFilterOperator {
	m := new(WSGCommonFilterOperator)
	return m
}

type WSGCommonFirmwareVersion string

func NewWSGCommonFirmwareVersion() *WSGCommonFirmwareVersion {
	m := new(WSGCommonFirmwareVersion)
	return m
}

type WSGCommonFQDN string

func NewWSGCommonFQDN() *WSGCommonFQDN {
	m := new(WSGCommonFQDN)
	return m
}

type WSGCommonFullTextSearch struct {
	// Fields
	// Specific fields to search
	Fields []string `json:"fields,omitempty"`

	// Type
	// Search logic operator
	// Constraints:
	//    - oneof:[AND,OR]
	Type *string `json:"type,omitempty" validate:"oneof=AND OR"`

	// Value
	// Text or number to search
	Value *string `json:"value,omitempty"`
}

func NewWSGCommonFullTextSearch() *WSGCommonFullTextSearch {
	m := new(WSGCommonFullTextSearch)
	return m
}

type WSGCommonGenericRef struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

func NewWSGCommonGenericRef() *WSGCommonGenericRef {
	m := new(WSGCommonGenericRef)
	return m
}

type WSGCommonHealthCheckPolicy struct {
	// ResponseFail
	// Response Fail
	ResponseFail *bool `json:"responseFail,omitempty"`

	// ResponseWindow
	// Response window
	// Constraints:
	//    - required
	//    - default:20
	//    - min:5
	//    - max:30
	ResponseWindow *int `json:"responseWindow" validate:"required,gte=5,lte=30"`

	// ReviveInterval
	// Revive interval
	// Constraints:
	//    - required
	//    - default:120
	//    - min:60
	//    - max:3600
	ReviveInterval *int `json:"reviveInterval" validate:"required,gte=60,lte=3600"`

	// ZombiePeriod
	// Zombie period
	// Constraints:
	//    - required
	//    - default:40
	//    - min:30
	//    - max:120
	ZombiePeriod *int `json:"zombiePeriod" validate:"required,gte=30,lte=120"`
}

func NewWSGCommonHealthCheckPolicy() *WSGCommonHealthCheckPolicy {
	m := new(WSGCommonHealthCheckPolicy)
	return m
}

type WSGCommonHTTPS string

func NewWSGCommonHTTPS() *WSGCommonHTTPS {
	m := new(WSGCommonHTTPS)
	return m
}

type WSGCommonIdList []string

func MakeWSGCommonIdList() WSGCommonIdList {
	m := make(WSGCommonIdList, 0)
	return m
}

type WSGCommonIpAddress string

func NewWSGCommonIpAddress() *WSGCommonIpAddress {
	m := new(WSGCommonIpAddress)
	return m
}

type WSGCommonIpMode string

func NewWSGCommonIpMode() *WSGCommonIpMode {
	m := new(WSGCommonIpMode)
	return m
}

type WSGCommonLanguageName string

func NewWSGCommonLanguageName() *WSGCommonLanguageName {
	m := new(WSGCommonLanguageName)
	return m
}

type WSGCommonLatitude float64

func NewWSGCommonLatitude() *WSGCommonLatitude {
	m := new(WSGCommonLatitude)
	return m
}

type WSGCommonLocation string

func NewWSGCommonLocation() *WSGCommonLocation {
	m := new(WSGCommonLocation)
	return m
}

type WSGCommonLocationAdditionalInfo string

func NewWSGCommonLocationAdditionalInfo() *WSGCommonLocationAdditionalInfo {
	m := new(WSGCommonLocationAdditionalInfo)
	return m
}

type WSGCommonLongitude float64

func NewWSGCommonLongitude() *WSGCommonLongitude {
	m := new(WSGCommonLongitude)
	return m
}

type WSGCommonLteBandLockChannel struct {
	// Channel3g
	// LTE 3G channels
	Channel3g *string `json:"channel3g,omitempty"`

	// Channel4g
	// LTE 4G channels
	Channel4g *string `json:"channel4g,omitempty"`

	// SimCardId
	// SIM card ID(Primary:0, Secondary:1)
	SimCardId *int `json:"simCardId,omitempty"`

	// Type
	// LTE chipset SKU type
	Type *string `json:"type,omitempty"`
}

func NewWSGCommonLteBandLockChannel() *WSGCommonLteBandLockChannel {
	m := new(WSGCommonLteBandLockChannel)
	return m
}

type WSGCommonMac string

func NewWSGCommonMac() *WSGCommonMac {
	m := new(WSGCommonMac)
	return m
}

type WSGCommonNormalName string

func NewWSGCommonNormalName() *WSGCommonNormalName {
	m := new(WSGCommonNormalName)
	return m
}

type WSGCommonNormalName2to64 string

func NewWSGCommonNormalName2to64() *WSGCommonNormalName2to64 {
	m := new(WSGCommonNormalName2to64)
	return m
}

type WSGCommonNormalName2to128 string

func NewWSGCommonNormalName2to128() *WSGCommonNormalName2to128 {
	m := new(WSGCommonNormalName2to128)
	return m
}

type WSGCommonNormalNameAllowBlank string

func NewWSGCommonNormalNameAllowBlank() *WSGCommonNormalNameAllowBlank {
	m := new(WSGCommonNormalNameAllowBlank)
	return m
}

type WSGCommonNormalURL string

func NewWSGCommonNormalURL() *WSGCommonNormalURL {
	m := new(WSGCommonNormalURL)
	return m
}

type WSGCommonOui string

func NewWSGCommonOui() *WSGCommonOui {
	m := new(WSGCommonOui)
	return m
}

type WSGCommonOverrideClientAdmissionControl struct {
	Enabled *bool `json:"enabled,omitempty"`

	// MaxRadioLoadPercent
	// Maximum radio load percentage.
	// Constraints:
	//    - default:75
	//    - min:50
	//    - max:100
	MaxRadioLoadPercent *int `json:"maxRadioLoadPercent,omitempty" validate:"gte=50,lte=100"`

	// MinClientCount
	// Minimum client count number.
	// Constraints:
	//    - default:10
	//    - min:0
	//    - max:100
	MinClientCount *int `json:"minClientCount,omitempty" validate:"gte=0,lte=100"`

	// MinClientThroughputMbps
	// Minimum client throughput in Mbps.
	// Constraints:
	//    - default:0
	//    - min:0.000000
	//    - max:100.000000
	MinClientThroughputMbps *float64 `json:"minClientThroughputMbps,omitempty" validate:"gte=0.000000,lte=100.000000"`
}

func NewWSGCommonOverrideClientAdmissionControl() *WSGCommonOverrideClientAdmissionControl {
	m := new(WSGCommonOverrideClientAdmissionControl)
	return m
}

type WSGCommonOverrideGenericRef struct {
	Enabled *bool `json:"enabled,omitempty"`

	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

func NewWSGCommonOverrideGenericRef() *WSGCommonOverrideGenericRef {
	m := new(WSGCommonOverrideGenericRef)
	return m
}

type WSGCommonOverrideSmartMonitor struct {
	Enabled *bool `json:"enabled,omitempty"`

	// IntervalInSec
	// Interval in seconds. This is required if smartMonitor is enabled
	// Constraints:
	//    - default:10
	//    - min:5
	//    - max:60
	IntervalInSec *int `json:"intervalInSec,omitempty" validate:"gte=5,lte=60"`

	// RetryThreshold
	// Retry threshold. This is required if smartMonitor is enabled
	// Constraints:
	//    - default:3
	//    - min:1
	//    - max:10
	RetryThreshold *int `json:"retryThreshold,omitempty" validate:"gte=1,lte=10"`
}

func NewWSGCommonOverrideSmartMonitor() *WSGCommonOverrideSmartMonitor {
	m := new(WSGCommonOverrideSmartMonitor)
	return m
}

type WSGCommonPoeModeSetting string

func NewWSGCommonPoeModeSetting() *WSGCommonPoeModeSetting {
	m := new(WSGCommonPoeModeSetting)
	return m
}

type WSGCommonPortalCustomization struct {
	// Language
	// Constraints:
	//    - required
	Language *WSGCommonPortalLanguage `json:"language" validate:"required,oneof=en_US zh_TW zh_CN nl_NL fr_FR de_DE ja_JP es_ES se_SE ar_SA cz_CZ da_DK tr_TR pt_BR"`

	// Logo
	// logo
	Logo *string `json:"logo,omitempty"`

	// TermsAndConditionsRequired
	// Terms and conditions is required or not
	TermsAndConditionsRequired *bool `json:"termsAndConditionsRequired,omitempty"`

	// TermsAndConditionsText
	// Terms and conditions text
	// Constraints:
	//    - default:'Terms of Use
	//
	// By accepting this agreement and accessing the wireless network, you acknowledge that you are of legal age, you have read and understood, and agree to be bound by this agreement.
	// (*) The wireless network service is provided by the property owners and is completely at their discretion. Your access to the network may be blocked, suspended, or terminated at any time for any reason.
	// (*) You agree not to use the wireless network for any purpose that is unlawful or otherwise prohibited and you are fully responsible for your use.
	// (*) The wireless network is provided "as is" without warranties of any kind, either expressed or implied.
	//
	// This wireless network is powered by Ruckus Wireless.'
	//    - max:16000
	//    - min:0
	TermsAndConditionsText *string `json:"termsAndConditionsText,omitempty" validate:"max=16000,min=0"`

	// Title
	// Title
	// Constraints:
	//    - max:63
	//    - min:0
	Title *string `json:"title,omitempty" validate:"max=63,min=0"`
}

func NewWSGCommonPortalCustomization() *WSGCommonPortalCustomization {
	m := new(WSGCommonPortalCustomization)
	return m
}

// WSGCommonPortalLanguage
//
// Language
// Constraints:
//    - default:'en_US'
//    - oneof:[en_US,zh_TW,zh_CN,nl_NL,fr_FR,de_DE,ja_JP,es_ES,se_SE,ar_SA,cz_CZ,da_DK,tr_TR,pt_BR]
type WSGCommonPortalLanguage string

func NewWSGCommonPortalLanguage() *WSGCommonPortalLanguage {
	m := new(WSGCommonPortalLanguage)
	return m
}

type WSGCommonProtectionMode string

func NewWSGCommonProtectionMode() *WSGCommonProtectionMode {
	m := new(WSGCommonProtectionMode)
	return m
}

type WSGCommonQinq struct {
	// Cvlan
	// Constraints:
	//    - min:2
	//    - max:4094
	Cvlan *int `json:"cvlan,omitempty" validate:"gte=2,lte=4094"`

	// Svlan
	// Constraints:
	//    - min:2
	//    - max:4094
	Svlan *int `json:"svlan,omitempty" validate:"gte=2,lte=4094"`
}

func NewWSGCommonQinq() *WSGCommonQinq {
	m := new(WSGCommonQinq)
	return m
}

type WSGCommonQueryCriteria struct {
	// Attributes
	// Get specific columns only
	Attributes []string `json:"attributes,omitempty"`

	// Criteria
	// Add backward compatibility for UI framework
	Criteria *string `json:"criteria,omitempty"`

	// ExpandDomains
	// Whether to expand domains into sub domains/ zones or not
	ExpandDomains *bool `json:"expandDomains,omitempty"`

	// ExtraFilters
	// "AND" condition for multiple filters
	ExtraFilters []*WSGCommonQueryCriteriaExtraFiltersType `json:"extraFilters,omitempty"`

	// ExtraNotFilters
	// "NOT" condition for multiple filters
	ExtraNotFilters []*WSGCommonQueryCriteriaExtraNotFiltersType `json:"extraNotFilters,omitempty"`

	ExtraTimeRange *WSGCommonTimeRange `json:"extraTimeRange,omitempty"`

	// Filters
	// Filters used to select specific resource scope
	Filters []*WSGCommonQueryCriteriaFiltersType `json:"filters,omitempty"`

	FullTextSearch *WSGCommonFullTextSearch `json:"fullTextSearch,omitempty"`

	// Limit
	// Size of one page
	// Constraints:
	//    - min:1
	Limit *int `json:"limit,omitempty" validate:"gte=1"`

	// Options
	// Specified feature required information
	Options *WSGCommonQueryCriteriaOptionsType `json:"options,omitempty"`

	// Page
	// Page number to get
	// Constraints:
	//    - min:1
	Page *int `json:"page,omitempty" validate:"gte=1"`

	// Query
	// Add backward compatibility for UI framework
	Query *string `json:"query,omitempty"`

	// SortInfo
	// About sorting
	SortInfo *WSGCommonQueryCriteriaSortInfoType `json:"sortInfo,omitempty"`
}

func NewWSGCommonQueryCriteria() *WSGCommonQueryCriteria {
	m := new(WSGCommonQueryCriteria)
	return m
}

type WSGCommonQueryCriteriaExtraFiltersType struct {
	Operator *WSGCommonFilterOperator `json:"operator,omitempty"`

	// Type
	// Filters for specific attributes
	Type *string `json:"type,omitempty"`

	// Value
	// value to search
	Value *string `json:"value,omitempty"`
}

func NewWSGCommonQueryCriteriaExtraFiltersType() *WSGCommonQueryCriteriaExtraFiltersType {
	m := new(WSGCommonQueryCriteriaExtraFiltersType)
	return m
}

type WSGCommonQueryCriteriaExtraNotFiltersType struct {
	// Type
	// Filters for specific attribute
	Type *string `json:"type,omitempty"`

	// Value
	// value not to search
	Value *string `json:"value,omitempty"`
}

func NewWSGCommonQueryCriteriaExtraNotFiltersType() *WSGCommonQueryCriteriaExtraNotFiltersType {
	m := new(WSGCommonQueryCriteriaExtraNotFiltersType)
	return m
}

type WSGCommonQueryCriteriaFiltersType struct {
	Operator *WSGCommonFilterOperator `json:"operator,omitempty"`

	// Type
	// Group type
	Type *string `json:"type,omitempty"`

	// Value
	// Group ID
	Value *string `json:"value,omitempty"`
}

func NewWSGCommonQueryCriteriaFiltersType() *WSGCommonQueryCriteriaFiltersType {
	m := new(WSGCommonQueryCriteriaFiltersType)
	return m
}

// WSGCommonQueryCriteriaOptionsType
//
// Specified feature required information
type WSGCommonQueryCriteriaOptionsType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGCommonQueryCriteriaOptionsType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGCommonQueryCriteriaOptionsType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGCommonQueryCriteriaOptionsType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGCommonQueryCriteriaOptionsType() *WSGCommonQueryCriteriaOptionsType {
	m := new(WSGCommonQueryCriteriaOptionsType)
	return m
}

// WSGCommonQueryCriteriaSortInfoType
//
// About sorting
type WSGCommonQueryCriteriaSortInfoType struct {
	// Dir
	// Constraints:
	//    - oneof:[ASC,DESC]
	Dir *string `json:"dir,omitempty" validate:"oneof=ASC DESC"`

	SortColumn *string `json:"sortColumn,omitempty"`
}

func NewWSGCommonQueryCriteriaSortInfoType() *WSGCommonQueryCriteriaSortInfoType {
	m := new(WSGCommonQueryCriteriaSortInfoType)
	return m
}

type WSGCommonQueryCriteriaSuperSet struct {
	// Attributes
	// Get specific columns only
	Attributes []string `json:"attributes,omitempty"`

	// Criteria
	// Add backward compatibility for UI framework
	Criteria *string `json:"criteria,omitempty"`

	// ExpandDomains
	// Whether to expand domains into sub domains/ zones or not
	ExpandDomains *bool `json:"expandDomains,omitempty"`

	// ExtraFilters
	// "AND" condition for multiple filters
	ExtraFilters []*WSGCommonQueryCriteriaExtraFiltersType `json:"extraFilters,omitempty"`

	// ExtraNotFilters
	// "NOT" condition for multiple filters
	ExtraNotFilters []*WSGCommonQueryCriteriaExtraNotFiltersType `json:"extraNotFilters,omitempty"`

	ExtraTimeRange *WSGCommonTimeRange `json:"extraTimeRange,omitempty"`

	// Filters
	// Filters used to select specific resource scope
	Filters []*WSGCommonQueryCriteriaFiltersType `json:"filters,omitempty"`

	FullTextSearch *WSGCommonFullTextSearch `json:"fullTextSearch,omitempty"`

	// Limit
	// Size of one page
	// Constraints:
	//    - min:1
	Limit *int `json:"limit,omitempty" validate:"gte=1"`

	// Options
	// Specified feature required information
	Options *WSGCommonQueryCriteriaOptionsType `json:"options,omitempty"`

	// Page
	// Page number to get
	// Constraints:
	//    - min:1
	Page *int `json:"page,omitempty" validate:"gte=1"`

	// Query
	// Add backward compatibility for UI framework
	Query *string `json:"query,omitempty"`

	// SortInfo
	// About sorting
	SortInfo *WSGCommonQueryCriteriaSortInfoType `json:"sortInfo,omitempty"`
}

func NewWSGCommonQueryCriteriaSuperSet() *WSGCommonQueryCriteriaSuperSet {
	m := new(WSGCommonQueryCriteriaSuperSet)
	return m
}

type WSGCommonRadio24 struct {
	// AutoCellSizing
	// Auto Cell Sizing
	AutoCellSizing *bool `json:"autoCellSizing,omitempty"`

	// Channel
	// Channel number
	Channel *int `json:"channel,omitempty"`

	// ChannelRange
	// Channel range options
	ChannelRange []int `json:"channelRange,omitempty"`

	// ChannelWidth
	// Channel width. Zero (0) means Auto.
	// Constraints:
	//    - default:0
	//    - oneof:[0,20,40]
	ChannelWidth *int `json:"channelWidth,omitempty" validate:"oneof=0 20 40"`

	TxPower *WSGCommonTxPower `json:"txPower,omitempty"`
}

func NewWSGCommonRadio24() *WSGCommonRadio24 {
	m := new(WSGCommonRadio24)
	return m
}

type WSGCommonRadio24SuperSet struct {
	// AutoCellSizing
	// Auto Cell Sizing
	AutoCellSizing *bool `json:"autoCellSizing,omitempty"`

	// AvailableChannelRange
	// Available channel range options
	AvailableChannelRange []int `json:"availableChannelRange,omitempty"`

	// Channel
	// Channel number
	Channel *int `json:"channel,omitempty"`

	// ChannelRange
	// Channel range options
	ChannelRange []int `json:"channelRange,omitempty"`

	// ChannelWidth
	// Channel width. Zero (0) means Auto.
	// Constraints:
	//    - oneof:[0,20,40]
	ChannelWidth *int `json:"channelWidth,omitempty" validate:"oneof=0 20 40"`

	TxPower *WSGCommonTxPower `json:"txPower,omitempty"`
}

func NewWSGCommonRadio24SuperSet() *WSGCommonRadio24SuperSet {
	m := new(WSGCommonRadio24SuperSet)
	return m
}

type WSGCommonRadio50 struct {
	// AutoCellSizing
	// Auto Cell Sizing
	AutoCellSizing *bool `json:"autoCellSizing,omitempty"`

	// ChannelWidth
	// Channel width. Zero (0) means Auto. 8080 means 80+80MHz
	// Constraints:
	//    - default:0
	//    - oneof:[0,20,40,80,8080,160]
	ChannelWidth *int `json:"channelWidth,omitempty" validate:"oneof=0 20 40 80 8080 160"`

	// IndoorChannel
	// Channel number for Indoor AP
	IndoorChannel *int `json:"indoorChannel,omitempty"`

	// IndoorChannelRange
	// Channel range options for Indoor AP
	IndoorChannelRange []int `json:"indoorChannelRange,omitempty"`

	// IndoorSecondaryChannel
	// Secondary channel number for Indoor AP (channelWidth is 80+80MHz only)
	IndoorSecondaryChannel *int `json:"indoorSecondaryChannel,omitempty"`

	// OutdoorChannel
	// Channel number for Outdoor AP
	OutdoorChannel *int `json:"outdoorChannel,omitempty"`

	// OutdoorChannelRange
	// Channel range options for outdoor AP
	OutdoorChannelRange []int `json:"outdoorChannelRange,omitempty"`

	// OutdoorSecondaryChannel
	// Secondary channel number for outdoor AP (channelWidth is 80+80MHz only)
	OutdoorSecondaryChannel *int `json:"outdoorSecondaryChannel,omitempty"`

	TxPower *WSGCommonTxPower `json:"txPower,omitempty"`
}

func NewWSGCommonRadio50() *WSGCommonRadio50 {
	m := new(WSGCommonRadio50)
	return m
}

type WSGCommonRadio50SuperSet struct {
	// AutoCellSizing
	// Auto Cell Sizing
	AutoCellSizing *bool `json:"autoCellSizing,omitempty"`

	// AvailableIndoorChannelRange
	// Available channel range options
	AvailableIndoorChannelRange []int `json:"availableIndoorChannelRange,omitempty"`

	// AvailableOutdoorChannelRange
	// Available channel range options
	AvailableOutdoorChannelRange []int `json:"availableOutdoorChannelRange,omitempty"`

	// ChannelWidth
	// Channel width. Zero (0) means Auto. 8080 means 80+80MHz
	// Constraints:
	//    - oneof:[0,20,40,80,8080,160]
	ChannelWidth *int `json:"channelWidth,omitempty" validate:"oneof=0 20 40 80 8080 160"`

	// IndoorChannel
	// Channel number for Indoor AP
	IndoorChannel *int `json:"indoorChannel,omitempty"`

	// IndoorChannelRange
	// Channel range options for Indoor AP
	IndoorChannelRange []int `json:"indoorChannelRange,omitempty"`

	// IndoorSecondaryChannel
	// Secondary channel number for Indoor AP (channelWidth is 80+80MHz only)
	IndoorSecondaryChannel *int `json:"indoorSecondaryChannel,omitempty"`

	// OutdoorChannel
	// Channel number for Outdoor AP
	OutdoorChannel *int `json:"outdoorChannel,omitempty"`

	// OutdoorChannelRange
	// Channel range options for outdoor AP
	OutdoorChannelRange []int `json:"outdoorChannelRange,omitempty"`

	// OutdoorSecondaryChannel
	// Secondary channel number for outdoor AP (channelWidth is 80+80MHz only)
	OutdoorSecondaryChannel *int `json:"outdoorSecondaryChannel,omitempty"`

	TxPower *WSGCommonTxPower `json:"txPower,omitempty"`
}

func NewWSGCommonRadio50SuperSet() *WSGCommonRadio50SuperSet {
	m := new(WSGCommonRadio50SuperSet)
	return m
}

type WSGCommonRadiusServer struct {
	// Ip
	// Constraints:
	//    - required
	Ip *WSGCommonIpAddress `json:"ip" validate:"required"`

	// Port
	// Server port
	// Constraints:
	//    - required
	//    - min:1
	//    - max:65535
	Port *int `json:"port" validate:"required,gte=1,lte=65535"`

	// SharedSecret
	// Server shared secrect
	// Constraints:
	//    - required
	SharedSecret *string `json:"sharedSecret" validate:"required"`
}

func NewWSGCommonRadiusServer() *WSGCommonRadiusServer {
	m := new(WSGCommonRadiusServer)
	return m
}

type WSGCommonRateLimiting struct {
	// MaxOutstandingRequestsPerServer
	// Maximum outstanding requests (MOR), value should be 0 or between 10 and 4096
	// Constraints:
	//    - required
	//    - default:0
	MaxOutstandingRequestsPerServer *int `json:"maxOutstandingRequestsPerServer" validate:"required"`

	// SanityTimer
	// Sanity timer
	// Constraints:
	//    - required
	//    - default:10
	//    - min:1
	//    - max:3600
	SanityTimer *int `json:"sanityTimer" validate:"required,gte=1,lte=3600"`

	// Threshold
	// Threshold, value should be 0 if MOR is 0, or between 10 and 90 if MOR is between 10 and 4096
	// Constraints:
	//    - required
	//    - default:0
	Threshold *int `json:"threshold" validate:"required"`
}

func NewWSGCommonRateLimiting() *WSGCommonRateLimiting {
	m := new(WSGCommonRateLimiting)
	return m
}

type WSGCommonRbacMetadata struct {
	RbacMetadata []string `json:"rbacMetadata,omitempty"`
}

func NewWSGCommonRbacMetadata() *WSGCommonRbacMetadata {
	m := new(WSGCommonRbacMetadata)
	return m
}

type WSGCommonRealm string

func NewWSGCommonRealm() *WSGCommonRealm {
	m := new(WSGCommonRealm)
	return m
}

type WSGCommonRecoverySsid struct {
	// RecoverySsidEnabled
	// recovery ssid enable/disable
	RecoverySsidEnabled *bool `json:"recoverySsidEnabled,omitempty"`
}

func NewWSGCommonRecoverySsid() *WSGCommonRecoverySsid {
	m := new(WSGCommonRecoverySsid)
	return m
}

type WSGCommonSmartMonitor struct {
	// IntervalInSec
	// Interval in seconds. This is required if smartMonitor is enabled
	// Constraints:
	//    - default:10
	//    - min:5
	//    - max:60
	IntervalInSec *int `json:"intervalInSec,omitempty" validate:"gte=5,lte=60"`

	// RetryThreshold
	// Retry threshold. This is required if smartMonitor is enabled
	// Constraints:
	//    - default:3
	//    - min:1
	//    - max:10
	RetryThreshold *int `json:"retryThreshold,omitempty" validate:"gte=1,lte=10"`
}

func NewWSGCommonSmartMonitor() *WSGCommonSmartMonitor {
	m := new(WSGCommonSmartMonitor)
	return m
}

type WSGCommonSnmpCommunity struct {
	// CommunityName
	// name of the SNMP Community.
	// Constraints:
	//    - required
	CommunityName *string `json:"communityName" validate:"required"`

	// NotificationEnabled
	// notification privilege of the SNMP Coummunity
	NotificationEnabled *bool `json:"notificationEnabled,omitempty"`

	// NotificationTarget
	// Trap List of the SNMP Coummunity
	NotificationTarget []*WSGCommonTargetConfig `json:"notificationTarget,omitempty"`

	// NotificationType
	// type of the notification privilege
	// Constraints:
	//    - oneof:[TRAP,INFORM]
	NotificationType *string `json:"notificationType,omitempty" validate:"oneof=TRAP INFORM"`

	// ReadEnabled
	// read privilege of the SNMP Coummunity
	ReadEnabled *bool `json:"readEnabled,omitempty"`

	// WriteEnabled
	// write privilege of the SNMP Coummunity
	WriteEnabled *bool `json:"writeEnabled,omitempty"`
}

func NewWSGCommonSnmpCommunity() *WSGCommonSnmpCommunity {
	m := new(WSGCommonSnmpCommunity)
	return m
}

type WSGCommonSnmpUser struct {
	// AuthPassword
	// authPassword of the SNMP User.
	// Constraints:
	//    - min:8
	AuthPassword *string `json:"authPassword,omitempty" validate:"min=8"`

	// AuthProtocol
	// authProtocol of the SNMP User.
	// Constraints:
	//    - oneof:[MD5,SHA]
	AuthProtocol *string `json:"authProtocol,omitempty" validate:"oneof=MD5 SHA"`

	// NotificationEnabled
	// notification privilege of the SNMP User
	NotificationEnabled *bool `json:"notificationEnabled,omitempty"`

	// NotificationTarget
	// Trap List of the SNMP User
	NotificationTarget []*WSGCommonTargetConfig `json:"notificationTarget,omitempty"`

	// NotificationType
	// type of the notification privilege
	// Constraints:
	//    - oneof:[TRAP,INFORM]
	NotificationType *string `json:"notificationType,omitempty" validate:"oneof=TRAP INFORM"`

	// PrivPassword
	// privPassword of the SNMP User.
	// Constraints:
	//    - min:8
	PrivPassword *string `json:"privPassword,omitempty" validate:"min=8"`

	// PrivProtocol
	// privProtocol of the SNMP User.
	// Constraints:
	//    - oneof:[DES,AES]
	PrivProtocol *string `json:"privProtocol,omitempty" validate:"oneof=DES AES"`

	// ReadEnabled
	// read privilege of the SNMP User
	ReadEnabled *bool `json:"readEnabled,omitempty"`

	// UserName
	// name of the SNMP User.
	// Constraints:
	//    - required
	UserName *string `json:"userName" validate:"required"`

	// WriteEnabled
	// write privilege of the SNMP User
	WriteEnabled *bool `json:"writeEnabled,omitempty"`
}

func NewWSGCommonSnmpUser() *WSGCommonSnmpUser {
	m := new(WSGCommonSnmpUser)
	return m
}

type WSGCommonSubNetMask string

func NewWSGCommonSubNetMask() *WSGCommonSubNetMask {
	m := new(WSGCommonSubNetMask)
	return m
}

type WSGCommonTargetConfig struct {
	// Address
	// address of the SNMP Trap
	// Constraints:
	//    - required
	Address *string `json:"address" validate:"required"`

	// Port
	// port number of the SNMP Trap
	// Constraints:
	//    - required
	//    - min:1
	//    - max:65535
	Port *int `json:"port" validate:"required,gte=1,lte=65535"`
}

func NewWSGCommonTargetConfig() *WSGCommonTargetConfig {
	m := new(WSGCommonTargetConfig)
	return m
}

type WSGCommonTimeRange struct {
	// End
	// end time for collecting data
	End *int `json:"end,omitempty"`

	// Field
	// time field for collecting data
	// Constraints:
	//    - oneof:[insertionTime]
	Field *string `json:"field,omitempty" validate:"oneof=insertionTime"`

	// Interval
	// time interval in second
	Interval *int `json:"interval,omitempty"`

	// Start
	// start time for collecting data
	Start *int `json:"start,omitempty"`
}

func NewWSGCommonTimeRange() *WSGCommonTimeRange {
	m := new(WSGCommonTimeRange)
	return m
}

// WSGCommonTimeUnitStore
//
// time unit
// Constraints:
//    - oneof:[second,minute,hour,day]
type WSGCommonTimeUnitStore string

func NewWSGCommonTimeUnitStore() *WSGCommonTimeUnitStore {
	m := new(WSGCommonTimeUnitStore)
	return m
}

type WSGCommonTrafficClassProfileRef struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the Traffic Class Profile
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	TrafficClasses []*WSGCommonTrafficClassRef `json:"trafficClasses,omitempty"`

	// ZoneId
	// Zone Id of Traffic Class Profile
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGCommonTrafficClassProfileRef() *WSGCommonTrafficClassProfileRef {
	m := new(WSGCommonTrafficClassProfileRef)
	return m
}

type WSGCommonTrafficClassRef struct {
	// Id
	// Identifier of the Traffic Class
	Id *string `json:"id,omitempty"`

	// Whitelists
	// White list of the Traffic Class Profile. The multiple entries need to be separated by comma (,)
	Whitelists *string `json:"whitelists,omitempty"`
}

func NewWSGCommonTrafficClassRef() *WSGCommonTrafficClassRef {
	m := new(WSGCommonTrafficClassRef)
	return m
}

type WSGCommonTxPower string

func NewWSGCommonTxPower() *WSGCommonTxPower {
	m := new(WSGCommonTxPower)
	return m
}

type WSGCommonWebAuthenticationPortalCustomization struct {
	// Logo
	// Logo encoded with base64, format is "data:image/png;base64,the base64 encoded logo"
	// Constraints:
	//    - required
	Logo *string `json:"logo" validate:"required"`

	// Title
	// Title of the custom portal
	// Constraints:
	//    - max:63
	//    - min:0
	Title *string `json:"title,omitempty" validate:"max=63,min=0"`
}

func NewWSGCommonWebAuthenticationPortalCustomization() *WSGCommonWebAuthenticationPortalCustomization {
	m := new(WSGCommonWebAuthenticationPortalCustomization)
	return m
}

// WSGCommonWildFQDN
//
// Compare with FQDN, it could start with '*.'
type WSGCommonWildFQDN string

func NewWSGCommonWildFQDN() *WSGCommonWildFQDN {
	m := new(WSGCommonWildFQDN)
	return m
}

// WSGCommonZoneTunnelType
//
// Tunnel type configuration of the zone. No_Tunneled is for IPv6 mode
// Constraints:
//    - default:'RuckusGRE'
//    - oneof:[No_Tunneled,RuckusGRE,SoftGRE,SoftGREIpsec]
type WSGCommonZoneTunnelType string

func NewWSGCommonZoneTunnelType() *WSGCommonZoneTunnelType {
	m := new(WSGCommonZoneTunnelType)
	return m
}
