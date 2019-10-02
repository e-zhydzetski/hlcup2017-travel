package http

import "github.com/e-zhydzetski/hlcup2017-travel/internal/domain"

type Location struct {
	ID       uint32 `json:"id"`
	Place    string `json:"place"`
	Country  string `json:"country"`
	City     string `json:"city"`
	Distance uint32 `json:"distance"`
}

type UserViewDTO struct {
	ID        uint32 `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	BirthDate int64  `json:"birth_date"`
}

func newUserViewDTOFromDomain(d *domain.User) *UserViewDTO {
	return &UserViewDTO{
		ID:        d.ID,
		Email:     d.Email,
		FirstName: d.FirstName,
		LastName:  d.LastName,
		Gender:    d.Gender,
		BirthDate: d.BirthDate,
	}
}

type Visit struct {
	ID         uint32 `json:"id"`
	LocationID uint32 `json:"location"`
	UserID     uint32 `json:"user"`
	VisitedAt  int64  `json:"visited_at"`
	Mark       int    `json:"mark"`
}
