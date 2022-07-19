package app

import "net/http"

// trw test response writer to check output
// do not use in prod
type trw struct {
	response   string
	statusCode int
}

func (t *trw) Header() http.Header        { panic("implement me") }
func (t *trw) WriteHeader(statusCode int) { t.statusCode = statusCode }
func (t *trw) Write(bytes []byte) (int, error) {
	t.response = string(bytes)
	return 0, nil
}
