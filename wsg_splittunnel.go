package bigdog

// API Version: v9_1

import (
	"encoding/json"
	"net/http"
)

// WSGSplitTunnelCreateSplitTunnelProfile
//
// Definition: splitTunnel_createSplitTunnelProfile
type WSGSplitTunnelCreateSplitTunnelProfile struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// Rules
	// Destination rule of split tunnel profile
	// Constraints:
	//    - required
	Rules []*WSGSplitTunnelIpMaskRule `json:"rules"`
}

func NewWSGSplitTunnelCreateSplitTunnelProfile() *WSGSplitTunnelCreateSplitTunnelProfile {
	m := new(WSGSplitTunnelCreateSplitTunnelProfile)
	return m
}

// WSGSplitTunnelModifySplitTunnelProfile
//
// Definition: splitTunnel_modifySplitTunnelProfile
type WSGSplitTunnelModifySplitTunnelProfile struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Rules
	// Destination rule of split tunnel profile
	Rules []*WSGSplitTunnelIpMaskRule `json:"rules,omitempty"`
}

func NewWSGSplitTunnelModifySplitTunnelProfile() *WSGSplitTunnelModifySplitTunnelProfile {
	m := new(WSGSplitTunnelModifySplitTunnelProfile)
	return m
}

// WSGSplitTunnelIpMaskRule
//
// Definition: splitTunnel_SplitTunnelIpMaskRule
type WSGSplitTunnelIpMaskRule struct {
	// DestinationIp
	// Destination IP of split tunnel profile rule
	// Constraints:
	//    - required
	DestinationIp *string `json:"destinationIp"`

	// DestinationIpMask
	// Destination IP mask of split tunnel profile rule
	// Constraints:
	//    - required
	DestinationIpMask *string `json:"destinationIpMask"`
}

func NewWSGSplitTunnelIpMaskRule() *WSGSplitTunnelIpMaskRule {
	m := new(WSGSplitTunnelIpMaskRule)
	return m
}

// WSGSplitTunnelProfile
//
// Definition: splitTunnel_splitTunnelProfile
type WSGSplitTunnelProfile struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	DomainId *string `json:"domainId,omitempty"`

	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Rules
	// Destination rule of split tunnel profile
	Rules []*WSGSplitTunnelIpMaskRule `json:"rules,omitempty"`

	TenantId *string `json:"tenantId,omitempty"`

	ZoneId *string `json:"zoneId,omitempty"`
}

type WSGSplitTunnelProfileAPIResponse struct {
	*RawAPIResponse
	Data *WSGSplitTunnelProfile
}

func newWSGSplitTunnelProfileAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGSplitTunnelProfileAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGSplitTunnelProfileAPIResponse) Hydrate() error {
	r.Data = new(WSGSplitTunnelProfile)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSplitTunnelProfile() *WSGSplitTunnelProfile {
	m := new(WSGSplitTunnelProfile)
	return m
}

// WSGSplitTunnelProfileList
//
// Definition: splitTunnel_splitTunnelProfileList
type WSGSplitTunnelProfileList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSplitTunnelProfileListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGSplitTunnelProfileListAPIResponse struct {
	*RawAPIResponse
	Data *WSGSplitTunnelProfileList
}

func newWSGSplitTunnelProfileListAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGSplitTunnelProfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGSplitTunnelProfileListAPIResponse) Hydrate() error {
	r.Data = new(WSGSplitTunnelProfileList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSplitTunnelProfileList() *WSGSplitTunnelProfileList {
	m := new(WSGSplitTunnelProfileList)
	return m
}

// WSGSplitTunnelProfileListType
//
// Definition: splitTunnel_splitTunnelProfileListType
type WSGSplitTunnelProfileListType struct {
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGSplitTunnelProfileListType() *WSGSplitTunnelProfileListType {
	m := new(WSGSplitTunnelProfileListType)
	return m
}

// WSGSplitTunnelProfileQuery
//
// Definition: splitTunnel_splitTunnelProfileQuery
type WSGSplitTunnelProfileQuery struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSplitTunnelProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGSplitTunnelProfileQueryAPIResponse struct {
	*RawAPIResponse
	Data *WSGSplitTunnelProfileQuery
}

func newWSGSplitTunnelProfileQueryAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGSplitTunnelProfileQueryAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGSplitTunnelProfileQueryAPIResponse) Hydrate() error {
	r.Data = new(WSGSplitTunnelProfileQuery)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSplitTunnelProfileQuery() *WSGSplitTunnelProfileQuery {
	m := new(WSGSplitTunnelProfileQuery)
	return m
}
