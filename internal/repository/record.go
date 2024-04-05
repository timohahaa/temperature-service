package repository

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/timohahaa/postgres"
	"github.com/timohahaa/temperature-service/internal/entity"
)

type recordRepository struct {
	db  *postgres.Postgres
	log *logrus.Logger
}

func NewPostgresRepository(pg *postgres.Postgres, logger *logrus.Logger) *recordRepository {
	return &recordRepository{
		db:  pg,
		log: logger,
	}
}

func (r *recordRepository) SaveRecord(ctx context.Context, record entity.Record) error {
	sql, args, _ := r.db.Builder.
		Insert("records").
		Columns("timestamp", "value").
		Values(record.Timestamp, record.Value).
		ToSql()

	_, err := r.db.ConnPool.Exec(ctx, sql, args...)
	if err != nil {
		r.log.Errorf("recordRepository.SaveRecord - Exec: %v", err)
		return err
	}

	return nil
}
func (r *recordRepository) GetAvgTempByDay(ctx context.Context, day time.Time) (float32, error) {
	sql, args, _ := r.db.Builder.
		Select("AVG(value)").
		From("records").
		Where("timestamp::date = ?::date", day).
		ToSql()

	var avg float32
	err := r.db.ConnPool.QueryRow(ctx, sql, args...).Scan(&avg)
	if err != nil {
		r.log.Errorf("recordRepository.GetAvgTempByDay - QueryRow: %v", err)
		return 0, err
	}

	return avg, nil
}
