package util

import "time"

func GetLastDayMonth(ano int, mes int) string {
	ultimoDia := time.Date(ano, time.Month(mes+1), 0, 0, 0, 0, 0, time.UTC)
	return ultimoDia.Format("2006-01-02")
}
