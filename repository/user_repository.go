package repository

import (
	"github.com/notblessy/mini-wallet/model"
	"github.com/notblessy/mini-wallet/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository :nodoc:
func NewUserRepository(d *gorm.DB) model.UserRepository {
	return &userRepository{
		db: d,
	}
}

// Create :nodoc:
func (u *userRepository) Create(user *model.User) error {
	logger := log.WithFields(log.Fields{
		"user": utils.Encode(user),
	})

	err := u.db.Create(user).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
