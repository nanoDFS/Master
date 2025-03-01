package test

import (
	"context"
	"testing"

	"github.com/charmbracelet/log"
	"github.com/nanoDFS/Master/server"
	fileserver "github.com/nanoDFS/Master/server/proto"
	"github.com/nanoDFS/Master/utils"
	"google.golang.org/grpc"
)

func TestNewMasterServerRunner(t *testing.T) {
	port := utils.RandLocalAddr()
	master, err := server.NewMasterServerRunner(port)
	if err != nil {
		t.Errorf("failed to start server: %v", err)
	}
	master.Listen()
}

func TestRegister(t *testing.T) {
	port := utils.RandLocalAddr()
	master, _ := server.NewMasterServerRunner(port)
	master.Listen()

	conn, err := grpc.NewClient(port, grpc.WithInsecure())
	if err != nil {
		log.Errorf("did not connect: %v", err)
	}
	defer conn.Close()

	client := fileserver.NewFileServiceClient(conn)
	resp, _ := client.Register(context.Background(), &fileserver.ChunkServerRegisterReq{
		Address: utils.RandLocalAddr(),
	})
	expected := fileserver.RegisterResp{Status: true}
	if resp.String() != expected.String() {
		t.Errorf("expected %s , got %s", expected.String(), resp.String())
	}

}
func TestUploadAndDelete(t *testing.T) {
	port := utils.RandLocalAddr()
	master, _ := server.NewMasterServerRunner(port)
	master.Listen()

	conn, err := grpc.NewClient(port, grpc.WithInsecure())
	if err != nil {
		t.Errorf("did not connect: %v", err)
	}
	defer conn.Close()
	client := fileserver.NewFileServiceClient(conn)

	client.Register(context.Background(), &fileserver.ChunkServerRegisterReq{
		Address: utils.RandLocalAddr(),
	})
	client.Register(context.Background(), &fileserver.ChunkServerRegisterReq{
		Address: utils.RandLocalAddr(),
	})

	_, err = client.UploadFile(context.Background(), &fileserver.FileUploadReq{
		FileId: "some-user-generated-file-id",
		UserId: "some-user-id",
		Size:   180,
	})

	if err != nil {
		t.Errorf("failed to upload file, %v", err)
	}

	_, err = client.DownloadFile(context.Background(), &fileserver.FileDownloadReq{
		FileId: "some-user-generated-file-id",
	})
	if err != nil {
		t.Errorf("failed to delete file, %v", err)
	}
	_, err = client.DeleteFile(context.Background(), &fileserver.FileDeleteReq{
		FileId: "some-user-generated-file-id",
	})
	if err != nil {
		t.Errorf("failed to delete file, %v", err)
	}

}
