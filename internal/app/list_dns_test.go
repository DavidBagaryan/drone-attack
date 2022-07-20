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

func TestImplementation_ListDNS(t *testing.T) {
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

	dns := dns_storage.New()
	impl := New(sector_storage.New(), dns)

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			rw := &trw{}
			impl.ListDNS(rw, tc.req)

			assert.Equal(t, tc.response, rw.response)
			assert.Equal(t, tc.statusCode, rw.statusCode)
		})
	}
}

func TestImplementation_ListDNS_NotEmpty(t *testing.T) {
	dns := dns_storage.New()
	dns.Set(&dto.DNSResp{SectorID: 12, Location: 21.001})
	dns.Set(&dto.DNSResp{SectorID: 12, Location: 211.21})
	dns.Set(&dto.DNSResp{SectorID: 51, Location: 777.2})
	impl := New(sector_storage.New(), dns)

	rw := &trw{}
	impl.ListDNS(rw, &http.Request{Method: "GET"})

	got := dto.ListDNSResp{}
	err := json.Unmarshal([]byte(rw.response), &got)
	assert.NoError(t, err)

	want := dto.ListDNSResp{{Location: 21.001}, {Location: 211.21}, {Location: 777.2}}
	sort.Slice(got, func(i, j int) bool { return got[i].Location < got[j].Location })
	sort.Slice(want, func(i, j int) bool { return want[i].Location < want[j].Location })

	assert.Equal(t, want, got)
}
