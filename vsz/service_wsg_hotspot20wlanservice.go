package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/portalservice"
	"gopkg.in/go-playground/validator.v9"
)

type WSGHotspot20WLANServiceService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGHotspot20WLANServiceService(c *APIClient) *WSGHotspot20WLANServiceService {
	s := new(WSGHotspot20WLANServiceService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGHotspot20WLANServiceService() *WSGHotspot20WLANServiceService {
	return NewWSGHotspot20WLANServiceService(ss.apiClient)
}

// AddRkszonesHs20sByZoneId
//
// Use this API command to create a new Hotspot 2.0 WLAN profile of a zone.
//
// Request Body:
//	 - body *portalservice.CreateHotspot20WlanProfile
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGHotspot20WLANServiceService) AddRkszonesHs20sByZoneId(ctx context.Context, body *portalservice.CreateHotspot20WlanProfile, pZoneId string) (*common.CreateResult, error) {
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

// DeleteRkszonesHs20sById
//
// Use this API command to delete a Hotspot 2.0 WLAN Profile of a zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGHotspot20WLANServiceService) DeleteRkszonesHs20sById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesHs20sById
//
// Use this API command to retrieve a Hotspot 2.0 WLAN profile of a zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGHotspot20WLANServiceService) FindRkszonesHs20sById(ctx context.Context, pId string, pZoneId string) (*portalservice.Hotspot20WlanProfile, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesHs20sByZoneId
//
// Use this API command to retrieve a list of Hotspot 2.0 WLAN profiles of a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGHotspot20WLANServiceService) FindRkszonesHs20sByZoneId(ctx context.Context, pZoneId string) (*portalservice.PortalServiceList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateRkszonesHs20sById
//
// Use this API command to modify the basic information on Hotspot 2.0 WLAN profile of a zone.
//
// Request Body:
//	 - body *portalservice.ModifyHotspot20WlanProfile
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGHotspot20WLANServiceService) PartialUpdateRkszonesHs20sById(ctx context.Context, body *portalservice.ModifyHotspot20WlanProfile, pId string, pZoneId string) (*common.EmptyResult, error) {
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
