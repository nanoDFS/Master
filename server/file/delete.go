package filemetadata

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"

	"github.com/nanoDFS/Master/controller/auth"
	"github.com/nanoDFS/Master/controller/metadata"
	fms "github.com/nanoDFS/Master/server/file/proto"
)

func (t Server) DeleteFile(ctx context.Context, req *fms.FileDeleteReq) (*fms.DeleteResp, error) {
	fileHandler := metadata.GetFileController()
	file, err := fileHandler.Delete(req.FileId)
	if err != nil {
		return nil, fmt.Errorf("failed to delete error, %v", err)
	}

	token, err := auth.NewAuth().AuthorizeDelete(req.UserId, *file, *file.GetACL(), file.Size.Get())
	if err != nil {
		return &fms.DeleteResp{
			Success: false,
		}, err
	}
	chunk_servers, err := getChunkServers(file)
	if err != nil {
		return &fms.DeleteResp{
			Success: false,
		}, err
	}

	log.Infof("File delete has been initiated successfully for fileId: %s", req.GetFileId())
	log.Debugf("Selected chunk servers: %s", chunk_servers)
	return &fms.DeleteResp{
		Success:      true,
		ChunkServers: chunk_servers,
		AccessToken:  token,
	}, nil
}
