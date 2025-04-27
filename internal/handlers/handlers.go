package handlers

import (
	"dshusdock/go_project/config"
	"dshusdock/go_project/internal/constants"
	"dshusdock/go_project/internal/services/messagebus"
	"dshusdock/go_project/internal/services/renderview"
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	MBS messagebus.MessageBusSvc
}

func initRouteHandlers() {
	// Register the views
	// Repo.App.ViewCache["basevw"] = base.AppBaseVw.RegisterView(Repo.App)
	// Repo.App.ViewCache["lyoutvw"] = layoutvw.AppLayoutVw.RegisterView(Repo.App)
	// Repo.App.ViewCache["sidenavvw"] = sidenavvw.AppSideNavVw.RegisterView(Repo.App)

}

// http.ResponseWriter, r *http.Request NewRepo creates a new repository
func NewRepo(app *config.AppConfig) *Repository {
	return &Repository{
		App: app,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
	initRouteHandlers()
}

func (m *Repository) Base(w http.ResponseWriter, r *http.Request) {
	slog.Info("Base Handler PATH-", "value", r.URL.Path)

	// ////////////////TEMPORARY//////////////////////
	// token, _ := jwtauthsvc.CreateToken("dbutil")
	// http.SetCookie(w, &http.Cookie{
	// 	HttpOnly: true,
	// 	Expires: time.Now().Add(7 * 24 * time.Hour),
	// 	SameSite: http.SameSiteLaxMode,
	// 	// Uncomment below for HTTPS:
	// 	Secure: true,
	// 	// Must be named "jwt" or else the token cannot be 
	// 	// searched for by jwtauth.Verifier.
	// 	Name:  "jwt", 
	// 	Value: token,
	// })
	// err := session.SessionSvc.SessionMgr.RenewToken(r.Context())
	// if err != nil {
	// 	slog.Info("Error renewing token: ", err)
	// }
	// ///////////////////////////////////////////////

	renderview.RenderViewSvc.ProcessInit(w, r)

}

func (m *Repository) StateChange(w http.ResponseWriter, r *http.Request) {
	slog.Info("State Change Handler PATH-", "path", r.URL.Path)
	renderview.RenderViewSvc.ProcessStateChange(w, r)
}

func (m *Repository) HandleEvents(w http.ResponseWriter, r *http.Request) {

	event_type := constants.EVENT_CLICK
	if r.URL.Path == "/event/element/change" {
		event_type = constants.EVENT_CHANGE
	}

	m.checkSessionHandler(w, r)
	
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	data := r.PostForm
	data.Add("event", event_type)
	v_id := data.Get("view")

	if v_id == "" {
		_ = fmt.Errorf("no handler for route")
		return
	}
	slog.Info("[Handlers] View ID - ", "view", v_id)

	renderview.RenderViewSvc.ProcessEvent(w, r)
}

// func (m *Repository) HandleChangeEvents(w http.ResponseWriter, r *http.Request) {
// 	val := true//m.App.SessionManager.Get(r.Context(), "LoggedIn")
// 	slog.Info("[Handlers] Logged In -  ", val)

// 	if val != true {
// 		http.Redirect(w, r, "/", http.StatusSeeOther)
// 		return
// 	} 
	
// 	err := r.ParseForm()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	data := r.PostForm
// 	data.Add("event", constants.EVENT_CHANGE)
// 	v_id := data.Get("view_id")

// 	if v_id == "" {
// 		_ = fmt.Errorf("no handler for route")
// 		return
// 	}
// 	slog.Info("[Handlers] View ID - ", v_id)

// 	//renderview.RenderViewSvc.ProcessRequest(w, r, v_id)
// 	renderview.RenderViewSvc.ProcessEvent(w, r)
// }

func (m *Repository) FileUpload(w http.ResponseWriter, r *http.Request) {
	//  Check if user is logged in

/*
	val := m.App.SessionManager.Get(r.Context(), "LoggedIn")
	if val != true {
		// m.App.LoggedIn = false
		slog.Info("User not logged in")
		return
	} else {
		// m.App.LoggedIn = true
	}
*/
	// Parse our multipart form, 10 << 20 specifies a maximum of 10MB of memory
	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("myFile")
	if err != nil {
		slog.Info("Error Retrieving the File")
		slog.Info(err.Error())
		return
	}
	defer file.Close()

	//p := file_upload.ProcessLabInfo(file)
	//tables.PopulateMasterLayoutFromFile(file)
}



func (m *Repository) checkSessionHandler(w http.ResponseWriter, r *http.Request) {
	// if m.App.SessionManager.Exists(r.Context(), "session_id") {
    //     fmt.Println(w, "Session is active")
    // } else {
    //     fmt.Println(w, "Session has expired or does not exist")
    // }

	keys := m.App.SessionManager.Keys(r.Context())
	for _, key := range keys {
		slog.Debug("Key: ", "", key)
	}

}
