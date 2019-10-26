package dump

import (
	"github.com/e-zhydzetski/hlcup2017-travel/pkg/service/internal/app"
)

type User struct {
	ID        uint32 `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	BirthDate int64  `json:"birth_date"`
}

func (u *User) toDomainCreateDTO() *app.UserCreateDTO {
	return &app.UserCreateDTO{
		ID:        u.ID,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Gender:    u.Gender,
		BirthDate: u.BirthDate,
	}
}

type UserList struct {
	Users []*User `json:"users"`
}

type Location struct {
	ID       uint32 `json:"id"`
	Place    string `json:"place"`
	Country  string `json:"country"`
	City     string `json:"city"`
	Distance uint32 `json:"distance"`
}

func (l *Location) toDomainCreateDTO() *app.LocationCreateDTO {
	return &app.LocationCreateDTO{
		ID:       l.ID,
		Place:    l.Place,
		Country:  l.Country,
		City:     l.City,
		Distance: l.Distance,
	}
}

type LocationList struct {
	Locations []*Location `json:"locations"`
}

type Visit struct {
	ID         uint32 `json:"id"`
	LocationID uint32 `json:"location"`
	UserID     uint32 `json:"user"`
	VisitedAt  int64  `json:"visited_at"`
	Mark       int    `json:"mark"`
}

func (v *Visit) toDomainCreateDTO() *app.VisitCreateDTO {
	return &app.VisitCreateDTO{
		ID:         v.ID,
		LocationID: v.LocationID,
		UserID:     v.UserID,
		VisitedAt:  v.VisitedAt,
		Mark:       v.Mark,
	}
}

type VisitList struct {
	Visits []*Visit `json:"visits"`
}
