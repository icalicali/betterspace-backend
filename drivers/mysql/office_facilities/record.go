package officefacilities

import (
	officeFacilityUseCase "backend/businesses/office_facilities"
	"backend/drivers/mysql/facilities"
	"backend/drivers/mysql/offices"
)

type OfficeFacility struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	FacilitiesID uint         `json:"-"`
	Facility     facilities.Facility `json:"facility,omitempty" gorm:"foreignKey:FacilitiesID;references:ID"`
	OfficeID     uint           `json:"-"`
	Office       offices.Office `json:"office" gorm:"foreignKey:OfficeID;references:ID"`
}

func FromDomain(domain *officeFacilityUseCase.Domain) *OfficeFacility {
	return &OfficeFacility{
		ID:           domain.ID,
		FacilitiesID: domain.FacilitiesID,
		OfficeID:     domain.OfficeID,
	}
}

func (rec *OfficeFacility) ToDomain() officeFacilityUseCase.Domain {
	return officeFacilityUseCase.Domain{
		ID:           rec.ID,
		FacilitiesID: rec.FacilitiesID,
		OfficeID:     rec.OfficeID,
	}
}
