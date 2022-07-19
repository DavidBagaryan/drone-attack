package sector_storage

import (
	"errors"
	"sort"
	"sync"
	"testing"

	"github.com/DavidBagaryan/drone-attack/internal/dto"
	"github.com/stretchr/testify/assert"
)

func TestIMDB_Add(t *testing.T) {
	sectors := New()
	clientReq := []dto.ListSectorReq{
		{{AvailableDNS: 12}, {AvailableDNS: 11}, {AvailableDNS: 9}},
		{{AvailableDNS: 4}, {AvailableDNS: 8}, {AvailableDNS: 5}, {AvailableDNS: 6}},
		{{AvailableDNS: 1}, {AvailableDNS: 3}},
	}

	var wg sync.WaitGroup
	wg.Add(len(clientReq))

	for _, req := range clientReq {
		req := req
		go func() {
			defer wg.Done()
			resp := sectors.Add(req)
			compareListSectorReqResp(t, req, resp)
		}()
	}

	wg.Wait()

	var allRequests dto.ListSectorReq
	for _, req := range clientReq {
		allRequests = append(allRequests, req...)
	}

	compareListSectorReqResp(t, allRequests, sectors.List())
}

func TestIMDB_List(t *testing.T) {
	sectors := New()
	listReq := dto.ListSectorReq{{AvailableDNS: 12}, {AvailableDNS: 11}, {AvailableDNS: 9}}
	resp := sectors.Add(listReq)
	compareListSectorReqResp(t, listReq, resp)

	t.Parallel()
	resp = sectors.List()
	compareListSectorReqResp(t, listReq, resp)
}

func TestIMDB_Get(t *testing.T) {
	sectors := New()
	req := dto.ListSectorReq{{AvailableDNS: 12}, {AvailableDNS: 11}, {AvailableDNS: 9}}
	resp := sectors.Add(req)
	compareListSectorReqResp(t, req, resp)

	t.Run("client-read-1", func(t *testing.T) {
		t.Parallel()
		want := resp[0]
		got, err := sectors.Get(want.ID)
		assert.NoError(t, err)
		assert.Equal(t, &dto.SectorResp{
			ID:           want.ID,
			AvailableDNS: want.AvailableDNS,
		}, got)
	})
	t.Run("client-read-wrong", func(t *testing.T) {
		t.Parallel()
		got, err := sectors.Get(666)
		assert.Nil(t, got)
		assert.Equal(t, errors.New("sector id 666 not found"), err)
	})
	t.Run("client-read-2", func(t *testing.T) {
		t.Parallel()
		want := resp[1]
		got, err := sectors.Get(want.ID)
		assert.NoError(t, err)
		assert.Equal(t, &dto.SectorResp{
			ID:           want.ID,
			AvailableDNS: want.AvailableDNS,
		}, got)
	})
	t.Run("client-read-3", func(t *testing.T) {
		t.Parallel()
		want := resp[2]
		got, err := sectors.Get(want.ID)
		assert.NoError(t, err)
		assert.Equal(t, &dto.SectorResp{
			ID:           want.ID,
			AvailableDNS: want.AvailableDNS,
		}, got)
	})
}

func compareListSectorReqResp(t *testing.T, req dto.ListSectorReq, resp dto.ListSectorResp) {
	assert.Equal(t, len(req), len(resp))

	sort.Slice(req, func(i, j int) bool { return req[i].AvailableDNS < req[j].AvailableDNS })
	sort.Slice(resp, func(i, j int) bool { return resp[i].AvailableDNS < resp[j].AvailableDNS })

	for i := 0; i < len(req); i++ {
		assert.Equal(t, req[i].AvailableDNS, resp[i].AvailableDNS)
	}
}
