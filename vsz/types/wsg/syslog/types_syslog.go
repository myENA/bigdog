package syslog

// API Version: v8_1

type ModifySyslogSettings struct {
	// AppLogFacility
	// appLogFacility: Local0-Local7
	AppLogFacility *string `json:"appLogFacility,omitempty"`

	// AppLogSeverity
	// appLogSeverity: Emerg, Alert, Crit, Error, Warning, Notice, Info, Debug
	AppLogSeverity *string `json:"appLogSeverity,omitempty"`

	// AuditLogFacility
	// auditLogFacility: Local0-Local7
	AuditLogFacility *string `json:"auditLogFacility,omitempty"`

	// AuditLogSeverity
	// auditLogSeverity: Emerg, Alert, Crit, Error, Warning, Notice, Info, Debug
	AuditLogSeverity *string `json:"auditLogSeverity,omitempty"`

	// Enabled
	// enable logging to remote syslog server
	Enabled *bool `json:"enabled,omitempty"`

	// EventFilter
	// Event Filter, 0 : All events, 1 : All events except client association/disassociation events, 2 : All
	// events above a severity
	EventFilter *int `json:"eventFilter,omitempty" validate:"omitempty,gte=0,lte=2"`

	// EventFilterSeverity
	// Event Filter Severity, This only applies when the Event Filter is set to 2
	EventFilterSeverity *string `json:"eventFilterSeverity,omitempty" validate:"omitempty,oneof=Critical Major Minor Warning Informational Debug"`

	// EventLogFacility
	// Facility for Event
	EventLogFacility *string `json:"eventLogFacility,omitempty"`

	// ForwardUEEventsMsgFormatType
	// forwardUEEventsMsgFormatType : COMMON, ZD
	ForwardUEEventsMsgFormatType *string `json:"forwardUEEventsMsgFormatType,omitempty"`

	// OtherLogSeverity
	// otherLogSeverity: Emerg, Alert, Crit, Error, Warning, Notice, Info, Debug
	OtherLogSeverity *string `json:"otherLogSeverity,omitempty"`

	PrimaryServer *PrimaryServer `json:"primaryServer,omitempty"`

	Priority *Priority `json:"priority,omitempty"`

	SecondaryServer *SecondaryServer `json:"secondaryServer,omitempty"`
}

type PrimaryServer struct {
	// Host
	// address of the syslog server.
	Host *string `json:"host,omitempty"`

	// Port
	// port number of the syslog server
	Port *int `json:"port,omitempty"`

	// Protocol
	// protocol of the syslog server
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=UDP TCP"`
}

type Priority struct {
	// Critical
	// Event severity
	Critical *string `json:"critical,omitempty" validate:"omitempty,oneof=ERROR WARN INFO DEBUG"`

	// Debug
	// Event severity
	Debug *string `json:"debug,omitempty" validate:"omitempty,oneof=ERROR WARN INFO DEBUG"`

	// Informational
	// Event severity
	Informational *string `json:"informational,omitempty" validate:"omitempty,oneof=ERROR WARN INFO DEBUG"`

	// Major
	// Event severity
	Major *string `json:"major,omitempty" validate:"omitempty,oneof=ERROR WARN INFO DEBUG"`

	// Minor
	// Event severity
	Minor *string `json:"minor,omitempty" validate:"omitempty,oneof=ERROR WARN INFO DEBUG"`

	// Warning
	// Event severity
	Warning *string `json:"warning,omitempty" validate:"omitempty,oneof=ERROR WARN INFO DEBUG"`
}

type SecondaryServer struct {
	// Host
	// address of the syslog server.
	Host *string `json:"host,omitempty"`

	// Port
	// port number of the syslog server
	Port *int `json:"port,omitempty"`

	// Protocol
	// protocol of the syslog server
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=UDP TCP"`

	// RedundancyMode
	// The redundancy mode of syslog server
	RedundancyMode *string `json:"redundancyMode,omitempty" validate:"omitempty,oneof=active_active primary_backup"`
}

type SyslogServerSetting struct {
	// AppLogFacility
	// appLogFacility: Local0-Local7
	AppLogFacility *string `json:"appLogFacility,omitempty"`

	// AppLogSeverity
	// appLogSeverity: Emerg, Alert, Crit, Error, Warning, Notice, Info, Debug
	AppLogSeverity *string `json:"appLogSeverity,omitempty"`

	// AuditLogFacility
	// auditLogFacility: Local0-Local7
	AuditLogFacility *string `json:"auditLogFacility,omitempty"`

	// AuditLogSeverity
	// auditLogSeverity: Emerg, Alert, Crit, Error, Warning, Notice, Info, Debug
	AuditLogSeverity *string `json:"auditLogSeverity,omitempty"`

	// Enabled
	// enable logging to remote syslog server
	Enabled *bool `json:"enabled,omitempty"`

	// EventFilter
	// Event Filter, 0 : All events, 1 : All events except client association/disassociation events, 2 : All
	// events above a severity
	EventFilter *int `json:"eventFilter,omitempty" validate:"omitempty,gte=0,lte=2"`

	// EventFilterSeverity
	// Event Filter Severity, This only applies when the Event Filter is set to 2
	EventFilterSeverity *string `json:"eventFilterSeverity,omitempty" validate:"omitempty,oneof=Critical Major Minor Warning Informational Debug"`

	// EventLogFacility
	// auditLogFacility: Local0-Local7
	EventLogFacility *string `json:"eventLogFacility,omitempty"`

	// ForwardUEEventsMsgFormatType
	// forwardUEEventsMsgFormatType : COMMON, ZD
	ForwardUEEventsMsgFormatType *string `json:"forwardUEEventsMsgFormatType,omitempty"`

	// OtherLogSeverity
	// otherLogSeverity: Emerg, Alert, Crit, Error, Warning, Notice, Info, Debug
	OtherLogSeverity *string `json:"otherLogSeverity,omitempty"`

	PrimaryServer *PrimaryServer `json:"primaryServer,omitempty"`

	Priority *Priority `json:"priority,omitempty"`

	SecondaryServer *SecondaryServer `json:"secondaryServer,omitempty"`
}
