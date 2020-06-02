package vsz

// API Version: v9_0

type SwitchMDeploylogConfigurationHistoryQueryResult struct {
	// FirstIndex
	// Index of the first config returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more configs after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// The list of configs in this response
	List []*SwitchMDeploylog `json:"list,omitempty"`

	// TotalCount
	// Total records count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMDeploylogConfigurationHistoryQueryResult() *SwitchMDeploylogConfigurationHistoryQueryResult {
	m := new(SwitchMDeploylogConfigurationHistoryQueryResult)
	return m
}

type SwitchMDeploylog struct {
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

	Summary *SwitchMDeploylogStatusSummary `json:"summary,omitempty"`

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

func NewSwitchMDeploylog() *SwitchMDeploylog {
	m := new(SwitchMDeploylog)
	return m
}

type SwitchMDeploylogStatusSummary struct {
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

func NewSwitchMDeploylogStatusSummary() *SwitchMDeploylogStatusSummary {
	m := new(SwitchMDeploylogStatusSummary)
	return m
}
