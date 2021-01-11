package bigdog

// API Version: 1.0.0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SCIAirtimeUtilizationReportService struct {
	apiClient *SCIClient
}

func NewSCIAirtimeUtilizationReportService(c *SCIClient) *SCIAirtimeUtilizationReportService {
	s := new(SCIAirtimeUtilizationReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIAirtimeUtilizationReportService() *SCIAirtimeUtilizationReportService {
	return NewSCIAirtimeUtilizationReportService(ss.apiClient)
}

// SCIAirtimeUtilizationReport1overviewData
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_1_overview_Data
type SCIAirtimeUtilizationReport1overviewData []*SCIAirtimeUtilizationReport1overviewDataType

func MakeSCIAirtimeUtilizationReport1overviewData() SCIAirtimeUtilizationReport1overviewData {
	m := make(SCIAirtimeUtilizationReport1overviewData, 0)
	return m
}

// SCIAirtimeUtilizationReport1overviewDataType
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_1_overview_DataType
type SCIAirtimeUtilizationReport1overviewDataType struct {
	AirtimeUtilizationAvg *float64 `json:"airtimeUtilizationAvg,omitempty"`

	Radio *string `json:"radio,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAirtimeUtilizationReport1overviewDataType) UnmarshalJSON(b []byte) error {
	type _SCIAirtimeUtilizationReport1overviewDataType SCIAirtimeUtilizationReport1overviewDataType
	tmpType := new(_SCIAirtimeUtilizationReport1overviewDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "airtimeUtilizationAvg")
	delete(tmpType.XAdditionalProperties, "radio")
	*t = SCIAirtimeUtilizationReport1overviewDataType(*tmpType)
	return nil
}

func (t *SCIAirtimeUtilizationReport1overviewDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.AirtimeUtilizationAvg != nil {
		tmp["airtimeUtilizationAvg"] = t.AirtimeUtilizationAvg
	}
	if t.Radio != nil {
		tmp["radio"] = t.Radio
	}
	return json.Marshal(tmp)
}

func NewSCIAirtimeUtilizationReport1overviewDataType() *SCIAirtimeUtilizationReport1overviewDataType {
	m := new(SCIAirtimeUtilizationReport1overviewDataType)
	return m
}

// SCIAirtimeUtilizationReport2topChartData
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_2_topChart_Data
type SCIAirtimeUtilizationReport2topChartData []*SCIAirtimeUtilizationReport2topChartDataType

func MakeSCIAirtimeUtilizationReport2topChartData() SCIAirtimeUtilizationReport2topChartData {
	m := make(SCIAirtimeUtilizationReport2topChartData, 0)
	return m
}

// SCIAirtimeUtilizationReport2topChartDataType
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_2_topChart_DataType
type SCIAirtimeUtilizationReport2topChartDataType struct {
	Key *string `json:"key,omitempty"`

	Label *string `json:"label,omitempty"`

	Value *float64 `json:"value,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAirtimeUtilizationReport2topChartDataType) UnmarshalJSON(b []byte) error {
	type _SCIAirtimeUtilizationReport2topChartDataType SCIAirtimeUtilizationReport2topChartDataType
	tmpType := new(_SCIAirtimeUtilizationReport2topChartDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "key")
	delete(tmpType.XAdditionalProperties, "label")
	delete(tmpType.XAdditionalProperties, "value")
	*t = SCIAirtimeUtilizationReport2topChartDataType(*tmpType)
	return nil
}

func (t *SCIAirtimeUtilizationReport2topChartDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Key != nil {
		tmp["key"] = t.Key
	}
	if t.Label != nil {
		tmp["label"] = t.Label
	}
	if t.Value != nil {
		tmp["value"] = t.Value
	}
	return json.Marshal(tmp)
}

func NewSCIAirtimeUtilizationReport2topChartDataType() *SCIAirtimeUtilizationReport2topChartDataType {
	m := new(SCIAirtimeUtilizationReport2topChartDataType)
	return m
}

// SCIAirtimeUtilizationReport3topAPsByAirtime24TableData
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_3_topAPsByAirtime24Table_Data
type SCIAirtimeUtilizationReport3topAPsByAirtime24TableData []*SCIAirtimeUtilizationReport3topAPsByAirtime24TableDataType

func MakeSCIAirtimeUtilizationReport3topAPsByAirtime24TableData() SCIAirtimeUtilizationReport3topAPsByAirtime24TableData {
	m := make(SCIAirtimeUtilizationReport3topAPsByAirtime24TableData, 0)
	return m
}

// SCIAirtimeUtilizationReport3topAPsByAirtime24TableDataType
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_3_topAPsByAirtime24Table_DataType
type SCIAirtimeUtilizationReport3topAPsByAirtime24TableDataType struct {
	AirtimeBusyAvg *float64 `json:"airtimeBusyAvg,omitempty"`

	AirtimeIdleAvg *float64 `json:"airtimeIdleAvg,omitempty"`

	AirtimeRxAvg *float64 `json:"airtimeRxAvg,omitempty"`

	AirtimeTxAvg *float64 `json:"airtimeTxAvg,omitempty"`

	AirtimeUtilizationAvg *float64 `json:"airtimeUtilizationAvg,omitempty"`

	ApIp *string `json:"apIp,omitempty"`

	ApMac *string `json:"apMac,omitempty"`

	ApName *string `json:"apName,omitempty"`

	ClientCount *float64 `json:"clientCount,omitempty"`

	CtrlMac *string `json:"ctrlMac,omitempty"`

	CtrlName *string `json:"ctrlName,omitempty"`

	CtrlSerial *string `json:"ctrlSerial,omitempty"`

	Index *float64 `json:"index,omitempty"`

	MgmtRxBytes *float64 `json:"mgmtRxBytes,omitempty"`

	MgmtTraffic *float64 `json:"mgmtTraffic,omitempty"`

	MgmtTxBytes *float64 `json:"mgmtTxBytes,omitempty"`

	SessionCount *float64 `json:"sessionCount,omitempty"`

	TotalRxTraffic *float64 `json:"totalRxTraffic,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	TotalTxTraffic *float64 `json:"totalTxTraffic,omitempty"`

	UserRxBytes *float64 `json:"userRxBytes,omitempty"`

	UserTraffic *float64 `json:"userTraffic,omitempty"`

	UserTxBytes *float64 `json:"userTxBytes,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAirtimeUtilizationReport3topAPsByAirtime24TableDataType) UnmarshalJSON(b []byte) error {
	type _SCIAirtimeUtilizationReport3topAPsByAirtime24TableDataType SCIAirtimeUtilizationReport3topAPsByAirtime24TableDataType
	tmpType := new(_SCIAirtimeUtilizationReport3topAPsByAirtime24TableDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "airtimeBusyAvg")
	delete(tmpType.XAdditionalProperties, "airtimeIdleAvg")
	delete(tmpType.XAdditionalProperties, "airtimeRxAvg")
	delete(tmpType.XAdditionalProperties, "airtimeTxAvg")
	delete(tmpType.XAdditionalProperties, "airtimeUtilizationAvg")
	delete(tmpType.XAdditionalProperties, "apIp")
	delete(tmpType.XAdditionalProperties, "apMac")
	delete(tmpType.XAdditionalProperties, "apName")
	delete(tmpType.XAdditionalProperties, "clientCount")
	delete(tmpType.XAdditionalProperties, "ctrlMac")
	delete(tmpType.XAdditionalProperties, "ctrlName")
	delete(tmpType.XAdditionalProperties, "ctrlSerial")
	delete(tmpType.XAdditionalProperties, "index")
	delete(tmpType.XAdditionalProperties, "mgmtRxBytes")
	delete(tmpType.XAdditionalProperties, "mgmtTraffic")
	delete(tmpType.XAdditionalProperties, "mgmtTxBytes")
	delete(tmpType.XAdditionalProperties, "sessionCount")
	delete(tmpType.XAdditionalProperties, "totalRxTraffic")
	delete(tmpType.XAdditionalProperties, "totalTraffic")
	delete(tmpType.XAdditionalProperties, "totalTxTraffic")
	delete(tmpType.XAdditionalProperties, "userRxBytes")
	delete(tmpType.XAdditionalProperties, "userTraffic")
	delete(tmpType.XAdditionalProperties, "userTxBytes")
	*t = SCIAirtimeUtilizationReport3topAPsByAirtime24TableDataType(*tmpType)
	return nil
}

func (t *SCIAirtimeUtilizationReport3topAPsByAirtime24TableDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.AirtimeBusyAvg != nil {
		tmp["airtimeBusyAvg"] = t.AirtimeBusyAvg
	}
	if t.AirtimeIdleAvg != nil {
		tmp["airtimeIdleAvg"] = t.AirtimeIdleAvg
	}
	if t.AirtimeRxAvg != nil {
		tmp["airtimeRxAvg"] = t.AirtimeRxAvg
	}
	if t.AirtimeTxAvg != nil {
		tmp["airtimeTxAvg"] = t.AirtimeTxAvg
	}
	if t.AirtimeUtilizationAvg != nil {
		tmp["airtimeUtilizationAvg"] = t.AirtimeUtilizationAvg
	}
	if t.ApIp != nil {
		tmp["apIp"] = t.ApIp
	}
	if t.ApMac != nil {
		tmp["apMac"] = t.ApMac
	}
	if t.ApName != nil {
		tmp["apName"] = t.ApName
	}
	if t.ClientCount != nil {
		tmp["clientCount"] = t.ClientCount
	}
	if t.CtrlMac != nil {
		tmp["ctrlMac"] = t.CtrlMac
	}
	if t.CtrlName != nil {
		tmp["ctrlName"] = t.CtrlName
	}
	if t.CtrlSerial != nil {
		tmp["ctrlSerial"] = t.CtrlSerial
	}
	if t.Index != nil {
		tmp["index"] = t.Index
	}
	if t.MgmtRxBytes != nil {
		tmp["mgmtRxBytes"] = t.MgmtRxBytes
	}
	if t.MgmtTraffic != nil {
		tmp["mgmtTraffic"] = t.MgmtTraffic
	}
	if t.MgmtTxBytes != nil {
		tmp["mgmtTxBytes"] = t.MgmtTxBytes
	}
	if t.SessionCount != nil {
		tmp["sessionCount"] = t.SessionCount
	}
	if t.TotalRxTraffic != nil {
		tmp["totalRxTraffic"] = t.TotalRxTraffic
	}
	if t.TotalTraffic != nil {
		tmp["totalTraffic"] = t.TotalTraffic
	}
	if t.TotalTxTraffic != nil {
		tmp["totalTxTraffic"] = t.TotalTxTraffic
	}
	if t.UserRxBytes != nil {
		tmp["userRxBytes"] = t.UserRxBytes
	}
	if t.UserTraffic != nil {
		tmp["userTraffic"] = t.UserTraffic
	}
	if t.UserTxBytes != nil {
		tmp["userTxBytes"] = t.UserTxBytes
	}
	return json.Marshal(tmp)
}

func NewSCIAirtimeUtilizationReport3topAPsByAirtime24TableDataType() *SCIAirtimeUtilizationReport3topAPsByAirtime24TableDataType {
	m := new(SCIAirtimeUtilizationReport3topAPsByAirtime24TableDataType)
	return m
}

// SCIAirtimeUtilizationReport3topAPsByAirtime24TableMetaData
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_3_topAPsByAirtime24Table_MetaData
type SCIAirtimeUtilizationReport3topAPsByAirtime24TableMetaData struct {
	MaxValues *SCIAirtimeUtilizationReport3topAPsByAirtime24TableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAirtimeUtilizationReport3topAPsByAirtime24TableMetaData) UnmarshalJSON(b []byte) error {
	type _SCIAirtimeUtilizationReport3topAPsByAirtime24TableMetaData SCIAirtimeUtilizationReport3topAPsByAirtime24TableMetaData
	tmpType := new(_SCIAirtimeUtilizationReport3topAPsByAirtime24TableMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "maxValues")
	*t = SCIAirtimeUtilizationReport3topAPsByAirtime24TableMetaData(*tmpType)
	return nil
}

func (t *SCIAirtimeUtilizationReport3topAPsByAirtime24TableMetaData) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.MaxValues != nil {
		tmp["maxValues"] = t.MaxValues
	}
	return json.Marshal(tmp)
}

func NewSCIAirtimeUtilizationReport3topAPsByAirtime24TableMetaData() *SCIAirtimeUtilizationReport3topAPsByAirtime24TableMetaData {
	m := new(SCIAirtimeUtilizationReport3topAPsByAirtime24TableMetaData)
	return m
}

// SCIAirtimeUtilizationReport3topAPsByAirtime24TableMetaDataMaxValuesType
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_3_topAPsByAirtime24Table_MetaDataMaxValuesType
type SCIAirtimeUtilizationReport3topAPsByAirtime24TableMetaDataMaxValuesType struct {
	MgmtRxBytes *float64 `json:"mgmtRxBytes,omitempty"`

	MgmtTraffic *float64 `json:"mgmtTraffic,omitempty"`

	MgmtTxBytes *float64 `json:"mgmtTxBytes,omitempty"`

	TotalRxTraffic *float64 `json:"totalRxTraffic,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	TotalTxTraffic *float64 `json:"totalTxTraffic,omitempty"`

	UserRxBytes *float64 `json:"userRxBytes,omitempty"`

	UserTraffic *float64 `json:"userTraffic,omitempty"`

	UserTxBytes *float64 `json:"userTxBytes,omitempty"`
}

func NewSCIAirtimeUtilizationReport3topAPsByAirtime24TableMetaDataMaxValuesType() *SCIAirtimeUtilizationReport3topAPsByAirtime24TableMetaDataMaxValuesType {
	m := new(SCIAirtimeUtilizationReport3topAPsByAirtime24TableMetaDataMaxValuesType)
	return m
}

// SCIAirtimeUtilizationReport4topAPsByAirtime5TableData
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_4_topAPsByAirtime5Table_Data
type SCIAirtimeUtilizationReport4topAPsByAirtime5TableData []*SCIAirtimeUtilizationReport4topAPsByAirtime5TableDataType

func MakeSCIAirtimeUtilizationReport4topAPsByAirtime5TableData() SCIAirtimeUtilizationReport4topAPsByAirtime5TableData {
	m := make(SCIAirtimeUtilizationReport4topAPsByAirtime5TableData, 0)
	return m
}

// SCIAirtimeUtilizationReport4topAPsByAirtime5TableDataType
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_4_topAPsByAirtime5Table_DataType
type SCIAirtimeUtilizationReport4topAPsByAirtime5TableDataType struct {
	AirtimeBusyAvg *float64 `json:"airtimeBusyAvg,omitempty"`

	AirtimeIdleAvg *float64 `json:"airtimeIdleAvg,omitempty"`

	AirtimeRxAvg *float64 `json:"airtimeRxAvg,omitempty"`

	AirtimeTxAvg *float64 `json:"airtimeTxAvg,omitempty"`

	AirtimeUtilizationAvg *float64 `json:"airtimeUtilizationAvg,omitempty"`

	ApIp *string `json:"apIp,omitempty"`

	ApMac *string `json:"apMac,omitempty"`

	ApName *string `json:"apName,omitempty"`

	ClientCount *float64 `json:"clientCount,omitempty"`

	CtrlMac *string `json:"ctrlMac,omitempty"`

	CtrlName *string `json:"ctrlName,omitempty"`

	CtrlSerial *string `json:"ctrlSerial,omitempty"`

	Index *float64 `json:"index,omitempty"`

	MgmtRxBytes *float64 `json:"mgmtRxBytes,omitempty"`

	MgmtTraffic *float64 `json:"mgmtTraffic,omitempty"`

	MgmtTxBytes *float64 `json:"mgmtTxBytes,omitempty"`

	SessionCount *float64 `json:"sessionCount,omitempty"`

	TotalRxTraffic *float64 `json:"totalRxTraffic,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	TotalTxTraffic *float64 `json:"totalTxTraffic,omitempty"`

	UserRxBytes *float64 `json:"userRxBytes,omitempty"`

	UserTraffic *float64 `json:"userTraffic,omitempty"`

	UserTxBytes *float64 `json:"userTxBytes,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAirtimeUtilizationReport4topAPsByAirtime5TableDataType) UnmarshalJSON(b []byte) error {
	type _SCIAirtimeUtilizationReport4topAPsByAirtime5TableDataType SCIAirtimeUtilizationReport4topAPsByAirtime5TableDataType
	tmpType := new(_SCIAirtimeUtilizationReport4topAPsByAirtime5TableDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "airtimeBusyAvg")
	delete(tmpType.XAdditionalProperties, "airtimeIdleAvg")
	delete(tmpType.XAdditionalProperties, "airtimeRxAvg")
	delete(tmpType.XAdditionalProperties, "airtimeTxAvg")
	delete(tmpType.XAdditionalProperties, "airtimeUtilizationAvg")
	delete(tmpType.XAdditionalProperties, "apIp")
	delete(tmpType.XAdditionalProperties, "apMac")
	delete(tmpType.XAdditionalProperties, "apName")
	delete(tmpType.XAdditionalProperties, "clientCount")
	delete(tmpType.XAdditionalProperties, "ctrlMac")
	delete(tmpType.XAdditionalProperties, "ctrlName")
	delete(tmpType.XAdditionalProperties, "ctrlSerial")
	delete(tmpType.XAdditionalProperties, "index")
	delete(tmpType.XAdditionalProperties, "mgmtRxBytes")
	delete(tmpType.XAdditionalProperties, "mgmtTraffic")
	delete(tmpType.XAdditionalProperties, "mgmtTxBytes")
	delete(tmpType.XAdditionalProperties, "sessionCount")
	delete(tmpType.XAdditionalProperties, "totalRxTraffic")
	delete(tmpType.XAdditionalProperties, "totalTraffic")
	delete(tmpType.XAdditionalProperties, "totalTxTraffic")
	delete(tmpType.XAdditionalProperties, "userRxBytes")
	delete(tmpType.XAdditionalProperties, "userTraffic")
	delete(tmpType.XAdditionalProperties, "userTxBytes")
	*t = SCIAirtimeUtilizationReport4topAPsByAirtime5TableDataType(*tmpType)
	return nil
}

func (t *SCIAirtimeUtilizationReport4topAPsByAirtime5TableDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.AirtimeBusyAvg != nil {
		tmp["airtimeBusyAvg"] = t.AirtimeBusyAvg
	}
	if t.AirtimeIdleAvg != nil {
		tmp["airtimeIdleAvg"] = t.AirtimeIdleAvg
	}
	if t.AirtimeRxAvg != nil {
		tmp["airtimeRxAvg"] = t.AirtimeRxAvg
	}
	if t.AirtimeTxAvg != nil {
		tmp["airtimeTxAvg"] = t.AirtimeTxAvg
	}
	if t.AirtimeUtilizationAvg != nil {
		tmp["airtimeUtilizationAvg"] = t.AirtimeUtilizationAvg
	}
	if t.ApIp != nil {
		tmp["apIp"] = t.ApIp
	}
	if t.ApMac != nil {
		tmp["apMac"] = t.ApMac
	}
	if t.ApName != nil {
		tmp["apName"] = t.ApName
	}
	if t.ClientCount != nil {
		tmp["clientCount"] = t.ClientCount
	}
	if t.CtrlMac != nil {
		tmp["ctrlMac"] = t.CtrlMac
	}
	if t.CtrlName != nil {
		tmp["ctrlName"] = t.CtrlName
	}
	if t.CtrlSerial != nil {
		tmp["ctrlSerial"] = t.CtrlSerial
	}
	if t.Index != nil {
		tmp["index"] = t.Index
	}
	if t.MgmtRxBytes != nil {
		tmp["mgmtRxBytes"] = t.MgmtRxBytes
	}
	if t.MgmtTraffic != nil {
		tmp["mgmtTraffic"] = t.MgmtTraffic
	}
	if t.MgmtTxBytes != nil {
		tmp["mgmtTxBytes"] = t.MgmtTxBytes
	}
	if t.SessionCount != nil {
		tmp["sessionCount"] = t.SessionCount
	}
	if t.TotalRxTraffic != nil {
		tmp["totalRxTraffic"] = t.TotalRxTraffic
	}
	if t.TotalTraffic != nil {
		tmp["totalTraffic"] = t.TotalTraffic
	}
	if t.TotalTxTraffic != nil {
		tmp["totalTxTraffic"] = t.TotalTxTraffic
	}
	if t.UserRxBytes != nil {
		tmp["userRxBytes"] = t.UserRxBytes
	}
	if t.UserTraffic != nil {
		tmp["userTraffic"] = t.UserTraffic
	}
	if t.UserTxBytes != nil {
		tmp["userTxBytes"] = t.UserTxBytes
	}
	return json.Marshal(tmp)
}

func NewSCIAirtimeUtilizationReport4topAPsByAirtime5TableDataType() *SCIAirtimeUtilizationReport4topAPsByAirtime5TableDataType {
	m := new(SCIAirtimeUtilizationReport4topAPsByAirtime5TableDataType)
	return m
}

// SCIAirtimeUtilizationReport4topAPsByAirtime5TableMetaData
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_4_topAPsByAirtime5Table_MetaData
type SCIAirtimeUtilizationReport4topAPsByAirtime5TableMetaData struct {
	MaxValues *SCIAirtimeUtilizationReport4topAPsByAirtime5TableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAirtimeUtilizationReport4topAPsByAirtime5TableMetaData) UnmarshalJSON(b []byte) error {
	type _SCIAirtimeUtilizationReport4topAPsByAirtime5TableMetaData SCIAirtimeUtilizationReport4topAPsByAirtime5TableMetaData
	tmpType := new(_SCIAirtimeUtilizationReport4topAPsByAirtime5TableMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "maxValues")
	*t = SCIAirtimeUtilizationReport4topAPsByAirtime5TableMetaData(*tmpType)
	return nil
}

func (t *SCIAirtimeUtilizationReport4topAPsByAirtime5TableMetaData) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.MaxValues != nil {
		tmp["maxValues"] = t.MaxValues
	}
	return json.Marshal(tmp)
}

func NewSCIAirtimeUtilizationReport4topAPsByAirtime5TableMetaData() *SCIAirtimeUtilizationReport4topAPsByAirtime5TableMetaData {
	m := new(SCIAirtimeUtilizationReport4topAPsByAirtime5TableMetaData)
	return m
}

// SCIAirtimeUtilizationReport4topAPsByAirtime5TableMetaDataMaxValuesType
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_4_topAPsByAirtime5Table_MetaDataMaxValuesType
type SCIAirtimeUtilizationReport4topAPsByAirtime5TableMetaDataMaxValuesType struct {
	MgmtRxBytes *float64 `json:"mgmtRxBytes,omitempty"`

	MgmtTraffic *float64 `json:"mgmtTraffic,omitempty"`

	MgmtTxBytes *float64 `json:"mgmtTxBytes,omitempty"`

	TotalRxTraffic *float64 `json:"totalRxTraffic,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	TotalTxTraffic *float64 `json:"totalTxTraffic,omitempty"`

	UserRxBytes *float64 `json:"userRxBytes,omitempty"`

	UserTraffic *float64 `json:"userTraffic,omitempty"`

	UserTxBytes *float64 `json:"userTxBytes,omitempty"`
}

func NewSCIAirtimeUtilizationReport4topAPsByAirtime5TableMetaDataMaxValuesType() *SCIAirtimeUtilizationReport4topAPsByAirtime5TableMetaDataMaxValuesType {
	m := new(SCIAirtimeUtilizationReport4topAPsByAirtime5TableMetaDataMaxValuesType)
	return m
}

// SCIAirtimeUtilizationReport5trendChartData
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_5_trendChart_Data
type SCIAirtimeUtilizationReport5trendChartData [][]*SCIAirtimeUtilizationReport5trendChartDataTypeType

func MakeSCIAirtimeUtilizationReport5trendChartData() SCIAirtimeUtilizationReport5trendChartData {
	m := make(SCIAirtimeUtilizationReport5trendChartData, 0)
	return m
}

// SCIAirtimeUtilizationReport5trendChartDataType
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_5_trendChart_DataType
type SCIAirtimeUtilizationReport5trendChartDataType []*SCIAirtimeUtilizationReport5trendChartDataTypeType

func MakeSCIAirtimeUtilizationReport5trendChartDataType() SCIAirtimeUtilizationReport5trendChartDataType {
	m := make(SCIAirtimeUtilizationReport5trendChartDataType, 0)
	return m
}

// SCIAirtimeUtilizationReport5trendChartDataTypeType
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_5_trendChart_DataTypeType
type SCIAirtimeUtilizationReport5trendChartDataTypeType struct {
	AirtimeBusyAvg *float64 `json:"airtimeBusyAvg,omitempty"`

	AirtimeIdleAvg *float64 `json:"airtimeIdleAvg,omitempty"`

	AirtimeRxAvg *float64 `json:"airtimeRxAvg,omitempty"`

	AirtimeTxAvg *float64 `json:"airtimeTxAvg,omitempty"`

	AirtimeUtilizationAvg *float64 `json:"airtimeUtilizationAvg,omitempty"`

	End *string `json:"end,omitempty"`

	Start *string `json:"start,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAirtimeUtilizationReport5trendChartDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCIAirtimeUtilizationReport5trendChartDataTypeType SCIAirtimeUtilizationReport5trendChartDataTypeType
	tmpType := new(_SCIAirtimeUtilizationReport5trendChartDataTypeType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "airtimeBusyAvg")
	delete(tmpType.XAdditionalProperties, "airtimeIdleAvg")
	delete(tmpType.XAdditionalProperties, "airtimeRxAvg")
	delete(tmpType.XAdditionalProperties, "airtimeTxAvg")
	delete(tmpType.XAdditionalProperties, "airtimeUtilizationAvg")
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "start")
	*t = SCIAirtimeUtilizationReport5trendChartDataTypeType(*tmpType)
	return nil
}

func (t *SCIAirtimeUtilizationReport5trendChartDataTypeType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.AirtimeBusyAvg != nil {
		tmp["airtimeBusyAvg"] = t.AirtimeBusyAvg
	}
	if t.AirtimeIdleAvg != nil {
		tmp["airtimeIdleAvg"] = t.AirtimeIdleAvg
	}
	if t.AirtimeRxAvg != nil {
		tmp["airtimeRxAvg"] = t.AirtimeRxAvg
	}
	if t.AirtimeTxAvg != nil {
		tmp["airtimeTxAvg"] = t.AirtimeTxAvg
	}
	if t.AirtimeUtilizationAvg != nil {
		tmp["airtimeUtilizationAvg"] = t.AirtimeUtilizationAvg
	}
	if t.End != nil {
		tmp["end"] = t.End
	}
	if t.Start != nil {
		tmp["start"] = t.Start
	}
	return json.Marshal(tmp)
}

func NewSCIAirtimeUtilizationReport5trendChartDataTypeType() *SCIAirtimeUtilizationReport5trendChartDataTypeType {
	m := new(SCIAirtimeUtilizationReport5trendChartDataTypeType)
	return m
}

// SCIAirtimeUtilizationReport6trendTableData
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_6_trendTable_Data
type SCIAirtimeUtilizationReport6trendTableData []*SCIAirtimeUtilizationReport6trendTableDataType

func MakeSCIAirtimeUtilizationReport6trendTableData() SCIAirtimeUtilizationReport6trendTableData {
	m := make(SCIAirtimeUtilizationReport6trendTableData, 0)
	return m
}

// SCIAirtimeUtilizationReport6trendTableDataType
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_6_trendTable_DataType
type SCIAirtimeUtilizationReport6trendTableDataType struct {
	AirtimeBusyAvg24 *float64 `json:"airtimeBusyAvg_2-4,omitempty"`

	AirtimeBusyAvg5 *float64 `json:"airtimeBusyAvg_5,omitempty"`

	AirtimeIdleAvg24 *float64 `json:"airtimeIdleAvg_2-4,omitempty"`

	AirtimeIdleAvg5 *float64 `json:"airtimeIdleAvg_5,omitempty"`

	AirtimeRxAvg24 *float64 `json:"airtimeRxAvg_2-4,omitempty"`

	AirtimeRxAvg5 *float64 `json:"airtimeRxAvg_5,omitempty"`

	AirtimeTxAvg24 *float64 `json:"airtimeTxAvg_2-4,omitempty"`

	AirtimeTxAvg5 *float64 `json:"airtimeTxAvg_5,omitempty"`

	AirtimeUtilizationAvg24 *float64 `json:"airtimeUtilizationAvg_2-4,omitempty"`

	AirtimeUtilizationAvg5 *float64 `json:"airtimeUtilizationAvg_5,omitempty"`

	ClientCount interface{} `json:"clientCount,omitempty"`

	End *string `json:"end,omitempty"`

	MgmtRxBytes *float64 `json:"mgmtRxBytes,omitempty"`

	MgmtTraffic *float64 `json:"mgmtTraffic,omitempty"`

	MgmtTxBytes *float64 `json:"mgmtTxBytes,omitempty"`

	SessionCount interface{} `json:"sessionCount,omitempty"`

	Start *string `json:"start,omitempty"`

	TotalRxTraffic *float64 `json:"totalRxTraffic,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	TotalTxTraffic *float64 `json:"totalTxTraffic,omitempty"`

	UserRxBytes *float64 `json:"userRxBytes,omitempty"`

	UserTraffic *float64 `json:"userTraffic,omitempty"`

	UserTxBytes *float64 `json:"userTxBytes,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAirtimeUtilizationReport6trendTableDataType) UnmarshalJSON(b []byte) error {
	type _SCIAirtimeUtilizationReport6trendTableDataType SCIAirtimeUtilizationReport6trendTableDataType
	tmpType := new(_SCIAirtimeUtilizationReport6trendTableDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "airtimeBusyAvg_2-4")
	delete(tmpType.XAdditionalProperties, "airtimeBusyAvg_5")
	delete(tmpType.XAdditionalProperties, "airtimeIdleAvg_2-4")
	delete(tmpType.XAdditionalProperties, "airtimeIdleAvg_5")
	delete(tmpType.XAdditionalProperties, "airtimeRxAvg_2-4")
	delete(tmpType.XAdditionalProperties, "airtimeRxAvg_5")
	delete(tmpType.XAdditionalProperties, "airtimeTxAvg_2-4")
	delete(tmpType.XAdditionalProperties, "airtimeTxAvg_5")
	delete(tmpType.XAdditionalProperties, "airtimeUtilizationAvg_2-4")
	delete(tmpType.XAdditionalProperties, "airtimeUtilizationAvg_5")
	delete(tmpType.XAdditionalProperties, "clientCount")
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "mgmtRxBytes")
	delete(tmpType.XAdditionalProperties, "mgmtTraffic")
	delete(tmpType.XAdditionalProperties, "mgmtTxBytes")
	delete(tmpType.XAdditionalProperties, "sessionCount")
	delete(tmpType.XAdditionalProperties, "start")
	delete(tmpType.XAdditionalProperties, "totalRxTraffic")
	delete(tmpType.XAdditionalProperties, "totalTraffic")
	delete(tmpType.XAdditionalProperties, "totalTxTraffic")
	delete(tmpType.XAdditionalProperties, "userRxBytes")
	delete(tmpType.XAdditionalProperties, "userTraffic")
	delete(tmpType.XAdditionalProperties, "userTxBytes")
	*t = SCIAirtimeUtilizationReport6trendTableDataType(*tmpType)
	return nil
}

func (t *SCIAirtimeUtilizationReport6trendTableDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.AirtimeBusyAvg24 != nil {
		tmp["airtimeBusyAvg_2-4"] = t.AirtimeBusyAvg24
	}
	if t.AirtimeBusyAvg5 != nil {
		tmp["airtimeBusyAvg_5"] = t.AirtimeBusyAvg5
	}
	if t.AirtimeIdleAvg24 != nil {
		tmp["airtimeIdleAvg_2-4"] = t.AirtimeIdleAvg24
	}
	if t.AirtimeIdleAvg5 != nil {
		tmp["airtimeIdleAvg_5"] = t.AirtimeIdleAvg5
	}
	if t.AirtimeRxAvg24 != nil {
		tmp["airtimeRxAvg_2-4"] = t.AirtimeRxAvg24
	}
	if t.AirtimeRxAvg5 != nil {
		tmp["airtimeRxAvg_5"] = t.AirtimeRxAvg5
	}
	if t.AirtimeTxAvg24 != nil {
		tmp["airtimeTxAvg_2-4"] = t.AirtimeTxAvg24
	}
	if t.AirtimeTxAvg5 != nil {
		tmp["airtimeTxAvg_5"] = t.AirtimeTxAvg5
	}
	if t.AirtimeUtilizationAvg24 != nil {
		tmp["airtimeUtilizationAvg_2-4"] = t.AirtimeUtilizationAvg24
	}
	if t.AirtimeUtilizationAvg5 != nil {
		tmp["airtimeUtilizationAvg_5"] = t.AirtimeUtilizationAvg5
	}
	if t.ClientCount != nil {
		tmp["clientCount"] = t.ClientCount
	}
	if t.End != nil {
		tmp["end"] = t.End
	}
	if t.MgmtRxBytes != nil {
		tmp["mgmtRxBytes"] = t.MgmtRxBytes
	}
	if t.MgmtTraffic != nil {
		tmp["mgmtTraffic"] = t.MgmtTraffic
	}
	if t.MgmtTxBytes != nil {
		tmp["mgmtTxBytes"] = t.MgmtTxBytes
	}
	if t.SessionCount != nil {
		tmp["sessionCount"] = t.SessionCount
	}
	if t.Start != nil {
		tmp["start"] = t.Start
	}
	if t.TotalRxTraffic != nil {
		tmp["totalRxTraffic"] = t.TotalRxTraffic
	}
	if t.TotalTraffic != nil {
		tmp["totalTraffic"] = t.TotalTraffic
	}
	if t.TotalTxTraffic != nil {
		tmp["totalTxTraffic"] = t.TotalTxTraffic
	}
	if t.UserRxBytes != nil {
		tmp["userRxBytes"] = t.UserRxBytes
	}
	if t.UserTraffic != nil {
		tmp["userTraffic"] = t.UserTraffic
	}
	if t.UserTxBytes != nil {
		tmp["userTxBytes"] = t.UserTxBytes
	}
	return json.Marshal(tmp)
}

func NewSCIAirtimeUtilizationReport6trendTableDataType() *SCIAirtimeUtilizationReport6trendTableDataType {
	m := new(SCIAirtimeUtilizationReport6trendTableDataType)
	return m
}

// SCIAirtimeUtilizationReport6trendTableMetaData
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_6_trendTable_MetaData
type SCIAirtimeUtilizationReport6trendTableMetaData struct {
	MaxValues *SCIAirtimeUtilizationReport6trendTableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIAirtimeUtilizationReport6trendTableMetaData) UnmarshalJSON(b []byte) error {
	type _SCIAirtimeUtilizationReport6trendTableMetaData SCIAirtimeUtilizationReport6trendTableMetaData
	tmpType := new(_SCIAirtimeUtilizationReport6trendTableMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "maxValues")
	*t = SCIAirtimeUtilizationReport6trendTableMetaData(*tmpType)
	return nil
}

func (t *SCIAirtimeUtilizationReport6trendTableMetaData) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.MaxValues != nil {
		tmp["maxValues"] = t.MaxValues
	}
	return json.Marshal(tmp)
}

func NewSCIAirtimeUtilizationReport6trendTableMetaData() *SCIAirtimeUtilizationReport6trendTableMetaData {
	m := new(SCIAirtimeUtilizationReport6trendTableMetaData)
	return m
}

// SCIAirtimeUtilizationReport6trendTableMetaDataMaxValuesType
//
// Definition: AirtimeUtilizationReport_AirtimeUtilizationReport_6_trendTable_MetaDataMaxValuesType
type SCIAirtimeUtilizationReport6trendTableMetaDataMaxValuesType struct {
	MgmtRxBytes *float64 `json:"mgmtRxBytes,omitempty"`

	MgmtTraffic *float64 `json:"mgmtTraffic,omitempty"`

	MgmtTxBytes *float64 `json:"mgmtTxBytes,omitempty"`

	TotalRxTraffic *float64 `json:"totalRxTraffic,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	TotalTxTraffic *float64 `json:"totalTxTraffic,omitempty"`

	UserRxBytes *float64 `json:"userRxBytes,omitempty"`

	UserTraffic *float64 `json:"userTraffic,omitempty"`

	UserTxBytes *float64 `json:"userTxBytes,omitempty"`
}

func NewSCIAirtimeUtilizationReport6trendTableMetaDataMaxValuesType() *SCIAirtimeUtilizationReport6trendTableMetaDataMaxValuesType {
	m := new(SCIAirtimeUtilizationReport6trendTableMetaDataMaxValuesType)
	return m
}

// ReportAirtimeUtilizationReport1Overview
//
// Airtime Utilization Report - Overview
//
// Operation ID: report_AirtimeUtilizationReport_1_overview
// Operation path: /reports/5/sections/1/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAirtimeUtilizationReportService) ReportAirtimeUtilizationReport1Overview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAirtimeUtilizationReport1overview200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportAirtimeUtilizationReport1overview200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportAirtimeUtilizationReport1overview200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportAirtimeUtilizationReport1Overview, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportAirtimeUtilizationReport1overview200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportAirtimeUtilizationReport1overview200ResponseTypeAPIResponse), err
}

