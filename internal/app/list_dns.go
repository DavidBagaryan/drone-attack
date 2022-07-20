package app

import (
	"net/http"
)

// ListDNS fetches all existed dns
func (i Implementation) ListDNS(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		response405(writer)
		return
	}

	response200(writer, i.dns.List())
}
