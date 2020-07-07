package account

import (
	"github.com/shiningrush/demo-service/internal/domain/entity"
)

// Repo Account repository
type Repo interface {
	Get(id int) (*entity.Account, error)
	GetWithPhone(phone string) (*entity.Account, error)
	Create(account entity.Account) (*entity.Account, error)
	Update(account entity.Account) (*entity.Account, error)
	Delete(id int) error
}
