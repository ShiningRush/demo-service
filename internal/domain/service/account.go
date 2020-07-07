package service

import (
	"fmt"
	"github.com/shiningrush/demo-service/internal/domain/entity"
	"github.com/shiningrush/demo-service/internal/domain/repo/account"
	"github.com/shiningrush/demo-service/internal/domain/repo/bill"
)

type Account interface {
	NewAccount(phone string) (*entity.Account, error)
	ChangePhone(accId int, newPhone string) error
	Transfer(fromAccId, toAccId int, amount float64, desc string) error
}

type accountService struct {
	accRepo  account.Repo
	billRepo bill.Repo
}

func NewAccountSvc(accRepo account.Repo, billRepo bill.Repo) Account {
	return &accountService{
		accRepo:  accRepo,
		billRepo: billRepo,
	}
}

func (svc *accountService) NewAccount(phone string) (*entity.Account, error) {
	if err := svc.isPhoneAlreadyInUsed(phone); err != nil {
		return nil, err
	}

	newAcc := entity.NewAccount(phone)
	return &newAcc, nil

}

func (svc *accountService) ChangePhone(accId int, newPhone string) error {
	if err := svc.isPhoneAlreadyInUsed(newPhone); err != nil {
		return err
	}

	targetAcc, err := svc.accRepo.Get(accId)
	if err != nil {
		return err
	}
	if targetAcc == nil {
		return fmt.Errorf("account id: %d not found", accId)
	}

	targetAcc.Phone = newPhone
	_, err = svc.accRepo.Update(*targetAcc)
	if err != nil {
		return err
	}
	return nil
}

func (svc *accountService) isPhoneAlreadyInUsed(newPhone string) error {
	acc, err := svc.accRepo.GetWithPhone(newPhone)
	if err != nil {
		return err
	}

	if acc != nil {
		return fmt.Errorf("phone alreay be used")
	}
	return nil
}

func (svc *accountService) Transfer(fromAccId, toAccId int, amount float64, desc string) error {
	fromAcc, err := svc.getAccountAndCheckIt(fromAccId)
	if err != nil {
		return err
	}
	toAcc, err := svc.getAccountAndCheckIt(toAccId)
	if err != nil {
		return err
	}

	costBill, err := fromAcc.Cost(amount, desc)
	if err != nil {
		return err
	}
	earnBill, err := toAcc.Earn(amount, desc)
	if err != nil {
		return err
	}

	if _, err = svc.accRepo.Update(*fromAcc); err != nil {
		return err
	}
	if _, err = svc.accRepo.Update(*toAcc); err != nil {
		return err
	}

	if _, err = svc.billRepo.Create(*costBill); err != nil {
		return err
	}
	if _, err = svc.billRepo.Create(*earnBill); err != nil {
		return err
	}

	return nil
}

func (svc *accountService) getAccountAndCheckIt(accId int) (*entity.Account, error) {
	acc, err := svc.accRepo.Get(accId)
	if err != nil {
		return nil, err
	}
	if acc == nil || !acc.IsNormal() {
		return nil, fmt.Errorf("account: %d is not normal or not existed", accId)
	}

	return acc, nil
}
