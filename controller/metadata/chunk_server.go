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
		mu:    sync.RWMutex{},
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
	mu           sync.RWMutex
	ChunkServers map[string]*ChunkServer
}

var chunkServerMetadataInstance = &ChunkServerMetadata{ChunkServers: make(map[string]*ChunkServer), mu: sync.RWMutex{}}

func GetChunkServerMetadata() *ChunkServerMetadata {
	return chunkServerMetadataInstance
}

func (t *ChunkServerMetadata) Register(addr string, space int64) {
	tcp_addr, _ := net.ResolveTCPAddr("tcp", addr)
	t.mu.Lock()
	t.ChunkServers[addr] = NewChunkServer(tcp_addr, space)
	t.mu.Unlock()
}

func (t *ChunkServerMetadata) Drop(addr string) {
	t.mu.Lock()
	delete(t.ChunkServers, addr)
	t.mu.Unlock()
}

func (t *ChunkServerMetadata) GetAllChunkServers() []*ChunkServer {
	var servers []*ChunkServer
	t.mu.RLock()
	for _, v := range t.ChunkServers {
		servers = append(servers, v)
	}
	t.mu.RUnlock()
	return servers
}

func (t *ChunkServerMetadata) GetChunkServer(addr string) (*ChunkServer, error) {
	t.mu.RLock()
	if _, ok := t.ChunkServers[addr]; !ok {
		t.mu.RUnlock()
		return nil, fmt.Errorf("failed to fetch chunk server")
	}
	res := t.ChunkServers[addr]
	t.mu.RUnlock()
	return res, nil
}
