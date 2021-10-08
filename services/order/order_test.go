package order

import (
	"github.com/MaksimDzhangirov/tavern/domain/customer"
	"github.com/MaksimDzhangirov/tavern/domain/product"
	"testing"

	"github.com/google/uuid"
)

func init_products(t *testing.T) []product.Product {
	beer, err := product.NewProduct("Beer", "Healthy Beverage", 1.99)
	if err != nil {
		t.Error(err)
	}
	peanuts, err := product.NewProduct("Peanuts", "Healthy Snacks", 0.99)
	if err != nil {
		t.Error(err)
	}
	wine, err := product.NewProduct("Wine", "Healthy Snacks", 0.99)
	if err != nil {
		t.Error(err)
	}
	products := []product.Product{
		beer, peanuts, wine,
	}
	return products
}
func TestOrder_NewOrderService(t *testing.T) {
	// Создаём несколько товаров для вставки в репозиторий
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Error(err)
	}

	// Добавляем клиента
	cust, err := customer.NewCustomer("Percy")
	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}

	// Заказываем одно пиво
	order := []uuid.UUID{
		products[0].GetID(),
	}

	_, err = os.CreateOrder(cust.GetID(), order)

	if err != nil {
		t.Error(err)
	}

}
