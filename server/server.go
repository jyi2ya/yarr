package server

import (
	"log"
	"net/http"
	"sync"

	"github.com/jgkawell/yarr/storage"
	"github.com/jgkawell/yarr/worker"
)

type Server struct {
	Addr        string
	db          *storage.Storage
	worker      *worker.Worker
	cache       map[string]interface{}
	cache_mutex *sync.Mutex

	BasePath string

	// auth
	Username string
	Password string
}

func NewServer(db *storage.Storage, addr string) *Server {
	return &Server{
		db:          db,
		Addr:        addr,
		worker:      worker.NewWorker(db),
		cache:       make(map[string]interface{}),
		cache_mutex: &sync.Mutex{},
	}
}

func (h *Server) GetAddr() string {
	proto := "http"
	return proto + "://" + h.Addr + h.BasePath
}

func (s *Server) Start() {
	refreshRate := s.db.GetSettingsValueInt64("refresh_rate")
	s.worker.FindFavicons()
	s.worker.StartFeedCleaner()
	s.worker.SetRefreshRate(refreshRate)
	if refreshRate > 0 {
		s.worker.RefreshFeeds()
	}

	httpserver := &http.Server{Addr: s.Addr, Handler: s.handler()}
	err := httpserver.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
