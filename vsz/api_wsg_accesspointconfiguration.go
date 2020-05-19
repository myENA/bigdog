package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGAccessPointConfigurationService struct {
	apiClient *APIClient
}

func NewWSGAccessPointConfigurationService(c *APIClient) *WSGAccessPointConfigurationService {
	s := new(WSGAccessPointConfigurationService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAccessPointConfigurationService() *WSGAccessPointConfigurationService {
	return NewWSGAccessPointConfigurationService(ss.apiClient)
}

// AddAps
//
// Use this API command to create a new access point.
//
// Request Body:
//	 - body *WSGAPCreateAP
func (s *WSGAccessPointConfigurationService) AddAps(ctx context.Context, body *WSGAPCreateAP) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddAps, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusCreated, httpResp, nil, err)
	return rm, err
}

// AddApsPictureByApMac
//
// Use this API command to upload a new AP picture.
//
// Request Body:
//	 - body []byte
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) AddApsPictureByApMac(ctx context.Context, body []byte, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddApsPictureByApMac, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// AddSwapAps
//
// Use this API command to swap in specific AP
//
// Request Body:
//	 - body *WSGAPSwapApConfigure
func (s *WSGAccessPointConfigurationService) AddSwapAps(ctx context.Context, body *WSGAPSwapApConfigure) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddSwapAps, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsAltitudeByApMac
//
// Use this API command to disable AP level override of altitude. The access point will take its group's configuration or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsAltitudeByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsAltitudeByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsApMgmtVlanByApMac
//
// Disable AP Management Vlan Override of an AP.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsApMgmtVlanByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsApMgmtVlanByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsAutoChannelSelection24ByApMac
//
// Use this API command to disable the AP level override of auto channel selection on the 2.4GHz radio. The access point will take its group's or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsAutoChannelSelection24ByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsAutoChannelSelection24ByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsAutoChannelSelection50ByApMac
//
// Use this API command to disable the AP level override of auto channel selection on the 5GHz radio. The access point will take its group's or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsAutoChannelSelection50ByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsAutoChannelSelection50ByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsBonjourGatewayByApMac
//
// Use this API command to disable AP level override of bonjour gateway. The access point will take its group's configuration or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsBonjourGatewayByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsBonjourGatewayByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsByApMac
//
// Use this API command to delete an access point.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsChannelEvaluationIntervalByApMac
//
// Disable AP lChannel Evaluation Interval. The access point will take its group's configuration or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsChannelEvaluationIntervalByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsChannelEvaluationIntervalByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsClientAdmissionControl24ByApMac
//
// Use this API command to disable AP level override of client admission control 2.4GHz radio configuration. The access point will take its group's configuration or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsClientAdmissionControl24ByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsClientAdmissionControl24ByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsClientAdmissionControl50ByApMac
//
// Use this API command to disable AP level override of client admission control 5GHz radio configuration. The access point will take its group's configuration or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsClientAdmissionControl50ByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsClientAdmissionControl50ByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsDirectedMulticastFromNetworkEnabledByApMac
//
// Use this API command to disable Directed Multicast from network to wired/wireless client configuration override.The AP will apply its group's or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsDirectedMulticastFromNetworkEnabledByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsDirectedMulticastFromNetworkEnabledByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsDirectedMulticastFromWiredClientEnabledByApMac
//
// Use this API command to disable Island SSID Broadcast enabled configuration override.The AP will apply its group's or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsDirectedMulticastFromWiredClientEnabledByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsDirectedMulticastFromWiredClientEnabledByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsDirectedMulticastFromWirelessClientEnabledByApMac
//
// Use this API command to disable Island SSID Broadcast enabled configuration override.The AP will apply its group's or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsDirectedMulticastFromWirelessClientEnabledByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsDirectedMulticastFromWirelessClientEnabledByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsGpsCoordinatesByApMac
//
// Disable AP Management GPS Cooordinates of an AP.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsGpsCoordinatesByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsGpsCoordinatesByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsLocationAdditionalInfoByApMac
//
// Use this API command to disable AP level override of location additionalInfo. The access point will take its group's configuration or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsLocationAdditionalInfoByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsLocationAdditionalInfoByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsLocationByApMac
//
// Use this API command to disable AP level override of location. The access point will take its group's configuration or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsLocationByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsLocationByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsLoginByApMac
//
// Use this API command to disable the AP-level logon override. The AP will apply its group's or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsLoginByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsLoginByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsLteBandLockChannelsByApMac
//
// Use this API command to disable LTE band lock channel override. The AP will apply its group's or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsLteBandLockChannelsByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsLteBandLockChannelsByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsMeshOptionsByApMac
//
// Use this API command to disable mesh options.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsMeshOptionsByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsMeshOptionsByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsPictureByApMac
//
// Use this API command to delete an AP picture.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsPictureByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsPictureByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsProtectionMode24ByApMac
//
// Use this API command to disable 2.4GHz radio protection mode configuration override.The AP will apply its group's or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsProtectionMode24ByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsProtectionMode24ByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsRecoverySsidByApMac
//
// Use this API command to disable Recovery SSID configuration override.The AP will apply its group's or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsRecoverySsidByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsRecoverySsidByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsRksGreForwardBroadcastByApMac
//
// Use this API command to disable Ruckus GRE Broadcast packet forwarding override. The AP will apply its group's or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsRksGreForwardBroadcastByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsRksGreForwardBroadcastByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsRogueApAggressivenessModeByApMac
//
// Use this API command to disable rogue AP aggressiveness mode override. The AP will apply its group's or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsRogueApAggressivenessModeByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsRogueApAggressivenessModeByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsRogueApJammingThresholdByApMac
//
// Use this API command to disable rogue AP jamming threshold override. The AP will apply its group's or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsRogueApJammingThresholdByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsRogueApJammingThresholdByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsRogueApReportThresholdByApMac
//
// Use this API command to disable rogue AP report threshold override. The AP will apply its group's or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsRogueApReportThresholdByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsRogueApReportThresholdByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsSmartMonitorByApMac
//
// Use this API command to disable AP level override of smart monitor. The access point will take its group's configuration or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsSmartMonitorByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsSmartMonitorByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsSpecificByApMac
//
// Use this API command to disable specific configuration override from AP group or zone.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsSpecificByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsSpecificByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsSyslogByApMac
//
// Use this API command to disable the AP level syslog override. The access point will take its group's or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsSyslogByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsSyslogByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsVenueProfileByApMac
//
// Use this API command to disable AP level override of venue profile. The access point will take its group's configuration or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsVenueProfileByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsVenueProfileByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsWifi24ByApMac
//
// Use this API command to disable the AP level override of the 2.4GHz radio configuration. The access point will take its group's configuration or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsWifi24ByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsWifi24ByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsWifi24ChannelByApMac
//
// Use this API command to disable the AP level override of the 2.4GHz radio channel. The access point will take its group's configuration or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsWifi24ChannelByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsWifi24ChannelByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsWifi24ChannelRangeByApMac
//
// Use this API command to disable the AP level override of the 2.4GHz radio channelRange. The access point will take its group's configuration or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsWifi24ChannelRangeByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsWifi24ChannelRangeByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsWifi24ChannelWidthByApMac
//
// Use this API command to disable the AP level override of the 2.4GHz radio channelWidth. The access point will take its group's configuration or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsWifi24ChannelWidthByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsWifi24ChannelWidthByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsWifi24TxPowerByApMac
//
// Use this API command to disable the AP level override of the 2.4GHz radio txPower. The access point will take its group's configuration or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsWifi24TxPowerByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsWifi24TxPowerByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsWifi50ByApMac
//
// Use this API command to disable the AP level override of 5GHz radio configuration. The access point will take its group's or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsWifi50ByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsWifi50ByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsWifi50ChannelByApMac
//
// Use this API command to disable the AP level override of 5GHz radio channel. The access point will take its group's or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsWifi50ChannelByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsWifi50ChannelByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsWifi50ChannelRangeByApMac
//
// Use this API command to disable the AP level override of 5GHz radio channelRange. The access point will take its group's or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsWifi50ChannelRangeByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsWifi50ChannelRangeByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsWifi50ChannelWidthByApMac
//
// Use this API command to disable the AP level override of 5GHz radio channelWidth. The access point will take its group's or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsWifi50ChannelWidthByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsWifi50ChannelWidthByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsWifi50TxPowerByApMac
//
// Use this API command to disable the AP level override of 5GHz radio txPower. The access point will take its group's or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsWifi50TxPowerByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsWifi50TxPowerByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsWlanGroup24ByApMac
//
// Use this API command to disable the AP level override of WLAN group configuration on 2.4GHz radio. The access point will take its group's or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsWlanGroup24ByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsWlanGroup24ByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteApsWlanGroup50ByApMac
//
// Use this API command to disable the AP level override of WLAN group on the 5GHz radio. The access point will take its group's or zone's configuration.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) DeleteApsWlanGroup50ByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteApsWlanGroup50ByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindAps
//
// Use this API command to retrieve the list of APs that belong to a zone or a domain.
//
// Optional Parameters:
// - domainId string
//		- nullable
// - index string
//		- nullable
// - listSize string
//		- nullable
// - zoneId string
//		- nullable
func (s *WSGAccessPointConfigurationService) FindAps(ctx context.Context, optionalParams map[string][]string) (*WSGAPListEntry, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAPListEntry
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindAps, true)
	if v, ok := optionalParams["domainId"]; ok && len(v) > 0 {
		req.SetQueryParameter("domainId", v)
	}
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.SetQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.SetQueryParameter("listSize", v)
	}
	if v, ok := optionalParams["zoneId"]; ok && len(v) > 0 {
		req.SetQueryParameter("zoneId", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAPListEntry()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindApsByApMac
//
// Use this API command to retrieve the configuration of an AP.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) FindApsByApMac(ctx context.Context, apMac string) (*WSGAPConfiguration, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAPConfiguration
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApsByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAPConfiguration()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindApsPictureByApMac
//
// Use this API command to retrieve the current AP picture.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) FindApsPictureByApMac(ctx context.Context, apMac string) ([]byte, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     []byte
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApsPictureByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = make([]byte, 0)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindApsSupportLogByApMac
//
// Use this API command to download AP support log.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) FindApsSupportLogByApMac(ctx context.Context, apMac string) ([]byte, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     []byte
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApsSupportLogByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = make([]byte, 0)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindMeshZeroTouch
//
// Use this API command to retrieve a list of unapproved AP.
func (s *WSGAccessPointConfigurationService) FindMeshZeroTouch(ctx context.Context) (*WSGMeshNodeInfoList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGMeshNodeInfoList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindMeshZeroTouch, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGMeshNodeInfoList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateApsByApMac
//
// Use this API command to modify the configuration of an AP.
//
// Request Body:
//	 - body *WSGAPModifyAP
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) PartialUpdateApsByApMac(ctx context.Context, body *WSGAPModifyAP, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateApsByApMac, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// UpdateApsByApMac
//
// Use this API command to modify the entire information of an AP.
//
// Request Body:
//	 - body *WSGAPModifyAP
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) UpdateApsByApMac(ctx context.Context, body *WSGAPModifyAP, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateApsByApMac, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// UpdateApsRebootByApMac
//
// reboot an access point.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) UpdateApsRebootByApMac(ctx context.Context, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateApsRebootByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// UpdateApsSpecificByApMac
//
// Use this API command to modify specific configuration.
//
// Request Body:
//	 - body *WSGAPModel
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointConfigurationService) UpdateApsSpecificByApMac(ctx context.Context, body *WSGAPModel, apMac string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateApsSpecificByApMac, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// UpdateMeshZeroTouch
//
// Use this API command to approve/reject unapproved AP. Recommend to deploy 20 island APs to join per batch at the same time.
//
// Request Body:
//	 - body *WSGMeshNodeInfoUpdateAPZeroTouch
func (s *WSGAccessPointConfigurationService) UpdateMeshZeroTouch(ctx context.Context, body *WSGMeshNodeInfoUpdateAPZeroTouch) (interface{}, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateMeshZeroTouch, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}