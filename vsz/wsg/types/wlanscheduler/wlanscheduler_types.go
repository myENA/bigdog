package wlanscheduler

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/wsg/types/common"
)

type CreateWLANScheduler struct {
	Description *common.Description `json:"description,omitempty"`

	// Fri
	// Schedules on Friday
	Fri []string `json:"fri,omitempty"`

	// Mon
	// Schedules on Monday
	Mon []string `json:"mon,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Sat
	// Schedules on Saturday
	Sat []string `json:"sat,omitempty"`

	// Sun
	// Schedules on Sunday
	Sun []string `json:"sun,omitempty"`

	// Thu
	// Schedules on Thursday
	Thu []string `json:"thu,omitempty"`

	// Tue
	// Schedules on Tuesday
	Tue []string `json:"tue,omitempty"`

	// Wed
	// Schedules on Wednesday
	Wed []string `json:"wed,omitempty"`
}

type ModifyWLANScheduler struct {
	Description *common.Description `json:"description,omitempty"`

	// Fri
	// Schedules on Friday
	Fri []string `json:"fri,omitempty"`

	// Mon
	// Schedules on Monday
	Mon []string `json:"mon,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Sat
	// Schedules on Saturday
	Sat []string `json:"sat,omitempty"`

	// Sun
	// Schedules on Sunday
	Sun []string `json:"sun,omitempty"`

	// Thu
	// Schedules on Thursday
	Thu []string `json:"thu,omitempty"`

	// Tue
	// Schedules on Tuesday
	Tue []string `json:"tue,omitempty"`

	// Wed
	// schedules on Wednesday
	Wed []string `json:"wed,omitempty"`
}

type WLANSchedule struct {
	Description *common.Description `json:"description,omitempty"`

	// Fri
	// Schedules on Friday
	Fri []string `json:"fri,omitempty"`

	// ID
	// Identifier of the WLAN schedule
	ID *string `json:"id,omitempty"`

	// Mon
	// Schedules on Monday
	Mon []string `json:"mon,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Sat
	// Schedules on Saturday
	Sat []string `json:"sat,omitempty"`

	// Sun
	// Schedules on Sunday
	Sun []string `json:"sun,omitempty"`

	// Thu
	// Schedules on Thursday
	Thu []string `json:"thu,omitempty"`

	// Tue
	// Schedules on Tuesday
	Tue []string `json:"tue,omitempty"`

	// Wed
	// Schedules on Wednesday
	Wed []string `json:"wed,omitempty"`

	// ZoneID
	// Identifier of the zone to which the WLAN schedule belongs
	ZoneID *string `json:"zoneId,omitempty"`
}

type WLANScheduleList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WLANSchedule `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}
