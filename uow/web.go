package uow

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/google/uuid"
)

type Server struct {
	http.Server
	service *Service
}

func NewServer(service *Service) Server {
	mux := http.NewServeMux()
	return Server{
		Server: http.Server{
			Addr:    ":8080",
			Handler: mux,
		},
		service: service,
	}
}

func (s *Server) FindAll(w http.ResponseWriter, r *http.Request) {
	companies, err := s.service.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(companies); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) FindByID(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile(`^/companies/(\s+)$`)
	matches := re.FindAllStringSubmatch(r.URL.Path, -1)
	if matches == nil {
		http.NotFound(w, r)
		return
	}
	id, err := uuid.Parse(matches[1][0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	company, err := s.service.FIndByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := json.NewEncoder(w).Encode(company); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
