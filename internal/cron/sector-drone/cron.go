package sector_drone_cron

import (
	"math/rand"
	"time"

	"drone-attack/internal/dto"
)

type sectors interface {
	List() (list dto.ListSectorResp)
}

// this cron is about adding rand count of drones on sector
type cron struct {
	duration time.Duration
	sectors  sectors
}

// Run starts sector-drone cron
func (c cron) Run() {
	min, max := 1, 101
	ticker := time.NewTicker(c.duration)
	for {
		select {
		case <-ticker.C:
			for _, sector := range c.sectors.List() {
				rand.Seed(time.Now().UnixNano())
				randDronesNumber := rand.Intn(max-min) + min
				sector.DroneCount = uint64(randDronesNumber)
			}
		}
	}
}

// New constructor
func New(duration time.Duration, sectors sectors) *cron {
	return &cron{
		duration: duration,
		sectors:  sectors,
	}
}
