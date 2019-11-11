package vsz

// API Version: v8_1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGVendorSpecificAttributeProfileService struct {
	apiClient *APIClient
}

func NewWSGVendorSpecificAttributeProfileService(c *APIClient) *WSGVendorSpecificAttributeProfileService {
	s := new(WSGVendorSpecificAttributeProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGVendorSpecificAttributeProfileService() *WSGVendorSpecificAttributeProfileService {
	return NewWSGVendorSpecificAttributeProfileService(ss.apiClient)
}

type WSGVendorSpecificAttributeProfileCreateResult interface{}

type WSGVendorSpecificAttributeProfileDeleteBulk struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

type WSGVendorSpecificAttributeProfileEmptyResult struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGVendorSpecificAttributeProfileEmptyResult) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGVendorSpecificAttributeProfileEmptyResult{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGVendorSpecificAttributeProfileEmptyResult) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

type WSGVendorSpecificAttributeProfileGet struct {
	// Attributes
	// Vendor specific attribute list for Radius protocol
	Attributes []*WSGVendorSpecificAttributeProfileVendorSpecificAttribute `json:"attributes,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of the vendor specific attribute profile
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// ZoneId
	// Zone Id
	ZoneId *string `json:"zoneId,omitempty"`
}

type WSGVendorSpecificAttributeProfileList struct {
	// FirstIndex
	// Index of the first profile returned out of the profile list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more profiles after the list that is currently displayed
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Information list of the vendor specific attribute profile
	List []*WSGVendorSpecificAttributeProfileListType `json:"list,omitempty"`

	// TotalCount
	// Total number of the profiles
	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGVendorSpecificAttributeProfileListType struct {
	// Id
	// Identifier of the vendor specific attribute profile
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

type WSGVendorSpecificAttributeProfilePersist struct {
	// Attributes
	// Vendor specific attribute list for Radius protocol
	// Constraints:
	//    - required
	Attributes []*WSGVendorSpecificAttributeProfileVendorSpecificAttribute `json:"attributes" validate:"required,dive,required"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`
}

type WSGVendorSpecificAttributeProfileQueryCriteriaResult struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGVendorSpecificAttributeProfileGet `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGVendorSpecificAttributeProfileVendorSpecificAttribute struct {
	// KeyId
	// Key ID of vendor specific attribute
	// Constraints:
	//    - required
	KeyId *int `json:"keyId" validate:"required"`

	// SupportedRadiusProtocol
	// The radius protocol to which this given vendor specific attribute will attach
	// Constraints:
	//    - required
	//    - oneof:[ACCOUNTING,AUTHENTICATION,BOTH_AUTHENTICATION_AND_ACCOUNTING]
	SupportedRadiusProtocol *string `json:"supportedRadiusProtocol" validate:"required,oneof=ACCOUNTING AUTHENTICATION BOTH_AUTHENTICATION_AND_ACCOUNTING"`

	// Type
	// Type of vendor specific attribute
	// Constraints:
	//    - required
	//    - oneof:[INTEGER,STRING]
	Type *string `json:"type" validate:"required,oneof=INTEGER STRING"`

	// Value
	// Value of vendor specific attribute
	// Constraints:
	//    - required
	Value *string `json:"value" validate:"required"`

	// VendorId
	// Vendor ID of vendor specific attribute
	// Constraints:
	//    - required
	VendorId *int `json:"vendorId" validate:"required"`
}

// AddRkszonesVendorSpecificAttributeProfilesByZoneId
//
// Create a vendor specific attribute profile.
//
// Request Body:
//	 - body *WSGVendorSpecificAttributeProfilePersist
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGVendorSpecificAttributeProfileService) AddRkszonesVendorSpecificAttributeProfilesByZoneId(ctx context.Context, body *WSGVendorSpecificAttributeProfilePersist, pZoneId string) (WSGVendorSpecificAttributeProfileCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesVendorSpecificAttributeProfilesById
//
// Use this API command to delete a vendor specific attribute profile by ID.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGVendorSpecificAttributeProfileService) DeleteRkszonesVendorSpecificAttributeProfilesById(ctx context.Context, pId string, pZoneId string) (*WSGVendorSpecificAttributeProfileEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesVendorSpecificAttributeProfilesByZoneId
//
// Use this API command to delete a list of vendor specific attribute profile.
//
// Request Body:
//	 - body *WSGVendorSpecificAttributeProfileDeleteBulk
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGVendorSpecificAttributeProfileService) DeleteRkszonesVendorSpecificAttributeProfilesByZoneId(ctx context.Context, body *WSGVendorSpecificAttributeProfileDeleteBulk, pZoneId string) (*WSGVendorSpecificAttributeProfileEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesVendorSpecificAttributeProfilesById
//
// Get a vendor specific attribute profile by ID.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGVendorSpecificAttributeProfileService) FindRkszonesVendorSpecificAttributeProfilesById(ctx context.Context, pId string, pZoneId string) (*WSGVendorSpecificAttributeProfileGet, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesVendorSpecificAttributeProfilesByQueryCriteria
//
// Use this API command to retrieve a list of vendor specific attribute profile by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGVendorSpecificAttributeProfileService) FindRkszonesVendorSpecificAttributeProfilesByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGVendorSpecificAttributeProfileQueryCriteriaResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesVendorSpecificAttributeProfilesByZoneId
//
// Get a ID list of vendor specific attribute profile in this Zone.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGVendorSpecificAttributeProfileService) FindRkszonesVendorSpecificAttributeProfilesByZoneId(ctx context.Context, pZoneId string) (*WSGVendorSpecificAttributeProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateRkszonesVendorSpecificAttributeProfilesById
//
// Use this API command to modify entire information of a vendor specific attribute profile.
//
// Request Body:
//	 - body *WSGVendorSpecificAttributeProfilePersist
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGVendorSpecificAttributeProfileService) UpdateRkszonesVendorSpecificAttributeProfilesById(ctx context.Context, body *WSGVendorSpecificAttributeProfilePersist, pId string, pZoneId string) (*WSGVendorSpecificAttributeProfileEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
