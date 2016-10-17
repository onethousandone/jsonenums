// generated by jsonenums -type=WeekDay -type-prefix Day -lowercase; DO NOT EDIT
// TypePrefix: Day
// ToLowerCase: true

package main

import (
	"encoding/json"
	"fmt"
)

var (
	_WeekDayNameToValue = map[string]WeekDay{

		"monday":    DayMonday,
		"tuesday":   DayTuesday,
		"wednesday": DayWednesday,
		"thursday":  DayThursday,
		"friday":    DayFriday,
		"saturday":  DaySaturday,
		"sunday":    DaySunday,
	}

	_WeekDayValueToName = map[WeekDay]string{

		DayMonday:    "monday",
		DayTuesday:   "tuesday",
		DayWednesday: "wednesday",
		DayThursday:  "thursday",
		DayFriday:    "friday",
		DaySaturday:  "saturday",
		DaySunday:    "sunday",
	}
)

func init() {
	var v WeekDay
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_WeekDayNameToValue = map[string]WeekDay{
			interface{}(DayMonday).(fmt.Stringer).String():    DayMonday,
			interface{}(DayTuesday).(fmt.Stringer).String():   DayTuesday,
			interface{}(DayWednesday).(fmt.Stringer).String(): DayWednesday,
			interface{}(DayThursday).(fmt.Stringer).String():  DayThursday,
			interface{}(DayFriday).(fmt.Stringer).String():    DayFriday,
			interface{}(DaySaturday).(fmt.Stringer).String():  DaySaturday,
			interface{}(DaySunday).(fmt.Stringer).String():    DaySunday,
		}
	}
}

// MarshalJSON is generated so WeekDay satisfies json.Marshaler.
func (r WeekDay) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _WeekDayValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid WeekDay: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so WeekDay satisfies json.Unmarshaler.
func (r *WeekDay) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("WeekDay should be a string, got %s", data)
	}
	v, ok := _WeekDayNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid WeekDay %q", s)
	}
	*r = v
	return nil
}
