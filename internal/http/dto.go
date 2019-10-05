package http

import "github.com/e-zhydzetski/hlcup2017-travel/internal/domain"

type UserCreateDTO struct {
	ID        uint32 `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	BirthDate int64  `json:"birth_date"`
}

func (d *UserCreateDTO) toDomain() *domain.UserCreateDTO {
	return &domain.UserCreateDTO{
		ID:        d.ID,
		Email:     d.Email,
		FirstName: d.FirstName,
		LastName:  d.LastName,
		Gender:    d.Gender,
		BirthDate: d.BirthDate,
	}
}

type UserUpdateDTO struct {
	Email     *string `json:"email"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Gender    *string `json:"gender"`
	BirthDate *int64  `json:"birth_date"`
}

func (d *UserUpdateDTO) toDomain() *domain.UserUpdateDTO {
	return &domain.UserUpdateDTO{
		Email:     d.Email,
		FirstName: d.FirstName,
		LastName:  d.LastName,
		Gender:    d.Gender,
		BirthDate: d.BirthDate,
	}
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

type LocationCreateDTO struct {
	ID       uint32 `json:"id"`
	Place    string `json:"place"`
	Country  string `json:"country"`
	City     string `json:"city"`
	Distance uint32 `json:"distance"`
}

type LocationUpdateDTO struct {
	Place    *string `json:"place"`
	Country  *string `json:"country"`
	City     *string `json:"city"`
	Distance *uint32 `json:"distance"`
}

type LocationViewDTO struct {
	ID       uint32 `json:"id"`
	Place    string `json:"place"`
	Country  string `json:"country"`
	City     string `json:"city"`
	Distance uint32 `json:"distance"`
}

func newLocationViewDTOFromDomain(d *domain.Location) *LocationViewDTO {
	return &LocationViewDTO{
		ID:       d.ID,
		Place:    d.Place,
		Country:  d.Country,
		City:     d.City,
		Distance: d.Distance,
	}
}

type VisitCreateDTO struct {
	ID         uint32 `json:"id"`
	LocationID uint32 `json:"location"`
	UserID     uint32 `json:"user"`
	VisitedAt  int64  `json:"visited_at"`
	Mark       int    `json:"mark"`
}

type VisitUpdateDTO struct {
	LocationID *uint32 `json:"location"`
	UserID     *uint32 `json:"user"`
	VisitedAt  *int64  `json:"visited_at"`
	Mark       *int    `json:"mark"`
}

type VisitViewDTO struct {
	ID         uint32 `json:"id"`
	LocationID uint32 `json:"location"`
	UserID     uint32 `json:"user"`
	VisitedAt  int64  `json:"visited_at"`
	Mark       int    `json:"mark"`
}

func newVisitViewDTOFromDomain(d *domain.Visit) *VisitViewDTO {
	return &VisitViewDTO{
		ID:         d.ID,
		LocationID: d.LocationID,
		UserID:     d.UserID,
		VisitedAt:  d.VisitedAt,
		Mark:       d.Mark,
	}
}

type UserVisitDTO struct {
	Mark      int    `json:"mark"`
	VisitedAt int64  `json:"visited_at"`
	Place     string `json:"place"`
}

type UserVisitsDTO struct {
	Visits []*UserVisitDTO `json:"visits"`
}

func newUserVisitsDTOFromDomain(d *domain.UserVisits) *UserVisitsDTO {
	visitsDTO := &UserVisitsDTO{
		Visits: make([]*UserVisitDTO, len(d.Visits)),
	}
	for i, v := range d.Visits {
		visitsDTO.Visits[i] = &UserVisitDTO{
			Mark:      v.Mark,
			VisitedAt: v.VisitedAt,
			Place:     v.Place,
		}
	}
	return visitsDTO
}
