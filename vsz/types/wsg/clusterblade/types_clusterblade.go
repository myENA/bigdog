package clusterblade

// API Version: v8_1

type BladeProgress struct {
	// BladeUUID
	// bladeUUID of bladeProgress
	BladeUUID *string `json:"bladeUUID,omitempty"`

	// HostName
	// hostName of bladeProgress
	HostName *string `json:"hostName,omitempty"`

	// IterationName
	// iterationName of bladeProgress
	IterationName *string `json:"iterationName,omitempty"`

	// Progress
	// progress of bladeProgress
	Progress *int `json:"progress,omitempty"`

	// State
	// state of bladeProgress
	State *string `json:"state,omitempty"`
}

type ClusterOperationProgress struct {
	Operation *Operation `json:"operation,omitempty"`

	// OverallProgress
	// overallProgress of clusterOperationProgress
	OverallProgress *int `json:"overallProgress,omitempty"`

	PreviousOperationRecord *PreviousOperationRecord `json:"previousOperationRecord,omitempty"`
}

type ClusterState struct {
	// ClusterName
	// cluster name
	ClusterName *string `json:"clusterName,omitempty"`

	// ClusterRole
	// The cluster role of the current controller node
	ClusterRole *string `json:"clusterRole,omitempty" validate:"oneof=Leader Follower"`

	// ClusterState
	// cluster state
	ClusterState *string `json:"clusterState,omitempty" validate:"oneof=In_Service Out_Of_Service Maintenance Read_Only NetworkPartitionSuspected"`

	// CurrentNodeId
	// Identifier of the current controller node
	CurrentNodeId *string `json:"currentNodeId,omitempty"`

	// CurrentNodeName
	// The name of the current controller node
	CurrentNodeName *string `json:"currentNodeName,omitempty"`

	ManagementServiceStateList []*ClusterStateManagementServiceStateListType `json:"managementServiceStateList,omitempty"`

	NodeStateList []*ClusterStateNodeStateListType `json:"nodeStateList,omitempty"`
}

type ClusterStateManagementServiceStateListType struct {
	// ManagementServiceState
	// management service state
	ManagementServiceState *string `json:"managementServiceState,omitempty" validate:"oneof=Out_Of_Service In_Service"`

	// NodeId
	// Identifier of the controller node
	NodeId *string `json:"nodeId,omitempty"`

	// NodeName
	// node name
	NodeName *string `json:"nodeName,omitempty"`
}

type ClusterStateNodeStateListType struct {
	// NodeId
	// Identifier of the controller node
	NodeId *string `json:"nodeId,omitempty"`

	NodeName *string `json:"nodeName,omitempty"`

	// NodeState
	// node state
	NodeState *string `json:"nodeState,omitempty" validate:"oneof=Out_Of_Service In_Service"`
}

type ClusterStatus struct {
	// ClusterStatus
	// progress of bladeProgress
	ClusterStatus *string `json:"clusterStatus,omitempty" validate:"oneof=In_Service Out_Of_Service Maintenance Read_Only NetworkPartitionSuspected"`
}

type ClusterUpgradeProgress struct {
	// BladeProgresss
	// bladeProgressMap of clusterOperationProgress
	BladeProgresss []*BladeProgress `json:"bladeProgresss,omitempty"`

	// ClusterOperationBlockUI
	// clusterOperationBlockUI of clusterOperationProgress
	ClusterOperationBlockUI *bool `json:"clusterOperationBlockUI,omitempty"`

	// ClusterOperationDisplayMsg
	// clusterOperationDisplayMsg of clusterOperationProgress
	ClusterOperationDisplayMsg *string `json:"clusterOperationDisplayMsg,omitempty"`

	// ClusterSubTaskState
	// clusterSubTaskState of clusterOperationProgress
	ClusterSubTaskState *string `json:"clusterSubTaskState,omitempty" validate:"oneof=None Running Failed Completed"`

	// IsSelfBladeRebooting
	// isSelfBladeRebooting of clusterOperationProgress
	IsSelfBladeRebooting *bool `json:"isSelfBladeRebooting,omitempty"`

	Operation *Operation `json:"operation,omitempty"`

	// OverallProgress
	// overallProgress of clusterOperationProgress
	OverallProgress *int `json:"overallProgress,omitempty"`

	PreviousOperationRecord *PreviousOperationRecord `json:"previousOperationRecord,omitempty"`
}

type ControlNodeStatus struct {
	NodeStatusList []*ControlNodeStatusNodeStatusListType `json:"nodeStatusList,omitempty"`
}

type ControlNodeStatusNodeStatusListType struct {
	// NodeId
	// Identifier of the controller node
	NodeId *string `json:"nodeId,omitempty"`

	// NodeStatus
	// node status
	NodeStatus *string `json:"nodeStatus,omitempty" validate:"oneof=Out_Of_Service Bootstrapping Got_WSG_Version WSG_FW_Upgrading Initializing_Database Syncing_Configurations Changing_Configurations Launching_Apps In_Service Shutting_Down_Apps"`
}

type Operation string

type PreviousOperationRecord struct {
	// ErrorMsg
	// errorMsg of previousOperationRecord
	ErrorMsg *string `json:"errorMsg,omitempty"`

	Operation *Operation `json:"operation,omitempty"`

	// Success
	// success of previousOperationRecord
	Success *bool `json:"success,omitempty"`
}

type UploadPatchInfo struct {
	// AllowVersions
	// allowVersions of uploadPatchInfo
	AllowVersions []string `json:"allowVersions,omitempty"`

	// ApVersion
	// apVersion of uploadPatchInfo
	ApVersion *string `json:"apVersion,omitempty"`

	// ControlbladeVersion
	// controlbladeVersion of uploadPatchInfo
	ControlbladeVersion *string `json:"controlbladeVersion,omitempty"`

	// DatabladeVersion
	// databladeVersion of uploadPatchInfo
	DatabladeVersion *string `json:"databladeVersion,omitempty"`

	// FileName
	// fileName of uploadPatchInfo
	FileName *string `json:"fileName,omitempty"`

	// FileSize
	// fileSize of uploadPatchInfo
	FileSize *float64 `json:"fileSize,omitempty"`

	// FileUploadPath
	// fileUploadPath of uploadPatchInfo
	FileUploadPath *string `json:"fileUploadPath,omitempty"`

	// Version
	// version of uploadPatchInfo
	Version *string `json:"version,omitempty"`
}
