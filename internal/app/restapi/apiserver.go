package restapi

import (
	"io"
	"net/http"

	"github.com/danya1off/http-rest-api/internal/app/config"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer ...
type APIServer struct {
	config *config.Config
	router *mux.Router
	logger *logrus.Logger
}

// New ...
func New(config *config.Config) APIServer {
	return APIServer{
		config: config,
		router: mux.NewRouter(),
		logger: logrus.New(),
	}
}

// Start ...
func (s APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s APIServer) configureLogger() error {
	lvl, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(lvl)
	return nil
}

func (s APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.helloHandler())
}

func (s APIServer) helloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
