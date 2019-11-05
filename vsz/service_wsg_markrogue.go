package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/ap"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"gopkg.in/go-playground/validator.v9"
)

type WSGMarkRogueService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGMarkRogueService(c *APIClient) *WSGMarkRogueService {
	s := new(WSGMarkRogueService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGMarkRogueService() *WSGMarkRogueService {
	return NewWSGMarkRogueService(ss.apiClient)
}

// AddRogueMarkIgnore
//
// Mark a rogue AP as ignore.
//
// Request Body:
//	 - body *ap.ModifyRogueType
func (s *WSGMarkRogueService) AddRogueMarkIgnore(ctx context.Context, body *ap.ModifyRogueType) (*common.EmptyResult, error) {
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

// AddRogueMarkKnown
//
// Mark a rogue AP as know.
//
// Request Body:
//	 - body *ap.ModifyRogueType
func (s *WSGMarkRogueService) AddRogueMarkKnown(ctx context.Context, body *ap.ModifyRogueType) (*common.EmptyResult, error) {
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

// AddRogueMarkMalicious
//
// Mark a rogue AP as malicious.
//
// Request Body:
//	 - body *ap.ModifyRogueType
func (s *WSGMarkRogueService) AddRogueMarkMalicious(ctx context.Context, body *ap.ModifyRogueType) (*common.EmptyResult, error) {
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

// AddRogueMarkRogue
//
// Mark a rogue AP as rogue.
//
// Request Body:
//	 - body *ap.ModifyRogueType
func (s *WSGMarkRogueService) AddRogueMarkRogue(ctx context.Context, body *ap.ModifyRogueType) (*common.EmptyResult, error) {
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

// AddRogueUnMark
//
// Use this API command to remove the manual admin classification marking.
//
// Request Body:
//	 - body *ap.ModifyRogueType
func (s *WSGMarkRogueService) AddRogueUnMark(ctx context.Context, body *ap.ModifyRogueType) (*common.EmptyResult, error) {
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

// FindRogueMarkKnown
//
// Get Known Rogue AP list.
func (s *WSGMarkRogueService) FindRogueMarkKnown(ctx context.Context) (*ap.ModifyRogueType, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
