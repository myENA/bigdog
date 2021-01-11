package bigdog

// API Version: v9_1

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

type SwitchMPortsService struct {
	apiClient *VSZClient
}

func NewSwitchMPortsService(c *VSZClient) *SwitchMPortsService {
	s := new(SwitchMPortsService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMPortsService() *SwitchMPortsService {
	return NewSwitchMPortsService(ss.apiClient)
}

// AddSwitchPortsDetails
//
// Use this API command to retrieve all the switch ports and its details currently managed by SmartZone.
//
// Operation ID: addSwitchPortsDetails
// Operation path: /switch/ports/details
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMPortsService) AddSwitchPortsDetails(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMSwitchPortDetailsQueryResultListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMSwitchPortDetailsQueryResultListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMSwitchPortDetailsQueryResultListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddSwitchPortsDetails, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMSwitchPortDetailsQueryResultListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMSwitchPortDetailsQueryResultListAPIResponse), err
}

// AddSwitchPortsDetailsExport
//
// Download CSV of Switch Port Details
//
// Operation ID: addSwitchPortsDetailsExport
// Operation path: /switch/ports/details/export
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMPortsService) AddSwitchPortsDetailsExport(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*FileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newFileAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*FileAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddSwitchPortsDetailsExport, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "application/x-www-form-urlencoded")
	req.Header.Set(headerKeyAccept, "*/*")
	if b, err := json.Marshal(body); err != nil {
		return resp.(*FileAPIResponse), err
	} else if err = req.SetBody(bytes.NewBufferString((url.Values{"json": []string{string(b)}}).Encode())); err != nil {
		return resp.(*FileAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*FileAPIResponse), err
}

// AddSwitchPortsSummary
//
// Use this API command to retrieve ports summary based on status, speed of a switch, currently managed by SmartZone.
//
// Operation ID: addSwitchPortsSummary
// Operation path: /switch/ports/summary
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMPortsService) AddSwitchPortsSummary(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMSwitchPortsSummaryQueryResultListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMSwitchPortsSummaryQueryResultListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMSwitchPortsSummaryQueryResultListAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddSwitchPortsSummary, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMSwitchPortsSummaryQueryResultListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMSwitchPortsSummaryQueryResultListAPIResponse), err
}
