package dto

import "github.com/adenapila/market-go/errs"

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	Accounttype string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("To open a new account you need to deposit atleast 5000.00")
	}
	if r.Accounttype != "saving" && r.Accounttype != "checking" {
		return errs.NewValidationError("Account type should be checking or saving")
	}
	return nil
}
