package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGBlockclientService struct {
	apiClient *APIClient
}

func NewWSGBlockclientService(c *APIClient) *WSGBlockclientService {
	s := new(WSGBlockclientService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGBlockclientService() *WSGBlockclientService {
	return NewWSGBlockclientService(ss.apiClient)
}

// AddBlockClient
//
// Create new Block Clients by list.
//
// Request Body:
//	 - body *WSGProfileBulkBlockClient
func (s *WSGBlockclientService) AddBlockClient(ctx context.Context, body *WSGProfileBulkBlockClient) (WSGProfileCreateResultList, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddBlockClient, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = MakeWSGProfileCreateResultList()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, &resp, err)
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
func (s *WSGBlockclientService) AddBlockClientByApMacByApMac(ctx context.Context, body *WSGProfileBlockClient, apMac string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddBlockClientByApMacByApMac, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteBlockClient
//
// Delete Block Client List.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGBlockclientService) DeleteBlockClient(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteBlockClient, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteBlockClientById
//
// Delete a Block Client.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGBlockclientService) DeleteBlockClientById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteBlockClientById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindBlockClientById
//
// Retrieve a Block Client.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGBlockclientService) FindBlockClientById(ctx context.Context, id string) (*WSGProfileBlockClient, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindBlockClientById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileBlockClient()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindBlockClientByQueryCriteria
//
// Retrieve a list of Block Client.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGBlockclientService) FindBlockClientByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileBlockClientList, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindBlockClientByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileBlockClientList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindBlockClientByZoneByZoneId
//
// Retrieve a list of Block Client.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGBlockclientService) FindBlockClientByZoneByZoneId(ctx context.Context, zoneId string) (*WSGProfileBlockClientList, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindBlockClientByZoneByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileBlockClientList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
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
func (s *WSGBlockclientService) PartialUpdateBlockClientById(ctx context.Context, body *WSGProfileModifyBlockClient, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateBlockClientById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
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
func (s *WSGBlockclientService) UpdateBlockClientById(ctx context.Context, body *WSGProfileModifyBlockClient, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateBlockClientById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
