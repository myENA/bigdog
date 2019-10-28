package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/aaa"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGZoneAAAService struct {
	client *Client
}

func NewWSGZoneAAAService(client *Client) *WSGZoneAAAService {
	s := new(WSGZoneAAAService)
	s.client = client
	return s
}

func (ss *WSGService) WSGZoneAAAService() *WSGZoneAAAService {
	serv := new(WSGZoneAAAService)
	serv.client = ss.client
	return serv
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
