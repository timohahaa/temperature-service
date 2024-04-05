package v1

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/timohahaa/temperature-service/internal/service"
	pb "github.com/timohahaa/temperature-service/proto/record"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type recordServer struct {
	pb.UnimplementedRecordServer
	recordService service.RecordService
	log           *logrus.Logger
}

var _ pb.RecordServer = &recordServer{}

func NewRecordServer(recordService service.RecordService, logger *logrus.Logger) *recordServer {
	return &recordServer{
		recordService: recordService,
		log:           logger,
	}
}

func (s *recordServer) GetAvgTemp(ctx context.Context, req *pb.GetAvgTempReq) (*pb.GetAvgTempResp, error) {
	day, err := time.Parse(time.RFC3339, req.GetDate())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "specify a date in a RFC3339 format")
	}

	avg, err := s.recordService.GetAvgTempByDay(ctx, day)
	if err != nil {
		s.log.Errorf("recordServer.GetAvgTemp - recordService.GetAvgTempByDay: %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.GetAvgTempResp{Temperature: avg}, nil
}
