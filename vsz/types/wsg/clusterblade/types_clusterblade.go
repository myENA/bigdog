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

func NewBladeProgress() *BladeProgress {
	bladeProgressType := new(BladeProgress)
	return bladeProgressType
}

func NewBladeProgressWithDefaults() *BladeProgress {
	bladeProgressType := new(BladeProgress)
	return bladeProgressType
}

type ClusterOperationProgress struct {
	Operation *Operation `json:"operation,omitempty"`

	// OverallProgress
	// overallProgress of clusterOperationProgress
	OverallProgress *int `json:"overallProgress,omitempty"`

	PreviousOperationRecord *PreviousOperationRecord `json:"previousOperationRecord,omitempty"`
}

func NewClusterOperationProgress() *ClusterOperationProgress {
	clusterOperationProgressType := new(ClusterOperationProgress)
	return clusterOperationProgressType
}

func NewClusterOperationProgressWithDefaults() *ClusterOperationProgress {
	clusterOperationProgressType := new(ClusterOperationProgress)
	return clusterOperationProgressType
}

type ClusterState struct {
	// ClusterName
	// cluster name
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
	CurrentNodeId *string `json:"currentNodeId,omitempty"`

	// CurrentNodeName
	// The name of the current controller node
	CurrentNodeName *string `json:"currentNodeName,omitempty"`

	ManagementServiceStateList []*ClusterStateManagementServiceStateListType `json:"managementServiceStateList,omitempty"`

	NodeStateList []*ClusterStateNodeStateListType `json:"nodeStateList,omitempty"`
}

func NewClusterState() *ClusterState {
	clusterStateType := new(ClusterState)
	return clusterStateType
}

func NewClusterStateWithDefaults() *ClusterState {
	clusterStateType := new(ClusterState)
	return clusterStateType
}

type ClusterStateManagementServiceStateListType struct {
	// ManagementServiceState
	// management service state
	// Constraints:
	//    - nullable
	//    - oneof:[Out_Of_Service,In_Service]
	ManagementServiceState *string `json:"managementServiceState,omitempty" validate:"omitempty,oneof=Out_Of_Service In_Service"`

	// NodeId
	// Identifier of the controller node
	NodeId *string `json:"nodeId,omitempty"`

	// NodeName
	// node name
	NodeName *string `json:"nodeName,omitempty"`
}

func NewClusterStateManagementServiceStateListType() *ClusterStateManagementServiceStateListType {
	clusterStateManagementServiceStateListTypeType := new(ClusterStateManagementServiceStateListType)
	return clusterStateManagementServiceStateListTypeType
}

func NewClusterStateManagementServiceStateListTypeWithDefaults() *ClusterStateManagementServiceStateListType {
	clusterStateManagementServiceStateListTypeType := new(ClusterStateManagementServiceStateListType)
	return clusterStateManagementServiceStateListTypeType
}

type ClusterStateNodeStateListType struct {
	// NodeId
	// Identifier of the controller node
	NodeId *string `json:"nodeId,omitempty"`

	NodeName *string `json:"nodeName,omitempty"`

	// NodeState
	// node state
	// Constraints:
	//    - nullable
	//    - oneof:[Out_Of_Service,In_Service]
	NodeState *string `json:"nodeState,omitempty" validate:"omitempty,oneof=Out_Of_Service In_Service"`
}

func NewClusterStateNodeStateListType() *ClusterStateNodeStateListType {
	clusterStateNodeStateListTypeType := new(ClusterStateNodeStateListType)
	return clusterStateNodeStateListTypeType
}

func NewClusterStateNodeStateListTypeWithDefaults() *ClusterStateNodeStateListType {
	clusterStateNodeStateListTypeType := new(ClusterStateNodeStateListType)
	return clusterStateNodeStateListTypeType
}

type ClusterStatus struct {
	// ClusterStatus
	// progress of bladeProgress
	// Constraints:
	//    - nullable
	//    - oneof:[In_Service,Out_Of_Service,Maintenance,Read_Only,NetworkPartitionSuspected]
	ClusterStatus *string `json:"clusterStatus,omitempty" validate:"omitempty,oneof=In_Service Out_Of_Service Maintenance Read_Only NetworkPartitionSuspected"`
}

func NewClusterStatus() *ClusterStatus {
	clusterStatusType := new(ClusterStatus)
	return clusterStatusType
}

