package response

import facilities "backend/businesses/facilities"

type Facility struct {
	ID          uint   `json:"id" form:"id"`
	Description string `json:"description" form:"description"`
	Slug		string `json:"slug" form:"slug"`
}

func FromDomain(domain facilities.Domain) Facility {
	return Facility{
		ID:          domain.ID,
		Description: domain.Description,
		Slug: domain.Slug,
	}
}
