package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/DavidBagaryan/drone-attack/internal/dto"
)

// LocateDNS locates dns by coordinates and given sectorID in url query
func (i Implementation) LocateDNS(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(writer, "method not allowed")
		return
	}

	strSectorID := request.URL.Query().Get("id")
	if strSectorID == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "sectorID is undefined")
		return
	}

	const logFormat = "[LOCATE DNS]: %s"

	intSectorID, err := strconv.ParseInt(strSectorID, 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		log.Printf(logFormat, err)
		fmt.Fprint(writer, "sectorID is invalid")
		return
	}

	if intSectorID < 0 {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "sectorID must be greater then or equal 0")
		return
	}

	decoder := json.NewDecoder(request.Body)
	defer request.Body.Close()

	data := new(dto.DNSReq)
	err = decoder.Decode(data)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		log.Printf(logFormat, err)
		fmt.Fprint(writer, "an error occurred")
		return
	}

	sector, err := i.sectors.Get(uint64(intSectorID))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, err)
		return
	}

	sectorID, err := sector.Book()
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, err)
		return
	}

	resp := data.DNSRespWithSectorID(sectorID)
	i.dns.Set(resp)
	fmt.Fprint(writer, resp)
}
