package main

import (
	"GauGau/agen/clean-code/api/routers"
	"GauGau/agen/clean-code/configs"
	"GauGau/agen/clean-code/services"
	"log"
)

func main() {
	log.Println("begin....")

	cfg := configs.LoadConfig()
	log.Println("demo config = ", cfg.Demo)

	s := services.New()

	routers.NewGin(s)

	log.Println("end....")
}
