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

	Id *Mac `json:"id,omitempty"`

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

func NewAlarm() *Alarm {
	alarmType := new(Alarm)
	return alarmType
}

func NewAlarmWithDefaults() *Alarm {
	alarmType := new(Alarm)
	return alarmType
}

type Altitude struct {
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

func NewAltitude() *Altitude {
	altitudeType := new(Altitude)
	return altitudeType
}

func NewAltitudeWithDefaults() *Altitude {
	altitudeType := new(Altitude)
	altitudeUnitField := `meters`
	altitudeType.AltitudeUnit = &altitudeUnitField
	return altitudeType
}

// ApGpsSource
//
// GPS Source of the AP
// Constraints:
//    - nullable
//    - oneof:[GPS,MANUAL]
type ApGpsSource string

func NewApGpsSource() *ApGpsSource {
	apGpsSourceType := new(ApGpsSource)
	return apGpsSourceType
}

func NewApGpsSourceWithDefaults() *ApGpsSource {
	apGpsSourceType := new(ApGpsSource)
	return apGpsSourceType
}

type ApLatencyInterval struct {
	// PingEnabled
	// AP ping latency enabled
	PingEnabled *bool `json:"pingEnabled,omitempty"`
}

func NewApLatencyInterval() *ApLatencyInterval {
	apLatencyIntervalType := new(ApLatencyInterval)
	return apLatencyIntervalType
}

func NewApLatencyIntervalWithDefaults() *ApLatencyInterval {
	apLatencyIntervalType := new(ApLatencyInterval)
	pingEnabledField := false
	apLatencyIntervalType.PingEnabled = &pingEnabledField
	return apLatencyIntervalType
}

type ApLoginName string

func NewApLoginName() *ApLoginName {
	apLoginNameType := new(ApLoginName)
	return apLoginNameType
}

func NewApLoginNameWithDefaults() *ApLoginName {
	apLoginNameType := new(ApLoginName)
	return apLoginNameType
}

type ApLoginPassword string

func NewApLoginPassword() *ApLoginPassword {
	apLoginPasswordType := new(ApLoginPassword)
	return apLoginPasswordType
}

func NewApLoginPasswordWithDefaults() *ApLoginPassword {
	apLoginPasswordType := new(ApLoginPassword)
	return apLoginPasswordType
}

type ApManagementVlan struct {
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

func NewApManagementVlan() *ApManagementVlan {
	apManagementVlanType := new(ApManagementVlan)
	return apManagementVlanType
}

func NewApManagementVlanWithDefaults() *ApManagementVlan {
	apManagementVlanType := new(ApManagementVlan)
	idField := 1
	apManagementVlanType.Id = &idField
	modeField := `KEEP`
	apManagementVlanType.Mode = &modeField
	return apManagementVlanType
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
	// Constraints:
	//    - nullable
	//    - oneof:[0,20,40,80,8080,160]
	ChannelWidth *int `json:"channelWidth,omitempty" validate:"omitempty,oneof=0 20 40 80 8080 160"`

	// SecondaryChannel
	// channel number (channelWidth is 80+80MHz only)
	SecondaryChannel *int `json:"secondaryChannel,omitempty"`

	TxPower *TxPower `json:"txPower,omitempty"`
}

func NewApRadio50() *ApRadio50 {
	apRadio50Type := new(ApRadio50)
	return apRadio50Type
}

func NewApRadio50WithDefaults() *ApRadio50 {
	apRadio50Type := new(ApRadio50)
	return apRadio50Type
}

type ApRebootTimeout struct {
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

func NewApRebootTimeout() *ApRebootTimeout {
	apRebootTimeoutType := new(ApRebootTimeout)
	return apRebootTimeoutType
}

func NewApRebootTimeoutWithDefaults() *ApRebootTimeout {
	apRebootTimeoutType := new(ApRebootTimeout)
	gatewayLossTimeoutInSecField := 1800
	apRebootTimeoutType.GatewayLossTimeoutInSec = &gatewayLossTimeoutInSecField
	serverLossTimeoutInSecField := 7200
	apRebootTimeoutType.ServerLossTimeoutInSec = &serverLossTimeoutInSecField
	return apRebootTimeoutType
}

type AutoChannelSelection struct {
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

func NewAutoChannelSelection() *AutoChannelSelection {
	autoChannelSelectionType := new(AutoChannelSelection)
	return autoChannelSelectionType
}

func NewAutoChannelSelectionWithDefaults() *AutoChannelSelection {
	autoChannelSelectionType := new(AutoChannelSelection)
	channelFlyMtbcField := 480
	autoChannelSelectionType.ChannelFlyMtbc = &channelFlyMtbcField
	channelSelectModeField := `BackgroundScanning`
	autoChannelSelectionType.ChannelSelectMode = &channelSelectModeField
	return autoChannelSelectionType
}

type BaseServiceInfo struct {
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

func NewBaseServiceInfo() *BaseServiceInfo {
	baseServiceInfoType := new(BaseServiceInfo)
	return baseServiceInfoType
}

func NewBaseServiceInfoWithDefaults() *BaseServiceInfo {
	baseServiceInfoType := new(BaseServiceInfo)
	return baseServiceInfoType
}

type BulkDeleteRequest struct {
	IdList IdList `json:"idList,omitempty"`
}

func NewBulkDeleteRequest() *BulkDeleteRequest {
	bulkDeleteRequestType := new(BulkDeleteRequest)
	return bulkDeleteRequestType
}

func NewBulkDeleteRequestWithDefaults() *BulkDeleteRequest {
	bulkDeleteRequestType := new(BulkDeleteRequest)
	return bulkDeleteRequestType
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

	IpAddress *IpAddress `json:"ipAddress,omitempty"`

	Ipv6Address *IpAddress `json:"ipv6Address,omitempty"`

	Mac *Mac `json:"mac,omitempty"`

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

func NewClient() *Client {
	clientType := new(Client)
	return clientType
}

func NewClientWithDefaults() *Client {
	clientType := new(Client)
	return clientType
}

type ClientAdmissionControl struct {
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

func NewClientAdmissionControl() *ClientAdmissionControl {
	clientAdmissionControlType := new(ClientAdmissionControl)
	return clientAdmissionControlType
}

func NewClientAdmissionControlWithDefaults() *ClientAdmissionControl {
	clientAdmissionControlType := new(ClientAdmissionControl)
	maxRadioLoadPercentField := 75
	clientAdmissionControlType.MaxRadioLoadPercent = &maxRadioLoadPercentField
	minClientCountField := 10
	clientAdmissionControlType.MinClientCount = &minClientCountField
	minClientThroughputMbpsField := 0.000000
	clientAdmissionControlType.MinClientThroughputMbps = &minClientThroughputMbpsField
	return clientAdmissionControlType
}

type CreateResult struct {
	Id *string `json:"id,omitempty"`
}

func NewCreateResult() *CreateResult {
	createResultType := new(CreateResult)
	return createResultType
}

func NewCreateResultWithDefaults() *CreateResult {
	createResultType := new(CreateResult)
	return createResultType
}

type CreateResultIdName struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

func NewCreateResultIdName() *CreateResultIdName {
	createResultIdNameType := new(CreateResultIdName)
	return createResultIdNameType
}

func NewCreateResultIdNameWithDefaults() *CreateResultIdName {
	createResultIdNameType := new(CreateResultIdName)
	return createResultIdNameType
}

type Description string

func NewDescription() *Description {
	descriptionType := new(Description)
	return descriptionType
}

func NewDescriptionWithDefaults() *Description {
	descriptionType := new(Description)
	return descriptionType
}

type DescriptionTo128 string

func NewDescriptionTo128() *DescriptionTo128 {
	descriptionTo128Type := new(DescriptionTo128)
	return descriptionTo128Type
}

func NewDescriptionTo128WithDefaults() *DescriptionTo128 {
	descriptionTo128Type := new(DescriptionTo128)
	return descriptionTo128Type
}

type DhcpProfileRef struct {
	Description *Description `json:"description,omitempty"`

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

	Name *NormalName `json:"name,omitempty"`

	PoolEndIp *IpAddress `json:"poolEndIp,omitempty"`

	PoolStartIp *IpAddress `json:"poolStartIp,omitempty"`

	PrimaryDnsIp *IpAddress `json:"primaryDnsIp,omitempty"`

	SecondaryDnsIp *IpAddress `json:"secondaryDnsIp,omitempty"`

	SubnetMask *IpAddress `json:"subnetMask,omitempty"`

	SubnetNetworkIp *IpAddress `json:"subnetNetworkIp,omitempty"`

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

func NewDhcpProfileRef() *DhcpProfileRef {
	dhcpProfileRefType := new(DhcpProfileRef)
	return dhcpProfileRefType
}

func NewDhcpProfileRefWithDefaults() *DhcpProfileRef {
	dhcpProfileRefType := new(DhcpProfileRef)
	return dhcpProfileRefType
}

type DhcpSiteConfigListRef struct {
	// DwpdEnabled
	// DHCP Service Dynamic WAN Port Detection
	DwpdEnabled *bool `json:"dwpdEnabled,omitempty"`

	Eth0ProfileId *int `json:"eth0ProfileId,omitempty"`

	Eth1ProfileId *int `json:"eth1ProfileId,omitempty"`

	// ManualSelect
	// DHCP Service AP Selection Mode
	ManualSelect *bool `json:"manualSelect,omitempty"`

	SiteAps []*DhcpSiteConfigListRefSiteApsType `json:"siteAps,omitempty"`

	// SiteEnabled
	// DHCP Service Enabling Status
	SiteEnabled *bool `json:"siteEnabled,omitempty"`

	// SiteMode
	// DHCP Service mode
	// Constraints:
	//    - nullable
	//    - oneof:[EnableOnEachAPs,EnableOnMultipleAPs,EnableOnHierarchicalAPs]
	SiteMode *string `json:"siteMode,omitempty" validate:"omitempty,oneof=EnableOnEachAPs EnableOnMultipleAPs EnableOnHierarchicalAPs"`

	SiteProfiles []*DhcpProfileRef `json:"siteProfiles,omitempty"`

	// ZoneName
	// DHCP Service Zone Name
	ZoneName *string `json:"zoneName,omitempty"`
}

func NewDhcpSiteConfigListRef() *DhcpSiteConfigListRef {
	dhcpSiteConfigListRefType := new(DhcpSiteConfigListRef)
	return dhcpSiteConfigListRefType
}

func NewDhcpSiteConfigListRefWithDefaults() *DhcpSiteConfigListRef {
	dhcpSiteConfigListRefType := new(DhcpSiteConfigListRef)
	return dhcpSiteConfigListRefType
}

// DhcpSiteConfigListRefSiteApsType
//
// DHCP Site selected APs
type DhcpSiteConfigListRefSiteApsType struct {
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

func NewDhcpSiteConfigListRefSiteApsType() *DhcpSiteConfigListRefSiteApsType {
	dhcpSiteConfigListRefSiteApsTypeType := new(DhcpSiteConfigListRefSiteApsType)
	return dhcpSiteConfigListRefSiteApsTypeType
}

func NewDhcpSiteConfigListRefSiteApsTypeWithDefaults() *DhcpSiteConfigListRefSiteApsType {
	dhcpSiteConfigListRefSiteApsTypeType := new(DhcpSiteConfigListRefSiteApsType)
	return dhcpSiteConfigListRefSiteApsTypeType
}

type DhcpSiteConfigRef struct {
	// DwpdEnabled
	// DHCP Service Dynamic WAN Port Detection
	DwpdEnabled *bool `json:"dwpdEnabled,omitempty"`

	Eth0ProfileId *int `json:"eth0ProfileId,omitempty"`

	Eth1ProfileId *int `json:"eth1ProfileId,omitempty"`

	// ManualSelect
	// DHCP Service AP Selection Mode. This value is effective when the siteMode is EnableOnMultipleAPs.
	ManualSelect *bool `json:"manualSelect,omitempty"`

	SiteAps []*DhcpSiteConfigRefSiteApsType `json:"siteAps,omitempty"`

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

func NewDhcpSiteConfigRef() *DhcpSiteConfigRef {
	dhcpSiteConfigRefType := new(DhcpSiteConfigRef)
	return dhcpSiteConfigRefType
}

func NewDhcpSiteConfigRefWithDefaults() *DhcpSiteConfigRef {
	dhcpSiteConfigRefType := new(DhcpSiteConfigRef)
	return dhcpSiteConfigRefType
}

// DhcpSiteConfigRefSiteApsType
//
// DHCP Site selected APs. The content is effective when the siteMode is EnableOnMultipleAPs.
type DhcpSiteConfigRefSiteApsType struct {
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

func NewDhcpSiteConfigRefSiteApsType() *DhcpSiteConfigRefSiteApsType {
	dhcpSiteConfigRefSiteApsTypeType := new(DhcpSiteConfigRefSiteApsType)
	return dhcpSiteConfigRefSiteApsTypeType
}

func NewDhcpSiteConfigRefSiteApsTypeWithDefaults() *DhcpSiteConfigRefSiteApsType {
	dhcpSiteConfigRefSiteApsTypeType := new(DhcpSiteConfigRefSiteApsType)
	return dhcpSiteConfigRefSiteApsTypeType
}

type DoAssignIp struct {
	// DwpdEnabled
	// DHCP Service Dynamic WAN Port Detection
	DwpdEnabled *bool `json:"dwpdEnabled,omitempty"`

	// ManualSelect
	// DHCP Service AP Selection Mode. This value is effective when the siteMode is EnableOnMultipleAPs.
	ManualSelect *bool `json:"manualSelect,omitempty"`

	SiteAps []*DoAssignIpSiteApsType `json:"siteAps,omitempty"`

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

func NewDoAssignIp() *DoAssignIp {
	doAssignIpType := new(DoAssignIp)
	return doAssignIpType
}

func NewDoAssignIpWithDefaults() *DoAssignIp {
	doAssignIpType := new(DoAssignIp)
	return doAssignIpType
}

// DoAssignIpSiteApsType
//
// DHCP Site selected APs. The content is effective when the siteMode is EnableOnMultipleAPs.
type DoAssignIpSiteApsType struct {
	ApMac *string `json:"apMac,omitempty"`

	ApServerEnabled *bool `json:"apServerEnabled,omitempty"`

	ApServerPrimary *bool `json:"apServerPrimary,omitempty"`
}

func NewDoAssignIpSiteApsType() *DoAssignIpSiteApsType {
	doAssignIpSiteApsTypeType := new(DoAssignIpSiteApsType)
	return doAssignIpSiteApsTypeType
}

func NewDoAssignIpSiteApsTypeWithDefaults() *DoAssignIpSiteApsType {
	doAssignIpSiteApsTypeType := new(DoAssignIpSiteApsType)
	return doAssignIpSiteApsTypeType
}

type Email string

func NewEmail() *Email {
	emailType := new(Email)
	return emailType
}

func NewEmailWithDefaults() *Email {
	emailType := new(Email)
	return emailType
}

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

func NewEmptyResult() *EmptyResult {
	emptyResultType := new(EmptyResult)
	return emptyResultType
}

func NewEmptyResultWithDefaults() *EmptyResult {
	emptyResultType := new(EmptyResult)
	return emptyResultType
}

type FilterOperator string

func NewFilterOperator() *FilterOperator {
	filterOperatorType := new(FilterOperator)
	return filterOperatorType
}

func NewFilterOperatorWithDefaults() *FilterOperator {
	filterOperatorType := new(FilterOperator)
	return filterOperatorType
}

type FirmwareVersion string

func NewFirmwareVersion() *FirmwareVersion {
	firmwareVersionType := new(FirmwareVersion)
	return firmwareVersionType
}

func NewFirmwareVersionWithDefaults() *FirmwareVersion {
	firmwareVersionType := new(FirmwareVersion)
	return firmwareVersionType
}

type FQDN string

func NewFQDN() *FQDN {
	fQDNType := new(FQDN)
	return fQDNType
}

func NewFQDNWithDefaults() *FQDN {
	fQDNType := new(FQDN)
	return fQDNType
}

type FullTextSearch struct {
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

func NewFullTextSearch() *FullTextSearch {
	fullTextSearchType := new(FullTextSearch)
	return fullTextSearchType
}

func NewFullTextSearchWithDefaults() *FullTextSearch {
	fullTextSearchType := new(FullTextSearch)
	return fullTextSearchType
}

type GenericRef struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

func NewGenericRef() *GenericRef {
	genericRefType := new(GenericRef)
	return genericRefType
}

func NewGenericRefWithDefaults() *GenericRef {
	genericRefType := new(GenericRef)
	return genericRefType
}

type HealthCheckPolicy struct {
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

func NewHealthCheckPolicy() *HealthCheckPolicy {
	healthCheckPolicyType := new(HealthCheckPolicy)
	return healthCheckPolicyType
}

func NewHealthCheckPolicyWithDefaults() *HealthCheckPolicy {
	healthCheckPolicyType := new(HealthCheckPolicy)
	responseFailField := false
	healthCheckPolicyType.ResponseFail = &responseFailField
	responseWindowField := 20
	healthCheckPolicyType.ResponseWindow = &responseWindowField
	reviveIntervalField := 120
	healthCheckPolicyType.ReviveInterval = &reviveIntervalField
	zombiePeriodField := 40
	healthCheckPolicyType.ZombiePeriod = &zombiePeriodField
	return healthCheckPolicyType
}

type HTTPS string

func NewHTTPS() *HTTPS {
	hTTPSType := new(HTTPS)
	return hTTPSType
}

func NewHTTPSWithDefaults() *HTTPS {
	hTTPSType := new(HTTPS)
	return hTTPSType
}

type IdList []string

func NewIdList() *IdList {
	idListType := make(IdList, 0)
	return &idListType
}

func NewIdListWithDefaults() *IdList {
	idListType := make(IdList, 0)
	return &idListType
}

type IpAddress string

func NewIpAddress() *IpAddress {
	ipAddressType := new(IpAddress)
	return ipAddressType
}

func NewIpAddressWithDefaults() *IpAddress {
	ipAddressType := new(IpAddress)
	return ipAddressType
}

type IpMode string

func NewIpMode() *IpMode {
	ipModeType := new(IpMode)
	return ipModeType
}

func NewIpModeWithDefaults() *IpMode {
	ipModeType := new(IpMode)
	return ipModeType
}

type LanguageName string

func NewLanguageName() *LanguageName {
	languageNameType := new(LanguageName)
	return languageNameType
}

func NewLanguageNameWithDefaults() *LanguageName {
	languageNameType := new(LanguageName)
	return languageNameType
}

type Latitude float64

func NewLatitude() *Latitude {
	latitudeType := new(Latitude)
	return latitudeType
}

func NewLatitudeWithDefaults() *Latitude {
	latitudeType := new(Latitude)
	return latitudeType
}

type Location string

func NewLocation() *Location {
	locationType := new(Location)
	return locationType
}

func NewLocationWithDefaults() *Location {
	locationType := new(Location)
	return locationType
}

type LocationAdditionalInfo string

func NewLocationAdditionalInfo() *LocationAdditionalInfo {
	locationAdditionalInfoType := new(LocationAdditionalInfo)
	return locationAdditionalInfoType
}

func NewLocationAdditionalInfoWithDefaults() *LocationAdditionalInfo {
	locationAdditionalInfoType := new(LocationAdditionalInfo)
	return locationAdditionalInfoType
}

type Longitude float64

func NewLongitude() *Longitude {
	longitudeType := new(Longitude)
	return longitudeType
}

func NewLongitudeWithDefaults() *Longitude {
	longitudeType := new(Longitude)
	return longitudeType
}

type LteBandLockChannel struct {
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

func NewLteBandLockChannel() *LteBandLockChannel {
	lteBandLockChannelType := new(LteBandLockChannel)
	return lteBandLockChannelType
}

func NewLteBandLockChannelWithDefaults() *LteBandLockChannel {
	lteBandLockChannelType := new(LteBandLockChannel)
	return lteBandLockChannelType
}

type Mac string

func NewMac() *Mac {
	macType := new(Mac)
	return macType
}

func NewMacWithDefaults() *Mac {
	macType := new(Mac)
	return macType
}

type NormalName string

func NewNormalName() *NormalName {
	normalNameType := new(NormalName)
	return normalNameType
}

func NewNormalNameWithDefaults() *NormalName {
	normalNameType := new(NormalName)
	return normalNameType
}

type NormalName2to64 string

func NewNormalName2to64() *NormalName2to64 {
	normalName2to64Type := new(NormalName2to64)
	return normalName2to64Type
}

func NewNormalName2to64WithDefaults() *NormalName2to64 {
	normalName2to64Type := new(NormalName2to64)
	return normalName2to64Type
}

type NormalName2to128 string

func NewNormalName2to128() *NormalName2to128 {
	normalName2to128Type := new(NormalName2to128)
	return normalName2to128Type
}

func NewNormalName2to128WithDefaults() *NormalName2to128 {
	normalName2to128Type := new(NormalName2to128)
	return normalName2to128Type
}

type NormalNameAllowBlank string

func NewNormalNameAllowBlank() *NormalNameAllowBlank {
	normalNameAllowBlankType := new(NormalNameAllowBlank)
	return normalNameAllowBlankType
}

func NewNormalNameAllowBlankWithDefaults() *NormalNameAllowBlank {
	normalNameAllowBlankType := new(NormalNameAllowBlank)
	return normalNameAllowBlankType
}

type NormalURL string

func NewNormalURL() *NormalURL {
	normalURLType := new(NormalURL)
	return normalURLType
}

func NewNormalURLWithDefaults() *NormalURL {
	normalURLType := new(NormalURL)
	return normalURLType
}

type OverrideClientAdmissionControl struct {
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

func NewOverrideClientAdmissionControl() *OverrideClientAdmissionControl {
	overrideClientAdmissionControlType := new(OverrideClientAdmissionControl)
	return overrideClientAdmissionControlType
}

func NewOverrideClientAdmissionControlWithDefaults() *OverrideClientAdmissionControl {
	overrideClientAdmissionControlType := new(OverrideClientAdmissionControl)
	maxRadioLoadPercentField := 75
	overrideClientAdmissionControlType.MaxRadioLoadPercent = &maxRadioLoadPercentField
	minClientCountField := 10
	overrideClientAdmissionControlType.MinClientCount = &minClientCountField
	minClientThroughputMbpsField := 0.000000
	overrideClientAdmissionControlType.MinClientThroughputMbps = &minClientThroughputMbpsField
	return overrideClientAdmissionControlType
}

type OverrideGenericRef struct {
	Enabled *bool `json:"enabled,omitempty"`

	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

func NewOverrideGenericRef() *OverrideGenericRef {
	overrideGenericRefType := new(OverrideGenericRef)
	return overrideGenericRefType
}

func NewOverrideGenericRefWithDefaults() *OverrideGenericRef {
	overrideGenericRefType := new(OverrideGenericRef)
	return overrideGenericRefType
}

type OverrideSmartMonitor struct {
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

func NewOverrideSmartMonitor() *OverrideSmartMonitor {
	overrideSmartMonitorType := new(OverrideSmartMonitor)
	return overrideSmartMonitorType
}

func NewOverrideSmartMonitorWithDefaults() *OverrideSmartMonitor {
	overrideSmartMonitorType := new(OverrideSmartMonitor)
	intervalInSecField := 10
	overrideSmartMonitorType.IntervalInSec = &intervalInSecField
	retryThresholdField := 3
	overrideSmartMonitorType.RetryThreshold = &retryThresholdField
	return overrideSmartMonitorType
}

type PortalCustomization struct {
	// Language
	// Constraints:
	//    - required
	Language *PortalLanguage `json:"language" validate:"required"`

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

func NewPortalCustomization() *PortalCustomization {
	portalCustomizationType := new(PortalCustomization)
	return portalCustomizationType
}

func NewPortalCustomizationWithDefaults() *PortalCustomization {
	portalCustomizationType := new(PortalCustomization)
	termsAndConditionsRequiredField := false
	portalCustomizationType.TermsAndConditionsRequired = &termsAndConditionsRequiredField
	termsAndConditionsTextField := `Terms of Use

By accepting this agreement and accessing the wireless network, you acknowledge that you are of legal age, you have read and understood, and agree to be bound by this agreement.
(*) The wireless network service is provided by the property owners and is completely at their discretion. Your access to the network may be blocked, suspended, or terminated at any time for any reason.
(*) You agree not to use the wireless network for any purpose that is unlawful or otherwise prohibited and you are fully responsible for your use.
(*) The wireless network is provided "as is" without warranties of any kind, either expressed or implied.

This wireless network is powered by Ruckus Wireless.`
	portalCustomizationType.TermsAndConditionsText = &termsAndConditionsTextField
	return portalCustomizationType
}

// PortalLanguage
//
// Language
// Constraints:
//    - nullable
//    - default:'en_US'
//    - oneof:[en_US,zh_TW,zh_CN,nl_NL,fr_FR,de_DE,ja_JP,es_ES,se_SE,ar_SA,cz_CZ,da_DK,tr_TR,pt_BR]
type PortalLanguage string

func NewPortalLanguage() *PortalLanguage {
	portalLanguageType := new(PortalLanguage)
	return portalLanguageType
}

func NewPortalLanguageWithDefaults() *PortalLanguage {
	portalLanguageType := new(PortalLanguage)
	*portalLanguageType = `en_US`
	return portalLanguageType
}

type ProtectionMode string

func NewProtectionMode() *ProtectionMode {
	protectionModeType := new(ProtectionMode)
	return protectionModeType
}

func NewProtectionModeWithDefaults() *ProtectionMode {
	protectionModeType := new(ProtectionMode)
	*protectionModeType = `RTS_CTS`
	return protectionModeType
}

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
	// Constraints:
	//    - nullable
	//    - min:1
	Limit *int `json:"limit,omitempty" validate:"omitempty,gte=1"`

	// Options
	// Specified feature required information
	Options *QueryCriteriaOptionsType `json:"options,omitempty"`

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
	SortInfo *QueryCriteriaSortInfoType `json:"sortInfo,omitempty"`
}

func NewQueryCriteria() *QueryCriteria {
	queryCriteriaType := new(QueryCriteria)
	return queryCriteriaType
}

func NewQueryCriteriaWithDefaults() *QueryCriteria {
	queryCriteriaType := new(QueryCriteria)
	return queryCriteriaType
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

func NewQueryCriteriaExtraFiltersType() *QueryCriteriaExtraFiltersType {
	queryCriteriaExtraFiltersTypeType := new(QueryCriteriaExtraFiltersType)
	return queryCriteriaExtraFiltersTypeType
}

func NewQueryCriteriaExtraFiltersTypeWithDefaults() *QueryCriteriaExtraFiltersType {
	queryCriteriaExtraFiltersTypeType := new(QueryCriteriaExtraFiltersType)
	return queryCriteriaExtraFiltersTypeType
}

type QueryCriteriaExtraNotFiltersType struct {
	// Type
	// Filters for specific attribute
	Type *string `json:"type,omitempty"`

	// Value
	// value not to search
	Value *string `json:"value,omitempty"`
}

func NewQueryCriteriaExtraNotFiltersType() *QueryCriteriaExtraNotFiltersType {
	queryCriteriaExtraNotFiltersTypeType := new(QueryCriteriaExtraNotFiltersType)
	return queryCriteriaExtraNotFiltersTypeType
}

func NewQueryCriteriaExtraNotFiltersTypeWithDefaults() *QueryCriteriaExtraNotFiltersType {
	queryCriteriaExtraNotFiltersTypeType := new(QueryCriteriaExtraNotFiltersType)
	return queryCriteriaExtraNotFiltersTypeType
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

func NewQueryCriteriaFiltersType() *QueryCriteriaFiltersType {
	queryCriteriaFiltersTypeType := new(QueryCriteriaFiltersType)
	return queryCriteriaFiltersTypeType
}

func NewQueryCriteriaFiltersTypeWithDefaults() *QueryCriteriaFiltersType {
	queryCriteriaFiltersTypeType := new(QueryCriteriaFiltersType)
	return queryCriteriaFiltersTypeType
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

func NewQueryCriteriaOptionsType() *QueryCriteriaOptionsType {
	queryCriteriaOptionsTypeType := new(QueryCriteriaOptionsType)
	return queryCriteriaOptionsTypeType
}

func NewQueryCriteriaOptionsTypeWithDefaults() *QueryCriteriaOptionsType {
	queryCriteriaOptionsTypeType := new(QueryCriteriaOptionsType)
	return queryCriteriaOptionsTypeType
}

// QueryCriteriaSortInfoType
//
// About sorting
type QueryCriteriaSortInfoType struct {
	// Dir
	// Constraints:
	//    - nullable
	//    - oneof:[ASC,DESC]
	Dir *string `json:"dir,omitempty" validate:"omitempty,oneof=ASC DESC"`

	SortColumn *string `json:"sortColumn,omitempty"`
}

func NewQueryCriteriaSortInfoType() *QueryCriteriaSortInfoType {
	queryCriteriaSortInfoTypeType := new(QueryCriteriaSortInfoType)
	return queryCriteriaSortInfoTypeType
}

func NewQueryCriteriaSortInfoTypeWithDefaults() *QueryCriteriaSortInfoType {
	queryCriteriaSortInfoTypeType := new(QueryCriteriaSortInfoType)
	return queryCriteriaSortInfoTypeType
}

type QueryCriteriaSuperSet struct{}

func NewQueryCriteriaSuperSet() *QueryCriteriaSuperSet {
	queryCriteriaSuperSetType := new(QueryCriteriaSuperSet)
	return queryCriteriaSuperSetType
}

func NewQueryCriteriaSuperSetWithDefaults() *QueryCriteriaSuperSet {
	queryCriteriaSuperSetType := new(QueryCriteriaSuperSet)
	return queryCriteriaSuperSetType
}

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
	// Constraints:
	//    - nullable
	//    - default:0
	//    - oneof:[0,20,40]
	ChannelWidth *int `json:"channelWidth,omitempty" validate:"omitempty,oneof=0 20 40"`

	TxPower *TxPower `json:"txPower,omitempty"`
}

func NewRadio24() *Radio24 {
	radio24Type := new(Radio24)
	return radio24Type
}

func NewRadio24WithDefaults() *Radio24 {
	radio24Type := new(Radio24)
	channelWidthField := 0
	radio24Type.ChannelWidth = &channelWidthField
	return radio24Type
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
	// Constraints:
	//    - nullable
	//    - oneof:[0,20,40]
	ChannelWidth *int `json:"channelWidth,omitempty" validate:"omitempty,oneof=0 20 40"`

	TxPower *TxPower `json:"txPower,omitempty"`
}

func NewRadio24SuperSet() *Radio24SuperSet {
	radio24SuperSetType := new(Radio24SuperSet)
	return radio24SuperSetType
}

func NewRadio24SuperSetWithDefaults() *Radio24SuperSet {
	radio24SuperSetType := new(Radio24SuperSet)
	return radio24SuperSetType
}

type Radio50 struct {
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

	TxPower *TxPower `json:"txPower,omitempty"`
}

func NewRadio50() *Radio50 {
	radio50Type := new(Radio50)
	return radio50Type
}

func NewRadio50WithDefaults() *Radio50 {
	radio50Type := new(Radio50)
	channelWidthField := 0
	radio50Type.ChannelWidth = &channelWidthField
	return radio50Type
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

	TxPower *TxPower `json:"txPower,omitempty"`
}

func NewRadio50SuperSet() *Radio50SuperSet {
	radio50SuperSetType := new(Radio50SuperSet)
	return radio50SuperSetType
}

func NewRadio50SuperSetWithDefaults() *Radio50SuperSet {
	radio50SuperSetType := new(Radio50SuperSet)
	return radio50SuperSetType
}

type RadiusServer struct {
	// Ip
	// Constraints:
	//    - required
	Ip *IpAddress `json:"ip" validate:"required"`

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

func NewRadiusServer() *RadiusServer {
	radiusServerType := new(RadiusServer)
	return radiusServerType
}

func NewRadiusServerWithDefaults() *RadiusServer {
	radiusServerType := new(RadiusServer)
	return radiusServerType
}

type RateLimiting struct {
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

func NewRateLimiting() *RateLimiting {
	rateLimitingType := new(RateLimiting)
	return rateLimitingType
}

func NewRateLimitingWithDefaults() *RateLimiting {
	rateLimitingType := new(RateLimiting)
	maxOutstandingRequestsPerServerField := 0
	rateLimitingType.MaxOutstandingRequestsPerServer = &maxOutstandingRequestsPerServerField
	sanityTimerField := 10
	rateLimitingType.SanityTimer = &sanityTimerField
	thresholdField := 0
	rateLimitingType.Threshold = &thresholdField
	return rateLimitingType
}

type RbacMetadata struct {
	RbacMetadata []string `json:"rbacMetadata,omitempty"`
}

func NewRbacMetadata() *RbacMetadata {
	rbacMetadataType := new(RbacMetadata)
	return rbacMetadataType
}

func NewRbacMetadataWithDefaults() *RbacMetadata {
	rbacMetadataType := new(RbacMetadata)
	return rbacMetadataType
}

type Realm string

func NewRealm() *Realm {
	realmType := new(Realm)
	return realmType
}

func NewRealmWithDefaults() *Realm {
	realmType := new(Realm)
	return realmType
}

type RecoverySsid struct {
	// RecoverySsidEnabled
	// recovery ssid enable/disable
	RecoverySsidEnabled *bool `json:"recoverySsidEnabled,omitempty"`
}

func NewRecoverySsid() *RecoverySsid {
	recoverySsidType := new(RecoverySsid)
	return recoverySsidType
}

func NewRecoverySsidWithDefaults() *RecoverySsid {
	recoverySsidType := new(RecoverySsid)
	return recoverySsidType
}

type SmartMonitor struct {
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

func NewSmartMonitor() *SmartMonitor {
	smartMonitorType := new(SmartMonitor)
	return smartMonitorType
}

func NewSmartMonitorWithDefaults() *SmartMonitor {
	smartMonitorType := new(SmartMonitor)
	intervalInSecField := 10
	smartMonitorType.IntervalInSec = &intervalInSecField
	retryThresholdField := 3
	smartMonitorType.RetryThreshold = &retryThresholdField
	return smartMonitorType
}

type SnmpCommunity struct {
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
	NotificationTarget []*TargetConfig `json:"notificationTarget,omitempty"`

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

func NewSnmpCommunity() *SnmpCommunity {
	snmpCommunityType := new(SnmpCommunity)
	return snmpCommunityType
}

func NewSnmpCommunityWithDefaults() *SnmpCommunity {
	snmpCommunityType := new(SnmpCommunity)
	return snmpCommunityType
}

type SnmpUser struct {
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
	NotificationTarget []*TargetConfig `json:"notificationTarget,omitempty"`

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

func NewSnmpUser() *SnmpUser {
	snmpUserType := new(SnmpUser)
	return snmpUserType
}

func NewSnmpUserWithDefaults() *SnmpUser {
	snmpUserType := new(SnmpUser)
	return snmpUserType
}

type SubNetMask string

func NewSubNetMask() *SubNetMask {
	subNetMaskType := new(SubNetMask)
	return subNetMaskType
}

func NewSubNetMaskWithDefaults() *SubNetMask {
	subNetMaskType := new(SubNetMask)
	return subNetMaskType
}

type TargetConfig struct {
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

func NewTargetConfig() *TargetConfig {
	targetConfigType := new(TargetConfig)
	return targetConfigType
}

func NewTargetConfigWithDefaults() *TargetConfig {
	targetConfigType := new(TargetConfig)
	return targetConfigType
}

type TimeRange struct {
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

func NewTimeRange() *TimeRange {
	timeRangeType := new(TimeRange)
	return timeRangeType
}

func NewTimeRangeWithDefaults() *TimeRange {
	timeRangeType := new(TimeRange)
	return timeRangeType
}

// TimeUnitStore
//
// time unit
// Constraints:
//    - nullable
//    - oneof:[second,minute,hour,day]
type TimeUnitStore string

func NewTimeUnitStore() *TimeUnitStore {
	timeUnitStoreType := new(TimeUnitStore)
	return timeUnitStoreType
}

func NewTimeUnitStoreWithDefaults() *TimeUnitStore {
	timeUnitStoreType := new(TimeUnitStore)
	return timeUnitStoreType
}

type TrafficClassProfileRef struct {
	Description *Description `json:"description,omitempty"`

	// Id
	// Identifier of the Traffic Class Profile
	Id *string `json:"id,omitempty"`

	Name *NormalName `json:"name,omitempty"`

	TrafficClasses []*TrafficClassRef `json:"trafficClasses,omitempty"`

	// ZoneId
	// Zone Id of Traffic Class Profile
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewTrafficClassProfileRef() *TrafficClassProfileRef {
	trafficClassProfileRefType := new(TrafficClassProfileRef)
	return trafficClassProfileRefType
}

func NewTrafficClassProfileRefWithDefaults() *TrafficClassProfileRef {
	trafficClassProfileRefType := new(TrafficClassProfileRef)
	return trafficClassProfileRefType
}

type TrafficClassRef struct {
	// Id
	// Identifier of the Traffic Class
	Id *string `json:"id,omitempty"`

	// Whitelists
	// White list of the Traffic Class Profile. The multiple entries need to be separated by comma (,)
	Whitelists *string `json:"whitelists,omitempty"`
}

func NewTrafficClassRef() *TrafficClassRef {
	trafficClassRefType := new(TrafficClassRef)
	return trafficClassRefType
}

func NewTrafficClassRefWithDefaults() *TrafficClassRef {
	trafficClassRefType := new(TrafficClassRef)
	return trafficClassRefType
}

type TxPower string

func NewTxPower() *TxPower {
	txPowerType := new(TxPower)
	return txPowerType
}

func NewTxPowerWithDefaults() *TxPower {
	txPowerType := new(TxPower)
	*txPowerType = `Full`
	return txPowerType
}

type WebAuthenticationPortalCustomization struct {
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

func NewWebAuthenticationPortalCustomization() *WebAuthenticationPortalCustomization {
	webAuthenticationPortalCustomizationType := new(WebAuthenticationPortalCustomization)
	return webAuthenticationPortalCustomizationType
}

func NewWebAuthenticationPortalCustomizationWithDefaults() *WebAuthenticationPortalCustomization {
	webAuthenticationPortalCustomizationType := new(WebAuthenticationPortalCustomization)
	return webAuthenticationPortalCustomizationType
}

// WildFQDN
//
// Compare with FQDN, it could start with '*.'
type WildFQDN string

func NewWildFQDN() *WildFQDN {
	wildFQDNType := new(WildFQDN)
	return wildFQDNType
}

func NewWildFQDNWithDefaults() *WildFQDN {
	wildFQDNType := new(WildFQDN)
	return wildFQDNType
}

// ZoneTunnelType
//
// Tunnel type configuration of the zone. No_Tunneled is for IPv6 mode
// Constraints:
//    - nullable
//    - default:'RuckusGRE'
//    - oneof:[No_Tunneled,RuckusGRE,SoftGRE,SoftGREIpsec]
type ZoneTunnelType string

func NewZoneTunnelType() *ZoneTunnelType {
	zoneTunnelTypeType := new(ZoneTunnelType)
	return zoneTunnelTypeType
}

func NewZoneTunnelTypeWithDefaults() *ZoneTunnelType {
	zoneTunnelTypeType := new(ZoneTunnelType)
	*zoneTunnelTypeType = `RuckusGRE`
	return zoneTunnelTypeType
}
