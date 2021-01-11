package bigdog

// API Version: v9_1

import (
	"errors"
	"io"
)

// WSGCommonAltitude
//
// Definition: common_altitude
type WSGCommonAltitude struct {
	// AltitudeUnit
	// altitude unit
	// Constraints:
	//    - default:'meters'
	//    - oneof:[meters,floor]
	AltitudeUnit *string `json:"altitudeUnit,omitempty"`

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
// Definition: common_apGpsSource
//
// GPS Source of the AP
// Constraints:
//    - oneof:[GPS,MANUAL]
type WSGCommonApGpsSource string

func NewWSGCommonApGpsSource() *WSGCommonApGpsSource {
	m := new(WSGCommonApGpsSource)
	return m
}

// WSGCommonApLatencyInterval
//
// Definition: common_apLatencyInterval
type WSGCommonApLatencyInterval struct {
	// PingEnabled
	// AP ping latency enabled
	PingEnabled *bool `json:"pingEnabled,omitempty"`
}

func NewWSGCommonApLatencyInterval() *WSGCommonApLatencyInterval {
	m := new(WSGCommonApLatencyInterval)
	return m
}

// WSGCommonApLoginName
//
// Definition: common_apLoginName
//
// Constraints:
//    - max:64
type WSGCommonApLoginName string

func NewWSGCommonApLoginName() *WSGCommonApLoginName {
	m := new(WSGCommonApLoginName)
	return m
}

// WSGCommonApLoginPassword
//
// Definition: common_apLoginPassword
//
// Constraints:
//    - max:64
type WSGCommonApLoginPassword string

func NewWSGCommonApLoginPassword() *WSGCommonApLoginPassword {
	m := new(WSGCommonApLoginPassword)
	return m
}

// WSGCommonApManagementVlan
//
// Definition: common_apManagementVlan
type WSGCommonApManagementVlan struct {
	// Id
	// Vlan id of the zone
	Id *int `json:"id,omitempty"`

	// Mode
	// Vlan Mode of the zone
	// Constraints:
	//    - default:'KEEP'
	//    - oneof:[KEEP,USER_DEFINED]
	Mode *string `json:"mode,omitempty"`
}

func NewWSGCommonApManagementVlan() *WSGCommonApManagementVlan {
	m := new(WSGCommonApManagementVlan)
	return m
}

// WSGCommonApRadio50
//
// Definition: common_apRadio50
type WSGCommonApRadio50 struct {
	// AutoCellSizing
	// Auto Cell Sizing.(When Auto Cell Sizing is enabled, The TX Power settings will invalid.)
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
	ChannelWidth *int `json:"channelWidth,omitempty"`

	// SecondaryChannel
	// channel number (channelWidth is 80+80MHz only)
	SecondaryChannel *int `json:"secondaryChannel,omitempty"`

	TxPower *WSGCommonTxPower `json:"txPower,omitempty"`
}

func NewWSGCommonApRadio50() *WSGCommonApRadio50 {
	m := new(WSGCommonApRadio50)
	return m
}

// WSGCommonApRebootTimeout
//
// Definition: common_apRebootTimeout
type WSGCommonApRebootTimeout struct {
	// GatewayLossTimeoutInSec
	// Gateway loss timeout in second
	// Constraints:
	//    - default:1800
	//    - oneof:[0,1800,3600,5400,7200,9000,10800,12600,14400,16200,18000,19800,23400,25200,27000,28800,30600,32400,34200,36000,37800,39600,41400,43200,45000,46800,48600,50400,52200,54000,55800,57600,59400,61200,63000,64800,66600,68400,70200,72000,73800,75600,77400,79200,81000,82800,84600,86400]
	GatewayLossTimeoutInSec *int `json:"gatewayLossTimeoutInSec,omitempty"`

	// ServerLossTimeoutInSec
	// Server loss timeout in second
	// Constraints:
	//    - default:7200
	//    - oneof:[0,7200,14400,21600,28800,36000,43200,50400,57600,64800,72000,79200,86400]
	ServerLossTimeoutInSec *int `json:"serverLossTimeoutInSec,omitempty"`
}

func NewWSGCommonApRebootTimeout() *WSGCommonApRebootTimeout {
	m := new(WSGCommonApRebootTimeout)
	return m
}

// WSGCommonAutoChannelSelection
//
// Definition: common_autoChannelSelection
type WSGCommonAutoChannelSelection struct {
	// ChannelFlyChangeFrequency
	// ChannelFly Change Frequency
	// Constraints:
	//    - default:33
	//    - min:1
	//    - max:100
	ChannelFlyChangeFrequency *int `json:"channelFlyChangeFrequency,omitempty"`

	// ChannelFlyMtbc
	// ChannelFly MTBC
	// Constraints:
	//    - default:480
	//    - min:100
	//    - max:1440
	ChannelFlyMtbc *int `json:"channelFlyMtbc,omitempty"`

	// ChannelFlyOptimizationTimePeriod
	// ChannelFly Optimization Time Period
	ChannelFlyOptimizationTimePeriod []string `json:"channelFlyOptimizationTimePeriod,omitempty"`

	// ChannelSelectMode
	// Channel Select Mode
	// Constraints:
	//    - default:'BackgroundScanning'
	//    - oneof:[None,BackgroundScanning,ChannelFly]
	ChannelSelectMode *string `json:"channelSelectMode,omitempty"`
}

func NewWSGCommonAutoChannelSelection() *WSGCommonAutoChannelSelection {
	m := new(WSGCommonAutoChannelSelection)
	return m
}

// WSGCommonAwsVenue
//
// Definition: common_awsVenue
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

// WSGCommonBaseServiceInfo
//
// Definition: common_baseServiceInfo
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

// WSGCommonBulkDeleteRequest
//
// Definition: common_bulkDeleteRequest
type WSGCommonBulkDeleteRequest struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGCommonBulkDeleteRequest() *WSGCommonBulkDeleteRequest {
	m := new(WSGCommonBulkDeleteRequest)
	return m
}

// WSGCommonClientAdmissionControl
//
// Definition: common_clientAdmissionControl
type WSGCommonClientAdmissionControl struct {
	// MaxRadioLoadPercent
	// Maximum radio load percentage.
	// Constraints:
	//    - default:75
	//    - min:50
	//    - max:100
	MaxRadioLoadPercent *int `json:"maxRadioLoadPercent,omitempty"`

	// MinClientCount
	// Minimum client count number.
	// Constraints:
	//    - default:10
	//    - min:0
	//    - max:100
	MinClientCount *int `json:"minClientCount,omitempty"`

	// MinClientThroughputMbps
	// Minimum client throughput in Mbps.
	// Constraints:
	//    - default:0
	//    - min:0.000000
	//    - max:100.000000
	MinClientThroughputMbps *float64 `json:"minClientThroughputMbps,omitempty"`
}

func NewWSGCommonClientAdmissionControl() *WSGCommonClientAdmissionControl {
	m := new(WSGCommonClientAdmissionControl)
	return m
}

// WSGCommonCreateResult
//
// Definition: common_createResult
type WSGCommonCreateResult struct {
	Id *string `json:"id,omitempty"`
}

type WSGCommonCreateResultAPIResponse struct {
	*RawAPIResponse
	Data *WSGCommonCreateResult
}

func newWSGCommonCreateResultAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGCommonCreateResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGCommonCreateResultAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGCommonCreateResult)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGCommonCreateResult() *WSGCommonCreateResult {
	m := new(WSGCommonCreateResult)
	return m
}

