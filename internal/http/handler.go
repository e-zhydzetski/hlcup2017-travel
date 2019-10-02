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
		_ = json.NewEncoder(w).Encode(viewDTO)
		return nil
	}))
	r.Get("/locations/{id}", func(w http.ResponseWriter, r *http.Request) {

	})
	r.Get("/visits/{id}", func(w http.ResponseWriter, r *http.Request) {

	})
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
