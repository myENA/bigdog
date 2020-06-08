package ruckus

// API Version: v9_0

type WSGSplitTunnelCreateSplitTunnelProfile struct {
	Description **WSGSplitTunnelCreateSplitTunnelProfile `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name **WSGSplitTunnelCreateSplitTunnelProfile `json:"name"`

	// Rules
	// Destination rule of split tunnel profile
	// Constraints:
	//    - required
	Rules []**WSGSplitTunnelCreateSplitTunnelProfile `json:"rules"`
}

func NewWSGSplitTunnelCreateSplitTunnelProfile() *WSGSplitTunnelCreateSplitTunnelProfile {
	m := new(WSGSplitTunnelCreateSplitTunnelProfile)
	return m
}

type WSGSplitTunnelModifySplitTunnelProfile struct {
	Description **WSGSplitTunnelModifySplitTunnelProfile `json:"description,omitempty"`

	Name **WSGSplitTunnelModifySplitTunnelProfile `json:"name,omitempty"`

	// Rules
	// Destination rule of split tunnel profile
	Rules []**WSGSplitTunnelModifySplitTunnelProfile `json:"rules,omitempty"`
}

func NewWSGSplitTunnelModifySplitTunnelProfile() *WSGSplitTunnelModifySplitTunnelProfile {
	m := new(WSGSplitTunnelModifySplitTunnelProfile)
	return m
}

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

type WSGSplitTunnelProfile struct {
	Description **WSGSplitTunnelProfile `json:"description,omitempty"`

	DomainId *string `json:"domainId,omitempty"`

	Id *string `json:"id,omitempty"`

	Name **WSGSplitTunnelProfile `json:"name,omitempty"`

	// Rules
	// Destination rule of split tunnel profile
	Rules []**WSGSplitTunnelProfile `json:"rules,omitempty"`

	TenantId *string `json:"tenantId,omitempty"`

	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGSplitTunnelProfile() *WSGSplitTunnelProfile {
	m := new(WSGSplitTunnelProfile)
	return m
}

type WSGSplitTunnelProfileList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []**WSGSplitTunnelProfileList `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGSplitTunnelProfileList() *WSGSplitTunnelProfileList {
	m := new(WSGSplitTunnelProfileList)
	return m
}

type WSGSplitTunnelProfileListType struct {
	Id *string `json:"id,omitempty"`

	Name **WSGSplitTunnelProfileListType `json:"name,omitempty"`
}

func NewWSGSplitTunnelProfileListType() *WSGSplitTunnelProfileListType {
	m := new(WSGSplitTunnelProfileListType)
	return m
}

type WSGSplitTunnelProfileQuery struct {
	Extra **WSGSplitTunnelProfileQuery `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []**WSGSplitTunnelProfileQuery `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGSplitTunnelProfileQuery() *WSGSplitTunnelProfileQuery {
	m := new(WSGSplitTunnelProfileQuery)
	return m
}
