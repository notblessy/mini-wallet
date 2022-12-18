package repository

import (
	"github.com/notblessy/mini-wallet/model"
	"github.com/notblessy/mini-wallet/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type depositRepository struct {
	db         *gorm.DB
	walletRepo model.WalletRepository
}

// NewDepositRepository :nodoc:
func NewDepositRepository(d *gorm.DB, w model.WalletRepository) model.DepositRepository {
	return &depositRepository{
		db:         d,
		walletRepo: w,
	}
}

// Create :nodoc:
func (u *depositRepository) Create(deposit *model.Deposit, wallet *model.Wallet) error {
	logger := log.WithFields(log.Fields{
		"deposit": utils.Encode(deposit),
	})

	tx := u.db.Begin()
	err := tx.Create(deposit).Error
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
func (u *depositRepository) FindByReference(refID *string) (deposit *model.Deposit, err error) {
	logger := log.WithFields(log.Fields{
		"referenceID": refID,
	})

	err = u.db.Where("reference_id = ?", refID).First(&deposit).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error(err)
		return nil, err
	}

	return nil, nil
}
