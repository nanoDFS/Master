package metadata

import (
	"sync"

	"fmt"

	"github.com/nanoDFS/Master/controller/acl"
	loadbalancing "github.com/nanoDFS/Master/controller/metadata/load_balancing"
	"github.com/nanoDFS/Master/utils/config"
)

type File struct {
	id      string
	ownerId string
	acl     *acl.ACL
	Size    int64
	Chunks  [][]*ChunkServer

	mu sync.RWMutex
}

func newFile(id string, userId string, acl *acl.ACL, size int64) *File {

	var chunks [][]*ChunkServer
	for _, k := range generateChunks(id, size) {
		chunks = append(chunks, []*ChunkServer{k})
	}
	return &File{
		id:      id,
		ownerId: userId,
		acl:     acl,
		Size:    size,
		Chunks:  chunks,
		mu:      sync.RWMutex{},
	}
}

func (t *File) GetACL() *acl.ACL {
	t.mu.RLock()
	acl_ := t.acl
	t.mu.RUnlock()
	return acl_
}
func (t *File) GetOwnerID() string {
	t.mu.RLock()
	id := t.ownerId
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

// FileController is a singleton class, provides API for file system metadata
type FileController struct {
	mu    sync.RWMutex
	files map[string]*File
}

var fileController = &FileController{files: make(map[string]*File), mu: sync.RWMutex{}}

func GetFileController() *FileController {
	return fileController
}

func (t *FileController) Create(id string, userId string, acl *acl.ACL, size int64) *File {
	file := newFile(id, userId, acl, size)
	t.mu.Lock()
	t.files[id] = file
	t.mu.Unlock()
	return file
}

func (t *FileController) Delete(id string) (*File, error) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if _, ok := t.files[id]; !ok {
		return nil, fmt.Errorf("failed to fetch file with fileId: %s", id)
	}
	file := t.files[id]
	delete(t.files, id)
	return file, nil
}

func (t *FileController) Get(id string) (*File, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	if _, ok := t.files[id]; !ok {
		return nil, fmt.Errorf("failed to fetch file with fileId: %s", id)
	}
	res := t.files[id]
	return res, nil
}
