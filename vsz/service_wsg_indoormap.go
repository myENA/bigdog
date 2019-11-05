package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/indoormap"
	"gopkg.in/go-playground/validator.v9"
)

type WSGIndoorMapService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGIndoorMapService(c *APIClient) *WSGIndoorMapService {
	s := new(WSGIndoorMapService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGIndoorMapService() *WSGIndoorMapService {
	return NewWSGIndoorMapService(ss.apiClient)
}

// AddMaps
//
// Use this API command to create indoorMap.
//
// Request Body:
//	 - body *indoormap.IndoorMap
func (s *WSGIndoorMapService) AddMaps(ctx context.Context, body *indoormap.IndoorMap) (*indoormap.IndooMapAuditId, error) {
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

// DeleteMapsByIndoorMapId
//
// Use this API command to delete indoor map.
//
// Path Parameters:
// - pIndoorMapId string
//		- required
func (s *WSGIndoorMapService) DeleteMapsByIndoorMapId(ctx context.Context, pIndoorMapId string) (*indoormap.IndooMapAuditId, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindMaps
//
// Use this API command to get indoor map list.
//
// Query Parameters:
// - qGroupId string
//		- required
// - qGroupType string
//		- required
func (s *WSGIndoorMapService) FindMaps(ctx context.Context, qGroupId string, qGroupType string) (*indoormap.IndoorMapList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindMapsByIndoorMapId
//
// Use this API command to get indoor maps.
//
// Path Parameters:
// - pIndoorMapId string
//		- required
func (s *WSGIndoorMapService) FindMapsByIndoorMapId(ctx context.Context, pIndoorMapId string) (*indoormap.IndoorMap, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindMapsByQueryCriteria
//
// Use this API command to query indoorMap.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGIndoorMapService) FindMapsByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*indoormap.IndoorMapList, error) {
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

// PartialUpdateMapsByIndoorMapId
//
// Use this API command to update specific indoor map.
//
// Request Body:
//	 - body *indoormap.IndoorMap
//
// Path Parameters:
// - pIndoorMapId string
//		- required
func (s *WSGIndoorMapService) PartialUpdateMapsByIndoorMapId(ctx context.Context, body *indoormap.IndoorMap, pIndoorMapId string) (*indoormap.IndooMapAuditId, error) {
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

// UpdateMapsApsByIndoorMapId
//
// Use this API command to put Aps in indoor map.
//
// Request Body:
//	 - body indoormap.AccessPointList
//
// Path Parameters:
// - pIndoorMapId string
//		- required
func (s *WSGIndoorMapService) UpdateMapsApsByIndoorMapId(ctx context.Context, body indoormap.AccessPointList, pIndoorMapId string) (*indoormap.IndooMapAuditId, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if len(body) == 0 {
		return nil, errors.New("body cannot be empty")
	}
}
