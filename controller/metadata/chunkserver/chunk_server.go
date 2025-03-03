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
	Addr   net.Addr
	status *dm.ConcurrentValue[Status]
	space  *dm.ConcurrentValue[int64] // in bytes
}

func NewChunkServer(addr net.Addr, space int64) *ChunkServer {
	return &ChunkServer{
		Addr:   addr,
		status: dm.NewConcurrentValue(Active),
		space:  dm.NewConcurrentValue(space),
	}
}

func (t *ChunkServer) GetStatus() Status {
	return t.status.Get()
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

func (t *ChunkServerMetadata) Register(addr string, space int64) {
	tcp_addr, _ := net.ResolveTCPAddr("tcp", addr)
	t.chunkServers.Set(addr, NewChunkServer(tcp_addr, space))
}

func (t *ChunkServerMetadata) Drop(addr string) {
	t.chunkServers.Delete(addr)
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
