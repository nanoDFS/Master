package chunkservermetadata

import (
	"context"
	"testing"

	"github.com/charmbracelet/log"
	cms "github.com/nanoDFS/Master/server/chunkserver"
	cms_pb "github.com/nanoDFS/Master/server/chunkserver/proto"
	"github.com/nanoDFS/Master/utils"
	"google.golang.org/grpc"
)

func TestNewMasterServerRunner(t *testing.T) {
	port := utils.RandLocalAddr()
	master, err := cms.NewMasterServerRunner(port)
	if err != nil {
		t.Errorf("failed to start server: %v", err)
	}
	master.Listen()
}

func TestRegister(t *testing.T) {
	port := utils.RandLocalAddr()
	master, _ := cms.NewMasterServerRunner(port)
	master.Listen()

	conn, err := grpc.NewClient(port, grpc.WithInsecure())
	if err != nil {
		log.Errorf("did not connect: %v", err)
	}
	defer conn.Close()

	client := cms_pb.NewChunkServerRegisterServiceClient(conn)
	resp, _ := client.Register(context.Background(), &cms_pb.ChunkServerRegisterReq{
		Address: utils.RandLocalAddr(),
	})
	expected := cms_pb.RegisterResp{Status: true}
	if resp.String() != expected.String() {
		t.Errorf("expected %s , got %s", expected.String(), resp.String())
	}
}
