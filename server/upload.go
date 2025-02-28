package server

import (
	"context"
	"log"

	"github.com/nanoDFS/Master/controller/acl"
	"github.com/nanoDFS/Master/controller/metadata"
	fileserver "github.com/nanoDFS/Master/server/proto"
)

func (t Server) UploadFile(ctx context.Context, req *fileserver.FileUploadReq) (*fileserver.UploadResp, error) {
	fileHandler := metadata.GetFileController()
	access := acl.NewACL(req.UserId)
	file := fileHandler.Create(req.FileId, req.UserId, access, req.Size)
	token, _ := acl.NewJWT().Generate(&acl.Claims{UserId: req.UserId, Access: *access, Size: req.Size})

	var chunk_servers []*fileserver.ChunkServer
	for _, server := range file.Chunks {
		chunk_servers = append(chunk_servers, &fileserver.ChunkServer{
			Address: server[0].Addr.String(),
		})
	}
	log.Printf("File upload has be initiated successfully for fileId: %s, userId: %s", req.GetFileId(), req.GetUserId())
	log.Printf("Selected chunk servers: %s", chunk_servers)
	return &fileserver.UploadResp{
		Success:      true,
		ChunkServers: chunk_servers,
		AccessToken:  token,
	}, nil
}
