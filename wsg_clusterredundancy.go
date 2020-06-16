package bigdog

// API Version: v9_0

// WSGClusterRedundancyActiveCluster
//
// Definition: clusterRedundancy_activeCluster
type WSGClusterRedundancyActiveCluster struct {
	// ManagementEntryList
	// Management entry list of target active cluster
	ManagementEntryList []*WSGClusterRedundancyManagementEntry `json:"managementEntryList,omitempty"`

	// Priority
	// Priority of target active cluster
	Priority *int `json:"priority,omitempty"`

	// TargetClusterAdminPassword
	// Password of admin account of target active cluster
	TargetClusterAdminPassword *string `json:"targetClusterAdminPassword,omitempty"`
}

func NewWSGClusterRedundancyActiveCluster() *WSGClusterRedundancyActiveCluster {
	m := new(WSGClusterRedundancyActiveCluster)
	return m
}

// WSGClusterRedundancySettings
//
// Definition: clusterRedundancy_clusterRedundancySettings
type WSGClusterRedundancySettings struct {
	// ActiveClusterList
	// A list of target active clusters (Active-Active only)
	ActiveClusterList []*WSGClusterRedundancyActiveCluster `json:"activeClusterList,omitempty"`

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
	// Constraints:
	//    - oneof:[SUNDAY,MONDAY,TUESDAY,WEDNESDAY,THURSDAY,FRIDAY,SATURDAY]
	DayOfWeek *string `json:"dayOfWeek,omitempty"`

	// Hour
	// Schedule sync time in daily hour (0-23)
	Hour *int `json:"hour,omitempty"`

	// Interval
	// Schedule interval (Active-Active only)
	// Constraints:
	//    - oneof:[MONTHLY,WEEKLY,DAILY,HOURLY]
	Interval *string `json:"interval,omitempty"`

	// ManagementEntryList
	// Management entry list of standby cluster (Active-Standby only)
	ManagementEntryList []*WSGClusterRedundancyManagementEntry `json:"managementEntryList,omitempty"`

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

func NewWSGClusterRedundancySettings() *WSGClusterRedundancySettings {
	m := new(WSGClusterRedundancySettings)
	return m
}

// WSGClusterRedundancyManagementEntry
//
// Definition: clusterRedundancy_managementEntry
type WSGClusterRedundancyManagementEntry struct {
	// Ip
	// Management IP
	Ip *string `json:"ip,omitempty"`

	// Port
	// Management port
	Port *string `json:"port,omitempty"`
}

func NewWSGClusterRedundancyManagementEntry() *WSGClusterRedundancyManagementEntry {
	m := new(WSGClusterRedundancyManagementEntry)
	return m
}

// WSGClusterRedundancyUpdateClusterRedundancy
//
// Definition: clusterRedundancy_updateClusterRedundancy
type WSGClusterRedundancyUpdateClusterRedundancy struct {
	// ActiveClusterList
	// A list of target active clusters (Active-Active only)
	ActiveClusterList []*WSGClusterRedundancyActiveCluster `json:"activeClusterList,omitempty"`

	// ClusterRedundancyEnabled
	// Cluster redundancy enabled
	// Constraints:
	//    - required
	ClusterRedundancyEnabled *bool `json:"clusterRedundancyEnabled"`

	// ClusterRedundancyType
	// Cluster redundancy type (Active-Standby, or Active-Active)
	ClusterRedundancyType *string `json:"clusterRedundancyType,omitempty"`

	// DateOfMonth
	// Scheduled date of the month (1-31) (Active-Active only)
	DateOfMonth *int `json:"dateOfMonth,omitempty"`

	// DayOfWeek
	// Scheduled day of the week (Active-Active only)
	// Constraints:
	//    - oneof:[SUNDAY,MONDAY,TUESDAY,WEDNESDAY,THURSDAY,FRIDAY,SATURDAY]
	DayOfWeek *string `json:"dayOfWeek,omitempty"`

	// Hour
	// Schedule sync time in daily hour (0-23)
	Hour *int `json:"hour,omitempty"`

	// Interval
	// Schedule interval (Active-Active only)
	// Constraints:
	//    - oneof:[MONTHLY,WEEKLY,DAILY,HOURLY]
	Interval *string `json:"interval,omitempty"`

	// ManagementEntryList
	// Management entry list of standby cluster (Active-Standby only)
	ManagementEntryList []*WSGClusterRedundancyManagementEntry `json:"managementEntryList,omitempty"`

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

func NewWSGClusterRedundancyUpdateClusterRedundancy() *WSGClusterRedundancyUpdateClusterRedundancy {
	m := new(WSGClusterRedundancyUpdateClusterRedundancy)
	return m
}
