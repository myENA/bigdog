package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/service"
)

type WSGAuthenticationServiceService struct {
	client *Client
}

func NewWSGAuthenticationServiceService(client *Client) *WSGAuthenticationServiceService {
	s := new(WSGAuthenticationServiceService)
	s.client = client
	return s
}

func (ss *WSGService) WSGAuthenticationServiceService() *WSGAuthenticationServiceService {
	serv := new(WSGAuthenticationServiceService)
	serv.client = ss.client
	return serv
}

func (s *WSGAuthenticationServiceService) AddServicesAuthTestById(ctx context.Context, id string) error {
}

func (s *WSGAuthenticationServiceService) DeleteServicesAuth(ctx context.Context) (*common.EmptyResult, error) {
}

func (s *WSGAuthenticationServiceService) DeleteServicesAuthById(ctx context.Context, id string) (*common.EmptyResult, error) {
}

func (s *WSGAuthenticationServiceService) DeleteServicesAuthRadiusSecondaryById(ctx context.Context, id string) (*common.EmptyResult, error) {
}

func (s *WSGAuthenticationServiceService) DeleteServicesAuthRadiusStandbyPrimaryById(ctx context.Context, id string) (*common.EmptyResult, error) {
}

func (s *WSGAuthenticationServiceService) FindServicesAuthAd(ctx context.Context) (*service.ActiveDirectoryServiceList, error) {
}

func (s *WSGAuthenticationServiceService) FindServicesAuthAdById(ctx context.Context, id string) (*service.ActiveDirectoryService, error) {
}

func (s *WSGAuthenticationServiceService) FindServicesAuthAdByQueryCriteria(ctx context.Context) (*service.ActiveDirectoryServiceList, error) {
}

func (s *WSGAuthenticationServiceService) FindServicesAuthByQueryCriteria(ctx context.Context) (*service.CommonAuthenticationServiceList, error) {
}

func (s *WSGAuthenticationServiceService) FindServicesAuthGuestById(ctx context.Context, id string) (*service.CommonAuthenticationService, error) {
}

func (s *WSGAuthenticationServiceService) FindServicesAuthHlr(ctx context.Context) (*service.HlrServiceList, error) {
}

func (s *WSGAuthenticationServiceService) FindServicesAuthHlrById(ctx context.Context, id string) (*service.HlrService, error) {
}

func (s *WSGAuthenticationServiceService) FindServicesAuthHlrByQueryCriteria(ctx context.Context) (*service.HlrServiceList, error) {
}

func (s *WSGAuthenticationServiceService) FindServicesAuthLdap(ctx context.Context) (*service.LDAPServiceList, error) {
}

func (s *WSGAuthenticationServiceService) FindServicesAuthLdapById(ctx context.Context, id string) (*service.LDAPService, error) {
}

func (s *WSGAuthenticationServiceService) FindServicesAuthLdapByQueryCriteria(ctx context.Context) (*service.LDAPServiceList, error) {
}

func (s *WSGAuthenticationServiceService) FindServicesAuthLocal_dbById(ctx context.Context, id string) (*service.CommonAuthenticationService, error) {
}

func (s *WSGAuthenticationServiceService) FindServicesAuthRadius(ctx context.Context) (*service.RadiusAuthenticationServiceList, error) {
}

func (s *WSGAuthenticationServiceService) FindServicesAuthRadiusById(ctx context.Context, id string) (*service.RadiusAuthenticationService, error) {
}

func (s *WSGAuthenticationServiceService) FindServicesAuthRadiusByQueryCriteria(ctx context.Context) (*service.RadiusAuthenticationServiceList, error) {
}
