package service

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/timohahaa/temperature-service/internal/entity"
	"github.com/timohahaa/temperature-service/internal/repository"
)

type recordService struct {
	recordRepo repository.RecordRepository
	log        *logrus.Logger
}

func NewRecordService(repo repository.RecordRepository, logger *logrus.Logger) *recordService {
	return &recordService{
		recordRepo: repo,
		log:        logger,
	}
}

func (s *recordService) SaveRecord(ctx context.Context, record entity.Record) error {
	return s.recordRepo.SaveRecord(ctx, record)
}

func (s *recordService) GetAvgTempByDay(ctx context.Context, day time.Time) (float32, error) {
	return s.recordRepo.GetAvgTempByDay(ctx, day)
}
