package appointment

import (
	"time"

	"github.com/geraldofigueiredo/cadastro-ccz/pkg/animal"
	"github.com/geraldofigueiredo/cadastro-ccz/pkg/user"
)

type Appointment struct {
	ID         int
	Reserved   *time.Time
	Day        time.Time
	Shift      DayShift
	AnimalType animal.AnimalType
	AnimalName *string
	User       user.User

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
