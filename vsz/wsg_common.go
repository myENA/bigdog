package vsz

// API Version: v8_1

import (
	"encoding/json"
)

type WSGCommonAlarm struct {
	// AcknowledgedTime
	// Time the alarm was acknowledged
	AcknowledgedTime *int `json:"acknowledgedTime,omitempty"`

	// Category
	// Alarm category
	Category *string `json:"category,omitempty"`

	// ClearedTime
	// Time that alarm was cleared
	ClearedTime *int `json:"clearedTime,omitempty"`

	// Code
	// Alarm code
	Code *string `json:"code,omitempty"`

	// Description
	// Alarm description
	Description *string `json:"description,omitempty"`

	Id *WSGCommonMac `json:"id,omitempty"`

	// Severity
	// Alarm severity
	// Constraints:
	//    - nullable
	//    - oneof:[Critical,Major,Minor,Warning,Informational]
	Severity *string `json:"severity,omitempty" validate:"omitempty,oneof=Critical Major Minor Warning Informational"`

	// Status
	// Alarm status
	// Constraints:
	//    - nullable
	//    - oneof:[Outstanding,Acknowledged,Cleared]
	Status *string `json:"status,omitempty" validate:"omitempty,oneof=Outstanding Acknowledged Cleared"`

	// Time
	// Time of the alarm
	Time *string `json:"time,omitempty"`

	// Type
	// Alarm type
	Type *string `json:"type,omitempty"`
}

type WSGCommonAltitude struct {
	// AltitudeUnit
	// altitude unit
	// Constraints:
	//    - nullable
	//    - default:'meters'
	//    - oneof:[meters,floor]
	AltitudeUnit *string `json:"altitudeUnit,omitempty" validate:"omitempty,oneof=meters floor"`

	// AltitudeValue
	// altitude value
	AltitudeValue *int `json:"altitudeValue,omitempty"`
}

// WSGCommonApGpsSource
//
// GPS Source of the AP
// Constraints:
//    - nullable
//    - oneof:[GPS,MANUAL]
type WSGCommonApGpsSource string

type WSGCommonApLatencyInterval struct {
	// PingEnabled
	// AP ping latency enabled
	PingEnabled *bool `json:"pingEnabled,omitempty"`
}

type WSGCommonApLoginName string

type WSGCommonApLoginPassword string

type WSGCommonApManagementVlan struct {
	// Id
	// Vlan id of the zone
	Id *int `json:"id,omitempty"`

	// Mode
	// Vlan Mode of the zone
	// Constraints:
	//    - nullable
	//    - default:'KEEP'
	//    - oneof:[KEEP,USER_DEFINED]
	Mode *string `json:"mode,omitempty" validate:"omitempty,oneof=KEEP USER_DEFINED"`
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
	//    - nullable
	//    - oneof:[0,20,40,80,8080,160]
	ChannelWidth *int `json:"channelWidth,omitempty" validate:"omitempty,oneof=0 20 40 80 8080 160"`

	// SecondaryChannel
	// channel number (channelWidth is 80+80MHz only)
	SecondaryChannel *int `json:"secondaryChannel,omitempty"`

	TxPower *WSGCommonTxPower `json:"txPower,omitempty"`
}

type WSGCommonApRebootTimeout struct {
	// GatewayLossTimeoutInSec
	// Gateway loss timeout in second
	// Constraints:
	//    - nullable
	//    - default:1800
	//    - oneof:[0,1800,3600,5400,7200,9000,10800,12600,14400,16200,18000,19800,23400,25200,27000,28800,30600,32400,34200,36000,37800,39600,41400,43200,45000,46800,48600,50400,52200,54000,55800,57600,59400,61200,63000,64800,66600,68400,70200,72000,73800,75600,77400,79200,81000,82800,84600,86400]
	GatewayLossTimeoutInSec *int `json:"gatewayLossTimeoutInSec,omitempty" validate:"omitempty,oneof=0 1800 3600 5400 7200 9000 10800 12600 14400 16200 18000 19800 23400 25200 27000 28800 30600 32400 34200 36000 37800 39600 41400 43200 45000 46800 48600 50400 52200 54000 55800 57600 59400 61200 63000 64800 66600 68400 70200 72000 73800 75600 77400 79200 81000 82800 84600 86400"`

	// ServerLossTimeoutInSec
	// Server loss timeout in second
	// Constraints:
	//    - nullable
	//    - default:7200
	//    - oneof:[0,7200,14400,21600,28800,36000,43200,50400,57600,64800,72000,79200,86400]
	ServerLossTimeoutInSec *int `json:"serverLossTimeoutInSec,omitempty" validate:"omitempty,oneof=0 7200 14400 21600 28800 36000 43200 50400 57600 64800 72000 79200 86400"`
}

