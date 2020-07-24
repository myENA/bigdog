package bigdog

// API Version: 1.0.0

import (
	"encoding/json"
)

// SCISwitchDetailsReport126switchResourceUtilizationData
//
// Definition: SwitchDetailsReport_SwitchDetailsReport_126_switchResourceUtilization_Data
type SCISwitchDetailsReport126switchResourceUtilizationData []*SCISwitchDetailsReport126switchResourceUtilizationDataType

func MakeSCISwitchDetailsReport126switchResourceUtilizationData() SCISwitchDetailsReport126switchResourceUtilizationData {
	m := make(SCISwitchDetailsReport126switchResourceUtilizationData, 0)
	return m
}

// SCISwitchDetailsReport126switchResourceUtilizationDataType
//
// Definition: SwitchDetailsReport_SwitchDetailsReport_126_switchResourceUtilization_DataType
type SCISwitchDetailsReport126switchResourceUtilizationDataType struct {
	CpuPer *float64 `json:"cpuPer,omitempty"`

	End *string `json:"end,omitempty"`

	MemoryPer *float64 `json:"memoryPer,omitempty"`

	PoePer *float64 `json:"poePer,omitempty"`

	Start *string `json:"start,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCISwitchDetailsReport126switchResourceUtilizationDataType) UnmarshalJSON(b []byte) error {
	type _SCISwitchDetailsReport126switchResourceUtilizationDataType SCISwitchDetailsReport126switchResourceUtilizationDataType
	tmpType := new(_SCISwitchDetailsReport126switchResourceUtilizationDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "cpuPer")
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "memoryPer")
	delete(tmpType.XAdditionalProperties, "poePer")
	delete(tmpType.XAdditionalProperties, "start")
	*t = SCISwitchDetailsReport126switchResourceUtilizationDataType(*tmpType)
	return nil
}

func (t *SCISwitchDetailsReport126switchResourceUtilizationDataType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.CpuPer != nil {
		tmp["cpuPer"] = t.CpuPer
	}
	if t.End != nil {
		tmp["end"] = t.End
	}
	if t.MemoryPer != nil {
		tmp["memoryPer"] = t.MemoryPer
	}
	if t.PoePer != nil {
		tmp["poePer"] = t.PoePer
	}
	if t.Start != nil {
		tmp["start"] = t.Start
	}
	return json.Marshal(tmp)
}

func NewSCISwitchDetailsReport126switchResourceUtilizationDataType() *SCISwitchDetailsReport126switchResourceUtilizationDataType {
	m := new(SCISwitchDetailsReport126switchResourceUtilizationDataType)
	return m
}

// SCISwitchDetailsReport129topSwitchPortsByTrafficChartData
//
// Definition: SwitchDetailsReport_SwitchDetailsReport_129_topSwitchPortsByTrafficChart_Data
type SCISwitchDetailsReport129topSwitchPortsByTrafficChartData [][]*SCISwitchDetailsReport129topSwitchPortsByTrafficChartDataTypeType

func MakeSCISwitchDetailsReport129topSwitchPortsByTrafficChartData() SCISwitchDetailsReport129topSwitchPortsByTrafficChartData {
	m := make(SCISwitchDetailsReport129topSwitchPortsByTrafficChartData, 0)
	return m
}

// SCISwitchDetailsReport129topSwitchPortsByTrafficChartDataType
//
// Definition: SwitchDetailsReport_SwitchDetailsReport_129_topSwitchPortsByTrafficChart_DataType
type SCISwitchDetailsReport129topSwitchPortsByTrafficChartDataType []*SCISwitchDetailsReport129topSwitchPortsByTrafficChartDataTypeType

func MakeSCISwitchDetailsReport129topSwitchPortsByTrafficChartDataType() SCISwitchDetailsReport129topSwitchPortsByTrafficChartDataType {
	m := make(SCISwitchDetailsReport129topSwitchPortsByTrafficChartDataType, 0)
	return m
}

// SCISwitchDetailsReport129topSwitchPortsByTrafficChartDataTypeType
//
// Definition: SwitchDetailsReport_SwitchDetailsReport_129_topSwitchPortsByTrafficChart_DataTypeType
type SCISwitchDetailsReport129topSwitchPortsByTrafficChartDataTypeType struct {
	End *string `json:"end,omitempty"`

	PortName *string `json:"portName,omitempty"`

	PortNumber *string `json:"portNumber,omitempty"`

	Start *string `json:"start,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCISwitchDetailsReport129topSwitchPortsByTrafficChartDataTypeType) UnmarshalJSON(b []byte) error {
	type _SCISwitchDetailsReport129topSwitchPortsByTrafficChartDataTypeType SCISwitchDetailsReport129topSwitchPortsByTrafficChartDataTypeType
	tmpType := new(_SCISwitchDetailsReport129topSwitchPortsByTrafficChartDataTypeType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "portName")
	delete(tmpType.XAdditionalProperties, "portNumber")
	delete(tmpType.XAdditionalProperties, "start")
	delete(tmpType.XAdditionalProperties, "totalTraffic")
	*t = SCISwitchDetailsReport129topSwitchPortsByTrafficChartDataTypeType(*tmpType)
	return nil
}

