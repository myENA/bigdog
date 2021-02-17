package bigdog

// API Version: v9_1

import (
	"errors"
	"io"
)

// WSGPortalDetectionProfileCreatePortalDetectionProfile
//
// Definition: portalDetectionProfile_createPortalDetectionProfile
type WSGPortalDetectionProfileCreatePortalDetectionProfile struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// PortalDetectionPatterns
	// The pattern profiles for portal detection and suppression
	PortalDetectionPatterns []*WSGPortalDetectionProfilePortalDetectionPattern `json:"portalDetectionPatterns,omitempty"`
}

func NewWSGPortalDetectionProfileCreatePortalDetectionProfile() *WSGPortalDetectionProfileCreatePortalDetectionProfile {
	m := new(WSGPortalDetectionProfileCreatePortalDetectionProfile)
	return m
}

// WSGPortalDetectionProfilePortalDetectionPattern
//
// Definition: portalDetectionProfile_portalDetectionPattern
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
	Name *WSGCommonNormalName `json:"name"`

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

// WSGPortalDetectionProfile
//
// Definition: portalDetectionProfile_portalDetectionProfile
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

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// PortalDetectionPatterns
	// The pattern profiles for portal detection and suppression
	PortalDetectionPatterns []*WSGPortalDetectionProfilePortalDetectionPattern `json:"portalDetectionPatterns,omitempty"`

	// ZoneId
	// Zone ID
	ZoneId *string `json:"zoneId,omitempty"`
}

type WSGPortalDetectionProfileAPIResponse struct {
	*RawAPIResponse
	Data *WSGPortalDetectionProfile
}

func newWSGPortalDetectionProfileAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGPortalDetectionProfileAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGPortalDetectionProfileAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGPortalDetectionProfile)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGPortalDetectionProfile() *WSGPortalDetectionProfile {
	m := new(WSGPortalDetectionProfile)
	return m
}

// WSGPortalDetectionProfileList
//
// Definition: portalDetectionProfile_portalDetectionProfileList
type WSGPortalDetectionProfileList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGPortalDetectionProfile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGPortalDetectionProfileListAPIResponse struct {
	*RawAPIResponse
	Data *WSGPortalDetectionProfileList
}

func newWSGPortalDetectionProfileListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGPortalDetectionProfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGPortalDetectionProfileListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGPortalDetectionProfileList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGPortalDetectionProfileList() *WSGPortalDetectionProfileList {
	m := new(WSGPortalDetectionProfileList)
	return m
}
