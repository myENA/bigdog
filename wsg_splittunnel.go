package ruckus

// API Version: v9_0

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

func NewWSGSplitTunnelProfile() *WSGSplitTunnelProfile {
	m := new(WSGSplitTunnelProfile)
	return m
}

type WSGSplitTunnelProfileList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSplitTunnelProfileListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGSplitTunnelProfileList() *WSGSplitTunnelProfileList {
	m := new(WSGSplitTunnelProfileList)
	return m
}

type WSGSplitTunnelProfileListType struct {
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGSplitTunnelProfileListType() *WSGSplitTunnelProfileListType {
	m := new(WSGSplitTunnelProfileListType)
	return m
}

type WSGSplitTunnelProfileQuery struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSplitTunnelProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGSplitTunnelProfileQuery() *WSGSplitTunnelProfileQuery {
	m := new(WSGSplitTunnelProfileQuery)
	return m
}
