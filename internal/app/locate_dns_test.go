package app

import (
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/DavidBagaryan/drone-attack/internal/dto"
	dns_storage "github.com/DavidBagaryan/drone-attack/internal/storage/dns"
	sector_storage "github.com/DavidBagaryan/drone-attack/internal/storage/sector"
	"github.com/stretchr/testify/assert"
)

func TestImplementation_LocateDNS(t *testing.T) {
	badReq1, err := http.NewRequest("POST", "test.me/sector/locate?id=12", strings.NewReader(""))
	assert.NoError(t, err)

	badReq2, err := http.NewRequest(
		"POST",
		"test.me/sector/locate?id=666",
		strings.NewReader(`{"x": "123.12","y": "456.56","z": "789.89","vel": "20.0"}`),
	)
	assert.NoError(t, err)

	badReq3, err := http.NewRequest(
		"POST",
		"test.me/sector/locate?id=2",
		strings.NewReader(`{"x": "123.12","y": "456.56","z": "789.89","vel": "20.0"}`),
	)
	assert.NoError(t, err)

	badReq4, err := http.NewRequest(
		"POST",
		"test.me/sector/locate?id=2",
		strings.NewReader(`{"x": "123.12","y": "456.56","z": "789.89","vel": "20.0"}`),
	)
	assert.NoError(t, err)

	okReq, err := http.NewRequest(
		"POST",
		"test.me/sector/locate?id=1",
		strings.NewReader(`{"x": "123.12","y": "456.56","z": "789.89","vel": "20.0"}`),
	)
	assert.NoError(t, err)

	tt := map[string]struct {
		statusCode int
		response   string
		req        *http.Request
	}{
		"method not allowed": {
			statusCode: 405,
			response:   "method not allowed",
			req:        &http.Request{Method: "GET"},
		},
		"sectorID is undefined": {
			statusCode: 400,
			response:   "sectorID is undefined",
			req:        &http.Request{Method: "POST", URL: &url.URL{}},
		},
		"sectorID is invalid": {
			statusCode: 400,
			response:   "sectorID is invalid",
			req:        &http.Request{Method: "POST", URL: &url.URL{RawQuery: "id=test"}},
		},
		"sectorID must be greater then or equal 0": {
			statusCode: 400,
			response:   "sectorID must be greater then or equal 0",
			req:        &http.Request{Method: "POST", URL: &url.URL{RawQuery: "id=-12"}},
		},
		"bad request": {
			statusCode: 400,
			response:   "an error occurred",
			req:        badReq1,
		},
		"sector not found": {
			statusCode: 400,
			response:   "sector id 666 not found",
			req:        badReq2,
		},
		"no vacancies found": {
			statusCode: 400,
			response:   "no vacancies found",
			req:        badReq4,
		},
		"all ok": {
			response: `{"loc":1389.5700000000002}`,
			req:      okReq,
		},
	}

	sectors := sector_storage.New()
	sectors.Add(dto.ListSectorReq{{AvailableDNS: 41}, {AvailableDNS: 2}, {AvailableDNS: 0}})
	dns := dns_storage.New()
	impl := New(sectors, dns)

	impl.LocateDNS(&trw{}, badReq3) // to make AvailableDNS counter to 0

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			rw := &trw{}
			impl.LocateDNS(rw, tc.req)

			assert.Equal(t, tc.response, rw.response)
			assert.Equal(t, tc.statusCode, rw.statusCode)
		})
	}
}
