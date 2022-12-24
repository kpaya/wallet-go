package entity

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Client struct {
	ID        string `json:"id" validate:"uuid"`
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"email,required"`
	Accounts  []*Account
	CreatedAt time.Time `json:"created_at" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" validate:"required"`
}

func NewClient(name string, email string) (*Client, error) {
	client := &Client{
		ID:        uuid.NewString(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := client.Validate()
	if err != nil {
		return nil, err
	}
	return &Client{}, err
}

func (c *Client) AddAccount(account *Account) error {
	if account.Client.ID != c.ID {
		return errors.New("this account doesn't belongs you")
	}
	c.Accounts = append(c.Accounts, account)
	return nil
}

func (c *Client) Validate() error {
	var result string
	err := Validate.Struct(c)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			result += fmt.Sprintln(`O campo ` + e.Translate(Trans))
		}
		return errors.New(result)
	}
	return nil
}
