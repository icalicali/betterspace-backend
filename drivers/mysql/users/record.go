package users

import (
	"backend/businesses/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	FullName  string         `json:"full_name" gorm:"type:varchar(100)"`
	Email     string         `json:"email" gorm:"type:varchar(100);unique" faker:"email"`
	Password  string         `json:"password" faker:"password"`
	Gender    string         `json:"gender"`
	Photo     string         `json:"photo" gorm:"type:varchar(100)"`
	Roles     string         `gorm:"type:enum('user', 'admin')" json:"roles"`
}

func FromDomain(domain *users.Domain) *User {
	return &User{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		FullName:  domain.FullName,
		Gender:    domain.Gender,
		Email:     domain.Email,
		Password:  domain.Email,
		Photo:     domain.Photo,
		Roles:     domain.Roles,
	}
}

func (rec *User) ToDomain() users.Domain {
	return users.Domain{
		ID:        rec.ID,
		FullName:  rec.FullName,
		Email:     rec.Email,
		Password:  rec.Password,
		Gender:    rec.Gender,
		Photo:     rec.Photo,
		Roles:     rec.Roles,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
	}
}