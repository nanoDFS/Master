package filemetadata

import (
	"context"
	"testing"

	"github.com/charmbracelet/log"
	cms "github.com/nanoDFS/Master/server/chunkserver"
	cms_pb "github.com/nanoDFS/Master/server/chunkserver/proto"
	fms "github.com/nanoDFS/Master/server/file"
	fms_pb "github.com/nanoDFS/Master/server/file/proto"
	"github.com/nanoDFS/Master/utils"
	"google.golang.org/grpc"
)

func registerCS() *cms_pb.ChunkServerRegisterServiceClient {
	port := utils.RandLocalAddr()
	master, _ := cms.NewCSMetadataServerRunner(port)
	master.Listen()
	conn, err := grpc.NewClient(port, grpc.WithInsecure())
	if err != nil {
		log.Errorf("did not connect: %v", err)
	}
	defer conn.Close()

	client := cms_pb.NewChunkServerRegisterServiceClient(conn)
	client.Register(context.Background(), &cms_pb.ChunkServerRegisterReq{
		Address: utils.RandLocalAddr(),
	})
	return &client
}

func TestNewMasterServerRunner(t *testing.T) {
	port := utils.RandLocalAddr()
	master, err := fms.NewFileMetadataServerRunner(port)
	if err != nil {
		t.Errorf("failed to start server: %v", err)
	}
	master.Listen()
}

func TestUploadAndDelete(t *testing.T) {
	port := utils.RandLocalAddr()
	master, _ := fms.NewFileMetadataServerRunner(port)
	master.Listen()

	registerCS()
	conn, err := grpc.NewClient(port, grpc.WithInsecure())
	if err != nil {
		log.Errorf("did not connect: %v", err)
	}
	defer conn.Close()

	client := fms_pb.NewFileMetadataServiceClient(conn)

	_, err = client.UploadFile(context.Background(), &fms_pb.FileUploadReq{
		FileId: "some-user-generated-file-id",
		UserId: "some-user-id",
		Size:   180,
	})

	if err != nil {
		t.Errorf("failed to upload file, %v", err)
	}

	downloadRes, err := client.DownloadFile(context.Background(), &fms_pb.FileDownloadReq{
		UserId: "some-user-id",
		FileId: "some-user-generated-file-id",
	})
	if err != nil || !downloadRes.Success {
		t.Errorf("failed to download file, %v", err)
	}
	delRes, err := client.DeleteFile(context.Background(), &fms_pb.FileDeleteReq{
		UserId: "some-user-id",
		FileId: "some-user-generated-file-id",
	})
	if err != nil || !delRes.Success {
		t.Errorf("failed to delete file, %v", err)
	}

}
