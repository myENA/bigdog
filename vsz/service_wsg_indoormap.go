package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/indoormap"
)

type WSGIndoorMapService struct {
	apiClient *APIClient
}

func NewWSGIndoorMapService(c *APIClient) *WSGIndoorMapService {
	s := new(WSGIndoorMapService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGIndoorMapService() *WSGIndoorMapService {
	serv := new(WSGIndoorMapService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddMaps
//
// Use this API command to create indoorMap.
//
// Request Body:
//	 - body *indoormap.IndoorMap
func (s *WSGIndoorMapService) AddMaps(ctx context.Context, body *indoormap.IndoorMap) (*indoormap.IndooMapAuditId, error) {
}

// DeleteMapsByIndoorMapId
//
// Use this API command to delete indoor map.
//
// Path Parameters:
// - pIndoorMapId string
//		- required
func (s *WSGIndoorMapService) DeleteMapsByIndoorMapId(ctx context.Context, pIndoorMapId string) (*indoormap.IndooMapAuditId, error) {
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
}

// FindMapsByIndoorMapId
//
// Use this API command to get indoor maps.
//
// Path Parameters:
// - pIndoorMapId string
//		- required
func (s *WSGIndoorMapService) FindMapsByIndoorMapId(ctx context.Context, pIndoorMapId string) (*indoormap.IndoorMap, error) {
}

// FindMapsByQueryCriteria
//
// Use this API command to query indoorMap.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGIndoorMapService) FindMapsByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*indoormap.IndoorMapList, error) {
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
}
