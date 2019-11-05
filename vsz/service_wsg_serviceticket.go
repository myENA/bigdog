package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/serviceticket"
	"gopkg.in/go-playground/validator.v9"
)

type WSGServiceTicketService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGServiceTicketService(c *APIClient) *WSGServiceTicketService {
	s := new(WSGServiceTicketService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGServiceTicketService() *WSGServiceTicketService {
	return NewWSGServiceTicketService(ss.apiClient)
}

// AddServiceTicket
//
// Use this API command to log on to the controller and acquire a valid service ticket.
//
// Request Body:
//	 - body *serviceticket.LoginRequest
func (s *WSGServiceTicketService) AddServiceTicket(ctx context.Context, body *serviceticket.LoginRequest) (*serviceticket.LoginResponse, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}

// DeleteServiceTicket
//
// Use this API command to log off of the controller.
//
// Query Parameters:
// - qServiceTicket string
//		- required
func (s *WSGServiceTicketService) DeleteServiceTicket(ctx context.Context, qServiceTicket string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}
