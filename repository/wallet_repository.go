package repository

import (
	"github.com/notblessy/mini-wallet/model"
	"github.com/notblessy/mini-wallet/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type walletRepository struct {
	db *gorm.DB
}

// NewWalletRepository :nodoc:
func NewWalletRepository(d *gorm.DB) model.WalletRepository {
	return &walletRepository{
		db: d,
	}
}

// Create :nodoc:
func (w *walletRepository) Create(wallet *model.Wallet) error {
	logger := log.WithFields(log.Fields{
		"wallet": utils.Encode(wallet),
	})

	err := w.db.Create(wallet).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

// FindByID :nodoc:
func (w *walletRepository) FindByOwner(id *string) (wallet *model.Wallet, err error) {
	logger := log.WithFields(log.Fields{
		"walletID": id,
	})

	err = w.db.Where("owned_by = ?", id).First(&wallet).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error(err)
		return nil, err
	}

	return wallet, err
}

// Save :nodoc:
func (w *walletRepository) Save(wallet *model.Wallet, tx *gorm.DB) error {
	logger := log.WithFields(log.Fields{
		"wallet": utils.Encode(wallet),
	})

	err := tx.Save(wallet).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
