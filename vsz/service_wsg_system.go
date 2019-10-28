package vsz

// API Version: v8_1

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

func NewWSGSystemService(client *Client) *WSGSystemService {
	s := new(WSGSystemService)
	s.client = client
	return s
}

func (ss *WSGService) WSGSystemService() *WSGSystemService {
	serv := new(WSGSystemService)
	serv.client = ss.client
	return serv
}

// AddSystemAp_balance
//
// Execute ap balance.
func (s *WSGSystemService) AddSystemAp_balance(ctx context.Context) error {
}

// AddSystemApRoutineStatusIntervalSlowdown
//
// Use this API command to set AP routine status interval setting to 900 seconds.
func (s *WSGSystemService) AddSystemApRoutineStatusIntervalSlowdown(ctx context.Context) error {
}

// AddSystemApRoutineStatusIntervalSpeedup
//
// Use this API command to set AP routine status interval setting to 180 seconds.
func (s *WSGSystemService) AddSystemApRoutineStatusIntervalSpeedup(ctx context.Context) error {
}

// FindController
//
// Use this API command to retrieve the system summary.
func (s *WSGSystemService) FindController(ctx context.Context) (*system.ControllerList, error) {
}

// FindControllerStatisticsById
//
// Use this API command to retrieve the system statistics.
//
// Path Parameters:
// - pId string
//		- required
//
// Query Parameters:
// - qInterval string
// - qSize float64
func (s *WSGSystemService) FindControllerStatisticsById(ctx context.Context, pId string, qInterval string, qSize float64) (system.StatisticList, error) {
}

// FindSystem
//
// Use this API command to get settings of system. Currently, Only can get settings about AP number limit.
func (s *WSGSystemService) FindSystem(ctx context.Context) (*system.SystemSettings, error) {
}

// FindSystemApmodels
//
// Use this API command to retrieve support AP models for the current installed SZ version's default AP firmware.
func (s *WSGSystemService) FindSystemApmodels(ctx context.Context) error {
}

// FindSystemApmodelsByFirmwareVersion
//
// Use this API command to retrieve support AP models from input firmware version.
//
// Path Parameters:
// - pFirmwareVersion string
//		- required
func (s *WSGSystemService) FindSystemApmodelsByFirmwareVersion(ctx context.Context, pFirmwareVersion string) error {
}

// FindSystemApRoutineConfigInterval
//
// Use this API command to get AP routine configuration interval setting.
func (s *WSGSystemService) FindSystemApRoutineConfigInterval(ctx context.Context) (*aproutineconfiginterval.ApRoutineConfigIntervalRsp, error) {
}

// FindSystemApRoutineStatusInterval
//
// Use this API command to get AP routine status interval setting.
func (s *WSGSystemService) FindSystemApRoutineStatusInterval(ctx context.Context) (*aproutinestatusinterval.ApRoutineStatusIntervalRsp, error) {
}

// FindSystemByQueryCriteria
//
// Use this API command to query settings of system. Currently, Only can get settings about AP number limit by query criteria with domain and zone filters.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGSystemService) FindSystemByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*system.SystemSettings, error) {
}

// FindSystemDevicesSummary
//
// Use this API command to retrieve devices summary.
func (s *WSGSystemService) FindSystemDevicesSummary(ctx context.Context) (*devicecapacity.DevicesSummary, error) {
}

// FindSystemGatewayAdvanced
//
// Use this API command to retrieve gateway advanced setting.
func (s *WSGSystemService) FindSystemGatewayAdvanced(ctx context.Context) (*system.GatewayAdvanced, error) {
}

// FindSystemInventory
//
// Use this API command to retrieve the system inventory with current logon user domain.
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGSystemService) FindSystemInventory(ctx context.Context, qIndex string, qListSize string) (*system.InventoryList, error) {
}

// FindSystemNbi
//
// Use this API command to retrieve user information by Northbound Portal Interface.
//
// Query Parameters:
// - qDomainId string
func (s *WSGSystemService) FindSystemNbi(ctx context.Context, qDomainId string) (*system.NorthboundInterface, error) {
}

// FindSystemSystemTime
//
// Retrieve System Time Setting.
func (s *WSGSystemService) FindSystemSystemTime(ctx context.Context) (*system.SystemTimeSetting, error) {
}

// PartialUpdateSystemCaptcha
//
// Use this API command to modify the CAPTCHA setting.
//
// Request Body:
//	 - body *system.CaptchaSetting
func (s *WSGSystemService) PartialUpdateSystemCaptcha(ctx context.Context, body *system.CaptchaSetting) error {
}
