package vsz

// API Version: v9_0

type WSGSyslogModifySyslogSettings struct {
	// AppLogFacility
	// appLogFacility: Local0-Local7
	// Constraints:
	//    - nullable
	AppLogFacility *string `json:"appLogFacility,omitempty"`

	// AppLogSeverity
	// appLogSeverity: Emerg, Alert, Crit, Error, Warning, Notice, Info, Debug
	// Constraints:
	//    - nullable
	AppLogSeverity *string `json:"appLogSeverity,omitempty"`

	// AuditLogFacility
	// auditLogFacility: Local0-Local7
	// Constraints:
	//    - nullable
	AuditLogFacility *string `json:"auditLogFacility,omitempty"`

	// AuditLogSeverity
	// auditLogSeverity: Emerg, Alert, Crit, Error, Warning, Notice, Info, Debug
	// Constraints:
	//    - nullable
	AuditLogSeverity *string `json:"auditLogSeverity,omitempty"`

	// Enabled
	// enable logging to remote syslog server
	// Constraints:
	//    - nullable
	Enabled *bool `json:"enabled,omitempty"`

	// EventFilter
	// Event Filter, 0 : All events, 1 : All events except client association/disassociation events, 2 : All events above a severity
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:2
	EventFilter *int `json:"eventFilter,omitempty" validate:"omitempty,gte=0,lte=2"`

	// EventFilterSeverity
	// Event Filter Severity, This only applies when the Event Filter is set to 2
	// Constraints:
	//    - nullable
	//    - default:'Debug'
	//    - oneof:[Critical,Major,Minor,Warning,Informational,Debug]
	EventFilterSeverity *string `json:"eventFilterSeverity,omitempty" validate:"omitempty,oneof=Critical Major Minor Warning Informational Debug"`

	// EventLogFacility
	// Facility for Event
	// Constraints:
	//    - nullable
	EventLogFacility *string `json:"eventLogFacility,omitempty"`

	// ForwardUEEventsMsgFormatType
	// forwardUEEventsMsgFormatType : COMMON, ZD
	// Constraints:
	//    - nullable
	ForwardUEEventsMsgFormatType *string `json:"forwardUEEventsMsgFormatType,omitempty"`

	// OtherLogSeverity
	// otherLogSeverity: Emerg, Alert, Crit, Error, Warning, Notice, Info, Debug
	// Constraints:
	//    - nullable
	OtherLogSeverity *string `json:"otherLogSeverity,omitempty"`

	// PrimaryServer
	// Constraints:
	//    - nullable
	PrimaryServer *WSGSyslogPrimaryServer `json:"primaryServer,omitempty"`

	// Priority
	// Constraints:
	//    - nullable
	Priority *WSGSyslogPriority `json:"priority,omitempty"`

	// SecondaryServer
	// Constraints:
	//    - nullable
	SecondaryServer *WSGSyslogSecondaryServer `json:"secondaryServer,omitempty"`
}

func NewWSGSyslogModifySyslogSettings() *WSGSyslogModifySyslogSettings {
	m := new(WSGSyslogModifySyslogSettings)
	return m
}

type WSGSyslogPrimaryServer struct {
	// Host
	// address of the syslog server.
	// Constraints:
	//    - nullable
	Host *string `json:"host,omitempty"`

	// Port
	// port number of the syslog server
	// Constraints:
	//    - nullable
	Port *int `json:"port,omitempty"`

	// Protocol
	// protocol of the syslog server
	// Constraints:
	//    - nullable
	//    - oneof:[UDP,TCP]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=UDP TCP"`
}

func NewWSGSyslogPrimaryServer() *WSGSyslogPrimaryServer {
	m := new(WSGSyslogPrimaryServer)
	return m
}

type WSGSyslogPriority struct {
	// Critical
	// Event severity
	// Constraints:
	//    - nullable
	//    - oneof:[ERROR,WARN,INFO,DEBUG]
	Critical *string `json:"critical,omitempty" validate:"omitempty,oneof=ERROR WARN INFO DEBUG"`

	// Debug
	// Event severity
	// Constraints:
	//    - nullable
	//    - oneof:[ERROR,WARN,INFO,DEBUG]
	Debug *string `json:"debug,omitempty" validate:"omitempty,oneof=ERROR WARN INFO DEBUG"`

	// Informational
	// Event severity
	// Constraints:
	//    - nullable
	//    - oneof:[ERROR,WARN,INFO,DEBUG]
	Informational *string `json:"informational,omitempty" validate:"omitempty,oneof=ERROR WARN INFO DEBUG"`

	// Major
	// Event severity
	// Constraints:
	//    - nullable
	//    - oneof:[ERROR,WARN,INFO,DEBUG]
	Major *string `json:"major,omitempty" validate:"omitempty,oneof=ERROR WARN INFO DEBUG"`

	// Minor
	// Event severity
	// Constraints:
	//    - nullable
	//    - oneof:[ERROR,WARN,INFO,DEBUG]
	Minor *string `json:"minor,omitempty" validate:"omitempty,oneof=ERROR WARN INFO DEBUG"`

	// Warning
	// Event severity
	// Constraints:
	//    - nullable
	//    - oneof:[ERROR,WARN,INFO,DEBUG]
	Warning *string `json:"warning,omitempty" validate:"omitempty,oneof=ERROR WARN INFO DEBUG"`
}

