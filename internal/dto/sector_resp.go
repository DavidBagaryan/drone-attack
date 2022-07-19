package dto

import (
	"encoding/json"
	"errors"
	"sync"
)

// ListSectorResp sector list response
type ListSectorResp []*SectorResp

func (s ListSectorResp) String() string {
	resp, _ := json.Marshal(s)
	if string(resp) == "null" {
		return "[]"
	}

	return string(resp)
}

// SectorResp sector in storage
type SectorResp struct {
	ID           uint64 `json:"id"`
	AvailableDNS uint64 `json:"available_dns"`
	DeployedDNS  uint64 `json:"deployed_dns"`
	DroneCount   uint64 `json:"drone_count"`
	Rotation     int    `json:"-"`
	sync.RWMutex
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

	s.Lock()
	defer s.Unlock()

	if s.AvailableDNS == 0 {
		return 0, errors.New("no vacancies found")
	}

	s.AvailableDNS--
	s.DeployedDNS++

	return s.ID, nil
}
