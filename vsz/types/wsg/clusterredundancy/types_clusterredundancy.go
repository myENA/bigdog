package clusterredundancy

// API Version: v8_1

type ActiveCluster struct {
	// ManagementEntryList
	// Management entry list of target active cluster
	ManagementEntryList []*ManagementEntry `json:"managementEntryList,omitempty"`

	// Priority
	// Priority of target active cluster
	Priority *int `json:"priority,omitempty"`

	// TargetClusterAdminPassword
	// Password of admin account of target active cluster
	TargetClusterAdminPassword *string `json:"targetClusterAdminPassword,omitempty"`
}

type ClusterRedundancySettings struct {
	// ActiveClusterList
	// A list of target active clusters (Active-Active only)
	ActiveClusterList []*ActiveCluster `json:"activeClusterList,omitempty"`

	// ClusterRedundancyEnabled
	// Cluster redundancy enabled
	ClusterRedundancyEnabled *bool `json:"clusterRedundancyEnabled,omitempty"`

	// ClusterRedundancyType
	// Cluster redundancy type
	ClusterRedundancyType *string `json:"clusterRedundancyType,omitempty"`

	// DateOfMonth
	// Scheduled date of the month (1-31) (Active-Active only)
	DateOfMonth *int `json:"dateOfMonth,omitempty"`

	// DayOfWeek
	// Scheduled day of the week (Active-Active only)
	DayOfWeek *string `json:"dayOfWeek,omitempty"`

	// Hour
	// Schedule sync time in daily hour (0-23)
	Hour *int `json:"hour,omitempty"`

	// Interval
	// Schedule interval (Active-Active only)
	Interval *string `json:"interval,omitempty"`

	// ManagementEntryList
	// Management entry list of standby cluster (Active-Standby only)
	ManagementEntryList []*ManagementEntry `json:"managementEntryList,omitempty"`

	// Minute
	// Schedule sync time in minute (Active-Active only)
	Minute *int `json:"minute,omitempty"`

	// ScheduleSyncUpEnabled
	// Scheduled configuration sync enabled
	ScheduleSyncUpEnabled *bool `json:"scheduleSyncUpEnabled,omitempty"`

	// StandbyAdminPassword
	// Password of admin account of standby cluster (Active-Standby only)
	StandbyAdminPassword *string `json:"standbyAdminPassword,omitempty"`
}

type ManagementEntry struct {
	// Ip
	// Management IP
	Ip *string `json:"ip,omitempty"`

	// Port
	// Management port
	Port *string `json:"port,omitempty"`
}

type UpdateClusterRedundancy struct {
	// ActiveClusterList
	// A list of target active clusters (Active-Active only)
	ActiveClusterList []*ActiveCluster `json:"activeClusterList,omitempty"`

	// ClusterRedundancyEnabled
	// Cluster redundancy enabled
	ClusterRedundancyEnabled *bool `json:"clusterRedundancyEnabled,omitempty"`

	// ClusterRedundancyType
	// Cluster redundancy type (Active-Standby, or Active-Active)
	ClusterRedundancyType *string `json:"clusterRedundancyType,omitempty"`

	// DateOfMonth
	// Scheduled date of the month (1-31) (Active-Active only)
	DateOfMonth *int `json:"dateOfMonth,omitempty"`

	// DayOfWeek
	// Scheduled day of the week (Active-Active only)
	DayOfWeek *string `json:"dayOfWeek,omitempty"`

	// Hour
	// Schedule sync time in daily hour (0-23)
	Hour *int `json:"hour,omitempty"`

	// Interval
	// Schedule interval (Active-Active only)
	Interval *string `json:"interval,omitempty"`

	// ManagementEntryList
	// Management entry list of standby cluster (Active-Standby only)
	ManagementEntryList []*ManagementEntry `json:"managementEntryList,omitempty"`

	// Minute
	// Schedule sync time in minute (Active-Active only)
	Minute *int `json:"minute,omitempty"`

	// ScheduleSyncUpEnabled
	// Scheduled configuration sync enabled
	ScheduleSyncUpEnabled *bool `json:"scheduleSyncUpEnabled,omitempty"`

	// StandbyAdminPassword
	// Password of admin account of standby cluster (Active-Standby only)
	StandbyAdminPassword *string `json:"standbyAdminPassword,omitempty"`
}
