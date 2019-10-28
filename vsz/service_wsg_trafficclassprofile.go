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

// FindRkszonesTrafficClassProfileById
//
// Use this API command to retrieve a Traffic Class Profile of zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGTrafficClassProfileService) FindRkszonesTrafficClassProfileById(ctx context.Context, pId string, pZoneId string) (*common.TrafficClassProfileRef, error) {
}

// FindRkszonesTrafficClassProfileByZoneId
//
// Use this API command to retrieve a list of Traffic Class Profile of a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGTrafficClassProfileService) FindRkszonesTrafficClassProfileByZoneId(ctx context.Context, pZoneId string) (*profile.TrafficClassProfileList, error) {
}

// FindServicesTrafficClassProfileByQueryCriteria
//
// Retrieve a list of Traffic Class Profile.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGTrafficClassProfileService) FindServicesTrafficClassProfileByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.TrafficClassProfileList, error) {
}