type WSGCommonAutoChannelSelection struct {
	// ChannelFlyMtbc
	// ChannelFly MTBC
	// Constraints:
	//    - nullable
	//    - default:480
	//    - min:100
	//    - max:1440
	ChannelFlyMtbc *int `json:"channelFlyMtbc,omitempty" validate:"omitempty,gte=100,lte=1440"`

	// ChannelSelectMode
	// Channel Select Mode
	// Constraints:
	//    - nullable
	//    - default:'BackgroundScanning'
	//    - oneof:[None,BackgroundScanning,ChannelFly]
	ChannelSelectMode *string `json:"channelSelectMode,omitempty" validate:"omitempty,oneof=None BackgroundScanning ChannelFly"`
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

type WSGCommonBulkDeleteRequest struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

type WSGCommonClient struct {
	// ApTxDataRate
	// AP Tx Data Rate
	ApTxDataRate *string `json:"apTxDataRate,omitempty"`

	// Channel
	// Channel
	Channel *string `json:"channel,omitempty"`

	// ConnectedSince
	// Connected since (in milliseconds)
	ConnectedSince *int `json:"connectedSince,omitempty"`

	// FromClientBytes
	// From client bytes
	FromClientBytes *int `json:"fromClientBytes,omitempty"`

	// FromClientPkts
	// From client package frames
	FromClientPkts *int `json:"fromClientPkts,omitempty"`

	// HostName
	// Host name
	HostName *string `json:"hostName,omitempty"`

	IpAddress *WSGCommonIpAddress `json:"ipAddress,omitempty"`

	Ipv6Address *WSGCommonIpAddress `json:"ipv6Address,omitempty"`

	Mac *WSGCommonMac `json:"mac,omitempty"`

	// OsType
	// OS type
	OsType *string `json:"osType,omitempty"`

	// RadioId
	// Radio inditifier
	RadioId *string `json:"radioId,omitempty"`

	// RadioMode
	// Radio mode
	RadioMode *string `json:"radioMode,omitempty"`

	// Rssi
	// RSSI(dBm)
	Rssi *string `json:"rssi,omitempty"`

	// RxAvgByteRate
	// RX Avg Byte Rate
	RxAvgByteRate *int `json:"rxAvgByteRate,omitempty"`

	// RxByteRate
	// RX Byte Rate
	RxByteRate *int `json:"rxByteRate,omitempty"`

	// Snr
	// SNR(dB)
	Snr *string `json:"snr,omitempty"`

	// Ssid
	// SSID
	Ssid *string `json:"ssid,omitempty"`

	// Status
	// Status
	Status *string `json:"status,omitempty"`

	// ToClientBytes
	// To client bytes
	ToClientBytes *int `json:"toClientBytes,omitempty"`

	// ToClientDroppedPkts
	// To client dropped packages
	ToClientDroppedPkts *int `json:"toClientDroppedPkts,omitempty"`

	// ToClientPkts
	// To client package frames
	ToClientPkts *int `json:"toClientPkts,omitempty"`

	// TxAvgByteRate
	// TX Avg Byte Rate
	TxAvgByteRate *int `json:"txAvgByteRate,omitempty"`

	// TxByteRate
	// TX Byte Rate
	TxByteRate *int `json:"txByteRate,omitempty"`

	// User
	// User
	User *string `json:"user,omitempty"`

	// Vlan
	// VLAN id
	Vlan *string `json:"vlan,omitempty"`

	// WlanId
	// WLAN inditifier
	WlanId *string `json:"wlanId,omitempty"`
}

type WSGCommonClientAdmissionControl struct {
	// MaxRadioLoadPercent
	// Maximum radio load percentage.
	// Constraints:
	//    - nullable
	//    - default:75
	//    - min:50
	//    - max:100
	MaxRadioLoadPercent *int `json:"maxRadioLoadPercent,omitempty" validate:"omitempty,gte=50,lte=100"`

	// MinClientCount
	// Minimum client count number.
	// Constraints:
	//    - nullable
	//    - default:10
	//    - min:0
	//    - max:100
	MinClientCount *int `json:"minClientCount,omitempty" validate:"omitempty,gte=0,lte=100"`

	// MinClientThroughputMbps
	// Minimum client throughput in Mbps.
	// Constraints:
	//    - nullable
	//    - default:0
	//    - min:0.000000
	//    - max:100.000000
	MinClientThroughputMbps *float64 `json:"minClientThroughputMbps,omitempty" validate:"omitempty,gte=0.000000,lte=100.000000"`
}

type WSGCommonCreateResult struct {
	Id *string `json:"id,omitempty"`
}

type WSGCommonCreateResultIdName struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

type WSGCommonDescription string

type WSGCommonDescriptionTo128 string

type WSGCommonDhcpProfileRef struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the DHCP Profile
	Id *string `json:"id,omitempty"`

	// LeaseTimeHours
	// Lease time in hours of the DHCP Profile
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:24
	LeaseTimeHours *int `json:"leaseTimeHours,omitempty" validate:"omitempty,gte=0,lte=24"`

	// LeaseTimeMinutes
	// Lease time in minutes of the DHCP Profile
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:59
	LeaseTimeMinutes *int `json:"leaseTimeMinutes,omitempty" validate:"omitempty,gte=0,lte=59"`

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
	//    - nullable
	//    - min:1
	//    - max:4094
	VlanId *int `json:"vlanId,omitempty" validate:"omitempty,gte=1,lte=4094"`

	// ZoneId
	// Zone Id of DHCP Profile
	ZoneId *string `json:"zoneId,omitempty"`
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
	//    - nullable
	//    - oneof:[EnableOnEachAPs,EnableOnMultipleAPs,EnableOnHierarchicalAPs]
	SiteMode *string `json:"siteMode,omitempty" validate:"omitempty,oneof=EnableOnEachAPs EnableOnMultipleAPs EnableOnHierarchicalAPs"`

	SiteProfiles []*WSGCommonDhcpProfileRef `json:"siteProfiles,omitempty"`

	// ZoneName
	// DHCP Service Zone Name
	ZoneName *string `json:"zoneName,omitempty"`
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
	//    - nullable
	//    - oneof:[Online,Offline,Flagged]
	ApStatus *string `json:"apStatus,omitempty" validate:"omitempty,oneof=Online Offline Flagged"`
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
	//    - nullable
	//    - oneof:[EnableOnEachAPs,EnableOnMultipleAPs,EnableOnHierarchicalAPs]
	SiteMode *string `json:"siteMode,omitempty" validate:"omitempty,oneof=EnableOnEachAPs EnableOnMultipleAPs EnableOnHierarchicalAPs"`

	SiteProfileIds []string `json:"siteProfileIds,omitempty"`
}

