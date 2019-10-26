package dump

import (
	"archive/zip"
	"encoding/json"
	"github.com/e-zhydzetski/hlcup2017-travel/pkg/service/internal/app"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Dump struct {
	users     []*User
	locations []*Location
	visits    []*Visit
}

func (d *Dump) load(name string, reader io.Reader) error {
	if strings.HasPrefix(name, "users_") {
		data, err := ioutil.ReadAll(reader)
		if err != nil {
			return err
		}
		var list UserList
		if err := json.Unmarshal(data, &list); err != nil {
			return err
		}
		d.users = append(d.users, list.Users...)
		return nil
	}
	if strings.HasPrefix(name, "locations_") {
		data, err := ioutil.ReadAll(reader)
		if err != nil {
			return err
		}
		var list LocationList
		if err := json.Unmarshal(data, &list); err != nil {
			return err
		}
		d.locations = append(d.locations, list.Locations...)
		return nil
	}
	if strings.HasPrefix(name, "visits_") {
		data, err := ioutil.ReadAll(reader)
		if err != nil {
			return err
		}
		var list VisitList
		if err := json.Unmarshal(data, &list); err != nil {
			return err
		}
		d.visits = append(d.visits, list.Visits...)
		return nil
	}
	return nil
}

func LoadFromFolder(path string) (*Dump, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	dump := &Dump{}

	for _, file := range files {
		f, err := os.Open(path + "/" + file.Name())
		if err != nil {
			return nil, err
		}
		err = dump.load(file.Name(), f)
		_ = f.Close()
		if err != nil {
			return nil, err
		}
	}

	return dump, nil
}

func LoadFromZip(path string) (*Dump, error) {
	r, err := zip.OpenReader(path)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	dump := &Dump{}

	for _, file := range r.File {
		rc, err := file.Open()
		if err != nil {
			return nil, err
		}
		err = dump.load(file.Name, rc)
		_ = rc.Close()
		if err != nil {
			return nil, err
		}
	}

	return dump, nil
}

func (d *Dump) ToDomain() *app.Dump {
	res := &app.Dump{} // TODO preallocate slices
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