package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGTrafficClassProfileService struct {
	client *Client
}

func NewWSGTrafficClassProfileService(client *Client) *WSGTrafficClassProfileService {
	s := new(WSGTrafficClassProfileService)
	s.client = client
	return s
}

func (ss *WSGService) WSGTrafficClassProfileService() *WSGTrafficClassProfileService {
	serv := new(WSGTrafficClassProfileService)
	serv.client = ss.client
	return serv
}

func (s *WSGTrafficClassProfileService) FindRkszonesTrafficClassProfileById(ctx context.Context, pId string, pZoneId string) (*common.TrafficClassProfileRef, error) {
}

func (s *WSGTrafficClassProfileService) FindRkszonesTrafficClassProfileByZoneId(ctx context.Context, pZoneId string) (*profile.TrafficClassProfileList, error) {
}

func (s *WSGTrafficClassProfileService) FindServicesTrafficClassProfileByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.TrafficClassProfileList, error) {
}
