package repository

import (
	"github.com/notblessy/mini-wallet/model"
	"github.com/notblessy/mini-wallet/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type withdrawalRepository struct {
	db         *gorm.DB
	walletRepo model.WalletRepository
}

// NewWithdrawalRepository :nodoc:
func NewWithdrawalRepository(d *gorm.DB, w model.WalletRepository) model.WithdrawalRepository {
	return &withdrawalRepository{
		db:         d,
		walletRepo: w,
	}
}

// Create :nodoc:
func (u *withdrawalRepository) Create(withdrawal *model.Withdrawal, wallet *model.Wallet) error {
	logger := log.WithFields(log.Fields{
		"withdrawal": utils.Encode(withdrawal),
	})

	tx := u.db.Begin()
	err := tx.Create(withdrawal).Error
	if err != nil {
		logger.Error(err)
		tx.Rollback()
		return err
	}

	err = u.walletRepo.Save(wallet, tx)
	if err != nil {
		logger.Error(err)
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// FindByReference :nodoc:
func (w *withdrawalRepository) FindByReference(refID *string) (withdrawal *model.Withdrawal, err error) {
	logger := log.WithFields(log.Fields{
		"referenceID": refID,
	})

	err = w.db.Where("reference_id = ?", refID).First(&withdrawal).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error(err)
		return nil, err
	}

	return nil, nil
}
