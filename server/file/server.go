package filemetadata

import (
	"net"

	"github.com/charmbracelet/log"
	"github.com/nanoDFS/Master/controller/metadata"
	fms "github.com/nanoDFS/Master/server/file/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	fms.UnimplementedFileMetadataServiceServer
}

type MasterServer struct {
	Addr     net.Addr
	listener *net.Listener
	server   *grpc.Server
}

func NewMasterServerRunner(addr string) (*MasterServer, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	s := grpc.NewServer()
	fms.RegisterFileMetadataServiceServer(s, Server{})
	reflection.Register(s)
	return &MasterServer{
		Addr:     listener.Addr(),
		listener: &listener,
		server:   s,
	}, nil
}

func (t *MasterServer) Listen() error {
	go func() {
		log.Infof("started file metadata service, listening on port: %s", t.Addr)
		if err := t.server.Serve(*t.listener); err != nil {
			log.Fatalf("failed to listen on port %s", t.Addr)
		}
	}()
	return nil
}

func (t *MasterServer) Stop() {
	t.server.Stop()
}

func getChunkServers(file *metadata.File) []*fms.ChunkServer {
	var chunk_servers []*fms.ChunkServer
	for _, addr := range file.GetChunkServers() {
		chunk_servers = append(chunk_servers, &fms.ChunkServer{
			Address: addr,
		})
	}
	return chunk_servers
}
