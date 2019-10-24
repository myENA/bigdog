package deploylog

// API Version: v8_1

type ConfigurationHistoryQueryResult struct {
	// FirstIndex
	// Index of the first config returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more configs after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// The list of configs in this response
	List []*DeployLog `json:"list,omitempty"`

	// TotalCount
	// Total records count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

type DeployLog struct {
	// ConfigType
	// Config Type
	// Constraints:
	//    - nullable
	//    - oneof:[PROVISIONING,GLOBAL,COMMON,MODEL,SWITCH_SETTINGS,PORT_SETTINGS]
	ConfigType *string `json:"configType,omitempty" validate:"omitempty,oneof=PROVISIONING GLOBAL COMMON MODEL SWITCH_SETTINGS PORT_SETTINGS"`

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

	// Scheduled
	// Schedule
	Scheduled *int `json:"scheduled,omitempty"`

	// StartTime
	// Deploy Start Time
	StartTime *int `json:"startTime,omitempty"`

	Summary *DeployLogStatusSummary `json:"summary,omitempty"`

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

type DeployLogStatusSummary struct {
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
