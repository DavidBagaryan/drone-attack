package sector_drone_cron

import (
	"context"
	"math/rand"
	"time"

	"github.com/DavidBagaryan/drone-attack/internal/config"
	"github.com/DavidBagaryan/drone-attack/internal/dto"
)

type sectors interface {
	List() dto.ListSectorResp
}

// Service is about adding rand count of drones on sector
type Service struct {
	duration time.Duration
	sectors  sectors
}

// Run starts sector-drone Service
func (c Service) Run(ctx context.Context) {
	min, max := config.MinDroneCount, config.MaxDroneCount
	ticker := time.NewTicker(c.duration)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			for _, sector := range c.sectors.List() {
				rand.Seed(time.Now().UnixNano())
				randDronesNumber := rand.Intn(max-min) + min
				sector.DroneCount = uint64(randDronesNumber)
				sector.Rotation++
			}
		}
	}
}

// New constructor
func New(duration time.Duration, sectors sectors) *Service {
	return &Service{
		duration: duration,
		sectors:  sectors,
	}
}
