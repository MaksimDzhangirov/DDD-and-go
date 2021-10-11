// Пакет main запускает Таверну и осуществляет заказ
package main

import (
	"github.com/MaksimDzhangirov/tavern/domain/product"
	"github.com/MaksimDzhangirov/tavern/services/order"
	servicetavern "github.com/MaksimDzhangirov/tavern/services/tavern"
	"github.com/google/uuid"
)

func main() {

	products := productInventory()
	// Создаём OrderService, чтобы использовать в Таверне
	os, err := order.NewOrderService(
		order.WithMongoCustomerRepository("mongodb://localhost:27017"),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		panic(err)
	}
	// Создаём сервис Tavern
	tavern, err := servicetavern.NewTavern(
		servicetavern.WithOrderService(os),
	)
	if err != nil {
		panic(err)
	}

	uid, err := os.AddCustomer("Percy")
	if err != nil {
		panic(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}
	// Выполняем заказ
	err = tavern.Order(uid, order)
	if err != nil {
		panic(err)
	}
}

func productInventory() []product.Product {
	beer, err := product.NewProduct("Beer", "Healthy Beverage", 1.99)
	if err != nil {
		panic(err)
	}
	peanuts, err := product.NewProduct("Peanuts", "Healthy Snacks", 0.99)
	if err != nil {
		panic(err)
	}
	wine, err := product.NewProduct("Wine", "Healthy Snacks", 0.99)
	if err != nil {
		panic(err)
	}
	products := []product.Product{
		beer, peanuts, wine,
	}
	return products
}
