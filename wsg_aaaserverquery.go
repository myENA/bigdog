package ruckus

// API Version: v9_0

type WSGAAAServerQueryList struct {
	Extra *WSGAAAServerQueryList `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAAAServerQueryList `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAAAServerQueryList() *WSGAAAServerQueryList {
	m := new(WSGAAAServerQueryList)
	return m
}

type WSGAAAServerQueryCreateAaaServer struct {
	AdminDomainName *string `json:"adminDomainName,omitempty"`

	// AuthType
	// Constraints:
	//    - oneof:[WSG,WLAN]
	AuthType *string `json:"authType,omitempty"`

	CreateOn *int `json:"createOn,omitempty"`

	CreatorUUID *string `json:"creatorUUID,omitempty"`

	Description *string `json:"description,omitempty"`

	DomainId *string `json:"domainId,omitempty"`

	DomainName *string `json:"domainName,omitempty"`

	EnableSecondaryRadius *int `json:"enableSecondaryRadius,omitempty"`

	GlobalCatalog *bool `json:"globalCatalog,omitempty"`

	Id *string `json:"id,omitempty"`

	Ip *string `json:"ip,omitempty"`

	Ipv6 *string `json:"ipv6,omitempty"`

	IsConflict *int `json:"isConflict,omitempty"`

	Key *string `json:"key,omitempty"`

	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *string `json:"name,omitempty"`

	Port *int `json:"port,omitempty"`

	RadiusIP *string `json:"radiusIP,omitempty"`

	RadiusIPv6 *string `json:"radiusIPv6,omitempty"`

	RadiusPort *int `json:"radiusPort,omitempty"`

	RadiusRealm *string `json:"radiusRealm,omitempty"`

	SecondaryRadiusIP *string `json:"secondaryRadiusIP,omitempty"`

	SecondaryRadiusIPv6 *string `json:"secondaryRadiusIPv6,omitempty"`

	SecondaryRadiusPort *int `json:"secondaryRadiusPort,omitempty"`

	TacacsService *string `json:"tacacsService,omitempty"`

	TenantUUID *string `json:"tenantUUID,omitempty"`

	// Type
	// Constraints:
	//    - oneof:[RADIUS,AD,LDAP,RADIUSAcct,TACACS]
	Type *string `json:"type,omitempty"`

	WindowsDomainName *string `json:"windowsDomainName,omitempty"`

	ZoneUUID *string `json:"zoneUUID,omitempty"`
}

func NewWSGAAAServerQueryCreateAaaServer() *WSGAAAServerQueryCreateAaaServer {
	m := new(WSGAAAServerQueryCreateAaaServer)
	return m
}
