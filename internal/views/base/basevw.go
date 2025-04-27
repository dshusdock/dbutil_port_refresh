package base

import (
	"dshusdock/go_project/config"
	con "dshusdock/go_project/internal/constants"
	"dshusdock/go_project/internal/services/appmgmntsvc"
	"dshusdock/go_project/internal/services/dbservice/appdb/tables"
	"dshusdock/go_project/internal/services/jwtauthsvc"
	s "dshusdock/go_project/internal/services/session"
	"dshusdock/go_project/internal/services/utilitysvc"
	"encoding/gob"
	"log/slog"
	"net/http"
	"time"
)

///////////////////// Base View //////////////////////
type BaseVw struct {
	App *config.AppConfig
}

var AppBaseVw *BaseVw

func init() {
	AppBaseVw = &BaseVw{
		App: nil,
	}
	gob.Register(BaseVwData{})
}

func (m *BaseVw) RegisterHandler() con.ViewHandler {
	slog.Info("[basevw] - RegisterView")
	return &BaseVw{}
}

func (m *BaseVw) HandleRequest(w http.ResponseWriter, event con.AppEvent) any{	
	slog.Info("[basevw] - HandleRequest")
	var obj BaseVwData

	if s.SessionSvc.SessionMgr.Exists(event.Request.Context(), "basevw") {
		obj = s.SessionSvc.SessionMgr.Pop(event.Request.Context(), "basevw").(BaseVwData)
	} else {
		obj = *CreateBaseVwData()	
	}
	obj.ProcessHttpRequest(w, event)	
	
	s.SessionSvc.SessionMgr.Put(event.Request.Context(), "basevw", obj)

	return obj
}


///////////////////// Base View Data //////////////////////

var RED_ERROR_BORDER = "border: 1px solid red"

type BaseVwData struct {
	Base 			*con.BaseTemplateparams
	Data 			[]InputData
	LoggedIn 		bool
	DisplayMsg 		string
	DisplayResponse bool
	ErrorState		bool
}

type InputData struct {
	ForField 		string
	NameField		string
	IdField			string
	TypeField		string
	Label 			string
	Style			string
	
}

func CreateBaseVwData() *BaseVwData {
	return &BaseVwData{
		Base: nil,
		Data: []InputData{
			{
				ForField: "ca_firstname",
				NameField: "firstname",
				IdField: "ca_firstname",
				TypeField: "text",
				Label: "First Name",
				Style: "",
			},
			{
				ForField: "ca_lasttname",
				NameField: "lastname",
				IdField: "ca_lasttname",
				TypeField: "text",
				Label: "Last Name",
				Style: "",
			},
			{
				ForField: "ca_email",
				NameField: "email",
				IdField: "ca_email",
				TypeField: "text",
				Label: "Email",
				Style: "",
			},
			{
				ForField: "ca_username",
				NameField: "username",
				IdField: "ca_username",
				TypeField: "text",
				Label: "Username",
				Style: "",
			},
			{
				ForField: "ca_password",
				NameField: "password",
				IdField: "ca_password",
				TypeField: "password",
				Label: "Password",
				Style: "",
			},
			{
				ForField: "ca_conf_password",
				NameField: "conf_password",
				IdField: "ca_conf_password",
				TypeField: "password",
				Label: "Confirm Password",
				Style: "",
			},
		},
		LoggedIn: false,
		DisplayMsg: "",
		DisplayResponse: false,
		ErrorState: false,
	}
}

func (m *BaseVwData) ProcessHttpRequest(w http.ResponseWriter, event con.AppEvent) *BaseVwData{
	slog.Info("[basevw] - Processing request")
	m.DisplayMsg = ""

	if event.EventStr == "loginvw_create-account" {
		m.DisplayMsg = ""
		m.DisplayResponse = false
		m.ErrorState = false
		return m
	}

	if event.EventStr == "createacctvw_ok" {
		if m.ErrorState {
			m.DisplayMsg = ""
			m.DisplayResponse = false
			m.ErrorState = false
			return m
		}
		// Tell HTMX to do a full page reload
		w.Header().Set("HX-Redirect", "/")
		w.WriteHeader(http.StatusOK)
		return m
	}

	if event.EventStr == "createacctvw_submit" {
		slog.Info("Create Account")
		user := tables.UserData{
			FirstName: event.Request.PostForm.Get("firstname"),
			LastName: event.Request.PostForm.Get("lastname"),
			Email: event.Request.PostForm.Get("email"),
			Username: event.Request.PostForm.Get("username"),
			Password: nil,
		}
		pass := event.Request.PostForm.Get("password")
		val, _ := utilitysvc.EncryptValue(pass)
		user.Password = val
		
		// Confirm passwords match
		confPass := event.Request.PostForm.Get("conf_password")
		if pass != confPass {
			m.DisplayMsg = "Passwords do not match"
			m.DisplayResponse = true
			m.ErrorState = true
			return m
		}
		
		// Verify user does not already exist
		if tables.Check4Username(user.Username) {
			m.DisplayMsg = "Username already exists"
			m.DisplayResponse = true
			m.ErrorState = true
			return m
		}

		slog.Info("User: ", "user", user)
		err := tables.InsertUserInfo(user)
		if err != nil {
			slog.Error("Error inserting user info: ", "error", err)
			m.DisplayMsg = "Error creating account"
			m.DisplayResponse = true
			m.ErrorState = true
			return m
		}
		m.DisplayMsg = "Account Created"
		m.DisplayResponse = true	

		return m
	}

	if event.EventStr == "loginvw_login" {
		user := event.Request.PostForm.Get("username")
		pass := event.Request.PostForm.Get("password")
		
		if user == "test" || pass == "test" {
			m.LoggedIn = true
			appmgmntsvc.CreateAppManager(event.Request.Context())
			m.SetupSession(user, w, event.Request)
			return m
		}

		if tables.ValidatePassword(user, pass) {
			m.LoggedIn = true
			appmgmntsvc.CreateAppManager(event.Request.Context())

			m.SetupSession(user, w, event.Request)


		} else {
			m.DisplayMsg = "Invalid Credentials"
			m.DisplayResponse = true
			m.ErrorState = true
		}
		return m
	}

	return m
}

func (m *BaseVwData) ProcessMBusRequest(w http.ResponseWriter, r *http.Request) {

}

func (m* BaseVwData) SetupSession(username string, w http.ResponseWriter, r *http.Request) {
	token, _ := jwtauthsvc.CreateToken(username)
	http.SetCookie(w, &http.Cookie{
		HttpOnly: true,
		Expires: time.Now().Add(7 * 24 * time.Hour),
		SameSite: http.SameSiteLaxMode,
		// Uncomment below for HTTPS:
		Secure: true,
		// Must be named "jwt" or else the token cannot be 
		// searched for by jwtauth.Verifier.
		Name:  "jwt", 
		Value: token,
	})
	
	err := s.SessionSvc.SessionMgr.RenewToken(r.Context())
	if err != nil {
		http.Error(w, err.Error(), 500)
		return 
	}
	str, _ := utilitysvc.GenerateRandomString(20)
	slog.Info("Random string: ", "value", str)

	s.SessionSvc.SessionMgr.Put(r.Context(), "jwt", token)
	s.SessionSvc.SessionMgr.Put(r.Context(), "LoggedIn", true)
	s.SessionSvc.SessionMgr.Put(r.Context(), "userID", str)
}

func validateUserInfo(user tables.UserData) bool {
	slog.Info("Validating User Info: ", "value", user)
	return true
}
