package vsz

// API Version: v9_0

type WSGAAAServerQueryList struct {
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
	List []*WSGAAAServerQueryCreateAaaServer `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAAAServerQueryList() *WSGAAAServerQueryList {
	m := new(WSGAAAServerQueryList)
	return m
}

type WSGAAAServerQueryCreateAaaServer struct {
	// AdminDomainName
	// Constraints:
	//    - nullable
	AdminDomainName *string `json:"adminDomainName,omitempty"`

	// AuthType
	// Constraints:
	//    - nullable
	//    - oneof:[WSG,WLAN]
	AuthType *string `json:"authType,omitempty" validate:"omitempty,oneof=WSG WLAN"`

	// CreateOn
	// Constraints:
	//    - nullable
	CreateOn *int `json:"createOn,omitempty"`

	// CreatorUUID
	// Constraints:
	//    - nullable
	CreatorUUID *string `json:"creatorUUID,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *string `json:"description,omitempty"`

	// DomainId
	// Constraints:
	//    - nullable
	DomainId *string `json:"domainId,omitempty"`

	// DomainName
	// Constraints:
	//    - nullable
	DomainName *string `json:"domainName,omitempty"`

	// EnableSecondaryRadius
	// Constraints:
	//    - nullable
	EnableSecondaryRadius *int `json:"enableSecondaryRadius,omitempty"`

	// GlobalCatalog
	// Constraints:
	//    - nullable
	GlobalCatalog *bool `json:"globalCatalog,omitempty"`

	// Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// Ip
	// Constraints:
	//    - nullable
	Ip *string `json:"ip,omitempty"`

	// Ipv6
	// Constraints:
	//    - nullable
	Ipv6 *string `json:"ipv6,omitempty"`

	// IsConflict
	// Constraints:
	//    - nullable
	IsConflict *int `json:"isConflict,omitempty"`

	// Key
	// Constraints:
	//    - nullable
	Key *string `json:"key,omitempty"`

	// ModifiedDateTime
	// Constraints:
	//    - nullable
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierUsername
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *string `json:"name,omitempty"`

	// Port
	// Constraints:
	//    - nullable
	Port *int `json:"port,omitempty"`

	// RadiusIP
	// Constraints:
	//    - nullable
	RadiusIP *string `json:"radiusIP,omitempty"`

	// RadiusIPv6
	// Constraints:
	//    - nullable
	RadiusIPv6 *string `json:"radiusIPv6,omitempty"`

	// RadiusPort
	// Constraints:
	//    - nullable
	RadiusPort *int `json:"radiusPort,omitempty"`

	// RadiusRealm
	// Constraints:
	//    - nullable
	RadiusRealm *string `json:"radiusRealm,omitempty"`

	// SecondaryRadiusIP
	// Constraints:
	//    - nullable
	SecondaryRadiusIP *string `json:"secondaryRadiusIP,omitempty"`

	// SecondaryRadiusIPv6
	// Constraints:
	//    - nullable
	SecondaryRadiusIPv6 *string `json:"secondaryRadiusIPv6,omitempty"`

	// SecondaryRadiusPort
	// Constraints:
	//    - nullable
	SecondaryRadiusPort *int `json:"secondaryRadiusPort,omitempty"`

	// TacacsService
	// Constraints:
	//    - nullable
	TacacsService *string `json:"tacacsService,omitempty"`

	// TenantUUID
	// Constraints:
	//    - nullable
	TenantUUID *string `json:"tenantUUID,omitempty"`

	// Type
	// Constraints:
	//    - nullable
	//    - oneof:[RADIUS,AD,LDAP,RADIUSAcct,TACACS]
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=RADIUS AD LDAP RADIUSAcct TACACS"`

	// WindowsDomainName
	// Constraints:
	//    - nullable
	WindowsDomainName *string `json:"windowsDomainName,omitempty"`

	// ZoneUUID
	// Constraints:
	//    - nullable
	ZoneUUID *string `json:"zoneUUID,omitempty"`
}

func NewWSGAAAServerQueryCreateAaaServer() *WSGAAAServerQueryCreateAaaServer {
	m := new(WSGAAAServerQueryCreateAaaServer)
	return m
}
