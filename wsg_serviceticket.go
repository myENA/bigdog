package bigdog

// API Version: v9_1

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"
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

func newWSGServiceTicketLoginResponseAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGServiceTicketLoginResponseAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGServiceTicketLoginResponseAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGServiceTicketLoginResponse)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGServiceTicketLoginResponse() *WSGServiceTicketLoginResponse {
	m := new(WSGServiceTicketLoginResponse)
	return m
}

// AddServiceTicket
//
// Use this API command to log on to the controller and acquire a valid service ticket.
//
// Operation ID: addServiceTicket
// Operation path: /serviceTicket
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGServiceTicketLoginRequest
func (s *WSGServiceTicketService) AddServiceTicket(ctx context.Context, body *WSGServiceTicketLoginRequest, mutators ...RequestMutator) (*WSGServiceTicketLoginResponseAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGServiceTicketLoginResponseAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddServiceTicket, false)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGServiceTicketLoginResponseAPIResponse), err
}

// DeleteServiceTicket
//
// Use this API command to log off of the controller.
//
// Operation ID: deleteServiceTicket
// Operation path: /serviceTicket
// Success code: 200 (OK)
//
// Required parameters:
// - serviceTicket string
//		- required
func (s *WSGServiceTicketService) DeleteServiceTicket(ctx context.Context, serviceTicket string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteServiceTicket, false)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.QueryParams.Set("serviceTicket", serviceTicket)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*RawAPIResponse), err
}
