package booking

import (
	"time"
	"fmt"
)

// Schedule parses a textual representation of an appointment date into time.Time.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}
	}
	return t
}

// HasPassed checks if the given appointment date was in the past.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		return false
	}
	return t.Before(time.Now())
}

// IsAfternoonAppointment checks if the appointment is in the afternoon (>= 12:00 and < 18:00).
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		return false
	}
	hour := t.Hour()
	return hour >= 12 && hour < 18
}

// Description returns a description of the appointment date and time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		return ""
	}
	formatted := t.Format("Monday, January 2, 2006, at 15:04.")
	return fmt.Sprintf("You have an appointment on %s", formatted)
}

// AnniversaryDate returns the anniversary date of the salon's opening for the current year in UTC.
func AnniversaryDate() time.Time {
	year := time.Now().Year()
	return time.Date(year, time.September, 15, 0, 0, 0, 0, time.UTC)
}
