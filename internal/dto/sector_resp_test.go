package dto

import (
	"errors"
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSectorResp_Book(t *testing.T) {
	sector := &SectorResp{ID: 777, AvailableDNS: 1001}
	attempts := 1000

	var wg sync.WaitGroup
	wg.Add(attempts)

	for i := 0; i < attempts; i++ {
		i := i
		go func() {
			defer wg.Done()
			t.Log(fmt.Sprintf("writer-%d process...", i))
			sID, err := sector.Book()
			assert.Equal(t, sector.ID, sID)
			assert.NoError(t, err)
		}()
	}

	wg.Wait()
	assert.Equal(t, &SectorResp{ID: 777, AvailableDNS: 1, DeployedDNS: 1000, DroneCount: 0}, sector)

	// second try

	sector = &SectorResp{ID: 666, AvailableDNS: 1}
	sID, err := sector.Book()
	assert.Equal(t, sector.ID, sID)
	assert.NoError(t, err)
	assert.Equal(t, &SectorResp{ID: 666, AvailableDNS: 0, DeployedDNS: 1, DroneCount: 0}, sector)

	sID, err = sector.Book()
	assert.Equal(t, uint64(0), sID)
	assert.Equal(t, errors.New("no vacancies found"), err)
	assert.Equal(t, &SectorResp{ID: 666, AvailableDNS: 0, DeployedDNS: 1, DroneCount: 0}, sector)

	sector = nil
	sID, err = sector.Book()
	assert.Equal(t, uint64(0), sID)
	assert.Equal(t, errors.New("sector is undefined"), err)
}
