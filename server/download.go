package server

import (
	"context"

	fileserver "github.com/nanoDFS/Master/server/proto"
)

func (t Server) DownloadFile(context.Context, *fileserver.FileDownloadReq) (*fileserver.DownloadResp, error) {
	return nil, nil
}
