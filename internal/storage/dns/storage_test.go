package dns_storage

import (
	"testing"

	"github.com/DavidBagaryan/drone-attack/internal/dto"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestIMDB_List(t *testing.T) {
	dns := New()
	var err error

	dns1 := new(dto.DNSResp)
	err = gofakeit.Struct(dns1)
	assert.NoError(t, err)

	dns2 := new(dto.DNSResp)
	err = gofakeit.Struct(dns2)
	assert.NoError(t, err)

	dns3 := new(dto.DNSResp)
	err = gofakeit.Struct(dns3)
	assert.NoError(t, err)

	dnsData := data{
		dns1.Location: dns1,
		dns2.Location: dns2,
		dns3.Location: dns3,
	}

	dns.data = dnsData
	listResp := dns.List()
	compareDNSStorageDataAndRespList(t, dnsData, listResp)
}

func TestIMDB_Set(t *testing.T) {
	dns := New()
	general := new(dto.DNSResp)
	_ = gofakeit.Struct(general)
	dns.Set(general)

	listResp := dns.List()
	compareDNSStorageDataAndRespList(t, data{general.Location: general}, listResp)

	upd := new(dto.DNSResp)
	_ = gofakeit.Struct(upd)
	upd.Location = general.Location
	dns.Set(upd)
	listResp = dns.List()
	compareDNSStorageDataAndRespList(t, data{upd.Location: upd}, listResp)
}

func compareDNSStorageDataAndRespList(t *testing.T, data data, resp dto.ListDNSResp) {
	assert.Equal(t, len(data), len(resp))

	for _, dnsResp := range resp {
		got, ok := data[dnsResp.Location]
		assert.Equal(t, dnsResp, got)
		assert.True(t, ok)
	}
}
