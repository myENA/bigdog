package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGServiceTicketService struct {
	apiClient *APIClient
}

func NewWSGServiceTicketService(c *APIClient) *WSGServiceTicketService {
	s := new(WSGServiceTicketService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGServiceTicketService() *WSGServiceTicketService {
	return NewWSGServiceTicketService(ss.apiClient)
}

type WSGServiceTicketLoginRequest struct {
	// Password
	// Logon password
	// Constraints:
	//    - required
	Password *string `json:"password" validate:"required"`

	// Username
	// Logon user name
	// Constraints:
	//    - required
	Username *string `json:"username" validate:"required"`
}

func NewWSGServiceTicketLoginRequest() *WSGServiceTicketLoginRequest {
	m := new(WSGServiceTicketLoginRequest)
	return m
}

type WSGServiceTicketLoginResponse struct {
	ControllerVersion *string `json:"controllerVersion,omitempty"`

	// ServiceTicket
	// Logon authentication successful, the server generates a service ticket
	ServiceTicket *string `json:"serviceTicket,omitempty"`
}

func NewWSGServiceTicketLoginResponse() *WSGServiceTicketLoginResponse {
	m := new(WSGServiceTicketLoginResponse)
	return m
}

// AddServiceTicket
//
// Use this API command to log on to the controller and acquire a valid service ticket.
//
// Request Body:
//	 - body *WSGServiceTicketLoginRequest
func (s *WSGServiceTicketService) AddServiceTicket(ctx context.Context, body *WSGServiceTicketLoginRequest) (*WSGServiceTicketLoginResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGServiceTicketLoginResponse
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddServiceTicket, false)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGServiceTicketLoginResponse()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteServiceTicket
//
// Use this API command to log off of the controller.
//
// Required Parameters:
// - serviceTicket string
//		- required
func (s *WSGServiceTicketService) DeleteServiceTicket(ctx context.Context, serviceTicket string) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, serviceTicket, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteServiceTicket, false)
	req.SetQueryParameter("serviceTicket", []string{serviceTicket})
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
