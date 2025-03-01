package filemetadata

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"

	"github.com/nanoDFS/Master/controller/acl"
	"github.com/nanoDFS/Master/controller/metadata"
	fms "github.com/nanoDFS/Master/server/file/proto"
)

func (t Server) DownloadFile(ctx context.Context, req *fms.FileDownloadReq) (*fms.DownloadResp, error) {
	fileHandler := metadata.GetFileController()
	file, err := fileHandler.Get(req.FileId)
	if err != nil {
		return nil, fmt.Errorf("failed to delete file, %v", err)
	}

	token, _ := acl.NewJWT().Generate(&acl.Claims{UserId: file.GetOwnerID(), FileId: file.GetID(), Access: *file.GetACL(), Size: file.Size})
	chunk_servers := getChunkServers(file)

	log.Infof("File download has been initiated successfully for fileId: %s", req.GetFileId())
	log.Debugf("Selected chunk servers: %s", chunk_servers)
	return &fms.DownloadResp{
		ChunkServers: chunk_servers,
		AccessToken:  token,
	}, nil
}
