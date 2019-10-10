package common

// API Version: v8_1

import (
	"encoding/json"
)

type Alarm struct {
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

	ID *Mac `json:"id,omitempty"`

	// Severity
	// Alarm severity
	Severity *string `json:"severity,omitempty" validate:"oneof=Critical Major Minor Warning Informational"`

	// Status
	// Alarm status
	Status *string `json:"status,omitempty" validate:"oneof=Outstanding Acknowledged Cleared"`

	// Time
	// Time of the alarm
	Time *string `json:"time,omitempty"`

	// Type
	// Alarm type
	Type *string `json:"type,omitempty"`
}

type Altitude struct {
	// AltitudeUnit
	// altitude unit
	AltitudeUnit *string `json:"altitudeUnit,omitempty" validate:"oneof=meters floor"`

	// AltitudeValue
	// altitude value
	AltitudeValue *int `json:"altitudeValue,omitempty"`
}

// ApGpsSource
//
// GPS Source of the AP
type ApGpsSource string

type ApLatencyInterval struct {
	// PingEnabled
	// AP ping latency enabled
	PingEnabled *bool `json:"pingEnabled,omitempty"`
}

type ApLoginName string

type ApLoginPassword string

type ApManagementVlan struct {
	// ID
	// Vlan id of the zone
	ID *int `json:"id,omitempty"`

	// Mode
	// Vlan Mode of the zone
	Mode *string `json:"mode,omitempty" validate:"oneof=KEEP USER_DEFINED"`
}

type ApRadio50 struct {
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
	ChannelWidth *int `json:"channelWidth,omitempty" validate:"oneof=0 20 40 80 8080 160"`

	// SecondaryChannel
	// channel number (channelWidth is 80+80MHz only)
	SecondaryChannel *int `json:"secondaryChannel,omitempty"`

	TxPower *TxPower `json:"txPower,omitempty"`
}

type ApRebootTimeout struct {
	// GatewayLossTimeoutInSec
	// Gateway loss timeout in second
	GatewayLossTimeoutInSec *int `json:"gatewayLossTimeoutInSec,omitempty" validate:"oneof=0 1800 3600 5400 7200 9000 10800 12600 14400 16200 18000 19800 23400 25200 27000 28800 30600 32400 34200 36000 37800 39600 41400 43200 45000 46800 48600 50400 52200 54000 55800 57600 59400 61200 63000 64800 66600 68400 70200 72000 73800 75600 77400 79200 81000 82800 84600 86400"`

	// ServerLossTimeoutInSec
	// Server loss timeout in second
	ServerLossTimeoutInSec *int `json:"serverLossTimeoutInSec,omitempty" validate:"oneof=0 7200 14400 21600 28800 36000 43200 50400 57600 64800 72000 79200 86400"`
}

type AutoChannelSelection struct {
	// ChannelFlyMtbc
	// ChannelFly MTBC
	ChannelFlyMtbc *int `json:"channelFlyMtbc,omitempty" validate:"gte=100,lte=1440"`

	// ChannelSelectMode
	// Channel Select Mode
	ChannelSelectMode *string `json:"channelSelectMode,omitempty" validate:"oneof=None BackgroundScanning ChannelFly"`
}

type BaseServiceInfo struct {
	// ID
	// ID of service
	ID *string `json:"id,omitempty"`

	// ServiceID
	// ID of service
	ServiceID *string `json:"serviceId,omitempty"`

	// ServiceName
	// Name of service
	ServiceName *string `json:"serviceName,omitempty"`

	// ServiceType
	// Type of service
	ServiceType *string `json:"serviceType,omitempty"`
}

type BulkDeleteRequest struct {
	IDList IDList `json:"idList,omitempty"`
}

type Client struct {
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

	IPAddress *IPAddress `json:"ipAddress,omitempty"`

	Ipv6Address *IPAddress `json:"ipv6Address,omitempty"`

	Mac *Mac `json:"mac,omitempty"`

	// OsType
	// OS type
	OsType *string `json:"osType,omitempty"`

	// RadioID
	// Radio inditifier
	RadioID *string `json:"radioId,omitempty"`

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

	// WLANID
	// WLAN inditifier
	WLANID *string `json:"wlanId,omitempty"`
}

