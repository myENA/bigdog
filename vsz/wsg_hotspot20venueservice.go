package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGHotspot20VenueServiceService struct {
	apiClient *APIClient
}

func NewWSGHotspot20VenueServiceService(c *APIClient) *WSGHotspot20VenueServiceService {
	s := new(WSGHotspot20VenueServiceService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGHotspot20VenueServiceService() *WSGHotspot20VenueServiceService {
	return NewWSGHotspot20VenueServiceService(ss.apiClient)
}

// AddRkszonesHs20VenuesByZoneId
//
// Use this API command to create a new Hotspot 2.0 venue profile of a zone.
//
// Request Body:
//	 - body *WSGPortalServiceCreateHotspot20VenueProfile
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGHotspot20VenueServiceService) AddRkszonesHs20VenuesByZoneId(ctx context.Context, body *WSGPortalServiceCreateHotspot20VenueProfile, pZoneId string) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesHs20VenuesById
//
// Use this API command to delete Hotspot 2.0 venue profile of a zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGHotspot20VenueServiceService) DeleteRkszonesHs20VenuesById(ctx context.Context, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesHs20VenuesById
//
// Use this API command to retrieve a Hotspot 2.0 venue profile of a zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGHotspot20VenueServiceService) FindRkszonesHs20VenuesById(ctx context.Context, pId string, pZoneId string) (*WSGPortalServiceHotspot20VeuneProfile, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesHs20VenuesByZoneId
//
// Use this API command to retrieve a list of Hotspot 2.0 venue profile of a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGHotspot20VenueServiceService) FindRkszonesHs20VenuesByZoneId(ctx context.Context, pZoneId string) (*WSGPortalServiceList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateRkszonesHs20VenuesById
//
// Use this API command to modify the basic information on Hotspot 2.0 venue profile of a zone.
//
// Request Body:
//	 - body *WSGPortalServiceModifyHotspot20VenueProfile
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGHotspot20VenueServiceService) PartialUpdateRkszonesHs20VenuesById(ctx context.Context, body *WSGPortalServiceModifyHotspot20VenueProfile, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
