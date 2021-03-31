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
