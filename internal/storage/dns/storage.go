package dns_storage

import (
	"sync"

	"github.com/DavidBagaryan/drone-attack/internal/dto"
)

// IMDB drone navigation service runtime storage
type IMDB struct {
	sync.RWMutex
	data data
}

type data map[float64]*dto.DNSResp

// List composes dto.ListDNSResp
func (s *IMDB) List() (list dto.ListDNSResp) {
	s.RLock()
	defer s.RUnlock()

	for _, d := range s.data {
		list = append(list, d)
	}

	return
}

// Set adds IMDB to storage
func (s *IMDB) Set(dns *dto.DNSResp) {
	s.Lock()
	defer s.Unlock()

	s.data[dns.Location] = dns
}

// New runtime IMDB storage
func New() *IMDB {
	return &IMDB{
		data: make(data),
	}
}
