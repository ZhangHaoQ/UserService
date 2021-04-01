package repository

import (
	"UserService/Domain/model"
	"context"
	xerrors "github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserDB(ctx context.Context) *UserRepository {
	return &UserRepository{db: GetDB(ctx)}
}

func (this *UserRepository) Create(M model.UserModel) error {
	return this.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&M).Error; err != nil {
			return xerrors.Wrapf(err, "Create model fail")
		}
		return nil
	})
}

func (this *UserRepository) ExistUserByName(name string) bool {
	User := &model.UserModel{}
	this.db.Select("id").Where("user_name = ?", name).First(User)
	if User.ID > 0 {
		return true
	}
	return false
}

func (this *UserRepository) GetUserByMobile(Mobile string) (*model.UserModel, error) {
	User := &model.UserModel{}
	return User, this.db.Select("password").Where("mobile = ?", Mobile).First(User).Error
}

func (this *UserRepository) GetUserByName(UserName string) (*model.UserModel, error) {
	User := &model.UserModel{}
	return User, this.db.Where("user_name = ?", UserName).First(User).Error
}
