package clusterredundancy

// API Version: v8_0

type ActiveCluster struct {
	ManagementEntryList        []*ManagementEntry `json:"managementEntryList,omitempty"`
	Priority                   *int               `json:"priority,omitempty"`
	TargetClusterAdminPassword *string            `json:"targetClusterAdminPassword,omitempty"`
}

type ClusterRedundancySettings struct {
	ActiveClusterList        []*ActiveCluster   `json:"activeClusterList,omitempty"`
	ClusterRedundancyEnabled *bool              `json:"clusterRedundancyEnabled,omitempty"`
	ClusterRedundancyType    *string            `json:"clusterRedundancyType,omitempty"`
	DateOfMonth              *int               `json:"dateOfMonth,omitempty"`
	DayOfWeek                *string            `json:"dayOfWeek,omitempty"`
	Hour                     *int               `json:"hour,omitempty"`
	Interval                 *string            `json:"interval,omitempty"`
	ManagementEntryList      []*ManagementEntry `json:"managementEntryList,omitempty"`
	Minute                   *int               `json:"minute,omitempty"`
	ScheduleSyncUpEnabled    *bool              `json:"scheduleSyncUpEnabled,omitempty"`
	StandbyAdminPassword     *string            `json:"standbyAdminPassword,omitempty"`
}

type ManagementEntry struct {
	IP   *string `json:"ip,omitempty"`
	Port *string `json:"port,omitempty"`
}

type UpdateClusterRedundancy struct {
	ActiveClusterList        []*ActiveCluster   `json:"activeClusterList,omitempty"`
	ClusterRedundancyEnabled *bool              `json:"clusterRedundancyEnabled,omitempty"`
	ClusterRedundancyType    *string            `json:"clusterRedundancyType,omitempty"`
	DateOfMonth              *int               `json:"dateOfMonth,omitempty"`
	DayOfWeek                *string            `json:"dayOfWeek,omitempty"`
	Hour                     *int               `json:"hour,omitempty"`
	Interval                 *string            `json:"interval,omitempty"`
	ManagementEntryList      []*ManagementEntry `json:"managementEntryList,omitempty"`
	Minute                   *int               `json:"minute,omitempty"`
	ScheduleSyncUpEnabled    *bool              `json:"scheduleSyncUpEnabled,omitempty"`
	StandbyAdminPassword     *string            `json:"standbyAdminPassword,omitempty"`
}