// WSGCommonCreateResultIdName
//
// Definition: common_createResultIdName
type WSGCommonCreateResultIdName struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

type WSGCommonCreateResultIdNameAPIResponse struct {
	*RawAPIResponse
	Data *WSGCommonCreateResultIdName
}

func newWSGCommonCreateResultIdNameAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGCommonCreateResultIdNameAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGCommonCreateResultIdNameAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGCommonCreateResultIdName)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGCommonCreateResultIdName() *WSGCommonCreateResultIdName {
	m := new(WSGCommonCreateResultIdName)
	return m
}

// WSGCommonDescription
//
// Definition: common_description
//
// Constraints:
//    - nullable
//    - max:64
type WSGCommonDescription string

func NewWSGCommonDescription() *WSGCommonDescription {
	m := new(WSGCommonDescription)
	return m
}

// WSGCommonDescriptionTo128
//
// Definition: common_descriptionto128
//
// Constraints:
//    - nullable
//    - max:128
type WSGCommonDescriptionTo128 string

func NewWSGCommonDescriptionTo128() *WSGCommonDescriptionTo128 {
	m := new(WSGCommonDescriptionTo128)
	return m
}

// WSGCommonDhcpProfileRef
//
// Definition: common_dhcpProfileRef
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
	LeaseTimeHours *int `json:"leaseTimeHours,omitempty"`

	// LeaseTimeMinutes
	// Lease time in minutes of the DHCP Profile
	// Constraints:
	//    - min:0
	//    - max:59
	LeaseTimeMinutes *int `json:"leaseTimeMinutes,omitempty"`

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
	VlanId *int `json:"vlanId,omitempty"`

	// ZoneId
	// Zone Id of DHCP Profile
	ZoneId *string `json:"zoneId,omitempty"`
}

type WSGCommonDhcpProfileRefAPIResponse struct {
	*RawAPIResponse
	Data *WSGCommonDhcpProfileRef
}

func newWSGCommonDhcpProfileRefAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGCommonDhcpProfileRefAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGCommonDhcpProfileRefAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGCommonDhcpProfileRef)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGCommonDhcpProfileRef() *WSGCommonDhcpProfileRef {
	m := new(WSGCommonDhcpProfileRef)
	return m
}

// WSGCommonDhcpSiteConfigListRef
//
// Definition: common_dhcpSiteConfigListRef
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
	SiteMode *string `json:"siteMode,omitempty"`

	SiteProfiles []*WSGCommonDhcpProfileRef `json:"siteProfiles,omitempty"`

	// ZoneName
	// DHCP Service Zone Name
	ZoneName *string `json:"zoneName,omitempty"`
}

type WSGCommonDhcpSiteConfigListRefAPIResponse struct {
	*RawAPIResponse
	Data *WSGCommonDhcpSiteConfigListRef
}

func newWSGCommonDhcpSiteConfigListRefAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGCommonDhcpSiteConfigListRefAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGCommonDhcpSiteConfigListRefAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGCommonDhcpSiteConfigListRef)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGCommonDhcpSiteConfigListRef() *WSGCommonDhcpSiteConfigListRef {
	m := new(WSGCommonDhcpSiteConfigListRef)
	return m
}

// WSGCommonDhcpSiteConfigListRefSiteApsType
//
// Definition: common_dhcpSiteConfigListRefSiteApsType
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
	ApStatus *string `json:"apStatus,omitempty"`
}

func NewWSGCommonDhcpSiteConfigListRefSiteApsType() *WSGCommonDhcpSiteConfigListRefSiteApsType {
	m := new(WSGCommonDhcpSiteConfigListRefSiteApsType)
	return m
}

// WSGCommonDhcpSiteConfigRef
//
// Definition: common_dhcpSiteConfigRef
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
	SiteMode *string `json:"siteMode,omitempty"`

	SiteProfileIds []string `json:"siteProfileIds,omitempty"`
}

func NewWSGCommonDhcpSiteConfigRef() *WSGCommonDhcpSiteConfigRef {
	m := new(WSGCommonDhcpSiteConfigRef)
	return m
}

// WSGCommonDhcpSiteConfigRefSiteApsType
//
// Definition: common_dhcpSiteConfigRefSiteApsType
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
	ApServerIp *string `json:"apServerIp,omitempty"`

	ApServerPrimary *bool `json:"apServerPrimary,omitempty"`

	ApServerType *string `json:"apServerType,omitempty"`

	// ApStatus
	// Constraints:
	//    - oneof:[Online,Offline,Flagged]
	ApStatus *string `json:"apStatus,omitempty"`
}

