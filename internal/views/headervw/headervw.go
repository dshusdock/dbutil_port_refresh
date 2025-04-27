package headervw

import (
	"dshusdock/go_project/config"
	con "dshusdock/go_project/internal/constants"
	"dshusdock/go_project/internal/services/session"
	"encoding/gob"
	"encoding/json"
	"log/slog"
	"net/http"
)

type HeaderVw struct {
	App *config.AppConfig
}

var AppHeaderVw *HeaderVw

func init() {
	AppHeaderVw = &HeaderVw{
		App: nil,
	}
	gob.Register(AppHeaderVwData{})
	gob.Register(HeaderItem{})
}

func (m *HeaderVw) RegisterHandler() con.ViewHandler {
	slog.Info("[HeaderVw] - RegisterView")
	return &HeaderVw{}
}

func (m *HeaderVw) HandleRequest(w http.ResponseWriter, event con.AppEvent) any {
	slog.Info("[HeaderVw] - HandleRequest")
	var obj AppHeaderVwData

	if session.SessionSvc.SessionMgr.Exists(event.Request.Context(), "HeaderVw") {
		obj = session.SessionSvc.SessionMgr.Pop(event.Request.Context(), "HeaderVw").(AppHeaderVwData)
	} else {
		obj = *CreateHeaderVwData()	
	}

	obj.ProcessHttpRequest(w, event)	
	session.SessionSvc.SessionMgr.Put(event.Request.Context(), "HeaderVw", obj)

	return obj
}
 
///////////////////// Layout View Data //////////////////////

type AppHeaderVwData struct {
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

func CreateHeaderVwData() *AppHeaderVwData {
	return &AppHeaderVwData{
		Base: nil,
		Data: []HeaderItem{
			{"BTN_Home", "Home", "btn_init"},		
			{"BTN_Query", "Query", ""},		
			{"BTN_Audit", "Audit", ""},		
			{"BTN_LoadDB", "Load DB", ""},		
			{"BTN_Queries", "DB Queries", ""},	
			{"BTN_RemTasks", "Rem Tasks", ""},		
			{"BTN_TableMapper", "Table Mapper", ""},		
			{"BTN_ColumnFinder", "Column Finder", ""},		
			{"BTN_Logout", "Logout", ""},		
			
		},
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

func (m *AppHeaderVwData) ProcessHttpRequest(w http.ResponseWriter, event con.AppEvent) *AppHeaderVwData{
	slog.Info("[HeaderVwData] - ProcessHttpRequest")
	slog.Info("event.Src: ", "value", event.Src)
	if event.EventStr == "headervw_btn_Logout" {
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


