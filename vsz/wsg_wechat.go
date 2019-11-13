package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGWechatService struct {
	apiClient *APIClient
}

func NewWSGWechatService(c *APIClient) *WSGWechatService {
	s := new(WSGWechatService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGWechatService() *WSGWechatService {
	return NewWSGWechatService(ss.apiClient)
}

// AddRkszonesPortalsWechatByZoneId
//
// Use this API command to create wechat portal.
//
// Request Body:
//	 - body *WSGPortalServiceCreateWechat
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGWechatService) AddRkszonesPortalsWechatByZoneId(ctx context.Context, body *WSGPortalServiceCreateWechat, pZoneId string) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesPortalsWechatById
//
// Use this API command to delete wechat portal.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWechatService) DeleteRkszonesPortalsWechatById(ctx context.Context, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesPortalsWechatById
//
// Use this API command to retrieve wechat portal by ID.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWechatService) FindRkszonesPortalsWechatById(ctx context.Context, pId string, pZoneId string) (*WSGPortalServiceWechatConfiguration, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesPortalsWechatByZoneId
//
// Use this API command to retrieve a list of wechat portal.
//
// Path Parameters:
// - pZoneId string
//		- required
//
// Query Parameters:
// - qIndex string
//		- nullable
// - qListSize string
//		- nullable
func (s *WSGWechatService) FindRkszonesPortalsWechatByZoneId(ctx context.Context, pZoneId string, qIndex string, qListSize string) (*WSGPortalServiceList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateRkszonesPortalsWechatById
//
// Use this API command to modify the basic information of wechat portal.
//
// Request Body:
//	 - body *WSGPortalServiceModifyWechat
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWechatService) PartialUpdateRkszonesPortalsWechatById(ctx context.Context, body *WSGPortalServiceModifyWechat, pId string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}
