package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
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
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddRkszonesDhcpSiteDhcpSiteConfigDoAssignIpByZoneId
//
// Use this API command to get the DHCP/NAT service IP assignment when selecting with "Enable on Multiple APs". In the Manually Select AP mode (the manualSelect is true), the body should contain the selected APs (include the siteAps array). Otherwise, there is no need to include the selected APs in the Auto Select AP mode (see samples).
//
// Request Body:
//	 - body *WSGCommonDoAssignIp
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) AddRkszonesDhcpSiteDhcpSiteConfigDoAssignIpByZoneId(ctx context.Context, body *WSGCommonDoAssignIp, pZoneId string) (*WSGCommonDhcpSiteConfigListRef, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddRkszonesDual
//
// Use this API command to create a new Ruckus Wireless AP zone of IPv4/IPv6.
//
// Request Body:
//	 - body *WSGZoneCreateZone
func (s *WSGRuckusWirelessAPZoneService) AddRkszonesDual(ctx context.Context, body *WSGZoneCreateZone) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddRkszonesIpv6
//
// Use this API command to create a new Ruckus Wireless AP zone of IPv6.
//
// Request Body:
//	 - body *WSGZoneCreateZone
func (s *WSGRuckusWirelessAPZoneService) AddRkszonesIpv6(ctx context.Context, body *WSGZoneCreateZone) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesAltitudeById
//
// Use this API command to disable altitude configuration of zone.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesAltitudeById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesBackgroundScanning24ById
//
// Use this API command to disable background scanning 2.4GHz radio configuration for APs that belong to a zone.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesBackgroundScanning24ById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesBackgroundScanning50ById
//
// Use this API command to disable background scanning 5GHz radio configuration for APs that belong to a zone.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesBackgroundScanning50ById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesBandBalancingById
//
// Use this API command to disable band balancing for APs that belong to a zone.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesBandBalancingById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesById
//
// Use this API command to delete a zone.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesClientAdmissionControl24ById
//
// Use this API command to disable client admission control 2.4GHz radio configuration for APs that belong to a zone.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesClientAdmissionControl24ById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesClientAdmissionControl50ById
//
// Use this API command to disable client admission control 5GHz radio configuration for APs that belong to a zone.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesClientAdmissionControl50ById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesClientLoadBalancing24ById
//
// Use this API command to disable client load balancing 2.4GHz radio configuration for APs that belong to a zone.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesClientLoadBalancing24ById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesClientLoadBalancing50ById
//
// Use this API command to disable client load balancing 5GHz radio configuration for APs that belong to a zone.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesClientLoadBalancing50ById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesIpsecProfilesById
//
// Use this API command to Delete IPsec profiles.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesIpsecProfilesById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesLocationBasedServiceById
//
// Use this API command to disable location based service for APs that belong to a zone.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesLocationBasedServiceById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesMeshById
//
// Use this API command to disable mesh networking.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesMeshById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesRecoverySsidById
//
// Use this API command to clear recovery ssid setting of a zone.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesRecoverySsidById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesRogueById
//
// Use this API command to disable rogue AP detection for APs that belong to a zone.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesRogueById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesSmartMonitorById
//
// Use this API command to disable smart monitor for APs that belong to a zone.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesSmartMonitorById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesSnmpAgentById
//
// Use this API command to clear SNMPv2 and SNMPv3 agent that belong to a zone.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesSnmpAgentById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesSoftGreTunnelProfliesById
//
// Use this API command to Delete IPsec profiles.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesSoftGreTunnelProfliesById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesSyslogById
//
// Use this API command to disable syslog configuration for APs that belong to a zone.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesSyslogById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesVenueProfileById
//
// Use this API command to clear Hotspot 2.0 venue profile for APs that belong to a zone.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) DeleteRkszonesVenueProfileById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszones
//
// Use this API command to retrieve the list of Ruckus Wireless AP zones that belong to a domain.
//
// Query Parameters:
// - qDomainId string
//		- nullable
// - qIndex string
//		- nullable
// - qListSize string
//		- nullable
func (s *WSGRuckusWirelessAPZoneService) FindRkszones(ctx context.Context, qDomainId string, qIndex string, qListSize string) (*WSGZoneList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesApFirmwareByZoneId
//
// Use this API command to retrieve AP Firmware the list that belong to a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) FindRkszonesApFirmwareByZoneId(ctx context.Context, pZoneId string) (*WSGZoneApFirmwareList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesApmodelByModel
//
// Use this API command to retrieve AP model specific configuration that belong to a zone.
//
// Path Parameters:
// - pModel string
//		- required
// - pZoneId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) FindRkszonesApmodelByModel(ctx context.Context, pModel string, pZoneId string) (*WSGZoneAPModelApModel, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesApmodelCommonAttributeByModel
//
// Use this API command to retrieve AP model common attribute that belong to a zone.
//
// Path Parameters:
// - pModel string
//		- required
// - pZoneId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) FindRkszonesApmodelCommonAttributeByModel(ctx context.Context, pModel string, pZoneId string) (*WSGAPModelCommonAttribute, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesAvailableIpsecProfilesByZoneId
//
// Get available IPSec tunnel profiles of this Zone.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) FindRkszonesAvailableIpsecProfilesByZoneId(ctx context.Context, pZoneId string) (*WSGZoneAvailableTunnelProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesAvailableTunnelProfilesByZoneId
//
// Get available GRE tunnel profiles of this Zone.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) FindRkszonesAvailableTunnelProfilesByZoneId(ctx context.Context, pZoneId string) (*WSGZoneAvailableTunnelProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesById
//
// Use this API command to retrieve Ruckus Wireless AP zones configuration.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) FindRkszonesById(ctx context.Context, pId string) (*WSGZoneConfiguration, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesMeshById
//
// Use this API command to retrieve the mesh configuration of a zone.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) FindRkszonesMeshById(ctx context.Context, pId string) (*WSGZoneMeshConfiguration, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesServicesDhcpSiteConfigByQueryCriteria
//
// Use this API command to modify DHCP/NAT service configuration of Domain.
//
// Request Body:
//	 - body *WSGZoneQueryCriteria
func (s *WSGRuckusWirelessAPZoneService) FindRkszonesServicesDhcpSiteConfigByQueryCriteria(ctx context.Context, body *WSGZoneQueryCriteria) (*WSGZoneDhcpSiteConfigList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// PartialUpdateRkszonesById
//
// Use this API command to modify the basic information of a zone.
//
// Request Body:
//	 - body *WSGZoneModifyZone
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) PartialUpdateRkszonesById(ctx context.Context, body *WSGZoneModifyZone, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateRkszonesApFirmwareByZoneId
//
// Use this API command to change the AP Firmware that belong to a zone.
//
// Request Body:
//	 - body *WSGZoneModfiyApFirmware
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) UpdateRkszonesApFirmwareByZoneId(ctx context.Context, body *WSGZoneModfiyApFirmware, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateRkszonesApmodelByModel
//
// Use this API command to modify the AP model specific configuration that belong to a zone.
//
// Request Body:
//	 - body *WSGZoneAPModelApModel
//
// Path Parameters:
// - pModel string
//		- required
// - pZoneId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) UpdateRkszonesApmodelByModel(ctx context.Context, body *WSGZoneAPModelApModel, pModel string, pZoneId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateRkszonesById
//
// Use this API command to modify entire information of a zone.
//
// Request Body:
//	 - body *WSGZoneModifyZone
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGRuckusWirelessAPZoneService) UpdateRkszonesById(ctx context.Context, body *WSGZoneModifyZone, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
