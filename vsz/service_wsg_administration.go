package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/administration"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"gopkg.in/go-playground/validator.v9"
)

type WSGAdministrationService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGAdministrationService(c *APIClient) *WSGAdministrationService {
	s := new(WSGAdministrationService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGAdministrationService() *WSGAdministrationService {
	return NewWSGAdministrationService(ss.apiClient)
}

// AddAdminaaa
//
// Use this API command to create a new Admin AAA server
//
// Request Body:
//	 - body *administration.CreateAdminAAAServer
func (s *WSGAdministrationService) AddAdminaaa(ctx context.Context, body *administration.CreateAdminAAAServer) (*common.CreateResult, error) {
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

// AddRestart
//
// Use this API command to restart the controller.
func (s *WSGAdministrationService) AddRestart(ctx context.Context) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// AddShutdown
//
// Use this API command to shut down the controller.
func (s *WSGAdministrationService) AddShutdown(ctx context.Context) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteAdminaaaById
//
// Use this API command to delete an existing Admin AAA server
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAdministrationService) DeleteAdminaaaById(ctx context.Context, pId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindAdminaaa
//
// Use this API command to retrieve the list of Admin AAA server
//
// Query Parameters:
// - qType string
//		- required
func (s *WSGAdministrationService) FindAdminaaa(ctx context.Context, qType string) (*administration.RetrieveAdminAAAServerList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindAdminaaaById
//
// Use this API command to retrieve an existing Admin AAA server
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAdministrationService) FindAdminaaaById(ctx context.Context, pId string) (*administration.RetrieveAdminAAAServer, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindLicenses
//
// Use this API command to get all licenses currently assign in SCG.
func (s *WSGAdministrationService) FindLicenses(ctx context.Context) (*administration.LicensesList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindLicenseServer
//
// Use this API command to get license server configuration.
func (s *WSGAdministrationService) FindLicenseServer(ctx context.Context) (*administration.LicenseServer, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindLicensesSummary
//
// Use this API command to get licenses summary information.
func (s *WSGAdministrationService) FindLicensesSummary(ctx context.Context) (*administration.LicensesSummaryList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindLicensesSyncLogs
//
// Use this API command to get licenses synchronize logs.
func (s *WSGAdministrationService) FindLicensesSyncLogs(ctx context.Context) (*administration.LicensesSyncLogsList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateAdminaaaById
//
// Use this API command to modify an existing Admin AAA server
//
// Request Body:
//	 - body *administration.ModifyAdminAAAServer
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAdministrationService) UpdateAdminaaaById(ctx context.Context, body *administration.ModifyAdminAAAServer, pId string) (*common.EmptyResult, error) {
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

// UpdateLicenseServer
//
// Use this API command to update license server configuration.
//
// Request Body:
//	 - body *administration.ModfiyLicenseServer
func (s *WSGAdministrationService) UpdateLicenseServer(ctx context.Context, body *administration.ModfiyLicenseServer) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return err
	}
}

// UpdateLicensesSync
//
// Use this API command to ask all SCG in cluster to sync licenses from license server.
func (s *WSGAdministrationService) UpdateLicensesSync(ctx context.Context) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}
