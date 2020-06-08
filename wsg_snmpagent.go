package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGSnmpagentService struct {
	apiClient *VSZClient
}

func NewWSGSnmpagentService(c *VSZClient) *WSGSnmpagentService {
	s := new(WSGSnmpagentService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSnmpagentService() *WSGSnmpagentService {
	return NewWSGSnmpagentService(ss.apiClient)
}

// FindSystemSnmpAgent
//
// Retrieve SNMP Agent sertting.
func (s *WSGSnmpagentService) FindSystemSnmpAgent(ctx context.Context) (*WSGSystemSnmpAgentConfiguration, *APIResponseMeta, error) {
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
func (s *WSGSnmpagentService) UpdateSystemSnmpAgent(ctx context.Context, body *WSGSystemModifySnmpAgent) (*APIResponseMeta, error) {
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
