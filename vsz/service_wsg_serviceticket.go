package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/serviceticket"
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
	serv := new(WSGServiceTicketService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddServiceTicket
//
// Use this API command to log on to the controller and acquire a valid service ticket.
//
// Request Body:
//	 - body *serviceticket.LoginRequest
func (s *WSGServiceTicketService) AddServiceTicket(ctx context.Context, body *serviceticket.LoginRequest) (*serviceticket.LoginResponse, error) {
}

// DeleteServiceTicket
//
// Use this API command to log off of the controller.
//
// Query Parameters:
// - qServiceTicket string
//		- required
func (s *WSGServiceTicketService) DeleteServiceTicket(ctx context.Context, qServiceTicket string) error {
}
