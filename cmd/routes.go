package cmd

import (
	"golang-api/pkg/utils"
)

func (s *Server) InitializeRoutes() {
	s.Router.HandleFunc("/api/users", utils.HeaderJSON(CreateUser)).Methods("POST")
	s.Router.HandleFunc("/api/users", utils.HeaderJSON(GetUsers)).Methods("GET", "OPTIONS")
	s.Router.HandleFunc("/api/users/{id}", utils.HeaderJSON(GetUser)).Methods("GET", "OPTIONS")
	s.Router.HandleFunc("/api/users/{id}", utils.HeaderJSON(UpdateUser)).Methods("PATCH")
	s.Router.HandleFunc("/api/users/{id}", utils.HeaderJSON(DeleteUser)).Methods("DELETE")

	s.Router.HandleFunc("/api/users/{id}/friends", utils.HeaderJSON(GetFriends)).Methods("GET")
	s.Router.HandleFunc("/api/users/{id}/friends", utils.HeaderJSON(CreateFriend)).Methods("POST")
	s.Router.HandleFunc("/api/users/{uid}/friends/{fid}", utils.HeaderJSON(DeleteFriend)).Methods("DELETE")
}
