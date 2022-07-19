package sector_drone_cron

import (
	"context"
	"testing"
	"time"

	"github.com/DavidBagaryan/drone-attack/internal/dto"
	sector_storage "github.com/DavidBagaryan/drone-attack/internal/storage/sector"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestService_Run(t *testing.T) {
	sectorIMDB := sector_storage.New()
	req := genTestListSectorReq(15)
	sectorIMDB.Add(req)

	svc := New(time.Second, sectorIMDB)

	ctxTimeout := time.Second
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()
	svc.Run(ctx)

	time.Sleep(ctxTimeout) // wait for it
	for _, sector := range sectorIMDB.List() {
		assert.Equal(t, 1, sector.Rotation)
	}

	ctxTimeout = 2 * time.Second
	ctx, cancel = context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()
	svc.Run(ctx)

	time.Sleep(ctxTimeout) // again wait for it
	for _, sector := range sectorIMDB.List() {
		assert.Equal(t, 3, sector.Rotation)
	}
}

func genTestListSectorReq(sectorCount int) dto.ListSectorReq {
	list := make(dto.ListSectorReq, sectorCount)
	for i := 0; i < sectorCount; i++ {
		s := new(dto.SectorReq)
		_ = gofakeit.Struct(s)
		list[i] = s
	}

	return list
}
