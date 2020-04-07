package vsz

// API Version: v9_0

type WSGSplitTunnelCreateSplitTunnelProfile struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// Rules
	// Destination rule of split tunnel profile
	// Constraints:
	//    - required
	Rules []*WSGSplitTunnelIpMaskRule `json:"rules" validate:"required,dive"`
}

func NewWSGSplitTunnelCreateSplitTunnelProfile() *WSGSplitTunnelCreateSplitTunnelProfile {
	m := new(WSGSplitTunnelCreateSplitTunnelProfile)
	return m
}

type WSGSplitTunnelModifySplitTunnelProfile struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Rules
	// Destination rule of split tunnel profile
	// Constraints:
	//    - nullable
	Rules []*WSGSplitTunnelIpMaskRule `json:"rules,omitempty" validate:"omitempty,dive"`
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
	DestinationIp *string `json:"destinationIp" validate:"required"`

	// DestinationIpMask
	// Destination IP mask of split tunnel profile rule
	// Constraints:
	//    - required
	DestinationIpMask *string `json:"destinationIpMask" validate:"required"`
}

func NewWSGSplitTunnelIpMaskRule() *WSGSplitTunnelIpMaskRule {
	m := new(WSGSplitTunnelIpMaskRule)
	return m
}

type WSGSplitTunnelProfile struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Rules
	// Destination rule of split tunnel profile
	// Constraints:
	//    - nullable
	Rules []*WSGSplitTunnelIpMaskRule `json:"rules,omitempty" validate:"omitempty,dive"`

	// TenantId
	// Constraints:
	//    - nullable
	TenantId *string `json:"tenantId,omitempty"`

	// ZoneId
	// Constraints:
	//    - nullable
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGSplitTunnelProfile() *WSGSplitTunnelProfile {
	m := new(WSGSplitTunnelProfile)
	return m
}

type WSGSplitTunnelProfileList struct {
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
	List []*WSGSplitTunnelProfileListType `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGSplitTunnelProfileList() *WSGSplitTunnelProfileList {
	m := new(WSGSplitTunnelProfileList)
	return m
}

type WSGSplitTunnelProfileListType struct {
	// Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGSplitTunnelProfileListType() *WSGSplitTunnelProfileListType {
	m := new(WSGSplitTunnelProfileListType)
	return m
}

type WSGSplitTunnelProfileQuery struct {
	// Extra
	// Constraints:
	//    - nullable
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

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
	List []*WSGSplitTunnelProfile `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGSplitTunnelProfileQuery() *WSGSplitTunnelProfileQuery {
	m := new(WSGSplitTunnelProfileQuery)
	return m
}
