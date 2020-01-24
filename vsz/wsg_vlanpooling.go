package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
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
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	// Pool
	// VLANs of the VLAN pooling profile
	// Constraints:
	//    - required
	Pool *string `json:"pool" validate:"required"`
}

func NewWSGVLANPoolingCreateVlanPooling() *WSGVLANPoolingCreateVlanPooling {
	m := new(WSGVLANPoolingCreateVlanPooling)
	return m
}

type WSGVLANPoolingDeleteBulkVlanPooling struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGVLANPoolingDeleteBulkVlanPooling() *WSGVLANPoolingDeleteBulkVlanPooling {
	m := new(WSGVLANPoolingDeleteBulkVlanPooling)
	return m
}

type WSGVLANPoolingModifyVlanPooling struct {
	// Algo
	// Algorithm of the VLAN pooling profile
	// Constraints:
	//    - oneof:[MAC_HASH]
	Algo *string `json:"algo,omitempty" validate:"oneof=MAC_HASH"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// identifier of the domain
	DomainId *string `json:"domainId,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Pool
	// VLANs of the VLAN pooling profile
	Pool *string `json:"pool,omitempty"`
}

func NewWSGVLANPoolingModifyVlanPooling() *WSGVLANPoolingModifyVlanPooling {
	m := new(WSGVLANPoolingModifyVlanPooling)
	return m
}

type WSGVLANPooling struct {
	// Algo
	// Algorithm of the VLAN pooling profile
	// Constraints:
	//    - oneof:[MAC_HASH]
	Algo *string `json:"algo,omitempty" validate:"oneof=MAC_HASH"`

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

func NewWSGVLANPooling() *WSGVLANPooling {
	m := new(WSGVLANPooling)
	return m
}

type WSGVLANPoolingList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGVLANPoolingListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGVLANPoolingList() *WSGVLANPoolingList {
	m := new(WSGVLANPoolingList)
	return m
}

type WSGVLANPoolingListType struct {
	// Algo
	// Algorithm of the VLAN pooling profile
	// Constraints:
	//    - default:'MAC_HASH'
	//    - oneof:[MAC_HASH]
	Algo *string `json:"algo,omitempty" validate:"oneof=MAC_HASH"`

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

func NewWSGVLANPoolingListType() *WSGVLANPoolingListType {
	m := new(WSGVLANPoolingListType)
	return m
}

// AddVlanpoolings
//
// Use this API command to create new VLAN pooling.
//
// Request Body:
//	 - body *WSGVLANPoolingCreateVlanPooling
func (s *WSGVLANPoolingService) AddVlanpoolings(ctx context.Context, body *WSGVLANPoolingCreateVlanPooling) (*WSGCommonCreateResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonCreateResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddVlanpoolings, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteVlanpoolings
//
// Use this API command to bulk delete VLAN pooling.
//
// Request Body:
//	 - body *WSGVLANPoolingDeleteBulkVlanPooling
func (s *WSGVLANPoolingService) DeleteVlanpoolings(ctx context.Context, body *WSGVLANPoolingDeleteBulkVlanPooling) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteVlanpoolings, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteVlanpoolingsById
//
// Use this API command to delete VLAN pooling.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGVLANPoolingService) DeleteVlanpoolingsById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteVlanpoolingsById, true)
}

// FindVlanpoolingsById
//
// Use this API command to retrieve VLAN pooling.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGVLANPoolingService) FindVlanpoolingsById(ctx context.Context, id string) (*WSGVLANPooling, error) {
	var (
		req  *APIRequest
		resp *WSGVLANPooling
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindVlanpoolingsById, true)
}

// FindVlanpoolingsByQueryCriteria
//
// Use this API command to retrieve a list of VLAN poolings.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGVLANPoolingService) FindVlanpoolingsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGVLANPoolingList, error) {
	var (
		req  *APIRequest
		resp *WSGVLANPoolingList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindVlanpoolingsByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// PartialUpdateVlanpoolingsById
//
// Use this API command to modify the basic information on VLAN pooling.
//
// Request Body:
//	 - body *WSGVLANPoolingModifyVlanPooling
//
// Required Parameters:
// - id string
//		- required
func (s *WSGVLANPoolingService) PartialUpdateVlanpoolingsById(ctx context.Context, body *WSGVLANPoolingModifyVlanPooling, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateVlanpoolingsById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}
