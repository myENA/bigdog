package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGHotspotServiceService struct {
	apiClient *APIClient
}

func NewWSGHotspotServiceService(c *APIClient) *WSGHotspotServiceService {
	s := new(WSGHotspotServiceService)
	s.apiClient = c
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
//	 - body *WSGPortalServiceCreateHotspotExternal
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGHotspotServiceService) AddRkszonesPortalsHotspotExternalByZoneId(ctx context.Context, body *WSGPortalServiceCreateHotspotExternal, pZoneId string) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddRkszonesPortalsHotspotInternalByZoneId
//
// Use this API command to create a new Hotspot (WISPr) with internal logon URL of a zone.MacAddressFormat.
//
// Request Body:
//	 - body *WSGPortalServiceCreateHotspotInternal
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGHotspotServiceService) AddRkszonesPortalsHotspotInternalByZoneId(ctx context.Context, body *WSGPortalServiceCreateHotspotInternal, pZoneId string) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddRkszonesPortalsHotspotSmartClientOnlyByZoneId
//
// Use this API command to create a new Hotspot (WISPr) with smart client only of a zone.MacAddressFormat.
//
// Request Body:
//	 - body *WSGPortalServiceCreateHotspotSmartClientOnly
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGHotspotServiceService) AddRkszonesPortalsHotspotSmartClientOnlyByZoneId(ctx context.Context, body *WSGPortalServiceCreateHotspotSmartClientOnly, pZoneId string) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
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
func (s *WSGHotspotServiceService) DeleteRkszonesPortalsHotspotById(ctx context.Context, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
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
func (s *WSGHotspotServiceService) FindRkszonesPortalsHotspotById(ctx context.Context, pId string, pZoneId string) (*WSGPortalServiceHotspot, error) {
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
func (s *WSGHotspotServiceService) FindRkszonesPortalsHotspotByZoneId(ctx context.Context, pZoneId string) (*WSGPortalServiceList, error) {
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
//	 - body *WSGPortalServiceModifyHotspot
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGHotspotServiceService) PartialUpdateRkszonesPortalsHotspotById(ctx context.Context, body *WSGPortalServiceModifyHotspot, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}
