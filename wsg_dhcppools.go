package bigdog

// API Version: v9_1

import (
	"errors"
	"io"
)

// WSGDHCPPoolsDhcpClientInfo
//
// Definition: dhcppools_dhcpClientInfo
//
// DHCP Pool Client Information List
type WSGDHCPPoolsDhcpClientInfo struct {
	ClientIp *string `json:"clientIp,omitempty"`

	ClientMac *string `json:"clientMac,omitempty"`

	LeaseExpiryTime *int `json:"leaseExpiryTime,omitempty"`

	LeaseTime *int `json:"leaseTime,omitempty"`

	LeaseTimeHours *int `json:"leaseTimeHours,omitempty"`

	LeaseTimeMinutes *int `json:"leaseTimeMinutes,omitempty"`
}

func NewWSGDHCPPoolsDhcpClientInfo() *WSGDHCPPoolsDhcpClientInfo {
	m := new(WSGDHCPPoolsDhcpClientInfo)
	return m
}

// WSGDHCPPoolsDhcpPoolInfo
//
// Definition: dhcppools_dhcpPoolInfo
//
// DHCP Pool Information List
type WSGDHCPPoolsDhcpPoolInfo struct {
	ApIp *string `json:"apIp,omitempty"`

	AvailableIpCount *int `json:"availableIpCount,omitempty"`

	ClientInfoList *WSGDHCPPoolsDhcpPoolInfoClientInfoListType `json:"clientInfoList,omitempty"`

	Name *string `json:"name,omitempty"`

	PoolEndIp *string `json:"poolEndIp,omitempty"`

	PoolIndex *int `json:"poolIndex,omitempty"`

	PoolStartIp *string `json:"poolStartIp,omitempty"`

	SubnetMask *string `json:"subnetMask,omitempty"`

	TotalIpCount *int `json:"totalIpCount,omitempty"`

	UsedIpCount *int `json:"usedIpCount,omitempty"`

	VlanId *int `json:"vlanId,omitempty"`
}

type WSGDHCPPoolsDhcpPoolInfoAPIResponse struct {
	*RawAPIResponse
	Data *WSGDHCPPoolsDhcpPoolInfo
}

func newWSGDHCPPoolsDhcpPoolInfoAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGDHCPPoolsDhcpPoolInfoAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGDHCPPoolsDhcpPoolInfoAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGDHCPPoolsDhcpPoolInfo)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGDHCPPoolsDhcpPoolInfo() *WSGDHCPPoolsDhcpPoolInfo {
	m := new(WSGDHCPPoolsDhcpPoolInfo)
	return m
}

// WSGDHCPPoolsDhcpPoolInfoClientInfoListType
//
// Definition: dhcppools_dhcpPoolInfoClientInfoListType
type WSGDHCPPoolsDhcpPoolInfoClientInfoListType struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGDHCPPoolsDhcpClientInfo `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGDHCPPoolsDhcpPoolInfoClientInfoListType() *WSGDHCPPoolsDhcpPoolInfoClientInfoListType {
	m := new(WSGDHCPPoolsDhcpPoolInfoClientInfoListType)
	return m
}

// WSGDHCPPools
//
// Definition: dhcppools_dhcpPools
//
// DHCP Pools Usage Per AP
type WSGDHCPPools struct {
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

	PoolInfoList []*WSGDHCPPoolsDhcpPoolInfo `json:"poolInfoList,omitempty"`

	// TenantId
	// Constraints:
	//    - nullable
	TenantId *string `json:"tenantId,omitempty"`
}

type WSGDHCPPoolsAPIResponse struct {
	*RawAPIResponse
	Data *WSGDHCPPools
}

func newWSGDHCPPoolsAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGDHCPPoolsAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGDHCPPoolsAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGDHCPPools)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGDHCPPools() *WSGDHCPPools {
	m := new(WSGDHCPPools)
	return m
}
