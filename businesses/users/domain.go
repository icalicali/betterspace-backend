package users

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        	string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	gorm.DeletedAt
	FullName  	string
	Gender		string
	Email     	string
	Password  	string
	Image		string
	Roles		bool
}

type LoginDomain struct {
	Email     	string
	Password  	string
}

type Usecase interface {
	Register(userDomain *Domain) Domain
	Login(userDomain *LoginDomain) string
}

type Repository interface {
	Register(userDomain *Domain) Domain
	GetByEmail(userDomain *LoginDomain) Domain
}
