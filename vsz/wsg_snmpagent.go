package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGSNMPAgentService struct {
	apiClient *APIClient
}

func NewWSGSNMPAgentService(c *APIClient) *WSGSNMPAgentService {
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
func (s *WSGSNMPAgentService) FindSystemSnmpAgent(ctx context.Context) (*WSGSystemSnmpAgentConfiguration, error) {
	var (
		req  *APIRequest
		resp *WSGSystemSnmpAgentConfiguration
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindSystemSnmpAgent, true)
}

// UpdateSystemSnmpAgent
//
// Modify syslog server setting.
//
// Request Body:
//	 - body *WSGSystemModifySnmpAgent
func (s *WSGSNMPAgentService) UpdateSystemSnmpAgent(ctx context.Context, body *WSGSystemModifySnmpAgent) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateSystemSnmpAgent, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}
