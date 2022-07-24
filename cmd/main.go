package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/DavidBagaryan/drone-attack/internal/app"
	"github.com/DavidBagaryan/drone-attack/internal/config"
	sector_drone_cron "github.com/DavidBagaryan/drone-attack/internal/cron/sector-drone"
	dns_storage "github.com/DavidBagaryan/drone-attack/internal/storage/dns"
	sector_storage "github.com/DavidBagaryan/drone-attack/internal/storage/sector"
	"github.com/gorilla/mux"
)

func main() {
	sectors := sector_storage.New()
	impl := app.New(sectors, dns_storage.New())

	sectorDroneCron := sector_drone_cron.New(config.SectorDroneCronDuration, sectors)
	go sectorDroneCron.Run(context.Background())

	router := mux.NewRouter()
	router.HandleFunc("/sectors", impl.AddSectors).Methods(http.MethodPost)
	router.HandleFunc("/sectors", impl.ListSectors).Methods(http.MethodGet)
	router.HandleFunc("/sector/{id:[0-9]+}/locate", impl.LocateDNS).Methods(http.MethodPost)
	router.HandleFunc("/dns", impl.ListDNS).Methods(http.MethodGet)

	apiPort := os.Getenv("API_PORT")
	if apiPort == "" {
		apiPort = config.APIPortDefault
	}

	log.Fatal(http.ListenAndServe(":"+apiPort, router))
}
