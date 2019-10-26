package http

import (
	"github.com/e-zhydzetski/hlcup2017-travel/pkg/service/internal/app"
	"github.com/e-zhydzetski/hlcup2017-travel/pkg/service/internal/domain"
)

type UserCreateDTO struct {
	ID        uint32 `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	BirthDate int64  `json:"birth_date"`
}

func (d *UserCreateDTO) toValidDomain() (*app.UserCreateDTO, error) {
	if d.ID == 0 {
		return nil, app.ErrIllegalArgument
	}
	return &app.UserCreateDTO{
		ID:        d.ID,
		Email:     d.Email,
		FirstName: d.FirstName,
		LastName:  d.LastName,
		Gender:    d.Gender,
		BirthDate: d.BirthDate,
	}, nil
}

type UserUpdateDTO map[string]interface{}

func (d UserUpdateDTO) toValidDomain() (*app.UserUpdateDTO, error) {
	res := &app.UserUpdateDTO{}
	if _, exists := d["id"]; exists {
		return nil, app.ErrIllegalArgument
	}
	if val, exists := d["email"]; exists {
		if val == nil {
			return nil, app.ErrIllegalArgument
		}
		s := val.(string)
		res.Email = &s
	}
	if val, exists := d["first_name"]; exists {
		if val == nil {
			return nil, app.ErrIllegalArgument
		}
		s := val.(string)
		res.FirstName = &s
	}
	if val, exists := d["last_name"]; exists {
		if val == nil {
			return nil, app.ErrIllegalArgument
		}
		s := val.(string)
		res.LastName = &s
	}
	if val, exists := d["gender"]; exists {
		if val == nil {
			return nil, app.ErrIllegalArgument
		}
		s := val.(string)
		res.Gender = &s
	}
	if val, exists := d["birth_date"]; exists {
		if val == nil {
			return nil, app.ErrIllegalArgument
		}
		i := int64(val.(float64))
		res.BirthDate = &i
	}
	return res, nil
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

func (d *LocationCreateDTO) toValidDomain() (*app.LocationCreateDTO, error) {
	if d.ID == 0 {
		return nil, app.ErrIllegalArgument
	}
	return &app.LocationCreateDTO{
		ID:       d.ID,
		Place:    d.Place,
		Country:  d.Country,
		City:     d.City,
		Distance: d.Distance,
	}, nil
}

type LocationUpdateDTO map[string]interface{}

func (d LocationUpdateDTO) toValidDomain() (*app.LocationUpdateDTO, error) {
	res := &app.LocationUpdateDTO{}
	if _, exists := d["id"]; exists {
		return nil, app.ErrIllegalArgument
	}
	if val, exists := d["place"]; exists {
		if val == nil {
			return nil, app.ErrIllegalArgument
		}
		s := val.(string)
		res.Place = &s
	}
	if val, exists := d["country"]; exists {
		if val == nil {
			return nil, app.ErrIllegalArgument
		}
		s := val.(string)
		res.Country = &s
	}
	if val, exists := d["city"]; exists {
		if val == nil {
			return nil, app.ErrIllegalArgument
		}
		s := val.(string)
		res.City = &s
	}
	if val, exists := d["distance"]; exists {
		if val == nil {
			return nil, app.ErrIllegalArgument
		}
		i := uint32(val.(float64))
		res.Distance = &i
	}
	return res, nil
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

func (d *VisitCreateDTO) toValidDomain() (*app.VisitCreateDTO, error) {
	if d.ID == 0 {
		return nil, app.ErrIllegalArgument
	}
	return &app.VisitCreateDTO{
		ID:         d.ID,
		LocationID: d.LocationID,
		UserID:     d.UserID,
		VisitedAt:  d.VisitedAt,
		Mark:       d.Mark,
	}, nil
}

type VisitUpdateDTO map[string]interface{}

func (d VisitUpdateDTO) toValidDomain() (*app.VisitUpdateDTO, error) {
	res := &app.VisitUpdateDTO{}
	if _, exists := d["id"]; exists {
		return nil, app.ErrIllegalArgument
	}
	if val, exists := d["location"]; exists {
		if val == nil {
			return nil, app.ErrIllegalArgument
		}
		i := uint32(val.(float64))
		res.LocationID = &i
	}
	if val, exists := d["user"]; exists {
		if val == nil {
			return nil, app.ErrIllegalArgument
		}
		i := uint32(val.(float64))
		res.UserID = &i
	}
	if val, exists := d["visited_at"]; exists {
		if val == nil {
			return nil, app.ErrIllegalArgument
		}
		i := int64(val.(float64))
		res.VisitedAt = &i
	}
	if val, exists := d["mark"]; exists {
		if val == nil {
			return nil, app.ErrIllegalArgument
		}
		i := int(val.(float64))
		res.Mark = &i
	}
	return res, nil
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

func newUserVisitsDTOFromDomain(d *app.UserVisits) *UserVisitsDTO {
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

type LocationAvgDTO struct {
	Avg app.LocationAvg `json:"avg"`
}

func newLocationAvgDTOFromDomain(avg *app.LocationAvg) *LocationAvgDTO {
	return &LocationAvgDTO{
		Avg: *avg,
	}
}
