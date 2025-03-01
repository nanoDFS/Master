package chunkservermetadata

import (
	"context"

	"github.com/charmbracelet/log"

	chunk_server "github.com/nanoDFS/Master/controller/metadata"
	css "github.com/nanoDFS/Master/server/chunk_server_metadata/proto"
)

func (t Server) Register(ctx context.Context, req *css.ChunkServerRegisterReq) (*css.RegisterResp, error) {
	metadata_ctl := chunk_server.GetChunkServerMetadata()
	metadata_ctl.Register(req.Address, req.Space)
	log.Infof("Registered %s as a chunk server", req.Address)
	return &css.RegisterResp{
		Status: true,
	}, nil
}
