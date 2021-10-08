// Пакет aggregate
// Файл: product.go
// Product - это агрегат, описывающий товар
package aggregate

import (
	"errors"
	"github.com/MaksimDzhangirov/DDD-and-go/entity"
	"github.com/google/uuid"
)

var (
	// ErrMissingValues возвращается, когда товар создаётся без названия или описания
	ErrMissingValues = errors.New("missing values")
)

// Product - это агрегат, объединяющий позицию в меню, цену и количество
type Product struct {
	// item - это корневая сущность, которой является Item
	item *entity.Item
	price float64
	// quantity - количество товара на складе
	quantity int
}

// NewProduct создаст новый продукт
// вернет ошибку, если название или описание будет пустым
func NewProduct(name, description string, price float64) (Product, error) {
	if name == "" || description == "" {
		return Product{}, ErrMissingValues
	}

	return Product{
		item: &entity.Item{
			ID: uuid.New(),
			Name: name,
			Description: description,
		},
		price: price,
		quantity: 0,
	}, nil
}

func (p Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p Product) GetItem() *entity.Item {
	return p.item
}

func (p Product) GetPrice() float64 {
	return p.price
}