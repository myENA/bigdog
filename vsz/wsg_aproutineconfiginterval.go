package vsz

// API Version: v8_1

type WSGAPRoutineConfigIntervalReq struct {
	ApRoutineConfigInterval *int `json:"apRoutineConfigInterval,omitempty"`
}

type WSGAPRoutineConfigIntervalRsp struct {
	ApRoutineConfigInterval *int `json:"apRoutineConfigInterval,omitempty"`

	LowerBound *int `json:"lowerBound,omitempty"`

	UpperBound *int `json:"upperBound,omitempty"`
}
