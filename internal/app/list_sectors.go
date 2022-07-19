package app

import (
	"fmt"
	"net/http"
)

// ListSectors fetches all existed sectors
func (i Implementation) ListSectors(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		fmt.Fprint(writer, "method not allowed")
		return
	}

	fmt.Fprint(writer, i.sectors.List())
}
