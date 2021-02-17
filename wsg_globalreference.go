package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
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
// Use this API command to get friendly name of usable language for profile: Hotspot2.0 Identity Provider.
//
// Operation ID: findGlobalSettingsFriendlyNameLang
// Operation path: /globalSettings/friendlyNameLang
// Success code: 200 (OK)
func (s *WSGGlobalReferenceService) FindGlobalSettingsFriendlyNameLang(ctx context.Context, mutators ...RequestMutator) (*WSGSystemFriendlyNameLangListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGSystemFriendlyNameLangListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindGlobalSettingsFriendlyNameLang, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGSystemFriendlyNameLangListAPIResponse), err
}

// FindGlobalSettingsPortalLang
//
// Use this API command to get friendly name of usable language for profile: Guest Access (Language in General Options), Web Auth (Language in General Options).
//
// Operation ID: findGlobalSettingsPortalLang
// Operation path: /globalSettings/portalLang
// Success code: 200 (OK)
func (s *WSGGlobalReferenceService) FindGlobalSettingsPortalLang(ctx context.Context, mutators ...RequestMutator) (*WSGSystemPortalLangListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGSystemPortalLangListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindGlobalSettingsPortalLang, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGSystemPortalLangListAPIResponse), err
}
