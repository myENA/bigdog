package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
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
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMPortsService) AddSwitchPortsDetails(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchPortDetailsQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchPortDetailsQueryResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddSwitchPortsDetails, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchPortDetailsQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddSwitchPortsSummary
//
// Use this API command to retrieve ports summary based on status, speed of a switch, currently managed by SmartZone.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMPortsService) AddSwitchPortsSummary(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchPortsSummaryQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchPortsSummaryQueryResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddSwitchPortsSummary, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchPortsSummaryQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
