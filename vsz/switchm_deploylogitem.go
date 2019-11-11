package vsz

// API Version: v8_1

type SwitchMDeployLogItemConfigurationHistoryDetailQueryResult struct {
	// FirstIndex
	// Indicator of whether there are more configs after the current displayed list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more configs after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Indicator of whether there are more configs after the current displayed list
	List []*SwitchMDeployLogItem `json:"list,omitempty"`

	// TotalCount
	// Total records count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchMDeployLogItem struct {
	// Clis
	// CLI Command
	Clis *string `json:"clis,omitempty"`

	// DispatchFailedReason
	// Deployment Fail Description
	DispatchFailedReason *SwitchMDeployLogItemDispatchFailedReasonType `json:"dispatchFailedReason,omitempty"`

	// DispatchStatus
	// Status of Deployment
	// Constraints:
	//    - nullable
	//    - oneof:[PENDING,IN_PROGRESS,SUCCESS,FAILED]
	DispatchStatus *string `json:"dispatchStatus,omitempty" validate:"omitempty,oneof=PENDING IN_PROGRESS SUCCESS FAILED"`

	// EndTime
	// Deployment end datetime
	EndTime *int `json:"endTime,omitempty"`

	// Id
	// Switch Deployment History Id
	Id *string `json:"id,omitempty"`

	// SerialNumber
	// Serical Number of Switch
	SerialNumber *string `json:"serialNumber,omitempty"`

	// StartTime
	// Deployment begin datetime
	StartTime *int `json:"startTime,omitempty"`

	// SwitchId
	// Switch Id
	SwitchId *string `json:"switchId,omitempty"`

	// SwitchName
	// Switch Name
	SwitchName *string `json:"switchName,omitempty"`

	// TransactionId
	// Transaction Id of depolyment
	TransactionId *string `json:"transactionId,omitempty"`
}

// SwitchMDeployLogItemDispatchFailedReasonType
//
// Deployment Fail Description
type SwitchMDeployLogItemDispatchFailedReasonType struct {
	// LineNumber
	// Line Number Fail to Execute Cmd
	LineNumber *int `json:"lineNumber,omitempty"`

	// Message
	// Status
	Message *int `json:"message,omitempty"`
}
