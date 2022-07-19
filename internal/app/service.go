package app

import "github.com/DavidBagaryan/drone-attack/internal/dto"

// DNS data storage for drone navigation service
type DNS interface {
	Set(dns *dto.DNSResp)
	List() dto.ListDNSResp
}

// Sectors storage interface
type Sectors interface {
	Add(req dto.ListSectorReq) dto.ListSectorResp
	List() dto.ListSectorResp
	Get(id uint64) (*dto.SectorResp, error)
}

// Implementation drone navigating service implementation
type Implementation struct {
	dns     DNS
	sectors Sectors
}

// New app implementation constructor
func New(sectors Sectors, dns DNS) *Implementation {
	return &Implementation{
		dns:     dns,
		sectors: sectors,
	}
}
