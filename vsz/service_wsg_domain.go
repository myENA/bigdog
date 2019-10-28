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

// FindDomains
//
// Use this API command to retrieve a list of domain under Administration Domain.
//
// Query Parameters:
// - qExcludeRegularDomain string
// - qIncludeSelf string
// - qIndex string
// - qListSize string
// - qRecursively string
func (s *WSGDomainService) FindDomains(ctx context.Context, qExcludeRegularDomain string, qIncludeSelf string, qIndex string, qListSize string, qRecursively string) (*domain.DomainList, error) {
}

// FindDomainsById
//
// Use this API command to retrieve domain by specified Domain ID.
//
// Path Parameters:
// - pId string
//		- required
//
// Query Parameters:
// - qRecursively string
func (s *WSGDomainService) FindDomainsById(ctx context.Context, pId string, qRecursively string) (*domain.DomainConfiguration, error) {
}

// FindDomainsSubdomainById
//
// Use this API command to retrieve a list of subdomain by specified Domain ID.
//
// Path Parameters:
// - pId string
//		- required
//
// Query Parameters:
// - qExcludeRegularDomain string
// - qIncludeSelf string
// - qIndex string
// - qListSize string
// - qRecursively string
func (s *WSGDomainService) FindDomainsSubdomainById(ctx context.Context, pId string, qExcludeRegularDomain string, qIncludeSelf string, qIndex string, qListSize string, qRecursively string) (*domain.DomainList, error) {
}

// FindRkszonesDomains
//
// Use this API command to retrieve a list of domain under Administration Domain.
//
// Query Parameters:
// - qExcludeRegularDomain string
// - qIncludeSelf string
// - qIndex string
// - qListSize string
// - qRecursively string
func (s *WSGDomainService) FindRkszonesDomains(ctx context.Context, qExcludeRegularDomain string, qIncludeSelf string, qIndex string, qListSize string, qRecursively string) (*domain.DomainList, error) {
}

// FindRkszonesDomainsById
//
// Use this API command to retrieve domain by specified Domain ID.
//
// Path Parameters:
// - pId string
//		- required
//
// Query Parameters:
// - qRecursively string
func (s *WSGDomainService) FindRkszonesDomainsById(ctx context.Context, pId string, qRecursively string) (*domain.DomainConfiguration, error) {
}

// FindRkszonesDomainsSubdomainById
//
// Use this API command to retrieve a list of subdomain by specified Domain ID.
//
// Path Parameters:
// - pId string
//		- required
//
// Query Parameters:
// - qExcludeRegularDomain string
// - qIncludeSelf string
// - qIndex string
// - qListSize string
// - qRecursively string
func (s *WSGDomainService) FindRkszonesDomainsSubdomainById(ctx context.Context, pId string, qExcludeRegularDomain string, qIncludeSelf string, qIndex string, qListSize string, qRecursively string) (*domain.DomainList, error) {
}