// WSGCommonDhcpSiteConfigRefSiteApsType
//
// DHCP Site selected APs. The content is effective when the siteMode is EnableOnMultipleAPs.
type WSGCommonDhcpSiteConfigRefSiteApsType struct {
	ApGatewayIp *string `json:"apGatewayIp,omitempty"`

	ApMac *string `json:"apMac,omitempty"`

	ApName *string `json:"apName,omitempty"`

	ApServerEnabled *bool `json:"apServerEnabled,omitempty"`

	ApServerIp *string `json:"apServerIp,omitempty"`

	ApServerPrimary *bool `json:"apServerPrimary,omitempty"`

	ApServerType *string `json:"apServerType,omitempty"`

	// ApStatus
	// Constraints:
	//    - nullable
	//    - oneof:[Online,Offline,Flagged]
	ApStatus *string `json:"apStatus,omitempty" validate:"omitempty,oneof=Online Offline Flagged"`
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
	//    - nullable
	//    - oneof:[EnableOnEachAPs,EnableOnMultipleAPs,EnableOnHierarchicalAPs]
	SiteMode *string `json:"siteMode,omitempty" validate:"omitempty,oneof=EnableOnEachAPs EnableOnMultipleAPs EnableOnHierarchicalAPs"`

	SiteProfileIds []string `json:"siteProfileIds,omitempty"`
}

// WSGCommonDoAssignIpSiteApsType
//
// DHCP Site selected APs. The content is effective when the siteMode is EnableOnMultipleAPs.
type WSGCommonDoAssignIpSiteApsType struct {
	ApMac *string `json:"apMac,omitempty"`

	ApServerEnabled *bool `json:"apServerEnabled,omitempty"`

	ApServerPrimary *bool `json:"apServerPrimary,omitempty"`
}

