package bill

import (
	"github.com/shiningrush/demo-service/internal/domain/entity"
)

// Repo bill repository
type Repo interface {
	Get(id int) (*entity.Bill, error)
	Create(account entity.Bill) (*entity.Bill, error)
}
