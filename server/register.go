package server

import (
	"context"

	chunk_server "github.com/nanoDFS/Master/controller/metadata"
	fileserver "github.com/nanoDFS/Master/server/proto"
)

func (t Server) Register(ctx context.Context, req *fileserver.ChunkServerRegisterReq) (*fileserver.RegisterResp, error) {
	metadata_ctl := chunk_server.GetChunkServerMetadata()
	metadata_ctl.Register(req.Address)
	return &fileserver.RegisterResp{
		SecretKey: []byte("some random secrete key"),
	}, nil
}
