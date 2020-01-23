package vsz

// API Version: v8_1

type WSGDPEmptyResult struct {
	DpemptyResult *string `json:"dp_emptyResult,omitempty"`
}

func NewWSGDPEmptyResult() *WSGDPEmptyResult {
	m := new(WSGDPEmptyResult)
	return m
}

type WSGDPSwitchoverDp struct {
	// ClusterName
	// Name of destination cluster, Notice: System will take ipOrFqdn as 1st priority, and clusterName as 2nd.
	ClusterName *string `json:"clusterName,omitempty"`

	// DeleteRecord
	// Flag to delete DP record after switchover cluster. Default value is false.
	DeleteRecord *bool `json:"deleteRecord,omitempty"`

	// DpIdList
	// DP ID list
	DpIdList []string `json:"dpIdList,omitempty"`

	// IpOrFqdn
	// IP or FQDN address of destination cluster, Notice: System will take ipOrFqdn as 1st priority, and clusterName as 2nd.
	IpOrFqdn *string `json:"ipOrFqdn,omitempty"`
}

func NewWSGDPSwitchoverDp() *WSGDPSwitchoverDp {
	m := new(WSGDPSwitchoverDp)
	return m
}
