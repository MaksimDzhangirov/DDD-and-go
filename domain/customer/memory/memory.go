// Пакет memory - это реализация репозитория для хранения клиентов в памяти
package memory

import (
	"fmt"
	"github.com/MaksimDzhangirov/DDD-and-go/aggregate"
	"github.com/MaksimDzhangirov/DDD-and-go/domain/customer"
	"github.com/google/uuid"
	"sync"
)

// MemoryRepository удовлетворяет интерфейсу CustomerRepository
type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

// New - это функция-фабрика для создания нового репозитория
func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

// Get ищет клиента по ID
func (mr *MemoryRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	if c, ok := mr.customers[id]; ok {
		return c, nil
	}
	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

// Add добавляет нового клиента в репозиторий
func (mr *MemoryRepository) Add(c aggregate.Customer) error {
	if mr.customers == nil {
		// Дополнительная проверка для случая, если customers не были созданы по какой-то причине.
		// Такого никогда не должно происходить, если использовалась фабрика. Тем не менее никогда не говори "никогда"
		mr.Lock()
		mr.customers = make(map[uuid.UUID]aggregate.Customer)
		mr.Unlock()
	}
	// Убеждаемся, что Customer не был ещё добавлен в репозиторий
	if _, ok := mr.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists: %w", customer.ErrFailedToAddCustomer)
	}
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}

// Update заменяет существующую информацию о клиенте на новую
func (mr *MemoryRepository) Update(c aggregate.Customer) error {
	// Убеждаемся, что такой Customer существует в репозитории
	if _, ok := mr.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exists: %w", customer.ErrUpdateCustomer)
	}
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}