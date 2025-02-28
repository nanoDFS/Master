package server

import (
	"context"
	"log"

	chunk_server "github.com/nanoDFS/Master/controller/metadata"
	fileserver "github.com/nanoDFS/Master/server/proto"
)

func (t Server) Register(ctx context.Context, req *fileserver.ChunkServerRegisterReq) (*fileserver.RegisterResp, error) {
	metadata_ctl := chunk_server.GetChunkServerMetadata()
	metadata_ctl.Register(req.Address, req.Space)
	log.Printf("Registered %s as a chunk server", req.Address)
	return &fileserver.RegisterResp{
		Status: true,
	}, nil
}
