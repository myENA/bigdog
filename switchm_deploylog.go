package bigdog

// API Version: v9_1

import (
	"encoding/json"
	"net/http"
)

// SwitchMDeployLogConfigurationHistoryQueryResult
//
// Definition: deployLog_ConfigurationHistoryQueryResult
type SwitchMDeployLogConfigurationHistoryQueryResult struct {
	// FirstIndex
	// Index of the first config returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more configs after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// The list of configs in this response
	List []*SwitchMDeployLog `json:"list,omitempty"`

	// TotalCount
	// Total records count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchMDeployLogConfigurationHistoryQueryResultAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMDeployLogConfigurationHistoryQueryResult
}

func newSwitchMDeployLogConfigurationHistoryQueryResultAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(SwitchMDeployLogConfigurationHistoryQueryResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *SwitchMDeployLogConfigurationHistoryQueryResultAPIResponse) Hydrate() error {
	r.Data = new(SwitchMDeployLogConfigurationHistoryQueryResult)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSwitchMDeployLogConfigurationHistoryQueryResult() *SwitchMDeployLogConfigurationHistoryQueryResult {
	m := new(SwitchMDeployLogConfigurationHistoryQueryResult)
	return m
}

// SwitchMDeployLog
//
// Definition: deployLog_deployLog
type SwitchMDeployLog struct {
	// ConfigType
	// Config Type
	// Constraints:
	//    - oneof:[PROVISIONING,GLOBAL,AAA_SETTING,AAA_SERVER,COMMON,MODEL,SWITCH_SETTINGS,PORT_SETTINGS]
	ConfigType *string `json:"configType,omitempty"`

	// DeployScope
	// Deploy Scope
	// Constraints:
	//    - oneof:[PROVISION,GROUP,PORT,SWITCH]
	DeployScope *string `json:"deployScope,omitempty"`

	// DeployStatus
	// Deploy Status
	// Constraints:
	//    - oneof:[PENDING,STARTED,END]
	DeployStatus *string `json:"deployStatus,omitempty"`

	// EndTime
	// Deploy End Time
	EndTime *int `json:"endTime,omitempty"`

	// GroupId
	// Switch Group ID
	GroupId *string `json:"groupId,omitempty"`

	// Id
	// ID
	Id *string `json:"id,omitempty"`

	// ModelFamily
	// Switch Model Family
	ModelFamily *string `json:"modelFamily,omitempty"`

	NodeId *string `json:"nodeId,omitempty"`

	NodeName *string `json:"nodeName,omitempty"`

	// Scheduled
	// Schedule
	Scheduled *int `json:"scheduled,omitempty"`

	// StartTime
	// Deploy Start Time
	StartTime *int `json:"startTime,omitempty"`

	Summary *SwitchMDeployLogStatusSummary `json:"summary,omitempty"`

	// SwitchIds
	// Switch ID
	SwitchIds *string `json:"switchIds,omitempty"`

	// TransactionId
	// Deploy Transaction ID
	TransactionId *string `json:"transactionId,omitempty"`

	// Yang
	// YANG Model
	Yang *string `json:"yang,omitempty"`
}

func NewSwitchMDeployLog() *SwitchMDeployLog {
	m := new(SwitchMDeployLog)
	return m
}

// SwitchMDeployLogStatusSummary
//
// Definition: deployLog_deployLogStatusSummary
type SwitchMDeployLogStatusSummary struct {
	// Failed
	// Deployment Fail Counter
	Failed *int `json:"failed,omitempty"`

	// FailedNoResponse
	// Deployment Fail No Response Counter
	FailedNoResponse *int `json:"failedNoResponse,omitempty"`

	// FailedSaveFlash
	// Deployment Fail to Save Flash
	FailedSaveFlash *int `json:"failedSaveFlash,omitempty"`

	// Success
	// Deployment Success Counter
	Success *int `json:"success,omitempty"`
}

func NewSwitchMDeployLogStatusSummary() *SwitchMDeployLogStatusSummary {
	m := new(SwitchMDeployLogStatusSummary)
	return m
}
