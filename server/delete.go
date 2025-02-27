package server

import (
	"context"

	fileserver "github.com/nanoDFS/Master/server/proto"
)

func (t Server) DeleteFile(context.Context, *fileserver.FileDeleteReq) (*fileserver.DeleteResp, error) {
	return nil, nil
}
