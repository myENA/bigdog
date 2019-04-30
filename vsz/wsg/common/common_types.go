package common

// API Version: v8_0

type Alarm struct {
	AcknowledgedTime *int    `json:"acknowledgedTime,omitempty"`
	Category         *string `json:"category,omitempty"`
	ClearedTime      *int    `json:"clearedTime,omitempty"`
	Code             *string `json:"code,omitempty"`
	Description      *string `json:"description,omitempty"`
	ID               *string `json:"id,omitempty"`
	Severity         *string `json:"severity,omitempty"`
	Status           *string `json:"status,omitempty"`
	Time             *string `json:"time,omitempty"`
	Type             *string `json:"type,omitempty"`
}

type Altitude struct {
	AltitudeUnit  *string `json:"altitudeUnit,omitempty"`
	AltitudeValue *int    `json:"altitudeValue,omitempty"`
}

type ApLatencyInterval struct {
	PingEnabled *bool `json:"pingEnabled,omitempty"`
}

type ApManagementVlan struct {
	ID   *int    `json:"id,omitempty"`
	Mode *string `json:"mode,omitempty"`
}

type ApRadio50 struct {
	AutoCellSizing   *bool   `json:"autoCellSizing,omitempty"`
	Channel          *int    `json:"channel,omitempty"`
	ChannelRange     []int   `json:"channelRange,omitempty"`
	ChannelWidth     *int    `json:"channelWidth,omitempty"`
	SecondaryChannel *int    `json:"secondaryChannel,omitempty"`
	TxPower          *string `json:"txPower,omitempty"`
}

type ApRebootTimeout struct {
	GatewayLossTimeoutInSec *int `json:"gatewayLossTimeoutInSec,omitempty"`
	ServerLossTimeoutInSec  *int `json:"serverLossTimeoutInSec,omitempty"`
}

type AutoChannelSelection struct {
	ChannelFlyMtbc    *int    `json:"channelFlyMtbc,omitempty"`
	ChannelSelectMode *string `json:"channelSelectMode,omitempty"`
}

type BaseServiceInfo struct {
	ID          *string `json:"id,omitempty"`
	ServiceID   *string `json:"serviceId,omitempty"`
	ServiceName *string `json:"serviceName,omitempty"`
	ServiceType *string `json:"serviceType,omitempty"`
}

type BulkDeleteRequest struct {
	IDList IDList `json:"idList,omitempty"`
}

type Client struct {
	ApTxDataRate        *string `json:"apTxDataRate,omitempty"`
	Channel             *string `json:"channel,omitempty"`
	ConnectedSince      *int    `json:"connectedSince,omitempty"`
	FromClientBytes     *int    `json:"fromClientBytes,omitempty"`
	FromClientPkts      *int    `json:"fromClientPkts,omitempty"`
	HostName            *string `json:"hostName,omitempty"`
	IPAddress           *string `json:"ipAddress,omitempty"`
	Ipv6Address         *string `json:"ipv6Address,omitempty"`
	Mac                 *string `json:"mac,omitempty"`
	OsType              *string `json:"osType,omitempty"`
	RadioID             *string `json:"radioId,omitempty"`
	RadioMode           *string `json:"radioMode,omitempty"`
	Rssi                *string `json:"rssi,omitempty"`
	RxAvgByteRate       *int    `json:"rxAvgByteRate,omitempty"`
	RxByteRate          *int    `json:"rxByteRate,omitempty"`
	Snr                 *string `json:"snr,omitempty"`
	Ssid                *string `json:"ssid,omitempty"`
	Status              *string `json:"status,omitempty"`
	ToClientBytes       *int    `json:"toClientBytes,omitempty"`
	ToClientDroppedPkts *int    `json:"toClientDroppedPkts,omitempty"`
	ToClientPkts        *int    `json:"toClientPkts,omitempty"`
	TxAvgByteRate       *int    `json:"txAvgByteRate,omitempty"`
	TxByteRate          *int    `json:"txByteRate,omitempty"`
	User                *string `json:"user,omitempty"`
	Vlan                *string `json:"vlan,omitempty"`
	WLANID              *string `json:"wlanId,omitempty"`
}

type ClientAdmissionControl struct {
	MaxRadioLoadPercent     *int     `json:"maxRadioLoadPercent,omitempty"`
	MinClientCount          *int     `json:"minClientCount,omitempty"`
	MinClientThroughputMbps *float64 `json:"minClientThroughputMbps,omitempty"`
}

type CreateResult struct {
	ID *string `json:"id,omitempty"`
}

