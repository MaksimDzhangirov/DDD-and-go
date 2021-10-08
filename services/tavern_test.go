package services

import (
	"testing"

	"github.com/MaksimDzhangirov/DDD-and-go/aggregate"
	"github.com/google/uuid"
)

func Test_Tavern(t *testing.T) {
	// Создаём OrderService
	products := init_products(t)

	os, err := NewOrderService(
		WithMongoCustomerRepository("mongodb://localhost:27017"),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Error(err)
	}

	cust, err := aggregate.NewCustomer("Percy")
	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}
	// Выполняем заказ
	err = tavern.Order(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}

}
