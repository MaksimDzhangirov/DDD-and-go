// Пакет Customer содержит всю логику, связанную с предметной областью "Клиент"
package customer

import (
	"errors"
	"github.com/MaksimDzhangirov/DDD-and-go/aggregate"
	"github.com/google/uuid"
)

var (
	// ErrCustomerNotFound возвращается если клиент не был найден.
	ErrCustomerNotFound = errors.New("the customer was not found in the repository")
	// ErrFailedToAddCustomer возвращается, когда клиента нельзя добавить в хранилище.
	ErrFailedToAddCustomer = errors.New("failed to add the customer to the repository")
	// ErrUpdateCustomer возвращается, когда клиента нельзя обновить в хранилище.
	ErrUpdateCustomer = errors.New("failed to update the customer in the repository")
)

// CustomerRepository - это интерфейс, определяющий правила, которым должен
// удовлетворять репозиторий для хранения клиентов
type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
