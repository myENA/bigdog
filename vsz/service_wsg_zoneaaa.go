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

func (s *WSGZoneAAAService) DeleteRkszonesAaaById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGZoneAAAService) DeleteRkszonesAaaByZoneId(ctx context.Context, body *aaa.DeleteBulkAAAServerList, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGZoneAAAService) DeleteRkszonesAaaRadiusSecondaryById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGZoneAAAService) DeleteRkszonesAaaRadiusStandbyPrimaryById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

func (s *WSGZoneAAAService) FindRkszonesAaaAdById(ctx context.Context, pId string, pZoneId string) (*aaa.ActiveDirectory, error) {
}

func (s *WSGZoneAAAService) FindRkszonesAaaAdByZoneId(ctx context.Context, pZoneId string) (*aaa.ActiveDirectoryList, error) {
}

func (s *WSGZoneAAAService) FindRkszonesAaaLdapById(ctx context.Context, pId string, pZoneId string) (*aaa.LDAPServer, error) {
}

func (s *WSGZoneAAAService) FindRkszonesAaaLdapByZoneId(ctx context.Context, pZoneId string) (*aaa.LDAPServerList, error) {
}

func (s *WSGZoneAAAService) FindRkszonesAaaRadiusById(ctx context.Context, pId string, pZoneId string) (*aaa.AuthenticationServer, error) {
}

func (s *WSGZoneAAAService) FindRkszonesAaaRadiusByZoneId(ctx context.Context, pZoneId string) (*aaa.AuthenticationServerList, error) {
}
