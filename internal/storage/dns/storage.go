package dns_storage

import (
	"sync"

	"github.com/DavidBagaryan/drone-attack/internal/dto"
)

// DNS drone navigation service runtime storage
type DNS struct {
	sync.RWMutex
	data data
}

type data map[float64]*dto.DNSResp

// List composes dto.ListDNSResp
func (s *DNS) List() (list dto.ListDNSResp) {
	s.RLock()
	defer s.RUnlock()

	for _, d := range s.data {
		list = append(list, d)
	}

	return
}

// Set adds DNS to storage
func (s *DNS) Set(dns *dto.DNSResp) {
	s.Lock()
	defer s.Unlock()

	s.data[dns.Location] = dns
}

// New runtime DNS storage
func New() *DNS {
	return &DNS{
		data: make(data),
	}
}
