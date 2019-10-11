package portaldetectionprofile

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type CreatePortalDetectionProfile struct {
	Description *common.Description `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	// PortalDetectionPatterns
	// The pattern profiles for portal detection and suppression
	PortalDetectionPatterns []*PortalDetectionPattern `json:"portalDetectionPatterns,omitempty"`
}

func NewCreatePortalDetectionProfile() *CreatePortalDetectionProfile {
	createPortalDetectionProfileType := new(CreatePortalDetectionProfile)
	return createPortalDetectionProfileType
}

func NewCreatePortalDetectionProfileWithDefaults() *CreatePortalDetectionProfile {
	createPortalDetectionProfileType := new(CreatePortalDetectionProfile)
	return createPortalDetectionProfileType
}

type PortalDetectionPattern struct {
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
	Name *common.NormalName `json:"name" validate:"required"`

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

func NewPortalDetectionPattern() *PortalDetectionPattern {
	portalDetectionPatternType := new(PortalDetectionPattern)
	return portalDetectionPatternType
}

func NewPortalDetectionPatternWithDefaults() *PortalDetectionPattern {
	portalDetectionPatternType := new(PortalDetectionPattern)
	return portalDetectionPatternType
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

	Name *common.NormalName `json:"name,omitempty"`

	// PortalDetectionPatterns
	// The pattern profiles for portal detection and suppression
	PortalDetectionPatterns []*PortalDetectionPattern `json:"portalDetectionPatterns,omitempty"`

	// ZoneId
	// Zone ID
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewPortalDetectionProfile() *PortalDetectionProfile {
	portalDetectionProfileType := new(PortalDetectionProfile)
	return portalDetectionProfileType
}

func NewPortalDetectionProfileWithDefaults() *PortalDetectionProfile {
	portalDetectionProfileType := new(PortalDetectionProfile)
	return portalDetectionProfileType
}

type PortalDetectionProfileList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*PortalDetectionProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewPortalDetectionProfileList() *PortalDetectionProfileList {
	portalDetectionProfileListType := new(PortalDetectionProfileList)
	return portalDetectionProfileListType
}

func NewPortalDetectionProfileListWithDefaults() *PortalDetectionProfileList {
	portalDetectionProfileListType := new(PortalDetectionProfileList)
	return portalDetectionProfileListType
}
