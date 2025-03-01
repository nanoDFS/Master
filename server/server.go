package server

import (
	cms "github.com/nanoDFS/Master/server/chunk_server_metadata"
	fms "github.com/nanoDFS/Master/server/file_metadata"
)

type MasterServer struct {
	CMS *cms.MasterServer
	FMS *fms.MasterServer
}

func NewMasterServerRunner(faddr string, caddr string) (*MasterServer, error) {
	cms, _ := cms.NewMasterServerRunner(caddr)
	fms, _ := fms.NewMasterServerRunner(faddr)

	return &MasterServer{
		CMS: cms,
		FMS: fms,
	}, nil
}

func (t *MasterServer) Listen() error {
	err := t.FMS.Listen()
	if err != nil {
		return err
	}
	err = t.CMS.Listen()
	if err != nil {
		return err
	}
	return nil
}

func (t *MasterServer) Stop() {
	t.FMS.Stop()
	t.CMS.Stop()
}
