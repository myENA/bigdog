package dhcppools

// API Version: v8_1

// DhcpClientInfo
//
// DHCP Pool Client Information List
type DhcpClientInfo struct {
	ClientIp *string `json:"clientIp,omitempty"`

	ClientMac *string `json:"clientMac,omitempty"`

	LeaseExpiryTime *int `json:"leaseExpiryTime,omitempty"`

	LeaseTime *int `json:"leaseTime,omitempty"`

	LeaseTimeHours *int `json:"leaseTimeHours,omitempty"`

	LeaseTimeMinutes *int `json:"leaseTimeMinutes,omitempty"`
}

func NewDhcpClientInfo() *DhcpClientInfo {
	dhcpClientInfoType := new(DhcpClientInfo)
	return dhcpClientInfoType
}

func NewDhcpClientInfoWithDefaults() *DhcpClientInfo {
	dhcpClientInfoType := new(DhcpClientInfo)
	return dhcpClientInfoType
}

// DhcpPoolInfo
//
// DHCP Pool Information List
type DhcpPoolInfo struct {
	ApIp *string `json:"apIp,omitempty"`

	AvailableIpCount *int `json:"availableIpCount,omitempty"`

	ClientInfoList *DhcpPoolInfoClientInfoListType `json:"clientInfoList,omitempty"`

	Name *string `json:"name,omitempty"`

	PoolEndIp *string `json:"poolEndIp,omitempty"`

	PoolIndex *int `json:"poolIndex,omitempty"`

	PoolStartIp *string `json:"poolStartIp,omitempty"`

	SubnetMask *string `json:"subnetMask,omitempty"`

	TotalIpCount *int `json:"totalIpCount,omitempty"`

	UsedIpCount *int `json:"usedIpCount,omitempty"`

	VlanId *int `json:"vlanId,omitempty"`
}

func NewDhcpPoolInfo() *DhcpPoolInfo {
	dhcpPoolInfoType := new(DhcpPoolInfo)
	return dhcpPoolInfoType
}

func NewDhcpPoolInfoWithDefaults() *DhcpPoolInfo {
	dhcpPoolInfoType := new(DhcpPoolInfo)
	return dhcpPoolInfoType
}

type DhcpPoolInfoClientInfoListType struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*DhcpClientInfo `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewDhcpPoolInfoClientInfoListType() *DhcpPoolInfoClientInfoListType {
	dhcpPoolInfoClientInfoListTypeType := new(DhcpPoolInfoClientInfoListType)
	return dhcpPoolInfoClientInfoListTypeType
}

func NewDhcpPoolInfoClientInfoListTypeWithDefaults() *DhcpPoolInfoClientInfoListType {
	dhcpPoolInfoClientInfoListTypeType := new(DhcpPoolInfoClientInfoListType)
	return dhcpPoolInfoClientInfoListTypeType
}

// DhcpPools
//
// DHCP Pools Usage Per AP
type DhcpPools struct {
	ApMac *string `json:"apMac,omitempty"`

	DomainEntity *string `json:"domainEntity,omitempty"`

	DomainId *string `json:"domainId,omitempty"`

	Id *string `json:"id,omitempty"`

	PoolInfoList []*DhcpPoolInfo `json:"poolInfoList,omitempty"`

	TenantId *string `json:"tenantId,omitempty"`
}

func NewDhcpPools() *DhcpPools {
	dhcpPoolsType := new(DhcpPools)
	return dhcpPoolsType
}

func NewDhcpPoolsWithDefaults() *DhcpPools {
	dhcpPoolsType := new(DhcpPools)
	return dhcpPoolsType
}
