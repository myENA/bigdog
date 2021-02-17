package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
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
// Retrieve SNMP Agent sertting.
//
// Operation ID: findSystemSnmpAgent
// Operation path: /system/snmpAgent
// Success code: 200 (OK)
func (s *WSGSNMPAgentService) FindSystemSnmpAgent(ctx context.Context, mutators ...RequestMutator) (*WSGSystemSnmpAgentConfigurationAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGSystemSnmpAgentConfigurationAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindSystemSnmpAgent, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGSystemSnmpAgentConfigurationAPIResponse), err
}

// UpdateSystemSnmpAgent
//
// Modify syslog server setting.
//
// Operation ID: updateSystemSnmpAgent
// Operation path: /system/snmpAgent
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGSystemModifySnmpAgent
func (s *WSGSNMPAgentService) UpdateSystemSnmpAgent(ctx context.Context, body *WSGSystemModifySnmpAgent, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateSystemSnmpAgent, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}
