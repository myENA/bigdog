package vsz

// API Version: v9_0

type WSGClusterRedundancyActiveCluster struct {
	// ManagementEntryList
	// Management entry list of target active cluster
	// Constraints:
	//    - nullable
	ManagementEntryList []*WSGClusterRedundancyManagementEntry `json:"managementEntryList,omitempty" validate:"omitempty,dive"`

	// Priority
	// Priority of target active cluster
	// Constraints:
	//    - nullable
	Priority *int `json:"priority,omitempty"`

	// TargetClusterAdminPassword
	// Password of admin account of target active cluster
	// Constraints:
	//    - nullable
	TargetClusterAdminPassword *string `json:"targetClusterAdminPassword,omitempty"`
}

func NewWSGClusterRedundancyActiveCluster() *WSGClusterRedundancyActiveCluster {
	m := new(WSGClusterRedundancyActiveCluster)
	return m
}

type WSGClusterRedundancySettings struct {
	// ActiveClusterList
	// A list of target active clusters (Active-Active only)
	// Constraints:
	//    - nullable
	ActiveClusterList []*WSGClusterRedundancyActiveCluster `json:"activeClusterList,omitempty" validate:"omitempty,dive"`

	// ClusterRedundancyEnabled
	// Cluster redundancy enabled
	// Constraints:
	//    - nullable
	ClusterRedundancyEnabled *bool `json:"clusterRedundancyEnabled,omitempty"`

	// ClusterRedundancyType
	// Cluster redundancy type
	// Constraints:
	//    - nullable
	ClusterRedundancyType *string `json:"clusterRedundancyType,omitempty"`

	// DateOfMonth
	// Scheduled date of the month (1-31) (Active-Active only)
	// Constraints:
	//    - nullable
	DateOfMonth *int `json:"dateOfMonth,omitempty"`

	// DayOfWeek
	// Scheduled day of the week (Active-Active only)
	// Constraints:
	//    - nullable
	//    - oneof:[SUNDAY,MONDAY,TUESDAY,WEDNESDAY,THURSDAY,FRIDAY,SATURDAY]
	DayOfWeek *string `json:"dayOfWeek,omitempty" validate:"omitempty,oneof=SUNDAY MONDAY TUESDAY WEDNESDAY THURSDAY FRIDAY SATURDAY"`

	// Hour
	// Schedule sync time in daily hour (0-23)
	// Constraints:
	//    - nullable
	Hour *int `json:"hour,omitempty"`

	// Interval
	// Schedule interval (Active-Active only)
	// Constraints:
	//    - nullable
	//    - oneof:[MONTHLY,WEEKLY,DAILY,HOURLY]
	Interval *string `json:"interval,omitempty" validate:"omitempty,oneof=MONTHLY WEEKLY DAILY HOURLY"`

	// ManagementEntryList
	// Management entry list of standby cluster (Active-Standby only)
	// Constraints:
	//    - nullable
	ManagementEntryList []*WSGClusterRedundancyManagementEntry `json:"managementEntryList,omitempty" validate:"omitempty,dive"`

	// Minute
	// Schedule sync time in minute (Active-Active only)
	// Constraints:
	//    - nullable
	Minute *int `json:"minute,omitempty"`

	// ScheduleSyncUpEnabled
	// Scheduled configuration sync enabled
	// Constraints:
	//    - nullable
	ScheduleSyncUpEnabled *bool `json:"scheduleSyncUpEnabled,omitempty"`

	// StandbyAdminPassword
	// Password of admin account of standby cluster (Active-Standby only)
	// Constraints:
	//    - nullable
	StandbyAdminPassword *string `json:"standbyAdminPassword,omitempty"`
}

func NewWSGClusterRedundancySettings() *WSGClusterRedundancySettings {
	m := new(WSGClusterRedundancySettings)
	return m
}

type WSGClusterRedundancyManagementEntry struct {
	// Ip
	// Management IP
	// Constraints:
	//    - nullable
	Ip *string `json:"ip,omitempty"`

	// Port
	// Management port
	// Constraints:
	//    - nullable
	Port *string `json:"port,omitempty"`
}

func NewWSGClusterRedundancyManagementEntry() *WSGClusterRedundancyManagementEntry {
	m := new(WSGClusterRedundancyManagementEntry)
	return m
}

type WSGClusterRedundancyUpdateClusterRedundancy struct {
	// ActiveClusterList
	// A list of target active clusters (Active-Active only)
	// Constraints:
	//    - nullable
	ActiveClusterList []*WSGClusterRedundancyActiveCluster `json:"activeClusterList,omitempty" validate:"omitempty,dive"`

	// ClusterRedundancyEnabled
	// Cluster redundancy enabled
	// Constraints:
	//    - required
	ClusterRedundancyEnabled *bool `json:"clusterRedundancyEnabled" validate:"required"`

	// ClusterRedundancyType
	// Cluster redundancy type (Active-Standby, or Active-Active)
	// Constraints:
	//    - nullable
	ClusterRedundancyType *string `json:"clusterRedundancyType,omitempty"`

	// DateOfMonth
	// Scheduled date of the month (1-31) (Active-Active only)
	// Constraints:
	//    - nullable
	DateOfMonth *int `json:"dateOfMonth,omitempty"`

	// DayOfWeek
	// Scheduled day of the week (Active-Active only)
	// Constraints:
	//    - nullable
	//    - oneof:[SUNDAY,MONDAY,TUESDAY,WEDNESDAY,THURSDAY,FRIDAY,SATURDAY]
	DayOfWeek *string `json:"dayOfWeek,omitempty" validate:"omitempty,oneof=SUNDAY MONDAY TUESDAY WEDNESDAY THURSDAY FRIDAY SATURDAY"`

	// Hour
	// Schedule sync time in daily hour (0-23)
	// Constraints:
	//    - nullable
	Hour *int `json:"hour,omitempty"`

	// Interval
	// Schedule interval (Active-Active only)
	// Constraints:
	//    - nullable
	//    - oneof:[MONTHLY,WEEKLY,DAILY,HOURLY]
	Interval *string `json:"interval,omitempty" validate:"omitempty,oneof=MONTHLY WEEKLY DAILY HOURLY"`

	// ManagementEntryList
	// Management entry list of standby cluster (Active-Standby only)
	// Constraints:
	//    - nullable
	ManagementEntryList []*WSGClusterRedundancyManagementEntry `json:"managementEntryList,omitempty" validate:"omitempty,dive"`

	// Minute
	// Schedule sync time in minute (Active-Active only)
	// Constraints:
	//    - nullable
	Minute *int `json:"minute,omitempty"`

	// ScheduleSyncUpEnabled
	// Scheduled configuration sync enabled
	// Constraints:
	//    - nullable
	ScheduleSyncUpEnabled *bool `json:"scheduleSyncUpEnabled,omitempty"`

	// StandbyAdminPassword
	// Password of admin account of standby cluster (Active-Standby only)
	// Constraints:
	//    - nullable
	StandbyAdminPassword *string `json:"standbyAdminPassword,omitempty"`
}

func NewWSGClusterRedundancyUpdateClusterRedundancy() *WSGClusterRedundancyUpdateClusterRedundancy {
	m := new(WSGClusterRedundancyUpdateClusterRedundancy)
	return m
}
