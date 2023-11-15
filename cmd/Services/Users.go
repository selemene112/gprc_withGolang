package services

import (
	"context"
	UsersPb "gprc_protocolbuffer/pkg"
)

type UserServices struct {
	UsersPb.UnimplementedProductServiceServer
}

// func GetAllUsers() *UsersPb.Users {
// 	return &UsersPb.Users{}
// }

func (p *UserServices) GetAll(context.Context, *UsersPb.Empty) (*UsersPb.Users, error) {
	Users := &UsersPb.Users{
		Id:          19,
		Username:    "test",
		Password:    "test",
		Email:       "test",
		Name:        "testing",
		PhotoProfil: "test",
	}

	return Users, nil
}
