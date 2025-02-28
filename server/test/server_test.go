package test

import (
	"context"
	"log"
	"testing"

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
	if err := master.Listen(); err != nil {
		t.Errorf("failed to start listening at port: %s : %v", port, err)
	}
}

func TestRegister(t *testing.T) {
	port := utils.RandLocalAddr()
	master, _ := server.NewMasterServerRunner(port)
	if err := master.Listen(); err != nil {
		t.Errorf("failed to start listening at port: %s : %v", port, err)
	}

	conn, err := grpc.NewClient(port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
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
func TestUpload(t *testing.T) {
	port := utils.RandLocalAddr()
	master, _ := server.NewMasterServerRunner(port)
	if err := master.Listen(); err != nil {
		t.Errorf("failed to start listening at port: %s : %v", port, err)
	}

	conn, err := grpc.NewClient(port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
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

}

func TestStop(t *testing.T) {
	port := utils.RandLocalAddr()
	master, _ := server.NewMasterServerRunner(port)
	if err := master.Listen(); err != nil {
		t.Errorf("failed to start listening at port: %s : %v", port, err)
	}

	master.Stop()

	conn, err := grpc.NewClient(port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := fileserver.NewFileServiceClient(conn)
	_, err = client.Register(context.Background(), &fileserver.ChunkServerRegisterReq{
		Address: utils.RandLocalAddr(),
	})
	if err == nil {
		t.Errorf("failed to close server")
	}

}
