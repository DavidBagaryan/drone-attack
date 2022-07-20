package app

import (
	"encoding/json"
	"net/http"
	"sort"
	"testing"

	"github.com/DavidBagaryan/drone-attack/internal/dto"
	dns_storage "github.com/DavidBagaryan/drone-attack/internal/storage/dns"
	sector_storage "github.com/DavidBagaryan/drone-attack/internal/storage/sector"
	"github.com/stretchr/testify/assert"
)

func TestImplementation_ListSectors(t *testing.T) {
	tt := map[string]struct {
		statusCode int
		response   string
		req        *http.Request
	}{
		"method not allowed": {
			statusCode: 405,
			response:   "method not allowed",
			req:        &http.Request{Method: "POST"},
		},
		"all ok, empty list": {
			statusCode: 200,
			response:   "[]",
			req:        &http.Request{Method: "GET"},
		},
	}

	sectors := sector_storage.New()
	impl := New(sectors, dns_storage.New())

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			rw := &trw{}
			impl.ListSectors(rw, tc.req)

			assert.Equal(t, tc.response, rw.response)
			assert.Equal(t, tc.statusCode, rw.statusCode)
		})
	}
}

func TestImplementation_ListSectors_NotEmpty(t *testing.T) {
	sectors := sector_storage.New()
	impl := New(sectors, dns_storage.New())

	sectors.Add(dto.ListSectorReq{{AvailableDNS: 2}, {AvailableDNS: 3}, {AvailableDNS: 7}})
	rw := &trw{}
	impl.ListSectors(rw, &http.Request{Method: "GET"})

	got := dto.ListSectorResp{}
	err := json.Unmarshal([]byte(rw.response), &got)
	assert.NoError(t, err)

	want := dto.ListSectorResp{{ID: 0, AvailableDNS: 2}, {ID: 1, AvailableDNS: 3}, {ID: 2, AvailableDNS: 7}}
	sort.Slice(got, func(i, j int) bool { return got[i].AvailableDNS < got[j].AvailableDNS })
	sort.Slice(want, func(i, j int) bool { return want[i].AvailableDNS < want[j].AvailableDNS })

	assert.Equal(t, want, got)
}
