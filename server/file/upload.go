package filemetadata

import (
	"context"

	"github.com/charmbracelet/log"

	"github.com/nanoDFS/Master/controller/auth"
	"github.com/nanoDFS/Master/controller/auth/acl"
	"github.com/nanoDFS/Master/controller/metadata"
	fms "github.com/nanoDFS/Master/server/file/proto"
)

func (t Server) UploadFile(ctx context.Context, req *fms.FileUploadReq) (*fms.UploadResp, error) {
	fileHandler := metadata.GetFileController()
	access := acl.NewACL(req.UserId)
	log.Debugf("Waiting : %s", req.FileId)
	file := fileHandler.Create(req.FileId, req.UserId, access, req.Size)

	token, _ := auth.NewAuth().AuthorizeRead(req.UserId, *file, *file.GetACL(), file.Size.Get())
	chunk_servers := getChunkServers(file)

	log.Infof("File upload has been initiated successfully for fileId: %s, userId: %s", req.GetFileId(), req.GetUserId())
	log.Debugf("Selected chunk servers: %s", chunk_servers)
	return &fms.UploadResp{
		Success:      true,
		ChunkServers: chunk_servers,
		AccessToken:  token,
	}, nil
}
