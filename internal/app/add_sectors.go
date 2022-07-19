package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"drone-attack/internal/dto"
)

// AddSectors adds sectors by giving params
func (i Implementation) AddSectors(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		fmt.Fprint(writer, "method not allowed")
		return
	}

	decoder := json.NewDecoder(request.Body)
	var data dto.ListSectorReq
	err := decoder.Decode(&data)
	if err != nil {
		log.Printf("[ADD SECTORS]: %s", err)
		fmt.Fprint(writer, "an error occurred")
		return
	}

	fmt.Fprint(writer, i.sectors.Add(data))
}
