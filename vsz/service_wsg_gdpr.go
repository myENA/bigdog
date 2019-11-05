package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/gdpr"
)

type WSGGDPRService struct {
	apiClient *APIClient
}

func NewWSGGDPRService(c *APIClient) *WSGGDPRService {
	s := new(WSGGDPRService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGGDPRService() *WSGGDPRService {
	serv := new(WSGGDPRService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddGdprReport
//
// Use this API command to execute a client-related data search or delete task and upload a report to FTP. Also use this API to check task progress or to interrupt it.
//
// Request Body:
//	 - body *gdpr.Report
func (s *WSGGDPRService) AddGdprReport(ctx context.Context, body *gdpr.Report) error {
}
