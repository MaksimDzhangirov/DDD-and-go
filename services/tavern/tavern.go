package tavern

import (
	"github.com/MaksimDzhangirov/tavern/services/order"
	"github.com/google/uuid"
	"log"
)

// TavernConfiguration - это псевдоним для функции, которая принимает указатель и модифицирует Tavern
type TavernConfiguration func(os *Tavern) error

type Tavern struct {
	// OrderService используется для работы с заказами
	OrderService *order.OrderService
	// BillingService используется для работы со счетами
	// Вы можете реализовать его сами
	BillingService interface{}
}

// NewTavern принимает на вход произвольное число TavernConfiguration и создаёт Tavern
func NewTavern(cfgs ...TavernConfiguration) (*Tavern, error) {
	// Создаём таверну
	t := &Tavern{}
	// Применяем все переданные TavernConfiguration
	for _, cfg := range cfgs {
		// Передаем сервис в функцию конфигурации
		err := cfg(t)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

// WithOrderService передаёт заданный OrderService в Tavern
func WithOrderService(os *order.OrderService) TavernConfiguration {
	// возвращает функцию, чья сигнатура совпадает с TavernConfiguration
	return func(t *Tavern) error {
		t.OrderService = os
		return nil
	}
}

// Order осуществляет заказ для клиента
func (t *Tavern) Order(customer uuid.UUID, products []uuid.UUID) error {
	price, err := t.OrderService.CreateOrder(customer, products)
	if err != nil {
		return err
	}
	log.Printf("Bill the Customer: %0.0f", price)

	// Выставить счёт клиенту
	// err = t.BillingService(customer, price)
	return nil
}