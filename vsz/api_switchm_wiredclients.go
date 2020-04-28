package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type SwitchMWiredClientsService struct {
	apiClient *APIClient
}

func NewSwitchMWiredClientsService(c *APIClient) *SwitchMWiredClientsService {
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
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMWiredClientsService) AddSwitchClients(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchConnectedDevicesQueryList, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
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
func (s *SwitchMWiredClientsService) AddSwitchClientsAp(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchConnectedAPsQueryList, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
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
