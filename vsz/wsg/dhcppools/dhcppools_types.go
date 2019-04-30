package dhcppools

// API Version: v8_0

// DHCPClientInfo
//
// DHCP Pool Client Information List
type DHCPClientInfo struct {
	ClientIP *string `json:"clientIp,omitempty"`

	ClientMac *string `json:"clientMac,omitempty"`

	LeaseExpiryTime *int `json:"leaseExpiryTime,omitempty"`

	LeaseTime *int `json:"leaseTime,omitempty"`

	LeaseTimeHours *int `json:"leaseTimeHours,omitempty"`

	LeaseTimeMinutes *int `json:"leaseTimeMinutes,omitempty"`
}

// DHCPPoolInfo
//
// DHCP Pool Information List
type DHCPPoolInfo struct {
	ApIP *string `json:"apIp,omitempty"`

	AvailableIPCount *int `json:"availableIpCount,omitempty"`

	ClientInfoList *DHCPPoolInfoClientInfoListType `json:"clientInfoList,omitempty"`

	Name *string `json:"name,omitempty"`

	PoolEndIP *string `json:"poolEndIp,omitempty"`

	PoolIndex *int `json:"poolIndex,omitempty"`

	PoolStartIP *string `json:"poolStartIp,omitempty"`

	SubnetMask *string `json:"subnetMask,omitempty"`

	TotalIPCount *int `json:"totalIpCount,omitempty"`

	UsedIPCount *int `json:"usedIpCount,omitempty"`

	VlanID *int `json:"vlanId,omitempty"`
}

type DHCPPoolInfoClientInfoListType struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*DHCPClientInfo `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

// DHCPPools
//
// DHCP Pools Usage Per AP
type DHCPPools struct {
	ApMac *string `json:"apMac,omitempty"`

	DomainEntity *string `json:"domainEntity,omitempty"`

	DomainID *string `json:"domainId,omitempty"`

	ID *string `json:"id,omitempty"`

	PoolInfoList []*DHCPPoolInfo `json:"poolInfoList,omitempty"`

	TenantID *string `json:"tenantId,omitempty"`
}