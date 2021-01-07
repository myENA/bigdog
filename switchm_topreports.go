package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
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
// Operation ID: addSwitchTopByFirmware
//
// Use this API command to retrieves top N switch count based on firmware version.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMTopReportsService) AddSwitchTopByFirmware(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMSwitchTopSwitchesByFirmwareQueryResultListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMSwitchTopSwitchesByFirmwareQueryResultListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMSwitchTopSwitchesByFirmwareQueryResultListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMAddSwitchTopByFirmware, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMSwitchTopSwitchesByFirmwareQueryResultListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMSwitchTopSwitchesByFirmwareQueryResultListAPIResponse), err
}

// AddSwitchTopByModel
//
// Operation ID: addSwitchTopByModel
//
// Use this API command to retrieve top N switch count based on switch model.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMTopReportsService) AddSwitchTopByModel(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMSwitchTopSwitchesByModelQueryResultListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMSwitchTopSwitchesByModelQueryResultListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMSwitchTopSwitchesByModelQueryResultListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMAddSwitchTopByModel, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMSwitchTopSwitchesByModelQueryResultListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMSwitchTopSwitchesByModelQueryResultListAPIResponse), err
}