type CreateResultIDName struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type DHCPProfileRef struct {
	Description      *string `json:"description,omitempty"`
	ID               *string `json:"id,omitempty"`
	LeaseTimeHours   *int    `json:"leaseTimeHours,omitempty"`
	LeaseTimeMinutes *int    `json:"leaseTimeMinutes,omitempty"`
	Name             *string `json:"name,omitempty"`
	PoolEndIP        *string `json:"poolEndIp,omitempty"`
	PoolStartIP      *string `json:"poolStartIp,omitempty"`
	PrimaryDNSIP     *string `json:"primaryDnsIp,omitempty"`
	SecondaryDNSIP   *string `json:"secondaryDnsIp,omitempty"`
	SubnetMask       *string `json:"subnetMask,omitempty"`
	SubnetNetworkIP  *string `json:"subnetNetworkIp,omitempty"`
	VlanID           *int    `json:"vlanId,omitempty"`
	ZoneID           *string `json:"zoneId,omitempty"`
}

type DHCPSiteConfigListRef struct {
	DwpdEnabled   *bool                               `json:"dwpdEnabled,omitempty"`
	Eth0ProfileID *int                                `json:"eth0ProfileId,omitempty"`
	Eth1ProfileID *int                                `json:"eth1ProfileId,omitempty"`
	ManualSelect  *bool                               `json:"manualSelect,omitempty"`
	SiteAps       []*DHCPSiteConfigListRefSiteApsType `json:"siteAps,omitempty"`
	SiteEnabled   *bool                               `json:"siteEnabled,omitempty"`
	SiteMode      *string                             `json:"siteMode,omitempty"`
	SiteProfiles  []*DHCPProfileRef                   `json:"siteProfiles,omitempty"`
	ZoneName      *string                             `json:"zoneName,omitempty"`
}

type DHCPSiteConfigListRefSiteApsType struct {
	ApGatewayIP     *string `json:"apGatewayIp,omitempty"`
	ApMac           *string `json:"apMac,omitempty"`
	ApName          *string `json:"apName,omitempty"`
	ApServerEnabled *bool   `json:"apServerEnabled,omitempty"`
	ApServerIP      *string `json:"apServerIp,omitempty"`
	ApServerPrimary *bool   `json:"apServerPrimary,omitempty"`
	ApStatus        *string `json:"apStatus,omitempty"`
}

type DHCPSiteConfigRef struct {
	DwpdEnabled    *bool                           `json:"dwpdEnabled,omitempty"`
	Eth0ProfileID  *int                            `json:"eth0ProfileId,omitempty"`
	Eth1ProfileID  *int                            `json:"eth1ProfileId,omitempty"`
	ManualSelect   *bool                           `json:"manualSelect,omitempty"`
	SiteAps        []*DHCPSiteConfigRefSiteApsType `json:"siteAps,omitempty"`
	SiteEnabled    *bool                           `json:"siteEnabled,omitempty"`
	SiteMode       *string                         `json:"siteMode,omitempty"`
	SiteProfileIds []string                        `json:"siteProfileIds,omitempty"`
}

type DHCPSiteConfigRefSiteApsType struct {
	ApGatewayIP     *string `json:"apGatewayIp,omitempty"`
	ApMac           *string `json:"apMac,omitempty"`
	ApName          *string `json:"apName,omitempty"`
	ApServerEnabled *bool   `json:"apServerEnabled,omitempty"`
	ApServerIP      *string `json:"apServerIp,omitempty"`
	ApServerPrimary *bool   `json:"apServerPrimary,omitempty"`
	ApStatus        *string `json:"apStatus,omitempty"`
}

type DoAssignIP struct {
	DwpdEnabled    *bool                    `json:"dwpdEnabled,omitempty"`
	ManualSelect   *bool                    `json:"manualSelect,omitempty"`
	SiteAps        []*DoAssignIPSiteApsType `json:"siteAps,omitempty"`
	SiteEnabled    *bool                    `json:"siteEnabled,omitempty"`
	SiteMode       *string                  `json:"siteMode,omitempty"`
	SiteProfileIds []string                 `json:"siteProfileIds,omitempty"`
}

type DoAssignIPSiteApsType struct {
	ApMac           *string `json:"apMac,omitempty"`
	ApServerEnabled *bool   `json:"apServerEnabled,omitempty"`
	ApServerPrimary *bool   `json:"apServerPrimary,omitempty"`
}

type FullTextSearch struct {
	Fields []string `json:"fields,omitempty"`
	Type   *string  `json:"type,omitempty"`
	Value  *string  `json:"value,omitempty"`
}