func NewWSGCommonDhcpSiteConfigRefSiteApsType() *WSGCommonDhcpSiteConfigRefSiteApsType {
	m := new(WSGCommonDhcpSiteConfigRefSiteApsType)
	return m
}

// WSGCommonDoAssignIp
//
// Definition: common_doAssignIp
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
	SiteMode *string `json:"siteMode,omitempty"`

	SiteProfileIds []string `json:"siteProfileIds,omitempty"`
}

func NewWSGCommonDoAssignIp() *WSGCommonDoAssignIp {
	m := new(WSGCommonDoAssignIp)
	return m
}

// WSGCommonDoAssignIpSiteApsType
//
// Definition: common_doAssignIpSiteApsType
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

// WSGCommonEmail
//
// Definition: common_email
type WSGCommonEmail string

func NewWSGCommonEmail() *WSGCommonEmail {
	m := new(WSGCommonEmail)
	return m
}

// WSGCommonEtherType
//
// Definition: common_etherType
//
// Constraints:
//    - oneof:[0x0800,0x86DD,0x0806,0x8808,0x8809,0x8863,0x8864,0x888E,0x88CC]
type WSGCommonEtherType string

func NewWSGCommonEtherType() *WSGCommonEtherType {
	m := new(WSGCommonEtherType)
	return m
}

// WSGCommonFilterOperator
//
// Definition: common_filterOperator
//
// Constraints:
//    - oneof:[eq,gt,lt,gte,lte]
type WSGCommonFilterOperator string

func NewWSGCommonFilterOperator() *WSGCommonFilterOperator {
	m := new(WSGCommonFilterOperator)
	return m
}

// WSGCommonFirmwareVersion
//
// Definition: common_firmwareVersion
type WSGCommonFirmwareVersion string

func NewWSGCommonFirmwareVersion() *WSGCommonFirmwareVersion {
	m := new(WSGCommonFirmwareVersion)
	return m
}

// WSGCommonFQDN
//
// Definition: common_FQDN
type WSGCommonFQDN string

func NewWSGCommonFQDN() *WSGCommonFQDN {
	m := new(WSGCommonFQDN)
	return m
}

// WSGCommonFullTextSearch
//
// Definition: common_fullTextSearch
type WSGCommonFullTextSearch struct {
	// Fields
	// Specific fields to search
	Fields []string `json:"fields,omitempty"`

	// Type
	// Search logic operator
	// Constraints:
	//    - oneof:[AND,OR]
	Type *string `json:"type,omitempty"`

	// Value
	// Text or number to search
	Value *string `json:"value,omitempty"`
}

func NewWSGCommonFullTextSearch() *WSGCommonFullTextSearch {
	m := new(WSGCommonFullTextSearch)
	return m
}

// WSGCommonGenericRef
//
// Definition: common_genericRef
type WSGCommonGenericRef struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

func NewWSGCommonGenericRef() *WSGCommonGenericRef {
	m := new(WSGCommonGenericRef)
	return m
}

// WSGCommonHealthCheckPolicy
//
// Definition: common_healthCheckPolicy
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
	ResponseWindow *int `json:"responseWindow"`

	// ReviveInterval
	// Revive interval
	// Constraints:
	//    - required
	//    - default:120
	//    - min:60
	//    - max:3600
	ReviveInterval *int `json:"reviveInterval"`

	// ZombiePeriod
	// Zombie period
	// Constraints:
	//    - required
	//    - default:40
	//    - min:30
	//    - max:120
	ZombiePeriod *int `json:"zombiePeriod"`
}

func NewWSGCommonHealthCheckPolicy() *WSGCommonHealthCheckPolicy {
	m := new(WSGCommonHealthCheckPolicy)
	return m
}

// WSGCommonHTTPS
//
// Definition: common_HTTPS
//
// Constraints:
//    - nullable
type WSGCommonHTTPS string

func NewWSGCommonHTTPS() *WSGCommonHTTPS {
	m := new(WSGCommonHTTPS)
	return m
}

// WSGCommonIdList
//
// Definition: common_idList
type WSGCommonIdList []string

func MakeWSGCommonIdList() WSGCommonIdList {
	m := make(WSGCommonIdList, 0)
	return m
}

// WSGCommonIpAddress
//
// Definition: common_ipAddress
type WSGCommonIpAddress string

func NewWSGCommonIpAddress() *WSGCommonIpAddress {
	m := new(WSGCommonIpAddress)
	return m
}

// WSGCommonIpMode
//
// Definition: common_ipMode
//
// Constraints:
//    - oneof:[IPV4,IPV6,IPV4_IPV6]
type WSGCommonIpMode string

func NewWSGCommonIpMode() *WSGCommonIpMode {
	m := new(WSGCommonIpMode)
	return m
}

// WSGCommonLanguageName
//
// Definition: common_languageName
//
// Constraints:
//    - oneof:[English,Chinese,Czech,Danish,Dutch,French,German,Japanese,Spanish,Korean,Swedish,Turkish,eng,chi,cze,dan,dut,fre,ger,jpn,kor,spa,swe,tur]
type WSGCommonLanguageName string

func NewWSGCommonLanguageName() *WSGCommonLanguageName {
	m := new(WSGCommonLanguageName)
	return m
}

// WSGCommonLocation
//
// Definition: common_location
type WSGCommonLocation string

func NewWSGCommonLocation() *WSGCommonLocation {
	m := new(WSGCommonLocation)
	return m
}

// WSGCommonLocationAdditionalInfo
//
// Definition: common_locationAdditionalInfo
type WSGCommonLocationAdditionalInfo string

func NewWSGCommonLocationAdditionalInfo() *WSGCommonLocationAdditionalInfo {
	m := new(WSGCommonLocationAdditionalInfo)
	return m
}

// WSGCommonLteBandLockChannel
//
// Definition: common_lteBandLockChannel
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