type WSGCommonEmail string

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

type WSGCommonFilterOperator string

type WSGCommonFirmwareVersion string

type WSGCommonFQDN string

type WSGCommonFullTextSearch struct {
	// Fields
	// Specific fields to search
	Fields []string `json:"fields,omitempty"`

	// Type
	// Search logic operator
	// Constraints:
	//    - nullable
	//    - oneof:[AND,OR]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=AND OR"`

	// Value
	// Text or number to search
	Value *string `json:"value,omitempty"`
}

type WSGCommonGenericRef struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
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

type WSGCommonHTTPS string

type WSGCommonIdList []string

type WSGCommonIpAddress string

type WSGCommonIpMode string

type WSGCommonLanguageName string

type WSGCommonLatitude float64

type WSGCommonLocation string

type WSGCommonLocationAdditionalInfo string

type WSGCommonLongitude float64

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

type WSGCommonMac string

type WSGCommonNormalName string

type WSGCommonNormalName2to64 string

type WSGCommonNormalName2to128 string

type WSGCommonNormalNameAllowBlank string

type WSGCommonNormalURL string

type WSGCommonOverrideClientAdmissionControl struct {
	Enabled *bool `json:"enabled,omitempty"`

	// MaxRadioLoadPercent
	// Maximum radio load percentage.
	// Constraints:
	//    - nullable
	//    - default:75
	//    - min:50
	//    - max:100
	MaxRadioLoadPercent *int `json:"maxRadioLoadPercent,omitempty" validate:"omitempty,gte=50,lte=100"`

	// MinClientCount
	// Minimum client count number.
	// Constraints:
	//    - nullable
	//    - default:10
	//    - min:0
	//    - max:100
	MinClientCount *int `json:"minClientCount,omitempty" validate:"omitempty,gte=0,lte=100"`

	// MinClientThroughputMbps
	// Minimum client throughput in Mbps.
	// Constraints:
	//    - nullable
	//    - default:0
	//    - min:0.000000
	//    - max:100.000000
	MinClientThroughputMbps *float64 `json:"minClientThroughputMbps,omitempty" validate:"omitempty,gte=0.000000,lte=100.000000"`
}

