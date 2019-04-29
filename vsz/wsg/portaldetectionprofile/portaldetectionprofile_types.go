package portaldetectionprofile

// API Version: v8_0

type CreatePortalDetectionProfile struct {
	Description             *string                    `json:"description,omitempty"`
	Name                    *string                    `json:"name,omitempty"`
	PortalDetectionPatterns []*PortalDetectionPatterns `json:"portalDetectionPatterns,omitempty"`
}

type PortalDetectionPattern struct {
	HTTPCode         *int    `json:"httpCode,omitempty"`
	HTTPResponseBody *string `json:"httpResponseBody,omitempty"`
	Name             *string `json:"name,omitempty"`
	PatternType      *string `json:"patternType,omitempty"`
	UserAgentPattern *string `json:"userAgentPattern,omitempty"`
}

type PortalDetectionProfile struct {
	CreateDateTime          *int                       `json:"createDateTime,omitempty"`
	CreatorID               *string                    `json:"creatorId,omitempty"`
	CreatorUsername         *string                    `json:"creatorUsername,omitempty"`
	Description             *string                    `json:"description,omitempty"`
	ID                      *string                    `json:"id,omitempty"`
	ModifiedDateTime        *int                       `json:"modifiedDateTime,omitempty"`
	ModifierID              *string                    `json:"modifierId,omitempty"`
	ModifierUsername        *string                    `json:"modifierUsername,omitempty"`
	Name                    *string                    `json:"name,omitempty"`
	PortalDetectionPatterns []*PortalDetectionPatterns `json:"portalDetectionPatterns,omitempty"`
	ZoneID                  *string                    `json:"zoneId,omitempty"`
}

type PortalDetectionProfileList struct {
	Extra      *common.RBACMetadata `json:"extra,omitempty"`
	FirstIndex *int                 `json:"firstIndex,omitempty"`
	HasMore    *bool                `json:"hasMore,omitempty"`
	List       []*List              `json:"list,omitempty"`
	TotalCount *int                 `json:"totalCount,omitempty"`
}
