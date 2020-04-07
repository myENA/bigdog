package vsz

// API Version: v9_0

type WSGClusterBladeBladeProgress struct {
	// BladeUUID
	// bladeUUID of bladeProgress
	// Constraints:
	//    - nullable
	BladeUUID *string `json:"bladeUUID,omitempty"`

	// HostName
	// hostName of bladeProgress
	// Constraints:
	//    - nullable
	HostName *string `json:"hostName,omitempty"`

	// IterationName
	// iterationName of bladeProgress
	// Constraints:
	//    - nullable
	IterationName *string `json:"iterationName,omitempty"`

	// Progress
	// progress of bladeProgress
	// Constraints:
	//    - nullable
	Progress *int `json:"progress,omitempty"`

	// State
	// state of bladeProgress
	// Constraints:
	//    - nullable
	State *string `json:"state,omitempty"`
}

func NewWSGClusterBladeBladeProgress() *WSGClusterBladeBladeProgress {
	m := new(WSGClusterBladeBladeProgress)
	return m
}

type WSGClusterBladeClusterOperationProgress struct {
	// Operation
	// Constraints:
	//    - nullable
	Operation *WSGClusterBladeOperation `json:"operation,omitempty"`

	// OverallProgress
	// overallProgress of clusterOperationProgress
	// Constraints:
	//    - nullable
	OverallProgress *int `json:"overallProgress,omitempty"`

	// PreviousOperationRecord
	// Constraints:
	//    - nullable
	PreviousOperationRecord *WSGClusterBladePreviousOperationRecord `json:"previousOperationRecord,omitempty"`
}

func NewWSGClusterBladeClusterOperationProgress() *WSGClusterBladeClusterOperationProgress {
	m := new(WSGClusterBladeClusterOperationProgress)
	return m
}

type WSGClusterBladeClusterState struct {
	// ClusterName
	// cluster name
	// Constraints:
	//    - nullable
	ClusterName *string `json:"clusterName,omitempty"`

	// ClusterRole
	// The cluster role of the current controller node
	// Constraints:
	//    - nullable
	//    - oneof:[Leader,Follower]
	ClusterRole *string `json:"clusterRole,omitempty" validate:"omitempty,oneof=Leader Follower"`

	// ClusterState
	// cluster state
	// Constraints:
	//    - nullable
	//    - oneof:[In_Service,Out_Of_Service,Maintenance,Read_Only,NetworkPartitionSuspected]
	ClusterState *string `json:"clusterState,omitempty" validate:"omitempty,oneof=In_Service Out_Of_Service Maintenance Read_Only NetworkPartitionSuspected"`

	// CurrentNodeId
	// Identifier of the current controller node
	// Constraints:
	//    - nullable
	CurrentNodeId *string `json:"currentNodeId,omitempty"`

	// CurrentNodeName
	// The name of the current controller node
	// Constraints:
	//    - nullable
	CurrentNodeName *string `json:"currentNodeName,omitempty"`

	// ManagementServiceStateList
	// Constraints:
	//    - nullable
	ManagementServiceStateList []*WSGClusterBladeClusterStateManagementServiceStateListType `json:"managementServiceStateList,omitempty" validate:"omitempty,dive"`

	// NodeStateList
	// Constraints:
	//    - nullable
	NodeStateList []*WSGClusterBladeClusterStateNodeStateListType `json:"nodeStateList,omitempty" validate:"omitempty,dive"`
}

func NewWSGClusterBladeClusterState() *WSGClusterBladeClusterState {
	m := new(WSGClusterBladeClusterState)
	return m
}

type WSGClusterBladeClusterStateManagementServiceStateListType struct {
	// ManagementServiceState
	// management service state
	// Constraints:
	//    - nullable
	//    - oneof:[Out_Of_Service,In_Service]
	ManagementServiceState *string `json:"managementServiceState,omitempty" validate:"omitempty,oneof=Out_Of_Service In_Service"`

	// NodeId
	// Identifier of the controller node
	// Constraints:
	//    - nullable
	NodeId *string `json:"nodeId,omitempty"`

	// NodeName
	// node name
	// Constraints:
	//    - nullable
	NodeName *string `json:"nodeName,omitempty"`
}

func NewWSGClusterBladeClusterStateManagementServiceStateListType() *WSGClusterBladeClusterStateManagementServiceStateListType {
	m := new(WSGClusterBladeClusterStateManagementServiceStateListType)
	return m
}

type WSGClusterBladeClusterStateNodeStateListType struct {
	// NodeId
	// Identifier of the controller node
	// Constraints:
	//    - nullable
	NodeId *string `json:"nodeId,omitempty"`

	// NodeName
	// Constraints:
	//    - nullable
	NodeName *string `json:"nodeName,omitempty"`

	// NodeState
	// node state
	// Constraints:
	//    - nullable
	//    - oneof:[Out_Of_Service,In_Service]
	NodeState *string `json:"nodeState,omitempty" validate:"omitempty,oneof=Out_Of_Service In_Service"`
}

