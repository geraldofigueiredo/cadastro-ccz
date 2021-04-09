package user

import "time"

type UserType string
type UserGender string

const (
	Adm   UserType = "ADM"
	Ong   UserType = "ONG/PROTETOR(a)"
	Civil UserType = "CIVIL"
)

const (
	Male   UserGender = "MALE"
	Female UserGender = "FEMALE"
)

type User struct {
	ID        int
	Name      string
	CPF       string
	BirthDate time.Time
	Type      UserType
	Gender    UserGender

	Phone1   string
	Phone2   string
	Email    string
	Address  Address
	Password string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
