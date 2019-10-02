package domain

type LocationCreateDTO struct {
	ID       uint32
	Place    string
	Country  string
	City     string
	Distance uint32
}

func (dto *LocationCreateDTO) toDomain() *Location {
	return &Location{
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

func (dto *UserCreateDTO) toDomain() *User {
	return &User{
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

func (dto *VisitCreateDTO) toDomain() *Visit {
	return &Visit{
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
