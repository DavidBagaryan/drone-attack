package sector_storage

import (
	"fmt"
	"sync"

	"github.com/DavidBagaryan/drone-attack/internal/dto"
)

// IMDB runtime storage
type IMDB struct {
	sync.RWMutex
	data data
}

type data map[uint64]*dto.SectorResp

// Add adds butch of sectors
func (s *IMDB) Add(req dto.ListSectorReq) (added dto.ListSectorResp) {
	s.Lock()
	defer s.Unlock()

	for _, sectorReq := range req {
		nextID := uint64(len(s.data))
		sector := sectorReq.SectorRespWithID(nextID)
		s.data[nextID] = sector
		added = append(added, sector)
	}

	return
}

// Get fetches sector by id
func (s *IMDB) Get(id uint64) (*dto.SectorResp, error) {
	s.RLock()
	resp, ok := s.data[id]
	s.RUnlock()

	if !ok {
		return nil, fmt.Errorf("sector id %d not found", id)
	}

	return resp, nil
}

// List lists existed sectors
func (s *IMDB) List() (list dto.ListSectorResp) {
	s.RLock()
	defer s.RUnlock()

	for _, sector := range s.data {
		list = append(list, sector)
	}

	return
}

//New sector IMDB constructor
func New() *IMDB {
	s := &IMDB{
		data: make(data),
	}
	return s
}
