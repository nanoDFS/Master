package metadata

import (
	"sync"

	"fmt"

	"github.com/nanoDFS/Master/controller/acl"
	loadbalancing "github.com/nanoDFS/Master/controller/metadata/load_balancing"
	"github.com/nanoDFS/Master/utils/config"
)

type File struct {
	id     string
	userId string
	acl    *acl.ACL
	Size   int64
	Chunks [][]*ChunkServer

	mu sync.RWMutex
}

func newFile(id string, userId string, acl *acl.ACL, size int64) *File {
	var chunks [][]*ChunkServer
	for _, k := range generateChunks(id, size) {
		chunks = append(chunks, []*ChunkServer{k})
	}
	return &File{
		id:     id,
		userId: userId,
		acl:    acl,
		Size:   size,
		Chunks: chunks,
	}
}

func (t *File) GetACL() *acl.ACL {
	t.mu.RLock()
	acl_ := t.acl
	t.mu.RUnlock()
	return acl_
}
func (t *File) GetUserID() string {
	t.mu.RLock()
	id := t.userId
	t.mu.RUnlock()
	return id
}
func (t *File) GetID() string {
	t.mu.RLock()
	id := t.id
	t.mu.RUnlock()
	return id
}

func generateChunks(fileId string, size int64) []*ChunkServer {
	chunkSize := config.LoadConfig().Chunk.Size
	count := size / chunkSize
	if size%chunkSize != 0 {
		count++
	}

	var servers []*ChunkServer
	allChunkServers := GetChunkServerMetadata().GetAllChunkServers()
	loadbalancer := loadbalancing.NewConsistentHashing()
	for i := range count {
		index := loadbalancer.GetIndex(fileId+fmt.Sprint(i), len(allChunkServers))
		servers = append(servers, allChunkServers[index])
	}
	return servers
}

// FileController provide API for file system metadata
type FileController struct {
	files map[string]*File
}

var fileController = &FileController{files: make(map[string]*File)}

func GetFileController() *FileController {
	return fileController
}

func (t *FileController) Create(id string, userId string, acl *acl.ACL, size int64) *File {
	file := newFile(id, userId, acl, size)
	t.files[id] = file
	return file
}

func (t *FileController) Delete(id string) (*File, error) {
	if _, ok := t.files[id]; !ok {
		return nil, fmt.Errorf("failed to fetch file with fileId: %s", id)
	}
	file := t.files[id]
	delete(t.files, id)
	return file, nil
}

func (t *FileController) Get(id string) (*File, error) {
	if _, ok := t.files[id]; !ok {
		return nil, fmt.Errorf("failed to fetch file with fileId: %s", id)
	}
	return t.files[id], nil
}
