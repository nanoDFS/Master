package filemetadata

import (
	"context"
	"testing"

	"github.com/charmbracelet/log"
	cms "github.com/nanoDFS/Master/server/chunk_server_metadata"
	cms_pb "github.com/nanoDFS/Master/server/chunk_server_metadata/proto"
	fms "github.com/nanoDFS/Master/server/file_metadata"
	fms_pb "github.com/nanoDFS/Master/server/file_metadata/proto"
	"github.com/nanoDFS/Master/utils"
	"google.golang.org/grpc"
)

func registerCS() *cms_pb.ChunkServerRegisterServiceClient {
	port := utils.RandLocalAddr()
	master, _ := cms.NewMasterServerRunner(port)
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
	master, err := fms.NewMasterServerRunner(port)
	if err != nil {
		t.Errorf("failed to start server: %v", err)
	}
	master.Listen()
}

func TestUploadAndDelete(t *testing.T) {
	port := utils.RandLocalAddr()
	master, _ := fms.NewMasterServerRunner(port)
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

	_, err = client.DownloadFile(context.Background(), &fms_pb.FileDownloadReq{
		FileId: "some-user-generated-file-id",
	})
	if err != nil {
		t.Errorf("failed to delete file, %v", err)
	}
	_, err = client.DeleteFile(context.Background(), &fms_pb.FileDeleteReq{
		FileId: "some-user-generated-file-id",
	})
	if err != nil {
		t.Errorf("failed to delete file, %v", err)
	}

}
