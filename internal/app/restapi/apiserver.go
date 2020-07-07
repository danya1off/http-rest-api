package restapi

import (
	"io"
	"net/http"

	"github.com/danya1off/http-rest-api/internal/app/config"
	"github.com/danya1off/http-rest-api/internal/app/model/user"
	"github.com/danya1off/http-rest-api/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer ...
type APIServer struct {
	config  *config.Config
	router  *mux.Router
	logger  *logrus.Logger
	db      store.DB
	userDAO user.DAO
}

// New ...
func New(config *config.Config) APIServer {
	return APIServer{
		config: config,
		router: mux.NewRouter(),
		logger: config.Logger,
	}
}

// Start ...
func (s APIServer) Start() error {
	if err := s.configureDB(); err != nil {
		return err
	}
	s.configureUserRepository()

	s.configureRouter()
	s.logger.Info("Starting the server at port: " + s.config.BindAddr)
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.helloHandler)
}

func (s APIServer) configureDB() error {
	db := store.New(s.config.DBConfig)
	if err := db.Open(); err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s APIServer) configureUserRepository() {
	dao := user.Repository{
		DB:     s.db,
		Logger: s.logger,
	}
	s.userDAO = dao
}

func (s APIServer) helloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := s.userDAO.FindByEmail("")
	if err != nil {
		s.logger.WithError(err).Fatal("Error!")
	}
	io.WriteString(w, "Hello")
}