type ClientAdmissionControl struct {
	// MaxRadioLoadPercent
	// Maximum radio load percentage.
	MaxRadioLoadPercent *int `json:"maxRadioLoadPercent,omitempty" validate:"gte=50,lte=100"`

	// MinClientCount
	// Minimum client count number.
	MinClientCount *int `json:"minClientCount,omitempty" validate:"gte=0,lte=100"`

	// MinClientThroughputMbps
	// Minimum client throughput in Mbps.
	MinClientThroughputMbps *float64 `json:"minClientThroughputMbps,omitempty" validate:"gte=0.000000,lte=100.000000"`
}

type CreateResult struct {
	ID *string `json:"id,omitempty"`
}

type CreateResultIDName struct {
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

type Description string

type DescriptionTo128 string

type DHCPProfileRef struct {
	Description *Description `json:"description,omitempty"`

	// ID
	// Identifier of the DHCP Profile
	ID *string `json:"id,omitempty"`

	// LeaseTimeHours
	// Lease time in hours of the DHCP Profile
	LeaseTimeHours *int `json:"leaseTimeHours,omitempty" validate:"gte=0,lte=24"`

	// LeaseTimeMinutes
	// Lease time in minutes of the DHCP Profile
	LeaseTimeMinutes *int `json:"leaseTimeMinutes,omitempty" validate:"gte=0,lte=59"`

	Name *NormalName `json:"name,omitempty"`

	PoolEndIP *IPAddress `json:"poolEndIp,omitempty"`

	PoolStartIP *IPAddress `json:"poolStartIp,omitempty"`

	PrimaryDNSIP *IPAddress `json:"primaryDnsIp,omitempty"`

	SecondaryDNSIP *IPAddress `json:"secondaryDnsIp,omitempty"`

	SubnetMask *IPAddress `json:"subnetMask,omitempty"`

	SubnetNetworkIP *IPAddress `json:"subnetNetworkIp,omitempty"`

	// VlanID
	// VLAN ID of the DHCP Profile
	VlanID *int `json:"vlanId,omitempty" validate:"gte=1,lte=4094"`

	// ZoneID
	// Zone Id of DHCP Profile
	ZoneID *string `json:"zoneId,omitempty"`
}

type DHCPSiteConfigListRef struct {
	// DWPDEnabled
	// DHCP Service Dynamic WAN Port Detection
	DWPDEnabled *bool `json:"dwpdEnabled,omitempty"`

	Eth0ProfileID *int `json:"eth0ProfileId,omitempty"`

	Eth1ProfileID *int `json:"eth1ProfileId,omitempty"`

	// ManualSelect
	// DHCP Service AP Selection Mode
	ManualSelect *bool `json:"manualSelect,omitempty"`

	SiteAps []*DHCPSiteConfigListRefSiteApsType `json:"siteAps,omitempty"`

	// SiteEnabled
	// DHCP Service Enabling Status
	SiteEnabled *bool `json:"siteEnabled,omitempty"`

	// SiteMode
	// DHCP Service mode
	SiteMode *string `json:"siteMode,omitempty" validate:"oneof=EnableOnEachAPs EnableOnMultipleAPs EnableOnHierarchicalAPs"`

	SiteProfiles []*DHCPProfileRef `json:"siteProfiles,omitempty"`

	// ZoneName
	// DHCP Service Zone Name
	ZoneName *string `json:"zoneName,omitempty"`
}

// DHCPSiteConfigListRefSiteApsType
//
// DHCP Site selected APs
type DHCPSiteConfigListRefSiteApsType struct {
	ApGatewayIP *string `json:"apGatewayIp,omitempty"`

	ApMac *string `json:"apMac,omitempty"`

	ApName *string `json:"apName,omitempty"`

	ApServerEnabled *bool `json:"apServerEnabled,omitempty"`

	ApServerIP *string `json:"apServerIp,omitempty"`

	ApServerPrimary *bool `json:"apServerPrimary,omitempty"`

	ApServerType *string `json:"apServerType,omitempty"`

	ApStatus *string `json:"apStatus,omitempty" validate:"oneof=Online Offline Flagged"`
}

type DHCPSiteConfigRef struct {
	// DWPDEnabled
	// DHCP Service Dynamic WAN Port Detection
	DWPDEnabled *bool `json:"dwpdEnabled,omitempty"`

	Eth0ProfileID *int `json:"eth0ProfileId,omitempty"`

	Eth1ProfileID *int `json:"eth1ProfileId,omitempty"`

	// ManualSelect
	// DHCP Service AP Selection Mode. This value is effective when the siteMode is EnableOnMultipleAPs.
	ManualSelect *bool `json:"manualSelect,omitempty"`

	SiteAps []*DHCPSiteConfigRefSiteApsType `json:"siteAps,omitempty"`

	// SiteEnabled
	// DHCP Service Enabling Status
	SiteEnabled *bool `json:"siteEnabled,omitempty"`

	// SiteMode
	// DHCP Service mode
	SiteMode *string `json:"siteMode,omitempty" validate:"oneof=EnableOnEachAPs EnableOnMultipleAPs EnableOnHierarchicalAPs"`

	SiteProfileIds []string `json:"siteProfileIds,omitempty"`
}

// DHCPSiteConfigRefSiteApsType
//
// DHCP Site selected APs. The content is effective when the siteMode is EnableOnMultipleAPs.
type DHCPSiteConfigRefSiteApsType struct {
	ApGatewayIP *string `json:"apGatewayIp,omitempty"`

	ApMac *string `json:"apMac,omitempty"`

	ApName *string `json:"apName,omitempty"`

	ApServerEnabled *bool `json:"apServerEnabled,omitempty"`

	ApServerIP *string `json:"apServerIp,omitempty"`

	ApServerPrimary *bool `json:"apServerPrimary,omitempty"`

	ApServerType *string `json:"apServerType,omitempty"`

	ApStatus *string `json:"apStatus,omitempty" validate:"oneof=Online Offline Flagged"`
}

type DoAssignIP struct {
	// DWPDEnabled
	// DHCP Service Dynamic WAN Port Detection
	DWPDEnabled *bool `json:"dwpdEnabled,omitempty"`

	// ManualSelect
	// DHCP Service AP Selection Mode. This value is effective when the siteMode is EnableOnMultipleAPs.
	ManualSelect *bool `json:"manualSelect,omitempty"`

	SiteAps []*DoAssignIPSiteApsType `json:"siteAps,omitempty"`

	// SiteEnabled
	// DHCP Service Enabling Status
	SiteEnabled *bool `json:"siteEnabled,omitempty"`

	// SiteMode
	// DHCP Service mode
	SiteMode *string `json:"siteMode,omitempty" validate:"oneof=EnableOnEachAPs EnableOnMultipleAPs EnableOnHierarchicalAPs"`

	SiteProfileIds []string `json:"siteProfileIds,omitempty"`
}

// DoAssignIPSiteApsType
//
// DHCP Site selected APs. The content is effective when the siteMode is EnableOnMultipleAPs.
type DoAssignIPSiteApsType struct {
	ApMac *string `json:"apMac,omitempty"`

	ApServerEnabled *bool `json:"apServerEnabled,omitempty"`

	ApServerPrimary *bool `json:"apServerPrimary,omitempty"`
}

type Email string

type EmptyResult struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *EmptyResult) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = EmptyResult{XAdditionalProperties: tmp}
	return nil
}

