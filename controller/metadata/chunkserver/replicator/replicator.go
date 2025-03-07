package replicator

import (
	"fmt"

	cs "github.com/nanoDFS/Master/controller/metadata/chunkserver"
	dm "github.com/nanoDFS/Master/utils/datamodel"
)

// Replicas can be modified concurrently to point to different chunk servers
type Replicas struct {
	Primary   *dm.ConcurrentValue[*cs.ChunkServer]
	Secondary *dm.ConcurrentValue[*cs.ChunkServer]
	Tertiary  *dm.ConcurrentValue[*cs.ChunkServer]
}

func NewReplicas(p, s, t *cs.ChunkServer) *Replicas {
	return &Replicas{
		Primary:   dm.NewConcurrentValue(p),
		Secondary: dm.NewConcurrentValue(s),
		Tertiary:  dm.NewConcurrentValue(t),
	}
}

func (t *Replicas) GetActiveReplica() (*cs.ChunkServer, error) {
	if t.Primary.Get().IsActive() {
		return t.Primary.Get(), nil
	} else if t.Secondary.Get().IsActive() {
		return t.Secondary.Get(), nil
	} else if t.Tertiary.Get().IsActive() {
		return t.Tertiary.Get(), nil
	} else {
		return nil, fmt.Errorf("couldn't find any active servers")
	}
}
