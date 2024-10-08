package v1

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server struct {
	handler *Handler
	router  *chi.Mux
	address string
}

func NewServer(handler *Handler, router *chi.Mux) *Server {
	return &Server{
		handler: handler,
		router:  router,
	}
}

func (s *Server) Router() {
	s.router.Use(middleware.Logger)

	s.router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Post("/register", s.handler.RegisterHandler)
			r.Post("/login", s.handler.LoginHandler)

			r.Group(func(r chi.Router) {

				r.Use(JWTMiddleware(s.handler.services.Auth))
				r.Post("/create", s.handler.CreateNoteHandler)
				r.Get("/notes", s.handler.GetNotesHandler)
			})
		})
	})

	http.ListenAndServe(":3000", s.router)
}
