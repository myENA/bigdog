package bigdog

// API Version: 1.0.0

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
	PortName *string `json:"portName,omitempty"`

	PortNumber *string `json:"portNumber,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`
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
}

func NewSCISwitchDetailsReport138switchUptimeHistoryMetaData() *SCISwitchDetailsReport138switchUptimeHistoryMetaData {
	m := new(SCISwitchDetailsReport138switchUptimeHistoryMetaData)
	return m
}
