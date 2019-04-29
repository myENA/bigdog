package dhcppools

// API Version: v8_0

type DHCPClientInfo struct {
	ClientIP         *string `json:"clientIp,omitempty"`
	ClientMac        *string `json:"clientMac,omitempty"`
	LeaseExpiryTime  *int    `json:"leaseExpiryTime,omitempty"`
	LeaseTime        *int    `json:"leaseTime,omitempty"`
	LeaseTimeHours   *int    `json:"leaseTimeHours,omitempty"`
	LeaseTimeMinutes *int    `json:"leaseTimeMinutes,omitempty"`
}

type DHCPPoolInfo struct {
	ApIP             *string         `json:"apIp,omitempty"`
	AvailableIPCount *int            `json:"availableIpCount,omitempty"`
	ClientInfoList   *ClientInfoList `json:"clientInfoList,omitempty"`
	Name             *string         `json:"name,omitempty"`
	PoolEndIP        *string         `json:"poolEndIp,omitempty"`
	PoolIndex        *int            `json:"poolIndex,omitempty"`
	PoolStartIP      *string         `json:"poolStartIp,omitempty"`
	SubnetMask       *string         `json:"subnetMask,omitempty"`
	TotalIPCount     *int            `json:"totalIpCount,omitempty"`
	UsedIPCount      *int            `json:"usedIpCount,omitempty"`
	VlanID           *int            `json:"vlanId,omitempty"`
}

type DHCPPoolInfoClientInfoListType struct {
	FirstIndex *int    `json:"firstIndex,omitempty"`
	HasMore    *bool   `json:"hasMore,omitempty"`
	List       []*List `json:"list,omitempty"`
	TotalCount *int    `json:"totalCount,omitempty"`
}

type DHCPPools struct {
	ApMac        *string         `json:"apMac,omitempty"`
	DomainEntity *string         `json:"domainEntity,omitempty"`
	DomainID     *string         `json:"domainId,omitempty"`
	ID           *string         `json:"id,omitempty"`
	PoolInfoList []*PoolInfoList `json:"poolInfoList,omitempty"`
	TenantID     *string         `json:"tenantId,omitempty"`
}
