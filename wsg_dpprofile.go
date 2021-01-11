package bigdog

// API Version: v9_1

import (
	"errors"
	"io"
)

// WSGDPProfileBulkDelete
//
// Definition: dpProfile_bulkDelete
type WSGDPProfileBulkDelete struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGDPProfileBulkDelete() *WSGDPProfileBulkDelete {
	m := new(WSGDPProfileBulkDelete)
	return m
}

// WSGDPProfileDpDhcpProfileBasicBO
//
// Definition: dpProfile_dpDhcpProfileBasicBO
type WSGDPProfileDpDhcpProfileBasicBO struct {
	// DefaultLeaseTime
	// defaultLeaseTime
	DefaultLeaseTime *int `json:"defaultLeaseTime,omitempty"`

	// Description
	// description
	Description *string `json:"description,omitempty"`

	// DomainName
	// domainName
	DomainName *string `json:"domainName,omitempty"`

	// PrimaryDnsServer
	// primaryDnsServer
	PrimaryDnsServer *string `json:"primaryDnsServer,omitempty"`

	// ProfileId
	// profileId
	ProfileId *string `json:"profileId,omitempty"`

	// ProfileName
	// profileName
	ProfileName *string `json:"profileName,omitempty"`

	// SecondaryDnsServer
	// secondaryDnsServer
	SecondaryDnsServer *string `json:"secondaryDnsServer,omitempty"`
}

type WSGDPProfileDpDhcpProfileBasicBOAPIResponse struct {
	*RawAPIResponse
	Data *WSGDPProfileDpDhcpProfileBasicBO
}

func newWSGDPProfileDpDhcpProfileBasicBOAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGDPProfileDpDhcpProfileBasicBOAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGDPProfileDpDhcpProfileBasicBOAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGDPProfileDpDhcpProfileBasicBO)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGDPProfileDpDhcpProfileBasicBO() *WSGDPProfileDpDhcpProfileBasicBO {
	m := new(WSGDPProfileDpDhcpProfileBasicBO)
	return m
}

// WSGDPProfileDpDhcpProfileBasicBOList
//
// Definition: dpProfile_dpDhcpProfileBasicBOList
type WSGDPProfileDpDhcpProfileBasicBOList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGDPProfileDpDhcpProfileBasicBO `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGDPProfileDpDhcpProfileBasicBOListAPIResponse struct {
	*RawAPIResponse
	Data *WSGDPProfileDpDhcpProfileBasicBOList
}

func newWSGDPProfileDpDhcpProfileBasicBOListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGDPProfileDpDhcpProfileBasicBOListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGDPProfileDpDhcpProfileBasicBOListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGDPProfileDpDhcpProfileBasicBOList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGDPProfileDpDhcpProfileBasicBOList() *WSGDPProfileDpDhcpProfileBasicBOList {
	m := new(WSGDPProfileDpDhcpProfileBasicBOList)
	return m
}

// WSGDPProfileDpDhcpProfileHostBO
//
// Definition: dpProfile_dpDhcpProfileHostBO
type WSGDPProfileDpDhcpProfileHostBO struct {
	// BroadcastAddress
	// broadcastAddress
	BroadcastAddress *string `json:"broadcastAddress,omitempty"`

	// Description
	// description
	Description *string `json:"description,omitempty"`

	// DnsServers
	// dnsServers
	DnsServers []string `json:"dnsServers,omitempty"`

	// DomainName
	// domainName
	DomainName *string `json:"domainName,omitempty"`

	// FixedAddress
	// fixedAddress
	FixedAddress *string `json:"fixedAddress,omitempty"`

	// HardwareEthernet
	// hardwareEthernet
	HardwareEthernet *string `json:"hardwareEthernet,omitempty"`

	// HostId
	// hostId
	HostId *string `json:"hostId,omitempty"`

	// HostName
	// hostName
	HostName *string `json:"hostName,omitempty"`

	// LeaseTime
	// leaseTime
	LeaseTime *int `json:"leaseTime,omitempty"`

	// Name
	// name
	Name *string `json:"name,omitempty"`

	// ProfileId
	// profileId
	ProfileId *string `json:"profileId,omitempty"`

	// Routers
	// routers
	Routers []string `json:"routers,omitempty"`
}

type WSGDPProfileDpDhcpProfileHostBOAPIResponse struct {
	*RawAPIResponse
	Data *WSGDPProfileDpDhcpProfileHostBO
}

func newWSGDPProfileDpDhcpProfileHostBOAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGDPProfileDpDhcpProfileHostBOAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGDPProfileDpDhcpProfileHostBOAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGDPProfileDpDhcpProfileHostBO)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGDPProfileDpDhcpProfileHostBO() *WSGDPProfileDpDhcpProfileHostBO {
	m := new(WSGDPProfileDpDhcpProfileHostBO)
	return m
}

// WSGDPProfileDpDhcpProfileHostBOList
//
// Definition: dpProfile_dpDhcpProfileHostBOList
type WSGDPProfileDpDhcpProfileHostBOList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGDPProfileDpDhcpProfileHostBO `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGDPProfileDpDhcpProfileHostBOListAPIResponse struct {
	*RawAPIResponse
	Data *WSGDPProfileDpDhcpProfileHostBOList
}

