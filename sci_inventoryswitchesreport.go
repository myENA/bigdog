package bigdog

// API Version: 1.0.0

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type SCIInventorySwitchesReportService struct {
	apiClient *SCIClient
}

func NewSCIInventorySwitchesReportService(c *SCIClient) *SCIInventorySwitchesReportService {
	s := new(SCIInventorySwitchesReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIInventorySwitchesReportService() *SCIInventorySwitchesReportService {
	return NewSCIInventorySwitchesReportService(ss.apiClient)
}

// SCIInventorySwitchesReport113overviewData
//
// Definition: InventorySwitchesReport_InventorySwitchesReport_113_overview_Data
type SCIInventorySwitchesReport113overviewData []*SCIInventorySwitchesReport113overviewDataType

func MakeSCIInventorySwitchesReport113overviewData() SCIInventorySwitchesReport113overviewData {
	m := make(SCIInventorySwitchesReport113overviewData, 0)
	return m
}

// SCIInventorySwitchesReport113overviewDataType
//
// Definition: InventorySwitchesReport_InventorySwitchesReport_113_overview_DataType
type SCIInventorySwitchesReport113overviewDataType struct {
	DownPortCount *float64 `json:"downPortCount,omitempty"`

	OfflineSwitchCount *float64 `json:"offlineSwitchCount,omitempty"`

	OnlineSwitchCount *float64 `json:"onlineSwitchCount,omitempty"`

	OtherSwitchCount *float64 `json:"otherSwitchCount,omitempty"`

	PortCount *float64 `json:"portCount,omitempty"`

	SwitchCount *float64 `json:"switchCount,omitempty"`

	SwitchUnitCount *float64 `json:"switchUnitCount,omitempty"`

	UpPortCount *float64 `json:"upPortCount,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIInventorySwitchesReport113overviewDataType) UnmarshalJSON(b []byte) error {
	type _SCIInventorySwitchesReport113overviewDataType SCIInventorySwitchesReport113overviewDataType
	tmpType := new(_SCIInventorySwitchesReport113overviewDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "downPortCount")
	delete(tmpType.XAdditionalProperties, "offlineSwitchCount")
	delete(tmpType.XAdditionalProperties, "onlineSwitchCount")
	delete(tmpType.XAdditionalProperties, "otherSwitchCount")
	delete(tmpType.XAdditionalProperties, "portCount")
	delete(tmpType.XAdditionalProperties, "switchCount")
	delete(tmpType.XAdditionalProperties, "switchUnitCount")
	delete(tmpType.XAdditionalProperties, "upPortCount")
	*t = SCIInventorySwitchesReport113overviewDataType(*tmpType)
	return nil
}

func (t *SCIInventorySwitchesReport113overviewDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.DownPortCount != nil {
		tmp["downPortCount"] = t.DownPortCount
	}
	if t.OfflineSwitchCount != nil {
		tmp["offlineSwitchCount"] = t.OfflineSwitchCount
	}
	if t.OnlineSwitchCount != nil {
		tmp["onlineSwitchCount"] = t.OnlineSwitchCount
	}
	if t.OtherSwitchCount != nil {
		tmp["otherSwitchCount"] = t.OtherSwitchCount
	}
	if t.PortCount != nil {
		tmp["portCount"] = t.PortCount
	}
	if t.SwitchCount != nil {
		tmp["switchCount"] = t.SwitchCount
	}
	if t.SwitchUnitCount != nil {
		tmp["switchUnitCount"] = t.SwitchUnitCount
	}
	if t.UpPortCount != nil {
		tmp["upPortCount"] = t.UpPortCount
	}
	return json.Marshal(tmp)
}

func NewSCIInventorySwitchesReport113overviewDataType() *SCIInventorySwitchesReport113overviewDataType {
	m := new(SCIInventorySwitchesReport113overviewDataType)
	return m
}

// SCIInventorySwitchesReport116switchCountTrendData
//
// Definition: InventorySwitchesReport_InventorySwitchesReport_116_switchCountTrend_Data
type SCIInventorySwitchesReport116switchCountTrendData []*SCIInventorySwitchesReport116switchCountTrendDataType

func MakeSCIInventorySwitchesReport116switchCountTrendData() SCIInventorySwitchesReport116switchCountTrendData {
	m := make(SCIInventorySwitchesReport116switchCountTrendData, 0)
	return m
}

// SCIInventorySwitchesReport116switchCountTrendDataType
//
// Definition: InventorySwitchesReport_InventorySwitchesReport_116_switchCountTrend_DataType
type SCIInventorySwitchesReport116switchCountTrendDataType struct {
	Count *float64 `json:"count,omitempty"`

	End *string `json:"end,omitempty"`

	OfflineCount *float64 `json:"offlineCount,omitempty"`

	OnlineCount *float64 `json:"onlineCount,omitempty"`

	Start *string `json:"start,omitempty"`

	SwitchUnitCount *float64 `json:"switchUnitCount,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIInventorySwitchesReport116switchCountTrendDataType) UnmarshalJSON(b []byte) error {
	type _SCIInventorySwitchesReport116switchCountTrendDataType SCIInventorySwitchesReport116switchCountTrendDataType
	tmpType := new(_SCIInventorySwitchesReport116switchCountTrendDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "count")
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "offlineCount")
	delete(tmpType.XAdditionalProperties, "onlineCount")
	delete(tmpType.XAdditionalProperties, "start")
	delete(tmpType.XAdditionalProperties, "switchUnitCount")
	*t = SCIInventorySwitchesReport116switchCountTrendDataType(*tmpType)
	return nil
}

func (t *SCIInventorySwitchesReport116switchCountTrendDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Count != nil {
		tmp["count"] = t.Count
	}
	if t.End != nil {
		tmp["end"] = t.End
	}
	if t.OfflineCount != nil {
		tmp["offlineCount"] = t.OfflineCount
	}
	if t.OnlineCount != nil {
		tmp["onlineCount"] = t.OnlineCount
	}
	if t.Start != nil {
		tmp["start"] = t.Start
	}
	if t.SwitchUnitCount != nil {
		tmp["switchUnitCount"] = t.SwitchUnitCount
	}
	return json.Marshal(tmp)
}

func NewSCIInventorySwitchesReport116switchCountTrendDataType() *SCIInventorySwitchesReport116switchCountTrendDataType {
	m := new(SCIInventorySwitchesReport116switchCountTrendDataType)
	return m
}

// SCIInventorySwitchesReport117top10SwitchVersionChartData
//
// Definition: InventorySwitchesReport_InventorySwitchesReport_117_top10SwitchVersionChart_Data
type SCIInventorySwitchesReport117top10SwitchVersionChartData [][]*SCIInventorySwitchesReport117top10SwitchVersionChartDataTypeType

func MakeSCIInventorySwitchesReport117top10SwitchVersionChartData() SCIInventorySwitchesReport117top10SwitchVersionChartData {
	m := make(SCIInventorySwitchesReport117top10SwitchVersionChartData, 0)
	return m
}

// SCIInventorySwitchesReport117top10SwitchVersionChartDataType
//
// Definition: InventorySwitchesReport_InventorySwitchesReport_117_top10SwitchVersionChart_DataType
type SCIInventorySwitchesReport117top10SwitchVersionChartDataType []*SCIInventorySwitchesReport117top10SwitchVersionChartDataTypeType

func MakeSCIInventorySwitchesReport117top10SwitchVersionChartDataType() SCIInventorySwitchesReport117top10SwitchVersionChartDataType {
	m := make(SCIInventorySwitchesReport117top10SwitchVersionChartDataType, 0)
	return m
}

// SCIInventorySwitchesReport117top10SwitchVersionChartDataTypeType
//
// Definition: InventorySwitchesReport_InventorySwitchesReport_117_top10SwitchVersionChart_DataTypeType
type SCIInventorySwitchesReport117top10SwitchVersionChartDataTypeType struct {
	End *string `json:"end,omitempty"`

	Start *string `json:"start,omitempty"`

	SwitchCount *float64 `json:"switchCount,omitempty"`

	SwitchFirmware *string `json:"switchFirmware,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIInventorySwitchesReport117top10SwitchVersionChartDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCIInventorySwitchesReport117top10SwitchVersionChartDataTypeType SCIInventorySwitchesReport117top10SwitchVersionChartDataTypeType
	tmpType := new(_SCIInventorySwitchesReport117top10SwitchVersionChartDataTypeType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "start")
	delete(tmpType.XAdditionalProperties, "switchCount")
	delete(tmpType.XAdditionalProperties, "switchFirmware")
	*t = SCIInventorySwitchesReport117top10SwitchVersionChartDataTypeType(*tmpType)
	return nil
}

func (t *SCIInventorySwitchesReport117top10SwitchVersionChartDataTypeType) MarshalJSON() ([]byte, error) {
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
	if t.Start != nil {
		tmp["start"] = t.Start
	}
	if t.SwitchCount != nil {
		tmp["switchCount"] = t.SwitchCount
	}
	if t.SwitchFirmware != nil {
		tmp["switchFirmware"] = t.SwitchFirmware
	}
	return json.Marshal(tmp)
}

func NewSCIInventorySwitchesReport117top10SwitchVersionChartDataTypeType() *SCIInventorySwitchesReport117top10SwitchVersionChartDataTypeType {
	m := new(SCIInventorySwitchesReport117top10SwitchVersionChartDataTypeType)
	return m
}

// SCIInventorySwitchesReport117top10SwitchVersionChartMetaData
//
// Definition: InventorySwitchesReport_InventorySwitchesReport_117_top10SwitchVersionChart_MetaData
type SCIInventorySwitchesReport117top10SwitchVersionChartMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIInventorySwitchesReport117top10SwitchVersionChartMetaData) UnmarshalJSON(b []byte) error {
	type _SCIInventorySwitchesReport117top10SwitchVersionChartMetaData SCIInventorySwitchesReport117top10SwitchVersionChartMetaData
	tmpType := new(_SCIInventorySwitchesReport117top10SwitchVersionChartMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "colorKeys")
	*t = SCIInventorySwitchesReport117top10SwitchVersionChartMetaData(*tmpType)
	return nil
}

func (t *SCIInventorySwitchesReport117top10SwitchVersionChartMetaData) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(tmp)
}

func NewSCIInventorySwitchesReport117top10SwitchVersionChartMetaData() *SCIInventorySwitchesReport117top10SwitchVersionChartMetaData {
	m := new(SCIInventorySwitchesReport117top10SwitchVersionChartMetaData)
	return m
}

// SCIInventorySwitchesReport118topSwitchVersionsData
//
// Definition: InventorySwitchesReport_InventorySwitchesReport_118_topSwitchVersions_Data
type SCIInventorySwitchesReport118topSwitchVersionsData []*SCIInventorySwitchesReport118topSwitchVersionsDataType

func MakeSCIInventorySwitchesReport118topSwitchVersionsData() SCIInventorySwitchesReport118topSwitchVersionsData {
	m := make(SCIInventorySwitchesReport118topSwitchVersionsData, 0)
	return m
}

// SCIInventorySwitchesReport118topSwitchVersionsDataType
//
// Definition: InventorySwitchesReport_InventorySwitchesReport_118_topSwitchVersions_DataType
type SCIInventorySwitchesReport118topSwitchVersionsDataType struct {
	Index *float64 `json:"index,omitempty"`

	SwitchCount *float64 `json:"switchCount,omitempty"`

	SwitchFirmware *string `json:"switchFirmware,omitempty"`

	SwitchPercent *float64 `json:"switchPercent,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIInventorySwitchesReport118topSwitchVersionsDataType) UnmarshalJSON(b []byte) error {
	type _SCIInventorySwitchesReport118topSwitchVersionsDataType SCIInventorySwitchesReport118topSwitchVersionsDataType
	tmpType := new(_SCIInventorySwitchesReport118topSwitchVersionsDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "index")
	delete(tmpType.XAdditionalProperties, "switchCount")
	delete(tmpType.XAdditionalProperties, "switchFirmware")
	delete(tmpType.XAdditionalProperties, "switchPercent")
	*t = SCIInventorySwitchesReport118topSwitchVersionsDataType(*tmpType)
	return nil
}

func (t *SCIInventorySwitchesReport118topSwitchVersionsDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Index != nil {
		tmp["index"] = t.Index
	}
	if t.SwitchCount != nil {
		tmp["switchCount"] = t.SwitchCount
	}
	if t.SwitchFirmware != nil {
		tmp["switchFirmware"] = t.SwitchFirmware
	}
	if t.SwitchPercent != nil {
		tmp["switchPercent"] = t.SwitchPercent
	}
	return json.Marshal(tmp)
}

func NewSCIInventorySwitchesReport118topSwitchVersionsDataType() *SCIInventorySwitchesReport118topSwitchVersionsDataType {
	m := new(SCIInventorySwitchesReport118topSwitchVersionsDataType)
	return m
}

// SCIInventorySwitchesReport121topSwitchModelsChartData
//
// Definition: InventorySwitchesReport_InventorySwitchesReport_121_topSwitchModelsChart_Data
type SCIInventorySwitchesReport121topSwitchModelsChartData [][]*SCIInventorySwitchesReport121topSwitchModelsChartDataTypeType

func MakeSCIInventorySwitchesReport121topSwitchModelsChartData() SCIInventorySwitchesReport121topSwitchModelsChartData {
	m := make(SCIInventorySwitchesReport121topSwitchModelsChartData, 0)
	return m
}

// SCIInventorySwitchesReport121topSwitchModelsChartDataType
//
// Definition: InventorySwitchesReport_InventorySwitchesReport_121_topSwitchModelsChart_DataType
type SCIInventorySwitchesReport121topSwitchModelsChartDataType []*SCIInventorySwitchesReport121topSwitchModelsChartDataTypeType

func MakeSCIInventorySwitchesReport121topSwitchModelsChartDataType() SCIInventorySwitchesReport121topSwitchModelsChartDataType {
	m := make(SCIInventorySwitchesReport121topSwitchModelsChartDataType, 0)
	return m
}

// SCIInventorySwitchesReport121topSwitchModelsChartDataTypeType
//
// Definition: InventorySwitchesReport_InventorySwitchesReport_121_topSwitchModelsChart_DataTypeType
type SCIInventorySwitchesReport121topSwitchModelsChartDataTypeType struct {
	End *string `json:"end,omitempty"`

	Start *string `json:"start,omitempty"`

	SwitchCount *float64 `json:"switchCount,omitempty"`

	SwitchModel *string `json:"switchModel,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIInventorySwitchesReport121topSwitchModelsChartDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCIInventorySwitchesReport121topSwitchModelsChartDataTypeType SCIInventorySwitchesReport121topSwitchModelsChartDataTypeType
	tmpType := new(_SCIInventorySwitchesReport121topSwitchModelsChartDataTypeType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "start")
	delete(tmpType.XAdditionalProperties, "switchCount")
	delete(tmpType.XAdditionalProperties, "switchModel")
	*t = SCIInventorySwitchesReport121topSwitchModelsChartDataTypeType(*tmpType)
	return nil
}

func (t *SCIInventorySwitchesReport121topSwitchModelsChartDataTypeType) MarshalJSON() ([]byte, error) {
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
	if t.Start != nil {
		tmp["start"] = t.Start
	}
	if t.SwitchCount != nil {
		tmp["switchCount"] = t.SwitchCount
	}
	if t.SwitchModel != nil {
		tmp["switchModel"] = t.SwitchModel
	}
	return json.Marshal(tmp)
}

func NewSCIInventorySwitchesReport121topSwitchModelsChartDataTypeType() *SCIInventorySwitchesReport121topSwitchModelsChartDataTypeType {
	m := new(SCIInventorySwitchesReport121topSwitchModelsChartDataTypeType)
	return m
}

// SCIInventorySwitchesReport121topSwitchModelsChartMetaData
//
// Definition: InventorySwitchesReport_InventorySwitchesReport_121_topSwitchModelsChart_MetaData
type SCIInventorySwitchesReport121topSwitchModelsChartMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIInventorySwitchesReport121topSwitchModelsChartMetaData) UnmarshalJSON(b []byte) error {
	type _SCIInventorySwitchesReport121topSwitchModelsChartMetaData SCIInventorySwitchesReport121topSwitchModelsChartMetaData
	tmpType := new(_SCIInventorySwitchesReport121topSwitchModelsChartMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "colorKeys")
	*t = SCIInventorySwitchesReport121topSwitchModelsChartMetaData(*tmpType)
	return nil
}

func (t *SCIInventorySwitchesReport121topSwitchModelsChartMetaData) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(tmp)
}

func NewSCIInventorySwitchesReport121topSwitchModelsChartMetaData() *SCIInventorySwitchesReport121topSwitchModelsChartMetaData {
	m := new(SCIInventorySwitchesReport121topSwitchModelsChartMetaData)
	return m
}

// SCIInventorySwitchesReport122topSwitchModelsData
//
// Definition: InventorySwitchesReport_InventorySwitchesReport_122_topSwitchModels_Data
type SCIInventorySwitchesReport122topSwitchModelsData []*SCIInventorySwitchesReport122topSwitchModelsDataType

func MakeSCIInventorySwitchesReport122topSwitchModelsData() SCIInventorySwitchesReport122topSwitchModelsData {
	m := make(SCIInventorySwitchesReport122topSwitchModelsData, 0)
	return m
}

// SCIInventorySwitchesReport122topSwitchModelsDataType
//
// Definition: InventorySwitchesReport_InventorySwitchesReport_122_topSwitchModels_DataType
type SCIInventorySwitchesReport122topSwitchModelsDataType struct {
	Index *float64 `json:"index,omitempty"`

	SwitchCount *float64 `json:"switchCount,omitempty"`

	SwitchModel *string `json:"switchModel,omitempty"`

	SwitchPercent *float64 `json:"switchPercent,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIInventorySwitchesReport122topSwitchModelsDataType) UnmarshalJSON(b []byte) error {
	type _SCIInventorySwitchesReport122topSwitchModelsDataType SCIInventorySwitchesReport122topSwitchModelsDataType
	tmpType := new(_SCIInventorySwitchesReport122topSwitchModelsDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "index")
	delete(tmpType.XAdditionalProperties, "switchCount")
	delete(tmpType.XAdditionalProperties, "switchModel")
	delete(tmpType.XAdditionalProperties, "switchPercent")
	*t = SCIInventorySwitchesReport122topSwitchModelsDataType(*tmpType)
	return nil
}

func (t *SCIInventorySwitchesReport122topSwitchModelsDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Index != nil {
		tmp["index"] = t.Index
	}
	if t.SwitchCount != nil {
		tmp["switchCount"] = t.SwitchCount
	}
	if t.SwitchModel != nil {
		tmp["switchModel"] = t.SwitchModel
	}
	if t.SwitchPercent != nil {
		tmp["switchPercent"] = t.SwitchPercent
	}
	return json.Marshal(tmp)
}

func NewSCIInventorySwitchesReport122topSwitchModelsDataType() *SCIInventorySwitchesReport122topSwitchModelsDataType {
	m := new(SCIInventorySwitchesReport122topSwitchModelsDataType)
	return m
}

// SCIInventorySwitchesReport132portStatusTrendData
//
// Definition: InventorySwitchesReport_InventorySwitchesReport_132_portStatusTrend_Data
type SCIInventorySwitchesReport132portStatusTrendData [][]*SCIInventorySwitchesReport132portStatusTrendDataTypeType

func MakeSCIInventorySwitchesReport132portStatusTrendData() SCIInventorySwitchesReport132portStatusTrendData {
	m := make(SCIInventorySwitchesReport132portStatusTrendData, 0)
	return m
}

// SCIInventorySwitchesReport132portStatusTrendDataType
//
// Definition: InventorySwitchesReport_InventorySwitchesReport_132_portStatusTrend_DataType
type SCIInventorySwitchesReport132portStatusTrendDataType []*SCIInventorySwitchesReport132portStatusTrendDataTypeType

func MakeSCIInventorySwitchesReport132portStatusTrendDataType() SCIInventorySwitchesReport132portStatusTrendDataType {
	m := make(SCIInventorySwitchesReport132portStatusTrendDataType, 0)
	return m
}

// SCIInventorySwitchesReport132portStatusTrendDataTypeType
//
// Definition: InventorySwitchesReport_InventorySwitchesReport_132_portStatusTrend_DataTypeType
type SCIInventorySwitchesReport132portStatusTrendDataTypeType struct {
	Down *float64 `json:"Down,omitempty"`

	End *string `json:"end,omitempty"`

	PortCount *float64 `json:"portCount,omitempty"`

	PortStatus *string `json:"portStatus,omitempty"`

	Start *string `json:"start,omitempty"`

	Up *float64 `json:"Up,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIInventorySwitchesReport132portStatusTrendDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCIInventorySwitchesReport132portStatusTrendDataTypeType SCIInventorySwitchesReport132portStatusTrendDataTypeType
	tmpType := new(_SCIInventorySwitchesReport132portStatusTrendDataTypeType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "Down")
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "portCount")
	delete(tmpType.XAdditionalProperties, "portStatus")
	delete(tmpType.XAdditionalProperties, "start")
	delete(tmpType.XAdditionalProperties, "Up")
	*t = SCIInventorySwitchesReport132portStatusTrendDataTypeType(*tmpType)
	return nil
}

func (t *SCIInventorySwitchesReport132portStatusTrendDataTypeType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Down != nil {
		tmp["Down"] = t.Down
	}
	if t.End != nil {
		tmp["end"] = t.End
	}
	if t.PortCount != nil {
		tmp["portCount"] = t.PortCount
	}
	if t.PortStatus != nil {
		tmp["portStatus"] = t.PortStatus
	}
	if t.Start != nil {
		tmp["start"] = t.Start
	}
	if t.Up != nil {
		tmp["Up"] = t.Up
	}
	return json.Marshal(tmp)
}

func NewSCIInventorySwitchesReport132portStatusTrendDataTypeType() *SCIInventorySwitchesReport132portStatusTrendDataTypeType {
	m := new(SCIInventorySwitchesReport132portStatusTrendDataTypeType)
	return m
}

// SCIInventorySwitchesReport132portStatusTrendMetaData
//
// Definition: InventorySwitchesReport_InventorySwitchesReport_132_portStatusTrend_MetaData
type SCIInventorySwitchesReport132portStatusTrendMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIInventorySwitchesReport132portStatusTrendMetaData) UnmarshalJSON(b []byte) error {
	type _SCIInventorySwitchesReport132portStatusTrendMetaData SCIInventorySwitchesReport132portStatusTrendMetaData
	tmpType := new(_SCIInventorySwitchesReport132portStatusTrendMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "colorKeys")
	*t = SCIInventorySwitchesReport132portStatusTrendMetaData(*tmpType)
	return nil
}

func (t *SCIInventorySwitchesReport132portStatusTrendMetaData) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(tmp)
}

func NewSCIInventorySwitchesReport132portStatusTrendMetaData() *SCIInventorySwitchesReport132portStatusTrendMetaData {
	m := new(SCIInventorySwitchesReport132portStatusTrendMetaData)
	return m
}

// ReportInventorySwitchesReport113Overview
//
// Inventory - Switches Report - Overview
//
// Operation ID: report_InventorySwitchesReport_113_overview
// Operation path: /reports/16/sections/113/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventorySwitchesReportService) ReportInventorySwitchesReport113Overview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventorySwitchesReport113overview200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportInventorySwitchesReport113overview200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportInventorySwitchesReport113Overview, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportInventorySwitchesReport113overview200ResponseTypeAPIResponse), err
}

// ReportInventorySwitchesReport116SwitchCountTrend
//
// Inventory - Switches Report - Switch Count Trend
//
// Operation ID: report_InventorySwitchesReport_116_switchCountTrend
// Operation path: /reports/16/sections/116/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventorySwitchesReportService) ReportInventorySwitchesReport116SwitchCountTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventorySwitchesReport116switchCountTrend200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportInventorySwitchesReport116switchCountTrend200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportInventorySwitchesReport116SwitchCountTrend, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportInventorySwitchesReport116switchCountTrend200ResponseTypeAPIResponse), err
}

// ReportInventorySwitchesReport117Top10SwitchVersionChart
//
// Inventory - Switches Report - Top Switch Software Versions
//
// Operation ID: report_InventorySwitchesReport_117_top10SwitchVersionChart
// Operation path: /reports/16/sections/117/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventorySwitchesReportService) ReportInventorySwitchesReport117Top10SwitchVersionChart(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventorySwitchesReport117top10SwitchVersionChart200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportInventorySwitchesReport117top10SwitchVersionChart200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportInventorySwitchesReport117Top10SwitchVersionChart, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportInventorySwitchesReport117top10SwitchVersionChart200ResponseTypeAPIResponse), err
}

// ReportInventorySwitchesReport118TopSwitchVersions
//
// Inventory - Switches Report - Top Switch Software Versions
//
// Operation ID: report_InventorySwitchesReport_118_topSwitchVersions
// Operation path: /reports/16/sections/118/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventorySwitchesReportService) ReportInventorySwitchesReport118TopSwitchVersions(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventorySwitchesReport118topSwitchVersions200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportInventorySwitchesReport118topSwitchVersions200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportInventorySwitchesReport118TopSwitchVersions, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportInventorySwitchesReport118topSwitchVersions200ResponseTypeAPIResponse), err
}

// ReportInventorySwitchesReport121TopSwitchModelsChart
//
// Inventory - Switches Report - Top Switch Models
//
// Operation ID: report_InventorySwitchesReport_121_topSwitchModelsChart
// Operation path: /reports/16/sections/121/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventorySwitchesReportService) ReportInventorySwitchesReport121TopSwitchModelsChart(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventorySwitchesReport121topSwitchModelsChart200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportInventorySwitchesReport121topSwitchModelsChart200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportInventorySwitchesReport121TopSwitchModelsChart, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportInventorySwitchesReport121topSwitchModelsChart200ResponseTypeAPIResponse), err
}

// ReportInventorySwitchesReport122TopSwitchModels
//
// Inventory - Switches Report - Top Switch Models
//
// Operation ID: report_InventorySwitchesReport_122_topSwitchModels
// Operation path: /reports/16/sections/122/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventorySwitchesReportService) ReportInventorySwitchesReport122TopSwitchModels(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventorySwitchesReport122topSwitchModels200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportInventorySwitchesReport122topSwitchModels200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportInventorySwitchesReport122TopSwitchModels, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportInventorySwitchesReport122topSwitchModels200ResponseTypeAPIResponse), err
}

// ReportInventorySwitchesReport132PortStatusTrend
//
// Inventory - Switches Report - Port Status Trends
//
// Operation ID: report_InventorySwitchesReport_132_portStatusTrend
// Operation path: /reports/16/sections/132/data
// Success code: 200 (OK)
//
// Request body:
//	 - body *SCICommonReportQueryBody
func (s *SCIInventorySwitchesReportService) ReportInventorySwitchesReport132PortStatusTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportInventorySwitchesReport132portStatusTrend200ResponseTypeAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSCIReportInventorySwitchesReport132portStatusTrend200ResponseTypeAPIResponse
	)
	req = apiRequestFromPool(APISourceSCI, http.MethodPost, RouteSCIReportInventorySwitchesReport132PortStatusTrend, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SCIReportInventorySwitchesReport132portStatusTrend200ResponseTypeAPIResponse), err
}
