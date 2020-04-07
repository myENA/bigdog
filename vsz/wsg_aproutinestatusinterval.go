package vsz

// API Version: v9_0

type WSGAPRoutineStatusIntervalRsp struct {
	// ApRoutineStatusInterval
	// Constraints:
	//    - nullable
	ApRoutineStatusInterval *int `json:"apRoutineStatusInterval,omitempty"`
}

func NewWSGAPRoutineStatusIntervalRsp() *WSGAPRoutineStatusIntervalRsp {
	m := new(WSGAPRoutineStatusIntervalRsp)
	return m
}
