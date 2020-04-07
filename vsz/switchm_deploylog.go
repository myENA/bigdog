package vsz

// API Version: v9_0

type SwitchMDeployLogConfigurationHistoryQueryResult struct {
	// FirstIndex
	// Index of the first config returned out of the complete list
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more configs after the current displayed list
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// The list of configs in this response
	// Constraints:
	//    - nullable
	List []*SwitchMDeployLog `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Total records count in this response
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMDeployLogConfigurationHistoryQueryResult() *SwitchMDeployLogConfigurationHistoryQueryResult {
	m := new(SwitchMDeployLogConfigurationHistoryQueryResult)
	return m
}

type SwitchMDeployLog struct {
	// ConfigType
	// Config Type
	// Constraints:
	//    - nullable
	//    - oneof:[PROVISIONING,GLOBAL,AAA_SETTING,AAA_SERVER,COMMON,MODEL,SWITCH_SETTINGS,PORT_SETTINGS]
	ConfigType *string `json:"configType,omitempty" validate:"omitempty,oneof=PROVISIONING GLOBAL AAA_SETTING AAA_SERVER COMMON MODEL SWITCH_SETTINGS PORT_SETTINGS"`

	// DeployScope
	// Deploy Scope
	// Constraints:
	//    - nullable
	//    - oneof:[PROVISION,GROUP,PORT,SWITCH]
	DeployScope *string `json:"deployScope,omitempty" validate:"omitempty,oneof=PROVISION GROUP PORT SWITCH"`

	// DeployStatus
	// Deploy Status
	// Constraints:
	//    - nullable
	//    - oneof:[PENDING,STARTED,END]
	DeployStatus *string `json:"deployStatus,omitempty" validate:"omitempty,oneof=PENDING STARTED END"`

	// EndTime
	// Deploy End Time
	// Constraints:
	//    - nullable
	EndTime *int `json:"endTime,omitempty"`

	// GroupId
	// Switch Group ID
	// Constraints:
	//    - nullable
	GroupId *string `json:"groupId,omitempty"`

	// Id
	// ID
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// ModelFamily
	// Switch Model Family
	// Constraints:
	//    - nullable
	ModelFamily *string `json:"modelFamily,omitempty"`

	// NodeId
	// Constraints:
	//    - nullable
	NodeId *string `json:"nodeId,omitempty"`

	// NodeName
	// Constraints:
	//    - nullable
	NodeName *string `json:"nodeName,omitempty"`

	// Scheduled
	// Schedule
	// Constraints:
	//    - nullable
	Scheduled *int `json:"scheduled,omitempty"`

	// StartTime
	// Deploy Start Time
	// Constraints:
	//    - nullable
	StartTime *int `json:"startTime,omitempty"`

	// Summary
	// Constraints:
	//    - nullable
	Summary *SwitchMDeployLogStatusSummary `json:"summary,omitempty"`

	// SwitchIds
	// Switch ID
	// Constraints:
	//    - nullable
	SwitchIds *string `json:"switchIds,omitempty"`

	// TransactionId
	// Deploy Transaction ID
	// Constraints:
	//    - nullable
	TransactionId *string `json:"transactionId,omitempty"`

	// Yang
	// YANG Model
	// Constraints:
	//    - nullable
	Yang *string `json:"yang,omitempty"`
}

func NewSwitchMDeployLog() *SwitchMDeployLog {
	m := new(SwitchMDeployLog)
	return m
}

type SwitchMDeployLogStatusSummary struct {
	// Failed
	// Deployment Fail Counter
	// Constraints:
	//    - nullable
	Failed *int `json:"failed,omitempty"`

	// FailedNoResponse
	// Deployment Fail No Response Counter
	// Constraints:
	//    - nullable
	FailedNoResponse *int `json:"failedNoResponse,omitempty"`

	// FailedSaveFlash
	// Deployment Fail to Save Flash
	// Constraints:
	//    - nullable
	FailedSaveFlash *int `json:"failedSaveFlash,omitempty"`

	// Success
	// Deployment Success Counter
	// Constraints:
	//    - nullable
	Success *int `json:"success,omitempty"`
}

func NewSwitchMDeployLogStatusSummary() *SwitchMDeployLogStatusSummary {
	m := new(SwitchMDeployLogStatusSummary)
	return m
}
