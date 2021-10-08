// Пакет Product содержит репозиторий и реализации ProductRepository
package product

import (
	"errors"
	"github.com/MaksimDzhangirov/DDD-and-go/aggregate"
	"github.com/google/uuid"
)

var (
	// ErrProductNotFound возвращается, когда товар не найден
	ErrProductNotFound = errors.New("the product was not found")
	// ErrProductAlreadyExist возвращается при попытке добавить продукт, который уже существует
	ErrProductAlreadyExist = errors.New("the product already exists")
)

// ProductRepository - это интерфейс, которому должен удовлетворять репозиторий, использующий агрегат товара
type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (aggregate.Product, error)
	Add(product aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(id uuid.UUID) error
}
