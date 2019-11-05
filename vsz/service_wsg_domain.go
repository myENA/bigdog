package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/domain"
)

type WSGDomainService struct {
	apiClient *APIClient
}

func NewWSGDomainService(c *APIClient) *WSGDomainService {
	s := new(WSGDomainService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDomainService() *WSGDomainService {
	serv := new(WSGDomainService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddDomains
//
// Use this API command to create new domain.
//
// Request Body:
//	 - body *domain.CreateDomain
//
// Query Parameters:
// - qParentDomainId string
func (s *WSGDomainService) AddDomains(ctx context.Context, body *domain.CreateDomain, qParentDomainId string) (*common.CreateResult, error) {
}

// AddRkszonesDomains
//
// Use this API command to create new domain.
//
// Request Body:
//	 - body *domain.CreateDomain
//
// Query Parameters:
// - qParentDomainId string
func (s *WSGDomainService) AddRkszonesDomains(ctx context.Context, body *domain.CreateDomain, qParentDomainId string) (*common.CreateResult, error) {
}

// DeleteDomainsById
//
// Use this API command to delete domain.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDomainService) DeleteDomainsById(ctx context.Context, pId string) error {
}

// DeleteRkszonesDomainsById
//
// Use this API command to delete domain.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDomainService) DeleteRkszonesDomainsById(ctx context.Context, pId string) error {
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

// PartialUpdateDomainsById
//
// Use this API command to modify the basic information of domain.
//
// Request Body:
//	 - body *domain.ModifyDomain
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDomainService) PartialUpdateDomainsById(ctx context.Context, body *domain.ModifyDomain, pId string) error {
}

// PartialUpdateRkszonesDomainsById
//
// Use this API command to modify the basic information of domain.
//
// Request Body:
//	 - body *domain.ModifyDomain
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDomainService) PartialUpdateRkszonesDomainsById(ctx context.Context, body *domain.ModifyDomain, pId string) error {
}
