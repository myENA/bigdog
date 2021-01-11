package bigdog

// API Version: 1.0.0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SCIWLANsReportService struct {
	apiClient *SCIClient
}

func NewSCIWLANsReportService(c *SCIClient) *SCIWLANsReportService {
	s := new(SCIWLANsReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIWLANsReportService() *SCIWLANsReportService {
	return NewSCIWLANsReportService(ss.apiClient)
}

// SCIWLANsReport35overviewData
//
// Definition: WLANsReport_WLANsReport_35_overview_Data
type SCIWLANsReport35overviewData []*SCIWLANsReport35overviewDataType

func MakeSCIWLANsReport35overviewData() SCIWLANsReport35overviewData {
	m := make(SCIWLANsReport35overviewData, 0)
	return m
}

// SCIWLANsReport35overviewDataType
//
// Definition: WLANsReport_WLANsReport_35_overview_DataType
type SCIWLANsReport35overviewDataType struct {
	TotalAddedClientCount *float64 `json:"totalAddedClientCount,omitempty"`

	TotalAddedSsids *float64 `json:"totalAddedSsids,omitempty"`

	TotalClientCount *float64 `json:"totalClientCount,omitempty"`

	TotalDeletedClientCount *float64 `json:"totalDeletedClientCount,omitempty"`

	TotalDeletedSsids *float64 `json:"totalDeletedSsids,omitempty"`

	TotalSsids *float64 `json:"totalSsids,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIWLANsReport35overviewDataType) UnmarshalJSON(b []byte) error {
	type _SCIWLANsReport35overviewDataType SCIWLANsReport35overviewDataType
	tmpType := new(_SCIWLANsReport35overviewDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "totalAddedClientCount")
	delete(tmpType.XAdditionalProperties, "totalAddedSsids")
	delete(tmpType.XAdditionalProperties, "totalClientCount")
	delete(tmpType.XAdditionalProperties, "totalDeletedClientCount")
	delete(tmpType.XAdditionalProperties, "totalDeletedSsids")
	delete(tmpType.XAdditionalProperties, "totalSsids")
	*t = SCIWLANsReport35overviewDataType(*tmpType)
	return nil
}

func (t *SCIWLANsReport35overviewDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.TotalAddedClientCount != nil {
		tmp["totalAddedClientCount"] = t.TotalAddedClientCount
	}
	if t.TotalAddedSsids != nil {
		tmp["totalAddedSsids"] = t.TotalAddedSsids
	}
	if t.TotalClientCount != nil {
		tmp["totalClientCount"] = t.TotalClientCount
	}
	if t.TotalDeletedClientCount != nil {
		tmp["totalDeletedClientCount"] = t.TotalDeletedClientCount
	}
	if t.TotalDeletedSsids != nil {
		tmp["totalDeletedSsids"] = t.TotalDeletedSsids
	}
	if t.TotalSsids != nil {
		tmp["totalSsids"] = t.TotalSsids
	}
	return json.Marshal(tmp)
}

func NewSCIWLANsReport35overviewDataType() *SCIWLANsReport35overviewDataType {
	m := new(SCIWLANsReport35overviewDataType)
	return m
}

// SCIWLANsReport37activeSsidsTrendData
//
// Definition: WLANsReport_WLANsReport_37_activeSsidsTrend_Data
type SCIWLANsReport37activeSsidsTrendData [][]*SCIWLANsReport37activeSsidsTrendDataTypeType

func MakeSCIWLANsReport37activeSsidsTrendData() SCIWLANsReport37activeSsidsTrendData {
	m := make(SCIWLANsReport37activeSsidsTrendData, 0)
	return m
}

// SCIWLANsReport37activeSsidsTrendDataType
//
// Definition: WLANsReport_WLANsReport_37_activeSsidsTrend_DataType
type SCIWLANsReport37activeSsidsTrendDataType []*SCIWLANsReport37activeSsidsTrendDataTypeType

func MakeSCIWLANsReport37activeSsidsTrendDataType() SCIWLANsReport37activeSsidsTrendDataType {
	m := make(SCIWLANsReport37activeSsidsTrendDataType, 0)
	return m
}

// SCIWLANsReport37activeSsidsTrendDataTypeType
//
// Definition: WLANsReport_WLANsReport_37_activeSsidsTrend_DataTypeType
type SCIWLANsReport37activeSsidsTrendDataTypeType struct {
	End *string `json:"end,omitempty"`

	Ssid *float64 `json:"ssid,omitempty"`

	Start *string `json:"start,omitempty"`

	Traffic *float64 `json:"traffic,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIWLANsReport37activeSsidsTrendDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCIWLANsReport37activeSsidsTrendDataTypeType SCIWLANsReport37activeSsidsTrendDataTypeType
	tmpType := new(_SCIWLANsReport37activeSsidsTrendDataTypeType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "ssid")
	delete(tmpType.XAdditionalProperties, "start")
	delete(tmpType.XAdditionalProperties, "traffic")
	*t = SCIWLANsReport37activeSsidsTrendDataTypeType(*tmpType)
	return nil
}

func (t *SCIWLANsReport37activeSsidsTrendDataTypeType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.End != nil {
		tmp["end"] = t.End
	}
	if t.Ssid != nil {
		tmp["ssid"] = t.Ssid
	}
	if t.Start != nil {
		tmp["start"] = t.Start
	}
	if t.Traffic != nil {
		tmp["traffic"] = t.Traffic
	}
	return json.Marshal(tmp)
}

func NewSCIWLANsReport37activeSsidsTrendDataTypeType() *SCIWLANsReport37activeSsidsTrendDataTypeType {
	m := new(SCIWLANsReport37activeSsidsTrendDataTypeType)
	return m
}

// SCIWLANsReport38top10SsidsByClientCountData
//
// Definition: WLANsReport_WLANsReport_38_top10SsidsByClientCount_Data
type SCIWLANsReport38top10SsidsByClientCountData [][]*SCIWLANsReport38top10SsidsByClientCountDataTypeType

func MakeSCIWLANsReport38top10SsidsByClientCountData() SCIWLANsReport38top10SsidsByClientCountData {
	m := make(SCIWLANsReport38top10SsidsByClientCountData, 0)
	return m
}

// SCIWLANsReport38top10SsidsByClientCountDataType
//
// Definition: WLANsReport_WLANsReport_38_top10SsidsByClientCount_DataType
type SCIWLANsReport38top10SsidsByClientCountDataType []*SCIWLANsReport38top10SsidsByClientCountDataTypeType

func MakeSCIWLANsReport38top10SsidsByClientCountDataType() SCIWLANsReport38top10SsidsByClientCountDataType {
	m := make(SCIWLANsReport38top10SsidsByClientCountDataType, 0)
	return m
}

// SCIWLANsReport38top10SsidsByClientCountDataTypeType
//
// Definition: WLANsReport_WLANsReport_38_top10SsidsByClientCount_DataTypeType
type SCIWLANsReport38top10SsidsByClientCountDataTypeType struct {
	End *string `json:"end,omitempty"`

	Ssid *string `json:"ssid,omitempty"`

	Start *string `json:"start,omitempty"`

	UniqueUsers *float64 `json:"uniqueUsers,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIWLANsReport38top10SsidsByClientCountDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCIWLANsReport38top10SsidsByClientCountDataTypeType SCIWLANsReport38top10SsidsByClientCountDataTypeType
	tmpType := new(_SCIWLANsReport38top10SsidsByClientCountDataTypeType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "ssid")
	delete(tmpType.XAdditionalProperties, "start")
	delete(tmpType.XAdditionalProperties, "uniqueUsers")
	*t = SCIWLANsReport38top10SsidsByClientCountDataTypeType(*tmpType)
	return nil
}

func (t *SCIWLANsReport38top10SsidsByClientCountDataTypeType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.End != nil {
		tmp["end"] = t.End
	}
	if t.Ssid != nil {
		tmp["ssid"] = t.Ssid
	}
	if t.Start != nil {
		tmp["start"] = t.Start
	}
	if t.UniqueUsers != nil {
		tmp["uniqueUsers"] = t.UniqueUsers
	}
	return json.Marshal(tmp)
}

func NewSCIWLANsReport38top10SsidsByClientCountDataTypeType() *SCIWLANsReport38top10SsidsByClientCountDataTypeType {
	m := new(SCIWLANsReport38top10SsidsByClientCountDataTypeType)
	return m
}

// SCIWLANsReport38top10SsidsByClientCountMetaData
//
// Definition: WLANsReport_WLANsReport_38_top10SsidsByClientCount_MetaData
type SCIWLANsReport38top10SsidsByClientCountMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`

	TotalClients *float64 `json:"totalClients,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIWLANsReport38top10SsidsByClientCountMetaData) UnmarshalJSON(b []byte) error {
	type _SCIWLANsReport38top10SsidsByClientCountMetaData SCIWLANsReport38top10SsidsByClientCountMetaData
	tmpType := new(_SCIWLANsReport38top10SsidsByClientCountMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "colorKeys")
	delete(tmpType.XAdditionalProperties, "totalClients")
	*t = SCIWLANsReport38top10SsidsByClientCountMetaData(*tmpType)
	return nil
}

func (t *SCIWLANsReport38top10SsidsByClientCountMetaData) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.ColorKeys != nil {
		tmp["colorKeys"] = t.ColorKeys
	}
	if t.TotalClients != nil {
		tmp["totalClients"] = t.TotalClients
	}
	return json.Marshal(tmp)
}

func NewSCIWLANsReport38top10SsidsByClientCountMetaData() *SCIWLANsReport38top10SsidsByClientCountMetaData {
	m := new(SCIWLANsReport38top10SsidsByClientCountMetaData)
	return m
}

// SCIWLANsReport39ssidChangesOverTimeData
//
// Definition: WLANsReport_WLANsReport_39_ssidChangesOverTime_Data
type SCIWLANsReport39ssidChangesOverTimeData []*SCIWLANsReport39ssidChangesOverTimeDataType

func MakeSCIWLANsReport39ssidChangesOverTimeData() SCIWLANsReport39ssidChangesOverTimeData {
	m := make(SCIWLANsReport39ssidChangesOverTimeData, 0)
	return m
}

// SCIWLANsReport39ssidChangesOverTimeDataType
//
// Definition: WLANsReport_WLANsReport_39_ssidChangesOverTime_DataType
type SCIWLANsReport39ssidChangesOverTimeDataType struct {
	ClientCount *float64 `json:"clientCount,omitempty"`

	Ssid *string `json:"ssid,omitempty"`

	Status *string `json:"status,omitempty"`

	Time *float64 `json:"time,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIWLANsReport39ssidChangesOverTimeDataType) UnmarshalJSON(b []byte) error {
	type _SCIWLANsReport39ssidChangesOverTimeDataType SCIWLANsReport39ssidChangesOverTimeDataType
	tmpType := new(_SCIWLANsReport39ssidChangesOverTimeDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "clientCount")
	delete(tmpType.XAdditionalProperties, "ssid")
	delete(tmpType.XAdditionalProperties, "status")
	delete(tmpType.XAdditionalProperties, "time")
	*t = SCIWLANsReport39ssidChangesOverTimeDataType(*tmpType)
	return nil
}

func (t *SCIWLANsReport39ssidChangesOverTimeDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.ClientCount != nil {
		tmp["clientCount"] = t.ClientCount
	}
	if t.Ssid != nil {
		tmp["ssid"] = t.Ssid
	}
	if t.Status != nil {
		tmp["status"] = t.Status
	}
	if t.Time != nil {
		tmp["time"] = t.Time
	}
	return json.Marshal(tmp)
}

func NewSCIWLANsReport39ssidChangesOverTimeDataType() *SCIWLANsReport39ssidChangesOverTimeDataType {
	m := new(SCIWLANsReport39ssidChangesOverTimeDataType)
	return m
}

// SCIWLANsReport40topSsidsByTrafficTableData
//
// Definition: WLANsReport_WLANsReport_40_topSsidsByTrafficTable_Data
type SCIWLANsReport40topSsidsByTrafficTableData []*SCIWLANsReport40topSsidsByTrafficTableDataType

func MakeSCIWLANsReport40topSsidsByTrafficTableData() SCIWLANsReport40topSsidsByTrafficTableData {
	m := make(SCIWLANsReport40topSsidsByTrafficTableData, 0)
	return m
}

// SCIWLANsReport40topSsidsByTrafficTableDataType
//
// Definition: WLANsReport_WLANsReport_40_topSsidsByTrafficTable_DataType
type SCIWLANsReport40topSsidsByTrafficTableDataType struct {
	ApCount *float64 `json:"apCount,omitempty"`

	ClientCount *float64 `json:"clientCount,omitempty"`

	Index *float64 `json:"index,omitempty"`

	MgmtRxBytes *float64 `json:"mgmtRxBytes,omitempty"`

	MgmtTraffic *float64 `json:"mgmtTraffic,omitempty"`

	MgmtTxBytes *float64 `json:"mgmtTxBytes,omitempty"`

	SessionCount *float64 `json:"sessionCount,omitempty"`

	Ssid *string `json:"ssid,omitempty"`

	TotalRxTraffic *float64 `json:"totalRxTraffic,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	TotalTxTraffic *float64 `json:"totalTxTraffic,omitempty"`

	UserRxBytes *float64 `json:"userRxBytes,omitempty"`

	UserTraffic *float64 `json:"userTraffic,omitempty"`

	UserTxBytes *float64 `json:"userTxBytes,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIWLANsReport40topSsidsByTrafficTableDataType) UnmarshalJSON(b []byte) error {
	type _SCIWLANsReport40topSsidsByTrafficTableDataType SCIWLANsReport40topSsidsByTrafficTableDataType
	tmpType := new(_SCIWLANsReport40topSsidsByTrafficTableDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "apCount")
	delete(tmpType.XAdditionalProperties, "clientCount")
	delete(tmpType.XAdditionalProperties, "index")
	delete(tmpType.XAdditionalProperties, "mgmtRxBytes")
	delete(tmpType.XAdditionalProperties, "mgmtTraffic")
	delete(tmpType.XAdditionalProperties, "mgmtTxBytes")
	delete(tmpType.XAdditionalProperties, "sessionCount")
	delete(tmpType.XAdditionalProperties, "ssid")
	delete(tmpType.XAdditionalProperties, "totalRxTraffic")
	delete(tmpType.XAdditionalProperties, "totalTraffic")
	delete(tmpType.XAdditionalProperties, "totalTxTraffic")
	delete(tmpType.XAdditionalProperties, "userRxBytes")
	delete(tmpType.XAdditionalProperties, "userTraffic")
	delete(tmpType.XAdditionalProperties, "userTxBytes")
	*t = SCIWLANsReport40topSsidsByTrafficTableDataType(*tmpType)
	return nil
}

func (t *SCIWLANsReport40topSsidsByTrafficTableDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.ApCount != nil {
		tmp["apCount"] = t.ApCount
	}
	if t.ClientCount != nil {
		tmp["clientCount"] = t.ClientCount
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
	if t.Ssid != nil {
		tmp["ssid"] = t.Ssid
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

func NewSCIWLANsReport40topSsidsByTrafficTableDataType() *SCIWLANsReport40topSsidsByTrafficTableDataType {
	m := new(SCIWLANsReport40topSsidsByTrafficTableDataType)
	return m
}

// SCIWLANsReport40topSsidsByTrafficTableMetaData
//
// Definition: WLANsReport_WLANsReport_40_topSsidsByTrafficTable_MetaData
type SCIWLANsReport40topSsidsByTrafficTableMetaData struct {
	MaxValues *SCIWLANsReport40topSsidsByTrafficTableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	Name *string `json:"name,omitempty"`

	Percentage *float64 `json:"percentage,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	Traffic *float64 `json:"traffic,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIWLANsReport40topSsidsByTrafficTableMetaData) UnmarshalJSON(b []byte) error {
	type _SCIWLANsReport40topSsidsByTrafficTableMetaData SCIWLANsReport40topSsidsByTrafficTableMetaData
	tmpType := new(_SCIWLANsReport40topSsidsByTrafficTableMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "maxValues")
	delete(tmpType.XAdditionalProperties, "name")
	delete(tmpType.XAdditionalProperties, "percentage")
	delete(tmpType.XAdditionalProperties, "totalTraffic")
	delete(tmpType.XAdditionalProperties, "traffic")
	*t = SCIWLANsReport40topSsidsByTrafficTableMetaData(*tmpType)
	return nil
}

func (t *SCIWLANsReport40topSsidsByTrafficTableMetaData) MarshalJSON() ([]byte, error) {
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
	if t.Name != nil {
		tmp["name"] = t.Name
	}
	if t.Percentage != nil {
		tmp["percentage"] = t.Percentage
	}
	if t.TotalTraffic != nil {
		tmp["totalTraffic"] = t.TotalTraffic
	}
	if t.Traffic != nil {
		tmp["traffic"] = t.Traffic
	}
	return json.Marshal(tmp)
}

func NewSCIWLANsReport40topSsidsByTrafficTableMetaData() *SCIWLANsReport40topSsidsByTrafficTableMetaData {
	m := new(SCIWLANsReport40topSsidsByTrafficTableMetaData)
	return m
}

// SCIWLANsReport40topSsidsByTrafficTableMetaDataMaxValuesType
//
// Definition: WLANsReport_WLANsReport_40_topSsidsByTrafficTable_MetaDataMaxValuesType
type SCIWLANsReport40topSsidsByTrafficTableMetaDataMaxValuesType struct {
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

func NewSCIWLANsReport40topSsidsByTrafficTableMetaDataMaxValuesType() *SCIWLANsReport40topSsidsByTrafficTableMetaDataMaxValuesType {
	m := new(SCIWLANsReport40topSsidsByTrafficTableMetaDataMaxValuesType)
	return m
}

// SCIWLANsReport41topSsidsByClientsTableData
//
// Definition: WLANsReport_WLANsReport_41_topSsidsByClientsTable_Data
type SCIWLANsReport41topSsidsByClientsTableData []*SCIWLANsReport41topSsidsByClientsTableDataType

func MakeSCIWLANsReport41topSsidsByClientsTableData() SCIWLANsReport41topSsidsByClientsTableData {
	m := make(SCIWLANsReport41topSsidsByClientsTableData, 0)
	return m
}

// SCIWLANsReport41topSsidsByClientsTableDataType
//
// Definition: WLANsReport_WLANsReport_41_topSsidsByClientsTable_DataType
type SCIWLANsReport41topSsidsByClientsTableDataType struct {
	ApCount *float64 `json:"apCount,omitempty"`

	ClientCount *float64 `json:"clientCount,omitempty"`

	Index *float64 `json:"index,omitempty"`

	MgmtRxBytes *float64 `json:"mgmtRxBytes,omitempty"`

	MgmtTraffic *float64 `json:"mgmtTraffic,omitempty"`

	MgmtTxBytes *float64 `json:"mgmtTxBytes,omitempty"`

	SessionCount *float64 `json:"sessionCount,omitempty"`

	Ssid *string `json:"ssid,omitempty"`

	TotalRxTraffic *float64 `json:"totalRxTraffic,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	TotalTxTraffic *float64 `json:"totalTxTraffic,omitempty"`

	UserRxBytes *float64 `json:"userRxBytes,omitempty"`

	UserTraffic *float64 `json:"userTraffic,omitempty"`

	UserTxBytes *float64 `json:"userTxBytes,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIWLANsReport41topSsidsByClientsTableDataType) UnmarshalJSON(b []byte) error {
	type _SCIWLANsReport41topSsidsByClientsTableDataType SCIWLANsReport41topSsidsByClientsTableDataType
	tmpType := new(_SCIWLANsReport41topSsidsByClientsTableDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "apCount")
	delete(tmpType.XAdditionalProperties, "clientCount")
	delete(tmpType.XAdditionalProperties, "index")
	delete(tmpType.XAdditionalProperties, "mgmtRxBytes")
	delete(tmpType.XAdditionalProperties, "mgmtTraffic")
	delete(tmpType.XAdditionalProperties, "mgmtTxBytes")
	delete(tmpType.XAdditionalProperties, "sessionCount")
	delete(tmpType.XAdditionalProperties, "ssid")
	delete(tmpType.XAdditionalProperties, "totalRxTraffic")
	delete(tmpType.XAdditionalProperties, "totalTraffic")
	delete(tmpType.XAdditionalProperties, "totalTxTraffic")
	delete(tmpType.XAdditionalProperties, "userRxBytes")
	delete(tmpType.XAdditionalProperties, "userTraffic")
	delete(tmpType.XAdditionalProperties, "userTxBytes")
	*t = SCIWLANsReport41topSsidsByClientsTableDataType(*tmpType)
	return nil
}

func (t *SCIWLANsReport41topSsidsByClientsTableDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.ApCount != nil {
		tmp["apCount"] = t.ApCount
	}
	if t.ClientCount != nil {
		tmp["clientCount"] = t.ClientCount
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
	if t.Ssid != nil {
		tmp["ssid"] = t.Ssid
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

func NewSCIWLANsReport41topSsidsByClientsTableDataType() *SCIWLANsReport41topSsidsByClientsTableDataType {
	m := new(SCIWLANsReport41topSsidsByClientsTableDataType)
	return m
}

// SCIWLANsReport41topSsidsByClientsTableMetaData
//
// Definition: WLANsReport_WLANsReport_41_topSsidsByClientsTable_MetaData
type SCIWLANsReport41topSsidsByClientsTableMetaData struct {
	MaxValues *SCIWLANsReport41topSsidsByClientsTableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	Name *string `json:"name,omitempty"`

	Percentage *float64 `json:"percentage,omitempty"`

	TotalTraffic *float64 `json:"totalTraffic,omitempty"`

	Traffic *float64 `json:"traffic,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIWLANsReport41topSsidsByClientsTableMetaData) UnmarshalJSON(b []byte) error {
	type _SCIWLANsReport41topSsidsByClientsTableMetaData SCIWLANsReport41topSsidsByClientsTableMetaData
	tmpType := new(_SCIWLANsReport41topSsidsByClientsTableMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "maxValues")
	delete(tmpType.XAdditionalProperties, "name")
	delete(tmpType.XAdditionalProperties, "percentage")
	delete(tmpType.XAdditionalProperties, "totalTraffic")
	delete(tmpType.XAdditionalProperties, "traffic")
	*t = SCIWLANsReport41topSsidsByClientsTableMetaData(*tmpType)
	return nil
}

func (t *SCIWLANsReport41topSsidsByClientsTableMetaData) MarshalJSON() ([]byte, error) {
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
	if t.Name != nil {
		tmp["name"] = t.Name
	}
	if t.Percentage != nil {
		tmp["percentage"] = t.Percentage
	}
	if t.TotalTraffic != nil {
		tmp["totalTraffic"] = t.TotalTraffic
	}
	if t.Traffic != nil {
		tmp["traffic"] = t.Traffic
	}
	return json.Marshal(tmp)
}

func NewSCIWLANsReport41topSsidsByClientsTableMetaData() *SCIWLANsReport41topSsidsByClientsTableMetaData {
	m := new(SCIWLANsReport41topSsidsByClientsTableMetaData)
	return m
}

// SCIWLANsReport41topSsidsByClientsTableMetaDataMaxValuesType
//
// Definition: WLANsReport_WLANsReport_41_topSsidsByClientsTable_MetaDataMaxValuesType
type SCIWLANsReport41topSsidsByClientsTableMetaDataMaxValuesType struct {
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

func NewSCIWLANsReport41topSsidsByClientsTableMetaDataMaxValuesType() *SCIWLANsReport41topSsidsByClientsTableMetaDataMaxValuesType {
	m := new(SCIWLANsReport41topSsidsByClientsTableMetaDataMaxValuesType)
	return m
}

// ReportWLANsReport35Overview
//
// WLANs Report - Overview
//
// Operation ID: report_WLANsReport_35_overview
// Operation path: /reports/4/sections/35/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWLANsReportService) ReportWLANsReport35Overview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWLANsReport35overview200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportWLANsReport35overview200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportWLANsReport35overview200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportWLANsReport35Overview, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportWLANsReport35overview200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportWLANsReport35overview200ResponseTypeAPIResponse), err
}

// ReportWLANsReport36Top10SsidsByTraffic
//
// WLANs Report - Top SSIDs by Traffic
//
// Operation ID: report_WLANsReport_36_top10SsidsByTraffic
// Operation path: /reports/4/sections/36/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWLANsReportService) ReportWLANsReport36Top10SsidsByTraffic(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWLANsReport36top10SsidsByTraffic200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportWLANsReport36top10SsidsByTraffic200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportWLANsReport36top10SsidsByTraffic200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportWLANsReport36Top10SsidsByTraffic, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportWLANsReport36top10SsidsByTraffic200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportWLANsReport36top10SsidsByTraffic200ResponseTypeAPIResponse), err
}

// ReportWLANsReport37ActiveSsidsTrend
//
// WLANs Report - Active SSIDs Trend
//
// Operation ID: report_WLANsReport_37_activeSsidsTrend
// Operation path: /reports/4/sections/37/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWLANsReportService) ReportWLANsReport37ActiveSsidsTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWLANsReport37activeSsidsTrend200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportWLANsReport37activeSsidsTrend200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportWLANsReport37activeSsidsTrend200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportWLANsReport37ActiveSsidsTrend, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportWLANsReport37activeSsidsTrend200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportWLANsReport37activeSsidsTrend200ResponseTypeAPIResponse), err
}

// ReportWLANsReport38Top10SsidsByClientCount
//
// WLANs Report - Top SSIDs by Client Count
//
// Operation ID: report_WLANsReport_38_top10SsidsByClientCount
// Operation path: /reports/4/sections/38/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWLANsReportService) ReportWLANsReport38Top10SsidsByClientCount(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWLANsReport38top10SsidsByClientCount200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportWLANsReport38top10SsidsByClientCount200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportWLANsReport38top10SsidsByClientCount200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportWLANsReport38Top10SsidsByClientCount, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportWLANsReport38top10SsidsByClientCount200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportWLANsReport38top10SsidsByClientCount200ResponseTypeAPIResponse), err
}

// ReportWLANsReport39SsidChangesOverTime
//
// WLANs Report - SSID Changes Over Time
//
// Operation ID: report_WLANsReport_39_ssidChangesOverTime
// Operation path: /reports/4/sections/39/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWLANsReportService) ReportWLANsReport39SsidChangesOverTime(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWLANsReport39ssidChangesOverTime200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportWLANsReport39ssidChangesOverTime200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportWLANsReport39ssidChangesOverTime200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportWLANsReport39SsidChangesOverTime, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportWLANsReport39ssidChangesOverTime200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportWLANsReport39ssidChangesOverTime200ResponseTypeAPIResponse), err
}

// ReportWLANsReport40TopSsidsByTrafficTable
//
// WLANs Report - Top SSIDs by Traffic
//
// Operation ID: report_WLANsReport_40_topSsidsByTrafficTable
// Operation path: /reports/4/sections/40/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWLANsReportService) ReportWLANsReport40TopSsidsByTrafficTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWLANsReport40topSsidsByTrafficTable200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportWLANsReport40topSsidsByTrafficTable200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportWLANsReport40topSsidsByTrafficTable200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportWLANsReport40TopSsidsByTrafficTable, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportWLANsReport40topSsidsByTrafficTable200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportWLANsReport40topSsidsByTrafficTable200ResponseTypeAPIResponse), err
}

// ReportWLANsReport41TopSsidsByClientsTable
//
// WLANs Report - Top SSIDs by Client Count
//
// Operation ID: report_WLANsReport_41_topSsidsByClientsTable
// Operation path: /reports/4/sections/41/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIWLANsReportService) ReportWLANsReport41TopSsidsByClientsTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportWLANsReport41topSsidsByClientsTable200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSCIReportWLANsReport41topSsidsByClientsTable200ResponseTypeAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SCIReportWLANsReport41topSsidsByClientsTable200ResponseTypeAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSCIReportWLANsReport41TopSsidsByClientsTable, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SCIReportWLANsReport41topSsidsByClientsTable200ResponseTypeAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportWLANsReport41topSsidsByClientsTable200ResponseTypeAPIResponse), err
}
