package response

import (
	transactions "backend/businesses/transactions"

	"gorm.io/gorm"
)

type Transaction struct {
	ID            uint           `json:"id"`
	CreatedAt     string         `json:"created_at"`
	UpdatedAt     string         `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
	Duration      int            `json:"duration"`
	CheckIn       checkIn        `json:"check_in"`
	CheckOut      checkOut       `json:"check_out"`
	Price         uint           `json:"price"`
	Drink         string         `json:"drink"`
	Status        string         `json:"status"`
	PaymentMethod string         `json:"payment_method"`
	User          user           `json:"user"`
	Office        office         `json:"office"`
}

func FromDomain(domain transactions.Domain) Transaction {
	return Transaction{
		ID:            domain.ID,
		CreatedAt:     domain.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt:     domain.UpdatedAt.Format("02-01-2006 15:04:05"),
		DeletedAt:     domain.DeletedAt,
		Duration:      domain.Duration,
		Price:         domain.Price,
		Drink:         domain.Drink,
		Status:        domain.Status,
		PaymentMethod: domain.PaymentMethod,
		CheckIn: checkIn{
			Date: domain.CheckIn.Format("02-01-2006"),
			Time: domain.CheckIn.Format("15:04"),
		},
		CheckOut: checkOut{
			Date: domain.CheckOut.Format("02-01-2006"),
			Time: domain.CheckOut.Format("15:04"),
		},
		User: user{
			UserID:   domain.UserID,
			FullName: domain.UserFullName,
			Email:    domain.UserEmail,
		},
		Office: office{
			OfficeID:   domain.OfficeID,
			OfficeName: domain.OfficeName,
			OfficeType: domain.OfficeType,
		},
	}
}