// WSGCommonMac
//
// Definition: common_mac
type WSGCommonMac string

func NewWSGCommonMac() *WSGCommonMac {
	m := new(WSGCommonMac)
	return m
}

// WSGCommonMonitoringSummary
//
// Definition: common_monitoringSummary
type WSGCommonMonitoringSummary struct {
	FlaggedCount *int `json:"flaggedCount,omitempty"`

	OfflineCount *int `json:"offlineCount,omitempty"`

	OnlineCount *int `json:"onlineCount,omitempty"`
}

type WSGCommonMonitoringSummaryAPIResponse struct {
	*RawAPIResponse
	Data *WSGCommonMonitoringSummary
}

func newWSGCommonMonitoringSummaryAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGCommonMonitoringSummaryAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGCommonMonitoringSummaryAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGCommonMonitoringSummary)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGCommonMonitoringSummary() *WSGCommonMonitoringSummary {
	m := new(WSGCommonMonitoringSummary)
	return m
}

// WSGCommonMyRuckusConfig
//
// Definition: common_myRuckusConfig
type WSGCommonMyRuckusConfig struct {
	// AclForTunnelWlanAndVlanEnable
	// My.Ruckus support for tunnel-wlan/vlan
	// Constraints:
	//    - required
	AclForTunnelWlanAndVlanEnable *bool `json:"aclForTunnelWlanAndVlanEnable"`
}

func NewWSGCommonMyRuckusConfig() *WSGCommonMyRuckusConfig {
	m := new(WSGCommonMyRuckusConfig)
	return m
}

// WSGCommonNormalName
//
// Definition: common_normalName
//
// Constraints:
//    - max:32
//    - min:2
type WSGCommonNormalName string

func NewWSGCommonNormalName() *WSGCommonNormalName {
	m := new(WSGCommonNormalName)
	return m
}

// WSGCommonNormalName2to64
//
// Definition: common_normalName2to64
//
// Constraints:
//    - max:64
//    - min:2
type WSGCommonNormalName2to64 string

func NewWSGCommonNormalName2to64() *WSGCommonNormalName2to64 {
	m := new(WSGCommonNormalName2to64)
	return m
}

// WSGCommonNormalName2to128
//
// Definition: common_normalName2to128
//
// Constraints:
//    - max:128
//    - min:2
type WSGCommonNormalName2to128 string

func NewWSGCommonNormalName2to128() *WSGCommonNormalName2to128 {
	m := new(WSGCommonNormalName2to128)
	return m
}

// WSGCommonNormalNameAllowBlank
//
// Definition: common_normalNameAllowBlank
//
// Constraints:
//    - nullable
//    - max:64
type WSGCommonNormalNameAllowBlank string

func NewWSGCommonNormalNameAllowBlank() *WSGCommonNormalNameAllowBlank {
	m := new(WSGCommonNormalNameAllowBlank)
	return m
}

// WSGCommonNormalURL
//
// Definition: common_normalURL
type WSGCommonNormalURL string

func NewWSGCommonNormalURL() *WSGCommonNormalURL {
	m := new(WSGCommonNormalURL)
	return m
}

// WSGCommonOui
//
// Definition: common_oui
type WSGCommonOui string

func NewWSGCommonOui() *WSGCommonOui {
	m := new(WSGCommonOui)
	return m
}

// WSGCommonOverrideClientAdmissionControl
//
// Definition: common_overrideClientAdmissionControl
type WSGCommonOverrideClientAdmissionControl struct {
	Enabled *bool `json:"enabled,omitempty"`

	// MaxRadioLoadPercent
	// Maximum radio load percentage.
	// Constraints:
	//    - default:75
	//    - min:50
	//    - max:100
	MaxRadioLoadPercent *int `json:"maxRadioLoadPercent,omitempty"`

	// MinClientCount
	// Minimum client count number.
	// Constraints:
	//    - default:10
	//    - min:0
	//    - max:100
	MinClientCount *int `json:"minClientCount,omitempty"`

	// MinClientThroughputMbps
	// Minimum client throughput in Mbps.
	// Constraints:
	//    - default:0
	//    - min:0.000000
	//    - max:100.000000
	MinClientThroughputMbps *float64 `json:"minClientThroughputMbps,omitempty"`
}

func NewWSGCommonOverrideClientAdmissionControl() *WSGCommonOverrideClientAdmissionControl {
	m := new(WSGCommonOverrideClientAdmissionControl)
	return m
}

