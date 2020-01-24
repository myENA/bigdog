package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGFlexiVPNService struct {
	apiClient *APIClient
}

func NewWSGFlexiVPNService(c *APIClient) *WSGFlexiVPNService {
	s := new(WSGFlexiVPNService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGFlexiVPNService() *WSGFlexiVPNService {
	return NewWSGFlexiVPNService(ss.apiClient)
}

type WSGFlexiVPNSetting struct {
	// ZoneAffinityId
	// Zone Affinity ID
	// Constraints:
	//    - required
	ZoneAffinityId *string `json:"zoneAffinityId" validate:"required"`
}

func NewWSGFlexiVPNSetting() *WSGFlexiVPNSetting {
	m := new(WSGFlexiVPNSetting)
	return m
}

// DeleteRkszonesWlansFlexiVpnProfileById
//
// Use this API command to delete Flexi-VPN on WLAN
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGFlexiVPNService) DeleteRkszonesWlansFlexiVpnProfileById(ctx context.Context, id string, zoneId string) error {
	var (
		req *APIRequest
		err error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesWlansFlexiVpnProfileById, true)
}

// FindServicesFlexiVpnProfileByQueryCriteria
//
// Use this API command to query Flexi-VPN profiles.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGFlexiVPNService) FindServicesFlexiVpnProfileByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileFlexiVpnProfileList, error) {
	var (
		req  *APIRequest
		resp *WSGProfileFlexiVpnProfileList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesFlexiVpnProfileByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}
