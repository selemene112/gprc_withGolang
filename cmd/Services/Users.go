package services

import (
	"context"
	"fmt"
	UsersPb "gprc_protocolbuffer/pkg"

	"gorm.io/gorm"
)

type UserServices struct {
	UsersPb.UnimplementedProductServiceServer
	DB *gorm.DB
}

// func GetAllUsers() *UsersPb.Users {
// 	return &UsersPb.Users{}
// }

// func (p *UserServices) GetAll(context.Context, *UsersPb.Empty) (*UsersPb.Users, error) {
// 	GetUser := []*UsersPb.Users{}
// 	p.DB.Find(&GetUser)
// 	return &UsersPb.Users{}, nil
// }

func (p *UserServices) GetAll(ctx context.Context, req *UsersPb.Empty) (*UsersPb.UsersList, error) {
	GetAllUser := []*UsersPb.Users{}

	if err := p.DB.Find(&GetAllUser).Error; err != nil {
		return nil, err
	}

	fmt.Println(GetAllUser)

	// Anda dapat menyalin langsung data GetAllUser ke UsersList
	result := &UsersPb.UsersList{
		Data: GetAllUser,
	}

	return result, nil
}
