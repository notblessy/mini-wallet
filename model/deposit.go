package model

import "time"

type DepositStatus int32

const (
	DepositStatus_Success DepositStatus = 1
	DepositStatus_Pending DepositStatus = 2
)

// Enum value maps for DepositStatus.
var (
	DepositStatus_label = map[DepositStatus]string{
		1: "success",
		2: "pending",
	}
)

// DepositRepository :nodoc:
type DepositRepository interface {
	Create(deposit *Deposit, wallet *Wallet) error
	FindByReference(refID *string) (*Deposit, error)
}

type Deposit struct {
	ID          string        `json:"id" gorm:"primary_key"`
	DepositedBy string        `json:"deposited_by"`
	Status      DepositStatus `json:"status"`
	Amount      int64         `json:"amount"`
	ReferenceID string        `json:"reference_id"`
	DepositedAt *time.Time    `json:"deposit_at"`
}

type DepositResponse struct {
	ID          string     `json:"id" gorm:"primary_key"`
	DepositedBy string     `json:"deposited_by"`
	Status      string     `json:"status"`
	Amount      int64      `json:"amount"`
	ReferenceID string     `json:"reference_id"`
	DepositedAt *time.Time `json:"deposit_at"`
}

func (d *Deposit) NewResponse() *DepositResponse {
	if d != nil {
		return &DepositResponse{
			ID:          d.ID,
			DepositedBy: d.DepositedBy,
			Status:      DepositStatus_label[d.Status],
			Amount:      d.Amount,
			DepositedAt: d.DepositedAt,
			ReferenceID: d.ReferenceID,
		}
	}

	return nil
}
