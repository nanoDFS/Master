package metadata

import (
	"fmt"
	"net"
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
}

func NewChunkServer(addr net.Addr) *ChunkServer {
	return &ChunkServer{
		Addr: addr,
	}
}

type ChunkServerMetadata struct {
	ChunkServers map[string]*ChunkServer
}

var chunkServerMetadataInstance = &ChunkServerMetadata{ChunkServers: make(map[string]*ChunkServer)}

func GetChunkServerMetadata() *ChunkServerMetadata {
	return chunkServerMetadataInstance
}

func (t *ChunkServerMetadata) Register(addr string) {
	tcp_addr, _ := net.ResolveTCPAddr("tcp", addr)
	t.ChunkServers[addr] = NewChunkServer(tcp_addr)
}

func (t *ChunkServerMetadata) Drop(addr string) {
	delete(t.ChunkServers, addr)
}

func (t *ChunkServerMetadata) GetStatus(addr string) (Status, error) {
	if _, ok := t.ChunkServers[addr]; !ok {
		return 0, fmt.Errorf("failed to fetch chunk server")
	}
	return t.ChunkServers[addr].Status, nil
}
func (t *ChunkServerMetadata) SetStatus(addr string, status Status) error {
	if _, ok := t.ChunkServers[addr]; !ok {
		return fmt.Errorf("failed to fetch chunk server")
	}
	t.ChunkServers[addr].Status = status
	return nil
}
