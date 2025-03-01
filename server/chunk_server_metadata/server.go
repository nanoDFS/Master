package chunkservermetadata

import (
	"net"

	"github.com/charmbracelet/log"
	css "github.com/nanoDFS/Master/server/chunk_server_metadata/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	css.UnimplementedChunkServerRegisterServiceServer
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
	css.RegisterChunkServerRegisterServiceServer(s, Server{})
	reflection.Register(s)
	return &MasterServer{
		Addr:     listener.Addr(),
		listener: &listener,
		server:   s,
	}, nil
}

func (t *MasterServer) Listen() error {
	go func() {
		log.Infof("started chunk server metadata service, listening on port: %s", t.Addr)
		if err := t.server.Serve(*t.listener); err != nil {
			log.Fatalf("failed to listen on port %s", t.Addr)
		}
	}()
	return nil
}

func (t *MasterServer) Stop() {
	t.server.Stop()
}
