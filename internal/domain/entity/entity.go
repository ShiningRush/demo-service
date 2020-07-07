package entity

import (
	"fmt"
	"time"
)

type Account struct {
	ID        int
	Phone     string
	Status    AccountStatus
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (acc *Account) TableName() string {
	return "accounts"
}

type AccountStatus string

const (
	AccountStatusNormal AccountStatus = "normal"
	AccountStatusClosed AccountStatus = "closed"
)

func NewAccount(phone string) Account {
	return Account{
		Phone:   phone,
		Status:  AccountStatusNormal,
		Balance: 0,
	}
}

func (acc *Account) IsNormal() bool {
	return acc.Status == AccountStatusNormal
}

func (acc *Account) Close() error {
	if acc.Status != AccountStatusNormal {
		return fmt.Errorf("only normal account can close")
	}

	acc.Status = AccountStatusClosed
	return nil
}

func (acc *Account) Cost(amount float64, desc string) (*Bill, error) {
	if !acc.IsNormal() {
		return nil, fmt.Errorf("only normal account can cost")
	}
	if amount > acc.Balance {
		return nil, fmt.Errorf("balance not enough")
	}

	acc.Balance -= amount
	return &Bill{
		Desc:      desc,
		AccountId: acc.ID,
		Amount:    amount,
		CreatedAt: time.Time{},
	}, nil
}

func (acc *Account) Earn(amount float64, desc string) (*Bill, error) {
	if !acc.IsNormal() {
		return nil, fmt.Errorf("only normal account can cost")
	}

	acc.Balance += amount
	return &Bill{
		Type:      BillTypeEarn,
		Desc:      desc,
		AccountId: acc.ID,
		Amount:    amount,
		CreatedAt: time.Time{},
	}, nil
}

type Bill struct {
	ID        int
	Desc      string
	Type      BillType
	AccountId int
	Amount    float64
	CreatedAt time.Time
}

type BillType string

const (
	BillTypeCost = "cost"
	BillTypeEarn = "earn"
)
