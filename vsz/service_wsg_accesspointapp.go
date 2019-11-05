package vsz

// API Version: v8_1

import (
	"context"
	"encoding/json"
	"github.com/myENA/ruckus-client/vsz/types/wsg/ap"
)

type WSGAccessPointAPPService struct {
	apiClient *APIClient
}

func NewWSGAccessPointAPPService(c *APIClient) *WSGAccessPointAPPService {
	s := new(WSGAccessPointAPPService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAccessPointAPPService() *WSGAccessPointAPPService {
	serv := new(WSGAccessPointAPPService)
	serv.apiClient = ss.apiClient
	return serv
}

// FindApsLineman
//
// Use this API command to retrieve the summary information of an AP. This is used by the Ruckus Wireless AP mobile app.
//
// Query Parameters:
// - qDomainId string
// - qIndex string
// - qListSize string
// - qShowAlarm string
// - qZoneId string
func (s *WSGAccessPointAPPService) FindApsLineman(ctx context.Context, qDomainId string, qIndex string, qListSize string, qShowAlarm string, qZoneId string) (*ap.ApLinemanSummary, error) {
}

// FindApsTotalCount
//
// Use this API command to retrieve the total AP count within a zone or a domain.
//
// Query Parameters:
// - qDomainId string
// - qZoneId string
func (s *WSGAccessPointAPPService) FindApsTotalCount(ctx context.Context, qDomainId string, qZoneId string) error {
}

// FindLinemanWorkflow
//
// Use this API command to download the workflow file used by the Ruckus Wireless AP mobile app.
func (s *WSGAccessPointAPPService) FindLinemanWorkflow(ctx context.Context) (json.RawMessage, error) {
}

// UpdateLinemanWorkflow
//
// Use this API command to upload a workflow file used by the Ruckus Wireless AP mobile app.
func (s *WSGAccessPointAPPService) UpdateLinemanWorkflow(ctx context.Context) error {
}
