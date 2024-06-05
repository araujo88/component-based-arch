package server

import (
	"component-based-arch/internal/config"
	"component-based-arch/internal/logger"
	"component-based-arch/internal/userdao"
	"encoding/json"
	"net/http"
	"strconv"
)

// Server manages HTTP requests and utilizes other components.
type Server struct {
	config  *config.Config
	userDAO *userdao.UserDAO
	logger  *logger.Logger
}

// NewServer creates a new HTTP server with dependencies.
func NewServer(cfg *config.Config, userDao *userdao.UserDAO, logger *logger.Logger) *Server {
	return &Server{
		config:  cfg,
		userDAO: userDao,
		logger:  logger,
	}
}

// Serve starts the HTTP server.
func (s *Server) Serve() error {
	http.HandleFunc("POST /users", s.createUser)
	http.HandleFunc("GET /users/{id}", s.getUser)
	http.HandleFunc("PUT /users", s.updateUser)
	http.HandleFunc("DELETE /users/{id}", s.deleteUser)
	s.logger.Info("Server listening on port", s.config.ServerPort)
	return http.ListenAndServe(s.config.ServerPort, nil)
}

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	var user userdao.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		s.logger.Error("Error decoding user:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	s.userDAO.CreateUser(&user)
	s.logger.Info("Created user:", user)
	json.NewEncoder(w).Encode(user)
}

func (s *Server) getUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		s.logger.Error("Invalid user ID")
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	user, exists := s.userDAO.GetUser(id)
	if !exists {
		http.NotFound(w, r)
		return
	}
	s.logger.Info("Fetched user:", user)
	json.NewEncoder(w).Encode(user)
}

func (s *Server) updateUser(w http.ResponseWriter, r *http.Request) {
	var user userdao.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		s.logger.Error("Error decoding user:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if err := s.userDAO.UpdateUser(&user); err != nil {
		s.logger.Error("Error updating user:", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	s.logger.Info("Updated user:", user)
	json.NewEncoder(w).Encode(user)
}

func (s *Server) deleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		s.logger.Error("Invalid user ID")
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	_, exists := s.userDAO.GetUser(id)
	if !exists {
		http.NotFound(w, r)
		return
	}
	s.userDAO.DeleteUser(id)
	s.logger.Info("Deleted user with ID:", id)
	w.WriteHeader(http.StatusNoContent)
}