func newWSGDPProfileDpDhcpProfileHostBOListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGDPProfileDpDhcpProfileHostBOListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGDPProfileDpDhcpProfileHostBOListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGDPProfileDpDhcpProfileHostBOList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGDPProfileDpDhcpProfileHostBOList() *WSGDPProfileDpDhcpProfileHostBOList {
	m := new(WSGDPProfileDpDhcpProfileHostBOList)
	return m
}

// WSGDPProfileDpDhcpProfileOptionBO
//
// Definition: dpProfile_dpDhcpProfileOptionBO
type WSGDPProfileDpDhcpProfileOptionBO struct {
	// CodeNumber
	// codeNumber
	CodeNumber *int `json:"codeNumber,omitempty"`

	// FunctionName
	// functionName
	FunctionName *string `json:"functionName,omitempty"`

	// Type
	// type
	Type *string `json:"type,omitempty"`

	// Value
	// value
	Value *string `json:"value,omitempty"`
}

func NewWSGDPProfileDpDhcpProfileOptionBO() *WSGDPProfileDpDhcpProfileOptionBO {
	m := new(WSGDPProfileDpDhcpProfileOptionBO)
	return m
}

// WSGDPProfileDpDhcpProfileOptionInstance
//
// Definition: dpProfile_dpDhcpProfileOptionInstance
type WSGDPProfileDpDhcpProfileOptionInstance struct {
	FunctionName *string `json:"functionName,omitempty"`

	Value *string `json:"value,omitempty"`
}

func NewWSGDPProfileDpDhcpProfileOptionInstance() *WSGDPProfileDpDhcpProfileOptionInstance {
	m := new(WSGDPProfileDpDhcpProfileOptionInstance)
	return m
}

// WSGDPProfileDpDhcpProfileOptionSpaceApplyToBO
//
// Definition: dpProfile_dpDhcpProfileOptionSpaceApplyToBO
type WSGDPProfileDpDhcpProfileOptionSpaceApplyToBO struct {
	// AppliedPoolNames
	// appliedPoolNames
	AppliedPoolNames []string `json:"appliedPoolNames,omitempty"`

	// Description
	// description
	Description *string `json:"description,omitempty"`

	// Name
	// name
	Name *string `json:"name,omitempty"`

	// Options
	// options
	Options []*WSGDPProfileDpDhcpProfileOptionBO `json:"options,omitempty"`

	SpaceId *string `json:"spaceId,omitempty"`
}

type WSGDPProfileDpDhcpProfileOptionSpaceApplyToBOAPIResponse struct {
	*RawAPIResponse
	Data *WSGDPProfileDpDhcpProfileOptionSpaceApplyToBO
}

