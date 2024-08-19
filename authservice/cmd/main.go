package main

import (
	"auth_service/api"
	"auth_service/api/handler"
	"auth_service/config"
	"auth_service/genproto/user"
	"auth_service/logger"
	"auth_service/service"
	"auth_service/storage/postgres"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

// func main() {
// 	// Ma'lumotlar bazasiga ulanish
// 	db, err := postgres.ConnectDB()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	// Logger yaratish
// 	log := logger.NewLogger()

// 	// Servislarni yaratish
// 	userRepo := postgres.NewAuthRepo(db)
// 	authService := service.NewUserService(userRepo, log)
// 	authHandler := handler.NewAuthenticaionHandler(authService, log)

// 	// gRPC serverini yaratish
// 	s := grpc.NewServer()
// 	user.RegisterUsersServer(s, authService)

// 	// var wg sync.WaitGroup

// 	// HTTP serverni alohida goroutine'da ishga tushurish
// 	// wg.Add(1)
// 	go func() {
// 		// defer wg.Done()
// 		fmt.Printf("Server is listening on port %s\n", config.Load().USER_SERVICE)
// 		auth := api.NewServer(authHandler)
// 		router := auth.NewRouter()
// 		if err := router.Run(config.Load().USER_SERVICE); err != nil {
// 			log.Error("server error", "Error while running HTTP server: %v", err)
// 		}
// 	}()

// 	// gRPC serverni TCP portda tinglash
// 	go func() {

// 		liss, err := net.Listen("tcp", "localhost"+config.Load().HTTP_PORT)
// 		if err != nil {
// 			log.Error("...", "Error while listening on TCP: %v", err)
// 			return
// 		}

// 		// gRPC serverni ishga tushurish
// 		if err := s.Serve(liss); err != nil {
// 			log.Error("service not listening", "Failed to serve: %v", err)
// 		}
// 	}()

// 	quit := make(chan os.Signal, 1)
// 	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
// 	<-quit
// 	// Goroutine'larni kutish
// 	// wg.Wait()
// }

func main() {
	// Ma'lumotlar bazasiga ulanish
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Logger yaratish
	log := logger.NewLogger()

	// Servislarni yaratish
	userRepo := postgres.NewAuthRepo(db)
	authService := service.NewUserService(userRepo, log)
	authHandler := handler.NewAuthenticaionHandler(authService, log)

	// gRPC serverini yaratish
	s := grpc.NewServer()
	user.RegisterUsersServer(s, authService)

	// HTTP serverni alohida goroutine'da ishga tushurish
	go func() {
		fmt.Printf("HTTP Server is listening on port %s\n", config.Load().USER_SERVICE)
		// fmt.Printf("HTTP Server is listening on port %s\n", config.Load().HTTP_PORT)
		auth := api.NewServer(authHandler)
		router := auth.NewRouter()
		if err := router.Run(config.Load().USER_SERVICE); err != nil {
			log.Error("server error", "Error while running HTTP server: %v", err)
		}
	}()

	// gRPC serverni TCP portda tinglash
	go func() {
		liss, err := net.Listen("tcp", config.Load().HTTP_PORT)
		if err != nil {
			log.Error("grpc error", "Error while listening on TCP: %v", err)
			return
		}

		// gRPC serverni ishga tushurish
		fmt.Println("Server is started on", liss.Addr())
		if err := s.Serve(liss); err != nil {
			log.Error("grpc error", "Failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
