package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/e-zhydzetski/hlcup2017-travel/internal/domain"
)

func NewHandler(service domain.Service) http.Handler {
	emptyJSON := map[string]interface{}{}

	r := chi.NewRouter()
	r.Use(translateConnectionHeader())
	r.Get("/users/{id}", errorAware(func(w http.ResponseWriter, r *http.Request) error {
		id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
		if err != nil {
			return domain.ErrNotFound
		}
		user, err := service.GetUser(uint32(id))
		if err != nil {
			return err
		}
		dto := newUserViewDTOFromDomain(user)
		return jsonResponse(w, dto)
	}))
	r.Post("/users/{id}", errorAware(func(w http.ResponseWriter, r *http.Request) error {
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
		return jsonResponse(w, emptyJSON)
	}))
	r.Get("/locations/{id}", errorAware(func(w http.ResponseWriter, r *http.Request) error {
		id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
		if err != nil {
			return domain.ErrNotFound
		}
		loc, err := service.GetLocation(uint32(id))
		if err != nil {
			return err
		}
		dto := newLocationViewDTOFromDomain(loc)
		return jsonResponse(w, dto)
	}))
	r.Post("/locations/{id}", errorAware(func(w http.ResponseWriter, r *http.Request) error {
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
		return jsonResponse(w, emptyJSON)
	}))
	r.Get("/visits/{id}", errorAware(func(w http.ResponseWriter, r *http.Request) error {
		id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)
		if err != nil {
			return domain.ErrNotFound
		}
		visit, err := service.GetVisit(uint32(id))
		if err != nil {
			return err
		}
		dto := newVisitViewDTOFromDomain(visit)
		return jsonResponse(w, dto)
	}))
	r.Post("/visits/{id}", errorAware(func(w http.ResponseWriter, r *http.Request) error {
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
		return jsonResponse(w, emptyJSON)
	}))
	r.Get("/users/{id}/visits", errorAware(func(w http.ResponseWriter, r *http.Request) error {
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
		return jsonResponse(w, dto)
	}))
	r.Get("/locations/{id}/avg", errorAware(func(w http.ResponseWriter, r *http.Request) error {
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
		return jsonResponse(w, dto)
	}))
	return r
}

func jsonResponse(w http.ResponseWriter, data interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(b))) // to disable chunked response
	_, err = w.Write(b)
	return err
}

func errorAware(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
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

func translateConnectionHeader() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			connection := r.Header.Get("Connection")
			if connection != "" {
				w.Header().Set("Connection", connection)
			}
			next.ServeHTTP(w, r)
		})
	}
}
