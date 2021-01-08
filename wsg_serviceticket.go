package bigdog

// API Version: v9_1

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type WSGServiceTicketService struct {
	apiClient *VSZClient
}

func NewWSGServiceTicketService(c *VSZClient) *WSGServiceTicketService {
	s := new(WSGServiceTicketService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGServiceTicketService() *WSGServiceTicketService {
	return NewWSGServiceTicketService(ss.apiClient)
}

// WSGServiceTicketLoginRequest
//
// Definition: serviceTicket_loginRequest
type WSGServiceTicketLoginRequest struct {
	// Password
	// Logon password
	// Constraints:
	//    - required
	Password *string `json:"password"`

	// Username
	// Logon user name
	// Constraints:
	//    - required
	Username *string `json:"username"`
}

func NewWSGServiceTicketLoginRequest() *WSGServiceTicketLoginRequest {
	m := new(WSGServiceTicketLoginRequest)
	return m
}

// WSGServiceTicketLoginResponse
//
// Definition: serviceTicket_loginResponse
type WSGServiceTicketLoginResponse struct {
	ControllerVersion *string `json:"controllerVersion,omitempty"`

	// ServiceTicket
	// Logon authentication successful, the server generates a service ticket
	ServiceTicket *string `json:"serviceTicket,omitempty"`
}

type WSGServiceTicketLoginResponseAPIResponse struct {
	*RawAPIResponse
	Data *WSGServiceTicketLoginResponse
}

func newWSGServiceTicketLoginResponseAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGServiceTicketLoginResponseAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGServiceTicketLoginResponseAPIResponse) Hydrate() error {
	r.Data = new(WSGServiceTicketLoginResponse)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGServiceTicketLoginResponse() *WSGServiceTicketLoginResponse {
	m := new(WSGServiceTicketLoginResponse)
	return m
}

// AddServiceTicket
//
// Operation ID: addServiceTicket
//
// Use this API command to log on to the controller and acquire a valid service ticket.
//
// Request Body:
//	 - body *WSGServiceTicketLoginRequest
func (s *WSGServiceTicketService) AddServiceTicket(ctx context.Context, body *WSGServiceTicketLoginRequest, mutators ...RequestMutator) (*WSGServiceTicketLoginResponseAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGServiceTicketLoginResponseAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGServiceTicketLoginResponseAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddServiceTicket, false)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGServiceTicketLoginResponseAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGServiceTicketLoginResponseAPIResponse), err
}

// DeleteServiceTicket
//
// Operation ID: deleteServiceTicket
//
// Use this API command to log off of the controller.
//
// Required Parameters:
// - serviceTicket string
//		- required
func (s *WSGServiceTicketService) DeleteServiceTicket(ctx context.Context, serviceTicket string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteServiceTicket, false)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.QueryParams.Set("serviceTicket", serviceTicket)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
