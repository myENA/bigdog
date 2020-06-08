package ruckus

// API Version: 1.0.0

type SCIPciprofileprototypecountreports200ResponseType struct {
	Count *float64 `json:"count,omitempty"`
}

func NewSCIPciprofileprototypecountreports200ResponseType() *SCIPciprofileprototypecountreports200ResponseType {
	m := new(SCIPciprofileprototypecountreports200ResponseType)
	return m
}

type SCIPciprofileprototypegetreports200ResponseType []*SCIModelsPciReport

func MakeSCIPciprofileprototypegetreports200ResponseType() SCIPciprofileprototypegetreports200ResponseType {
	m := make(SCIPciprofileprototypegetreports200ResponseType, 0)
	return m
}
