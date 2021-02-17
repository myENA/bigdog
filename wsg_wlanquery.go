package bigdog

// API Version: v9_1

import (
	"errors"
	"io"
)

// WSGWLANQueryApWlanBssid
//
// Definition: wlanQuery_apWlanBssid
type WSGWLANQueryApWlanBssid struct {
	ApMac *string `json:"apMac,omitempty"`

	DeviceName *string `json:"deviceName,omitempty"`

	WlanBssids []*WSGWLANQueryWlanBssid `json:"wlanBssids,omitempty"`
}

func NewWSGWLANQueryApWlanBssid() *WSGWLANQueryApWlanBssid {
	m := new(WSGWLANQueryApWlanBssid)
	return m
}

// WSGWLANQueryApWlanBssidQueryList
//
// Definition: wlanQuery_apWlanBssidQueryList
type WSGWLANQueryApWlanBssidQueryList struct {
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of first index in current page
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Has more data or not
	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGWLANQueryApWlanBssid `json:"list,omitempty"`

	// TotalCount
	// Total matched AP count
	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGWLANQueryApWlanBssidQueryListAPIResponse struct {
	*RawAPIResponse
	Data *WSGWLANQueryApWlanBssidQueryList
}

func newWSGWLANQueryApWlanBssidQueryListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGWLANQueryApWlanBssidQueryListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGWLANQueryApWlanBssidQueryListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGWLANQueryApWlanBssidQueryList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGWLANQueryApWlanBssidQueryList() *WSGWLANQueryApWlanBssidQueryList {
	m := new(WSGWLANQueryApWlanBssidQueryList)
	return m
}

// WSGWLANQueryCreateWlanQuery
//
// Definition: wlanQuery_createWlanQuery
type WSGWLANQueryCreateWlanQuery struct {
	Alerts *int `json:"alerts,omitempty"`

	// ApplicationVisibility
	// Constraints:
	//    - oneof:[Enabled,Disabled]
	ApplicationVisibility *string `json:"applicationVisibility,omitempty"`

	AuthMethod *string `json:"authMethod,omitempty"`

	AuthType *string `json:"authType,omitempty"`

	Clients *int `json:"clients,omitempty"`

	Description *string `json:"description,omitempty"`

	DomainName *string `json:"domainName,omitempty"`

	// Enability11k
	// Constraints:
	//    - oneof:[Enabled,Disabled]
	Enability11k *string `json:"enability11k,omitempty"`

	// Enability11r
	// Constraints:
	//    - oneof:[Enabled,Disabled]
	Enability11r *string `json:"enability11r,omitempty"`

	EncryptionMethod *string `json:"encryptionMethod,omitempty"`

	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	Name *string `json:"name,omitempty"`

	Ssid *string `json:"ssid,omitempty"`

	// Status
	// Constraints:
	//    - oneof:[Online,Flagged,Offline]
	Status *string `json:"status,omitempty"`

	TenantDomainName *string `json:"tenantDomainName,omitempty"`

	TenantId *string `json:"tenantId,omitempty"`

	Traffic *int `json:"traffic,omitempty"`

	TrafficDownlink *int `json:"trafficDownlink,omitempty"`

	TrafficUplink *int `json:"trafficUplink,omitempty"`

	// Tunneled
	// Constraints:
	//    - oneof:[Tunneled,APBridged]
	Tunneled *string `json:"tunneled,omitempty"`

	Utp *string `json:"utp,omitempty"`

	Vlan *int `json:"vlan,omitempty"`

	WepEncryptionStrength *int `json:"wepEncryptionStrength,omitempty"`

	WlanId *string `json:"wlanId,omitempty"`

	WpaVersion *string `json:"wpaVersion,omitempty"`

	// ZeroITEnabled
	// Constraints:
	//    - oneof:[Enabled,Disabled]
	ZeroITEnabled *string `json:"zeroITEnabled,omitempty"`

	// ZeroITOnboard
	// Constraints:
	//    - oneof:[Enabled,Disabled]
	ZeroITOnboard *string `json:"zeroITOnboard,omitempty"`

	ZoneId *string `json:"zoneId,omitempty"`

	ZoneName *string `json:"zoneName,omitempty"`
}

func NewWSGWLANQueryCreateWlanQuery() *WSGWLANQueryCreateWlanQuery {
	m := new(WSGWLANQueryCreateWlanQuery)
	return m
}

// WSGWLANQueryWlanBssid
//
// Definition: wlanQuery_wlanBssid
type WSGWLANQueryWlanBssid struct {
	Bssid *string `json:"bssid,omitempty"`

	RadioId *int `json:"radioId,omitempty"`

	WlanId *int `json:"wlanId,omitempty"`

	WlanName *string `json:"wlanName,omitempty"`
}

func NewWSGWLANQueryWlanBssid() *WSGWLANQueryWlanBssid {
	m := new(WSGWLANQueryWlanBssid)
	return m
}

// WSGWLANQueryList
//
// Definition: wlanQuery_wlanQueryList
type WSGWLANQueryList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGWLANQueryCreateWlanQuery `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGWLANQueryListAPIResponse struct {
	*RawAPIResponse
	Data *WSGWLANQueryList
}

func newWSGWLANQueryListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGWLANQueryListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGWLANQueryListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGWLANQueryList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGWLANQueryList() *WSGWLANQueryList {
	m := new(WSGWLANQueryList)
	return m
}
