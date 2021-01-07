package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGSNMPAgentService struct {
	apiClient *VSZClient
}

func NewWSGSNMPAgentService(c *VSZClient) *WSGSNMPAgentService {
	s := new(WSGSNMPAgentService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSNMPAgentService() *WSGSNMPAgentService {
	return NewWSGSNMPAgentService(ss.apiClient)
}

// FindSystemSnmpAgent
//
// Operation ID: findSystemSnmpAgent
//
// Retrieve SNMP Agent sertting.
func (s *WSGSNMPAgentService) FindSystemSnmpAgent(ctx context.Context, mutators ...RequestMutator) (*WSGSystemSnmpAgentConfigurationAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSystemSnmpAgentConfigurationAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSystemSnmpAgentConfigurationAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindSystemSnmpAgent, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemSnmpAgentConfigurationAPIResponse), err
}

// UpdateSystemSnmpAgent
//
// Operation ID: updateSystemSnmpAgent
//
// Modify syslog server setting.
//
// Request Body:
//	 - body *WSGSystemModifySnmpAgent
func (s *WSGSNMPAgentService) UpdateSystemSnmpAgent(ctx context.Context, body *WSGSystemModifySnmpAgent, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateSystemSnmpAgent, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
