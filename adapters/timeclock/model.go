package timeclock

import "time"

type TimeClock struct {
	Date             time.Time `json:"date"`
	StartTime        time.Time `json:"start_time"`
	EndTime          time.Time `json:"end_time"`
	TotalHoursWorked float64   `json:"total_hours_worked"`
}
