package vsz

// API Version: v9_0

type WSGPortalDetectionProfileCreatePortalDetectionProfile struct {
	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// PortalDetectionPatterns
	// The pattern profiles for portal detection and suppression
	// Constraints:
	//    - nullable
	PortalDetectionPatterns []*WSGPortalDetectionProfilePortalDetectionPattern `json:"portalDetectionPatterns,omitempty" validate:"omitempty,dive"`
}

func NewWSGPortalDetectionProfileCreatePortalDetectionProfile() *WSGPortalDetectionProfileCreatePortalDetectionProfile {
	m := new(WSGPortalDetectionProfileCreatePortalDetectionProfile)
	return m
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

func NewWSGPortalDetectionProfilePortalDetectionPattern() *WSGPortalDetectionProfilePortalDetectionPattern {
	m := new(WSGPortalDetectionProfilePortalDetectionPattern)
	return m
}

type WSGPortalDetectionProfile struct {
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
	// Creator name
	// Constraints:
	//    - nullable
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// Constraints:
	//    - nullable
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the portal detection and suppression profile
	// Constraints:
	//    - nullable
	//    - max:64
	Id *string `json:"id,omitempty" validate:"omitempty,max=64"`

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
	// Modifier name
	// Constraints:
	//    - nullable
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// Constraints:
	//    - nullable
	Name *WSGCommonNormalName `json:"name,omitempty"`

	// PortalDetectionPatterns
	// The pattern profiles for portal detection and suppression
	// Constraints:
	//    - nullable
	PortalDetectionPatterns []*WSGPortalDetectionProfilePortalDetectionPattern `json:"portalDetectionPatterns,omitempty" validate:"omitempty,dive"`

	// ZoneId
	// Zone ID
	// Constraints:
	//    - nullable
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGPortalDetectionProfile() *WSGPortalDetectionProfile {
	m := new(WSGPortalDetectionProfile)
	return m
}

type WSGPortalDetectionProfileList struct {
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
	List []*WSGPortalDetectionProfile `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGPortalDetectionProfileList() *WSGPortalDetectionProfileList {
	m := new(WSGPortalDetectionProfileList)
	return m
}
