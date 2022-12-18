package model

import "time"

type WithdrawalStatus int32

const (
	WithdrawalStatus_Success WithdrawalStatus = 1
	WithdrawalStatus_Pending WithdrawalStatus = 2
)

// Enum value maps for WithdrawalStatus.
var (
	WithdrawalStatus_label = map[WithdrawalStatus]string{
		1: "success",
		2: "pending",
	}
)

// WithdrawalRepository :nodoc:
type WithdrawalRepository interface {
	Create(deposit *Withdrawal, wallet *Wallet) error
	FindByReference(refID *string) (*Withdrawal, error)
}

type Withdrawal struct {
	ID          string           `json:"id" gorm:"primary_key"`
	WithdrawnBy string           `json:"withdrawn_by"`
	Status      WithdrawalStatus `json:"status"`
	Amount      int64            `json:"amount"`
	ReferenceID string           `json:"reference_id"`
	WithdrawnAt *time.Time       `json:"withdrawn_at"`
}

type WithdrawalResponse struct {
	ID          string     `json:"id" gorm:"primary_key"`
	WithdrawnBy string     `json:"withdrawn_by"`
	Status      string     `json:"status"`
	Amount      int64      `json:"amount"`
	ReferenceID string     `json:"reference_id"`
	WithdrawnAt *time.Time `json:"withdrawn_at"`
}

func (w *Withdrawal) NewResponse() *WithdrawalResponse {
	if w != nil {
		return &WithdrawalResponse{
			ID:          w.ID,
			WithdrawnBy: w.WithdrawnBy,
			Status:      WithdrawalStatus_label[w.Status],
			Amount:      w.Amount,
			WithdrawnAt: w.WithdrawnAt,
			ReferenceID: w.ReferenceID,
		}
	}

	return nil
}
