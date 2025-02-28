package acl

type ACL struct {
	user_id string
	read    bool
	write   bool
	delete  bool
	update  bool
}

func NewACL(user_id string) *ACL {
	return &ACL{user_id: user_id}
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
