package app

import (
	"fmt"
	"log"
	"net/http"
)

func response200(writer http.ResponseWriter, resp interface{}) {
	response(http.StatusOK, resp, writer)
}

func response400(writer http.ResponseWriter) {
	response(http.StatusBadRequest, "an error occurred", writer)
}

func response400custom(writer http.ResponseWriter, resp interface{}) {
	response(http.StatusBadRequest, resp, writer)
}

func response405(writer http.ResponseWriter) {
	response(http.StatusMethodNotAllowed, "method not allowed", writer)
}

func response(code int, resp interface{}, writer http.ResponseWriter) {
	writer.WriteHeader(code)
	_, err := fmt.Fprint(writer, resp)
	if err != nil {
		log.Printf("[RESPONSE]: %s", err)
	}
}
