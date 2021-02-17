package bigdog

// API Version: v9_1

import (
	"errors"
	"io"
)

// SwitchMDeployLogItemConfigurationHistoryDetailQueryResult
//
// Definition: deployLogItem_ConfigurationHistoryDetailQueryResult
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

type SwitchMDeployLogItemConfigurationHistoryDetailQueryResultAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMDeployLogItemConfigurationHistoryDetailQueryResult
}

func newSwitchMDeployLogItemConfigurationHistoryDetailQueryResultAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMDeployLogItemConfigurationHistoryDetailQueryResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMDeployLogItemConfigurationHistoryDetailQueryResultAPIResponse) Hydrate() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return nil
		}
		return r.err
	}
	data := new(SwitchMDeployLogItemConfigurationHistoryDetailQueryResult)
	if err := r.doHydrate(data); err != nil {
		return err
	}
	r.Data = data
	return nil
}
func NewSwitchMDeployLogItemConfigurationHistoryDetailQueryResult() *SwitchMDeployLogItemConfigurationHistoryDetailQueryResult {
	m := new(SwitchMDeployLogItemConfigurationHistoryDetailQueryResult)
	return m
}

// SwitchMDeployLogItem
//
// Definition: deployLogItem_deployLogItem
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
// Definition: deployLogItem_deployLogItemDispatchFailedReasonType
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
