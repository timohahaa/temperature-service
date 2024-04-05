package service

import (
	"context"
	"time"

	"github.com/timohahaa/temperature-service/internal/entity"
)

type RecordService interface {
	SaveRecord(ctx context.Context, record entity.Record) error
	GetAvgTempByDay(ctx context.Context, day time.Time) (float32, error)
}

var _ RecordService = &recordService{}
