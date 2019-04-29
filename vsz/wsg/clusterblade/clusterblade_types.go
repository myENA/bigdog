package clusterblade

// API Version: v8_0

type BladeProgress struct {
	BladeUUID     *string `json:"bladeUUID,omitempty"`
	HostName      *string `json:"hostName,omitempty"`
	IterationName *string `json:"iterationName,omitempty"`
	Progress      *int    `json:"progress,omitempty"`
	State         *string `json:"state,omitempty"`
}

type ClusterApPatchOperationProgress struct {
	Operation               *string                  `json:"operation,omitempty"`
	OverallProgress         *int                     `json:"overallProgress,omitempty"`
	PreviousOperationRecord *PreviousOperationRecord `json:"previousOperationRecord,omitempty"`
}

type ClusterOperationProgress struct {
	BladeProgresss             []*BladeProgresss        `json:"bladeProgresss,omitempty"`
	ClusterOperationBlockUI    *bool                    `json:"clusterOperationBlockUI,omitempty"`
	ClusterOperationDisplayMsg *string                  `json:"clusterOperationDisplayMsg,omitempty"`
	ClusterSubTaskState        *string                  `json:"clusterSubTaskState,omitempty"`
	IsSelfBladeRebooting       *bool                    `json:"isSelfBladeRebooting,omitempty"`
	Operation                  *string                  `json:"operation,omitempty"`
	OverallProgress            *int                     `json:"overallProgress,omitempty"`
	PreviousOperationRecord    *PreviousOperationRecord `json:"previousOperationRecord,omitempty"`
}

type ClusterState struct {
	ClusterName                *string                       `json:"clusterName,omitempty"`
	ClusterRole                *string                       `json:"clusterRole,omitempty"`
	ClusterState               *string                       `json:"clusterState,omitempty"`
	CurrentNodeID              *string                       `json:"currentNodeId,omitempty"`
	CurrentNodeName            *string                       `json:"currentNodeName,omitempty"`
	ManagementServiceStateList []*ManagementServiceStateList `json:"managementServiceStateList,omitempty"`
	NodeStateList              []*NodeStateList              `json:"nodeStateList,omitempty"`
}

type ClusterStateManagementServiceStateListType struct {
	ManagementServiceState *string `json:"managementServiceState,omitempty"`
	NodeID                 *string `json:"nodeId,omitempty"`
	NodeName               *string `json:"nodeName,omitempty"`
}

type ClusterStateNodeStateListType struct {
	NodeID    *string `json:"nodeId,omitempty"`
	NodeName  *string `json:"nodeName,omitempty"`
	NodeState *string `json:"nodeState,omitempty"`
}

type ClusterStatus struct {
	ClusterStatus *string `json:"clusterStatus,omitempty"`
}

type ControlNodeStatus struct {
	NodeStatusList []*NodeStatusList `json:"nodeStatusList,omitempty"`
}

type ControlNodeStatusNodeStatusListType struct {
	NodeID     *string `json:"nodeId,omitempty"`
	NodeStatus *string `json:"nodeStatus,omitempty"`
}

type PreviousOperationRecord struct {
	ErrorMsg  *string `json:"errorMsg,omitempty"`
	Operation *string `json:"operation,omitempty"`
	Success   *bool   `json:"success,omitempty"`
}

type UploadPatchInfo struct {
	AllowVersions       []string `json:"allowVersions,omitempty"`
	ApVersion           *string  `json:"apVersion,omitempty"`
	ControlbladeVersion *string  `json:"controlbladeVersion,omitempty"`
	DatabladeVersion    *string  `json:"databladeVersion,omitempty"`
	FileName            *string  `json:"fileName,omitempty"`
	FileSize            *float64 `json:"fileSize,omitempty"`
	FileUploadPath      *string  `json:"fileUploadPath,omitempty"`
	Version             *string  `json:"version,omitempty"`
}
