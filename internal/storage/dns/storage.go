package dns_storage

import (
	"sync"

	"github.com/DavidBagaryan/drone-attack/internal/dto"
)

type dns struct {
	sync.RWMutex
	data data
}

type data map[float64]*dto.DNSResp

// Set adds dns to storage
func (s *dns) Set(dns *dto.DNSResp) {
	s.RLock()
	defer s.RUnlock()

	s.data[dns.Location] = dns
}

// List composes dto.ListDNSResp
func (s *dns) List() (list dto.ListDNSResp) {
	s.RLock()
	defer s.RUnlock()

	for _, d := range s.data {
		list = append(list, d)
	}

	return
}

// New runtime dns storage
func New() *dns {
	return &dns{
		data: make(data),
	}
}
