package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGApplicationLogAndStatusService struct {
	apiClient *APIClient
}

func NewWSGApplicationLogAndStatusService(c *APIClient) *WSGApplicationLogAndStatusService {
	s := new(WSGApplicationLogAndStatusService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGApplicationLogAndStatusService() *WSGApplicationLogAndStatusService {
	return NewWSGApplicationLogAndStatusService(ss.apiClient)
}

// FindApplicationsByBladeUUID
//
// Use this API command to retrieve a list of application log and status.
//
// Required Parameters:
// - bladeUUID string
//		- required
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGApplicationLogAndStatusService) FindApplicationsByBladeUUID(ctx context.Context, bladeUUID string, optionalParams map[string]interface{}) (*WSGAdministrationApplicationLogAndStatusList, error) {
	var (
		req  *APIRequest
		resp *WSGAdministrationApplicationLogAndStatusList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, bladeUUID, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApplicationsByBladeUUID, true)
	req.SetPathParameter("bladeUUID", bladeUUID)
	if v, ok := optionalParams["index"]; ok {
		req.AddQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok {
		req.AddQueryParameter("listSize", v)
	}
}

// FindApplicationsDownloadByBladeUUID
//
// Use this API command to download logs of the application.
//
// Required Parameters:
// - appName string
//		- required
// - bladeUUID string
//		- required
//
// Optional Parameters:
// - logFileName string
//		- nullable
func (s *WSGApplicationLogAndStatusService) FindApplicationsDownloadByBladeUUID(ctx context.Context, appName string, bladeUUID string, optionalParams map[string]interface{}) ([]byte, error) {
	var (
		req  *APIRequest
		resp []byte
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, bladeUUID, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, appName, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApplicationsDownloadByBladeUUID, true)
	req.SetQueryParameter("appName", appName)
	req.SetPathParameter("bladeUUID", bladeUUID)
	if v, ok := optionalParams["logFileName"]; ok {
		req.AddQueryParameter("logFileName", v)
	}
}

// FindApplicationsDownloadsnapByBladeUUID
//
// Use this API command to download snapshot logs.
//
// Required Parameters:
// - bladeUUID string
//		- required
func (s *WSGApplicationLogAndStatusService) FindApplicationsDownloadsnapByBladeUUID(ctx context.Context, bladeUUID string) ([]byte, error) {
	var (
		req  *APIRequest
		resp []byte
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, bladeUUID, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApplicationsDownloadsnapByBladeUUID, true)
	req.SetPathParameter("bladeUUID", bladeUUID)
}

// PartialUpdateApplications
//
// Use this API command to modify log level of specified application.
//
// Request Body:
//	 - body *WSGAdministrationModifyLogLevel
func (s *WSGApplicationLogAndStatusService) PartialUpdateApplications(ctx context.Context, body *WSGAdministrationModifyLogLevel) error {
	var (
		req *APIRequest
		err error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateApplications, true)
	if err = req.SetBody(body); err != nil {
		return err
	}
}
