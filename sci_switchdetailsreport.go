package bigdog

// API Version: 1.0.0

// SCISwitchDetailsReport126switchResourceUtilizationData
//
// Definition: SwitchDetailsReport.SwitchDetailsReport.126.switchResourceUtilization.Data
type SCISwitchDetailsReport126switchResourceUtilizationData []*SCISwitchDetailsReport126switchResourceUtilizationDataType

func MakeSCISwitchDetailsReport126switchResourceUtilizationData() SCISwitchDetailsReport126switchResourceUtilizationData {
	m := make(SCISwitchDetailsReport126switchResourceUtilizationData, 0)
	return m
}

// SCISwitchDetailsReport126switchResourceUtilizationDataType
//
// Definition: SwitchDetailsReport.SwitchDetailsReport.126.switchResourceUtilization.DataType
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
// Definition: SwitchDetailsReport.SwitchDetailsReport.129.topSwitchPortsByTrafficChart.Data
type SCISwitchDetailsReport129topSwitchPortsByTrafficChartData []*SCISwitchDetailsReport129topSwitchPortsByTrafficChartDataType

func MakeSCISwitchDetailsReport129topSwitchPortsByTrafficChartData() SCISwitchDetailsReport129topSwitchPortsByTrafficChartData {
	m := make(SCISwitchDetailsReport129topSwitchPortsByTrafficChartData, 0)
	return m
}

// SCISwitchDetailsReport129topSwitchPortsByTrafficChartDataType
//
// Definition: SwitchDetailsReport.SwitchDetailsReport.129.topSwitchPortsByTrafficChart.DataType
type SCISwitchDetailsReport129topSwitchPortsByTrafficChartDataType struct {
	PortName *string `json:"portName,omitempty"`

	PortNumber *string `json:"portNumber,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`
}

func NewSCISwitchDetailsReport129topSwitchPortsByTrafficChartDataType() *SCISwitchDetailsReport129topSwitchPortsByTrafficChartDataType {
	m := new(SCISwitchDetailsReport129topSwitchPortsByTrafficChartDataType)
	return m
}

// SCISwitchDetailsReport129topSwitchPortsByTrafficChartMetaData
//
// Definition: SwitchDetailsReport.SwitchDetailsReport.129.topSwitchPortsByTrafficChart.MetaData
type SCISwitchDetailsReport129topSwitchPortsByTrafficChartMetaData struct {
	ColorKeys []string `json:"colorKeys,omitempty"`
}

func NewSCISwitchDetailsReport129topSwitchPortsByTrafficChartMetaData() *SCISwitchDetailsReport129topSwitchPortsByTrafficChartMetaData {
	m := new(SCISwitchDetailsReport129topSwitchPortsByTrafficChartMetaData)
	return m
}

// SCISwitchDetailsReport130topSwitchPortsByTrafficTableData
//
// Definition: SwitchDetailsReport.SwitchDetailsReport.130.topSwitchPortsByTrafficTable.Data
type SCISwitchDetailsReport130topSwitchPortsByTrafficTableData []*SCISwitchDetailsReport130topSwitchPortsByTrafficTableDataType

func MakeSCISwitchDetailsReport130topSwitchPortsByTrafficTableData() SCISwitchDetailsReport130topSwitchPortsByTrafficTableData {
	m := make(SCISwitchDetailsReport130topSwitchPortsByTrafficTableData, 0)
	return m
}

// SCISwitchDetailsReport130topSwitchPortsByTrafficTableDataType
//
// Definition: SwitchDetailsReport.SwitchDetailsReport.130.topSwitchPortsByTrafficTable.DataType
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
// Definition: SwitchDetailsReport.SwitchDetailsReport.130.topSwitchPortsByTrafficTable.MetaData
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
// Definition: SwitchDetailsReport.SwitchDetailsReport.130.topSwitchPortsByTrafficTable.MetaDataMaxValuesType
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
// Definition: SwitchDetailsReport.SwitchDetailsReport.138.switchUptimeHistory.Data
type SCISwitchDetailsReport138switchUptimeHistoryData []*SCISwitchDetailsReport138switchUptimeHistoryDataType

func MakeSCISwitchDetailsReport138switchUptimeHistoryData() SCISwitchDetailsReport138switchUptimeHistoryData {
	m := make(SCISwitchDetailsReport138switchUptimeHistoryData, 0)
	return m
}

// SCISwitchDetailsReport138switchUptimeHistoryDataType
//
// Definition: SwitchDetailsReport.SwitchDetailsReport.138.switchUptimeHistory.DataType
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
// Definition: SwitchDetailsReport.SwitchDetailsReport.138.switchUptimeHistory.MetaData
type SCISwitchDetailsReport138switchUptimeHistoryMetaData struct {
	TotalDowntime *int `json:"totalDowntime,omitempty"`

	TotalUptime *int `json:"totalUptime,omitempty"`
}

func NewSCISwitchDetailsReport138switchUptimeHistoryMetaData() *SCISwitchDetailsReport138switchUptimeHistoryMetaData {
	m := new(SCISwitchDetailsReport138switchUptimeHistoryMetaData)
	return m
}