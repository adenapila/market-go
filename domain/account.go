package domain

import (
	"github.com/adenapila/market-go/dto"
	"github.com/adenapila/market-go/errs"
)

type Account struct {
	AccountId   string
	CostumerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{a.AccountId}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
