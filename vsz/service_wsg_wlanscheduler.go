package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/wlanscheduler"
	"gopkg.in/go-playground/validator.v9"
)

type WSGWLANSchedulerService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGWLANSchedulerService(c *APIClient) *WSGWLANSchedulerService {
	s := new(WSGWLANSchedulerService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGWLANSchedulerService() *WSGWLANSchedulerService {
	return NewWSGWLANSchedulerService(ss.apiClient)
}

// AddRkszonesWlanSchedulersByZoneId
//
// Use this API command to create a new WLAN schedule.
//
// Request Body:
//	 - body *wlanscheduler.CreateWlanScheduler
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGWLANSchedulerService) AddRkszonesWlanSchedulersByZoneId(ctx context.Context, body *wlanscheduler.CreateWlanScheduler, pZoneId string) (*common.CreateResult, error) {
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

// DeleteRkszonesWlanSchedulersById
//
// Use this API command to delete a WLAN schedule.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWLANSchedulerService) DeleteRkszonesWlanSchedulersById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesWlanSchedulersById
//
// Use this API command to retrieve a WLAN schedule.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWLANSchedulerService) FindRkszonesWlanSchedulersById(ctx context.Context, pId string, pZoneId string) (*wlanscheduler.WlanSchedule, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesWlanSchedulersByZoneId
//
// Use this API command to retrieve the list of WLAN schedule from a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGWLANSchedulerService) FindRkszonesWlanSchedulersByZoneId(ctx context.Context, pZoneId string, qIndex string, qListSize string) (*wlanscheduler.WlanScheduleList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateRkszonesWlanSchedulersById
//
// Use this API command to modify the basic information of a WLAN schedule.
//
// Request Body:
//	 - body *wlanscheduler.ModifyWlanScheduler
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWLANSchedulerService) PartialUpdateRkszonesWlanSchedulersById(ctx context.Context, body *wlanscheduler.ModifyWlanScheduler, pId string, pZoneId string) (*common.EmptyResult, error) {
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
