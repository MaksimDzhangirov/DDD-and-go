// Пакет Service содержит все сервисы, которые объединяя репозитории, создают бизнес-логику
package services

import (
	"context"
	"github.com/MaksimDzhangirov/DDD-and-go/aggregate"
	"github.com/MaksimDzhangirov/DDD-and-go/domain/customer"
	"github.com/MaksimDzhangirov/DDD-and-go/domain/customer/memory"
	"github.com/MaksimDzhangirov/DDD-and-go/domain/customer/mongo"
	"github.com/MaksimDzhangirov/DDD-and-go/domain/product"
	prodmemory "github.com/MaksimDzhangirov/DDD-and-go/domain/product/memory"
	"github.com/google/uuid"
	"log"
)

// OrderConfiguration - это псевдоним для функции, которая принимает указатель на OrderService и модифицирует его
type OrderConfiguration func(os *OrderService) error

// OrderService - это реализация OrderService
type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
}

// NewOrderService принимает на вход произвольное число функций OrderConfiguration и возвращает новый OrderService
// Каждая OrderConfiguration вызывается в том порядке, в котором она передавалась
func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	// Создаём OrderService
	os := &OrderService{}
	// Применяем все переданные OrderConfiguration
	for _, cfg := range cfgs {
		// Передаем сервис в функцию конфигурации
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

// WithCustomerRepository передаёт заданный репозиторий CustomerRepository сервису OrderService
func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	// возвращает функцию, чья сигнатура совпадает с OrderConfiguration
	// нужно возвращать именно её, чтобы родительская функция могла принять все необходимые параметры
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

// WithMemoryCustomerRepository передаёт MemoryCustomerRepository в OrderService
func WithMemoryCustomerRepository() OrderConfiguration {
	// Создаёт репозиторий, хранящий данные в памяти, если нам нужно задать
	// параметры, например, соединения с базой данных, они могут быть добавлены
	// здесь
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func WithMongoCustomerRepository(connectionString string) OrderConfiguration {
	return func(os *OrderService) error {
		// Создаёт MongoDB репозиторий, если нам нужно задать
		// параметры, например, соединения с базой данных, они могут быть добавлены
		// здесь
		cr, err := mongo.New(context.Background(), connectionString)
		if err != nil {
			return err
		}
		os.customers = cr
		return nil
	}
}

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		// Создаём репозиторий, хранящий данные в памяти, если нам нужно задать
		// параметры, например, соединения с базой данных, они могут быть добавлены
		// здесь
		pr := prodmemory.New()

		// добавляем Items в репозиторий
		for _, p := range products {
			err := pr.Add(p)
			if err != nil {
				return err
			}
		}
		os.products = pr
		return nil
	}
}

// CreateOrder объединит несколько репозиториев и позволит создать заказ для клиента
// вернёт общую стоимость всех товаров в заказе
func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {
	// Получаем информацию о клиенте
	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}

	// Получаем информацию о каждом товаре, ой, нам нужен ProductRepository
	var products []aggregate.Product
	var price float64
	for _, id := range productIDs {
		p, err := o.products.GetByID(id)
		if err != nil {
			return 0, err
		}
		products = append(products, p)
		price += p.GetPrice()
	}

	// Все товары есть на складе, теперь мы можем создать заказ
	log.Printf("Customer: %s has ordered %d products", c.GetID(), len(products))

	return price, nil
}
