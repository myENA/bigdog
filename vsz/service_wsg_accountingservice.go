package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/service"
	"gopkg.in/go-playground/validator.v9"
)

type WSGAccountingServiceService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGAccountingServiceService(c *APIClient) *WSGAccountingServiceService {
	s := new(WSGAccountingServiceService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGAccountingServiceService() *WSGAccountingServiceService {
	return NewWSGAccountingServiceService(ss.apiClient)
}

// AddServicesAcctRadius
//
// Use this API command to create a new RADIUS accounting service.
//
// Request Body:
//	 - body *service.CreateRadiusAccounting
func (s *WSGAccountingServiceService) AddServicesAcctRadius(ctx context.Context, body *service.CreateRadiusAccounting) (*common.CreateResult, error) {
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

// AddServicesAcctTestById
//
// Use this API command to test an accounting service.
//
// Request Body:
//	 - body *service.TestingConfig
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAccountingServiceService) AddServicesAcctTestById(ctx context.Context, body *service.TestingConfig, pId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return errors.New("body cannot be empty")
	}
}

// DeleteServicesAcct
//
// Use this API command to delete a list of accounting service.
//
// Request Body:
//	 - body *service.DeleteBulkAccountingService
func (s *WSGAccountingServiceService) DeleteServicesAcct(ctx context.Context, body *service.DeleteBulkAccountingService) (*common.EmptyResult, error) {
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

// DeleteServicesAcctById
//
// Use this API command to delete an accounting service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAccountingServiceService) DeleteServicesAcctById(ctx context.Context, pId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteServicesAcctRadiusById
//
// Use this API command to delete a RADIUS accounting service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAccountingServiceService) DeleteServicesAcctRadiusById(ctx context.Context, pId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteServicesAcctRadiusSecondaryById
//
// Use this API command to disable secondary RADIUS server of a RADIUS accounting service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAccountingServiceService) DeleteServicesAcctRadiusSecondaryById(ctx context.Context, pId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteServicesAcctRadiusStandbyPrimaryById
//
// Use this API command to disable Standby secondary RADIUS server of a RADIUS Accounting service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAccountingServiceService) DeleteServicesAcctRadiusStandbyPrimaryById(ctx context.Context, pId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindServicesAcctByQueryCriteria
//
// Use this API command to retrieve a list of accounting services by query criteria.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGAccountingServiceService) FindServicesAcctByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*service.CommonAccountingServiceList, error) {
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

// FindServicesAcctRadius
//
// Use this API command to retrieve a list of RADIUS accounting services.
func (s *WSGAccountingServiceService) FindServicesAcctRadius(ctx context.Context) (*service.RadiusAccountingServiceList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindServicesAcctRadiusById
//
// Use this API command to retrieve a RADIUS accounting service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAccountingServiceService) FindServicesAcctRadiusById(ctx context.Context, pId string) (*service.RadiusAccountingService, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindServicesAcctRadiusByQueryCriteria
//
// Use this API command to retrieve a list of Radius accounting services by query criteria.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGAccountingServiceService) FindServicesAcctRadiusByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*service.RadiusAccountingServiceList, error) {
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

// PartialUpdateServicesAcctRadiusById
//
// Use this API command to modify the basic information of a RADIUS accounting service.
//
// Request Body:
//	 - body *service.ModifyRadiusAccounting
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAccountingServiceService) PartialUpdateServicesAcctRadiusById(ctx context.Context, body *service.ModifyRadiusAccounting, pId string) (*common.EmptyResult, error) {
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
