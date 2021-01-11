package bigdog

// API Version: v9_1

import (
	"errors"
	"io"
)

// WSGClusterBladeBladeProgress
//
// Definition: clusterblade_bladeProgress
type WSGClusterBladeBladeProgress struct {
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

func NewWSGClusterBladeBladeProgress() *WSGClusterBladeBladeProgress {
	m := new(WSGClusterBladeBladeProgress)
	return m
}

// WSGClusterBladeClusterOperationProgress
//
// Definition: clusterblade_clusterOperationProgress
type WSGClusterBladeClusterOperationProgress struct {
	Operation *WSGClusterBladeOperation `json:"operation,omitempty"`

	// OverallProgress
	// overallProgress of clusterOperationProgress
	OverallProgress *int `json:"overallProgress,omitempty"`

	PreviousOperationRecord *WSGClusterBladePreviousOperationRecord `json:"previousOperationRecord,omitempty"`
}

func NewWSGClusterBladeClusterOperationProgress() *WSGClusterBladeClusterOperationProgress {
	m := new(WSGClusterBladeClusterOperationProgress)
	return m
}

// WSGClusterBladeClusterState
//
// Definition: clusterblade_clusterState
type WSGClusterBladeClusterState struct {
	// ClusterName
	// cluster name
	ClusterName *string `json:"clusterName,omitempty"`

	// ClusterRole
	// The cluster role of the current controller node
	// Constraints:
	//    - oneof:[Leader,Follower]
	ClusterRole *string `json:"clusterRole,omitempty"`

	// ClusterState
	// cluster state
	// Constraints:
	//    - oneof:[In_Service,Out_Of_Service,Maintenance,Read_Only,NetworkPartitionSuspected]
	ClusterState *string `json:"clusterState,omitempty"`

	// CurrentNodeId
	// Identifier of the current controller node
	CurrentNodeId *string `json:"currentNodeId,omitempty"`

	// CurrentNodeName
	// The name of the current controller node
	CurrentNodeName *string `json:"currentNodeName,omitempty"`

	ManagementServiceStateList []*WSGClusterBladeClusterStateManagementServiceStateListType `json:"managementServiceStateList,omitempty"`

	NodeStateList []*WSGClusterBladeClusterStateNodeStateListType `json:"nodeStateList,omitempty"`
}

type WSGClusterBladeClusterStateAPIResponse struct {
	*RawAPIResponse
	Data *WSGClusterBladeClusterState
}

func newWSGClusterBladeClusterStateAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGClusterBladeClusterStateAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGClusterBladeClusterStateAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGClusterBladeClusterState)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGClusterBladeClusterState() *WSGClusterBladeClusterState {
	m := new(WSGClusterBladeClusterState)
	return m
}

// WSGClusterBladeClusterStateManagementServiceStateListType
//
// Definition: clusterblade_clusterStateManagementServiceStateListType
type WSGClusterBladeClusterStateManagementServiceStateListType struct {
	// ManagementServiceState
	// management service state
	// Constraints:
	//    - oneof:[Out_Of_Service,In_Service]
	ManagementServiceState *string `json:"managementServiceState,omitempty"`

	// NodeId
	// Identifier of the controller node
	NodeId *string `json:"nodeId,omitempty"`

	// NodeName
	// node name
	NodeName *string `json:"nodeName,omitempty"`
}

func NewWSGClusterBladeClusterStateManagementServiceStateListType() *WSGClusterBladeClusterStateManagementServiceStateListType {
	m := new(WSGClusterBladeClusterStateManagementServiceStateListType)
	return m
}

// WSGClusterBladeClusterStateNodeStateListType
//
// Definition: clusterblade_clusterStateNodeStateListType
type WSGClusterBladeClusterStateNodeStateListType struct {
	// NodeId
	// Identifier of the controller node
	NodeId *string `json:"nodeId,omitempty"`

	NodeName *string `json:"nodeName,omitempty"`

	// NodeState
	// node state
	// Constraints:
	//    - oneof:[Out_Of_Service,In_Service]
	NodeState *string `json:"nodeState,omitempty"`
}

func NewWSGClusterBladeClusterStateNodeStateListType() *WSGClusterBladeClusterStateNodeStateListType {
	m := new(WSGClusterBladeClusterStateNodeStateListType)
	return m
}

// WSGClusterBladeClusterUpgradeProgress
//
// Definition: clusterblade_clusterUpgradeProgress
type WSGClusterBladeClusterUpgradeProgress struct {
	// BladeProgresss
	// bladeProgressMap of clusterOperationProgress
	BladeProgresss []*WSGClusterBladeBladeProgress `json:"bladeProgresss,omitempty"`

	// ClusterOperationBlockUI
	// clusterOperationBlockUI of clusterOperationProgress
	ClusterOperationBlockUI *bool `json:"clusterOperationBlockUI,omitempty"`

	// ClusterOperationDisplayMsg
	// clusterOperationDisplayMsg of clusterOperationProgress
	ClusterOperationDisplayMsg *string `json:"clusterOperationDisplayMsg,omitempty"`

	// ClusterSubTaskState
	// clusterSubTaskState of clusterOperationProgress
	// Constraints:
	//    - oneof:[None,Running,Failed,Completed]
	ClusterSubTaskState *string `json:"clusterSubTaskState,omitempty"`

	// IsSelfBladeRebooting
	// isSelfBladeRebooting of clusterOperationProgress
	IsSelfBladeRebooting *bool `json:"isSelfBladeRebooting,omitempty"`

	Operation *WSGClusterBladeOperation `json:"operation,omitempty"`

	// OverallProgress
	// overallProgress of clusterOperationProgress
	OverallProgress *int `json:"overallProgress,omitempty"`

	PreviousOperationRecord *WSGClusterBladePreviousOperationRecord `json:"previousOperationRecord,omitempty"`
}

func NewWSGClusterBladeClusterUpgradeProgress() *WSGClusterBladeClusterUpgradeProgress {
	m := new(WSGClusterBladeClusterUpgradeProgress)
	return m
}

// WSGClusterBladeOperation
//
// Definition: clusterblade_operation
//
// Constraints:
//    - oneof:[None,Join,Backup,Restore,Upload,Upgrade,BackupAndUpgrade,RemoveBlade,ChangeBladeIP,DiskSnapshot,UploadVDPFirmware,UploadAPFirmware,AddAPFirmware,AddLteFirmware]
type WSGClusterBladeOperation string

func NewWSGClusterBladeOperation() *WSGClusterBladeOperation {
	m := new(WSGClusterBladeOperation)
	return m
}

// WSGClusterBladePreviousOperationRecord
//
// Definition: clusterblade_previousOperationRecord
type WSGClusterBladePreviousOperationRecord struct {
	// ErrorMsg
	// errorMsg of previousOperationRecord
	ErrorMsg *string `json:"errorMsg,omitempty"`

	Operation *WSGClusterBladeOperation `json:"operation,omitempty"`

	// Success
	// success of previousOperationRecord
	Success *bool `json:"success,omitempty"`
}

func NewWSGClusterBladePreviousOperationRecord() *WSGClusterBladePreviousOperationRecord {
	m := new(WSGClusterBladePreviousOperationRecord)
	return m
}

// WSGClusterBladeUploadPatchInfo
//
// Definition: clusterblade_uploadPatchInfo
type WSGClusterBladeUploadPatchInfo struct {
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

func NewWSGClusterBladeUploadPatchInfo() *WSGClusterBladeUploadPatchInfo {
	m := new(WSGClusterBladeUploadPatchInfo)
	return m
}
