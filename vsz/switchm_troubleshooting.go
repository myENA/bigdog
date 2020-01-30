package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type SwitchMTroubleShootingService struct {
	apiClient *APIClient
}

func NewSwitchMTroubleShootingService(c *APIClient) *SwitchMTroubleShootingService {
	s := new(SwitchMTroubleShootingService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMTroubleShootingService() *SwitchMTroubleShootingService {
	return NewSwitchMTroubleShootingService(ss.apiClient)
}

type SwitchMTroubleShootingSupportLogStatus struct {
	// CreatedTime
	// Created Time of this SupportLog Request
	CreatedTime *string `json:"createdTime,omitempty"`

	// DownloadStatus
	// SupportLog Download Status (DOWNLOADING, DONE, FAILED)
	// Constraints:
	//    - oneof:[DOWNLOADING,DONE,TIMEOUT,FAILED]
	DownloadStatus *string `json:"downloadStatus,omitempty" validate:"oneof=DOWNLOADING DONE TIMEOUT FAILED"`

	// SerialNumber
	// Switch Serial Number
	SerialNumber *string `json:"serialNumber,omitempty"`

	// SwitchId
	// Switch MAC Address
	SwitchId *string `json:"switchId,omitempty"`
}

func NewSwitchMTroubleShootingSupportLogStatus() *SwitchMTroubleShootingSupportLogStatus {
	m := new(SwitchMTroubleShootingSupportLogStatus)
	return m
}

// FindSupportLogBySwitchId
//
// Use this API to request ICX to prepare support log.
//
// Required Parameters:
// - switchId string
//		- required
func (s *SwitchMTroubleShootingService) FindSupportLogBySwitchId(ctx context.Context, switchId string) (*SwitchMCommonCreateResult, error) {
	var (
		req      *APIRequest
		resp     *SwitchMCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, switchId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindSupportLogBySwitchId, true)
	req.SetPathParameter("switchId", switchId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMCommonCreateResult()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindSupportLogDownloadBySwitchId
//
// Use this API to download support log.
//
// Required Parameters:
// - switchId string
//		- required
func (s *SwitchMTroubleShootingService) FindSupportLogDownloadBySwitchId(ctx context.Context, switchId string) (*SwitchMCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *SwitchMCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, switchId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindSupportLogDownloadBySwitchId, true)
	req.SetPathParameter("switchId", switchId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMCommonEmptyResult()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindSupportLogStatusBySwitchId
//
// Use this API to get the status of current support log request.
//
// Required Parameters:
// - switchId string
//		- required
func (s *SwitchMTroubleShootingService) FindSupportLogStatusBySwitchId(ctx context.Context, switchId string) (*SwitchMTroubleShootingSupportLogStatus, error) {
	var (
		req      *APIRequest
		resp     *SwitchMTroubleShootingSupportLogStatus
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, switchId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindSupportLogStatusBySwitchId, true)
	req.SetPathParameter("switchId", switchId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMTroubleShootingSupportLogStatus()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}
