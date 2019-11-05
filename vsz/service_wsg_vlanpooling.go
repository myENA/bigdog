package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/vlanpooling"
	"gopkg.in/go-playground/validator.v9"
)

type WSGVlanPoolingService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGVlanPoolingService(c *APIClient) *WSGVlanPoolingService {
	s := new(WSGVlanPoolingService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGVlanPoolingService() *WSGVlanPoolingService {
	return NewWSGVlanPoolingService(ss.apiClient)
}

// AddVlanpoolings
//
// Use this API command to create new VLAN pooling.
//
// Request Body:
//	 - body *vlanpooling.CreateVlanPooling
func (s *WSGVlanPoolingService) AddVlanpoolings(ctx context.Context, body *vlanpooling.CreateVlanPooling) (*common.CreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}

// DeleteVlanpoolings
//
// Use this API command to bulk delete VLAN pooling.
//
// Request Body:
//	 - body *vlanpooling.DeleteBulkVlanPooling
func (s *WSGVlanPoolingService) DeleteVlanpoolings(ctx context.Context, body *vlanpooling.DeleteBulkVlanPooling) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
}

// DeleteVlanpoolingsById
//
// Use this API command to delete VLAN pooling.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGVlanPoolingService) DeleteVlanpoolingsById(ctx context.Context, pId string) (*common.EmptyResult, error) {
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
func (s *WSGVlanPoolingService) FindVlanpoolingsById(ctx context.Context, pId string) (*vlanpooling.VlanPooling, error) {
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
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGVlanPoolingService) FindVlanpoolingsByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*vlanpooling.VlanPoolingList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
}

// PartialUpdateVlanpoolingsById
//
// Use this API command to modify the basic information on VLAN pooling.
//
// Request Body:
//	 - body *vlanpooling.ModifyVlanPooling
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGVlanPoolingService) PartialUpdateVlanpoolingsById(ctx context.Context, body *vlanpooling.ModifyVlanPooling, pId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}