func newWSGDPProfileDpDhcpProfileOptionSpaceApplyToBOAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGDPProfileDpDhcpProfileOptionSpaceApplyToBOAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGDPProfileDpDhcpProfileOptionSpaceApplyToBOAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGDPProfileDpDhcpProfileOptionSpaceApplyToBO)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGDPProfileDpDhcpProfileOptionSpaceApplyToBO() *WSGDPProfileDpDhcpProfileOptionSpaceApplyToBO {
	m := new(WSGDPProfileDpDhcpProfileOptionSpaceApplyToBO)
	return m
}

// WSGDPProfileDpDhcpProfileOptionSpaceApplyToBOList
//
// Definition: dpProfile_dpDhcpProfileOptionSpaceApplyToBOList
type WSGDPProfileDpDhcpProfileOptionSpaceApplyToBOList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGDPProfileDpDhcpProfileOptionSpaceApplyToBO `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGDPProfileDpDhcpProfileOptionSpaceApplyToBOListAPIResponse struct {
	*RawAPIResponse
	Data *WSGDPProfileDpDhcpProfileOptionSpaceApplyToBOList
}

func newWSGDPProfileDpDhcpProfileOptionSpaceApplyToBOListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGDPProfileDpDhcpProfileOptionSpaceApplyToBOListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGDPProfileDpDhcpProfileOptionSpaceApplyToBOListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGDPProfileDpDhcpProfileOptionSpaceApplyToBOList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGDPProfileDpDhcpProfileOptionSpaceApplyToBOList() *WSGDPProfileDpDhcpProfileOptionSpaceApplyToBOList {
	m := new(WSGDPProfileDpDhcpProfileOptionSpaceApplyToBOList)
	return m
}

// WSGDPProfileDpDhcpProfileOptionSpaceBO
//
// Definition: dpProfile_dpDhcpProfileOptionSpaceBO
type WSGDPProfileDpDhcpProfileOptionSpaceBO struct {
	// Description
	// description
	Description *string `json:"description,omitempty"`

	// Name
	// name
	Name *string `json:"name,omitempty"`

	// Options
	// options
	Options []*WSGDPProfileDpDhcpProfileOptionBO `json:"options,omitempty"`

	// SpaceId
	// spaceId
	SpaceId *string `json:"spaceId,omitempty"`
}

func NewWSGDPProfileDpDhcpProfileOptionSpaceBO() *WSGDPProfileDpDhcpProfileOptionSpaceBO {
	m := new(WSGDPProfileDpDhcpProfileOptionSpaceBO)
	return m
}

// WSGDPProfileDpDhcpProfileOptionSpaceInstance
//
// Definition: dpProfile_dpDhcpProfileOptionSpaceInstance
type WSGDPProfileDpDhcpProfileOptionSpaceInstance struct {
	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`

	Options []*WSGDPProfileDpDhcpProfileOptionInstance `json:"options,omitempty"`

	SpaceId *string `json:"spaceId,omitempty"`
}

func NewWSGDPProfileDpDhcpProfileOptionSpaceInstance() *WSGDPProfileDpDhcpProfileOptionSpaceInstance {
	m := new(WSGDPProfileDpDhcpProfileOptionSpaceInstance)
	return m
}

