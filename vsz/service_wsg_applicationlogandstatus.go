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

// FindApplicationsByBladeUUID
//
// Use this API command to retrieve a list of application log and status.
//
// Path Parameters:
// - pBladeUUID string
//		- required
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGApplicationLogAndStatusService) FindApplicationsByBladeUUID(ctx context.Context, pBladeUUID string, qIndex string, qListSize string) (*administration.ApplicationLogAndStatusList, error) {
}

// FindApplicationsDownloadByBladeUUID
//
// Use this API command to download logs of the application.
//
// Path Parameters:
// - pBladeUUID string
//		- required
//
// Query Parameters:
// - qAppName string
//		- required
// - qLogFileName string
func (s *WSGApplicationLogAndStatusService) FindApplicationsDownloadByBladeUUID(ctx context.Context, pBladeUUID string, qAppName string, qLogFileName string) (json.RawMessage, error) {
}

// FindApplicationsDownloadsnapByBladeUUID
//
// Use this API command to download snapshot logs.
//
// Path Parameters:
// - pBladeUUID string
//		- required
func (s *WSGApplicationLogAndStatusService) FindApplicationsDownloadsnapByBladeUUID(ctx context.Context, pBladeUUID string) (json.RawMessage, error) {
}

// PartialUpdateApplications
//
// Use this API command to modify log level of specified application.
//
// Request Body:
//	 - body *administration.ModifyLogLevel
func (s *WSGApplicationLogAndStatusService) PartialUpdateApplications(ctx context.Context, body *administration.ModifyLogLevel) error {
}
