package app

import (
	"fmt"
	"net/http"
)

// ListDNS fetches all existed dns
func (i Implementation) ListDNS(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		fmt.Fprint(writer, "method not allowed")
		return
	}

	fmt.Fprint(writer, i.dns.List())
}
