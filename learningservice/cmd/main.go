package main

import (
	"log"
	"net"

	pb "learn_service/genproto/learning"
	"learn_service/pkg/logger"
	service "learn_service/service"
	postgres "learn_service/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.ConnectionDb()
	if err != nil {
		log.Fatalf("Error while connecting to the database: %v", err)
	}

	liss, err := net.Listen("tcp", ":8070")
	if err != nil {
		log.Fatalf("Error while listening on TCP: %v", err)
	}
	logger := logger.NewLogger()
	s := grpc.NewServer()
	lMan := postgres.NewLearnRepository(db)
	pb.RegisterLearningServiceServer(s, service.NewLearningService(lMan, logger).(pb.LearningServiceServer))

	log.Printf("Server listening at %v", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
