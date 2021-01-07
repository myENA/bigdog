package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGGDPRService struct {
	apiClient *VSZClient
}

func NewWSGGDPRService(c *VSZClient) *WSGGDPRService {
	s := new(WSGGDPRService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGGDPRService() *WSGGDPRService {
	return NewWSGGDPRService(ss.apiClient)
}

// WSGGDPRFtp
//
// Definition: gdpr_ftp
type WSGGDPRFtp struct {
	// FtpHost
	// IP/DN of FTP
	// Constraints:
	//    - required
	FtpHost *string `json:"ftpHost"`

	// FtpPassword
	// Password for FTP login
	FtpPassword *string `json:"ftpPassword,omitempty"`

	// FtpPort
	// Port used by FTP
	// Constraints:
	//    - required
	//    - min:21
	//    - max:65535
	FtpPort *int `json:"ftpPort"`

	// FtpProtocol
	// Protocol used
	// Constraints:
	//    - oneof:[FTP,SFTP]
	FtpProtocol *string `json:"ftpProtocol,omitempty"`

	// FtpRemoteDirectory
	// Destination directory used for file upload
	FtpRemoteDirectory *string `json:"ftpRemoteDirectory,omitempty"`

	// FtpUserName
	// Username for FTP login
	FtpUserName *string `json:"ftpUserName,omitempty"`
}

func NewWSGGDPRFtp() *WSGGDPRFtp {
	m := new(WSGGDPRFtp)
	return m
}

// WSGGDPRReport
//
// Definition: gdpr_report
type WSGGDPRReport struct {
	// Action
	// Request action
	// Constraints:
	//    - required
	//    - oneof:[SEARCH,DELETE,INTERRUPT,PROGRESS]
	Action *string `json:"action"`

	// ClientMac
	// Client mac
	// Constraints:
	//    - required
	ClientMac *string `json:"clientMac"`

	Ftp *WSGGDPRFtp `json:"ftp,omitempty"`
}

func NewWSGGDPRReport() *WSGGDPRReport {
	m := new(WSGGDPRReport)
	return m
}

// AddGdprReport
//
// Operation ID: addGdprReport
//
// Use this API command to execute a client-related data search or delete task and upload a report to FTP. Also use this API to check task progress or to interrupt it.
//
// Request Body:
//	 - body *WSGGDPRReport
func (s *WSGGDPRService) AddGdprReport(ctx context.Context, body *WSGGDPRReport, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawAPIResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddGdprReport, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawAPIResponse)
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
