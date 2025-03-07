package health

import (
	"time"

	"github.com/charmbracelet/log"
	cs "github.com/nanoDFS/Master/controller/metadata/chunkserver"
	"github.com/nanoDFS/p2p/p2p/transport"
)

type HealthClient struct {
	client              *transport.TCPTransport
	chunkServerMetadata *cs.ChunkServerMetadata
	quitChan            chan struct{}
}

func NewHealthClient(addr string) (*HealthClient, error) {
	client, err := transport.NewTCPTransport(addr)
	chunkServerMetadata := cs.GetChunkServerMetadata()
	return &HealthClient{
		client:              client,
		chunkServerMetadata: chunkServerMetadata,
		quitChan:            make(chan struct{}),
	}, err
}

func (t *HealthClient) Start() {
	go t.start()
}

func (t *HealthClient) start() {
	for {
		select {
		case <-t.quitChan:
			return
		default:
			log.Debugf("PING: started")
			for _, s := range t.chunkServerMetadata.GetAllChunkServers() {
				go func(s *cs.ChunkServer) {
					t.ping(s)
				}(s)
			}
		}
		time.Sleep(time.Second * 3) // TODO: read from config
	}
}

func (t *HealthClient) ping(server *cs.ChunkServer) {
	if err := t.client.Send(server.MonitorAddr.String(), "health/heartbeat"); err != nil {
		server.SetStatus(cs.Inactive)
		log.Warnf("PING: found inactive server: %s", server.MonitorAddr)
	}
	t.client.Close(server.MonitorAddr.String())
}

func (t *HealthClient) Stop() {
	defer log.Info("sutting down monitor server")
	t.quitChan <- struct{}{}
}
