package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/domain"
)

type WSGDomainService struct {
	client *Client
}

func NewWSGDomainService(client *Client) *WSGDomainService {
	s := new(WSGDomainService)
	s.client = client
	return s
}

func (ss *WSGService) WSGDomainService() *WSGDomainService {
	serv := new(WSGDomainService)
	serv.client = ss.client
	return serv
}

func (s *WSGDomainService) FindDomains(ctx context.Context, qExcludeRegularDomain string, qIncludeSelf string, qIndex string, qListSize string, qRecursively string) (*domain.DomainList, error) {
}

func (s *WSGDomainService) FindDomainsById(ctx context.Context, pId string, qRecursively string) (*domain.DomainConfiguration, error) {
}

func (s *WSGDomainService) FindDomainsSubdomainById(ctx context.Context, pId string, qExcludeRegularDomain string, qIncludeSelf string, qIndex string, qListSize string, qRecursively string) (*domain.DomainList, error) {
}

func (s *WSGDomainService) FindRkszonesDomains(ctx context.Context, qExcludeRegularDomain string, qIncludeSelf string, qIndex string, qListSize string, qRecursively string) (*domain.DomainList, error) {
}

func (s *WSGDomainService) FindRkszonesDomainsById(ctx context.Context, pId string, qRecursively string) (*domain.DomainConfiguration, error) {
}

func (s *WSGDomainService) FindRkszonesDomainsSubdomainById(ctx context.Context, pId string, qExcludeRegularDomain string, qIncludeSelf string, qIndex string, qListSize string, qRecursively string) (*domain.DomainList, error) {
}
