package vsz

// API Version: v9_0

import (
	"context"
	"fmt"
	"net/http"
)

type WSGGlobalreferenceService struct {
	apiClient *APIClient
}

func NewWSGGlobalreferenceService(c *APIClient) *WSGGlobalreferenceService {
	s := new(WSGGlobalreferenceService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGGlobalreferenceService() *WSGGlobalreferenceService {
	return NewWSGGlobalreferenceService(ss.apiClient)
}

// FindGlobalSettingsFriendlyNameLang
//
// Use this API command to get friendly name of usable language for profile: Hotspot2.0 Identity Provider.
func (s *WSGGlobalreferenceService) FindGlobalSettingsFriendlyNameLang(ctx context.Context) (*WSGSystemFriendlyNameLangList, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGFindGlobalSettingsFriendlyNameLang), true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemFriendlyNameLangList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindGlobalSettingsPortalLang
//
// Use this API command to get friendly name of usable language for profile: Guest Access (Language in General Options), Web Auth (Language in General Options).
func (s *WSGGlobalreferenceService) FindGlobalSettingsPortalLang(ctx context.Context) (*WSGSystemPortalLangList, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodGet, fmt.Sprintf("/%s%s", s.apiClient.wsgPath, RouteWSGFindGlobalSettingsPortalLang), true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemPortalLangList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}
