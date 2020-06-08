package ruckus

// API Version: v9_0

type WSGWIFICallingCreateWifiCallingPolicy struct {
	Description *WSGWIFICallingCreateWifiCallingPolicy `json:"description,omitempty"`

	// DomainId
	// Identifier of the System (root) domain or partner managed domain to which the Wi-Fi calling policy belongs
	DomainId *string `json:"domainId,omitempty"`

	// Epdgs
	// ePDG list of the Wi-Fi calling policy
	// Constraints:
	//    - required
	Epdgs []*WSGWIFICallingCreateWifiCallingPolicy `json:"epdgs"`

	// Name
	// Constraints:
	//    - required
	Name *WSGWIFICallingCreateWifiCallingPolicy `json:"name"`

	// Priority
	// QoS priority of the Wi-Fi calling policy
	// Constraints:
	//    - required
	//    - oneof:[BACKGROUND,BEST_EFFORT,VIDEO,VOICE]
	Priority *string `json:"priority"`
}

func NewWSGWIFICallingCreateWifiCallingPolicy() *WSGWIFICallingCreateWifiCallingPolicy {
	m := new(WSGWIFICallingCreateWifiCallingPolicy)
	return m
}

type WSGWIFICallingDeleteBulk struct {
	IdList *WSGWIFICallingDeleteBulk `json:"idList,omitempty"`
}

func NewWSGWIFICallingDeleteBulk() *WSGWIFICallingDeleteBulk {
	m := new(WSGWIFICallingDeleteBulk)
	return m
}

type WSGWIFICallingEpdg struct {
	// Fqdn
	// Fully qualified domain name of ePDG
	Fqdn *string `json:"fqdn,omitempty"`

	Ip *string `json:"ip,omitempty"`
}

func NewWSGWIFICallingEpdg() *WSGWIFICallingEpdg {
	m := new(WSGWIFICallingEpdg)
	return m
}

type WSGWIFICallingModifyWifiCallingPolicy struct {
	Description *WSGWIFICallingModifyWifiCallingPolicy `json:"description,omitempty"`

	// Epdgs
	// ePDG list of the Wi-Fi calling policy
	Epdgs []*WSGWIFICallingModifyWifiCallingPolicy `json:"epdgs,omitempty"`

	Name *WSGWIFICallingModifyWifiCallingPolicy `json:"name,omitempty"`

	// Priority
	// QoS priority of the Wi-Fi calling policy
	// Constraints:
	//    - oneof:[BACKGROUND,BEST_EFFORT,VIDEO,VOICE]
	Priority *string `json:"priority,omitempty"`
}

func NewWSGWIFICallingModifyWifiCallingPolicy() *WSGWIFICallingModifyWifiCallingPolicy {
	m := new(WSGWIFICallingModifyWifiCallingPolicy)
	return m
}

type WSGWIFICallingPolicy struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGWIFICallingPolicy `json:"description,omitempty"`

	// DomainId
	// Identifier of the System (root) domain or partner managed domain to which the Wi-Fi calling policy belongs
	DomainId *string `json:"domainId,omitempty"`

	// Epdgs
	// ePDG list of the Wi-Fi calling policy
	Epdgs []*WSGWIFICallingPolicy `json:"epdgs,omitempty"`

	// Id
	// Identifier of the Wi-Fi calling policy
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *WSGWIFICallingPolicy `json:"name,omitempty"`

	// Priority
	// QoS priority of the Wi-Fi calling policy
	// Constraints:
	//    - oneof:[BACKGROUND,BEST_EFFORT,VIDEO,VOICE]
	Priority *string `json:"priority,omitempty"`

	// TenantId
	// Tenant Id
	TenantId *string `json:"tenantId,omitempty"`
}

func NewWSGWIFICallingPolicy() *WSGWIFICallingPolicy {
	m := new(WSGWIFICallingPolicy)
	return m
}

type WSGWIFICallingPolicyList struct {
	Extra *WSGWIFICallingPolicyList `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGWIFICallingPolicyList `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGWIFICallingPolicyList() *WSGWIFICallingPolicyList {
	m := new(WSGWIFICallingPolicyList)
	return m
}
