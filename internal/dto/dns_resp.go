package dto

import "encoding/json"

// ListDNSResp DNSResp list
type ListDNSResp []*DNSResp

// DNSResp drone navigation service response data
type DNSResp struct {
	SectorID uint64  `json:"-"`
	Location float64 `json:"loc"`
}

func (d DNSResp) String() string {
	resp, _ := json.Marshal(d)
	return string(resp)
}
