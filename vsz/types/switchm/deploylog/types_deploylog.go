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
	ConfigType *string `json:"configType,omitempty" validate:"oneof=PROVISIONING GLOBAL COMMON MODEL SWITCH_SETTINGS PORT_SETTINGS"`

	// DeployScope
	// Deploy Scope
	DeployScope *string `json:"deployScope,omitempty" validate:"oneof=PROVISION GROUP PORT SWITCH"`

	// DeployStatus
	// Deploy Status
	DeployStatus *string `json:"deployStatus,omitempty" validate:"oneof=PENDING STARTED END"`

	// EndTime
	// Deploy End Time
	EndTime *int `json:"endTime,omitempty"`

	// GroupID
	// Switch Group ID
	GroupID *string `json:"groupId,omitempty"`

	// ID
	// ID
	ID *string `json:"id,omitempty"`

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

	// TransactionID
	// Deploy Transaction ID
	TransactionID *string `json:"transactionId,omitempty"`

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
