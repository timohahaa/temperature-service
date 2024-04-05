package app

import (
	"context"
	"net"

	"github.com/sirupsen/logrus"
	"github.com/timohahaa/postgres"
	v1 "github.com/timohahaa/temperature-service/internal/controllers/grpc/v1"
	"github.com/timohahaa/temperature-service/internal/delivery/nats"
	"github.com/timohahaa/temperature-service/internal/repository"
	"github.com/timohahaa/temperature-service/internal/service"
	"google.golang.org/grpc"

	pb "github.com/timohahaa/temperature-service/proto/record"
)

func Run() {
	logger := logrus.New()

	pg, err := postgres.New("url here")
	if err != nil {
		logger.Fatal(err)
	}
	recordRepo := repository.NewPostgresRepository(pg, logger)
	recordService := service.NewRecordService(recordRepo, logger)

	recordServer := v1.NewRecordServer(recordService, logger)
	recordConsumer := nats.NewRecordConsumer(recordService, "0.0.0.0:4222", "record", 1, 5, logger)

	err = recordConsumer.StartRecordConsumer(context.Background())
	if err != nil {
		logger.Fatal(err)
	}

	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		logger.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterRecordServer(grpcServer, recordServer)
	grpcServer.Serve(lis)
}
