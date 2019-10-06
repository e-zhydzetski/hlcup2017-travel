package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/e-zhydzetski/hlcup2017-travel/internal/domain"
)

func NewHandler(service domain.Service) http.Handler {
	r := chi.NewRouter()
	r.Get("/users/{id}", ErrorAware(func(w http.ResponseWriter, r *http.Request) error {
		id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
		if err != nil {
			return domain.ErrNotFound
		}
		user, err := service.GetUser(uint32(id))
		if err != nil {
			return err
		}
		viewDTO := newUserViewDTOFromDomain(user)
		return json.NewEncoder(w).Encode(viewDTO)
	}))
	r.Post("/users/{id}", ErrorAware(func(w http.ResponseWriter, r *http.Request) error {
		idStr := chi.URLParam(r, "id")
		if idStr == "new" { // create
			var dto UserCreateDTO
			err := json.NewDecoder(r.Body).Decode(&dto)
			if err != nil {
				return domain.ErrIllegalArgument
			}
			create, err := dto.toValidDomain()
			if err != nil {
				return err
			}
			err = service.CreateUser(create)
			if err != nil {
				return err
			}
		} else { // update
			id, err := strconv.ParseUint(idStr, 10, 32)
			if err != nil {
				return domain.ErrNotFound
			}
			var dto UserUpdateDTO
			err = json.NewDecoder(r.Body).Decode(&dto)
			if err != nil {
				return domain.ErrIllegalArgument
			}
			update, err := dto.toValidDomain()
			if err != nil {
				return err
			}
			err = service.UpdateUser(uint32(id), update)
			if err != nil {
				return err
			}
		}

		_, _ = w.Write([]byte("{}"))
		return nil
	}))
	r.Get("/locations/{id}", ErrorAware(func(w http.ResponseWriter, r *http.Request) error {
		id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
		if err != nil {
			return domain.ErrNotFound
		}
		loc, err := service.GetLocation(uint32(id))
		if err != nil {
			return err
		}
		viewDTO := newLocationViewDTOFromDomain(loc)
		return json.NewEncoder(w).Encode(viewDTO)
	}))
	r.Post("/locations/{id}", ErrorAware(func(w http.ResponseWriter, r *http.Request) error {
		idStr := chi.URLParam(r, "id")
		if idStr == "new" { // create
			var dto LocationCreateDTO
			err := json.NewDecoder(r.Body).Decode(&dto)
			if err != nil {
				return domain.ErrIllegalArgument
			}
			create, err := dto.toValidDomain()
			if err != nil {
				return err
			}
			err = service.CreateLocation(create)
			if err != nil {
				return err
			}
		} else { // update
			id, err := strconv.ParseUint(idStr, 10, 32)
			if err != nil {
				return domain.ErrNotFound
			}
			var dto LocationUpdateDTO
			err = json.NewDecoder(r.Body).Decode(&dto)
			if err != nil {
				return domain.ErrIllegalArgument
			}
			update, err := dto.toValidDomain()
			if err != nil {
				return err
			}
			err = service.UpdateLocation(uint32(id), update)
			if err != nil {
				return err
			}
		}

		_, _ = w.Write([]byte("{}"))
		return nil
	}))
	r.Get("/visits/{id}", ErrorAware(func(w http.ResponseWriter, r *http.Request) error {
		id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
		if err != nil {
			return domain.ErrNotFound
		}
		visit, err := service.GetVisit(uint32(id))
		if err != nil {
			return err
		}
		viewDTO := newVisitViewDTOFromDomain(visit)
		return json.NewEncoder(w).Encode(viewDTO)
	}))
	r.Post("/visits/{id}", ErrorAware(func(w http.ResponseWriter, r *http.Request) error {
		idStr := chi.URLParam(r, "id")
		if idStr == "new" { // create
			var dto VisitCreateDTO
			err := json.NewDecoder(r.Body).Decode(&dto)
			if err != nil {
				return domain.ErrIllegalArgument
			}
			create, err := dto.toValidDomain()
			if err != nil {
				return err
			}
			err = service.CreateVisit(create)
			if err != nil {
				return err
			}
		} else { // update
			id, err := strconv.ParseUint(idStr, 10, 32)
			if err != nil {
				return domain.ErrNotFound
			}
			var dto VisitUpdateDTO
			err = json.NewDecoder(r.Body).Decode(&dto)
			if err != nil {
				return domain.ErrIllegalArgument
			}
			update, err := dto.toValidDomain()
			if err != nil {
				return err
			}
			err = service.UpdateVisit(uint32(id), update)
			if err != nil {
				return err
			}
		}

		_, _ = w.Write([]byte("{}"))
		return nil
	}))
	r.Get("/users/{id}/visits", ErrorAware(func(w http.ResponseWriter, r *http.Request) error {
		params := &domain.GetUserVisitsParams{}
		id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
		if err != nil {
			return domain.ErrNotFound
		}
		params.UserID = uint32(id)
		q := r.URL.Query()
		fromDate := q.Get("fromDate")
		if fromDate != "" {
			d, err := strconv.ParseInt(fromDate, 10, 64)
			if err != nil {
				return domain.ErrIllegalArgument
			}
			params.FromDate = &d
		}
		toDate := q.Get("toDate")
		if toDate != "" {
			d, err := strconv.ParseInt(toDate, 10, 64)
			if err != nil {
				return domain.ErrIllegalArgument
			}
			params.ToDate = &d
		}
		country := q.Get("country")
		if country != "" {
			params.Country = &country
		}
		toDist := q.Get("toDistance")
		if toDist != "" {
			t, err := strconv.ParseUint(toDist, 10, 32)
			if err != nil {
				return domain.ErrIllegalArgument
			}
			d := uint32(t)
			params.ToDistance = &d
		}
		visits, err := service.GetUserVisits(params)
		if err != nil {
			return err
		}
		dto := newUserVisitsDTOFromDomain(visits)
		return json.NewEncoder(w).Encode(dto)
	}))
	r.Get("/locations/{id}/avg", ErrorAware(func(w http.ResponseWriter, r *http.Request) error {
		params := &domain.GetLocationAvgParams{}
		id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
		if err != nil {
			return domain.ErrNotFound
		}
		params.LocationID = uint32(id)
		q := r.URL.Query()
		fromDate := q.Get("fromDate")
		if fromDate != "" {
			d, err := strconv.ParseInt(fromDate, 10, 64)
			if err != nil {
				return domain.ErrIllegalArgument
			}
			params.FromDate = &d
		}
		toDate := q.Get("toDate")
		if toDate != "" {
			d, err := strconv.ParseInt(toDate, 10, 64)
			if err != nil {
				return domain.ErrIllegalArgument
			}
			params.ToDate = &d
		}
		fromAge := q.Get("fromAge")
		if fromAge != "" {
			t, err := strconv.ParseUint(fromAge, 10, 32)
			if err != nil {
				return domain.ErrIllegalArgument
			}
			d := uint(t)
			params.FromAge = &d
		}
		toAge := q.Get("toAge")
		if toAge != "" {
			t, err := strconv.ParseUint(toAge, 10, 32)
			if err != nil {
				return domain.ErrIllegalArgument
			}
			d := uint(t)
			params.ToAge = &d
		}
		gender := q.Get("gender")
		if gender != "" {
			if gender != "m" && gender != "f" {
				return domain.ErrIllegalArgument
			}
			params.Gender = &gender
		}
		avg, err := service.GetLocationAvg(params)
		if err != nil {
			return err
		}
		dto := newLocationAvgDTOFromDomain(avg)
		return json.NewEncoder(w).Encode(dto)
	}))
	return r
}

func ErrorAware(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			switch err {
			case domain.ErrIllegalArgument:
				w.WriteHeader(http.StatusBadRequest)
			case domain.ErrNotFound:
				w.WriteHeader(http.StatusNotFound)
			default:
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte(err.Error()))
			}
		}
	}
}
