package sector_storage

import (
	"fmt"
	"sync"

	"github.com/DavidBagaryan/drone-attack/internal/dto"
)

type storage struct {
	sync.RWMutex
	data data
}

type data map[uint64]*dto.SectorResp

func (s *storage) Add(req dto.ListSectorReq) (added dto.ListSectorResp) {
	s.RLock()
	defer s.RUnlock()

	for _, sectorReq := range req {
		id := uint64(len(s.data))
		sector := sectorReq.SectorRespWithID(id)
		s.data[id] = sector
		added = append(added, sector)
	}

	return
}
func (s *storage) List() (list dto.ListSectorResp) {
	s.RLock()
	defer s.RUnlock()

	for _, sector := range s.data {
		list = append(list, sector)
	}

	return
}

func (s *storage) Get(id uint64) (*dto.SectorResp, error) {
	resp, ok := s.data[id]
	if !ok {
		return nil, fmt.Errorf("sector id %d not found", id)
	}

	return resp, nil
}

//New sector storage constructor
func New() *storage {
	s := &storage{
		data: make(data),
	}
	return s
}
