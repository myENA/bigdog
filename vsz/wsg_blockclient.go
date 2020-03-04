package vsz

// API Version: v9_0

import (
	"context"
	"fmt"
	"net/http"
)

type WSGBlockClientService struct {
	apiClient *APIClient
}

func NewWSGBlockClientService(c *APIClient) *WSGBlockClientService {
	s := new(WSGBlockClientService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGBlockClientService() *WSGBlockClientService {
	return NewWSGBlockClientService(ss.apiClient)
}

// AddBlockClient
//
// Create new Block Clients by list.
//
// Request Body:
//	 - body *WSGProfileBulkBlockClient
func (s *WSGBlockClientService) AddBlockClient(ctx context.Context, body *WSGProfileBulkBlockClient) (WSGProfileCreateResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     WSGProfileCreateResultList
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
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGAddBlockClient), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = MakeWSGProfileCreateResultList()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// AddBlockClientByApMacByApMac
//
// Create a new Block Client by AP MAC.
//
// Request Body:
//	 - body *WSGProfileBlockClient
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGBlockClientService) AddBlockClientByApMacByApMac(ctx context.Context, body *WSGProfileBlockClient, apMac string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResult
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
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGAddBlockClientByApMacByApMac), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, &resp, err)
	return resp, rm, err
}

// DeleteBlockClient
//
// Delete Block Client List.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGBlockClientService) DeleteBlockClient(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonEmptyResult
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
	req = NewAPIRequest(http.MethodDelete, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGDeleteBlockClient), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
}

// DeleteBlockClientById
//
// Delete a Block Client.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGBlockClientService) DeleteBlockClientById(ctx context.Context, id string) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGDeleteBlockClientById), true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
}

// FindBlockClientById
//
// Retrieve a Block Client.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGBlockClientService) FindBlockClientById(ctx context.Context, id string) (*WSGProfileBlockClient, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileBlockClient
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGFindBlockClientById), true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileBlockClient()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindBlockClientByQueryCriteria
//
// Retrieve a list of Block Client.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGBlockClientService) FindBlockClientByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileBlockClientList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileBlockClientList
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
	req = NewAPIRequest(http.MethodPost, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGFindBlockClientByQueryCriteria), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileBlockClientList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindBlockClientByZoneByZoneId
//
// Retrieve a list of Block Client.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGBlockClientService) FindBlockClientByZoneByZoneId(ctx context.Context, zoneId string) (*WSGProfileBlockClientList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileBlockClientList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGFindBlockClientByZoneByZoneId), true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileBlockClientList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// PartialUpdateBlockClientById
//
// Modify a specific Block Client basic.
//
// Request Body:
//	 - body *WSGProfileModifyBlockClient
//
// Required Parameters:
// - id string
//		- required
func (s *WSGBlockClientService) PartialUpdateBlockClientById(ctx context.Context, body *WSGProfileModifyBlockClient, id string) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonEmptyResult
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
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPatch, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGPartialUpdateBlockClientById), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
}

// UpdateBlockClientById
//
// Modify a specific Block Client basic.
//
// Request Body:
//	 - body *WSGProfileModifyBlockClient
//
// Required Parameters:
// - id string
//		- required
func (s *WSGBlockClientService) UpdateBlockClientById(ctx context.Context, body *WSGProfileModifyBlockClient, id string) (*WSGCommonEmptyResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonEmptyResult
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
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, fmt.Sprintf("%s%s", s.apiClient.wsgPath, RouteWSGUpdateBlockClientById), true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
	return resp, rm, err
}
