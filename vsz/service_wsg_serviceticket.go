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

func (s *WSGServiceTicketService) AddServiceTicket(ctx context.Context) (*serviceticket.LoginResponse, error) {
}
