package dto

import (
	"encoding/json"
	"errors"
	"sync/atomic"
)

// ListSectorResp sector list response
type ListSectorResp []*SectorResp

// SectorResp sector in storage
type SectorResp struct {
	ID           uint64
	AvailableDNS uint64
	DeployedDNS  uint64
	DroneCount   uint64
}

func (s *SectorResp) String() string {
	res, _ := json.Marshal(s)
	return string(res)
}

// Book fetches sector ID if sector has vacancy to place a new dns
func (s *SectorResp) Book() (uint64, error) {
	if s == nil {
		return 0, errors.New("sector is undefined")
	}
	if s.AvailableDNS == 0 {
		return 0, errors.New("no vacancies found")
	}

	atomic.AddUint64(&s.AvailableDNS, ^uint64(0))
	atomic.AddUint64(&s.DeployedDNS, 1)

	return s.ID, nil
}
