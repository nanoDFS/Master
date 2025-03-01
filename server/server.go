package server

import (
	"net"

	"github.com/charmbracelet/log"

	"github.com/nanoDFS/Master/controller/metadata"
	fileserver "github.com/nanoDFS/Master/server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	fileserver.UnimplementedFileServiceServer
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
	fileserver.RegisterFileServiceServer(s, Server{})
	reflection.Register(s)
	return &MasterServer{
		Addr:     listener.Addr(),
		listener: &listener,
		server:   s,
	}, nil
}

func (t *MasterServer) Listen() error {
	go func() {
		log.Infof("started listening on port: %s", t.Addr)
		if err := t.server.Serve(*t.listener); err != nil {
			log.Fatalf("failed to listen on port %s", t.Addr)
		}
	}()
	return nil
}

func (t *MasterServer) Stop() {
	t.server.Stop()
}

func getChunkServers(file *metadata.File) []*fileserver.ChunkServer {
	var chunk_servers []*fileserver.ChunkServer
	for _, server := range file.Chunks {
		chunk_servers = append(chunk_servers, &fileserver.ChunkServer{
			Address: server[0].Addr.String(),
		})
	}
	return chunk_servers
}
