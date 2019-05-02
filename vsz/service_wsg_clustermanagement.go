package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/administration"
	"github.com/myENA/ruckus-client/vsz/types/wsg/clusterblade"
	"github.com/myENA/ruckus-client/vsz/types/wsg/clusterredundancy"
)

type WSGClusterManagementService struct {
    client *Client
}

func NewWSGClusterManagementService (client *Client) *WSGClusterManagementService {
    s := new(WSGClusterManagementService)
    s.client = client
    return s
}

func (ss *WSGService) WSGClusterManagementService () *WSGClusterManagementService {
    serv := new(WSGClusterManagementService)
    serv.client = ss.client
    return serv
}

func (s *WSGClusterManagementService) AddApPatchFile (ctx context.Context) error {
}

func (s *WSGClusterManagementService) AddClusterBackup (ctx context.Context) error {
}

func (s *WSGClusterManagementService) AddClusterRestoreById (ctx context.Context, id string) error {
}

func (s *WSGClusterManagementService) AddConfigurationBackup (ctx context.Context) error {
}

func (s *WSGClusterManagementService) AddConfigurationRestoreById (ctx context.Context, id string) error {
}

func (s *WSGClusterManagementService) AddConfigurationUpload (ctx context.Context) error {
}

func (s *WSGClusterManagementService) AddUpgrade (ctx context.Context) (administration.UpgradeStatus, error) {
}

func (s *WSGClusterManagementService) AddUpgradeUpload (ctx context.Context) (administration.UpgradeStatus, error) {
}

func (s *WSGClusterManagementService) DeleteClusterById (ctx context.Context, id string) error {
}

func (s *WSGClusterManagementService) DeleteConfigurationById (ctx context.Context, id string) error {
}

func (s *WSGClusterManagementService) FindApPatch (ctx context.Context) (administration.ApPatchInfo, error) {
}

func (s *WSGClusterManagementService) FindApPatchHistory (ctx context.Context) (administration.ApPatchHistoryList, error) {
}

func (s *WSGClusterManagementService) FindApPatchStatus (ctx context.Context) (administration.ApPatchStatus, error) {
}

func (s *WSGClusterManagementService) FindCluster (ctx context.Context) (administration.ClusterBackupList, error) {
}

func (s *WSGClusterManagementService) FindClusterGeoRedundancy (ctx context.Context) (clusterredundancy.ClusterRedundancySettings, error) {
}

func (s *WSGClusterManagementService) FindClusterNodeStatus (ctx context.Context) (clusterblade.ControlNodeStatus, error) {
}

func (s *WSGClusterManagementService) FindClusterState (ctx context.Context) (clusterblade.ClusterState, error) {
}

func (s *WSGClusterManagementService) FindClusterStatus (ctx context.Context) (clusterblade.ClusterStatus, error) {
}

func (s *WSGClusterManagementService) FindConfiguration (ctx context.Context) (administration.ConfigurationBackupList, error) {
}

func (s *WSGClusterManagementService) FindConfigurationDownload (ctx context.Context) (json.RawMessage, error) {
}

func (s *WSGClusterManagementService) FindConfigurationSettingsAutoExportBackup (ctx context.Context) (administration.AutoExportBackup, error) {
}

func (s *WSGClusterManagementService) FindConfigurationSettingsScheduleBackup (ctx context.Context) (administration.ScheduleBackup, error) {
}

func (s *WSGClusterManagementService) FindUpgradeHistory (ctx context.Context) (administration.UpgradeHistoryList, error) {
}

func (s *WSGClusterManagementService) FindUpgradePatch (ctx context.Context) (administration.UpgradePatchInfo, error) {
}

func (s *WSGClusterManagementService) FindUpgradeStatus (ctx context.Context) (administration.UpgradeStatus, error) {
}