func NewWSGSyslogPriority() *WSGSyslogPriority {
	m := new(WSGSyslogPriority)
	return m
}

type WSGSyslogSecondaryServer struct {
	// Host
	// address of the syslog server.
	// Constraints:
	//    - nullable
	Host *string `json:"host,omitempty"`

	// Port
	// port number of the syslog server
	// Constraints:
	//    - nullable
	Port *int `json:"port,omitempty"`

	// Protocol
	// protocol of the syslog server
	// Constraints:
	//    - nullable
	//    - oneof:[UDP,TCP]
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,oneof=UDP TCP"`

	// RedundancyMode
	// The redundancy mode of syslog server
	// Constraints:
	//    - nullable
	//    - oneof:[active_active,primary_backup]
	RedundancyMode *string `json:"redundancyMode,omitempty" validate:"omitempty,oneof=active_active primary_backup"`
}

func NewWSGSyslogSecondaryServer() *WSGSyslogSecondaryServer {
	m := new(WSGSyslogSecondaryServer)
	return m
}

type WSGSyslogServerSetting struct {
	// AppLogFacility
	// appLogFacility: Local0-Local7
	// Constraints:
	//    - nullable
	AppLogFacility *string `json:"appLogFacility,omitempty"`

	// AppLogSeverity
	// appLogSeverity: Emerg, Alert, Crit, Error, Warning, Notice, Info, Debug
	// Constraints:
	//    - nullable
	AppLogSeverity *string `json:"appLogSeverity,omitempty"`

	// AuditLogFacility
	// auditLogFacility: Local0-Local7
	// Constraints:
	//    - nullable
	AuditLogFacility *string `json:"auditLogFacility,omitempty"`

	// AuditLogSeverity
	// auditLogSeverity: Emerg, Alert, Crit, Error, Warning, Notice, Info, Debug
	// Constraints:
	//    - nullable
	AuditLogSeverity *string `json:"auditLogSeverity,omitempty"`

	// Enabled
	// enable logging to remote syslog server
	// Constraints:
	//    - nullable
	Enabled *bool `json:"enabled,omitempty"`

	// EventFilter
	// Event Filter, 0 : All events, 1 : All events except client association/disassociation events, 2 : All events above a severity
	// Constraints:
	//    - nullable
	//    - min:0
	//    - max:2
	EventFilter *int `json:"eventFilter,omitempty" validate:"omitempty,gte=0,lte=2"`

	// EventFilterSeverity
	// Event Filter Severity, This only applies when the Event Filter is set to 2
	// Constraints:
	//    - nullable
	//    - oneof:[Critical,Major,Minor,Warning,Informational,Debug]
	EventFilterSeverity *string `json:"eventFilterSeverity,omitempty" validate:"omitempty,oneof=Critical Major Minor Warning Informational Debug"`

	// EventLogFacility
	// auditLogFacility: Local0-Local7
	// Constraints:
	//    - nullable
	EventLogFacility *string `json:"eventLogFacility,omitempty"`

	// ForwardUEEventsMsgFormatType
	// forwardUEEventsMsgFormatType : COMMON, ZD
	// Constraints:
	//    - nullable
	ForwardUEEventsMsgFormatType *string `json:"forwardUEEventsMsgFormatType,omitempty"`

	// OtherLogSeverity
	// otherLogSeverity: Emerg, Alert, Crit, Error, Warning, Notice, Info, Debug
	// Constraints:
	//    - nullable
	OtherLogSeverity *string `json:"otherLogSeverity,omitempty"`

	// PrimaryServer
	// Constraints:
	//    - nullable
	PrimaryServer *WSGSyslogPrimaryServer `json:"primaryServer,omitempty"`

	// Priority
	// Constraints:
	//    - nullable
	Priority *WSGSyslogPriority `json:"priority,omitempty"`

	// SecondaryServer
	// Constraints:
	//    - nullable
	SecondaryServer *WSGSyslogSecondaryServer `json:"secondaryServer,omitempty"`
}

func NewWSGSyslogServerSetting() *WSGSyslogServerSetting {
	m := new(WSGSyslogServerSetting)
	return m
}
