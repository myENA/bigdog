package vsz

// API Version: v8_1

import (
	"context"
	"encoding/json"
	"github.com/myENA/ruckus-client/vsz/types/wsg/administration"
)

type WSGApplicationLogAndStatusService struct {
	client *Client
}

func NewWSGApplicationLogAndStatusService(client *Client) *WSGApplicationLogAndStatusService {
	s := new(WSGApplicationLogAndStatusService)
	s.client = client
	return s
}

func (ss *WSGService) WSGApplicationLogAndStatusService() *WSGApplicationLogAndStatusService {
	serv := new(WSGApplicationLogAndStatusService)
	serv.client = ss.client
	return serv
}

func (s *WSGApplicationLogAndStatusService) FindApplicationsByBladeUUID(ctx context.Context, pBladeUUID string, qIndex string, qListSize string) (*administration.ApplicationLogAndStatusList, error) {
}

func (s *WSGApplicationLogAndStatusService) FindApplicationsDownloadByBladeUUID(ctx context.Context, pBladeUUID string, qAppName string, qLogFileName string) (json.RawMessage, error) {
}

func (s *WSGApplicationLogAndStatusService) FindApplicationsDownloadsnapByBladeUUID(ctx context.Context, pBladeUUID string) (json.RawMessage, error) {
}

func (s *WSGApplicationLogAndStatusService) PartialUpdateApplications(ctx context.Context, body *administration.ModifyLogLevel) error {
}
