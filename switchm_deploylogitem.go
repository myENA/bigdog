package ruckus

// API Version: v9_0

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

func NewSwitchMDeployLogItemConfigurationHistoryDetailQueryResult() *SwitchMDeployLogItemConfigurationHistoryDetailQueryResult {
	m := new(SwitchMDeployLogItemConfigurationHistoryDetailQueryResult)
	return m
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
	//    - oneof:[PENDING,IN_PROGRESS,SUCCESS,FAILED]
	DispatchStatus *string `json:"dispatchStatus,omitempty"`

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

func NewSwitchMDeployLogItem() *SwitchMDeployLogItem {
	m := new(SwitchMDeployLogItem)
	return m
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

func NewSwitchMDeployLogItemDispatchFailedReasonType() *SwitchMDeployLogItemDispatchFailedReasonType {
	m := new(SwitchMDeployLogItemDispatchFailedReasonType)
	return m
}
