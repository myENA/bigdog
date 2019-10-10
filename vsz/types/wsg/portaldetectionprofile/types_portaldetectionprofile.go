package portaldetectionprofile

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type CreatePortalDetectionProfile struct {
	Description *common.Description `json:"description,omitempty"`

	Name *common.NormalName `json:"name,omitempty" validate:"required"`

	// PortalDetectionPatterns
	// The pattern profiles for portal detection and suppression
	PortalDetectionPatterns []*PortalDetectionPattern `json:"portalDetectionPatterns,omitempty"`
}

type PortalDetectionPattern struct {
	// HttpCode
	// HTTP status codes
	HttpCode *int `json:"httpCode,omitempty" validate:"required,gte=100,lte=599"`

	// HttpResponseBody
	// HTTP response body
	HttpResponseBody *string `json:"httpResponseBody,omitempty" validate:"max=1024"`

	Name *common.NormalName `json:"name,omitempty" validate:"required"`

	// PatternType
	// Portal detection and suppression pattern type
	PatternType *string `json:"patternType,omitempty" validate:"oneof=USER_AGENT"`

	// UserAgentPattern
	// Portal detection and suppression pattern for user agent
	UserAgentPattern *string `json:"userAgentPattern,omitempty" validate:"required,max=256"`
}

type PortalDetectionProfile struct {
	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// Id
	// Identifier of the portal detection and suppression profile
	Id *string `json:"id,omitempty" validate:"max=64"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// PortalDetectionPatterns
	// The pattern profiles for portal detection and suppression
	PortalDetectionPatterns []*PortalDetectionPattern `json:"portalDetectionPatterns,omitempty"`

	// ZoneId
	// Zone ID
	ZoneId *string `json:"zoneId,omitempty"`
}

type PortalDetectionProfileList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*PortalDetectionProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}
