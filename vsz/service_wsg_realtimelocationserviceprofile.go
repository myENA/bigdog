package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGRealTimeLocationServiceProfileService struct {
	apiClient *APIClient
}

func NewWSGRealTimeLocationServiceProfileService(c *APIClient) *WSGRealTimeLocationServiceProfileService {
	s := new(WSGRealTimeLocationServiceProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGRealTimeLocationServiceProfileService() *WSGRealTimeLocationServiceProfileService {
	serv := new(WSGRealTimeLocationServiceProfileService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddRkszonesRealTimeLocationServiceByZoneId
//
// Use this API command to create RTLS Profile.
//
// Request Body:
//	 - body *profile.CreateRtlsProfile
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGRealTimeLocationServiceProfileService) AddRkszonesRealTimeLocationServiceByZoneId(ctx context.Context, body *profile.CreateRtlsProfile, pZoneId string) (*common.CreateResult, error) {
}

// DeleteRkszonesRealTimeLocationServiceById
//
// Use this API command to Delete RTLS Profile by profile's ID.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGRealTimeLocationServiceProfileService) DeleteRkszonesRealTimeLocationServiceById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

// FindRkszonesRealTimeLocationServiceById
//
// Use this API command to Get RTLS Profile by profile's ID.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGRealTimeLocationServiceProfileService) FindRkszonesRealTimeLocationServiceById(ctx context.Context, pId string, pZoneId string) (*profile.CreateRtlsProfile, error) {
}

// FindRkszonesRealTimeLocationServiceByZoneId
//
// Use this API command to Get RTLS Profile by zone ID.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGRealTimeLocationServiceProfileService) FindRkszonesRealTimeLocationServiceByZoneId(ctx context.Context, pZoneId string) (*profile.RtlsProfileList, error) {
}

// UpdateRkszonesRealTimeLocationServiceById
//
// Use this API command to Modify RTLS Profile by profile's ID.
//
// Request Body:
//	 - body *profile.UpdateRtlsProfile
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGRealTimeLocationServiceProfileService) UpdateRkszonesRealTimeLocationServiceById(ctx context.Context, body *profile.UpdateRtlsProfile, pId string, pZoneId string) (*common.EmptyResult, error) {
}
