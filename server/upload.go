package server

import (
	"context"
	"fmt"

	fileserver "github.com/nanoDFS/Master/server/proto"
)

func (t Server) UploadFile(context.Context, *fileserver.FileUploadReq) (*fileserver.UploadResp, error) {
	fmt.Print("HI")
	return nil, nil
}
