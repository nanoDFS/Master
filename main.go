package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"

	"github.com/nanoDFS/Master/monitor"
	server "github.com/nanoDFS/Master/server"
	"github.com/nanoDFS/Master/utils"
)

func createSingleMaster(faddr string, caddr string) {
	master, _ := server.NewMasterServerRunner(faddr, caddr)
	if err := master.Listen(); err != nil {
		fmt.Printf("failed to start listening %v", err)
	}
}

func main() {
	utils.InitLog()

	createSingleMaster(os.Getenv("FADDR"), os.Getenv("CADDR"))

	port := utils.RandLocalAddr()
	m, err := monitor.NewMonitor(port)
	m.Start()
	if err != nil {
		log.Errorf("failed to create monitor , %v", err)
	}

	select {}
}
