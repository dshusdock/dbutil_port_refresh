package dbutilvw

import (
	"dshusdock/go_project/config"
	con "dshusdock/go_project/internal/constants"
	"dshusdock/go_project/internal/services/session"
	"encoding/gob"
	"encoding/json"
	"log/slog"
	"net/http"
)

type DBUtilVw struct {
	App *config.AppConfig
}

var AppDBUtilVw *DBUtilVw

func init() {
	AppDBUtilVw = &DBUtilVw{
		App: nil,
	}
	gob.Register(AppDBUtilVwData{})
	gob.Register(HeaderItem{})
}

func (m *DBUtilVw) RegisterHandler() con.ViewHandler {
	slog.Info("[DBUtilVw] - RegisterView")
	return &DBUtilVw{}
}

func (m *DBUtilVw) HandleRequest(w http.ResponseWriter, event con.AppEvent) any {
	slog.Info("[DBUtilVw] - HandleRequest")
	var obj AppDBUtilVwData

	if session.SessionSvc.SessionMgr.Exists(event.Request.Context(), "DBUtilVw") {
		obj = session.SessionSvc.SessionMgr.Pop(event.Request.Context(), "DBUtilVw").(AppDBUtilVwData)
	} else {
		obj = *CreateDBUtilVwData()	
	}

	obj.ProcessHttpRequest(w, event)	
	session.SessionSvc.SessionMgr.Put(event.Request.Context(), "DBUtilVw", obj)

	return obj
}
 
///////////////////// Layout View Data //////////////////////

type AppDBUtilVwData struct {
	Base 		*con.BaseTemplateparams
	Data 		[]HeaderItem
	QueryVw 	bool
	AuditVw 	bool
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

func CreateDBUtilVwData() *AppDBUtilVwData {
	return &AppDBUtilVwData{
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
		QueryVw: false,
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

func (m *AppDBUtilVwData) ProcessHttpRequest(w http.ResponseWriter, event con.AppEvent) *AppDBUtilVwData{
	slog.Info("[DBUtilVwData] - ProcessHttpRequest")
	slog.Info("event.Src: ", "value", event.Src)
	
	if event.EventId == con.EVENT_STARTUP {
		return m
	}

	if event.EventStr == "queryvw_btn-execute" {
		slog.Info("event.EventStr: ", "value", event.EventStr)
		m.QueryVw = true
	}

	if event.EventStr == "dbsidenavvw_audit" {
		m.AuditVw = true
	}
	return m
}


