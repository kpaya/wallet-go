package entity

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Transaction struct {
	ID          string    `json:"id" validate:"required,uuid"`
	AccountFrom *Account  `json:"accountFrom" validate:"-"`
	AccountTo   *Account  `json:"accountTo" validate:"-"`
	Amount      float64   `json:"amount" validate:"required,number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewTransaction(accountFrom *Account, accountTo *Account, amount float64) (*Transaction, error) {
	transaction := &Transaction{
		ID:          uuid.NewString(),
		AccountFrom: accountFrom,
		AccountTo:   accountTo,
		Amount:      amount,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err := transaction.Validate()
	if err != nil {
		return &Transaction{}, err
	}
	err = transaction.Commit()
	if err != nil {
		return &Transaction{}, err
	}
	return transaction, nil
}

func (t *Transaction) Commit() error {
	t.AccountFrom.Debit(t.Amount)
	t.AccountTo.Credit(t.Amount)
	return nil
}

func (t *Transaction) Validate() error {
	var result string
	err := Validate.Struct(t)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			result += fmt.Sprintln(`O campo ` + e.Translate(Trans))
		}
		return errors.New(result)
	}

	if t.AccountFrom.Balance < t.Amount {
		return errors.New("the source account doesn't has enough funds to the transaction")
	}

	if t.Amount <= 0 {
		return errors.New("the amount must be greater than 0")
	}
	return nil
}
