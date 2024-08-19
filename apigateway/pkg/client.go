package pkg

import (
	"api_service/config"
	pbuLearning "api_service/genproto/learning"
	pbuAuth "api_service/genproto/user"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAuthenticationClient(cfg *config.Config) pbuAuth.UsersClient {
	fmt.Println("--------------------------------------------------", cfg.AUTH_SERVICE)
	conn, err := grpc.NewClient(fmt.Sprintf("auth-service:%s", cfg.AUTH_SERVICE), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("error while connecting authentication service ", err)
	}

	return pbuAuth.NewUsersClient(conn)

}

func NewLearningClient(cfg *config.Config) pbuLearning.LearningServiceClient {
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++", cfg.LEARNING_SERVICE)
	// Create the gRPC connection
	conn, err := grpc.NewClient(fmt.Sprintf("learning_service:%s", cfg.LEARNING_SERVICE), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to product service: %v", err)
	}
	
	// Return the new ProductServiceClient
	return pbuLearning.NewLearningServiceClient(conn)
}
