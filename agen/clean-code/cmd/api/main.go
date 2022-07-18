package main

import (
	"{{PROJECT_NAME}}/api/routers"
	"{{PROJECT_NAME}}/configs"
	"{{PROJECT_NAME}}/services"
	"{{PROJECT_NAME}}/pkg/mongodb"
	"{{PROJECT_NAME}}/repository"
	"log"
	"fmt"
)

func main() {
	log.Println("begin....")

	cfg := configs.LoadConfig()
	log.Println("demo config = ", cfg.Demo)


	mongodbHandler, err := mongodb.Init(mongodb.Config{})
	if err != nil {
		fmt.Printf("cannot connect to mongodb - error: [%s]", err)
	}

	s := services.
	New(){{SET_REPO_SERVICE}}.Build()

	routers.New(s)

	log.Println("end....")
}