// ReportAirtimeUtilizationReport2TopChart
//
// Airtime Utilization Report - Top 10 APs by Airtime Utilization
//
// Operation ID: report_AirtimeUtilizationReport_2_topChart
// Operation path: /reports/5/sections/2/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAirtimeUtilizationReportService) ReportAirtimeUtilizationReport2TopChart(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAirtimeUtilizationReport2topChart200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportAirtimeUtilizationReport2topChart200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportAirtimeUtilizationReport2topChart200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportAirtimeUtilizationReport2TopChart, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportAirtimeUtilizationReport2topChart200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportAirtimeUtilizationReport2topChart200ResponseTypeAPIResponse), err
}

// ReportAirtimeUtilizationReport3TopAPsByAirtime24Table
//
// Airtime Utilization Report - Top APs by Airtime Utilization for 2.4 GHz
//
// Operation ID: report_AirtimeUtilizationReport_3_topAPsByAirtime24Table
// Operation path: /reports/5/sections/3/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAirtimeUtilizationReportService) ReportAirtimeUtilizationReport3TopAPsByAirtime24Table(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAirtimeUtilizationReport3topAPsByAirtime24Table200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportAirtimeUtilizationReport3topAPsByAirtime24Table200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportAirtimeUtilizationReport3topAPsByAirtime24Table200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportAirtimeUtilizationReport3TopAPsByAirtime24Table, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportAirtimeUtilizationReport3topAPsByAirtime24Table200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportAirtimeUtilizationReport3topAPsByAirtime24Table200ResponseTypeAPIResponse), err
}