type GenericRef struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type HealthCheckPolicy struct {
	ResponseFail   *bool `json:"responseFail,omitempty"`
	ResponseWindow *int  `json:"responseWindow,omitempty"`
	ReviveInterval *int  `json:"reviveInterval,omitempty"`
	ZombiePeriod   *int  `json:"zombiePeriod,omitempty"`
}

type IDList []string

type OverrideClientAdmissionControl struct {
	Enabled                 *bool    `json:"enabled,omitempty"`
	MaxRadioLoadPercent     *int     `json:"maxRadioLoadPercent,omitempty"`
	MinClientCount          *int     `json:"minClientCount,omitempty"`
	MinClientThroughputMbps *float64 `json:"minClientThroughputMbps,omitempty"`
}

type OverrideGenericRef struct {
	Enabled *bool   `json:"enabled,omitempty"`
	ID      *string `json:"id,omitempty"`
	Name    *string `json:"name,omitempty"`
}

type OverrideSmartMonitor struct {
	Enabled        *bool `json:"enabled,omitempty"`
	IntervalInSec  *int  `json:"intervalInSec,omitempty"`
	RetryThreshold *int  `json:"retryThreshold,omitempty"`
}

type PortalCustomization struct {
	Language                   *string `json:"language,omitempty"`
	Logo                       *string `json:"logo,omitempty"`
	TermsAndConditionsRequired *bool   `json:"termsAndConditionsRequired,omitempty"`
	TermsAndConditionsText     *string `json:"termsAndConditionsText,omitempty"`
	Title                      *string `json:"title,omitempty"`
}

type QueryCriteria struct {
	Attributes      []string                            `json:"attributes,omitempty"`
	Criteria        *string                             `json:"criteria,omitempty"`
	ExpandDomains   *bool                               `json:"expandDomains,omitempty"`
	ExtraFilters    []*QueryCriteriaExtraFiltersType    `json:"extraFilters,omitempty"`
	ExtraNotFilters []*QueryCriteriaExtraNotFiltersType `json:"extraNotFilters,omitempty"`
	ExtraTimeRange  *TimeRange                          `json:"extraTimeRange,omitempty"`
	Filters         []*QueryCriteriaFiltersType         `json:"filters,omitempty"`
	FullTextSearch  *FullTextSearch                     `json:"fullTextSearch,omitempty"`
	Limit           *int                                `json:"limit,omitempty"`
	Options         map[string]interface{}              `json:"options,omitempty"`
	Page            *int                                `json:"page,omitempty"`
	Query           *string                             `json:"query,omitempty"`
	SortInfo        *QueryCriteriaSortInfoType          `json:"sortInfo,omitempty"`
}

type QueryCriteriaExtraFiltersType struct {
	Operator *string `json:"operator,omitempty"`
	Type     *string `json:"type,omitempty"`
	Value    *string `json:"value,omitempty"`
}

