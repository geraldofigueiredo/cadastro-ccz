package appointment

import (
	"time"

	"github.com/geraldofigueiredo/cadastro-ccz/pkg/user"
)

type DayShift string

const (
	Matutino   DayShift = "Matutino"
	Vespertino DayShift = "Vespertino"
)

type Appointment struct {
	ID         int
	Reserved   *time.Time
	Day        time.Time
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