func (t *EmptyResult) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

type FilterOperator string

type FirmwareVersion string

type FQDN string

type FullTextSearch struct {
	// Fields
	// Specific fields to search
	Fields []string `json:"fields,omitempty"`

	// Type
	// Search logic operator
	Type *string `json:"type,omitempty" validate:"oneof=AND OR"`

	// Value
	// Text or number to search
	Value *string `json:"value,omitempty"`
}

type GenericRef struct {
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

type HealthCheckPolicy struct {
	// ResponseFail
	// Response Fail
	ResponseFail *bool `json:"responseFail,omitempty"`

	// ResponseWindow
	// Response window
	ResponseWindow *int `json:"responseWindow,omitempty" validate:"required,gte=5,lte=30"`

	// ReviveInterval
	// Revive interval
	ReviveInterval *int `json:"reviveInterval,omitempty" validate:"required,gte=60,lte=3600"`

	// ZombiePeriod
	// Zombie period
	ZombiePeriod *int `json:"zombiePeriod,omitempty" validate:"required,gte=30,lte=120"`
}

type HTTPS string

type IDList []string

type IPAddress string

type IPMode string

type LanguageName string

type Latitude float64

type Location string

type LocationAdditionalInfo string

type Longitude float64

type LteBandLockChannel struct {
	// Channel3G
	// LTE 3G channels
	Channel3G *string `json:"channel3g,omitempty"`

	// Channel4G
	// LTE 4G channels
	Channel4G *string `json:"channel4g,omitempty"`

	// SimCardID
	// SIM card ID(Primary:0, Secondary:1)
	SimCardID *int `json:"simCardId,omitempty"`

	// Type
	// LTE chipset SKU type
	Type *string `json:"type,omitempty"`
}

type Mac string

type NormalName string

type NormalName2To64 string

type NormalName2To128 string

type NormalNameAllowBlank string

type NormalURL string

type OverrideClientAdmissionControl struct {
	Enabled *bool `json:"enabled,omitempty"`

	// MaxRadioLoadPercent
	// Maximum radio load percentage.
	MaxRadioLoadPercent *int `json:"maxRadioLoadPercent,omitempty" validate:"gte=50,lte=100"`

	// MinClientCount
	// Minimum client count number.
	MinClientCount *int `json:"minClientCount,omitempty" validate:"gte=0,lte=100"`

	// MinClientThroughputMbps
	// Minimum client throughput in Mbps.
	MinClientThroughputMbps *float64 `json:"minClientThroughputMbps,omitempty" validate:"gte=0.000000,lte=100.000000"`
}

type OverrideGenericRef struct {
	Enabled *bool `json:"enabled,omitempty"`

	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

type OverrideSmartMonitor struct {
	Enabled *bool `json:"enabled,omitempty"`

	// IntervalInSec
	// Interval in seconds. This is required if smartMonitor is enabled
	IntervalInSec *int `json:"intervalInSec,omitempty" validate:"gte=5,lte=60"`

	// RetryThreshold
	// Retry threshold. This is required if smartMonitor is enabled
	RetryThreshold *int `json:"retryThreshold,omitempty" validate:"gte=1,lte=10"`
}

type PortalCustomization struct {
	Language *PortalLanguage `json:"language,omitempty" validate:"required"`

	// Logo
	// logo
	Logo *string `json:"logo,omitempty"`

	// TermsAndConditionsRequired
	// Terms and conditions is required or not
	TermsAndConditionsRequired *bool `json:"termsAndConditionsRequired,omitempty"`

	// TermsAndConditionsText
	// Terms and conditions text
	TermsAndConditionsText *string `json:"termsAndConditionsText,omitempty" validate:"max=3999,min=0"`

	// Title
	// Title
	Title *string `json:"title,omitempty" validate:"max=63,min=0"`
}

// PortalLanguage
//
// Language
type PortalLanguage string

type ProtectionMode string

type QueryCriteria struct {
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
	ExtraFilters []*QueryCriteriaExtraFiltersType `json:"extraFilters,omitempty"`

	// ExtraNotFilters
	// "NOT" condition for multiple filters
	ExtraNotFilters []*QueryCriteriaExtraNotFiltersType `json:"extraNotFilters,omitempty"`

	ExtraTimeRange *TimeRange `json:"extraTimeRange,omitempty"`

	// Filters
	// Filters used to select specific resource scope
	Filters []*QueryCriteriaFiltersType `json:"filters,omitempty"`

	FullTextSearch *FullTextSearch `json:"fullTextSearch,omitempty"`

	// Limit
	// Size of one page
	Limit *int `json:"limit,omitempty" validate:"gte=1"`

	// Options
	// Specified feature required information
	Options *QueryCriteriaOptionsType `json:"options,omitempty"`

	// Page
	// Page number to get
	Page *int `json:"page,omitempty" validate:"gte=1"`

	// Query
	// Add backward compatibility for UI framework
	Query *string `json:"query,omitempty"`

	// SortInfo
	// About sorting
	SortInfo *QueryCriteriaSortInfoType `json:"sortInfo,omitempty"`
}

type QueryCriteriaExtraFiltersType struct {
	Operator *FilterOperator `json:"operator,omitempty"`

	// Type
	// Filters for specific attributes
	Type *string `json:"type,omitempty"`

	// Value
	// value to search
	Value *string `json:"value,omitempty"`
}

type QueryCriteriaExtraNotFiltersType struct {
	// Type
	// Filters for specific attribute
	Type *string `json:"type,omitempty"`

	// Value
	// value not to search
	Value *string `json:"value,omitempty"`
}

type QueryCriteriaFiltersType struct {
	Operator *FilterOperator `json:"operator,omitempty"`

	// Type
	// Group type
	Type *string `json:"type,omitempty"`

	// Value
	// Group ID
	Value *string `json:"value,omitempty"`
}

// QueryCriteriaOptionsType
//
// Specified feature required information
type QueryCriteriaOptionsType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *QueryCriteriaOptionsType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = QueryCriteriaOptionsType{XAdditionalProperties: tmp}
	return nil
}

func (t *QueryCriteriaOptionsType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

// QueryCriteriaSortInfoType
//
// About sorting
type QueryCriteriaSortInfoType struct {
	Dir *string `json:"dir,omitempty" validate:"oneof=ASC DESC"`

	SortColumn *string `json:"sortColumn,omitempty"`
}

type QueryCriteriaSuperSet struct{}

type Radio24 struct {
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
	ChannelWidth *int `json:"channelWidth,omitempty" validate:"oneof=0 20 40"`

	TxPower *TxPower `json:"txPower,omitempty"`
}

type Radio24SuperSet struct {
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
	ChannelWidth *int `json:"channelWidth,omitempty" validate:"oneof=0 20 40"`

	TxPower *TxPower `json:"txPower,omitempty"`
}

type Radio50 struct {
	// AutoCellSizing
	// Auto Cell Sizing
	AutoCellSizing *bool `json:"autoCellSizing,omitempty"`

	// ChannelWidth
	// Channel width. Zero (0) means Auto. 8080 means 80+80MHz
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

	TxPower *TxPower `json:"txPower,omitempty"`
}

type Radio50SuperSet struct {
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

	TxPower *TxPower `json:"txPower,omitempty"`
}

type RadiusServer struct {
	IP *IPAddress `json:"ip,omitempty" validate:"required"`

	// Port
	// Server port
	Port *int `json:"port,omitempty" validate:"required,gte=1,lte=65535"`

	// SharedSecret
	// Server shared secrect
	SharedSecret *string `json:"sharedSecret,omitempty" validate:"required"`
}

type RateLimiting struct {
	// MaxOutstandingRequestsPerServer
	// Maximum outstanding requests (MOR), value should be 0 or between 10 and 4096
	MaxOutstandingRequestsPerServer *int `json:"maxOutstandingRequestsPerServer,omitempty" validate:"required"`

	// SanityTimer
	// Sanity timer
	SanityTimer *int `json:"sanityTimer,omitempty" validate:"required,gte=1,lte=3600"`

	// Threshold
	// Threshold, value should be 0 if MOR is 0, or between 10 and 90 if MOR is between 10 and 4096
	Threshold *int `json:"threshold,omitempty" validate:"required"`
}

type RBACMetadata struct {
	RBACMetadata []string `json:"rbacMetadata,omitempty"`
}

type Realm string

type RecoverySsid struct {
	// RecoverySsidEnabled
	// recovery ssid enable/disable
	RecoverySsidEnabled *bool `json:"recoverySsidEnabled,omitempty"`
}

type SmartMonitor struct {
	// IntervalInSec
	// Interval in seconds. This is required if smartMonitor is enabled
	IntervalInSec *int `json:"intervalInSec,omitempty" validate:"gte=5,lte=60"`

	// RetryThreshold
	// Retry threshold. This is required if smartMonitor is enabled
	RetryThreshold *int `json:"retryThreshold,omitempty" validate:"gte=1,lte=10"`
}

type SNMPCommunity struct {
	// CommunityName
	// name of the SNMP Community.
	CommunityName *string `json:"communityName,omitempty" validate:"required"`

	// NotificationEnabled
	// notification privilege of the SNMP Coummunity
	NotificationEnabled *bool `json:"notificationEnabled,omitempty"`

	// NotificationTarget
	// Trap List of the SNMP Coummunity
	NotificationTarget []*TargetConfig `json:"notificationTarget,omitempty"`

	// NotificationType
	// type of the notification privilege
	NotificationType *string `json:"notificationType,omitempty" validate:"oneof=TRAP INFORM"`

	// ReadEnabled
	// read privilege of the SNMP Coummunity
	ReadEnabled *bool `json:"readEnabled,omitempty"`

	// WriteEnabled
	// write privilege of the SNMP Coummunity
	WriteEnabled *bool `json:"writeEnabled,omitempty"`
}

type SNMPUser struct {
	// AuthPassword
	// authPassword of the SNMP User.
	AuthPassword *string `json:"authPassword,omitempty" validate:"min=8"`

	// AuthProtocol
	// authProtocol of the SNMP User.
	AuthProtocol *string `json:"authProtocol,omitempty" validate:"oneof=MD5 SHA"`

	// NotificationEnabled
	// notification privilege of the SNMP User
	NotificationEnabled *bool `json:"notificationEnabled,omitempty"`

	// NotificationTarget
	// Trap List of the SNMP User
	NotificationTarget []*TargetConfig `json:"notificationTarget,omitempty"`

	// NotificationType
	// type of the notification privilege
	NotificationType *string `json:"notificationType,omitempty" validate:"oneof=TRAP INFORM"`

	// PrivPassword
	// privPassword of the SNMP User.
	PrivPassword *string `json:"privPassword,omitempty" validate:"min=8"`

	// PrivProtocol
	// privProtocol of the SNMP User.
	PrivProtocol *string `json:"privProtocol,omitempty" validate:"oneof=DES AES"`

	// ReadEnabled
	// read privilege of the SNMP User
	ReadEnabled *bool `json:"readEnabled,omitempty"`

	// UserName
	// name of the SNMP User.
	UserName *string `json:"userName,omitempty" validate:"required"`

	// WriteEnabled
	// write privilege of the SNMP User
	WriteEnabled *bool `json:"writeEnabled,omitempty"`
}

type SubNetMask string

type TargetConfig struct {
	// Address
	// address of the SNMP Trap
	Address *string `json:"address,omitempty" validate:"required"`

	// Port
	// port number of the SNMP Trap
	Port *int `json:"port,omitempty" validate:"required,gte=1,lte=65535"`
}

type TimeRange struct {
	// End
	// end time for collecting data
	End *float64 `json:"end,omitempty"`

	// Field
	// time field for collecting data
	Field *string `json:"field,omitempty" validate:"oneof=insertionTime"`

	// Interval
	// time interval in second
	Interval *float64 `json:"interval,omitempty"`

	// Start
	// start time for collecting data
	Start *float64 `json:"start,omitempty"`
}

// TimeUnitStore
//
// time unit
type TimeUnitStore string

type TrafficClassProfileRef struct {
	Description *Description `json:"description,omitempty"`

	// ID
	// Identifier of the Traffic Class Profile
	ID *string `json:"id,omitempty"`

	Name *NormalName `json:"name,omitempty"`

	TrafficClasses []*TrafficClassRef `json:"trafficClasses,omitempty"`

	// ZoneID
	// Zone Id of Traffic Class Profile
	ZoneID *string `json:"zoneId,omitempty"`
}

type TrafficClassRef struct {
	// ID
	// Identifier of the Traffic Class
	ID *string `json:"id,omitempty"`

	// Whitelists
	// White list of the Traffic Class Profile. The multiple entries need to be separated by comma (,)
	Whitelists *string `json:"whitelists,omitempty"`
}

type TxPower string

type WebAuthenticationPortalCustomization struct {
	// Logo
	// Logo encoded with base64, format is "data:image/png;base64,the base64 encoded logo"
	Logo *string `json:"logo,omitempty" validate:"required"`

	// Title
	// Title of the custom portal
	Title *string `json:"title,omitempty" validate:"max=63,min=0"`
}

// WildFQDN
//
// Compare with FQDN, it could start with '*.'
type WildFQDN string

// ZoneTunnelType
//
// Tunnel type configuration of the zone. No_Tunneled is for IPv6 mode
type ZoneTunnelType string
