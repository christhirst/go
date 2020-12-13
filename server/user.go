package server

import (
	"context"
	protos "go-postgres/proto/user"

	"github.com/hashicorp/go-hclog"
)

type User struct {
	log hclog.Logger
}

func NewUser(l hclog.Logger) *User {
	return &User{l}
}

func (c *User) GetUser(ctx context.Context, rr *protos.UserRequest) (*protos.DataResponse, error) {
	c.log.Info("Handle GetUser", "user", rr.GetUserName(), "pw", rr.GetUserPW())
	return &protos.DataResponse{User: "22", Id: 1}, nil
}
