package entity

import "github.com/google/uuid"

// Item представляет собой элемент для всех подобластей
type Item struct {
	ID          uuid.UUID
	Name        string
	Description string
}
