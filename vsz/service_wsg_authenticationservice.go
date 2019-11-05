package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/service"
)

type WSGAuthenticationServiceService struct {
	apiClient *APIClient
}

func NewWSGAuthenticationServiceService(c *APIClient) *WSGAuthenticationServiceService {
	s := new(WSGAuthenticationServiceService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAuthenticationServiceService() *WSGAuthenticationServiceService {
	serv := new(WSGAuthenticationServiceService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddServicesAuthAd
//
// Use this API command to create a new active directory authentication service.
//
// Request Body:
//	 - body *service.CreateActiveDirectoryAuthentication
func (s *WSGAuthenticationServiceService) AddServicesAuthAd(ctx context.Context, body *service.CreateActiveDirectoryAuthentication) (*common.CreateResult, error) {
}

// AddServicesAuthHlr
//
// Use this API command to create a new hlr authentication service.
//
// Request Body:
//	 - body *service.CreateHlrAuthentication
func (s *WSGAuthenticationServiceService) AddServicesAuthHlr(ctx context.Context, body *service.CreateHlrAuthentication) (*common.CreateResult, error) {
}

// AddServicesAuthLdap
//
// Use this API command to create a new LDAP authentication service.
//
// Request Body:
//	 - body *service.CreateLDAPAuthentication
func (s *WSGAuthenticationServiceService) AddServicesAuthLdap(ctx context.Context, body *service.CreateLDAPAuthentication) (*common.CreateResult, error) {
}

// AddServicesAuthRadius
//
// Use this API command to create a new RADIUS authentication service.
//
// Request Body:
//	 - body *service.CreateRadiusAuthentication
func (s *WSGAuthenticationServiceService) AddServicesAuthRadius(ctx context.Context, body *service.CreateRadiusAuthentication) (*common.CreateResult, error) {
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
}

// DeleteServicesAuth
//
// Use this API command to delete a list of authentication service.
//
// Request Body:
//	 - body *service.DeleteBulkAuthenticationService
func (s *WSGAuthenticationServiceService) DeleteServicesAuth(ctx context.Context, body *service.DeleteBulkAuthenticationService) (*common.EmptyResult, error) {
}

// DeleteServicesAuthAdById
//
// Use this API command to delete an active directory authentication service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthAdById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteServicesAuthById
//
// Use this API command to delete an authentication service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteServicesAuthHlrById
//
// Use this API command to delete a hlr authentication service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthHlrById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteServicesAuthLdapById
//
// Use this API command to delete a LDAP authentication service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthLdapById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteServicesAuthRadiusById
//
// Use this API command to delete a RADIUS authentication service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthRadiusById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteServicesAuthRadiusSecondaryById
//
// Use this API command to disable secondary RADIUS server of a RADIUS authentication service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthRadiusSecondaryById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteServicesAuthRadiusStandbyPrimaryById
//
// Use this API command to disable Standby secondary RADIUS server of a RADIUS authentication service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) DeleteServicesAuthRadiusStandbyPrimaryById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// FindServicesAuthAd
//
// Use this API command to retrieve a list of active directory authentication services.
func (s *WSGAuthenticationServiceService) FindServicesAuthAd(ctx context.Context) (*service.ActiveDirectoryServiceList, error) {
}

// FindServicesAuthAdById
//
// Use this API command to retrieve an active directory authentication service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) FindServicesAuthAdById(ctx context.Context, pId string) (*service.ActiveDirectoryService, error) {
}

// FindServicesAuthAdByQueryCriteria
//
// Use this API command to retrieve a list of AD Authentication services by query criteria.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGAuthenticationServiceService) FindServicesAuthAdByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*service.ActiveDirectoryServiceList, error) {
}

// FindServicesAuthByQueryCriteria
//
// Use this API command to retrieve a list of Authentication services by query criteria.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGAuthenticationServiceService) FindServicesAuthByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*service.CommonAuthenticationServiceList, error) {
}

// FindServicesAuthGuestById
//
// Use this API command to retrieve a Guest authentication service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) FindServicesAuthGuestById(ctx context.Context, pId string) (*service.CommonAuthenticationService, error) {
}

// FindServicesAuthHlr
//
// Use this API command to retrieve a list of hlr authentication services.
func (s *WSGAuthenticationServiceService) FindServicesAuthHlr(ctx context.Context) (*service.HlrServiceList, error) {
}

// FindServicesAuthHlrById
//
// Use this API command to retrieve a hlr authentication service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) FindServicesAuthHlrById(ctx context.Context, pId string) (*service.HlrService, error) {
}

// FindServicesAuthHlrByQueryCriteria
//
// Use this API command to retrieve a list of hlr Authentication services by query criteria.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGAuthenticationServiceService) FindServicesAuthHlrByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*service.HlrServiceList, error) {
}

// FindServicesAuthLdap
//
// Use this API command to retrieve a list of LDAP authentication services.
func (s *WSGAuthenticationServiceService) FindServicesAuthLdap(ctx context.Context) (*service.LDAPServiceList, error) {
}

// FindServicesAuthLdapById
//
// Use this API command to retrieve a LDAP authentication service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) FindServicesAuthLdapById(ctx context.Context, pId string) (*service.LDAPService, error) {
}

// FindServicesAuthLdapByQueryCriteria
//
// Use this API command to retrieve a list of LDAP Authentication services by query criteria.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGAuthenticationServiceService) FindServicesAuthLdapByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*service.LDAPServiceList, error) {
}

// FindServicesAuthLocal_dbById
//
// Use this API command to retrieve a LocalDB authentication service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) FindServicesAuthLocal_dbById(ctx context.Context, pId string) (*service.CommonAuthenticationService, error) {
}

// FindServicesAuthRadius
//
// Use this API command to retrieve a list of RADIUS authentication services.
func (s *WSGAuthenticationServiceService) FindServicesAuthRadius(ctx context.Context) (*service.RadiusAuthenticationServiceList, error) {
}

// FindServicesAuthRadiusById
//
// Use this API command to retrieve a RADIUS authentication service.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAuthenticationServiceService) FindServicesAuthRadiusById(ctx context.Context, pId string) (*service.RadiusAuthenticationService, error) {
}

// FindServicesAuthRadiusByQueryCriteria
//
// Use this API command to retrieve a list of radius Authentication services by query criteria.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGAuthenticationServiceService) FindServicesAuthRadiusByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*service.RadiusAuthenticationServiceList, error) {
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
}
