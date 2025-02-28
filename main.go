package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nanoDFS/Master/server"
	fileserver "github.com/nanoDFS/Master/server/proto"
	"github.com/nanoDFS/Master/utils"
	"google.golang.org/grpc"
)

func createSingleMaster(port string) {
	master, _ := server.NewMasterServerRunner(port)
	if err := master.Listen(); err != nil {
		fmt.Printf("failed to start listening at port: %s : %v", port, err)
	}
}

func main() {
	createSingleMaster(":9000")
	createSingleMaster(":8000")
	createSingleMaster(":8004")

	go test(":9000")
	go test(":8000")
	go test(":8004")
	select {}
}

func test(port string) {
	conn, err := grpc.NewClient(port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := fileserver.NewFileServiceClient(conn)

	go func() {
		for range 10 {
			client.Register(context.Background(), &fileserver.ChunkServerRegisterReq{
				Address: utils.RandLocalAddr(),
			})
		}
	}()

	for i := range 4 {
		client.UploadFile(context.Background(), &fileserver.FileUploadReq{
			FileId: fmt.Sprintf("some-user-generated-file-id-%d", i),
			UserId: fmt.Sprintf("some-user-id-%d", i),
			Size:   180,
		})
	}
}
