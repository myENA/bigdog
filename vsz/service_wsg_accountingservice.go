package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/service"
)

type WSGAccountingServiceService struct {
	apiClient *APIClient
}

func NewWSGAccountingServiceService(c *APIClient) *WSGAccountingServiceService {
	s := new(WSGAccountingServiceService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAccountingServiceService() *WSGAccountingServiceService {
	serv := new(WSGAccountingServiceService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddServicesAcctRadius
//
// Use this API command to create a new RADIUS accounting service.
//
// Request Body:
//	 - body *service.CreateRadiusAccounting
func (s *WSGAccountingServiceService) AddServicesAcctRadius(ctx context.Context, body *service.CreateRadiusAccounting) (*common.CreateResult, error) {
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
}

// DeleteServicesAcct
//
// Use this API command to delete a list of accounting service.
//
// Request Body:
//	 - body *service.DeleteBulkAccountingService
func (s *WSGAccountingServiceService) DeleteServicesAcct(ctx context.Context, body *service.DeleteBulkAccountingService) (*common.EmptyResult, error) {
}

// DeleteServicesAcctById
//
// Use this API command to delete an accounting service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAccountingServiceService) DeleteServicesAcctById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteServicesAcctRadiusById
//
// Use this API command to delete a RADIUS accounting service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAccountingServiceService) DeleteServicesAcctRadiusById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteServicesAcctRadiusSecondaryById
//
// Use this API command to disable secondary RADIUS server of a RADIUS accounting service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAccountingServiceService) DeleteServicesAcctRadiusSecondaryById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteServicesAcctRadiusStandbyPrimaryById
//
// Use this API command to disable Standby secondary RADIUS server of a RADIUS Accounting service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAccountingServiceService) DeleteServicesAcctRadiusStandbyPrimaryById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// FindServicesAcctByQueryCriteria
//
// Use this API command to retrieve a list of accounting services by query criteria.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGAccountingServiceService) FindServicesAcctByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*service.CommonAccountingServiceList, error) {
}

// FindServicesAcctRadius
//
// Use this API command to retrieve a list of RADIUS accounting services.
func (s *WSGAccountingServiceService) FindServicesAcctRadius(ctx context.Context) (*service.RadiusAccountingServiceList, error) {
}

// FindServicesAcctRadiusById
//
// Use this API command to retrieve a RADIUS accounting service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAccountingServiceService) FindServicesAcctRadiusById(ctx context.Context, pId string) (*service.RadiusAccountingService, error) {
}

// FindServicesAcctRadiusByQueryCriteria
//
// Use this API command to retrieve a list of Radius accounting services by query criteria.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGAccountingServiceService) FindServicesAcctRadiusByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*service.RadiusAccountingServiceList, error) {
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
}
