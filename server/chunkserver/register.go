package chunkservermetadata

import (
	"context"

	"github.com/charmbracelet/log"

	cs "github.com/nanoDFS/Master/controller/metadata/chunkserver"
	css "github.com/nanoDFS/Master/server/chunkserver/proto"
)

func (t Server) Register(ctx context.Context, req *css.ChunkServerRegisterReq) (*css.RegisterResp, error) {
	metadata_ctl := cs.GetChunkServerMetadata()
	metadata_ctl.Register(req.MonitorAddress, req.StreamingAddress, req.Space)
	log.Infof("Registered %s as a chunk server with streaming on %s", req.MonitorAddress, req.StreamingAddress)
	return &css.RegisterResp{
		Success: true,
	}, nil
}
