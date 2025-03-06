package filemetadata

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"

	"github.com/nanoDFS/Master/controller/auth"
	"github.com/nanoDFS/Master/controller/metadata"
	fms "github.com/nanoDFS/Master/server/file/proto"
)

func (t Server) DownloadFile(ctx context.Context, req *fms.FileDownloadReq) (*fms.DownloadResp, error) {
	fileHandler := metadata.GetFileController()
	file, err := fileHandler.Get(req.FileId)
	if err != nil {
		return nil, fmt.Errorf("failed to read file, %v", err)
	}

	token, err := auth.NewAuth().AuthorizeRead(req.UserId, *file, *file.GetACL(), file.Size.Get()) // TODO: sending *file might cause concurrency issue
	if err != nil {
		return &fms.DownloadResp{
			Success: false,
		}, err
	}

	chunk_servers := getChunkServers(file)

	log.Infof("File download has been initiated successfully for fileId: %s", req.GetFileId())
	log.Debugf("Selected chunk servers: %s", chunk_servers)
	return &fms.DownloadResp{
		Success:      true,
		ChunkServers: chunk_servers,
		AccessToken:  token,
	}, nil
}