type WSGCommonOverrideGenericRef struct {
	Enabled *bool `json:"enabled,omitempty"`

	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

type WSGCommonOverrideSmartMonitor struct {
	Enabled *bool `json:"enabled,omitempty"`

	// IntervalInSec
	// Interval in seconds. This is required if smartMonitor is enabled
	// Constraints:
	//    - nullable
	//    - default:10
	//    - min:5
	//    - max:60
	IntervalInSec *int `json:"intervalInSec,omitempty" validate:"omitempty,gte=5,lte=60"`

	// RetryThreshold
	// Retry threshold. This is required if smartMonitor is enabled
	// Constraints:
	//    - nullable
	//    - default:3
	//    - min:1
	//    - max:10
	RetryThreshold *int `json:"retryThreshold,omitempty" validate:"omitempty,gte=1,lte=10"`
}

type WSGCommonPortalCustomization struct {
	// Language
	// Constraints:
	//    - required
	Language *WSGCommonPortalLanguage `json:"language" validate:"required"`

	// Logo
	// logo
	Logo *string `json:"logo,omitempty"`

	// TermsAndConditionsRequired
	// Terms and conditions is required or not
	TermsAndConditionsRequired *bool `json:"termsAndConditionsRequired,omitempty"`

	// TermsAndConditionsText
	// Terms and conditions text
	// Constraints:
	//    - nullable
	//    - default:'Terms of Use
	//
	// By accepting this agreement and accessing the wireless network, you acknowledge that you are of legal age, you have read and understood, and agree to be bound by this agreement.
	// (*) The wireless network service is provided by the property owners and is completely at their discretion. Your access to the network may be blocked, suspended, or terminated at any time for any reason.
	// (*) You agree not to use the wireless network for any purpose that is unlawful or otherwise prohibited and you are fully responsible for your use.
	// (*) The wireless network is provided "as is" without warranties of any kind, either expressed or implied.
	//
	// This wireless network is powered by Ruckus Wireless.'
	//    - max:3999
	//    - min:0
	TermsAndConditionsText *string `json:"termsAndConditionsText,omitempty" validate:"omitempty,max=3999,min=0"`

	// Title
	// Title
	// Constraints:
	//    - nullable
	//    - max:63
	//    - min:0
	Title *string `json:"title,omitempty" validate:"omitempty,max=63,min=0"`
}

// WSGCommonPortalLanguage
//
// Language
// Constraints:
//    - nullable
//    - default:'en_US'
//    - oneof:[en_US,zh_TW,zh_CN,nl_NL,fr_FR,de_DE,ja_JP,es_ES,se_SE,ar_SA,cz_CZ,da_DK,tr_TR,pt_BR]
type WSGCommonPortalLanguage string

type WSGCommonProtectionMode string

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
	//    - nullable
	//    - min:1
	Limit *int `json:"limit,omitempty" validate:"omitempty,gte=1"`

	// Options
	// Specified feature required information
	Options *WSGCommonQueryCriteriaOptionsType `json:"options,omitempty"`

	// Page
	// Page number to get
	// Constraints:
	//    - nullable
	//    - min:1
	Page *int `json:"page,omitempty" validate:"omitempty,gte=1"`

	// Query
	// Add backward compatibility for UI framework
	Query *string `json:"query,omitempty"`

	// SortInfo
	// About sorting
	SortInfo *WSGCommonQueryCriteriaSortInfoType `json:"sortInfo,omitempty"`
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

type WSGCommonQueryCriteriaExtraNotFiltersType struct {
	// Type
	// Filters for specific attribute
	Type *string `json:"type,omitempty"`

	// Value
	// value not to search
	Value *string `json:"value,omitempty"`
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

// WSGCommonQueryCriteriaSortInfoType
//
// About sorting
type WSGCommonQueryCriteriaSortInfoType struct {
	// Dir
	// Constraints:
	//    - nullable
	//    - oneof:[ASC,DESC]
	Dir *string `json:"dir,omitempty" validate:"omitempty,oneof=ASC DESC"`

	SortColumn *string `json:"sortColumn,omitempty"`
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
	//    - nullable
	//    - min:1
	Limit *int `json:"limit,omitempty" validate:"omitempty,gte=1"`

	// Options
	// Specified feature required information
	Options *WSGCommonQueryCriteriaOptionsType `json:"options,omitempty"`

	// Page
	// Page number to get
	// Constraints:
	//    - nullable
	//    - min:1
	Page *int `json:"page,omitempty" validate:"omitempty,gte=1"`

	// Query
	// Add backward compatibility for UI framework
	Query *string `json:"query,omitempty"`

	// SortInfo
	// About sorting
	SortInfo *WSGCommonQueryCriteriaSortInfoType `json:"sortInfo,omitempty"`
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
	//    - nullable
	//    - default:0
	//    - oneof:[0,20,40]
	ChannelWidth *int `json:"channelWidth,omitempty" validate:"omitempty,oneof=0 20 40"`

	TxPower *WSGCommonTxPower `json:"txPower,omitempty"`
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
	//    - nullable
	//    - oneof:[0,20,40]
	ChannelWidth *int `json:"channelWidth,omitempty" validate:"omitempty,oneof=0 20 40"`

	TxPower *WSGCommonTxPower `json:"txPower,omitempty"`
}

type WSGCommonRadio50 struct {
	// AutoCellSizing
	// Auto Cell Sizing
	AutoCellSizing *bool `json:"autoCellSizing,omitempty"`

	// ChannelWidth
	// Channel width. Zero (0) means Auto. 8080 means 80+80MHz
	// Constraints:
	//    - nullable
	//    - default:0
	//    - oneof:[0,20,40,80,8080,160]
	ChannelWidth *int `json:"channelWidth,omitempty" validate:"omitempty,oneof=0 20 40 80 8080 160"`

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
	//    - nullable
	//    - oneof:[0,20,40,80,8080,160]
	ChannelWidth *int `json:"channelWidth,omitempty" validate:"omitempty,oneof=0 20 40 80 8080 160"`

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

type WSGCommonRbacMetadata struct {
	RbacMetadata []string `json:"rbacMetadata,omitempty"`
}

type WSGCommonRealm string

type WSGCommonRecoverySsid struct {
	// RecoverySsidEnabled
	// recovery ssid enable/disable
	RecoverySsidEnabled *bool `json:"recoverySsidEnabled,omitempty"`
}

type WSGCommonSmartMonitor struct {
	// IntervalInSec
	// Interval in seconds. This is required if smartMonitor is enabled
	// Constraints:
	//    - nullable
	//    - default:10
	//    - min:5
	//    - max:60
	IntervalInSec *int `json:"intervalInSec,omitempty" validate:"omitempty,gte=5,lte=60"`

	// RetryThreshold
	// Retry threshold. This is required if smartMonitor is enabled
	// Constraints:
	//    - nullable
	//    - default:3
	//    - min:1
	//    - max:10
	RetryThreshold *int `json:"retryThreshold,omitempty" validate:"omitempty,gte=1,lte=10"`
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
	//    - nullable
	//    - oneof:[TRAP,INFORM]
	NotificationType *string `json:"notificationType,omitempty" validate:"omitempty,oneof=TRAP INFORM"`

	// ReadEnabled
	// read privilege of the SNMP Coummunity
	ReadEnabled *bool `json:"readEnabled,omitempty"`

	// WriteEnabled
	// write privilege of the SNMP Coummunity
	WriteEnabled *bool `json:"writeEnabled,omitempty"`
}

type WSGCommonSnmpUser struct {
	// AuthPassword
	// authPassword of the SNMP User.
	// Constraints:
	//    - nullable
	//    - min:8
	AuthPassword *string `json:"authPassword,omitempty" validate:"omitempty,min=8"`

	// AuthProtocol
	// authProtocol of the SNMP User.
	// Constraints:
	//    - nullable
	//    - oneof:[MD5,SHA]
	AuthProtocol *string `json:"authProtocol,omitempty" validate:"omitempty,oneof=MD5 SHA"`

	// NotificationEnabled
	// notification privilege of the SNMP User
	NotificationEnabled *bool `json:"notificationEnabled,omitempty"`

	// NotificationTarget
	// Trap List of the SNMP User
	NotificationTarget []*WSGCommonTargetConfig `json:"notificationTarget,omitempty"`

	// NotificationType
	// type of the notification privilege
	// Constraints:
	//    - nullable
	//    - oneof:[TRAP,INFORM]
	NotificationType *string `json:"notificationType,omitempty" validate:"omitempty,oneof=TRAP INFORM"`

	// PrivPassword
	// privPassword of the SNMP User.
	// Constraints:
	//    - nullable
	//    - min:8
	PrivPassword *string `json:"privPassword,omitempty" validate:"omitempty,min=8"`

	// PrivProtocol
	// privProtocol of the SNMP User.
	// Constraints:
	//    - nullable
	//    - oneof:[DES,AES]
	PrivProtocol *string `json:"privProtocol,omitempty" validate:"omitempty,oneof=DES AES"`

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

type WSGCommonSubNetMask string

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

type WSGCommonTimeRange struct {
	// End
	// end time for collecting data
	End *float64 `json:"end,omitempty"`

	// Field
	// time field for collecting data
	// Constraints:
	//    - nullable
	//    - oneof:[insertionTime]
	Field *string `json:"field,omitempty" validate:"omitempty,oneof=insertionTime"`

	// Interval
	// time interval in second
	Interval *float64 `json:"interval,omitempty"`

	// Start
	// start time for collecting data
	Start *float64 `json:"start,omitempty"`
}

// WSGCommonTimeUnitStore
//
// time unit
// Constraints:
//    - nullable
//    - oneof:[second,minute,hour,day]
type WSGCommonTimeUnitStore string

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

type WSGCommonTrafficClassRef struct {
	// Id
	// Identifier of the Traffic Class
	Id *string `json:"id,omitempty"`

	// Whitelists
	// White list of the Traffic Class Profile. The multiple entries need to be separated by comma (,)
	Whitelists *string `json:"whitelists,omitempty"`
}

type WSGCommonTxPower string

type WSGCommonWebAuthenticationPortalCustomization struct {
	// Logo
	// Logo encoded with base64, format is "data:image/png;base64,the base64 encoded logo"
	// Constraints:
	//    - required
	Logo *string `json:"logo" validate:"required"`

	// Title
	// Title of the custom portal
	// Constraints:
	//    - nullable
	//    - max:63
	//    - min:0
	Title *string `json:"title,omitempty" validate:"omitempty,max=63,min=0"`
}

// WSGCommonWildFQDN
//
// Compare with FQDN, it could start with '*.'
type WSGCommonWildFQDN string

// WSGCommonZoneTunnelType
//
// Tunnel type configuration of the zone. No_Tunneled is for IPv6 mode
// Constraints:
//    - nullable
//    - default:'RuckusGRE'
//    - oneof:[No_Tunneled,RuckusGRE,SoftGRE,SoftGREIpsec]
type WSGCommonZoneTunnelType string
