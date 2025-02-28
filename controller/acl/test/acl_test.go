package acl_test

import (
	"testing"

	"github.com/nanoDFS/Master/controller/acl"
)

func TestJWTGenerate(t *testing.T) {
	_, err := acl.NewJWT().Generate(&acl.Claims{
		UserId: "test",
		Access: *acl.NewACL("test"),
	})

	if err != nil {
		t.Errorf("failed to generate JWT, %v", err)
	}
}
