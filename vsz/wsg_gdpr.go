package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
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
	return NewWSGGDPRService(ss.apiClient)
}

type WSGGDPRFtp struct {
	// FtpHost
	// IP/DN of FTP
	FtpHost *string `json:"ftpHost,omitempty"`

	// FtpPassword
	// Password for FTP login
	FtpPassword *string `json:"ftpPassword,omitempty"`

	// FtpPort
	// Port used by FTP
	// Constraints:
	//    - nullable
	//    - min:21
	//    - max:65535
	FtpPort *int `json:"ftpPort,omitempty" validate:"omitempty,gte=21,lte=65535"`

	// FtpProtocol
	// Protocol used
	// Constraints:
	//    - nullable
	//    - oneof:[FTP,SFTP]
	FtpProtocol *string `json:"ftpProtocol,omitempty" validate:"omitempty,oneof=FTP SFTP"`

	// FtpRemoteDirectory
	// Destination directory used for file upload
	FtpRemoteDirectory *string `json:"ftpRemoteDirectory,omitempty"`

	// FtpUserName
	// Username for FTP login
	FtpUserName *string `json:"ftpUserName,omitempty"`
}

type WSGGDPRReport struct {
	// Action
	// Request action
	// Constraints:
	//    - nullable
	//    - oneof:[SEARCH,DELETE,INTERRUPT,PROGRESS]
	Action *string `json:"action,omitempty" validate:"omitempty,oneof=SEARCH DELETE INTERRUPT PROGRESS"`

	// ClientMac
	// Client mac
	ClientMac *string `json:"clientMac,omitempty"`

	Ftp *WSGGDPRFtp `json:"ftp,omitempty"`
}

// AddGdprReport
//
// Use this API command to execute a client-related data search or delete task and upload a report to FTP. Also use this API to check task progress or to interrupt it.
//
// Request Body:
//	 - body *WSGGDPRReport
func (s *WSGGDPRService) AddGdprReport(ctx context.Context, body *WSGGDPRReport) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}
