package apiserver

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type User struct {
	name string
}

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.logger.Info("starting api server")
	s.configureRouter()

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/cards", s.handleCards())
	s.router.HandleFunc("/home", s.handleHome())
	s.router.HandleFunc("/gym", s.handleGym())
	s.router.HandleFunc("/spa", s.handleSpa())
	s.router.HandleFunc("/lpg", s.handleLpg())
	s.router.HandleFunc("/massage", s.handleMassage())
	s.router.HandleFunc("/solarium", s.handleSolarium())
	s.router.HandleFunc("/kedr", s.handleKedr())
	s.router.HandleFunc("/infrcabina", s.handleInfrcabina())
	s.router.HandleFunc("/sauna", s.handleSauna())
}

func (s *APIServer) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		HomeInfo := User{"Bob"}
		tmpl, _ := template.ParseFiles("templates/home2page.html")
		tmpl.Execute(w, HomeInfo)
	}
}

func (s *APIServer) handleCards() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		HomeInfo := User{"Bob"}
		tmpl, _ := template.ParseFiles("templates/cards.html")
		tmpl.Execute(w, HomeInfo)
	}
}

func (s *APIServer) handleGym() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		HomeInfo := User{"Bob"}
		tmpl, _ := template.ParseFiles("templates/gym.html")
		tmpl.Execute(w, HomeInfo)
	}
}

func (s *APIServer) handleSpa() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		HomeInfo := User{"Bob"}
		tmpl, _ := template.ParseFiles("templates/spa.html")
		tmpl.Execute(w, HomeInfo)
	}
}

func (s *APIServer) handleLpg() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		HomeInfo := User{"Bob"}
		tmpl, _ := template.ParseFiles("templates/lpg.html")
		tmpl.Execute(w, HomeInfo)
	}
}

func (s *APIServer) handleMassage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		HomeInfo := User{"Bob"}
		tmpl, _ := template.ParseFiles("templates/massage.html")
		tmpl.Execute(w, HomeInfo)
	}
}

func (s *APIServer) handleSolarium() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		HomeInfo := User{"Bob"}
		tmpl, _ := template.ParseFiles("templates/solarium.html")
		tmpl.Execute(w, HomeInfo)
	}
}

func (s *APIServer) handleKedr() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		HomeInfo := User{"Bob"}
		tmpl, _ := template.ParseFiles("templates/kedr.html")
		tmpl.Execute(w, HomeInfo)
	}
}

func (s *APIServer) handleInfrcabina() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		HomeInfo := User{"Bob"}
		tmpl, _ := template.ParseFiles("templates/infrcab.html")
		tmpl.Execute(w, HomeInfo)
	}
}

func (s *APIServer) handleSauna() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		HomeInfo := User{"Bob"}
		tmpl, _ := template.ParseFiles("templates/sauna.html")
		tmpl.Execute(w, HomeInfo)
	}
}
