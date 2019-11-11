package vsz

// API Version: v8_1

type WSGPortalDetectionProfileCreatePortalDetectionProfile struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// PortalDetectionPatterns
	// The pattern profiles for portal detection and suppression
	PortalDetectionPatterns []*WSGPortalDetectionProfilePortalDetectionPattern `json:"portalDetectionPatterns,omitempty"`
}

type WSGPortalDetectionProfilePortalDetectionPattern struct {
	// HttpCode
	// HTTP status codes
	// Constraints:
	//    - required
	//    - min:100
	//    - max:599
	HttpCode *int `json:"httpCode" validate:"required,gte=100,lte=599"`

	// HttpResponseBody
	// HTTP response body
	// Constraints:
	//    - nullable
	//    - max:1024
	HttpResponseBody *string `json:"httpResponseBody,omitempty" validate:"omitempty,max=1024"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// PatternType
	// Portal detection and suppression pattern type
	// Constraints:
	//    - nullable
	//    - oneof:[USER_AGENT]
	PatternType *string `json:"patternType,omitempty" validate:"omitempty,oneof=USER_AGENT"`

	// UserAgentPattern
	// Portal detection and suppression pattern for user agent
	// Constraints:
	//    - required
	//    - max:256
	UserAgentPattern *string `json:"userAgentPattern" validate:"required,max=256"`
}

type WSGPortalDetectionProfile struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the portal detection and suppression profile
	// Constraints:
	//    - nullable
	//    - max:64
	Id *string `json:"id,omitempty" validate:"omitempty,max=64"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// PortalDetectionPatterns
	// The pattern profiles for portal detection and suppression
	PortalDetectionPatterns []*WSGPortalDetectionProfilePortalDetectionPattern `json:"portalDetectionPatterns,omitempty"`

	// ZoneId
	// Zone ID
	ZoneId *string `json:"zoneId,omitempty"`
}

type WSGPortalDetectionProfileList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGPortalDetectionProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}
