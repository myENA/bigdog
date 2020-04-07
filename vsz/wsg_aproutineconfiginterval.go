package vsz

// API Version: v9_0

type WSGAPRoutineConfigIntervalReq struct {
	// ApRoutineConfigInterval
	// Constraints:
	//    - nullable
	ApRoutineConfigInterval *int `json:"apRoutineConfigInterval,omitempty"`
}

func NewWSGAPRoutineConfigIntervalReq() *WSGAPRoutineConfigIntervalReq {
	m := new(WSGAPRoutineConfigIntervalReq)
	return m
}

type WSGAPRoutineConfigIntervalRsp struct {
	// ApRoutineConfigInterval
	// Constraints:
	//    - nullable
	ApRoutineConfigInterval *int `json:"apRoutineConfigInterval,omitempty"`

	// LowerBound
	// Constraints:
	//    - nullable
	LowerBound *int `json:"lowerBound,omitempty"`

	// UpperBound
	// Constraints:
	//    - nullable
	UpperBound *int `json:"upperBound,omitempty"`
}

func NewWSGAPRoutineConfigIntervalRsp() *WSGAPRoutineConfigIntervalRsp {
	m := new(WSGAPRoutineConfigIntervalRsp)
	return m
}
