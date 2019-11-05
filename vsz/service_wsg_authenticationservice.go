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

type WSGAuthenticationServiceService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGAuthenticationServiceService(c *APIClient) *WSGAuthenticationServiceService {
	s := new(WSGAuthenticationServiceService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGAuthenticationServiceService() *WSGAuthenticationServiceService {
	return NewWSGAuthenticationServiceService(ss.apiClient)
}

// AddServicesAuthAd
//
// Use this API command to create a new active directory authentication service.
//
// Request Body:
//	 - body *service.CreateActiveDirectoryAuthentication
func (s *WSGAuthenticationServiceService) AddServicesAuthAd(ctx context.Context, body *service.CreateActiveDirectoryAuthentication) (*common.CreateResult, error) {
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

// AddServicesAuthHlr
//
// Use this API command to create a new hlr authentication service.
//
// Request Body:
//	 - body *service.CreateHlrAuthentication
func (s *WSGAuthenticationServiceService) AddServicesAuthHlr(ctx context.Context, body *service.CreateHlrAuthentication) (*common.CreateResult, error) {
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

// AddServicesAuthLdap
//
// Use this API command to create a new LDAP authentication service.
//
// Request Body:
//	 - body *service.CreateLDAPAuthentication
func (s *WSGAuthenticationServiceService) AddServicesAuthLdap(ctx context.Context, body *service.CreateLDAPAuthentication) (*common.CreateResult, error) {
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

// AddServicesAuthRadius
//
// Use this API command to create a new RADIUS authentication service.
//
// Request Body:
//	 - body *service.CreateRadiusAuthentication
func (s *WSGAuthenticationServiceService) AddServicesAuthRadius(ctx context.Context, body *service.CreateRadiusAuthentication) (*common.CreateResult, error) {
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

// AddServicesAuthTestById
//
// Use this API command to test an authentication service.
//
// Request Body:
//	 - body *service.TestingConfig
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) AddServicesAuthTestById(ctx context.Context, body *service.TestingConfig, pId string) error {
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

// DeleteServicesAuth
//
// Use this API command to delete a list of authentication service.
//
// Request Body:
//	 - body *service.DeleteBulkAuthenticationService
func (s *WSGAuthenticationServiceService) DeleteServicesAuth(ctx context.Context, body *service.DeleteBulkAuthenticationService) (*common.EmptyResult, error) {
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

// DeleteServicesAuthAdById
//
// Use this API command to delete an active directory authentication service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthAdById(ctx context.Context, pId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteServicesAuthById
//
// Use this API command to delete an authentication service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthById(ctx context.Context, pId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteServicesAuthHlrById
//
// Use this API command to delete a hlr authentication service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthHlrById(ctx context.Context, pId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteServicesAuthLdapById
//
// Use this API command to delete a LDAP authentication service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthLdapById(ctx context.Context, pId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteServicesAuthRadiusById
//
// Use this API command to delete a RADIUS authentication service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthRadiusById(ctx context.Context, pId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteServicesAuthRadiusSecondaryById
//
// Use this API command to disable secondary RADIUS server of a RADIUS authentication service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthRadiusSecondaryById(ctx context.Context, pId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteServicesAuthRadiusStandbyPrimaryById
//
// Use this API command to disable Standby secondary RADIUS server of a RADIUS authentication service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthRadiusStandbyPrimaryById(ctx context.Context, pId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindServicesAuthAd
//
// Use this API command to retrieve a list of active directory authentication services.
func (s *WSGAuthenticationServiceService) FindServicesAuthAd(ctx context.Context) (*service.ActiveDirectoryServiceList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindServicesAuthAdById
//
// Use this API command to retrieve an active directory authentication service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) FindServicesAuthAdById(ctx context.Context, pId string) (*service.ActiveDirectoryService, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindServicesAuthAdByQueryCriteria
//
// Use this API command to retrieve a list of AD Authentication services by query criteria.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGAuthenticationServiceService) FindServicesAuthAdByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*service.ActiveDirectoryServiceList, error) {
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

// FindServicesAuthByQueryCriteria
//
// Use this API command to retrieve a list of Authentication services by query criteria.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGAuthenticationServiceService) FindServicesAuthByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*service.CommonAuthenticationServiceList, error) {
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

// FindServicesAuthGuestById
//
// Use this API command to retrieve a Guest authentication service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) FindServicesAuthGuestById(ctx context.Context, pId string) (*service.CommonAuthenticationService, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindServicesAuthHlr
//
// Use this API command to retrieve a list of hlr authentication services.
func (s *WSGAuthenticationServiceService) FindServicesAuthHlr(ctx context.Context) (*service.HlrServiceList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindServicesAuthHlrById
//
// Use this API command to retrieve a hlr authentication service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) FindServicesAuthHlrById(ctx context.Context, pId string) (*service.HlrService, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindServicesAuthHlrByQueryCriteria
//
// Use this API command to retrieve a list of hlr Authentication services by query criteria.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGAuthenticationServiceService) FindServicesAuthHlrByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*service.HlrServiceList, error) {
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

// FindServicesAuthLdap
//
// Use this API command to retrieve a list of LDAP authentication services.
func (s *WSGAuthenticationServiceService) FindServicesAuthLdap(ctx context.Context) (*service.LDAPServiceList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindServicesAuthLdapById
//
// Use this API command to retrieve a LDAP authentication service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) FindServicesAuthLdapById(ctx context.Context, pId string) (*service.LDAPService, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindServicesAuthLdapByQueryCriteria
//
// Use this API command to retrieve a list of LDAP Authentication services by query criteria.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGAuthenticationServiceService) FindServicesAuthLdapByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*service.LDAPServiceList, error) {
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

// FindServicesAuthLocal_dbById
//
// Use this API command to retrieve a LocalDB authentication service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) FindServicesAuthLocal_dbById(ctx context.Context, pId string) (*service.CommonAuthenticationService, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindServicesAuthRadius
//
// Use this API command to retrieve a list of RADIUS authentication services.
func (s *WSGAuthenticationServiceService) FindServicesAuthRadius(ctx context.Context) (*service.RadiusAuthenticationServiceList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindServicesAuthRadiusById
//
// Use this API command to retrieve a RADIUS authentication service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) FindServicesAuthRadiusById(ctx context.Context, pId string) (*service.RadiusAuthenticationService, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindServicesAuthRadiusByQueryCriteria
//
// Use this API command to retrieve a list of radius Authentication services by query criteria.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGAuthenticationServiceService) FindServicesAuthRadiusByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*service.RadiusAuthenticationServiceList, error) {
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

// PartialUpdateServicesAuthAdById
//
// Use this API command to modify the basic information of an active directory authentication service.
//
// Request Body:
//	 - body *service.ModifyActiveDirectoryAuthentication
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) PartialUpdateServicesAuthAdById(ctx context.Context, body *service.ModifyActiveDirectoryAuthentication, pId string) (*common.EmptyResult, error) {
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

// PartialUpdateServicesAuthHlrById
//
// Use this API command to modify the basic information of a hlr authentication service.
//
// Request Body:
//	 - body *service.ModifyHlrAuthentication
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) PartialUpdateServicesAuthHlrById(ctx context.Context, body *service.ModifyHlrAuthentication, pId string) (*common.EmptyResult, error) {
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

// PartialUpdateServicesAuthLdapById
//
// Use this API command to modify the basic information of a LDAP authentication service.
//
// Request Body:
//	 - body *service.ModifyLDAPAuthentication
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) PartialUpdateServicesAuthLdapById(ctx context.Context, body *service.ModifyLDAPAuthentication, pId string) (*common.EmptyResult, error) {
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

// PartialUpdateServicesAuthLocal_dbById
//
// Use this API command to update LocalDB authentication service.
//
// Request Body:
//	 - body *service.ModifyLocalDbAuthentication
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) PartialUpdateServicesAuthLocal_dbById(ctx context.Context, body *service.ModifyLocalDbAuthentication, pId string) (*common.EmptyResult, error) {
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

// PartialUpdateServicesAuthRadiusById
//
// Use this API command to modify the basic information of a RADIUS authentication service.
//
// Request Body:
//	 - body *service.ModifyRadiusAuthentication
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) PartialUpdateServicesAuthRadiusById(ctx context.Context, body *service.ModifyRadiusAuthentication, pId string) (*common.EmptyResult, error) {
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