// WSGDPProfileDpDhcpProfilePoolBO
//
// Definition: dpProfile_dpDhcpProfilePoolBO
type WSGDPProfileDpDhcpProfilePoolBO struct {
	// BroadcastAddress
	// broadcastAddress
	BroadcastAddress *string `json:"broadcastAddress,omitempty"`

	// Description
	// description
	Description *string `json:"description,omitempty"`

	// DomainName
	// domainName
	DomainName *string `json:"domainName,omitempty"`

	// ExcludeAddressRange
	// excludeAddressRange
	ExcludeAddressRange *string `json:"excludeAddressRange,omitempty"`

	// HostName
	// hostName
	HostName *string `json:"hostName,omitempty"`

	// Ip
	// ip
	Ip *string `json:"ip,omitempty"`

	// IpRange
	// ipRange
	IpRange *string `json:"ipRange,omitempty"`

	// LeaseTime
	// leaseTime
	LeaseTime *int `json:"leaseTime,omitempty"`

	// NetMask
	// netMask
	NetMask *string `json:"netMask,omitempty"`

	// PoolId
	// poolId
	PoolId *string `json:"poolId,omitempty"`

	// PoolName
	// poolName
	PoolName *string `json:"poolName,omitempty"`

	// PrimaryDnsServer
	// primaryDnsServer
	PrimaryDnsServer *string `json:"primaryDnsServer,omitempty"`

	// PrimaryRouter
	// primaryRouter
	PrimaryRouter *string `json:"primaryRouter,omitempty"`

	// ProfileId
	// profileId
	ProfileId *string `json:"profileId,omitempty"`

	// QinqVlanRanges
	// qinqVlanRanges
	QinqVlanRanges []*WSGDPProfileDpDhcpProfileQinqVlanRangeBO `json:"qinqVlanRanges,omitempty"`

	// SecondaryDnsServer
	// secondaryDnsServer
	SecondaryDnsServer *string `json:"secondaryDnsServer,omitempty"`

	// SecondaryRouter
	// secondaryRouter
	SecondaryRouter *string `json:"secondaryRouter,omitempty"`

	// SubOptionSpaces
	// subOptionSpaces
	SubOptionSpaces []*WSGDPProfileDpDhcpProfileOptionSpaceInstance `json:"subOptionSpaces,omitempty"`

	// VlanRange
	// vlanRange
	VlanRange *string `json:"vlanRange,omitempty"`
}

type WSGDPProfileDpDhcpProfilePoolBOAPIResponse struct {
	*RawAPIResponse
	Data *WSGDPProfileDpDhcpProfilePoolBO
}

func newWSGDPProfileDpDhcpProfilePoolBOAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGDPProfileDpDhcpProfilePoolBOAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGDPProfileDpDhcpProfilePoolBOAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGDPProfileDpDhcpProfilePoolBO)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGDPProfileDpDhcpProfilePoolBO() *WSGDPProfileDpDhcpProfilePoolBO {
	m := new(WSGDPProfileDpDhcpProfilePoolBO)
	return m
}

// WSGDPProfileDpDhcpProfilePoolBOList
//
// Definition: dpProfile_dpDhcpProfilePoolBOList
type WSGDPProfileDpDhcpProfilePoolBOList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGDPProfileDpDhcpProfilePoolBO `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGDPProfileDpDhcpProfilePoolBOListAPIResponse struct {
	*RawAPIResponse
	Data *WSGDPProfileDpDhcpProfilePoolBOList
}

func newWSGDPProfileDpDhcpProfilePoolBOListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGDPProfileDpDhcpProfilePoolBOListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGDPProfileDpDhcpProfilePoolBOListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGDPProfileDpDhcpProfilePoolBOList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGDPProfileDpDhcpProfilePoolBOList() *WSGDPProfileDpDhcpProfilePoolBOList {
	m := new(WSGDPProfileDpDhcpProfilePoolBOList)
	return m
}

// WSGDPProfileDpDhcpProfileQinqVlanRangeBO
//
// Definition: dpProfile_dpDhcpProfileQinqVlanRangeBO
type WSGDPProfileDpDhcpProfileQinqVlanRangeBO struct {
	// Cvlan
	// cvlan
	Cvlan *string `json:"cvlan,omitempty"`

	// Svlan
	// svlan
	Svlan *string `json:"svlan,omitempty"`
}

func NewWSGDPProfileDpDhcpProfileQinqVlanRangeBO() *WSGDPProfileDpDhcpProfileQinqVlanRangeBO {
	m := new(WSGDPProfileDpDhcpProfileQinqVlanRangeBO)
	return m
}

// WSGDPProfileDpNatProfileBasicBO
//
// Definition: dpProfile_dpNatProfileBasicBO
type WSGDPProfileDpNatProfileBasicBO struct {
	// AppliedDpKey
	// The applied DP keys
	AppliedDpKey *string `json:"appliedDpKey,omitempty"`

	// Description
	// description
	Description *string `json:"description,omitempty"`

	NatPublicSubnetId *WSGDPProfileDpNatProfilePublicSubnetIdBO `json:"natPublicSubnetId,omitempty"`

	// NatPublicVlanId
	// natPublicVlanId
	NatPublicVlanId *int `json:"natPublicVlanId,omitempty"`

	// PrimaryNatDefaultRouteGateway
	// primaryNatDefaultRouteGateway
	PrimaryNatDefaultRouteGateway *string `json:"primaryNatDefaultRouteGateway,omitempty"`

	// ProfileId
	// profileId
	ProfileId *string `json:"profileId,omitempty"`

	// ProfileName
	// profileName
	ProfileName *string `json:"profileName,omitempty"`

	// SecondaryNatDefaultRouteGateway
	// secondaryNatDefaultRouteGateway
	SecondaryNatDefaultRouteGateway *string `json:"secondaryNatDefaultRouteGateway,omitempty"`
}

type WSGDPProfileDpNatProfileBasicBOAPIResponse struct {
	*RawAPIResponse
	Data *WSGDPProfileDpNatProfileBasicBO
}

func newWSGDPProfileDpNatProfileBasicBOAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGDPProfileDpNatProfileBasicBOAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGDPProfileDpNatProfileBasicBOAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGDPProfileDpNatProfileBasicBO)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGDPProfileDpNatProfileBasicBO() *WSGDPProfileDpNatProfileBasicBO {
	m := new(WSGDPProfileDpNatProfileBasicBO)
	return m
}

// WSGDPProfileDpNatProfileBasicBOList
//
// Definition: dpProfile_dpNatProfileBasicBOList
type WSGDPProfileDpNatProfileBasicBOList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGDPProfileDpNatProfileBasicBO `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGDPProfileDpNatProfileBasicBOListAPIResponse struct {
	*RawAPIResponse
	Data *WSGDPProfileDpNatProfileBasicBOList
}

