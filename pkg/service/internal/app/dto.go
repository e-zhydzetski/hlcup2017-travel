package app

import "github.com/e-zhydzetski/hlcup2017-travel/pkg/service/internal/domain"

type LocationCreateDTO struct {
	ID       uint32
	Place    string
	Country  string
	City     string
	Distance uint32
}

func (dto *LocationCreateDTO) toDomain() *domain.Location {
	return &domain.Location{
		ID:       dto.ID,
		Place:    dto.Place,
		Country:  dto.Country,
		City:     dto.City,
		Distance: dto.Distance,
	}
}

type LocationUpdateDTO struct {
	Place    *string
	Country  *string
	City     *string
	Distance *uint32
}

type UserCreateDTO struct {
	ID        uint32
	Email     string
	FirstName string
	LastName  string
	Gender    string
	BirthDate int64
}

func (dto *UserCreateDTO) toDomain() *domain.User {
	return &domain.User{
		ID:        dto.ID,
		Email:     dto.Email,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Gender:    dto.Gender,
		BirthDate: dto.BirthDate,
	}
}

type UserUpdateDTO struct {
	Email     *string
	FirstName *string
	LastName  *string
	Gender    *string
	BirthDate *int64
}

type VisitCreateDTO struct {
	ID         uint32
	LocationID uint32
	UserID     uint32
	VisitedAt  int64
	Mark       int
}

func (dto *VisitCreateDTO) toDomain() *domain.Visit {
	return &domain.Visit{
		ID:         dto.ID,
		LocationID: dto.LocationID,
		UserID:     dto.UserID,
		VisitedAt:  dto.VisitedAt,
		Mark:       dto.Mark,
	}
}

type VisitUpdateDTO struct {
	LocationID *uint32
	UserID     *uint32
	VisitedAt  *int64
	Mark       *int
}
