// Сущности внутри пакета являются общими для всех подобластей
package entity

import (
	"github.com/google/uuid"
)

// Person - это сущность, которая описывает человека во всех предметных областях
type Person struct {
	// ID - идентификатор сущности, ID общий для всех подобластей
	ID uuid.UUID
	// Name - это имя человека
	Name string
	// Age - это возраст человека
	Age int
}