func newWSGDPProfileDpNatProfileBasicBOListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGDPProfileDpNatProfileBasicBOListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGDPProfileDpNatProfileBasicBOListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGDPProfileDpNatProfileBasicBOList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGDPProfileDpNatProfileBasicBOList() *WSGDPProfileDpNatProfileBasicBOList {
	m := new(WSGDPProfileDpNatProfileBasicBOList)
	return m
}

// WSGDPProfileDpNatProfilePoolBO
//
// Definition: dpProfile_dpNatProfilePoolBO
type WSGDPProfileDpNatProfilePoolBO struct {
	// Description
	// description
	Description *string `json:"description,omitempty"`

	// NatPortRange
	// natPortRange
	NatPortRange *string `json:"natPortRange,omitempty"`

	// PoolId
	// poolId
	PoolId *string `json:"poolId,omitempty"`

	// PoolName
	// poolName
	PoolName *string `json:"poolName,omitempty"`

	// PrivateQinqVlanRange
	// privateQinqVlanRange
	PrivateQinqVlanRange []*WSGDPProfileDpNatProfilePrivateQinqVlanRangeBO `json:"privateQinqVlanRange,omitempty"`

	// PrivateVlanRange
	// privateVlanRange
	PrivateVlanRange []string `json:"privateVlanRange,omitempty"`

	// ProfileId
	// profileId
	ProfileId *string `json:"profileId,omitempty"`

	// PublicAddressRange
	// publicAddressRange
	PublicAddressRange []string `json:"publicAddressRange,omitempty"`

	// PublicPrefix
	// publicPrefix
	PublicPrefix *int `json:"publicPrefix,omitempty"`

	// PublicVlan
	// publicVlan
	PublicVlan *int `json:"publicVlan,omitempty"`
}

type WSGDPProfileDpNatProfilePoolBOAPIResponse struct {
	*RawAPIResponse
	Data *WSGDPProfileDpNatProfilePoolBO
}

func newWSGDPProfileDpNatProfilePoolBOAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGDPProfileDpNatProfilePoolBOAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGDPProfileDpNatProfilePoolBOAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGDPProfileDpNatProfilePoolBO)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGDPProfileDpNatProfilePoolBO() *WSGDPProfileDpNatProfilePoolBO {
	m := new(WSGDPProfileDpNatProfilePoolBO)
	return m
}

