package vsz

// API Version: v9_0

type WSGDPSwitchoverDp struct {
	// ClusterName
	// Name of destination cluster, Notice: System will take ipOrFqdn as 1st priority, and clusterName as 2nd.
	// Constraints:
	//    - nullable
	ClusterName *string `json:"clusterName,omitempty"`

	// DeleteRecord
	// Flag to delete DP record after switchover cluster. Default value is false.
	// Constraints:
	//    - nullable
	DeleteRecord *bool `json:"deleteRecord,omitempty"`

	// DpIdList
	// DP ID list
	// Constraints:
	//    - nullable
	DpIdList []string `json:"dpIdList,omitempty" validate:"omitempty,dive"`

	// IpOrFqdn
	// IP or FQDN address of destination cluster, Notice: System will take ipOrFqdn as 1st priority, and clusterName as 2nd.
	// Constraints:
	//    - nullable
	IpOrFqdn *string `json:"ipOrFqdn,omitempty"`
}

func NewWSGDPSwitchoverDp() *WSGDPSwitchoverDp {
	m := new(WSGDPSwitchoverDp)
	return m
}
