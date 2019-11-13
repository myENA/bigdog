package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
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
	return NewWSGZoneAAAService(ss.apiClient)
}

// AddRkszonesAaaAdByZoneId
//
// Use this API command to create a new active directory server of a zone.
//
// Request Body:
//	 - body *WSGAAACreateActiveDirectoryServer
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) AddRkszonesAaaAdByZoneId(ctx context.Context, body *WSGAAACreateActiveDirectoryServer, pZoneId string) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddRkszonesAaaLdapByZoneId
//
// Use this API command to create a new LDAP server of a zone.
//
// Request Body:
//	 - body *WSGAAACreateLDAPServer
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) AddRkszonesAaaLdapByZoneId(ctx context.Context, body *WSGAAACreateLDAPServer, pZoneId string) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddRkszonesAaaRadiusByZoneId
//
// Use this API command to create a new radius server of a zone.
//
// Request Body:
//	 - body *WSGAAACreateAuthenticationServer
//
// Path Parameters:
// - pZoneId string
//		- required
//
// Query Parameters:
// - qForAccounting string
//		- nullable
func (s *WSGZoneAAAService) AddRkszonesAaaRadiusByZoneId(ctx context.Context, body *WSGAAACreateAuthenticationServer, pZoneId string, qForAccounting string) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
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
func (s *WSGZoneAAAService) DeleteRkszonesAaaAdById(ctx context.Context, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
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
func (s *WSGZoneAAAService) DeleteRkszonesAaaById(ctx context.Context, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesAaaByZoneId
//
// Use this API command to delete a list of AAA server.
//
// Request Body:
//	 - body *WSGAAADeleteBulkAAAServerList
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) DeleteRkszonesAaaByZoneId(ctx context.Context, body *WSGAAADeleteBulkAAAServerList, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
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
func (s *WSGZoneAAAService) DeleteRkszonesAaaLdapById(ctx context.Context, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
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
func (s *WSGZoneAAAService) DeleteRkszonesAaaRadiusById(ctx context.Context, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
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
func (s *WSGZoneAAAService) DeleteRkszonesAaaRadiusSecondaryById(ctx context.Context, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
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
func (s *WSGZoneAAAService) DeleteRkszonesAaaRadiusStandbyPrimaryById(ctx context.Context, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
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
func (s *WSGZoneAAAService) FindRkszonesAaaAdById(ctx context.Context, pId string, pZoneId string) (*WSGAAAActiveDirectory, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesAaaAdByZoneId
//
// Use this API command to retrieve a list of active directory servers of a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) FindRkszonesAaaAdByZoneId(ctx context.Context, pZoneId string) (*WSGAAAActiveDirectoryList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
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
func (s *WSGZoneAAAService) FindRkszonesAaaLdapById(ctx context.Context, pId string, pZoneId string) (*WSGAAALDAPServer, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesAaaLdapByZoneId
//
// Use this API command to retrieve a list of LDAP servers of a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) FindRkszonesAaaLdapByZoneId(ctx context.Context, pZoneId string) (*WSGAAALDAPServerList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
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
func (s *WSGZoneAAAService) FindRkszonesAaaRadiusById(ctx context.Context, pId string, pZoneId string) (*WSGAAAAuthenticationServer, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesAaaRadiusByZoneId
//
// Use this API command to retrieve a list of radius servers of a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) FindRkszonesAaaRadiusByZoneId(ctx context.Context, pZoneId string) (*WSGAAAAuthenticationServerList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateRkszonesAaaAdById
//
// Use this API command to modify the basic information on active directory server of a zone.
//
// Request Body:
//	 - body *WSGAAAModifyActiveDirectoryServer
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) PartialUpdateRkszonesAaaAdById(ctx context.Context, body *WSGAAAModifyActiveDirectoryServer, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateRkszonesAaaLdapById
//
// Use this API command to modify the basic information on LDAP server of a zone.
//
// Request Body:
//	 - body *WSGAAAModifyLDAPServer
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) PartialUpdateRkszonesAaaLdapById(ctx context.Context, body *WSGAAAModifyLDAPServer, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateRkszonesAaaRadiusById
//
// Use this API command to modify the basic information on radius server of a zone.
//
// Request Body:
//	 - body *WSGAAAModifyAuthenticationServer
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) PartialUpdateRkszonesAaaRadiusById(ctx context.Context, body *WSGAAAModifyAuthenticationServer, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateRkszonesAaaAdById
//
// Use this API command to modify the basic information on active directory server of a zone by complete attributes.
//
// Request Body:
//	 - body *WSGAAAModifyActiveDirectoryServer
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) UpdateRkszonesAaaAdById(ctx context.Context, body *WSGAAAModifyActiveDirectoryServer, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateRkszonesAaaLdapById
//
// Use this API command to modify the basic information on LDAP server of a zone by complete attributes.
//
// Request Body:
//	 - body *WSGAAAModifyLDAPServer
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) UpdateRkszonesAaaLdapById(ctx context.Context, body *WSGAAAModifyLDAPServer, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateRkszonesAaaRadiusById
//
// Use this API command to modify the basic information on radius server of a zone by complete attributes.
//
// Request Body:
//	 - body *WSGAAAModifyAuthenticationServer
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGZoneAAAService) UpdateRkszonesAaaRadiusById(ctx context.Context, body *WSGAAAModifyAuthenticationServer, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}
