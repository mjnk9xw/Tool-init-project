package main

import (
	"{{PROJECT_NAME}}/api/routers"
	"{{PROJECT_NAME}}/configs"
	"{{PROJECT_NAME}}/services"
	"log"
)

func main() {
	log.Println("begin....")

	cfg := configs.LoadConfig()
	log.Println("demo config = ", cfg.Demo)

	s := services.New()

	routers.New(s)

	log.Println("end....")
}
