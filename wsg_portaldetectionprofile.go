package ruckus

// API Version: v9_0

type WSGPortalDetectionProfileCreatePortalDetectionProfile struct {
	Description **WSGPortalDetectionProfileCreatePortalDetectionProfile `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name **WSGPortalDetectionProfileCreatePortalDetectionProfile `json:"name"`

	// PortalDetectionPatterns
	// The pattern profiles for portal detection and suppression
	PortalDetectionPatterns []**WSGPortalDetectionProfileCreatePortalDetectionProfile `json:"portalDetectionPatterns,omitempty"`
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
	HttpCode *int `json:"httpCode"`

	// HttpResponseBody
	// HTTP response body
	// Constraints:
	//    - max:1024
	HttpResponseBody *string `json:"httpResponseBody,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name **WSGPortalDetectionProfilePortalDetectionPattern `json:"name"`

	// PatternType
	// Portal detection and suppression pattern type
	// Constraints:
	//    - oneof:[USER_AGENT]
	PatternType *string `json:"patternType,omitempty"`

	// UserAgentPattern
	// Portal detection and suppression pattern for user agent
	// Constraints:
	//    - required
	//    - max:256
	UserAgentPattern *string `json:"userAgentPattern"`
}

func NewWSGPortalDetectionProfilePortalDetectionPattern() *WSGPortalDetectionProfilePortalDetectionPattern {
	m := new(WSGPortalDetectionProfilePortalDetectionPattern)
	return m
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

	Description **WSGPortalDetectionProfile `json:"description,omitempty"`

	// Id
	// Identifier of the portal detection and suppression profile
	// Constraints:
	//    - max:64
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name **WSGPortalDetectionProfile `json:"name,omitempty"`

	// PortalDetectionPatterns
	// The pattern profiles for portal detection and suppression
	PortalDetectionPatterns []**WSGPortalDetectionProfile `json:"portalDetectionPatterns,omitempty"`

	// ZoneId
	// Zone ID
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGPortalDetectionProfile() *WSGPortalDetectionProfile {
	m := new(WSGPortalDetectionProfile)
	return m
}

type WSGPortalDetectionProfileList struct {
	Extra **WSGPortalDetectionProfileList `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []**WSGPortalDetectionProfileList `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGPortalDetectionProfileList() *WSGPortalDetectionProfileList {
	m := new(WSGPortalDetectionProfileList)
	return m
}
