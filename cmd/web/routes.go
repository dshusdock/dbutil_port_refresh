package main

import (
	"context"
	"dshusdock/go_project/config"
	"dshusdock/go_project/internal/handlers"
	"dshusdock/go_project/internal/services/jwtauthsvc"
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth/v5"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Heartbeat("/health"))
	//mux.Use(middleware.Logger)
	
	// Protecting the routes
	mux.Group(func(r chi.Router) {
		// r.Use(MyMiddleware)		
		r.Use(jwtauth.Verifier(jwtauthsvc.GetToken()))
		r.Use(jwtauth.Authenticator(jwtauthsvc.GetToken()))	
			
		r.Post("/event/element/click", handlers.Repo. HandleEvents)
		r.Post("/event/element/change", handlers.Repo.HandleEvents)
		r.Post("/event/element/search", handlers.Repo.HandleEvents)
		r.Post("/upload", handlers.Repo.FileUpload)
		
	})

	// Unprotected routes
	mux.Get("/", handlers.Repo.Base)
	mux.Post("/app/state/change", handlers.Repo.StateChange)

	fileServer := http.FileServer(http.Dir("./ui/html/"))
	mux.Handle("/html/*", http.StripPrefix("/html", fileServer))

	return mux
}

/////////////////////// MiddleWare ////////////////////////////////////////

// HTTP middleware (example) setting a value on the request context
func MyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	  // create new context from `r` request context, and assign key `"user"`
	  // to value of `"123"`
	  ctx := context.WithValue(r.Context(), "user", "123")
  
	  // call the next handler in the chain, passing the response writer and
	  // the updated request object with the new context value.
	  //
	  // note: context.Context values are nested, so any previously set
	  // values will be accessible as well, and the new `"user"` key
	  // will be accessible from this point forward.
	  next.ServeHTTP(w, r.WithContext(ctx))
	})
  }

  func MyHandler(w http.ResponseWriter, r *http.Request) {
    // here we read from the request context and fetch out `"user"` key set in
    // the MyMiddleware example above.
    user := r.Context().Value("user").(string)

    // respond to the client
    w.Write([]byte(fmt.Sprintf("hi %s", user)))
}

  
