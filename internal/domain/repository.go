package domain

import "log"

type Dump struct {
	Users     []*UserCreateDTO
	Locations []*LocationCreateDTO
	Visits    []*VisitCreateDTO
}

func NewRepositoryFromDump(dump *Dump) *Repository {
	repo := &Repository{
		users:     make(map[uint32]*User),
		locations: make(map[uint32]*Location),
		visits:    make(map[uint32]*Visit),
	}

	for _, u := range dump.Users {
		d := u.toDomain()
		repo.users[d.ID] = d
	}
	log.Println("Users count:", len(repo.users))

	for _, l := range dump.Locations {
		d := l.toDomain()
		repo.locations[d.ID] = d
	}
	log.Println("Locations count:", len(repo.locations))

	for _, v := range dump.Visits {
		d := v.toDomain()
		repo.visits[d.ID] = d
	}
	log.Println("Visits count:", len(repo.visits))

	log.Println("Dump loaded to repository")

	return repo
}

type Repository struct {
	users     map[uint32]*User
	locations map[uint32]*Location
	visits    map[uint32]*Visit
}

func (r Repository) CreateUser(createDTO *UserCreateDTO) error {
	_, ok := r.users[createDTO.ID]
	if ok {
		return ErrIllegalArgument
	}
	u := createDTO.toDomain()
	r.users[u.ID] = u
	return nil
}

func (r Repository) UpdateUser(id uint32, updateDTO *UserUpdateDTO) error {
	u, ok := r.users[id]
	if !ok {
		return ErrNotFound
	}
	if updateDTO.Email != nil {
		u.Email = *updateDTO.Email
	}
	if updateDTO.FirstName != nil {
		u.FirstName = *updateDTO.FirstName
	}
	if updateDTO.LastName != nil {
		u.LastName = *updateDTO.LastName
	}
	if updateDTO.Gender != nil {
		u.Gender = *updateDTO.Gender
	}
	if updateDTO.BirthDate != nil {
		u.BirthDate = *updateDTO.BirthDate
	}
	return nil
}

func (r Repository) GetUser(id uint32) (*User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, ErrNotFound
	}
	return u, nil
}

func (r Repository) CreateLocation(createDTO *LocationCreateDTO) error {
	_, ok := r.locations[createDTO.ID]
	if ok {
		return ErrIllegalArgument
	}
	l := createDTO.toDomain()
	r.locations[l.ID] = l
	return nil
}

func (r Repository) UpdateLocation(id uint32, updateDTO *LocationUpdateDTO) error {
	l, ok := r.locations[id]
	if !ok {
		return ErrNotFound
	}
	if updateDTO.Place != nil {
		l.Place = *updateDTO.Place
	}
	if updateDTO.Country != nil {
		l.Country = *updateDTO.Country
	}
	if updateDTO.City != nil {
		l.City = *updateDTO.City
	}
	if updateDTO.Place != nil {
		l.Distance = *updateDTO.Distance
	}
	return nil
}

func (r Repository) GetLocation(id uint32) (*Location, error) {
	l, ok := r.locations[id]
	if !ok {
		return nil, ErrNotFound
	}
	return l, nil
}

func (r Repository) CreateVisit(createDTO *VisitCreateDTO) error {
	_, ok := r.visits[createDTO.ID]
	if ok {
		return ErrIllegalArgument
	}
	v := createDTO.toDomain()
	r.visits[v.ID] = v
	return nil
}

func (r Repository) UpdateVisit(id uint32, updateDTO *VisitUpdateDTO) error {
	v, ok := r.visits[id]
	if !ok {
		return ErrNotFound
	}
	if updateDTO.LocationID != nil {
		v.LocationID = *updateDTO.LocationID
	}
	if updateDTO.UserID != nil {
		v.UserID = *updateDTO.UserID
	}
	if updateDTO.VisitedAt != nil {
		v.VisitedAt = *updateDTO.VisitedAt
	}
	if updateDTO.Mark != nil {
		v.Mark = *updateDTO.Mark
	}
	return nil
}

func (r Repository) GetVisit(id uint32) (*Visit, error) {
	v, ok := r.visits[id]
	if !ok {
		return nil, ErrNotFound
	}
	return v, nil
}

func (r Repository) GetUserVisits(params *GetUserVisitsParams) (*UserVisits, error) {
	return &UserVisits{Visits: []*UserVisit{}}, nil
}

func (r Repository) GetLocationAvg(params *GetLocationAvgParams) (*LocationAvg, error) {
	avg := LocationAvg(0)
	return &avg, nil
}
