package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGWirelessclientService struct {
	apiClient *VSZClient
}

func NewWSGWirelessclientService(c *VSZClient) *WSGWirelessclientService {
	s := new(WSGWirelessclientService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGWirelessclientService() *WSGWirelessclientService {
	return NewWSGWirelessclientService(ss.apiClient)
}

// AddClientsBulkDeauth
//
// Use this API command to bulk deauth client.
//
// Request Body:
//	 - body *WSGClientDeAuthClientList
func (s *WSGWirelessclientService) AddClientsBulkDeauth(ctx context.Context, body *WSGClientDeAuthClientList) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddClientsBulkDeauth, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// AddClientsBulkDisconnect
//
// Use this API command to bulk disconnect client.
//
// Request Body:
//	 - body *WSGClientDisconnectClientList
func (s *WSGWirelessclientService) AddClientsBulkDisconnect(ctx context.Context, body *WSGClientDisconnectClientList) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddClientsBulkDisconnect, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
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
func (s *WSGWirelessclientService) AddClientsByWlanNameByWlanname(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, wlanname string) (*WSGClientQueryList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGClientQueryList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddClientsByWlanNameByWlanname, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("wlanname", wlanname)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGClientQueryList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddClientsDeauth
//
// Use this API command to deauth client.
//
// Request Body:
//	 - body *WSGClientDeAuthClient
func (s *WSGWirelessclientService) AddClientsDeauth(ctx context.Context, body *WSGClientDeAuthClient) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddClientsDeauth, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// AddClientsDisconnect
//
// Use this API command to disconnect client.
//
// Request Body:
//	 - body *WSGClientDisconnectClient
func (s *WSGWirelessclientService) AddClientsDisconnect(ctx context.Context, body *WSGClientDisconnectClient) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddClientsDisconnect, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindApsOperationalClientTotalCountByApMac
//
// Use this API command to retrieve the total client count per AP.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGWirelessclientService) FindApsOperationalClientTotalCountByApMac(ctx context.Context, apMac string) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApsOperationalClientTotalCountByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindClientByQueryCriteria
//
// Query clients with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGWirelessclientService) FindClientByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGClientQueryList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGClientQueryList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindClientByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGClientQueryList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindHistoricalclientByQueryCriteria
//
// Use this API command to retrive historical client.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGWirelessclientService) FindHistoricalclientByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGClientHistoricalClientList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGClientHistoricalClientList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindHistoricalclientByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGClientHistoricalClientList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
