package appointment

import (
	"time"

	"github.com/geraldofigueiredo/cadastro-ccz/internal/domains/user"
)

type DayShift string

const (
	Matutino   DayShift = "Matutino"
	Vespertino DayShift = "Vespertino"
)

type Appointment struct {
	ID         int
	Reserved   *time.Time
	Day        uint
	Month      uint
	Year       uint
	Shift      DayShift
	AnimalType AnimalType
	AnimalName *string
	User       user.User

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (Appointment) TableName() string {
	return "appointment"
}
