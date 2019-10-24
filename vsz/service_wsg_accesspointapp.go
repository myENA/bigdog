package vsz

// API Version: v8_1

import (
	"context"
	"encoding/json"
	"github.com/myENA/ruckus-client/vsz/types/wsg/ap"
)

type WSGAccessPointAPPService struct {
	client *Client
}

func NewWSGAccessPointAPPService(client *Client) *WSGAccessPointAPPService {
	s := new(WSGAccessPointAPPService)
	s.client = client
	return s
}

func (ss *WSGService) WSGAccessPointAPPService() *WSGAccessPointAPPService {
	serv := new(WSGAccessPointAPPService)
	serv.client = ss.client
	return serv
}

// FindApsLineman
//
// Use this API command to retrieve the summary information of an AP. This is used by the Ruckus Wireless AP mobile app.
func (s *WSGAccessPointAPPService) FindApsLineman(ctx context.Context, qDomainId string, qIndex string, qListSize string, qShowAlarm string, qZoneId string) (*ap.ApLinemanSummary, error) {
}

// FindApsTotalCount
//
// Use this API command to retrieve the total AP count within a zone or a domain.
func (s *WSGAccessPointAPPService) FindApsTotalCount(ctx context.Context, qDomainId string, qZoneId string) error {
}

// FindLinemanWorkflow
//
// Use this API command to download the workflow file used by the Ruckus Wireless AP mobile app.
func (s *WSGAccessPointAPPService) FindLinemanWorkflow(ctx context.Context) (json.RawMessage, error) {
}
