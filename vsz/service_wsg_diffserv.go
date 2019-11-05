package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/zone"
)

type WSGDiffServService struct {
	apiClient *APIClient
}

func NewWSGDiffServService(c *APIClient) *WSGDiffServService {
	s := new(WSGDiffServService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDiffServService() *WSGDiffServService {
	serv := new(WSGDiffServService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddRkszonesDiffservByZoneId
//
// Use this API command to create DiffServ profile.
//
// Request Body:
//	 - body *zone.CreateDiffServProfile
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGDiffServService) AddRkszonesDiffservByZoneId(ctx context.Context, body *zone.CreateDiffServProfile, pZoneId string) (*common.CreateResult, error) {
}

// DeleteRkszonesDiffservById
//
// Use this API command to delete DiffServ profile.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGDiffServService) DeleteRkszonesDiffservById(ctx context.Context, pId string, pZoneId string) error {
}

// FindRkszonesDiffservById
//
// Use this API command to retrieve DiffServ profile.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGDiffServService) FindRkszonesDiffservById(ctx context.Context, pId string, pZoneId string) (*zone.DiffServConfiguration, error) {
}

// FindRkszonesDiffservByZoneId
//
// Use this API command to retrieve a list of DiffServ profile.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGDiffServService) FindRkszonesDiffservByZoneId(ctx context.Context, pZoneId string) (*zone.DiffServList, error) {
}

// PartialUpdateRkszonesDiffservById
//
// Use this API command to modify the basic information of DiffServ profile.
//
// Request Body:
//	 - body *zone.ModifyDiffServProfile
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGDiffServService) PartialUpdateRkszonesDiffservById(ctx context.Context, body *zone.ModifyDiffServProfile, pId string, pZoneId string) error {
}
