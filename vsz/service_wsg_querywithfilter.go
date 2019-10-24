package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/aaaserverquery"
	"github.com/myENA/ruckus-client/vsz/types/wsg/apquery"
	"github.com/myENA/ruckus-client/vsz/types/wsg/clientquery"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/dpsk"
	"github.com/myENA/ruckus-client/vsz/types/wsg/indoormap"
	"github.com/myENA/ruckus-client/vsz/types/wsg/meshneighborinfo"
	"github.com/myENA/ruckus-client/vsz/types/wsg/meshnodeinfo"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
	"github.com/myENA/ruckus-client/vsz/types/wsg/rogueinfo"
	"github.com/myENA/ruckus-client/vsz/types/wsg/wlanquery"
)

type WSGQueryWithFilterService struct {
	client *Client
}

func NewWSGQueryWithFilterService(client *Client) *WSGQueryWithFilterService {
	s := new(WSGQueryWithFilterService)
	s.client = client
	return s
}

func (ss *WSGService) WSGQueryWithFilterService() *WSGQueryWithFilterService {
	serv := new(WSGQueryWithFilterService)
	serv.client = ss.client
	return serv
}

// FindApByQueryCriteria
//
// Query APs with specified filters
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindApByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*apquery.ApQueryList, error) {
}

// FindApWlanByQueryCriteria
//
// Use this API command to retrieve AP Wlan list with BSSID information by QueryCriteria
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindApWlanByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*wlanquery.ApWlanBssidQueryList, error) {
}

// FindClientByQueryCriteria
//
// Query clients with specified filters.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindClientByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*clientquery.ClientQueryList, error) {
}

// FindDpskByQueryCriteria
//
// Query DPSKs with specified filters.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindDpskByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*dpsk.DpskQueryList, error) {
}

// FindIndoorMapByQueryCriteria
//
// Query indoorMap with specified filters.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindIndoorMapByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*indoormap.IndoorMapSummaryList, error) {
}

// FindMeshNeighborByApMacByQueryCriteria
//
// Use this API command to retrieve a list of neighbor access points on mesh AP.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindMeshNeighborByApMacByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet, pApMac string) (*meshneighborinfo.MeshNeighborInfoList, error) {
}

// FindMeshTopologyByApMacByQueryCriteria
//
// Use this API command to retrieve a list of topology on mesh AP.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindMeshTopologyByApMacByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet, pApMac string) (meshnodeinfo.MeshNodeInfoArray, error) {
}

// FindMeshTopologyByQueryCriteria
//
// Use this API command to retrieve a list of topology on zone.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindMeshTopologyByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*meshnodeinfo.MeshNodeInfoList, error) {
}

// FindRoguesInfoListByQueryCriteria
//
// Use this API command to retrieve a list of rogue access points.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindRoguesInfoListByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*rogueinfo.RogueInfoList, error) {
}

// FindServicesAaaServerAcctByQueryCriteria
//
// Query Accounting AAAServers with specified filters.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindServicesAaaServerAcctByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*aaaserverquery.AaaServerQueryList, error) {
}

// FindServicesAaaServerAuthByQueryCriteria
//
// Query Authentication AAAServers with specified filters.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindServicesAaaServerAuthByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*aaaserverquery.AaaServerQueryList, error) {
}

// FindServicesAaaServerByQueryCriteria
//
// Query AAAServers with specified filters.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindServicesAaaServerByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*aaaserverquery.AaaServerQueryList, error) {
}

// FindServicesBonjourPolicyByQueryCriteria
//
// Query bonjourPolicy Profiles with specified filters.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindServicesBonjourPolicyByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) error {
}

// FindServicesDevicePolicyByQueryCriteria
//
// Query Device Policy Profiles with specified filters.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindServicesDevicePolicyByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) error {
}

// FindServicesDhcpProfileByQueryCriteria
//
// Query DHCP Profiles with specified filters.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindServicesDhcpProfileByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.DhcpProfileList, error) {
}

// FindServicesDscpProfileByQueryCriteria
//
// Query DSCP Profiles with specified filters.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindServicesDscpProfileByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) error {
}

// FindServicesEthernetPortProfileByQueryCriteria
//
// Query Ethernet Port Profiles with specified filters.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindServicesEthernetPortProfileByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) error {
}

// FindServicesGuessAccessByQueryCriteria
//
// Query Guess Access Profiles with specified filters.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindServicesGuessAccessByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) error {
}

// FindServicesGuestAccessByQueryCriteria
//
// Query Guest Access Profiles with specified filters.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindServicesGuestAccessByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) error {
}

// FindServicesHotspot20ProfileByQueryCriteria
//
// Query Hotspot20 Profiles with specified filters.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindServicesHotspot20ProfileByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) error {
}

// FindServicesHotspotByQueryCriteria
//
// Query Hotspot Profiles with specified filters.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindServicesHotspotByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) error {
}

// FindServicesL2AccessControlByQueryCriteria
//
// Query L2 AccessControl Profiles with specified filters.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindServicesL2AccessControlByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) error {
}

// FindServicesVenueProfileByQueryCriteria
//
// Query Venue Profiles with specified filters.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindServicesVenueProfileByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) error {
}

// FindServicesVlanPoolingByQueryCriteria
//
// Query Vlan Pooling Profiles with specified filters.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindServicesVlanPoolingByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) error {
}

// FindServicesWebAuthenticationByQueryCriteria
//
// Query Web Authentications with specified filters.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindServicesWebAuthenticationByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) error {
}

// FindServicesWechatProfileByQueryCriteria
//
// Query Wechat Portals with specified filters.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindServicesWechatProfileByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) error {
}

// FindServicesWlanSchedulerByQueryCriteria
//
// Query Wlan Schedulers with specified filters.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindServicesWlanSchedulerByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) error {
}

// FindUserByQueryCriteria
//
// Query users with specified filters.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindUserByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) error {
}

// FindWlanByQueryCriteria
//
// Query WLANs with specified filters.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGQueryWithFilterService) FindWlanByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*wlanquery.WlanQueryList, error) {
}
