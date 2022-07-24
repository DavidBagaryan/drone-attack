package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DavidBagaryan/drone-attack/internal/dto"
	"github.com/gorilla/mux"
)

// LocateDNS locates dns by coordinates and given sectorID in url query
func (i Implementation) LocateDNS(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		response405(writer)
		return
	}

	strSectorID := request.URL.Query().Get("id")
	if strSectorID == "" {
		var ok bool
		query := mux.Vars(request)
		strSectorID, ok = query["id"]
		if !ok {
			response400custom(writer, "sectorID is undefined")
			return
		}
	}

	intSectorID, err := strconv.ParseInt(strSectorID, 10, 64)
	if err != nil {
		response400custom(writer, "sectorID is invalid")
		return
	}

	if intSectorID < 0 {
		response400custom(writer, "sectorID must be greater then or equal 0")
		return
	}

	decoder := json.NewDecoder(request.Body)
	defer request.Body.Close()

	data := new(dto.DNSReq)
	err = decoder.Decode(data)
	if err != nil {
		response400(writer)
		return
	}

	sector, err := i.sectors.Get(uint64(intSectorID))
	if err != nil {
		response400custom(writer, err)
		return
	}

	sectorID, err := sector.Book()
	if err != nil {
		response400custom(writer, err)
		return
	}

	resp := data.DNSRespWithSectorID(sectorID)
	i.dns.Set(resp)

	response200(writer, resp)
}
