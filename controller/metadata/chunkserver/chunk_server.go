package chunkserver

import (
	"net"

	dm "github.com/nanoDFS/Master/utils/datamodel"
)

type Status int

const (
	Active   Status = iota // 0
	Inactive               // 1
)

type ChunkServer struct {
	MonitorAddr   net.Addr
	StreamingAddr net.Addr
	status        *dm.ConcurrentValue[Status]
	space         *dm.ConcurrentValue[int64] // in bytes
}

func NewChunkServer(monitorAddr net.Addr, streamingAddr net.Addr, space int64) *ChunkServer {
	return &ChunkServer{
		MonitorAddr:   monitorAddr,
		StreamingAddr: streamingAddr,
		status:        dm.NewConcurrentValue(Active),
		space:         dm.NewConcurrentValue(space),
	}
}

func (t *ChunkServer) GetStatus() Status {
	return t.status.Get()
}

func (t *ChunkServer) IsActive() bool {
	return t.status.Get() == Active
}

func (t *ChunkServer) SetStatus(status Status) {
	t.status.Set(status) // TODO: Should I keep ChunkServer when it becomes inactive
}
func (t *ChunkServer) GetSpace() int64 {
	return t.space.Get()
}

func (t *ChunkServer) SetSpaces(space int64) {
	t.space.Set(space)
}

type ChunkServerMetadata struct {
	chunkServers *dm.ConcurrentMap[string, *ChunkServer]
}

var chunkServerMetadataInstance = &ChunkServerMetadata{chunkServers: dm.NewConcurrentMap[string, *ChunkServer]()}

func GetChunkServerMetadata() *ChunkServerMetadata {
	return chunkServerMetadataInstance
}

func (t *ChunkServerMetadata) Register(monitorAddr string, streamingAddr string, space int64) {
	tcp_addr, _ := net.ResolveTCPAddr("tcp", monitorAddr)
	grpc_addr, _ := net.ResolveTCPAddr("tcp", streamingAddr)
	t.chunkServers.Set(monitorAddr, NewChunkServer(tcp_addr, grpc_addr, space))
}

func (t *ChunkServerMetadata) Drop(addr string) {
	t.chunkServers.Delete(addr)
}

func (t *ChunkServerMetadata) GetAllActiveChunkServers() []*ChunkServer {
	allCs := t.chunkServers.Values()
	activeCs := []*ChunkServer{}
	for _, cs := range allCs {
		if cs.GetStatus() == Active {
			activeCs = append(activeCs, cs)
		}
	}
	return activeCs
}

func (t *ChunkServerMetadata) GetAllChunkServers() []*ChunkServer {
	return t.chunkServers.Values()
}

func (t *ChunkServerMetadata) GetChunkServer(addr string) (*ChunkServer, error) {
	res, err := t.chunkServers.Get(addr)
	if err != nil {
		return nil, err
	}
	return res, nil
}
