package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/vrischmann/envconfig"

	"github.com/weather-project/config"
	"github.com/weather-project/handler"

	_ "github.com/lib/pq"
)

func main() {
	log.Println("Starting service...")

	var cfg config.Service
	if err := envconfig.Init(&cfg); err != nil {
		log.Panicf("Error loading main configuration %v\n", err.Error())
	}
	log.Print(cfg)

	if err := startService(cfg.Port); err != nil {
		log.Fatal("Unable to start server", err)
	}
}

func startService(port string) error {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/weather", handler.GetWeather).Methods(http.MethodGet)

	log.Printf("Starting server on port %s ", port)

	c := cors.AllowAll()
	return http.ListenAndServe(":"+port, c.Handler(router))
}
