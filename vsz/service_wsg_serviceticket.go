package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/serviceticket"
)

type WSGServiceTicketService struct {
	client *Client
}

func NewWSGServiceTicketService(client *Client) *WSGServiceTicketService {
	s := new(WSGServiceTicketService)
	s.client = client
	return s
}

func (ss *WSGService) WSGServiceTicketService() *WSGServiceTicketService {
	serv := new(WSGServiceTicketService)
	serv.client = ss.client
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