// WSGCommonOverrideGenericRef
//
// Definition: common_overrideGenericRef
type WSGCommonOverrideGenericRef struct {
	Enabled *bool `json:"enabled,omitempty"`

	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

func NewWSGCommonOverrideGenericRef() *WSGCommonOverrideGenericRef {
	m := new(WSGCommonOverrideGenericRef)
	return m
}

// WSGCommonOverrideSmartMonitor
//
// Definition: common_overrideSmartMonitor
type WSGCommonOverrideSmartMonitor struct {
	Enabled *bool `json:"enabled,omitempty"`

	// IntervalInSec
	// Interval in seconds. This is required if smartMonitor is enabled
	// Constraints:
	//    - default:10
	//    - min:5
	//    - max:60
	IntervalInSec *int `json:"intervalInSec,omitempty"`

	// RetryThreshold
	// Retry threshold. This is required if smartMonitor is enabled
	// Constraints:
	//    - default:3
	//    - min:1
	//    - max:10
	RetryThreshold *int `json:"retryThreshold,omitempty"`
}

func NewWSGCommonOverrideSmartMonitor() *WSGCommonOverrideSmartMonitor {
	m := new(WSGCommonOverrideSmartMonitor)
	return m
}

// WSGCommonPoeModeSetting
//
// Definition: common_poeModeSetting
//
// Constraints:
//    - nullable
//    - oneof:[Auto,_802_3af,_802_3at,_802_3atPlus,_802_3bt_Class_5,_802_3bt_Class_6,_802_3bt_Class_7,_802_3bt_Class_8]
type WSGCommonPoeModeSetting string

func NewWSGCommonPoeModeSetting() *WSGCommonPoeModeSetting {
	m := new(WSGCommonPoeModeSetting)
	return m
}

// WSGCommonPortalCustomization
//
// Definition: common_portalCustomization
type WSGCommonPortalCustomization struct {
	// Language
	// Constraints:
	//    - required
	Language *WSGCommonPortalLanguage `json:"language"`

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
	TermsAndConditionsText *string `json:"termsAndConditionsText,omitempty"`

	// Title
	// Title
	// Constraints:
	//    - max:63
	//    - min:0
	Title *string `json:"title,omitempty"`
}

func NewWSGCommonPortalCustomization() *WSGCommonPortalCustomization {
	m := new(WSGCommonPortalCustomization)
	return m
}

// WSGCommonPortalLanguage
//
// Definition: common_portalLanguage
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

// WSGCommonProtectionMode
//
// Definition: common_protectionMode
//
// Constraints:
//    - default:'RTS_CTS'
//    - oneof:[NONE,CTS_ONLY,RTS_CTS]
type WSGCommonProtectionMode string

func NewWSGCommonProtectionMode() *WSGCommonProtectionMode {
	m := new(WSGCommonProtectionMode)
	return m
}

// WSGCommonQinq
//
// Definition: common_qinq
type WSGCommonQinq struct {
	// Cvlan
	// Constraints:
	//    - min:2
	//    - max:4094
	Cvlan *int `json:"cvlan,omitempty"`

	// Svlan
	// Constraints:
	//    - min:2
	//    - max:4094
	Svlan *int `json:"svlan,omitempty"`
}

func NewWSGCommonQinq() *WSGCommonQinq {
	m := new(WSGCommonQinq)
	return m
}

// WSGCommonQueryCriteria
//
// Definition: common_queryCriteria
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
	Limit *int `json:"limit,omitempty"`

	// Options
	// Specified feature required information
	Options interface{} `json:"options,omitempty"`

	// Page
	// Page number to get
	// Constraints:
	//    - min:1
	Page *int `json:"page,omitempty"`

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

// WSGCommonQueryCriteriaExtraFiltersType
//
// Definition: common_queryCriteriaExtraFiltersType
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

// WSGCommonQueryCriteriaExtraNotFiltersType
//
// Definition: common_queryCriteriaExtraNotFiltersType
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

// WSGCommonQueryCriteriaFiltersType
//
// Definition: common_queryCriteriaFiltersType
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

// WSGCommonQueryCriteriaSortInfoType
//
// Definition: common_queryCriteriaSortInfoType
//
// About sorting
type WSGCommonQueryCriteriaSortInfoType struct {
	// Dir
	// Constraints:
	//    - oneof:[ASC,DESC]
	Dir *string `json:"dir,omitempty"`

	SortColumn *string `json:"sortColumn,omitempty"`
}

func NewWSGCommonQueryCriteriaSortInfoType() *WSGCommonQueryCriteriaSortInfoType {
	m := new(WSGCommonQueryCriteriaSortInfoType)
	return m
}

// WSGCommonQueryCriteriaSuperSet
//
// Definition: common_queryCriteriaSuperSet
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
	ExtraFilters []*WSGCommonQueryCriteriaSuperSetExtraFiltersType `json:"extraFilters,omitempty"`

	// ExtraNotFilters
	// "NOT" condition for multiple filters
	ExtraNotFilters []*WSGCommonQueryCriteriaSuperSetExtraNotFiltersType `json:"extraNotFilters,omitempty"`

	ExtraTimeRange *WSGCommonTimeRange `json:"extraTimeRange,omitempty"`

	// Filters
	// Filters used to select specific resource scope
	Filters []*WSGCommonQueryCriteriaSuperSetFiltersType `json:"filters,omitempty"`

	FullTextSearch *WSGCommonFullTextSearch `json:"fullTextSearch,omitempty"`

	// Limit
	// Size of one page
	// Constraints:
	//    - min:1
	Limit *int `json:"limit,omitempty"`

	// Options
	// Specified feature required information
	Options *WSGCommonQueryCriteriaSuperSetOptionsType `json:"options,omitempty"`

	// Page
	// Page number to get
	// Constraints:
	//    - min:1
	Page *int `json:"page,omitempty"`

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

// WSGCommonQueryCriteriaSuperSetExtraFiltersType
//
// Definition: common_queryCriteriaSuperSetExtraFiltersType
type WSGCommonQueryCriteriaSuperSetExtraFiltersType struct {
	Operator *WSGCommonFilterOperator `json:"operator,omitempty"`

	// Type
	// Filters for specific attribute
	// Constraints:
	//    - oneof:[CONTROLBLADE,DATABLADE,DOMAIN,ZONE,THIRD_PARTY_ZONE,APGROUP,WLANGROUP,INDOORMAP,AP,WLAN,ProtocolType,TIMERANGE,RADIOID,WLANID,CATEGORY,CLIENT,CP,DP,CLUSTER,NODE,BLADE,SYNCEDSTATUS,OSTYPE,APP,PORT,STATUS,REGISTRATIONSTATE,GATEWAY,APIPADDRESS,CLIENTIPADDRESS,CLIENTIPV6ADDRESS,SEVERITY,ACKNOWLEDGED,MVNOID,USER,USERID,WLANNAME,AUDITIPADDRESS,AUDITUSERUUID,AUDITOBJECT,AUDITACTION,AUDITTENANTUUID,AUDITOBJECTUUID,AUTHTYPE,AUDITTYPE,H20SuppportEnabled,AaaSuppportEnabled,GppSuppportEnabled,Type,RogueMac,SSID,ALARMSTATE,DEVICENAME,SWITCH,ZoneAffinityProfileId,MONITORINGENABLED]
	Type *string `json:"type,omitempty"`

	// Value
	// value to search
	Value *string `json:"value,omitempty"`
}

func NewWSGCommonQueryCriteriaSuperSetExtraFiltersType() *WSGCommonQueryCriteriaSuperSetExtraFiltersType {
	m := new(WSGCommonQueryCriteriaSuperSetExtraFiltersType)
	return m
}

// WSGCommonQueryCriteriaSuperSetExtraNotFiltersType
//
// Definition: common_queryCriteriaSuperSetExtraNotFiltersType
type WSGCommonQueryCriteriaSuperSetExtraNotFiltersType struct {
	// Type
	// Filters for specific attribute
	// Constraints:
	//    - oneof:[CONTROLBLADE,DATABLADE,DOMAIN,ZONE,THIRD_PARTY_ZONE,APGROUP,WLANGROUP,INDOORMAP,AP,WLAN,ProtocolType,TIMERANGE,RADIOID,WLANID,CATEGORY,CLIENT,CP,DP,CLUSTER,NODE,BLADE,SYNCEDSTATUS,OSTYPE,APP,PORT,STATUS,REGISTRATIONSTATE,GATEWAY,APIPADDRESS,CLIENTIPADDRESS,SEVERITY,ACKNOWLEDGED,MVNOID,USER,USERID,WLANNAME,AUDITIPADDRESS,AUDITUSERUUID,AUDITOBJECT,AUDITACTION,AUDITTENANTUUID,AUDITOBJECTUUID,AUTHTYPE,AUDITTYPE,H20SuppportEnabled,AaaSuppportEnabled,GppSuppportEnabled,MONITORINGENABLED]
	Type *string `json:"type,omitempty"`

	// Value
	// value not to search
	Value *string `json:"value,omitempty"`
}

func NewWSGCommonQueryCriteriaSuperSetExtraNotFiltersType() *WSGCommonQueryCriteriaSuperSetExtraNotFiltersType {
	m := new(WSGCommonQueryCriteriaSuperSetExtraNotFiltersType)
	return m
}

// WSGCommonQueryCriteriaSuperSetFiltersType
//
// Definition: common_queryCriteriaSuperSetFiltersType
type WSGCommonQueryCriteriaSuperSetFiltersType struct {
	Operator *WSGCommonFilterOperator `json:"operator,omitempty"`

	// Type
	// Group type
	// Constraints:
	//    - oneof:[SYSTEM,CATEGORY,CONTROLBLADE,DATABLADE,DOMAIN,ZONE,THIRD_PARTY_ZONE,APGROUP,WLANGROUP,INDOORMAP,AP,WLAN,BLADE,SYNCEDSTATUS,REGISTRATIONSTATE,STATUS,SWITCH_GROUP]
	Type *string `json:"type,omitempty"`

	// Value
	// Group ID
	Value *string `json:"value,omitempty"`
}

func NewWSGCommonQueryCriteriaSuperSetFiltersType() *WSGCommonQueryCriteriaSuperSetFiltersType {
	m := new(WSGCommonQueryCriteriaSuperSetFiltersType)
	return m
}

// WSGCommonQueryCriteriaSuperSetOptionsType
//
// Definition: common_queryCriteriaSuperSetOptionsType
//
// Specified feature required information
type WSGCommonQueryCriteriaSuperSetOptionsType struct {
	// AcctincludeNa
	// include Not Available acct service option while returning result
	AcctincludeNa *bool `json:"acct_includeNa,omitempty"`

	// AccttestableOnly
	// only get testable service type
	AccttestableOnly *bool `json:"acct_testableOnly,omitempty"`

	// Accttype
	// accounting service types to get, use comma to separate, Ex: RADIUS,CGF
	Accttype *string `json:"acct_type,omitempty"`

	// AuthhostedAaaSupportedEnabled
	// Indicate if Hosted AAA Support is enabled
	AuthhostedAaaSupportedEnabled *bool `json:"auth_hostedAaaSupportedEnabled,omitempty"`

	// AuthincludeAdGlobal
	// If AD is in list, include only AD with Global Catalog configured
	AuthincludeAdGlobal *bool `json:"auth_includeAdGlobal,omitempty"`

	// AuthincludeGuest
	// include Guest auth service while returning result
	AuthincludeGuest *bool `json:"auth_includeGuest,omitempty"`

	// AuthincludeLocalDb
	// include LocalDB auth service while returning result
	AuthincludeLocalDb *bool `json:"auth_includeLocalDb,omitempty"`

	// AuthincludeNa
	// include Not Available auth service option while returning result
	AuthincludeNa *bool `json:"auth_includeNa,omitempty"`

	// AuthplmnIdentifierEnabled
	// Indicate if Configure PLMN identifier is enabled
	AuthplmnIdentifierEnabled *bool `json:"auth_plmnIdentifierEnabled,omitempty"`

	// AuthrealmType
	// To get specific authentication service information for configuring realm based authentication profile
	// Constraints:
	//    - oneof:[ALL,RADIUS]
	AuthrealmType *string `json:"auth_realmType,omitempty"`

	// AuthtestableOnly
	// only get testable service type
	AuthtestableOnly *bool `json:"auth_testableOnly,omitempty"`

	// Authtype
	// authentication service types to get, use comma to separate, Ex: RADIUS,AD
	Authtype *string `json:"auth_type,omitempty"`

	// Forwardingtype
	// forwarding service types to get, use comma to separate, Ex: L2oGRE,TTGPDG,Bridge,Advanced
	Forwardingtype *string `json:"forwarding_type,omitempty"`

	// GlobalFilterId
	// Specify GlobalFilter ID for query.
	GlobalFilterId *string `json:"globalFilterId,omitempty"`

	// IncludeSharedResources
	// Whether to include the resources of parent domain or not.
	IncludeSharedResources *bool `json:"includeSharedResources,omitempty"`

	// IncludeUserClickNode
	// Can be used when group tree rendering needs include user clicked node.
	IncludeUserClickNode *bool `json:"includeUserClickNode,omitempty"`

	// IncludeUsers
	// Should also retrieve users or not
	IncludeUsers *bool `json:"includeUsers,omitempty"`

	// INCLUDERBACMETADATA
	// Whether to include RBAC metadata or not.
	INCLUDERBACMETADATA *bool `json:"INCLUDE_RBAC_METADATA,omitempty"`

	// InMap
	// Specify inMap status for query.
	InMap *bool `json:"inMap,omitempty"`

	// TENANTID
	// Specify Tenant ID for query.
	TENANTID *string `json:"TENANT_ID,omitempty"`
}

func NewWSGCommonQueryCriteriaSuperSetOptionsType() *WSGCommonQueryCriteriaSuperSetOptionsType {
	m := new(WSGCommonQueryCriteriaSuperSetOptionsType)
	return m
}

// WSGCommonRadio24
//
// Definition: common_radio24
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
	ChannelWidth *int `json:"channelWidth,omitempty"`

	TxPower *WSGCommonTxPower `json:"txPower,omitempty"`
}