// ReportAirtimeUtilizationReport4TopAPsByAirtime5Table
//
// Airtime Utilization Report - Top APs by Airtime Utilization for 5 GHz
//
// Operation ID: report_AirtimeUtilizationReport_4_topAPsByAirtime5Table
// Operation path: /reports/5/sections/4/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAirtimeUtilizationReportService) ReportAirtimeUtilizationReport4TopAPsByAirtime5Table(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAirtimeUtilizationReport4topAPsByAirtime5Table200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportAirtimeUtilizationReport4topAPsByAirtime5Table200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportAirtimeUtilizationReport4topAPsByAirtime5Table200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportAirtimeUtilizationReport4TopAPsByAirtime5Table, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportAirtimeUtilizationReport4topAPsByAirtime5Table200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportAirtimeUtilizationReport4topAPsByAirtime5Table200ResponseTypeAPIResponse), err
}

// ReportAirtimeUtilizationReport5TrendChart
//
// Airtime Utilization Report - Airtime Utilization Trend
//
// Operation ID: report_AirtimeUtilizationReport_5_trendChart
// Operation path: /reports/5/sections/5/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAirtimeUtilizationReportService) ReportAirtimeUtilizationReport5TrendChart(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAirtimeUtilizationReport5trendChart200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportAirtimeUtilizationReport5trendChart200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportAirtimeUtilizationReport5trendChart200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportAirtimeUtilizationReport5TrendChart, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportAirtimeUtilizationReport5trendChart200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportAirtimeUtilizationReport5trendChart200ResponseTypeAPIResponse), err
}

// ReportAirtimeUtilizationReport6TrendTable
//
// Airtime Utilization Report - Airtime Utilization Over Time
//
// Operation ID: report_AirtimeUtilizationReport_6_trendTable
// Operation path: /reports/5/sections/6/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAirtimeUtilizationReportService) ReportAirtimeUtilizationReport6TrendTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAirtimeUtilizationReport6trendTable200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportAirtimeUtilizationReport6trendTable200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportAirtimeUtilizationReport6trendTable200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportAirtimeUtilizationReport6TrendTable, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportAirtimeUtilizationReport6trendTable200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportAirtimeUtilizationReport6trendTable200ResponseTypeAPIResponse), err
}
