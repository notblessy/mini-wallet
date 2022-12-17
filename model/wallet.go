package model

import "time"

type WalletStatus int32

const (
	WalletStatus_Enabled  WalletStatus = 0
	WalletStatus_Disabled WalletStatus = 1
)

// Enum value maps for WalletStatus.
var (
	WalletStatus_label = map[WalletStatus]string{
		0: "enabled",
		1: "disabled",
	}
)

// WalletRepository :nodoc:
type WalletRepository interface {
	Create(wallet *Wallet) error
	FindByID(walletID *string) (*Wallet, error)
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
