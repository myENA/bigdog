package vsz

// API Version: v8_1

type WSGSyslogModifySyslogSettings struct {
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
	// Event Filter, 0 : All events, 1 : All events except client association/disassociation events, 2 : All events above a severity
	// Constraints:
	//    - min:0
	//    - max:2
	EventFilter *int `json:"eventFilter,omitempty" validate:"gte=0,lte=2"`

	// EventFilterSeverity
	// Event Filter Severity, This only applies when the Event Filter is set to 2
	// Constraints:
	//    - default:'Debug'
	//    - oneof:[Critical,Major,Minor,Warning,Informational,Debug]
	EventFilterSeverity *string `json:"eventFilterSeverity,omitempty" validate:"oneof=Critical Major Minor Warning Informational Debug"`

	// EventLogFacility
	// Facility for Event
	EventLogFacility *string `json:"eventLogFacility,omitempty"`

	// ForwardUEEventsMsgFormatType
	// forwardUEEventsMsgFormatType : COMMON, ZD
	ForwardUEEventsMsgFormatType *string `json:"forwardUEEventsMsgFormatType,omitempty"`

	// OtherLogSeverity
	// otherLogSeverity: Emerg, Alert, Crit, Error, Warning, Notice, Info, Debug
	OtherLogSeverity *string `json:"otherLogSeverity,omitempty"`

	PrimaryServer *WSGSyslogPrimaryServer `json:"primaryServer,omitempty"`

	Priority *WSGSyslogPriority `json:"priority,omitempty"`

	SecondaryServer *WSGSyslogSecondaryServer `json:"secondaryServer,omitempty"`
}

func NewWSGSyslogModifySyslogSettings() *WSGSyslogModifySyslogSettings {
	m := new(WSGSyslogModifySyslogSettings)
	return m
}

type WSGSyslogPrimaryServer struct {
	// Host
	// address of the syslog server.
	Host *string `json:"host,omitempty"`

	// Port
	// port number of the syslog server
	Port *int `json:"port,omitempty"`

	// Protocol
	// protocol of the syslog server
	// Constraints:
	//    - oneof:[UDP,TCP]
	Protocol *string `json:"protocol,omitempty" validate:"oneof=UDP TCP"`
}

func NewWSGSyslogPrimaryServer() *WSGSyslogPrimaryServer {
	m := new(WSGSyslogPrimaryServer)
	return m
}

type WSGSyslogPriority struct {
	// Critical
	// Event severity
	// Constraints:
	//    - oneof:[ERROR,WARN,INFO,DEBUG]
	Critical *string `json:"critical,omitempty" validate:"oneof=ERROR WARN INFO DEBUG"`

	// Debug
	// Event severity
	// Constraints:
	//    - oneof:[ERROR,WARN,INFO,DEBUG]
	Debug *string `json:"debug,omitempty" validate:"oneof=ERROR WARN INFO DEBUG"`

	// Informational
	// Event severity
	// Constraints:
	//    - oneof:[ERROR,WARN,INFO,DEBUG]
	Informational *string `json:"informational,omitempty" validate:"oneof=ERROR WARN INFO DEBUG"`

	// Major
	// Event severity
	// Constraints:
	//    - oneof:[ERROR,WARN,INFO,DEBUG]
	Major *string `json:"major,omitempty" validate:"oneof=ERROR WARN INFO DEBUG"`

	// Minor
	// Event severity
	// Constraints:
	//    - oneof:[ERROR,WARN,INFO,DEBUG]
	Minor *string `json:"minor,omitempty" validate:"oneof=ERROR WARN INFO DEBUG"`

	// Warning
	// Event severity
	// Constraints:
	//    - oneof:[ERROR,WARN,INFO,DEBUG]
	Warning *string `json:"warning,omitempty" validate:"oneof=ERROR WARN INFO DEBUG"`
}

func NewWSGSyslogPriority() *WSGSyslogPriority {
	m := new(WSGSyslogPriority)
	return m
}

type WSGSyslogSecondaryServer struct {
	// Host
	// address of the syslog server.
	Host *string `json:"host,omitempty"`

	// Port
	// port number of the syslog server
	Port *int `json:"port,omitempty"`

	// Protocol
	// protocol of the syslog server
	// Constraints:
	//    - oneof:[UDP,TCP]
	Protocol *string `json:"protocol,omitempty" validate:"oneof=UDP TCP"`

	// RedundancyMode
	// The redundancy mode of syslog server
	// Constraints:
	//    - oneof:[active_active,primary_backup]
	RedundancyMode *string `json:"redundancyMode,omitempty" validate:"oneof=active_active primary_backup"`
}

func NewWSGSyslogSecondaryServer() *WSGSyslogSecondaryServer {
	m := new(WSGSyslogSecondaryServer)
	return m
}

type WSGSyslogServerSetting struct {
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
	// Event Filter, 0 : All events, 1 : All events except client association/disassociation events, 2 : All events above a severity
	// Constraints:
	//    - min:0
	//    - max:2
	EventFilter *int `json:"eventFilter,omitempty" validate:"gte=0,lte=2"`

	// EventFilterSeverity
	// Event Filter Severity, This only applies when the Event Filter is set to 2
	// Constraints:
	//    - oneof:[Critical,Major,Minor,Warning,Informational,Debug]
	EventFilterSeverity *string `json:"eventFilterSeverity,omitempty" validate:"oneof=Critical Major Minor Warning Informational Debug"`

	// EventLogFacility
	// auditLogFacility: Local0-Local7
	EventLogFacility *string `json:"eventLogFacility,omitempty"`

	// ForwardUEEventsMsgFormatType
	// forwardUEEventsMsgFormatType : COMMON, ZD
	ForwardUEEventsMsgFormatType *string `json:"forwardUEEventsMsgFormatType,omitempty"`

	// OtherLogSeverity
	// otherLogSeverity: Emerg, Alert, Crit, Error, Warning, Notice, Info, Debug
	OtherLogSeverity *string `json:"otherLogSeverity,omitempty"`

	PrimaryServer *WSGSyslogPrimaryServer `json:"primaryServer,omitempty"`

	Priority *WSGSyslogPriority `json:"priority,omitempty"`

	SecondaryServer *WSGSyslogSecondaryServer `json:"secondaryServer,omitempty"`
}

func NewWSGSyslogServerSetting() *WSGSyslogServerSetting {
	m := new(WSGSyslogServerSetting)
	return m
}