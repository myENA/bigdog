package bigdog

// API Version: v9_1

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type SwitchMWiredClientsService struct {
	apiClient *VSZClient
}

func NewSwitchMWiredClientsService(c *VSZClient) *SwitchMWiredClientsService {
	s := new(SwitchMWiredClientsService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMWiredClientsService() *SwitchMWiredClientsService {
	return NewSwitchMWiredClientsService(ss.apiClient)
}

// AddSwitchClients
//
// Use this API command to retrieve all the wired clients connected to switch, currently managed by SmartZone.
//
// Operation ID: addSwitchClients
// Operation path: /switch/clients
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMWiredClientsService) AddSwitchClients(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMSwitchConnectedDevicesQueryListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMSwitchConnectedDevicesQueryListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddSwitchClients, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*SwitchMSwitchConnectedDevicesQueryListAPIResponse), err
}

// AddSwitchClientsAp
//
// Use this API command to retrieve all the Ruckus APs connected to switch, currently managed by SmartZone.
//
// Operation ID: addSwitchClientsAp
// Operation path: /switch/clients/ap
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMWiredClientsService) AddSwitchClientsAp(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMSwitchConnectedAPsQueryListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMSwitchConnectedAPsQueryListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddSwitchClientsAp, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*SwitchMSwitchConnectedAPsQueryListAPIResponse), err
}

// AddSwitchClientsAPExport
//
// Download CSV of AP's discovered via LLDP
//
// Operation ID: addSwitchClientsAPExport
// Operation path: /switch/clients/ap/export
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMWiredClientsService) AddSwitchClientsAPExport(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*FileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newFileAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddSwitchClientsAPExport, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "application/x-www-form-urlencoded")
	req.Header.Set(headerKeyAccept, "*/*")
	if b, err := json.Marshal(body); err != nil {
		return resp.(*FileAPIResponse), err
	} else {
		req.SetBody(url.Values{"json": []string{string(b)}})
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*FileAPIResponse), err
}

// AddSwitchClientsExport
//
// Download CSV of wired clients discovered via LLDP
//
// Operation ID: addSwitchClientsExport
// Operation path: /switch/clients/export
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMWiredClientsService) AddSwitchClientsExport(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*FileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newFileAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddSwitchClientsExport, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "application/x-www-form-urlencoded")
	req.Header.Set(headerKeyAccept, "*/*")
	if b, err := json.Marshal(body); err != nil {
		return resp.(*FileAPIResponse), err
	} else {
		req.SetBody(url.Values{"json": []string{string(b)}})
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*FileAPIResponse), err
}
