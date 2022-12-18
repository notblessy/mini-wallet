package model

import (
	"time"

	"gorm.io/gorm"
)

type WalletStatus int32

const (
	WalletStatus_Enabled  WalletStatus = 1
	WalletStatus_Disabled WalletStatus = 2
)

// Enum value maps for WalletStatus.
var (
	WalletStatus_label = map[WalletStatus]string{
		1: "enabled",
		2: "disabled",
	}
)

// WalletRepository :nodoc:
type WalletRepository interface {
	Create(wallet *Wallet) error
	FindByOwner(walletID *string) (*Wallet, error)
	Save(wallet *Wallet, tx *gorm.DB) error
	ToggleStatus(wallet *Wallet) error
}

type Wallet struct {
	ID         string       `json:"id" gorm:"primary_key"`
	OwnedBy    string       `json:"owned_by"`
	Status     WalletStatus `json:"status"`
	Balance    int64        `json:"balance"`
	EnabledAt  *time.Time   `json:"enabled_at"`
	DisabledAt *time.Time   `json:"disabled_at"`
}

type WalletResponse struct {
	ID         string     `json:"id" gorm:"primary_key"`
	OwnedBy    string     `json:"owned_by"`
	Status     string     `json:"status"`
	Balance    int64      `json:"balance"`
	EnabledAt  *time.Time `json:"enabled_at,omitempty"`
	DisabledAt *time.Time `json:"disabled_at,omitempty"`
}

type WalletStatusRequest struct {
	IsDisabled bool `json:"is_disabled"`
}

func (w *Wallet) NewResponse() *WalletResponse {
	if w != nil {
		return &WalletResponse{
			ID:         w.ID,
			OwnedBy:    w.OwnedBy,
			Status:     WalletStatus_label[w.Status],
			Balance:    w.Balance,
			EnabledAt:  w.EnabledAt,
			DisabledAt: w.DisabledAt,
		}
	}

	return nil
}
