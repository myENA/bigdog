package portaldetectionprofile

// API Version: v8_0

type CreatePortalDetectionProfile struct {
	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`

	// PortalDetectionPatterns
	// The pattern profiles for portal detection and suppression
	PortalDetectionPatterns []*PortalDetectionPattern `json:"portalDetectionPatterns,omitempty"`
}

type PortalDetectionPattern struct {
	// HTTPCode
	// HTTP status codes
	HTTPCode *int `json:"httpCode,omitempty"`

	// HTTPResponseBody
	// HTTP response body
	HTTPResponseBody *string `json:"httpResponseBody,omitempty"`

	Name *string `json:"name,omitempty"`

	// PatternType
	// Portal detection and suppression pattern type
	PatternType *string `json:"patternType,omitempty"`

	// UserAgentPattern
	// Portal detection and suppression pattern for user agent
	UserAgentPattern *string `json:"userAgentPattern,omitempty"`
}

type PortalDetectionProfile struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorID
	// Creator ID
	CreatorID *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *string `json:"description,omitempty"`

	// ID
	// Identifier of the portal detection and suppression profile
	ID *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierID
	// Modifier ID
	ModifierID *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *string `json:"name,omitempty"`

	// PortalDetectionPatterns
	// The pattern profiles for portal detection and suppression
	PortalDetectionPatterns []*PortalDetectionPattern `json:"portalDetectionPatterns,omitempty"`

	// ZoneID
	// Zone ID
	ZoneID *string `json:"zoneId,omitempty"`
}

type PortalDetectionProfileList struct {
	Extra *common.RBACMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*PortalDetectionProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}
