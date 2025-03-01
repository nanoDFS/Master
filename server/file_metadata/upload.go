package filemetadata

import (
	"context"

	"github.com/charmbracelet/log"

	"github.com/nanoDFS/Master/controller/acl"
	"github.com/nanoDFS/Master/controller/metadata"
	fms "github.com/nanoDFS/Master/server/file_metadata/proto"
)

func (t Server) UploadFile(ctx context.Context, req *fms.FileUploadReq) (*fms.UploadResp, error) {
	fileHandler := metadata.GetFileController()
	access := acl.NewACL(req.UserId)
	log.Debugf("Waiting : %s", req.FileId)
	file := fileHandler.Create(req.FileId, req.UserId, access, req.Size)

	token, _ := acl.NewJWT().Generate(&acl.Claims{UserId: req.UserId, Access: *file.GetACL(), Size: req.Size})

	chunk_servers := getChunkServers(file)

	log.Infof("File upload has been initiated successfully for fileId: %s, userId: %s", req.GetFileId(), req.GetUserId())
	log.Debugf("Selected chunk servers: %s", chunk_servers)
	return &fms.UploadResp{
		Success:      true,
		ChunkServers: chunk_servers,
		AccessToken:  token,
	}, nil
}
