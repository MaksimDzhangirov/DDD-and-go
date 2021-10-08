// Агрегаты пакета хранят агрегаты, объединяющие несколько сущностей в один объект
package aggregate

import (
	"errors"
	"github.com/google/uuid"

	"github.com/MaksimDzhangirov/DDD-and-go/entity"
	"github.com/MaksimDzhangirov/DDD-and-go/valueobject"
)

var (
	// ErrInvalidPerson возвращается, когда нельзя создать экземпляр person в фабрике NewCustomer
	ErrInvalidPerson = errors.New("a customer has to have a valid person")
)

// Customer - это агрегат, объединяющий все сущности, необходимые для описания
// клиента в предметной области
type Customer struct {
	// person - это корневая сущность клиента
	// т. е. person.ID - это основной идентификатор для этого агрегата
	person *entity.Person
	// клиент может купить несколько товаров
	products []*entity.Item
	// клиент может осуществлять множество транзакций
	transactions []valueobject.Transaction
}

// NewCustomer - это фабрика для создания нового агрегата Customer
// Она проверит, что передано не пустое имя
func NewCustomer(name string) (Customer, error) {
	// Проверяем, что name не пустая строка
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	// Создаём новый экземпляр person и генерируем ID
	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}
	// Создаём объект Customer и инициализируем все значения,
	// чтобы избежать исключений, связанные со ссылкой на нулевой указатель
	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}

// GetID возвращает ID корневой сущности клиента
func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

// SetID присваивает ID корневой сущности
func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.ID = id
}

// SetName назначает имя для клиента
func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.Name = name
}

// GetName возвращает имя клиента
func (c *Customer) GetName() string {
	return c.person.Name
}