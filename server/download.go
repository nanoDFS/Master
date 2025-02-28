package server

import (
	"context"
	"fmt"
	"log"

	"github.com/nanoDFS/Master/controller/acl"
	"github.com/nanoDFS/Master/controller/metadata"
	fileserver "github.com/nanoDFS/Master/server/proto"
)

func (t Server) DownloadFile(ctx context.Context, req *fileserver.FileDownloadReq) (*fileserver.DownloadResp, error) {
	fileHandler := metadata.GetFileController()
	file, err := fileHandler.Get(req.FileId)
	if err != nil {
		return nil, fmt.Errorf("failed to delete error, %v", err)
	}

	token, _ := acl.NewJWT().Generate(&acl.Claims{UserId: file.GetUserID(), Access: *file.GetACL(), Size: file.Size})
	chunk_servers := getChunkServers(file)

	log.Printf("File download has been initiated successfully for fileId: %s", req.GetFileId())
	log.Printf("Selected chunk servers: %s", chunk_servers)
	return &fileserver.DownloadResp{
		ChunkServers: chunk_servers,
		AccessToken:  token,
	}, nil
}
