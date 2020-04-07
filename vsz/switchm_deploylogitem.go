package vsz

// API Version: v9_0

type SwitchMDeployLogItemConfigurationHistoryDetailQueryResult struct {
	// FirstIndex
	// Indicator of whether there are more configs after the current displayed list
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more configs after the current displayed list
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Indicator of whether there are more configs after the current displayed list
	// Constraints:
	//    - nullable
	List []*SwitchMDeployLogItem `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Total records count in this response
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMDeployLogItemConfigurationHistoryDetailQueryResult() *SwitchMDeployLogItemConfigurationHistoryDetailQueryResult {
	m := new(SwitchMDeployLogItemConfigurationHistoryDetailQueryResult)
	return m
}

type SwitchMDeployLogItem struct {
	// Clis
	// CLI Command
	// Constraints:
	//    - nullable
	Clis *string `json:"clis,omitempty"`

	// DispatchFailedReason
	// Deployment Fail Description
	// Constraints:
	//    - nullable
	DispatchFailedReason *SwitchMDeployLogItemDispatchFailedReasonType `json:"dispatchFailedReason,omitempty"`

	// DispatchStatus
	// Status of Deployment
	// Constraints:
	//    - nullable
	//    - oneof:[PENDING,IN_PROGRESS,SUCCESS,FAILED]
	DispatchStatus *string `json:"dispatchStatus,omitempty" validate:"omitempty,oneof=PENDING IN_PROGRESS SUCCESS FAILED"`

	// EndTime
	// Deployment end datetime
	// Constraints:
	//    - nullable
	EndTime *int `json:"endTime,omitempty"`

	// Id
	// Switch Deployment History Id
	// Constraints:
	//    - nullable
	Id *string `json:"id,omitempty"`

	// SerialNumber
	// Serical Number of Switch
	// Constraints:
	//    - nullable
	SerialNumber *string `json:"serialNumber,omitempty"`

	// StartTime
	// Deployment begin datetime
	// Constraints:
	//    - nullable
	StartTime *int `json:"startTime,omitempty"`

	// SwitchId
	// Switch Id
	// Constraints:
	//    - nullable
	SwitchId *string `json:"switchId,omitempty"`

	// SwitchName
	// Switch Name
	// Constraints:
	//    - nullable
	SwitchName *string `json:"switchName,omitempty"`

	// TransactionId
	// Transaction Id of depolyment
	// Constraints:
	//    - nullable
	TransactionId *string `json:"transactionId,omitempty"`
}

func NewSwitchMDeployLogItem() *SwitchMDeployLogItem {
	m := new(SwitchMDeployLogItem)
	return m
}

// SwitchMDeployLogItemDispatchFailedReasonType
//
// Deployment Fail Description
// Constraints:
//    - nullable
type SwitchMDeployLogItemDispatchFailedReasonType struct {
	// LineNumber
	// Line Number Fail to Execute Cmd
	// Constraints:
	//    - nullable
	LineNumber *int `json:"lineNumber,omitempty"`

	// Message
	// Status
	// Constraints:
	//    - nullable
	Message *int `json:"message,omitempty"`
}

func NewSwitchMDeployLogItemDispatchFailedReasonType() *SwitchMDeployLogItemDispatchFailedReasonType {
	m := new(SwitchMDeployLogItemDispatchFailedReasonType)
	return m
}
