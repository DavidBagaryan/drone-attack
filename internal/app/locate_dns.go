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
		fmt.Fprint(writer, "method not allowed")
		return
	}

	strSectorID := request.URL.Query().Get("id")
	if strSectorID == "" {
		fmt.Fprint(writer, "sectorID is undefined")
		return
	}

	const logFormat = "[LOCATE DNS]: %s"

	intSectorID, err := strconv.ParseInt(strSectorID, 10, 64)
	if err != nil {
		log.Printf(logFormat, err)
		fmt.Fprint(writer, "an error occurred")
		return
	}

	if intSectorID < 0 {
		fmt.Fprint(writer, "sectorID must be greater then or equal 0")
		return
	}

	decoder := json.NewDecoder(request.Body)
	data := new(dto.DNSReq)
	err = decoder.Decode(data)
	if err != nil {
		log.Printf(logFormat, err)
		fmt.Fprint(writer, "an error occurred")
		return
	}

	sector, err := i.sectors.Get(uint64(intSectorID))
	if err != nil {
		fmt.Fprint(writer, err)
		return
	}

	sectorID, err := sector.Book()
	if err != nil {
		fmt.Fprint(writer, err)
		return
	}

	resp := data.DNSRespWithSectorID(sectorID)
	i.dns.Set(resp)
	fmt.Fprint(writer, resp)
}