func NewWSGClusterBladeClusterStateNodeStateListType() *WSGClusterBladeClusterStateNodeStateListType {
	m := new(WSGClusterBladeClusterStateNodeStateListType)
	return m
}

type WSGClusterBladeClusterUpgradeProgress struct {
	// BladeProgresss
	// bladeProgressMap of clusterOperationProgress
	// Constraints:
	//    - nullable
	BladeProgresss []*WSGClusterBladeBladeProgress `json:"bladeProgresss,omitempty" validate:"omitempty,dive"`

	// ClusterOperationBlockUI
	// clusterOperationBlockUI of clusterOperationProgress
	// Constraints:
	//    - nullable
	ClusterOperationBlockUI *bool `json:"clusterOperationBlockUI,omitempty"`

	// ClusterOperationDisplayMsg
	// clusterOperationDisplayMsg of clusterOperationProgress
	// Constraints:
	//    - nullable
	ClusterOperationDisplayMsg *string `json:"clusterOperationDisplayMsg,omitempty"`

	// ClusterSubTaskState
	// clusterSubTaskState of clusterOperationProgress
	// Constraints:
	//    - nullable
	//    - oneof:[None,Running,Failed,Completed]
	ClusterSubTaskState *string `json:"clusterSubTaskState,omitempty" validate:"omitempty,oneof=None Running Failed Completed"`

	// IsSelfBladeRebooting
	// isSelfBladeRebooting of clusterOperationProgress
	// Constraints:
	//    - nullable
	IsSelfBladeRebooting *bool `json:"isSelfBladeRebooting,omitempty"`

	// Operation
	// Constraints:
	//    - nullable
	Operation *WSGClusterBladeOperation `json:"operation,omitempty"`

	// OverallProgress
	// overallProgress of clusterOperationProgress
	// Constraints:
	//    - nullable
	OverallProgress *int `json:"overallProgress,omitempty"`

	// PreviousOperationRecord
	// Constraints:
	//    - nullable
	PreviousOperationRecord *WSGClusterBladePreviousOperationRecord `json:"previousOperationRecord,omitempty"`
}

func NewWSGClusterBladeClusterUpgradeProgress() *WSGClusterBladeClusterUpgradeProgress {
	m := new(WSGClusterBladeClusterUpgradeProgress)
	return m
}

type WSGClusterBladeOperation string

func NewWSGClusterBladeOperation() *WSGClusterBladeOperation {
	m := new(WSGClusterBladeOperation)
	return m
}

type WSGClusterBladePreviousOperationRecord struct {
	// ErrorMsg
	// errorMsg of previousOperationRecord
	// Constraints:
	//    - nullable
	ErrorMsg *string `json:"errorMsg,omitempty"`

	// Operation
	// Constraints:
	//    - nullable
	Operation *WSGClusterBladeOperation `json:"operation,omitempty"`

	// Success
	// success of previousOperationRecord
	// Constraints:
	//    - nullable
	Success *bool `json:"success,omitempty"`
}

func NewWSGClusterBladePreviousOperationRecord() *WSGClusterBladePreviousOperationRecord {
	m := new(WSGClusterBladePreviousOperationRecord)
	return m
}

type WSGClusterBladeUploadPatchInfo struct {
	// AllowVersions
	// allowVersions of uploadPatchInfo
	// Constraints:
	//    - nullable
	AllowVersions []string `json:"allowVersions,omitempty" validate:"omitempty,dive"`

	// ApVersion
	// apVersion of uploadPatchInfo
	// Constraints:
	//    - nullable
	ApVersion *string `json:"apVersion,omitempty"`

	// ControlbladeVersion
	// controlbladeVersion of uploadPatchInfo
	// Constraints:
	//    - nullable
	ControlbladeVersion *string `json:"controlbladeVersion,omitempty"`

	// DatabladeVersion
	// databladeVersion of uploadPatchInfo
	// Constraints:
	//    - nullable
	DatabladeVersion *string `json:"databladeVersion,omitempty"`

	// FileName
	// fileName of uploadPatchInfo
	// Constraints:
	//    - nullable
	FileName *string `json:"fileName,omitempty"`

	// FileSize
	// fileSize of uploadPatchInfo
	// Constraints:
	//    - nullable
	FileSize *float64 `json:"fileSize,omitempty"`

	// FileUploadPath
	// fileUploadPath of uploadPatchInfo
	// Constraints:
	//    - nullable
	FileUploadPath *string `json:"fileUploadPath,omitempty"`

	// Version
	// version of uploadPatchInfo
	// Constraints:
	//    - nullable
	Version *string `json:"version,omitempty"`
}

func NewWSGClusterBladeUploadPatchInfo() *WSGClusterBladeUploadPatchInfo {
	m := new(WSGClusterBladeUploadPatchInfo)
	return m
}
