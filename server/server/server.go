package server

import (
	"database/sql"
	"innovateafitc/config"
	"innovateafitc/routes"

	// this package is the database driver.
	_ "github.com/lib/pq"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Server struct {
	routes *routes.Routes
	cfg    *config.Config
}

// New creates something new.
func New(cfg *config.Config) (*Server, error) {
	boil.DebugMode = true

	db, err := sql.Open("postgres", cfg.ConnectionString)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	r, err := routes.New(cfg, db) //, db

	if err != nil {
		return nil, err
	}

	return &Server{
		routes: r,
		cfg:    cfg,
	}, nil
}

// Start runs the application and connects to the database
func (s *Server) Start() error {
	return s.routes.Router.Run(s.cfg.ListenAddr)
}
