package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type SwitchMTopReportsService struct {
	apiClient *APIClient
}

func NewSwitchMTopReportsService(c *APIClient) *SwitchMTopReportsService {
	s := new(SwitchMTopReportsService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMTopReportsService() *SwitchMTopReportsService {
	return NewSwitchMTopReportsService(ss.apiClient)
}

// AddSwitchTopByFirmware
//
// Use this API command to retrieves top N switch count based on firmware version.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMTopReportsService) AddSwitchTopByFirmware(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchTopSwitchesByFirmwareQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchTopSwitchesByFirmwareQueryResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddSwitchTopByFirmware, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchTopSwitchesByFirmwareQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddSwitchTopByModel
//
// Use this API command to retrieve top N switch count based on switch model.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMTopReportsService) AddSwitchTopByModel(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchTopSwitchesByModelQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchTopSwitchesByModelQueryResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddSwitchTopByModel, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchTopSwitchesByModelQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
