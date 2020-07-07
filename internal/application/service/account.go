package service

import (
	"github.com/shiningrush/demo-service/internal/domain/entity"
	"github.com/shiningrush/demo-service/internal/domain/repo/account"
	"github.com/shiningrush/demo-service/internal/domain/repo/bill"
	"github.com/shiningrush/demo-service/internal/domain/service"
	"github.com/shiningrush/demo-service/internal/infrastructure/repo"
	"github.com/shiningrush/demo-service/internal/utils/dao"
)

func NewAccountAppSvc(ar account.Repo, br bill.Repo) *AccountService {
	return &AccountService{
		ar: ar,
		br: br,
	}
}

type AccountService struct {
	ar account.Repo
	br bill.Repo
}

func (svc *AccountService) Create(phone string) (*entity.Account, error) {
	acc := entity.NewAccount(phone)
	saveAcc, err := svc.ar.Create(acc)
	if err != nil {
		return nil, err
	}

	return saveAcc, nil
}

func (svc *AccountService) Get(id int) (*entity.Account, error) {
	return svc.Get(id)
}

func (svc *AccountService) Delete(id int) error {
	return svc.ar.Delete(id)
}

func (svc *AccountService) ChangePhone(id int, newPhone string) error {
	dSvc := service.NewAccountSvc(repo.NewAccountRepo(dao.DefaultDB), repo.NewBillRepo(dao.DefaultDB))
	return dSvc.ChangePhone(id, newPhone)
}

func (svc *AccountService) Pay(id int, amount float64, desc string) error {
	acc, err := svc.ar.Get(id)
	if err != nil {
		return err
	}

	bl, err := acc.Cost(amount, desc)
	if err != nil {
		return err
	}

	_, err = svc.br.Create(*bl)
	if err != nil {
		return err
	}
	_, err = svc.ar.Update(*acc)
	if err != nil {
		return err
	}
	return nil
}

func (svc *AccountService) Transfer(fromId, toId int, amount float64, desc string) error {
	dSvc := service.NewAccountSvc(svc.ar, svc.br)
	err := dSvc.Transfer(fromId, toId, amount, desc)
	if err != nil {
		return err
	}

	return nil
}

func (svc *AccountService) Close(id int) error {
	acc, err := svc.ar.Get(id)
	if err != nil {
		return err
	}

	err = acc.Close()
	if err != nil {
		return err
	}

	_, err = svc.ar.Update(*acc)
	if err != nil {
		return err
	}

	return nil
}
