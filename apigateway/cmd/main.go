package main

import (
	"api_service/api"
	_ "api_service/api/docs"
	"api_service/casbin"
	"api_service/config"
	"api_service/pkg/logger"
	"fmt"
	"log"
)

func main() {
	cfg := config.Load()
	logger, err := logger.NewLoger()
	if err != nil {
		fmt.Println(err)
	}

	enforcer, err := casbin.CasbinEnforcer(logger)
	fmt.Println(err)
	if err != nil {
		logger.Error(err.Error())
		log.Fatal(err)
	}
	fmt.Println(enforcer, "dfadsfasdifjadsiuf")
	res := api.NewRouter(cfg, logger, enforcer)

	err = res.Run("api-service:8080")
	if err != nil {
		logger.Error(err.Error())
	}

}
