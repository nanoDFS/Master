package chunkservermetadata

import (
	"net"

	"github.com/charmbracelet/log"
	css "github.com/nanoDFS/Master/server/chunkserver/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	css.UnimplementedChunkServerRegisterServiceServer
}

type CSMetadataServer struct {
	Addr     net.Addr
	listener *net.Listener
	server   *grpc.Server
}

func NewCSMetadataServerRunner(addr string) (*CSMetadataServer, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	s := grpc.NewServer()
	css.RegisterChunkServerRegisterServiceServer(s, Server{})
	reflection.Register(s)
	return &CSMetadataServer{
		Addr:     listener.Addr(),
		listener: &listener,
		server:   s,
	}, nil
}

func (t *CSMetadataServer) Listen() error {
	go func() {
		log.Infof("started chunk server metadata service, listening on port: %s", t.Addr)
		if err := t.server.Serve(*t.listener); err != nil {
			log.Fatalf("failed to listen on port %s", t.Addr)
		}
	}()
	return nil
}

func (t *CSMetadataServer) Stop() {
	t.server.Stop()
}