func NewClusterStatusWithDefaults() *ClusterStatus {
	clusterStatusType := new(ClusterStatus)
	return clusterStatusType
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
	// Constraints:
	//    - nullable
	//    - oneof:[None,Running,Failed,Completed]
	ClusterSubTaskState *string `json:"clusterSubTaskState,omitempty" validate:"omitempty,oneof=None Running Failed Completed"`

	// IsSelfBladeRebooting
	// isSelfBladeRebooting of clusterOperationProgress
	IsSelfBladeRebooting *bool `json:"isSelfBladeRebooting,omitempty"`

	Operation *Operation `json:"operation,omitempty"`

	// OverallProgress
	// overallProgress of clusterOperationProgress
	OverallProgress *int `json:"overallProgress,omitempty"`

	PreviousOperationRecord *PreviousOperationRecord `json:"previousOperationRecord,omitempty"`
}

func NewClusterUpgradeProgress() *ClusterUpgradeProgress {
	clusterUpgradeProgressType := new(ClusterUpgradeProgress)
	return clusterUpgradeProgressType
}

func NewClusterUpgradeProgressWithDefaults() *ClusterUpgradeProgress {
	clusterUpgradeProgressType := new(ClusterUpgradeProgress)
	return clusterUpgradeProgressType
}

type ControlNodeStatus struct {
	NodeStatusList []*ControlNodeStatusNodeStatusListType `json:"nodeStatusList,omitempty"`
}

func NewControlNodeStatus() *ControlNodeStatus {
	controlNodeStatusType := new(ControlNodeStatus)
	return controlNodeStatusType
}

func NewControlNodeStatusWithDefaults() *ControlNodeStatus {
	controlNodeStatusType := new(ControlNodeStatus)
	return controlNodeStatusType
}

type ControlNodeStatusNodeStatusListType struct {
	// NodeId
	// Identifier of the controller node
	NodeId *string `json:"nodeId,omitempty"`

	// NodeStatus
	// node status
	// Constraints:
	//    - nullable
	//    - oneof:[Out_Of_Service,Bootstrapping,Got_WSG_Version,WSG_FW_Upgrading,Initializing_Database,Syncing_Configurations,Changing_Configurations,Launching_Apps,In_Service,Shutting_Down_Apps]
	NodeStatus *string `json:"nodeStatus,omitempty" validate:"omitempty,oneof=Out_Of_Service Bootstrapping Got_WSG_Version WSG_FW_Upgrading Initializing_Database Syncing_Configurations Changing_Configurations Launching_Apps In_Service Shutting_Down_Apps"`
}

func NewControlNodeStatusNodeStatusListType() *ControlNodeStatusNodeStatusListType {
	controlNodeStatusNodeStatusListTypeType := new(ControlNodeStatusNodeStatusListType)
	return controlNodeStatusNodeStatusListTypeType
}

func NewControlNodeStatusNodeStatusListTypeWithDefaults() *ControlNodeStatusNodeStatusListType {
	controlNodeStatusNodeStatusListTypeType := new(ControlNodeStatusNodeStatusListType)
	return controlNodeStatusNodeStatusListTypeType
}

type Operation string

func NewOperation() *Operation {
	operationType := new(Operation)
	return operationType
}

func NewOperationWithDefaults() *Operation {
	operationType := new(Operation)
	return operationType
}

type PreviousOperationRecord struct {
	// ErrorMsg
	// errorMsg of previousOperationRecord
	ErrorMsg *string `json:"errorMsg,omitempty"`

	Operation *Operation `json:"operation,omitempty"`

	// Success
	// success of previousOperationRecord
	Success *bool `json:"success,omitempty"`
}

func NewPreviousOperationRecord() *PreviousOperationRecord {
	previousOperationRecordType := new(PreviousOperationRecord)
	return previousOperationRecordType
}

func NewPreviousOperationRecordWithDefaults() *PreviousOperationRecord {
	previousOperationRecordType := new(PreviousOperationRecord)
	return previousOperationRecordType
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

func NewUploadPatchInfo() *UploadPatchInfo {
	uploadPatchInfoType := new(UploadPatchInfo)
	return uploadPatchInfoType
}

func NewUploadPatchInfoWithDefaults() *UploadPatchInfo {
	uploadPatchInfoType := new(UploadPatchInfo)
	return uploadPatchInfoType
}
