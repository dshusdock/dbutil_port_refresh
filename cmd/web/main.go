package main

import (
	"dshusdock/go_project/config"
	"dshusdock/go_project/internal/constants"
	"dshusdock/go_project/internal/handlers"
	"dshusdock/go_project/internal/services/renderview"
	"dshusdock/go_project/internal/services/session"
	"dshusdock/go_project/internal/services/utilitysvc"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
	"github.com/alexedwards/scs/v2"
)

var app config.AppConfig

func main() {
	app = config.AppConfig{}
	_ip, _port, _log_level := utilitysvc.ReadJsonConfigFile()
	
	var programLevel = new(slog.LevelVar) // Info by default
	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))
	
	// Set the level based on the config file
	switch _log_level {
	case "debug":
		programLevel.Set(slog.LevelDebug)
	case "info":
		programLevel.Set(slog.LevelInfo)
	case "error":
		programLevel.Set(slog.LevelError)
	default:
		programLevel.Set(slog.LevelInfo)
	}

	app.InProduction = false
	app.ViewCache = make(map[string]constants.ViewInterface)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render := renderview.NewRenderViewSvc(&app)
	renderview.MapRenderViewSvc(render)
	renderview.InitRouteHandlers()
	
	// Setup Session Manager
	app.SessionManager = scs.New()
	app.SessionManager.Lifetime = 24 * time.Hour
	app.SessionManager.IdleTimeout = 480 * time.Minute
	app.SessionManager.Cookie.Name = "session_id"
	app.SessionManager.Cookie.Domain = _ip
	app.SessionManager.Cookie.HttpOnly = true
	app.SessionManager.Cookie.Persist = true
	app.SessionManager.Cookie.Secure = true
	// app.SessionManager.Cookie.Path = "/exops/"
	// app.SessionManager.Cookie.SameSite = http.SameSiteStrictMode
	session.SessionSvc.RegisterSessionManager(app.SessionManager)
	slog.Debug("Cookie Domain: ", "Domain", app.SessionManager.Cookie.Domain)

	buildDate := os.Getenv("BUILD_DATE")
	slog.Info("Build", "Date", buildDate)
	slog.Info("Starting application -", "Port", _port)
	srv := &http.Server{
		Addr:    ":" + _port,
		Handler: app.SessionManager.LoadAndSave(routes(&app)),
	}

	err := srv.ListenAndServeTLS("dev_cert.crt", "private.key")
	if err != nil {
		log.Fatal(err)
	}
}