type QueryCriteriaExtraNotFiltersType struct {
	Type  *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

type QueryCriteriaFiltersType struct {
	Operator *string `json:"operator,omitempty"`
	Type     *string `json:"type,omitempty"`
	Value    *string `json:"value,omitempty"`
}

type QueryCriteriaSortInfoType struct {
	Dir        *string `json:"dir,omitempty"`
	SortColumn *string `json:"sortColumn,omitempty"`
}

type QueryCriteriaSuperSet struct {
}

type Radio24 struct {
	AutoCellSizing *bool   `json:"autoCellSizing,omitempty"`
	Channel        *int    `json:"channel,omitempty"`
	ChannelRange   []int   `json:"channelRange,omitempty"`
	ChannelWidth   *int    `json:"channelWidth,omitempty"`
	TxPower        *string `json:"txPower,omitempty"`
}

type Radio24SuperSet struct {
	AutoCellSizing        *bool   `json:"autoCellSizing,omitempty"`
	AvailableChannelRange []int   `json:"availableChannelRange,omitempty"`
	Channel               *int    `json:"channel,omitempty"`
	ChannelRange          []int   `json:"channelRange,omitempty"`
	ChannelWidth          *int    `json:"channelWidth,omitempty"`
	TxPower               *string `json:"txPower,omitempty"`
}

type Radio50 struct {
	AutoCellSizing          *bool   `json:"autoCellSizing,omitempty"`
	ChannelWidth            *int    `json:"channelWidth,omitempty"`
	IndoorChannel           *int    `json:"indoorChannel,omitempty"`
	IndoorChannelRange      []int   `json:"indoorChannelRange,omitempty"`
	IndoorSecondaryChannel  *int    `json:"indoorSecondaryChannel,omitempty"`
	OutdoorChannel          *int    `json:"outdoorChannel,omitempty"`
	OutdoorChannelRange     []int   `json:"outdoorChannelRange,omitempty"`
	OutdoorSecondaryChannel *int    `json:"outdoorSecondaryChannel,omitempty"`
	TxPower                 *string `json:"txPower,omitempty"`
}

type Radio50SuperSet struct {
	AutoCellSizing               *bool   `json:"autoCellSizing,omitempty"`
	AvailableIndoorChannelRange  []int   `json:"availableIndoorChannelRange,omitempty"`
	AvailableOutdoorChannelRange []int   `json:"availableOutdoorChannelRange,omitempty"`
	ChannelWidth                 *int    `json:"channelWidth,omitempty"`
	IndoorChannel                *int    `json:"indoorChannel,omitempty"`
	IndoorChannelRange           []int   `json:"indoorChannelRange,omitempty"`
	IndoorSecondaryChannel       *int    `json:"indoorSecondaryChannel,omitempty"`
	OutdoorChannel               *int    `json:"outdoorChannel,omitempty"`
	OutdoorChannelRange          []int   `json:"outdoorChannelRange,omitempty"`
	OutdoorSecondaryChannel      *int    `json:"outdoorSecondaryChannel,omitempty"`
	TxPower                      *string `json:"txPower,omitempty"`
}

type RadiusServer struct {
	IP           *string `json:"ip,omitempty"`
	Port         *int    `json:"port,omitempty"`
	SharedSecret *string `json:"sharedSecret,omitempty"`
}

type RateLimiting struct {
	MaxOutstandingRequestsPerServer *int `json:"maxOutstandingRequestsPerServer,omitempty"`
	SanityTimer                     *int `json:"sanityTimer,omitempty"`
	Threshold                       *int `json:"threshold,omitempty"`
}

type RBACMetadata struct {
	RBACMetadata []string `json:"rbacMetadata,omitempty"`
}

type RecoverySsid struct {
	RecoverySsidEnabled *bool `json:"recoverySsidEnabled,omitempty"`
}

type SmartMonitor struct {
	IntervalInSec  *int `json:"intervalInSec,omitempty"`
	RetryThreshold *int `json:"retryThreshold,omitempty"`
}

type SNMPCommunity struct {
	CommunityName       *string         `json:"communityName,omitempty"`
	NotificationEnabled *bool           `json:"notificationEnabled,omitempty"`
	NotificationTarget  []*TargetConfig `json:"notificationTarget,omitempty"`
	NotificationType    *string         `json:"notificationType,omitempty"`
	ReadEnabled         *bool           `json:"readEnabled,omitempty"`
	WriteEnabled        *bool           `json:"writeEnabled,omitempty"`
}

type SNMPUser struct {
	AuthPassword        *string         `json:"authPassword,omitempty"`
	AuthProtocol        *string         `json:"authProtocol,omitempty"`
	NotificationEnabled *bool           `json:"notificationEnabled,omitempty"`
	NotificationTarget  []*TargetConfig `json:"notificationTarget,omitempty"`
	NotificationType    *string         `json:"notificationType,omitempty"`
	PrivPassword        *string         `json:"privPassword,omitempty"`
	PrivProtocol        *string         `json:"privProtocol,omitempty"`
	ReadEnabled         *bool           `json:"readEnabled,omitempty"`
	UserName            *string         `json:"userName,omitempty"`
	WriteEnabled        *bool           `json:"writeEnabled,omitempty"`
}

type TargetConfig struct {
	Address *string `json:"address,omitempty"`
	Port    *int    `json:"port,omitempty"`
}

type TimeRange struct {
	End      *float64 `json:"end,omitempty"`
	Field    *string  `json:"field,omitempty"`
	Interval *float64 `json:"interval,omitempty"`
	Start    *float64 `json:"start,omitempty"`
}

type TrafficClassProfileRef struct {
	Description    *string            `json:"description,omitempty"`
	ID             *string            `json:"id,omitempty"`
	Name           *string            `json:"name,omitempty"`
	TrafficClasses []*TrafficClassRef `json:"trafficClasses,omitempty"`
	ZoneID         *string            `json:"zoneId,omitempty"`
}

type TrafficClassRef struct {
	ID         *string `json:"id,omitempty"`
	Whitelists *string `json:"whitelists,omitempty"`
}
