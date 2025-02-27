package server

import (
	"net"

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
		if err := t.server.Serve(*t.listener); err != nil {
		}
	}()
	return nil
}

func (t *MasterServer) Stop() {
	t.server.Stop()
}
