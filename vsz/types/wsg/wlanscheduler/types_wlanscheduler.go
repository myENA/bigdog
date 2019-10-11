package wlanscheduler

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type CreateWlanScheduler struct {
	Description *common.Description `json:"description,omitempty"`

	// Fri
	// Schedules on Friday
	Fri []string `json:"fri,omitempty"`

	// Mon
	// Schedules on Monday
	Mon []string `json:"mon,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

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

func NewCreateWlanScheduler() *CreateWlanScheduler {
	createWlanSchedulerType := new(CreateWlanScheduler)
	return createWlanSchedulerType
}

func NewCreateWlanSchedulerWithDefaults() *CreateWlanScheduler {
	createWlanSchedulerType := new(CreateWlanScheduler)
	return createWlanSchedulerType
}

type ModifyWlanScheduler struct {
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

func NewModifyWlanScheduler() *ModifyWlanScheduler {
	modifyWlanSchedulerType := new(ModifyWlanScheduler)
	return modifyWlanSchedulerType
}

func NewModifyWlanSchedulerWithDefaults() *ModifyWlanScheduler {
	modifyWlanSchedulerType := new(ModifyWlanScheduler)
	return modifyWlanSchedulerType
}

type WlanSchedule struct {
	Description *common.Description `json:"description,omitempty"`

	// Fri
	// Schedules on Friday
	Fri []string `json:"fri,omitempty"`

	// Id
	// Identifier of the WLAN schedule
	Id *string `json:"id,omitempty"`

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

	// ZoneId
	// Identifier of the zone to which the WLAN schedule belongs
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWlanSchedule() *WlanSchedule {
	wlanScheduleType := new(WlanSchedule)
	return wlanScheduleType
}

func NewWlanScheduleWithDefaults() *WlanSchedule {
	wlanScheduleType := new(WlanSchedule)
	return wlanScheduleType
}

type WlanScheduleList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WlanSchedule `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWlanScheduleList() *WlanScheduleList {
	wlanScheduleListType := new(WlanScheduleList)
	return wlanScheduleListType
}

func NewWlanScheduleListWithDefaults() *WlanScheduleList {
	wlanScheduleListType := new(WlanScheduleList)
	return wlanScheduleListType
}