func NewWSGCommonRadio24() *WSGCommonRadio24 {
	m := new(WSGCommonRadio24)
	return m
}

// WSGCommonRadio24SuperSet
//
// Definition: common_radio24SuperSet
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
	ChannelWidth *int `json:"channelWidth,omitempty"`

	TxPower *WSGCommonTxPower `json:"txPower,omitempty"`
}

func NewWSGCommonRadio24SuperSet() *WSGCommonRadio24SuperSet {
	m := new(WSGCommonRadio24SuperSet)
	return m
}

// WSGCommonRadio50
//
// Definition: common_radio50
type WSGCommonRadio50 struct {
	// AutoCellSizing
	// Auto Cell Sizing
	AutoCellSizing *bool `json:"autoCellSizing,omitempty"`

	// ChannelWidth
	// Channel width. Zero (0) means Auto. 8080 means 80+80MHz
	// Constraints:
	//    - default:0
	//    - oneof:[0,20,40,80,8080,160]
	ChannelWidth *int `json:"channelWidth,omitempty"`

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

// WSGCommonRadio50SuperSet
//
// Definition: common_radio50SuperSet
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
	ChannelWidth *int `json:"channelWidth,omitempty"`

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

// WSGCommonRadiusServer
//
// Definition: common_radiusServer
type WSGCommonRadiusServer struct {
	// Ip
	// Constraints:
	//    - required
	Ip *WSGCommonIpAddress `json:"ip"`

	// Port
	// Server port
	// Constraints:
	//    - required
	//    - min:1
	//    - max:65535
	Port *int `json:"port"`

	// SharedSecret
	// Server shared secrect
	// Constraints:
	//    - required
	SharedSecret *string `json:"sharedSecret"`
}

func NewWSGCommonRadiusServer() *WSGCommonRadiusServer {
	m := new(WSGCommonRadiusServer)
	return m
}

// WSGCommonRateLimiting
//
// Definition: common_rateLimiting
type WSGCommonRateLimiting struct {
	// MaxOutstandingRequestsPerServer
	// Maximum outstanding requests (MOR), value should be 0 or between 10 and 4096
	// Constraints:
	//    - required
	//    - default:0
	MaxOutstandingRequestsPerServer *int `json:"maxOutstandingRequestsPerServer"`

	// SanityTimer
	// Sanity timer
	// Constraints:
	//    - required
	//    - default:10
	//    - min:1
	//    - max:3600
	SanityTimer *int `json:"sanityTimer"`

	// Threshold
	// Threshold, value should be 0 if MOR is 0, or between 10 and 90 if MOR is between 10 and 4096
	// Constraints:
	//    - required
	//    - default:0
	Threshold *int `json:"threshold"`
}

func NewWSGCommonRateLimiting() *WSGCommonRateLimiting {
	m := new(WSGCommonRateLimiting)
	return m
}

// WSGCommonRealm
//
// Definition: common_realm
//
// Constraints:
//    - max:255
type WSGCommonRealm string

func NewWSGCommonRealm() *WSGCommonRealm {
	m := new(WSGCommonRealm)
	return m
}

// WSGCommonRecoverySsid
//
// Definition: common_recoverySsid
type WSGCommonRecoverySsid struct {
	// RecoverySsidEnabled
	// recovery ssid enable/disable
	RecoverySsidEnabled *bool `json:"recoverySsidEnabled,omitempty"`
}

func NewWSGCommonRecoverySsid() *WSGCommonRecoverySsid {
	m := new(WSGCommonRecoverySsid)
	return m
}

// WSGCommonSmartMonitor
//
// Definition: common_smartMonitor
type WSGCommonSmartMonitor struct {
	// IntervalInSec
	// Interval in seconds. This is required if smartMonitor is enabled
	// Constraints:
	//    - default:10
	//    - min:5
	//    - max:60
	IntervalInSec *int `json:"intervalInSec,omitempty"`

	// RetryThreshold
	// Retry threshold. This is required if smartMonitor is enabled
	// Constraints:
	//    - default:3
	//    - min:1
	//    - max:10
	RetryThreshold *int `json:"retryThreshold,omitempty"`
}

func NewWSGCommonSmartMonitor() *WSGCommonSmartMonitor {
	m := new(WSGCommonSmartMonitor)
	return m
}

// WSGCommonSnmpCommunity
//
// Definition: common_snmpCommunity
type WSGCommonSnmpCommunity struct {
	// CommunityName
	// name of the SNMP Community.
	// Constraints:
	//    - required
	CommunityName *string `json:"communityName"`

	// NotificationEnabled
	// notification privilege of the SNMP Coummunity
	// Constraints:
	//    - required
	NotificationEnabled *bool `json:"notificationEnabled"`

	// NotificationTarget
	// Trap List of the SNMP Coummunity
	NotificationTarget []*WSGCommonTargetConfig `json:"notificationTarget,omitempty"`

	// NotificationType
	// type of the notification privilege
	// Constraints:
	//    - oneof:[TRAP,INFORM]
	NotificationType *string `json:"notificationType,omitempty"`

	// ReadEnabled
	// read privilege of the SNMP Coummunity
	// Constraints:
	//    - required
	ReadEnabled *bool `json:"readEnabled"`

	// WriteEnabled
	// write privilege of the SNMP Coummunity
	// Constraints:
	//    - required
	WriteEnabled *bool `json:"writeEnabled"`
}

func NewWSGCommonSnmpCommunity() *WSGCommonSnmpCommunity {
	m := new(WSGCommonSnmpCommunity)
	return m
}

// WSGCommonSnmpUser
//
// Definition: common_snmpUser
type WSGCommonSnmpUser struct {
	// AuthPassword
	// authPassword of the SNMP User.
	// Constraints:
	//    - min:8
	AuthPassword *string `json:"authPassword,omitempty"`

	// AuthProtocol
	// authProtocol of the SNMP User.
	// Constraints:
	//    - oneof:[MD5,SHA]
	AuthProtocol *string `json:"authProtocol,omitempty"`

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
	NotificationType *string `json:"notificationType,omitempty"`

	// PrivPassword
	// privPassword of the SNMP User.
	// Constraints:
	//    - min:8
	PrivPassword *string `json:"privPassword,omitempty"`

	// PrivProtocol
	// privProtocol of the SNMP User.
	// Constraints:
	//    - oneof:[DES,AES]
	PrivProtocol *string `json:"privProtocol,omitempty"`

	// ReadEnabled
	// read privilege of the SNMP User
	ReadEnabled *bool `json:"readEnabled,omitempty"`

	// UserName
	// name of the SNMP User.
	// Constraints:
	//    - required
	UserName *string `json:"userName"`

	// WriteEnabled
	// write privilege of the SNMP User
	WriteEnabled *bool `json:"writeEnabled,omitempty"`
}

func NewWSGCommonSnmpUser() *WSGCommonSnmpUser {
	m := new(WSGCommonSnmpUser)
	return m
}

// WSGCommonSubNetMask
//
// Definition: common_subNetMask
type WSGCommonSubNetMask string

func NewWSGCommonSubNetMask() *WSGCommonSubNetMask {
	m := new(WSGCommonSubNetMask)
	return m
}

// WSGCommonTargetConfig
//
// Definition: common_targetConfig
type WSGCommonTargetConfig struct {
	// Address
	// address of the SNMP Trap
	// Constraints:
	//    - required
	Address *string `json:"address"`

	// Port
	// port number of the SNMP Trap
	// Constraints:
	//    - required
	//    - min:1
	//    - max:65535
	Port *int `json:"port"`
}

func NewWSGCommonTargetConfig() *WSGCommonTargetConfig {
	m := new(WSGCommonTargetConfig)
	return m
}

// WSGCommonTimeRange
//
// Definition: common_timeRange
type WSGCommonTimeRange struct {
	// End
	// end time for collecting data
	End *float64 `json:"end,omitempty"`

	// Field
	// time field for collecting data
	// Constraints:
	//    - oneof:[insertionTime]
	Field *string `json:"field,omitempty"`

	// Interval
	// time interval in second
	Interval *float64 `json:"interval,omitempty"`

	// Start
	// start time for collecting data
	Start *float64 `json:"start,omitempty"`
}

func NewWSGCommonTimeRange() *WSGCommonTimeRange {
	m := new(WSGCommonTimeRange)
	return m
}

// WSGCommonTimeUnitStore
//
// Definition: common_timeUnitStore
//
// time unit
// Constraints:
//    - oneof:[second,minute,hour,day]
type WSGCommonTimeUnitStore string

func NewWSGCommonTimeUnitStore() *WSGCommonTimeUnitStore {
	m := new(WSGCommonTimeUnitStore)
	return m
}

// WSGCommonTrafficClassProfileRef
//
// Definition: common_trafficClassProfileRef
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

type WSGCommonTrafficClassProfileRefAPIResponse struct {
	*RawAPIResponse
	Data *WSGCommonTrafficClassProfileRef
}

func newWSGCommonTrafficClassProfileRefAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGCommonTrafficClassProfileRefAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGCommonTrafficClassProfileRefAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGCommonTrafficClassProfileRef)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGCommonTrafficClassProfileRef() *WSGCommonTrafficClassProfileRef {
	m := new(WSGCommonTrafficClassProfileRef)
	return m
}

