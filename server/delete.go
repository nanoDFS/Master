package server

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"

	"github.com/nanoDFS/Master/controller/acl"
	"github.com/nanoDFS/Master/controller/metadata"
	fileserver "github.com/nanoDFS/Master/server/proto"
)

func (t Server) DeleteFile(ctx context.Context, req *fileserver.FileDeleteReq) (*fileserver.DeleteResp, error) {
	fileHandler := metadata.GetFileController()
	file, err := fileHandler.Delete(req.FileId)
	if err != nil {
		return nil, fmt.Errorf("failed to delete error, %v", err)
	}

	token, _ := acl.NewJWT().Generate(&acl.Claims{UserId: file.GetUserID(), Access: *file.GetACL(), Size: file.Size})

	chunk_servers := getChunkServers(file)

	log.Infof("File delete has been initiated successfully for fileId: %s", req.GetFileId())
	log.Debugf("Selected chunk servers: %s", chunk_servers)
	return &fileserver.DeleteResp{
		ChunkServers: chunk_servers,
		AccessToken:  token,
	}, nil
}
