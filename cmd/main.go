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

	// this (subject/action) looks a little ugly, but I don't want to implement dispatcher
	// or add smth like gorilla in 1.0.0 version
	// but v1.1.0 will contain proper REST API approach
	// just keep prev endpoint to backwards capability (and for history)
	// DEPRECATED and will remove in ^2.0.0
	router.HandleFunc("/sectors/add", impl.AddSectors)
	router.HandleFunc("/sectors/list", impl.ListSectors)
	router.HandleFunc("/sector/locate", impl.LocateDNS) // sectorID passes as a query param
	router.HandleFunc("/dns/list", impl.ListDNS)

	// new API
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
