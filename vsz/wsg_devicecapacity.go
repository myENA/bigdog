package vsz

// API Version: v9_0

type WSGDeviceCapacityDevicesSummary struct {
	// ApCapacity
	// maximum ap capacity of this node.
	// Constraints:
	//    - nullable
	ApCapacity *int `json:"apCapacity,omitempty"`

	// Aps
	// connected aps in this node.
	// Constraints:
	//    - nullable
	Aps *int `json:"aps,omitempty"`

	// MaxApOfCluster
	// maximum ap capacity per cluster.
	// Constraints:
	//    - nullable
	MaxApOfCluster *int `json:"maxApOfCluster,omitempty"`

	// MaxSwitchOfCluster
	// maximum switch capacity per cluster.
	// Constraints:
	//    - nullable
	MaxSwitchOfCluster *int `json:"maxSwitchOfCluster,omitempty"`

	// SwitchCapacity
	// maximum switch capacity of this node.
	// Constraints:
	//    - nullable
	SwitchCapacity *int `json:"switchCapacity,omitempty"`

	// Switches
	// connected switches in this node.
	// Constraints:
	//    - nullable
	Switches *int `json:"switches,omitempty"`

	// TotalApCapacity
	// maximum total ap capacity of this node.
	// Constraints:
	//    - nullable
	TotalApCapacity *int `json:"totalApCapacity,omitempty"`

	// TotalAps
	// total connected switches in the cluster.
	// Constraints:
	//    - nullable
	TotalAps *int `json:"totalAps,omitempty"`

	// TotalRemainingApCapacity
	// total remaining ap capacity of this node.
	// Constraints:
	//    - nullable
	TotalRemainingApCapacity *int `json:"totalRemainingApCapacity,omitempty"`

	// TotalRemainingSwitchCapacity
	// total remaining switch capacity of this node.
	// Constraints:
	//    - nullable
	TotalRemainingSwitchCapacity *int `json:"totalRemainingSwitchCapacity,omitempty"`

	// TotalSwitchCapacity
	// maximum total switch capacity of this node.
	// Constraints:
	//    - nullable
	TotalSwitchCapacity *int `json:"totalSwitchCapacity,omitempty"`

	// TotalSwitches
	// total connected switches in the cluster.
	// Constraints:
	//    - nullable
	TotalSwitches *int `json:"totalSwitches,omitempty"`
}

func NewWSGDeviceCapacityDevicesSummary() *WSGDeviceCapacityDevicesSummary {
	m := new(WSGDeviceCapacityDevicesSummary)
	return m
}
