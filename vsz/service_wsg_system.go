package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/aproutineconfiginterval"
	"github.com/myENA/ruckus-client/vsz/types/wsg/aproutinestatusinterval"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/devicecapacity"
	"github.com/myENA/ruckus-client/vsz/types/wsg/system"
)

type WSGSystemService struct {
    client *Client
}

func NewWSGSystemService (client *Client) *WSGSystemService {
    s := new(WSGSystemService)
    s.client = client
    return s
}

func (ss *WSGService) WSGSystemService () *WSGSystemService {
    serv := new(WSGSystemService)
    serv.client = ss.client
    return serv
}

func (s *WSGSystemService) AddSystemAp_balance (ctx context.Context) error {
}

func (s *WSGSystemService) AddSystemApRoutineStatusIntervalSlowdown (ctx context.Context) error {
}

func (s *WSGSystemService) AddSystemApRoutineStatusIntervalSpeedup (ctx context.Context) error {
}

func (s *WSGSystemService) FindController (ctx context.Context) (*system.ControllerList, error) {
}

func (s *WSGSystemService) FindControllerStatisticsById (ctx context.Context, id string) (system.StatisticList, error) {
}

func (s *WSGSystemService) FindSystem (ctx context.Context) (*system.SystemSettings, error) {
}

func (s *WSGSystemService) FindSystemApmodels (ctx context.Context) error {
}

func (s *WSGSystemService) FindSystemApmodelsByFirmwareVersion (ctx context.Context, firmwareVersion:.+ string) error {
}

func (s *WSGSystemService) FindSystemApRoutineConfigInterval (ctx context.Context) (*aproutineconfiginterval.ApRoutineConfigIntervalRsp, error) {
}

func (s *WSGSystemService) FindSystemApRoutineStatusInterval (ctx context.Context) (*aproutinestatusinterval.ApRoutineStatusIntervalRsp, error) {
}

func (s *WSGSystemService) FindSystemByQueryCriteria (ctx context.Context) (*system.SystemSettings, error) {
}

func (s *WSGSystemService) FindSystemDevicesSummary (ctx context.Context) (*devicecapacity.DevicesSummary, error) {
}

func (s *WSGSystemService) FindSystemGatewayAdvanced (ctx context.Context) (*system.GatewayAdvanced, error) {
}

func (s *WSGSystemService) FindSystemInventory (ctx context.Context) (*system.InventoryList, error) {
}

func (s *WSGSystemService) FindSystemNbi (ctx context.Context) (*system.NorthboundInterface, error) {
}

func (s *WSGSystemService) FindSystemSystemTime (ctx context.Context) (*system.SystemTimeSetting, error) {
}

func (s *WSGSystemService) PartialUpdateSystemCaptcha (ctx context.Context) error {
}

