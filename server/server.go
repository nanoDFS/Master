package server

import (
	cms "github.com/nanoDFS/Master/server/chunkserver"
	fms "github.com/nanoDFS/Master/server/file"
)

type MasterServer struct {
	CMS *cms.CSMetadataServer
	FMS *fms.FileMetadataServer
}

func NewMasterServerRunner(faddr string, caddr string) (*MasterServer, error) {
	cms, _ := cms.NewCSMetadataServerRunner(caddr)
	fms, _ := fms.NewFileMetadataServerRunner(faddr)

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
