package config

import "time"

// it is better to have etcd with watchers on vars,
// but I think it is enough for this pet

const (
	LocalDeployPort         = ":2022"
	SectorDroneCronDuration = 9 * time.Second
)
