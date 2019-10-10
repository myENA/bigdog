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

func (s *WSGIndoorMapService) FindMaps(ctx context.Context) (*indoormap.IndoorMapList, error) {
}

func (s *WSGIndoorMapService) FindMapsByIndoorMapId(ctx context.Context, indoorMapId string) (*indoormap.IndoorMap, error) {
}

func (s *WSGIndoorMapService) FindMapsByQueryCriteria(ctx context.Context) (*indoormap.IndoorMapList, error) {
}

func (s *WSGIndoorMapService) UpdateMapsApsByIndoorMapId(ctx context.Context, indoorMapId string) (*indoormap.IndooMapAuditId, error) {
}
