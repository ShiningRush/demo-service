package repo

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/shiningrush/demo-service/internal/domain/entity"
	"github.com/shiningrush/demo-service/internal/domain/repo/bill"
	"github.com/shiningrush/demo-service/internal/utils/dao"
	"time"
)

func NewBillRepo(db *gorm.DB) bill.Repo {
	if db == nil {
		db = dao.DefaultDB
	}
	return &billRepo{
		db: db,
	}
}

type billRepo struct {
	db *gorm.DB
}

func (r *billRepo) Get(id int) (*entity.Bill, error) {
	ret := entity.Bill{}
	if err := r.db.Find(&ret, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("execute query failed: %w", err)
	}

	return &ret, nil
}

func (r *billRepo) Create(bill entity.Bill) (*entity.Bill, error) {
	bill.CreatedAt = time.Now()

	if err := dao.DefaultDB.Create(&bill).Error; err != nil {
		return nil, fmt.Errorf("execute create failed: %w", err)
	}

	return &bill, nil
}
