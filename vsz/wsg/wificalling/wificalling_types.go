package wificalling

// API Version: v8_0

type CreateWifiCallingPolicy struct {
	Description *string `json:"description,omitempty"`

	// DomainID
	// Identifier of the System (root) domain or partner managed domain to which the Wi-Fi calling policy
	// belongs
	DomainID *string `json:"domainId,omitempty"`

	// Epdgs
	// ePDG list of the Wi-Fi calling policy
	Epdgs []*Epdg `json:"epdgs,omitempty"`

	Name *string `json:"name,omitempty"`

	// Priority
	// QoS priority of the Wi-Fi calling policy
	Priority *string `json:"priority,omitempty"`
}

type DeleteBulk struct {
	IDList common.IDList `json:"idList,omitempty"`
}

type Epdg struct {
	// Fqdn
	// Fully qualified domain name of ePDG
	Fqdn *string `json:"fqdn,omitempty"`

	IP *string `json:"ip,omitempty"`
}

type ModifyEntireWifiCallingPolicy struct {
	*ModifyWifiCallingPolicy
}

type ModifyWifiCallingPolicy struct {
	Description *string `json:"description,omitempty"`

	// Epdgs
	// ePDG list of the Wi-Fi calling policy
	Epdgs []*Epdg `json:"epdgs,omitempty"`

	Name *string `json:"name,omitempty"`

	// Priority
	// QoS priority of the Wi-Fi calling policy
	Priority *string `json:"priority,omitempty"`
}

type WifiCallingPolicy struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *string `json:"description,omitempty"`

	// DomainID
	// Identifier of the System (root) domain or partner managed domain to which the Wi-Fi calling policy
	// belongs
	DomainID *string `json:"domainId,omitempty"`

	// Epdgs
	// ePDG list of the Wi-Fi calling policy
	Epdgs []*Epdg `json:"epdgs,omitempty"`

	// ID
	// Identifier of the Wi-Fi calling policy
	ID *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *string `json:"name,omitempty"`

	// Priority
	// QoS priority of the Wi-Fi calling policy
	Priority *string `json:"priority,omitempty"`

	// TenantID
	// Tenant Id
	TenantID *string `json:"tenantId,omitempty"`
}

type WifiCallingPolicyList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WifiCallingPolicy `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}
