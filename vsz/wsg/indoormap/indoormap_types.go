package indoormap

// API Version: v8_0

type BasicIndoorMap struct {
	Address       *string  `json:"address,omitempty"`
	ApGroupID     *string  `json:"apGroupId,omitempty"`
	Description   *string  `json:"description,omitempty"`
	DomainID      *string  `json:"domainId,omitempty"`
	GroupType     *string  `json:"groupType,omitempty"`
	ID            *string  `json:"id,omitempty"`
	ImageFileName *string  `json:"imageFileName,omitempty"`
	Latitude      *float64 `json:"latitude,omitempty"`
	Longitude     *float64 `json:"longitude,omitempty"`
	Name          *string  `json:"name,omitempty"`
	Orientation   *string  `json:"orientation,omitempty"`
	Scale         *Scale   `json:"scale,omitempty"`
	TenantID      *string  `json:"tenantId,omitempty"`
	ZoneID        *string  `json:"zoneId,omitempty"`
}

type IndooMapAuditID struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type IndoorMap struct {
	Address       *string  `json:"address,omitempty"`
	ApGroupID     *string  `json:"apGroupId,omitempty"`
	Description   *string  `json:"description,omitempty"`
	DomainID      *string  `json:"domainId,omitempty"`
	GroupType     *string  `json:"groupType,omitempty"`
	ID            *string  `json:"id,omitempty"`
	ImageData     *string  `json:"imageData,omitempty"`
	ImageFileName *string  `json:"imageFileName,omitempty"`
	Latitude      *float64 `json:"latitude,omitempty"`
	Longitude     *float64 `json:"longitude,omitempty"`
	Name          *string  `json:"name,omitempty"`
	Orientation   *string  `json:"orientation,omitempty"`
	Scale         *Scale   `json:"scale,omitempty"`
	TenantID      *string  `json:"tenantId,omitempty"`
	ZoneID        *string  `json:"zoneId,omitempty"`
}

type IndoorMapAp struct {
	IndoorMapXy *IndoorMapXy `json:"indoorMapXy,omitempty"`
	Mac         *string      `json:"mac,omitempty"`
}

type IndoorMapList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}

type IndoorMapSummary struct {
	Address       *string  `json:"address,omitempty"`
	ApCount       *float64 `json:"apCount,omitempty"`
	ApGroupID     *string  `json:"apGroupId,omitempty"`
	Description   *string  `json:"description,omitempty"`
	DomainID      *string  `json:"domainId,omitempty"`
	GroupType     *string  `json:"groupType,omitempty"`
	ID            *string  `json:"id,omitempty"`
	ImageFileName *string  `json:"imageFileName,omitempty"`
	Key           *string  `json:"key,omitempty"`
	Latitude      *float64 `json:"latitude,omitempty"`
	Longitude     *float64 `json:"longitude,omitempty"`
	Name          *string  `json:"name,omitempty"`
	Scale         *Scale   `json:"scale,omitempty"`
	TenantID      *string  `json:"tenantId,omitempty"`
	ZoneID        *string  `json:"zoneId,omitempty"`
}

type IndoorMapSummaryList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}

type IndoorMapXy struct {
	X *float64 `json:"x,omitempty"`
	Y *float64 `json:"y,omitempty"`
}

type Scale struct {
	A        *A       `json:"a,omitempty"`
	B        *B       `json:"b,omitempty"`
	Distance *float64 `json:"distance,omitempty"`
	Unit     *string  `json:"unit,omitempty"`
}
