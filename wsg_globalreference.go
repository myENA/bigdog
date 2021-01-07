package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGGlobalReferenceService struct {
	apiClient *VSZClient
}

func NewWSGGlobalReferenceService(c *VSZClient) *WSGGlobalReferenceService {
	s := new(WSGGlobalReferenceService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGGlobalReferenceService() *WSGGlobalReferenceService {
	return NewWSGGlobalReferenceService(ss.apiClient)
}

// FindGlobalSettingsFriendlyNameLang
//
// Operation ID: findGlobalSettingsFriendlyNameLang
//
// Use this API command to get friendly name of usable language for profile: Hotspot2.0 Identity Provider.
func (s *WSGGlobalReferenceService) FindGlobalSettingsFriendlyNameLang(ctx context.Context, mutators ...RequestMutator) (*WSGSystemFriendlyNameLangListAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSystemFriendlyNameLangList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindGlobalSettingsFriendlyNameLang, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGSystemFriendlyNameLangList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}

// FindGlobalSettingsPortalLang
//
// Operation ID: findGlobalSettingsPortalLang
//
// Use this API command to get friendly name of usable language for profile: Guest Access (Language in General Options), Web Auth (Language in General Options).
func (s *WSGGlobalReferenceService) FindGlobalSettingsPortalLang(ctx context.Context, mutators ...RequestMutator) (*WSGSystemPortalLangListAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSystemPortalLangList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindGlobalSettingsPortalLang, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGSystemPortalLangList()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, s.apiClient.autoHydrate, err)
	return resp, rm, err
}
