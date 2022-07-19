package dto

// DNSReq drone navigation service request data
type DNSReq struct {
	X        float64 `json:"x,string"`
	Y        float64 `json:"y,string"`
	Z        float64 `json:"z,string"`
	Velocity float64 `json:"vel,string"`
}

// Location super business logics
func (d DNSReq) Location(sectorID uint64) float64 {
	floatSectorID := float64(sectorID)
	return d.X*floatSectorID + d.Y*floatSectorID + d.Z*floatSectorID + d.Velocity
}

// DNSRespWithSectorID DNSResp creator
func (d DNSReq) DNSRespWithSectorID(sectorID uint64) *DNSResp {
	return &DNSResp{
		SectorID: sectorID,
		Location: d.Location(sectorID),
	}
}
