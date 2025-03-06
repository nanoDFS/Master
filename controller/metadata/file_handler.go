package metadata

import (
	"github.com/nanoDFS/Master/controller/auth/acl"
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
		chunk_servers = append(chunk_servers, res.Primary.Get().StreamingAddr.String())
	}
	return chunk_servers
}

// FileController is a singleton class, provides API for file system metadata
type FileController struct {
	files *dm.ConcurrentMap[string, *File]
}

var fileController = &FileController{files: dm.NewConcurrentMap[string, *File]()}

func GetFileController() *FileController {
	return fileController
}

func (t *FileController) Create(id string, userId string, acl *acl.ACL, size int64) *File {
	file := newFile(id, userId, acl, size)
	t.files.Set(id, file)
	return file
}

func (t *FileController) Delete(id string) (*File, error) {
	return t.files.Delete(id)
}

func (t *FileController) Get(id string) (*File, error) {
	return t.files.Get(id)
}
