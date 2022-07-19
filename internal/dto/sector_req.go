package dto

// ListSectorReq sector list request
type ListSectorReq []*SectorReq

// SectorReq sector request with dns and drones count
type SectorReq struct {
	AvailableDNS uint64 `json:"count_dns"`
}

// SectorRespWithID creator
func (r SectorReq) SectorRespWithID(id uint64) *SectorResp {
	ss := &SectorResp{
		ID:           id,
		AvailableDNS: r.AvailableDNS,
	}

	if ss.AvailableDNS == 0 {
		ss.AvailableDNS = 1
	}

	return ss
}
