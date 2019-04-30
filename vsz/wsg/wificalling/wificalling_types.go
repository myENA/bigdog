package wificalling

// API Version: v8_0

type CreateWifiCallingPolicy struct {
	Description *string `json:"description,omitempty"`
	DomainID    *string `json:"domainId,omitempty"`
	Epdgs       []*Epdg `json:"epdgs,omitempty"`
	Name        *string `json:"name,omitempty"`
	Priority    *string `json:"priority,omitempty"`
}

type DeleteBulk struct {
	IDList common.IDList `json:"idList,omitempty"`
}

type Epdg struct {
	Fqdn *string `json:"fqdn,omitempty"`
	IP   *string `json:"ip,omitempty"`
}

type ModifyEntireWifiCallingPolicy struct {
	*ModifyWifiCallingPolicy
}

type ModifyWifiCallingPolicy struct {
	Description *string `json:"description,omitempty"`
	Epdgs       []*Epdg `json:"epdgs,omitempty"`
	Name        *string `json:"name,omitempty"`
	Priority    *string `json:"priority,omitempty"`
}

type WifiCallingPolicy struct {
	CreateDateTime   *int    `json:"createDateTime,omitempty"`
	CreatorID        *string `json:"creatorId,omitempty"`
	CreatorUsername  *string `json:"creatorUsername,omitempty"`
	Description      *string `json:"description,omitempty"`
	DomainID         *string `json:"domainId,omitempty"`
	Epdgs            []*Epdg `json:"epdgs,omitempty"`
	ID               *string `json:"id,omitempty"`
	ModifiedDateTime *int    `json:"modifiedDateTime,omitempty"`
	ModifierID       *string `json:"modifierId,omitempty"`
	ModifierUsername *string `json:"modifierUsername,omitempty"`
	Name             *string `json:"name,omitempty"`
	Priority         *string `json:"priority,omitempty"`
	TenantID         *string `json:"tenantId,omitempty"`
}

type WifiCallingPolicyList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*WifiCallingPolicy `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}
