package expense

import (
	"time"
	"github.com/google/uuid"
)

type Expense struct {
	ID uuid.UUID
	Amount float64
	Category string
	Note string
	CreatedAt time.Time
}
