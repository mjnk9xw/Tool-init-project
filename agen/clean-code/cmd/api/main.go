package main

import (
	"{{PROJECT_NAME}}/api/routers"
	"{{PROJECT_NAME}}/configs"
	"{{PROJECT_NAME}}/services"
	{{MONGO_PKG}}
	{{REDIS_PKG}}
	{{LOG_PKG}}
	"{{PROJECT_NAME}}/repository"
	"log"
	"fmt"
)

func main() {
	log.Println("begin....")

	cfg := configs.LoadConfig()
	log.Println("demo config = ", cfg.Demo)

	logger := logger.New()
	defer logger.Sync()
	
	{{MONGO_HANDLER}}
	{{REDIS_HANDLER}}

	s := services.
	New().SetLog(logger){{SET_REPO_SERVICE}}{{SET_REDIS_SERVICE}}.Build()

	routers.New(s)

	log.Println("end....")
}
