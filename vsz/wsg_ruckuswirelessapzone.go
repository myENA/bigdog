package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGRuckusWirelessAPZoneService struct {
	apiClient *APIClient
}

func NewWSGRuckusWirelessAPZoneService(c *APIClient) *WSGRuckusWirelessAPZoneService {
	s := new(WSGRuckusWirelessAPZoneService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGRuckusWirelessAPZoneService() *WSGRuckusWirelessAPZoneService {
	return NewWSGRuckusWirelessAPZoneService(ss.apiClient)
}

// AddRkszones
//
// Use this API command to create a new Ruckus Wireless AP zone.
//
// Request Body:
//	 - body *WSGZoneCreateZone
func (s *WSGRuckusWirelessAPZoneService) AddRkszones(ctx context.Context, body *WSGZoneCreateZone) (*WSGCommonCreateResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonCreateResult
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszones, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// AddRkszonesDhcpSiteDhcpSiteConfigDoAssignIpByZoneId
//
// Use this API command to get the DHCP/NAT service IP assignment when selecting with "Enable on Multiple APs". In the Manually Select AP mode (the manualSelect is true), the body should contain the selected APs (include the siteAps array). Otherwise, there is no need to include the selected APs in the Auto Select AP mode (see samples).
//
// Request Body:
//	 - body *WSGCommonDoAssignIp
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) AddRkszonesDhcpSiteDhcpSiteConfigDoAssignIpByZoneId(ctx context.Context, body *WSGCommonDoAssignIp, zoneId string) (*WSGCommonDhcpSiteConfigListRef, error) {
	var (
		req  *APIRequest
		resp *WSGCommonDhcpSiteConfigListRef
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
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesDhcpSiteDhcpSiteConfigDoAssignIpByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// AddRkszonesDual
//
// Use this API command to create a new Ruckus Wireless AP zone of IPv4/IPv6.
//
// Request Body:
//	 - body *WSGZoneCreateZone
func (s *WSGRuckusWirelessAPZoneService) AddRkszonesDual(ctx context.Context, body *WSGZoneCreateZone) (*WSGCommonCreateResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonCreateResult
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesDual, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// AddRkszonesIpv6
//
// Use this API command to create a new Ruckus Wireless AP zone of IPv6.
//
// Request Body:
//	 - body *WSGZoneCreateZone
func (s *WSGRuckusWirelessAPZoneService) AddRkszonesIpv6(ctx context.Context, body *WSGZoneCreateZone) (*WSGCommonCreateResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonCreateResult
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesIpv6, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteRkszonesAltitudeById
//
// Use this API command to disable altitude configuration of zone.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesAltitudeById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesAltitudeById, true)
}

// DeleteRkszonesBackgroundScanning24ById
//
// Use this API command to disable background scanning 2.4GHz radio configuration for APs that belong to a zone.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesBackgroundScanning24ById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesBackgroundScanning24ById, true)
}

// DeleteRkszonesBackgroundScanning50ById
//
// Use this API command to disable background scanning 5GHz radio configuration for APs that belong to a zone.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesBackgroundScanning50ById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesBackgroundScanning50ById, true)
}

// DeleteRkszonesBandBalancingById
//
// Use this API command to disable band balancing for APs that belong to a zone.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesBandBalancingById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesBandBalancingById, true)
}

// DeleteRkszonesById
//
// Use this API command to delete a zone.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesById, true)
}

// DeleteRkszonesClientAdmissionControl24ById
//
// Use this API command to disable client admission control 2.4GHz radio configuration for APs that belong to a zone.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesClientAdmissionControl24ById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesClientAdmissionControl24ById, true)
}

// DeleteRkszonesClientAdmissionControl50ById
//
// Use this API command to disable client admission control 5GHz radio configuration for APs that belong to a zone.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesClientAdmissionControl50ById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesClientAdmissionControl50ById, true)
}

// DeleteRkszonesClientLoadBalancing24ById
//
// Use this API command to disable client load balancing 2.4GHz radio configuration for APs that belong to a zone.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesClientLoadBalancing24ById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesClientLoadBalancing24ById, true)
}

// DeleteRkszonesClientLoadBalancing50ById
//
// Use this API command to disable client load balancing 5GHz radio configuration for APs that belong to a zone.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesClientLoadBalancing50ById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesClientLoadBalancing50ById, true)
}

// DeleteRkszonesIpsecProfilesById
//
// Use this API command to Delete IPsec profiles.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesIpsecProfilesById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesIpsecProfilesById, true)
}

// DeleteRkszonesLocationBasedServiceById
//
// Use this API command to disable location based service for APs that belong to a zone.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesLocationBasedServiceById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesLocationBasedServiceById, true)
}

// DeleteRkszonesMeshById
//
// Use this API command to disable mesh networking.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesMeshById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesMeshById, true)
}

// DeleteRkszonesRecoverySsidById
//
// Use this API command to clear recovery ssid setting of a zone.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesRecoverySsidById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesRecoverySsidById, true)
}

// DeleteRkszonesRogueById
//
// Use this API command to disable rogue AP detection for APs that belong to a zone.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesRogueById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesRogueById, true)
}

// DeleteRkszonesSmartMonitorById
//
// Use this API command to disable smart monitor for APs that belong to a zone.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesSmartMonitorById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesSmartMonitorById, true)
}

// DeleteRkszonesSnmpAgentById
//
// Use this API command to clear SNMPv2 and SNMPv3 agent that belong to a zone.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesSnmpAgentById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesSnmpAgentById, true)
}

// DeleteRkszonesSoftGreTunnelProfliesById
//
// Use this API command to Delete IPsec profiles.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesSoftGreTunnelProfliesById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesSoftGreTunnelProfliesById, true)
}

