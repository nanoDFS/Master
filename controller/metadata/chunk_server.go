package metadata

import (
	"fmt"
	"net"
	"sync"
)

type Status int

const (
	Pending    Status = iota // 0
	InProgress               // 1
	Completed                // 2
	Failed                   // 3
)

type ChunkServer struct {
	Addr   net.Addr
	Status Status
	Space  int64 // in bytes
	mu     sync.RWMutex
}

func NewChunkServer(addr net.Addr, space int64) *ChunkServer {
	return &ChunkServer{
		Addr:  addr,
		Space: space,
	}
}

func (t *ChunkServer) GetStatus() Status {
	t.mu.RLock()
	status := t.Status
	t.mu.RUnlock()
	return status
}

func (t *ChunkServer) SetStatus(status Status) {
	t.mu.Lock()
	t.Status = status
	t.mu.Unlock()
}
func (t *ChunkServer) GetSpace() int64 {
	t.mu.RLock()
	space := t.Space
	t.mu.RUnlock()
	return space
}

func (t *ChunkServer) SetSpaces(space int64) {
	t.mu.Lock()
	t.Space = space
	t.mu.Unlock()
}

type ChunkServerMetadata struct {
	ChunkServers map[string]*ChunkServer
}

var chunkServerMetadataInstance = &ChunkServerMetadata{ChunkServers: make(map[string]*ChunkServer)}

func GetChunkServerMetadata() *ChunkServerMetadata {
	return chunkServerMetadataInstance
}

func (t *ChunkServerMetadata) Register(addr string, space int64) {
	tcp_addr, _ := net.ResolveTCPAddr("tcp", addr)
	t.ChunkServers[addr] = NewChunkServer(tcp_addr, space)
}

func (t *ChunkServerMetadata) Drop(addr string) {
	delete(t.ChunkServers, addr)
}

func (t *ChunkServerMetadata) GetAllChunkServers() []*ChunkServer {
	var servers []*ChunkServer
	for _, v := range t.ChunkServers {
		servers = append(servers, v)
	}
	return servers
}

func (t *ChunkServerMetadata) GetChunkServer(addr string) (*ChunkServer, error) {
	if _, ok := t.ChunkServers[addr]; !ok {
		return nil, fmt.Errorf("failed to fetch chunk server")
	}
	return t.ChunkServers[addr], nil
}