func (t *SCISwitchDetailsReport129topSwitchPortsByTrafficChartDataTypeType) MarshalJSON() ([]byte, error) {
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
	if t.PortName != nil {
		tmp["portName"] = t.PortName
	}
	if t.PortNumber != nil {
		tmp["portNumber"] = t.PortNumber
	}
	if t.Start != nil {
		tmp["start"] = t.Start
	}
	if t.TotalTraffic != nil {
		tmp["totalTraffic"] = t.TotalTraffic
	}
	return json.Marshal(tmp)
}

func NewSCISwitchDetailsReport129topSwitchPortsByTrafficChartDataTypeType() *SCISwitchDetailsReport129topSwitchPortsByTrafficChartDataTypeType {
	m := new(SCISwitchDetailsReport129topSwitchPortsByTrafficChartDataTypeType)
	return m
}

// SCISwitchDetailsReport129topSwitchPortsByTrafficChartMetaData
//
// Definition: SwitchDetailsReport_SwitchDetailsReport_129_topSwitchPortsByTrafficChart_MetaData
type SCISwitchDetailsReport129topSwitchPortsByTrafficChartMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCISwitchDetailsReport129topSwitchPortsByTrafficChartMetaData) UnmarshalJSON(b []byte) error {
	type _SCISwitchDetailsReport129topSwitchPortsByTrafficChartMetaData SCISwitchDetailsReport129topSwitchPortsByTrafficChartMetaData
	tmpType := new(_SCISwitchDetailsReport129topSwitchPortsByTrafficChartMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "colorKeys")
	*t = SCISwitchDetailsReport129topSwitchPortsByTrafficChartMetaData(*tmpType)
	return nil
}

