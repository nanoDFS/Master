package test

import (
	"context"
	"log"
	"testing"

	"github.com/nanoDFS/Master/server"
	fileserver "github.com/nanoDFS/Master/server/proto"
	"google.golang.org/grpc"
)

func TestNewMasterServerRunner(t *testing.T) {
	port := ":8000"
	master, err := server.NewMasterServerRunner(port)
	if err != nil {
		t.Errorf("failed to start server: %v", err)
	}
	if err := master.Listen(); err != nil {
		t.Errorf("failed to start listening at port: %s : %v", port, err)
	}
}

func TestRegister(t *testing.T) {
	port := ":8003"
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
		Address: ":8000",
	})
	expected := fileserver.RegisterResp{Status: true}
	if resp.String() != expected.String() {
		t.Errorf("expected %s , got %s", expected.String(), resp.String())
	}

}
func TestUpload(t *testing.T) {
	port := ":8003"
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
		Address: ":8000",
	})
	client.Register(context.Background(), &fileserver.ChunkServerRegisterReq{
		Address: ":8002",
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
	port := ":8001"
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
		Address: ":8000",
	})
	if err == nil {
		t.Errorf("failed to close server")
	}

}
