package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/sci"
)

type WSGSCIService struct {
	client *Client
}

func NewWSGSCIService(client *Client) *WSGSCIService {
	s := new(WSGSCIService)
	s.client = client
	return s
}

func (ss *WSGService) WSGSCIService() *WSGSCIService {
	serv := new(WSGSCIService)
	serv.client = ss.client
	return serv
}

func (s *WSGSCIService) AddSciSciProfileSciPriority(ctx context.Context) error {
}

func (s *WSGSCIService) FindSciSciEventCode(ctx context.Context) (*sci.SciEventCode, error) {
}

func (s *WSGSCIService) FindSciSciProfile(ctx context.Context) (*sci.SciProfileList, error) {
}

func (s *WSGSCIService) FindSciSciProfileById(ctx context.Context, id string) (*sci.SciProfile, error) {
}

func (s *WSGSCIService) PartialUpdateSciSciEnabled(ctx context.Context) error {
}
