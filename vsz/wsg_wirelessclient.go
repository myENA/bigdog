package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGWirelessClientService struct {
	apiClient *APIClient
}

func NewWSGWirelessClientService(c *APIClient) *WSGWirelessClientService {
	s := new(WSGWirelessClientService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGWirelessClientService() *WSGWirelessClientService {
	return NewWSGWirelessClientService(ss.apiClient)
}

// AddClientsBulkDeauth
//
// Use this API command to bulk deauth client.
//
// Request Body:
//	 - body *WSGClientDeAuthClientList
func (s *WSGWirelessClientService) AddClientsBulkDeauth(ctx context.Context, body *WSGClientDeAuthClientList) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddClientsBulkDeauth, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// AddClientsBulkDisconnect
//
// Use this API command to bulk disconnect client.
//
// Request Body:
//	 - body *WSGClientDisconnectClientList
func (s *WSGWirelessClientService) AddClientsBulkDisconnect(ctx context.Context, body *WSGClientDisconnectClientList) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddClientsBulkDisconnect, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// AddClientsByWlanNameByWlanname
//
// Use this API command to query client by wlan name.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required Parameters:
// - wlanname string
//		- required
func (s *WSGWirelessClientService) AddClientsByWlanNameByWlanname(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, wlanname string) (*WSGClientQueryList, error) {
	var (
		req      *APIRequest
		resp     *WSGClientQueryList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, wlanname, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddClientsByWlanNameByWlanname, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("wlanname", wlanname)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGClientQueryList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// AddClientsDeauth
//
// Use this API command to deauth client.
//
// Request Body:
//	 - body *WSGClientDeAuthClient
func (s *WSGWirelessClientService) AddClientsDeauth(ctx context.Context, body *WSGClientDeAuthClient) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddClientsDeauth, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// AddClientsDisconnect
//
// Use this API command to disconnect client.
//
// Request Body:
//	 - body *WSGClientDisconnectClient
func (s *WSGWirelessClientService) AddClientsDisconnect(ctx context.Context, body *WSGClientDisconnectClient) (*WSGCommonEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *WSGCommonEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddClientsDisconnect, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonEmptyResult()
	return resp, handleResponse(req, http.StatusNoContent, httpResp, &resp, err)
}

// FindApsOperationalClientByApMac
//
// Use this API command to retrieve the client list per AP.
//
// Required Parameters:
// - apMac string
//		- required
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGWirelessClientService) FindApsOperationalClientByApMac(ctx context.Context, apMac string, optionalParams map[string][]string) (*WSGAPClientList, error) {
	var (
		req      *APIRequest
		resp     *WSGAPClientList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApsOperationalClientByApMac, true)
	req.SetPathParameter("apMac", apMac)
	if v, ok := optionalParams["index"]; ok {
		req.AddQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok {
		req.AddQueryParameter("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAPClientList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindApsOperationalClientTotalCountByApMac
//
// Use this API command to retrieve the total client count per AP.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGWirelessClientService) FindApsOperationalClientTotalCountByApMac(ctx context.Context, apMac string) (interface{}, error) {
	var (
		req      *APIRequest
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApsOperationalClientTotalCountByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	return resp, handleResponse(req, http.StatusOK, httpResp, resp, err)
}

// FindHistoricalclientByQueryCriteria
//
// Use this API command to retrive historical client.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGWirelessClientService) FindHistoricalclientByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGClientHistoricalClientList, error) {
	var (
		req      *APIRequest
		resp     *WSGClientHistoricalClientList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindHistoricalclientByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGClientHistoricalClientList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}
