package auth

import (
	"fmt"

	"github.com/nanoDFS/Master/controller/auth/acl"
	md "github.com/nanoDFS/Master/controller/metadata"
)

type Auth struct {
}

func NewAuth() Auth {
	return Auth{}
}

func (t Auth) authorize(userId string, fileId string, access acl.ACL, size int64) []byte {
	token, _ := acl.NewJWT().Generate(&acl.Claims{UserId: userId, FileId: fileId, Access: access, Size: size})
	return token
}

func (t Auth) AuthorizeRead(userId string, file md.File, access acl.ACL, size int64) ([]byte, error) {
	if !access.CanRead() || file.GetOwnerID() != userId {
		return nil, fmt.Errorf("do not have access to read")
	}
	return t.authorize(userId, file.GetID(), access, size), nil
}

func (t Auth) AuthorizeDelete(userId string, file md.File, access acl.ACL, size int64) ([]byte, error) {
	if !access.CanDelete() || file.GetOwnerID() != userId {
		return nil, fmt.Errorf("do not have access to read")
	}
	return t.authorize(userId, file.GetID(), access, size), nil
}
