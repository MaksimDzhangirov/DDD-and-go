package tavern

import (
	"github.com/google/uuid"
	"time"
)

// Transaction - информация об оплате
type Transaction struct {
	// все значения заданы заданы в нижнем регистре, поскольку неизменяемы
	amount    int
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}
