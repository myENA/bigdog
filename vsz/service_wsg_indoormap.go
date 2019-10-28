package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/indoormap"
)

type WSGIndoorMapService struct {
	client *Client
}

func NewWSGIndoorMapService(client *Client) *WSGIndoorMapService {
	s := new(WSGIndoorMapService)
	s.client = client
	return s
}

func (ss *WSGService) WSGIndoorMapService() *WSGIndoorMapService {
	serv := new(WSGIndoorMapService)
	serv.client = ss.client
	return serv
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
