package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/domain"
)

type WSGDomainService struct {
    client *Client
}

func NewWSGDomainService (client *Client) *WSGDomainService {
    s := new(WSGDomainService)
    s.client = client
    return s
}

func (ss *WSGService) WSGDomainService () *WSGDomainService {
    serv := new(WSGDomainService)
    serv.client = ss.client
    return serv
}

func (s *WSGDomainService) FindDomains (ctx context.Context) (*domain.DomainList, error) {
}

func (s *WSGDomainService) FindDomainsById (ctx context.Context, id string) (*domain.DomainConfiguration, error) {
}

func (s *WSGDomainService) FindDomainsSubdomainById (ctx context.Context, id string) (*domain.DomainList, error) {
}

func (s *WSGDomainService) FindRkszonesDomains (ctx context.Context) (*domain.DomainList, error) {
}

func (s *WSGDomainService) FindRkszonesDomainsById (ctx context.Context, id string) (*domain.DomainConfiguration, error) {
}

func (s *WSGDomainService) FindRkszonesDomainsSubdomainById (ctx context.Context, id string) (*domain.DomainList, error) {
}

