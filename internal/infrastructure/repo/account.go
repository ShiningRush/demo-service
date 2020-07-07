package repo

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/shiningrush/demo-service/internal/domain/entity"
	"github.com/shiningrush/demo-service/internal/domain/repo/account"
	"github.com/shiningrush/demo-service/internal/utils/dao"
	"time"
)

func NewAccountRepo(db *gorm.DB) account.Repo {
	if db == nil {
		db = dao.DefaultDB
	}
	return &accountRepo{
		db: db,
	}
}

type accountRepo struct {
	db *gorm.DB
}

func (r *accountRepo) Get(id int) (*entity.Account, error) {
	ret := entity.Account{}
	if err := r.db.First(&ret, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("execute query failed: %w", err)
	}

	return &ret, nil
}

func (r *accountRepo) GetWithPhone(phone string) (*entity.Account, error) {
	ret := entity.Account{}
	if err := r.db.Find(&ret, "phone = ?", phone).Error; err != nil {
		return nil, fmt.Errorf("execute query failed: %w", err)
	}

	return &ret, nil
}

func (r *accountRepo) Create(account entity.Account) (*entity.Account, error) {
	account.UpdatedAt = time.Now()
	if err := r.db.Create(&account).Error; err != nil {
		return nil, fmt.Errorf("execute create failed: %w", err)
	}

	return &account, nil
}

func (r *accountRepo) Update(account entity.Account) (*entity.Account, error) {
	account.UpdatedAt = time.Now()

	if err := r.db.Model(&entity.Account{}).Update(&account).Error; err != nil {
		return nil, fmt.Errorf("execute update failed: %w", err)
	}

	return &account, nil
}

func (r *accountRepo) Delete(id int) error {
	if err := r.db.Delete(&entity.Account{}, "id = ?", id).Error; err != nil {
		return fmt.Errorf("execute delete failed: %w", err)
	}

	return nil
}
