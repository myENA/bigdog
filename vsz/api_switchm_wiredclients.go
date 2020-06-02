package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type SwitchMWiredclientsService struct {
	apiClient *APIClient
}

func NewSwitchMWiredclientsService(c *APIClient) *SwitchMWiredclientsService {
	s := new(SwitchMWiredclientsService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMWiredclientsService() *SwitchMWiredclientsService {
	return NewSwitchMWiredclientsService(ss.apiClient)
}

// AddSwitchClients
//
// Use this API command to retrieve all the wired clients connected to switch, currently managed by SmartZone.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMWiredclientsService) AddSwitchClients(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchConnectedDevicesQueryList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchConnectedDevicesQueryList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddSwitchClients, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchConnectedDevicesQueryList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddSwitchClientsAp
//
// Use this API command to retrieve all the Ruckus APs connected to switch, currently managed by SmartZone.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMWiredclientsService) AddSwitchClientsAp(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchConnectedAPsQueryList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchConnectedAPsQueryList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddSwitchClientsAp, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchConnectedAPsQueryList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
