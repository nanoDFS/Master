package acl

type Mode int

const (
	Read  Mode = iota // 0
	Write             // 1
)

type ACL struct {
	user_id string
	read    bool
	write   bool
	delete  bool
	update  bool
}

func NewACL(user_id string) *ACL {
	access := &ACL{user_id: user_id}
	access.SetFullAccess()
	return access
}

func (t *ACL) SetFullAccess() {
	t.read, t.write, t.delete, t.update = true, true, true, true
}

func (t *ACL) CanRead() bool {
	return t.read
}
func (t *ACL) CanWrite() bool {
	return t.write
}
func (t *ACL) CanDelete() bool {
	return t.delete
}
