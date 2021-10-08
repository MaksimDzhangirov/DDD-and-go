// Пакет product
// Product - это агрегат, описывающий товар
package product

import (
	"errors"

	"github.com/MaksimDzhangirov/tavern"
	"github.com/google/uuid"
)

var (
	// ErrMissingValues возвращается, когда товар создаётся без названия или описания
	ErrMissingValues = errors.New("missing values")
)

// Product - это агрегат, объединяющий позицию в меню, цену и количество
type Product struct {
	// item - это корневая сущность, которой является Item
	item *tavern.Item
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
		item: &tavern.Item{
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

func (p Product) GetItem() *tavern.Item {
	return p.item
}

func (p Product) GetPrice() float64 {
	return p.price
}