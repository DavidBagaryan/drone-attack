package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DavidBagaryan/drone-attack/internal/dto"
)

// AddSectors adds sectors by giving params
func (i Implementation) AddSectors(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		response405(writer)
		return
	}

	decoder := json.NewDecoder(request.Body)
	defer request.Body.Close()

	var data dto.ListSectorReq
	err := decoder.Decode(&data)
	if err != nil {
		log.Printf("[ADD SECTORS]: %s", err)
		response400(writer)
		return
	}

	response200(writer, i.sectors.Add(data))
}
