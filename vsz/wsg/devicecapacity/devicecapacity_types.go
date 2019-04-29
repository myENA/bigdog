package devicecapacity

// API Version: v8_0

type DevicesSummary struct {
	ApCapacity                   *int `json:"apCapacity,omitempty"`
	Aps                          *int `json:"aps,omitempty"`
	MaxApOfCluster               *int `json:"maxApOfCluster,omitempty"`
	MaxSwitchOfCluster           *int `json:"maxSwitchOfCluster,omitempty"`
	SwitchCapacity               *int `json:"switchCapacity,omitempty"`
	Switches                     *int `json:"switches,omitempty"`
	TotalApCapacity              *int `json:"totalApCapacity,omitempty"`
	TotalAps                     *int `json:"totalAps,omitempty"`
	TotalRemainingApCapacity     *int `json:"totalRemainingApCapacity,omitempty"`
	TotalRemainingSwitchCapacity *int `json:"totalRemainingSwitchCapacity,omitempty"`
	TotalSwitchCapacity          *int `json:"totalSwitchCapacity,omitempty"`
	TotalSwitches                *int `json:"totalSwitches,omitempty"`
}
