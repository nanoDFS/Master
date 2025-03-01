package main

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"

	"github.com/nanoDFS/Master/health"
	server "github.com/nanoDFS/Master/server"
	cms_pb "github.com/nanoDFS/Master/server/chunk_server_metadata/proto"
	fms_pb "github.com/nanoDFS/Master/server/file_metadata/proto"
	"github.com/nanoDFS/Master/utils"
	"google.golang.org/grpc"
)

func createSingleMaster(faddr string, caddr string) {
	master, _ := server.NewMasterServerRunner(faddr, caddr)
	if err := master.Listen(); err != nil {
		fmt.Printf("failed to start listening %v", err)
	}
}

func main() {
	utils.InitLog()

	createSingleMaster(":9000", ":9001")
	createSingleMaster(":8001", ":8002")
	createSingleMaster(":8003", ":8004")

	addCS(":9001")
	addCS(":9002")
	addCS(":8004")

	go test(":9000", 0)
	go test(":8001", 1)
	go test(":8003", 2)

	port := utils.RandLocalAddr()
	monitor, err := health.NewHealthMonitor(port)
	monitor.Listen()
	if err != nil {
		log.Errorf("failed to create monitor , %v", err)
	}

	select {}
}

func addCS(caddr string) {
	conn, err := grpc.NewClient(caddr, grpc.WithInsecure())
	if err != nil {
		log.Errorf("did not connect: %v", err)
	}
	defer conn.Close()
	client := cms_pb.NewChunkServerRegisterServiceClient(conn)
	func() {
		for range 3 {
			client.Register(context.Background(), &cms_pb.ChunkServerRegisterReq{
				Address: utils.RandLocalAddr(),
			})
		}
	}()
}

func test(port string, rid int) {
	conn, err := grpc.NewClient(port, grpc.WithInsecure())
	if err != nil {
		log.Errorf("did not connect: %v", err)
	}
	defer conn.Close()
	client := fms_pb.NewFileMetadataServiceClient(conn)

	for i := range 4 {
		client.UploadFile(context.Background(), &fms_pb.FileUploadReq{
			FileId: fmt.Sprintf("some-user-generated-file-id-%d-%d", i, rid),
			UserId: fmt.Sprintf("some-user-id-%d-%d", i, rid),
			Size:   180,
		})
	}
}
