package main

import (
	"log"
	"net/http"

	"drone-attack/internal/app"
	"drone-attack/internal/config"
	sector_drone_cron "drone-attack/internal/cron/sector-drone"
	dns_storage "drone-attack/internal/storage/dns"
	sector_storage "drone-attack/internal/storage/sector"
)

func main() {
	sectors := sector_storage.New()
	impl := app.New(sectors, dns_storage.New())

	sectorDroneCron := sector_drone_cron.New(config.SectorDroneCronDuration, sectors)
	go sectorDroneCron.Run()

	// this (subject/action) looks a little ugly, but I don't want to implement dispatcher
	// or add smth like gorilla
	http.HandleFunc("/sectors/add", impl.AddSectors)
	http.HandleFunc("/sectors/list", impl.ListSectors)
	http.HandleFunc("/sector/locate", impl.LocateDNS) // sectorID passes as a query param
	http.HandleFunc("/dns/list", impl.ListDNS) // sectorID passes as a query param

	log.Fatal(http.ListenAndServe(config.LocalDeployPort, nil))
}
