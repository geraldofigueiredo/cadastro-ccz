package user

import "time"

type UserType string

const (
	Adm   UserType = "ADM"
	Ong   UserType = "ONG/PROTETOR(a)"
	Civil UserType = "CIVIL"
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
