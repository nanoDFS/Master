package metadata

import (
	"sync"

	"fmt"

	"github.com/nanoDFS/Master/controller/acl"
	cs "github.com/nanoDFS/Master/controller/metadata/chunkserver"
	lb "github.com/nanoDFS/Master/controller/metadata/chunkserver/loadbalancer"
	repl "github.com/nanoDFS/Master/controller/metadata/chunkserver/replicator"
	dm "github.com/nanoDFS/Master/utils/datamodel"
)

type File struct {
	id      string
	ownerId string
	acl     *acl.ACL
	Size    *dm.ConcurrentValue[int64]
	chunks  *dm.ConcurrentList[*repl.Replicas]
}

func newFile(id string, userId string, acl *acl.ACL, size int64) *File {

	var chunks *dm.ConcurrentList[*repl.Replicas] = dm.NewConcurrentList[*repl.Replicas]()
	chunker := cs.NewChunker(lb.ConsistentHashing{})
	for _, k := range chunker.Generate(id, size) {
		chunks.Append(repl.NewReplicas(k, nil, nil))
	}
	return &File{
		id:      id,
		ownerId: userId,
		acl:     acl,
		Size:    dm.NewConcurrentValue(size),
		chunks:  chunks,
	}
}

func (t *File) GetACL() *acl.ACL {
	return t.acl
}
func (t *File) GetOwnerID() string {
	return t.ownerId
}
func (t *File) GetID() string {
	return t.id
}

func (t *File) GetChunkServers() []string {
	var chunk_servers []string
	for i := range t.chunks.Size() {
		res, _ := t.chunks.Get(i)
		chunk_servers = append(chunk_servers, res.Primary.Get().Addr.String())
	}
	return chunk_servers
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
