package bigdog

// API Version: v9_0

// WSGWIFICallingCreateWifiCallingPolicy
//
// Definition: wifiCalling_createWifiCallingPolicy
type WSGWIFICallingCreateWifiCallingPolicy struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Identifier of the System (root) domain or partner managed domain to which the Wi-Fi calling policy belongs
	DomainId *string `json:"domainId,omitempty"`

	// Epdgs
	// ePDG list of the Wi-Fi calling policy
	// Constraints:
	//    - required
	Epdgs []*WSGWIFICallingEpdg `json:"epdgs"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

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

// WSGWIFICallingDeleteBulk
//
// Definition: wifiCalling_deleteBulk
type WSGWIFICallingDeleteBulk struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGWIFICallingDeleteBulk() *WSGWIFICallingDeleteBulk {
	m := new(WSGWIFICallingDeleteBulk)
	return m
}

// WSGWIFICallingEpdg
//
// Definition: wifiCalling_epdg
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

// WSGWIFICallingModifyWifiCallingPolicy
//
// Definition: wifiCalling_modifyWifiCallingPolicy
type WSGWIFICallingModifyWifiCallingPolicy struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Epdgs
	// ePDG list of the Wi-Fi calling policy
	Epdgs []*WSGWIFICallingEpdg `json:"epdgs,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

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

// WSGWIFICallingPolicy
//
// Definition: wifiCalling_wifiCallingPolicy
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

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Identifier of the System (root) domain or partner managed domain to which the Wi-Fi calling policy belongs
	DomainId *string `json:"domainId,omitempty"`

	// Epdgs
	// ePDG list of the Wi-Fi calling policy
	Epdgs []*WSGWIFICallingEpdg `json:"epdgs,omitempty"`

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

	Name *WSGCommonNormalName `json:"name,omitempty"`

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

// WSGWIFICallingPolicyList
//
// Definition: wifiCalling_wifiCallingPolicyList
type WSGWIFICallingPolicyList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGWIFICallingPolicy `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGWIFICallingPolicyList() *WSGWIFICallingPolicyList {
	m := new(WSGWIFICallingPolicyList)
	return m
}
