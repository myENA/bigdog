package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGHotspot20WiFiOperatorProfileService struct {
	apiClient *APIClient
}

func NewWSGHotspot20WiFiOperatorProfileService(c *APIClient) *WSGHotspot20WiFiOperatorProfileService {
	s := new(WSGHotspot20WiFiOperatorProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGHotspot20WiFiOperatorProfileService() *WSGHotspot20WiFiOperatorProfileService {
	return NewWSGHotspot20WiFiOperatorProfileService(ss.apiClient)
}

// AddProfilesHs20Operators
//
// Use this API command to create a new Hotspot 2.0 Wi-Fi operator.
//
// Request Body:
//	 - body *WSGProfileHs20Operator
func (s *WSGHotspot20WiFiOperatorProfileService) AddProfilesHs20Operators(ctx context.Context, body *WSGProfileHs20Operator) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteProfilesHs20Operators
//
// Use this API command to delete multiple Hotspot 2.0 Wi-Fi operators.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGHotspot20WiFiOperatorProfileService) DeleteProfilesHs20Operators(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteProfilesHs20OperatorsById
//
// Use this API command to delete a Hotspot 2.0 Wi-Fi operator.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGHotspot20WiFiOperatorProfileService) DeleteProfilesHs20OperatorsById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteProfilesHs20OperatorsCertificateById
//
// Use this API command to disable certificate of a Hotspot 2.0 Wi-Fi operator.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGHotspot20WiFiOperatorProfileService) DeleteProfilesHs20OperatorsCertificateById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesHs20Operators
//
// Use this API command to retrieve list of Hotspot 2.0 Wi-Fi Operators.
//
// Query Parameters:
// - qIndex string
//		- nullable
// - qListSize string
//		- nullable
func (s *WSGHotspot20WiFiOperatorProfileService) FindProfilesHs20Operators(ctx context.Context, qIndex string, qListSize string) (*WSGProfileHs20OperatorList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesHs20OperatorsById
//
// Use this API command to retrieve a Hotspot 2.0 Wi-Fi operator.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGHotspot20WiFiOperatorProfileService) FindProfilesHs20OperatorsById(ctx context.Context, pId string) (*WSGProfileHs20Operator, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindProfilesHs20OperatorsByQueryCriteria
//
// Query hotspot 2.0 Wi-Fi operators.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGHotspot20WiFiOperatorProfileService) FindProfilesHs20OperatorsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileHs20OperatorList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateProfilesHs20OperatorsById
//
// Use this API command to modify the basic information of a Hotspot 2.0 Wi-Fi operator.
//
// Request Body:
//	 - body *WSGProfileModifyHS20Operator
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGHotspot20WiFiOperatorProfileService) PartialUpdateProfilesHs20OperatorsById(ctx context.Context, body *WSGProfileModifyHS20Operator, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateProfilesHs20OperatorsById
//
// Use this API command to modify entire configuration of a Hotspot 2.0 Wi-Fi operator.
//
// Request Body:
//	 - body *WSGProfileHs20Operator
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGHotspot20WiFiOperatorProfileService) UpdateProfilesHs20OperatorsById(ctx context.Context, body *WSGProfileHs20Operator, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}
