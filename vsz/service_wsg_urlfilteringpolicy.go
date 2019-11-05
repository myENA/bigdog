package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/urlfiltering"
	"gopkg.in/go-playground/validator.v9"
)

type WSGURLFilteringPolicyService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGURLFilteringPolicyService(c *APIClient) *WSGURLFilteringPolicyService {
	s := new(WSGURLFilteringPolicyService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGURLFilteringPolicyService() *WSGURLFilteringPolicyService {
	return NewWSGURLFilteringPolicyService(ss.apiClient)
}

// AddUrlFilteringUrlFilteringPolicy
//
// Use this API command to create a URL Filtering policy.
//
// Request Body:
//	 - body *urlfiltering.CreateUrlFilteringPolicy
func (s *WSGURLFilteringPolicyService) AddUrlFilteringUrlFilteringPolicy(ctx context.Context, body *urlfiltering.CreateUrlFilteringPolicy) (*common.CreateResult, error) {
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

// DeleteUrlFilteringUrlFilteringPolicy
//
// Use this API command to delete bulk URL Filtering policies.
//
// Request Body:
//	 - body *urlfiltering.DeleteBulk
func (s *WSGURLFilteringPolicyService) DeleteUrlFilteringUrlFilteringPolicy(ctx context.Context, body *urlfiltering.DeleteBulk) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return errors.New("body cannot be empty")
	}
}

// DeleteUrlFilteringUrlFilteringPolicyById
//
// Use this API command to delete a URL Filtering policy.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGURLFilteringPolicyService) DeleteUrlFilteringUrlFilteringPolicyById(ctx context.Context, pId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindUrlFilteringBlockCategories
//
// Use this API command to retrieve the block categories of URL Filtering.
func (s *WSGURLFilteringPolicyService) FindUrlFilteringBlockCategories(ctx context.Context) (*urlfiltering.UrlFilteringBlockCategoriesList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindUrlFilteringByQueryCriteria
//
// Use this API command to retrieve a list of URL Filtering policies by query criteria.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGURLFilteringPolicyService) FindUrlFilteringByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*urlfiltering.UrlFilteringPolicyList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
}

// FindUrlFilteringUrlFilteringPolicy
//
// Use this API command to retrieve list of URL Filtering policies.
//
// Query Parameters:
// - qDomainId string
// - qIndex string
// - qListSize string
func (s *WSGURLFilteringPolicyService) FindUrlFilteringUrlFilteringPolicy(ctx context.Context, qDomainId string, qIndex string, qListSize string) (*urlfiltering.UrlFilteringPolicyList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindUrlFilteringUrlFilteringPolicyById
//
// Use this API command to retrieve an URL Filtering policy.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGURLFilteringPolicyService) FindUrlFilteringUrlFilteringPolicyById(ctx context.Context, pId string) (*urlfiltering.UrlFilteringPolicy, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateUrlFilteringUrlFilteringPolicyById
//
// Use this API command to patch a URL Filtering policy.
//
// Request Body:
//	 - body *urlfiltering.ModifyUrlFilteringPolicy
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGURLFilteringPolicyService) PartialUpdateUrlFilteringUrlFilteringPolicyById(ctx context.Context, body *urlfiltering.ModifyUrlFilteringPolicy, pId string) (*common.EmptyResult, error) {
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