func (t *SCISwitchDetailsReport129topSwitchPortsByTrafficChartMetaData) MarshalJSON() ([]byte, error) {
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

func NewSCISwitchDetailsReport129topSwitchPortsByTrafficChartMetaData() *SCISwitchDetailsReport129topSwitchPortsByTrafficChartMetaData {
	m := new(SCISwitchDetailsReport129topSwitchPortsByTrafficChartMetaData)
	return m
}

// SCISwitchDetailsReport130topSwitchPortsByTrafficTableData
//
// Definition: SwitchDetailsReport_SwitchDetailsReport_130_topSwitchPortsByTrafficTable_Data
type SCISwitchDetailsReport130topSwitchPortsByTrafficTableData []*SCISwitchDetailsReport130topSwitchPortsByTrafficTableDataType

func MakeSCISwitchDetailsReport130topSwitchPortsByTrafficTableData() SCISwitchDetailsReport130topSwitchPortsByTrafficTableData {
	m := make(SCISwitchDetailsReport130topSwitchPortsByTrafficTableData, 0)
	return m
}

// SCISwitchDetailsReport130topSwitchPortsByTrafficTableDataType
//
// Definition: SwitchDetailsReport_SwitchDetailsReport_130_topSwitchPortsByTrafficTable_DataType
type SCISwitchDetailsReport130topSwitchPortsByTrafficTableDataType struct {
	Index *int `json:"index,omitempty"`

	PortName *string `json:"portName,omitempty"`

	PortNumber *string `json:"portNumber,omitempty"`

	TotalRxTraffic *int `json:"totalRxTraffic,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	TotalTxTraffic *int `json:"totalTxTraffic,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCISwitchDetailsReport130topSwitchPortsByTrafficTableDataType) UnmarshalJSON(b []byte) error {
	type _SCISwitchDetailsReport130topSwitchPortsByTrafficTableDataType SCISwitchDetailsReport130topSwitchPortsByTrafficTableDataType
	tmpType := new(_SCISwitchDetailsReport130topSwitchPortsByTrafficTableDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "index")
	delete(tmpType.XAdditionalProperties, "portName")
	delete(tmpType.XAdditionalProperties, "portNumber")
	delete(tmpType.XAdditionalProperties, "totalRxTraffic")
	delete(tmpType.XAdditionalProperties, "totalTraffic")
	delete(tmpType.XAdditionalProperties, "totalTxTraffic")
	*t = SCISwitchDetailsReport130topSwitchPortsByTrafficTableDataType(*tmpType)
	return nil
}

func (t *SCISwitchDetailsReport130topSwitchPortsByTrafficTableDataType) MarshalJSON() ([]byte, error) {
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
	if t.PortName != nil {
		tmp["portName"] = t.PortName
	}
	if t.PortNumber != nil {
		tmp["portNumber"] = t.PortNumber
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
	return json.Marshal(tmp)
}

func NewSCISwitchDetailsReport130topSwitchPortsByTrafficTableDataType() *SCISwitchDetailsReport130topSwitchPortsByTrafficTableDataType {
	m := new(SCISwitchDetailsReport130topSwitchPortsByTrafficTableDataType)
	return m
}

// SCISwitchDetailsReport130topSwitchPortsByTrafficTableMetaData
//
// Definition: SwitchDetailsReport_SwitchDetailsReport_130_topSwitchPortsByTrafficTable_MetaData
type SCISwitchDetailsReport130topSwitchPortsByTrafficTableMetaData struct {
	MaxValues *SCISwitchDetailsReport130topSwitchPortsByTrafficTableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	Name *string `json:"name,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCISwitchDetailsReport130topSwitchPortsByTrafficTableMetaData) UnmarshalJSON(b []byte) error {
	type _SCISwitchDetailsReport130topSwitchPortsByTrafficTableMetaData SCISwitchDetailsReport130topSwitchPortsByTrafficTableMetaData
	tmpType := new(_SCISwitchDetailsReport130topSwitchPortsByTrafficTableMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "maxValues")
	delete(tmpType.XAdditionalProperties, "name")
	*t = SCISwitchDetailsReport130topSwitchPortsByTrafficTableMetaData(*tmpType)
	return nil
}

func (t *SCISwitchDetailsReport130topSwitchPortsByTrafficTableMetaData) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(tmp)
}

func NewSCISwitchDetailsReport130topSwitchPortsByTrafficTableMetaData() *SCISwitchDetailsReport130topSwitchPortsByTrafficTableMetaData {
	m := new(SCISwitchDetailsReport130topSwitchPortsByTrafficTableMetaData)
	return m
}

// SCISwitchDetailsReport130topSwitchPortsByTrafficTableMetaDataMaxValuesType
//
// Definition: SwitchDetailsReport_SwitchDetailsReport_130_topSwitchPortsByTrafficTable_MetaDataMaxValuesType
type SCISwitchDetailsReport130topSwitchPortsByTrafficTableMetaDataMaxValuesType struct {
	TotalRxTraffic *int `json:"totalRxTraffic,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	TotalTxTraffic *int `json:"totalTxTraffic,omitempty"`
}

func NewSCISwitchDetailsReport130topSwitchPortsByTrafficTableMetaDataMaxValuesType() *SCISwitchDetailsReport130topSwitchPortsByTrafficTableMetaDataMaxValuesType {
	m := new(SCISwitchDetailsReport130topSwitchPortsByTrafficTableMetaDataMaxValuesType)
	return m
}

// SCISwitchDetailsReport138switchUptimeHistoryData
//
// Definition: SwitchDetailsReport_SwitchDetailsReport_138_switchUptimeHistory_Data
type SCISwitchDetailsReport138switchUptimeHistoryData []*SCISwitchDetailsReport138switchUptimeHistoryDataType

func MakeSCISwitchDetailsReport138switchUptimeHistoryData() SCISwitchDetailsReport138switchUptimeHistoryData {
	m := make(SCISwitchDetailsReport138switchUptimeHistoryData, 0)
	return m
}

// SCISwitchDetailsReport138switchUptimeHistoryDataType
//
// Definition: SwitchDetailsReport_SwitchDetailsReport_138_switchUptimeHistory_DataType
type SCISwitchDetailsReport138switchUptimeHistoryDataType struct {
	End *string `json:"end,omitempty"`

	IsSwitchUp *int `json:"isSwitchUp,omitempty"`

	Start *string `json:"start,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCISwitchDetailsReport138switchUptimeHistoryDataType) UnmarshalJSON(b []byte) error {
	type _SCISwitchDetailsReport138switchUptimeHistoryDataType SCISwitchDetailsReport138switchUptimeHistoryDataType
	tmpType := new(_SCISwitchDetailsReport138switchUptimeHistoryDataType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "end")
	delete(tmpType.XAdditionalProperties, "isSwitchUp")
	delete(tmpType.XAdditionalProperties, "start")
	*t = SCISwitchDetailsReport138switchUptimeHistoryDataType(*tmpType)
	return nil
}

func (t *SCISwitchDetailsReport138switchUptimeHistoryDataType) MarshalJSON() ([]byte, error) {
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
	if t.IsSwitchUp != nil {
		tmp["isSwitchUp"] = t.IsSwitchUp
	}
	if t.Start != nil {
		tmp["start"] = t.Start
	}
	return json.Marshal(tmp)
}

func NewSCISwitchDetailsReport138switchUptimeHistoryDataType() *SCISwitchDetailsReport138switchUptimeHistoryDataType {
	m := new(SCISwitchDetailsReport138switchUptimeHistoryDataType)
	return m
}

// SCISwitchDetailsReport138switchUptimeHistoryMetaData
//
// Definition: SwitchDetailsReport_SwitchDetailsReport_138_switchUptimeHistory_MetaData
type SCISwitchDetailsReport138switchUptimeHistoryMetaData struct {
	TotalDowntime *int `json:"totalDowntime,omitempty"`

	TotalUptime *int `json:"totalUptime,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCISwitchDetailsReport138switchUptimeHistoryMetaData) UnmarshalJSON(b []byte) error {
	type _SCISwitchDetailsReport138switchUptimeHistoryMetaData SCISwitchDetailsReport138switchUptimeHistoryMetaData
	tmpType := new(_SCISwitchDetailsReport138switchUptimeHistoryMetaData)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "totalDowntime")
	delete(tmpType.XAdditionalProperties, "totalUptime")
	*t = SCISwitchDetailsReport138switchUptimeHistoryMetaData(*tmpType)
	return nil
}

func (t *SCISwitchDetailsReport138switchUptimeHistoryMetaData) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.TotalDowntime != nil {
		tmp["totalDowntime"] = t.TotalDowntime
	}
	if t.TotalUptime != nil {
		tmp["totalUptime"] = t.TotalUptime
	}
	return json.Marshal(tmp)
}

func NewSCISwitchDetailsReport138switchUptimeHistoryMetaData() *SCISwitchDetailsReport138switchUptimeHistoryMetaData {
	m := new(SCISwitchDetailsReport138switchUptimeHistoryMetaData)
	return m
}
