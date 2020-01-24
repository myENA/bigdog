package vsz

// API Version: v8_1

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
func (s *WSGServiceTicketService) AddServiceTicket(ctx context.Context, body *WSGServiceTicketLoginRequest) (*WSGServiceTicketLoginResponse, error) {
	var (
		req  *APIRequest
		resp *WSGServiceTicketLoginResponse
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddServiceTicket, false)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteServiceTicket
//
// Use this API command to log off of the controller.
//
// Required Parameters:
// - serviceTicket string
//		- required
func (s *WSGServiceTicketService) DeleteServiceTicket(ctx context.Context, serviceTicket string) (interface{}, error) {
	var (
		req  *APIRequest
		resp interface{}
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, serviceTicket, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteServiceTicket, false)
}
