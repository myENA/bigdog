package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGVLANPoolingService struct {
	apiClient *APIClient
}

func NewWSGVLANPoolingService(c *APIClient) *WSGVLANPoolingService {
	s := new(WSGVLANPoolingService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGVLANPoolingService() *WSGVLANPoolingService {
	return NewWSGVLANPoolingService(ss.apiClient)
}

type WSGVLANPoolingCreateVlanPooling struct {
	// Algo
	// Algorithm of the VLAN pooling profile
	// Constraints:
	//    - required
	//    - default:'MAC_HASH'
	//    - oneof:[MAC_HASH]
	Algo *string `json:"algo" validate:"required,oneof=MAC_HASH"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// identifier of the domain
	DomainId *string `json:"domainId,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required"`

	// Pool
	// VLANs of the VLAN pooling profile
	// Constraints:
	//    - required
	Pool *string `json:"pool" validate:"required"`
}

type WSGVLANPoolingDeleteBulkVlanPooling struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

type WSGVLANPoolingModifyVlanPooling struct {
	// Algo
	// Algorithm of the VLAN pooling profile
	// Constraints:
	//    - nullable
	//    - oneof:[MAC_HASH]
	Algo *string `json:"algo,omitempty" validate:"omitempty,oneof=MAC_HASH"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// identifier of the domain
	DomainId *string `json:"domainId,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Pool
	// VLANs of the VLAN pooling profile
	Pool *string `json:"pool,omitempty"`
}

type WSGVLANPooling struct {
	// Algo
	// Algorithm of the VLAN pooling profile
	// Constraints:
	//    - nullable
	//    - oneof:[MAC_HASH]
	Algo *string `json:"algo,omitempty" validate:"omitempty,oneof=MAC_HASH"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Identifier of the domain which the VLAN pooling profile belongs to
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of the VLAN pooling profile
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Pool
	// VLANs of the VLAN pooling profile
	Pool *string `json:"pool,omitempty"`
}

type WSGVLANPoolingList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGVLANPoolingListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGVLANPoolingListType struct {
	// Algo
	// Algorithm of the VLAN pooling profile
	// Constraints:
	//    - nullable
	//    - default:'MAC_HASH'
	//    - oneof:[MAC_HASH]
	Algo *string `json:"algo,omitempty" validate:"omitempty,oneof=MAC_HASH"`

	// Description
	// Description of the service
	Description *string `json:"description,omitempty"`

	// DomainId
	// identifier of the domain
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of the service
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Pool
	// VLANs of the VLAN pooling profile
	Pool *string `json:"pool,omitempty"`
}

// AddVlanpoolings
//
// Use this API command to create new VLAN pooling.
//
// Request Body:
//	 - body *WSGVLANPoolingCreateVlanPooling
func (s *WSGVLANPoolingService) AddVlanpoolings(ctx context.Context, body *WSGVLANPoolingCreateVlanPooling) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteVlanpoolings
//
// Use this API command to bulk delete VLAN pooling.
//
// Request Body:
//	 - body *WSGVLANPoolingDeleteBulkVlanPooling
func (s *WSGVLANPoolingService) DeleteVlanpoolings(ctx context.Context, body *WSGVLANPoolingDeleteBulkVlanPooling) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteVlanpoolingsById
//
// Use this API command to delete VLAN pooling.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGVLANPoolingService) DeleteVlanpoolingsById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindVlanpoolingsById
//
// Use this API command to retrieve VLAN pooling.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGVLANPoolingService) FindVlanpoolingsById(ctx context.Context, pId string) (*WSGVLANPooling, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindVlanpoolingsByQueryCriteria
//
// Use this API command to retrieve a list of VLAN poolings.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGVLANPoolingService) FindVlanpoolingsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGVLANPoolingList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateVlanpoolingsById
//
// Use this API command to modify the basic information on VLAN pooling.
//
// Request Body:
//	 - body *WSGVLANPoolingModifyVlanPooling
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGVLANPoolingService) PartialUpdateVlanpoolingsById(ctx context.Context, body *WSGVLANPoolingModifyVlanPooling, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
