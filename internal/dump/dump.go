package dump

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/e-zhydzetski/hlcup2017-travel/internal/domain"
)

type Dump struct {
	users     []*User
	locations []*Location
	visits    []*Visit
}

func LoadFromFolder(path string) (*Dump, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	dump := &Dump{}

	for _, file := range files {
		fileName := file.Name()
		if strings.HasPrefix(fileName, "users_") {
			data, err := ioutil.ReadFile(path + "/" + fileName)
			if err != nil {
				return nil, err
			}
			var list UserList
			if err = json.Unmarshal(data, &list); err != nil {
				return nil, err
			}
			dump.users = append(dump.users, list.Users...)
			continue
		}
		if strings.HasPrefix(fileName, "locations_") {
			data, err := ioutil.ReadFile(path + "/" + fileName)
			if err != nil {
				return nil, err
			}
			var list LocationList
			if err = json.Unmarshal(data, &list); err != nil {
				return nil, err
			}
			dump.locations = append(dump.locations, list.Locations...)
			continue
		}
		if strings.HasPrefix(fileName, "visits_") {
			data, err := ioutil.ReadFile(path + "/" + fileName)
			if err != nil {
				return nil, err
			}
			var list VisitList
			if err = json.Unmarshal(data, &list); err != nil {
				return nil, err
			}
			dump.visits = append(dump.visits, list.Visits...)
			continue
		}
	}

	return dump, nil
}

func (d *Dump) ToDomain() *domain.Dump {
	res := &domain.Dump{} // TODO preallocate slices
	for _, u := range d.users {
		res.Users = append(res.Users, u.toDomainCreateDTO())
	}
	for _, l := range d.locations {
		res.Locations = append(res.Locations, l.toDomainCreateDTO())
	}
	for _, v := range d.visits {
		res.Visits = append(res.Visits, v.toDomainCreateDTO())
	}
	return res
}