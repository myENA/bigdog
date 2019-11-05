package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/systemipsec"
	"gopkg.in/go-playground/validator.v9"
)

type WSGSystemIPsecService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGSystemIPsecService(c *APIClient) *WSGSystemIPsecService {
	s := new(WSGSystemIPsecService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGSystemIPsecService() *WSGSystemIPsecService {
	return NewWSGSystemIPsecService(ss.apiClient)
}

// FindSystemIpsec
//
// Use this API command to retrieve the System IPSec.
func (s *WSGSystemIPsecService) FindSystemIpsec(ctx context.Context) (*systemipsec.GetResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateSystemIpsec
//
// Use this API command to modify the System IPSec.
//
// Request Body:
//	 - body *systemipsec.Update
func (s *WSGSystemIPsecService) UpdateSystemIpsec(ctx context.Context, body *systemipsec.Update) (*common.CreateResult, error) {
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
