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

func (s *WSGQueryWithFilterService) FindApByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*apquery.ApQueryList, error) {
}

func (s *WSGQueryWithFilterService) FindApWlanByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*wlanquery.ApWlanBssidQueryList, error) {
}

func (s *WSGQueryWithFilterService) FindClientByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*clientquery.ClientQueryList, error) {
}

func (s *WSGQueryWithFilterService) FindDpskByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*dpsk.DpskQueryList, error) {
}

func (s *WSGQueryWithFilterService) FindIndoorMapByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*indoormap.IndoorMapSummaryList, error) {
}

func (s *WSGQueryWithFilterService) FindMeshNeighborByApMacByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet, pApMac string) (*meshneighborinfo.MeshNeighborInfoList, error) {
}

func (s *WSGQueryWithFilterService) FindMeshTopologyByApMacByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet, pApMac string) (meshnodeinfo.MeshNodeInfoArray, error) {
}

func (s *WSGQueryWithFilterService) FindMeshTopologyByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*meshnodeinfo.MeshNodeInfoList, error) {
}

func (s *WSGQueryWithFilterService) FindRoguesInfoListByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*rogueinfo.RogueInfoList, error) {
}

func (s *WSGQueryWithFilterService) FindServicesAaaServerAcctByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*aaaserverquery.AaaServerQueryList, error) {
}

func (s *WSGQueryWithFilterService) FindServicesAaaServerAuthByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*aaaserverquery.AaaServerQueryList, error) {
}

func (s *WSGQueryWithFilterService) FindServicesAaaServerByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*aaaserverquery.AaaServerQueryList, error) {
}

func (s *WSGQueryWithFilterService) FindServicesBonjourPolicyByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) error {
}

func (s *WSGQueryWithFilterService) FindServicesDevicePolicyByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) error {
}

func (s *WSGQueryWithFilterService) FindServicesDhcpProfileByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.DhcpProfileList, error) {
}

func (s *WSGQueryWithFilterService) FindServicesDscpProfileByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) error {
}

func (s *WSGQueryWithFilterService) FindServicesEthernetPortProfileByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) error {
}

func (s *WSGQueryWithFilterService) FindServicesGuessAccessByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) error {
}

func (s *WSGQueryWithFilterService) FindServicesGuestAccessByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) error {
}

func (s *WSGQueryWithFilterService) FindServicesHotspot20ProfileByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) error {
}

func (s *WSGQueryWithFilterService) FindServicesHotspotByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) error {
}

func (s *WSGQueryWithFilterService) FindServicesL2AccessControlByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) error {
}

func (s *WSGQueryWithFilterService) FindServicesVenueProfileByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) error {
}

func (s *WSGQueryWithFilterService) FindServicesVlanPoolingByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) error {
}

func (s *WSGQueryWithFilterService) FindServicesWebAuthenticationByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) error {
}

func (s *WSGQueryWithFilterService) FindServicesWechatProfileByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) error {
}

func (s *WSGQueryWithFilterService) FindServicesWlanSchedulerByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) error {
}

func (s *WSGQueryWithFilterService) FindUserByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) error {
}

func (s *WSGQueryWithFilterService) FindWlanByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*wlanquery.WlanQueryList, error) {
}
