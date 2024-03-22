package timeclock

import (
	"fmt"
	"strings"
	"time"
)

type TimeClock struct {
	Date             time.Time `json:"date"`
	StartTime        time.Time `json:"start_time"`
	EndTime          time.Time `json:"end_time"`
	TotalHoursWorked float64   `json:"total_hours_worked"`
}

func FormatTimeClocksForEmail(timeClocks []TimeClock) string {
	var builder strings.Builder

	for _, tc := range timeClocks {
		timeClockStr := fmt.Sprintf("Data: %s\n", tc.Date.Format("2006-01-02"))
		timeClockStr += fmt.Sprintf("Hora de início: %s\n", tc.StartTime.Format("15:04:05"))
		timeClockStr += fmt.Sprintf("Hora de término: %s\n", tc.EndTime.Format("15:04:05"))
		timeClockStr += fmt.Sprintf("Total de horas trabalhadas: %.2f\n", tc.TotalHoursWorked)

		builder.WriteString(timeClockStr)
		builder.WriteString("\n")
	}

	return builder.String()
}
