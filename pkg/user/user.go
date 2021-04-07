package user

import "time"

type UserType string

const (
	ADM   UserType = "ADM"
	ONG   UserType = "ONG/PROTETOR(a)"
	Civil UserType = "Civil"
)

type User struct {
	ID    int
	Name  string
	CPF   int
	Phone string
	Type  UserType

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
