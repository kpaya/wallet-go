package entity

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Account struct {
	ID        string    `json:"id" validate:"required"`
	Client    *Client   `json:"client" validate:"-"`
	Balance   float64   `json:"balance" validate:"numeric"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" validate:"required"`
}

func NewAccount(client *Client) (*Account, error) {
	account := &Account{
		ID:        uuid.NewString(),
		Client:    client,
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := account.Validate()
	if err != nil {
		return &Account{}, err
	}
	return account, nil
}

func (a *Account) Credit(value float64) error {
	if value <= 0 {
		return errors.New("the value must be greater than 0")
	}
	a.Balance += value
	a.UpdatedAt = time.Now()
	return nil
}

func (a *Account) Debit(value float64) error {
	if value > a.Balance {
		return errors.New("you do not have enough funds")
	}
	if value <= 0 {
		return errors.New("provide to us a valid value to withdraw")
	}
	a.Balance -= value
	a.UpdatedAt = time.Now()
	return nil
}

func (a *Account) Validate() error {
	var result string
	err := Validate.Struct(a)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			result += fmt.Sprintln(`O campo ` + e.Translate(Trans))
		}
		return errors.New(result)
	}
	return nil
}
