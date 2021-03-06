package bigdog

// API Version: v9_1

import (
	"errors"
	"io"
)

// WSGAPRulesApRuleConfiguration
//
// Definition: aprules_apRuleConfiguration
type WSGAPRulesApRuleConfiguration struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	GpsCoordinates *WSGAPRulesGpsCoordinates `json:"gpsCoordinates,omitempty"`

	// Id
	// Identifier of the AP Registration Rules
	Id *string `json:"id,omitempty"`

	IpAddressRange *WSGAPRulesIpAddressRange `json:"ipAddressRange,omitempty"`

	MobilityZone *WSGCommonGenericRef `json:"mobilityZone,omitempty"`

	// Priority
	// priority of the AP Registration Rules
	Priority *int `json:"priority,omitempty"`

	// ProvisionTag
	// ProvisionTag of the AP Registration Rules
	ProvisionTag *string `json:"provisionTag,omitempty"`

	Subnet *WSGAPRulesSubnet `json:"subnet,omitempty"`

	// Type
	// type of the AP Registration Rules
	// Constraints:
	//    - oneof:[IPAddressRange,Subnet,GPSCoordinates,ProvisionTag]
	Type *string `json:"type,omitempty"`
}

type WSGAPRulesApRuleConfigurationAPIResponse struct {
	*RawAPIResponse
	Data *WSGAPRulesApRuleConfiguration
}

func newWSGAPRulesApRuleConfigurationAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAPRulesApRuleConfigurationAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAPRulesApRuleConfigurationAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAPRulesApRuleConfiguration)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAPRulesApRuleConfiguration() *WSGAPRulesApRuleConfiguration {
	m := new(WSGAPRulesApRuleConfiguration)
	return m
}

// WSGAPRulesApRuleList
//
// Definition: aprules_apRuleList
type WSGAPRulesApRuleList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAPRulesApRuleListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAPRulesApRuleListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAPRulesApRuleList
}

func newWSGAPRulesApRuleListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAPRulesApRuleListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAPRulesApRuleListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGAPRulesApRuleList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGAPRulesApRuleList() *WSGAPRulesApRuleList {
	m := new(WSGAPRulesApRuleList)
	return m
}

// WSGAPRulesApRuleListType
//
// Definition: aprules_apRuleListType
type WSGAPRulesApRuleListType struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the AP Registration Rules
	Id *string `json:"id,omitempty"`

	// Priority
	// priority of the AP Registration Rules
	Priority *int `json:"priority,omitempty"`
}

func NewWSGAPRulesApRuleListType() *WSGAPRulesApRuleListType {
	m := new(WSGAPRulesApRuleListType)
	return m
}

// WSGAPRulesCreateApRule
//
// Definition: aprules_createApRule
type WSGAPRulesCreateApRule struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	GpsCoordinates *WSGAPRulesGpsCoordinates `json:"gpsCoordinates,omitempty"`

	IpAddressRange *WSGAPRulesIpAddressRange `json:"ipAddressRange,omitempty"`

	// MobilityZone
	// Constraints:
	//    - required
	MobilityZone *WSGCommonGenericRef `json:"mobilityZone"`

	// ProvisionTag
	// ProvisionTag of the AP Registration Rules
	ProvisionTag *string `json:"provisionTag,omitempty"`

	Subnet *WSGAPRulesSubnet `json:"subnet,omitempty"`

	// Type
	// type of the AP Registration Rules
	// Constraints:
	//    - required
	//    - oneof:[IPAddressRange,Subnet,GPSCoordinates,ProvisionTag]
	Type *string `json:"type"`
}

func NewWSGAPRulesCreateApRule() *WSGAPRulesCreateApRule {
	m := new(WSGAPRulesCreateApRule)
	return m
}

// WSGAPRulesGpsCoordinates
//
// Definition: aprules_gpsCoordinates
type WSGAPRulesGpsCoordinates struct {
	// Distance
	// distance
	Distance *float64 `json:"distance,omitempty"`

	Latitude interface{} `json:"latitude,omitempty"`

	Longitude interface{} `json:"longitude,omitempty"`
}

func NewWSGAPRulesGpsCoordinates() *WSGAPRulesGpsCoordinates {
	m := new(WSGAPRulesGpsCoordinates)
	return m
}

// WSGAPRulesIpAddressRange
//
// Definition: aprules_ipAddressRange
type WSGAPRulesIpAddressRange struct {
	FromIp *WSGCommonIpAddress `json:"fromIp,omitempty"`

	ToIp *WSGCommonIpAddress `json:"toIp,omitempty"`
}

func NewWSGAPRulesIpAddressRange() *WSGAPRulesIpAddressRange {
	m := new(WSGAPRulesIpAddressRange)
	return m
}

// WSGAPRulesModifyApRule
//
// Definition: aprules_modifyApRule
type WSGAPRulesModifyApRule struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	GpsCoordinates *WSGAPRulesGpsCoordinates `json:"gpsCoordinates,omitempty"`

	IpAddressRange *WSGAPRulesIpAddressRange `json:"ipAddressRange,omitempty"`

	MobilityZone *WSGCommonGenericRef `json:"mobilityZone,omitempty"`

	// ProvisionTag
	// ProvisionTag of the AP Registration Rules
	ProvisionTag *string `json:"provisionTag,omitempty"`

	Subnet *WSGAPRulesSubnet `json:"subnet,omitempty"`

	// Type
	// type of the AP Registration Rules
	// Constraints:
	//    - oneof:[IPAddressRange,Subnet,GPSCoordinates,ProvisionTag]
	Type *string `json:"type,omitempty"`
}

func NewWSGAPRulesModifyApRule() *WSGAPRulesModifyApRule {
	m := new(WSGAPRulesModifyApRule)
	return m
}

// WSGAPRulesSubnet
//
// Definition: aprules_subnet
type WSGAPRulesSubnet struct {
	NetworkAddress *WSGCommonIpAddress `json:"networkAddress,omitempty"`

	// SubnetMask
	// subnetMask
	SubnetMask *string `json:"subnetMask,omitempty"`
}

func NewWSGAPRulesSubnet() *WSGAPRulesSubnet {
	m := new(WSGAPRulesSubnet)
	return m
}
