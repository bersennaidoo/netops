package server

import (
	"net/http"

	"github.com/bersennaidoo/netops/application/rest/handlers"
	"github.com/bersennaidoo/netops/application/rest/mid"
	"github.com/gorilla/mux"
	"github.com/kataras/golog"
	"github.com/spf13/viper"
)

type HttpServer struct {
	router  *mux.Router
	handler *handlers.Handler
	config  *viper.Viper
	log     *golog.Logger
	mid     *mid.Middleware
}

func New(handler *handlers.Handler, config *viper.Viper, log *golog.Logger, mid *mid.Middleware) *HttpServer {
	return &HttpServer{
		router:  mux.NewRouter(),
		handler: handler,
		config:  config,
		log:     log,
		mid:     mid,
	}
}

func (s *HttpServer) InitRouter() {

	s.router.Use(s.mid.RecoverPanic, s.mid.LogRequest)
	s.router.HandleFunc("/create/config", s.handler.CreateConfig).Methods("POST")
}

func (s *HttpServer) Start() {

	addr := s.config.GetString("http.http_addr")
	srv := &http.Server{
		Addr:    addr,
		Handler: s.router,
	}

	s.log.Debugf("Server Starting on :3000")

	err := srv.ListenAndServe()
	if err != nil {
		s.log.Fatal(err)
	}
}
