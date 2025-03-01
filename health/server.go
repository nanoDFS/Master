package health

import (
	"time"

	"github.com/charmbracelet/log"
	"github.com/nanoDFS/Master/controller/metadata"
	"github.com/nanoDFS/p2p/p2p/transport"
)

type HealthMonitor struct {
	server              *transport.TCPTransport
	chunkServerMetadata *metadata.ChunkServerMetadata
	quitChan            chan struct{}
}

func NewHealthMonitor(addr string) (*HealthMonitor, error) {
	server, err := transport.NewTCPTransport(addr)
	chunkServerMetadata := metadata.GetChunkServerMetadata()
	return &HealthMonitor{
		server:              server,
		chunkServerMetadata: chunkServerMetadata,
		quitChan:            make(chan struct{}),
	}, err
}

func (t *HealthMonitor) Listen() {
	go t.listen()
}

func (t *HealthMonitor) listen() {
	for {
		select {
		case <-t.quitChan:
			return
		default:
			log.Info("PING: started")
			for _, s := range t.chunkServerMetadata.GetAllChunkServers() {
				go func(s *metadata.ChunkServer) {
					t.ping(s)
				}(s)
			}
		}
		time.Sleep(time.Second * 3) // TODO: read from config
	}
}

func (t *HealthMonitor) ping(server *metadata.ChunkServer) {
	if err := t.server.Send(server.Addr.String(), "health/heartbeat"); err != nil {
		server.SetStatus(metadata.Inactive)
		log.Warnf("PING: found inactive server: %s", server.Addr)
	}
	t.server.Close(server.Addr.String())
}

func (t *HealthMonitor) Stop() {
	log.Info("sutting down monitor server")
	t.quitChan <- struct{}{}
}
