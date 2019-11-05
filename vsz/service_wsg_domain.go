package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/domain"
	"gopkg.in/go-playground/validator.v9"
)

type WSGDomainService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGDomainService(c *APIClient) *WSGDomainService {
	s := new(WSGDomainService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGDomainService() *WSGDomainService {
	return NewWSGDomainService(ss.apiClient)
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
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
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
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}

// DeleteDomainsById
//
// Use this API command to delete domain.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDomainService) DeleteDomainsById(ctx context.Context, pId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesDomainsById
//
// Use this API command to delete domain.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGDomainService) DeleteRkszonesDomainsById(ctx context.Context, pId string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
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
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
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
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
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
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
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
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
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
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
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
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
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
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return err
	}
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
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return err
	}
}
