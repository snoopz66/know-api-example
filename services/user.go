package services

import (
	uuid "github.com/satori/go.uuid"
	"github.com/snoopz66/know-api-example/models"
	"github.com/snoopz66/know-api-example/repositories/userRepo"
	"go.uber.org/zap"
)

type User struct {
	UserDbRepo *userRepo.DynDB
	Log        *zap.Logger
}

func New(userRepo *userRepo.DynDB, logger *zap.Logger) *User {
	return &User{
		UserDbRepo: userRepo,
		Log:        logger,
	}
}

func (u *User) CreateUser(user *models.User) (*models.User, error) {
	user.UUID = uuid.NewV4().String()
	err := u.UserDbRepo.CreateUser(user)
	if err != nil {
		u.Log.Error("failed to create user in db")
		return nil, err
	}
	u.Log.Info("user created")
	return user, nil
}

func (u *User) GetUser(uuid string) (*models.User, error) {
	user, err := u.UserDbRepo.GetUser(uuid)
	if err != nil {
		u.Log.Error("could not find user")
		return nil, err
	}
	u.Log.Info("found user " + user.UUID)
	return user, nil
}
