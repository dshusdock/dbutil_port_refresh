package appsidenavvw

import (
	"dshusdock/go_project/config"
	con "dshusdock/go_project/internal/constants"
	"dshusdock/go_project/internal/services/session"
	"encoding/gob"
	"encoding/json"
	"log/slog"
	"net/http"
)

type AppSideNavVw struct {
	App *config.AppConfig
}

var AppAppSideNavVw *AppSideNavVw

func init() {
	AppAppSideNavVw = &AppSideNavVw{
		App: nil,
	}
	gob.Register(AppAppSideNavVwData{})
	gob.Register(HeaderItem{})
}

func (m *AppSideNavVw) RegisterHandler() con.ViewHandler {
	slog.Info("[AppSideNavVw] - RegisterView")
	return &AppSideNavVw{}
}

func (m *AppSideNavVw) HandleRequest(w http.ResponseWriter, event con.AppEvent) any {
	slog.Info("[AppSideNavVw] - HandleRequest")
	var obj AppAppSideNavVwData

	if session.SessionSvc.SessionMgr.Exists(event.Request.Context(), "AppSideNavVw") {
		obj = session.SessionSvc.SessionMgr.Pop(event.Request.Context(), "AppSideNavVw").(AppAppSideNavVwData)
	} else {
		obj = *CreateAppSideNavVwData()	
	}

	obj.ProcessHttpRequest(w, event)	
	session.SessionSvc.SessionMgr.Put(event.Request.Context(), "AppSideNavVw", obj)

	return obj
}
 
///////////////////// Layout View Data //////////////////////

type AppAppSideNavVwData struct {
	Base *con.BaseTemplateparams
	Data []HeaderItem
}

type HeaderItem struct {
	ID 				string
	Lbl     		string
	Class			string
}

type DropDownData struct {
	Class string
	Style string
}

func CreateAppSideNavVwData() *AppAppSideNavVwData {
	return &AppAppSideNavVwData{
		Base: nil,
		Data: nil,
	}
}

type SomeStruct struct{
	Result string	`json:"result"`
}

func SomeHandler(w http.ResponseWriter, r *http.Request) {
    data := SomeStruct{ Result: "success" }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(data)
}

func (m *AppAppSideNavVwData) ProcessHttpRequest(w http.ResponseWriter, event con.AppEvent) *AppAppSideNavVwData{
	slog.Info("[AppSideNavVwData] - ProcessHttpRequest")
	
	if event.EventStr == "AppSideNavVw_btn_Logout" {
		// Tell HTMX to do a full page reload
		w.Header().Set("HX-Redirect", "/")
		w.WriteHeader(http.StatusOK)
		return m
	}
	if event.EventId == con.EVENT_STARTUP {
		return m
	}
	return m
}