// DeleteRkszonesSyslogById
//
// Use this API command to disable syslog configuration for APs that belong to a zone.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesSyslogById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesSyslogById, true)
}

// DeleteRkszonesVenueProfileById
//
// Use this API command to clear Hotspot 2.0 venue profile for APs that belong to a zone.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesVenueProfileById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesVenueProfileById, true)
}

// FindRkszones
//
// Use this API command to retrieve the list of Ruckus Wireless AP zones that belong to a domain.
//
// Optional Parameters:
// - domainId string
//		- nullable
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGRuckusWirelessAPZoneService) FindRkszones(ctx context.Context, optionalParams map[string]interface{}) (*WSGZoneList, error) {
	var (
		req  *APIRequest
		resp *WSGZoneList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszones, true)
}

// FindRkszonesApFirmwareByZoneId
//
// Use this API command to retrieve AP Firmware the list that belong to a zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) FindRkszonesApFirmwareByZoneId(ctx context.Context, zoneId string) (*WSGZoneApFirmwareList, error) {
	var (
		req  *APIRequest
		resp *WSGZoneApFirmwareList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesApFirmwareByZoneId, true)
}

// FindRkszonesApmodelByModel
//
// Use this API command to retrieve AP model specific configuration that belong to a zone.
//
// Required Parameters:
// - model string
//		- required
// - zoneId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) FindRkszonesApmodelByModel(ctx context.Context, model string, zoneId string) (*WSGZoneAPModelApModel, error) {
	var (
		req  *APIRequest
		resp *WSGZoneAPModelApModel
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, model, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesApmodelByModel, true)
}

// FindRkszonesApmodelCommonAttributeByModel
//
// Use this API command to retrieve AP model common attribute that belong to a zone.
//
// Required Parameters:
// - model string
//		- required
// - zoneId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) FindRkszonesApmodelCommonAttributeByModel(ctx context.Context, model string, zoneId string) (*WSGAPModelCommonAttribute, error) {
	var (
		req  *APIRequest
		resp *WSGAPModelCommonAttribute
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, model, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesApmodelCommonAttributeByModel, true)
}

// FindRkszonesAvailableIpsecProfilesByZoneId
//
// Get available IPSec tunnel profiles of this Zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) FindRkszonesAvailableIpsecProfilesByZoneId(ctx context.Context, zoneId string) (*WSGZoneAvailableTunnelProfileList, error) {
	var (
		req  *APIRequest
		resp *WSGZoneAvailableTunnelProfileList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesAvailableIpsecProfilesByZoneId, true)
}

// FindRkszonesAvailableTunnelProfilesByZoneId
//
// Get available GRE tunnel profiles of this Zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) FindRkszonesAvailableTunnelProfilesByZoneId(ctx context.Context, zoneId string) (*WSGZoneAvailableTunnelProfileList, error) {
	var (
		req  *APIRequest
		resp *WSGZoneAvailableTunnelProfileList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesAvailableTunnelProfilesByZoneId, true)
}

// FindRkszonesById
//
// Use this API command to retrieve Ruckus Wireless AP zones configuration.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) FindRkszonesById(ctx context.Context, id string) (*WSGZoneConfiguration, error) {
	var (
		req  *APIRequest
		resp *WSGZoneConfiguration
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesById, true)
}

// FindRkszonesMeshById
//
// Use this API command to retrieve the mesh configuration of a zone.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) FindRkszonesMeshById(ctx context.Context, id string) (*WSGZoneMeshConfiguration, error) {
	var (
		req  *APIRequest
		resp *WSGZoneMeshConfiguration
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesMeshById, true)
}

// FindRkszonesServicesDhcpSiteConfigByQueryCriteria
//
// Use this API command to modify DHCP/NAT service configuration of Domain.
//
// Request Body:
//	 - body *WSGZoneQueryCriteria
func (s *WSGRuckusWirelessAPZoneService) FindRkszonesServicesDhcpSiteConfigByQueryCriteria(ctx context.Context, body *WSGZoneQueryCriteria) (*WSGZoneDhcpSiteConfigList, error) {
	var (
		req  *APIRequest
		resp *WSGZoneDhcpSiteConfigList
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindRkszonesServicesDhcpSiteConfigByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// PartialUpdateRkszonesById
//
// Use this API command to modify the basic information of a zone.
//
// Request Body:
//	 - body *WSGZoneModifyZone
//
// Required Parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) PartialUpdateRkszonesById(ctx context.Context, body *WSGZoneModifyZone, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
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
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// UpdateRkszonesApFirmwareByZoneId
//
// Use this API command to change the AP Firmware that belong to a zone.
//
// Request Body:
//	 - body *WSGZoneModfiyApFirmware
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) UpdateRkszonesApFirmwareByZoneId(ctx context.Context, body *WSGZoneModfiyApFirmware, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
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
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateRkszonesApFirmwareByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// UpdateRkszonesApmodelByModel
//
// Use this API command to modify the AP model specific configuration that belong to a zone.
//
// Request Body:
//	 - body *WSGZoneAPModelApModel
//
// Required Parameters:
// - model string
//		- required
// - zoneId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) UpdateRkszonesApmodelByModel(ctx context.Context, body *WSGZoneAPModelApModel, model string, zoneId string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
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
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, model, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateRkszonesApmodelByModel, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// UpdateRkszonesById
//
// Use this API command to modify entire information of a zone.
//
// Request Body:
//	 - body *WSGZoneModifyZone
//
// Required Parameters:
// - id string
//		- required
func (s *WSGRuckusWirelessAPZoneService) UpdateRkszonesById(ctx context.Context, body *WSGZoneModifyZone, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
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
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateRkszonesById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}
