package bigdog

// API Version: v9_0

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
// Retrieve SNMP Agent sertting.
func (s *WSGSNMPAgentService) FindSystemSnmpAgent(ctx context.Context) (*WSGSystemSnmpAgentConfiguration, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSystemSnmpAgentConfiguration
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindSystemSnmpAgent, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemSnmpAgentConfiguration()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateSystemSnmpAgent
//
// Modify syslog server setting.
//
// Request Body:
//	 - body *WSGSystemModifySnmpAgent
func (s *WSGSNMPAgentService) UpdateSystemSnmpAgent(ctx context.Context, body *WSGSystemModifySnmpAgent) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateSystemSnmpAgent, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
