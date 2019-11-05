package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/aaa"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGZoneAAAService struct {
	apiClient *APIClient
}

func NewWSGZoneAAAService(c *APIClient) *WSGZoneAAAService {
	s := new(WSGZoneAAAService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGZoneAAAService() *WSGZoneAAAService {
	serv := new(WSGZoneAAAService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddRkszonesAaaAdByZoneId
//
// Use this API command to create a new active directory server of a zone.
//
// Request Body:
//	 - body *aaa.CreateActiveDirectoryServer
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) AddRkszonesAaaAdByZoneId(ctx context.Context, body *aaa.CreateActiveDirectoryServer, pZoneId string) (*common.CreateResult, error) {
}

// AddRkszonesAaaLdapByZoneId
//
// Use this API command to create a new LDAP server of a zone.
//
// Request Body:
//	 - body *aaa.CreateLDAPServer
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) AddRkszonesAaaLdapByZoneId(ctx context.Context, body *aaa.CreateLDAPServer, pZoneId string) (*common.CreateResult, error) {
}

// AddRkszonesAaaRadiusByZoneId
//
// Use this API command to create a new radius server of a zone.
//
// Request Body:
//	 - body *aaa.CreateAuthenticationServer
//
// Path Parameters:
// - pZoneId string
//		- required
//
// Query Parameters:
// - qForAccounting string
func (s *WSGZoneAAAService) AddRkszonesAaaRadiusByZoneId(ctx context.Context, body *aaa.CreateAuthenticationServer, pZoneId string, qForAccounting string) (*common.CreateResult, error) {
}

// DeleteRkszonesAaaAdById
//
// Use this API command to delete an active directory server of a zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) DeleteRkszonesAaaAdById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesAaaById
//
// Use this API command to delete an AAA server.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) DeleteRkszonesAaaById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesAaaByZoneId
//
// Use this API command to delete a list of AAA server.
//
// Request Body:
//	 - body *aaa.DeleteBulkAAAServerList
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) DeleteRkszonesAaaByZoneId(ctx context.Context, body *aaa.DeleteBulkAAAServerList, pZoneId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesAaaLdapById
//
// Use this API command to delete a LDAP server of a zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) DeleteRkszonesAaaLdapById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesAaaRadiusById
//
// Use this API command to delete a radius server of a zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) DeleteRkszonesAaaRadiusById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesAaaRadiusSecondaryById
//
// Use this API command to disable secondary server on radius server of a zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) DeleteRkszonesAaaRadiusSecondaryById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesAaaRadiusStandbyPrimaryById
//
// Use this API command to disable primary server on radius server of a zone for Standby Cluster.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) DeleteRkszonesAaaRadiusStandbyPrimaryById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

// FindRkszonesAaaAdById
//
// Use this API command to retrieve an active directory server of a zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) FindRkszonesAaaAdById(ctx context.Context, pId string, pZoneId string) (*aaa.ActiveDirectory, error) {
}

// FindRkszonesAaaAdByZoneId
//
// Use this API command to retrieve a list of active directory servers of a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) FindRkszonesAaaAdByZoneId(ctx context.Context, pZoneId string) (*aaa.ActiveDirectoryList, error) {
}

// FindRkszonesAaaLdapById
//
// Use this API command to retrieve a LDAP server of a zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) FindRkszonesAaaLdapById(ctx context.Context, pId string, pZoneId string) (*aaa.LDAPServer, error) {
}

// FindRkszonesAaaLdapByZoneId
//
// Use this API command to retrieve a list of LDAP servers of a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) FindRkszonesAaaLdapByZoneId(ctx context.Context, pZoneId string) (*aaa.LDAPServerList, error) {
}

// FindRkszonesAaaRadiusById
//
// Use this API command to retrieve a radius server of a zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) FindRkszonesAaaRadiusById(ctx context.Context, pId string, pZoneId string) (*aaa.AuthenticationServer, error) {
}

// FindRkszonesAaaRadiusByZoneId
//
// Use this API command to retrieve a list of radius servers of a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) FindRkszonesAaaRadiusByZoneId(ctx context.Context, pZoneId string) (*aaa.AuthenticationServerList, error) {
}

// PartialUpdateRkszonesAaaAdById
//
// Use this API command to modify the basic information on active directory server of a zone.
//
// Request Body:
//	 - body *aaa.ModifyActiveDirectoryServer
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) PartialUpdateRkszonesAaaAdById(ctx context.Context, body *aaa.ModifyActiveDirectoryServer, pId string, pZoneId string) (*common.EmptyResult, error) {
}

// PartialUpdateRkszonesAaaLdapById
//
// Use this API command to modify the basic information on LDAP server of a zone.
//
// Request Body:
//	 - body *aaa.ModifyLDAPServer
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) PartialUpdateRkszonesAaaLdapById(ctx context.Context, body *aaa.ModifyLDAPServer, pId string, pZoneId string) (*common.EmptyResult, error) {
}

// PartialUpdateRkszonesAaaRadiusById
//
// Use this API command to modify the basic information on radius server of a zone.
//
// Request Body:
//	 - body *aaa.ModifyAuthenticationServer
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) PartialUpdateRkszonesAaaRadiusById(ctx context.Context, body *aaa.ModifyAuthenticationServer, pId string, pZoneId string) (*common.EmptyResult, error) {
}

// UpdateRkszonesAaaAdById
//
// Use this API command to modify the basic information on active directory server of a zone by complete attributes.
//
// Request Body:
//	 - body *aaa.ModifyActiveDirectoryServer
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) UpdateRkszonesAaaAdById(ctx context.Context, body *aaa.ModifyActiveDirectoryServer, pId string, pZoneId string) (*common.EmptyResult, error) {
}

// UpdateRkszonesAaaLdapById
//
// Use this API command to modify the basic information on LDAP server of a zone by complete attributes.
//
// Request Body:
//	 - body *aaa.ModifyLDAPServer
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) UpdateRkszonesAaaLdapById(ctx context.Context, body *aaa.ModifyLDAPServer, pId string, pZoneId string) (*common.EmptyResult, error) {
}

// UpdateRkszonesAaaRadiusById
//
// Use this API command to modify the basic information on radius server of a zone by complete attributes.
//
// Request Body:
//	 - body *aaa.ModifyAuthenticationServer
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) UpdateRkszonesAaaRadiusById(ctx context.Context, body *aaa.ModifyAuthenticationServer, pId string, pZoneId string) (*common.EmptyResult, error) {
}
