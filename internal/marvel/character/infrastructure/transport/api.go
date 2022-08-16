package transport

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jamsxd/marvel-api/internal/marvel/character/application"
)

func NewServer(endpoint application.CharacterEndpoint) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Heartbeat("/health"))
	r.Post("/character", makePostCharacter(endpoint))
	return r
}

func makePostCharacter(endpoint application.CharacterEndpoint) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		req := &application.CreateCharacterRequest{}

		if err := json.NewDecoder(r.Body).Decode(&req.Character); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res, err := endpoint.CreateCharacter(r.Context(), *req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}
