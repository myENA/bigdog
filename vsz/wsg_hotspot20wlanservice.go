package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGHotspot20WLANServiceService struct {
	apiClient *APIClient
}

func NewWSGHotspot20WLANServiceService(c *APIClient) *WSGHotspot20WLANServiceService {
	s := new(WSGHotspot20WLANServiceService)
	s.apiClient = c
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
//	 - body *WSGPortalServiceCreateHotspot20WlanProfile
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGHotspot20WLANServiceService) AddRkszonesHs20sByZoneId(ctx context.Context, body *WSGPortalServiceCreateHotspot20WlanProfile, pZoneId string) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
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
func (s *WSGHotspot20WLANServiceService) DeleteRkszonesHs20sById(ctx context.Context, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
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
func (s *WSGHotspot20WLANServiceService) FindRkszonesHs20sById(ctx context.Context, pId string, pZoneId string) (*WSGPortalServiceHotspot20WlanProfile, error) {
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
func (s *WSGHotspot20WLANServiceService) FindRkszonesHs20sByZoneId(ctx context.Context, pZoneId string) (*WSGPortalServiceList, error) {
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
//	 - body *WSGPortalServiceModifyHotspot20WlanProfile
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGHotspot20WLANServiceService) PartialUpdateRkszonesHs20sById(ctx context.Context, body *WSGPortalServiceModifyHotspot20WlanProfile, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}
