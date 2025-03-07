package chunkserver

import (
	"fmt"

	lb "github.com/nanoDFS/Master/controller/metadata/chunkserver/loadbalancer"
	"github.com/nanoDFS/Master/utils/config"
)

type Chunker struct {
	startegy lb.LoadBalancingStrategy
}

func NewChunker(strategy lb.LoadBalancingStrategy) *Chunker {
	return &Chunker{
		startegy: strategy,
	}
}

func (t *Chunker) Generate(fileId string, size int64) []*ChunkServer {
	chunkSize := config.LoadConfig().Chunk.Size

	count := size / chunkSize
	if size%chunkSize != 0 {
		count++
	}

	var servers []*ChunkServer

	allChunkServers := GetChunkServerMetadata().GetAllActiveChunkServers()
	for i := range count {
		index, err := t.startegy.GetIndex(lb.Opts{Key: fileId + fmt.Sprint(i), Length: len(allChunkServers)})
		if err != nil {
			return servers
		}
		servers = append(servers, allChunkServers[index])
	}

	return servers
}
