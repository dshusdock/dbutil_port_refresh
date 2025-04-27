package config

import (
	"dshusdock/go_project/internal/constants"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	UseCache      		bool
	ViewCache     		map[string]constants.ViewInterface
	InProduction  		bool
	SessionManager 	 	*scs.SessionManager
	
}