// WSGDPProfileDpNatProfilePoolBOList
//
// Definition: dpProfile_dpNatProfilePoolBOList
type WSGDPProfileDpNatProfilePoolBOList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGDPProfileDpNatProfilePoolBO `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGDPProfileDpNatProfilePoolBOListAPIResponse struct {
	*RawAPIResponse
	Data *WSGDPProfileDpNatProfilePoolBOList
}

func newWSGDPProfileDpNatProfilePoolBOListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGDPProfileDpNatProfilePoolBOListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGDPProfileDpNatProfilePoolBOListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGDPProfileDpNatProfilePoolBOList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGDPProfileDpNatProfilePoolBOList() *WSGDPProfileDpNatProfilePoolBOList {
	m := new(WSGDPProfileDpNatProfilePoolBOList)
	return m
}

// WSGDPProfileDpNatProfilePrivateQinqVlanRangeBO
//
// Definition: dpProfile_dpNatProfilePrivateQinqVlanRangeBO
type WSGDPProfileDpNatProfilePrivateQinqVlanRangeBO struct {
	// Cvlan
	// cvlan
	Cvlan *string `json:"cvlan,omitempty"`

	// Svlan
	// svlan
	Svlan *string `json:"svlan,omitempty"`
}

func NewWSGDPProfileDpNatProfilePrivateQinqVlanRangeBO() *WSGDPProfileDpNatProfilePrivateQinqVlanRangeBO {
	m := new(WSGDPProfileDpNatProfilePrivateQinqVlanRangeBO)
	return m
}

// WSGDPProfileDpNatProfilePublicSubnetIdBO
//
// Definition: dpProfile_dpNatProfilePublicSubnetIdBO
type WSGDPProfileDpNatProfilePublicSubnetIdBO struct {
	// Ip
	// ip
	Ip *string `json:"ip,omitempty"`

	// PrefixLength
	// prefixLength
	PrefixLength *int `json:"prefixLength,omitempty"`
}

func NewWSGDPProfileDpNatProfilePublicSubnetIdBO() *WSGDPProfileDpNatProfilePublicSubnetIdBO {
	m := new(WSGDPProfileDpNatProfilePublicSubnetIdBO)
	return m
}

// WSGDPProfileSettingBO
//
// Definition: dpProfile_dpProfileSettingBO
type WSGDPProfileSettingBO struct {
	// Description
	// description
	Description *string `json:"description,omitempty"`

	// DhcpProfileId
	// dhcpProfileId
	DhcpProfileId *string `json:"dhcpProfileId,omitempty"`

	// DhcpProfileName
	// dhcpProfileName
	DhcpProfileName *string `json:"dhcpProfileName,omitempty"`

	// DpKey
	// dpKey
	DpKey *string `json:"dpKey,omitempty"`

	// DpName
	// dpName
	DpName *string `json:"dpName,omitempty"`

	// DpVersion
	// dpVersion
	DpVersion *string `json:"dpVersion,omitempty"`

	// NatProfileId
	// natProfileId
	NatProfileId *string `json:"natProfileId,omitempty"`

	// NatProfileName
	// natProfileName
	NatProfileName *string `json:"natProfileName,omitempty"`
}

type WSGDPProfileSettingBOAPIResponse struct {
	*RawAPIResponse
	Data *WSGDPProfileSettingBO
}

func newWSGDPProfileSettingBOAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGDPProfileSettingBOAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGDPProfileSettingBOAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGDPProfileSettingBO)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGDPProfileSettingBO() *WSGDPProfileSettingBO {
	m := new(WSGDPProfileSettingBO)
	return m
}

// WSGDPProfileSettingBOList
//
// Definition: dpProfile_dpProfileSettingBOList
type WSGDPProfileSettingBOList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGDPProfileSettingBO `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGDPProfileSettingBOListAPIResponse struct {
	*RawAPIResponse
	Data *WSGDPProfileSettingBOList
}

func newWSGDPProfileSettingBOListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGDPProfileSettingBOListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGDPProfileSettingBOListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGDPProfileSettingBOList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGDPProfileSettingBOList() *WSGDPProfileSettingBOList {
	m := new(WSGDPProfileSettingBOList)
	return m
}
