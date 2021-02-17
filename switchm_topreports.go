package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
)

type SwitchMTopReportsService struct {
	apiClient *VSZClient
}

func NewSwitchMTopReportsService(c *VSZClient) *SwitchMTopReportsService {
	s := new(SwitchMTopReportsService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMTopReportsService() *SwitchMTopReportsService {
	return NewSwitchMTopReportsService(ss.apiClient)
}

// AddSwitchTopByFirmware
//
// Use this API command to retrieves top N switch count based on firmware version.
//
// Operation ID: addSwitchTopByFirmware
// Operation path: /switch/top/byFirmware
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMTopReportsService) AddSwitchTopByFirmware(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMSwitchTopSwitchesByFirmwareQueryResultListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMSwitchTopSwitchesByFirmwareQueryResultListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddSwitchTopByFirmware, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*SwitchMSwitchTopSwitchesByFirmwareQueryResultListAPIResponse), err
}

// AddSwitchTopByModel
//
// Use this API command to retrieve top N switch count based on switch model.
//
// Operation ID: addSwitchTopByModel
// Operation path: /switch/top/byModel
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMTopReportsService) AddSwitchTopByModel(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMSwitchTopSwitchesByModelQueryResultListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMSwitchTopSwitchesByModelQueryResultListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddSwitchTopByModel, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*SwitchMSwitchTopSwitchesByModelQueryResultListAPIResponse), err
}
