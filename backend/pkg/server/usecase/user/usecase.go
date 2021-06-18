package model

import (
	"fmt"

	"github.com/google/uuid"
	model "github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/model/user"
	ur "github.com/littletake/supporterz_hackathon_2021/pkg/server/domain/repository/user"
)

type UserUsecase interface {
	GetUserByUserID(userID string) (*model.User, error)
	RegisterUser(userName string, password string) (string, error)
	// UpdateUser(userID string, userName string) (*model.User, error)
}

type userUsecase struct {
	userRepo   ur.UserRepo
	createUUID func() (uuid.UUID, error)
}

func NewUserUsecase(ur ur.UserRepo, f func() (uuid.UUID, error)) UserUsecase {
	return &userUsecase{
		userRepo:   ur,
		createUUID: f,
	}
}

func (uu *userUsecase) GetUserByUserID(userID string) (*model.User, error) {
	user, err := uu.userRepo.SelectUserByUserID(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		err := fmt.Errorf("user not found. userID=%s", userID)
		return nil, err
	}
	return user, nil
}

func (uu *userUsecase) RegisterUser(userName string, password string) (string, error) {
	// userID
	userID, err := uu.createUUID()
	if err != nil {
		return "", err
	}
	// token
	token, err := uu.createUUID()
	if err != nil {
		return "", err
	}

	user := &model.User{
		UserID:    userID.String(),
		UserName:  userName,
		Password:  password,
		AuthToken: token.String(),
	}
	// 登録
	if err := uu.userRepo.InsertUser(user); err != nil {
		return "", err
	}
	return user.AuthToken, nil
}