// WSGCommonTrafficClassRef
//
// Definition: common_trafficClassRef
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

// WSGCommonTxPower
//
// Definition: common_txPower
//
// Constraints:
//    - default:'Full'
//    - oneof:[Full,-1dB,-2dB,-3dB(1/2),-4dB,-5dB,-6dB(1/4),-7dB,-8dB,-9dB(1/8),-10dB,Min]
type WSGCommonTxPower string

func NewWSGCommonTxPower() *WSGCommonTxPower {
	m := new(WSGCommonTxPower)
	return m
}

// WSGCommonWebAuthenticationPortalCustomization
//
// Definition: common_webAuthenticationPortalCustomization
type WSGCommonWebAuthenticationPortalCustomization struct {
	// Logo
	// Logo encoded with base64, format is "data:image/png;base64,the base64 encoded logo"
	// Constraints:
	//    - required
	Logo *string `json:"logo"`

	// Title
	// Title of the custom portal
	// Constraints:
	//    - max:63
	//    - min:0
	Title *string `json:"title,omitempty"`
}

func NewWSGCommonWebAuthenticationPortalCustomization() *WSGCommonWebAuthenticationPortalCustomization {
	m := new(WSGCommonWebAuthenticationPortalCustomization)
	return m
}

// WSGCommonWildFQDN
//
// Definition: common_wildFQDN
//
// Compare with FQDN, it could start with '*.'
type WSGCommonWildFQDN string

func NewWSGCommonWildFQDN() *WSGCommonWildFQDN {
	m := new(WSGCommonWildFQDN)
	return m
}

// WSGCommonZoneTunnelType
//
// Definition: common_zoneTunnelType
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
