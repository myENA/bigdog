package deploylogitem

// API Version: v8_1

type ConfigurationHistoryDetailQueryResult struct {
	// FirstIndex
	// Indicator of whether there are more configs after the current displayed list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more configs after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Indicator of whether there are more configs after the current displayed list
	List []*DeployLogItem `json:"list,omitempty"`

	// TotalCount
	// Total records count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

type DeployLogItem struct {
	// Clis
	// CLI Command
	Clis *string `json:"clis,omitempty"`

	// DispatchFailedReason
	// Deployment Fail Description
	DispatchFailedReason *DeployLogItemDispatchFailedReasonType `json:"dispatchFailedReason,omitempty"`

	// DispatchStatus
	// Status of Deployment
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

// DeployLogItemDispatchFailedReasonType
//
// Deployment Fail Description
type DeployLogItemDispatchFailedReasonType struct {
	// LineNumber
	// Line Number Fail to Execute Cmd
	LineNumber *int `json:"lineNumber,omitempty"`

	// Message
	// Status
	Message *int `json:"message,omitempty"`
}
