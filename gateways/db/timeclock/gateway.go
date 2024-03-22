package timeclock

import (
	"log"
	"time"
	timeClockAdapter "time-clock/adapters/timeclock"
	"time-clock/entities"

	"gorm.io/gorm"
)

type Repository struct {
	repository *gorm.DB
}

func NewGateway(repository *gorm.DB) *Repository {
	return &Repository{repository: repository}
}

func (r *Repository) ClockIn(userId uint32) (*entities.TimeClock, error) {
	timeClock := entities.TimeClock{
		UserID:    userId,
		ClockIn:   time.Now(),
		CreatedAt: time.Now(),
	}
	result := r.repository.Create(&timeClock)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &timeClock, nil
}

func (r *Repository) Report(userID uint32, startDate, endDate time.Time) ([]timeClockAdapter.TimeClock, error) {
	var timeClocks []timeClockAdapter.TimeClock

	result := r.repository.Raw(
		"SELECT "+
			"   date, "+
			"	MIN(clock_in) AS start_time, "+
			"	MAX(clock_out) AS end_time, "+
			"	SUM(EXTRACT(EPOCH FROM intervalo_trabalhado) / 3600) AS total_hours_worked "+
			" FROM "+
			" ( SELECT "+
			"   DATE(clock_in) AS date, "+
			"	clock_in, "+
			"	LEAD(clock_in) OVER (PARTITION BY DATE(clock_in) ORDER BY clock_in) AS clock_out, "+
			"	LEAD(clock_in) OVER (PARTITION BY DATE(clock_in) ORDER BY clock_in) - clock_in AS intervalo_trabalhado "+
			" FROM time_clocks WHERE deleted_at IS NULL AND DATE(clock_in) >= ? AND DATE(clock_in) < ? AND user_id = ? "+
			" ) AS intervalo_hora_trabalhada GROUP BY date ORDER BY date", startDate, endDate, userID).Scan(&timeClocks)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return timeClocks, nil
}
