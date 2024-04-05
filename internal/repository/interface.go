package repository

import (
	"context"
	"time"

	"github.com/timohahaa/temperature-service/internal/entity"
)

type RecordRepository interface {
	SaveRecord(ctx context.Context, record entity.Record) error
	GetAvgTempByDay(ctx context.Context, day time.Time) (float32, error)
}

var _ RecordRepository = &recordRepository{}
