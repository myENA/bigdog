package dp

// API Version: v8_0

type EmptyResult struct {
	DpEmptyResult *string `json:"dp_emptyResult,omitempty"`
}

type SwitchoverDp struct {
	ClusterName  *string  `json:"clusterName,omitempty"`
	DeleteRecord *bool    `json:"deleteRecord,omitempty"`
	DpIDList     []string `json:"dpIdList,omitempty"`
	IPOrFqdn     *string  `json:"ipOrFqdn,omitempty"`
}
