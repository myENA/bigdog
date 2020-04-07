package vsz

// API Version: v9_0

// WSGDHCPPoolsDhcpClientInfo
//
// DHCP Pool Client Information List
// Constraints:
//    - nullable
type WSGDHCPPoolsDhcpClientInfo struct {
	// ClientIp
	// Constraints:
	//    - nullable
	ClientIp *string `json:"clientIp,omitempty"`

	// ClientMac
	// Constraints:
	//    - nullable
	ClientMac *string `json:"clientMac,omitempty"`

	// LeaseExpiryTime
	// Constraints:
	//    - nullable
	LeaseExpiryTime *int `json:"leaseExpiryTime,omitempty"`

	// LeaseTime
	// Constraints:
	//    - nullable
	LeaseTime *int `json:"leaseTime,omitempty"`

	// LeaseTimeHours
	// Constraints:
	//    - nullable
	LeaseTimeHours *int `json:"leaseTimeHours,omitempty"`

	// LeaseTimeMinutes
	// Constraints:
	//    - nullable
	LeaseTimeMinutes *int `json:"leaseTimeMinutes,omitempty"`
}

func NewWSGDHCPPoolsDhcpClientInfo() *WSGDHCPPoolsDhcpClientInfo {
	m := new(WSGDHCPPoolsDhcpClientInfo)
	return m
}

// WSGDHCPPoolsDhcpPoolInfo
//
// DHCP Pool Information List
// Constraints:
//    - nullable
type WSGDHCPPoolsDhcpPoolInfo struct {
	// ApIp
	// Constraints:
	//    - nullable
	ApIp *string `json:"apIp,omitempty"`

	// AvailableIpCount
	// Constraints:
	//    - nullable
	AvailableIpCount *int `json:"availableIpCount,omitempty"`

	// ClientInfoList
	// Constraints:
	//    - nullable
	ClientInfoList *WSGDHCPPoolsDhcpPoolInfoClientInfoListType `json:"clientInfoList,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// PoolEndIp
	// Constraints:
	//    - nullable
	PoolEndIp *string `json:"poolEndIp,omitempty"`

	// PoolIndex
	// Constraints:
	//    - nullable
	PoolIndex *int `json:"poolIndex,omitempty"`

	// PoolStartIp
	// Constraints:
	//    - nullable
	PoolStartIp *string `json:"poolStartIp,omitempty"`

	// SubnetMask
	// Constraints:
	//    - nullable
	SubnetMask *string `json:"subnetMask,omitempty"`

	// TotalIpCount
	// Constraints:
	//    - nullable
	TotalIpCount *int `json:"totalIpCount,omitempty"`

	// UsedIpCount
	// Constraints:
	//    - nullable
	UsedIpCount *int `json:"usedIpCount,omitempty"`

	// VlanId
	// Constraints:
	//    - nullable
	VlanId *int `json:"vlanId,omitempty"`
}

func NewWSGDHCPPoolsDhcpPoolInfo() *WSGDHCPPoolsDhcpPoolInfo {
	m := new(WSGDHCPPoolsDhcpPoolInfo)
	return m
}

type WSGDHCPPoolsDhcpPoolInfoClientInfoListType struct {
	// FirstIndex
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGDHCPPoolsDhcpClientInfo `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGDHCPPoolsDhcpPoolInfoClientInfoListType() *WSGDHCPPoolsDhcpPoolInfoClientInfoListType {
	m := new(WSGDHCPPoolsDhcpPoolInfoClientInfoListType)
	return m
}

// WSGDHCPPools
//
// DHCP Pools Usage Per AP
// Constraints:
//    - nullable
type WSGDHCPPools struct {
	// ApMac
	// Constraints:
	//    - nullable
	ApMac *string `json:"apMac,omitempty"`

	// DomainEntity
	// Constraints:
	//    - nullable
	DomainEntity *string `json:"domainEntity,omitempty"`

	// DomainId
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// PoolInfoList
	// Constraints:
	//    - nullable
	PoolInfoList []*WSGDHCPPoolsDhcpPoolInfo `json:"poolInfoList,omitempty" validate:"omitempty,dive"`

	// TenantId
	// Constraints:
	//    - nullable
	TenantId *string `json:"tenantId,omitempty"`
}

func NewWSGDHCPPools() *WSGDHCPPools {
	m := new(WSGDHCPPools)
	return m
}
