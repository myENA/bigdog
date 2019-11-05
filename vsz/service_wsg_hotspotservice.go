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

type WSGHotspotServiceService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGHotspotServiceService(c *APIClient) *WSGHotspotServiceService {
	s := new(WSGHotspotServiceService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGHotspotServiceService() *WSGHotspotServiceService {
	return NewWSGHotspotServiceService(ss.apiClient)
}

// AddRkszonesPortalsHotspotExternalByZoneId
//
// Use this API command to create a new Hotspot (WISPr) with external logon URL of a zone.MacAddressFormat.
//
// Request Body:
//	 - body *portalservice.CreateHotspotExternal
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGHotspotServiceService) AddRkszonesPortalsHotspotExternalByZoneId(ctx context.Context, body *portalservice.CreateHotspotExternal, pZoneId string) (*common.CreateResult, error) {
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

// AddRkszonesPortalsHotspotInternalByZoneId
//
// Use this API command to create a new Hotspot (WISPr) with internal logon URL of a zone.MacAddressFormat.
//
// Request Body:
//	 - body *portalservice.CreateHotspotInternal
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGHotspotServiceService) AddRkszonesPortalsHotspotInternalByZoneId(ctx context.Context, body *portalservice.CreateHotspotInternal, pZoneId string) (*common.CreateResult, error) {
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

// AddRkszonesPortalsHotspotSmartClientOnlyByZoneId
//
// Use this API command to create a new Hotspot (WISPr) with smart client only of a zone.MacAddressFormat.
//
// Request Body:
//	 - body *portalservice.CreateHotspotSmartClientOnly
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGHotspotServiceService) AddRkszonesPortalsHotspotSmartClientOnlyByZoneId(ctx context.Context, body *portalservice.CreateHotspotSmartClientOnly, pZoneId string) (*common.CreateResult, error) {
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

// DeleteRkszonesPortalsHotspotById
//
// Use this API command to delete a Hotspot (WISPr) of a zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGHotspotServiceService) DeleteRkszonesPortalsHotspotById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesPortalsHotspotById
//
// Use this API command to retrieve a Hotspot (WISPr) of zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGHotspotServiceService) FindRkszonesPortalsHotspotById(ctx context.Context, pId string, pZoneId string) (*portalservice.Hotspot, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesPortalsHotspotByZoneId
//
// Use this API command to retrieve a list of Hotspot (WISPr) of a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGHotspotServiceService) FindRkszonesPortalsHotspotByZoneId(ctx context.Context, pZoneId string) (*portalservice.PortalServiceList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateRkszonesPortalsHotspotById
//
// Use this API command to modify the basic information on Hotspot (WISPr) of a zone.MacAddressFormat.
//
// Request Body:
//	 - body *portalservice.ModifyHotspot
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGHotspotServiceService) PartialUpdateRkszonesPortalsHotspotById(ctx context.Context, body *portalservice.ModifyHotspot, pId string, pZoneId string) (*common.EmptyResult, error) {
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
