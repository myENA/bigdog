package vsz

// API Version: v9_0

type WSGWIFICallingCreateWifiCallingPolicy struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Identifier of the System (root) domain or partner managed domain to which the Wi-Fi calling policy belongs
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Epdgs
	// ePDG list of the Wi-Fi calling policy
	// Constraints:
	//    - required
	Epdgs []*WSGWIFICallingEpdg `json:"epdgs" validate:"required,dive"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// Priority
	// QoS priority of the Wi-Fi calling policy
	// Constraints:
	//    - required
	//    - oneof:[BACKGROUND,BEST_EFFORT,VIDEO,VOICE]
	Priority *string `json:"priority" validate:"required,oneof=BACKGROUND BEST_EFFORT VIDEO VOICE"`
}

func NewWSGWIFICallingCreateWifiCallingPolicy() *WSGWIFICallingCreateWifiCallingPolicy {
	m := new(WSGWIFICallingCreateWifiCallingPolicy)
	return m
}

type WSGWIFICallingDeleteBulk struct {
	// IdList
	// Constraints:
	//    - nullable
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGWIFICallingDeleteBulk() *WSGWIFICallingDeleteBulk {
	m := new(WSGWIFICallingDeleteBulk)
	return m
}

type WSGWIFICallingEpdg struct {
	// Fqdn
	// Fully qualified domain name of ePDG
	// Constraints:
	//    - nullable
	Fqdn *string `json:"fqdn,omitempty"`

	// Ip
	// Constraints:
	//    - nullable
	Ip *string `json:"ip,omitempty"`
}

func NewWSGWIFICallingEpdg() *WSGWIFICallingEpdg {
	m := new(WSGWIFICallingEpdg)
	return m
}

type WSGWIFICallingModifyWifiCallingPolicy struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Epdgs
	// ePDG list of the Wi-Fi calling policy
	// Constraints:
	//    - nullable
	Epdgs []*WSGWIFICallingEpdg `json:"epdgs,omitempty" validate:"omitempty,dive"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Priority
	// QoS priority of the Wi-Fi calling policy
	// Constraints:
	//    - nullable
	//    - oneof:[BACKGROUND,BEST_EFFORT,VIDEO,VOICE]
	Priority *string `json:"priority,omitempty" validate:"omitempty,oneof=BACKGROUND BEST_EFFORT VIDEO VOICE"`
}

func NewWSGWIFICallingModifyWifiCallingPolicy() *WSGWIFICallingModifyWifiCallingPolicy {
	m := new(WSGWIFICallingModifyWifiCallingPolicy)
	return m
}

type WSGWIFICallingPolicy struct {
	// CreateDateTime
	// Timestamp of being created
	// Constraints:
	//    - nullable
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	// Constraints:
	//    - nullable
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Identifier of the System (root) domain or partner managed domain to which the Wi-Fi calling policy belongs
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// Epdgs
	// ePDG list of the Wi-Fi calling policy
	// Constraints:
	//    - nullable
	Epdgs []*WSGWIFICallingEpdg `json:"epdgs,omitempty" validate:"omitempty,dive"`

	// Id
	// Identifier of the Wi-Fi calling policy
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	// Constraints:
	//    - nullable
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Priority
	// QoS priority of the Wi-Fi calling policy
	// Constraints:
	//    - nullable
	//    - oneof:[BACKGROUND,BEST_EFFORT,VIDEO,VOICE]
	Priority *string `json:"priority,omitempty" validate:"omitempty,oneof=BACKGROUND BEST_EFFORT VIDEO VOICE"`

	// TenantId
	// Tenant Id
	// Constraints:
	//    - nullable
	TenantId *string `json:"tenantId,omitempty"`
}

func NewWSGWIFICallingPolicy() *WSGWIFICallingPolicy {
	m := new(WSGWIFICallingPolicy)
	return m
}

type WSGWIFICallingPolicyList struct {
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
	List []*WSGWIFICallingPolicy `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGWIFICallingPolicyList() *WSGWIFICallingPolicyList {
	m := new(WSGWIFICallingPolicyList)
	return m
}
