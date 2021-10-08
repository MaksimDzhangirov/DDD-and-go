// Пакет Product содержит репозиторий и реализации ProductRepository
package product

import (
	"errors"
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
	GetAll() ([]Product, error)
	GetByID(id uuid.UUID) (Product, error)
	Add(product Product) error
	Update(product Product) error
	Delete(id uuid.UUID) error
}
