package domain

import "github.com/e-zhydzetski/hlcup2017-travel/internal/x/xerror"

type GetUserVisitsParams struct {
	UserID     uint32
	FromDate   *int64
	ToDate     *int64
	Country    *string
	ToDistance *uint32
}

type UserVisit struct {
	Mark      int
	VisitedAt int64
	Place     string
}

type UserVisits struct {
	Visits []*UserVisit
}

type GetLocationAvgParams struct {
	LocationID uint32
	FromDate   *int64
	ToDate     *int64
	FromAge    *uint
	ToAge      *uint
	Gender     *string
}

type LocationAvg float32

const ErrNotFound = xerror.Error("not found")
const ErrIllegalArgument = xerror.Error("illegal argument")

type Service interface {
	CreateUser(createDTO *UserCreateDTO) error
	UpdateUser(ID uint32, updateDTO *UserUpdateDTO) error
	GetUser(id uint32) (*User, error)

	CreateLocation(createDTO *LocationCreateDTO) error
	UpdateLocation(ID uint32, updateDTO *LocationUpdateDTO) error
	GetLocation(id uint32) (*Location, error)

	CreateVisit(createDTO *VisitCreateDTO) error
	UpdateVisit(ID uint32, updateDTO *VisitUpdateDTO) error
	GetVisit(id uint32) (*Visit, error)

	GetUserVisits(params *GetUserVisitsParams) (*UserVisits, error)
	GetLocationAvg(params *GetLocationAvgParams) (*LocationAvg, error)
}
